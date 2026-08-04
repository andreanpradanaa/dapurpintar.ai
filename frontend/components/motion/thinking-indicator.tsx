"use client";

import { motion } from "framer-motion";
import { useLanguage } from "@/components/providers/language-provider";
import { ChefHat, Heart } from "lucide-react";

export function ThinkingIndicator() {
  const { t } = useLanguage();
  return (
    <div className="flex items-center gap-2.5">
      <div className="relative flex h-9 w-9 items-center justify-center rounded-lg bg-bg-section">
        <motion.div
          className="absolute inset-0 rounded-lg bg-accent-soft"
          animate={{ opacity: [0.3, 0.6, 0.3] }}
          transition={{ duration: 1.6, repeat: Infinity }}
        />
        <ChefHat className="h-4 w-4 text-accent relative z-10" strokeWidth={1.75} />
      </div>
      <div className="flex items-center gap-1 text-sm text-text-secondary">
        <span>{t.app.generate.thinking}</span>
        <span className="flex gap-0.5">
          <motion.span
            className="h-1 w-1 rounded-full bg-accent"
            animate={{ opacity: [0.3, 1, 0.3] }}
            transition={{ duration: 1, repeat: Infinity, delay: 0 }}
          />
          <motion.span
            className="h-1 w-1 rounded-full bg-accent"
            animate={{ opacity: [0.3, 1, 0.3] }}
            transition={{ duration: 1, repeat: Infinity, delay: 0.2 }}
          />
          <motion.span
            className="h-1 w-1 rounded-full bg-accent"
            animate={{ opacity: [0.3, 1, 0.3] }}
            transition={{ duration: 1, repeat: Infinity, delay: 0.4 }}
          />
        </span>
      </div>
    </div>
  );
}

export function FavoriteHeart({ active, onClick }: { active: boolean; onClick?: () => void }) {
  return (
    <button
      type="button"
      onClick={onClick}
      className="group flex h-10 w-10 items-center justify-center rounded-full border border-border bg-bg-card hover:border-border-strong transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
      aria-label={active ? "Remove from favorites" : "Add to favorites"}
    >
      <Heart
        className="h-4 w-4 transition-colors"
        strokeWidth={1.75}
        fill={active ? "#A8553A" : "transparent"}
        color={active ? "#A8553A" : "currentColor"}
      />
    </button>
  );
}
