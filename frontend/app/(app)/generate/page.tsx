"use client";

import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { useSearchParams } from "next/navigation";
import { Sparkles, RefreshCw, Heart, Share2, Clock, Users, Flame, ChefHat } from "lucide-react";
import { Button } from "@/components/ui/button";
import { CircularProgress } from "@/components/ui/progress";
import { IngredientInput } from "@/components/app/ingredient-input";
import { GenerationLoader } from "@/components/app/generation-loader";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { generateRecipe } from "@/lib/generate";
import { ApiError } from "@/lib/api";
import { toast } from "sonner";
import type { Recipe, Dietary } from "@/lib/types";
import { photoForRecipe } from "@/lib/photo";
import Image from "next/image";

const DIETARY_OPTIONS: { value: Dietary; labelKey: keyof typeof import("@/lib/i18n").translations.en.app.generate.options }[] = [
  { value: "vegetarian", labelKey: "vegetarian" },
  { value: "vegan", labelKey: "vegan" },
  { value: "halal", labelKey: "halal" },
  { value: "gluten-free", labelKey: "gluten-free" },
  { value: "spicy", labelKey: "spicy" },
  { value: "low-carb", labelKey: "low-carb" },
];

export default function GeneratePage() {
  const { t } = useLanguage();
  const params = useSearchParams();

  const initialIngredients = (params.get("ingredients") ?? "")
    .split(",")
    .map((s) => s.trim())
    .filter(Boolean);

  const [ingredients, setIngredients] = useState<string[]>(initialIngredients);
  const [dietary, setDietary] = useState<Dietary[]>([]);
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<Recipe | null>(null);

  const addHistory = useStore((s) => s.addHistory);
  const isFav = useStore((s) => (result ? s.favorites.includes(result.id) : false));
  const toggleFav = useStore((s) => s.toggleFavorite);

  const onGenerate = async () => {
    if (ingredients.length === 0) {
      toast.error("Add at least one ingredient first");
      return;
    }
    setLoading(true);
    setResult(null);

    try {
      const recipe = await generateRecipe(ingredients, dietary);
      setResult(recipe);
      addHistory({
        recipeId: recipe.id,
        ingredients,
        dietary,
      });
      toast.success("Here's your recipe.");
    } catch (err) {
      // Surface a clear, actionable message to the user
      const message =
        err instanceof ApiError
          ? err.body.message
          : err instanceof Error
          ? err.message
          : "Something went wrong. Please try again.";
      toast.error(message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-8 max-w-4xl mx-auto">
      <div>
        <h1 className="font-display text-3xl sm:text-4xl font-medium tracking-tight text-text-primary">
          {t.app.generate.title}
        </h1>
        <p className="mt-1.5 text-sm text-text-muted">{t.app.generate.subtitle}</p>
      </div>

      <div className="rounded-2xl border border-border bg-bg-card p-5 lg:p-6 space-y-5">
        <div>
          <label className="text-[10px] font-semibold uppercase tracking-[0.12em] text-text-muted mb-2 block">
            Your ingredients
          </label>
          <IngredientInput value={ingredients} onChange={setIngredients} />
        </div>

        <div>
          <label className="text-[10px] font-semibold uppercase tracking-[0.12em] text-text-muted mb-2 block">
            {t.app.generate.dietary}
          </label>
          <div className="flex flex-wrap gap-2">
            {DIETARY_OPTIONS.map((opt) => {
              const active = dietary.includes(opt.value);
              return (
                <button
                  key={opt.value}
                  type="button"
                  onClick={() => {
                    setDietary((d) =>
                      active ? d.filter((v) => v !== opt.value) : [...d, opt.value]
                    );
                  }}
                  className={
                    "px-3 py-1.5 rounded-full text-sm font-medium border transition-colors duration-200 " +
                    (active
                      ? "bg-accent text-accent-fg border-accent"
                      : "bg-bg-card text-text-secondary border-border hover:border-border-strong hover:text-text-primary")
                  }
                >
                  {t.app.generate.options[opt.labelKey]}
                </button>
              );
            })}
          </div>
        </div>

        <div className="flex flex-wrap items-center gap-3 pt-2 border-t border-border">
          <div className="text-xs text-text-muted">
            {ingredients.length} ingredient{ingredients.length !== 1 ? "s" : ""} · {dietary.length} preferences
          </div>
          <div className="ml-auto flex items-center gap-2">
            {result && !loading && (
              <Button variant="ghost" onClick={onGenerate} className="group">
                <RefreshCw className="h-3.5 w-3.5 group-hover:rotate-180 transition-transform duration-500" />
                {t.app.generate.regenerate}
              </Button>
            )}
            <Button onClick={onGenerate} loading={loading} className="group" size="md">
              <Sparkles className="h-4 w-4" />
              {t.app.generate.generate}
            </Button>
          </div>
        </div>
      </div>

      <AnimatePresence mode="wait">
        {loading && (
          <motion.div key="loading" exit={{ opacity: 0, y: -8 }} transition={{ duration: 0.24 }}>
            <GenerationLoader active={loading} done={false} />
          </motion.div>
        )}
        {result && !loading && (
          <motion.div
            key="result"
            initial={{ opacity: 0, y: 16 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5, ease: [0.22, 1, 0.36, 1] }}
          >
            <RecipeResult recipe={result} isFav={isFav} toggleFav={toggleFav} />
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}

function RecipeResult({
  recipe,
  isFav,
  toggleFav,
}: {
  recipe: Recipe;
  isFav: boolean;
  toggleFav: (id: string) => void;
}) {
  return (
    <div className="space-y-5">
      <div className="rounded-2xl border border-border bg-bg-card overflow-hidden">
        <div className="relative h-48 sm:h-64 bg-bg-section">
          <Image
            src={photoForRecipe(recipe.slug)}
            alt={recipe.title}
            fill
            sizes="(min-width: 1024px) 70vw, 100vw"
            className="object-cover photo-warm"
          />
          <div className="absolute top-4 left-4 right-4 flex items-center justify-between">
            <span className="inline-flex items-center rounded-md bg-bg-card/90 backdrop-blur-sm border border-border text-text-primary text-[10px] font-medium px-2 py-0.5">
              {recipe.cuisine}
            </span>
            <div className="flex items-center gap-2">
              <button
                type="button"
                onClick={() => toggleFav(recipe.id)}
                className="flex h-9 w-9 items-center justify-center rounded-md bg-bg-card/90 backdrop-blur-sm text-text-primary hover:bg-bg-card border border-border transition-colors"
                aria-label={isFav ? "Unfavorite" : "Favorite"}
              >
                <Heart
                  className="h-4 w-4"
                  fill={isFav ? "#A8553A" : "transparent"}
                  color={isFav ? "#A8553A" : "currentColor"}
                />
              </button>
              <button
                type="button"
                className="flex h-9 w-9 items-center justify-center rounded-md bg-bg-card/90 backdrop-blur-sm text-text-primary hover:bg-bg-card border border-border transition-colors"
                aria-label="Share"
              >
                <Share2 className="h-4 w-4" />
              </button>
            </div>
          </div>
          <div className="absolute inset-x-0 bottom-0 h-32 bg-gradient-to-t from-black/55 to-transparent pointer-events-none" />
          <div className="absolute bottom-4 left-4 right-4">
            <h2 className="text-2xl sm:text-3xl font-serif font-medium text-text-primary">
              {recipe.title}
            </h2>
            <p className="mt-1 text-sm text-text-primary/85 line-clamp-2 max-w-2xl">
              {recipe.description}
            </p>
          </div>
        </div>

        <div className="grid grid-cols-2 sm:grid-cols-4 border-b border-border divide-x divide-border">
          <Meta icon={<Clock className="h-3.5 w-3.5" />} label="Prep + Cook" value={`${recipe.prepTime + recipe.cookTime}m`} />
          <Meta icon={<Users className="h-3.5 w-3.5" />} label="Servings" value={String(recipe.servings)} />
          <Meta icon={<Flame className="h-3.5 w-3.5" />} label="Difficulty" value={recipe.difficulty} />
          <Meta icon={<ChefHat className="h-3.5 w-3.5" />} label="Calories" value={`${recipe.nutrition.calories}`} />
        </div>

        <div className="p-5 border-b border-border">
          <div className="text-[10px] font-semibold uppercase tracking-[0.12em] text-text-muted mb-4">
            Nutrition per serving
          </div>
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
            {[
              { label: "Calories", value: recipe.nutrition.calories, unit: "kcal", percent: 28 },
              { label: "Protein", value: recipe.nutrition.protein, unit: "g", percent: 56 },
              { label: "Carbs", value: recipe.nutrition.carbs, unit: "g", percent: 24 },
              { label: "Fat", value: recipe.nutrition.fat, unit: "g", percent: 23 },
            ].map((n) => (
              <div key={n.label} className="flex flex-col items-center gap-2 p-3 rounded-xl border border-border bg-bg-card">
                <CircularProgress value={n.percent} size={56} strokeWidth={4} />
                <div className="text-center">
                  <div className="text-sm font-semibold text-text-primary tabular-nums">
                    {n.value}
                    <span className="text-xs text-text-muted ml-0.5">{n.unit}</span>
                  </div>
                  <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">{n.label}</div>
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="grid md:grid-cols-2">
          <div className="p-5 md:border-r border-border">
            <div className="text-[10px] font-semibold uppercase tracking-[0.12em] text-text-muted mb-3">
              Ingredients
            </div>
            <ul className="space-y-2">
              {recipe.ingredients.map((ing, i) => (
                <li key={i} className="flex items-center justify-between gap-2 text-sm py-1">
                  <span className="flex items-center gap-2 text-text-primary">
                    <span className="h-1.5 w-1.5 rounded-full bg-accent shrink-0" />
                    {ing.name}
                  </span>
                  <span className="text-xs text-text-muted tabular-nums shrink-0">{ing.amount}</span>
                </li>
              ))}
            </ul>
          </div>
          <div className="p-5">
            <div className="text-[10px] font-semibold uppercase tracking-[0.12em] text-text-muted mb-3">
              Steps
            </div>
            <ol className="space-y-3">
              {recipe.steps.map((step) => (
                <li key={step.order} className="flex gap-3 text-sm leading-[1.6]">
                  <span className="flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-bg-section text-text-secondary text-xs font-semibold tabular-nums">
                    {step.order}
                  </span>
                  <p className="text-text-secondary pt-0.5">{step.text}</p>
                </li>
              ))}
            </ol>
          </div>
        </div>

        <div className="p-5 border-t border-border flex flex-wrap items-center gap-2 bg-bg-section/40">
          <a href={`/recipes/${recipe.slug}`}>
            <Button>Open full recipe</Button>
          </a>
          <Button variant="ghost" onClick={() => toggleFav(recipe.id)}>
            <Heart
              className="h-4 w-4"
              fill={isFav ? "#A8553A" : "transparent"}
              color={isFav ? "#A8553A" : "currentColor"}
            />
            {isFav ? "Saved" : "Save"}
          </Button>
          <Button variant="ghost">
            <Share2 className="h-4 w-4" />
            Share
          </Button>
        </div>
      </div>
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
  value: string;
}) {
  return (
    <div className="flex flex-col items-center gap-1 p-4">
      <div className="text-text-muted">{icon}</div>
      <div className="text-sm font-semibold text-text-primary capitalize tabular-nums">{value}</div>
      <div className="text-[10px] uppercase tracking-[0.1em] text-text-subtle">{label}</div>
    </div>
  );
}
