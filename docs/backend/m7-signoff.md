# DapurPintar AI M7 Backend Foundation - Sign-off

## Document Control

| Item | Value |
|---|---|
| Milestone | M7 - Backend Foundation |
| Deliverable | M7-001 Go/Fiber Backend Foundation |
| Status | Draft - Awaiting Cross-Functional Review |
| Parent documents | `docs/architecture/implementation-readiness.md`, `docs/architecture/m4-decision-register.md` |
| Scope | Go module, Fiber application skeleton, platform components, adapters, and authentication foundation |

## Purpose

This document records the M7 backend foundation deliverables. It is the first implementation milestone: a running Go/Fiber application skeleton with shared platform components, database and cache adapters, telemetry, and the authentication foundation. Business features are intentionally out of scope and land in M9.

## Deliverables

### M7-001 Go module and application skeleton

- Go module: `github.com/andreanpradanaa/dapurpintar.ai/backend`.
- Entrypoint: `backend/cmd/api/main.go` wires configuration, logging, telemetry, database, migrations, and the HTTP server.
- Fiber application with global `RequestID` and `Recover` middleware.
- Route groups matching the `/api/v1` contract: `GET /health`, and the `/api/v1/accounts` public and authenticated groups.

### M7-002 Configuration, logging, and error envelope

- `internal/config`: environment-driven configuration with required-value enforcement (`DATABASE_URL`, `JWT_SECRET`).
- `internal/platform/logger`: structured `slog` logger, JSON in production, text in development.
- `internal/platform/errors`: stable error codes from the M6-002 catalog with HTTP status mapping.
- `internal/platform/response`: the M6 success and error envelopes, including pagination and field-level details.

### M7-003 Database and cache adapters

- `internal/platform/database`: `pgxpool` connection pool; PostgreSQL remains the system of record.
- Goose migration runner (`internal/platform/database/migrate.go`) and the `backend/migrations/` directory with the M5-001 baseline placeholder.
- `internal/platform/cache`: Redis client; supporting infrastructure only, never an authorization source of truth.

### M7-004 Authentication foundation

- `internal/auth/token.go`: short-lived JWT access tokens with minimal claims and rotating refresh tokens.
- `internal/auth/password.go`: Argon2id password hashing with timing-safe verification.
- `internal/http/middleware/auth.go`: bearer-token authentication middleware storing the authenticated subject.
- Session transport and CSRF details remain pending the M4-DEC-003 and M4-DEC-004 decisions.

### M7-005 Telemetry foundation

- `internal/platform/telemetry`: OpenTelemetry OTLP/HTTP trace exporter with graceful shutdown, disabled by default locally.

## Verified

- `go build ./...` passes.
- `go vet ./...` passes.
- `gofmt -l .` clean.
- `go test -race ./...` passes for the auth, errors, and http packages.

## Dependencies on Pending Decisions

- M4-DEC-003 (token and session lifetimes) and M4-DEC-004 (cookie/CSRF policy) affect the final session transport; the foundation supports rotation and revocation, but browser cookie details await those decisions.
- M4-DEC-002 (credential recovery) affects account lifecycle endpoints, which land in M9.
- The M5 schema DDL placeholder will be replaced by the full migration set during M9 feature work.

## Exit Criteria

M7 is complete when:

- The Go/Fiber skeleton builds, vets, and tests clean.
- Configuration, logging, errors, and response envelopes match the approved contracts.
- PostgreSQL, Redis, Goose, and SQLC foundations are wired with authoritative-behavior boundaries.
- Authentication foundation exists and is tested.
- Telemetry instrumentation is initialized and gracefully shuts down.

## Related Documents

- `docs/architecture/implementation-readiness.md`
- `docs/architecture/m4-decision-register.md`
- `docs/database/m5-schema.md`
- `docs/api/m6-error-catalog.md`
- `docs/api/openapi.yaml`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/observability-architecture.md`
