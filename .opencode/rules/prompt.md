# Prompt Rules

> Aturan cara menulis dan menggunakan prompt untuk AI assistant di DapurPintar AI.

## Prinsip Prompt

- Tujuan harus jelas dan spesifik.
- Sertakan konteks yang relevan dari dokumen approved (bukan dari asumsi).
- Batasi ruang lingkup agar fokus pada satu milestone/deliverable.
- Selalu rujuk decision register untuk keputusan yang terkunci.

## Struktur Prompt

- **Tugas:** apa yang harus dilakukan.
- **Konteks:** informasi pendukung dari dokumen/artefak.
- **Batasan:** constraint, teknologi, standar proyek.
- **Output:** format yang diharapkan.

## Prompt per Peran

| Peran | File Prompt |
|-------|-------------|
| Product Manager | `prompts/product/` |
| Backend | `prompts/backend/` |
| Frontend | `prompts/frontend/` |
| Architecture | `prompts/architecture/` |
| Business | `prompts/business/` |
| AI | `prompts/ai/` |
| DevOps | `prompts/devops/` |
| Database | `prompts/database/` |

## Checklist

- [ ] Tujuan jelas?
- [ ] Konteks cukup (dari dokumen approved)?
- [ ] Batasan disebutkan?
- [ ] Output terdefinisi?
- [ ] Decision register yang relevan dirujuk?
- [ ] Lingkup sesuai milestone?

## Anti-Patterns

- Meminta perubahan tanpa menyebut dokumen yang mengunci keputusan.
- Meminta implementasi fitur yang masih diblokir decision pending.
- Meminta AI menulis dokumen yang bertentangan dengan ADR yang disetujui.
