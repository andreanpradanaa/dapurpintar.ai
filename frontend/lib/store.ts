"use client";

import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import type { HistoryEntry, PantryItem, User, Dietary } from "@/lib/types";

type StoreState = {
  // Auth (mocked)
  isAuthenticated: boolean;
  user: User | null;
  signIn: (user?: Partial<User>) => void;
  signOut: () => void;
  updateUser: (patch: Partial<User>) => void;

  // Favorites
  favorites: string[];
  toggleFavorite: (recipeId: string) => void;
  isFavorite: (recipeId: string) => boolean;

  // History
  history: HistoryEntry[];
  addHistory: (entry: Omit<HistoryEntry, "id" | "createdAt">) => void;
  removeHistory: (id: string) => void;
  clearHistory: () => void;

  // Pantry
  pantry: PantryItem[];
  addPantryItem: (item: Omit<PantryItem, "id" | "addedAt">) => void;
  removePantryItem: (id: string) => void;
  clearPantry: () => void;

  // Preferences
  preferences: {
    defaultServings: number;
    units: "metric" | "imperial";
    dietaryDefaults: Dietary[];
    notifications: {
      weekly: boolean;
      newFeature: boolean;
      marketing: boolean;
    };
  };
  setPreference: <K extends keyof StoreState["preferences"]>(
    key: K,
    value: StoreState["preferences"][K]
  ) => void;
  setNotification: (key: keyof StoreState["preferences"]["notifications"], value: boolean) => void;
};

const DEFAULT_USER: User = {
  id: "u-001",
  name: "Gusti Adistriani",
  email: "gusti@dapurpintar.ai",
  bio: "Home cook exploring Indonesian cuisine with AI. 🍳",
  plan: "pro",
  joinedAt: "2025-08-12",
  streak: 23,
  recipesGenerated: 147,
  favoritesCount: 0,
};

export const useStore = create<StoreState>()(
  persist(
    (set, get) => ({
      isAuthenticated: true, // dev: start authenticated so dashboard is browsable
      user: DEFAULT_USER,
      signIn: (patch) =>
        set({
          isAuthenticated: true,
          user: { ...(get().user ?? DEFAULT_USER), ...(patch ?? {}) },
        }),
      signOut: () => set({ isAuthenticated: false }),
      updateUser: (patch) =>
        set((s) => (s.user ? { user: { ...s.user, ...patch } } : {})),

      favorites: ["r-001", "r-003", "r-007", "r-018"],
      toggleFavorite: (id) =>
        set((s) => {
          const exists = s.favorites.includes(id);
          return {
            favorites: exists ? s.favorites.filter((f) => f !== id) : [id, ...s.favorites],
          };
        }),
      isFavorite: (id) => get().favorites.includes(id),

      history: [
        {
          id: "h-001",
          recipeId: "r-001",
          ingredients: ["chicken", "rice", "shallot", "garlic", "egg", "kecap manis"],
          dietary: ["halal"],
          createdAt: "2026-07-28T08:14:00Z",
        },
        {
          id: "h-002",
          recipeId: "r-006",
          ingredients: ["chicken thigh", "honey", "kecap", "garlic", "lime"],
          dietary: ["halal"],
          createdAt: "2026-07-26T19:42:00Z",
        },
        {
          id: "h-003",
          recipeId: "r-021",
          ingredients: ["flour", "egg", "milk", "sugar", "cheese"],
          dietary: ["vegetarian"],
          createdAt: "2026-07-25T15:08:00Z",
        },
        {
          id: "h-004",
          recipeId: "r-010",
          ingredients: ["tempeh", "kecap", "garlic", "chili"],
          dietary: ["vegan", "halal"],
          createdAt: "2026-07-24T11:30:00Z",
        },
        {
          id: "h-005",
          recipeId: "r-005",
          ingredients: ["egg noodle", "egg", "cabbage", "kecap", "shallot"],
          dietary: ["halal"],
          createdAt: "2026-07-22T20:15:00Z",
        },
      ],
      addHistory: (entry) =>
        set((s) => ({
          history: [
            { ...entry, id: `h-${Date.now()}`, createdAt: new Date().toISOString() },
            ...s.history,
          ].slice(0, 100),
        })),
      removeHistory: (id) => set((s) => ({ history: s.history.filter((h) => h.id !== id) })),
      clearHistory: () => set({ history: [] }),

      pantry: [
        { id: "p-001", name: "Chicken", nameId: "Ayam", category: "protein", addedAt: "2026-07-30" },
        { id: "p-002", name: "Rice", nameId: "Nasi", category: "grain", addedAt: "2026-07-30" },
        { id: "p-003", name: "Garlic", nameId: "Bawang putih", category: "vegetable", addedAt: "2026-07-29" },
        { id: "p-004", name: "Shallot", nameId: "Bawang merah", category: "vegetable", addedAt: "2026-07-29" },
        { id: "p-005", name: "Egg", nameId: "Telur", category: "protein", addedAt: "2026-07-28" },
        { id: "p-006", name: "Sweet soy sauce", nameId: "Kecap manis", category: "sauce", addedAt: "2026-07-28" },
        { id: "p-007", name: "Chili", nameId: "Cabai", category: "spice", addedAt: "2026-07-27" },
        { id: "p-008", name: "Coconut milk", nameId: "Santan", category: "dairy", addedAt: "2026-07-27" },
        { id: "p-009", name: "Tempeh", nameId: "Tempe", category: "protein", addedAt: "2026-07-26" },
        { id: "p-010", name: "Cabbage", nameId: "Kubis", category: "vegetable", addedAt: "2026-07-25" },
      ],
      addPantryItem: (item) =>
        set((s) => ({
          pantry: [
            { ...item, id: `p-${Date.now()}`, addedAt: new Date().toISOString().slice(0, 10) },
            ...s.pantry,
          ],
        })),
      removePantryItem: (id) => set((s) => ({ pantry: s.pantry.filter((p) => p.id !== id) })),
      clearPantry: () => set({ pantry: [] }),

      preferences: {
        defaultServings: 2,
        units: "metric",
        dietaryDefaults: [],
        notifications: {
          weekly: true,
          newFeature: true,
          marketing: false,
        },
      },
      setPreference: (key, value) =>
        set((s) => ({ preferences: { ...s.preferences, [key]: value } })),
      setNotification: (key, value) =>
        set((s) => ({
          preferences: {
            ...s.preferences,
            notifications: { ...s.preferences.notifications, [key]: value },
          },
        })),
    }),
    {
      name: "dapur-pintar-store",
      storage: createJSONStorage(() => localStorage),
      partialize: (state) => ({
        isAuthenticated: state.isAuthenticated,
        user: state.user,
        favorites: state.favorites,
        history: state.history,
        pantry: state.pantry,
        preferences: state.preferences,
      }),
    }
  )
);
