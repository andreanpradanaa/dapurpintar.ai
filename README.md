# DapurPintar AI

Decide dinner with what you have. An AI-assisted kitchen companion for Indonesian home cooks.

## Quick Start

```bash
# 1. Start dependencies
docker compose up -d postgres

# 2. Apply migrations
cd backend
go install github.com/pressly/goose/v3/cmd/goose@v3.27.3
goose -dir migrations postgres "postgres://dapurpintar:dapurpintar@localhost:5432/dapurpintar?sslmode=disable" up

# 3. Configure env
cp .env.example .env
# Edit JWT_SECRET (min 32 random chars)

# 4. Start backend
go run ./cmd/api

# 5. Start frontend (in a new terminal)
cd ../frontend
echo "NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1" > .env.local
npm install
npm run dev
```

## Commands

```bash
# Backend (cd backend)
make build      # Build the API binary
make test       # Run unit tests
make itest      # Run integration tests (requires Postgres)
make lint       # go vet + go build
make fmt        # gofmt
make migrate-up # Apply pending migrations
make run        # Build and start

# Frontend (cd frontend)
npm run dev     # Start dev server on :3000
npm run build   # Production build
npm run lint    # ESLint
```

## Architecture

| Layer | Stack | Directory |
|---|---|---|
| Backend API | Go 1.26, Fiber v2, pgx v5, SQLC | `backend/` |
| Database | PostgreSQL 16, Goose migrations | `backend/migrations/` |
| Cache | Redis 7 (optional) | `docker-compose.yml` |
| AI Gateway | OpenAI adapter (optional) | `backend/internal/ai/` |
| Frontend | Next.js 14, TypeScript, Tailwind | `frontend/` |

### Bounded Contexts (Backend)

| Context | Package | Endpoints |
|---|---|---|
| Identity & Access | `internal/identity` | `/accounts`, `/profile` |
| Pantry Management | `internal/pantry` | `/pantry` |
| Recipe Experience | `internal/recipes` | `/recipes`, `/favorites` |
| Meal Planning | `internal/mealplan` | `/meal-plans` |
| Shopping Optimization | `internal/shopping` | `/shopping-lists` |
| AI Recommendations | `internal/recommendation` | `/recommendations` |

## Design System

The frontend implements the approved M3 design tokens:

- **Palette**: ink (dark text), steel (borders), paper (canvas), action-primary (decisions)
- **Typography**: Inter for UI, JetBrains Mono for data
- **States**: loading skeletons, empty states, error toasts on every screen

## Testing

```bash
# Backend unit tests
cd backend && go test ./...

# Backend integration tests (Postgres required)
TEST_DATABASE_URL="postgres://dapurpintar:dapurpintar@localhost:5432/dapurpintar?sslmode=disable" go test ./...

# Frontend
cd frontend && npm run lint
```
