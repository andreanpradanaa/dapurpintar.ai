"use client";

import { Sparkles, BarChart3, ListOrdered, Lightbulb, Heart, ArrowUpRight } from "lucide-react";
import { useLanguage } from "@/components/providers/language-provider";
import { Stagger, StaggerItem } from "@/components/motion/reveal";

const ICONS = [Sparkles, BarChart3, ListOrdered, Lightbulb, Heart];

export function Features() {
  const { t } = useLanguage();
  const items = t.features.items;

  return (
    <section id="features" className="relative py-32">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="max-w-2xl mb-16">
          <p className="text-xs font-semibold uppercase tracking-[0.12em] text-text-muted mb-4">
            {t.features.eyebrow}
          </p>
          <h2 className="font-display text-balance text-4xl sm:text-5xl lg:text-6xl leading-[1.05] tracking-[-0.025em] text-text-primary">
            {t.features.title}
          </h2>
          <p className="mt-5 text-lg text-text-muted leading-[1.6]">
            {t.features.subtitle}
          </p>
        </div>

        <Stagger className="grid md:grid-cols-2 lg:grid-cols-3 gap-5" staggerChildren={0.08}>
          {Object.values(items).map((item, i) => {
            const Icon = ICONS[i];
            return (
              <StaggerItem
                key={i}
                className={
                  i === 0
                    ? "md:col-span-2 lg:col-span-2 lg:row-span-2"
                    : ""
                }
              >
                <FeatureCard
                  title={item.title}
                  desc={item.desc}
                  icon={<Icon className="h-5 w-5" strokeWidth={1.5} />}
                  highlight={i === 0}
                />
              </StaggerItem>
            );
          })}
        </Stagger>
      </div>
    </section>
  );
}

function FeatureCard({
  title,
  desc,
  icon,
  highlight,
}: {
  title: string;
  desc: string;
  icon: React.ReactNode;
  highlight?: boolean;
}) {
  return (
    <div
      className={
        "group relative h-full rounded-2xl border border-border bg-bg-card p-6 lg:p-8 transition-all duration-300 overflow-hidden " +
        "hover:border-border-strong hover:-translate-y-0.5 hover:shadow-md " +
        (highlight ? "lg:min-h-[320px]" : "")
      }
    >
      <div className="relative h-full flex flex-col">
        <div
          className={
            "inline-flex h-10 w-10 items-center justify-center rounded-lg border " +
            (highlight
              ? "bg-accent-soft text-accent-active border-accent-soft-strong"
              : "bg-bg-section text-text-secondary border-border")
          }
        >
          {icon}
        </div>
        <h3 className="mt-6 text-lg lg:text-xl font-semibold tracking-tight text-text-primary">
          {title}
        </h3>
        <p className="mt-2.5 text-[15px] text-text-muted leading-[1.6] flex-1">
          {desc}
        </p>
        <div className="mt-6 flex items-center gap-1 text-xs font-medium text-text-muted opacity-0 group-hover:opacity-100 group-hover:text-accent-active transition-all duration-300">
          <span>Read more</span>
          <ArrowUpRight className="h-3.5 w-3.5" />
        </div>
      </div>
    </div>
  );
}
