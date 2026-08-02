---
name: backend-go
description: Use when implementing or reviewing Go backend code for DapurPintar AI. Covers the Fiber application skeleton, Clean Architecture module layout, M6 error codes and response envelopes, bounded-context ownership, SQLC/Goose/PostgreSQL/Redis rules, and the M4 Definition of Ready/Done. Trigger on Go, Fiber, repository, use case, handler, migration, SQLC, or backend feature work.
---

# Backend Go Skill

## Purpose

Guide Go backend implementation and review for DapurPintar AI so every change follows the approved architecture: Clean Architecture over a modular monolith, DDD bounded-context ownership, a versioned REST contract under `/api/v1`, and PostgreSQL as the system of record.

## Responsibilities

- Implement backend features as vertical business slices owned by one bounded context.
- Keep domain behavior testable without HTTP, PostgreSQL, Redis, OpenAI, or a browser.
- Expose business capabilities through the Fiber HTTP layer, never database tables.
- Enforce authorization server-side; never trust client-supplied ownership.
- Preserve the M6 error contract and response envelope.
- Keep AI output as decision support, never as business truth.

## Inputs

- `backend/` existing module layout.
- `docs/architecture/api-design.md` and `docs/api/openapi.yaml` for the public contract.
- `docs/api/m6-error-catalog.md` for stable error codes.
- `docs/architecture/bounded-context.md` and `docs/architecture/tactical-ddd.md` for ownership.
- `docs/database/m5-schema.md`, `docs/database/m5-migrations.md`, `docs/database/m5-sqlc.md` for persistence.
- `docs/architecture/m4-decision-register.md` and `docs/architecture/implementation-readiness.md` for decision gates and DoR/DoD.
- `.opencode/AI_RULES.md` for project-wide rules.

## Outputs

- Go code that builds, vets, and passes `go test -race ./...`.
- Repository interfaces and SQLC queries within the owning bounded context.
- Fiber handlers returning the approved response envelope.
- Tests covering domain behavior, adapters, and HTTP contracts.
- Updated `PROJECT_ROADMAP.md` and milestone-list when milestones change.

## Dependencies

- Go, Fiber, pgx, SQLC, Goose, Redis, OpenTelemetry.
- The M5/M6/M7 artifacts must be read before writing persistence or API code.
- Pending decisions in the Decision Register must be referenced, not silently resolved.

## Status

Active - aligned with M7/M9 backend implementation.
