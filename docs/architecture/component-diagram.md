# DapurPintar AI Backend Component Diagram

## Executive Summary

The C4 Level 3 Component Diagram decomposes only the Backend Modular Monolith into logical components. It shows how transport, application orchestration, domain rules, repository abstractions, technical adapters, AI integration, read models, event publication, authentication, and observability collaborate while preserving Clean Architecture and Domain Driven Design.

The Domain Layer remains the center of the model. It owns aggregates, value objects, domain services, domain events, policies, and business invariants. The Application Layer coordinates use cases and transaction boundaries around those domain decisions. Infrastructure implements technical concerns behind abstractions and never becomes a dependency of business rules.

The diagram does not represent Go packages, folders, classes, deployment units, microservices, message brokers, database schema, or API endpoint details. The Backend Modular Monolith remains one MVP logical application.

## Purpose

This component view provides a shared logical model for:

- Understanding the internal responsibilities of the Backend Modular Monolith.
- Protecting the dependency direction required by Clean Architecture.
- Showing where business rules, orchestration, persistence access, AI integration, and cross-cutting concerns belong.
- Preventing direct client-to-data-store or client-to-AI-provider communication.
- Giving future implementation work a stable responsibility map without prescribing source structure.

## Component Overview

| Component | Responsibility category | Primary role |
|---|---|---|
| API Layer | Boundary | Translates user-facing requests and business outcomes at the application boundary. |
| Authentication & Authorization Component | Cross-cutting business protection | Establishes trusted identity and access scope for protected use cases. |
| Application Layer | Orchestration | Coordinates commands, queries, policies, aggregates, repositories, AI assistance, and domain events. |
| Domain Layer | Core business model | Owns aggregates, value objects, domain services, policies, invariants, and domain events. |
| Repository Abstraction | Business persistence boundary | Defines business-facing contracts for retrieving and changing owned aggregates and context. |
| Read Model Component | Query representation | Prepares business views for user and decision questions without owning business truth. |
| AI Integration Component | Application integration boundary | Converts business context into AI assistance requests and maps responses back into product meaning. |
| Event Publishing Component | Internal business boundary | Publishes domain events inside the application boundary without a message broker. |
| Infrastructure Layer | Technical implementation boundary | Implements repository and technical concerns for approved external stores and supporting services. |
| Observability Component | Cross-cutting operational concern | Captures safe traces, metrics, and structured operational signals. |

## Component Responsibilities

### API Layer

The API Layer is the outer application boundary used by the Web Application and future approved consumers.

Responsibilities:

- Accepts user intent and identifies the appropriate application use case.
- Performs boundary-level request interpretation and response mapping.
- Passes authentication context to the Authentication & Authorization Component.
- Returns business outcomes, validation results, and business errors without exposing internal models.
- Exposes read models prepared by the Read Model Component.
- Does not contain business rules, aggregate invariants, repository logic, prompt logic, or provider calls.

### Authentication & Authorization Component

This component protects the business system's identity and access boundary.

Responsibilities:

- Establishes and validates trusted account participation.
- Determines whether a User may access or change personal kitchen context.
- Applies future household participation and visibility rules when that capability exists.
- Provides authorization decisions to application use cases.
- Supports account access lifecycle decisions.
- Does not own User Profile preferences, Pantry truth, Recipe meaning, Meal Plan intent, Shopping intent, or Recommendation suitability.

### Application Layer

The Application Layer is the orchestrator of business use cases and the coordination point between the outer boundary and the Domain Layer.

Responsibilities:

- Coordinates registration, profile management, pantry, recipe, recommendation, meal-planning, and shopping use cases.
- Establishes business transaction boundaries.
- Loads the context needed for a decision through Repository Abstraction contracts.
- Invokes Domain Layer behavior and Domain Services.
- Requests AI assistance through the AI Integration Component.
- Requests Read Model preparation for business queries.
- Sends identified Domain Events to the Event Publishing Component.
- Applies authorization decisions before protected business changes.
- Does not implement business invariants that belong to aggregates or policies.

### Domain Layer

The Domain Layer is the center of the architecture and the owner of business meaning.

Responsibilities:

- Owns the Account, User Profile, Pantry, Recipe, Meal Plan, Shopping List, and Kitchen Recommendation aggregates.
- Owns entities and value objects such as Pantry Item, Planned Meal, Shopping Item, Ingredient Quantity, Expiry Date, Meal Slot, Cooking Time, and Recommendation Status.
- Owns domain services such as Kitchen Context Assessment, Recommendation Suitability, Ingredient Waste Prioritization, Meal Planning Guidance, and Shopping Need Assessment.
- Owns domain policies such as preference-aware recommendation, waste reduction, duplicate-shopping prevention, user decision, privacy boundary, and commercial neutrality.
- Enforces aggregate invariants and business validation rules.
- Defines Domain Events that represent meaningful business facts.
- Owns the business-facing repository abstractions required by the Domain and Application layers.
- Remains independent of Fiber, PostgreSQL, Redis, OpenAI, the AI Gateway, and Web Application concerns.

### Repository Abstraction

Repository Abstraction is the business-facing persistence boundary owned by the Domain/Application side of the architecture.

Responsibilities:

- Defines contracts for Account Records, User Context Records, Pantry Records, Recipe Knowledge, Meal Plan Records, Shopping List Records, and Kitchen Recommendation Records.
- Expresses what business information a use case needs, not how information is stored.
- Preserves aggregate ownership and bounded-context language.
- Prevents SQL, database-specific behavior, or generated persistence types from entering the Domain Layer.
- Supports authoritative access to PostgreSQL through Infrastructure implementations.

It is an abstraction, not a database and not a generic data-access utility.

### Read Model Component

The Read Model Component prepares business views for user and decision questions.

Responsibilities:

- Prepares Pantry Dashboard, Recommendation View, Weekly Meal View, Shopping View, Home Dashboard, Recipe Discovery View, and Profile Context View.
- Combines approved business information for a specific question.
- Preserves traceability to the owning context of each business fact.
- Supports query-focused views without changing aggregate ownership.
- Does not create alternative business truth or make domain decisions.

### AI Integration Component

The AI Integration Component is the Backend Modular Monolith's logical boundary for requesting AI assistance.

Responsibilities:

- Receives an AI-oriented use-case request from the Application Layer.
- Gathers only the approved business context needed for the decision.
- Translates Domain concepts into a provider-independent AI assistance request.
- Communicates only with the AI Gateway container.
- Interprets the Gateway result as candidate product guidance for Domain validation.
- Preserves the distinction between generated assistance and accepted user decisions.
- Does not communicate directly with OpenAI or another External AI Provider.
- Does not own recommendation acceptance, Pantry truth, Recipe truth, Meal Plan intent, or Shopping intent.

### Event Publishing Component

The Event Publishing Component represents publication of Domain Events inside the application boundary.

Responsibilities:

- Receives Domain Events from completed application use cases.
- Preserves event ownership and business names from the Domain Layer.
- Makes events available to approved in-application consumers such as Read Model preparation, product measurement, or future notification decisions.
- Supports business traceability without changing event meaning.
- Does not introduce Kafka, RabbitMQ, NATS, Redis Streams, or another messaging technology.
- Does not become a second business owner for an event or aggregate.

### Infrastructure Layer

The Infrastructure Layer implements technical concerns behind the abstractions owned by the Domain/Application side.

Responsibilities:

- Implements Repository Abstraction for PostgreSQL using the approved SQLC and Goose approach.
- Provides approved Redis access for cache, session management, rate limiting, and temporary state.
- Supplies technical configuration and external technical adapters required by the application boundary.
- Supports infrastructure-level failures and technical error translation.
- Implements interfaces; it does not define business rules or change aggregate ownership.
- Does not call the External AI Provider directly for product AI behavior; AI access remains behind the AI Gateway through AI Integration.

### Observability Component

The Observability Component is cross-cutting and operational rather than business-owning.

Responsibilities:

- Captures request, use-case, repository, Redis, AI, and event-publication signals at appropriate boundaries.
- Correlates business outcomes with latency, availability, error, and dependency behavior.
- Supports AI quality and operational measurements such as response time, error rate, acceptance, and cost signals.
- Applies safe attributes, sampling, and redaction.
- Does not change Domain state or expose sensitive data by default.

## Component Relationships

| From | To | Relationship | Responsibility preserved |
|---|---|---|---|
| API Layer | Authentication & Authorization | Requests identity and access decision | API does not implement access policy. |
| API Layer | Application Layer | Invokes a business use case | Transport does not implement business rules. |
| Application Layer | Domain Layer | Executes aggregate, policy, and service decisions | Domain remains the business authority. |
| Application Layer | Repository Abstraction | Reads and changes owned business context | Persistence is accessed through business contracts. |
| Domain Layer | Repository Abstraction | Defines business persistence needs | Domain does not depend on technical storage. |
| Infrastructure Layer | Repository Abstraction | Implements repository contracts | Technical concerns depend inward. |
| Application Layer | AI Integration Component | Requests contextual AI assistance | Business orchestration remains in Application. |
| AI Integration Component | AI Gateway container | Sends provider-independent AI assistance request | No direct client or provider access. |
| Application Layer | Read Model Component | Requests business views | Read models do not own business truth. |
| Application Layer | Event Publishing Component | Publishes completed domain facts | Publication remains internal and logical. |
| API Layer | Read Model Component | Presents prepared business views | View preparation remains outside transport. |
| Infrastructure Layer | PostgreSQL | Implements authoritative data access | PostgreSQL remains the system of record. |
| Infrastructure Layer | Redis | Uses supporting temporary state | Redis remains non-authoritative. |
| API Layer | Observability Component | Emits boundary signals | Sensitive attributes are controlled. |
| Application Layer | Observability Component | Emits use-case signals | Business outcomes remain observable. |
| AI Integration Component | Observability Component | Emits AI interaction signals | AI quality and cost remain measurable. |
| Event Publishing Component | Observability Component | Emits publication signals | Event behavior is diagnosable without a broker. |

## Cross-Cutting Components

The following components apply across use cases without owning the business domain:

| Cross-cutting component | Applies to | Rule |
|---|---|---|
| Authentication & Authorization | API and Application boundaries | Every protected decision is made within the trusted access scope. |
| Observability | Boundary, application, AI, repository, infrastructure, and event publication | Instrumentation must be useful, safe, and privacy-conscious. |
| Validation and Error Translation | API and Application boundaries | External or technical failures become safe business-facing outcomes. |
| Configuration and Policy Selection | Application and AI Integration boundaries | Policy choices must not bypass Domain ownership. |

Configuration and validation are described as responsibilities of existing logical boundaries, not as additional containers.

## AI Integration Flow

The AI flow follows a strict direction:

1. The API Layer receives a user request for recommendation, AI chat, pantry analysis, or another approved AI-assisted capability.
2. The Application Layer verifies access and gathers the relevant context through Repository Abstraction contracts.
3. The Domain Layer and Domain Services determine the business purpose, constraints, and suitability rules.
4. The AI Integration Component minimizes and translates the approved context into a provider-independent assistance request.
5. The AI Integration Component communicates only with the AI Gateway container.
6. The AI Gateway applies provider integration, prompt orchestration, provider abstraction, safety checks, quotas, time limits, retries, and resilience policies.
7. The AI Gateway communicates with the External AI Provider and returns provider-independent assistance.
8. The AI Integration Component maps the assistance back into a candidate product recommendation.
9. The Application Layer asks the Domain Layer to validate the business meaning and recommendation state.
10. The user receives a recommendation that remains a proposal until accepted or confirmed.

The External AI Provider never becomes the owner of Pantry, Recipe, Meal Plan, Shopping List, User Profile, or Recommendation truth.

## Data Access Strategy

- PostgreSQL is the system of record for durable business data.
- The Domain/Application side owns Repository Abstraction contracts.
- Infrastructure implements those contracts for PostgreSQL using SQLC-generated access and Goose-managed schema evolution.
- The Application Layer coordinates business changes; it does not expose persistence details to the API Layer.
- The Domain Layer remains independent of PostgreSQL and SQLC types.
- Read Model Component queries are optimized for business questions but remain traceable to authoritative source contexts.
- Redis is accessed by Infrastructure for cache, session management, rate limiting, and temporary state only.
- Redis loss or eviction must not remove or redefine authoritative business facts.
- No component other than Infrastructure directly communicates with PostgreSQL or Redis.

## Dependency Rules

1. The Domain Layer is the center of business meaning.
2. Business rules must never depend on Infrastructure, Fiber, PostgreSQL, Redis, OpenAI, or the AI Gateway.
3. The Application Layer depends on Domain abstractions and orchestrates use cases; it does not own aggregate invariants.
4. The API Layer depends on Application use cases and boundary contracts; it does not call repositories, databases, Redis, or AI providers directly.
5. Repository abstractions are owned by the Domain/Application side; Infrastructure implements them.
6. Infrastructure dependencies point inward toward repository and application abstractions.
7. AI Integration communicates only with the AI Gateway and never directly with an External AI Provider.
8. The AI Gateway is not a source of business truth and cannot bypass the Backend Modular Monolith.
9. Read Model Component prepares views and cannot mutate or become authoritative over aggregates.
10. Event Publishing publishes Domain Events inside the application boundary only and introduces no message broker.
11. Cross-cutting components observe or protect behavior without changing Domain ownership.
12. No dependency may cause the MVP Backend Modular Monolith to become a collection of microservices.

## Security Considerations

- Authentication and authorization decisions are enforced before protected Application use cases run.
- User ownership and future household scope are validated in the Application and Domain boundaries.
- API Layer input is validated before it enters Application use cases.
- Repository Abstraction and Infrastructure enforce access to only the business scope authorized by the use case.
- AI Integration sends only minimized, purpose-limited context to the AI Gateway.
- AI Gateway protects provider credentials and applies AI-specific safety and resilience policies.
- AI output is validated and remains decision support, not authoritative business truth.
- Redis keys and temporary state must be scoped to the authorized user or business context.
- Domain Events and Read Models must not expose passwords, tokens, private prompts, or unnecessary pantry/profile data.
- Observability redacts sensitive data and avoids using raw user context as unrestricted telemetry.

## Assumptions

- The Backend Modular Monolith is one MVP logical application.
- The Domain Layer remains independent and central.
- The AI Gateway is a dedicated logical boundary, not a new microservice.
- Repository abstractions remain owned by Domain/Application responsibilities.
- PostgreSQL is authoritative and Redis is non-authoritative.
- Domain Events are published only inside the application boundary for the MVP.
- Read Models are views over owned business facts and do not introduce CQRS infrastructure as a separate system.
- The component model supports the existing strategic and tactical DDD boundaries without adding new bounded contexts.
- Future external systems remain behind their approved logical boundaries.

## Risks

- The Application Layer may become a second location for business rules instead of coordinating Domain behavior.
- The Domain Layer may become coupled to repository or technical types despite dependency rules.
- Infrastructure may bypass repository abstractions and leak storage concerns into business decisions.
- The AI Integration Component may become a second AI Gateway if provider-specific behavior is duplicated.
- The AI Gateway may be mistaken for a separately deployable microservice before measured need exists.
- Read models may accidentally become authoritative or embed business decisions.
- Event Publishing may evolve into an unapproved messaging architecture or hidden workflow engine.
- Redis may be treated as durable business state.
- Observability may leak sensitive prompts, profile context, pantry data, or credentials.
- Component boundaries may be expanded before MVP behavior and domain ownership are validated.

## Diagram Reference

The C4 Level 3 source is maintained in `docs/architecture/diagrams/component-diagram.mmd`.
