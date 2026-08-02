# DapurPintar AI Architecture Vision

## Document Control

| Item | Value |
|---|---|
| Status | Initial architecture vision |
| Scope | MVP foundation and future extension direction |
| Primary audience | Product, backend, frontend, AI, QA, DevOps, and AI-assisted development |
| Source of truth | `docs/product/` product planning documentation |
| Related diagram | `docs/architecture/diagrams/architecture-overview.mmd` |

## Executive Summary

DapurPintar AI is an AI-first SaaS kitchen companion for people and families who want to plan meals, manage pantry stock, discover suitable recipes, shop efficiently, and reduce food waste. The MVP focuses on registration and login, user profile and preferences, pantry CRUD, recipe search and detail, favorites, AI chat and recommendations, daily and weekly meal plans, and shopping lists with automatic generation.

The system will start as a modular monolith implemented with Go and Fiber. Clean Architecture and Domain Driven Design will keep business rules independent from HTTP, persistence, caching, and external AI providers. A versioned REST API will be the contract between the Next.js web application and the backend. PostgreSQL will be the system of record, Redis will provide selective low-latency and control-plane capabilities, and OpenTelemetry will make application, database, and AI-provider behavior observable from the beginning.

This architecture deliberately optimizes for a validated MVP while preserving seams for family workspaces, nutrition, notifications, OCR, grocery integrations, public APIs, premium capabilities, and eventual regional or enterprise growth. It does not introduce microservices, a marketplace, food delivery, hardware, or social-platform concerns into the MVP.

## Architecture Goals

- Deliver a small, coherent MVP that supports the product north star: weekly AI-assisted meals planned.
- Make pantry, recipe, meal-planning, and shopping context available to AI without duplicating domain ownership.
- Protect personal, household, preference, and AI-conversation data through privacy-by-design controls.
- Keep domain rules testable without a running web server, database, cache, or AI provider.
- Provide a stable API contract that can support the web client now and partner or public APIs later.
- Make AI provider replacement, prompt evolution, and response evaluation possible without changing core domains.
- Establish operational visibility for API latency, database behavior, cache behavior, AI latency/cost/errors, and business events.
- Support the initial Indonesia launch and future commercialization without prematurely distributing the system.

## Guiding Principles

1. **MVP first:** Build only the approved P0 capabilities needed to validate user value and product-market fit.
2. **AI simplifies work:** AI is a decision-support capability grounded in user context, not an uncontrolled source of truth.
3. **User value before feature count:** Architecture effort must follow measurable user outcomes such as time saved, food waste reduction, and recommendation acceptance.
4. **Domain ownership is explicit:** Each module owns its rules, use cases, data access, and API surface.
5. **Dependencies point inward:** Frameworks and providers depend on application and domain abstractions, never the reverse.
6. **API first:** Contracts, validation, error semantics, authorization, and versioning are defined before client or server implementation.
7. **Privacy and security by design:** Collect the minimum data necessary, isolate tenants, and make access decisions explicit.
8. **Observable by default:** Requests, important business actions, database calls, cache operations, and AI calls carry correlation context.
9. **Build for evolution, not speculation:** Use modular boundaries and ports before considering separate deployable services.
10. **Documentation before implementation:** Architecture, contracts, decisions, and assumptions remain discoverable for humans and AI agents.

## Quality Attributes

| Attribute | Architectural intent |
|---|---|
| Usability | A responsive web experience should make common actions, especially menu selection and shopping-list creation, direct and understandable. |
| Performance | Normal API operations target under 300 ms; database queries target under 100 ms; AI responses target under 3 seconds where the provider permits. |
| Availability | MVP services target 99.9% API and AI availability, with graceful degradation when optional dependencies fail. |
| Security | Authentication, authorization, tenant isolation, input validation, rate limiting, secret protection, and auditability are mandatory concerns. |
| Privacy | User data is purpose-limited, protected in transit and at rest, and excluded from provider prompts unless necessary and authorized. |
| Maintainability | Module boundaries, dependency inversion, generated SQL access, migrations, and automated tests reduce change cost. |
| Scalability | Stateless API processes, indexed relational data, cacheable reads, and provider abstraction support growth without an early microservice split. |
| Testability | Domain and application behavior can be tested independently; adapters are covered by integration and contract tests. |
| Operability | Telemetry supports diagnosis of user-facing, infrastructure, database, and AI failures. |

## Functional Scope

### MVP capabilities

- Account registration, login, JWT-based authentication, user profile, and preferences.
- Pantry item creation, update, deletion, category, quantity, and expiration date tracking.
- Recipe search, recipe detail, favorite recipes, and AI recipe recommendation.
- AI chat assistant and pantry-aware recommendation use cases.
- Daily and weekly meal planning.
- Shopping lists and automatically generated shopping lists.

### Future capabilities and extension points

Family workspaces and shared data, nutrition summaries and goals, notifications, barcode/OCR/image recognition, advanced analytics, administration, prompt management, premium plans, grocery partnerships, public APIs, smart-kitchen integrations, and enterprise use cases are outside the MVP runtime scope. Their future introduction must add modules or adapters through the same architectural boundaries rather than bypassing domain ownership.

### Explicitly out of scope

The architecture does not include food sales, food delivery, hardware or IoT development, cooking robots, social feeds, live streaming, POS, accounting, or ERP inventory behavior.

## Non Functional Requirements

- **Latency:** Meet the targets in the quality-attribute table for normal operations; long-running AI workflows must expose progress or a clear bounded response behavior.
- **Availability:** Maintain a 99.9% API target and isolate optional AI or cache failures from core authenticated data operations where feasible.
- **Data integrity:** PostgreSQL transactions protect related pantry, meal-plan, shopping-list, favorite, and account changes.
- **Consistency:** PostgreSQL is authoritative; Redis is never the sole source of user or business data.
- **Security:** All protected operations require authenticated identity and module-level authorization checks.
- **Observability:** Every request has trace and correlation context; failures include actionable structured attributes without secrets or sensitive prompt content.
- **Recoverability:** Database backups, migration discipline, restore procedures, and dependency failure handling are required before public launch.
- **Portability:** AI and infrastructure providers are accessed through ports so provider changes do not force domain changes.
- **Quality:** Automated unit, integration, API, AI evaluation, performance, and security testing should support the engineering target of at least 80% coverage and zero critical release defects.
- **Commercial readiness:** Tenant/workspace ownership, usage measurement, plan entitlements, and provider-cost visibility must have clear future insertion points, even where billing is not part of MVP.

## Technology Stack

| Concern | Technology / decision |
|---|---|
| Web client | Next.js, React, TypeScript, Tailwind CSS |
| Backend runtime | Go |
| HTTP/API | Fiber, versioned REST/JSON API |
| Architecture | Modular Monolith, Clean Architecture, Domain Driven Design |
| System of record | PostgreSQL |
| SQL access | SQLC-generated type-safe queries |
| Schema migration | Goose |
| Cache and fast coordination | Redis |
| AI provider | OpenAI through an internal AI provider abstraction |
| Telemetry | OpenTelemetry |
| Visualization | Grafana |

Technology choices are implementation constraints for the system design phase, not permission to add unapproved infrastructure or split the MVP into services.

## High Level Architecture

The browser client communicates only with the public REST API. Fiber handles transport concerns such as routing, authentication middleware, request validation, response mapping, and API errors. The application layer coordinates use cases and transactions. The domain layer contains entities, value objects, policies, and domain services. Repository ports are implemented by PostgreSQL adapters using SQLC-generated queries. Redis and OpenAI are accessed through explicit ports and adapters. OpenTelemetry instruments the flow end to end and exports operational data for Grafana.

The backend is deployed as one logical application in the MVP. Modules communicate through application contracts and domain events where useful, not through direct access to another module's tables or internal handlers. A later extraction into services is possible only after measured scaling or organizational needs justify the operational cost.

## Architectural Style

### Modular Monolith

One deployable backend contains cohesive business modules with explicit public interfaces. Each module owns its domain behavior and persistence boundary. This minimizes MVP operational complexity while preventing a shared, unstructured codebase.

### Clean Architecture

Dependencies flow from adapters and frameworks toward application use cases and domain rules. The domain does not know Fiber, PostgreSQL, Redis, OpenAI, or Next.js. Application ports define what external capabilities are needed; infrastructure implements them.

### Domain Driven Design

The model follows kitchen-management language: user, pantry item, recipe, favorite, meal plan, shopping list, and AI recommendation. Aggregates and invariants are defined during detailed system design. Module boundaries must follow business capabilities, not technical layers alone.

### API First

The REST contract is a versioned product interface. Request and response schemas, authentication requirements, authorization rules, pagination, idempotency where needed, validation, error codes, and compatibility policy are specified before client/server implementation.

## Core Modules

| Module | MVP responsibility | Main ownership |
|---|---|---|
| Authentication | Register, login, JWT lifecycle, credential security | Identity and authentication state |
| User | Profile and preferences used by product and AI | User-owned settings and preferences |
| Pantry | Items, categories, quantities, expiry dates, and history direction | Pantry inventory and lifecycle rules |
| Recipe | Search, detail, favorites, and recipe metadata | Recipe read model and user favorites |
| Meal Planner | Daily and weekly plans | Planned meals, dates, and plan invariants |
| Shopping | Manual and generated shopping lists | Shopping-list items, status, and generation results |
| AI Assistant | Chat, recipe recommendation, pantry analysis, and AI orchestration | AI request lifecycle, context assembly, safety, and provider abstraction |
| Shared Platform | Authorization, configuration, validation, errors, telemetry, and cross-cutting policies | Reusable technical capabilities, not business ownership |

Post-MVP modules such as Family Workspace, Nutrition, Notification, OCR, Analytics, and Administration should be added as bounded modules with their own contracts and ownership.

## API Strategy

- Use a versioned REST/JSON API, beginning with `/api/v1`.
- Model endpoints around product capabilities and resources, not database tables or provider APIs.
- Require authentication for user-owned data and enforce ownership/workspace scope in the application layer.
- Use consistent success envelopes, error codes, validation errors, pagination, sorting, and request correlation identifiers.
- Keep handlers thin: translate HTTP to an application command/query and map the result back to a response.
- Avoid exposing internal entities, SQL details, prompts, provider payloads, or credentials.
- Define rate limits and quotas for authentication and AI endpoints; make future plan-based entitlements possible.
- Treat AI requests as bounded, observable operations with explicit timeout, retry, fallback, and provider-error behavior.
- Publish an API specification as the contract artifact during system design; generated clients or server scaffolding must not become the domain model.

## Database Strategy

- PostgreSQL is the authoritative store for identity, users, pantry, recipes, favorites, meal plans, shopping lists, and AI request metadata required for product behavior.
- Use relational ownership, foreign keys, constraints, indexes, and transactions to protect invariants and tenant isolation.
- Use SQLC for typed access from reviewed SQL; repositories expose domain/application abstractions rather than SQLC types.
- Use Goose for ordered, forward-only schema migrations. Application versions and migrations are deployed deliberately and are observable.
- Keep module persistence boundaries explicit. A module must not write another module's tables directly.
- Store timestamps consistently and define the MVP timezone behavior for meal plans, expiration, and daily summaries.
- Store only the AI conversation and prompt/context data required by the product and evaluation policy; classify sensitive fields and define retention before implementation.
- Redis may cache read results, hold short-lived authentication or rate-limit state, and coordinate bounded work. Cache invalidation and expiry must be explicit, and cache loss must not lose authoritative data.

## AI Strategy

AI is an application capability, not a domain authority. The AI Assistant module assembles relevant context through application ports, including authorized pantry data, preferences, meal plans, recipes, and other approved signals. It sends a minimized, structured request through an OpenAI adapter and validates the result before returning a product response.

- Use an internal AI gateway/provider port to isolate OpenAI credentials, SDK details, model selection, retries, timeouts, quotas, and provider errors.
- Prefer structured outputs with schema validation for recommendations, meal plans, and shopping-list generation.
- Ground recommendations in persisted product data and clearly distinguish generated suggestions from user-confirmed facts.
- Apply input validation, prompt-injection defenses, output safety checks, and sensitive-data minimization.
- Version prompts and record model/prompt metadata needed for reproducibility and evaluation without exposing secrets.
- Cache only safe, appropriately scoped results; never allow one user's private context to leak through shared cache keys.
- Track latency, availability, error rate, token/cost signals, acceptance, and feedback against the product targets.
- Provide graceful failure: preserve core CRUD and explain when an AI-dependent operation cannot complete.
- Keep RAG, additional providers, computer vision, voice, and nutrition coaching as future adapters or modules, not MVP assumptions.

## Security Principles

- Authenticate with secure password handling and JWT-based access control; protect refresh or revocation state as appropriate to the final session design.
- Authorize every read and write against the authenticated user and, later, workspace membership and role.
- Enforce tenant isolation in application queries and repository filters; never trust client-supplied ownership identifiers.
- Validate and normalize all external input, including chat content, filters, identifiers, and generated AI output.
- Apply rate limiting to login, registration, chat, recommendation, and other abuse-prone endpoints using Redis-backed controls where appropriate.
- Encrypt traffic in transit and protect database, cache, provider, and signing secrets through deployment secret management.
- Do not log passwords, tokens, full private prompts, raw sensitive profile data, or unnecessary pantry details.
- Record security-relevant audit events such as authentication changes and privileged actions.
- Use least privilege for runtime identities, database roles, provider keys, and operational access.
- Treat prompt injection, provider compromise, data leakage, abuse, and cost exhaustion as security concerns.
- Define retention, deletion, export, and consent behavior before commercialization and regional expansion.

## Observability Strategy

OpenTelemetry is the instrumentation standard across the request path. A trace should connect the browser request where supported, Fiber handler, application use case, repository/database operation, Redis operation, and OpenAI call.

- Emit structured logs with severity, timestamp, environment, module, request ID, trace ID, user/workspace-safe identifiers, and outcome.
- Collect metrics for request count, latency, status/error rate, active users, database latency, cache hit rate, AI latency/errors/cost, and key product events.
- Capture distributed traces for slow or failed API, database, cache, and AI interactions.
- Use Grafana dashboards for API health, infrastructure dependencies, AI quality/operations, and product-operational indicators.
- Define alerts for availability, latency, error rates, database saturation, cache failures, AI provider failures, and unusual authentication or AI usage.
- Apply sampling and attribute redaction so telemetry remains useful without becoming a source of personal-data exposure.
- Instrument north-star and leading indicators such as AI chat usage, pantry items added, recipes saved, shopping lists generated, and meal plans created.

## Scalability Strategy

1. Keep the API stateless so multiple application instances can be added behind a load balancer.
2. Use PostgreSQL indexes, bounded queries, pagination, connection pooling, and measured query optimization before introducing another datastore.
3. Use Redis for hot, safe-to-cache reads, rate limiting, and short-lived coordination, with PostgreSQL fallback where applicable.
4. Bound and queue work that does not need to block the user request, especially future notifications, analytics, and large AI workflows.
5. Separate AI provider quotas, retries, and concurrency from normal CRUD traffic so provider pressure does not exhaust the API.
6. Preserve module interfaces so high-load modules can later be extracted without changing consumer contracts.
7. Add read models, asynchronous events, replicas, or service extraction only in response to measured bottlenecks and clear ownership needs.

## Constraints

- MVP delivery and product validation take priority over speculative platform complexity.
- Backend technology is Go with Fiber, PostgreSQL, Redis, SQLC, Goose, and OpenTelemetry.
- Frontend technology is Next.js, React, TypeScript, and Tailwind CSS.
- The initial architecture is a modular monolith and must not assume microservice operations.
- OpenAI is an external dependency with cost, latency, quota, availability, and policy risks; provider abstraction is required.
- The initial market is Indonesia, while future language, regional, SaaS plan, and compliance needs must not be blocked by MVP design.
- Product assumptions about AI adoption, digital pantry entry, recommendation value, and freemium conversion remain hypotheses to validate through beta analytics and testing.
- Existing product documentation is the source of truth; any scope change requires explicit product and architecture review.
- This document is an architecture vision, not a detailed schema, endpoint specification, deployment topology, or implementation plan.

## Success Criteria

The architecture vision is successful when:

- Engineers can identify the correct module and dependency direction for every MVP capability.
- The MVP can be delivered without introducing microservice, marketplace, hardware, or enterprise complexity.
- API, database, AI, security, and telemetry decisions are explicit enough to guide system design and AI-assisted development.
- User-owned data remains isolated and recoverable, and AI cannot silently become the source of truth.
- The system can be tested at domain, application, adapter, API, performance, security, and AI-evaluation levels.
- Operational teams can diagnose API, PostgreSQL, Redis, and OpenAI failures through OpenTelemetry and Grafana.
- The design supports the product targets of 1,000 registered users and 500 MAU for MVP, with a path toward later growth.
- Engineering measurements can be aligned with the product targets: under 300 ms normal API latency, under 100 ms database query time, at least 99.9% API availability, under 3 seconds average AI response time, under 1% AI error rate, and at least 80% test coverage.
- Future modules and commercialization can be added through stable boundaries rather than rewrites of the MVP core.

## Reference Product Documents

This vision was derived from the documents in `docs/product/`, including the Product Vision, Product Mission, Problem Statement, Target Users, User Personas, Value Proposition, Product Scope, Feature Inventory, Product Roadmap, Success Metrics, Business Goals, Risks & Assumptions, and M1 Sign-off. The approved M1 sign-off authorizes System Design as the next milestone.
