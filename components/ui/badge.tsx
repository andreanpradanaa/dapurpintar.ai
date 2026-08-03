import * as React from "react";
import { cn } from "@/lib/utils";

type BadgeVariant = "default" | "accent" | "success" | "warning" | "danger" | "info" | "outline";

const variantClasses: Record<BadgeVariant, string> = {
  default: "bg-bg-section text-text-secondary border-border",
  accent: "bg-accent-soft text-accent-active border-accent-soft-strong",
  success: "bg-[#5c7a4d]/10 text-[#3d5230] border-[#5c7a4d]/20",
  warning: "bg-[#b8804a]/10 text-[#7a4f24] border-[#b8804a]/20",
  danger: "bg-danger/10 text-danger border-danger/20",
  info: "bg-[#5a7a8a]/10 text-[#3a5566] border-[#5a7a8a]/20",
  outline: "bg-transparent text-text-muted border-border-strong",
};

export function Badge({
  variant = "default",
  className,
  ...props
}: React.HTMLAttributes<HTMLSpanElement> & { variant?: BadgeVariant }) {
  return (
    <span
      className={cn(
        "inline-flex items-center gap-1 rounded-md border px-2 py-0.5 text-xs font-medium",
        variantClasses[variant],
        className
      )}
      {...props}
    />
  );
}
