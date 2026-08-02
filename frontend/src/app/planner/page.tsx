"use client";
import { useEffect, useState, useCallback } from "react";
import { useAuth } from "@/lib/auth";
import { api, type MealPlan } from "@/lib/api";
import { useRouter } from "next/navigation";

const DAYS = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"];
const OCCASIONS = ["breakfast", "lunch", "dinner", "snack"] as const;

export default function PlannerPage() {
  const { account } = useAuth();
  const router = useRouter();
  const [plans, setPlans] = useState<MealPlan[]>([]);
  const [plan, setPlan] = useState<MealPlan | null>(null);
  const [meals, setMeals] = useState<Record<string, string[]>>({});
  const [loading, setLoading] = useState(false);
  const [err, setErr] = useState("");
  const [create, setCreate] = useState(false);
  const [title, setTitle] = useState("");
  const [start, setStart] = useState("");
  const [end, setEnd] = useState("");

  const loadPlans = useCallback(() => {
    api.mealPlans().then(r => setPlans(r.data)).catch(() => setErr("Failed to load plans"));
  }, []);

  useEffect(() => { if (account) loadPlans(); }, [account, loadPlans]);

  const openPlan = async (p: MealPlan) => {
    setPlan(p);
    setLoading(true);
    try {
      const r = await api.plannedMeals(p.id);
      const map: Record<string, string[]> = {};
      r.data.forEach(m => {
        const key = `${m.meal_date}|${m.meal_occasion}`;
        if (!map[key]) map[key] = [];
        map[key].push(m.recipe_id || "—");
      });
      setMeals(map);
    } catch { setErr("Failed to load meals"); }
    setLoading(false);
  };

  const submitPlan = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const r = await api.createMealPlan({ period_start: start, period_end: end, title: title || "Meal plan" });
      setCreate(false);
      setPlans(prev => [r.data, ...prev]);
      openPlan(r.data);
    } catch { setErr("Failed to create plan"); }
  };

  if (!account) { router.push("/login"); return null; }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-xl font-bold text-ink-900">Planner</h1>
        <button onClick={() => setCreate(!create)} className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-action-dark transition-colors">
          + New plan
        </button>
      </div>

      {create && (
        <form onSubmit={submitPlan} className="bg-white-000 border border-steel-200 rounded-lg p-4 space-y-3">
          <input type="text" placeholder="Title (optional)" value={title} onChange={e => setTitle(e.target.value)} className="w-full" />
          <div className="flex gap-2">
            <input type="date" value={start} onChange={e => setStart(e.target.value)} required className="flex-1" />
            <input type="date" value={end} onChange={e => setEnd(e.target.value)} required className="flex-1" />
          </div>
          {err && <p className="text-feedback-error text-sm">{err}</p>}
          <button type="submit" className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium">Create</button>
        </form>
      )}

      {!plan && (
        <div className="space-y-2">
          {plans.map(p => (
            <div key={p.id} onClick={() => openPlan(p)} className="bg-white-000 border border-steel-200 rounded-lg p-4 cursor-pointer hover:border-action-primary/30 transition-colors">
              <div className="flex justify-between items-center">
                <div>
                  <p className="font-semibold text-ink-900">{p.title || "Meal plan"}</p>
                  <p className="text-sm text-ink-700">{p.period_start} to {p.period_end}</p>
                </div>
                <span className={`text-xs px-2 py-0.5 rounded ${p.status === "completed" ? "bg-context-positive/20 text-context-positive-dark" : p.status === "cancelled" ? "bg-feedback-error/10 text-feedback-error" : "bg-feedback-info/10 text-feedback-info"}`}>{p.status}</span>
              </div>
            </div>
          ))}
          {plans.length === 0 && (
            <div className="text-center py-12 text-ink-700 bg-white-000 border border-steel-200 rounded-xl">
              <p className="text-lg">No meal plans yet.</p>
              <p className="text-sm text-steel-400 mt-1">Create a weekly plan to organize your meals.</p>
            </div>
          )}
        </div>
      )}

      {plan && loading && <p className="text-steel-400 text-center py-8">Loading meals...</p>}

      {plan && !loading && (
        <div>
          <div className="flex justify-between items-center mb-4">
            <button onClick={() => setPlan(null)} className="text-action-primary text-sm font-medium">← Back to plans</button>
            <span className="text-sm text-steel-400">{plan.period_start} to {plan.period_end}</span>
          </div>
          <div className="overflow-x-auto">
            <div className="grid grid-cols-7 gap-1 min-w-[600px]">
              {DAYS.map(d => <div key={d} className="text-center text-xs font-semibold text-ink-700 py-1">{d}</div>)}
              {DAYS.map(d => (
                <div key={d} className="border border-steel-200 rounded p-1 min-h-[80px] bg-white-000">
                  {OCCASIONS.map(o => {
                    const key = `${d}|${o}`; // simplified - in real app, compute actual date
                    return (
                      <div key={o} className="text-xs mb-1">
                        <span className="text-steel-400">{o[0].toUpperCase()}</span>
                        {meals[key] ? <span className="ml-1 text-context-positive-dark">•</span> : <span className="ml-1 text-steel-200">—</span>}
                      </div>
                    );
                  })}
                </div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
