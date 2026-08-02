# Backend Go Rules

These rules apply to all Go backend work in this repository.

## Layering

- Business rules belong in the Domain and Application layers, never in HTTP handlers, database adapters, or middleware.
- Handlers parse and validate HTTP input, call application use cases, and return the response envelope.
- Application use cases orchestrate domain behavior and coordinate cross-aggregate workflows.
- Domain layer owns aggregate invariants, lifecycle, and policies.
- Repositories are interfaces in the domain/application language; SQLC implements them behind the boundary.

## Bounded contexts

- Every change belongs to exactly one owning bounded context.
- Never access another context's tables directly. Use published business identity references.
- Never create a shared global model for Ingredient, Recipe, Meal, Recommendation, or Preference.
- Cross-context reads must preserve source ownership and never become write paths.

## API contract

- Use REST only. Every endpoint is versioned: `/api/v1/...`.
- Return the M6 response envelope via `internal/platform/response`.
- Use stable M6 error codes from `internal/platform/errors`. Never leak SQL, stack traces, prompts, credentials, or provider payloads.
- `Idempotency-Key` is honored on retry-sensitive `POST` commands.
- Collections use cursor pagination with documented `sort` and `order`.

## Persistence

- PostgreSQL is the system of record. SQLC generates queries from reviewed SQL.
- Goose manages migrations; migrations are forward-compatible and never destructive.
- Redis is supporting infrastructure only. Business data never lives in Redis.
- Soft-deleted rows are filtered in SQL (`deleted_at is null`), not after retrieval.
- Date-bounded queries receive the user timezone and compute local boundaries; never default to server UTC.

## Authentication and authorization

- Auth middleware establishes identity; the Application Layer authorizes each resource.
- Ownership scope is derived server-side. Never accept client-supplied `userId` or scope for authorization.
- Access tokens carry minimal claims and never include pantry, preference, or conversation context.

## Quality

- Every feature includes tests: domain invariants, adapter behavior, and HTTP contract.
- Every change includes error handling, logging, and observability signals.
- Code must pass `go build ./...`, `go vet ./...`, `gofmt -l .` clean, and `go test -race ./...`.
- Prefer small functions, dependency injection, composition, and explicit interfaces. Avoid global state and god objects.

## Security

- Never store secrets in code. Load from environment or deployment secret management.
- Never log passwords, tokens, refresh secrets, provider credentials, or private kitchen context.
- Validate all external input at the API boundary and again by business rules.
- AI APIs defend against prompt injection, unsafe output, and cost exhaustion.
