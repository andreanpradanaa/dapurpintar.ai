# DapurPintar AI Deployment Architecture

## Document Control

| Item | Value |
|---|---|
| Status | Proposed for M2 review |
| Scope | MVP runtime, environments, delivery, operations, recovery, and scaling |
| Primary audience | Product, backend, frontend, AI, QA, DevOps, and AI-assisted development |
| Related diagram | `docs/architecture/diagrams/deployment-architecture-flow.mmd` |

## Executive Summary

DapurPintar AI is deployed as a modular monolith for the MVP. The Next.js Web Application communicates with one versioned Go/Fiber Backend Application through the `/api/v1` REST contract. The Backend Application contains the business modules, authentication, authorization, AI Gateway, observability instrumentation, and application orchestration in one logical deployable unit.

PostgreSQL is the authoritative system of record and must be deployed with durable storage, controlled access, backups, migration discipline, and restore verification. Redis is a supporting dependency for cache, rate limiting, and bounded coordination only. OpenAI is accessed through the AI Gateway and never directly by the browser. OpenTelemetry signals are exported to the operational telemetry stack and visualized through Grafana.

The deployment design separates local development, shared validation, and production concerns without prematurely introducing microservices, a message broker, a service mesh, or a complex multi-region platform. It preserves stateless application scaling, explicit dependency boundaries, secret management, graceful degradation, and a measured path toward future growth.

## Deployment Goals

- Deploy the MVP with the smallest reliable operational surface.
- Keep business modules inside one Backend Application while preserving internal boundaries.
- Protect user data, credentials, signing keys, AI provider keys, and telemetry.
- Make PostgreSQL durable, recoverable, and authoritative.
- Ensure Redis failure cannot redefine durable business state.
- Provide TLS, health checks, controlled configuration, and observable operations.
- Support repeatable build, migration, release, rollback, and recovery processes.
- Scale stateless Web and Backend Application instances horizontally when measurements require it.
- Isolate AI provider latency, quota, cost, and failure from core CRUD behavior.

## Non-Goals

The MVP deployment does not introduce:

- Microservices or separately deployed bounded contexts.
- A service mesh or cluster orchestration requirement.
- Multi-region active-active deployment.
- A message broker as a mandatory runtime dependency.
- Kubernetes-specific architecture.
- Self-hosted AI model serving.
- Data warehouse, enterprise tenant isolation, or partner network topology.

These may become valid later only after measured scale, availability, organizational, or compliance needs justify their operational cost.

## Deployment Principles

- **One logical backend:** the Backend Modular Monolith remains one deployable application in the MVP.
- **Stateless request handling:** application instances do not hold authoritative user or business state locally.
- **Durable source of truth:** PostgreSQL owns durable business facts and recovery obligations.
- **Supporting cache only:** Redis loss must not remove or redefine Account, Pantry, Recipe, Meal Plan, Shopping List, Recommendation, or Profile state.
- **Private dependencies:** PostgreSQL, Redis, and telemetry stores are not directly reachable from browsers.
- **Provider isolation:** OpenAI is reachable only through the Backend AI Gateway and controlled egress.
- **Immutable release artifacts:** the same reviewed artifact is promoted between environments where practical.
- **Configuration separation:** environment configuration and secrets are injected at runtime, not committed to source control.
- **Observable operations:** every environment has enough telemetry to diagnose user-facing and dependency failures.
- **Recovery before scale:** backups, restore tests, migrations, and rollback behavior are required before public launch.

## Logical Runtime Components

| Component | MVP deployment responsibility | Authority or boundary |
|---|---|---|
| Next.js Web Application | Serves the browser experience and calls the REST API | Never accesses private stores or OpenAI directly |
| Backend Application | Runs Fiber transport, authentication, application use cases, domain modules, repositories, AI Gateway, and telemetry instrumentation | Single logical application and business boundary |
| PostgreSQL | Stores Account, User Profile, kitchen data, recommendation context, and required session or audit facts | Authoritative durable business state |
| Redis | Supports rate limiting, safe cache, and bounded transient coordination | Non-authoritative supporting service |
| AI Provider Egress | Allows controlled Backend-to-OpenAI communication | External dependency behind AI Gateway |
| Telemetry Collection | Receives OpenTelemetry traces, metrics, and logs | Operational signal boundary, not business truth |
| Grafana | Visualizes dashboards and alerts | Operational access boundary |
| Secret Management | Supplies database, signing, provider, and telemetry credentials | Sensitive configuration authority |

The exact cloud or hosting vendor is intentionally deferred. The logical boundaries and security requirements apply regardless of provider.

## Environment Model

### Local development

Local development provides a reproducible environment for the Web Application, Backend Application, PostgreSQL, Redis, and optional telemetry services. Local dependencies may run through Docker-compatible tooling, but local convenience must not weaken production security assumptions.

Local data is synthetic or explicitly disposable. Production secrets, real user data, production provider keys, and production telemetry are never used locally.

### Shared validation environment

The shared validation environment verifies integrated Web, Backend, PostgreSQL, Redis, AI Gateway, migration, authentication, and observability behavior. It uses isolated credentials, an isolated database, bounded AI usage, and synthetic or approved test data.

This environment is the target for API contract tests, integration tests, AI evaluation smoke tests, migration rehearsal, security checks, and release-candidate verification.

### Production

Production serves real users and must provide:

- Encrypted ingress and controlled egress.
- Private database and cache connectivity.
- Managed or operationally equivalent PostgreSQL durability and backup.
- Secret management and rotation.
- Minimum viable redundancy for the agreed availability target.
- Health checks, graceful shutdown, and observable release behavior.
- Restricted operational access and auditability.
- Tested restore and rollback procedures.

Production data must never be copied into lower environments without an approved privacy-preserving process.

## Network and Trust Boundaries

The deployment has these conceptual zones:

1. **Public edge:** browser-facing Web Application and HTTPS ingress.
2. **Application zone:** Backend Application receiving only approved API traffic.
3. **Data zone:** PostgreSQL and Redis reachable only by authorized application identities.
4. **Provider egress:** controlled outbound access from the AI Gateway to OpenAI.
5. **Operations zone:** telemetry collection, Grafana, deployment tooling, and restricted administrative access.

The browser must not connect directly to PostgreSQL, Redis, OpenAI, telemetry backends, or secret management. Operational interfaces must not be exposed as public application endpoints.

## Application Runtime

### Web Application

The Next.js Web Application is deployed as a separately managed frontend runtime or static/server-rendered hosting boundary according to the selected platform. It communicates with the Backend Application using the approved `/api/v1` contract and secure session transport.

The Web Application does not contain authorization decisions for protected resources. It may present user state and route users, but the Backend Application remains authoritative for authentication, authorization, and business outcomes.

### Backend Application

The Backend Application is packaged as one versioned release artifact. It contains:

- Fiber HTTP and API boundary.
- Authentication and authorization.
- Application use cases and transaction orchestration.
- Domain modules and aggregate policies.
- PostgreSQL repositories using SQLC-generated access.
- Redis adapters for approved supporting concerns.
- AI Gateway and provider adapter.
- OpenTelemetry instrumentation and safe operational signals.

The Backend Application should support graceful startup and shutdown:

- Validate required configuration and secrets before accepting traffic.
- Verify required database connectivity and migration compatibility.
- Register telemetry without blocking indefinitely on telemetry backend availability.
- Stop accepting new requests before termination.
- Allow bounded in-flight operations to finish or cancel by deadline.
- Close database, Redis, provider, and telemetry resources cleanly.

### Stateless scaling

Backend instances remain interchangeable. Durable sessions, business state, refresh revocation, and idempotency authority are not held only in process memory. If local memory is used for performance, its loss must be safe and its scope must be explicit.

Scale-out is triggered by measured request volume, latency, concurrency, database capacity, or AI workload pressure. Scaling the Backend Application does not imply splitting bounded contexts into services.

## Data and Dependency Deployment

### PostgreSQL

PostgreSQL is deployed with:

- Durable persistent storage.
- Encryption in transit and at rest where supported.
- Restricted network access and least-privilege database roles.
- Automated or scheduled backups appropriate to the recovery target.
- Point-in-time recovery capability where available and justified.
- Connection limits and pooling aligned with Backend Application replicas.
- Migration version tracking and forward-only Goose migration discipline.
- Restore verification in the shared validation or recovery environment.

Application startup must not silently apply unsafe schema changes. Migration execution is a deliberate release step with compatibility checks and observable outcome.

### Redis

Redis is deployed as a supporting service with authentication, restricted network access, bounded memory, and explicit expiry policies. Its use is limited to:

- Rate limiting and abuse controls.
- Safe, appropriately scoped cache.
- Short-lived session or coordination state where the authority remains defined elsewhere.
- Bounded transient work coordination.

Cache loss, eviction, restart, or temporary unavailability must not lose durable business facts. Security-critical behavior must fail closed or fall back to a safe bounded policy rather than granting access.

### OpenAI

The Backend AI Gateway connects to OpenAI through controlled outbound HTTPS. The deployment supplies provider credentials through secret management. The browser, Web Application runtime, domain layer, and ordinary API handlers never receive provider credentials or call the provider directly.

Provider timeouts, quotas, retries, cost controls, and egress failures are observable. Core non-AI operations remain usable when the provider is unavailable.

### Telemetry

The Backend Application emits OpenTelemetry traces, metrics, and structured logs to the approved telemetry collection boundary. Grafana consumes operational signals for dashboards and alerts. Telemetry backends are protected by operational access controls and do not become a business data store.

Telemetry collection failure must not block ordinary business requests indefinitely. High-priority security and audit signals follow their approved durable path.

## Configuration and Secrets

Configuration is separated into:

- Non-sensitive environment configuration such as service address, environment name, feature flags, timeouts, and safe limits.
- Sensitive secrets such as database credentials, Redis credentials, JWT signing keys, cookie secrets, OpenAI keys, and telemetry exporter credentials.

Secrets are supplied through an approved secret-management mechanism at runtime. They are not committed to Git, embedded in images, printed in logs, or returned by health endpoints.

Secret lifecycle requirements:

- Separate secrets by environment.
- Apply least privilege and narrow scope.
- Rotate signing, database, cache, provider, and telemetry secrets according to policy.
- Support planned rotation without exposing old or new values.
- Revoke compromised secrets and investigate related telemetry.
- Validate required secrets before the Backend Application accepts traffic.

## Delivery and Release Process

The deployment pipeline should progress through explicit stages:

1. Validate source, formatting, tests, security checks, API contract, and AI evaluation gates.
2. Build immutable Web and Backend release artifacts.
3. Scan dependencies and artifacts for known security issues.
4. Apply or rehearse compatible database migrations in the target environment.
5. Deploy the release candidate to shared validation.
6. Run smoke, integration, authentication, authorization, AI, observability, and recovery checks.
7. Promote the reviewed artifact to production with an approval boundary.
8. Monitor health, error rate, latency, dependency behavior, and business signals.
9. Roll back the application artifact or apply the approved forward recovery procedure when necessary.

Database migrations must preserve compatibility between the currently running application and the next release where rolling replacement is used. Destructive schema changes require a staged migration and explicit review.

## Health, Readiness, and Graceful Degradation

### Liveness

Liveness indicates that the process is running and able to execute its basic runtime loop. It must not fail solely because OpenAI, Redis, Grafana, or another optional dependency is temporarily unavailable.

### Readiness

Readiness indicates that the instance may receive traffic. It verifies the configuration and dependencies required for the instance's declared service level. It must distinguish:

- Core readiness, including compatible PostgreSQL access.
- Optional AI readiness.
- Supporting Redis readiness.
- Telemetry export health.

The deployment must not route traffic to an instance that cannot safely serve core authenticated operations. AI unavailability should produce bounded AI failure behavior rather than hide core CRUD readiness where the product policy permits.

### Degradation behavior

- Account and core CRUD operations remain available when AI is unavailable where feasible.
- Cache misses fall back to authoritative PostgreSQL reads where applicable.
- Redis failure triggers safe rate-limit or coordination behavior.
- Telemetry export failure does not corrupt business transactions.
- Provider timeout returns a safe product-level failure and does not expose provider internals.
- Database unavailability prevents operations requiring authoritative state and produces an observable dependency failure.

## Security and Compliance Boundaries

- Terminate HTTPS at an approved secure edge and protect internal traffic according to environment risk.
- Restrict ingress to the Web Application and approved operational paths.
- Restrict egress from the Backend Application to approved database, Redis, telemetry, and AI provider destinations.
- Use separate runtime identities for Web, Backend, database migration, telemetry, and operational access where practical.
- Apply least-privilege database roles and prevent application runtime access beyond its required schema operations.
- Keep backups, logs, telemetry, and images free of secrets and unnecessary personal data.
- Audit release, migration, secret, administrative, and recovery actions.
- Keep production access time-bounded and reviewed.

This architecture does not claim regulatory compliance by itself. Regional privacy, retention, consent, and deletion obligations require product, legal, and operational validation before launch.

## Backup, Recovery, and Disaster Readiness

### Recovery objectives

The product must define target Recovery Point Objective (RPO) and Recovery Time Objective (RTO) before production launch. Until approved, the deployment must not claim a stronger recovery guarantee than the tested backup and restore process supports.

### Required recovery capabilities

- Scheduled PostgreSQL backups with protected access.
- Backup encryption and environment separation.
- Restore tests using isolated infrastructure.
- Migration replay or compatibility verification after restore.
- Documented secret recovery and rotation process.
- Recovery procedure for Redis loss as a non-authoritative dependency.
- Recovery procedure for AI provider outage or credential rotation.
- Recovery procedure for telemetry backend outage.
- Verification that restored business data remains tenant- and ownership-safe.

Recovery tests must include Account, User Profile, Pantry, Meal Plan, Shopping List, Recommendation, and session-revocation facts that are required for safe operation.

## Scaling and Capacity

The initial deployment is sized for the MVP targets of approximately 1,000 registered users and 500 MAU, subject to load testing and measured operational limits. Capacity planning must observe:

- Backend request volume, concurrency, and latency.
- PostgreSQL connection, CPU, memory, storage, and query latency.
- Redis memory, hit rate, and rate-limit load.
- AI provider quota, concurrency, latency, and cost.
- Telemetry volume, cardinality, retention, and export pressure.

Scaling order should prefer measured, lower-complexity changes:

1. Optimize queries, indexes, pooling, bounded requests, and cache behavior.
2. Add Backend Application replicas behind a load balancer.
3. Tune PostgreSQL capacity and introduce replicas or read models where justified.
4. Isolate AI concurrency or asynchronous workloads.
5. Consider service extraction only after a bounded context has measured operational or organizational justification.

## Deployment Observability

Deployment operations emit signals for:

- Build and release result.
- Artifact version and source revision.
- Migration start, completion, and failure.
- Startup, readiness, shutdown, and restart.
- Instance count and replacement events.
- Backup completion, failure, and restore test result.
- Secret rotation outcome without secret values.
- Dependency health and configuration validation.

These signals connect to the M2-014 observability model and are visible to authorized operators through Grafana or the approved operational tooling.

## Rollback and Change Safety

Application rollback is preferred when the release artifact is faulty and the database schema remains backward-compatible. Database rollback is not assumed to be safe for every migration; forward-compatible corrective migrations are preferred for durable schema changes.

Every release must define:

- Compatibility with the previous release.
- Migration order and recovery behavior.
- Health and alert checks.
- Rollback or forward-fix decision criteria.
- Owner and approval boundary.
- User-impact communication path for material incidents.

Feature flags may reduce exposure of incomplete product capability, but they must not bypass authentication, authorization, tenant isolation, business invariants, or safety controls.

## Future Evolution

Future deployment extensions may include multiple Backend Application replicas, asynchronous workers, notification processing, read replicas, separate AI workloads, household or commercial tenancy, regional environments, and bounded service extraction. Each extension must preserve:

- API and application ownership boundaries.
- PostgreSQL authority and explicit data ownership.
- Secure secret and network boundaries.
- OpenTelemetry propagation and redaction.
- AI Gateway provider isolation.
- Tested backup, recovery, and operational access control.

## Risks and Assumptions

### Risks

- A simple MVP deployment may become a single point of failure if redundancy and recovery are not tested.
- Incorrect database connection pooling may exhaust PostgreSQL when application replicas increase.
- Redis may be accidentally treated as authoritative during a degraded operation.
- AI provider latency or cost may consume application capacity without isolation.
- Migration incompatibility may make rollback unsafe.
- Secrets or private data may leak through images, logs, backups, or telemetry.
- Vendor-specific deployment decisions may reduce portability before product scale justifies them.
- Recovery objectives may be promised before restore procedures are tested.

### Assumptions

- The MVP uses one logical Backend Application, not microservices.
- The Web Application and Backend Application are separate logical runtime boundaries.
- PostgreSQL is the authoritative durable store.
- Redis is supporting infrastructure only.
- OpenAI is accessed through the AI Gateway over controlled egress.
- OpenTelemetry and Grafana remain the operational observability standard.
- `/api/v1` remains the API contract prefix.
- Exact cloud provider, container orchestrator, regions, RPO, and RTO are deployment decisions to validate before production.

## Exit Criteria

M2-015 is ready for review when:

- Logical runtime components and trust boundaries are defined.
- Environment differences and production requirements are explicit.
- Modular monolith deployment and stateless scaling are preserved.
- PostgreSQL, Redis, OpenAI, secrets, and telemetry deployment boundaries are covered.
- Release, migration, health, degradation, rollback, backup, and recovery behavior are documented.
- Security, capacity, observability, and operational access requirements are defined.
- Future deployment evolution does not introduce premature infrastructure complexity.
- The deployment flow diagram reflects the approved architecture.

## Related Documents

- `docs/architecture/architecture-vision.md`
- `docs/architecture/observability-architecture.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/ai-architecture.md`
- `docs/architecture/adr/ADR-004-use-modular-monolith-for-mvp.md`
- `docs/architecture/adr/ADR-003-use-postgresql-as-system-of-record.md`
- `docs/architecture/diagrams/deployment-architecture-flow.mmd`
