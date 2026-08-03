"use client";

import { useState } from "react";
import { motion } from "framer-motion";
import { Search, Grid3x3, List } from "lucide-react";
import { Input } from "@/components/ui/input";
import { RecipeCard } from "@/components/app/recipe-card";
import { EmptyState } from "@/components/app/empty-state";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { RECIPES } from "@/lib/mock-data/recipes";
import { cn } from "@/lib/utils";

type View = "grid" | "list";

export default function FavoritesPage() {
  const { t } = useLanguage();
  const favorites = useStore((s) => s.favorites);
  const [view, setView] = useState<View>("grid");
  const [query, setQuery] = useState("");

  const favRecipes = favorites
    .map((id) => RECIPES.find((r) => r.id === id))
    .filter(Boolean) as typeof RECIPES;

  const filtered = favRecipes.filter((r) =>
    !query
      ? true
      : r.title.toLowerCase().includes(query.toLowerCase()) ||
        r.tags.some((t) => t.includes(query.toLowerCase()))
  );

  return (
    <div className="space-y-6">
      <div>
        <h1 className="font-display text-2xl sm:text-3xl font-medium tracking-tight text-text-primary">
          {t.app.favorites.title}
        </h1>
        <p className="mt-1.5 text-sm text-text-muted">{t.app.favorites.subtitle}</p>
      </div>

      {favRecipes.length === 0 ? (
        <EmptyState
          variant="heart"
          title={t.app.favorites.empty.title}
          description={t.app.favorites.empty.desc}
          action={{ label: t.app.favorites.empty.cta, href: "/generate" }}
        />
      ) : (
        <>
          <div className="flex flex-col sm:flex-row items-stretch sm:items-center gap-3">
            <div className="flex-1">
              <Input
                type="text"
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                placeholder="Search favorites…"
                leadingIcon={<Search className="h-4 w-4" />}
              />
            </div>
            <div className="inline-flex items-center gap-1 p-1 rounded-full bg-bg-section border border-border">
              {[
                { value: "grid" as const, icon: Grid3x3, label: "Grid" },
                { value: "list" as const, icon: List, label: "List" },
              ].map((opt) => (
                <button
                  key={opt.value}
                  type="button"
                  onClick={() => setView(opt.value)}
                  aria-label={opt.label}
                  className={cn(
                    "p-1.5 rounded-full transition-colors",
                    view === opt.value
                      ? "bg-bg-card text-text-primary"
                      : "text-text-muted hover:text-text-primary"
                  )}
                >
                  <opt.icon className="h-4 w-4" />
                </button>
              ))}
            </div>
          </div>

          {view === "grid" ? (
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              className="grid sm:grid-cols-2 lg:grid-cols-3 gap-4"
            >
              {filtered.map((r) => (
                <RecipeCard key={r.id} recipe={r} />
              ))}
            </motion.div>
          ) : (
            <div className="space-y-2">
              {filtered.map((r) => (
                <RecipeCard key={r.id} recipe={r} variant="compact" />
              ))}
            </div>
          )}
        </>
      )}
    </div>
  );
}
