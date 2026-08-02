# ADR-007 Use REST API

- Status: Accepted
- Date: 2026-08-02

## Context

The Next.js web application needs a stable interface to the Go backend. The product roadmap also anticipates future partner and public APIs. MVP clients need resource operations for authentication, users, pantry, recipes, meal plans, shopping lists, and AI requests, with consistent authorization and error behavior.

## Decision

Use a versioned REST/JSON API as the primary application interface, beginning with `/api/v1`. Endpoints will represent product capabilities and resources, while handlers remain thin and delegate to application use cases. The API contract will define schemas, errors, authentication, authorization, pagination, compatibility, and bounded AI behavior.

## Consequences

- REST is familiar and accessible for the Next.js client, testing tools, and future partners.
- HTTP caching, standard status codes, and resource-oriented contracts are available where appropriate.
- Versioning provides a compatibility boundary for future commercialization and public APIs.
- Complex AI workflows may need explicit operation semantics rather than pretending every action is simple CRUD.
- The team must maintain an API specification and prevent internal database or provider models from leaking into responses.

## Alternatives Considered

- **gRPC:** Strong contracts and efficient service-to-service communication, but less direct for browser clients and public API consumption in the MVP.
- **GraphQL:** Flexible client queries, but adds schema and resolver complexity before the product's resource boundaries are validated.
- **Unversioned ad hoc JSON endpoints:** Fast initially, but creates compatibility and client coupling risks.
