import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

const PUBLIC_PATHS = ["/login", "/api"];

function isPublic(pathname: string) {
  if (pathname === "/") return true;
  return PUBLIC_PATHS.some(p => pathname.startsWith(p));
}
const ASSET_PATTERNS = /\.(ico|png|svg|jpg|css|js|map)$/;

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;
  if (ASSET_PATTERNS.test(pathname)) return NextResponse.next();

  const isPub = isPublic(pathname);
  const hasSession = request.cookies.has("dp_session");

  if (!isPub && !hasSession) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  if (pathname.startsWith("/login") && hasSession) {
    return NextResponse.redirect(new URL("/today", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/((?!_next/static|_next/image|favicon.ico).*)"],
};
