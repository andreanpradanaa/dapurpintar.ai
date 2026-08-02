"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { useAuth } from "@/lib/auth";

export default function LoginPage() {
  const { login, register, account, loading } = useAuth();
  const router = useRouter();
  const [mode, setMode] = useState<"login" | "register">("login");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [error, setError] = useState("");

  if (loading) return (
    <div className="min-h-screen flex items-center justify-center bg-paper-050">
      <div className="animate-pulse text-steel-400 text-sm">Loading...</div>
    </div>
  );

  if (account) { router.replace("/today"); return null; }

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    try {
      if (mode === "login") await login(email, password);
      else await register(email, password, name);
      router.push("/today");
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : "Something went wrong");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-paper-050 px-4">
      <div className="w-full max-w-sm">
        <h1 className="text-2xl font-bold text-center mb-2 text-ink-900">DapurPintar</h1>
        <p className="text-center text-ink-700 text-sm mb-8">Decide dinner with what you have.</p>

        <form onSubmit={submit} className="space-y-4 bg-white-000 p-6 rounded-xl border border-steel-200">
          <h2 className="font-semibold text-ink-900">{mode === "login" ? "Log in" : "Create account"}</h2>

          {mode === "register" && (
            <input type="text" placeholder="Display name" value={name} onChange={e => setName(e.target.value)} required className="w-full" />
          )}
          <input type="email" placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} required className="w-full" />
          <input type="password" placeholder="Password (min 8 chars)" value={password} onChange={e => setPassword(e.target.value)} required minLength={8} className="w-full" />

          {error && <p className="text-feedback-error text-sm">{error}</p>}

          <button type="submit" className="w-full bg-action-primary text-white-000 py-2 rounded-lg font-medium hover:bg-action-dark transition-colors">
            {mode === "login" ? "Log in" : "Create account"}
          </button>

          <p className="text-sm text-center text-ink-700">
            {mode === "login" ? "Belum punya akun? " : "Sudah punya akun? "}
            <button type="button" onClick={() => setMode(mode === "login" ? "register" : "login")} className="text-action-primary font-medium">
              {mode === "login" ? "Daftar" : "Log in"}
            </button>
          </p>
        </form>
      </div>
    </div>
  );
}
