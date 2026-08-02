# DapurPintar AI Observability Architecture

## Document Control

| Item | Value |
|---|---|
| Status | Proposed for M2 review |
| Scope | MVP application, dependency, security, and product observability |
| Primary audience | Product, backend, frontend, AI, QA, DevOps, and AI-assisted development |
| Related diagram | `docs/architecture/diagrams/observability-architecture-flow.mmd` |

## Executive Summary

Observability for DapurPintar AI connects user-facing behavior, business outcomes, application work, and dependency health without exposing private user data. OpenTelemetry is the standard instrumentation model for traces, metrics, and structured logs. Grafana is the visualization and operational investigation surface.

The MVP observes the request path from the browser and Fiber API through authentication, application use cases, domain operations, PostgreSQL, Redis, and the AI Gateway or provider adapter. Business events and product metrics are correlated with technical signals so the team can distinguish infrastructure failure from poor recommendation quality, missing pantry context, or a product workflow problem.

Telemetry is not a second business database. It may contain safe references to business operations for diagnosis and measurement, but it must not become authoritative for Account, Pantry, Recipe, Meal Plan, Shopping List, Recommendation, or User Profile state. Redaction, bounded cardinality, sampling, access control, and retention are mandatory because telemetry can accidentally contain credentials, prompts, private context, or personal identifiers.

## Goals

- Diagnose slow, failed, and degraded API operations.
- Trace important work across the modular monolith and external dependencies.
- Measure API, database, cache, authentication, and AI reliability against approved targets.
- Connect technical health to product outcomes such as recommendations accepted and weekly AI-assisted meals planned.
- Detect security, abuse, privacy, cost, and provider incidents safely.
- Provide actionable dashboards and alerts without premature deployment topology.
- Preserve correlation across synchronous requests and future bounded asynchronous work.
- Make telemetry useful to humans and AI-assisted engineering without leaking sensitive data.

## Non-Goals

This document does not define:

- A specific hosted telemetry vendor or production deployment topology.
- Database schema, migration, or query implementation.
- Business event transport or a message broker.
- Full product analytics warehouse architecture.
- Logging every request body, prompt, conversation, or domain object.
- A replacement for application audit records or PostgreSQL business state.

M2-015 Deployment Architecture will define environment topology, runtime placement, scaling, and operational deployment details while preserving these observability contracts.

## Observability Principles

- **One correlation model:** traces, logs, metrics, and business signals share safe request and operation context.
- **Instrument boundaries, not noise:** capture API, application, domain outcome, repository, cache, AI, and security boundaries that explain behavior.
- **Privacy by default:** do not collect sensitive payloads simply because instrumentation can access them.
- **Business ownership remains authoritative:** telemetry references business facts but does not own them.
- **Actionable signals:** every alert should identify an impact, likely boundary, and next investigation path.
- **Bounded cardinality:** labels and attributes use controlled values rather than arbitrary user input or identifiers.
- **Graceful degradation:** telemetry failure must not make core business operations fail where safe buffering or dropping is possible.
- **Cost discipline:** sampling, aggregation, retention, and attribute limits are part of the design.
- **Consistent semantics:** modules use shared names for operation, outcome, dependency, and error categories.

## Telemetry Model

### Traces

Traces represent a request or bounded background operation as a causal tree of spans. A protected API trace should connect, where applicable:

```text
Browser request -> Fiber route -> Authentication and Authorization
-> Application use case -> Domain decision -> Repository/PostgreSQL
-> Redis -> AI Gateway -> Provider Adapter -> External AI Provider
```

Each span represents a meaningful boundary or dependency operation. Internal implementation details, raw SQL, prompts, credentials, and full payloads are not span attributes.

Traces should capture:

- Trace ID and parent relationship.
- Safe request and operation identifiers.
- Module and bounded-context name.
- Operation name and outcome.
- Duration and status.
- Dependency category and safe target name.
- Error category and retry outcome.
- AI capability, model/provider metadata, and validation result where applicable.

### Metrics

Metrics provide aggregated, low-cardinality measurements for trends, SLOs, alerts, and product operations. Metric labels must use controlled dimensions such as environment, route template, method, status class, module, dependency type, provider, capability, and outcome category.

Metrics must not use raw email addresses, user IDs, pantry item names, recipe names, prompts, conversation text, or arbitrary request values as labels.

### Structured logs

Logs provide event-level explanation for investigation. Every operational log should include, where available:

- Timestamp and severity.
- Environment and service identity.
- Trace ID and span ID.
- Request ID.
- Module and operation.
- Safe user or workspace reference when operationally justified.
- Outcome and error category.
- Duration or relevant bounded quantity.
- Dependency or provider category.

Logs must not include passwords, access tokens, refresh secrets, cookies, raw prompts, full conversation content, provider payloads, SQL with user values, or unnecessary pantry and profile details.

### Business and product signals

Business signals measure meaningful product outcomes and are emitted from application or domain boundaries after the relevant business outcome is known. Candidate MVP signals include:

- Account registration accepted.
- Login success or failure category.
- Profile completed.
- Pantry item added, adjusted, or removed.
- Recipe viewed or favorited.
- AI recommendation requested, created, presented, accepted, rejected, or unable to complete.
- Recommendation option accepted.
- Recommendation conversation started, continued, or closed.
- Meal Plan created, revised, completed, or cancelled.
- Shopping List generated, activated, completed, or cancelled.
- Shopping Item completed.

Product signals must use safe dimensions and must not duplicate private domain payloads into telemetry.

## Correlation and Context Propagation

### Request context

The API establishes or validates a request correlation context at the boundary. The context carries:

- Trace ID.
- Span ID and parent relationship.
- Request ID exposed by the API error and response contract.
- Authenticated principal reference in a redacted, access-controlled form.
- Environment and service version.

The API accepts standard trace propagation from approved clients only after validation. A client cannot choose arbitrary authorization, tenant, or business ownership context through trace metadata.

### Internal propagation

Application calls, repository operations, Redis operations, AI Gateway calls, and domain event publication carry the active context. Future asynchronous work must propagate a new child context with a causation reference rather than reusing a request context indefinitely.

### Business correlation

Where available, telemetry may include safe references to an Account, Recommendation, Meal Plan, or Shopping List operation. These references are for correlation only and must not make telemetry a permission bypass or a source of business truth.

## Instrumentation Boundaries

| Boundary | Required signals | Important safe attributes |
|---|---|---|
| Browser and API request | Trace, request metric, access log | Route template, method, status class, request ID |
| Authentication and authorization | Trace, metrics, security log | Operation, outcome category, account state, denial category |
| Application use case | Trace, duration metric, business signal | Module, use case, outcome, business operation reference |
| Domain decision | Trace or event-linked metric | Aggregate type, decision outcome, invariant category |
| PostgreSQL repository | Trace, latency/error metrics | Operation category, table family or repository, result category |
| Redis operation | Trace, hit/miss/error metrics | Operation category, cache or rate-limit purpose |
| AI Gateway | Trace, latency/error/cost metrics | Capability, provider, model class, validation and fallback outcome |
| External AI provider | Child dependency trace, error/latency metrics | Provider, model, status category, retry category |
| Domain event publication | Trace, publication metric, structured log | Event type, owning context, outcome, causation reference |
| Background operation | Trace, queue/work metric, structured log | Operation type, attempt, deadline, outcome |

The modular monolith remains one logical application. Instrumentation boundaries must not imply separate services or deployable units.

## Core Metric Catalog

### API and application metrics

- Request count by route template, method, module, and status class.
- Request latency distribution and slow-request count.
- Application use-case success, validation, conflict, and failure counts.
- In-flight requests and bounded work concurrency.
- Request timeout and cancellation count.

Targets from Architecture Vision:

- Normal API operations target under 300 ms.
- API availability target is 99.9% for the MVP service.

### Database metrics

- Query count and latency distribution by repository operation category.
- Database error and timeout count.
- Connection pool utilization and exhaustion.
- Transaction commit and rollback outcomes.
- Slow query count without logging raw query parameters.

The initial query target is under 100 ms where the operation and data shape permit it. Detailed database infrastructure metrics are refined during M2-015.

### Redis metrics

- Cache hit, miss, stale, and error counts by safe purpose.
- Rate-limit allow, deny, and fallback counts.
- Operation latency and timeout count.
- Connection and availability failures.

Redis telemetry must distinguish safe cache behavior from security-critical rate limiting and must never imply Redis is authoritative business storage.

### Authentication and security metrics

- Registration, login, refresh, and logout outcomes by safe category.
- Authentication failure and account restriction counts.
- Authorization denial counts by bounded reason category.
- Refresh-token reuse detection.
- Rate-limit denials and suspicious activity categories.

Security metrics must avoid account enumeration and must not expose identifiers in metric labels.

### AI metrics

- AI request count by capability, provider, and outcome.
- AI latency distribution and timeout count.
- Provider error, retry, quota, and fallback count.
- Output schema-validation failure count.
- Safety rejection and unsupported-claim count.
- Usage and cost signals where the provider supplies them.
- Recommendation presentation, acceptance, rejection, and unable-to-complete outcomes.
- AI error rate target below 1% where the provider and operation classification permit comparable measurement.
- AI response target under 3 seconds average for bounded user-facing operations.

AI quality metrics must distinguish provider availability from product relevance. Acceptance rate is not a direct provider health metric.

### Product metrics

- Active users and authenticated sessions using safe aggregation.
- Pantry items added.
- Recipes saved.
- Shopping Lists generated.
- Meal Plans created.
- Weekly AI-assisted meals planned as the north-star outcome.

Product metrics must have an explicit owner and event meaning. They must not be inferred solely from HTTP request counts when a domain outcome is available.

## Dashboards

Grafana dashboards should be organized around investigation questions rather than technology alone.

### API health dashboard

- Availability and error rate.
- Request volume and latency percentiles.
- Slowest route templates.
- Authentication and authorization failure trends.
- Current dependency impact.

### Dependency dashboard

- PostgreSQL latency, errors, connection pressure, and transaction outcomes.
- Redis availability, latency, hit rate, and rate-limit behavior.
- AI provider latency, quota, errors, retries, and cost signals.

### AI operations and quality dashboard

- Requests by AI capability.
- Validation, safety, fallback, and provider outcomes.
- Latency, error rate, usage, and cost.
- Recommendation presentation and acceptance trends.
- Missing-context and unable-to-complete trends.

### Security dashboard

- Login, refresh, revocation, authorization denial, and rate-limit trends.
- Suspicious refresh reuse and account restriction events.
- Access anomalies by bounded operational category.

### Product operations dashboard

- Pantry activity, recipes saved, Shopping Lists generated, Meal Plans created, and weekly AI-assisted meals planned.
- Funnel from recommendation request to presentation, option acceptance, and explicit planning.
- Error and abandonment points in the core user journey.

## Alerting Strategy

Alerts should represent user impact, dependency risk, security risk, or uncontrolled cost. Every alert must define severity, owner, signal, threshold or condition, runbook reference, and expected response.

### Availability and latency alerts

- Sustained API availability below the MVP target.
- Route or use-case latency materially above the 300 ms normal target.
- Database latency or timeout increase that threatens API behavior.
- Connection pool exhaustion or repeated transaction failure.

### Dependency alerts

- Redis failure affecting rate limits or safe cache behavior.
- AI provider error, quota, or timeout surge.
- AI concurrency or cost approaching an approved limit.
- Telemetry pipeline failure or excessive dropped signals.

### Security alerts

- Unusual login failure or refresh-reuse activity.
- Abnormal authorization denials.
- Rate-limit or abuse-control surge.
- Account restriction or privileged-operation anomaly.

Alerts must avoid including secrets or private payloads in notification content.

## Privacy, Redaction, and Access

### Data classification

Telemetry attributes and events are classified before collection:

- **Operational:** environment, module, route template, status class, latency, dependency category.
- **Business-safe:** bounded operation type, aggregate type, outcome, capability, and anonymized or access-controlled reference.
- **Sensitive:** account identifiers, profile values, pantry details, recipe personalizations, conversation text, prompts, tokens, credentials, and provider payloads.

Sensitive data is excluded by default. If a diagnostic workflow requires a sensitive value, it must use a time-bounded, access-controlled, redacted mechanism outside normal telemetry.

### Redaction controls

- Redact authorization headers and cookies at ingestion.
- Never log request or response bodies by default.
- Redact prompt, conversation, and provider payload fields.
- Normalize or hash identifiers only when correlation genuinely requires it.
- Keep metric labels low-cardinality and allowlisted.
- Apply sampling before high-volume detailed traces are retained.
- Test redaction with representative failure paths.

### Access and retention

Operational access follows least privilege and is separate from business-resource access. Retention periods, deletion behavior, export behavior, and regional hosting requirements must be finalized before production. Telemetry retention must not override a user's approved privacy deletion policy.

## Sampling and Cost Control

The telemetry system uses different strategies by signal:

- Metrics are aggregated and retained for trend and alert needs.
- Error and slow traces are prioritized for retention.
- Successful traces may be sampled according to traffic and investigation needs.
- Security events and critical audit references are retained according to policy.
- Debug-level logs are disabled or tightly bounded in production.

Sampling must preserve enough context to diagnose AI, database, authentication, and cross-module failures. Sampling must not be used to hide security events or distort product metrics.

## Failure Handling

Telemetry failure must not corrupt business state. The system should:

- Prefer non-blocking export and bounded local buffering.
- Drop low-priority signals before critical business operations fail.
- Preserve security-critical audit behavior through its approved durable path.
- Mark telemetry export failure without recursively generating unbounded telemetry.
- Keep request and business operation behavior independent from Grafana availability.

Observability of the telemetry pipeline itself includes export success, dropped signals, queue pressure, exporter latency, and backend availability.

## Testing and Verification

Observability must be verified through:

- Trace propagation tests across API, application, repository, Redis, and AI boundaries.
- Redaction tests for headers, cookies, bodies, prompts, and provider errors.
- Metric naming, label, and cardinality tests.
- Error and slow-trace retention tests.
- Dashboard query and alert condition tests.
- Failure tests for database, Redis, AI provider, and telemetry backend outages.
- Correlation tests linking technical operations to safe business events.
- Load tests that measure telemetry overhead and export backpressure.

Instrumentation must remain useful when optional dependencies fail and must not materially violate the API latency target.

## Future Evolution

Future additions may include asynchronous workers, notification workflows, household and commercial metrics, data warehouse analytics, multiple application instances, replicas, and service extraction. Each addition must preserve:

- OpenTelemetry context propagation.
- Safe semantic conventions and bounded cardinality.
- Ownership-aware business event measurement.
- Redaction and least-privilege access.
- Clear separation between operational telemetry and authoritative business data.

M2-015 will refine deployment, storage, backup, and scaling topology without changing the signal contracts defined here.

## Risks and Assumptions

### Risks

- High-cardinality attributes may increase cost and make dashboards unusable.
- Instrumentation may leak prompts, tokens, private kitchen context, or account identifiers.
- Technical metrics may be mistaken for product success without domain outcome signals.
- Sampling may hide rare security or provider failures if applied without priority rules.
- Telemetry pipeline failure may go undetected without self-observability.
- AI quality may degrade while availability and latency remain healthy.
- Over-instrumentation may affect API latency and database behavior.

### Assumptions

- OpenTelemetry is the standard instrumentation model.
- Grafana is the primary visualization and dashboard surface.
- PostgreSQL, Redis, and OpenAI remain the key MVP dependencies.
- The Backend Modular Monolith is one logical application in the MVP.
- Security and privacy policies prohibit secrets and unnecessary private content in telemetry.
- Product metrics are emitted from meaningful application or domain outcomes.
- `/api/v1` remains the API contract prefix.

## Exit Criteria

M2-014 is ready for review when:

- Traces, metrics, logs, and business signals are defined.
- Correlation across the request path and dependencies is explicit.
- API, database, Redis, authentication, AI, security, and product metrics are covered.
- Dashboards and alerts are organized around operational questions and user impact.
- Privacy, redaction, sampling, cardinality, access, and retention controls are defined.
- Telemetry failure and dependency failure behavior are defined.
- Testing and self-observability requirements are documented.
- The observability flow diagram reflects the approved architecture.

## Related Documents

- `docs/architecture/architecture-vision.md`
- `docs/architecture/component-diagram.md`
- `docs/architecture/ai-architecture.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/adr/ADR-009-use-opentelemetry-for-observability.md`
- `docs/architecture/diagrams/observability-architecture-flow.mmd`
