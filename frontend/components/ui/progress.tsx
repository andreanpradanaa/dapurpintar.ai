"use client";

import * as React from "react";
import { motion } from "framer-motion";
import { cn } from "@/lib/utils";

type ProgressProps = {
  value: number;
  max?: number;
  variant?: "default" | "accent";
  size?: "sm" | "md" | "lg";
  className?: string;
  showLabel?: boolean;
};

export function Progress({
  value,
  max = 100,
  variant = "accent",
  size = "md",
  className,
  showLabel = false,
}: ProgressProps) {
  const percent = Math.max(0, Math.min(100, (value / max) * 100));
  const sizeClass = { sm: "h-1", md: "h-1.5", lg: "h-2.5" }[size];

  return (
    <div className={cn("w-full", className)}>
      <div
        className={cn(
          "w-full overflow-hidden rounded-full bg-bg-section",
          sizeClass
        )}
        role="progressbar"
        aria-valuenow={value}
        aria-valuemin={0}
        aria-valuemax={max}
      >
        <motion.div
          className={cn(
            "h-full rounded-full",
            variant === "accent"
              ? "bg-accent"
              : "bg-text-muted"
          )}
          initial={{ width: 0 }}
          animate={{ width: `${percent}%` }}
          transition={{ duration: 0.6, ease: [0.22, 1, 0.36, 1] }}
        />
      </div>
      {showLabel && (
        <div className="mt-1 text-xs text-text-muted tabular-nums">{Math.round(percent)}%</div>
      )}
    </div>
  );
}

export function CircularProgress({
  value,
  size = 64,
  strokeWidth = 4,
  className,
}: {
  value: number;
  size?: number;
  strokeWidth?: number;
  className?: string;
}) {
  const radius = (size - strokeWidth) / 2;
  const circumference = 2 * Math.PI * radius;
  const offset = circumference - (value / 100) * circumference;
  const percent = Math.max(0, Math.min(100, value));

  return (
    <div className={cn("relative inline-flex items-center justify-center", className)}>
      <svg width={size} height={size} className="-rotate-90">
        <circle
          cx={size / 2}
          cy={size / 2}
          r={radius}
          fill="none"
          stroke="rgba(26, 22, 18, 0.08)"
          strokeWidth={strokeWidth}
        />
        <motion.circle
          cx={size / 2}
          cy={size / 2}
          r={radius}
          fill="none"
          stroke="#A8553A"
          strokeWidth={strokeWidth}
          strokeLinecap="round"
          strokeDasharray={circumference}
          initial={{ strokeDashoffset: circumference }}
          animate={{ strokeDashoffset: offset }}
          transition={{ duration: 0.8, ease: [0.22, 1, 0.36, 1] }}
        />
      </svg>
      <span className="absolute text-xs font-semibold tabular-nums text-text-primary">
        {Math.round(percent)}%
      </span>
    </div>
  );
}
