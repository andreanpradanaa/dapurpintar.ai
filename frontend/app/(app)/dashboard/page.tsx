"use client";

import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { QuickGenerateCard } from "@/components/app/quick-generate";
import { RecipeCard } from "@/components/app/recipe-card";
import { NutritionLineChart } from "@/components/app/nutrition-line-chart";
import { Stagger, StaggerItem } from "@/components/motion/reveal";
import { RECIPES } from "@/lib/mock-data/recipes";
import { Flame, TrendingUp, Utensils, Award } from "lucide-react";
import { INGREDIENT_CATEGORIES } from "@/lib/mock-data/ingredients";
import Link from "next/link";
import { ArrowRight } from "lucide-react";

const CHART_DATA = [
  { label: "Mon", value: 1820 },
  { label: "Tue", value: 2100 },
  { label: "Wed", value: 1680 },
  { label: "Thu", value: 2350 },
  { label: "Fri", value: 1980 },
  { label: "Sat", value: 2520 },
  { label: "Sun", value: 2240 },
];

const STATS = [
  {
    icon: Flame,
    label: "Calories avg",
    value: "2,098",
    delta: "+8%",
    color: "text-warning",
  },
  {
    icon: TrendingUp,
    label: "Protein",
    value: "112g",
    delta: "+12%",
    color: "text-info",
  },
  {
    icon: Utensils,
    label: "Recipes",
    value: "7",
    delta: "+3",
    color: "text-accent",
  },
  {
    icon: Award,
    label: "Streak",
    value: "23d",
    delta: "Personal best",
    color: "text-success",
  },
];

export default function DashboardPage() {
  const { t } = useLanguage();
  const user = useStore((s) => s.user);
  const pantry = useStore((s) => s.pantry);
  const favorites = useStore((s) => s.favorites);
  const history = useStore((s) => s.history);

  const recentRecipes = history
    .slice(0, 3)
    .map((h) => RECIPES.find((r) => r.id === h.recipeId))
    .filter(Boolean) as typeof RECIPES;

  const greeting =
    new Date().getHours() < 18 ? t.app.dashboard.greeting : t.app.dashboard.greetingEvening;

  return (
    <div className="space-y-8">
      {/* Greeting */}
      <div>
        <h1 className="font-display text-3xl sm:text-4xl font-medium tracking-tight text-text-primary">
          {greeting}, {user?.name?.split(" ")[0]}.
        </h1>
        <p className="mt-1.5 text-sm text-text-muted">{t.app.dashboard.sub}</p>
      </div>

      <QuickGenerateCard />

      <Stagger className="grid grid-cols-2 lg:grid-cols-4 gap-4" staggerChildren={0.06}>
        {STATS.map((s) => (
          <StaggerItem key={s.label}>
            <div className="rounded-2xl border border-border bg-bg-card p-4 hover:border-border-strong transition-colors">
              <div className="flex items-center justify-between mb-3">
                <div className={`flex h-8 w-8 items-center justify-center rounded-md bg-bg-section ${s.color}`}>
                  <s.icon className="h-4 w-4" strokeWidth={1.75} />
                </div>
                <span className="text-[10px] font-medium text-text-muted tabular-nums">
                  {s.delta}
                </span>
              </div>
              <div className="text-2xl font-serif font-medium text-text-primary tabular-nums">
                {s.value}
              </div>
              <div className="text-xs text-text-muted mt-0.5">{s.label}</div>
            </div>
          </StaggerItem>
        ))}
      </Stagger>

      <div className="grid lg:grid-cols-3 gap-6">
        <div className="lg:col-span-2 space-y-4">
          <div className="flex items-center justify-between">
            <h2 className="text-base font-serif font-medium text-text-primary">
              {t.app.dashboard.recent}
            </h2>
            <Link
              href="/history"
              className="text-xs text-text-muted hover:text-text-primary inline-flex items-center gap-1"
            >
              {t.app.dashboard.recentAll}
              <ArrowRight className="h-3 w-3" />
            </Link>
          </div>
          {recentRecipes.length > 0 ? (
            <Stagger className="grid sm:grid-cols-2 lg:grid-cols-3 gap-4" staggerChildren={0.05}>
              {recentRecipes.map((r) => (
                <StaggerItem key={r.id}>
                  <RecipeCard recipe={r} />
                </StaggerItem>
              ))}
            </Stagger>
          ) : (
            <div className="text-sm text-text-muted py-10 text-center border border-dashed border-border rounded-2xl">
              No recipes yet. Generate your first one above.
            </div>
          )}
        </div>

        <div className="space-y-6">
          <div className="rounded-2xl border border-border bg-bg-card p-5">
            <div className="flex items-center justify-between mb-3">
              <h3 className="text-sm font-medium text-text-primary font-serif">
                {t.app.dashboard.nutrition}
              </h3>
              <span className="text-[10px] text-text-muted uppercase tracking-[0.1em]">7d</span>
            </div>
            <NutritionLineChart data={CHART_DATA} />
            <div className="mt-3 grid grid-cols-3 gap-2 pt-3 border-t border-border">
              <div>
                <div className="text-[10px] text-text-muted uppercase tracking-[0.1em]">Avg</div>
                <div className="text-sm font-semibold text-text-primary tabular-nums">2,098</div>
              </div>
              <div>
                <div className="text-[10px] text-text-muted uppercase tracking-[0.1em]">Goal</div>
                <div className="text-sm font-semibold text-text-primary tabular-nums">2,200</div>
              </div>
              <div>
                <div className="text-[10px] text-text-muted uppercase tracking-[0.1em]">Δ</div>
                <div className="text-sm font-semibold text-success tabular-nums">-102</div>
              </div>
            </div>
          </div>

          <div className="rounded-2xl border border-border bg-bg-card p-5">
            <div className="flex items-center justify-between mb-3">
              <h3 className="text-sm font-medium text-text-primary font-serif">
                {t.app.dashboard.pantry}
              </h3>
              <span className="text-xs text-text-muted tabular-nums">{pantry.length}</span>
            </div>
            {pantry.length === 0 ? (
              <p className="text-xs text-text-muted">{t.app.dashboard.pantryEmpty}</p>
            ) : (
              <div className="space-y-3">
                {INGREDIENT_CATEGORIES.map((cat) => {
                  const items = pantry.filter((p) => p.category === cat.id).slice(0, 3);
                  if (items.length === 0) return null;
                  return (
                    <div key={cat.id}>
                      <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle mb-1.5">
                        {cat.label}
                      </div>
                      <div className="flex flex-wrap gap-1">
                        {items.map((it) => (
                          <span
                            key={it.id}
                            className="text-xs px-2 py-0.5 rounded-md bg-bg-section border border-border text-text-secondary"
                          >
                            {it.name}
                          </span>
                        ))}
                      </div>
                    </div>
                  );
                })}
              </div>
            )}
          </div>

          {favorites.length > 0 && (
            <div className="rounded-2xl border border-border bg-bg-card p-5">
              <div className="flex items-center justify-between mb-3">
                <h3 className="text-sm font-medium text-text-primary font-serif">Favorites</h3>
                <Link
                  href="/favorites"
                  className="text-[10px] text-text-muted hover:text-text-primary uppercase tracking-[0.1em]"
                >
                  See all
                </Link>
              </div>
              <div className="space-y-1.5">
                {favorites.slice(0, 3).map((id) => {
                  const r = RECIPES.find((x) => x.id === id);
                  if (!r) return null;
                  return (
                    <Link
                      key={id}
                      href={`/recipes/${r.slug}`}
                      className="flex items-center gap-2 p-1.5 -m-1.5 rounded-lg hover:bg-bg-section transition-colors"
                    >
                      <div className="text-xs font-medium text-text-primary truncate">
                        {r.title}
                      </div>
                    </Link>
                  );
                })}
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
