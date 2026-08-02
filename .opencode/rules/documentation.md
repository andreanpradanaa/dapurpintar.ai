# Documentation Rules

> Aturan penulisan dan pemeliharaan dokumentasi untuk DapurPintar AI.

## Prinsip

- Dokumentasi ditulis untuk pembaca (engineer dan arsitek berikutnya), bukan penulis.
- Selalu perbarui dokumen saat kode berubah, di perubahan yang sama.
- Deskripsikan keadaan yang disetujui (approved state), bukan keadaan yang dicita-citakan.
- Jika dokumen bertentangan dengan decision record, decision record yang menang.

## Struktur Dokumentasi

- Konteks: `contexts/`
- Keputusan: `templates/adr.md`, Decision Register `docs/architecture/m4-decision-register.md`
- API: `docs/api/openapi.yaml`, `docs/api/`
- Database: `docs/database/`
- Backend: `docs/backend/`
- Milestone: `PROJECT_ROADMAP.md`, `docs/project/milestone-list.md`

## Standar Format

- Gunakan Markdown.
- Judul hierarki (H1 → H2 → H3).
- Sertakan contoh kode bila perlu; contoh harus cocok dengan API/layout nyata.
- Faktanya di tabel, rationale di prosa.

## Konvensi

- Referensi keputusan pakai ID register: `M4-DEC-XXX`.
- Terminologi konsisten: `expiry`, bukan `expiration`; `Recommendation`, bukan `suggestion`.
- Nama milestone cocok dengan milestone list.
- Status konsisten dengan protokol review (In Review, Approved, Complete).
- Link antar dokumen pakai path relatif-repo.

## Review

- Tandai referensi rusak dan konten basi saat review.
- Perbarui dokumen dependen di perubahan yang sama dengan perubahan sumber.
