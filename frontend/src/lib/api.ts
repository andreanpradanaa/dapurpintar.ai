const BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

type RequestOptions = { method?: string; body?: unknown; headers?: Record<string, string> };

async function request<T>(path: string, opts: RequestOptions = {}): Promise<T> {
  const url = `${BASE}${path}`;
  const headers: Record<string, string> = { "Content-Type": "application/json", ...opts.headers };
  const res = await fetch(url, {
    method: opts.method || "GET",
    headers,
    body: opts.body ? JSON.stringify(opts.body) : undefined,
    credentials: "include",
  });
  if (!res.ok) {
    const err = await res.json().catch(() => ({ error: { message: "Network error" } }));
    throw new Error(err?.error?.message || `Request failed (${res.status})`);
  }
  if (res.status === 204) return undefined as T;
  return res.json() as Promise<T>;
}

export type Account = { id: string; email: string; status: string; timezone: string; email_verified: boolean; profile_id?: string; created_at: string };
export type Profile = { id: string; display_name: string; status: string; timezone?: string };
export type PantryItem = { id: string; ingredient_name: string; category: string; quantity: number; unit: string; expiry_date?: string; status: string };
export type PantrySummary = { total_items: number; expiring_soon_count: number; running_low_count: number };
export type Page<T> = { next_cursor?: string; has_more: boolean };
export type Collection<T> = { data: T[]; page: Page<T> };

export type Recipe = { id: string; title: string; summary: string; servings: number; prep_time_minutes?: number; cook_time_minutes?: number; ingredients?: {name:string;quantity:string}[]; instructions?: string[] };
export type MealPlan = { id: string; title: string; period_start: string; period_end: string; status: string; created_at: string };
export type PlannedMeal = { id: string; meal_plan_id: string; meal_date: string; meal_occasion: string; recipe_id?: string; recommendation_option_id?: string; status: string };
export type ShoppingList = { id: string; title: string; status: string; item_counts: { open: number; completed: number }; created_at: string };
export type ShoppingItem = { id: string; shopping_list_id: string; ingredient_name: string; quantity: number; unit: string; source: string; status: string };

export const api = {
  // Auth
  register: (body: { email: string; password: string; display_name: string; timezone: string }) =>
    request<{ data: Account }>("/accounts", { method: "POST", body }),
  login: (body: { email: string; password: string }) =>
    request<{ data: Account }>("/accounts/login", { method: "POST", body }),
  refresh: () => request<{ data: Account }>("/accounts/refresh", { method: "POST" }),
  logout: () => request<void>("/accounts/logout", { method: "POST" }),
  me: () => request<{ data: Account }>("/accounts/me"),

  // Profile
  getProfile: () => request<{ data: Profile }>("/profile"),
  updateProfile: (body: { display_name?: string; timezone?: string }) =>
    request<{ data: Profile }>("/profile", { method: "PATCH", body }),

  // Pantry
  pantrySummary: () => request<{ data: PantrySummary }>("/pantry"),
  pantryItems: (cursor?: string) =>
    request<Collection<PantryItem>>(`/pantry/items?limit=50&cursor=${cursor || ""}`),
  addPantryItem: (body: { ingredient_name: string; category: string; quantity: number; unit?: string; expiry_date?: string }) =>
    request<{ data: PantryItem }>("/pantry/items", { method: "POST", body }),
  expiringItems: (cursor?: string) =>
    request<Collection<PantryItem>>(`/pantry/expiry?limit=20&cursor=${cursor || ""}`),

  // Recipes
  recipes: (q?: string, cursor?: string) =>
    request<Collection<Recipe>>(`/recipes?limit=20${q ? "&q=" + encodeURIComponent(q) : ""}&cursor=${cursor || ""}`),
  recipe: (id: string) => request<{ data: Recipe }>(`/recipes/${id}`),
  favoriteRecipe: (recipeId: string) => request<void>(`/favorites/recipes/${recipeId}`, { method: "PUT" }),
  unfavoriteRecipe: (recipeId: string) => request<void>(`/favorites/recipes/${recipeId}`, { method: "DELETE" }),
  favorites: (cursor?: string) =>
    request<Collection<{ recipe: Recipe; created_at: string }>>(`/favorites?limit=20&cursor=${cursor || ""}`),

  // Meal Plans
  mealPlans: (cursor?: string) =>
    request<Collection<MealPlan>>(`/meal-plans?limit=20&cursor=${cursor || ""}`),
  createMealPlan: (body: { period_start: string; period_end: string; title: string }) =>
    request<{ data: MealPlan }>("/meal-plans", { method: "POST", body }),
  mealPlan: (id: string) => request<{ data: MealPlan }>(`/meal-plans/${id}`),
  plannedMeals: (planId: string, cursor?: string) =>
    request<Collection<PlannedMeal>>(`/meal-plans/${planId}/meals?limit=50&cursor=${cursor || ""}`),
  planMeal: (planId: string, body: { meal_date: string; meal_occasion: string; recipe_id?: string }) =>
    request<{ data: PlannedMeal }>(`/meal-plans/${planId}/meals`, { method: "POST", body }),
  completeMealPlan: (id: string) => request<{ data: MealPlan }>(`/meal-plans/${id}/complete`, { method: "POST" }),

  // Shopping
  shoppingLists: (cursor?: string) =>
    request<Collection<ShoppingList>>(`/shopping-lists?limit=20&cursor=${cursor || ""}`),
  createShoppingList: (body: { title: string }) =>
    request<{ data: ShoppingList }>("/shopping-lists", { method: "POST", body }),
  shoppingList: (id: string) => request<{ data: ShoppingList }>(`/shopping-lists/${id}`),
  shoppingItems: (listId: string, cursor?: string) =>
    request<Collection<ShoppingItem>>(`/shopping-lists/${listId}/items?limit=50&cursor=${cursor || ""}`),
  addShoppingItem: (listId: string, body: { ingredient_name: string; quantity?: number; unit?: string }) =>
    request<{ data: ShoppingItem }>(`/shopping-lists/${listId}/items`, { method: "POST", body }),
  activateShoppingList: (id: string) => request<{ data: ShoppingList }>(`/shopping-lists/${id}/activate`, { method: "POST" }),
  completeShoppingList: (id: string) => request<{ data: ShoppingList }>(`/shopping-lists/${id}/complete`, { method: "POST" }),
  completeShoppingItem: (listId: string, itemId: string) =>
    request<{ data: ShoppingItem }>(`/shopping-lists/${listId}/items/${itemId}/complete`, { method: "POST" }),
};

export async function getAuthenticated<T>(fn: () => Promise<T>): Promise<T> {
  try { return await fn(); } catch {
    try { await api.refresh(); return await fn(); } catch { throw new Error("Authentication required"); }
  }
}
