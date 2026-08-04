"use client";

import { Languages } from "lucide-react";
import { useLanguage } from "@/components/providers/language-provider";
import { motion, AnimatePresence } from "framer-motion";

export function LanguageToggle({ className }: { className?: string }) {
  const { lang, toggle } = useLanguage();
  return (
    <button
      type="button"
      onClick={toggle}
      className={
        "relative inline-flex h-9 items-center gap-1.5 rounded-md border border-border bg-bg-card px-2.5 text-xs font-medium text-text-secondary hover:text-text-primary hover:border-border-strong transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent " +
        (className ?? "")
      }
      aria-label="Toggle language"
    >
      <Languages className="h-3.5 w-3.5" strokeWidth={2} />
      <span className="relative h-3.5 w-7 overflow-hidden">
        <AnimatePresence mode="wait" initial={false}>
          <motion.span
            key={lang}
            initial={{ y: 8, opacity: 0 }}
            animate={{ y: 0, opacity: 1 }}
            exit={{ y: -8, opacity: 0 }}
            transition={{ duration: 0.18 }}
            className="absolute inset-0 flex items-center justify-center tabular-nums tracking-wider"
          >
            {lang === "en" ? "EN" : "ID"}
          </motion.span>
        </AnimatePresence>
      </span>
    </button>
  );
}
