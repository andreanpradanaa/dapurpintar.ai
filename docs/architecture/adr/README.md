# Architecture Decision Records

This directory contains the accepted architectural decisions for DapurPintar AI. These records complement `docs/architecture/architecture-vision.md` and provide the reasoning behind the technology and structural choices for the MVP.

## ADR Index

| ADR | Decision | Status |
|---|---|---|
| [ADR-001](ADR-001-use-golang-for-backend.md) | Use Golang for backend | Accepted |
| [ADR-002](ADR-002-use-fiber-as-http-framework.md) | Use Fiber as HTTP framework | Accepted |
| [ADR-003](ADR-003-use-postgresql-as-system-of-record.md) | Use PostgreSQL as system of record | Accepted |
| [ADR-004](ADR-004-use-modular-monolith-for-mvp.md) | Use Modular Monolith for MVP | Accepted |
| [ADR-005](ADR-005-use-clean-architecture.md) | Use Clean Architecture | Accepted |
| [ADR-006](ADR-006-use-domain-driven-design.md) | Use Domain Driven Design | Accepted |
| [ADR-007](ADR-007-use-rest-api.md) | Use REST API | Accepted |
| [ADR-008](ADR-008-use-sqlc-for-database-access.md) | Use SQLC for database access | Accepted |
| [ADR-009](ADR-009-use-opentelemetry-for-observability.md) | Use OpenTelemetry for observability | Accepted |
| [ADR-010](ADR-010-use-ai-gateway-abstraction.md) | Use AI Gateway abstraction for OpenAI integration | Accepted |

## Conventions

- ADRs are numbered in decision order and should not be renumbered after publication.
- A new decision that changes an accepted decision should create a new ADR and mark the superseded record accordingly.
- These records describe architectural intent and trade-offs; detailed schemas, endpoint specifications, and implementation plans belong in later System Design documents.
