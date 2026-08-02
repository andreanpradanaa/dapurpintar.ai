# Engineering Context

> Standar teknis, tooling, dan cara kerja pengembangan DapurPintar AI. Otoritas: `docs/architecture/` dan `AI_RULES.md`.

## Tech Stack

- **Backend:** Go (1.26), Fiber v3, pgxpool, SQLC, Goose, JWT (HS256), Argon2id, OpenTelemetry, slog.
- **Frontend:** Next.js, M3 design system (mengikuti milestone M10).
- **Database:** PostgreSQL (system of record), Redis (cache non-authoritative).
- **AI:** AI Gateway -> OpenAI, structured output, prompt/policy versioning (M8).

## Struktur Repository

```
├── .opencode/          # Config opencode, skill, rules, contexts, prompts, templates
├── backend/
│   ├── cmd/api/main.go
│   ├── internal/{http, auth, config, platform, <contexts>}
│   └── migrations/     # Goose migrations
├── docs/
│   ├── architecture/
│   ├── api/
│   ├── database/
│   ├── backend/
│   └── project/
├── PROJECT_ROADMAP.md
├── Makefile
└── .env.example
```

## Tooling

| Alat | Tujuan |
|------|--------|
| `go build ./...` | Compile backend |
| `go vet ./...` | Static check backend |
| `gofmt -l .` | Format check backend |
| `go test -race ./...` | Unit/integration test backend |
| `make run/build/test/lint/fmt` | Backend shortcuts |
| SQLC / Goose | Query generation / migrasi |
| OpenTelemetry | Observability |

## Perintah Penting (Backend)

- Build: `make build`
- Test: `make test`
- Lint/format: `make lint`, `make fmt`
- Dev: `make run`
- Modul path: `github.com/andreanpradanaa/dapurpintar.ai/backend` (import pakai prefix ini).

## Konvensi

- Clean Architecture lapisan domain/application/handler/adapter terpisah.
- M6 envelope response + kode error stabil (`internal/platform/errors`).
- Soft-delete difilter di SQL (`deleted_at IS NULL`).
- UTC disimpan; interpretasi di timezone pengguna (default `Asia/Jakarta`).
- Jangan pakai GORM/ORM; AI hanya lewat AI Gateway.
