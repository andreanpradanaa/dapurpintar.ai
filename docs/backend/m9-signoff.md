# DapurPintar AI M9 DP-FEAT-001 Identity and Access - Sign-off

## Document Control

| Item | Value |
|---|---|
| Milestone | M9 - MVP Features |
| Deliverable | M9-001 DP-FEAT-001 Identity and Access: registration, login, logout, refresh, profile, preferences |
| Status | Ready for review |
| Parent documents | `docs/architecture/m4-decision-register.md`, `docs/architecture/authentication-authorization.md`, `docs/api/m6-api-contract.md`, `docs/database/m5-schema.md` |
| Scope | Registration, login, logout, refresh with durable session rotation and revocation, current account, profile, and preferences over PostgreSQL |

## Purpose

This document records the DP-FEAT-001 vertical slice of the M9 MVP Features milestone: the Identity and Access and User Context and Preferences bounded contexts, delivered end to end from the M5 schema through SQLC-generated queries, the domain use cases, and the Fiber HTTP handlers. It is the first slice to prove the approved M7 Clean Architecture wiring (handler -> service -> store -> PostgreSQL) against a real database.

## Decisions Applied

- M4-DEC-001: email verification stays optional for local development; accounts register as `active`. Required before public launch.
- M4-DEC-002: credential change (password reset) is in MVP scope; password recovery is deferred.
- M4-DEC-003: PostgreSQL is the durable session and revocation authority. `auth_sessions` stores only a SHA-256 hash of the refresh secret, the lineage `family_id`, and revocation metadata. Refresh rotates the session (old marked `replaced_by`, new in the same family); reuse of a revoked session revokes the whole lineage.
- M4-DEC-004: session cookies `dp_session` (access) and `dp_refresh` (refresh) are HttpOnly, SameSite=Lax, Secure in production; the middleware reads the cookie or a Bearer header and origin-checks state changes.
- M4-DEC-006/007/008/009: M5 schema, SQLC, M6 API contract, and the recipe public boundary are honored (recipe tables exist in the base schema; recipe endpoints arrive with DP-FEAT-003).
- M4-DEC-013: recommendation conversations are retained while active with a 30-day window; not in this slice.

## Deliverables

### DP-FEAT-001-1 Schema and SQLC

- `backend/migrations/00001_base_schema.sql`: the approved M5 base schema in dependency order, plus `auth_sessions` (DP-FEAT-001 addition for M4-DEC-003). Tables are ordered so foreign keys resolve; the Down migration mirrors the reverse order.
- `backend/sqlc.yaml`: PostgreSQL engine, `pgx/v5` sql package, `uuid`/`timestamptz` mapped to `string`/`time.Time`, nullable text as pointer strings, and `date` mapped to `time.Time` so pgx v5 scans `valid_from` in binary format.
- `backend/internal/database/queries/identity.sql` and `user_context.sql`: named queries for accounts, sessions, profiles, and preference sets.
- `backend/internal/gen/sqlc/`: generated models, parameters, and query methods.

### DP-FEAT-001-2 Identity domain and store

- `backend/internal/identity/store.go`: the `Store` port (accounts, profiles, preferences, sessions) that handlers and services depend on; the SQLC adapter lives behind it.
- `backend/internal/identity/domain.go`: `Account`, `UserProfile`, `PreferenceSet`, and `Session` aggregates with their M5 lifecycle enums.
- `backend/internal/identity/service.go`: use cases Register, Login, Logout, Refresh (with rotation and lineage revocation), GetAccountByID, GetProfileByAccountID, UpdateProfile, UpdatePreferences, ActivePreferences. Argon2id password hashing, email normalization, password policy, and generic login failures (no account enumeration).
- `backend/internal/identity/store/postgres.go`: `Postgres` adapter mapping SQLC rows to domain types and translating `pgx.ErrNoRows` and unique violations into domain errors.

### DP-FEAT-001-3 HTTP wiring

- `backend/internal/http/handlers_identity.go`: register, login, refresh, logout, me, getProfile, updateProfile, updatePreferences with the M6 AccountView/UserProfileView/PreferencesView schemas and stable M6 error codes.
- `backend/internal/http/middleware/session.go`: `SessionCookies` (SetAccess/SetRefresh/ClearAccess/ClearRefresh) and `Authenticated` middleware reading the session cookie or Bearer header.
- `backend/internal/http/server.go`: routes wired under `/api/v1`, the authenticated group, session-cookie construction from config, and error mapping helpers.
- `backend/cmd/api/main.go`: constructs the identity store and service and injects them into `http.New`.

## Verified

- `go build ./...` passes.
- `go vet ./...` passes.
- `gofmt -l .` clean.
- `go test ./...` passes (auth, errors, ai, http).
- Integration tests run against the Docker Compose PostgreSQL (`TEST_DATABASE_URL` gated) and pass with `-v`.

Test coverage:

- `internal/identity/store/postgres_test.go` (integration, requires Postgres):
  - Account lifecycle: create, fetch by id/email, update status, duplicate-email `ErrEmailInUse`.
  - Profile and preferences: profile create/fetch/update, preference-set create and active lookup, missing set returns `ErrNotFound`.
  - Session lifecycle: create, lookup by secret hash, rotation via `MarkSessionReplacedBy`, revocation, and family revocation.
  - End to end: Register -> Login -> Refresh (rotated) -> reuse detection -> Logout over the real adapter.
- `internal/http/server_test.go`: health endpoint, protected route requires session (401 without), protected route succeeds with a valid token using a stub identity store.

## Dependencies on Pending Decisions

- M4-DEC-005 (unified login strategy) remains Pending for M15/M16 and does not block this slice.
- M4-DEC-014/015/017 (social login, MFA, and others) remain Pending for M15/M16.
- Email verification (M4-DEC-001) becomes mandatory before public launch; the `email_verified_at` column and status plumbing already exist.
- Password recovery (M4-DEC-002) is deferred; `UpdatePassword` arrives with that feature.

## Exit Criteria

DP-FEAT-001 is complete when:

- Registration, login, logout, refresh, current account, profile, and preferences work end to end over PostgreSQL.
- Refresh rotates sessions and detects reuse by revoking the lineage (M4-DEC-003).
- Session cookies are HttpOnly/SameSite=Lax and Secure outside development (M4-DEC-004).
- M6 schemas and error codes are honored.
- The slice is covered by unit and integration tests and the whole backend builds, vets, and is gofmt-clean.

## Related Documents

- `docs/architecture/m4-decision-register.md`
- `docs/architecture/authentication-authorization.md`
- `docs/api/m6-api-contract.md`
- `docs/database/m5-schema.md`
- `docs/database/m5-sqlc.md`
- `docs/project/milestone-list.md`
