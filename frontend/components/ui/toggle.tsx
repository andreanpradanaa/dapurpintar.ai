"use client";

import { cn } from "@/lib/utils";
import * as React from "react";

export function Switch({
  className,
  checked,
  onCheckedChange,
  disabled,
  id,
}: {
  className?: string;
  checked: boolean;
  onCheckedChange: (v: boolean) => void;
  disabled?: boolean;
  id?: string;
}) {
  return (
    <button
      id={id}
      type="button"
      role="switch"
      aria-checked={checked}
      disabled={disabled}
      onClick={() => onCheckedChange(!checked)}
      className={cn(
        "relative inline-flex h-6 w-11 shrink-0 cursor-pointer items-center rounded-full border transition-colors duration-200",
        "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg-base",
        "disabled:opacity-50 disabled:cursor-not-allowed",
        checked
          ? "bg-accent border-accent"
          : "bg-bg-section border-border-strong",
        className
      )}
    >
      <span
        className={cn(
          "pointer-events-none inline-block h-4 w-4 transform rounded-full bg-white shadow-sm transition-transform duration-200",
          checked ? "translate-x-[22px]" : "translate-x-[3px]"
        )}
      />
    </button>
  );
}

export function Checkbox({
  className,
  checked,
  onCheckedChange,
  label,
  disabled,
  id,
}: {
  className?: string;
  checked: boolean;
  onCheckedChange: (v: boolean) => void;
  label?: React.ReactNode;
  disabled?: boolean;
  id?: string;
}) {
  return (
    <label
      className={cn(
        "inline-flex items-center gap-2.5 cursor-pointer select-none",
        disabled && "opacity-50 cursor-not-allowed",
        className
      )}
    >
      <button
        id={id}
        type="button"
        role="checkbox"
        aria-checked={checked}
        disabled={disabled}
        onClick={() => onCheckedChange(!checked)}
        className={cn(
          "flex h-4 w-4 shrink-0 items-center justify-center rounded border transition-colors duration-150",
          "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg-base",
          checked
            ? "bg-accent border-accent"
            : "bg-bg-card border-border-strong"
        )}
      >
        {checked && (
          <svg viewBox="0 0 12 12" className="h-3 w-3 text-accent-fg" fill="none" stroke="currentColor" strokeWidth="2.5" strokeLinecap="round" strokeLinejoin="round">
            <path d="M2.5 6L5 8.5L9.5 3.5" />
          </svg>
        )}
      </button>
      {label && <span className="text-sm text-text-primary">{label}</span>}
    </label>
  );
}
