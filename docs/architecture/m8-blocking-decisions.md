# DapurPintar AI M8 Blocking Decision Records

## Document Control

| Item | Value |
|---|---|
| Milestone | M4 - Solution Architecture Refinement |
| Deliverable | M4-005 M8 Blocking Decision Records |
| Status | Draft - Awaiting Owner Approval |
| Parent documents | `docs/architecture/m4-decision-register.md`, `docs/architecture/implementation-backlog.md`, `docs/architecture/ai-architecture.md` |
| Scope | Decisions blocking M8 AI Foundation: M4-DEC-010, M4-DEC-011, M4-DEC-012, M4-DEC-016 |

## Purpose

This document records the decisions that block M8 AI Foundation, following the decision record format required by the Architecture Decision Register. Each record captures context, options considered, the recommended direction, consequences, affected documents, and revisit conditions. Approval converts the register item from `Pending` to `Decided`.

A record becomes final only after the owner approves it and affected disciplines confirm consequences. Approving these records unblocks DP-AI-001, DP-AI-002, DP-AI-003, and the technical spikes DP-SPK-003, DP-SPK-004, DP-SPK-006, DP-SPK-009 in the Implementation Backlog.

## Decision Record: M4-DEC-010 - Initial OpenAI model and capability profile

### Context and problem

M8 needs a concrete OpenAI model and capability profile behind the AI Gateway: which model family, how context is sized, which capabilities (structured output, tooling) are used, and which latency/cost/quality targets must hold. Without an explicit profile, provider selection, prompt design, and evaluation thresholds cannot start, and ad hoc model choice could silently change behavior and cost between environments.

### Options considered

1. **A fixed default model profile with versioned alternatives (recommended).** Select one primary OpenAI model (capability class: structured-output and reasoning capable) as the MVP default, configured by name and version through the Gateway. Profile carries documented latency, cost, and quality expectations plus fallback candidates. Model upgrades are a deliberate, evaluated change, not an implicit drift.
2. **Model auto-selection by the provider.** Simplest to operate, but produces nondeterministic latency, cost, and output behavior across requests, which breaks evaluation and quota control.
3. **Multi-provider routing at MVP.** Matches long-term AI Architecture goals but adds provider abstraction and routing work beyond MVP scope and contradicts the non-goal of a multi-provider routing platform at MVP.

### Recommended direction

Adopt **a fixed default model profile with versioned alternatives**. The AI Gateway exposes a model profile (primary model name/version, capability set, context budget, temperature/seed policy, fallback candidates). The initial profile is defined during M8 validation with DP-SPK-003 and targets: MVP-quality structured kitchen recommendations, bounded per-request latency, and cost within the M4-DEC-016 budget. Model and profile changes flow through the versioned prompt/policy pipeline (M4-DEC-011) and are regression-evaluated (M4-DEC-012) before promotion.

### Consequences and risks

- Deterministic, reviewable model behavior; evaluation and quota control are meaningful.
- A fixed default needs periodic re-evaluation; a model upgrade is a governed change, not automatic.
- Fallback candidates exist but are not load-balanced at MVP; switching is an explicit operation.
- Risk of profile drift if environment config diverges; the Gateway owns the single profile source.

### Affected documents and contracts

- `docs/architecture/ai-architecture.md` (Gateway and prompt/policy management)
- `docs/architecture/m4-decision-register.md` (M4-DEC-011, M4-DEC-012, M4-DEC-016)
- Backlog items DP-AI-001, DP-AI-002, DP-AI-003, DP-FEAT-006, DP-FEAT-008

### Owner and approval

- Owner: AI Engineering + Product
- Approval date: Pending

### Revisit condition

Revisit if a model upgrade materially changes latency, cost, or output quality, if a capability need (tooling, larger context) appears, or if provider pricing changes the budget arithmetic.

## Decision Record: M4-DEC-011 - Prompt, safety, and structured-output policy

### Context and problem

M8 implements prompt and policy versioning plus structured-output validation. Without an agreed policy, prompt changes could be deployed unreviewed, safety gates (injection, unsafe content, leakage) could behave inconsistently, and malformed provider output could reach product responses.

### Options considered

1. **Versioned prompt/safety/policy pipeline with layered validation (recommended).** Prompts, safety policy, and structured-output schemas are versioned artifacts selected by the Gateway. Validation is layered: transport/schema, product rules, and safety (unsupported claims, unsafe instructions, privacy leakage, prompt-injection output, prohibited content). Invalid or unsafe output is rejected or safely degraded, never surfaced as a commitment. Prompt changes require review and regression evaluation before promotion.
2. **Prompts as feature-handler strings.** Fastest to write, but unreviewable, unevaluable, and conflicts with AI Architecture's "controlled product and safety artifacts".
3. **Provider-side validation only.** Trusts the provider's moderation; does not protect against leakage of private context or prompt-injection output, and does not enforce product-shaped output.

### Recommended direction

Adopt **versioned prompt/safety/policy pipeline with layered validation**. Prompts, safety policy, and structured-output schemas are versioned and selected through the Gateway. Output is validated in layers (schema, product rules, safety) before a response or business aggregate is touched. Safety policy covers prompt injection, unsafe or prohibited content, unsupported claims, and privacy leakage. Any failure degrades safely to a bounded, honest error rather than an invented answer. Changes to prompts, policies, or schemas are reviewable and regression-evaluated before promotion.

### Consequences and risks

- Prompts and policies become governed artifacts; feature handlers cannot edit behavior inline.
- Layered validation adds Gateway complexity, justified by safety and reproducibility.
- A failure that rejects output must degrade gracefully, which requires explicit fallback messaging.
- Risk of over-engineering if validation layers are duplicated; the Gateway owns a single validation path.

### Affected documents and contracts

- `docs/architecture/ai-architecture.md` (Prompt and Policy Management, Safety and Trust, Structured Output)
- `docs/architecture/m4-decision-register.md` (M4-DEC-010, M4-DEC-012, M4-DEC-013, M4-DEC-016)
- Backlog items DP-AI-002, DP-FEAT-006, DP-FEAT-007, DP-FEAT-008

### Owner and approval

- Owner: AI Engineering + Security
- Approval date: Pending

### Revisit condition

Revisit if a new safety class (e.g., nutrition or medical advice) enters scope, if provider safety capabilities change the injection surface, or if a policy bypass is discovered in testing.

## Decision Record: M4-DEC-012 - AI evaluation dataset and acceptance rubric

### Context and problem

M8 needs an evaluation harness and representative scenarios (DP-AI-003), but the acceptance criteria for "good enough" AI are undefined. Without an agreed dataset and rubric, quality cannot be measured consistently, and prompt/model/policy changes cannot be regression-checked before release.

### Options considered

1. **A privacy-safe representative dataset with a published rubric and pass gates (recommended).** Build an evaluation dataset representative of Indonesian MVP usage (Sarah and Daniel scenarios: pantry, expiry, meal-plan acceptance, decline paths, unsafe/injection cases), stored apart from production data. Publish a rubric scoring relevance, accuracy, safety, and format compliance. A pass gate (threshold per dimension) blocks prompt/model/policy promotion until met.
2. **Manual spot-checking.** No deterministic quality signal; promotion decisions are subjective and cannot gate CI.
3. **Reuse production user data for evaluation.** Rich signal but violates privacy expectations and the "evaluation datasets separated from private production data" principle.

### Recommended direction

Adopt **a privacy-safe representative dataset with a published rubric and pass gates**. The dataset covers the MVP persona journeys, safety and injection cases, and failure modes. The rubric scores relevance, factual accuracy, safety compliance, and structured-output conformance. A pass gate defines minimum scores per dimension; promotions that miss the gate are rejected. Evaluation records store only metadata (scenario, rubric scores, decision) separate from business data, per the M4-DEC-013 retention posture. Acceptance, rejection, and rationale from real users remain durable business records used alongside the rubric.

### Consequences and risks

- Quality becomes measurable and CI-gateable; promotions are evidence-backed.
- Building a representative dataset is real work and must be privacy-reviewed.
- Risk of rubric gaming or over-fitting to the dataset; the dataset must be periodically refreshed and separated between training/eval concerns.
- Metrics must not reward invented facts; accuracy dimension treats unsupported claims as failures.

### Affected documents and contracts

- `docs/architecture/ai-architecture.md` (Evaluation and Quality, Data and Retention)
- `docs/architecture/m4-decision-register.md` (M4-DEC-013, M4-DEC-016)
- Backlog items DP-AI-003, DP-QA-002, DP-FEAT-006

### Owner and approval

- Owner: AI Engineering + Product
- Approval date: Pending

### Revisit condition

Revisit if the persona set or MVP journeys change, if a safety class enters scope, if evaluation shows the rubric no longer discriminates good from bad output, or if user feedback diverges from rubric scores.

## Decision Record: M4-DEC-016 - AI quota and cost budget

### Context and problem

AI calls have real per-token cost and failure modes (timeouts, retries, runaway loops). Without an explicit budget and enforcement, cost exhaustion or a runaway retry loop could generate unbounded spend or duplicate business commands.

### Options considered

1. **Explicit per-user and global budgets with enforced limits and alerts (recommended).** Define a global monthly AI cost budget plus per-user quota, enforced at the Gateway and API boundary. Timeouts and retries are bounded; retry loops can never create cost spikes or duplicate commands. Quota exhaustion degrades the AI feature to an honest "unavailable" state without breaking core non-AI operations. Alerts fire near thresholds.
2. **No budget, rely on provider limits.** Simple but uncontrolled; cost and plan limits remain unmanaged and a single hot path could exhaust spend.
3. **Hard block at zero with manual intervention.** Safest for cost but poor UX; better to degrade gracefully within budget while alerting.

### Recommended direction

Adopt **explicit per-user and global budgets with enforced limits and alerts**. The Gateway enforces per-request timeout and retry bounds, and the API layer enforces per-user quota. A global monthly cost budget is tracked and alerted (DP-SPK-009 defines concrete mechanics and alert thresholds). On quota exhaustion, AI recommendations return a bounded, honest unavailable state while pantry, meal, shopping, and account operations continue normally. Telemetry records usage and cost metadata per request, correlated to the originating API request.

### Consequences and risks

- Cost and plan limits are controlled; alerts prevent surprise spend.
- AI feature degrades gracefully on exhaustion, requiring explicit messaging and re-enablement path.
- Budget mechanics (metering accuracy, alert thresholds) need DP-SPK-009 validation.
- Risk of quota churn (users exhausting and being confused) requires clear UX for the degraded state.

### Affected documents and contracts

- `docs/architecture/ai-architecture.md` (Reliability, Availability, and Failure Handling; Observability)
- `docs/architecture/m4-decision-register.md` (M4-DEC-010, M4-DEC-012)
- Backlog items DP-AI-003, DP-SPK-009, DP-QA-002, DP-FEAT-006

### Owner and approval

- Owner: Product + Finance + AI
- Approval date: Pending

### Revisit condition

Revisit if the pricing model changes, if per-user plan limits are introduced, if a model change alters cost per recommendation, or if alerting shows the budget is either constantly hit or never approached.

## Approval Protocol

1. Owner reviews the recommended direction and evidence for each record.
2. Affected disciplines (AI Engineering, Security, Product, Finance, Backend) confirm consequences.
3. Architecture checks compatibility with bounded ownership and existing ADRs (ADR-010 AI Gateway).
4. Approved records update the register status to `Decided`.
5. M8 work (DP-AI-001, DP-AI-002, DP-AI-003) and its spikes may then start.

## Related Documents

- `docs/architecture/m4-decision-register.md`
- `docs/architecture/implementation-backlog.md`
- `docs/architecture/ai-architecture.md`
- `docs/architecture/m4-m5-blocking-decisions.md`
