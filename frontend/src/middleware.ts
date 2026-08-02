import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

const PUBLIC_PATHS = ["/login", "/api"];
const ASSET_PATTERNS = /\.(ico|png|svg|jpg|css|js|map)$/;

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;
  if (ASSET_PATTERNS.test(pathname)) return NextResponse.next();

  const isPublic = PUBLIC_PATHS.some(p => pathname.startsWith(p));
  const hasSession = request.cookies.has("dp_session");

  if (!isPublic && !hasSession) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  if (isPublic && hasSession) {
    return NextResponse.redirect(new URL("/today", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/((?!_next/static|_next/image|favicon.ico).*)"],
};
