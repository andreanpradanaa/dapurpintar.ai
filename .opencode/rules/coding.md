# Coding Rules

> Standar penulisan kode yang wajib diikuti untuk DapurPintar AI.

## Bahasa Pemrograman

- Backend: Go (lihat `rules/backend.md` dan skill `backend-go`).
- Frontend: Next.js/TypeScript (lihat `rules/frontend.md` dan skill `frontend-nextjs`).
- Database: SQL (PostgreSQL) via SQLC dan Goose.

## Gaya Kode (Code Style)

- Gunakan formatter bawaan bahasa: `gofmt` untuk Go, Prettier/ESLint untuk frontend.
- Ikuti konvensi penamaan yang sudah ada di proyek.
- Tanpa komentar yang berlebihan; komentar hanya untuk alasan non-sepele.

## Struktur File

- Satu tanggung jawab per file/modul.
- Ikuti pola direktori existing (`backend/internal/`, dst.).
- Clean Architecture lapisan: domain/application/handler/adapter dipisah jelas.

## Testing

- Setiap perubahan fitur menyertakan test.
- Jalankan test sebelum menyelesaikan tugas.
- Backend: `go test -race ./...`.
- Frontend: perintah test project (komponen/interaksi/a11y).

## Keamanan

- Jangan pernah menulis secret/key ke kode.
- Validasi input user di boundary API dan lagi di business rules.
- Ikuti prinsip least privilege; otorisasi selalu disimpulkan server-side.

## Anti-Patterns

- Copy-paste berlebihan → refactor.
- Komentar berlebihan untuk kode yang sudah jelas.
- Hardcoded value tanpa konfigurasi.
- ORM/general-purpose O/R mapper di backend → gunakan SQLC.
- AI output yang langsung menjadi commitment bisnis tanpa konfirmasi user.
