import type { RecipeIngredient } from "@/lib/types";

export type Ingredient = {
  name: string;
  nameId: string;
  emoji: string;
  category: RecipeIngredient["category"];
  common: boolean;
};

export const INGREDIENTS: Ingredient[] = [
  // Protein
  { name: "Chicken breast", nameId: "Dada ayam", emoji: "🍗", category: "protein", common: true },
  { name: "Chicken thigh", nameId: "Paha ayam", emoji: "🍗", category: "protein", common: true },
  { name: "Beef", nameId: "Daging sapi", emoji: "🥩", category: "protein", common: true },
  { name: "Shrimp", nameId: "Udang", emoji: "🦐", category: "protein", common: true },
  { name: "Egg", nameId: "Telur", emoji: "🥚", category: "protein", common: true },
  { name: "Tofu", nameId: "Tahu", emoji: "🟦", category: "protein", common: true },
  { name: "Tempeh", nameId: "Tempe", emoji: "🟫", category: "protein", common: true },
  { name: "Fish fillet", nameId: "Fillet ikan", emoji: "🐟", category: "protein", common: true },
  { name: "Pork belly", nameId: "Babi", emoji: "🥓", category: "protein", common: false },
  { name: "Squid", nameId: "Cumi", emoji: "🦑", category: "protein", common: false },

  // Vegetable
  { name: "Garlic", nameId: "Bawang putih", emoji: "🧄", category: "vegetable", common: true },
  { name: "Shallot", nameId: "Bawang merah", emoji: "🧅", category: "vegetable", common: true },
  { name: "Onion", nameId: "Bawang bombay", emoji: "🧅", category: "vegetable", common: true },
  { name: "Tomato", nameId: "Tomat", emoji: "🍅", category: "vegetable", common: true },
  { name: "Chili", nameId: "Cabai", emoji: "🌶️", category: "spice", common: true },
  { name: "Carrot", nameId: "Wortel", emoji: "🥕", category: "vegetable", common: true },
  { name: "Cabbage", nameId: "Kubis", emoji: "🥬", category: "vegetable", common: true },
  { name: "Spinach", nameId: "Bayam", emoji: "🥬", category: "vegetable", common: true },
  { name: "Bean sprouts", nameId: "Tauge", emoji: "🌱", category: "vegetable", common: true },
  { name: "Long bean", nameId: "Kacang panjang", emoji: "🫛", category: "vegetable", common: true },
  { name: "Cucumber", nameId: "Mentimun", emoji: "🥒", category: "vegetable", common: true },
  { name: "Eggplant", nameId: "Terong", emoji: "🍆", category: "vegetable", common: true },
  { name: "Potato", nameId: "Kentang", emoji: "🥔", category: "vegetable", common: true },
  { name: "Corn", nameId: "Jagung", emoji: "🌽", category: "vegetable", common: true },
  { name: "Lemongrass", nameId: "Serai", emoji: "🌾", category: "spice", common: true },
  { name: "Galangal", nameId: "Lengkuas", emoji: "🫚", category: "spice", common: true },
  { name: "Ginger", nameId: "Jahe", emoji: "🫚", category: "spice", common: true },
  { name: "Turmeric", nameId: "Kunyit", emoji: "🟡", category: "spice", common: true },
  { name: "Kaffir lime", nameId: "Daun jeruk", emoji: "🍃", category: "spice", common: true },
  { name: "Pandan leaf", nameId: "Daun pandan", emoji: "🌿", category: "spice", common: true },
  { name: "Basil", nameId: "Kemangi", emoji: "🌿", category: "spice", common: false },
  { name: "Lime", nameId: "Jeruk nipis", emoji: "🍋", category: "spice", common: true },
  { name: "Coconut", nameId: "Kelapa", emoji: "🥥", category: "vegetable", common: false },

  // Grain
  { name: "Rice", nameId: "Nasi", emoji: "🍚", category: "grain", common: true },
  { name: "Noodle", nameId: "Mie", emoji: "🍜", category: "grain", common: true },
  { name: "Vermicelli", nameId: "Bihun", emoji: "🍝", category: "grain", common: true },
  { name: "Flour", nameId: "Tepung", emoji: "🌾", category: "grain", common: true },
  { name: "Bread", nameId: "Roti", emoji: "🍞", category: "grain", common: true },

  // Dairy
  { name: "Coconut milk", nameId: "Santan", emoji: "🥥", category: "dairy", common: true },
  { name: "Milk", nameId: "Susu", emoji: "🥛", category: "dairy", common: true },
  { name: "Butter", nameId: "Mentega", emoji: "🧈", category: "dairy", common: true },

  // Sauce
  { name: "Soy sauce", nameId: "Kecap manis", emoji: "🍶", category: "sauce", common: true },
  { name: "Salt", nameId: "Garam", emoji: "🧂", category: "sauce", common: true },
  { name: "Sugar", nameId: "Gula", emoji: "🍬", category: "sauce", common: true },
  { name: "Vinegar", nameId: "Cuka", emoji: "🍾", category: "sauce", common: true },
  { name: "Oyster sauce", nameId: "Saus tiram", emoji: "🥫", category: "sauce", common: true },
  { name: "Chili sauce", nameId: "Sambal", emoji: "🌶️", category: "sauce", common: true },
  { name: "Peanut sauce", nameId: "Saus kacang", emoji: "🥜", category: "sauce", common: true },
];

export const INGREDIENT_CATEGORIES: { id: RecipeIngredient["category"]; label: string; labelId: string }[] = [
  { id: "protein", label: "Protein", labelId: "Protein" },
  { id: "vegetable", label: "Vegetables", labelId: "Sayuran" },
  { id: "spice", label: "Spices & Herbs", labelId: "Bumbu & Rempah" },
  { id: "grain", label: "Grains", labelId: "Biji-bijian" },
  { id: "dairy", label: "Dairy", labelId: "Susu" },
  { id: "sauce", label: "Sauces & Condiments", labelId: "Saus & Bumbu" },
  { id: "other", label: "Other", labelId: "Lainnya" },
];
