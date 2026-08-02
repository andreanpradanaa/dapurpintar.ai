---
name: technical-writer
description: Use when writing or reviewing documentation for DapurPintar AI. Covers markdown structure, document conventions, cross-referencing, and keeping docs in sync with the codebase and decisions. Trigger on documentation, doc writing, README, ADR, or doc review work.
---

# Technical Writer Skill

## Purpose

Produce documentation that engineers and architects can trust: accurate, cross-referenced, and current.

## Responsibilities

- Write and structure project documents consistently.
- Keep documents linked to the artifacts they describe.
- Flag stale content and broken references.

## Inputs

- Existing docs under `docs/`, `PROJECT_ROADMAP.md`, and milestone list.
- Decision registers and ADRs for authoritative positions.
- Actual code and contracts for what is true today.

## Outputs

- Markdown documents following the repo conventions.
- Cross-references between decision, contract, schema, and code artifacts.
- Reviews that catch staleness before it spreads.

## Dependencies

- Documents must describe the approved state; the author writes to the decision records.
- Code examples must match the real API and backend layout.

## Status

Active - supports documentation across all milestones.
