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
- M2–M8: Approved (reviewed and approved 2026-08-02).
- Current milestone: M9 MVP Features (DP-FEAT-001 Identity and Access in progress).
- Berikutnya: DP-FEAT-002 pantry, then DP-FEAT-003 recipe discovery.
