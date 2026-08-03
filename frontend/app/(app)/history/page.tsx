"use client";

import { useState, useMemo } from "react";
import Link from "next/link";
import { motion, AnimatePresence } from "framer-motion";
import { Search, Trash2, ChevronRight } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { EmptyState } from "@/components/app/empty-state";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { getRecipeById } from "@/lib/mock-data/recipes";
import { relativeTime } from "@/lib/utils";
import { toast } from "sonner";

type Filter = "all" | "favorites" | "thisWeek";

export default function HistoryPage() {
  const { t } = useLanguage();
  const history = useStore((s) => s.history);
  const favorites = useStore((s) => s.favorites);
  const removeHistory = useStore((s) => s.removeHistory);
  const clearHistory = useStore((s) => s.clearHistory);
  const [query, setQuery] = useState("");
  const [filter, setFilter] = useState<Filter>("all");

  const weekAgo = Date.now() - 7 * 24 * 60 * 60 * 1000;

  const filtered = useMemo(() => {
    return history
      .filter((h) => {
        if (filter === "favorites") return favorites.includes(h.recipeId);
        if (filter === "thisWeek") return new Date(h.createdAt).getTime() > weekAgo;
        return true;
      })
      .filter((h) => {
        if (!query) return true;
        const recipe = getRecipeById(h.recipeId);
        return (
          h.ingredients.some((i) => i.toLowerCase().includes(query.toLowerCase())) ||
          recipe?.title.toLowerCase().includes(query.toLowerCase())
        );
      });
  }, [history, filter, favorites, query, weekAgo]);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="font-display text-2xl sm:text-3xl font-medium tracking-tight text-text-primary">
          {t.app.history.title}
        </h1>
        <p className="mt-1.5 text-sm text-text-muted">{t.app.history.subtitle}</p>
      </div>

      <div className="flex flex-col sm:flex-row items-stretch sm:items-center gap-3">
        <div className="flex-1">
          <Input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder={t.app.history.search}
            leadingIcon={<Search className="h-4 w-4" />}
          />
        </div>
        <div className="inline-flex items-center gap-1 p-1 rounded-full bg-bg-section border border-border">
          {[
            { value: "all" as const, label: t.app.history.filterAll },
            { value: "thisWeek" as const, label: t.app.history.filterThisWeek },
            { value: "favorites" as const, label: t.app.history.filterFavorites },
          ].map((opt) => (
            <button
              key={opt.value}
              type="button"
              onClick={() => setFilter(opt.value)}
              className={
                "px-3 py-1.5 text-xs font-medium rounded-full transition-colors " +
                (filter === opt.value
                  ? "bg-bg-card text-text-primary shadow-sm"
                  : "text-text-muted hover:text-text-primary")
              }
            >
              {opt.label}
            </button>
          ))}
        </div>
        {history.length > 0 && (
          <Button
            variant="ghost"
            size="sm"
            onClick={() => {
              if (window.confirm("Clear all history?")) {
                clearHistory();
                toast.success("History cleared");
              }
            }}
          >
            <Trash2 className="h-3.5 w-3.5" />
            Clear all
          </Button>
        )}
      </div>

      {filtered.length === 0 ? (
        <EmptyState
          variant={history.length === 0 ? "chef" : "search"}
          title={
            history.length === 0
              ? t.app.history.empty.title
              : "No matches found"
          }
          description={
            history.length === 0
              ? t.app.history.empty.desc
              : "Try adjusting your search or filters."
          }
          action={
            history.length === 0
              ? { label: t.app.history.empty.cta, href: "/generate" }
              : undefined
          }
        />
      ) : (
        <div className="rounded-2xl border border-border bg-bg-card overflow-hidden">
          <ul className="divide-y divide-border">
            <AnimatePresence>
              {filtered.map((h, i) => {
                const recipe = getRecipeById(h.recipeId);
                if (!recipe) return null;
                const isFav = favorites.includes(h.recipeId);
                return (
                  <motion.li
                    key={h.id}
                    initial={{ opacity: 0, y: 4 }}
                    animate={{ opacity: 1, y: 0 }}
                    exit={{ opacity: 0, x: -10 }}
                    transition={{ delay: i * 0.03 }}
                    className="group flex items-center gap-4 p-4 hover:bg-bg-section transition-colors"
                  >
                    <Link href={`/recipes/${recipe.slug}`} className="flex items-center gap-4 flex-1 min-w-0">
                      <div className="min-w-0 flex-1">
                        <div className="flex items-center gap-2 mb-0.5">
                          <h3 className="text-sm font-medium text-text-primary truncate font-serif">
                            {recipe.title}
                          </h3>
                          {isFav && <Heart className="h-3 w-3 text-accent-active" fill="currentColor" />}
                        </div>
                        <div className="text-xs text-text-muted truncate">
                          {h.ingredients.slice(0, 5).join(" · ")}
                          {h.ingredients.length > 5 && ` +${h.ingredients.length - 5}`}
                        </div>
                        <div className="mt-1.5 flex items-center gap-2 text-[10px] text-text-subtle uppercase tracking-[0.1em]">
                          <span>{relativeTime(h.createdAt)}</span>
                          <span>·</span>
                          <span>{recipe.cuisine}</span>
                          {h.dietary.length > 0 && (
                            <>
                              <span>·</span>
                              <span>{h.dietary.length} filters</span>
                            </>
                          )}
                        </div>
                      </div>
                      <ChevronRight className="h-4 w-4 text-text-muted group-hover:text-text-primary transition-colors" />
                    </Link>
                    <div className="flex items-center gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                      <button
                        type="button"
                        onClick={() => {
                          removeHistory(h.id);
                          toast.success("Removed");
                        }}
                        className="flex h-8 w-8 items-center justify-center rounded-md text-text-muted hover:text-danger hover:bg-bg-section"
                        aria-label="Remove"
                      >
                        <Trash2 className="h-3.5 w-3.5" />
                      </button>
                    </div>
                  </motion.li>
                );
              })}
            </AnimatePresence>
          </ul>
        </div>
      )}
    </div>
  );
}

function Heart(props: { className?: string; fill?: string }) {
  return (
    <svg viewBox="0 0 24 24" className={props.className} fill={props.fill} stroke="currentColor" strokeWidth="2">
      <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z" />
    </svg>
  );
}
