"use client";

import Link from "next/link";
import { Menu, Search, Bell, Sparkles, Command } from "lucide-react";
import { Button } from "@/components/ui/button";
import { LanguageToggle } from "@/components/ui/language-toggle";
import { useLanguage } from "@/components/providers/language-provider";
import { usePathname } from "next/navigation";
import { Avatar } from "@/components/ui/avatar";
import { useStore } from "@/lib/store";

export function Topbar({ onMenuClick }: { onMenuClick: () => void }) {
  const { t } = useLanguage();
  const user = useStore((s) => s.user);
  const pathname = usePathname();

  return (
    <header className="sticky top-0 z-30 bg-bg-base/85 backdrop-blur-md border-b border-border">
      <div className="flex items-center justify-between gap-3 h-16 px-4 sm:px-6">
        <div className="flex items-center gap-2 flex-1 min-w-0">
          <button
            type="button"
            onClick={onMenuClick}
            className="lg:hidden flex h-9 w-9 items-center justify-center rounded-md text-text-muted hover:text-text-primary hover:bg-bg-section focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
            aria-label="Open menu"
          >
            <Menu className="h-5 w-5" />
          </button>

          {/* Page title */}
          <div className="hidden sm:block">
            <h1 className="text-base font-serif font-medium capitalize text-text-primary">
              {pathname.split("/").pop()?.replace("-", " ") || "dashboard"}
            </h1>
          </div>
        </div>

        <div className="flex items-center gap-2">
          <button
            type="button"
            className="hidden md:inline-flex items-center gap-2 h-9 px-3 rounded-md border border-border bg-bg-card text-sm text-text-muted hover:border-border-strong hover:text-text-primary transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
          >
            <Search className="h-3.5 w-3.5" />
            <span>Search…</span>
            <kbd className="ml-2 hidden lg:inline-flex items-center gap-0.5 rounded border border-border bg-bg-base px-1.5 text-[10px] font-mono">
              <Command className="h-2.5 w-2.5" />K
            </kbd>
          </button>

          <LanguageToggle className="hidden sm:inline-flex" />

          <button
            type="button"
            className="relative flex h-9 w-9 items-center justify-center rounded-md text-text-muted hover:text-text-primary hover:bg-bg-section focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
            aria-label="Notifications"
          >
            <Bell className="h-4 w-4" />
            <span className="absolute top-2 right-2 h-1.5 w-1.5 rounded-full bg-accent" />
          </button>

          <Link href="/generate" className="hidden sm:inline-flex">
            <Button size="sm" className="group">
              <Sparkles className="h-3.5 w-3.5" />
              {t.app.sidebar.generate}
            </Button>
          </Link>

          <Link href="/profile" className="lg:hidden">
            <Avatar name={user?.name ?? "User"} size="sm" />
          </Link>
        </div>
      </div>
    </header>
  );
}
