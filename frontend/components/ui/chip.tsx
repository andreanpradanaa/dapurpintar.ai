"use client";

import { motion, AnimatePresence } from "framer-motion";
import { X } from "lucide-react";
import { cn } from "@/lib/utils";
import * as React from "react";

type ChipProps = {
  children: React.ReactNode;
  onRemove?: () => void;
  variant?: "default" | "accent" | "outline";
  size?: "sm" | "md";
  className?: string;
  icon?: React.ReactNode;
};

export function Chip({
  children,
  onRemove,
  variant = "default",
  size = "md",
  className,
  icon,
}: ChipProps) {
  const variantClass = {
    default: "bg-bg-section text-text-primary border-border",
    accent: "bg-accent-soft text-accent-active border-accent-soft-strong",
    outline: "bg-transparent text-text-primary border-border-strong",
  }[variant];

  const sizeClass = {
    sm: "h-7 px-2.5 text-xs gap-1.5 rounded-md",
    md: "h-8 px-3 text-sm gap-1.5 rounded-md",
  }[size];

  return (
    <AnimatePresence mode="popLayout">
      <motion.span
        layout
        initial={{ opacity: 0, scale: 0.85 }}
        animate={{ opacity: 1, scale: 1 }}
        exit={{ opacity: 0, scale: 0.85 }}
        transition={{ duration: 0.18, ease: [0.22, 1, 0.36, 1] }}
        className={cn(
          "inline-flex items-center border font-medium select-none",
          variantClass,
          sizeClass,
          className
        )}
      >
        {icon && <span className="flex items-center">{icon}</span>}
        <span>{children}</span>
        {onRemove && (
          <button
            type="button"
            onClick={onRemove}
            className={cn(
              "ml-0.5 -mr-1 flex items-center justify-center rounded-full p-0.5",
              "text-text-muted hover:text-text-primary hover:bg-bg-base/60",
              "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent"
            )}
            aria-label="Remove"
          >
            <X className="h-3 w-3" strokeWidth={2.5} />
          </button>
        )}
      </motion.span>
    </AnimatePresence>
  );
}
