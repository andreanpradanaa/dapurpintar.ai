"use client";

import { useState } from "react";
import Link from "next/link";
import Image from "next/image";
import { motion } from "framer-motion";
import { ArrowRight, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Chip } from "@/components/ui/chip";
import { LIFESTYLE_PHOTOS } from "@/lib/photo";

const SUGGESTED_INGREDIENTS = [
  "Chicken",
  "Egg",
  "Garlic",
  "Coconut milk",
  "Lemongrass",
];

export function Hero() {
  const [input, setInput] = useState("");
  const [chips, setChips] = useState<string[]>([]);

  const add = (v: string) => {
    const t = v.trim();
    if (!t || chips.includes(t)) return;
    setChips((c) => [...c, t]);
    setInput("");
  };

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    add(input);
  };

  return (
    <section className="relative pt-36 sm:pt-44 pb-28 overflow-hidden">
      {/* Subtle paper grain texture */}
      <div className="absolute inset-0 -z-10 bg-paper opacity-50" />
      {/* Subtle warm wash on top */}
      <div className="absolute inset-x-0 top-0 h-[600px] -z-10 bg-warm-wash" />

      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="grid lg:grid-cols-12 gap-12 lg:gap-16 items-center">
          {/* Left column — copy + input */}
          <div className="lg:col-span-7 text-center lg:text-left">
            <motion.h1
              initial={{ opacity: 0, y: 12 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.6, ease: [0.22, 1, 0.36, 1] }}
              className="font-display text-balance text-[44px] sm:text-[64px] lg:text-[80px] leading-[0.95] tracking-[-0.035em] text-text-primary"
            >
              What&apos;s in your kitchen?
            </motion.h1>

            <motion.p
              initial={{ opacity: 0, y: 12 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.6, delay: 0.1, ease: [0.22, 1, 0.36, 1] }}
              className="mt-7 text-pretty text-lg sm:text-xl text-text-muted leading-[1.6] max-w-xl mx-auto lg:mx-0"
            >
              A quiet cooking companion. Tell us what you have, and we&apos;ll
              help you decide what to make tonight.
            </motion.p>

            {/* Inline ingredient input */}
            <motion.form
              initial={{ opacity: 0, y: 12 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.6, delay: 0.2, ease: [0.22, 1, 0.36, 1] }}
              onSubmit={onSubmit}
              className="mt-10"
            >
              <div className="flex flex-wrap items-center gap-2 rounded-2xl border border-border bg-bg-card p-2.5 transition-colors focus-within:border-border-strong">
                {chips.map((c) => (
                  <Chip
                    key={c}
                    variant="accent"
                    onRemove={() => setChips((p) => p.filter((x) => x !== c))}
                  >
                    {c}
                  </Chip>
                ))}
                <Plus className="h-4 w-4 text-text-muted ml-1" />
                <input
                  type="text"
                  value={input}
                  onChange={(e) => setInput(e.target.value)}
                  placeholder={chips.length === 0 ? "chicken, egg, garlic, rice…" : "add another"}
                  className="flex-1 min-w-[120px] bg-transparent px-2 py-1.5 text-[15px] text-text-primary placeholder:text-text-subtle focus:outline-none"
                />
                <Link href={`/generate?ingredients=${encodeURIComponent(chips.join(","))}`}>
                  <Button size="md" className="group">
                    Cook
                    <ArrowRight className="h-4 w-4 transition-transform group-hover:translate-x-0.5" />
                  </Button>
                </Link>
              </div>

              <div className="mt-4 flex flex-wrap items-center gap-2 text-xs text-text-muted">
                <span>Try</span>
                {SUGGESTED_INGREDIENTS.filter((s) => !chips.includes(s)).map((s) => (
                  <button
                    key={s}
                    type="button"
                    onClick={() => add(s)}
                    className="px-2.5 py-1 rounded-full border border-border hover:border-border-strong hover:bg-bg-card text-text-secondary transition-colors"
                  >
                    {s}
                  </button>
                ))}
              </div>
            </motion.form>

            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ duration: 0.6, delay: 0.4 }}
              className="mt-10 text-xs text-text-subtle"
            >
              <span>Loved by 8,200+ home cooks</span>
            </motion.div>
          </div>

          {/* Right column — photo */}
          <motion.div
            initial={{ opacity: 0, y: 24 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: 0.3, ease: [0.22, 1, 0.36, 1] }}
            className="lg:col-span-5 relative"
          >
            <div className="relative aspect-[4/5] rounded-2xl overflow-hidden bg-bg-section shadow-lg">
              <Image
                src={LIFESTYLE_PHOTOS.handWithBowl}
                alt="A quiet moment in the kitchen"
                fill
                priority
                sizes="(min-width: 1024px) 40vw, 100vw"
                className="object-cover photo-warm"
              />
              {/* Subtle warm overlay */}
              <div className="absolute inset-0 bg-gradient-to-tr from-[#a8553a]/[0.04] via-transparent to-transparent pointer-events-none" />
            </div>
            {/* Caption beneath */}
            <p className="mt-4 text-xs text-text-subtle text-center italic font-serif">
              Nasi Goreng, somewhere in Jakarta.
            </p>
          </motion.div>
        </div>
      </div>
    </section>
  );
}
