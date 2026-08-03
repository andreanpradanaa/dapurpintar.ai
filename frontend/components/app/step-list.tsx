"use client";

import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { Play, Pause, RotateCcw, Lightbulb } from "lucide-react";
import type { RecipeStep } from "@/lib/types";
import { cn } from "@/lib/utils";

function formatDuration(sec: number) {
  if (sec < 60) return `${sec}s`;
  const m = Math.floor(sec / 60);
  const s = sec % 60;
  return s === 0 ? `${m}m` : `${m}m ${s}s`;
}

export function StepList({ steps }: { steps: RecipeStep[] }) {
  const [activeStep, setActiveStep] = useState(0);
  const [running, setRunning] = useState(false);
  const [remaining, setRemaining] = useState(0);

  const step = steps[activeStep];
  const stepDuration = step?.durationSec ?? 60;

  const startTimer = () => {
    setRunning(true);
    setRemaining(stepDuration);
  };

  const pauseTimer = () => setRunning(false);

  const resetTimer = () => {
    setRunning(false);
    setRemaining(stepDuration);
  };

  if (running && remaining > 0) {
    setTimeout(() => setRemaining((r) => Math.max(0, r - 1)), 1000);
  } else if (running && remaining === 0) {
    setRunning(false);
    if (typeof window !== "undefined") {
      try {
        navigator.vibrate?.(100);
      } catch {}
    }
  }

  return (
    <div className="space-y-4">
      <ol className="space-y-2">
        {steps.map((s, i) => {
          const isActive = i === activeStep;
          const isDone = i < activeStep;
          return (
            <li key={s.order}>
              <button
                type="button"
                onClick={() => setActiveStep(i)}
                className={cn(
                  "w-full text-left flex gap-3 p-4 rounded-xl border transition-all",
                  isActive
                    ? "border-text-muted bg-bg-section"
                    : isDone
                    ? "border-border bg-bg-card opacity-60"
                    : "border-border bg-bg-card hover:border-border-strong"
                )}
              >
                <span
                  className={cn(
                    "flex h-7 w-7 shrink-0 items-center justify-center rounded-full text-xs font-semibold tabular-nums",
                    isActive
                      ? "bg-text-primary text-text-inverse"
                      : isDone
                      ? "bg-success/20 text-success"
                      : "bg-bg-section text-text-muted border border-border"
                  )}
                >
                  {s.order}
                </span>
                <div className="flex-1 min-w-0">
                  <p
                    className={cn(
                      "text-sm leading-[1.6]",
                      isDone ? "text-text-muted line-through" : "text-text-primary"
                    )}
                  >
                    {s.text}
                  </p>
                  {s.tip && (
                    <div className="mt-2 flex items-start gap-1.5 text-xs text-text-secondary bg-bg-section border border-border rounded-md p-2.5">
                      <Lightbulb className="h-3 w-3 mt-0.5 shrink-0 text-warning" />
                      <span>{s.tip}</span>
                    </div>
                  )}
                </div>
                {s.durationSec && (
                  <span className="text-xs text-text-muted tabular-nums shrink-0">
                    {formatDuration(s.durationSec)}
                  </span>
                )}
              </button>
            </li>
          );
        })}
      </ol>

      <AnimatePresence>
        {step?.durationSec && (
          <motion.div
            key={activeStep}
            initial={{ opacity: 0, y: 4 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -4 }}
            className="rounded-xl border border-border bg-bg-card p-4 flex items-center gap-4"
          >
            <div className="flex-1">
              <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">
                Step {step.order} timer
              </div>
              <div className="text-3xl font-serif font-medium text-text-primary tabular-nums tracking-tight">
                {formatDuration(remaining || stepDuration)}
              </div>
            </div>
            <div className="flex items-center gap-2">
              {running ? (
                <button
                  type="button"
                  onClick={pauseTimer}
                  className="flex h-10 w-10 items-center justify-center rounded-full bg-bg-section border border-border text-text-primary hover:border-border-strong"
                  aria-label="Pause"
                >
                  <Pause className="h-4 w-4" />
                </button>
              ) : (
                <button
                  type="button"
                  onClick={startTimer}
                  className="flex h-10 w-10 items-center justify-center rounded-full bg-accent text-accent-fg hover:bg-accent-hover"
                  aria-label="Start"
                >
                  <Play className="h-4 w-4 ml-0.5" fill="currentColor" />
                </button>
              )}
              <button
                type="button"
                onClick={resetTimer}
                className="flex h-10 w-10 items-center justify-center rounded-full bg-bg-section border border-border text-text-primary hover:border-border-strong"
                aria-label="Reset"
              >
                <RotateCcw className="h-4 w-4" />
              </button>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}
