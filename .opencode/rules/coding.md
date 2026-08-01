# Coding Rules

> Standar penulisan kode yang wajib diikuti.

## Bahasa Pemrograman

- ...

## Gaya Kode (Code Style)

- Gunakan formatter bawaan bahasa (gofmt, Prettier, dll.)
- Ikuti konvensi penamaan proyek yang sudah ada

## Struktur File

- Satu tanggung jawab per file/modul
- Ikuti pola direktori existing (`internal/`, `app/`, dst.)

## Testing

- Setiap perubahan fitur harus menyertakan test
- Jalankan test sebelum menyelesaikan tugas
- Command test: ...

## Keamanan

- Jangan pernah menulis secret/key ke kode
- Validasi input dari user
- Ikuti prinsip least privilege

## Anti-Patterns

- Copy-paste berlebihan → refactor
- Komentar berlebihan untuk kode yang sudah jelas
- Hardcoded value tanpa konfigurasi
