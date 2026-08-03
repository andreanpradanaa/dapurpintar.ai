// Thin typed fetch wrapper around the Go Fiber backend.
// Phase 1 endpoints:
//   POST /api/v1/recipes/generate
//   GET  /api/v1/recipes/:slug
//   GET  /api/v1/health

import type { Recipe, Difficulty, Dietary, RecipeIngredient, RecipeStep, Nutrition } from "./types";

const API_URL =
  process.env.NEXT_PUBLIC_API_URL?.replace(/\/$/, "") || "http://localhost:8080";

export type Language = "en" | "id";

export interface RecipeRef {
  id: string;
  slug: string;
  title: string;
  score?: number;
}

export interface GenerateResponse {
  recipe: Recipe;
  matchScore: number;
  sources: RecipeRef[];
  alternatives: RecipeRef[];
}

export interface HealthResponse {
  status: string;
  db: string;
  llm: string;
  recipes: number;
}

export class ApiError extends Error {
  constructor(
    public readonly status: number,
    public readonly body: {
      error: string;
      message: string;
      requestId?: string;
      fields?: Record<string, string>;
    }
  ) {
    super(body.message);
    this.name = "ApiError";
  }
}

async function request<T>(path: string, init: RequestInit = {}): Promise<T> {
  const url = `${API_URL}${path}`;
  let res: Response;
  try {
    res = await fetch(url, {
      ...init,
      headers: {
        "Content-Type": "application/json",
        ...(init.headers ?? {}),
      },
    });
  } catch (err) {
    // Network-level failure: server down, DNS error, CORS preflight
    // rejected, offline, etc. Convert to ApiError so the caller can
    // handle it uniformly with HTTP errors.
    const message =
      err instanceof Error
        ? `Cannot reach backend at ${API_URL}: ${err.message}`
        : `Cannot reach backend at ${API_URL}`;
    throw new ApiError(0, {
      error: "network_error",
      message,
    });
  }
  if (!res.ok) {
    let body: ApiError["body"] = {
      error: "unknown",
      message: res.statusText,
    };
    try {
      body = await res.json();
    } catch {
      // ignore
    }
    throw new ApiError(res.status, body);
  }
  return (await res.json()) as T;
}

export function generateRecipe(req: {
  ingredients: string[];
  dietary?: Dietary[];
  language?: Language;
}): Promise<GenerateResponse> {
  return request<GenerateResponse>("/api/v1/recipes/generate", {
    method: "POST",
    body: JSON.stringify(req),
  });
}

export function getRecipeBySlug(slug: string): Promise<Recipe> {
  return request<Recipe>(`/api/v1/recipes/${encodeURIComponent(slug)}`);
}

export function getHealth(): Promise<HealthResponse> {
  return request<HealthResponse>("/api/v1/health");
}

export const apiBaseUrl = API_URL;

// Re-export the canonical Recipe type so callers can import it from
// either lib/types or lib/api without divergence.
export type { Recipe, Difficulty, Dietary, RecipeIngredient, RecipeStep, Nutrition };
