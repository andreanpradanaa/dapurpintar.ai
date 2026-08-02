# DevOps Engineer Rules

## Secrets

- Secrets are never committed. `.env.example` holds keys only, with documented meaning and defaults.
- Use a secret store (env-injected or dedicated manager) in every non-local environment.
- Rotation is a documented process, not a fire drill.

## Configuration

- Configuration is explicit and validated at startup by `internal/platform/config`.
- Sensitive values are never logged; telemetry attributes are redacted before export.
- Timezone, TTLs, and provider keys have non-secret defaults documented in `.env.example`.

## Observability

- Traces, metrics, and logs share correlation attributes.
- AI calls, database queries, and HTTP requests are instrumented end-to-end.
- Alerts cover availability, error rate, latency, and AI quota/cost.

## Reliability

- PostgreSQL is the system of record; backups are verified by restore tests, not assumed.
- Redis holds only cache; a Redis outage must not lose business data.
- Recovery objectives (RPO/RTO) are explicit and tested.

## Delivery

- Every release is reviewable, versioned, and deployable independently.
- Migrations are applied before release of depending code; no destructive steps.
- Rollback path exists for application releases.

## Compliance and hygiene

- Dependencies are pinned and scanned.
- Logs and metrics respect privacy rules; personal kitchen context is redacted.
