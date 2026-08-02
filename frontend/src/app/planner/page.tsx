"use client";
import { useEffect, useState, useCallback } from "react";
import { useAuth } from "@/lib/auth";
import { api, getAuthenticated, type MealPlan, type PlannedMeal, type Recipe } from "@/lib/api";
import { useToast } from "@/lib/toast";
import { ListSkeleton, Skeleton } from "@/components/skeleton";

const OCCASIONS = ["breakfast", "lunch", "dinner", "snack"] as const;

function computeDates(periodStart: string): { d: string; label: string }[] {
  const start = new Date(periodStart + "T00:00:00");
  const days: { d: string; label: string }[] = [];
  for (let i = 0; i < 7; i++) {
    const d = new Date(start);
    d.setDate(start.getDate() + i);
    const iso = d.toISOString().slice(0, 10);
    const label = d.toLocaleDateString("en-US", { weekday: "short", month: "short", day: "numeric" });
    days.push({ d: iso, label });
  }
  return days;
}

export default function PlannerPage() {
  const { account } = useAuth();
  const { toast } = useToast();
  const [plans, setPlans] = useState<MealPlan[]>([]);
  const [plan, setPlan] = useState<MealPlan | null>(null);
  const [meals, setMeals] = useState<Record<string, PlannedMeal[]>>({});
  const [loading, setLoading] = useState(false);
  const [create, setCreate] = useState(false);
  const [title, setTitle] = useState("");
  const [start, setStart] = useState("");
  const [end, setEnd] = useState("");
  const [picker, setPicker] = useState<{ date: string; occasion: string } | null>(null);
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [recipeLoading, setRecipeLoading] = useState(false);

  const loadPlans = useCallback(() => {
    getAuthenticated(() => api.mealPlans()).then(r => setPlans(r.data)).catch(() => toast("error", "Failed to load plans"));
  }, [toast]);

  useEffect(() => { if (account) loadPlans(); }, [account, loadPlans]);

  const openPlan = async (p: MealPlan) => {
    setPlan(p);
    setLoading(true);
    try {
      const r = await getAuthenticated(() => api.plannedMeals(p.id));
      const map: Record<string, PlannedMeal[]> = {};
      r.data.forEach(m => {
        const key = `${m.meal_date}|${m.meal_occasion}`;
        if (!map[key]) map[key] = [];
        map[key].push(m);
      });
      setMeals(map);
    } catch { toast("error", "Failed to load meals"); }
    setLoading(false);
  };

  const submitPlan = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const r = await getAuthenticated(() => api.createMealPlan({ period_start: start, period_end: end, title: title || "Meal plan" }));
      setCreate(false);
      toast("success", "Plan created");
      setPlans(prev => [r.data, ...prev]);
      openPlan(r.data);
    } catch { toast("error", "Failed to create plan"); }
  };

  const openPicker = async (date: string, occasion: string) => {
    setPicker({ date, occasion });
    setRecipeLoading(true);
    try {
      const r = await api.recipes();
      setRecipes(r.data);
    } catch { toast("error", "Failed to load recipes"); }
    setRecipeLoading(false);
  };

  const pickRecipe = async (recipe: Recipe) => {
    if (!picker || !plan) return;
    try {
      const r = await getAuthenticated(() => api.planMeal(plan.id, { meal_date: picker.date, meal_occasion: picker.occasion, recipe_id: recipe.id }));
      toast("success", `${recipe.title} planned`);
      setMeals(prev => {
        const key = `${picker.date}|${picker.occasion}`;
        return { ...prev, [key]: [...(prev[key] || []), r.data] };
      });
    } catch { toast("error", "Failed to plan meal"); }
    setPicker(null);
  };

  const addEmptyMeal = async (date: string, occasion: string) => {
    if (!plan) return;
    try {
      const r = await getAuthenticated(() => api.planMeal(plan.id, { meal_date: date, meal_occasion: occasion }));
      toast("success", "Meal slot reserved");
      setMeals(prev => {
        const key = `${date}|${occasion}`;
        return { ...prev, [key]: [...(prev[key] || []), r.data] };
      });
    } catch { toast("error", "Failed to reserve slot"); }
  };

  if (!account) return null;

  const days = plan ? computeDates(plan.period_start) : [];

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-xl font-bold text-kuali-950">Planner</h1>
        <button onClick={() => setCreate(!create)} className="bg-rempah-500 text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-rempah-700 transition-colors focus:outline-none focus:ring-2 focus:ring-rempah-500/50">
          + New plan
        </button>
      </div>

      {create && (
        <form onSubmit={submitPlan} className="bg-white border border-bambu-200 rounded-lg p-4 space-y-3">
          <label htmlFor="plan-title" className="sr-only">Title</label>
          <input id="plan-title" type="text" placeholder="Title (optional)" value={title} onChange={e => setTitle(e.target.value)} className="w-full" />
          <div className="flex gap-2">
            <label htmlFor="plan-start" className="sr-only">Start date</label>
            <input id="plan-start" type="date" value={start} onChange={e => setStart(e.target.value)} required className="flex-1" />
            <label htmlFor="plan-end" className="sr-only">End date</label>
            <input id="plan-end" type="date" value={end} onChange={e => setEnd(e.target.value)} required className="flex-1" />
          </div>
          <button type="submit" className="bg-rempah-500 text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-rempah-700 transition-colors focus:outline-none focus:ring-2 focus:ring-rempah-500/50">Create</button>
        </form>
      )}

      {!plan && (
        <div>
          {plans.length === 0 && (
            <div className="text-center py-12 text-kuali-700 bg-white border border-bambu-200 rounded-xl">
              <p className="text-lg">No meal plans yet.</p>
              <p className="text-sm text-bambu-300 mt-1">Create a weekly plan to organize your meals.</p>
            </div>
          )}
          <div className="space-y-2">
            {plans.map(p => (
              <div key={p.id} onClick={() => openPlan(p)} className="bg-white border border-bambu-200 rounded-lg p-4 cursor-pointer hover:border-rempah-500/30 transition-colors">
                <div className="flex justify-between items-center">
                  <div>
                    <p className="font-semibold text-kuali-950">{p.title || "Meal plan"}</p>
                    <p className="text-sm text-kuali-700">{p.period_start} to {p.period_end}</p>
                  </div>
                  <span className={`text-xs px-2 py-0.5 rounded ${p.status === "completed" ? "bg-context-positive/20 text-context-positive-dark" : p.status === "cancelled" ? "bg-feedback-error/10 text-feedback-error" : "bg-feedback-info/10 text-feedback-info"}`}>{p.status}</span>
                </div>
              </div>
            ))}
          </div>
        </div>
      )}

      {plan && loading && <ListSkeleton count={4} />}

      {plan && !loading && (
        <div>
          <div className="flex justify-between items-center mb-4">
            <button onClick={() => setPlan(null)} className="text-rempah-500 text-sm font-medium hover:underline">← Back to plans</button>
            <span className="text-sm text-bambu-300">{plan.period_start} to {plan.period_end}</span>
          </div>
          <div className="overflow-x-auto">
            <div className="grid grid-cols-7 gap-1 min-w-[700px]">
              {days.map(dd => (
                <div key={dd.d} className="text-center text-xs font-semibold text-kuali-700 py-1 truncate">{dd.label}</div>
              ))}
              {days.map(dd => (
                <div key={dd.d} className="border border-bambu-200 rounded p-1 min-h-[90px] bg-white">
                  {OCCASIONS.map(o => {
                    const key = `${dd.d}|${o}`;
                    const slotMeals = meals[key] || [];
                    return (
                      <div key={o} className="text-xs mb-1">
                        <span className="text-bambu-300">{o.slice(0, 2)}</span>
                        {slotMeals.length > 0 ? (
                          slotMeals.map(m => (
                            <span key={m.id} className="ml-1 text-context-positive-dark font-medium truncate block">{m.recipe_id ? "🍽" : "—"}</span>
                          ))
                        ) : (
                          <button
                            type="button"
                            onClick={() => openPicker(dd.d, o)}
                            className="ml-1 text-bambu-300 hover:text-rempah-500 transition-colors"
                            aria-label={`Add ${o} on ${dd.d}`}
                          >+</button>
                        )}
                      </div>
                    );
                  })}
                </div>
              ))}
            </div>
          </div>
        </div>
      )}

      {picker && (
        <div className="fixed inset-0 bg-kuali-950/50 flex items-center justify-center p-4 z-50" onClick={() => setPicker(null)}>
          <div className="bg-white rounded-xl max-w-md w-full max-h-[70vh] overflow-y-auto p-6 space-y-3" onClick={e => e.stopPropagation()} aria-label="Select a recipe">
            <div className="flex justify-between items-center">
              <h2 className="font-semibold text-kuali-950">Pick a recipe for {picker.occasion} on {picker.date}</h2>
              <button onClick={() => setPicker(null)} className="text-bambu-300 hover:text-kuali-700 text-sm">✕</button>
            </div>
            <button
              type="button"
              onClick={() => { addEmptyMeal(picker.date, picker.occasion); setPicker(null); }}
              className="w-full text-left text-sm text-kuali-700 hover:bg-bambu-200 rounded-lg px-3 py-2 transition-colors"
            >
              Just reserve this slot (no recipe)
            </button>
            {recipeLoading ? (
              <div className="space-y-2">{Array.from({length:3}).map((_,i)=><Skeleton key={i} className="h-12 w-full"/>)}</div>
            ) : (
              <div className="space-y-1">
                {recipes.map(r => (
                  <button
                    key={r.id}
                    type="button"
                    onClick={() => pickRecipe(r)}
                    className="w-full text-left bg-bambu-200/50 hover:bg-bambu-200 rounded-lg px-3 py-2 transition-colors"
                  >
                    <span className="font-medium text-kuali-950 text-sm">{r.title}</span>
                    <span className="text-xs text-kuali-700 ml-2">{r.servings} servings</span>
                  </button>
                ))}
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
