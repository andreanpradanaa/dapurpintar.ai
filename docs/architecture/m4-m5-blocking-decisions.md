# DapurPintar AI M5 Blocking Decision Records

## Document Control

| Item | Value |
|---|---|
| Milestone | M4 - Solution Architecture Refinement |
| Deliverable | M4-004 M5 Blocking Decision Records |
| Status | Draft - Awaiting Owner Approval |
| Parent documents | `docs/architecture/m4-decision-register.md`, `docs/architecture/implementation-backlog.md` |
| Scope | Decisions blocking M5: M4-DEC-006, M4-DEC-007, M4-DEC-013 |

## Purpose

This document records the decisions that block M5 Database Design, following the decision record format required by the Architecture Decision Register. Each record captures context, options considered, the recommended direction, consequences, affected documents, and revisit conditions. Approval converts the register item from `Pending` to `Decided`.

A record becomes final only after the owner approves it and affected disciplines confirm consequences. Approving these records unblocks DP-CON-001 and DP-CON-002 in the Implementation Backlog.

## Decision Record: M4-DEC-006 - Concrete PostgreSQL schema and migration shape

### Context and problem

M5 needs to produce a concrete PostgreSQL schema, constraints, indexes, and a Goose migration plan. The conceptual model in `docs/architecture/database-design.md` defines aggregate ownership but no schema syntax. Without an agreed schema and migration shape, M7 repositories and M9 features cannot start safely, and ad hoc table design could silently erase bounded-context ownership.

### Options considered

1. **Single schema, aggregate-owned tables (recommended).** One PostgreSQL schema holds all MVP tables. Every table carries the owning bounded context in its naming or documentation. Aggregate roots and their child entities live together; cross-context references store only the business identity and purpose, never child data.
2. **One schema per bounded context.** Each owning context gets its own PostgreSQL schema. Strongest ownership separation, but adds cross-schema joins and migration coordination for MVP scale and slows the first implementation slice.
3. **Single shared table model without explicit ownership.** Fastest to write but recreates a global data model and violates Database Design ownership principles.

### Recommended direction

Adopt **single schema with aggregate-owned tables**. One PostgreSQL schema, forward-only Goose migrations, UUID primary keys, `timestamptz` audit columns (`created_at`, `updated_at`), soft-delete flags only where Database Design calls for them, and an explicit ownership marker per table. Cross-context references use the owning context's business identity; no table stores another context's child entities.

### Consequences and risks

- Ownership stays visible in the schema and maps cleanly to SQLC query boundaries.
- Cross-context reads (e.g., Recommendation referencing Recipe) may need joins across ownership boundaries; joins are acceptable for read views but must not become write paths.
- Soft delete must not make entities appear active in business views.
- UUID keys avoid sequence leaks but require index and FK design care in M5.
- Risk of a future global Ingredient or Recipe model is reduced by keeping per-context language.

### Affected documents and contracts

- `docs/architecture/database-design.md` (M5 will refine while preserving ownership)
- `docs/architecture/diagrams/erd.mmd`
- Backlog items DP-CON-001, DP-CON-002, DP-FND-002, DP-FEAT-002, DP-FEAT-004, DP-FEAT-005

### Owner and approval

- Owner: Database + Backend
- Approval date: Pending

### Revisit condition

Revisit if M5 reveals an aggregate or ownership ambiguity that changes table boundaries, or if a future Household Collaboration scope requires a shared ownership model.

## Decision Record: M4-DEC-007 - MVP timezone policy

### Context and problem

Meal dates, pantry expiry, daily views, and recommendation context must agree. Without an explicit timezone policy, a user's local "today" can disagree with server-computed dates, causing missed meals, wrong expiry attention, and confusing daily views.

### Options considered

1. **Store UTC, interpret in user timezone (recommended).** Store all instants as `timestamptz` in UTC. Meal dates and daily views are computed against the user's profile timezone. Expiry attention is evaluated relative to the user's local date.
2. **Store user-local timestamps.** Simplest to display, but ambiguous at DST boundaries and wrong when the user changes timezone.
3. **Store an offset alongside the instant.** Preserves the local representation, but adds a field and still breaks when the user's timezone changes.

### Recommended direction

Store all instants as `timestepstz` in UTC. The User Profile records the user's timezone (e.g., IANA identifier, default `Asia/Jakarta` for the Indonesian MVP audience). Meal dates are stored as the business date and computed for daily views in the user's timezone. Expiry and "today" evaluations use the user's local date boundary. Profile timezone changes recompute derived views without rewriting stored instants.

### Consequences and risks

- Daily views and expiry logic need a timezone context passed through the query layer.
- Defaulting to `Asia/Jakarta` keeps the MVP correct for the primary audience until profile timezone selection exists.
- Risk of mismatched boundaries if some queries default to server UTC; M5 must require the user timezone on date-bounded queries.

### Affected documents and contracts

- `docs/architecture/database-design.md`
- Backlog items DP-CON-001, DP-FEAT-002 (expiry), DP-FEAT-004 (meal dates)

### Owner and approval

- Owner: Product + Backend
- Approval date: Pending

### Revisit condition

Revisit if the MVP audience expands beyond a single primary timezone or if a future Household Collaboration spans timezones.

## Decision Record: M4-DEC-013 - Raw prompt and conversation retention policy

### Context and problem

AI assistance stores Recommendation Conversation context and may retain prompt and conversation history. Without an explicit retention policy, the system risks retaining more personal or prompt data than the product needs, and the schema cannot be finalized until the retention boundary is known.

### Options considered

1. **Minimize retention; keep conversation only for active decision support (recommended).** Store only the bounded Recommendation Conversation needed to understand and continue the current recommendation. Raw provider prompts and payloads are never stored. Retain conversation content only while the owning Recommendation is active, then anonymize or delete per a short fixed window unless evaluation explicitly opts in.
2. **Retain conversation history indefinitely.** Supports evaluation and personalization but increases privacy risk and contradicts the "minimize personal data" principle.
3. **Retain full prompt/response pairs with opt-in evaluation.** Useful for quality work but requires explicit user consent, separation from business data, and a second retention pipeline.

### Recommended direction

Store only the bounded Recommendation Conversation for the active Recommendation. Never store raw provider prompts or payloads. Delete or anonymize conversation content when the Recommendation is closed or after a fixed retention window (proposed: 30 days) unless a separately consented evaluation record opts in. Evaluation records are stored apart from business data and are subject to their own privacy review.

### Consequences and risks

- Recommendation quality and acceptance measurement need to work from the durable Recommendation record (identity, options, rationale, lifecycle, timestamps) rather than raw conversation text.
- Users lose access to long-ago AI conversation detail; acceptable for MVP decision-support scope.
- Risk that a future AI Chat capability needs independent retention rules; this is covered by the future aggregate path in Database Design.

### Affected documents and contracts

- `docs/architecture/database-design.md` (AI Data Strategy)
- `docs/architecture/ai-architecture.md`
- Backlog items DP-CON-001, DP-AI-002, DP-FEAT-007

### Owner and approval

- Owner: Product + Security
- Approval date: Pending

### Revisit condition

Revisit if AI Chat becomes an independent durable capability, if evaluation needs change, or if a privacy or regulatory review changes the retention boundary.

## Approval Protocol

1. Owner reviews the recommended direction and evidence.
2. Affected disciplines (Database, Backend, AI, Security, Product) confirm consequences.
3. Architecture checks compatibility with bounded ownership and existing ADRs.
4. Approved records update the register status to `Decided`.
5. M5 work (DP-CON-001, DP-CON-002) may then start.

## Related Documents

- `docs/architecture/m4-decision-register.md`
- `docs/architecture/implementation-backlog.md`
- `docs/architecture/database-design.md`
- `docs/architecture/diagrams/erd.mmd`
- `docs/architecture/ai-architecture.md`
