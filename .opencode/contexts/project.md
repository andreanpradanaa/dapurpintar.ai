# Project Context

> Ringkasan proyek dapurpintar.ai, tujuan, dan ruang lingkup. Otoritas: `PROJECT_ROADMAP.md`, `docs/project/milestone-list.md`, `docs/architecture/m4-decision-register.md`.

## Ringkasan

- **Nama Proyek:** dapurpintar.ai
- **Visi:** Membantu Sarah dan Daniel mengurangi pemborosan makanan dan membuat keputusan memasak lebih mudah, dengan AI sebagai decision support yang aman dan dapat dipercaya.

## Tujuan

- Sistem rekomendasi makanan berbasis pantry yang akurat dan aman.
- Kontrak API, skema database, dan arsitektur yang terdokumentasi penuh.
- Backend Go yang bersih (Clean Architecture + DDD) dengan kualitas teruji.
- Migrasi tanpa keputusan tersembunyi: setiap milestone diblokir oleh decision yang eksplisit.

## Ruang Lingkup

- In-scope: pantry management, recommendation (AI decision support), meal planning, preferensi, authentication/authorization, observability.
- Out-of-scope (MVP): multi-household sharing, AI auto-purchase/auto-mutate, integrasi pihak ketiga lain di luar OpenAI.

## Pemangku Kepentingan

| Peran | Nama / Tim |
|-------|-----------|
| Produk | User (Software Architect) |
| Teknik | User (Software Architect) |
| Bisnis | User (Software Architect) |

## Metrik Keberhasilan

- Rekomendasi akurat dan aman sesuai rubrik evaluasi AI (M4-DEC-012).
- API sesuai kontrak OpenAPI; error memakai kode stabil M6.
- Semua test lulus: `go build`, `go vet`, `gofmt`, `go test -race ./...`.
- Milestone ter-review dan disetujui dengan sign-off eksplisit.

## Status Saat Ini

- M0, M1: Complete.
- M2–M7: In Review (M7 Backend Foundation).
- Current milestone: M8 AI Foundation.
- Deliverable saat ini: M4-005 M8 Blocking Decision Records (In Review).
- Berikutnya: DP-AI-001..003, diblokir oleh M4-DEC-010/011/012/016 (`docs/architecture/m8-blocking-decisions.md`).
