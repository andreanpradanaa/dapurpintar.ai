# ADR-009 Use OpenTelemetry for Observability

- Status: Accepted
- Date: 2026-08-02

## Context

The system depends on a web client, Fiber API, application and domain workflows, PostgreSQL, Redis, and OpenAI. Product and engineering targets include API latency, database query time, API and AI availability, AI response time, and AI error rate. AI-provider and database failures must be diagnosable without exposing private user data.

## Decision

Use OpenTelemetry as the standard instrumentation and telemetry model. Instrumentation will cover the request path and important business and dependency operations, with structured logs, metrics, and traces exported to the operational stack and visualized through Grafana. Telemetry must apply redaction, safe attributes, and appropriate sampling.

## Consequences

- Traces can connect API requests to application work and database, cache, and AI-provider dependencies.
- Standardized telemetry supports consistent dashboards, alerts, and future backend changes.
- Product metrics such as AI usage, meal plans, pantry items, and shopping lists can be correlated with system health.
- Instrumentation and cardinality require discipline to control cost and noise.
- Privacy controls are mandatory because telemetry can accidentally contain personal data, prompts, tokens, or pantry details.

## Alternatives Considered

- **Vendor-specific instrumentation:** May offer deep integration, but increases vendor lock-in and makes future platform changes harder.
- **Logs only:** Simple to start, but insufficient for causal request tracing and reliable latency or availability measurement.
- **Separate libraries per signal:** Possible, but creates inconsistent propagation and operational conventions across the system.
