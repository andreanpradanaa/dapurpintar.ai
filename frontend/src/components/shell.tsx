"use client";
import { useEffect, useState } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useAuth } from "@/lib/auth";
import { api } from "@/lib/api";

const links = [
  { href: "/today", label: "Today" },
  { href: "/pantry", label: "Pantry" },
  { href: "/planner", label: "Planner" },
  { href: "/shopping", label: "Shopping" },
  { href: "/discover", label: "Discover" },
];

export default function Shell({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const { account, logout } = useAuth();
  const [attentionCount, setAttentionCount] = useState(0);

  useEffect(() => {
    if (!account) return;
    api.pantrySummary().then(r => {
      setAttentionCount((r.data.expiring_soon_count || 0) + (r.data.running_low_count || 0));
    }).catch(() => {});
  }, [account]);

  if (!account) return <>{children}</>;

  return (
    <div className="min-h-screen pb-16 md:pb-0">
      <header className="sticky top-0 z-30 bg-santan-050/80 backdrop-blur-md border-b border-bambu-200/50">
        <div className="max-w-5xl mx-auto px-4 h-14 flex items-center justify-between">
          <Link href="/today" className="font-display text-xl font-bold text-kuali-950 hover:text-rempah-500 transition-colors">
            DapurPintar
          </Link>
          <div className="flex items-center gap-4 text-sm">
            <Link href="/profile" className="text-kuali-700/60 hover:text-kuali-950 transition-colors">
              {account?.email}
            </Link>
            <button onClick={logout} className="text-bambu-300 hover:text-rempah-500 transition-colors text-xs font-medium">
              Keluar
            </button>
          </div>
        </div>
      </header>

      <main className="max-w-5xl mx-auto px-4 py-8">{children}</main>

      <nav className="fixed bottom-0 left-0 right-0 bg-santan-050/95 backdrop-blur-md border-t border-bambu-200/50 z-30 md:static md:border-t-0 md:border-b md:mt-0">
        <div className="flex justify-around max-w-5xl mx-auto">
          {links.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              className={`relative py-3 px-3 text-xs font-medium tracking-wide uppercase transition-colors ${
                pathname.startsWith(l.href)
                  ? "text-rempah-500"
                  : "text-kuali-700/40 hover:text-kuali-700"
              }`}
            >
              {l.label}
              {l.href === "/today" && attentionCount > 0 && (
                <span className="absolute -top-0.5 -right-1 bg-rempah-500 text-white text-[9px] font-bold rounded-full w-4 h-4 flex items-center justify-center leading-none">
                  {attentionCount > 9 ? "9+" : attentionCount}
                </span>
              )}
            </Link>
          ))}
        </div>
      </nav>
    </div>
  );
}
