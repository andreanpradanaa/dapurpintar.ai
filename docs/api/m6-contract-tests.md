# DapurPintar AI M6 API Design - Contract Tests and Compatibility Policy

## Document Control

| Item | Value |
|---|---|
| Milestone | M6 - API Design |
| Deliverable | M6-003 Contract Tests and Compatibility Policy |
| Status | Draft - Awaiting Cross-Functional Review |
| Parent documents | `docs/api/openapi.yaml`, `docs/api/m6-error-catalog.md` |
| Scope | Contract test strategy and version compatibility rules |

## Purpose

This document defines how the OpenAPI contract is verified and how `v1` evolves without breaking approved clients. It ensures the client and server agree on the contract before implementation, so no feature silently diverges from the documented behavior.

## Contract Test Strategy

### Layered verification

| Layer | What it verifies | When it runs |
|---|---|---|
| Schema lint | OpenAPI is structurally valid and self-consistent | Every change to `docs/api/openapi.yaml` |
| Request/response validation | Live requests and responses match the OpenAPI schemas | Continuous integration against the backend |
| State transition tests | Lifecycle commands honor the documented state machine | Feature implementation |
| Compatibility tests | New contract changes remain backward-compatible with `v1` | Every API change |
| Frontend contract tests | The frontend consumer matches the contract | Frontend build and integration |

### Test categories

#### Schema validation

- `docs/api/openapi.yaml` parses as valid OpenAPI 3.0.
- Every referenced schema exists; no dangling `$ref`.
- Every path has documented success and error responses.
- Enums in the contract match the M6-002 catalog and the M5 lifecycle constraints.

#### Behavioral contract tests

Behavioral tests run against a real backend instance with a disposable database:

- Every path from the endpoint catalog is exercised.
- Successful responses validate against the response schema.
- Error responses validate against the error envelope and catalog codes.
- Pagination is verified for every collection: default limit, max limit, malformed cursor, `has_more` boundary.
- Idempotency is verified for every retry-sensitive `POST`: same-key same-payload replay returns the original result; same-key different-payload returns a conflict.
- Authorization is verified: a resource owned by another account returns `404`/`403`, never the resource.
- Soft-deleted resources never appear active in business views.
- Date-bounded queries are verified with a user timezone other than the server default.

#### State transition tests

- Each aggregate lifecycle (Meal Plan, Shopping List, Kitchen Recommendation) is verified against its documented transitions.
- Invalid transitions return the documented `409` codes.
- Acceptance of a Recommendation Option only proceeds through a presented Recommendation.

## Compatibility Policy

### Version lifecycle

- The public contract is versioned by URI: `/api/v1`.
- `v1` is Stable for the MVP resources when approved.
- Additive changes are allowed within `v1`:
  - New response fields (optional additions).
  - New optional query parameters.
  - New resources.
  - New error codes for new failure modes.
- Breaking changes require a new major URI version (`/v2`):
  - Removing or renaming a resource.
  - Changing ownership semantics.
  - Changing required behavior or required fields.
  - Changing error meaning for an existing condition.
  - Changing a documented lifecycle transition.

### Change rules

- Every additive change must keep existing clients working without modification.
- Deprecated resources remain documented during a transition period and include a replacement direction.
- Future `/v2` must not silently change the meaning of existing user commitments or aggregate ownership.
- A change that alters API semantics, domain ownership, or product scope returns to the architecture or product review path before it is implemented.

### Contract ownership

- `docs/api/openapi.yaml` is the single source of truth for the wire contract.
- `docs/api/m6-error-catalog.md` is the single source of truth for error codes.
- Backend responses and frontend consumers both validate against these documents.
- M5 schema and M6 contract must stay consistent; a schema change that alters a response field triggers a contract review.

## CI Integration

- Contract validation runs in CI before backend merge.
- Generated clients or schemas are regenerated and compared with committed output.
- A failing contract test blocks merge.
- The full migration chain runs before behavioral contract tests so the schema under test matches the approved M5 baseline.

## Exit Criteria

M6-003 is complete when:

- Schema and behavioral contract test categories are defined.
- State transition and authorization cases are explicit.
- Compatibility policy separates additive from breaking changes.
- Contract ownership is assigned to the OpenAPI and error catalog documents.
- CI integration points are defined.

## Related Documents

- `docs/api/openapi.yaml`
- `docs/api/m6-error-catalog.md`
- `docs/architecture/api-design.md`
- `docs/database/m5-schema.md`
- `docs/database/m5-migrations.md`
- `docs/architecture/implementation-readiness.md`
