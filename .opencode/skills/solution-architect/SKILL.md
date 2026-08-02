---
name: solution-architect
description: Use when making or reviewing architecture decisions for DapurPintar AI. Covers the ADR protocol, the Decision Register (M4-DEC-*) workflow, cross-document consistency, bounded-context and DDD guidance, and milestone sign-off reviews. Trigger on architecture, decision, ADR, roadmap, milestone, or cross-cutting design work.
---

# Solution Architect Skill

## Purpose

Keep the architecture coherent and documented so each milestone moves the system forward without contradicting prior decisions. All architecture changes flow through the Decision Register and ADR protocol.

## Responsibilities

- Propose and review architecture decisions as ADRs.
- Maintain the Decision Register and cross-document consistency.
- Confirm every implementation milestone is blocked by named decisions, not silent assumptions.
- Own the roadmap, milestone list, and sign-off protocol.
- Keep the diagrams, contracts, and decision records aligned.

## Inputs

- `docs/architecture/` decision registers, backlog, readiness, ADRs, and design docs.
- `PROJECT_ROADMAP.md` and `docs/project/milestone-list.md`.
- `docs/api/openapi.yaml`, `docs/database/`, `docs/backend/` artifacts.
- `.opencode/contexts/architecture.md` and `.opencode/AI_RULES.md`.

## Outputs

- ADRs and updated Decision Register entries.
- Consistent, reviewable milestones with explicit decision blockers.
- Refreshed roadmap, milestone list, and sign-off records.

## Dependencies

- Approved decisions must not be reopened without a superseding ADR.
- Pending decisions must be resolved before their blocked milestone is implemented.

## Status

Active - aligned with milestone review and M8 planning.
