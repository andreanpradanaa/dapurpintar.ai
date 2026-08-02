# DapurPintar AI M8 AI Foundation - Sign-off

## Document Control

| Item | Value |
|---|---|
| Milestone | M8 - AI Foundation |
| Deliverable | M8-001 AI Gateway port and OpenAI provider adapter |
| Status | Draft - Awaiting Cross-Functional Review |
| Parent documents | `docs/architecture/ai-architecture.md`, `docs/architecture/m8-blocking-decisions.md`, `docs/architecture/m4-decision-register.md` |
| Scope | AI Gateway port, model profile, structured output validation, and the OpenAI provider adapter |

## Purpose

This document records the M8 AI Foundation deliverable M8-001. It establishes the AI Gateway boundary (ADR-010): a provider-neutral contract that business modules depend on, an OpenAI provider adapter that owns all provider SDK details, and the validation layer that keeps provider output from becoming a product commitment. Prompts, safety policy versioning (M8-002), and the evaluation harness (M8-003) land in later M8 deliverables.

## Decisions Applied

- M4-DEC-010: fixed default model profile with versioned alternatives. `ai.Profile` carries model name, provider, capabilities, context budget, token bound, temperature, seed, and timeout; `ai.DefaultProfile()` returns the local-development default.
- M4-DEC-011: versioned prompt/safety/policy pipeline with layered validation. Requests and results carry `PromptRev`, `SafetyRev`, and `SchemaRev`; output is validated in layers (transport, schema) before a result is returned.
- M4-DEC-016: bounded timeouts and retries. The adapter enforces a per-request operation deadline, bounded automatic retries, and provider quota classification into `AI_QUOTA_EXCEEDED`.

## Deliverables

### M8-001-1 AI Gateway port and types

- `internal/ai/gateway.go`: `ai.Gateway` interface (`Complete`), `Request`/`Result`/`Message`/`Usage`, and the `Purpose` vocabulary (`kitchen-recommendation`, `pantry-analysis`).
- Business modules depend on this contract; no handler, aggregate, or unrelated module calls the provider SDK.

### M8-001-2 Model profile

- `internal/ai/profile.go`: `ai.Profile` with validation and a redacted `String()` identifier for telemetry.
- `ai.DefaultProfile()` default model for local development; concrete model targets validated with DP-SPK-003 and recorded in deployment configuration.

### M8-001-3 Structured output validation

- `internal/ai/schema.go`: versioned `OutputSchema`, the schema-validation layer (`ValidateOutput`), and the Kitchen Recommendation structured-output contract.
- Provider output is never silently repaired into a business fact; invalid output is rejected with a safe error.

### M8-001-4 OpenAI provider adapter

- `internal/ai/openai/adapter.go`: the only package touching the OpenAI SDK.
- Owns credentials, request mapping, structured-output (`json_schema`) invocation, bounded timeout and retry, usage/latency capture, and provider error translation into stable M6 AI codes (`AI_UNAVAILABLE`, `AI_QUOTA_EXCEEDED`, `AI_SAFETY_REJECTED`).
- `translateError` classifies 429 as quota, 400 as safety-rejected, 5xx as unavailable, and deadline as unavailable.

### M8-001-5 Configuration and wiring

- `internal/config`: `AI_PROVIDER`, `AI_API_KEY`, `AI_MODEL`, `AI_TIMEOUT`, `AI_MAX_RETRIES`.
- `.env.example`: documented AI variables; empty `AI_PROVIDER` disables AI and core non-AI operations remain fully usable (fail-closed AI).
- `backend/cmd/api/main.go`: constructs the adapter when `AI_PROVIDER` is set; startup fails if configured AI cannot be built.

## Verified

- `go build ./...` passes.
- `go vet ./...` passes.
- `gofmt -l .` clean.
- `go test -race ./...` passes for the ai and ai/openai packages, plus existing packages.

Test coverage in `internal/ai/openai/adapter_test.go` and `internal/ai/profile_test.go`:

- Successful completion returns validated structured output and usage metadata.
- Request body carries `json_schema` structured output for the purpose.
- Empty content, content-filter finish reason, schema-invalid output, provider 429, and provider 5xx each map to the correct M6 error code.
- Unsupported purposes fail before any provider call.
- Adapter construction rejects a missing API key.
- Profile validation rejects incomplete profiles; schema validation rejects empty, malformed, non-object, and missing-field output.

## Dependencies on Pending Decisions

- M8-002 (prompt, safety, and structured-output policy versioning) will introduce the versioned policy store and the system-prompt assembly that feeds `Request.Messages`.
- M8-003 (AI evaluation harness) will consume the `Result` metadata (revisions, model, usage, latency) for regression evaluation per M4-DEC-012.
- M4-DEC-016 global/per-user budget enforcement and alert thresholds land with DP-SPK-009; the adapter already classifies quota failures and records usage.
- The `pantry-analysis` purpose is registered in the vocabulary but its output schema is finalized with DP-FEAT-008.

## Exit Criteria

M8-001 is complete when:

- The AI Gateway contract is provider-neutral and product-vocabulary-driven.
- The OpenAI adapter is the only provider SDK consumer and is not reachable from handlers or domain code.
- Requests and results carry policy revision metadata for reproducibility.
- Structured output validation rejects invalid output with safe M6 errors.
- Timeout, retry, and quota behavior are bounded and tested.
- Config, wiring, and documentation are consistent with M8-001 scope.

## Related Documents

- `docs/architecture/ai-architecture.md`
- `docs/architecture/m8-blocking-decisions.md`
- `docs/architecture/m4-decision-register.md`
- `docs/architecture/implementation-backlog.md`
- `docs/api/m6-error-catalog.md`
- `docs/database/m5-schema.md`
