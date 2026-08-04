"use client";

import { motion } from "framer-motion";
import { Soup, BookOpen, Utensils } from "lucide-react";
import { useLanguage } from "@/components/providers/language-provider";

const STEPS = [
  { key: "step1" as const, icon: Soup },
  { key: "step2" as const, icon: BookOpen },
  { key: "step3" as const, icon: Utensils },
];

export function HowItWorks() {
  const { t } = useLanguage();
  return (
    <section id="how" className="relative py-32 border-t border-border bg-bg-section">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="max-w-2xl mb-20">
          <p className="text-xs font-semibold uppercase tracking-[0.12em] text-text-muted mb-4">
            {t.how.eyebrow}
          </p>
          <h2 className="font-display text-balance text-4xl sm:text-5xl lg:text-6xl leading-[1.05] tracking-[-0.025em] text-text-primary">
            {t.how.title}
          </h2>
          <p className="mt-5 text-lg text-text-muted leading-[1.6]">
            {t.how.subtitle}
          </p>
        </div>

        <div className="relative grid md:grid-cols-3 gap-12 md:gap-8">
          {/* Hairline connector rule */}
          <div className="hidden md:block absolute top-7 left-[16%] right-[16%] h-px bg-border" />

          {STEPS.map((step, i) => {
            const Icon = step.icon;
            const data = t.how[step.key];
            return (
              <motion.div
                key={step.key}
                initial={{ opacity: 0, y: 16 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true, margin: "-50px" }}
                transition={{ duration: 0.6, delay: i * 0.15, ease: [0.22, 1, 0.36, 1] }}
                className="relative"
              >
                <div className="relative flex flex-col">
                  <div className="relative inline-flex">
                    <div className="flex h-14 w-14 items-center justify-center rounded-full bg-bg-card border border-border">
                      <Icon className="h-6 w-6 text-accent" strokeWidth={1.5} />
                    </div>
                    <span className="absolute -top-1 -right-1 flex h-6 w-6 items-center justify-center rounded-full bg-text-primary text-text-inverse text-[11px] font-medium tabular-nums">
                      {i + 1}
                    </span>
                  </div>
                  <h3 className="mt-6 text-xl font-serif font-medium tracking-tight text-text-primary">
                    {data.title}
                  </h3>
                  <p className="mt-2.5 text-[15px] text-text-muted leading-[1.6]">
                    {data.desc}
                  </p>
                </div>
              </motion.div>
            );
          })}
        </div>
      </div>
    </section>
  );
}
