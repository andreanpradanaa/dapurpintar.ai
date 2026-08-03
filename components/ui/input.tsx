"use client";

import * as React from "react";
import { cn } from "@/lib/utils";

type InputProps = React.InputHTMLAttributes<HTMLInputElement> & {
  invalid?: boolean;
  leadingIcon?: React.ReactNode;
  trailingIcon?: React.ReactNode;
};

export const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className, invalid, leadingIcon, trailingIcon, type = "text", ...props }, ref) => {
    return (
      <div
        className={cn(
          "group relative flex h-11 w-full items-center rounded-lg border bg-bg-card transition-colors duration-200",
          "focus-within:border-accent",
          invalid
            ? "border-danger/60"
            : "border-border hover:border-border-strong",
          className
        )}
      >
        {leadingIcon && (
          <span className="pl-3.5 pr-2 text-text-muted group-focus-within:text-text-secondary">
            {leadingIcon}
          </span>
        )}
        <input
          ref={ref}
          type={type}
          className={cn(
            "flex-1 bg-transparent text-[15px] text-text-primary placeholder:text-text-subtle",
            "focus:outline-none disabled:opacity-50",
            leadingIcon ? "pl-0" : "pl-3.5",
            trailingIcon ? "pr-2" : "pr-3.5",
            "h-full"
          )}
          {...props}
        />
        {trailingIcon && (
          <span className="pr-3.5 pl-2 text-text-muted group-focus-within:text-text-secondary">
            {trailingIcon}
          </span>
        )}
      </div>
    );
  }
);
Input.displayName = "Input";

type TextareaProps = React.TextareaHTMLAttributes<HTMLTextAreaElement> & { invalid?: boolean };

export const Textarea = React.forwardRef<HTMLTextAreaElement, TextareaProps>(
  ({ className, invalid, ...props }, ref) => (
    <textarea
      ref={ref}
      className={cn(
        "flex w-full rounded-lg border bg-bg-card px-3.5 py-3 text-[15px] text-text-primary placeholder:text-text-subtle",
        "transition-colors duration-200 resize-y min-h-[88px] leading-relaxed",
        "focus:outline-none focus:border-accent",
        invalid ? "border-danger/60" : "border-border hover:border-border-strong",
        className
      )}
      {...props}
    />
  )
);
Textarea.displayName = "Textarea";
