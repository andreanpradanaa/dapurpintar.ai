"use client";

import Link from "next/link";
import { useState, Suspense } from "react";
import { motion } from "framer-motion";
import { Mail, Lock, User, ArrowRight, Check } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Logo } from "@/components/ui/logo";
import { LanguageToggle } from "@/components/ui/language-toggle";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { useRouter, useSearchParams } from "next/navigation";
import { toast } from "sonner";

export default function RegisterPage() {
  return (
    <Suspense fallback={null}>
      <RegisterForm />
    </Suspense>
  );
}

function RegisterForm() {
  const { t } = useLanguage();
  const router = useRouter();
  const params = useSearchParams();
  const signIn = useStore((s) => s.signIn);
  const updateUser = useStore((s) => s.updateUser);
  const [name, setName] = useState("Andini Pradipta");
  const [email, setEmail] = useState("andini@dapurpintar.ai");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    await new Promise((r) => setTimeout(r, 700));
    signIn({ name, email });
    updateUser({ plan: params.get("plan") === "pro" ? "pro" : "free" });
    setLoading(false);
    toast.success("Welcome to Dapur Pintar.");
    router.push("/dashboard");
  };

  return (
    <div className="min-h-screen flex flex-col lg:flex-row">
      <div className="flex-1 flex flex-col p-6 sm:p-10">
        <div className="flex items-center justify-between">
          <Link href="/" className="inline-flex">
            <Logo size="md" />
          </Link>
          <LanguageToggle />
        </div>

        <div className="flex-1 flex items-center justify-center">
          <motion.div
            initial={{ opacity: 0, y: 12 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, ease: [0.22, 1, 0.36, 1] }}
            className="w-full max-w-sm"
          >
            <h1 className="font-display text-3xl sm:text-4xl font-medium tracking-tight text-text-primary text-balance">
              {t.auth.signup.title}
            </h1>
            <p className="mt-2 text-sm text-text-muted">{t.auth.signup.subtitle}</p>

            <form onSubmit={onSubmit} className="mt-8 space-y-4">
              <div>
                <label className="text-xs font-medium text-text-secondary mb-1.5 block">
                  {t.auth.signup.name}
                </label>
                <Input
                  required
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  leadingIcon={<User className="h-4 w-4" />}
                />
              </div>
              <div>
                <label className="text-xs font-medium text-text-secondary mb-1.5 block">
                  {t.auth.signup.email}
                </label>
                <Input
                  type="email"
                  required
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  leadingIcon={<Mail className="h-4 w-4" />}
                />
              </div>
              <div>
                <label className="text-xs font-medium text-text-secondary mb-1.5 block">
                  {t.auth.signup.password}
                </label>
                <Input
                  type="password"
                  required
                  minLength={6}
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  leadingIcon={<Lock className="h-4 w-4" />}
                />
              </div>
              <Button type="submit" className="w-full" loading={loading} size="md">
                {t.auth.signup.submit}
                <ArrowRight className="h-4 w-4" />
              </Button>
            </form>

            <ul className="mt-5 space-y-1.5 text-xs text-text-muted">
              {["Unlimited recipes", "Save favorites", "Personalized suggestions"].map((b) => (
                <li key={b} className="flex items-center gap-2">
                  <Check className="h-3 w-3 text-accent-active" strokeWidth={3} />
                  {b}
                </li>
              ))}
            </ul>

            <p className="mt-6 text-xs text-text-muted text-center">{t.auth.signup.terms}</p>
            <p className="mt-2 text-sm text-text-muted text-center">
              {t.auth.signup.haveAccount}{" "}
              <Link href="/login" className="text-accent-active hover:text-text-primary font-medium">
                {t.auth.signup.signin}
              </Link>
            </p>
          </motion.div>
        </div>
      </div>

      <div className="hidden lg:flex flex-1 relative items-center justify-center bg-bg-section p-12 border-l border-border overflow-hidden">
        <div className="relative max-w-md text-center">
          <p className="font-serif italic text-2xl text-text-primary leading-[1.4]">
            &ldquo;Cooking is not about being fancy. It&apos;s about paying attention.&rdquo;
          </p>
          <p className="mt-6 text-sm text-text-muted uppercase tracking-[0.12em]">
            Start your kitchen
          </p>
        </div>
      </div>
    </div>
  );
}
