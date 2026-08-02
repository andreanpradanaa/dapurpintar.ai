# Architecture Context

> Arsitektur sistem DapurPintar AI secara keseluruhan. Otoritas penuh: `docs/architecture/`.

## High-Level Architecture

- Modular monolith dengan Clean Architecture dan DDD bounded contexts.
- Satu API REST: `/api/v1` (OpenAPI: `docs/api/openapi.yaml`).
- PostgreSQL sebagai system of record; Redis sebagai infrastruktur pendukung (non-authoritative).
- AI Gateway sebagai satu-satunya pintu ke OpenAI; output AI adalah decision support.
- Otorisasi dan ownership selalu disimpulkan server-side.

## Komponen

| Komponen | Teknologi | Tanggung Jawab |
|----------|-----------|----------------|
| Backend API | Go / Fiber | Application scaffolding, auth, HTTP contract |
| Persistence | PostgreSQL / pgx | System of record, SQLC queries, Goose migrations |
| Cache | Redis | Supporting infra, non-authoritative |
| AI | AI Gateway -> OpenAI | Decision support, structured output |
| Frontend | Next.js | UI berbasis M3 design system |
| Observability | OpenTelemetry (OTLP) | Traces, metrics, logs |

## Bounded Contexts

Unit ownership data + perilaku. Lihat `docs/architecture/bounded-context.md`.

- **Identity & Profile** — akun, profil, preferensi pengguna.
- **Pantry Management** — item pantry, kategori, expirasi.
- **Recipe Management** — resep, komposisi bahan.
- **Meal Planning** — rencana makan, penerimaan rekomendasi.
- **Recommendation** — saran AI berbasis konteks pengguna.
- **Purchase & Waste** — (perlu konfirmasi) perencanaan belanja dan pemantauan pemborosan.

Aturan: tidak ada model global bersama; hanya referensi identitas bisnis antar context.

## Data Flow

1. Client (Next.js) -> API `/api/v1` (Fiber).
2. Middleware auth menetapkan identitas; Application Layer mengotorisasi.
3. Use cases memanggil domain behavior + repositories.
4. Repositories diimplementasikan SQLC di atas PostgreSQL.
5. AI dipanggil hanya lewat AI Gateway; hasil divalidasi structured-output.
6. Response mengembalikan envelope M6.

## Keputusan Arsitektural (ADR)

- Lihat `docs/architecture/m4-decision-register.md` (M4-DEC-*) dan `docs/architecture/adr/`.
- Template: `.opencode/templates/adr.md`.

## Pola & Prinsip

- Clean Architecture, DDD (tactical + strategic), GCU (Grub/Clean/Usable).
- SQLC bukan ORM; Goose untuk migrasi; migrasi non-destruktif.
- UTC disimpan, diinterpretasi di timezone pengguna (default `Asia/Jakarta`).

## Runbook / Infrastruktur

- Deployment: `docs/architecture/deployment-architecture.md` (M15).
- Observability: `docs/architecture/observability-architecture.md`.
