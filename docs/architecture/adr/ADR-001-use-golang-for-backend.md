# ADR-001 Use Golang for Backend

- Status: Accepted
- Date: 2026-08-02

## Context

DapurPintar AI needs a backend for an AI-first SaaS MVP with authentication, pantry management, recipes, meal planning, shopping lists, and AI-assisted workflows. The product roadmap requires a production API that can start simply and support future growth. The architecture also requires maintainability, testability, performance, and clear operational behavior.

## Decision

Use Golang as the backend programming language. Go will host the modular monolith and its application, domain, and infrastructure adapters, while keeping business rules independent from framework and provider details.

## Consequences

- The team gains a compiled, strongly typed runtime with straightforward concurrency and a small deployment footprint.
- Go is well suited to stateless APIs, database access, and bounded integration calls to AI providers.
- A consistent language across backend modules reduces cognitive overhead for a modular monolith.
- The team must establish Go coding standards, testing conventions, error handling, and dependency management.
- Developers without Go experience may require onboarding, and some product libraries may be less mature than their equivalents in larger ecosystems.

## Alternatives Considered

- **Node.js/TypeScript backend:** Strong ecosystem and shared language with the frontend, but Go better matches the required backend runtime and provides a simpler compiled operational footprint.
- **Python:** Strong AI ecosystem, but less suitable as the primary transactional API runtime for this MVP; AI access can remain behind an integration boundary.
- **Java/Kotlin:** Mature enterprise tooling, but higher runtime and operational complexity than needed for the MVP.
