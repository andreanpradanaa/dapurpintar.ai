# DapurPintar AI Milestone List

## Current Status

| Item | Status |
|---|---|
| Current milestone | M6 - API Design |
| Current deliverable | M6-003 Contract Tests and Compatibility Policy |
| Current status | In Review |
| Next action | Review M6 deliverables, then start M7 Backend Foundation (DP-FND-001) |
| Next implementation milestone | M10 - Frontend Development |

Milestone approval tetap mengikuti proses review. Artifact yang sudah dibuat tidak otomatis berarti milestone telah disetujui atau ditutup.

## Roadmap Overview

| ID | Milestone | Focus | Status |
|---|---|---|---|
| M0 | Foundation | Repository, workflow, documentation, OpenCode workspace | Complete |
| M1 | Product Planning | Product vision, scope, personas, roadmap, metrics | Complete |
| M2 | Solution Architecture | Domain, system, API, security, AI, observability, deployment | In Review |
| M3 | UX/UI Design | UX strategy, wireframes, design system, prototype, usability | In Review |
| M4 | Solution Architecture | Architecture refinement and implementation preparation | In Review |
| M5 | Database Design | Schema, migrations, indexes, constraints, query design | In Review |
| M6 | API Design | Detailed schemas, OpenAPI, contract validation | In Review |
| M6 | API Design | Detailed schemas, OpenAPI, contract validation | Planned |
| M7 | Backend Foundation | Go, Fiber, Clean Architecture, module scaffolding | Planned |
| M8 | AI Foundation | AI Gateway implementation, prompts, evaluation foundation | Planned |
| M9 | MVP Features | Backend implementation of approved MVP capabilities | Planned |
| M10 | Frontend Development | Next.js application implementation | Planned |
| M11 | Smart Kitchen Features | Smart pantry and kitchen enhancements | Planned |
| M12 | Advanced AI | Advanced recommendations and AI capabilities | Planned |
| M13 | SaaS Platform | Plans, entitlements, usage, commercial capabilities | Planned |
| M14 | Testing | Quality assurance, performance, security, AI evaluation | Planned |
| M15 | Production | DevOps, deployment hardening, monitoring, recovery | Planned |
| M16 | Launch | Release, onboarding, growth, and post-launch learning | Planned |

## M0 - Foundation

Status: **Complete**

- Repository and project structure.
- Git and branch workflow.
- Documentation standards.
- GitHub templates and project conventions.
- OpenCode workspace and AI development workflow.
- Development scripts and technology baseline.

Reference: `docs/project/m0-signoff.md`

## M1 - Product Planning

Status: **Complete**

- Product Vision and Mission.
- Problem Statement.
- Target Users and User Personas.
- Value Proposition.
- Product Scope and Feature Inventory.
- Product Roadmap and Success Metrics.
- Business Goals, Pricing, Risks, and Assumptions.

Reference: `docs/product/m1-signoff.md`

## M2 - Solution Architecture

Status: **In Review**

| ID | Deliverable | Artifact | Status |
|---|---|---|---|
| M2-001 | Architecture Vision | `docs/architecture/architecture-vision.md` | Complete |
| M2-002 | Architecture Decision Records | `docs/architecture/adr/` | Complete |
| M2-003 | Domain Discovery | `docs/architecture/domain-discovery.md` | Complete |
| M2-004 | Bounded Context | `docs/architecture/bounded-context.md` | Complete |
| M2-005 | Tactical DDD | `docs/architecture/tactical-ddd.md` | Complete |
| M2-006 | Event Storming | `docs/domain/event-storming.md` | Complete |
| M2-007 | System Context Diagram | `docs/architecture/system-context.md` | Complete |
| M2-008 | Container Diagram | `docs/architecture/container-diagram.md` | Complete |
| M2-009 | Component Diagram | `docs/architecture/component-diagram.md` | Complete |
| M2-010 | Database Design | `docs/architecture/database-design.md` | Complete |
| M2-011 | API Design | `docs/architecture/api-design.md` | Complete |
| M2-012 | Authentication and Authorization | `docs/architecture/authentication-authorization.md` | Complete |
| M2-013 | AI Architecture | `docs/architecture/ai-architecture.md` | Complete |
| M2-014 | Observability Architecture | `docs/architecture/observability-architecture.md` | Complete |
| M2-015 | Deployment Architecture | `docs/architecture/deployment-architecture.md` | Complete |
| M2-016 | Architecture Sign-off | `docs/architecture/m2-signoff.md` | In Review |

## M3 - UX/UI Design

Status: **In Review**

| ID | Deliverable | Artifact | Status |
|---|---|---|---|
| M3-001 | UX/UI Design Foundation | `docs/ux/ux-ui-design.md` | Complete |
| M3-002 | MVP Wireframe Specification | `docs/ux/mvp-wireframes.md` | Complete |
| M3-003 | Design System Foundation | `docs/ux/design-system.md` | Complete |
| M3-004 | High-Fidelity Screen Specification | `docs/ux/high-fidelity-screen-spec.md` | Complete |
| M3-005 | Usability Validation Plan | `docs/ux/usability-validation-plan.md` | Complete |
| M3-006 | UX/UI Design Sign-off | `docs/ux/m3-signoff.md` | In Review |

Prototype:

- `design/prototype/index.html`
- `design/prototype/styles.css`
- `design/prototype/app.js`
- `design/prototype/README.md`

## M4 - Solution Architecture Refinement

Status: **In Review**

- M4-001 Implementation Readiness: `docs/architecture/implementation-readiness.md`
- M4-002 Architecture Decision Register: `docs/architecture/m4-decision-register.md`
- M4-003 Implementation Backlog and Technical Spikes: `docs/architecture/implementation-backlog.md`
- M4-004 M5 Blocking Decision Records: `docs/architecture/m4-m5-blocking-decisions.md`
- Resolve architecture decisions deferred from M2.
- Refine implementation boundaries from M3 validation.
- Confirm frontend, backend, database, and AI implementation contracts.
- Prepare implementation backlog and technical spikes.

## M5 - Database Design

Status: **In Review**

| ID | Deliverable | Artifact | Status |
|---|---|---|---|
| M5-001 | Concrete PostgreSQL Schema | `docs/database/m5-schema.md` | In Review |
| M5-002 | Goose Migration Strategy | `docs/database/m5-migrations.md` | In Review |
| M5-003 | SQLC Query Contract | `docs/database/m5-sqlc.md` | In Review |
| M5-004 | Seed and Test Data Strategy | `docs/database/m5-seed-and-test-data.md` | In Review |

Scope:

- Detailed PostgreSQL schema.
- Goose migration strategy.
- Constraints, indexes, and query plans.
- SQLC query contract.
- Seed and test data strategy.

## M6 - API Design

Status: **In Review**

| ID | Deliverable | Artifact | Status |
|---|---|---|---|
| M6-001 | OpenAPI Contract | `docs/api/openapi.yaml` | In Review |
| M6-002 | Error, Validation, Pagination, Idempotency Catalog | `docs/api/m6-error-catalog.md` | In Review |
| M6-003 | Contract Tests and Compatibility Policy | `docs/api/m6-contract-tests.md` | In Review |

Scope:

- Detailed request and response schemas.
- OpenAPI specification.
- Contract tests.
- Generated documentation and client considerations.
- Final pagination, error, authentication, and compatibility validation.

## M7 - Backend Foundation

Status: **Planned**

- Go module and Fiber application.
- Clean Architecture and module structure.
- Configuration, errors, validation, and logging.
- PostgreSQL, Redis, Goose, and SQLC integration.
- Authentication and authorization foundation.
- OpenTelemetry instrumentation foundation.

## M8 - AI Foundation

Status: **Planned**

- AI Gateway implementation.
- OpenAI provider adapter.
- Prompt and policy versioning.
- Structured output validation.
- Safety, timeout, retry, quota, and cost controls.
- AI evaluation harness.

## M9 - MVP Features

Status: **Planned**

- Registration, login, profile, and preferences.
- Pantry CRUD and expiry context.
- Recipe search, detail, and favorites.
- AI recommendation and recommendation-scoped conversation.
- Daily and weekly meal planning.
- Shopping list and automatic generation.

## M10 - Frontend Development

Status: **Planned**

- Next.js application foundation.
- Approved design system implementation.
- Responsive MVP screens.
- API integration and authenticated flows.
- Loading, empty, error, permission, and degraded-AI states.
- Accessibility and frontend testing.

## M11 - Smart Kitchen Features

Status: **Planned**

- Barcode and OCR capabilities.
- Improved pantry intelligence.
- Smart reminders.
- Household and kitchen workflow enhancements after validation.

## M12 - Advanced AI

Status: **Planned**

- Advanced meal planning.
- Ingredient replacement and leftover recommendations.
- Cooking assistant improvements.
- Future nutrition and multimodal AI capabilities subject to scope approval.

## M13 - SaaS Platform

Status: **Planned**

- Plans and entitlements.
- Usage and quota management.
- Premium capabilities.
- Commercial operations and future workspace support.

## M14 - Testing

Status: **Planned**

- Unit, integration, API, contract, and end-to-end testing.
- Performance and load testing.
- Security testing.
- AI quality and regression evaluation.
- Release quality gates and zero critical defect target.

## M15 - Production

Status: **Planned**

- Production deployment hardening.
- CI/CD and release automation.
- Backup, restore, and disaster recovery.
- Operational dashboards, alerting, and incident response.
- Capacity, cost, security, and privacy readiness.

## M16 - Launch

Status: **Planned**

- Release and launch readiness.
- Onboarding and support.
- Product analytics and feedback loops.
- Growth experiments and retention learning.
- Post-launch roadmap revision based on validated evidence.

## Workflow Rule

For every milestone:

1. Define the deliverable and acceptance criteria.
2. Read the approved source documents before making changes.
3. Create the artifact within the milestone scope.
4. Perform internal consistency and quality checks.
5. Commit the result with a focused commit.
6. Submit the artifact for review.
7. Do not mark the milestone Approved or Closed before reviewer approval.

## Related Documents

- `PROJECT_ROADMAP.md`
- `docs/project/m0-signoff.md`
- `docs/product/m1-signoff.md`
- `docs/architecture/m2-signoff.md`
- `docs/ux/m3-signoff.md`
