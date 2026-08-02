# DevOps Engineer Examples

## Configuration surface from `.env.example`

```dotenv
DATABASE_URL=postgres://...            # primary PostgreSQL DSN
REDIS_ADDR=localhost:6379              # cache only, non-authoritative
ACCESS_TOKEN_TTL=15m                   # short-lived JWT access tokens
SESSION_COOKIE_NAME=dp_session         # session cookie name
OTLP_EXPORTER_ENDPOINT=                # optional OTLP/HTTP endpoint
AI_GATEWAY_MODEL=                      # model profile, M4-DEC-010
```

## Redacted telemetry

```go
attrs := []attribute.KeyValue{
    attribute.String("http.route", route),
    attribute.Bool("ai.fallback", fallback),
    // NEVER attribute.String("user.pantry", rawPantryJSON)
}
```

## Safe migration ordering

```text
1. Apply forward-compatible migration.
2. Deploy new application version.
3. Backfill / verify.
4. Deprecate old code path (never destructive in one step).
```
