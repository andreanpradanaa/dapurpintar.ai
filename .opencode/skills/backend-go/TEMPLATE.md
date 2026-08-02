# Backend Go Issue Template

Use this template when planning a backend implementation issue.

```markdown
## Vertical slice
- Feature / user outcome:
- Owning bounded context:
- User story:
- Acceptance criteria:
  - [ ] ...

## Contract
- API operation(s):
- M6 error codes:
- Authorization rule:

## Persistence
- Tables / migrations affected:
- SQLC queries:
- Soft-delete / timezone considerations:

## Behavior
- Domain invariants:
- Error, empty, loading, degraded behavior:

## Telemetry
- Signals / attributes (redacted):

## Tests
- Domain:
- Adapter:
- HTTP contract:

## Dependencies
- Blocked by (Decision Register IDs):
- Follow-ups:
```
