"use client";

import { cn } from "@/lib/utils";

type LogoSize = "sm" | "md" | "lg";

/**
 * Wordmark only. Fraunces 500, ink color.
 * No monogram, no icon. Replaceable.
 */
export function Logo({
  size = "md",
  className,
}: {
  size?: LogoSize;
  className?: string;
}) {
  const sizes = {
    sm: { text: "text-base" },
    md: { text: "text-lg" },
    lg: { text: "text-2xl" },
  }[size];

  return (
    <span
      className={cn(
        "font-serif font-medium tracking-tight text-text-primary",
        sizes.text,
        className
      )}
      style={{ fontVariationSettings: '"opsz" 144, "SOFT" 50' }}
    >
      Dapur Pintar
    </span>
  );
}

/**
 * Compact logo for tight spaces — wordmark with small hairline.
 */
export function LogoMark({ className }: { className?: string }) {
  return (
    <span
      className={cn(
        "font-serif font-medium tracking-tight text-text-primary text-base",
        className
      )}
      style={{ fontVariationSettings: '"opsz" 144, "SOFT" 50' }}
    >
      <span className="text-text-muted mr-2">·</span>
      Dapur Pintar
    </span>
  );
}
