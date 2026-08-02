# Backend Rules

> Backend-specific rules for DapurPintar AI. These complement `.opencode/AI_RULES.md` and the `backend-go` skill.

## Module layout

- The backend lives in `backend/`, module path `github.com/andreanpradanaa/dapurpintar.ai/backend`.
- Entry point: `backend/cmd/api/main.go`.
- Application scaffolding under `backend/internal/`: `http`, `auth`, `config`, `platform/*`.
- Each bounded context will own a package under `backend/internal/` for domain, application, and adapter code.

## Non-negotiables

- Clean Architecture over a modular monolith.
- DDD bounded contexts own their data and behavior.
- GORM and general-purpose ORMs are prohibited; use SQLC with reviewed SQL.
- PostgreSQL is the system of record; Redis is supporting infrastructure only.
- REST only, versioned under `/api/v1`.
- AI traffic flows through the AI Gateway only.

## API contract

- Responses use the M6 envelope from `internal/platform/response`.
- Errors use stable M6 codes from `internal/platform/errors`.
- Handlers validate input, call use cases, return envelope. No business logic in handlers.
- Session-cookie auth (`dp_session`) per M6; short-lived access tokens.
- Collections use cursor pagination.

## Persistence

- Goose migrations under `backend/migrations/`, forward-compatible and never destructive.
- SQLC queries defined per context in reviewed SQL.
- Soft-deleted rows filtered in SQL (`deleted_at IS NULL`).
- Timestamps stored UTC, interpreted in user timezone (`Asia/Jakarta` default).

## Quality gates

- `go build ./...`, `go vet ./...`, `gofmt -l .` clean.
- `go test -race ./...` passes.
- Tests cover domain invariants, adapters, and HTTP contracts.

## Security

- Secrets via environment/secret store only.
- Ownership and authorization derived server-side.
- Never log passwords, tokens, credentials, or private kitchen context.
