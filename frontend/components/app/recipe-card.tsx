"use client";

import Link from "next/link";
import Image from "next/image";
import { motion } from "framer-motion";
import { Heart, Star, Clock, Users } from "lucide-react";
import { useStore } from "@/lib/store";
import { cn } from "@/lib/utils";
import { photoForRecipe } from "@/lib/photo";
import type { Recipe } from "@/lib/types";

export function RecipeCard({
  recipe,
  variant = "default",
  className,
}: {
  recipe: Recipe;
  variant?: "default" | "compact";
  className?: string;
}) {
  const isFav = useStore((s) => s.favorites.includes(recipe.id));
  const toggleFav = useStore((s) => s.toggleFavorite);

  if (variant === "compact") {
    return (
      <Link
        href={`/recipes/${recipe.slug}`}
        className={cn(
          "group flex items-center gap-3 p-2.5 rounded-xl border border-border bg-bg-card hover:bg-bg-section hover:border-border-strong transition-all duration-200",
          className
        )}
      >
        <div className="relative h-12 w-12 shrink-0 rounded-lg overflow-hidden bg-bg-section">
          <Image
            src={photoForRecipe(recipe.slug)}
            alt={recipe.title}
            fill
            sizes="48px"
            className="object-cover photo-warm"
          />
        </div>
        <div className="min-w-0 flex-1">
          <div className="text-sm font-medium text-text-primary truncate">{recipe.title}</div>
          <div className="text-xs text-text-muted flex items-center gap-2 mt-0.5">
            <span className="flex items-center gap-0.5">
              <Clock className="h-3 w-3" />
              {recipe.prepTime + recipe.cookTime}m
            </span>
            <span className="flex items-center gap-0.5">
              <Star className="h-3 w-3 fill-accent text-accent" />
              {recipe.rating}
            </span>
          </div>
        </div>
        <button
          type="button"
          onClick={(e) => {
            e.preventDefault();
            toggleFav(recipe.id);
          }}
          className="flex h-8 w-8 items-center justify-center rounded-md text-text-muted hover:text-accent hover:bg-bg-section opacity-0 group-hover:opacity-100 transition-opacity"
          aria-label={isFav ? "Unfavorite" : "Favorite"}
        >
          <Heart className="h-4 w-4" fill={isFav ? "#A8553A" : "transparent"} color={isFav ? "#A8553A" : "currentColor"} />
        </button>
      </Link>
    );
  }

  return (
    <motion.div whileHover={{ y: -2 }} transition={{ duration: 0.2 }}>
      <Link
        href={`/recipes/${recipe.slug}`}
        className={cn(
          "group block rounded-2xl border border-border bg-bg-card overflow-hidden transition-all duration-300",
          "hover:border-border-strong hover:shadow-md",
          className
        )}
      >
        <div className="relative aspect-[4/3] overflow-hidden bg-bg-section">
          <Image
            src={photoForRecipe(recipe.slug)}
            alt={recipe.title}
            fill
            sizes="(min-width: 1024px) 33vw, (min-width: 640px) 50vw, 100vw"
            className="object-cover photo-warm transition-transform duration-500 group-hover:scale-105"
          />
          <div className="absolute top-3 left-3 right-3 flex items-center justify-between">
            <span className="inline-flex items-center gap-1 rounded-md bg-bg-card/90 backdrop-blur-sm border border-border text-text-primary text-[10px] font-medium px-2 py-0.5">
              {recipe.cuisine}
            </span>
            <button
              type="button"
              onClick={(e) => {
                e.preventDefault();
                toggleFav(recipe.id);
              }}
              className="flex h-8 w-8 items-center justify-center rounded-md bg-bg-card/90 backdrop-blur-sm text-text-primary hover:bg-bg-card border border-border transition-colors"
              aria-label={isFav ? "Unfavorite" : "Favorite"}
            >
              <Heart
                className="h-4 w-4"
                fill={isFav ? "#A8553A" : "transparent"}
                color={isFav ? "#A8553A" : "currentColor"}
              />
            </button>
          </div>
        </div>
        <div className="p-5">
          <h3 className="text-base font-serif font-medium text-text-primary tracking-tight">
            {recipe.title}
          </h3>
          <p className="mt-1.5 text-sm text-text-muted line-clamp-2 leading-[1.55]">
            {recipe.description}
          </p>
          <div className="mt-4 flex items-center justify-between text-xs text-text-muted">
            <span className="flex items-center gap-1">
              <Users className="h-3 w-3" />
              {recipe.servings} servings
            </span>
            <span className="flex items-center gap-1.5">
              <Clock className="h-3 w-3" />
              {recipe.prepTime + recipe.cookTime}m
            </span>
            <span className="flex items-center gap-1 text-text-secondary">
              <span className="h-1.5 w-1.5 rounded-full bg-accent" />
              {recipe.nutrition.calories} cal
            </span>
          </div>
        </div>
      </Link>
    </motion.div>
  );
}
