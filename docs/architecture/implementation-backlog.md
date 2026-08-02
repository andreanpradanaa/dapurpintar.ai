# DapurPintar AI Implementation Backlog and Technical Spikes

## Document Control

| Item | Value |
|---|---|
| Milestone | M4 - Solution Architecture Refinement |
| Deliverable | M4-003 Implementation Backlog and Technical Spikes |
| Status | Draft - Awaiting Cross-Functional Review |
| Parent documents | `docs/architecture/implementation-readiness.md`, `docs/architecture/m4-decision-register.md` |
| Related diagram | `docs/architecture/diagrams/implementation-dependency-map.mmd` |

## Purpose

This document turns the accepted implementation sequence into a concrete, prioritized backlog of implementation issues and technical spikes. It gives each work item an owner, a target milestone, a definition of ready, and a definition of done, and it makes visible which pending architecture decisions must be resolved before an item can start.

It does not replace milestone deliverables. M5, M6, M7, and M8 still produce the schema, API contract, backend foundation, and AI foundation. This document sequences that work so no implementation task silently resolves a cross-functional architecture decision.

## Backlog Principles

- Work is expressed as vertical business slices, not isolated technical layers.
- An item is ready when its owning bounded context, contract, authorization rule, persistence impact, and failure behavior are explicit.
- An item is blocked when a pending decision from the Decision Register affects it; the blocking decision ID is recorded on the item.
- Technical spikes exist to collect evidence needed by a decision, not to bypass the decision owner or the review protocol.
- Backlog items must trace to a product feature, domain rule, API operation, or operational requirement.
- No backlog item may change API semantics, domain ownership, or product scope without returning to the correct review path.

## Backlog Structure

| Field | Meaning |
|---|---|
| ID | Stable backlog identifier, e.g. `DP-BE-001` |
| Type | `Foundation`, `Feature`, `Contract`, `Test`, `Ops`, or `Spike` |
| Vertical slice | Business outcome the item delivers, matching Layer 3 of Implementation Readiness |
| Owning context | Bounded context that owns the behavior |
| Target milestone | Milestone where the item is implemented |
| Blocked by | Decision Register IDs that must be resolved before start |
| Definition of ready | Specific inputs that must exist before the item can start |
| Definition of done | Specific outputs that must exist for the item to close |

## Implementation Backlog

### Contracts and foundation

| ID | Type | Item | Target | Blocked by | Notes |
|---|---|---|---|---|---|
| DP-CON-001 | Contract | M5 PostgreSQL schema, constraints, and Goose migration plan | M5 | M4-DEC-006, M4-DEC-007, M4-DEC-013 | Inputs: Database Design, Decision Register defaults |
| DP-CON-002 | Contract | M5 SQLC query contract and seed/test data strategy | M5 | M4-DEC-006 | Follows M5 schema |
| DP-CON-003 | Contract | M6 OpenAPI contract and request/response schemas | M6 | M4-DEC-008, M4-DEC-009, M4-DEC-003, M4-DEC-004 | Inputs: API Design, UX actions and screen states |
| DP-CON-004 | Contract | M6 error, validation, pagination, and idempotency code catalog | M6 | M4-DEC-008 | Consistent with API Design error strategy |
| DP-CON-005 | Contract | M6 contract tests and compatibility policy | M6 | M4-DEC-008 | Validates the public contract |
| DP-FND-001 | Foundation | Go/Fiber application skeleton with lifecycle, config, and errors | M7 | none | Inputs: M5/M6 contracts |
| DP-FND-002 | Foundation | PostgreSQL, Redis, Goose, SQLC, and OpenTelemetry adapters | M7 | M4-DEC-006 | Redis remains non-authoritative |
| DP-FND-003 | Foundation | Authentication boundary and session transport | M7 | M4-DEC-002, M4-DEC-003, M4-DEC-004 | Consistent with Authentication Architecture |
| DP-AI-001 | Foundation | AI Gateway port and OpenAI provider adapter | M8 | M4-DEC-010 | Inputs: AI Architecture |
| DP-AI-002 | Foundation | Prompt, safety, and structured-output policy versioning | M8 | M4-DEC-011, M4-DEC-013 | Inputs: AI Architecture |
| DP-AI-003 | Foundation | AI evaluation harness and representative scenarios | M8 | M4-DEC-012, M4-DEC-016 | Accepts the AI acceptance rubric |

### Vertical feature slices (Layer 3 order)

| ID | Type | Item | Target | Blocked by | Owning context |
|---|---|---|---|---|---|
| DP-FEAT-001 | Feature | Registration, login, logout, refresh, profile, and preferences | M9 | M4-DEC-001, M4-DEC-002, M4-DEC-003 | Identity and Access, User Context and Preferences |
| DP-FEAT-002 | Feature | Pantry and Pantry Item CRUD, expiry attention view | M9 | M4-DEC-006, M4-DEC-007 | Pantry Management |
| DP-FEAT-003 | Feature | Recipe discovery, detail, and favorites | M9 | M4-DEC-009 | Culinary Knowledge and Recipe Experience |
| DP-FEAT-004 | Feature | Meal Plan and Planned Meal lifecycle | M9 | M4-DEC-006, M4-DEC-007, M4-DEC-008 | Meal Planning |
| DP-FEAT-005 | Feature | Shopping List generation, review, activation, and completion | M9 | M4-DEC-006, M4-DEC-008 | Shopping Optimization |
| DP-FEAT-006 | Feature | Kitchen Recommendation request, present, accept, reject, supersede | M9 | M4-DEC-008, M4-DEC-010, M4-DEC-011, M4-DEC-012 | AI-Assisted Kitchen Decision Support |
| DP-FEAT-007 | Feature | Recommendation-scoped conversation | M9 | M4-DEC-008, M4-DEC-011, M4-DEC-013 | AI-Assisted Kitchen Decision Support |
| DP-FEAT-008 | Feature | AI Pantry Analysis endpoint | M9 | M4-DEC-010, M4-DEC-011 | AI-Assisted Kitchen Decision Support |

### Operational and quality

| ID | Type | Item | Target | Blocked by | Notes |
|---|---|---|---|---|---|
| DP-OPS-001 | Ops | Telemetry correlation, redaction, and alerting baseline | M7/M15 | M4-DEC-017 | Inputs: Observability Architecture |
| DP-QA-001 | Test | Contract, integration, and end-to-end test strategy | M14 | M4-DEC-008 | Inputs: M6 contracts and M9 features |
| DP-QA-002 | Test | AI quality regression evaluation | M14 | M4-DEC-012, M4-DEC-016 | Uses M8 evaluation harness |
| DP-OPS-002 | Ops | Backup, restore, DR, RPO, and RTO verification | M15 | M4-DEC-014, M4-DEC-015 | Tested restore behavior, not aspirational target |
| DP-OPS-003 | Ops | Incident ownership and runbooks | M15 | M4-DEC-017 | Documented before public launch |

## Technical Spikes

Technical spikes collect evidence needed to resolve a pending decision or de-risk a high-uncertainty implementation. A spike has an owner, a duration bound, a concrete question, and an explicit outcome. A spike must not commit production code that changes API semantics, domain ownership, or product scope.

| ID | Question | Owner | Required by | Blocking | Expected outcome |
|---|---|---|---|---|---|
| DP-SPK-001 | What is the concrete PostgreSQL schema shape for accounts, pantry, recipes, meal plans, shopping, and recommendations while preserving aggregate ownership? | Database | M5 | M4-DEC-006 | Draft schema and migration order for review |
| DP-SPK-002 | How should MVP timezone policy be modeled so meal dates, expiry, and daily views agree? | Product + Backend | M5 | M4-DEC-007 | Timezone storage and boundary decision |
| DP-SPK-003 | What is the minimum OpenAI model and capability profile that meets MVP latency, cost, and quality targets? | AI Engineering | M8 | M4-DEC-010 | Model profile and capability recommendation |
| DP-SPK-004 | How should structured output validation and prompt safety gates behave under provider failure and injection attempts? | AI Engineering + Security | M8 | M4-DEC-011 | Validation and safety policy proposal |
| DP-SPK-005 | What is the browser session and CSRF behavior for the MVP on the target cookie/domain topology? | Security + Frontend | M7/M10 | M4-DEC-004 | Session transport and CSRF decision |
| DP-SPK-006 | What retention period for raw prompts and conversations satisfies privacy, evaluation, and product value? | Product + Security | M8/M15 | M4-DEC-013 | Retention policy proposal |
| DP-SPK-007 | Which hosting provider and production environment topology meets MVP deployment and cost constraints? | DevOps + Product | M15 | M4-DEC-014 | Provider and environment recommendation |
| DP-SPK-008 | What RPO and RTO can the chosen environment actually meet under a tested restore? | DevOps + Product | M15 | M4-DEC-015 | Tested restore evidence and target values |
| DP-SPK-009 | How should AI quota and cost budgets be enforced and alerted to avoid cost exhaustion? | Product + Finance + AI | M8/M15 | M4-DEC-016 | Quota and cost control proposal |

## Backlog Sequencing

### M5 - Database Design

Start DP-CON-001 after DP-SPK-001 and DP-SPK-002 resolve M4-DEC-006 and M4-DEC-007. DP-CON-002 follows DP-CON-001.

### M6 - API Design

Start DP-CON-003 after M4-DEC-008, M4-DEC-009, M4-DEC-003, and M4-DEC-004 are resolved or their defaults are confirmed. DP-CON-004 and DP-CON-005 follow.

### M7 - Backend Foundation

DP-FND-001 can start once M6 contracts exist. DP-FND-002 requires the M5 schema. DP-FND-003 requires the authentication contract decisions from M4-DEC-002, M4-DEC-003, and M4-DEC-004.

### M8 - AI Foundation

DP-AI-001 starts after M4-DEC-010 and the AI Architecture baseline. DP-AI-002 requires M4-DEC-011 and M4-DEC-013. DP-AI-003 requires M4-DEC-012 and M4-DEC-016.

### M9 - MVP Features

Vertical slices start in Layer 3 order: DP-FEAT-001 through DP-FEAT-008. Each feature inherits the blocking decisions listed in the backlog table and must meet the Cross-Cutting Definition of Ready from Implementation Readiness before implementation begins.

### M10 - Frontend Development

Frontend implementation starts only after M3 usability validation findings are accepted and M6 contracts exist. It consumes the approved design system and M3 prototype rather than creating ad hoc UI.

## Definition of Ready

An implementation issue is ready to start when:

- Scope and user outcome are clear.
- The owning bounded context is named.
- The API operation or internal contract is identified.
- Authorization and user scope are explicit.
- Persistence impact is known or intentionally deferred.
- Error, empty, loading, and degraded behavior is defined.
- Telemetry signals are identified.
- Test scenarios are listed.
- Dependencies and follow-up decisions are recorded.
- No blocking Decision Register item remains unresolved.

## Definition of Done

An implementation issue is done when:

- Domain behavior and invariants are tested.
- API and persistence adapters follow approved contracts.
- Ownership and authorization are enforced server-side.
- No implicit cross-context mutation is introduced.
- Error and degraded states are safe and actionable.
- Telemetry is correlated and redacted.
- Relevant tests pass.
- Documentation and decision records are updated.
- The change is reviewable as one focused unit.

## Exit Criteria

M4-003 is complete when:

- Every Layer 3 vertical slice has a backlog item with owner, target, and decision dependencies.
- Technical spikes are scoped with an owner, a duration bound, and an expected outcome.
- Backlog items reference the blocking Decision Register IDs they depend on.
- No implementation item can start while a required decision is still unresolved.
- The backlog is traceable to product features, domain rules, API operations, and operational requirements.

## Related Documents

- `docs/architecture/implementation-readiness.md`
- `docs/architecture/m4-decision-register.md`
- `docs/architecture/diagrams/implementation-dependency-map.mmd`
- `docs/architecture/api-design.md`
- `docs/architecture/database-design.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/ai-architecture.md`
- `docs/architecture/observability-architecture.md`
- `docs/architecture/deployment-architecture.md`
- `docs/ux/m3-signoff.md`
