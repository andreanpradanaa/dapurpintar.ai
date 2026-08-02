# DapurPintar AI — Remaining Work

Last updated: 2026-08-02. Status after M9 backend + M10 frontend complete.

## Testing

| Item | Priority | Notes |
|---|---|---|
| Frontend unit/integration tests | Medium | Jest + React Testing Library for components and hooks (AuthProvider, API client, views) |
| Frontend E2E tests | Medium | Playwright or Cypress for critical user journeys (register → add pantry → plan meal → shop) |
| Backend service unit tests | Medium | Currently only `identity/service_test.go` exists. Add for pantry, recipes, mealplan, shopping, recommendation services |
| Backend error path tests | Low | Test pgx errors, duplicate handling, constraint violations across all adapters |

## New Features (M11-M12)

| Item | Priority | Notes |
|---|---|---|
| Barcode scanner | Low | Requires `zxing-js/library` or native camera API. Scan product barcodes for quick pantry adds |
| OCR receipt scanner | Low | Requires Tesseract.js or cloud OCR. Parse shopping receipts to add multiple pantry items |
| Advanced meal planning AI | Planned | AI-generated weekly meal plans based on pantry contents and preferences. Requires AI provider configured |
| Ingredient replacement suggestions | Planned | AI-suggested substitutes when pantry is missing an ingredient |
| Recipe seed expansion | Low | Add more Indonesian recipes (20-30 total) covering more regions and dietary preferences |

## Polish

| Item | Priority | Notes |
|---|---|---|
| Page transition animations | Low | `framer-motion` for smooth page-to-page transitions |
| Shopping list generate from meal plan | Low | Auto-populate shopping list from planned meals' ingredients |
| Dark mode | Low | `next-themes` for light/dark toggle |
| PWA support | Low | Service worker + manifest for installable app |
| Email verification | Required pre-launch | M4-DEC-001: currently accounts are active immediately |

## Dependencies

| Item | Status | Notes |
|---|---|---|
| AI Gateway (OpenAI) | Optional, wired | Works when `AI_PROVIDER=openai` and `AI_API_KEY` are set. Falls back gracefully |
| Redis cache | Optional, not wired | Schema supports it but no caching logic implemented yet |
| OTel observability | Configured, disabled | `OTLP_DISABLE=true` by default. Enable for production monitoring |
