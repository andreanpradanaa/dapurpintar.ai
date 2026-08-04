"use client";

import { useState } from "react";
import Link from "next/link";
import { ArrowRight, Sparkles, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Chip } from "@/components/ui/chip";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { useRouter } from "next/navigation";

const SUGGESTED = ["Chicken", "Egg", "Rice", "Garlic", "Chili", "Coconut milk"];

export function QuickGenerateCard() {
  const { t } = useLanguage();
  const router = useRouter();
  const pantry = useStore((s) => s.pantry);
  const [chips, setChips] = useState<string[]>([]);
  const [input, setInput] = useState("");

  const addChip = (v: string) => {
    const trimmed = v.trim();
    if (!trimmed || chips.includes(trimmed)) return;
    setChips((c) => [...c, trimmed]);
    setInput("");
  };

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    addChip(input);
  };

  const onGenerate = () => {
    const ingredients = chips.length > 0 ? chips : pantry.map((p) => p.name);
    if (ingredients.length === 0) return;
    const params = new URLSearchParams({ ingredients: ingredients.join(",") });
    router.push(`/generate?${params.toString()}`);
  };

  return (
    <div className="relative rounded-2xl border border-border bg-bg-card p-6 lg:p-8">
      <div className="relative">
        <div className="flex items-center gap-2 mb-2">
          <span className="text-[10px] font-semibold uppercase tracking-[0.12em] text-text-muted">
            Quick start
          </span>
        </div>
        <h2 className="font-serif text-2xl sm:text-3xl font-medium tracking-tight text-text-primary">
          What are you cooking today?
        </h2>
        <p className="mt-1.5 text-sm text-text-muted">
          Add what you have. We&apos;ll handle the recipe.
        </p>

        <form onSubmit={onSubmit} className="mt-6">
          {chips.length > 0 && (
            <div className="flex flex-wrap gap-1.5 mb-2.5">
              {chips.map((c) => (
                <Chip
                  key={c}
                  variant="accent"
                  onRemove={() => setChips((prev) => prev.filter((x) => x !== c))}
                >
                  {c}
                </Chip>
              ))}
            </div>
          )}
          <div className="flex items-center gap-2 rounded-xl border border-border bg-bg-card focus-within:border-border-strong transition-colors p-1.5">
            <Plus className="h-4 w-4 text-text-muted ml-2" />
            <input
              type="text"
              value={input}
              onChange={(e) => setInput(e.target.value)}
              placeholder="Add an ingredient and press Enter…"
              className="flex-1 bg-transparent px-2 py-2 text-sm text-text-primary placeholder:text-text-subtle focus:outline-none"
            />
            <Button type="submit" variant="ghost" size="sm" onClick={onSubmit}>
              Add
            </Button>
          </div>
        </form>

        <div className="mt-3 flex flex-wrap gap-1.5">
          {SUGGESTED.filter((s) => !chips.includes(s)).slice(0, 4).map((s) => (
            <button
              key={s}
              type="button"
              onClick={() => addChip(s)}
              className="inline-flex items-center gap-1 text-xs text-text-muted hover:text-text-primary px-2.5 py-1 rounded-full border border-border hover:border-border-strong transition-colors"
            >
              <Plus className="h-3 w-3" />
              {s}
            </button>
          ))}
        </div>

        <div className="mt-6 flex flex-wrap items-center gap-2">
          <Button onClick={onGenerate} className="group">
            <Sparkles className="h-4 w-4" />
            {t.app.generate.generate}
            <ArrowRight className="h-4 w-4 transition-transform group-hover:translate-x-0.5" />
          </Button>
          <Link
            href="/generate"
            className="text-xs text-text-muted hover:text-text-primary inline-flex items-center gap-1"
          >
            Open full creator
            <ArrowRight className="h-3 w-3" />
          </Link>
        </div>
      </div>
    </div>
  );
}
