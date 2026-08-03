# Dapur Pintar AI

> Turn Ingredients Into Delicious Meals with AI.

A premium AI recipe generator. Enter what's in your kitchen, get a tailored recipe with steps, timing, and nutrition — in seconds.

## Quick start

```bash
npm install
npm run dev
```

Open http://localhost:3000.

## Stack

- **Next.js 15** (App Router) + **React 19**
- **Tailwind 4** with custom design tokens via `@theme`
- **Framer Motion** for micro-interactions
- **Zustand** for client state (persisted to localStorage)
- **Lucide React** for icons
- **Sonner** for toasts
- **Inter** (display + body) + **JetBrains Mono** (numerics)

## Routes

### Public (no auth)

| Route | Purpose |
|---|---|
| `/` | Landing page — 8 sections |
| `/pricing` | Standalone pricing |
| `/faq` | Standalone FAQ |
| `/login` | Sign in |
| `/register` | Create account |

### Authenticated (gated by sidebar app shell)

| Route | Purpose |
|---|---|
| `/dashboard` | Greeting + quick-generate + stats + recent |
| `/generate` | Recipe generator with AI loading flow |
| `/recipes/[slug]` | Recipe detail with steps, nutrition, timer |
| `/history` | List of past generations |
| `/favorites` | Saved recipes |
| `/settings` | Tabs: Account, Preferences, Notifications, Billing, Danger |
| `/profile` | User profile + activity |

### System

| Route | Purpose |
|---|---|
| `/not-found` | 404 |
| `/error` | Error boundary |
| `/loading` | Global loading |
| `/offline` | Offline fallback |

## Design system

All tokens live in `app/globals.css` under `@theme`. Type-safe re-exports in `lib/design-tokens.ts`. UI primitives in `components/ui/`.

- **Color**: `#09090B` base · `#18181B` card · `#1F2937` elevated · `#10B981` accent
- **Type**: Inter only, 6 weights
- **Radius**: 6 / 8 / 12 / 16 / 20 / 24 / 999
- **Easing**: `cubic-bezier(0.16, 1, 0.3, 1)` (Expo.out) for enters, `cubic-bezier(0.4, 0, 1, 1)` for exits
- **Motion**: 150ms micro · 200ms state · 300ms layout · 500ms reveal
- **Dark-mode first**; light mode supported via system preference

See `design-system/dapur-pintar-ai/MASTER.md` for the full design rationale.

## Mock data

- `lib/mock-data/recipes.ts` — 30+ realistic Indonesian recipes (Nasi Goreng, Soto Ayam, Rendang, etc.) with full ingredients, steps, timing, nutrition, and bilingual copy
- `lib/mock-data/ingredients.ts` — 50+ common Indonesian pantry items

## Scripts

```bash
npm run dev        # localhost:3000
npm run build      # production build
npm run start      # serve production
npm run lint       # eslint
npm run typecheck  # tsc --noEmit
```

## Notes

This build uses fully mocked data. To wire to a real backend, replace the implementation in `lib/generate.ts` and `lib/store.ts` actions with API calls. The store's `addHistory`, `toggleFavorite`, etc. are the only places that mutate state.
