"use client";

import { useState } from "react";
import Link from "next/link";
import { notFound } from "next/navigation";
import Image from "next/image";
import { motion } from "framer-motion";
import {
  Heart,
  Share2,
  Clock,
  Users,
  Flame,
  ChefHat,
  Star,
  Lightbulb,
  ChevronRight,
  Bookmark,
  Twitter,
  Link as LinkIcon,
  Check,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { CircularProgress } from "@/components/ui/progress";
import { StepList } from "@/components/app/step-list";
import { RecipeCard } from "@/components/app/recipe-card";
import { Stagger, StaggerItem } from "@/components/motion/reveal";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { getRecipeById, getRelatedRecipes } from "@/lib/mock-data/recipes";
import { photoForRecipe } from "@/lib/photo";
import { cn, formatDuration } from "@/lib/utils";
import { toast } from "sonner";

export function RecipeDetail({ slug }: { slug: string }) {
  const recipe = getRecipeById(slug);
  if (!recipe) notFound();

  const { t } = useLanguage();
  const isFav = useStore((s) => s.favorites.includes(recipe.id));
  const toggleFav = useStore((s) => s.toggleFavorite);
  const [checked, setChecked] = useState<Record<number, boolean>>({});
  const [shareOpen, setShareOpen] = useState(false);
  const [servings, setServings] = useState(recipe.servings);

  const related = getRelatedRecipes(recipe.id, 3);

  const checkedCount = Object.values(checked).filter(Boolean).length;
  const totalIngredients = recipe.ingredients.length;

  const handleShare = (platform?: string) => {
    if (platform === "twitter") {
      window.open(
        `https://twitter.com/intent/tweet?text=${encodeURIComponent(recipe.title)}&url=${encodeURIComponent(typeof window !== "undefined" ? window.location.href : "")}`,
        "_blank"
      );
    } else if (platform === "copy") {
      navigator.clipboard.writeText(typeof window !== "undefined" ? window.location.href : "");
      toast.success("Link copied to clipboard");
    }
    setShareOpen(false);
  };

  return (
    <div className="max-w-5xl mx-auto space-y-10">
      {/* Breadcrumb */}
      <nav className="flex items-center gap-1.5 text-xs text-text-muted">
        <Link href="/dashboard" className="hover:text-text-primary">Dashboard</Link>
        <ChevronRight className="h-3 w-3" />
        <Link href="/generate" className="hover:text-text-primary">Create</Link>
        <ChevronRight className="h-3 w-3" />
        <span className="text-text-secondary truncate">{recipe.title}</span>
      </nav>

      {/* Hero */}
      <div className="rounded-2xl border border-border bg-bg-card overflow-hidden">
        <div className="relative aspect-[16/7] sm:aspect-[16/6] bg-bg-section">
          <Image
            src={photoForRecipe(recipe.slug)}
            alt={recipe.title}
            fill
            priority
            sizes="(min-width: 1024px) 80vw, 100vw"
            className="object-cover photo-warm"
          />
          {/* Subtle bottom gradient for text legibility */}
          <div className="absolute inset-0 bg-gradient-to-t from-black/50 via-black/0 to-transparent" />
          <div className="absolute top-5 left-5 right-5 flex items-center justify-between">
            <div className="flex items-center gap-2">
              <span className="inline-flex items-center rounded-md bg-bg-card/90 backdrop-blur-sm border border-border text-text-primary text-[10px] font-medium px-2 py-0.5">
                {recipe.cuisine}
              </span>
              {recipe.dietary.slice(0, 2).map((d) => (
                <Badge key={d} variant="default" className="bg-bg-card/90 backdrop-blur-sm border-border text-text-primary capitalize">
                  {d}
                </Badge>
              ))}
            </div>
            <div className="flex items-center gap-2">
              <button
                type="button"
                onClick={() => toggleFav(recipe.id)}
                className="flex h-10 w-10 items-center justify-center rounded-md bg-bg-card/90 backdrop-blur-sm text-text-primary hover:bg-bg-card border border-border transition-colors"
                aria-label={isFav ? "Unfavorite" : "Favorite"}
              >
                <Heart
                  className="h-4 w-4"
                  fill={isFav ? "#A8553A" : "transparent"}
                  color={isFav ? "#A8553A" : "currentColor"}
                />
              </button>
              <div className="relative">
                <button
                  type="button"
                  onClick={() => setShareOpen((o) => !o)}
                  className="flex h-10 w-10 items-center justify-center rounded-md bg-bg-card/90 backdrop-blur-sm text-text-primary hover:bg-bg-card border border-border transition-colors"
                  aria-label="Share"
                >
                  <Share2 className="h-4 w-4" />
                </button>
                {shareOpen && (
                  <motion.div
                    initial={{ opacity: 0, y: 4 }}
                    animate={{ opacity: 1, y: 0 }}
                    className="absolute right-0 top-12 z-10 w-48 rounded-xl border border-border bg-bg-card p-1 shadow-lg"
                  >
                    <button
                      onClick={() => handleShare("twitter")}
                      className="w-full flex items-center gap-2 px-3 py-2 text-sm text-text-secondary hover:text-text-primary hover:bg-bg-section rounded-md text-left"
                    >
                      <Twitter className="h-4 w-4" /> Share on X
                    </button>
                    <button
                      onClick={() => handleShare("copy")}
                      className="w-full flex items-center gap-2 px-3 py-2 text-sm text-text-secondary hover:text-text-primary hover:bg-bg-section rounded-md text-left"
                    >
                      <LinkIcon className="h-4 w-4" /> Copy link
                    </button>
                  </motion.div>
                )}
              </div>
            </div>
          </div>
          <div className="absolute bottom-5 left-5 right-5">
            <h1 className="text-3xl sm:text-5xl font-serif font-medium text-white tracking-tight text-balance drop-shadow-sm">
              {recipe.title}
            </h1>
            <p className="mt-2 text-sm sm:text-base text-white/90 max-w-2xl leading-[1.55] drop-shadow-sm">
              {recipe.description}
            </p>
            <div className="mt-3 flex items-center gap-3 text-xs text-white/90 drop-shadow-sm">
              <span className="flex items-center gap-1">
                <Star className="h-3 w-3 fill-accent text-accent" />
                {recipe.rating} · {recipe.reviews} cooks
              </span>
            </div>
          </div>
        </div>

        {/* Meta strip */}
        <div className="grid grid-cols-2 sm:grid-cols-5 border-t border-border divide-x divide-border">
          <Meta icon={<Clock className="h-3.5 w-3.5" />} label="Prep" value={`${recipe.prepTime}m`} />
          <Meta icon={<Flame className="h-3.5 w-3.5" />} label="Cook" value={`${recipe.cookTime}m`} />
          <Meta
            icon={<Users className="h-3.5 w-3.5" />}
            label="Servings"
            value={
              <div className="flex items-center gap-1.5">
                <button
                  onClick={() => setServings((s) => Math.max(1, s - 1))}
                  className="h-5 w-5 rounded bg-bg-section text-text-muted hover:text-text-primary"
                  aria-label="Decrease servings"
                >
                  -
                </button>
                <span className="tabular-nums">{servings}</span>
                <button
                  onClick={() => setServings((s) => s + 1)}
                  className="h-5 w-5 rounded bg-bg-section text-text-muted hover:text-text-primary"
                  aria-label="Increase servings"
                >
                  +
                </button>
              </div>
            }
          />
          <Meta icon={<ChefHat className="h-3.5 w-3.5" />} label="Level" value={recipe.difficulty} />
          <Meta
            icon={<ChefHat className="h-3.5 w-3.5" />}
            label="Total"
            value={formatDuration(recipe.prepTime + recipe.cookTime)}
          />
        </div>
      </div>

      <div className="grid lg:grid-cols-3 gap-6">
        {/* Ingredients */}
        <div className="lg:col-span-1">
          <div className="rounded-2xl border border-border bg-bg-card p-5 lg:sticky lg:top-20">
            <div className="flex items-center justify-between mb-3">
              <h2 className="text-xs font-semibold uppercase tracking-[0.12em] text-text-secondary">
                Ingredients
              </h2>
              <span className="text-xs text-text-muted tabular-nums">
                {checkedCount}/{totalIngredients}
              </span>
            </div>
            <div className="h-1 rounded-full bg-bg-section overflow-hidden mb-4">
              <motion.div
                className="h-full bg-accent"
                initial={{ width: 0 }}
                animate={{ width: `${(checkedCount / totalIngredients) * 100}%` }}
                transition={{ duration: 0.3 }}
              />
            </div>
            <ul className="space-y-1.5">
              {recipe.ingredients.map((ing, i) => {
                const isChecked = !!checked[i];
                return (
                  <li key={i}>
                    <button
                      type="button"
                      onClick={() => setChecked((c) => ({ ...c, [i]: !c[i] }))}
                      className={cn(
                        "w-full flex items-center gap-2.5 p-2 rounded-lg border transition-colors text-left",
                        isChecked
                          ? "border-border bg-bg-section"
                          : "border-border bg-bg-card hover:border-border-strong"
                      )}
                    >
                      <span
                        className={cn(
                          "flex h-4 w-4 shrink-0 items-center justify-center rounded border transition-colors",
                          isChecked ? "bg-accent border-accent" : "border-border-strong"
                        )}
                      >
                        {isChecked && <Check className="h-3 w-3 text-accent-fg" strokeWidth={3} />}
                      </span>
                      <div className="flex-1 min-w-0">
                        <div
                          className={cn(
                            "text-sm",
                            isChecked ? "line-through text-text-subtle" : "text-text-primary"
                          )}
                        >
                          {ing.name}
                        </div>
                      </div>
                      <span className="text-xs text-text-muted tabular-nums shrink-0">
                        {ing.amount}
                      </span>
                    </button>
                  </li>
                );
              })}
            </ul>
          </div>
        </div>

        {/* Steps + Nutrition */}
        <div className="lg:col-span-2 space-y-6">
          <div>
            <h2 className="text-xs font-semibold uppercase tracking-[0.12em] text-text-secondary mb-3">
              Steps
            </h2>
            <StepList steps={recipe.steps} />
          </div>

          <div className="rounded-2xl border border-border bg-bg-card p-5">
            <h2 className="text-xs font-semibold uppercase tracking-[0.12em] text-text-secondary mb-4">
              Nutrition per serving
            </h2>
            <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
              {[
                { label: "Calories", value: recipe.nutrition.calories, unit: "kcal", percent: 28 },
                { label: "Protein", value: recipe.nutrition.protein, unit: "g", percent: 56 },
                { label: "Carbs", value: recipe.nutrition.carbs, unit: "g", percent: 24 },
                { label: "Fat", value: recipe.nutrition.fat, unit: "g", percent: 23 },
              ].map((n) => (
                <div key={n.label} className="flex flex-col items-center gap-2 p-3 rounded-xl border border-border bg-bg-card">
                  <CircularProgress value={n.percent} size={64} strokeWidth={5} />
                  <div className="text-center">
                    <div className="text-base font-semibold text-text-primary tabular-nums">
                      {n.value}
                      <span className="text-xs text-text-muted ml-0.5">{n.unit}</span>
                    </div>
                    <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">
                      {n.label}
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          {recipe.steps.some((s) => s.tip) && (
            <div className="rounded-2xl border border-border bg-bg-section p-5">
              <div className="flex items-center gap-2 mb-2">
                <Lightbulb className="h-4 w-4 text-warning" />
                <h3 className="text-sm font-medium text-text-primary">{t.app.recipe.tipsTitle}</h3>
              </div>
              <ul className="space-y-2 text-sm text-text-secondary">
                {recipe.steps
                  .filter((s) => s.tip)
                  .map((s) => (
                    <li key={s.order} className="flex gap-2">
                      <span className="text-text-muted">·</span>
                      <span>{s.tip}</span>
                    </li>
                  ))}
              </ul>
            </div>
          )}

          <div className="flex flex-wrap items-center gap-2">
            <Button>
              <Bookmark className="h-4 w-4" />
              {t.app.recipe.start}
            </Button>
            <Button
              variant={isFav ? "secondary" : "outline"}
              onClick={() => toggleFav(recipe.id)}
            >
              <Heart
                className="h-4 w-4"
                fill={isFav ? "#A8553A" : "transparent"}
                color={isFav ? "#A8553A" : "currentColor"}
              />
              {isFav ? t.app.recipe.saved : t.app.recipe.save}
            </Button>
            <Button variant="ghost">
              <Share2 className="h-4 w-4" />
              {t.app.recipe.share}
            </Button>
          </div>
        </div>
      </div>

      {related.length > 0 && (
        <section className="pt-10 border-t border-border">
          <h2 className="text-base font-serif font-medium text-text-primary mb-4">
            {t.app.recipe.related}
          </h2>
          <Stagger className="grid sm:grid-cols-2 lg:grid-cols-3 gap-4" staggerChildren={0.06}>
            {related.map((r) => (
              <StaggerItem key={r.id}>
                <RecipeCard recipe={r} />
              </StaggerItem>
            ))}
          </Stagger>
        </section>
      )}
    </div>
  );
}

function Meta({
  icon,
  label,
  value,
}: {
  icon: React.ReactNode;
  label: string;
  value: React.ReactNode;
}) {
  return (
    <div className="flex flex-col items-center gap-1 p-4">
      <div className="text-text-muted">{icon}</div>
      <div className="text-sm font-semibold text-text-primary capitalize">{value}</div>
      <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">{label}</div>
    </div>
  );
}
