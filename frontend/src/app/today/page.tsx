"use client";
import { useEffect, useState, useCallback } from "react";
import { useAuth } from "@/lib/auth";
import { api, type PantryAnalysis, type PantrySummary, type PantryItem } from "@/lib/api";
import { useToast } from "@/lib/toast";
import { StatSkeleton, ListSkeleton } from "@/components/skeleton";

export default function TodayPage() {
  const { account, loading } = useAuth();
  const { toast } = useToast();
  const [summary, setSummary] = useState<PantrySummary | null>(null);
  const [expiring, setExpiring] = useState<PantryItem[]>([]);
  const [loadingData, setLoadingData] = useState(true);
  const [analysis, setAnalysis] = useState<PantryAnalysis | null>(null);
  const [analysisLoading, setAnalysisLoading] = useState(false);
  const [analysisError, setAnalysisError] = useState("");
  const [loadError, setLoadError] = useState(false);

  const loadData = useCallback(() => {
    if (!account) return;
    setLoadingData(true);
    setLoadError(false);
    api.refreshPantryStatuses().catch(() => {});
    Promise.all([
      api.pantrySummary().then(r => setSummary(r.data)).catch(() => { throw new Error(); }),
      api.expiringItems().then(r => setExpiring(r.data)).catch(() => {}),
    ]).catch(() => setLoadError(true)).finally(() => setLoadingData(false));
  }, [account]);

  useEffect(() => { loadData(); }, [loadData]);

  if (loading) return <div className="py-12 space-y-4">{Array.from({length:3}).map((_,i)=><StatSkeleton key={i}/>)}</div>;
  if (!account) return null;

  const analyzePantry = async () => {
    setAnalysisLoading(true);
    setAnalysisError("");
    try {
      const result = await api.analyzePantry();
      setAnalysis(result.data);
    } catch {
      setAnalysisError("Pantry analysis is unavailable right now.");
    } finally {
      setAnalysisLoading(false);
    }
  };

  return (
    <div className="space-y-6">
      <h1 className="text-xl font-bold text-ink-900">Today</h1>

      {!loadingData && loadError && (
        <div className="text-center py-4 text-ink-700 bg-white-000 border border-steel-200 rounded-xl">
          <p className="text-sm">Failed to load pantry data.</p>
          <button onClick={loadData} className="mt-2 text-action-primary text-sm font-medium hover:underline">Retry</button>
        </div>
      )}

      {loadingData ? (
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <StatSkeleton /><StatSkeleton /><StatSkeleton />
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <StatCard label="Pantry items" value={summary?.total_items ?? "—"} color="ink" />
          <StatCard label="Expiring soon" value={summary?.expiring_soon_count ?? "—"} color="attention" />
          <StatCard label="Running low" value={summary?.running_low_count ?? "—"} color="info" />
        </div>
      )}

      {loadingData && expiring.length === 0 ? <ListSkeleton count={2} /> : expiring.length > 0 && (
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

      {!loadingData && summary && summary.total_items === 0 && (
        <div className="text-center py-8 text-ink-700">
          <p className="text-lg">Pantry kamu kosong.</p>
          <p className="text-sm text-steel-400 mt-1">Add ingredients to get started.</p>
        </div>
      )}

      <section className="bg-ink-900 text-white-000 rounded-xl p-5 space-y-4" aria-labelledby="analysis-title">
        <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-3">
          <div>
            <h2 id="analysis-title" className="font-semibold">Kitchen read</h2>
            <p className="text-sm text-steel-200 mt-1">See what to use first from your pantry.</p>
          </div>
          <button
            type="button"
            onClick={analyzePantry}
            disabled={analysisLoading}
            className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-action-dark disabled:opacity-60 focus:outline-none focus:ring-2 focus:ring-action-primary/50"
          >
            {analysisLoading ? "Analyzing..." : "Analyze pantry"}
          </button>
        </div>

        {analysisError && <p role="alert" className="text-sm text-red-200">{analysisError}</p>}
        {analysis && analysis.use_first_opportunities.length === 0 && analysis.optimization_suggestions.length === 0 && (
          <p className="text-sm text-steel-200">No suggestions yet. Add pantry items or try again when AI is available.</p>
        )}
        {analysis && analysis.use_first_opportunities.length > 0 && (
          <div>
            <h3 className="text-sm font-medium text-context-attention">Use first</h3>
            <ul className="mt-2 space-y-2">
              {analysis.use_first_opportunities.map(item => (
                <li key={item.pantry_item_id} className="text-sm">
                  <span className="font-medium">{item.ingredient_name}</span>
                  <span className="text-steel-200"> · {item.reason}</span>
                </li>
              ))}
            </ul>
          </div>
        )}
        {analysis && analysis.optimization_suggestions.length > 0 && (
          <div>
            <h3 className="text-sm font-medium text-context-positive">Ideas</h3>
            <ul className="mt-2 space-y-2">
              {analysis.optimization_suggestions.map((suggestion, index) => (
                <li key={`${suggestion.title}-${index}`} className="text-sm">
                  <span className="font-medium">{suggestion.title}</span>
                  <span className="text-steel-200"> · {suggestion.description}</span>
                </li>
              ))}
            </ul>
          </div>
        )}
      </section>
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
