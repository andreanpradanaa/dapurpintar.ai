"use client";
import { useState } from "react";
import { useAuth } from "@/lib/auth";
import { api } from "@/lib/api";
import { useRouter } from "next/navigation";

export default function ProfilePage() {
  const { account, profile, logout, refresh } = useAuth();
  const router = useRouter();
  const [name, setName] = useState(profile?.display_name || "");
  const [msg, setMsg] = useState("");

  if (!account) { router.push("/login"); return null; }

  const save = async () => {
    try {
      await api.updateProfile({ display_name: name });
      await refresh();
      setMsg("Profile updated.");
    } catch { setMsg("Failed to update."); }
  };

  return (
    <div className="max-w-md space-y-6">
      <h1 className="text-xl font-bold text-ink-900">Profile</h1>
      <div className="bg-white-000 border border-steel-200 rounded-xl p-6 space-y-4">
        <div>
          <p className="text-sm text-ink-700">Email</p>
          <p className="text-ink-900 font-medium">{account.email}</p>
        </div>
        <div>
          <label className="text-sm text-ink-700">Display name</label>
          <input type="text" value={name} onChange={e => setName(e.target.value)} className="w-full mt-1" />
        </div>
        {msg && <p className="text-sm text-context-positive-dark">{msg}</p>}
        <button onClick={save} className="bg-action-primary text-white-000 px-4 py-2 rounded-lg text-sm font-medium hover:bg-action-dark transition-colors">
          Save
        </button>
      </div>

      <button onClick={() => logout().then(() => router.push("/login"))} className="text-feedback-error text-sm font-medium">
        Log out
      </button>
    </div>
  );
}
