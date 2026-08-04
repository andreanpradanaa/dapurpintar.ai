# Dapur Pintar — Master Design System

> Last updated: Reskin phase complete. Source of truth for all design decisions.

---

## 1. North star

> When a user lands on Dapur Pintar, they should feel like they've stepped into a quiet, well-stocked kitchen — not a software product. AI is the mechanism, not the feature. The brand reads like an *appliance* and a *companion*, not a tool.

## 2. Brand references, decoded

| Reference | What we steal |
|---|---|
| **Apple** | Generous whitespace, one focal element per section, no decoration, photographic hero |
| **Aesop** | Editorial serif display, warm earth palette, single restrained accent, soft product photography, copy that reads like a thoughtful note |
| **MUJI** | Off-white paper background, no shadows, no glow, hairline rules, monospace for numeric data |
| **KitchenAid** | Warm cream + ink + one earth accent, ingredient-led copy, the *physical* feel of cooking |
| **Airbnb** | Calm, editorial layouts, real-feeling photography, copy that puts the *experience* first |

## 3. What we explicitly do NOT do

- No emerald, neon, or cyberpunk colors
- No purple/indigo gradients
- No glow effects, no halo orbs
- No glassmorphism, no backdrop blur
- No "AI-powered" language in user-facing copy
- No scale bounce on press
- No animated color gradients
- No emoji as icons
- No auto-playing anything
- No dark mode (light only — commitment to the lifestyle aesthetic)

## 4. Color system (locked, do not deviate)

### Surface

| Token | Hex | Use |
|---|---|---|
| `--color-bg-base` | `#FAF6F0` | warm cream paper — the default page background |
| `--color-bg-card` | `#FFFFFF` | white surface — cards, inputs, dialogs, sidebar |
| `--color-bg-elevated` | `#FFFFFF` | white with shadow (no separate "elevated" tone — same as card) |
| `--color-bg-section` | `#F4EFE5` | one shade darker — alternating section rhythm, footer |
| `--color-bg-overlay` | `rgba(26, 22, 18, 0.32)` | soft ink wash for modals |

### Text (deep warm ink scale)

| Token | Hex | Use |
|---|---|---|
| `--color-text-primary` | `#1A1612` | body & headings |
| `--color-text-secondary` | `#3D352D` | secondary, tab labels |
| `--color-text-muted` | `#6B6358` | captions, helper text |
| `--color-text-subtle` | `#9B9286` | very muted, hairlines-adjacent |
| `--color-text-inverse` | `#FAF6F0` | text on dark/accent surfaces |

### Accent — Terracotta (single, used sparingly)

| Token | Hex | Use |
|---|---|---|
| `--color-accent` | `#A8553A` | primary action, focus ring, brand mark |
| `--color-accent-hover` | `#8E4530` | hover |
| `--color-accent-active` | `#723627` | press |
| `--color-accent-soft` | `rgba(168, 85, 58, 0.08)` | soft wash, chip bg, subtle tints |
| `--color-accent-soft-strong` | `rgba(168, 85, 58, 0.14)` | hover wash, badge |
| `--color-accent-fg` | `#FFFFFF` | text on accent |

### Border (ink-tinted, low contrast)

| Token | Value |
|---|---|
| `--color-border` | `rgba(26, 22, 18, 0.08)` |
| `--color-border-strong` | `rgba(26, 22, 18, 0.14)` |
| `--color-border-accent` | `rgba(168, 85, 58, 0.4)` |

### Semantic (muted earth tones, never neon)

| Token | Hex | Note |
|---|---|---|
| `--color-success` | `#5C7A4D` | muted sage, not emerald |
| `--color-warning` | `#B8804A` | ochre, not amber |
| `--color-danger` | `#A14B3A` | brick red, not bright red |
| `--color-info` | `#5A7A8A` | dusty blue |

### Contrast verification (WCAG)

- `#1A1612` on `#FAF6F0` → **15.4:1** (AAA)
- `#3D352D` on `#FAF6F0` → **11.2:1** (AAA)
- `#6B6358` on `#FAF6F0` → **5.0:1** (AA)
- `#9B9286` on `#FAF6F0` → **2.7:1** (large text only — use sparingly)
- `#A8553A` on `#FAF6F0` → **4.6:1** (AA large text only)
- `#FFFFFF` on `#A8553A` → **5.4:1** (AA, used for primary button text)

## 5. Typography (Fraunces + Inter)

### Families

- **Display: Fraunces** (variable, Google Fonts) — used for h1/h2/h3, stat numbers, brand wordmark. Variable font with `opsz` and `SOFT` axes for editorial weight + humanist warmth.
- **Body & UI: Inter** (variable) — 300/400/500/600/700
- **Mono: JetBrains Mono** (variable) — for nutrition values, time, serving counts, stat numbers in tabular contexts

### Type scale (slightly larger and looser than typical SaaS)

| Token | Size / lh | Tracking | Use |
|---|---|---|---|
| Display | 80 / 88 | -0.035em | Hero h1 |
| H1 | 64 / 72 | -0.03em | Section h1 |
| H2 | 48 / 56 | -0.025em | Section h2 |
| H3 | 32 / 40 | -0.015em | Page-level h3 |
| H4 | 22 / 30 | -0.01em | Card title |
| Body L | 18 / 30 | 0 | Lead paragraph |
| Body | 16 / 26 | 0 | Default |
| Small | 14 / 22 | 0 | UI |
| Caption | 12 / 16 | +0.04em | Meta uppercase (12 → 10) |
| Mono | 13 / 20 | 0 | Numerics |

**Line-height deliberately increased** (1.5+ for body, ~1.0 for display). This is the editorial move.

### Display treatment

Display headings use Fraunces with:
- `font-variation-settings: "opsz" 144, "SOFT" 50` for editorial softness
- Tight tracking (-0.03em to -0.04em)

Use the `.font-display` utility class which sets these automatically.

## 6. Spacing

8pt-aligned scale, extended for marketing:

`4 · 8 · 12 · 16 · 20 · 24 · 32 · 40 · 48 · 64 · 80 · 96 · 128`

**Section vertical padding (desktop):**
- Marketing sections: `py-32` (128px top + bottom)
- Hero: `pt-36 sm:pt-44` to clear the floating nav
- App sections: `py-6` to `py-8` (24-32px)
- Cards internal: `p-5` to `p-6` (20-24px)

## 7. Radius

| Token | Value | Use |
|---|---|---|
| `rounded-md` | 8px | Buttons, inputs, chips |
| `rounded-lg` | 12px | Cards (small), tab pills |
| `rounded-xl` | 16px | Cards (default) |
| `rounded-2xl` | 20px | Cards (large), modals |
| `rounded-3xl` | 24px | Hero mockups (not used in reskin) |
| `rounded-full` | 9999px | Pills, avatars, badges, nav, toggle switches |

## 8. Shadow (warm-tinted, restrained)

```css
--shadow-xs: 0 1px 2px rgba(26, 22, 18, 0.04);
--shadow-sm: 0 1px 2px rgba(26, 22, 18, 0.04), 0 1px 1px rgba(26, 22, 18, 0.02);
--shadow-md: 0 4px 16px rgba(26, 22, 18, 0.06);
--shadow-lg: 0 12px 32px rgba(26, 22, 18, 0.08);
--shadow-xl: 0 24px 64px rgba(26, 22, 18, 0.10);
--shadow-accent: 0 4px 16px rgba(168, 85, 58, 0.16);  /* primary CTA only */
--shadow-accent-lg: 0 8px 32px rgba(168, 85, 58, 0.20); /* primary CTA hover */
```

**No glow halos. No `0 0 0 1px rgba(...)` accent rings. No colored drop shadows on cards. Cards have no shadow by default; shadow appears on hover only.**

## 9. Motion (calmed significantly from typical SaaS)

| Token | Duration | Use |
|---|---|---|
| Micro | 180ms | Hover, press |
| State | 240ms | State change |
| Layout | 360ms | Tab switch, sidebar |
| Reveal | 600ms | Page enter, scroll reveal |

**Easing:** `cubic-bezier(0.22, 1, 0.36, 1)` for all entries (calmer than the typical Expo.out)
**Press:** Subtle darken (`active:bg-accent-active`). **No scale bounce** — this is a lifestyle product, not a tech product.
**Hover:** `-translate-y-0.5` on button, shadow-md → shadow-lg on cards.

`prefers-reduced-motion: reduce` honored throughout via global CSS rule.

## 10. Iconography

**Source:** Lucide React. **No emoji as icons.** Stroke width 1.5–1.75 for the refined feel.

Common sizes: 12 (inline), 14 (UI), 16 (button), 20 (feature card icon).

## 11. Components

### Button variants

| Variant | Use |
|---|---|
| `primary` | Main CTA — terracotta bg, white text, soft accent shadow |
| `secondary` | Less-emphasized CTA — white card + hairline border |
| `ghost` | Tertiary action — text + hover bg |
| `outline` | Card-like CTA — border only |
| `destructive` | Delete/danger — brick red bg |
| `link` | Inline link — terracotta, underline on hover |

Sizes: `sm` (36px), `md` (44px), `lg` (48px), `icon` variants.

### Card variants

| Variant | Use |
|---|---|
| `default` | Most cards — white + hairline border, no shadow |
| `elevated` | Floating modals, popovers — adds shadow-md |
| `outline` | Glass-like over photo — transparent + border |

### Chip

Spring-animated entry/exit. `accent` variant uses terracotta-soft for active selection.

### Logo

**Wordmark only.** Fraunces 500, ink color, no icon. Replaceable.

## 12. Layout principles

### Container

- Marketing: `max-w-7xl mx-auto px-4 sm:px-6 lg:px-8`
- App: `max-w-7xl mx-auto px-4 sm:px-6 lg:px-8` (sidebar gutter on lg+)

### Breakpoints

- `sm`: 640px
- `md`: 768px (sidebar collapses to drawer)
- `lg`: 1024px (sidebar full)
- `xl`: 1280px (max content width)

### Grid patterns

- Marketing hero: `grid lg:grid-cols-12 gap-12 lg:gap-16` (text 7 cols, photo 5 cols)
- Features: `grid md:grid-cols-2 lg:grid-cols-3 gap-5`
- App cards: `grid sm:grid-cols-2 lg:grid-cols-3 gap-4`
- Stats: `grid grid-cols-2 lg:grid-cols-4 gap-4`
- Recipe detail: `grid lg:grid-cols-3 gap-6` (sidebar 1 col, main 2 cols)

## 13. Page rhythm

Marketing pages alternate between `--color-bg-base` (cream) and `--color-bg-section` (slightly darker cream) for section rhythm. No dark sections. No high-contrast color blocks.

## 14. Photography treatment

All photos receive a CSS filter:
```css
filter: saturate(0.85) contrast(1.02) brightness(1.02);
```

Plus a very subtle warm overlay (4% terracotta) when needed. Photos are sourced from Unsplash with curated search terms focused on warm, editorial, neutral-surface food imagery. All in `lib/photo.ts`.

## 15. Copy principles

- **The product is cooking. AI is the engine, not the headline.**
- The user is a "cook" or "home cook", not a "user" or "customer"
- The verb is "cook", "make", "try", "consider" — not "generate", "compute", "analyze"
- "Created" not "generated" for past-tense actions
- Section eyebrows are simple phrases ("What you can do", "Common questions") not jargon ("Capabilities", "FAQ")
- Hero CTAs are short, human verbs: "Cook" not "Generate Recipe"
- 404 page: "Lost in the pantry." — playful, on-brand
- Auth pages: replace generic motivational copy with single-line italic quotes about cooking
- Never use "AI-powered", "intelligent", "smart" (except in the user-facing way: "smart substitutions" is fine, "smart AI engine" is not)

## 16. Accessibility checklist (per build)

- [x] Color contrast 4.5:1+ for all body text (verified)
- [x] 44×44px minimum touch targets
- [x] Visible focus rings (terracotta, 2px outline, 2px offset)
- [x] `aria-label` on every icon-only button
- [x] `aria-expanded` on accordions, dropdowns
- [x] `role="progressbar"` with value/max on Progress
- [x] `role="tablist" / "tab" / "tabpanel"` on Tabs
- [x] Keyboard navigation
- [x] `prefers-reduced-motion: reduce` honored
- [x] Viewport meta with `width=device-width, initial-scale=1`
- [x] `color-scheme: light` declared

## 17. Performance budget

- First Load JS: ~195 KB (target: < 200 KB) ✅
- All routes pre-rendered except dynamic `/recipes/[slug]`
- Images use `next/image` with explicit `sizes`
- Fraunces loaded as variable font (single weight file)

## 18. What's intentionally NOT here

- No dark mode (light-only is a brand commitment)
- No second font family
- No glassmorphism, no backdrop blur
- No glow halos, no orb gradients
- No emerald, purple, neon, or cyberpunk colors
- No "AI" in visible landing copy
- No emoji as icons
- No glassmorphism on sticky nav
- No animated color gradients
- No 3D, no parallax, no scroll-jacking
- No auto-playing media
- No cookie banner (no real cookies)
- No analytics, no Sentry, no third-party scripts
