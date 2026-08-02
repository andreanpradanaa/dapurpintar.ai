# DapurPintar AI M4 Architecture Decision Register

## Document Control

| Item | Value |
|---|---|
| Milestone | M4 - Solution Architecture Refinement |
| Deliverable | M4-002 Architecture Decision Register |
| Status | Draft - Awaiting Cross-Functional Review |
| Parent document | `docs/architecture/implementation-readiness.md` |

## Purpose

This register tracks decisions deferred by M2 and M3 that can affect implementation sequencing. It prevents unresolved questions from being silently decided inside code, UI, migration, prompt, or deployment work.

The register distinguishes architectural direction from implementation detail. A decision is only marked `Decided` when the owner, affected disciplines, and required evidence have reviewed it.

## Status Definitions

| Status | Meaning |
|---|---|
| Confirmed | Existing architecture or product decision is sufficient and should be preserved |
| Pending | A decision is required before the target gate |
| Validating | Prototype, test, or operational evidence is being collected |
| Blocked | Work cannot safely proceed until the decision is made |
| Accepted exception | Work may proceed with an explicit, time-bounded exception |

## Confirmed Decisions

| ID | Decision | Source | Affected work |
|---|---|---|---|
| ADR-001 | Go is the backend runtime | ADR-001 | M7 and later |
| ADR-002 | Fiber is the HTTP framework | ADR-002 | M7 and later |
| ADR-003 | PostgreSQL is the system of record | ADR-003 | M5, M7, M9 |
| ADR-004 | MVP backend is a modular monolith | ADR-004 | M4-M15 |
| ADR-005 | Clean Architecture dependency direction | ADR-005 | M7 and later |
| ADR-006 | Domain Driven Design and bounded ownership | ADR-006 | M4-M10 |
| ADR-007 | Versioned REST API beginning `/api/v1` | ADR-007 and API Design | M6, M7, M10 |
| ADR-008 | SQLC for reviewed type-safe SQL access | ADR-008 | M5, M7 |
| ADR-009 | OpenTelemetry and Grafana observability | ADR-009 | M7, M15 |
| ADR-010 | AI Gateway abstraction around OpenAI | ADR-010 | M8, M9 |
| UX-001 | Context Strip is the signature UX element | M3 UX/UI Design | M10 |
| UX-002 | AI acceptance, planning, shopping, and pantry changes are distinct actions | M3 UX/UI Design | M6, M9, M10 |

These decisions should not be reopened during implementation without a new ADR or explicit milestone review.

## Pending Decision Register

| ID | Decision needed | Owner | Target gate | Status | Impact if delayed |
|---|---|---|---|---|---|
| M4-DEC-001 | Email verification required for MVP? | Product + Security | M7 Auth foundation | Decided | Account lifecycle and onboarding ambiguity |
| M4-DEC-002 | Password recovery and credential-change flow | Product + Security | M7 Auth foundation | Decided | Cannot complete account lifecycle safely |
| M4-DEC-003 | Access-token and refresh-session lifetimes | Security + Backend | M7 Auth foundation | Decided | Cookie, refresh, logout, and test behavior remain undefined |
| M4-DEC-004 | Final SameSite, cookie domain, and CSRF policy | Security + Frontend | M7 and M10 | Decided | Browser session implementation may be reworked |
| M4-DEC-005 | MFA requirement before public launch | Product + Security | M15/M16 | Pending | Launch security scope remains incomplete |
| M4-DEC-006 | Concrete PostgreSQL schema and migration shape | Database + Backend | M5 | Decided | M7 repositories and M9 features cannot safely start |
| M4-DEC-007 | MVP timezone policy | Product + Backend | M5 | Decided | Meal dates, expiry, and daily views may disagree |
| M4-DEC-008 | Detailed API request/response schemas | API + Backend + Frontend | M6 | Decided | Client and server contract cannot be implemented reliably |
| M4-DEC-009 | Public recipe access boundary | Product + Security | M6 | Decided | Public and authenticated UI/API behavior may diverge |
| M4-DEC-010 | Initial OpenAI model and capability profile | AI Engineering + Product | M8 | Decided | Latency, cost, output, and evaluation targets remain uncertain |
| M4-DEC-011 | Prompt, safety, and structured-output revisions | AI Engineering + Security | M8 | Decided | AI implementation cannot pass reliable safety and quality gates |
| M4-DEC-012 | AI evaluation dataset and acceptance rubric | AI Engineering + Product | M8 | Decided | Recommendation quality cannot be measured consistently |
| M4-DEC-013 | Raw prompt and conversation retention policy | Product + Security | M8/M15 | Decided | Privacy and reproducibility tradeoff remains unresolved |
| M4-DEC-014 | Hosting provider and production environment | DevOps + Product | M15 | Pending | Deployment, secrets, and recovery cannot be finalized |
| M4-DEC-015 | Production RPO and RTO | DevOps + Product | M15 | Pending | Backup and recovery readiness cannot be approved |
| M4-DEC-016 | AI quota and cost budget | Product + Finance + AI | M8/M15 | Decided | Cost exhaustion and plan limits remain uncontrolled |
| M4-DEC-017 | Incident response owners and runbooks | DevOps + Security + QA | M15 | Pending | Production failures lack accountable response |
| M4-DEC-018 | Today default after onboarding | Product + Design | M3 validation | Validating | First-session flow may be revised |
| M4-DEC-019 | Indonesian wording for option acceptance | Product + Design + Research | M3 validation | Validating | Users may misunderstand commitment boundaries |
| M4-DEC-020 | Planner layout: grid or day stack | Product + Design + Research | M3 validation | Validating | Responsive interaction may be redesigned |

## Decision Requirements

### M4 architecture refinement

M4 must not decide every pending item. It must make ownership and timing explicit, identify blockers, and prevent accidental decisions in downstream work.

M4 outputs:

- Each pending item has an owner.
- Each pending item has a target milestone or launch gate.
- Blockers for M5, M6, M7, M8, M10, and M15 are visible.
- Decisions that require user or operational evidence are marked `Validating`.
- Decisions that change API, domain ownership, or product scope trigger the correct review path.

### Decision record format

When a decision is made, record:

- Decision ID.
- Context and problem.
- Options considered.
- Chosen direction.
- Consequences and risks.
- Affected documents, contracts, or screens.
- Owner and approval date.
- Revisit condition.

An implementation issue must reference the decision ID when its behavior depends on a pending or confirmed decision.

## Blocking Rules

### Blocks M5

- M4-DEC-006 concrete PostgreSQL schema.
- M4-DEC-007 timezone policy.
- M4-DEC-013 retention policy for stored AI context where schema is affected.

### Blocks M6

- M4-DEC-008 API schemas.
- M4-DEC-009 public recipe access.
- M4-DEC-003 and M4-DEC-004 authentication contract details.

### Blocks M7

- M4-DEC-002 credential recovery boundary.
- M4-DEC-003 token and session lifetimes.
- M4-DEC-004 browser session and CSRF policy.
- M4-DEC-006 schema and migration contract.
- M4-DEC-008 API schemas.

### Blocks M8

- M4-DEC-010 model capability profile.
- M4-DEC-011 prompt, safety, and structured-output policy.
- M4-DEC-012 evaluation dataset and rubric.
- M4-DEC-016 AI quota and cost budget.

### Blocks M10

- M4-DEC-008 API schemas.
- M4-DEC-004 session transport and CSRF behavior.
- M4-DEC-018 and M4-DEC-019 if they affect primary onboarding or recommendation actions.

### Blocks M15

- M4-DEC-014 hosting and production environment.
- M4-DEC-015 RPO and RTO.
- M4-DEC-016 AI budget and cost alerts.
- M4-DEC-017 incident ownership and runbooks.

## Current Recommended Defaults

The following are safe working assumptions for planning only. They are not final decisions until the owner records them:

- MVP email verification may remain optional for early local development but must be decided before public launch.
- Access tokens should remain short-lived and refresh sessions revocable, consistent with Authentication Architecture.
- PostgreSQL should store durable session and revocation authority; Redis must remain supporting state.
- Public recipe access should return only general recipe content and never personalized context.
- The initial AI provider should use one approved model profile before considering multi-provider routing.
- Raw prompts and conversations should be minimized and retained only where evaluation or product value justifies it.
- RPO and RTO should be based on tested restore behavior rather than an aspirational target.
- UX wording and Planner layout should be validated with representative users before frontend lock-in.

## Recently Decided

- M4-DEC-010, M4-DEC-011, M4-DEC-012, M4-DEC-016 were approved via `docs/architecture/m8-blocking-decisions.md` (M4-005). Recommended directions: fixed default model profile with versioned alternatives; versioned prompt/safety/structured-output pipeline with layered validation; privacy-safe evaluation dataset with published rubric and pass gates; explicit per-user and global AI budgets with enforced limits and alerts.
- M4-DEC-001, M4-DEC-002, M4-DEC-003, M4-DEC-004, M4-DEC-006, M4-DEC-007, M4-DEC-008, M4-DEC-009, and M4-DEC-013 were approved as part of the M2-M8 milestone review sign-off. Chosen directions follow the "Current Recommended Defaults" below:
  - M4-DEC-001: Email verification stays optional for local development; it becomes required before public launch.
  - M4-DEC-002: Credential change (authenticated password update) is in MVP scope; forgotten-password recovery is deferred to pre-launch (M15/M16).
  - M4-DEC-003: Short-lived access sessions (default 15m) with revocable refresh sessions (default 30d); revocation authority stored in PostgreSQL.
  - M4-DEC-004: Session cookie `dp_session` with `HttpOnly`, `Secure` in production, `SameSite=Lax`; state-changing endpoints additionally verify the request origin to protect against CSRF.
  - M4-DEC-006: The concrete PostgreSQL schema and migration shape follow `docs/database/m5-schema.md` and `docs/database/m5-migrations.md` (M5 sign-off).
  - M4-DEC-007: Store UTC; interpret in the user timezone with `Asia/Jakarta` as the default (M5 sign-off).
  - M4-DEC-008: The API contract follows `docs/api/openapi.yaml` and `docs/api/m6-error-catalog.md` (M6 sign-off).
  - M4-DEC-009: Public recipe access returns only general recipe content; personalized context is never served to unauthenticated callers.
  - M4-DEC-013: Recommendation conversations are retained only while the recommendation is active; raw prompts are never stored; retained context uses a 30-day window (M5/M8 sign-off).
- M4-DEC-005 (MFA) and M4-DEC-014/015/017 remain Pending and gate M15/M16; they do not block M9 implementation.

## Review Protocol

1. Owner proposes a decision and records evidence.
2. Affected disciplines review consequences.
3. Architecture checks compatibility with bounded ownership and existing ADRs.
4. Product confirms scope and user impact.
5. Decision is recorded and linked from implementation work.
6. Any changed API, domain, privacy, or deployment contract is updated in its source document.

## Exit Criteria

M4-002 is complete when:

- All deferred M2/M3 decisions are listed.
- Every item has an owner and target gate.
- Blocking dependencies for M5-M10 and M15 are explicit.
- Confirmed decisions are separated from working assumptions.
- No implementation task can silently resolve a cross-functional architecture decision.

## Related Documents

- `docs/architecture/implementation-readiness.md`
- `docs/architecture/m2-signoff.md`
- `docs/ux/m3-signoff.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/database-design.md`
- `docs/architecture/api-design.md`
- `docs/architecture/ai-architecture.md`
- `docs/architecture/deployment-architecture.md`
