"use client";

import { useState } from "react";
import { motion } from "framer-motion";
import { ChefHat, BookOpen, ListOrdered, Utensils, Check } from "lucide-react";
import { useLanguage } from "@/components/providers/language-provider";
import { GENERATION_PHASES } from "@/lib/generate";
import { cn } from "@/lib/utils";

const ICONS = [BookOpen, ListOrdered, Utensils, ChefHat];

export function GenerationLoader({
  active,
  done,
}: {
  active: boolean;
  done: boolean;
}) {
  const { t } = useLanguage();
  const phaseLabels = t.app.generate.phases;
  const [completedCount, setCompletedCount] = useState(0);

  useState(() => {
    if (!active) return;
    const start = Date.now();
    const total = 3200;
    const interval = setInterval(() => {
      const elapsed = Date.now() - start;
      const c = Math.min(GENERATION_PHASES.length, Math.floor((elapsed / total) * GENERATION_PHASES.length) + 1);
      setCompletedCount(c);
      if (elapsed >= total) clearInterval(interval);
    }, 100);
    return () => clearInterval(interval);
  });

  return (
    <motion.div
      initial={{ opacity: 0, y: 12 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: -8 }}
      transition={{ duration: 0.36 }}
      className="rounded-2xl border border-border bg-bg-card p-6 lg:p-8"
    >
      <div className="flex items-center gap-3 mb-6">
        <div className="relative flex h-11 w-11 items-center justify-center rounded-xl bg-bg-section border border-border">
          <motion.div
            className="absolute inset-0 rounded-xl bg-accent-soft"
            animate={{ opacity: [0.3, 0.6, 0.3] }}
            transition={{ duration: 1.6, repeat: Infinity }}
          />
          <ChefHat className="h-5 w-5 text-accent relative z-10" strokeWidth={1.75} />
        </div>
        <div>
          <div className="text-sm font-medium text-text-primary">
            {t.app.generate.thinking}
            <span className="inline-flex ml-1 gap-0.5">
              <span className="h-1 w-1 rounded-full bg-accent pulse-soft" />
              <span className="h-1 w-1 rounded-full bg-accent pulse-soft" style={{ animationDelay: "0.2s" }} />
              <span className="h-1 w-1 rounded-full bg-accent pulse-soft" style={{ animationDelay: "0.4s" }} />
            </span>
          </div>
          <div className="text-xs text-text-muted">
            {done ? t.app.generate.result : "Composing your recipe…"}
          </div>
        </div>
      </div>

      <ol className="space-y-2">
        {GENERATION_PHASES.map((phase, i) => {
          const Icon = ICONS[i];
          const phaseDone = done || i < completedCount;
          const phaseActive = !done && i === completedCount - 1;
          const label = phaseLabels[Object.keys(phaseLabels)[i] as keyof typeof phaseLabels];

          return (
            <li
              key={phase.id}
              className={cn(
                "flex items-center gap-3 p-3 rounded-xl border transition-colors duration-300",
                phaseDone
                  ? "border-accent-soft-strong bg-accent-soft"
                  : phaseActive
                  ? "border-border-strong bg-bg-section"
                  : "border-border bg-bg-card"
              )}
            >
              <div
                className={cn(
                  "flex h-8 w-8 items-center justify-center rounded-lg transition-colors",
                  phaseDone
                    ? "bg-accent text-accent-fg"
                    : phaseActive
                    ? "bg-bg-card text-accent"
                    : "bg-bg-section text-text-subtle"
                )}
              >
                {phaseDone ? <Check className="h-4 w-4" strokeWidth={3} /> : <Icon className="h-4 w-4" strokeWidth={1.75} />}
              </div>
              <span
                className={cn(
                  "text-sm font-medium flex-1",
                  phaseDone || phaseActive ? "text-text-primary" : "text-text-muted"
                )}
              >
                {label}
              </span>
              {phaseActive && (
                <motion.span
                  className="h-1.5 w-1.5 rounded-full bg-accent"
                  animate={{ opacity: [0.3, 1, 0.3] }}
                  transition={{ duration: 1, repeat: Infinity }}
                />
              )}
            </li>
          );
        })}
      </ol>
    </motion.div>
  );
}
