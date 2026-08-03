# Dapur Pintar — Backend

Go 1.26 + Fiber v2 + PostgreSQL + OpenRouter (LLM gateway).

## Phase 1 scope

`POST /api/v1/recipes/generate` — given a list of ingredients, scores
the curated Postgres library by ingredient overlap, takes the top 3 as
style references, and asks the LLM to compose a fresh recipe.

`GET /api/v1/recipes/:slug` — returns a recipe from the library.

`GET /api/v1/health` — status + db + llm.

## Quick start

### Option A — local Go + Docker Postgres

```bash
# 1. Start Postgres
docker compose up -d postgres

# 2. Set env
cp .env.example .env
# edit .env: set OPENAI_API_KEY

# 3. Run (auto-migrates, auto-seeds on first boot)
make dev
```

### Option B — full Docker

```bash
cp .env.example .env
# set OPENAI_API_KEY in .env
docker compose up --build
```

## Endpoints

### `POST /api/v1/recipes/generate`

```json
// request
{ "ingredients": ["chicken","garlic","rice"], "dietary": ["halal"], "language": "en" }

// response
{
  "recipe": { "id":"gen_…", "title":"…", … },
  "matchScore": 9,
  "sources": [{ "id":"r-001", "slug":"nasi-goreng-ayam", "title":"Nasi Goreng Ayam" }, …],
  "alternatives": [{ "id":"r-002", "slug":"soto-ayam", "title":"Soto Ayam Bening", "score":6 }, …]
}
```

### `GET /api/v1/recipes/:slug`

Returns the full recipe document from the library.

### `GET /api/v1/health`

```json
{ "status":"ok", "db":"up", "llm":"openai", "recipes": 32 }
```

## Layout

```
backend/
├── cmd/
│   ├── server/         # main binary
│   └── seed/           # standalone seed command
├── internal/
│   ├── config/         # env → struct
│   ├── model/          # Recipe, Ingredient, Nutrition, …
│   ├── repo/           # RecipeRepo interface + pgx impl + seed helper
│   ├── service/        # matcher (scoring), generator (RAG orchestrator), llm/ (OpenAI client)
│   ├── handler/        # Fiber route handlers
│   ├── middleware/     # recover
│   └── router/         # route registration
├── data/
│   ├── recipes.json    # 32 base recipes (seed)
│   └── ingredients.json
├── migrations/          # 000001_init.{up,down}.sql
├── test/                # *_test.go (handler + matcher)
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
└── .env.example
```

## Architecture (RAG)

```
User ingredients
       ↓
POST /generate
       ↓
1. SELECT all recipes
2. score each by ingredient overlap + dietary + rating
3. take top 3 → "style references"
4. build OpenAI prompt (system + references + ingredients + language)
5. gpt-4o-mini returns fresh recipe (strict JSON schema)
6. validate, hydrate defaults (gradient, id, slug, createdAt)
7. return { recipe, matchScore, sources, alternatives }
```

## Testing

```bash
make test
```

Handler tests use an in-memory repo + stub LLM — no Postgres or OpenAI
key required.

## Migrations

```bash
make migrate-install  # one-time
make migrate-up
make migrate-create name=add_history_table  # new migration
```

The server also auto-runs migrations on boot.

## LLM provider flexibility

The backend speaks the OpenAI Chat Completions API. Point it at any
compatible endpoint via `OPENAI_BASE_URL`:

| Provider | Base URL | Notes |
|---|---|---|
| OpenAI (default) | `https://api.openai.com/v1` | `gpt-4o-mini`, `gpt-4o` |
| opencode.ai | `https://api.opencode.ai/v1` (verify) | If opencode.ai exposes an OpenAI-compatible API |
| Together | `https://api.together.xyz/v1` | `meta-llama/...` |
| Groq | `https://api.groq.com/openai/v1` | `llama-3.1-...` |
| OpenRouter | `https://openrouter.ai/api/v1` | Any model |
| Local llama.cpp | `http://localhost:8080/v1` | Self-hosted |

The model name in `OPENAI_MODEL` must match what the endpoint expects.

## Notes for future phases

- **Phase 2** — auth (JWT), `GET /api/v1/recipes` list, `GET /api/v1/ingredients` autocomplete
- **Phase 3** — user history, favorites (writes to DB)
- **Phase 4** — rate limiting, LLM response cache (Redis)
- **Phase 5** — streaming responses (SSE), image generation
- **Phase 6** — OpenTelemetry, structured logs → Loki
