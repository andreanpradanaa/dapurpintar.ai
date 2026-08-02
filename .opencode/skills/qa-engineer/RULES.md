# QA Engineer Rules

## Contract is the source of truth

- The OpenAPI contract and M6 error catalog define what must hold in production.
- Contract drift (API no longer matching `openapi.yaml`) is a blocking defect.
- Error responses must use stable codes from the error catalog.

## Layered test strategy

- Unit tests cover domain invariants and rules with no external services.
- Integration tests cover adapters against real PostgreSQL/Redis in a test environment.
- Contract tests verify the HTTP boundary shape for every operation.
- E2E covers the critical user journeys end to end.

## Traceability

- Every acceptance criterion maps to at least one test.
- Every test names the criterion or user story it protects.
- Regressions are added at the layer where the bug was introduced.

## Fixtures and data

- Test data uses the approved fixtures; production paths are never used for seeding.
- Tests are deterministic: time, timezone, and randomness are controlled.
- No test depends on shared mutable state or test execution order.

## Release gates

- A milestone ships only when its contract, acceptance, and regression tests pass.
- Performance, security, and accessibility checks are part of the gate where specified.
- `go test -race ./...` is part of every backend gate.

## Review

- Test reviews check assertions for meaning, not just coverage numbers.
- A passing test that only asserts the happy path is incomplete.
