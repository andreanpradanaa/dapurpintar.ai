"use client";
import { useState } from "react";
import Link from "next/link";
import { motion } from "framer-motion";
import {
  ChefHat, Search, ShoppingBasket, CalendarDays, Bot,
  ArrowRight, Shield, CheckCircle2, Star, Sparkles,
  Leaf, Clock, Users, ChevronDown, Menu, X, Globe,
} from "lucide-react";
import { api, type Recipe } from "@/lib/api";

const stagger = { container: { hidden: {}, show: { transition: { staggerChildren: 0.08 } } }, item: { hidden: { opacity: 0, y: 20 }, show: { opacity: 1, y: 0 } } };

function FeatureCard({ icon: Icon, title, desc }: { icon: React.ElementType; title: string; desc: string }) {
  return (
    <motion.div whileHover={{ y: -4 }} className="bg-surface rounded-2xl border border-line p-6 transition-shadow hover:shadow-lg hover:shadow-emerald-soft/50">
      <div className="w-10 h-10 rounded-xl bg-emerald-soft flex items-center justify-center mb-4"><Icon className="w-5 h-5 text-emerald" /></div>
      <h3 className="font-display text-lg font-semibold text-ink mb-2">{title}</h3>
      <p className="text-ink-muted text-sm leading-relaxed">{desc}</p>
    </motion.div>
  );
}

function DemoCard({ title, items }: { title: string; items: string[] }) {
  return (
    <motion.div whileHover={{ y: -2 }} className="bg-surface border border-line rounded-2xl p-6 transition-shadow hover:shadow-md">
      <p className="font-mono text-[10px] uppercase tracking-[0.15em] text-emerald mb-3">{title}</p>
      <div className="space-y-2">
        {items.map((s, i) => (
          <div key={i} className="flex items-center gap-2 text-sm text-ink">
            <CheckCircle2 className="w-4 h-4 text-emerald flex-shrink-0" />
            <span>{s}</span>
          </div>
        ))}
      </div>
    </motion.div>
  );
}

export default function LandingPage() {
  const [query, setQuery] = useState("");
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [loading, setLoading] = useState(false);
  const [chips, setChips] = useState<string[]>([]);
  const [hasSearched, setHasSearched] = useState(false);

  const handleSearch = () => {
    let current = chips;
    if (query.trim()) { current = [...chips, query.trim().replace(/,$/, "")]; setChips(current); setQuery(""); }
    const q = current.join(" ");
    if (!q.trim()) return;
    setHasSearched(true);
    setLoading(true);
    api.recipes(q).then(r => setRecipes(r.data.slice(0, 4))).catch(() => {}).finally(() => setLoading(false));
  };

  const removeChip = (i: number) => setChips(chips.filter((_, idx) => idx !== i));
  const addChip = (s: string) => setChips([...chips, s]);

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") { e.preventDefault(); handleSearch(); }
    if (e.key === ",") { e.preventDefault(); if (query.trim()) { setChips([...chips, query.trim().replace(/,$/, "")]); setQuery(""); } }
    if (e.key === "Backspace" && !query && chips.length > 0) removeChip(chips.length - 1);
  };

  return (
    <div className="bg-canvas min-h-screen">
      {/* ---- Header ---- */}
      <header className="sticky top-0 z-50 bg-canvas/80 backdrop-blur-xl border-b border-line/50">
        <div className="max-w-6xl mx-auto px-6 h-16 flex items-center justify-between">
          <Link href="/" className="font-display text-xl font-bold text-ink tracking-tight">DapurPintar</Link>
          <div className="flex items-center gap-4">
            <Link href="/login" className="text-sm text-ink-muted hover:text-ink transition-colors">Sign in</Link>
            <Link href="/login" className="text-sm bg-ink text-surface px-4 py-2 rounded-full font-medium hover:bg-ink/90 transition-colors">Try free</Link>
          </div>
        </div>
      </header>

      {/* ---- 1. Hero ---- */}
      <section className="max-w-6xl mx-auto px-6 pt-20 pb-16 md:pt-32 md:pb-24">
        <motion.div initial="hidden" animate="show" variants={stagger.container} className="text-center max-w-3xl mx-auto space-y-6">
          <motion.p variants={stagger.item} className="font-mono text-[11px] uppercase tracking-[0.2em] text-emerald font-semibold">Your kitchen, thinking ahead</motion.p>
          <motion.h1 variants={stagger.item} className="text-hero-mobile md:text-hero font-display text-ink">
            Your AI kitchen assistant
          </motion.h1>
          <motion.p variants={stagger.item} className="text-lg text-ink-muted max-w-xl mx-auto leading-relaxed">
            Know what to cook, use what you have, and shop with intention. DapurPintar reads your pantry so you don&apos;t have to guess.
          </motion.p>
          <motion.div variants={stagger.item} className="flex gap-3 justify-center">
            <Link href="/login" className="bg-emerald text-surface px-6 py-3.5 rounded-full font-semibold text-sm hover:bg-emerald-deep transition-colors inline-flex items-center gap-2">Get started <ArrowRight className="w-4 h-4" /></Link>
            <a href="#demo" className="text-ink-muted text-sm font-medium px-6 py-3.5 rounded-full border border-line hover:border-ink/20 transition-colors">See how it works</a>
          </motion.div>
        </motion.div>

        {/* Ingredient composer + AI preview */}
        <motion.div initial={{ opacity: 0, y: 40 }} animate={{ opacity: 1, y: 0 }} transition={{ delay: 0.4 }} className="mt-16 max-w-3xl mx-auto">
          <div className="bg-surface rounded-3xl border border-line shadow-xl shadow-ink/3 overflow-hidden">
            {/* Input area */}
            <div className="p-4 md:p-6 border-b border-line">
              <div className="flex flex-wrap gap-2 mb-3 min-h-[28px]">
                {chips.map((c, i) => (
                  <motion.span initial={{ scale: 0 }} animate={{ scale: 1 }} key={i} className="inline-flex items-center gap-1 bg-emerald-soft text-emerald-deep border border-emerald/20 px-3 py-1 rounded-full text-sm font-medium">
                    {c} <button onClick={() => removeChip(i)} className="hover:text-ink transition-colors ml-1">×</button>
                  </motion.span>
                ))}
              </div>
              <div className="flex gap-3">
                <input type="text" value={query} onChange={e => setQuery(e.target.value)} onKeyDown={handleKeyDown}
                  placeholder="ayam, santan, bawang putih..." className="flex-1 !border-0 !ring-0 !px-1 !py-1 text-ink placeholder:text-ink-soft text-lg focus:!ring-0 bg-transparent" autoFocus />
                <button onClick={handleSearch} className="bg-emerald text-surface px-5 py-2.5 rounded-full text-sm font-semibold hover:bg-emerald-deep transition-colors flex items-center gap-2 flex-shrink-0">
                  <Search className="w-4 h-4" /> Find
                </button>
              </div>
            </div>

            {/* Quick add */}
            {chips.length === 0 && !query && (
              <div className="px-6 pb-6 flex flex-wrap gap-2">
                {["ayam", "tahu", "tempe", "telur", "santan", "kecap"].map(s => (
                  <button key={s} onClick={() => addChip(s)} className="text-xs text-ink-muted hover:text-emerald border border-line hover:border-emerald/30 rounded-full px-4 py-1.5 transition-colors">{s}</button>
                ))}
              </div>
            )}
          </div>

          {/* Results */}
          {(hasSearched && !loading && recipes.length === 0) && (
            <p className="text-center text-sm text-ink-muted mt-6">No recipes found for these ingredients. Try different ones.</p>
          )}
          {loading && (
            <div className="mt-6 grid md:grid-cols-2 gap-3">
              {[1,2].map(i => <div key={i} className="bg-surface border border-line rounded-2xl p-5 animate-pulse"><div className="h-5 bg-canvas-alt rounded w-3/4 mb-2" /><div className="h-4 bg-canvas-alt rounded w-1/2" /></div>)}
            </div>
          )}
          {!loading && recipes.length > 0 && (
            <motion.div initial={{ opacity: 0, y: 10 }} animate={{ opacity: 1, y: 0 }} className="mt-6 grid md:grid-cols-2 gap-3">
              {recipes.map(r => (
                <div key={r.id} className="bg-surface border border-line rounded-2xl p-5 hover:shadow-md hover:border-emerald/20 transition-all">
                  <div className="flex items-start gap-3">
                    <div className="w-10 h-10 rounded-xl bg-emerald-soft flex items-center justify-center flex-shrink-0"><ChefHat className="w-5 h-5 text-emerald" /></div>
                    <div>
                      <h3 className="font-display font-semibold text-ink">{r.title}</h3>
                      <p className="text-sm text-ink-muted mt-1 line-clamp-2">{r.summary}</p>
                      <div className="flex gap-4 mt-3 text-xs text-ink-soft">
                        <span className="flex items-center gap-1"><Users className="w-3 h-3" /> {r.servings}</span>
                        {r.prep_time_minutes && <span className="flex items-center gap-1"><Clock className="w-3 h-3" /> {r.prep_time_minutes}m</span>}
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </motion.div>
          )}
        </motion.div>
      </section>

      {/* ---- 2. Social proof ---- */}
      <section className="border-y border-line bg-surface py-8">
        <motion.div initial={{ opacity: 0 }} whileInView={{ opacity: 1 }} viewport={{ once: true }} className="max-w-6xl mx-auto px-6 text-center">
          <p className="text-xs font-mono uppercase tracking-[0.15em] text-ink-soft mb-4">The modern kitchen runs on</p>
          <div className="flex flex-wrap justify-center gap-8 text-ink-muted text-sm font-medium">
            <span className="flex items-center gap-2"><Sparkles className="w-4 h-4 text-yellow" /> AI-powered suggestions</span>
            <span className="flex items-center gap-2"><Leaf className="w-4 h-4 text-emerald" /> 12+ Indonesian recipes</span>
            <span className="flex items-center gap-2"><Shield className="w-4 h-4 text-ink-soft" /> Your data stays private</span>
          </div>
        </motion.div>
      </section>

      {/* ---- 3. Features ---- */}
      <section id="demo" className="max-w-6xl mx-auto px-6 py-section">
        <motion.div initial={{ opacity: 0, y: 24 }} whileInView={{ opacity: 1, y: 0 }} viewport={{ once: true }} className="text-center max-w-xl mx-auto mb-16 space-y-3">
          <p className="font-mono text-[11px] uppercase tracking-[0.2em] text-emerald font-semibold">One assistant, four decisions</p>
          <h2 className="text-display font-display text-ink">Everything your kitchen needs</h2>
          <p className="text-ink-muted">From knowing what to cook to completing your shopping, one intelligent assistant handles it.</p>
        </motion.div>

        <motion.div initial="hidden" whileInView="show" viewport={{ once: true }} variants={stagger.container} className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
          {[
            { icon: ChefHat, title: "Know what to cook", desc: "AI reads your pantry and suggests meals you can make right now — no guesswork." },
            { icon: Box, title: "Track your ingredients", desc: "Add what you have. We track expiry, suggest what to use first, and keep your pantry smart." },
            { icon: CalendarDays, title: "Plan your week", desc: "Build weekly meal plans with a drag-and-drop grid. Assign recipes to days in seconds." },
            { icon: ShoppingBasket, title: "Shop with purpose", desc: "Generate shopping lists from your meal plan. Activate, complete, and never over-buy again." },
          ].map((f, i) => <FeatureCard key={i} {...f} />)}
        </motion.div>
      </section>

      {/* ---- 4-7. Demo sections ---- */}
      {[
        { title: "AI-powered recommendations", subtitle: "Type what's in your kitchen. Get recipe suggestions instantly.", items: ["Soto Ayam — 20 min · 4 servings", "Nasi Goreng — 10 min · 2 servings", "Sate Ayam — 15 min · 4 servings"] },
        { title: "Smart pantry management", subtitle: "Track ingredients, get expiry alerts, and know what to use first.", items: ["Auto-categorize: bumbu, protein, sayur, pokok", "Expiry reminders at 3 days", "Running low alerts for quantity ≤ 1"] },
        { title: "Weekly meal planner", subtitle: "Plan 7 days of meals with a visual grid. Assign recipes to slots.", items: ["Drag-and-drop recipe assignment", "Breakfast, lunch, dinner, snack slots", "One-click complete or cancel"] },
        { title: "Shopping list with intent", subtitle: "Generate lists from meal plans, activate in the store, complete as you go.", items: ["Auto-generate from planned meals", "Toggle items complete while shopping", "Draft → activated → completed lifecycle"] },
      ].map((d, i) => (
        <section key={i} className={`${i % 2 === 0 ? 'bg-surface border-y border-line' : ''}`}>
          <div className="max-w-6xl mx-auto px-6 py-16 md:py-24">
            <motion.div initial={{ opacity: 0, y: 24 }} whileInView={{ opacity: 1, y: 0 }} viewport={{ once: true }} className="grid md:grid-cols-2 gap-12 items-center">
              <div className={i % 2 === 1 ? "md:order-2" : ""}>
                <p className="font-mono text-[11px] uppercase tracking-[0.2em] text-emerald font-semibold mb-3">{"0" + (i + 1)}</p>
                <h3 className="font-display text-2xl md:text-3xl font-bold text-ink mb-3">{d.title}</h3>
                <p className="text-ink-muted mb-6">{d.subtitle}</p>
                <DemoCard title="SAMPLE OUTPUT" items={d.items} />
              </div>
              <div className={`bg-canvas-alt rounded-3xl border border-line p-8 ${i % 2 === 1 ? "md:order-1" : ""}`}>
                <div className="space-y-3">
                  {d.items.map((s, j) => (
                    <div key={j} className="bg-surface rounded-xl border border-line p-4 flex items-center gap-3">
                      <div className="w-8 h-8 rounded-lg bg-emerald-soft flex items-center justify-center flex-shrink-0"><CheckCircle2 className="w-4 h-4 text-emerald" /></div>
                      <span className="text-sm text-ink font-medium">{s}</span>
                    </div>
                  ))}
                </div>
              </div>
            </motion.div>
          </div>
        </section>
      ))}

      {/* ---- 8. AI Chat Preview ---- */}
      <section className="max-w-6xl mx-auto px-6 py-section">
        <motion.div initial={{ opacity: 0, y: 24 }} whileInView={{ opacity: 1, y: 0 }} viewport={{ once: true }} className="text-center max-w-xl mx-auto mb-12 space-y-3">
          <p className="font-mono text-[11px] uppercase tracking-[0.2em] text-emerald font-semibold">Conversational AI</p>
          <h2 className="text-display font-display text-ink">Ask your kitchen anything</h2>
          <p className="text-ink-muted">Discuss substitutions, cooking times, or get step-by-step guidance — all in one conversation.</p>
        </motion.div>
        <div className="max-w-2xl mx-auto bg-surface border border-line rounded-3xl p-6 space-y-4 shadow-lg shadow-ink/3">
          {[
            { role: "user", content: "Can I substitute tofu for chicken in soto?" },
            { role: "ai", content: "Yes! Firm tofu works well as a substitute. Marinate it in turmeric and salt for 10 minutes before adding to the broth. The texture will be different but the flavor profile stays authentic." },
          ].map((m, i) => (
            <div key={i} className={`flex gap-3 ${m.role === "user" ? "justify-end" : ""}`}>
              {m.role === "ai" && <div className="w-8 h-8 rounded-full bg-emerald-soft flex items-center justify-center flex-shrink-0"><Bot className="w-4 h-4 text-emerald" /></div>}
              <div className={`max-w-[80%] rounded-2xl px-4 py-3 ${m.role === "user" ? "bg-emerald text-surface" : "bg-canvas-alt text-ink"}`}>
                <p className="text-sm">{m.content}</p>
              </div>
            </div>
          ))}
          <div className="flex gap-2 pt-2">
            <input type="text" placeholder="Ask about a recipe..." className="flex-1 bg-canvas-alt !border-0 !ring-0 rounded-full px-4 py-2.5 text-sm" readOnly />
            <button className="bg-emerald text-surface p-2.5 rounded-full"><ArrowRight className="w-4 h-4" /></button>
          </div>
        </div>
      </section>

      {/* ---- 9. Benefits ---- */}
      <section className="bg-surface border-y border-line">
        <div className="max-w-6xl mx-auto px-6 py-section grid md:grid-cols-3 gap-8">
          {[
            { icon: Shield, title: "Private by design", desc: "Your pantry data never leaves your account. We don't sell or share anything." },
            { icon: Sparkles, title: "AI that learns", desc: "The more you cook, the better our suggestions. Preferences adapt over time." },
            { icon: Globe, title: "Indonesian-first", desc: "Built for Indonesian home cooks with local ingredients, recipes, and measurements." },
          ].map((b, i) => (
            <motion.div key={i} initial={{ opacity: 0, y: 20 }} whileInView={{ opacity: 1, y: 0 }} viewport={{ once: true }} transition={{ delay: i * 0.1 }} className="text-center space-y-3">
              <div className="w-12 h-12 rounded-2xl bg-emerald-soft flex items-center justify-center mx-auto"><b.icon className="w-6 h-6 text-emerald" /></div>
              <h3 className="font-display font-semibold text-ink">{b.title}</h3>
              <p className="text-sm text-ink-muted">{b.desc}</p>
            </motion.div>
          ))}
        </div>
      </section>

      {/* ---- 10. Testimonials ---- */}
      <section className="max-w-6xl mx-auto px-6 py-section">
        <motion.div initial={{ opacity: 0 }} whileInView={{ opacity: 1 }} viewport={{ once: true }} className="text-center max-w-xl mx-auto mb-12 space-y-3">
          <p className="font-mono text-[11px] uppercase tracking-[0.2em] text-emerald font-semibold">Loved by home cooks</p>
          <h2 className="text-display font-display text-ink">What people are saying</h2>
        </motion.div>
        <div className="grid md:grid-cols-3 gap-6">
          {[
            { name: "Rina W.", text: "Finally, an app that understands Indonesian cooking. The pantry suggestions are spot-on.", stars: 5 },
            { name: "Budi S.", text: "Meal planning for the week used to take hours. Now it takes 5 minutes.", stars: 5 },
            { name: "Dian K.", text: "My husband was skeptical, but the AI recommendations changed how we do dinner every night.", stars: 5 },
          ].map((t, i) => (
            <motion.div key={i} initial={{ opacity: 0, y: 20 }} whileInView={{ opacity: 1, y: 0 }} viewport={{ once: true }} transition={{ delay: i * 0.1 }} className="bg-surface border border-line rounded-2xl p-6 space-y-4">
              <div className="flex gap-1">{Array(t.stars).fill(null).map((_, j) => <Star key={j} className="w-4 h-4 fill-yellow text-yellow" />)}</div>
              <p className="text-sm text-ink leading-relaxed">&ldquo;{t.text}&rdquo;</p>
              <p className="text-xs text-ink-muted font-medium">{t.name}</p>
            </motion.div>
          ))}
        </div>
      </section>

      {/* ---- 11. Pricing ---- */}
      <section className="bg-surface border-y border-line">
        <div className="max-w-6xl mx-auto px-6 py-section text-center space-y-4">
          <p className="font-mono text-[11px] uppercase tracking-[0.2em] text-emerald font-semibold">Pricing</p>
          <h2 className="text-display font-display text-ink">Free during early access</h2>
          <p className="text-ink-muted max-w-md mx-auto">We&apos;re building the best kitchen assistant. Join now and get lifetime early access at no cost.</p>
          <Link href="/login" className="inline-flex items-center gap-2 bg-emerald text-surface px-6 py-3 rounded-full font-semibold text-sm hover:bg-emerald-deep transition-colors">Start for free <ArrowRight className="w-4 h-4" /></Link>
        </div>
      </section>

      {/* ---- 12. FAQ ---- */}
      <section className="max-w-3xl mx-auto px-6 py-section">
        <motion.div initial={{ opacity: 0 }} whileInView={{ opacity: 1 }} viewport={{ once: true }} className="text-center mb-12 space-y-3">
          <p className="font-mono text-[11px] uppercase tracking-[0.2em] text-emerald font-semibold">FAQ</p>
          <h2 className="text-display font-display text-ink">Questions you might have</h2>
        </motion.div>
        <div className="space-y-3">
          {[
            { q: "Do I need an account?", a: "Only to save your pantry, meal plans, and preferences. You can try the landing page without signing up." },
            { q: "Is the AI free?", a: "Yes, during early access. When we introduce pricing, early users will keep their grandfathered plan." },
            { q: "What data do you store?", a: "Only what you add: ingredients, recipes, plans, and shopping lists. We never share or sell your data." },
            { q: "Can I use this on my phone?", a: "Yes. DapurPintar works on desktop, tablet, and mobile with the same experience." },
          ].map((faq, i) => (
            <details key={i} className="bg-surface border border-line rounded-2xl group">
              <summary className="flex items-center justify-between px-6 py-4 cursor-pointer text-ink font-medium text-sm list-none">
                {faq.q} <ChevronDown className="w-4 h-4 text-ink-soft group-open:rotate-180 transition-transform" />
              </summary>
              <p className="px-6 pb-4 text-sm text-ink-muted leading-relaxed">{faq.a}</p>
            </details>
          ))}
        </div>
      </section>

      {/* ---- 13. Final CTA ---- */}
      <section className="bg-ink rounded-[40px] mx-6 mb-16 max-w-4xl md:mx-auto overflow-hidden relative">
        <div className="absolute inset-0 bg-gradient-to-br from-emerald/20 to-transparent" />
        <div className="relative px-8 py-16 md:py-24 text-center space-y-6">
          <motion.div initial={{ opacity: 0, y: 20 }} whileInView={{ opacity: 1, y: 0 }} viewport={{ once: true }} className="space-y-4">
            <h2 className="text-display font-display text-surface">Ready to cook smarter?</h2>
            <p className="text-surface/60 max-w-md mx-auto">Join thousands of home cooks who let AI handle the thinking — so they can focus on the cooking.</p>
            <div className="flex gap-3 justify-center">
              <Link href="/login" className="bg-emerald text-surface px-8 py-3.5 rounded-full font-semibold text-sm hover:bg-emerald/90 transition-colors inline-flex items-center gap-2">Get started free <ArrowRight className="w-4 h-4" /></Link>
            </div>
          </motion.div>
        </div>
      </section>

      {/* ---- 14. Footer ---- */}
      <footer className="border-t border-line bg-surface px-6 py-12">
        <div className="max-w-6xl mx-auto flex flex-col md:flex-row justify-between items-center gap-6">
          <div className="flex items-center gap-6 text-sm text-ink-muted">
            <Link href="/" className="font-display font-bold text-ink">DapurPintar</Link>
            <span>© 2026</span>
          </div>
          <div className="flex items-center gap-6 text-sm text-ink-muted">
            <Link href="/login" className="hover:text-ink transition-colors">Sign in</Link>
            <a href="#" className="hover:text-ink transition-colors">Privacy</a>
            <a href="#" className="hover:text-ink transition-colors">Terms</a>
          </div>
        </div>
      </footer>
    </div>
  );
}

function Box(props: React.ComponentProps<"svg">) {
  return (
    <svg {...props} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z" />
      <polyline points="3.27 6.96 12 12.01 20.73 6.96" /><line x1="12" y1="22.08" x2="12" y2="12" />
    </svg>
  );
}
