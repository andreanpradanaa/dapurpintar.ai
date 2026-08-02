"use client";
import { useEffect, useState, useCallback } from "react";
import { useAuth } from "@/lib/auth";
import { api, type ShoppingList, type ShoppingItem } from "@/lib/api";

export default function ShoppingPage() {
  const { account } = useAuth();
  const [lists, setLists] = useState<ShoppingList[]>([]);
  const [list, setList] = useState<ShoppingList | null>(null);
  const [items, setItems] = useState<ShoppingItem[]>([]);
  const [loading, setLoading] = useState(false);
  const [err, setErr] = useState("");
  const [create, setCreate] = useState(false);
  const [title, setTitle] = useState("");
  const [itemName, setItemName] = useState("");
  const [itemQty, setItemQty] = useState("1");
  const [itemUnit, setItemUnit] = useState("");

  const loadLists = useCallback(() => {
    api.shoppingLists().then(r => setLists(r.data)).catch(() => setErr("Failed to load"));
  }, []);

  useEffect(() => { if (account) loadLists(); }, [account, loadLists]);

  const openList = async (l: ShoppingList) => {
    setList(l);
    setLoading(true);
    try {
      const r = await api.shoppingItems(l.id);
      setItems(r.data);
    } catch { setErr("Failed to load items"); }
    setLoading(false);
  };

  const addItem = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!list) return;
    try {
      await api.addShoppingItem(list.id, { ingredient_name: itemName, quantity: parseFloat(itemQty) || 1, unit: itemUnit || "unit" });
      setItemName(""); setItemQty("1"); setItemUnit("");
      const r = await api.shoppingItems(list.id);
      setItems(r.data);
    } catch { setErr("Failed to add item"); }
  };

  const toggleItem = async (item: ShoppingItem) => {
    if (!list) return;
    try {
      await api.completeShoppingItem(list.id, item.id);
      const r = await api.shoppingItems(list.id);
      setItems(r.data);
    } catch { /* ignore */ }
  };

  const createList = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const r = await api.createShoppingList({ title });
      setCreate(false); setTitle("");
      openList(r.data);
      loadLists();
    } catch { setErr("Failed to create list"); }
  };

  const activate = async () => { if (list) { const r = await api.activateShoppingList(list.id); setList(r.data); } };
  const complete = async () => { if (list) { const r = await api.completeShoppingList(list.id); setList(r.data); } };

  if (!account) return null;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-xl font-bold text-ink-900">Shopping</h1>
        <button onClick={() => setCreate(!create)} className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-action-dark transition-colors">
          + New list
        </button>
      </div>

      {create && (
        <form onSubmit={createList} className="bg-white-000 border border-steel-200 rounded-lg p-4 space-y-3">
          <input type="text" placeholder="List title" value={title} onChange={e => setTitle(e.target.value)} required className="w-full" />
          {err && <p className="text-feedback-error text-sm">{err}</p>}
          <button type="submit" className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium">Create</button>
        </form>
      )}

      {!list && (
        <div className="space-y-2">
          {lists.map(l => (
            <div key={l.id} onClick={() => openList(l)} className="bg-white-000 border border-steel-200 rounded-lg p-4 cursor-pointer hover:border-action-primary/30 transition-colors">
              <div className="flex justify-between items-center">
                <div>
                  <p className="font-semibold text-ink-900">{l.title}</p>
                  <p className="text-xs text-ink-700">{l.item_counts.open} open · {l.item_counts.completed} done</p>
                </div>
                <span className={`text-xs px-2 py-0.5 rounded ${l.status === "active" ? "bg-context-positive/20 text-context-positive-dark" : l.status === "completed" ? "bg-steel-200 text-steel-400" : "bg-feedback-info/10 text-feedback-info"}`}>{l.status}</span>
              </div>
            </div>
          ))}
          {lists.length === 0 && (
            <div className="text-center py-12 text-ink-700 bg-white-000 border border-steel-200 rounded-xl">
              <p className="text-lg">No shopping lists yet.</p>
              <p className="text-sm text-steel-400 mt-1">Create a list to start tracking your groceries.</p>
            </div>
          )}
        </div>
      )}

      {list && (
        <div>
          <div className="flex items-center justify-between mb-4 gap-2">
            <button onClick={() => { setList(null); loadLists(); }} className="text-action-primary text-sm font-medium">← Back</button>
            <h2 className="text-lg font-semibold text-ink-900">{list.title}</h2>
            <div className="flex gap-2">
              {list.status === "draft" || list.status === "generated" || list.status === "reviewed" ? (
                <button onClick={activate} className="bg-context-positive text-context-positive-dark text-xs px-3 py-1 rounded font-medium">Activate</button>
              ) : null}
              {list.status === "active" ? (
                <button onClick={complete} className="bg-action-primary text-white-000 text-xs px-3 py-1 rounded font-medium">Complete</button>
              ) : null}
            </div>
          </div>

          <form onSubmit={addItem} className="flex gap-2 mb-4">
            <input type="text" placeholder="Item name" value={itemName} onChange={e => setItemName(e.target.value)} required className="flex-1" />
            <input type="number" placeholder="Qty" value={itemQty} onChange={e => setItemQty(e.target.value)} className="w-16" min="1" />
            <input type="text" placeholder="Unit" value={itemUnit} onChange={e => setItemUnit(e.target.value)} className="w-16" />
            <button type="submit" className="bg-action-primary text-white-000 px-3 py-2 rounded-lg text-sm font-medium">Add</button>
          </form>

          {loading && <p className="text-steel-400 text-center py-4">Loading...</p>}

          <div className="space-y-1">
            {items.map(item => (
              <div key={item.id} onClick={() => toggleItem(item)} className={`bg-white-000 border rounded-lg p-3 flex justify-between items-center cursor-pointer transition-colors ${item.status === "completed" ? "border-steel-200 opacity-50" : "border-steel-200 hover:border-action-primary/30"}`}>
                <div className="flex items-center gap-2">
                  <span className={`text-sm ${item.status === "completed" ? "text-context-positive-dark line-through" : "text-ink-900"}`}>{item.ingredient_name}</span>
                  <span className="text-xs text-steel-400">{item.quantity} {item.unit}</span>
                </div>
                <span className="text-xs text-steel-400">{item.status}</span>
              </div>
            ))}
          </div>

          {items.length === 0 && !loading && (
            <div className="text-center py-8 text-ink-700">
              <p className="text-sm">List is empty. Add items above.</p>
            </div>
          )}
        </div>
      )}
    </div>
  );
}
