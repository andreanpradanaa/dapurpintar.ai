# DapurPintar AI M6 API Design - Error, Validation, Pagination, and Idempotency Catalog

## Document Control

| Item | Value |
|---|---|
| Milestone | M6 - API Design |
| Deliverable | M6-002 Error, Validation, Pagination, and Idempotency Catalog |
| Status | Draft - Awaiting Cross-Functional Review |
| Parent documents | `docs/api/openapi.yaml`, `docs/architecture/api-design.md` |
| Scope | Stable error codes, validation behavior, pagination contract, and idempotency semantics |

## Purpose

This document defines the stable machine-readable error catalog, validation rules, pagination behavior, and idempotency semantics that back the OpenAPI contract. It gives backend, frontend, QA, and AI-assisted development one authoritative reference so behavior does not drift between implementations.

## Response Envelope

### Success envelope

```json
{
  "data": {},
  "page": null,
  "request_id": "req_01HZX..."
}
```

- `data`: the resource, collection, or business view.
- `page`: `null` for non-collection responses; populated for collections.
- `request_id`: correlation identifier also returned in the `X-Request-ID` header.

### Error envelope

```json
{
  "error": {
    "code": "PANTRY_ITEM_NOT_FOUND",
    "message": "The pantry item is not available in your pantry.",
    "details": [],
    "request_id": "req_01HZX..."
  }
}
```

- `code`: stable, machine-readable, documented in this catalog.
- `message`: user-safe explanation.
- `details`: field or query validation details; never secrets or internals.
- `request_id`: correlation identifier also returned in the `X-Request-ID` header.

## HTTP Status Categories

| Status | Meaning | Used for |
|---|---|---|
| 200 | Operation completed | Reads, lifecycle commands, responses with body |
| 201 | Resource created | POST creation operations |
| 202 | Accepted | Recommendation request that is being created asynchronously |
| 204 | No content | DELETE, PUT, or commands returning no body |
| 400 | Malformed or contract validation | Unknown filters, unsupported sorts, malformed identifiers |
| 401 | Missing/invalid authentication | Session or credential problems |
| 403 | Known but not authorized | Valid identity outside scope, restricted account |
| 404 | Not available in scope | May conceal another user's resource existence |
| 409 | State conflict | Invalid lifecycle transition, idempotency reuse conflict, duplicate creation |
| 422 | Understandable but violates domain rules | Business validation, invariant violations |
| 429 | Rate or abuse limit | Registration, login, refresh, AI-intensive endpoints |
| 503 | Dependency unavailable | AI provider failure; core operations remain usable |
| 5xx | Internal failure | Unexpected server errors |

## Error Code Catalog

Codes are grouped by concern. Each code is stable; renaming requires a new major version.

### Authentication and session

| Code | HTTP | Meaning |
|---|---|---|
| AUTH_CREDENTIALS_INVALID | 401 | Credentials did not match; generic, no enumeration |
| AUTH_SESSION_EXPIRED | 401 | Access session expired |
| AUTH_SESSION_INVALID | 401 | Access session malformed or invalid |
| AUTH_REFRESH_REUSED | 401 | Refresh credential reuse detected; lineage revoked |
| AUTH_ACCOUNT_RESTRICTED | 403 | Account is restricted and cannot perform the operation |
| AUTH_ACCOUNT_NOT_ACTIVE | 403 | Account is not active for normal participation |
| AUTH_SCOPE_FORBIDDEN | 403 | Identity valid but operation outside authorized scope |

### Validation

| Code | HTTP | Meaning |
|---|---|---|
| VALIDATION_FIELD_INVALID | 400 | A field failed contract-level validation |
| VALIDATION_QUERY_UNSUPPORTED | 400 | Unknown filter, sort, or order value |
| VALIDATION_PAGINATION_INVALID | 400 | Cursor malformed or limit out of range |
| VALIDATION_ID_INVALID | 400 | Resource identifier is not a valid UUID |
| VALIDATION_PAYLOAD_MALFORMED | 400 | Body is not valid JSON or does not match schema |

### Registration and account

| Code | HTTP | Meaning |
|---|---|---|
| ACCOUNT_EMAIL_IN_USE | 409 | Registration email already registered (subject to enumeration policy) |
| ACCOUNT_EMAIL_INVALID | 422 | Email does not satisfy account policy |
| ACCOUNT_PASSWORD_WEAK | 422 | Password fails policy |
| ACCOUNT_VERIFICATION_REQUIRED | 403 | Account pending verification under product policy |

### Domain and business

| Code | HTTP | Meaning |
|---|---|---|
| PANTRY_ITEM_NOT_FOUND | 404 | Pantry item not available in scope |
| PANTRY_QUANTITY_NEGATIVE | 422 | Quantity cannot be negative |
| PANTRY_EXPIRY_INVALID | 422 | Expiry date is invalid or contradictory |
| RECIPE_NOT_FOUND | 404 | Recipe not available |
| RECIPE_NOT_PUBLIC | 404 | Recipe detail not available in public scope |
| MEAL_PLAN_NOT_FOUND | 404 | Meal plan not available in scope |
| MEAL_PLAN_PERIOD_INVALID | 422 | Period end precedes period start |
| MEAL_SLOT_CONFLICT | 409 | Meal slot already occupied in the planning period |
| PLANNED_MEAL_NOT_FOUND | 404 | Planned meal not available in scope |
| SHOPPING_LIST_NOT_FOUND | 404 | Shopping list not available in scope |
| SHOPPING_ITEM_NOT_FOUND | 404 | Shopping item not available in scope |
| SHOPPING_STATE_CONFLICT | 409 | Lifecycle transition not allowed from current state |
| RECOMMENDATION_NOT_FOUND | 404 | Recommendation not available in scope |
| RECOMMENDATION_OPTION_NOT_FOUND | 404 | Recommendation option not available in scope |
| RECOMMENDATION_STATE_INVALID | 409 | Transition not allowed from current recommendation state |
| RECOMMENDATION_OPTION_NOT_ACCEPTABLE | 409 | Option cannot be accepted in its current state |
| CONVERSATION_NOT_FOUND | 404 | Conversation not available in scope |
| CONVERSATION_STATE_INVALID | 409 | Conversation closed or completed cannot be continued |

### AI dependency

| Code | HTTP | Meaning |
|---|---|---|
| AI_UNAVAILABLE | 503 | AI provider could not complete the operation |
| AI_SAFETY_REJECTED | 422 | Output rejected by safety policy |
| AI_QUOTA_EXCEEDED | 429 | AI usage quota or cost budget reached |

### Abuse and limits

| Code | HTTP | Meaning |
|---|---|---|
| RATE_LIMIT_EXCEEDED | 429 | Documented usage or abuse limit exceeded |

## Validation Rules

- Input validation happens at the API boundary and again by business rules. Contract validation failures are `400`; domain invariant violations are `422`.
- Unknown filters, unsupported sorts, and invalid order values are validation errors, never silently ignored.
- String lengths and required fields follow the OpenAPI schema.
- `format: date` values must be valid calendar dates in the user's accepted format; they are stored as user-local dates.
- Numeric quantities are non-negative (`minimum: 0`) and are validated before persistence.
- Enum values are restricted to the documented lifecycle values.

## Pagination Contract

### Cursor pagination

- Query parameters: `cursor` and `limit`.
- Default `limit` is 20; server-enforced maximum is 100.
- Response `page` object:

```json
{
  "next_cursor": "eyJhZnRlciI6InV1aWQifQ==",
  "has_more": true
}
```

- `next_cursor` is opaque to clients. Clients must treat it as an opaque string and never decode or construct it.
- `has_more` indicates whether a subsequent page exists.
- Pagination is stable for a user reviewing a collection: it must not silently skip or duplicate business items.
- Default ordering is documented per collection; only documented `sort` and `order` values are accepted.

### Pagination rules

- A malformed cursor is `VALIDATION_PAGINATION_INVALID` (400).
- `limit` outside 1-100 is `VALIDATION_PAGINATION_INVALID` (400).
- Collections with a stable ordering (expiry date, meal date, creation time) use keyset pagination; the cursor encodes the last seen ordering key.

## Idempotency Semantics

### Scope

- Applies to retry-sensitive `POST` commands: add pantry item, create meal plan, plan meal, create shopping list, generate shopping list, add shopping item, request recommendation, start conversation, continue conversation, pantry analysis.
- `PUT` and `DELETE` operations are already idempotent by design.

### Header

- Client supplies `Idempotency-Key` on the request.
- Server scopes the key to the authenticated user and endpoint.
- Server retains the key and original result for at least 24 hours.
- A retry with the same key and the same payload returns the original result.
- A retry with the same key and a different payload is `VALIDATION_FIELD_INVALID` (409-conflict semantics) to prevent contradictory state.
- Keys are opaque strings up to 64 characters.

### Guarantees

- Repeating an acceptance, completion, or removal command must not create contradictory business state even without a key.
- Idempotency does not change aggregate boundaries; each command still runs within its owning aggregate's consistency boundary.

## Related Documents

- `docs/api/openapi.yaml`
- `docs/architecture/api-design.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/database-design.md`
