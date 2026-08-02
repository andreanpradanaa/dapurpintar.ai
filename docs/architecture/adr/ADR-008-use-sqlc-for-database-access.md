# ADR-008 Use SQLC for Database Access

- Status: Accepted
- Date: 2026-08-02

## Context

PostgreSQL is the system of record, and the modular monolith requires explicit repository boundaries. The project needs type-safe database access, predictable SQL behavior, and maintainable queries without allowing generated persistence types to become the domain model. Goose remains the schema migration mechanism defined by the architecture vision.

## Decision

Use SQLC to generate type-safe Go database access from reviewed SQL. SQLC-generated code will be confined to PostgreSQL adapters and repositories. Application and domain layers will depend on repository ports and domain/application models, not SQLC-generated types. Goose will manage ordered schema migrations separately.

## Consequences

- SQL remains visible, reviewable, and optimizable while reducing manual row-mapping and type errors.
- Generated code improves consistency across repositories and supports PostgreSQL features directly.
- Repository boundaries remain explicit and prevent persistence concerns from leaking into business rules.
- SQLC generation adds a build and toolchain step that must remain reproducible.
- Query changes require coordination with schema migrations and repository contracts.

## Alternatives Considered

- **ORM:** Faster for conventional CRUD, but can hide query behavior and weaken control over performance and PostgreSQL-specific semantics.
- **Hand-written database access:** Maximum control, but more repetitive mapping and a larger surface for type and scanning errors.
- **Generic repository abstraction:** Reduces visible SQL, but risks losing query clarity and forcing an abstraction that does not match domain needs.
