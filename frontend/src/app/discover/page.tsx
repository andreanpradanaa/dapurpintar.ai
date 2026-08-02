"use client";
import { useEffect, useState } from "react";
import { useAuth } from "@/lib/auth";
import { api, type Recipe } from "@/lib/api";

export default function DiscoverPage() {
  const { account } = useAuth();
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [favorites, setFavorites] = useState<Set<string>>(new Set());
  const [query, setQuery] = useState("");
  const [loading, setLoading] = useState(false);
  const [err, setErr] = useState("");
  const [selected, setSelected] = useState<Recipe | null>(null);
  const [tab, setTab] = useState<"discover" | "favorites">("discover");

  const load = (q?: string) => {
    setLoading(true);
    const doLoad = async () => {
      if (tab === "favorites") {
        const r = await api.favorites();
        setRecipes(r.data.map(f => f.recipe).filter(Boolean));
      } else {
        const r = await api.recipes(q);
        setRecipes(r.data);
      }
    };
    doLoad().catch(() => setErr("Failed to load")).finally(() => setLoading(false));
  };

  const loadFavs = () => {
    api.favorites().then(r => setFavorites(new Set(r.data.map(f => f.recipe?.id).filter(Boolean)))).catch(() => {});
  };

  useEffect(() => { if (account) { load(); loadFavs(); } }, [account, tab]);

  const search = (e: React.FormEvent) => { e.preventDefault(); load(query); };
  const toggleFav = async (id: string) => {
    try {
      if (favorites.has(id)) await api.unfavoriteRecipe(id);
      else await api.favoriteRecipe(id);
      setFavorites(prev => { const s = new Set(prev); s.has(id) ? s.delete(id) : s.add(id); return s; });
    } catch { /* ignore */ }
  };

  if (!account) return null;

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <h1 className="text-xl font-bold text-kuali-950">Discover</h1>
        <div className="flex gap-2">
          <button onClick={() => setTab("discover")} className={`text-sm font-medium px-3 py-1 rounded ${tab === "discover" ? "bg-rempah-500 text-white-000" : "text-kuali-700 hover:bg-bambu-200"}`}>Recipes</button>
          <button onClick={() => setTab("favorites")} className={`text-sm font-medium px-3 py-1 rounded ${tab === "favorites" ? "bg-rempah-500 text-white-000" : "text-kuali-700 hover:bg-bambu-200"}`}>Favorites</button>
        </div>
      </div>

      {tab === "discover" && (
        <form onSubmit={search} className="flex gap-2">
          <input type="text" placeholder="Search recipes..." value={query} onChange={e => setQuery(e.target.value)} className="flex-1" />
          <button type="submit" className="bg-rempah-500 text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-rempah-700 transition-colors">Search</button>
        </form>
      )}

      {loading && <p className="text-bambu-300 text-center py-8">Loading...</p>}
      {err && <p className="text-feedback-error text-sm">{err}</p>}
      {!loading && !err && recipes.length === 0 && (
        <div className="text-center py-12 text-kuali-700 bg-white border border-bambu-200 rounded-xl">
          <p className="text-lg">{tab === "favorites" ? "No favorites yet." : "No recipes found."}</p>
          <p className="text-sm text-bambu-300 mt-1">{tab === "favorites" ? "Favorite recipes to see them here." : "Try a different search."}</p>
        </div>
      )}

      <div className="grid gap-3 md:grid-cols-2">
        {recipes.map(r => (
          <div key={r.id} className="bg-white border border-bambu-200 rounded-lg p-4 hover:border-rempah-500/30 cursor-pointer transition-colors" onClick={() => setSelected(r)}>
            <div className="flex justify-between items-start">
              <div className="flex-1">
                <h3 className="font-semibold text-kuali-950">{r.title}</h3>
                <p className="text-sm text-kuali-700 mt-1 line-clamp-2">{r.summary}</p>
                <div className="flex gap-3 mt-2 text-xs text-bambu-300">
                  <span>{r.servings} servings</span>
                  {r.prep_time_minutes && <span>{r.prep_time_minutes}m prep</span>}
                  {r.cook_time_minutes && <span>{r.cook_time_minutes}m cook</span>}
                </div>
              </div>
              <button onClick={e => { e.stopPropagation(); toggleFav(r.id); }} className={`ml-2 text-lg ${favorites.has(r.id) ? "text-rempah-500" : "text-bambu-300 hover:text-rempah-500"}`}>
                {favorites.has(r.id) ? "♥" : "♡"}
              </button>
            </div>
          </div>
        ))}
      </div>

      {selected && (
        <div className="fixed inset-0 bg-kuali-950/50 flex items-center justify-center p-4 z-50" onClick={() => setSelected(null)}>
          <div className="bg-white rounded-xl max-w-lg w-full max-h-[80vh] overflow-y-auto p-6 space-y-4" onClick={e => e.stopPropagation()}>
            <h2 className="text-xl font-bold text-kuali-950">{selected.title}</h2>
            <p className="text-kuali-700">{selected.summary}</p>
            <div className="flex gap-3 text-sm text-bambu-300">
              <span>{selected.servings} servings</span>
              {selected.prep_time_minutes && <span>{selected.prep_time_minutes}m prep</span>}
              {selected.cook_time_minutes && <span>{selected.cook_time_minutes}m cook</span>}
            </div>
            {selected.ingredients && selected.ingredients.length > 0 && (
              <div>
                <h3 className="font-semibold text-kuali-950 text-sm mb-1">Ingredients</h3>
                <ul className="text-sm text-kuali-700 space-y-1">
                  {selected.ingredients.map((ing, i) => <li key={i}>{ing.quantity} {ing.name}</li>)}
                </ul>
              </div>
            )}
            {selected.instructions && selected.instructions.length > 0 && (
              <div>
                <h3 className="font-semibold text-kuali-950 text-sm mb-1">Instructions</h3>
                <ol className="text-sm text-kuali-700 space-y-1 list-decimal list-inside">
                  {selected.instructions.map((s, i) => <li key={i}>{s}</li>)}
                </ol>
              </div>
            )}
            <button onClick={() => setSelected(null)} className="text-sm text-rempah-500 font-medium">Close</button>
          </div>
        </div>
      )}
    </div>
  );
}
