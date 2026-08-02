---
name: ai-engineer
description: Use when implementing or reviewing AI capabilities for DapurPintar AI, including the AI Gateway, OpenAI provider adapter, prompt and policy versioning, structured output validation, safety/retry/quota/cost controls, and the AI evaluation harness (M8). Trigger on AI Gateway, OpenAI, prompt, structured output, recommendation, evaluation, or AI safety work.
---

# AI Engineer Skill

## Purpose

Guide AI engineering work so AI remains an optional decision-support dependency, never a source of business truth. All AI traffic passes through the AI Gateway abstraction; business modules never call the provider SDK directly.

## Responsibilities

- Implement and maintain the AI Gateway port and provider adapters.
- Version prompts, safety policy, and structured-output schemas.
- Enforce safety, timeout, retry, quota, and cost controls.
- Build and run the AI evaluation harness.
- Keep AI output distinguishable from user-confirmed decisions.

## Inputs

- `docs/architecture/ai-architecture.md` for the gateway boundary.
- `docs/api/openapi.yaml` for AI operation contracts.
- `docs/database/m5-schema.md` and `docs/architecture/m4-m5-blocking-decisions.md` for the retention policy (M4-DEC-013).
- `docs/architecture/m4-decision-register.md` for AI decisions M4-DEC-010/011/012/016.
- `docs/architecture/tactical-ddd.md` for the Recommendation lifecycle and invariants.

## Outputs

- AI Gateway port and OpenAI adapter that build and pass tests.
- Versioned prompts, safety policy, and structured-output schemas.
- Timeout, retry, quota, and cost controls with fail-closed behavior.
- Evaluation harness and representative test scenarios.
- No raw provider prompts or payloads stored in business data.

## Dependencies

- Go backend foundation (M7) and the M6 API contract.
- M4-DEC-010 (model profile), M4-DEC-011 (safety/prompt policy), M4-DEC-012 (evaluation), M4-DEC-016 (budget) before production AI.
- Recommendation lifecycle rules from Tactical DDD.

## Status

Active - M8 AI Foundation current; decision records M4-DEC-010/011/012/016 in review.
