import { RECIPES } from "@/lib/mock-data/recipes";
import type { Recipe } from "@/lib/types";

/**
 * Simulated AI recipe generator.
 * Given a list of ingredients and dietary preferences, picks a matching recipe from the mock dataset.
 * Includes a streaming-like delay to feel real.
 */
export async function generateRecipe(
  ingredients: string[],
  dietary: string[] = []
): Promise<Recipe> {
  // 1.8-3.2s delay to feel realistic
  const delay = 1800 + Math.random() * 1400;
  await new Promise((resolve) => setTimeout(resolve, delay));

  const lowerIngredients = ingredients.map((i) => i.toLowerCase().trim()).filter(Boolean);

  // Score each recipe by overlap of ingredient keywords
  const scored = RECIPES.map((recipe) => {
    const recipeWords = [
      recipe.title,
      ...recipe.tags,
      ...recipe.ingredients.map((i) => i.name),
    ]
      .join(" ")
      .toLowerCase();

    let score = 0;
    for (const ing of lowerIngredients) {
      if (recipeWords.includes(ing)) score += 3;
      // partial match
      for (const word of ing.split(/\s+/)) {
        if (word.length >= 3 && recipeWords.includes(word)) score += 1;
      }
    }
    // Dietary alignment
    if (dietary.length > 0) {
      const matches = dietary.filter((d) => recipe.dietary.includes(d as never)).length;
      score += matches * 2;
      if (matches === dietary.length) score += 5;
    }
    // Prefer well-rated
    score += recipe.rating * 0.5;

    return { recipe, score };
  });

  scored.sort((a, b) => b.score - a.score);
  const winner = scored[0]?.recipe ?? RECIPES[0];

  return winner;
}

/**
 * Streaming-style phase emitter for the generation animation.
 */
export type GenerationPhase = {
  id: number;
  label: { en: string; id: string };
};

export const GENERATION_PHASES: GenerationPhase[] = [
  {
    id: 1,
    label: { en: "Analyzing your ingredients", id: "Menganalisis bahan Anda" },
  },
  {
    id: 2,
    label: { en: "Matching flavor profiles", id: "Mencocokkan profil rasa" },
  },
  {
    id: 3,
    label: { en: "Composing cooking steps", id: "Menyusun langkah masak" },
  },
  {
    id: 4,
    label: { en: "Plating the result", id: "Menyajikan hasil" },
  },
];
