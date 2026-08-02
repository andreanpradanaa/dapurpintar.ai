"use client";
import { useEffect, useState } from "react";
import { useAuth } from "@/lib/auth";
import { api, type PantrySummary, type PantryItem } from "@/lib/api";
import { useRouter } from "next/navigation";

export default function TodayPage() {
  const { account, loading } = useAuth();
  const router = useRouter();
  const [summary, setSummary] = useState<PantrySummary | null>(null);
  const [expiring, setExpiring] = useState<PantryItem[]>([]);
  const [err, setErr] = useState("");

  useEffect(() => {
    if (!account) return;
    api.pantrySummary().then(r => setSummary(r.data)).catch(() => setErr("Failed to load pantry"));
    api.expiringItems().then(r => setExpiring(r.data)).catch(() => {});
  }, [account]);

  if (loading) return <div className="py-12 text-center text-steel-400">Loading...</div>;
  if (!account) { router.push("/login"); return null; }

  return (
    <div className="space-y-6">
      <h1 className="text-xl font-bold text-ink-900">Today</h1>
      {err && <p className="text-feedback-error text-sm">{err}</p>}

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <StatCard label="Pantry items" value={summary?.total_items ?? "—"} color="ink" />
        <StatCard label="Expiring soon" value={summary?.expiring_soon_count ?? "—"} color="attention" />
        <StatCard label="Running low" value={summary?.running_low_count ?? "—"} color="info" />
      </div>

      {expiring.length > 0 && (
        <section>
          <h2 className="font-semibold text-ink-900 mb-3">Expiring soon</h2>
          <div className="space-y-2">
            {expiring.map(item => (
              <div key={item.id} className="bg-white-000 border border-steel-200 rounded-lg p-3 flex justify-between items-center">
                <div>
                  <p className="font-medium text-ink-900">{item.ingredient_name}</p>
                  <p className="text-sm text-ink-700">{item.quantity} {item.unit}</p>
                </div>
                {item.expiry_date && <span className="text-xs text-context-attention-dark bg-context-attention/20 px-2 py-1 rounded">{item.expiry_date}</span>}
              </div>
            ))}
          </div>
        </section>
      )}

      {summary && summary.total_items === 0 && (
        <div className="text-center py-8 text-ink-700">
          <p className="text-lg">Pantry kamu kosong.</p>
          <p className="text-sm text-steel-400 mt-1">Add ingredients to get started.</p>
        </div>
      )}
    </div>
  );
}

function StatCard({ label, value, color }: { label: string; value: number | string; color: string }) {
  const colors: Record<string, string> = { ink: "bg-ink-900", attention: "bg-context-attention", info: "bg-feedback-info" };
  return (
    <div className="bg-white-000 border border-steel-200 rounded-xl p-4">
      <p className="text-sm text-ink-700">{label}</p>
      <p className={`text-3xl font-bold mt-1 ${color === "ink" ? "text-ink-900" : color === "attention" ? "text-context-attention-dark" : "text-feedback-info"}`}>{value}</p>
    </div>
  );
}
