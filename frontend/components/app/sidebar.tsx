"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { motion } from "framer-motion";
import {
  LayoutDashboard,
  Sparkles,
  History,
  Heart,
  Settings,
  User,
  LogOut,
  X,
} from "lucide-react";
import { Logo } from "@/components/ui/logo";
import { Avatar } from "@/components/ui/avatar";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { cn } from "@/lib/utils";
import { APP_NAV } from "@/lib/site";
import { Badge } from "@/components/ui/badge";

const ICON_MAP = {
  LayoutDashboard,
  Sparkles,
  History,
  Heart,
  Settings,
  User,
} as const;

export function Sidebar({
  mobileOpen,
  onMobileClose,
}: {
  mobileOpen: boolean;
  onMobileClose: () => void;
}) {
  const pathname = usePathname();
  const { t } = useLanguage();
  const user = useStore((s) => s.user);
  const signOut = useStore((s) => s.signOut);
  const historyCount = useStore((s) => s.history.length);
  const favCount = useStore((s) => s.favorites.length);

  const NavContent = () => (
    <div className="flex flex-col h-full">
      <div className="h-16 flex items-center justify-between px-5 border-b border-border">
        <Link href="/dashboard" onClick={onMobileClose} className="flex items-center">
          <Logo size="md" />
        </Link>
        <button
          type="button"
          onClick={onMobileClose}
          className="lg:hidden flex h-8 w-8 items-center justify-center rounded-md text-text-muted hover:text-text-primary hover:bg-bg-section"
          aria-label="Close menu"
        >
          <X className="h-4 w-4" />
        </button>
      </div>

      <nav className="flex-1 px-3 py-5 overflow-y-auto">
        <ul className="space-y-0.5">
          {APP_NAV.map((item) => {
            const Icon = ICON_MAP[item.icon as keyof typeof ICON_MAP];
            const active = pathname === item.href || pathname.startsWith(item.href + "/");
            const labelKey = item.labelKey.split(".").pop() as keyof typeof t.app.sidebar;
            const count =
              item.icon === "History"
                ? historyCount
                : item.icon === "Heart"
                ? favCount
                : null;
            return (
              <li key={item.href}>
                <Link
                  href={item.href}
                  onClick={onMobileClose}
                  className={cn(
                    "group relative flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors duration-200",
                    "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent",
                    active
                      ? "text-text-primary"
                      : "text-text-muted hover:text-text-primary hover:bg-bg-section"
                  )}
                >
                  {active && (
                    <motion.span
                      layoutId="sidebar-active"
                      className="absolute inset-0 rounded-lg bg-bg-section border border-border"
                      transition={{ type: "spring", stiffness: 350, damping: 30 }}
                    />
                  )}
                  <span className="relative z-10 flex items-center gap-3 flex-1">
                    <Icon
                      className={cn(
                        "h-4 w-4",
                        active ? "text-accent" : "text-text-muted group-hover:text-text-primary"
                      )}
                      strokeWidth={1.75}
                    />
                    <span>{t.app.sidebar[labelKey]}</span>
                  </span>
                  {count !== null && count > 0 && (
                    <span className="relative z-10 text-[10px] font-medium text-text-muted tabular-nums">
                      {count}
                    </span>
                  )}
                </Link>
              </li>
            );
          })}
        </ul>

        {user?.plan === "pro" && (
          <div className="mt-8 rounded-xl border border-accent-soft-strong bg-accent-soft p-4">
            <div className="flex items-center gap-2 mb-2">
              <Sparkles className="h-3.5 w-3.5 text-accent-active" />
              <span className="text-[10px] font-semibold text-accent-active uppercase tracking-[0.12em]">
                Pro
              </span>
            </div>
            <p className="text-xs text-text-secondary leading-[1.6]">
              You have unlimited recipes and detailed nutrition breakdowns.
            </p>
            <Link
              href="/settings"
              onClick={onMobileClose}
              className="mt-3 inline-flex items-center gap-1 text-xs text-accent-active hover:text-text-primary font-medium"
            >
              Manage plan →
            </Link>
          </div>
        )}
      </nav>

      <div className="border-t border-border p-3">
        <Link
          href="/profile"
          onClick={onMobileClose}
          className="flex items-center gap-3 rounded-lg p-2 hover:bg-bg-section transition-colors"
        >
          <Avatar name={user?.name ?? "User"} size="sm" />
          <div className="flex-1 min-w-0">
            <div className="text-sm font-medium text-text-primary truncate">
              {user?.name}
            </div>
            <div className="text-xs text-text-muted truncate flex items-center gap-1">
              <Badge variant="accent" className="text-[9px] py-0 capitalize">
                {user?.plan}
              </Badge>
            </div>
          </div>
        </Link>
        <button
          type="button"
          onClick={signOut}
          className="mt-2 w-full flex items-center gap-2 rounded-lg px-3 py-2 text-xs text-text-muted hover:text-text-primary hover:bg-bg-section transition-colors"
        >
          <LogOut className="h-3.5 w-3.5" />
          Sign out
        </button>
      </div>
    </div>
  );

  return (
    <>
      {/* Desktop sidebar */}
      <aside className="hidden lg:flex w-64 shrink-0 border-r border-border bg-bg-section/40 sticky top-0 h-screen">
        <NavContent />
      </aside>

      {/* Mobile drawer */}
      {mobileOpen && (
        <div className="lg:hidden fixed inset-0 z-50 flex">
          <div
            className="absolute inset-0 bg-bg-overlay"
            onClick={onMobileClose}
          />
          <motion.aside
            initial={{ x: -300 }}
            animate={{ x: 0 }}
            exit={{ x: -300 }}
            transition={{ type: "spring", stiffness: 300, damping: 30 }}
            className="relative w-72 max-w-[85vw] bg-bg-base border-r border-border h-full"
          >
            <NavContent />
          </motion.aside>
        </div>
      )}
    </>
  );
}
