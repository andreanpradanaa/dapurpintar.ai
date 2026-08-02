# ADR-005 Use Clean Architecture

- Status: Accepted
- Date: 2026-08-02

## Context

The product combines domain workflows with external concerns including Fiber, PostgreSQL, Redis, OpenAI, and telemetry. AI provider policies and infrastructure details may change, while core rules for pantry, recipes, meal plans, shopping, and authorization must remain testable and maintainable.

## Decision

Use Clean Architecture with dependencies pointing inward. The domain layer contains business rules, the application layer coordinates use cases and ports, and infrastructure and transport adapters implement those ports. Frameworks and external providers must not become dependencies of the domain.

## Consequences

- Domain and application behavior can be tested without HTTP, databases, caches, or live AI providers.
- Provider and framework changes are localized to adapter or composition boundaries.
- The structure makes responsibilities and dependency direction explicit for future contributors and AI-assisted development.
- Additional interfaces and mapping code introduce up-front ceremony.
- Teams must avoid creating abstractions without a real boundary or allowing convenience imports to bypass the architecture.

## Alternatives Considered

- **Traditional layered architecture:** Easier to start, but often permits business rules to depend directly on frameworks and persistence.
- **Framework-first architecture:** Faster for simple CRUD, but creates stronger coupling and weakens provider replacement and testing.
- **Microservice-per-domain architecture:** Not appropriate for the MVP deployment and operational constraints; Clean Architecture provides the needed isolation inside one deployable system.
