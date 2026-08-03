export type Difficulty = "easy" | "medium" | "hard";
export type Dietary = "vegetarian" | "vegan" | "halal" | "gluten-free" | "spicy" | "low-carb";

export type RecipeIngredient = {
  name: string;
  nameId: string;
  amount: string;
  category: "protein" | "vegetable" | "spice" | "grain" | "dairy" | "sauce" | "other";
  optional?: boolean;
};

export type RecipeStep = {
  order: number;
  text: string;
  textId: string;
  durationSec?: number;
  tip?: string;
};

export type Nutrition = {
  calories: number;
  protein: number;
  carbs: number;
  fat: number;
  fiber: number;
};

export type Recipe = {
  id: string;
  slug: string;
  title: string;
  titleId: string;
  description: string;
  descriptionId: string;
  image: string;
  gradient: [string, string];
  cuisine: string;
  difficulty: Difficulty;
  prepTime: number;
  cookTime: number;
  servings: number;
  ingredients: RecipeIngredient[];
  steps: RecipeStep[];
  nutrition: Nutrition;
  tags: string[];
  dietary: Dietary[];
  rating: number;
  reviews: number;
  createdAt: string;
};

export type PantryItem = {
  id: string;
  name: string;
  nameId: string;
  category: RecipeIngredient["category"];
  addedAt: string;
};

export type HistoryEntry = {
  id: string;
  recipeId: string;
  ingredients: string[];
  dietary: Dietary[];
  createdAt: string;
};

export type User = {
  id: string;
  name: string;
  email: string;
  avatar?: string;
  bio?: string;
  plan: "free" | "pro" | "family";
  joinedAt: string;
  streak: number;
  recipesGenerated: number;
  favoritesCount: number;
};
