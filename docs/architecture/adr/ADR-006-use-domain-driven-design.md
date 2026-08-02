# ADR-006 Use Domain Driven Design

- Status: Accepted
- Date: 2026-08-02

## Context

The product is more than a recipe search application. It coordinates user preferences, pantry stock and expiry, recipes, favorites, daily and weekly meal plans, shopping lists, and AI recommendations. Future capabilities include family collaboration and nutrition. The architecture needs boundaries that reflect product language and ownership rather than technical layers alone.

## Decision

Use Domain Driven Design to shape the modular monolith around business capabilities and a shared product vocabulary. Initial bounded areas are Authentication, User, Pantry, Recipe, Meal Planner, Shopping, and AI Assistant. Detailed aggregates, invariants, domain events, and context maps will be defined during System Design.

## Consequences

- Module boundaries align engineering decisions with user and business capabilities.
- Domain rules such as ownership, expiry, plan dates, and list generation can remain explicit and testable.
- The model can evolve toward family workspaces, nutrition, notifications, and partner integrations without a single global domain model.
- DDD requires collaboration with Product and careful vocabulary management.
- Some early boundaries may change after user validation and Event Storming; the architecture must allow refinement without treating current boundaries as immutable.

## Alternatives Considered

- **Feature folders without domain modeling:** Lower initial ceremony, but risks duplicated rules and unclear ownership as features interact.
- **Database-first modeling:** Convenient for CRUD, but database tables would drive the business model rather than the user and domain language.
- **One shared domain model:** Simpler conceptually, but creates coupling across MVP and future business capabilities.
