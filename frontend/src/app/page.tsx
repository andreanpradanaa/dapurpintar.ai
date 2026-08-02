"use client";
import { useEffect, useState, useRef } from "react";
import Link from "next/link";
import { api, type Recipe } from "@/lib/api";

export default function LandingPage() {
  const [query, setQuery] = useState("");
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [loading, setLoading] = useState(false);
  const [chips, setChips] = useState<string[]>([]);
  const debounce = useRef<NodeJS.Timeout | null>(null);

  const searchFromChips = (chipsList: string[]) => {
    const q = chipsList.join(" ");
    if (!q.trim()) { setRecipes([]); return; }
    setLoading(true);
    api.recipes(q).then(r => setRecipes(r.data.slice(0, 6))).catch(() => {}).finally(() => setLoading(false));
  };

  const addChip = (ingredient: string) => {
    const cleaned = ingredient.trim().replace(/,$/, "");
    if (!cleaned) return;
    const newChips = [...chips, cleaned];
    setChips(newChips);
    setQuery("");
    searchFromChips(newChips);
  };

  const removeChip = (index: number) => {
    const newChips = chips.filter((_, i) => i !== index);
    setChips(newChips);
    searchFromChips(newChips);
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter" || e.key === ",") {
      e.preventDefault();
      addChip(query);
    }
    if (e.key === "Backspace" && query === "" && chips.length > 0) {
      removeChip(chips.length - 1);
    }
  };

  useEffect(() => {
    if (debounce.current) clearTimeout(debounce.current);
    if (!query.trim()) return;
    debounce.current = setTimeout(() => {
      api.recipes(query).then(r => setRecipes(r.data.slice(0, 4))).catch(() => {});
    }, 400);
    return () => { if (debounce.current) clearTimeout(debounce.current); };
  }, [query]);

  return (
    <main className="min-h-screen">
      {/* Hero */}
      <section className="relative overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-b from-santan-100/80 to-santan-050 pointer-events-none" />
        <div className="max-w-4xl mx-auto px-6 pt-24 pb-20 md:pt-32 md:pb-28 relative">
          <div className="text-center space-y-6 animate-fadeUp">
            <h1 className="font-display text-5xl md:text-7xl font-bold text-kuali-950 text-balance leading-tight">
              Decide dinner with <span className="italic text-rempah-500">what you have.</span>
            </h1>
            <p className="text-lg md:text-xl text-kuali-700/70 max-w-xl mx-auto text-balance">
              Type what's in your kitchen. We'll suggest recipes you can make right now — no account needed.
            </p>
          </div>

          {/* Ingredient input */}
          <div className="mt-12 max-w-2xl mx-auto">
            <div className="bg-white rounded-2xl border border-bambu-200 shadow-sm shadow-kuali-950/5 p-4 md:p-5 transition-shadow focus-within:shadow-md focus-within:border-rempah-500/30">
              <div className="flex flex-wrap gap-2 mb-3 min-h-[28px]">
                {chips.map((c, i) => (
                  <span key={i} className="inline-flex items-center gap-1 bg-santan-100 border border-bambu-200 text-kuali-700 px-3 py-1.5 rounded-full text-sm font-medium animate-fadeUp">
                    {c}
                    <button onClick={() => removeChip(i)} className="text-bambu-300 hover:text-rempah-500 ml-1 transition-colors" aria-label={`Remove ${c}`}>×</button>
                  </span>
                ))}
              </div>
              <input
                type="text"
                value={query}
                onChange={e => setQuery(e.target.value)}
                onKeyDown={handleKeyDown}
                onBlur={() => query.trim() && addChip(query)}
                placeholder={chips.length === 0 ? "e.g. ayam, bawang putih, santan, tahu" : "Add more ingredients..."}
                className="w-full !border-0 !ring-0 !px-1 !py-1 text-kuali-950 placeholder:text-bambu-300 text-lg focus:!ring-0"
                autoFocus
              />
            </div>

            {/* Quick suggestions */}
            {chips.length === 0 && query.length === 0 && (
              <div className="flex flex-wrap gap-2 mt-4 justify-center">
                {["ayam", "tahu", "tempe", "telur", "santan"].map(s => (
                  <button key={s} onClick={() => addChip(s)} className="text-sm text-bambu-300 hover:text-rempah-500 border border-transparent hover:border-bambu-200 rounded-full px-4 py-1.5 transition-all">
                    {s}
                  </button>
                ))}
              </div>
            )}
          </div>
        </div>
      </section>

      {/* Results */}
      {(recipes.length > 0 || loading) && (
        <section className="max-w-5xl mx-auto px-6 pb-24">
          {loading && (
            <div className="grid md:grid-cols-3 gap-4">
              {Array.from({ length: 3 }).map((_, i) => (
                <div key={i} className="bg-white border border-bambu-200 rounded-xl p-5 space-y-3 animate-pulse-soft">
                  <div className="h-5 bg-santan-200 rounded w-3/4" />
                  <div className="h-4 bg-santan-200 rounded w-1/2" />
                  <div className="h-3 bg-santan-200 rounded w-full" />
                </div>
              ))}
            </div>
          )}
          {!loading && (
            <div className="grid md:grid-cols-3 gap-4">
              {recipes.map((r, i) => (
                <div key={r.id} className="bg-white border border-bambu-200 rounded-xl p-5 hover:shadow-md hover:border-rempah-500/20 transition-all cursor-pointer animate-fadeUp" style={{ animationDelay: `${i * 80}ms`, animationFillMode: "both" }}>
                  <h3 className="font-display text-lg font-bold text-kuali-950">{r.title}</h3>
                  <p className="text-sm text-kuali-700/60 mt-1 line-clamp-2">{r.summary}</p>
                  <div className="flex gap-4 mt-3 text-xs text-bambu-300 font-medium">
                    <span>{r.servings} servings</span>
                    {r.prep_time_minutes && <span>{r.prep_time_minutes}m prep</span>}
                    {r.cook_time_minutes && <span>{r.cook_time_minutes}m cook</span>}
                  </div>
                </div>
              ))}
            </div>
          )}
        </section>
      )}

      {/* Footer nav */}
      {chips.length === 0 && recipes.length === 0 && (
        <section className="max-w-4xl mx-auto px-6 pb-20 text-center space-y-4">
          <p className="text-kuali-700/60 text-sm">
            Already have an account?{" "}
            <Link href="/login" className="text-rempah-500 font-medium hover:text-rempah-700 transition-colors">Log in</Link>
            {" "}or{" "}
            <Link href="/login" className="text-rempah-500 font-medium hover:text-rempah-700 transition-colors">sign up</Link>
          </p>
        </section>
      )}
    </main>
  );
}
