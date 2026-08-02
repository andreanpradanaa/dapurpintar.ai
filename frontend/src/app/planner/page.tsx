"use client";
import { useAuth } from "@/lib/auth";
import { useRouter } from "next/navigation";

export default function PlannerPage() {
  const { account } = useAuth();
  const router = useRouter();
  if (!account) { router.push("/login"); return null; }
  return (
    <div className="space-y-6">
      <h1 className="text-xl font-bold text-ink-900">Planner</h1>
      <div className="text-center py-12 text-ink-700 bg-white-000 border border-steel-200 rounded-xl">
        <p className="text-lg">Meal planning coming soon.</p>
        <p className="text-sm text-steel-400 mt-1">Create weekly meal plans based on what is in your pantry.</p>
      </div>
    </div>
  );
}
