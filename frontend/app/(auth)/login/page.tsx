"use client";

import Link from "next/link";
import { useState, Suspense } from "react";
import { motion } from "framer-motion";
import { Mail, Lock, ArrowRight } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Checkbox } from "@/components/ui/toggle";
import { Logo } from "@/components/ui/logo";
import { LanguageToggle } from "@/components/ui/language-toggle";
import { useLanguage } from "@/components/providers/language-provider";
import { useStore } from "@/lib/store";
import { useRouter } from "next/navigation";
import { toast } from "sonner";

export default function LoginPage() {
  return (
    <Suspense fallback={null}>
      <LoginForm />
    </Suspense>
  );
}

function LoginForm() {
  const { t } = useLanguage();
  const router = useRouter();
  const signIn = useStore((s) => s.signIn);
  const [email, setEmail] = useState("demo@dapurpintar.ai");
  const [password, setPassword] = useState("demo1234");
  const [remember, setRemember] = useState(true);
  const [loading, setLoading] = useState(false);

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    await new Promise((r) => setTimeout(r, 600));
    signIn();
    setLoading(false);
    toast.success("Welcome back.");
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
              {t.auth.signin.title}
            </h1>
            <p className="mt-2 text-sm text-text-muted">{t.auth.signin.subtitle}</p>

            <form onSubmit={onSubmit} className="mt-8 space-y-4">
              <div>
                <label className="text-xs font-medium text-text-secondary mb-1.5 block">
                  {t.auth.signin.email}
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
                  {t.auth.signin.password}
                </label>
                <Input
                  type="password"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  leadingIcon={<Lock className="h-4 w-4" />}
                />
              </div>
              <div className="flex items-center justify-between">
                <Checkbox
                  checked={remember}
                  onCheckedChange={setRemember}
                  label="Remember me"
                />
                <Link href="#" className="text-xs text-accent-active hover:text-text-primary">
                  {t.auth.signin.forgot}
                </Link>
              </div>
              <Button type="submit" className="w-full" loading={loading} size="md">
                {t.auth.signin.submit}
                <ArrowRight className="h-4 w-4" />
              </Button>
            </form>

            <div className="my-6 flex items-center gap-3 text-xs text-text-muted">
              <div className="flex-1 h-px bg-border" />
              {t.auth.signin.or}
              <div className="flex-1 h-px bg-border" />
            </div>

            <div className="grid grid-cols-2 gap-2">
              <Button variant="outline" type="button">
                <svg viewBox="0 0 24 24" className="h-4 w-4" aria-hidden>
                  <path fill="#FFC107" d="M21.8 10.2H12v3.9h5.6c-.2 1.4-1.7 4.1-5.6 4.1-3.4 0-6.1-2.8-6.1-6.2S8.6 5.8 12 5.8c1.9 0 3.2.8 4 1.5l2.7-2.6C16.9 3.2 14.6 2.2 12 2.2 6.5 2.2 2 6.7 2 12.2s4.5 10 10 10c5.8 0 9.6-4 9.6-9.7 0-.6-.1-1.1-.2-1.5z" />
                </svg>
                {t.auth.signin.google}
              </Button>
              <Button variant="outline" type="button">
                <svg viewBox="0 0 24 24" className="h-4 w-4" fill="currentColor" aria-hidden>
                  <path d="M12 .3a12 12 0 0 0-3.8 23.4c.6.1.8-.3.8-.6v-2.2c-3.3.7-4-1.6-4-1.6-.6-1.4-1.4-1.8-1.4-1.8-1.1-.7.1-.7.1-.7 1.2.1 1.9 1.3 1.9 1.3 1.1 1.8 2.8 1.3 3.6 1 .1-.8.4-1.3.8-1.6-2.7-.3-5.5-1.3-5.5-6 0-1.3.5-2.4 1.2-3.2-.1-.3-.5-1.5.1-3.2 0 0 1-.3 3.3 1.2a11.5 11.5 0 0 1 6 0c2.3-1.5 3.3-1.2 3.3-1.2.7 1.7.3 2.9.1 3.2.8.8 1.2 1.9 1.2 3.2 0 4.7-2.9 5.7-5.6 6 .5.4.8 1.1.8 2.3v3.3c0 .3.2.7.8.6A12 12 0 0 0 12 .3" />
                </svg>
                {t.auth.signin.github}
              </Button>
            </div>

            <p className="mt-6 text-sm text-text-muted text-center">
              {t.auth.signin.noAccount}{" "}
              <Link href="/register" className="text-accent-active hover:text-text-primary font-medium">
                {t.auth.signin.signup}
              </Link>
            </p>
          </motion.div>
        </div>
      </div>

      <div className="hidden lg:flex flex-1 relative items-center justify-center bg-bg-section p-12 border-l border-border overflow-hidden">
        <div className="relative max-w-md text-center">
          <p className="font-serif italic text-2xl text-text-primary leading-[1.4]">
            &ldquo;The quiet work of a kitchen is some of the most important work there is.&rdquo;
          </p>
          <p className="mt-6 text-sm text-text-muted uppercase tracking-[0.12em]">
            Welcome back
          </p>
        </div>
      </div>
    </div>
  );
}
