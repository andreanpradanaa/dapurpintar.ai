"use client";

import Link from "next/link";
import { useEffect, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { Menu, X } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Logo } from "@/components/ui/logo";
import { LanguageToggle } from "@/components/ui/language-toggle";
import { useLanguage } from "@/components/providers/language-provider";
import { cn } from "@/lib/utils";
import { NAV_LINKS } from "@/lib/site";

export function MarketingNav() {
  const { t } = useLanguage();
  const [scrolled, setScrolled] = useState(false);
  const [open, setOpen] = useState(false);

  useEffect(() => {
    const onScroll = () => setScrolled(window.scrollY > 12);
    onScroll();
    window.addEventListener("scroll", onScroll, { passive: true });
    return () => window.removeEventListener("scroll", onScroll);
  }, []);

  return (
    <>
      <motion.header
        initial={{ y: -8, opacity: 0 }}
        animate={{ y: 0, opacity: 1 }}
        transition={{ duration: 0.5, ease: [0.22, 1, 0.36, 1] }}
        className={cn(
          "fixed top-0 inset-x-0 z-50 transition-all duration-300",
          scrolled ? "py-3" : "py-5"
        )}
      >
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <nav
            className={cn(
              "flex items-center justify-between gap-4 rounded-full border px-5 sm:px-6 py-2.5 transition-all duration-300",
              scrolled
                ? "bg-bg-card/90 backdrop-blur-md border-border shadow-sm"
                : "bg-transparent border-transparent"
            )}
          >
            <Link href="/" className="flex items-center" aria-label="Dapur Pintar home">
              <Logo size="md" />
            </Link>

            <div className="hidden md:flex items-center gap-1">
              {NAV_LINKS.map((link) => (
                <Link
                  key={link.href}
                  href={link.href}
                  className="px-3 py-1.5 text-sm text-text-secondary hover:text-text-primary transition-colors duration-200 rounded-md focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
                >
                  {t.nav[link.labelKey.split(".").pop() as keyof typeof t.nav]}
                </Link>
              ))}
            </div>

            <div className="flex items-center gap-2">
              <LanguageToggle className="hidden sm:inline-flex" />
              <Link
                href="/login"
                className="hidden sm:inline-flex text-sm text-text-secondary hover:text-text-primary transition-colors px-3 py-1.5 rounded-md focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
              >
                {t.nav.signin}
              </Link>
              <Link href="/register" className="hidden sm:inline-flex">
                <Button size="sm">{t.nav.getstarted}</Button>
              </Link>
              <button
                type="button"
                onClick={() => setOpen((o) => !o)}
                className="md:hidden flex h-9 w-9 items-center justify-center rounded-md text-text-muted hover:text-text-primary hover:bg-bg-section focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
                aria-label="Toggle menu"
                aria-expanded={open}
              >
                {open ? <X className="h-5 w-5" /> : <Menu className="h-5 w-5" />}
              </button>
            </div>
          </nav>
        </div>
      </motion.header>

      <AnimatePresence>
        {open && (
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.18 }}
            className="fixed inset-0 z-40 md:hidden"
            onClick={() => setOpen(false)}
          >
            <div className="absolute inset-0 bg-bg-overlay" />
            <motion.div
              initial={{ y: -8, opacity: 0 }}
              animate={{ y: 0, opacity: 1 }}
              exit={{ y: -8, opacity: 0 }}
              className="absolute top-24 inset-x-4 bg-bg-card border border-border rounded-2xl p-4 shadow-lg"
              onClick={(e) => e.stopPropagation()}
            >
              <div className="flex flex-col gap-1">
                {NAV_LINKS.map((link) => (
                  <Link
                    key={link.href}
                    href={link.href}
                    onClick={() => setOpen(false)}
                    className="px-3 py-2.5 text-sm text-text-secondary hover:text-text-primary hover:bg-bg-section rounded-md"
                  >
                    {t.nav[link.labelKey.split(".").pop() as keyof typeof t.nav]}
                  </Link>
                ))}
                <div className="border-t border-border my-2" />
                <Link
                  href="/login"
                  onClick={() => setOpen(false)}
                  className="px-3 py-2.5 text-sm text-text-secondary hover:text-text-primary hover:bg-bg-section rounded-md"
                >
                  {t.nav.signin}
                </Link>
                <Link href="/register" onClick={() => setOpen(false)}>
                  <Button className="w-full mt-1">{t.nav.getstarted}</Button>
                </Link>
              </div>
            </motion.div>
          </motion.div>
        )}
      </AnimatePresence>
    </>
  );
}
