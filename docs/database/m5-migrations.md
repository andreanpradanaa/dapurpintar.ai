# DapurPintar AI M5 Database Design - Goose Migration Strategy

## Document Control

| Item | Value |
|---|---|
| Milestone | M5 - Database Design |
| Deliverable | M5-002 Goose Migration Strategy |
| Status | Draft - Awaiting Cross-Functional Review |
| Parent documents | `docs/database/m5-schema.md`, `docs/architecture/m4-m5-blocking-decisions.md` |
| Scope | Migration tooling, versioning, ordering, rollback policy, and release sequence |

## Purpose

This document defines how the M5-001 PostgreSQL schema is delivered and evolved using Goose. It establishes versioning, ordering, forward-compatibility, and failure behavior so that backend foundation work in M7 can rely on a stable migration baseline.

## Tooling Baseline

- **Tool:** Goose (Go SQL migration tool).
- **Storage:** Migration version state kept in the PostgreSQL database itself via Goose's version table.
- **Schema:** Migrations target the `public` schema.
- **Driver:** Goose PostgreSQL driver.
- **Location:** `migrations/` directory at the repository root.
- **Local workflow:** Migrations run through the development scripts established in M0; the same scripts drive test and CI environments.

## Versioning Convention

- Migration files use Goose's default format: `<version>_<name>.sql` for up and down pairs.
- The version is a strictly increasing integer. The MVP starts at `00001`.
- Up and down files for one version must be paired and reviewed together.
- Each migration is immutable once applied in any shared environment. Corrections are delivered as new migrations, never by editing an applied file.

## Migration Ordering

Migrations are ordered by the version number and follow the schema dependency order:

| Version | Purpose | Contents |
|---|---|---|
| 00001 | Base schema | Enable extensions (`uuid-ossp` or use `gen_random_uuid()` from core), create all tables, constraints, and indexes from M5-001 in dependency order |
| 00002 | Base data | Public recipe seed baseline and reference data required for discovery (subject to M5-004) |
| 00003 | Retention support | Conversation retention cleanup procedure and any supporting index |

The base schema migration itself must create tables in dependency order so foreign keys resolve:

1. `accounts`
2. `user_profiles`
3. `preference_sets`
4. `pantries`
5. `pantry_items`
6. `recipes`
7. `recipe_favorites`
8. `meal_plans`
9. `planned_meals`
10. `kitchen_recommendations`
11. `recommendation_options`
12. `recommendation_conversations`
13. `shopping_lists`
14. `shopping_items`

Note: `shopping_lists` references both `meal_plans` and `kitchen_recommendations`, and `planned_meals` references `recommendation_options`; those dependencies drive the ordering above.

## Migration Content Rules

- Every migration includes the up file and the down file.
- A migration must be forward-compatible with the release sequence. It may add nullable columns, new tables, or new indexes; it must not destroy data the running application still reads.
- `NOT NULL` additions with no default require a two-phase pattern (add nullable, backfill, set not null) unless the table is new.
- Soft-delete and lifecycle columns are added with safe defaults when introduced.
- No migration stores raw secrets, provider credentials, prompts, or production data.
- Seed data migrations must be idempotent or guarded so they do not duplicate on re-run.

## Rollback Policy

- Rollbacks are the exception, not the default. The primary path for a bad migration is forward: apply a corrective migration.
- A down file is still provided for every migration so a failed application can be unwound cleanly during development and local work.
- Rollback is never used as a substitute for a corrective migration in a shared environment once other migrations have been applied on top.
- Goose supports down-by-one and down-to-version; local development may use it, shared environments must follow the corrective-migration rule.

## Failure Behavior

- If a migration fails, Goose records the failure and stops. The database is left in the state of the failed transaction where possible.
- The backend must treat an out-of-date migration baseline as a startup error: refuse to start if pending migrations exist that are not part of the approved deploy.
- CI and test environments rebuild from zero on every run to verify the full migration chain from `00001` upward.
- A failing migration in CI blocks merge; the migration must be corrected in review, not force-applied.

## Migration Integrity Checks

- The full up chain must apply cleanly from an empty database.
- The down chain must rewind cleanly to zero in a disposable environment at least once per review.
- Every migration is reviewed for ownership preservation: no migration may create a global shared model or merge bounded-context ownership.
- Before merge, `goose status` must show the full chain and every applied version must match the committed files.

## Release Sequence

- Migrations are versioned independently of application releases. A migration may exist in the repository before the feature that uses it, as long as it is backward-compatible.
- Deployment applies migrations before starting new backend code that requires them.
- A migration that changes existing user-visible meaning (state transitions, ownership, privacy) must reference the decision or ADR that approved it.

## Exit Criteria

M5-002 is complete when:

- Goose tooling and directory layout are defined.
- Versioning, ordering, and dependency-order rules are explicit.
- Content rules prevent destructive or data-losing changes.
- Rollback and failure behavior are defined for local and shared environments.
- Integrity checks verify the migration chain in CI and test builds.

## Related Documents

- `docs/database/m5-schema.md`
- `docs/architecture/m4-m5-blocking-decisions.md`
- `docs/architecture/database-design.md`
- `docs/architecture/implementation-readiness.md`
