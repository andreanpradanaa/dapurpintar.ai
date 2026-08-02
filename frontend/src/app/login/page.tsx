"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
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
    <div className="min-h-screen flex items-center justify-center bg-santan-050">
      <div className="animate-pulse-soft text-kuali-700/40 text-sm">Loading...</div>
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
    <div className="min-h-screen flex bg-santan-050">
      {/* Left panel — visual */}
      <div className="hidden md:flex md:w-5/12 bg-gradient-to-br from-kuali-950 via-kuali-700 to-rempah-700 items-center justify-center relative overflow-hidden">
        <div className="absolute inset-0 opacity-10 bg-[radial-gradient(circle_at_30%_20%,_#fef9f0_0%,_transparent_50%)]" />
        <div className="relative text-center px-10 space-y-6">
          <p className="font-display text-5xl font-bold text-white leading-tight">
            Decide dinner with what you have.
          </p>
          <p className="text-santan-200/60 text-sm max-w-xs mx-auto text-balance">
            The AI kitchen companion that knows your pantry and your taste.
          </p>
        </div>
      </div>

      {/* Right panel — form */}
      <div className="flex-1 flex items-center justify-center px-6">
        <div className="w-full max-w-sm space-y-8">
          <div className="text-center md:text-left">
            <Link href="/" className="font-display text-3xl font-bold text-kuali-950 hover:text-rempah-500 transition-colors">DapurPintar</Link>
            <p className="text-kuali-700/50 text-sm mt-2">{mode === "login" ? "Welcome back." : "Start your kitchen journey."}</p>
          </div>

          <form onSubmit={submit} className="space-y-4">
            {mode === "register" && (
              <div>
                <label htmlFor="name" className="block text-xs font-medium text-kuali-700/60 mb-1.5">Display name</label>
                <input id="name" type="text" placeholder="Your name" value={name} onChange={e => setName(e.target.value)} required className="w-full" />
              </div>
            )}
            <div>
              <label htmlFor="email" className="block text-xs font-medium text-kuali-700/60 mb-1.5">Email</label>
              <input id="email" type="email" placeholder="you@example.com" value={email} onChange={e => setEmail(e.target.value)} required className="w-full" />
            </div>
            <div>
              <label htmlFor="password" className="block text-xs font-medium text-kuali-700/60 mb-1.5">Password</label>
              <input id="password" type="password" placeholder="Min 8 characters" value={password} onChange={e => setPassword(e.target.value)} required minLength={8} className="w-full" />
            </div>

            {error && <p className="text-rempah-500 text-sm bg-rempah-500/5 border border-rempah-500/10 rounded-lg px-3 py-2" role="alert">{error}</p>}

            <button type="submit" className="w-full bg-kuali-950 text-white py-3 rounded-xl font-medium hover:bg-kuali-700 transition-colors text-sm tracking-wide">
              {mode === "login" ? "Log in" : "Create account"}
            </button>
          </form>

          <p className="text-center text-sm text-kuali-700/50">
            {mode === "login" ? "No account yet?" : "Already have an account?"}{" "}
            <button type="button" onClick={() => { setMode(mode === "login" ? "register" : "login"); setError(""); }} className="text-rempah-500 font-medium hover:text-rempah-700 transition-colors">
              {mode === "login" ? "Sign up" : "Log in"}
            </button>
          </p>

          <p className="text-center">
            <Link href="/" className="text-xs text-kuali-700/30 hover:text-kuali-700/50 transition-colors">← Back to home</Link>
          </p>
        </div>
      </div>
    </div>
  );
}
