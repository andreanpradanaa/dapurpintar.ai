# Dapur Pintar

> A quiet cooking companion. Tell us what's in your kitchen — we'll help you decide what to make.

A premium recipe generator. The frontend (Next.js) talks to a Go + Fiber backend that uses a curated recipe library in Postgres plus OpenAI to compose fresh recipes (RAG pattern).

## Structure

```
dapurpintar.ai/
├── frontend/           # Next.js 15 app (the UI)
├── backend/            # Go 1.26 + Fiber v2 + Postgres + OpenAI
├── docker-compose.yml  # Postgres for local dev
├── design-system/      # design tokens, master doc
└── README.md
```

## Quick start (full stack)

```bash
# 1. Start Postgres
docker compose up -d postgres

# 2. Set backend env
cd backend
cp .env.example .env
# edit .env: set OPENAI_API_KEY

# 3. Start backend (auto-migrates, auto-seeds on first boot)
make dev

# 4. In another terminal, start the frontend
cd ../frontend
npm install
npm run dev

# 5. Open
#    http://localhost:3000  — frontend
#    http://localhost:8080/api/v1/health — backend health
```

## Stack

- **Frontend** — Next.js 15 + React 19, Tailwind 4, Framer Motion, Zustand, Lucide
- **Backend** — Go 1.26 + Fiber v2, pgx/v5, golang-migrate-compatible, go-playground/validator
- **Database** — PostgreSQL 16
- **AI** — OpenAI `gpt-4o-mini` with structured JSON output
- **Pattern** — RAG: top-3 from the curated library as style references, LLM composes a fresh recipe

## Phase 1 endpoints

| Method | Path | Purpose |
|---|---|---|
| `POST` | `/api/v1/recipes/generate` | Given ingredients, return a fresh recipe + match score + sources + alternatives |
| `GET`  | `/api/v1/recipes/:slug` | Fetch a recipe from the library by slug |
| `GET`  | `/api/v1/health` | Status, db, llm provider, recipe count |

See [`backend/README.md`](./backend/README.md) for the full backend docs, the RAG architecture, and migration commands.

## Documentation

- [`backend/README.md`](./backend/README.md) — backend architecture, RAG pipeline, migrations, env
- [`design-system/dapur-pintar-ai/MASTER.md`](./design-system/dapur-pintar-ai/MASTER.md) — design system, type scale, color tokens, accessibility

## What's NOT in Phase 1 (later phases)

- **Phase 2** — auth (JWT), `GET /api/v1/recipes` list, `GET /api/v1/ingredients` autocomplete
- **Phase 3** — user history, favorites (writes to DB)
- **Phase 4** — rate limiting, LLM response cache
- **Phase 5** — streaming responses, image generation
- **Phase 6** — observability
