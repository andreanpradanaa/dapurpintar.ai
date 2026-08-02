# DapurPintar AI Implementation Readiness

## Document Control

| Item | Value |
|---|---|
| Milestone | M4 - Solution Architecture Refinement |
| Status | Draft - Awaiting M2/M3 Review Closure |
| Scope | Implementation sequencing, dependency readiness, and decision gates |
| Parent documents | `docs/architecture/m2-signoff.md`, `docs/ux/m3-signoff.md`, `docs/project/milestone-list.md` |
| Related diagram | `docs/architecture/diagrams/implementation-dependency-map.mmd` |

## Purpose

This document translates the approved architecture and UX direction into an implementation-readiness map. It identifies what is ready, what is blocked by a decision, what must be implemented first, and which team owns the next action.

It does not replace detailed database design, OpenAPI schema design, technical task breakdown, or code implementation. It prevents the team from starting a downstream feature while its domain, contract, security, data, or UX dependency is still ambiguous.

## Readiness Summary

| Area | Readiness | Gate |
|---|---|---|
| Product scope | Directionally ready | M1 approval and M3 usability findings |
| Domain boundaries | Ready for implementation refinement | M2 architecture approval |
| UX direction | Prototype ready for validation | M3 usability validation |
| API resource vocabulary | Ready | Detailed schemas in M6 |
| Authentication model | Directionally ready | Final session, recovery, and CSRF policy |
| Database model | Conceptually ready | Detailed schema, migrations, indexes in M5 |
| AI architecture | Directionally ready | Model, prompt, safety, and evaluation decisions in M8 |
| Observability | Contract ready | Runtime and storage topology in M15 |
| Deployment | Logical topology ready | Provider, environment, RPO, and RTO decisions |
| Backend implementation | Not started | M7 foundation gates |
| Frontend implementation | Intentionally deferred | M3 approval and M10 |

## Implementation Principles

- Complete the smallest decision needed to unblock the next bounded slice.
- Implement vertical business capability boundaries, not isolated technical layers with no user outcome.
- Keep domain behavior testable without HTTP, PostgreSQL, Redis, OpenAI, or a browser.
- Keep API contracts ahead of client and server coupling.
- Keep migrations forward-compatible with the release sequence.
- Treat AI as an optional decision-support dependency, never as a source of truth.
- Preserve user ownership and explicit commitment actions in every implementation.
- Do not introduce a service, queue, or platform component before a measured need exists.
- Make every implementation task traceable to a product feature, domain rule, API operation, or operational requirement.

## Dependency Layers

### Layer 0: Review gates

Before implementation begins:

- M2 architecture sign-off is reviewed.
- M3 UX/UI sign-off is reviewed.
- M3 prototype validation is planned or completed according to the release decision.
- Deferred architecture decisions are assigned to an owner and milestone.

### Layer 1: Shared backend foundation

The backend foundation provides:

- Go module and application lifecycle.
- Fiber API boundary.
- Configuration and secret loading.
- Standard errors and response envelopes.
- Request correlation and OpenTelemetry base instrumentation.
- PostgreSQL connection and migration integration.
- Redis adapter and safe failure behavior.
- Authentication and authorization boundary.

No business module should create its own incompatible version of these concerns.

### Layer 2: Durable business contracts

Detailed data and API contracts establish:

- PostgreSQL schema and ownership constraints.
- Goose migration order.
- SQLC query contracts.
- Request and response schemas.
- Validation and error codes.
- Pagination, idempotency, and authorization behavior.

Layer 2 must preserve the conceptual ownership already defined by DDD and Database Design.

### Layer 3: Core vertical capabilities

The first user-visible backend slices should follow the product journey:

1. Account registration, login, profile, and preferences.
2. Pantry item creation, adjustment, removal, category, and expiry.
3. Recipe discovery, detail, and favorites.
4. Meal Plan creation and Planned Meal actions.
5. Shopping List creation, generation, review, and completion.
6. Recommendation request, presentation, option acceptance, rejection, and scoped conversation.

Each slice includes domain behavior, application use case, repository contract, API mapping, authorization, telemetry, tests, and failure behavior before moving to the next dependent slice.

### Layer 4: Frontend implementation

Frontend implementation consumes approved API contracts and M3 design artifacts. It must not create a parallel domain model or infer authorization from visual state.

The frontend starts with the approved shell, design tokens, core components, and validated Today-to-Pantry-to-Recommendation-to-Plan/Shop journey. M10 remains the implementation milestone in the current roadmap.

## Team Workstreams

| Workstream | Primary responsibility | First dependency |
|---|---|---|
| Product | Scope, acceptance criteria, unresolved user decisions | M3 validation findings |
| Architecture | Boundary decisions, ADRs, cross-document consistency | M2/M3 review closure |
| Backend | Go foundation and vertical business capabilities | M5/M6 contracts and M7 |
| Database | Schema, migrations, indexes, query performance | Database conceptual design |
| AI Engineering | Gateway, provider adapter, prompts, evaluation, safety | AI architecture and approved use cases |
| Frontend | Next.js implementation and API integration | M3 artifacts and M6 contract |
| QA | Test strategy, contract, integration, usability, and release gates | Acceptance criteria and contracts |
| DevOps | Environments, CI/CD, secrets, monitoring, backup, recovery | Deployment architecture |

## Readiness Gates by Roadmap Milestone

### M4 - Architecture refinement

Ready when:

- M2 and M3 review findings are recorded.
- Deferred decisions have owners and target milestones.
- Implementation sequence is accepted.
- No UX change silently alters domain or API ownership.

### M5 - Detailed database design

Required inputs:

- Approved aggregates and ownership.
- API resource vocabulary.
- Authentication session persistence needs.
- AI Recommendation and Conversation retention policy.

Required outputs:

- Concrete PostgreSQL schema.
- Constraints and indexes.
- Goose migration plan.
- SQLC query boundaries.
- Seed and test data strategy.

### M6 - Detailed API design

Required inputs:

- UX actions and screen states.
- Resource and command catalog.
- Authentication and authorization rules.
- Pagination, error, and idempotency conventions.

Required outputs:

- OpenAPI or equivalent contract.
- Request and response schemas.
- Validation and error code catalog.
- Contract tests and compatibility policy.

### M7 - Backend foundation

Required inputs:

- M5 schema and migration decisions.
- M6 API contract.
- Authentication architecture.
- Deployment configuration boundaries.

Required outputs:

- Running Go/Fiber application skeleton.
- Shared platform components.
- Database, Redis, telemetry, and authentication adapters.
- Test and local development workflow.

### M8 - AI foundation

Required inputs:

- AI use cases and Recommendation lifecycle.
- Context ownership and authorization.
- Provider abstraction and deployment secret policy.
- AI quality and safety metrics.

Required outputs:

- AI Gateway port and adapter.
- Structured output schemas.
- Prompt and policy revisions.
- Safety, retry, timeout, quota, and cost controls.
- Evaluation harness and representative test scenarios.

### M9 - MVP backend features

Each feature is ready for implementation only when its vertical slice includes:

- User story and acceptance criteria.
- Domain command or query.
- Aggregate and invariant ownership.
- API operation and schema.
- Database query and migration impact.
- Authorization rule.
- Telemetry and audit signals.
- Unit, integration, and contract test plan.
- Failure and degraded behavior.

### M10 - Frontend development

Required inputs:

- M3 usability findings and approved design.
- M6 API contract.
- Authentication transport and error behavior.
- Accessibility and responsive requirements.

Required outputs:

- Next.js application shell.
- Reusable design-system components.
- MVP screens and journeys.
- API integration and state handling.
- Frontend accessibility and end-to-end tests.

## Decision Register

| Decision | Current direction | Owner | Gate |
|---|---|---|---|
| API prefix | `/api/v1` | Architecture | Confirmed |
| Backend style | Modular monolith | Architecture | Confirmed |
| System of record | PostgreSQL | Architecture/Database | Confirmed |
| Supporting cache | Redis, non-authoritative | Backend/DevOps | Confirmed |
| AI provider boundary | AI Gateway around OpenAI | AI Engineering | Confirmed |
| Authentication | JWT access plus rotating refresh session | Security/Backend | Detail pending |
| Browser session transport | Secure HttpOnly SameSite cookies | Security/Frontend | Detail pending |
| API payload schemas | Detailed in M6 | API/Backend | Pending |
| PostgreSQL schema | Detailed in M5 | Database | Pending |
| Hosting provider | Vendor-neutral | DevOps | Pending |
| Production RPO/RTO | Must be approved before launch | DevOps/Product | Pending |
| Recommendation acceptance | Specific option, explicit user action | Product/Domain | Confirmed |

## Cross-Cutting Definition of Ready

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

## Cross-Cutting Definition of Done

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

## Risks and Assumptions

### Risks

- Starting M7 before M5/M6 contracts may create avoidable rework.
- Starting M10 before M3 validation may encode misunderstood user behavior.
- Deferred authentication policy may affect frontend session behavior.
- AI model and prompt decisions may alter latency, cost, and screen expectations.
- Detailed schema work may reveal aggregate or ownership ambiguities.
- M2 and M3 remaining in review may create parallel contradictory decisions.

### Assumptions

- The existing M2 and M3 artifacts remain the working baseline until review findings change them.
- M4 is refinement and implementation preparation, not permission to bypass milestone gates.
- M5 and M6 provide the concrete contracts needed before M7 foundation implementation.
- M10 consumes the M3 prototype and approved design system rather than replacing them with ad hoc UI.

## Exit Criteria

M4 implementation readiness is complete when:

- All deferred M2/M3 decisions have owners and target gates.
- The workstream sequence and dependency map are accepted.
- M5-M10 inputs and outputs are explicit.
- Definition of Ready and Definition of Done are adopted for implementation issues.
- No team begins implementation by bypassing domain ownership, API contracts, authorization, or UX validation.

## Related Documents

- `docs/project/milestone-list.md`
- `docs/architecture/m2-signoff.md`
- `docs/ux/m3-signoff.md`
- `docs/architecture/architecture-vision.md`
- `docs/architecture/component-diagram.md`
- `docs/architecture/api-design.md`
- `docs/architecture/database-design.md`
- `docs/architecture/ai-architecture.md`
- `docs/architecture/deployment-architecture.md`
- `docs/ux/usability-validation-plan.md`
- `docs/architecture/diagrams/implementation-dependency-map.mmd`
