# Solution Architect Examples

## Decision Register entry

```markdown
## M4-DEC-013 - Recommendation conversation retention
Status: Approved

Context: Conversations contain personal cooking context. Full retention
conflicts with privacy expectations and M4-DEC-007.

Decision: Retain conversation data only while a Recommendation is active,
and for no longer than the 30-day window. Never store raw prompts.

Consequences: Simpler privacy story; acceptance and rationale retained as
durable business records for evaluation.

References: `docs/architecture/m4-m5-blocking-decisions.md`
```

## Cross-document consistency check

```text
Rename ingredient "expiration" -> "expiry" in schema MUST also update:
  - docs/database/m5-schema.md
  - docs/api/openapi.yaml (Ingredient schema)
  - backend SQLC queries and view models
  - ai-architecture.md examples if they reference the field
```

## Superseding a decision

```markdown
M4-DEC-020 supersedes M4-DEC-016 for cost-control limits.
Status: Approved
References: M4-DEC-016 (superseded), M4-DEC-020 (this ADR).
```
