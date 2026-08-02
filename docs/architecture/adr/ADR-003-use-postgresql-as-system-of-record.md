# ADR-003 Use PostgreSQL as System of Record

- Status: Accepted
- Date: 2026-08-02

## Context

The MVP manages related user-owned data: identity, preferences, pantry items, recipes, favorites, meal plans, shopping lists, and AI request metadata. These records have relationships, ownership rules, and transactional consistency requirements. The technical assumptions explicitly identify PostgreSQL as the candidate system of record for the initial scale.

## Decision

Use PostgreSQL as the authoritative system of record. Relational constraints, transactions, indexes, and explicit ownership boundaries will protect the integrity and isolation of product data. Redis may support cache and short-lived coordination but will not be authoritative.

## Consequences

- Related kitchen data can be modeled and updated with strong transactional guarantees.
- PostgreSQL supports the MVP scale and provides a path to indexing, pooling, replicas, and measured optimization.
- A single authoritative store simplifies recovery, backup, migration, and operational reasoning.
- Schema design and query performance require deliberate review as modules and tenants grow.
- Highly specialized or very large future workloads may require read models or additional stores, but only after measured need.

## Alternatives Considered

- **Document database:** Flexible schemas could help evolving AI data, but the MVP's ownership, relationships, and transactional workflows favor relational integrity.
- **Redis as primary storage:** Low latency is useful for ephemeral data, but Redis is not appropriate as the durable source for core user and business records.
- **Multiple databases per module:** Stronger isolation is possible, but it adds operational complexity that is not justified for an MVP modular monolith.
