# DapurPintar AI Container Diagram

## Executive Summary

The C4 Level 2 Container Diagram decomposes the DapurPintar AI system boundary into logical containers and shows how those containers collaborate with actors, data stores, and external systems. It is the bridge between the Level 1 System Context view and later detailed design.

The MVP remains a Modular Monolith. The Backend Modular Monolith is the single owner of business rules, domain logic, application orchestration, authorization decisions, and coordination of the kitchen bounded contexts. The AI Gateway is a dedicated logical container boundary for provider integration and AI-specific policies; it is not an instruction to create a separate microservice or deployment unit.

The diagram intentionally excludes deployment topology, infrastructure environments, internal source structure, database schema, REST resource design, and message-broker behavior.

## Purpose

This diagram establishes:

- The logical responsibilities inside DapurPintar AI.
- The boundary between user interaction and business decision-making.
- The boundary between business orchestration and external AI-provider integration.
- The role of PostgreSQL as the system of record.
- The limited role of Redis as supporting temporary state and control information.
- The cross-cutting role of observability without making it a business owner.
- The approved communication direction for the MVP.

It should be used with the Level 1 System Context Diagram, the bounded-context documents, the tactical DDD model, and the domain event catalog.

## Container Overview

| Container | Type | MVP status | Primary responsibility |
|---|---|---|---|
| Web Application | User-facing application | MVP | Presents the kitchen experience and captures user intent. |
| Backend Modular Monolith | Logical business application | MVP | Owns business rules, domain logic, authorization, and orchestration across bounded contexts. |
| AI Gateway | Logical provider-boundary application | MVP | Owns AI provider integration, prompt orchestration, provider abstraction, safety boundary, and resilience policies. |
| PostgreSQL | Business data store | MVP | System of record for authoritative business data. |
| Redis | Supporting data store | MVP | Cache, session management, rate limiting, and temporary state only. |
| Observability | Cross-cutting operational container | MVP | Receives telemetry for diagnosis and operational insight; owns no business decisions or data. |

### Logical container interpretation

Containers represent logical responsibilities within the DapurPintar AI system boundary. The Backend Modular Monolith and AI Gateway may share the same MVP application boundary while remaining separate architectural responsibilities. No container in this view implies a microservice, separate deployment, or independent infrastructure environment.

## Container Responsibilities

### Web Application

The Web Application is the user-facing entry point for the DapurPintar AI experience.

Responsibilities:

- Supports registration, profile completion, pantry management, recipe discovery, meal planning, shopping, and AI interaction.
- Presents read models such as pantry, recommendation, weekly meal, shopping, and home views.
- Captures user commands and displays business outcomes.
- Makes the distinction between recommendations, plans, and shopping intentions visible to the user.
- Does not own business rules, aggregate invariants, pantry truth, recommendation suitability, or AI-provider behavior.
- Does not communicate directly with PostgreSQL, Redis, the AI Gateway, or external systems.

Technology choice: Next.js, React, TypeScript, and Tailwind CSS.

### Backend Modular Monolith

The Backend Modular Monolith is the central business container and the sole business orchestrator for the MVP.

Responsibilities:

- Owns the Authentication, User, Pantry, Recipe, Meal Planner, Shopping, and AI Assistant business modules.
- Executes application use cases and coordinates domain services.
- Owns business rules, aggregate invariants, policies, authorization decisions, and context ownership.
- Reads and writes authoritative business data through the repository boundary.
- Uses Redis only for approved supporting concerns.
- Requests AI assistance through the AI Gateway and validates the resulting product meaning.
- Maintains the distinction between external AI output and user-confirmed business decisions.
- Exposes the system's product-facing application boundary to the Web Application and future approved consumers.

Technology choice: Go and Fiber, organized as a Modular Monolith using Clean Architecture and Domain Driven Design.

### AI Gateway

The AI Gateway is the dedicated logical boundary between DapurPintar AI's business system and an External AI Provider.

Responsibilities:

- Receives AI assistance requests from the Backend Modular Monolith only.
- Owns provider abstraction and provider-specific integration behavior.
- Owns prompt orchestration, model selection policy, context minimization, and provider request shaping.
- Applies AI-specific validation, safety controls, provider error handling, time limits, retries, quotas, and resilience policies.
- Returns provider-independent assistance to the Backend Modular Monolith.
- Provides AI usage, quality, latency, and cost signals for business and operational observation.
- Does not own pantry truth, recipe truth, meal commitments, shopping commitments, user preferences, or recommendation acceptance.
- Does not communicate directly with the Web Application or Users.

Technology choice: OpenAI through an internal AI Gateway / provider abstraction. The technology is an implementation choice behind this logical boundary.

### PostgreSQL

PostgreSQL is the authoritative business data store.

Responsibilities:

- Holds durable business information owned by DapurPintar AI.
- Preserves relationships, ownership, and consistency for account, profile, pantry, recipe, favorite, meal-plan, shopping, and required AI business data.
- Supports the Backend Modular Monolith as the source of truth.

PostgreSQL does not make business decisions, call external systems, or become a shared data store for the Web Application or AI Gateway.

Technology choice: PostgreSQL, accessed through the approved repository and SQLC approach. Schema and query design are outside this document.

### Redis

Redis is a supporting data store and is never authoritative for business data.

Permitted responsibilities:

- Cache safe and appropriately scoped read information.
- Support short-lived session management state.
- Support rate limiting and abuse controls.
- Hold temporary coordination or bounded transient state.

Redis must not be used as the system of record for pantry, recipe, meal, shopping, account, or recommendation truth. Loss of Redis state must not create a business-data loss.

Technology choice: Redis.

### Observability

Observability is a cross-cutting logical container for operational visibility.

Responsibilities:

- Receives traces, metrics, and structured operational signals from relevant containers.
- Supports diagnosis of user-facing behavior, database behavior, Redis behavior, AI-provider behavior, latency, availability, and business outcomes.
- Supports business measures such as AI usage, recommendation acceptance, meal plans, pantry activity, and shopping-list generation.
- Applies privacy-conscious redaction and avoids secrets or unnecessary personal context.
- Does not own business facts, change domain state, or make kitchen decisions.

Technology choice: OpenTelemetry for instrumentation and Grafana for visualization, as specified by the approved architecture.

## Container Relationships

| From | To | Relationship | Business meaning |
|---|---|---|---|
| User | Web Application | Uses | Uses the kitchen companion to express intent and review outcomes. |
| Web Application | Backend Modular Monolith | Requests product capabilities | Sends user intent and receives business views and outcomes. |
| Backend Modular Monolith | PostgreSQL | Reads and writes authoritative business data | Maintains the source of truth for owned business contexts. |
| Backend Modular Monolith | Redis | Uses supporting temporary state | Improves access and controls abuse without owning business truth. |
| Backend Modular Monolith | AI Gateway | Requests AI assistance | Asks for contextual support while retaining business ownership. |
| AI Gateway | External AI Provider | Requests provider assistance | Exchanges minimized context for generated AI assistance. |
| Backend Modular Monolith | Observability | Emits operational and business signals | Makes decisions, failures, and outcomes diagnosable. |
| Web Application | Observability | Emits user-facing signals | Supports product and experience visibility where appropriate. |
| AI Gateway | Observability | Emits AI-provider signals | Supports AI quality, latency, availability, and cost visibility. |
| PostgreSQL | Observability | Emits data-store signals | Supports database health and performance visibility. |
| Redis | Observability | Emits supporting-store signals | Supports cache, session, rate-limit, and temporary-state visibility. |

### Future relationships

Future actors and external systems from the approved System Context view interact through the Backend Modular Monolith or an approved business boundary. They do not bypass the system's business ownership:

- Household Members will use the future household collaboration capabilities through the Backend Modular Monolith.
- Nutrition Professionals will participate through future Nutrition Guidance capabilities.
- Grocery Partners and Grocery Marketplaces will interact through future Shopping Optimization and commercial boundaries.
- Authentication Providers, Notification Providers, and Payment Providers will be introduced only for their approved future responsibilities.
- Commercial Operators will interact with future SaaS Commercial Operations without changing kitchen-domain meaning.

## External Systems

### External AI Provider (MVP)

The External AI Provider is accessed only through the AI Gateway. It supplies generative assistance for recipe recommendation, pantry analysis, AI chat, and future AI-assisted planning. It does not receive direct requests from the Web Application and does not own DapurPintar AI business decisions.

### Authentication Provider (Future)

An Authentication Provider may support future identity or verification scenarios. It remains outside the system boundary and cannot own User Profile, Pantry, Recipe, Meal Plan, Shopping List, or Recommendation meaning.

### Grocery Marketplace (Future)

A Grocery Marketplace may support future approved purchase-intent exchange. Ordering, offers, inventory, pricing, and fulfillment remain outside DapurPintar AI.

### Notification Provider (Future)

A Notification Provider may deliver future expiry or meal reminders. DapurPintar AI owns why a reminder matters; the provider owns external delivery.

### Payment Provider (Future)

A Payment Provider may process future subscription payments. DapurPintar AI owns product participation and entitlement meaning; the provider owns payment processing and settlement.

## Technology Choices (High-Level Only)

| Logical concern | Approved choice |
|---|---|
| User-facing application | Next.js, React, TypeScript, Tailwind CSS |
| Business application | Go with Fiber |
| Business architecture | Modular Monolith, Clean Architecture, Domain Driven Design |
| Product interface | Versioned REST/JSON API, represented here only as a logical communication boundary |
| Authoritative data store | PostgreSQL |
| Supporting temporary state | Redis |
| Database access approach | SQLC through repository boundaries |
| Schema change approach | Goose |
| AI provider boundary | AI Gateway abstraction with OpenAI as initial provider |
| Observability | OpenTelemetry and Grafana |

These choices describe technology responsibilities at a high level. They do not define source code structure, schemas, deployment, infrastructure, or operational topology.

## Communication Overview

### User interaction path

1. The User expresses an intention through the Web Application.
2. The Web Application asks the Backend Modular Monolith to perform a product capability.
3. The Backend Modular Monolith applies authorization, policies, aggregate rules, and business orchestration.
4. The Backend Modular Monolith reads or changes authoritative business information in PostgreSQL.
5. The Backend Modular Monolith may use Redis for approved supporting concerns.
6. When AI assistance is appropriate, the Backend Modular Monolith requests it from the AI Gateway.
7. The AI Gateway communicates with the External AI Provider and returns provider-independent assistance.
8. The Backend Modular Monolith validates and interprets the assistance as product meaning before returning an outcome to the Web Application.

### Communication constraints

- Web Application communicates with the Backend Modular Monolith, not directly with data stores or the AI Provider.
- Backend Modular Monolith communicates with the AI Gateway, not directly with the External AI Provider for product AI behavior.
- AI Gateway communicates with the External AI Provider, not with the User or Web Application.
- PostgreSQL is accessed by the Backend Modular Monolith as the authoritative store.
- Redis is accessed only by the Backend Modular Monolith for supporting temporary concerns.
- Future external systems must be introduced through the boundary that owns their business relationship.

## Boundary of Responsibility

### DapurPintar AI owns

- User and future household kitchen context within the approved access scope.
- Pantry availability and expiry meaning.
- Recipe and cooking-guidance meaning within the product.
- Recommendation relevance, rationale, safety, lifecycle, and acceptance state.
- Planned meal intent and shopping intent.
- Business rules, domain policies, authorization, and cross-context orchestration.
- The distinction between AI assistance and confirmed user decisions.
- Product-level commercial participation and future entitlements.

### DapurPintar AI does not own

- External AI model behavior, provider policies, or provider availability.
- External identity verification or authentication-provider operations.
- Grocery marketplace inventory, pricing, offers, ordering, or fulfillment.
- External notification delivery.
- Payment processing and settlement.
- A user's independent nutrition or medical responsibility.
- Internal behavior of future household members or commercial partners outside the product scope.
- Deployment environments, cloud infrastructure, network topology, or runtime operations in this Level 2 view.

## Security Considerations

- The Backend Modular Monolith is the authority for authentication, authorization, ownership, and future household scope.
- The Web Application must not bypass the Backend Modular Monolith to access business data or AI assistance.
- The AI Gateway must minimize personal and kitchen context sent to the External AI Provider.
- Provider credentials and external-system secrets remain behind their responsible logical boundary.
- AI output must be validated, safety-checked, and treated as decision support rather than business truth.
- PostgreSQL access is limited to the Backend Modular Monolith's business responsibility.
- Redis data must be scoped, short-lived where appropriate, and incapable of becoming authoritative business state.
- Observability signals must redact passwords, tokens, sensitive prompts, and unnecessary pantry or profile details.
- Future household sharing must enforce explicit participation and visibility rules.
- Rate limiting and abuse controls apply to account and AI-intensive business capabilities through approved supporting state.

## Assumptions

- The MVP has one logical Backend Modular Monolith and does not introduce microservices.
- The AI Gateway is a dedicated logical boundary, not a separate MVP deployment commitment.
- PostgreSQL remains the authoritative source for business data.
- Redis remains non-authoritative and supports only caching, session management, rate limiting, and temporary state.
- OpenAI is represented as the initial External AI Provider.
- Future external systems remain optional and do not block the MVP journey.
- The Web Application is the initial user-facing consumer.
- Internal bounded contexts remain owned by the Backend Modular Monolith even when represented as separate business modules.
- Container relationships describe logical responsibility, not deployment topology.

## Risks

- The AI Gateway could be mistaken for a separate microservice, causing premature distribution of the MVP.
- The Backend Modular Monolith could become a shared, unstructured application if bounded-context ownership is not enforced.
- Direct client-to-AI or client-to-data-store access would bypass business rules, privacy controls, and user decision boundaries.
- Redis could be misused as authoritative business storage, creating data-loss or consistency risks.
- External AI-provider latency, cost, availability, or policy changes could affect recommendation quality and user trust.
- Prompt orchestration could leak personal, household, pantry, or sensitive context if minimization is not maintained.
- Future partner and commercial relationships could expand the logical boundary beyond the approved product scope.
- Observability could expose sensitive business or personal information if redaction and purpose limitation are not enforced.
- Adding containers for future capabilities before validated business need would increase architectural complexity and contradict MVP-first principles.

## Diagram Reference

The C4 Level 2 source is maintained in `docs/architecture/diagrams/container-diagram.mmd`.
