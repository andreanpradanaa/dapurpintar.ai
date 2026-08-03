# DapurPintar AI — Landing Page Redesign from Zero

## Creative Brief

**Product:** AI-powered kitchen assistant for Indonesian home cooks.
**Single Job of This Page:** Make a visitor who has never heard of DapurPintar understand the product in 5 seconds, feel that it is premium and intelligent, and click "Get started" before scrolling past the hero.

**Design Thesis:** The kitchen is the last place in the house that still runs on memory and guesswork. Every other part of life has an assistant — email, calendar, fitness. DapurPintar brings that same intelligence to the most personal room in the home. The landing page should feel like meeting that assistant for the first time: calm, capable, and quietly brilliant.

**Aesthetic Risk:** Using warm, natural light as the dominant visual metaphor instead of the dark/tech aesthetic most AI products default to. The page should feel like morning sunlight through a kitchen window — bright, inviting, and human. This is the opposite of what most AI startups do (dark mode, neon accents), and it matches the product's domain: food, home, warmth.

---

## Design Tokens (Reference Only)

| Role | Value | Reasoning |
|---|---|---|
| Canvas | `#F9F7F2` | Warm off-white, like sunlight on a kitchen counter. Not generic `#FFFFFF` or cold `#FAFAFA`. |
| Surface | `#FFFFFF` | Cards and dashboard panels — pure white for contrast against warm canvas. |
| Ink | `#1C1B17` | Warm near-black for text. Not `#000` (too harsh) or `#17211B` (too green). Slight brown undertone. |
| Sage | `#5B7B5A` | Primary brand color — a muted, natural green. Feels organic, calm, intelligent. Not emerald-bright, not military. |
| Amber | `#E89143` | Accent for CTAs and attention states. Warm spice tone. Used sparingly. |
| Mist | `#E8E4DB` | Borders, dividers, subtle separations. Warm light gray. |

**Typography:**
- Display: **Instrument Serif** — a modern editorial serif with warmth and character. Used only for hero headline and section titles. This is the risk: a serif on a tech product.
- Body/UI: **Inter** — clean, neutral, excellent at small sizes. The workhorse.
- Mono: **JetBrains Mono** — for data labels, pantry counts, timestamps. Signals "system intelligence."

---

## Section-by-Section Design

---

### 1. Floating Glass Navigation

**Purpose:** Brand presence and primary navigation. Must be always accessible without competing with content.

**User Emotion:** "This is a real product, not a side project." Quiet confidence.

**Key Message:** DapurPintar is established, navigable, and has a clear entry point.

**Visual Direction:** A floating pill-shaped bar centered at the top with glass effect (`backdrop-blur`, subtle border, soft shadow). Logo wordmark on the left in Instrument Serif. Two links center: "How it works" and "Features". On the right: "Sign in" (ghost text) and "Get started" (amber pill button). On scroll, the bar stays but gains a faint canvas-colored background.

**Layout:**
```
      ┌─────────────────────────────────────────────────┐
      │  DapurPintar    How it works  Features    Sign in  [Get started] │
      └─────────────────────────────────────────────────┘
```

**Conversion Goal:** Always-visible CTA. The "Get started" button follows the user down the page so they can convert at any moment.

---

### 2. Hero — Full First Screen

**Purpose:** Make the visitor understand the product and feel its intelligence in under 5 seconds. This is the thesis of the entire page.

**User Emotion:** "Wow... this is a real AI product, and it's about MY kitchen." A mix of surprise, relief, and desire.

**Key Message:** DapurPintar is an AI assistant that knows what's in your pantry and tells you what to cook. It's not a recipe search engine. It thinks.

**Visual Direction:** Full viewport height (`100vh`). Bright warm canvas background with a subtle radial light coming from the top-right, simulating kitchen window light. The hero is split into two zones:

**Zone A — Narrative (left 45% on desktop, full width on mobile):**
- Eyebrow: `AI KITCHEN ASSISTANT` in mono, sage color, 11px, letter-spaced
- Headline in Instrument Serif, 64–80px: **"Your kitchen finally has a brain."**
- Subtitle in Inter, 20px, ink-muted: "DapurPintar reads your pantry, remembers what you have, and tells you exactly what to cook tonight — before anything expires."
- Two CTAs: Primary amber pill "Start cooking smarter →" and ghost "Watch demo"
- Social proof line below: "Join 2,000+ home cooks reducing food waste" with 5-star icons

**Zone B — Dashboard Preview (right 55% on desktop, below on mobile):**
A large, tilted (3D perspective) preview of the actual product dashboard. Not a screenshot — a styled mockup showing:
- Pantry panel with ingredient chips ("ayam", "santan", "bawang")
- AI response card: "You can make **Soto Ayam** tonight. It uses 3 of your ingredients and takes 20 minutes."
- A subtle floating sparkle/glow element near the AI response to signal intelligence
- The dashboard has a soft drop shadow and slight float animation (gentle up-down, 4s loop)

**Layout (Desktop):**
```
┌──────────────────────────────────────────────────────────────┐
│  AI KITCHEN ASSISTANT                                         │
│                                                               │
│  Your kitchen                        ┌──────────────────────┐ │
│  finally has                         │  YOUR PANTRY          │ │
│  a brain.                            │  ayam  santan  bawang │ │
│                                       │                       │ │
│  DapurPintar reads your pantry,       │  AI SUGGESTION        │ │
│  remembers what you have, and         │  "You can make        │ │
│  tells you exactly what to cook       │   Soto Ayam tonight."│ │
│  tonight — before anything expires.  │  20 min · 4 servings  │ │
│                                       └──────────────────────┘ │
│  [Start cooking smarter →]  Watch demo                        │
│                                                               │
│  ★★★★★  Join 2,000+ home cooks reducing food waste           │
└──────────────────────────────────────────────────────────────┘
```

**Layout (Mobile):**
```
┌──────────────────────────┐
│  AI KITCHEN ASSISTANT     │
│                          │
│  Your kitchen            │
│  finally has             │
│  a brain.                │
│                          │
│  DapurPintar reads your  │
│  pantry and tells you    │
│  what to cook tonight.   │
│                          │
│  [Start cooking smarter→]│
│  Watch demo              │
│                          │
│  ★★★★★ 2,000+ cooks     │
│                          │
│  ┌────────────────────┐  │
│  │  YOUR PANTRY        │  │
│  │  ayam santan bawang │  │
│  │  AI: "Soto Ayam!"   │  │
│  └────────────────────┘  │
└──────────────────────────┘
```

**Conversion Goal:** The dashboard preview IS the conversion. The visitor sees the product working, not a description of it. The amber CTA is the only amber element on screen, making it the single point of focus for action.

---

### 3. Problem Section — "The Invisible Problem"

**Purpose:** Make the visitor recognize a pain they've normalized. Don't sell the solution yet — make them feel the problem first.

**User Emotion:** "Oh... I do waste a lot of food." Recognition. A quiet sting.

**Key Message:** You buy ingredients, forget them, and throw them away. Everyone does. It's not your fault — it's that nothing in your kitchen reminds you.

**Visual Direction:** Full-width band with warm canvas background. No cards. Three large editorial statements in Instrument Serif, arranged horizontally with thin dividers between them. Each statement is a single sentence with a bold number.

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│                                                              │
│   1.3 billion tonnes         Rp 48 million                    │
│   of food wasted             per household                   │
│   every year globally.       over a lifetime in Indonesia.   │
│                                                              │
│   ─────────────  ─────────────  ─────────────                │
│                                                              │
│   You don't need another recipe app.                         │
│   You need something that knows what you already have.       │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** Create emotional tension. The visitor now WANTS a solution because they feel the problem is real and personal. This sets up the solution section as relief, not pitch.

---

### 4. Solution Section — "Meet DapurPintar"

**Purpose:** Resolve the tension from the problem section. Introduce the product as the answer.

**User Emotion:** "Finally, someone solved this." Relief. Curiosity about how.

**Key Message:** DapurPintar is an AI assistant that lives in your kitchen. It tracks what you have, suggests what to cook, plans your week, and builds your shopping list — all from one place.

**Visual Direction:** Dark sage background band (full-width, `#5B7B5A`). White text. A single large headline in Instrument Serif italic. Below it, a horizontal flow showing the three core actions: "See your pantry → Get a suggestion → Cook tonight." Each step has a small custom illustration (not an icon — a line-drawn kitchen object). The dark band creates visual rhythm contrast with the bright sections around it.

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│                    (dark sage background)                     │
│                                                              │
│              DapurPintar knows your kitchen                   │
│              so you don't have to think.  (italic serif)      │
│                                                              │
│    ┌──────────┐    ┌──────────┐    ┌──────────┐             │
│    │ See what │    │ Get a    │    │ Cook     │             │
│    │ you have │───→│ smart    │───→│ tonight  │             │
│    │          │    │ suggestion│    │          │             │
│    └──────────┘    └──────────┘    └──────────┘             │
│                                                              │
│         (line-drawn illustrations above each step)            │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** The dark band is a visual reset — it signals "this is the answer." The three-step flow makes the product feel simple, not complex. The visitor thinks "I can do three steps."

---

### 5. Feature Showcase — Immersive Storytelling, Not Card Grid

**Purpose:** Show depth without boredom. Each feature gets its own visual moment, not a card in a grid.

**User Emotion:** "This is more thoughtful than I expected." Growing respect for the product.

**Key Message:** DapurPintar handles four kitchen decisions: what to cook, what to track, what to plan, and what to buy.

**Visual Direction:** Four alternating full-width sections, each with a different layout pattern. NO card grids. NO icon-in-circle. Each section is a two-column layout that alternates left-right, with a large product mockup on one side and narrative text on the other. The mockups are real UI fragments, not illustrations.

**Layout pattern:**
```
Section A — AI Recommendations (text left, mockup right)
──────────────────────────────────────────────────────────────
  01                                                    ┌───┐
  AI RECOMMENDATIONS                                    │   │
                                                       │ UI│
  Type what's in your kitchen.                         │   │
  Get recipes you can make right now.                  │   │
                                                       │   │
  The AI considers your ingredients,                    │   │
  expiry dates, and cooking time —                    │   │
  then suggests three meals you can                    └───┘
  cook tonight.

──────────────────────────────────────────────────────────────
Section B — Smart Pantry (mockup left, text right)
──────────────────────────────────────────────────────────────
  ┌───┐                                                 02
  │   │                                          SMART PANTRY
  │UI │
  │   │     Track every ingredient with          Auto-categorize
  │   │     intelligence.                        ingredients, get
  │   │                                         expiry alerts before
  └───┘                                         anything goes bad.
──────────────────────────────────────────────────────────────
Section C — Weekly Planner (text left, mockup right)
──────────────────────────────────────────────────────────────
  03                                                    ┌───┐
  MEAL PLANNER                                          │   │
                                                       │ UI│
  Plan seven days of meals in                           │   │
  a visual grid. Assign recipes                         │   │
  to days, occasions, and moods.                       │   │
                                                       └───┘
──────────────────────────────────────────────────────────────
Section D — Shopping Lists (mockup left, text right)
──────────────────────────────────────────────────────────────
  ┌───┐                                                 04
  │   │                                          SHOPPING LIST
  │UI │
  │   │     Generate lists from plans,           Activate, shop,
  │   │     check off as you go,                 and complete —
  │   │     never buy what you                  all in one list.
  └───┘
```

**Conversion Goal:** Each section ends with a subtle inline CTA: "Try this feature →" linking to signup. The alternating layout creates rhythm — the eye knows something new is coming, not the same card again.

---

### 6. Interactive Dashboard Showcase — Visual Centerpiece

**Purpose:** This is the "wow" moment. The single section where the visitor stops scrolling and leans in. Show the full product experience in one composed frame.

**User Emotion:** "I want this. Now." Desire. The moment of conversion decision.

**Key Message:** One screen. One assistant. Your entire kitchen, organized.

**Visual Direction:** Full-width section. A large, centered "product canvas" — a styled representation of the DapurPintar dashboard that takes up 80% of the section width. It's not a screenshot, it's a composed scene:

- Left panel: pantry list with ingredient chips, expiry badges, quantity
- Center: AI conversation — "I have ayam, santan, and bawang. What should I cook?" → AI responds with a recommendation card
- Right panel: weekly meal planner grid with 3 recipes assigned
- Bottom strip: shopping list with items checked off

The scene has a soft, realistic shadow. Subtle floating animation. Around the canvas, small floating annotation labels: "AI suggests", "Expiry alert", "Drag to plan", "Auto-generated". These labels have small connecting lines to the UI elements they describe.

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│                                                              │
│                    One screen. One assistant.                 │
│                                                              │
│  ┌─────────┐  ┌──────────────────┐  ┌──────────────────┐   │
│  │ PANTRY  │  │  AI CONVERSATION │  │  WEEK PLANNER    │   │
│  │         │  │                  │  │                  │   │
│  │ ayam ✓ │  │  "I have ayam,   │  │  Mon: Soto       │   │
│  │ santan ✓│  │   santan,        │  │  Tue: Nasi Goreng│   │
│  │ bawang ✓│  │   bawang..."     │  │  Wed: Sate       │   │
│  │ telur ⚠ │  │                  │  │  Thu: —          │   │
│  │         │  │  → "Soto Ayam!   │  │  Fri: Gado-gado  │   │
│  │         │  │     20m, 4 srv"  │  │                  │   │
│  └─────────┘  └──────────────────┘  └──────────────────┘   │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  SHOPPING LIST:  ✓ santan  ✓ bawang  ☐ daun jeruk   │   │
│  └──────────────────────────────────────────────────────┘   │
│                                                              │
│         [Start your free pantry →]                            │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** This is the conversion anchor. The CTA below the dashboard is the strongest on the page because it follows the "wow" moment. The visitor has just seen the product work end-to-end. The button is amber, large, and the only action on screen.

---

### 7. How It Works — Simple 3-Step Flow

**Purpose:** Remove perceived complexity. Show that getting started takes minutes, not hours.

**User Emotion:** "That's it? I can do this." Confidence. Low friction feeling.

**Key Message:** Three steps: add your ingredients, get AI suggestions, cook and repeat.

**Visual Direction:** Horizontal timeline on a bright canvas background. Three nodes connected by a thin sage line. Each node is a circle with a number, not an icon. Below each number, a single sentence and a small UI fragment showing that step. The timeline reads left to right like a recipe.

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│                                                              │
│                    How it works                              │
│                                                              │
│    ●──────────●──────────●                                  │
│    1          2          3                                   │
│                                                              │
│    Add        Get         Cook &                            │
│    ingredients suggestions  repeat                          │
│                                                              │
│    Type what's  The AI      Use the suggestion,             │
│    in your      suggests    rate it, and the AI              │
│    kitchen      meals       learns your taste               │
│                                                              │
│    [chip chip]  [card]      [★ rate]                        │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** Friction reduction. The visitor who was interested but worried about setup complexity now feels it's easy. This directly precedes the testimonial section to provide social proof right after the "I can do this" feeling.

---

### 8. Testimonials — Real People, Real Kitchens

**Purpose:** Social proof that this product works for people like the visitor.

**User Emotion:** "Other people like me use this and love it." Validation. Trust.

**Key Message:** DapurPintar is used and loved by real Indonesian home cooks.

**Visual Direction:** Background shifts to warm canvas (`#F9F7F2`). Three testimonial cards, but NOT identical — one is larger (featured), two are smaller side cards. The featured testimonial has a longer quote, a name, a location, and a specific result ("reduced my food waste by 60%"). Cards have soft shadows and warm borders.

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│                    What home cooks are saying                 │
│                                                              │
│  ┌──────────────────────┐  ┌──────────────────┐            │
│  │ "I used to throw away│  │ "Meal planning    │            │
│  │  vegetables every     │  │  used to take     │            │
│  │  week. Now DapurPintar│  │  hours."         │            │
│  │  reminds me to use    │  │                  │            │
│  │  them first. I've cut │  │  — Budi, Jakarta │            │
│  │  my waste by 60%."    │  │                  │            │
│  │                      │  └──────────────────┘            │
│  │  — Rina, Bandung      │  ┌──────────────────┐            │
│  │  Home cook, 2 kids   │  │ "DapurPintar      │            │
│  │                      │  │  ngerti masakan   │            │
│  └──────────────────────┘  │  Indonesia."     │            │
│                              │  — Dian, Surabaya│            │
│                              └──────────────────┘            │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** Social validation. The specific, measurable result ("60% less waste") is more persuasive than a star rating. The Indonesian locations make it feel local and real.

---

### 9. Pricing Preview — "Coming Soon" Frame

**Purpose:** Establish that this will be a paid product (increasing perceived value) while offering free early access (removing friction).

**User Emotion:** "I'm getting something valuable for free." Smart. Early.

**Key Message:** DapurPintar will be a premium product. Right now, it's free during early access. Early users keep their access forever.

**Visual Direction:** A single centered card on canvas background. Not a pricing table — one plan. The card has a subtle amber glow border. Inside: "Early Access — Free" in large serif, three bullet points of what's included, and a single CTA. Below the card, in small muted text: "Paid plans launching 2026. Early users are grandfathered forever."

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│                                                              │
│                    Pricing                                   │
│                                                              │
│            ┌──────────────────────────┐                      │
│            │  (subtle amber glow)      │                      │
│            │                          │                      │
│            │  Early Access             │                      │
│            │  Free                     │                      │
│            │                          │                      │
│            │  ✓ Unlimited pantry       │                      │
│            │  ✓ AI recommendations     │                      │
│            │  ✓ Meal planner           │                      │
│            │  ✓ Shopping lists         │                      │
│            │  ✓ AI conversation       │                      │
│            │                          │                      │
│            │  [Get early access →]     │                      │
│            └──────────────────────────┘                      │
│                                                              │
│      Paid plans launching 2026.                             │
│      Early users keep their access — forever.               │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** Urgency through scarcity ("early access") without being pushy. The "forever" line is the closer — it makes signing up now feel like a smart investment, not a free trial.

---

### 10. FAQ — Calm, Accordion-Free

**Purpose:** Answer objections without making the page feel like a support doc.

**User Emotion:** "My question is answered." Reassurance.

**Key Message:** Common concerns are anticipated and addressed honestly.

**Visual Direction:** NOT an accordion. All questions and answers visible in a clean two-column layout — question in serif italic on the left, answer in sans on the right. Thin dividers between items. This feels editorial, not technical.

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│                    Questions                                  │
│                                                              │
│  Do I need an account?        Yes, to save your pantry       │
│                               and preferences. Try the       │
│                               search above without one.       │
│  ────────────────────────────────────────────────────        │
│  Is the AI free?              During early access, yes.       │
│                               Early users keep access.        │
│  ────────────────────────────────────────────────────        │
│  What about my data?          We store only what you add.    │
│                               No selling, no sharing.          │
│  ────────────────────────────────────────────────────        │
│  Does it work on my phone?    Yes — desktop, tablet, mobile. │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** Address the last objections. The visible answers (not hidden in accordions) mean the visitor doesn't have to work to find reassurance.

---

### 11. Final CTA — The Kitchen Door

**Purpose:** Last chance to convert. Must feel like an invitation, not a sales pitch.

**User Emotion:** "I'm ready." Decision made.

**Key Message:** Your kitchen is waiting. Open the door.

**Visual Direction:** Full-width section with warm sage background (`#5B7B5A`). Large Instrument Serif headline in white. A single amber button. Below, the social proof line again: "2,000+ home cooks. Free during early access."

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│                    (sage background)                          │
│                                                              │
│                                                              │
│              Your kitchen is waiting.                         │
│              (Instrument Serif, 48px, white, italic)          │
│                                                              │
│              Join 2,000+ home cooks who let AI                │
│              handle the thinking.                            │
│                                                              │
│              [Start cooking smarter →]                        │
│              (amber button, large)                            │
│                                                              │
│              Free during early access                        │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** This is the last impression. The sage background is calming, not aggressive. The headline is an invitation, not a command. The amber button is the only warm element on a cool background — it draws the eye naturally.

---

### 12. Footer — Quiet Sign-Off

**Purpose:** Brand closure. Navigation without noise.

**User Emotion:** "This is a real company." Confidence in post-signup experience.

**Key Message:** DapurPintar is a product, not a landing page.

**Visual Direction:** Simple, warm canvas background. Three columns: brand + tagline, product links, company links. Tiny copyright line at the bottom. No newsletter signup, no social media bar — just clean, confident closure.

**Layout:**
```
┌──────────────────────────────────────────────────────────────┐
│  DapurPintar        Product          Company                  │
│  AI kitchen         Features         About                    │
│  assistant          Pricing          Privacy                  │
│                    FAQ              Terms                     │
│                                                              │
│                    © 2026 DapurPintar AI                      │
└──────────────────────────────────────────────────────────────┘
```

**Conversion Goal:** Minimal — just trust through professionalism. A messy footer undermines everything above it.

---

## Complete Landing Page Hierarchy

```
1. Floating Glass Navigation (sticky, always visible)
2. Hero — full viewport, narrative left + dashboard preview right
3. Problem — emotional pain, full-width, three stat cards
4. Solution — dark sage band, three-step flow
5. Feature Showcase — four alternating sections, immersive
6. Interactive Dashboard — full-width centerpiece, three-panel composition
7. How It Works — three-step timeline
8. Testimonials — asymmetric card layout, real results
9. Pricing — single card, early access frame
10. FAQ — two-column editorial, visible answers
11. Final CTA — sage background, single amber button
12. Footer — three-column, quiet sign-off
```

---

## Estimated Scroll Flow

```
[100vh]  Hero — visitor understands product, sees dashboard, CTA visible
         ↓ gentle scroll, warm canvas continues
[60vh]   Problem — stats hit hard, emotional shift
         ↓ dark sage band appears = visual reset
[50vh]   Solution — three steps, relief
         ↓ back to bright canvas
[240vh]  Feature Showcase — four alternating sections (~60vh each)
         ↓ momentum builds, variety keeps scroll engaging
[80vh]   Dashboard Centerpiece — the "wow" moment, conversion anchor
         ↓ rhythm slows
[40vh]   How It Works — simple, calming
         ↓
[50vh]   Testimonials — social proof
         ↓
[40vh]   Pricing — urgency frame
         ↓
[40vh]   FAQ — objection handling
         ↓ dark sage = visual return to brand
[50vh]   Final CTA — the ask
         ↓
[20vh]   Footer — quiet exit
─────────────────
Total: ~7 screen heights of content
```

---

## Top 10 Reasons This Outperforms the Current Landing

| # | Reason | Why |
|---|---|---|
| 1 | Hero shows the product working, not a text description | Visitors convert when they see value, not when they read about it |
| 2 | Dashboard centerpiece is a composed scene, not card grid | One "wow" moment beats four "nice" moments |
| 3 | Problem section creates emotional tension before selling | People buy solutions to felt problems, not to feature lists |
| 4 | Dark sage bands create visual rhythm and brand contrast | Prevents scroll fatigue and page-blending that kills engagement |
| 5 | Instrument Serif for headlines creates brand identity | No AI startup uses a warm serif. This is instantly distinguishable |
| 6 | Amber is used ONLY for CTAs — single warm accent | Every amber element is clickable. The eye learns: amber = action |
| 7 | Inline CTAs after each feature section, not just hero + footer | Conversion happens at any scroll depth, not just top and bottom |
| 8 | FAQ shows all answers, no accordion friction | Friction in objection-handling = lost conversions |
| 9 | Pricing frames "early access" as investment, not free trial | "Grandfathered forever" creates urgency without pressure |
| 10 | Warm canvas (#F9F7F2) instead of cold white/gray | Kitchen = warmth. The page feels like the domain, not a generic SaaS template |
