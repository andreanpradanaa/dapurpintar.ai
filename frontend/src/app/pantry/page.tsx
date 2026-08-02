"use client";
import { useEffect, useState } from "react";
import { useAuth } from "@/lib/auth";
import { api, type PantryItem } from "@/lib/api";
import { useRouter } from "next/navigation";

export default function PantryPage() {
  const { account } = useAuth();
  const router = useRouter();
  const [items, setItems] = useState<PantryItem[]>([]);
  const [showAdd, setShowAdd] = useState(false);
  const [err, setErr] = useState("");

  const load = () => {
    api.pantryItems().then(r => setItems(r.data)).catch(() => setErr("Failed to load"));
  };
  useEffect(() => { if (account) load(); }, [account]);

  if (!account) { router.push("/login"); return null; }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-xl font-bold text-ink-900">Pantry</h1>
        <button onClick={() => setShowAdd(!showAdd)} className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-action-dark transition-colors">
          + Add item
        </button>
      </div>

      {showAdd && <AddItemForm onAdded={() => { setShowAdd(false); load(); }} />}

      {err && <p className="text-feedback-error text-sm">{err}</p>}

      {items.length === 0 && !showAdd && (
        <div className="text-center py-8 text-ink-700">
          <p className="text-lg">No ingredients yet.</p>
          <p className="text-sm text-steel-400 mt-1">Add your first ingredient above.</p>
        </div>
      )}

      <div className="grid gap-2">
        {items.map(item => (
          <div key={item.id} className="bg-white-000 border border-steel-200 rounded-lg p-3 flex justify-between items-center">
            <div>
              <p className="font-medium text-ink-900">{item.ingredient_name}</p>
              <p className="text-xs text-ink-700">{item.quantity} {item.unit} · {item.category}</p>
            </div>
            <div className="flex items-center gap-2">
              {item.expiry_date && <span className="text-xs text-steel-400">{item.expiry_date}</span>}
              <span className={`text-xs px-2 py-0.5 rounded ${item.status === "expiring_soon" ? "bg-context-attention/20 text-context-attention-dark" : item.status === "running_low" ? "bg-feedback-info/10 text-feedback-info" : "bg-context-positive/20 text-context-positive-dark"}`}>{item.status}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

function AddItemForm({ onAdded }: { onAdded: () => void }) {
  const [name, setName] = useState("");
  const [category, setCategory] = useState("");
  const [qty, setQty] = useState("1");
  const [unit, setUnit] = useState("");
  const [expiry, setExpiry] = useState("");
  const [err, setErr] = useState("");

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await api.addPantryItem({ ingredient_name: name, category: category || "other", quantity: parseFloat(qty) || 1, unit: unit || "unit", expiry_date: expiry || undefined });
      onAdded();
    } catch (ex: unknown) {
      setErr(ex instanceof Error ? ex.message : "Failed");
    }
  };

  return (
    <form onSubmit={submit} className="bg-white-000 border border-steel-200 rounded-lg p-4 space-y-3">
      <input type="text" placeholder="Ingredient name" value={name} onChange={e => setName(e.target.value)} required className="w-full" />
      <div className="flex gap-2">
        <input type="text" placeholder="Category" value={category} onChange={e => setCategory(e.target.value)} className="flex-1" />
        <input type="number" placeholder="Qty" value={qty} onChange={e => setQty(e.target.value)} className="w-20" min="0" />
        <input type="text" placeholder="Unit" value={unit} onChange={e => setUnit(e.target.value)} className="w-20" />
      </div>
      <input type="date" value={expiry} onChange={e => setExpiry(e.target.value)} className="w-full" />
      {err && <p className="text-feedback-error text-sm">{err}</p>}
      <button type="submit" className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-action-dark transition-colors">
        Save
      </button>
    </form>
  );
}
