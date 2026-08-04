"use client";

import { generateRecipe as apiGenerate, type Recipe } from "./api";

/**
 * Phase 1: thin wrapper around the Go Fiber backend.
 * The LLM-composed recipe comes back wrapped in { recipe, matchScore, sources, alternatives }.
 * The previous in-memory scoring fallback has been removed — the backend is the
 * single source of truth.
 *
 * The mock dataset in lib/mock-data/recipes.ts is still used for the recipe
 * detail and history pages until those phases ship.
 */
export async function generateRecipe(
  ingredients: string[],
  dietary: string[] = []
): Promise<Recipe> {
  const { recipe } = await apiGenerate({
    ingredients,
    dietary: dietary as never[],
    language: "en",
  });
  return recipe;
}
