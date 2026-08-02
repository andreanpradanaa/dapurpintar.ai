---
name: qa-engineer
description: Use when writing or reviewing tests, test strategy, or quality gates for DapurPintar AI. Covers M6 contract tests, M14 test strategy, acceptance criteria verification, and release gates. Trigger on testing, contract tests, test strategy, quality, regression, or acceptance work.
---

# QA Engineer Skill

## Purpose

Turn documented requirements into verified behavior. Quality is defined at the acceptance criteria and contract level, not by accident of implementation.

## Responsibilities

- Write and maintain contract tests against `docs/api/openapi.yaml`.
- Define and run the test strategy across unit, integration, contract, and E2E layers.
- Map acceptance criteria to test cases.
- Gate releases on quality signals.

## Inputs

- `docs/api/openapi.yaml` and `docs/api/m6-contract-tests.md`.
- `docs/api/m6-error-catalog.md` for error behavior verification.
- `docs/architecture/implementation-readiness.md` for DoR/DoD.
- `backend/` Go tests and `docs/database/` data fixtures.

## Outputs

- Contract tests that fail on contract drift.
- Acceptance criteria with traceable test cases.
- Release gate reports.

## Dependencies

- Backend HTTP contract and error catalog are authoritative.
- Test data must not flow through production paths.

## Status

Active - aligned with M6 contract testing and M14 test strategy planning.
