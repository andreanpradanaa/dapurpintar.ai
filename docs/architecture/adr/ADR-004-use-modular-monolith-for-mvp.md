# ADR-004 Use Modular Monolith for MVP

- Status: Accepted
- Date: 2026-08-02

## Context

DapurPintar AI must validate its product assumptions quickly while supporting an AI-first SaaS direction. The MVP has several related capabilities but a small initial scale and a single product team. The architecture vision requires clear module boundaries without prematurely introducing distributed-system operations.

## Decision

Use a modular monolith for the MVP. One deployable backend will contain cohesive modules with explicit contracts, application boundaries, domain ownership, and persistence boundaries. Modules must not bypass another module's interfaces or write another module's tables directly.

## Consequences

- Deployment, local development, testing, and operations remain simpler than with microservices.
- Cross-module workflows can use in-process application contracts while retaining business boundaries.
- The product can focus on validating weekly AI-assisted meal planning and other MVP outcomes.
- A poorly enforced monolith could become a shared, coupled codebase; module ownership and dependency rules must be reviewed continuously.
- Future service extraction remains possible where measured scaling, isolation, or team ownership justifies its cost.

## Alternatives Considered

- **Microservices:** Independent scaling and deployment are attractive, but network contracts, distributed tracing, deployment, and data ownership would add complexity before MVP validation.
- **Unstructured monolith:** Fast initially, but it would undermine maintainability and make future evolution harder.
- **Serverless functions:** Useful for isolated workloads, but less suitable as the primary structure for cohesive transactional domains and AI orchestration.
