# DapurPintar AI - Project Context

> Primary context source for AI assistants working on this repository.
> Authority: `PROJECT_ROADMAP.md`, `docs/project/milestone-list.md`, `docs/architecture/m4-decision-register.md`, `.opencode/AI_RULES.md`.

---

## Project Overview

- **Project Name:** DapurPintar AI
- **Project Type:** AI-powered Kitchen Management SaaS
- **Current Phase:** Implementation
- **Current Milestone:** M7 - Backend Foundation (In Review)
- **Progress:** M7-001 Go/Fiber Backend Foundation - In Review
- **Next Milestone:** M8 - AI Foundation (DP-AI-001..003, blocked by M4-DEC-010/011/012/016)

---

## Product Vision

DapurPintar AI helps Sarah and Daniel make smarter kitchen decisions by leveraging AI, safely and without taking over decision making.

The product aims to:

- Reduce food waste.
- Simplify meal planning.
- Optimize grocery shopping.
- Recommend recipes based on available ingredients and expiry.
- Personalize recommendations using user preferences.
- Keep AI as decision support: it proposes, the user decides.

---

## Technology Stack

| Layer | Technology |
|-------|-----------|
| Frontend | Next.js, React, TypeScript, M3 design system |
| Backend | Go (1.26), Fiber v3, pgxpool, JWT (HS256), Argon2id, slog |
| Database | PostgreSQL (system of record), SQLC, Goose |
| Cache | Redis (non-authoritative) |
| AI | AI Gateway -> OpenAI, structured output |
| Observability | OpenTelemetry (OTLP/HTTP) |
| Deployment | Kubernetes (M15), secrets via env/secret store |

---

## Architecture

The project follows:

- Modular Monolith
- Clean Architecture
- Domain Driven Design (strategic + tactical)
- Bounded contexts owning their data and behavior (no global shared models)
- REST API, versioned under `/api/v1` (OpenAPI)
- Repository pattern implemented with SQLC
- Goose migrations (forward-compatible, non-destructive)
- Session-cookie auth (`dp_session`) per M6
- M6 response envelope + stable error codes
- UTC stored, interpreted in user timezone (default `Asia/Jakarta`)

---

## Bounded Contexts

- **Identity & Profile** — accounts, profiles, preferences.
- **Pantry Management** — pantry items, categories, expiry.
- **Recipe Management** — recipes, ingredient composition.
- **Meal Planning** — meal plans, recommendation acceptance.
- **Recommendation** — AI proposals based on user context.
- **Purchase & Waste** — (to be confirmed) shopping planning and waste tracking.

---

## Development Principles

Always prioritize:

1. Business correctness
2. Maintainability
3. Simplicity
4. Testability
5. Scalability

Non-negotiables:

- No GORM/ORM; use SQLC with reviewed SQL.
- AI only through the AI Gateway; never auto-commits business state.
- Authorization always derived server-side.
- No implicit cross-context mutation.
- Decision Register (M4-DEC-*) is authoritative; pending decisions block milestones.

---

## Working Agreement

- Every change belongs to one owning bounded context.
- Milestones are reviewed before approval (In Review -> Approved with sign-off).
- Never mark a milestone Approved/Closed without reviewer approval.
- Keep documents consistent with decisions and code (see `contexts/project.md`).
- Backend quality gates: `go build ./...`, `go vet ./...`, `gofmt -l .`, `go test -race ./...`.
