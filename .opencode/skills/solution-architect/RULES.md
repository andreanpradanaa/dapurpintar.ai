# Solution Architect Rules

## Decision protocol

- Every architecture decision is recorded with a stable ID and status in the Decision Register.
- Approved decisions are not silently reopened. Superseding requires a new ADR that references the superseded ID.
- Pending decisions remain explicit blockers; they are never resolved silently inside implementation work.

## ADR structure

- Use the ADR format in `docs/architecture/` (Context, Decision, Consequences, Status, References).
- Each ADR states what changed and what is intentionally unchanged.
- ADRs cross-reference the decision register ID they implement or supersede.

## Cross-document consistency

- A change to a decision must update every document that depends on it: API contract, database schema, backend layout, AI architecture, roadmap.
- OpenAPI, schema, and backend artifacts must agree on names, ownership, and behavior.
- Milestones and the roadmap list the same set of deliverables as the implementation backlog.

## Ownership and boundaries

- Bounded contexts stay the unit of ownership for data and behavior.
- Shared models across contexts are prohibited; identity references only.
- Diagrams must match the codebase layout and the documentation.

## Review gates

- Each milestone is reviewed before the next starts; status flips to Approved only with explicit sign-off.
- A milestone is not approved while a blocking decision is pending.
- Review comments become follow-up issues, not silent amendments.

## Documentation

- Write for the next engineer and the next architect, not for the current mood.
- Prefer decisions over preferences: state rationale, alternatives considered, and consequences.
