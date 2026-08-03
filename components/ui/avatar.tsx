"use client";

import { cn, initials } from "@/lib/utils";
import * as React from "react";

type AvatarProps = {
  src?: string;
  name: string;
  size?: "xs" | "sm" | "md" | "lg" | "xl";
  className?: string;
};

const sizeMap = {
  xs: "h-6 w-6 text-[10px]",
  sm: "h-8 w-8 text-xs",
  md: "h-10 w-10 text-sm",
  lg: "h-12 w-12 text-base",
  xl: "h-16 w-16 text-lg",
};

export function Avatar({ src, name, size = "md", className }: AvatarProps) {
  return (
    <div
      className={cn(
        "relative inline-flex shrink-0 items-center justify-center overflow-hidden rounded-full",
        "bg-gradient-to-br from-bg-section to-bg-base border border-border",
        "text-text-primary font-semibold",
        sizeMap[size],
        className
      )}
      aria-label={name}
    >
      {src ? (
        // eslint-disable-next-line @next/next/no-img-element
        <img src={src} alt={name} className="h-full w-full object-cover" />
      ) : (
        <span>{initials(name)}</span>
      )}
    </div>
  );
}

export function AvatarGroup({
  avatars,
  max = 4,
  size = "sm",
}: {
  avatars: { name: string; src?: string }[];
  max?: number;
  size?: AvatarProps["size"];
}) {
  const visible = avatars.slice(0, max);
  const extra = avatars.length - visible.length;
  return (
    <div className="flex -space-x-2">
      {visible.map((a, i) => (
        <Avatar
          key={i}
          name={a.name}
          src={a.src}
          size={size}
          className="ring-2 ring-bg-base"
        />
      ))}
      {extra > 0 && (
        <div
          className={cn(
            "relative inline-flex shrink-0 items-center justify-center rounded-full",
            "bg-bg-section border border-border text-text-muted font-medium",
            "ring-2 ring-bg-base",
            sizeMap[size!]
          )}
        >
          +{extra}
        </div>
      )}
    </div>
  );
}
