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
    <div className="min-h-screen pb-14 md:pb-0">
      <header className="border-b border-steel-200 bg-white-000 px-4 py-3 flex items-center justify-between">
        <Link href="/today" className="font-semibold text-lg text-ink-900">
          DapurPintar
        </Link>
        <div className="flex items-center gap-3 text-sm">
          <Link href="/profile" className="text-ink-700 hover:text-ink-900">
            {account?.email}
          </Link>
          <button onClick={logout} className="text-steel-400 hover:text-ink-700">
            Keluar
          </button>
        </div>
      </header>
      <main className="max-w-4xl mx-auto px-4 py-6">{children}</main>
      <nav className="fixed bottom-0 left-0 right-0 border-t border-steel-200 bg-white-000 md:static md:border-t-0 md:border-b md:mt-0 z-40">
        <div className="flex justify-around max-w-4xl mx-auto">
          {links.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              className={`relative py-3 px-4 text-sm font-medium ${
                pathname.startsWith(l.href)
                  ? "text-action-primary border-b-2 border-action-primary -mb-[2px]"
                  : "text-ink-700 hover:text-ink-900"
              }`}
            >
              {l.label}
              {l.href === "/today" && attentionCount > 0 && (
                <span className="absolute -top-0.5 -right-0.5 bg-action-primary text-white-000 text-[10px] font-bold rounded-full w-4 h-4 flex items-center justify-center leading-none">
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
