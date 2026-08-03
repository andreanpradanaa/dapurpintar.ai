"use client";

import Link from "next/link";
import { motion, useScroll, useTransform } from "framer-motion";
import { useRef } from "react";
import { ArrowRight, Check, Sparkles, Plus, Clock, Users, Flame, Star, ChefHat } from "lucide-react";

const fadeUp = {
  hidden: { opacity: 0, y: 32 },
  visible: (delay: number = 0) => ({
    opacity: 1,
    y: 0,
    transition: { duration: 0.8, delay, ease: [0.16, 1, 0.3, 1] as const },
  }),
};

const stagger = {
  visible: { transition: { staggerChildren: 0.12 } },
};

export default function LandingPage() {
  const heroRef = useRef<HTMLDivElement>(null);
  const { scrollYProgress } = useScroll({
    target: heroRef,
    offset: ["start start", "end start"],
  });
  const heroY = useTransform(scrollYProgress, [0, 1], [0, 100]);
  const heroOpacity = useTransform(scrollYProgress, [0, 0.8], [1, 0]);

  return (
    <div className="bg-canvas min-h-screen">
      {/* Floating Navigation */}
      <header className="fixed top-6 left-1/2 -translate-x-1/2 z-50">
        <nav className="bg-surface/80 backdrop-blur-xl border border-mist/50 rounded-full px-2 py-2 shadow-product flex items-center gap-2">
          <Link href="/" className="font-display text-xl text-ink px-4">
            DapurPintar
          </Link>
          <div className="hidden md:flex items-center gap-1">
            <a href="#features" className="text-sm text-ink-muted hover:text-ink px-4 py-2 rounded-full transition-colors">Features</a>
            <a href="#product" className="text-sm text-ink-muted hover:text-ink px-4 py-2 rounded-full transition-colors">Product</a>
          </div>
          <div className="flex items-center gap-2">
            <Link href="/login" className="text-sm text-ink-muted hover:text-ink px-4 py-2 rounded-full transition-colors">
              Sign in
            </Link>
            <Link
              href="/login"
              className="bg-amber text-surface text-sm font-medium px-5 py-2 rounded-full hover:bg-amber-deep transition-colors"
            >
              Get started
            </Link>
          </div>
        </nav>
      </header>

      {/* Hero Section */}
      <section ref={heroRef} className="relative min-h-screen flex items-center pt-24 pb-20 overflow-hidden">
        <div
          className="absolute inset-0 pointer-events-none"
          style={{
            background: "radial-gradient(ellipse 80% 60% at 70% 20%, rgba(91, 123, 90, 0.06) 0%, transparent 60%), radial-gradient(ellipse 60% 40% at 20% 80%, rgba(232, 145, 67, 0.04) 0%, transparent 50%)",
          }}
        />

        <motion.div style={{ y: heroY, opacity: heroOpacity }} className="relative max-w-[1400px] mx-auto px-6 w-full">
          <div className="grid lg:grid-cols-[42%_1fr] gap-12 lg:gap-20 items-center">
            {/* Left — Narrative */}
            <motion.div variants={stagger} initial="hidden" animate="visible" className="space-y-8">
              <motion.p variants={fadeUp} className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em]">
                AI Kitchen Assistant
              </motion.p>

              <motion.h1 variants={fadeUp} custom={0.1} className="font-display text-hero text-ink text-balance">
                Your kitchen
                <br />
                <span className="italic text-sage">finally</span> has
                <br />a brain.
              </motion.h1>

              <motion.p variants={fadeUp} custom={0.2} className="text-body-xl text-ink-muted max-w-md leading-relaxed">
                DapurPintar reads your pantry, remembers what you have, and tells you exactly what to cook tonight — before anything expires.
              </motion.p>

              <motion.div variants={fadeUp} custom={0.3} className="flex flex-wrap items-center gap-4">
                <Link
                  href="/login"
                  className="group bg-amber text-surface px-8 py-4 rounded-full font-semibold text-sm inline-flex items-center gap-3 hover:bg-amber-deep transition-all shadow-glow-amber"
                >
                  Start cooking smarter
                  <ArrowRight className="w-4 h-4 group-hover:translate-x-1 transition-transform" />
                </Link>
                <a
                  href="#features"
                  className="text-sm text-ink-muted font-medium px-6 py-4 rounded-full border border-mist hover:border-ink/20 transition-colors"
                >
                  See how it works
                </a>
              </motion.div>

              <motion.div variants={fadeUp} custom={0.4} className="flex items-center gap-3 pt-4">
                <div className="flex -space-x-1">
                  {[0, 1, 2, 3, 4].map((i) => (
                    <Star key={i} className="w-4 h-4 fill-amber text-amber" />
                  ))}
                </div>
                <span className="text-data text-ink-soft">
                  Trusted by <span className="text-ink font-medium">2,000+</span> home cooks
                </span>
              </motion.div>
            </motion.div>

            {/* Right — Product Preview */}
            <motion.div
              initial={{ opacity: 0, y: 60, scale: 0.95 }}
              animate={{ opacity: 1, y: 0, scale: 1 }}
              transition={{ duration: 1, delay: 0.3, ease: [0.16, 1, 0.3, 1] }}
              className="relative"
            >
              <div className="absolute -inset-8 bg-gradient-to-br from-sage-soft/30 via-transparent to-amber-soft/20 rounded-[40px] blur-3xl" />

              <motion.div
                animate={{ y: [0, -8, 0] }}
                transition={{ duration: 6, repeat: Infinity, ease: "easeInOut" }}
                className="relative bg-surface rounded-4xl border border-mist shadow-product-xl overflow-hidden"
              >
                {/* Window Chrome */}
                <div className="flex items-center gap-2 px-6 py-4 border-b border-mist bg-canvas/50">
                  <div className="flex gap-1.5">
                    <div className="w-3 h-3 rounded-full bg-red-400/80" />
                    <div className="w-3 h-3 rounded-full bg-yellow-400/80" />
                    <div className="w-3 h-3 rounded-full bg-green-400/80" />
                  </div>
                  <div className="flex-1 text-center">
                    <span className="text-[11px] font-mono text-ink-soft">dapurpintar.ai/today</span>
                  </div>
                </div>

                {/* Dashboard Content */}
                <div className="p-6 space-y-5">
                  {/* Greeting */}
                  <div>
                    <p className="font-display text-2xl text-ink">Good evening, Sarah</p>
                    <p className="text-data text-ink-muted mt-1">3 ingredients expiring soon · 5 recipes you can make now</p>
                  </div>

                  {/* Pantry Chips */}
                  <div>
                    <div className="flex items-center justify-between mb-3">
                      <span className="font-mono text-[10px] uppercase tracking-[0.15em] text-sage font-semibold">Your Pantry</span>
                      <span className="text-[10px] font-mono text-ink-soft">18 items</span>
                    </div>
                    <div className="flex flex-wrap gap-2">
                      {[
                        { name: "Ayam", status: "fresh" },
                        { name: "Santan", status: "fresh" },
                        { name: "Bawang", status: "fresh" },
                        { name: "Telur", status: "low" },
                        { name: "Tempe", status: "expiring" },
                      ].map((item) => (
                        <span
                          key={item.name}
                          className="inline-flex items-center gap-1.5 bg-canvas border border-mist rounded-full pl-3 pr-3 py-1.5 text-data font-medium text-ink"
                        >
                          <span
                            className={`w-1.5 h-1.5 rounded-full ${
                              item.status === "fresh" ? "bg-sage" : item.status === "low" ? "bg-amber" : "bg-red-400"
                            }`}
                          />
                          {item.name}
                        </span>
                      ))}
                      <span className="inline-flex items-center gap-1 text-sage text-data font-medium pl-2">
                        <Plus className="w-3 h-3" /> 13 more
                      </span>
                    </div>
                  </div>

                  {/* AI Suggestion Card */}
                  <div className="bg-gradient-to-br from-sage-soft/50 to-sage-soft/20 rounded-2xl p-5 border border-sage/10">
                    <div className="flex items-center gap-2 mb-3">
                      <div className="w-6 h-6 rounded-full bg-sage flex items-center justify-center">
                        <Sparkles className="w-3 h-3 text-white" />
                      </div>
                      <span className="font-mono text-[10px] uppercase tracking-[0.15em] text-sage font-semibold">
                        AI Suggestion
                      </span>
                    </div>
                    <div className="flex items-start gap-4">
                      <div className="w-16 h-16 rounded-xl bg-gradient-to-br from-amber-soft to-amber-soft/50 flex items-center justify-center flex-shrink-0">
                        <ChefHat className="w-7 h-7 text-amber" />
                      </div>
                      <div className="flex-1 min-w-0">
                        <p className="font-display text-xl text-ink">Soto Ayam</p>
                        <p className="text-data text-ink-muted mt-0.5">Uses 4 of your ingredients</p>
                        <div className="flex items-center gap-3 mt-2">
                          <span className="inline-flex items-center gap-1 text-[10px] font-mono text-ink-soft">
                            <Clock className="w-3 h-3" /> 25 min
                          </span>
                          <span className="inline-flex items-center gap-1 text-[10px] font-mono text-ink-soft">
                            <Users className="w-3 h-3" /> 4 servings
                          </span>
                          <span className="text-[10px] font-mono text-sage font-bold">96% match</span>
                        </div>
                      </div>
                    </div>
                  </div>

                  {/* Quick Stats */}
                  <div className="grid grid-cols-3 gap-3">
                    {[
                      { label: "This week", value: "12 meals", sub: "planned" },
                      { label: "Saved", value: "Rp 85k", sub: "this month" },
                      { label: "Waste", value: "-62%", sub: "vs. last month" },
                    ].map((stat) => (
                      <div key={stat.label} className="bg-canvas rounded-xl p-3 border border-mist">
                        <p className="text-[10px] font-mono uppercase tracking-wider text-ink-soft">{stat.label}</p>
                        <p className="font-display text-lg text-ink mt-0.5">{stat.value}</p>
                        <p className="text-[10px] text-ink-soft">{stat.sub}</p>
                      </div>
                    ))}
                  </div>
                </div>
              </motion.div>
            </motion.div>
          </div>
        </motion.div>
      </section>

      {/* Problem / Editorial Statement */}
      <section className="py-section">
        <motion.div
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-100px" }}
          variants={stagger}
          className="max-w-[1200px] mx-auto px-6"
        >
          <motion.p variants={fadeUp} className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em] mb-8">
            The invisible problem
          </motion.p>

          <div className="grid md:grid-cols-3 gap-8 md:gap-0">
            {[
              { number: "1.3B", unit: "tonnes", desc: "of food wasted globally every year" },
              { number: "48M", unit: "rupiah", desc: "per household lost to expired ingredients in Indonesia" },
              { number: "37%", unit: "of groceries", desc: "thrown away before being cooked" },
            ].map((stat, i) => (
              <motion.div
                key={stat.number}
                variants={fadeUp}
                custom={i * 0.15}
                className={`py-8 ${i < 2 ? "md:border-r md:border-mist md:pr-8" : ""} ${i > 0 ? "md:pl-8" : ""}`}
              >
                <p className="font-display text-[clamp(3rem,6vw,5rem)] text-ink leading-none">
                  {stat.number}
                </p>
                <p className="font-mono text-data uppercase tracking-wider text-sage mt-2">{stat.unit}</p>
                <p className="text-body-lg text-ink-muted mt-3 max-w-xs">{stat.desc}</p>
              </motion.div>
            ))}
          </div>

          <motion.div
            variants={fadeUp}
            custom={0.5}
            className="mt-16 pt-12 border-t border-mist"
          >
            <p className="font-display text-display-sm text-ink max-w-2xl text-balance">
              You don&apos;t need another recipe app.
              <br />
              <span className="text-ink-muted">You need something that knows what you already have.</span>
            </p>
          </motion.div>
        </motion.div>
      </section>

      {/* Solution — Dark Sage Band */}
      <section className="bg-sage-deep py-section-sm relative overflow-hidden">
        <div
          className="absolute inset-0 pointer-events-none"
          style={{
            background: "radial-gradient(ellipse 60% 50% at 80% 50%, rgba(168, 196, 166, 0.08) 0%, transparent 60%)",
          }}
        />

        <motion.div
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-80px" }}
          variants={stagger}
          className="relative max-w-[1200px] mx-auto px-6 text-center"
        >
          <motion.h2 variants={fadeUp} className="font-display text-display text-white/95 text-balance">
            DapurPintar knows your kitchen
            <br />
            <span className="italic text-sage-glow">so you don&apos;t have to think.</span>
          </motion.h2>

          <motion.div variants={fadeUp} custom={0.2} className="mt-16 flex flex-col md:flex-row items-center justify-center gap-6 md:gap-4">
            {[
              { label: "See what you have", desc: "Pantry awareness" },
              { label: "Get a smart suggestion", desc: "AI recommendations" },
              { label: "Cook tonight", desc: "Meal planning" },
            ].map((step, i) => (
              <div key={step.label} className="flex items-center gap-4">
                <div className="text-left">
                  <p className="font-display text-xl text-white">{step.label}</p>
                  <p className="text-data text-white/50 mt-0.5">{step.desc}</p>
                </div>
                {i < 2 && <ArrowRight className="w-5 h-5 text-white/30 hidden md:block" />}
              </div>
            ))}
          </motion.div>
        </motion.div>
      </section>

      {/* Features */}
      <section id="features" className="py-section">
        {/* Feature 1 — AI Recommendations */}
        <FeatureSection
          number="01"
          eyebrow="AI Recommendations"
          title={<>Type what&apos;s in your kitchen.<br /><span className="italic text-sage">Get recipes you can make right now.</span></>}
          description="The AI considers your ingredients, expiry dates, and cooking time — then suggests meals you can cook tonight. No random recipes. No grocery runs."
          mockup={<AIRecommendationMockup />}
          reversed={false}
        />

        {/* Feature 2 — Smart Pantry */}
        <FeatureSection
          number="02"
          eyebrow="Smart Pantry"
          title={<>Track every ingredient<br /><span className="italic text-sage">with intelligence.</span></>}
          description="Auto-categorize by type, track expiry dates, and get alerts before anything goes bad. Your pantry remembers everything so you don't have to."
          mockup={<PantryMockup />}
          reversed={true}
        />

        {/* Feature 3 — Meal Planner */}
        <FeatureSection
          number="03"
          eyebrow="Weekly Planner"
          title={<>Plan seven days of meals<br /><span className="italic text-sage">in a visual grid.</span></>}
          description="Drag recommended meals into a weekly calendar. Breakfast, lunch, dinner — assigned to days in seconds. Leave the mental load to the AI."
          mockup={<PlannerMockup />}
          reversed={false}
        />

        {/* Feature 4 — Shopping */}
        <FeatureSection
          number="04"
          eyebrow="Shopping List"
          title={<>What you don&apos;t have becomes<br /><span className="italic text-sage">what you need to buy.</span></>}
          description="Generate shopping lists from your meal plans. Check off as you shop. Never forget an ingredient again."
          mockup={<ShoppingMockup />}
          reversed={true}
        />
      </section>

      {/* Product Canvas — Full Dashboard */}
      <section id="product" className="py-section bg-canvas-alt border-y border-mist">
        <motion.div
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-100px" }}
          variants={stagger}
          className="max-w-[1400px] mx-auto px-6"
        >
          <motion.div variants={fadeUp} className="text-center mb-16">
            <p className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em] mb-4">
              The complete kitchen
            </p>
            <h2 className="font-display text-display text-ink text-balance">
              One screen.
              <br />
              One assistant.
            </h2>
          </motion.div>

          <motion.div
            variants={fadeUp}
            custom={0.2}
            className="relative max-w-6xl mx-auto"
          >
            <div className="absolute -inset-12 bg-gradient-to-b from-sage-glow/10 via-transparent to-transparent rounded-[60px] blur-3xl" />

            <div className="relative bg-surface rounded-4xl border border-mist shadow-product-xl overflow-hidden">
              {/* Window Chrome */}
              <div className="flex items-center gap-2 px-6 py-4 border-b border-mist bg-canvas/50">
                <div className="flex gap-1.5">
                  <div className="w-3 h-3 rounded-full bg-red-400/80" />
                  <div className="w-3 h-3 rounded-full bg-yellow-400/80" />
                  <div className="w-3 h-3 rounded-full bg-green-400/80" />
                </div>
                <div className="flex-1 text-center">
                  <span className="text-[11px] font-mono text-ink-soft">dapurpintar.ai/today</span>
                </div>
              </div>

              {/* Dashboard Grid */}
              <div className="p-6 md:p-8">
                <div className="grid grid-cols-12 gap-5">
                  {/* Pantry Panel */}
                  <div className="col-span-12 md:col-span-3 bg-canvas rounded-2xl p-4 border border-mist">
                    <div className="flex items-center justify-between mb-4">
                      <span className="font-mono text-[10px] uppercase tracking-[0.15em] text-sage font-semibold">Pantry</span>
                      <span className="text-[10px] font-mono text-ink-soft">18 items</span>
                    </div>
                    <div className="space-y-2">
                      {[
                        { name: "Beras", qty: "2 kg", status: "fresh" },
                        { name: "Ayam", qty: "500g", status: "fresh" },
                        { name: "Santan", qty: "200ml", status: "fresh" },
                        { name: "Telur", qty: "4 butir", status: "low" },
                        { name: "Tempe", qty: "250g", status: "expiring" },
                      ].map((item) => (
                        <div key={item.name} className="flex items-center justify-between py-1.5">
                          <div className="flex items-center gap-2">
                            <span className={`w-1.5 h-1.5 rounded-full ${item.status === "fresh" ? "bg-sage" : item.status === "low" ? "bg-amber" : "bg-red-400"}`} />
                            <span className="text-data text-ink">{item.name}</span>
                          </div>
                          <span className="text-[10px] font-mono text-ink-soft">{item.qty}</span>
                        </div>
                      ))}
                    </div>
                  </div>

                  {/* AI Conversation */}
                  <div className="col-span-12 md:col-span-5 bg-canvas rounded-2xl p-4 border border-mist">
                    <div className="flex items-center gap-2 mb-4">
                      <div className="w-5 h-5 rounded-full bg-sage flex items-center justify-center">
                        <Sparkles className="w-2.5 h-2.5 text-white" />
                      </div>
                      <span className="font-mono text-[10px] uppercase tracking-[0.15em] text-sage font-semibold">AI Assistant</span>
                    </div>

                    <div className="space-y-3">
                      <div className="flex justify-end">
                        <div className="bg-sage text-white rounded-2xl rounded-tr-sm px-4 py-2.5 max-w-[85%]">
                          <p className="text-data">I have ayam, santan, bawang, and telur. What should I cook?</p>
                        </div>
                      </div>

                      <div className="flex gap-2">
                        <div className="w-6 h-6 rounded-full bg-sage/10 flex items-center justify-center flex-shrink-0 mt-0.5">
                          <Sparkles className="w-3 h-3 text-sage" />
                        </div>
                        <div className="bg-surface border border-mist rounded-2xl rounded-tl-sm px-4 py-3 flex-1">
                          <p className="text-data text-ink">Based on your pantry, I recommend <strong>Soto Ayam</strong>.</p>
                          <div className="mt-3 bg-canvas rounded-xl p-3 border border-mist">
                            <div className="flex items-center gap-2">
                              <ChefHat className="w-5 h-5 text-amber" />
                              <div>
                                <p className="text-data font-medium text-ink">Soto Ayam</p>
                                <div className="flex gap-2 mt-0.5">
                                  <span className="text-[10px] font-mono text-ink-soft">25 min</span>
                                  <span className="text-[10px] font-mono text-sage font-bold">96% match</span>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>

                  {/* Weekly Planner */}
                  <div className="col-span-12 md:col-span-4 bg-canvas rounded-2xl p-4 border border-mist">
                    <div className="flex items-center justify-between mb-4">
                      <span className="font-mono text-[10px] uppercase tracking-[0.15em] text-sage font-semibold">This Week</span>
                      <span className="text-[10px] font-mono text-ink-soft">7 meals planned</span>
                    </div>
                    <div className="space-y-2">
                      {[
                        { day: "Mon", meal: "Soto Ayam", time: "Dinner" },
                        { day: "Tue", meal: "Nasi Goreng", time: "Lunch" },
                        { day: "Wed", meal: "Rendang", time: "Dinner" },
                        { day: "Thu", meal: "Gado-gado", time: "Lunch" },
                        { day: "Fri", meal: "Sate Ayam", time: "Dinner" },
                      ].map((item) => (
                        <div key={item.day} className="flex items-center gap-3 py-1.5">
                          <span className="w-8 text-[10px] font-mono text-ink-soft">{item.day}</span>
                          <div className="flex-1 bg-surface border border-mist rounded-lg px-3 py-1.5">
                            <p className="text-data text-ink font-medium">{item.meal}</p>
                          </div>
                          <span className="text-[10px] font-mono text-ink-soft">{item.time}</span>
                        </div>
                      ))}
                    </div>
                  </div>

                  {/* Shopping List */}
                  <div className="col-span-12 bg-canvas rounded-2xl p-4 border border-mist">
                    <div className="flex items-center justify-between mb-4">
                      <span className="font-mono text-[10px] uppercase tracking-[0.15em] text-sage font-semibold">Shopping List</span>
                      <span className="text-[10px] font-mono text-ink-soft">3 of 8 checked</span>
                    </div>
                    <div className="flex flex-wrap gap-x-6 gap-y-2">
                      {[
                        { name: "Daun salam", done: true },
                        { name: "Serai", done: true },
                        { name: "Kunyit", done: true },
                        { name: "Lengkuas", done: false },
                        { name: "Kecap manis", done: false },
                        { name: "Bawang merah", done: false },
                        { name: "Cabai rawit", done: false },
                        { name: "Tomat", done: false },
                      ].map((item) => (
                        <div key={item.name} className="flex items-center gap-2">
                          <div className={`w-4 h-4 rounded border-2 flex items-center justify-center ${item.done ? "bg-sage border-sage" : "border-mist"}`}>
                            {item.done && <Check className="w-2.5 h-2.5 text-white" />}
                          </div>
                          <span className={`text-data ${item.done ? "line-through text-ink-soft" : "text-ink"}`}>{item.name}</span>
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
              </div>
            </div>

            {/* Floating Annotations */}
            <div className="hidden lg:block">
              {[
                { text: "AI suggests", x: "52%", y: "28%" },
                { text: "Expiry alert", x: "8%", y: "65%" },
                { text: "Drag to plan", x: "78%", y: "50%" },
                { text: "Auto-generated", x: "35%", y: "90%" },
              ].map((label) => (
                <div
                  key={label.text}
                  className="absolute bg-surface border border-mist rounded-full px-3 py-1 shadow-product text-[10px] font-mono text-ink-muted whitespace-nowrap"
                  style={{ left: label.x, top: label.y, transform: "translate(-50%, -50%)" }}
                >
                  {label.text}
                </div>
              ))}
            </div>
          </motion.div>

          <motion.div variants={fadeUp} custom={0.4} className="text-center mt-12">
            <Link
              href="/login"
              className="group bg-amber text-surface px-8 py-4 rounded-full font-semibold text-sm inline-flex items-center gap-3 hover:bg-amber-deep transition-all shadow-glow-amber"
            >
              Start your free pantry
              <ArrowRight className="w-4 h-4 group-hover:translate-x-1 transition-transform" />
            </Link>
          </motion.div>
        </motion.div>
      </section>

      {/* How It Works */}
      <section className="py-section">
        <motion.div
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-100px" }}
          variants={stagger}
          className="max-w-[1000px] mx-auto px-6"
        >
          <motion.div variants={fadeUp} className="text-center mb-20">
            <p className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em] mb-4">
              How it works
            </p>
            <h2 className="font-display text-display text-ink">Three steps. That&apos;s it.</h2>
          </motion.div>

          <div className="relative">
            {/* Timeline line */}
            <div className="absolute top-8 left-[10%] right-[10%] h-px bg-mist hidden md:block" />

            <div className="grid md:grid-cols-3 gap-12 md:gap-8">
              {[
                {
                  step: "1",
                  title: "Add ingredients",
                  desc: "Type what's in your kitchen. Your pantry becomes a living inventory.",
                },
                {
                  step: "2",
                  title: "Get suggestions",
                  desc: "The AI suggests meals you can cook right now, based on what you have.",
                },
                {
                  step: "3",
                  title: "Cook & repeat",
                  desc: "Rate what you make. The AI learns your taste. Suggestions get better.",
                },
              ].map((item, i) => (
                <motion.div key={item.step} variants={fadeUp} custom={i * 0.15} className="text-center relative">
                  <div className="w-16 h-16 rounded-full bg-surface border-2 border-sage flex items-center justify-center mx-auto mb-6 shadow-glow-sage relative z-10">
                    <span className="font-display text-2xl text-sage">{item.step}</span>
                  </div>
                  <h3 className="font-display text-display-sm text-ink mb-3">{item.title}</h3>
                  <p className="text-body-lg text-ink-muted max-w-xs mx-auto">{item.desc}</p>
                </motion.div>
              ))}
            </div>
          </div>
        </motion.div>
      </section>

      {/* Testimonials */}
      <section className="py-section bg-canvas">
        <motion.div
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-80px" }}
          variants={stagger}
          className="max-w-[1200px] mx-auto px-6"
        >
          <motion.p variants={fadeUp} className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em] mb-4 text-center">
            What home cooks are saying
          </motion.p>

          <div className="mt-16 grid md:grid-cols-[1.2fr_0.8fr_0.8fr] gap-5">
            {/* Featured */}
            <motion.div
              variants={fadeUp}
              custom={0.1}
              className="bg-surface border border-mist rounded-3xl p-8 shadow-product"
            >
              <div className="flex gap-1 mb-4">
                {[0, 1, 2, 3, 4].map((i) => (
                  <Star key={i} className="w-4 h-4 fill-amber text-amber" />
                ))}
              </div>
              <p className="font-display text-display-sm text-ink leading-snug text-balance">
                &ldquo;I used to throw away vegetables every week. Now DapurPintar reminds me to use them first. I&apos;ve cut my waste by 60%.&rdquo;
              </p>
              <div className="mt-8 flex items-center gap-3">
                <div className="w-10 h-10 rounded-full bg-sage-soft flex items-center justify-center">
                  <span className="font-display text-sage">R</span>
                </div>
                <div>
                  <p className="text-data font-medium text-ink">Rina, Bandung</p>
                  <p className="text-[10px] font-mono text-ink-soft">Home cook, 2 kids</p>
                </div>
              </div>
            </motion.div>

            {/* Side cards */}
            <motion.div
              variants={fadeUp}
              custom={0.2}
              className="bg-surface border border-mist rounded-3xl p-6 shadow-product"
            >
              <div className="flex gap-1 mb-4">
                {[0, 1, 2, 3, 4].map((i) => (
                  <Star key={i} className="w-3.5 h-3.5 fill-amber text-amber" />
                ))}
              </div>
              <p className="text-body-lg text-ink leading-snug">
                &ldquo;Meal planning used to take hours. Now it takes 5 minutes.&rdquo;
              </p>
              <div className="mt-6">
                <p className="text-data font-medium text-ink">Budi, Jakarta</p>
              </div>
            </motion.div>

            <motion.div
              variants={fadeUp}
              custom={0.3}
              className="bg-surface border border-mist rounded-3xl p-6 shadow-product"
            >
              <div className="flex gap-1 mb-4">
                {[0, 1, 2, 3, 4].map((i) => (
                  <Star key={i} className="w-3.5 h-3.5 fill-amber text-amber" />
                ))}
              </div>
              <p className="text-body-lg text-ink leading-snug">
                &ldquo;DapurPintar ngerti masakan Indonesia. Finally.&rdquo;
              </p>
              <div className="mt-6">
                <p className="text-data font-medium text-ink">Dian, Surabaya</p>
              </div>
            </motion.div>
          </div>
        </motion.div>
      </section>

      {/* Pricing */}
      <section className="py-section">
        <motion.div
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-80px" }}
          variants={stagger}
          className="max-w-lg mx-auto px-6 text-center"
        >
          <motion.p variants={fadeUp} className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em] mb-4">
            Pricing
          </motion.p>
          <motion.h2 variants={fadeUp} custom={0.1} className="font-display text-display text-ink">
            Early access.
            <br />
            <span className="italic">Free.</span>
          </motion.h2>

          <motion.div
            variants={fadeUp}
            custom={0.2}
            className="mt-12 relative"
          >
            <div className="absolute -inset-1 bg-gradient-to-br from-amber/20 via-transparent to-amber/20 rounded-4xl blur-xl" />

            <div className="relative bg-surface border border-mist rounded-3xl p-8 shadow-product-lg">
              <p className="font-display text-3xl text-ink">Early Access</p>
              <p className="font-display text-hero-mobile text-ink mt-2">
                Free
              </p>

              <div className="mt-8 space-y-3 text-left">
                {[
                  "Unlimited pantry tracking",
                  "AI-powered recommendations",
                  "Weekly meal planner",
                  "Smart shopping lists",
                  "AI conversation assistant",
                ].map((feature) => (
                  <div key={feature} className="flex items-center gap-3">
                    <div className="w-5 h-5 rounded-full bg-sage-soft flex items-center justify-center flex-shrink-0">
                      <Check className="w-3 h-3 text-sage" />
                    </div>
                    <span className="text-body-lg text-ink">{feature}</span>
                  </div>
                ))}
              </div>

              <Link
                href="/login"
                className="group mt-8 bg-amber text-surface w-full px-8 py-4 rounded-full font-semibold text-sm inline-flex items-center justify-center gap-3 hover:bg-amber-deep transition-all shadow-glow-amber"
              >
                Get early access
                <ArrowRight className="w-4 h-4 group-hover:translate-x-1 transition-transform" />
              </Link>
            </div>
          </motion.div>

          <motion.p variants={fadeUp} custom={0.4} className="mt-8 text-data text-ink-soft">
            Paid plans launching 2026. Early users keep their access — forever.
          </motion.p>
        </motion.div>
      </section>

      {/* FAQ */}
      <section className="py-section bg-canvas-alt border-y border-mist">
        <motion.div
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-80px" }}
          variants={stagger}
          className="max-w-3xl mx-auto px-6"
        >
          <motion.div variants={fadeUp} className="text-center mb-16">
            <p className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em] mb-4">
              Questions
            </p>
            <h2 className="font-display text-display text-ink">Things you might want to know</h2>
          </motion.div>

          <div className="space-y-0">
            {[
              {
                q: "Do I need an account?",
                a: "Yes — to save your pantry, meal plans, and preferences. The landing page is open to everyone.",
              },
              {
                q: "Is the AI really free?",
                a: "During early access, yes. When we introduce paid plans, early users are grandfathered in forever.",
              },
              {
                q: "What data do you store?",
                a: "Only what you add: ingredients, recipes, meal plans, and shopping lists. We never share or sell your data.",
              },
              {
                q: "Can I use this on my phone?",
                a: "Yes. DapurPintar works on desktop, tablet, and mobile — same experience everywhere.",
              },
              {
                q: "What makes this different from a recipe app?",
                a: "Recipe apps show you what you could make. DapurPintar tells you what you should make — based on what you already have.",
              },
            ].map((faq, i) => (
              <motion.div
                key={faq.q}
                variants={fadeUp}
                custom={i * 0.08}
                className="border-b border-mist"
              >
                <details className="group">
                  <summary className="flex items-center justify-between py-6 cursor-pointer list-none">
                    <span className="font-display text-display-sm text-ink">{faq.q}</span>
                    <Plus className="w-5 h-5 text-ink-soft group-open:rotate-45 transition-transform flex-shrink-0 ml-4" />
                  </summary>
                  <p className="pb-6 text-body-lg text-ink-muted leading-relaxed">{faq.a}</p>
                </details>
              </motion.div>
            ))}
          </div>
        </motion.div>
      </section>

      {/* Final CTA */}
      <section className="py-section">
        <motion.div
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-80px" }}
          variants={stagger}
          className="max-w-[900px] mx-auto px-6 text-center"
        >
          <motion.p variants={fadeUp} className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em] mb-4">
            Ready?
          </motion.p>
          <motion.h2 variants={fadeUp} custom={0.1} className="font-display text-hero text-ink text-balance">
            Your kitchen is
            <br />
            <span className="italic text-sage">waiting.</span>
          </motion.h2>
          <motion.p variants={fadeUp} custom={0.2} className="text-body-xl text-ink-muted max-w-md mx-auto mt-8">
            Join thousands of home cooks who let AI handle the thinking — so they can focus on the cooking.
          </motion.p>

          <motion.div variants={fadeUp} custom={0.3} className="mt-10">
            <Link
              href="/login"
              className="group bg-amber text-surface px-10 py-5 rounded-full font-semibold text-base inline-flex items-center gap-3 hover:bg-amber-deep transition-all shadow-glow-amber"
            >
              Start for free
              <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
            </Link>
          </motion.div>
        </motion.div>
      </section>

      {/* Footer */}
      <footer className="border-t border-mist bg-surface">
        <div className="max-w-[1200px] mx-auto px-6 py-12">
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-8">
            <div>
              <Link href="/" className="font-display text-2xl text-ink">DapurPintar</Link>
              <p className="text-data text-ink-soft mt-2">The AI kitchen assistant.</p>
            </div>

            <div className="flex flex-wrap gap-x-8 gap-y-2 text-sm text-ink-muted">
              <Link href="/login" className="hover:text-ink transition-colors">Sign in</Link>
              <a href="#" className="hover:text-ink transition-colors">Privacy</a>
              <a href="#" className="hover:text-ink transition-colors">Terms</a>
            </div>
          </div>

          <div className="mt-12 pt-8 border-t border-mist flex flex-col md:flex-row justify-between items-center gap-4">
            <p className="text-data text-ink-soft">&copy; 2026 DapurPintar AI. All rights reserved.</p>
            <p className="text-data text-ink-soft">Made for Indonesian home cooks.</p>
          </div>
        </div>
      </footer>
    </div>
  );
}

/* ────────────────────────────────────────────────────────────
   Feature Section Component
   ──────────────────────────────────────────────────────────── */

function FeatureSection({
  number,
  eyebrow,
  title,
  description,
  mockup,
  reversed,
}: {
  number: string;
  eyebrow: string;
  title: React.ReactNode;
  description: string;
  mockup: React.ReactNode;
  reversed: boolean;
}) {
  return (
    <div className={`max-w-[1400px] mx-auto px-6 ${reversed ? "py-section" : "py-section"} ${reversed ? "" : ""}`}>
      <motion.div
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-80px" }}
        variants={stagger}
        className={`grid lg:grid-cols-2 gap-12 lg:gap-20 items-center ${reversed ? "" : ""}`}
      >
        {/* Text */}
        <motion.div variants={fadeUp} className={`space-y-6 ${reversed ? "lg:order-2" : ""}`}>
          <p className="font-mono text-eyebrow uppercase text-sage font-semibold tracking-[0.2em]">
            {number} &middot; {eyebrow}
          </p>
          <h2 className="font-display text-display text-ink text-balance">
            {title}
          </h2>
          <p className="text-body-xl text-ink-muted max-w-md leading-relaxed">
            {description}
          </p>
          <Link
            href="/login"
            className="group inline-flex items-center gap-2 text-sm font-medium text-ink hover:text-sage transition-colors"
          >
            Try this feature
            <ArrowRight className="w-4 h-4 group-hover:translate-x-1 transition-transform" />
          </Link>
        </motion.div>

        {/* Mockup */}
        <motion.div
          variants={fadeUp}
          custom={0.2}
          className={`relative ${reversed ? "lg:order-1" : ""}`}
        >
          <div className="absolute -inset-6 bg-gradient-to-br from-sage-soft/30 via-transparent to-amber-soft/10 rounded-[40px] blur-3xl" />
          <div className="relative">
            {mockup}
          </div>
        </motion.div>
      </motion.div>
    </div>
  );
}

/* ────────────────────────────────────────────────────────────
   Mockup Components
   ──────────────────────────────────────────────────────────── */

function AIRecommendationMockup() {
  return (
    <div className="bg-surface rounded-4xl border border-mist shadow-product-xl overflow-hidden">
      <div className="flex items-center gap-2 px-6 py-4 border-b border-mist bg-canvas/50">
        <div className="flex gap-1.5">
          <div className="w-3 h-3 rounded-full bg-red-400/80" />
          <div className="w-3 h-3 rounded-full bg-yellow-400/80" />
          <div className="w-3 h-3 rounded-full bg-green-400/80" />
        </div>
        <span className="flex-1 text-center text-[11px] font-mono text-ink-soft">dapurpintar.ai/discover</span>
      </div>

      <div className="p-6 space-y-5">
        <div>
          <p className="font-mono text-[10px] uppercase tracking-[0.15em] text-sage font-semibold mb-3">Available ingredients</p>
          <div className="flex flex-wrap gap-2">
            {["Ayam", "Santan", "Bawang putih", "Kunyit", "Serai"].map((item) => (
              <span
                key={item}
                className="inline-flex items-center gap-1.5 bg-sage-soft border border-sage/20 rounded-full px-4 py-2 text-data font-medium text-sage-deep"
              >
                <span className="w-1.5 h-1.5 rounded-full bg-sage" />
                {item}
              </span>
            ))}
          </div>
        </div>

        <div className="border-t border-mist pt-5">
          <div className="flex items-center gap-2 mb-4">
            <div className="w-6 h-6 rounded-full bg-sage flex items-center justify-center">
              <Sparkles className="w-3 h-3 text-white" />
            </div>
            <span className="font-mono text-[10px] uppercase tracking-[0.15em] text-sage font-semibold">3 suggestions</span>
          </div>

          <div className="space-y-3">
            {[
              { name: "Soto Ayam", match: 96, time: 25, servings: 4, desc: "Uses 5 of your ingredients" },
              { name: "Opor Ayam", match: 88, time: 40, servings: 4, desc: "Uses 4 of your ingredients" },
              { name: "Ayam Goreng Bumbu Kuning", match: 82, time: 30, servings: 2, desc: "Uses 4 of your ingredients" },
            ].map((recipe, i) => (
              <div
                key={recipe.name}
                className={`flex items-center gap-4 p-4 rounded-2xl border transition-colors ${i === 0 ? "bg-sage-soft/30 border-sage/20" : "bg-canvas border-mist hover:border-sage/20"}`}
              >
                <div className="w-12 h-12 rounded-xl bg-gradient-to-br from-amber-soft to-amber-soft/50 flex items-center justify-center flex-shrink-0">
                  <ChefHat className="w-5 h-5 text-amber" />
                </div>
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2">
                    <p className="font-display text-lg text-ink">{recipe.name}</p>
                    {i === 0 && <Flame className="w-3.5 h-3.5 text-amber" />}
                  </div>
                  <p className="text-data text-ink-muted">{recipe.desc}</p>
                  <div className="flex items-center gap-3 mt-1.5">
                    <span className="inline-flex items-center gap-1 text-[10px] font-mono text-ink-soft">
                      <Clock className="w-3 h-3" /> {recipe.time} min
                    </span>
                    <span className="inline-flex items-center gap-1 text-[10px] font-mono text-ink-soft">
                      <Users className="w-3 h-3" /> {recipe.servings} servings
                    </span>
                    <span className="text-[10px] font-mono text-sage font-bold">{recipe.match}% match</span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}

function PantryMockup() {
  const items = [
    { name: "Beras", qty: "2 kg", category: "Pokok", status: "fresh" as const },
    { name: "Ayam kampung", qty: "500g", category: "Protein", status: "fresh" as const, expiring: "5 days" },
    { name: "Santan", qty: "200ml", category: "Bumbu", status: "expiring" as const, expiring: "2 days" },
    { name: "Telur", qty: "4 butir", category: "Protein", status: "low" as const },
    { name: "Bawang putih", qty: "8 siung", category: "Bumbu", status: "fresh" as const },
    { name: "Tempe", qty: "250g", category: "Protein", status: "expiring" as const, expiring: "1 day" },
    { name: "Cabai rawit", qty: "50g", category: "Bumbu", status: "fresh" as const },
  ];

  return (
    <div className="bg-surface rounded-4xl border border-mist shadow-product-xl overflow-hidden">
      <div className="flex items-center gap-2 px-6 py-4 border-b border-mist bg-canvas/50">
        <div className="flex gap-1.5">
          <div className="w-3 h-3 rounded-full bg-red-400/80" />
          <div className="w-3 h-3 rounded-full bg-yellow-400/80" />
          <div className="w-3 h-3 rounded-full bg-green-400/80" />
        </div>
        <span className="flex-1 text-center text-[11px] font-mono text-ink-soft">dapurpintar.ai/pantry</span>
      </div>

      <div className="p-6">
        <div className="flex items-center justify-between mb-5">
          <div>
            <p className="font-display text-2xl text-ink">My Pantry</p>
            <p className="text-data text-ink-muted mt-0.5">18 items · 2 expiring soon</p>
          </div>
          <button className="bg-sage text-white px-4 py-2 rounded-full text-data font-medium inline-flex items-center gap-1.5">
            <Plus className="w-3.5 h-3.5" /> Add item
          </button>
        </div>

        {/* Filter pills */}
        <div className="flex gap-2 mb-5">
          {["All", "Protein", "Bumbu", "Pokok"].map((filter, i) => (
            <span
              key={filter}
              className={`px-4 py-1.5 rounded-full text-data font-medium cursor-pointer transition-colors ${
                i === 0 ? "bg-sage text-white" : "bg-canvas text-ink-muted border border-mist hover:border-sage/30"
              }`}
            >
              {filter}
            </span>
          ))}
        </div>

        <div className="space-y-1">
          {items.map((item) => (
            <div
              key={item.name}
              className="flex items-center justify-between py-3 px-4 rounded-xl hover:bg-canvas/50 transition-colors"
            >
              <div className="flex items-center gap-3">
                <span
                  className={`w-2 h-2 rounded-full ${
                    item.status === "fresh" ? "bg-sage" : item.status === "low" ? "bg-amber" : "bg-red-400"
                  }`}
                />
                <div>
                  <p className="text-data-lg font-medium text-ink">{item.name}</p>
                  <p className="text-[10px] font-mono text-ink-soft">{item.qty} · {item.category}</p>
                </div>
              </div>
              <div className="flex items-center gap-3">
                {item.expiring && (
                  <span className="text-[10px] font-mono bg-amber-soft text-amber-deep px-2.5 py-1 rounded-full">
                    {item.expiring}
                  </span>
                )}
                {item.status === "low" && (
                  <span className="text-[10px] font-mono bg-amber-soft text-amber-deep px-2.5 py-1 rounded-full">
                    Running low
                  </span>
                )}
                {item.status === "fresh" && !item.expiring && (
                  <span className="w-1.5 h-1.5 rounded-full bg-sage" />
                )}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

function PlannerMockup() {
  const days = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"];
  const meals: Record<string, { lunch?: string; dinner?: string }> = {
    Mon: { dinner: "Soto Ayam" },
    Tue: { lunch: "Nasi Goreng" },
    Wed: { dinner: "Rendang" },
    Thu: { lunch: "Gado-gado" },
    Fri: { dinner: "Sate Ayam" },
    Sat: { lunch: "Soto Ayam", dinner: "Nasi Padang" },
    Sun: {},
  };

  return (
    <div className="bg-surface rounded-4xl border border-mist shadow-product-xl overflow-hidden">
      <div className="flex items-center gap-2 px-6 py-4 border-b border-mist bg-canvas/50">
        <div className="flex gap-1.5">
          <div className="w-3 h-3 rounded-full bg-red-400/80" />
          <div className="w-3 h-3 rounded-full bg-yellow-400/80" />
          <div className="w-3 h-3 rounded-full bg-green-400/80" />
        </div>
        <span className="flex-1 text-center text-[11px] font-mono text-ink-soft">dapurpintar.ai/planner</span>
      </div>

      <div className="p-6">
        <div className="flex items-center justify-between mb-6">
          <div>
            <p className="font-display text-2xl text-ink">This Week</p>
            <p className="text-data text-ink-muted mt-0.5">Jan 6 — Jan 12</p>
          </div>
          <button className="bg-sage text-white px-4 py-2 rounded-full text-data font-medium inline-flex items-center gap-1.5">
            <Sparkles className="w-3.5 h-3.5" /> Auto-fill
          </button>
        </div>

        <div className="space-y-2">
          {days.map((day) => {
            const dayMeals = meals[day] || {};
            return (
              <div key={day} className="flex items-center gap-3">
                <span className="w-10 text-data font-mono text-ink-soft font-medium">{day}</span>
                <div className="flex-1 grid grid-cols-2 gap-2">
                  <div className={`rounded-xl px-4 py-3 min-h-[48px] flex items-center ${dayMeals.lunch ? "bg-sage-soft/30 border border-sage/10" : "bg-canvas border border-dashed border-mist"}`}>
                    {dayMeals.lunch ? (
                      <div className="flex items-center gap-2">
                        <ChefHat className="w-4 h-4 text-sage flex-shrink-0" />
                        <div>
                          <p className="text-[10px] font-mono text-ink-soft uppercase">Lunch</p>
                          <p className="text-data font-medium text-ink">{dayMeals.lunch}</p>
                        </div>
                      </div>
                    ) : (
                      <span className="text-data text-ink-soft italic w-full text-center">Empty</span>
                    )}
                  </div>
                  <div className={`rounded-xl px-4 py-3 min-h-[48px] flex items-center ${dayMeals.dinner ? "bg-amber-soft/30 border border-amber/10" : "bg-canvas border border-dashed border-mist"}`}>
                    {dayMeals.dinner ? (
                      <div className="flex items-center gap-2">
                        <ChefHat className="w-4 h-4 text-amber flex-shrink-0" />
                        <div>
                          <p className="text-[10px] font-mono text-ink-soft uppercase">Dinner</p>
                          <p className="text-data font-medium text-ink">{dayMeals.dinner}</p>
                        </div>
                      </div>
                    ) : (
                      <span className="text-data text-ink-soft italic w-full text-center">Empty</span>
                    )}
                  </div>
                </div>
              </div>
            );
          })}
        </div>

        <div className="mt-5 pt-4 border-t border-mist flex items-center justify-between">
          <span className="text-data text-ink-muted">7 of 14 slots filled</span>
          <span className="text-[10px] font-mono text-sage font-semibold">50% planned</span>
        </div>
      </div>
    </div>
  );
}

function ShoppingMockup() {
  const items = [
    { name: "Daun salam", qty: "3 lembar", done: true },
    { name: "Serai", qty: "2 batang", done: true },
    { name: "Kunyit", qty: "2 ruas", done: true },
    { name: "Lengkuas", qty: "1 ruas", done: false },
    { name: "Kecap manis", qty: "1 botol", done: false },
    { name: "Bawang merah", qty: "10 siung", done: false },
    { name: "Cabai rawit", qty: "50g", done: false },
    { name: "Tomat", qty: "3 buah", done: false },
  ];

  const checked = items.filter((i) => i.done).length;

  return (
    <div className="bg-surface rounded-4xl border border-mist shadow-product-xl overflow-hidden">
      <div className="flex items-center gap-2 px-6 py-4 border-b border-mist bg-canvas/50">
        <div className="flex gap-1.5">
          <div className="w-3 h-3 rounded-full bg-red-400/80" />
          <div className="w-3 h-3 rounded-full bg-yellow-400/80" />
          <div className="w-3 h-3 rounded-full bg-green-400/80" />
        </div>
        <span className="flex-1 text-center text-[11px] font-mono text-ink-soft">dapurpintar.ai/shopping</span>
      </div>

      <div className="p-6">
        <div className="flex items-center justify-between mb-5">
          <div>
            <p className="font-display text-2xl text-ink">Soto Ayam Ingredients</p>
            <p className="text-data text-ink-muted mt-0.5">Generated from meal plan</p>
          </div>
        </div>

        {/* Progress */}
        <div className="mb-6">
          <div className="flex items-center justify-between mb-2">
            <span className="text-data text-ink-muted">{checked} of {items.length} items</span>
            <span className="text-[10px] font-mono text-sage font-semibold">{Math.round((checked / items.length) * 100)}%</span>
          </div>
          <div className="h-2 bg-canvas rounded-full overflow-hidden">
            <div
              className="h-full bg-sage rounded-full transition-all"
              style={{ width: `${(checked / items.length) * 100}%` }}
            />
          </div>
        </div>

        <div className="space-y-1">
          {items.map((item) => (
            <div
              key={item.name}
              className={`flex items-center gap-3 py-3 px-4 rounded-xl transition-colors ${item.done ? "opacity-50" : "hover:bg-canvas/50"}`}
            >
              <div
                className={`w-5 h-5 rounded-md border-2 flex items-center justify-center flex-shrink-0 transition-colors ${
                  item.done ? "bg-sage border-sage" : "border-mist"
                }`}
              >
                {item.done && <Check className="w-3 h-3 text-white" />}
              </div>
              <span className={`flex-1 text-data-lg ${item.done ? "line-through text-ink-soft" : "text-ink font-medium"}`}>
                {item.name}
              </span>
              <span className="text-data font-mono text-ink-soft">{item.qty}</span>
            </div>
          ))}
        </div>

        <div className="mt-5 pt-4 border-t border-mist flex items-center justify-between">
          <span className="text-data text-ink-muted">Estimated total</span>
          <span className="font-display text-xl text-ink">Rp 45.000</span>
        </div>
      </div>
    </div>
  );
}
