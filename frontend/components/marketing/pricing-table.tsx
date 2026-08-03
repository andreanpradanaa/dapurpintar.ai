"use client";

import { useState } from "react";
import Link from "next/link";
import { motion } from "framer-motion";
import { Check } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { useLanguage } from "@/components/providers/language-provider";
import { cn } from "@/lib/utils";

type Cycle = "monthly" | "yearly";

const PLAN_ORDER = ["free", "pro", "family"] as const;

export function PricingTable() {
  const { t } = useLanguage();
  const [cycle, setCycle] = useState<Cycle>("yearly");

  return (
    <section id="pricing" className="relative py-32 border-t border-border">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="max-w-2xl mb-12">
          <p className="text-xs font-semibold uppercase tracking-[0.12em] text-text-muted mb-4">
            {t.pricing.eyebrow}
          </p>
          <h2 className="font-display text-balance text-4xl sm:text-5xl lg:text-6xl leading-[1.05] tracking-[-0.025em] text-text-primary">
            {t.pricing.title}
          </h2>
          <p className="mt-5 text-lg text-text-muted leading-[1.6]">
            {t.pricing.subtitle}
          </p>
        </div>

        <div className="flex items-center justify-center mb-12">
          <div className="inline-flex items-center gap-1 p-1 rounded-full bg-bg-section border border-border">
            <button
              type="button"
              onClick={() => setCycle("monthly")}
              className={cn(
                "px-4 py-1.5 text-sm font-medium rounded-full transition-colors",
                cycle === "monthly"
                  ? "bg-bg-card text-text-primary shadow-sm"
                  : "text-text-muted hover:text-text-primary"
              )}
            >
              {t.pricing.monthly}
            </button>
            <button
              type="button"
              onClick={() => setCycle("yearly")}
              className={cn(
                "px-4 py-1.5 text-sm font-medium rounded-full transition-colors inline-flex items-center gap-1.5",
                cycle === "yearly"
                  ? "bg-bg-card text-text-primary shadow-sm"
                  : "text-text-muted hover:text-text-primary"
              )}
            >
              {t.pricing.yearly}
              <span className="text-[10px] font-semibold text-accent-active bg-accent-soft px-1.5 py-0.5 rounded uppercase tracking-[0.1em]">
                {t.pricing.save}
              </span>
            </button>
          </div>
        </div>

        <div className="grid lg:grid-cols-3 gap-5">
          {PLAN_ORDER.map((key, idx) => {
            const plan = t.pricing.plans[key];
            const featured = key === "pro";
            return (
              <motion.div
                key={key}
                initial={{ opacity: 0, y: 16 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ duration: 0.5, delay: idx * 0.08, ease: [0.22, 1, 0.36, 1] }}
                className={cn(
                  "relative rounded-2xl border p-7 transition-all duration-300",
                  featured
                    ? "border-accent/40 bg-accent-soft/30"
                    : "border-border bg-bg-card hover:border-border-strong"
                )}
              >
                {featured && "badge" in plan && plan.badge && (
                  <div className="absolute -top-3 left-7">
                    <Badge variant="accent" className="bg-accent text-accent-fg border-accent">
                      {plan.badge}
                    </Badge>
                  </div>
                )}

                <h3 className="text-base font-medium text-text-primary mb-1 font-serif">
                  {plan.name}
                </h3>
                <p className="text-sm text-text-muted leading-[1.6] mb-6 min-h-[44px]">
                  {plan.desc}
                </p>

                <div className="flex items-baseline gap-1 mb-7">
                  <span className="text-5xl font-serif font-medium text-text-primary tracking-tight tabular-nums">
                    {plan.price}
                  </span>
                  {key !== "free" && (
                    <span className="text-text-muted text-sm">{t.pricing.perMonth}</span>
                  )}
                </div>

                <Link href={featured ? "/register?plan=pro" : "/register"}>
                  <Button
                    variant={featured ? "primary" : "outline"}
                    className="w-full"
                    size="md"
                  >
                    {featured ? t.pricing.ctaPrimary : t.pricing.cta}
                  </Button>
                </Link>

                <ul className="mt-7 space-y-3">
                  {plan.features.map((f, i) => (
                    <li key={i} className="flex items-start gap-2.5 text-sm text-text-secondary">
                      <span
                        className={cn(
                          "mt-0.5 flex h-4 w-4 shrink-0 items-center justify-center rounded-full",
                          featured ? "bg-accent-soft text-accent-active" : "bg-bg-section text-text-muted"
                        )}
                      >
                        <Check className="h-2.5 w-2.5" strokeWidth={3} />
                      </span>
                      <span>{f}</span>
                    </li>
                  ))}
                </ul>
              </motion.div>
            );
          })}
        </div>
      </div>
    </section>
  );
}
