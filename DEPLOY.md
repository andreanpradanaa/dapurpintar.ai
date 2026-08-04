# Deploy Guide — Dapur Pintar (Gratis)

Stack: **Vercel** (Next.js) + **Fly.io** (Go/Fiber) + **Supabase** (Postgres).

## 0. Prerequisites

- [x] `flyctl` — `brew install flyctl`
- Akun Supabase, Vercel, Fly.io (semua gratis).
- Repo ter-push ke GitHub: `andreanpradanaa/dapurpintar.ai`.

## 1. Supabase — database Postgres

Via dashboard (paling gampang untuk one-time setup):

1. https://supabase.com → New Project.
2. Name: `dapur-pintar`, password: generate kuat (simpan!), region: **Singapore**.
3. Tunggu ~2 menit. Buka **Project Settings → Database → Connection string**.
4. Pilih tab **Session pooler** (port `5432`, host `*.pooler.supabase.com`).
5. Salin URL, append `?sslmode=require` di akhir → simpan sebagai `DATABASE_URL`.

Format akhir:
```
postgresql://postgres.<ref>:<password>@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres?sslmode=require
```

> Catatan: pakai **Session pooler** (5432) BUKAN Transaction pooler (6543) — pgx
> pakai prepared statements yang bentrok dengan transaction mode.

## 2. Fly.io — backend

```bash
# Dari root repo
cd backend

# Login (buka browser)
fly auth login

# Hapus app lama jika pernah coba
# fly apps destroy dapur-pintar-backend

# Launch (gunakan fly.toml yang sudah ada, jangan generate ulang)
fly launch --no-deploy --dockerfile Dockerfile --name dapur-pintar-backend --region sin

# Set secrets — ganti nilai di dalam { ... }
fly secrets set DATABASE_URL="postgresql://postgres.<ref>:<pass>@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres?sslmode=require"
fly secrets set OPENAI_API_KEY="sk-or-v1-..."           # OpenRouter free $5 credit
fly secrets set OPENAI_MODEL="anthropic/claude-3.5-sonnet"
fly secrets set OPENAI_BASE_URL="https://openrouter.ai/api/v1"
fly secrets set OPENAI_TIMEOUT="60s"
fly secrets set CORS_ORIGINS="*"                         # tighten setelah frontend live

# Deploy!
fly deploy

# Cek
fly status
fly logs
curl https://dapur-pintar-backend.fly.dev/api/v1/health
# 期待: {"status":"ok","db":"up","llm":"openai","recipes":32}
```

Konfigurasi `fly.toml`:
- `auto_stop_machines=true`, `min_machines_running=0` → full gratis, cold start ±15s.
- Size `shared-cpu-1x` 256MB — termasuk free allowance.
- Region `sin` (Singapore) — terdekat Indonesia.

## 3. Vercel — frontend

1. https://vercel.com/new → Import repo `andreanpradanaa/dapurpintar.ai`.
2. **Root Directory**: pilih `frontend/`.
3. Framework preset: Next.js (auto-detect).
4. **Environment Variables**:
   - `BACKEND_URL` = `https://dapur-pintar-backend.fly.dev`
5. Deploy → dapat domain `dapur-pintar-ai-xxx.vercel.app`.

## 4. Tighten CORS

Setelah frontend live:

```bash
cd backend
fly secrets set CORS_ORIGINS="https://dapur-pintar-ai-xxx.vercel.app"
fly deploy   # atau fly machines restart <id>
```

## 5. Verifikasi

1. `curl https://dapur-pintar-backend.fly.dev/api/v1/health` → `recipes:32`.
2. Buka Vercel domain → generate recipe → lihat `/api/v1/recipes/generate` ter-proxy ke Fly.
3. `fly logs` → "migrations applied" + "recipe library ready".

## Catatan biaya

- **Vercel Hobby**: free forever (100GB bandwidth, 100h build/bulan).
- **Fly.io free allowance** (per Org):
  - 3 shared-cpu 1x 256MB VM.
  - 3GB outbound bandwidth.
  - Asalkan app idle paling banyak waktu → 99% gratis.
  - Perlu tambah payment method untuk verifikasi (tidak akan charged selama di allowance).
- **Supabase free tier**: 500MB storage + 2 projetos, pausable 1 minggu idle -> resume sekali klik.
- **LLM**: OpenRouter beri $5 free credit, cukup ±500-2000 request untuk `gpt-4o-mini` / `claude-3.5-sonnet`.

## Troubleshooting

- **`migrations dir not found`** di log Fly: pastikan `Dockerfile` menyalin `migrations/` dan `data/` (sudah ada di `backend/Dockerfile`).
- **`OPENAI_API_KEY is required`**: `fly secrets list` untuk konfirmasi secret ter-set.
- **502 di Vercel**: cek `fly status`, kemungkinan mesin tidur → request pertama sengaja lambat. Tunggu 15-20s lalu refresh.
- **Supabase connection refused**: pakai Session pooler (5432) dan tambah `?sslmode=require`.
- **DB not seeding**: cek `fly logs` — seed hanya jalan saat `count==0`, kalau sudah ada data dari testing sebelumnya, hapus dulu di Supabase table editor.