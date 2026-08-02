"use client";
import { useEffect, useState } from "react";
import { useAuth } from "@/lib/auth";
import { api, type PantryItem } from "@/lib/api";
import { useToast } from "@/lib/toast";
import { useRouter } from "next/navigation";
import { ListSkeleton } from "@/components/skeleton";

export default function PantryPage() {
  const { account } = useAuth();
  const router = useRouter();
  const { toast } = useToast();
  const [items, setItems] = useState<PantryItem[]>([]);
  const [showAdd, setShowAdd] = useState(false);
  const [loading, setLoading] = useState(true);

  const load = () => {
    setLoading(true);
    api.pantryItems().then(r => setItems(r.data)).catch(() => toast("error", "Failed to load pantry")).finally(() => setLoading(false));
  };
  useEffect(() => { if (account) load(); }, [account]);

  if (!account) { router.push("/login"); return null; }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-xl font-bold text-ink-900" id="pantry-title">Pantry</h1>
        <button onClick={() => setShowAdd(!showAdd)} className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-action-dark transition-colors focus:outline-none focus:ring-2 focus:ring-action-primary/50" aria-label={showAdd ? "Cancel" : "Add item"}>
          {showAdd ? "Cancel" : "+ Add item"}
        </button>
      </div>

      {showAdd && <AddItemForm onAdded={() => { setShowAdd(false); toast("success", "Item added"); load(); }} />}

      {loading ? <ListSkeleton count={4} /> : items.length === 0 && !showAdd ? (
        <div className="text-center py-8 text-ink-700">
          <p className="text-lg">No ingredients yet.</p>
          <p className="text-sm text-steel-400 mt-1">Add your first ingredient above.</p>
        </div>
      ) : (
        <div className="grid gap-2" role="list" aria-labelledby="pantry-title">
          {items.map(item => (
            <div key={item.id} className="bg-white-000 border border-steel-200 rounded-lg p-3 flex justify-between items-center" role="listitem">
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
      )}
    </div>
  );
}

function AddItemForm({ onAdded }: { onAdded: () => void }) {
  const { toast } = useToast();
  const [name, setName] = useState("");
  const [category, setCategory] = useState("");
  const [qty, setQty] = useState("1");
  const [unit, setUnit] = useState("");
  const [expiry, setExpiry] = useState("");

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await api.addPantryItem({ ingredient_name: name, category: category || "other", quantity: parseFloat(qty) || 1, unit: unit || "unit", expiry_date: expiry || undefined });
      onAdded();
    } catch {
      toast("error", "Failed to add item");
    }
  };

  return (
    <form onSubmit={submit} className="bg-white-000 border border-steel-200 rounded-lg p-4 space-y-3" aria-label="Add pantry item">
      <label htmlFor="ingredient-name" className="sr-only">Ingredient name</label>
      <input id="ingredient-name" type="text" placeholder="Ingredient name" value={name} onChange={e => setName(e.target.value)} required className="w-full" />
      <div className="flex gap-2">
        <label htmlFor="category" className="sr-only">Category</label>
        <input id="category" type="text" placeholder="Category" value={category} onChange={e => setCategory(e.target.value)} className="flex-1" />
        <label htmlFor="qty" className="sr-only">Quantity</label>
        <input id="qty" type="number" placeholder="Qty" value={qty} onChange={e => setQty(e.target.value)} className="w-20" min="0" />
        <label htmlFor="unit" className="sr-only">Unit</label>
        <input id="unit" type="text" placeholder="Unit" value={unit} onChange={e => setUnit(e.target.value)} className="w-20" />
      </div>
      <label htmlFor="expiry" className="sr-only">Expiry date</label>
      <input id="expiry" type="date" value={expiry} onChange={e => setExpiry(e.target.value)} className="w-full" />
      <button type="submit" className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-action-dark transition-colors focus:outline-none focus:ring-2 focus:ring-action-primary/50">
        Save
      </button>
    </form>
  );
}
