"use client";

import * as React from "react";
import { motion, type HTMLMotionProps } from "framer-motion";
import { cn } from "@/lib/utils";

type Variant = "primary" | "secondary" | "ghost" | "outline" | "destructive" | "link";
type Size = "sm" | "md" | "lg" | "icon" | "icon-sm" | "icon-lg";

const variantClasses: Record<Variant, string> = {
  primary:
    "bg-cta text-cta-fg hover:bg-cta-hover active:bg-cta-active shadow-[0_4px_16px_rgba(5,150,105,0.16)] hover:shadow-[0_6px_20px_rgba(5,150,105,0.20)]",
  secondary:
    "bg-bg-card text-text-primary border border-border hover:bg-bg-base hover:border-border-strong",
  ghost:
    "text-text-primary hover:bg-bg-base active:bg-bg-section",
  outline:
    "border border-border-strong text-text-primary hover:bg-bg-card hover:border-text-muted",
  destructive:
    "bg-danger text-text-inverse hover:opacity-90 active:opacity-80",
  link: "text-accent hover:text-accent-hover underline-offset-4 hover:underline",
};

const sizeClasses: Record<Size, string> = {
  sm: "h-9 px-3.5 text-sm gap-1.5 rounded-md",
  md: "h-11 px-5 text-sm gap-2 rounded-lg",
  lg: "h-12 px-6 text-[15px] gap-2 rounded-lg",
  icon: "h-11 w-11 rounded-lg",
  "icon-sm": "h-9 w-9 rounded-md",
  "icon-lg": "h-12 w-12 rounded-lg",
};

type ButtonProps = Omit<HTMLMotionProps<"button">, "ref" | "children"> & {
  variant?: Variant;
  size?: Size;
  loading?: boolean;
  asChild?: boolean;
  children?: React.ReactNode;
};

export const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  (
    {
      className,
      variant = "primary",
      size = "md",
      loading = false,
      disabled,
      children,
      ...props
    },
    ref
  ) => {
    return (
      <motion.button
        ref={ref}
        whileHover={{ y: -1 }}
        whileTap={{ y: 0 }}
        transition={{ duration: 0.18, ease: [0.22, 1, 0.36, 1] }}
        disabled={disabled || loading}
        className={cn(
          "relative inline-flex items-center justify-center font-medium transition-colors duration-200 select-none",
          "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg-base",
          "disabled:opacity-50 disabled:pointer-events-none",
          "whitespace-nowrap",
          variantClasses[variant],
          sizeClasses[size],
          className
        )}
        {...props}
      >
        {loading && (
          <span className="absolute inset-0 flex items-center justify-center">
            <span className="h-4 w-4 rounded-full border-2 border-current border-r-transparent animate-spin" />
          </span>
        )}
        <span className={cn("inline-flex items-center justify-center", loading && "invisible")}>
          {children}
        </span>
      </motion.button>
    );
  }
);
Button.displayName = "Button";
