# Product Context

> Konteks produk DapurPintar AI: persona, masalah, dan alur pengguna. Otoritas: `docs/architecture/domain-discovery.md`, `docs/project/project-assumptions.md`.

## Persona Pengguna

- **Sarah** — ibu rumah tangga di Jakarta, ingin mengurangi pemborosan makanan, mengelola pantry, dan dapat saran menu yang memakai bahan yang tersedia sebelum expirasi.
- **Daniel** — profesional muda yang sibuk, ingin memasak lebih sering, memanfaatkan bahan yang ada, dan mengurangi keputusan "apa yang harus dimasak".

## Masalah yang Dipecahkan

- Bahan makanan terbuang karena lupa expirasi dan tidak ada rencana memakai.
- Kebingungan memutuskan menu dari bahan yang tersedia.
- Kurangnya dukungan yang akurat dan aman untuk menyarankan resep berbasis stok pantry.

## Value Proposition

- Rekomendasi makanan yang akurat dan dapat ditindaklanjuti, berbasis pantry yang selalu aktual, dengan AI sebagai decision support yang tidak pernah diam-diam mengubah data pengguna.

## Alur Pengguna (User Journey)

1. Pengguna menambah/memperbarui item pantry.
2. Sistem melacak expirasi dan stok.
3. Pengguna meminta rekomendasi menu.
4. AI mengusulkan opsi (berbasis konteks + keamanan); pengguna menerima/menolak.
5. Rencana makan dikonfirmasi; AI tidak pernah auto-commit.

## Fitur Utama

- Manajemen pantry (CRUD, kategori, expirasi).
- Rekomendasi berbasis AI (decision support).
- Meal planning (konfirmasi pengguna).
- Preferensi pengguna (alergi, diet, timezone).
- Pemantauan pemborosan/belanja (per konfirmasi).

## Non-Fitur (Constraint)

- AI tidak pernah otomatis membeli/menambah/mengubah pantry tanpa konfirmasi.
- Tidak ada model global bersama antar bounded context.
- Fokus MVP: Sarah dan Daniel, single-profile.
