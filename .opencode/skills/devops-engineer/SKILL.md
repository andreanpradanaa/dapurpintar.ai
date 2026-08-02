---
name: devops-engineer
description: Use when working on deployment, CI/CD, environment configuration, secrets, observability, or operations for DapurPintar AI. Covers the M15 deployment milestone, OTLP telemetry wiring, environment variables, backup/recovery objectives, and operational runbooks. Trigger on deployment, CI/CD, Docker, secrets, monitoring, logging, alerting, or ops work.
---

# DevOps Engineer Skill

## Purpose

Run DapurPintar AI reliably and observably. Every environment change keeps secrets managed, telemetry correlated, backups verified, and releases reviewable.

## Responsibilities

- Maintain the deployment pipeline and environment configuration.
- Wire OpenTelemetry (OTLP) traces, metrics, and logs.
- Manage secrets via environment and secret stores, never in code.
- Define and verify backup, restore, and recovery objectives.
- Write runbooks and incident responses.

## Inputs

- `backend/.env.example` for the required environment surface.
- `internal/platform/telemetry` and `internal/platform/config` in `backend/`.
- `docs/architecture/` and `PROJECT_ROADMAP.md` for operational milestones.
- `.env.example` and secrets conventions.

## Outputs

- Reproducible environments with documented variables and defaults.
- Correlated, redacted telemetry and alert rules.
- Backup/restore verification and recovery runbooks.

## Dependencies

- PostgreSQL backup (M15), Redis as non-authoritative cache, object storage for media.
- Secret management must be decided before production deployment.

## Status

Active - aligned with M15 Deployment and Operations planning.
