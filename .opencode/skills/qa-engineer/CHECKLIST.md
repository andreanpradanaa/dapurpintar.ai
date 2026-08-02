# QA Engineer Checklist

## Contract tests
- [ ] Every operation in `openapi.yaml` has a contract test.
- [ ] Success and error shapes match the M6 envelope and error catalog.
- [ ] Contract drift fails the gate.

## Acceptance criteria
- [ ] Each criterion maps to a test case.
- [ ] Happy path, error path, empty path, and edge cases covered.

## Quality gates
- [ ] Unit, integration, contract tests pass.
- [ ] `go test -race ./...` passes.
- [ ] Regression tests added at the correct layer.
- [ ] Tests are deterministic and isolated.
- [ ] Release gate report produced.
