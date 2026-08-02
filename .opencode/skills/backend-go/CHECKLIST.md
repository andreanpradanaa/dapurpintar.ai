# Backend Go Checklist

## Before starting an implementation issue

- [ ] Read `docs/architecture/implementation-readiness.md` and the M4 Definition of Ready.
- [ ] Confirm the owning bounded context is named.
- [ ] Confirm the API operation and M6 error codes that apply.
- [ ] Confirm authorization scope is explicit.
- [ ] Confirm persistence impact is known (tables, migrations, SQLC queries).
- [ ] Confirm error, empty, loading, and degraded behavior is defined.
- [ ] Confirm no blocking Decision Register item remains unresolved.

## While implementing

- [ ] Domain behavior and invariants are tested without external dependencies.
- [ ] Repositories are defined in the domain language and implemented with SQLC.
- [ ] Migrations are forward-compatible and reviewed for ownership preservation.
- [ ] Handlers return the M6 envelope and stable error codes.
- [ ] Authorization is enforced server-side for every protected operation.
- [ ] No implicit cross-context mutation is introduced.
- [ ] Telemetry is correlated and redacted.

## Before finishing

- [ ] `go build ./...` passes.
- [ ] `go vet ./...` passes.
- [ ] `gofmt -l .` is clean.
- [ ] `go test -race ./...` passes.
- [ ] Documentation and decision records are updated.
- [ ] The change is reviewable as one focused unit.

## Definition of Done (M4)

- [ ] Domain behavior and invariants are tested.
- [ ] API and persistence adapters follow approved contracts.
- [ ] Ownership and authorization are enforced server-side.
- [ ] No implicit cross-context mutation.
- [ ] Error and degraded states are safe and actionable.
- [ ] Telemetry is correlated and redacted.
- [ ] Relevant tests pass.
- [ ] Documentation and decision records are updated.
- [ ] The change is reviewable as one focused unit.
