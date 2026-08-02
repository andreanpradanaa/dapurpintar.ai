# Repository Rules

> Aturan terkait repository: struktur, git, dan manajemen kode untuk DapurPintar AI.

## Struktur Repository

- `docs/` — arsitektur, API, database, backend, project.
- `backend/` — aplikasi Go (module `github.com/andreanpradanaa/dapurpintar.ai/backend`).
- `prompts/` di `.opencode/` — prompt per peran.
- `.opencode/` — konfigurasi opencode, skill, rules, contexts, templates.
- Root: `PROJECT_ROADMAP.md`, `PROJECT_CONTEXT.md` (.opencode), `.env.example`, `Makefile`.

## Git Workflow

- Branch naming: `feature/`, `fix/`, `chore/`, `docs/`.
- Commit per unit kerja yang fokus dan ter-review; jangan commit perubahan yang tidak diminta.
- Inspect `git status`, `git diff`, `git log` sebelum commit.
- Jangan commit secret atau file env nyata.

## Commit Message

- Ikuti gaya konvensi yang sudah ada di repo.
- Ringkas, fokus pada apa dan mengapa.
- Sertakan referensi milestone/deliverable jika relevan (contoh: M7).

## Versioning

- Migrations diberi nomor berurutan (Goose `00001_...`).
- Prompt, policy, dan skema structured-output diberi versi.
- Artefak kontrak (OpenAPI) di-versioning mengikuti rilis API.

## Code Ownership

- Setiap perubahan dimiliki oleh satu bounded context.
- Jangan mengubah data/behavior context lain secara implisit.

## Merge & Review

- Perubahan di-review sebelum dianggap selesai.
- Milestone di-review (In Review) dan baru disetujui dengan sign-off eksplisit.
- Follow-up review menjadi issue, bukan amendemen diam-diam.
