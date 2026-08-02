# Milestone 2 Architecture Sign-off

## Document Control

| Item | Value |
|---|---|
| Milestone | M2 - Solution Architecture |
| Status | Draft - Awaiting Software Architect Review |
| Owner | Solution Architecture |
| Scope | Strategic domain design, system design, API, security, AI, observability, and deployment architecture |
| Next milestone | M3 - UX/UI Design |

## Purpose

This document records the completion state, architectural decisions, unresolved risks, and readiness of Milestone 2. It is the review gate between the approved product planning foundation and the UX/UI and later implementation milestones.

M2 defines the business and technical boundaries that future product, backend, frontend, AI, QA, and DevOps work must follow. It does not represent implementation completion. No M2 deliverable is considered finally approved until the Software Architect review is completed.

## Executive Summary

Milestone 2 established the architecture foundation for DapurPintar AI. The system is defined as a modular monolith with Clean Architecture and Domain Driven Design. The AI-Assisted Kitchen Decision Support context is the Core Domain, while User Context and Preferences, Pantry Management, Culinary Knowledge and Recipe Experience, Meal Planning, and Shopping Optimization provide supporting business capabilities.

The architecture uses PostgreSQL as the authoritative system of record, Redis for non-authoritative supporting capabilities, a versioned REST/JSON API beginning with `/api/v1`, secure authentication and ownership-based authorization, an AI Gateway abstraction around OpenAI, OpenTelemetry and Grafana for observability, and a vendor-neutral deployment model designed for MVP simplicity and recovery.

The recommended decision is **GO to M3**, subject to Software Architect approval and completion of the implementation-readiness actions listed in this document. M2 remains open until that review is recorded.

## M2 Deliverables

| ID | Deliverable | Primary artifact | Status |
|---|---|---|---|
| M2-001 | Architecture Vision | `docs/architecture/architecture-vision.md` | Complete - awaiting review |
| M2-002 | Architecture Decision Records | `docs/architecture/adr/` | Complete - awaiting review |
| M2-003 | Domain Discovery and Strategic DDD | `docs/architecture/domain-discovery.md` | Complete - awaiting review |
| M2-004 | Bounded Context and Context Map | `docs/architecture/bounded-context.md` | Complete - awaiting review |
| M2-005 | Tactical DDD | `docs/architecture/tactical-ddd.md` | Complete - awaiting review |
| M2-006 | Event Storming | `docs/domain/event-storming.md` | Complete - awaiting review |
| M2-007 | System Context Diagram | `docs/architecture/system-context.md` | Complete - awaiting review |
| M2-008 | Container Diagram | `docs/architecture/container-diagram.md` | Complete - awaiting review |
| M2-009 | Component Diagram | `docs/architecture/component-diagram.md` | Complete - awaiting review |
| M2-010 | Database Design | `docs/architecture/database-design.md` | Complete - awaiting review |
| M2-011 | API Design | `docs/architecture/api-design.md` | Complete - awaiting review |
| M2-012 | Authentication and Authorization | `docs/architecture/authentication-authorization.md` | Complete - awaiting review |
| M2-013 | AI Architecture | `docs/architecture/ai-architecture.md` | Complete - awaiting review |
| M2-014 | Observability Architecture | `docs/architecture/observability-architecture.md` | Complete - awaiting review |
| M2-015 | Deployment Architecture | `docs/architecture/deployment-architecture.md` | Complete - awaiting review |
| M2-016 | Architecture Sign-off | This document | In review |

Supporting diagrams are maintained beside their related architecture documents under `docs/architecture/diagrams/`.

## Architectural Decisions Confirmed

### Business and domain

- AI-Assisted Kitchen Decision Support is the Core Domain.
- AI proposes and explains cooking decisions; the user remains the final decision-maker.
- Each bounded context owns its language, business rules, aggregate lifecycle, and authoritative meaning.
- Cross-context references and recommendations do not create implicit cross-context mutations.
- Household collaboration, nutrition, notifications, commercial operations, and partner capabilities remain future extensions.

### Application structure

- The MVP is a Modular Monolith.
- Clean Architecture keeps domain and application rules independent from Fiber, PostgreSQL, Redis, OpenAI, and Next.js.
- The Backend Application is one logical deployable unit with explicit module contracts and persistence boundaries.
- Domain events and application contracts may coordinate modules without requiring a message broker.

### Data

- PostgreSQL is the authoritative store for durable identity and business state.
- Redis supports cache, rate limiting, session coordination, and bounded transient work only.
- Aggregate ownership prevents modules from writing another module's authoritative tables directly.
- AI and historical context are retained only where product value, evaluation, privacy, and retention policy justify it.

### API

- The public contract is versioned REST/JSON beginning with `/api/v1`.
- API resources represent business capabilities and bounded-context ownership, not database tables or provider payloads.
- Protected resources require authenticated identity and server-derived ownership authorization.
- The contract defines stable errors, response envelopes, cursor pagination, sorting, filtering, idempotency, and request correlation.
- AI output remains proposal data until explicit user acceptance and a later domain command.

### Security

- Authentication uses secure password handling, short-lived JWT access tokens, and rotating refresh sessions.
- Browser credentials use secure, HttpOnly, SameSite-aware cookies with CSRF protection for state-changing requests.
- Authorization is evaluated in the application and owning bounded context, never trusted from the frontend or client ownership fields.
- Account restriction, session revocation, rate limiting, auditability, and secret management are mandatory boundaries.

### AI

- OpenAI is accessed only through the AI Gateway and provider adapter.
- Context is authorized, purpose-specific, minimized, and translated into provider-independent requests.
- Structured output, source validation, business validation, safety validation, and product-level response mapping are required.
- AI failure must not make core non-AI CRUD operations unavailable where graceful degradation is possible.

### Operations

- OpenTelemetry is the standard for traces, metrics, and structured logs.
- Grafana provides dashboards and alerting for API, dependency, security, AI, and product operations.
- Deployment remains vendor-neutral, with explicit environment, network, secret, release, backup, recovery, and scaling boundaries.
- Production readiness requires tested restore, migration, rollback, health, and dependency-failure procedures.

## Architecture Quality Review

| Quality attribute | M2 architectural response | Review status |
|---|---|---|
| Usability | API and domain contracts preserve direct user actions and explicit commitments | Ready for review |
| Performance | Bounded queries, cursor pagination, stateless application scaling, and AI deadlines | Ready for review |
| Availability | Modular monolith simplicity, dependency degradation, health checks, and 99.9% target | Ready for review |
| Security | Authentication, ownership authorization, rate limiting, secret protection, and audit signals | Ready for review |
| Privacy | Context minimization, telemetry redaction, provider isolation, retention boundaries | Ready for review |
| Maintainability | DDD modules, Clean Architecture, ports, SQLC, Goose, and explicit ownership | Ready for review |
| Scalability | Stateless replicas, PostgreSQL optimization, bounded AI concurrency, measured extraction | Ready for review |
| Testability | Domain, application, API, integration, security, AI evaluation, and recovery test seams | Ready for review |
| Operability | OpenTelemetry, Grafana, dashboards, alerts, deployment signals, and runbook requirements | Ready for review |
| Recoverability | PostgreSQL backup, restore verification, migration discipline, and RPO/RTO gate | Conditional before production |

## Open Issues and Implementation Gates

The following items are intentionally deferred from M2 detail but must be resolved before the relevant implementation or launch gate:

- Define the complete request and response schemas in the API specification.
- Finalize email verification, password recovery, credential-change, and MFA policy.
- Finalize access-token and refresh-session lifetimes, cookie domain, SameSite mode, and CSRF mechanism.
- Define concrete PostgreSQL schema, indexes, constraints, migrations, and query plans.
- Select the hosting provider, deployment environment, operational regions, and telemetry storage.
- Approve production RPO and RTO targets and complete restore testing.
- Select initial OpenAI models, prompt revisions, safety policy revisions, and AI evaluation datasets.
- Define final retention, deletion, export, consent, and regional privacy behavior.
- Define production capacity limits, AI quota budgets, and cost alerts.
- Define incident response owners and runbooks for API, database, Redis, AI, security, and telemetry failures.

These issues are not architectural contradictions. They are implementation, product-policy, deployment, or operational decisions that must follow the boundaries established by M2.

## Accepted Risks

- MVP requirements may change after UX validation and beta feedback.
- AI quality, provider behavior, latency, quota, and cost may require prompt or model adjustments.
- Pantry data entered by users may be incomplete or stale and reduce recommendation relevance.
- Future household and commercial scopes may require new authorization and ownership models.
- A modular monolith may require later extraction if measured scale or team boundaries justify it.
- Vendor-neutral deployment decisions may be refined after infrastructure validation.
- Observability and AI evaluation require careful sampling and privacy controls to avoid cost or data exposure.

No accepted risk permits bypassing user ownership, domain authority, AI safety, or PostgreSQL source-of-truth rules.

## Exit Criteria

M2 is ready to close when:

- All M2-001 through M2-015 deliverables exist and are internally consistent.
- Architecture decisions have traceability to product requirements and accepted ADRs.
- Domain ownership, aggregate boundaries, API ownership, and authorization scope are explicit.
- PostgreSQL authority and Redis limitations are preserved across all documents.
- AI remains a bounded application capability and cannot silently create commitments.
- Observability, deployment, privacy, recovery, and security boundaries are documented.
- No critical architectural contradiction or unresolved blocker prevents M3 or implementation planning.
- Software Architect reviews and approves this sign-off.

## Go / No-Go Recommendation

### Recommendation

**GO - recommended, pending Software Architect approval.**

M2 provides sufficient architectural direction to begin M3 UX/UI Design and to prepare subsequent implementation planning. The open issues above remain controlled follow-up gates and do not require reopening M2 unless review identifies a contradiction with product scope or domain ownership.

### Current approval state

| Role | Status |
|---|---|
| Product Owner | Pending review |
| Software Architect | Pending review |
| Engineering | Pending review |
| AI Engineering | Pending review |
| Security | Pending review |
| QA | Pending review |
| DevOps | Pending review |

This document must not be changed to Approved until the designated reviewer confirms the M2 architecture and records any required follow-up decisions.

## Lessons Learned

- Architecture documents are most useful when each one declares ownership and non-goals.
- AI boundaries must be expressed as product and domain rules, not only provider integration details.
- API, authentication, AI, observability, and deployment decisions need explicit cross-document consistency checks.
- PostgreSQL authority and user decision control are recurring invariants across the architecture.
- Vendor-neutral deployment detail is sufficient for M2 while avoiding premature infrastructure commitment.
- OpenCode works more reliably when each milestone has a bounded deliverable, acceptance criteria, and review gate.

## Next Milestone

**M3 - UX/UI Design**

The next milestone should translate the approved product scope and architecture boundaries into:

- User journeys and information architecture.
- UX flows for onboarding, pantry, recipes, recommendations, meal planning, and shopping.
- UI design system and responsive interaction patterns.
- Error, loading, empty, degraded-AI, and permission states.
- Frontend-to-API interaction assumptions that remain consistent with `/api/v1` and ownership rules.

M3 must not introduce frontend behavior that bypasses backend authorization, creates implicit cross-context mutations, or treats AI output as authoritative business state.

## Related Documents

- `PROJECT_ROADMAP.md`
- `docs/product/m1-signoff.md`
- `docs/architecture/architecture-vision.md`
- `docs/architecture/architecture-context.md`
- `docs/architecture/domain-discovery.md`
- `docs/architecture/bounded-context.md`
- `docs/architecture/tactical-ddd.md`
- `docs/domain/event-storming.md`
- `docs/architecture/database-design.md`
- `docs/architecture/api-design.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/ai-architecture.md`
- `docs/architecture/observability-architecture.md`
- `docs/architecture/deployment-architecture.md`
