"use client";

import { motion } from "framer-motion";
import { ChefHat, Heart, Plus, Inbox, Search } from "lucide-react";
import { cn } from "@/lib/utils";

const ICONS = {
  inbox: Inbox,
  heart: Heart,
  chef: ChefHat,
  search: Search,
  plus: Plus,
};

type Variant = keyof typeof ICONS;

export function EmptyState({
  variant = "inbox",
  title,
  description,
  action,
  className,
}: {
  variant?: Variant;
  title: string;
  description?: string;
  action?: { label: string; href?: string; onClick?: () => void };
  className?: string;
}) {
  const Icon = ICONS[variant];

  return (
    <motion.div
      initial={{ opacity: 0, y: 8 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.36 }}
      className={cn(
        "flex flex-col items-center justify-center text-center py-20 px-6 rounded-2xl border border-dashed border-border bg-bg-card/40",
        className
      )}
    >
      <div className="relative inline-flex items-center justify-center mb-5">
        <div className="flex h-14 w-14 items-center justify-center rounded-full bg-bg-section border border-border">
          <Icon className="h-6 w-6 text-accent" strokeWidth={1.5} />
        </div>
      </div>
      <h3 className="text-base font-medium text-text-primary font-serif">{title}</h3>
      {description && (
        <p className="mt-2 text-sm text-text-muted max-w-sm leading-[1.6]">{description}</p>
      )}
      {action && (
        <div className="mt-6">
          {action.href ? (
            <a
              href={action.href}
              className="inline-flex items-center gap-1.5 rounded-lg bg-accent px-4 py-2 text-sm font-medium text-accent-fg hover:bg-accent-hover transition-colors"
            >
              {action.label}
            </a>
          ) : (
            <button
              type="button"
              onClick={action.onClick}
              className="inline-flex items-center gap-1.5 rounded-lg bg-accent px-4 py-2 text-sm font-medium text-accent-fg hover:bg-accent-hover transition-colors"
            >
              {action.label}
            </button>
          )}
        </div>
      )}
    </motion.div>
  );
}
