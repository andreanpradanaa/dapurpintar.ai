# ADR-002 Use Fiber as HTTP Framework

- Status: Accepted
- Date: 2026-08-02

## Context

The backend requires a stable HTTP boundary for the Next.js client and future API consumers. The MVP needs routing, middleware, authentication integration, validation, error mapping, and observability without allowing transport concerns to leak into domain logic.

## Decision

Use Fiber as the Go HTTP framework. Fiber will be confined to the transport and composition boundary: routes, middleware, request decoding, response mapping, and HTTP-specific error handling. Application use cases and domain rules will not depend on Fiber.

## Consequences

- The project receives a lightweight, productive framework aligned with the Go backend decision.
- Middleware provides a consistent place for authentication, request context, rate limiting, and telemetry integration.
- Thin handlers can preserve the Clean Architecture dependency direction.
- Fiber becomes a framework dependency that requires version management and framework-specific operational knowledge.
- A framework migration would still require adapter changes, but should not require domain changes if the boundary is respected.

## Alternatives Considered

- **Go standard library `net/http`:** Maximum minimalism and portability, but requires more application-level assembly for routing and common HTTP concerns.
- **Gin:** Mature Go web framework, but Fiber is the selected project constraint and provides an appropriate API framework for the MVP.
- **Echo:** Capable alternative, but introducing another framework would not provide enough benefit to justify divergence from the architecture vision.
