---
document_id: M1-008
title: Feature Inventory
owner: Product Manager
status: Draft
version: 1.0.0
last_updated: 2026-08-02
related_documents:
  - product-scope.md
  - mvp-definition.md
  - product-roadmap.md
---

# Feature Inventory

## Overview

### Purpose

Dokumen ini berisi daftar seluruh fitur DapurPintar AI yang akan menjadi acuan pengembangan backend, frontend, database, AI, testing, dan roadmap produk.

Setiap fitur memiliki identitas unik (Feature ID) sehingga mudah dilacak pada GitHub Issues, sprint, dokumentasi API, database, dan implementasi kode.

---

# Feature Status Legend

| Status | Description |
|---------|-------------|
| Planned | Belum dikerjakan |
| In Progress | Sedang dikembangkan |
| Completed | Sudah selesai |
| Deferred | Ditunda |
| Cancelled | Dibatalkan |

---

# Priority Legend

| Priority | Description |
|----------|-------------|
| P0 | Wajib untuk MVP |
| P1 | Setelah MVP |
| P2 | Nice to Have |
| P3 | Future Vision |

---

# Complexity Legend

| Complexity | Description |
|------------|-------------|
| Low | ≤ 2 hari |
| Medium | 3–5 hari |
| High | > 5 hari |
| Epic | Perlu dipecah menjadi beberapa fitur |

---

# Feature Inventory

| Feature ID | Module | Feature | Priority | MVP | AI | Complexity | Business Value | User Value | Status |
|------------|--------|---------|----------|-----|----|------------|----------------|------------|--------|
| AUTH-001 | Authentication | User Registration | P0 | ✅ | ❌ | Medium | High | High | Planned |
| AUTH-002 | Authentication | User Login | P0 | ✅ | ❌ | Medium | High | High | Planned |
| AUTH-003 | Authentication | JWT Authentication | P0 | ✅ | ❌ | Medium | High | High | Planned |
| USER-001 | User | User Profile | P0 | ✅ | ❌ | Low | High | High | Planned |
| USER-002 | User | User Preferences | P0 | ✅ | ❌ | Low | High | High | Planned |
| FAMILY-001 | Family | Family Workspace | P1 | ❌ | ❌ | Medium | High | High | Planned |
| PANTRY-001 | Pantry | Add Pantry Item | P0 | ✅ | ❌ | Medium | High | High | Planned |
| PANTRY-002 | Pantry | Update Pantry Item | P0 | ✅ | ❌ | Medium | High | High | Planned |
| PANTRY-003 | Pantry | Delete Pantry Item | P0 | ✅ | ❌ | Low | Medium | Medium | Planned |
| PANTRY-004 | Pantry | Pantry Categories | P0 | ✅ | ❌ | Low | Medium | High | Planned |
| PANTRY-005 | Pantry | Expiration Tracking | P0 | ✅ | ❌ | Medium | High | High | Planned |
| PANTRY-006 | Pantry | Pantry Dashboard | P1 | ❌ | ❌ | Medium | High | High | Planned |
| RECIPE-001 | Recipe | Recipe Search | P0 | ✅ | ❌ | Medium | High | High | Planned |
| RECIPE-002 | Recipe | Recipe Detail | P0 | ✅ | ❌ | Low | High | High | Planned |
| RECIPE-003 | Recipe | Favorite Recipes | P0 | ✅ | ❌ | Low | Medium | High | Planned |
| AI-001 | AI | AI Recipe Recommendation | P0 | ✅ | ✅ | High | High | High | Planned |
| AI-002 | AI | Pantry Analysis | P0 | ✅ | ✅ | High | High | High | Planned |
| AI-003 | AI | Ingredient Replacement | P1 | ❌ | ✅ | High | High | High | Planned |
| AI-004 | AI | Leftover Recommendation | P1 | ❌ | ✅ | High | High | High | Planned |
| AI-005 | AI | AI Chat Assistant | P0 | ✅ | ✅ | High | High | High | Planned |
| MEAL-001 | Meal Planner | Daily Meal Plan | P0 | ✅ | ❌ | Medium | High | High | Planned |
| MEAL-002 | Meal Planner | Weekly Meal Plan | P0 | ✅ | ❌ | Medium | High | High | Planned |
| MEAL-003 | Meal Planner | Auto Meal Planning | P1 | ❌ | ✅ | High | High | High | Planned |
| SHOP-001 | Shopping | Shopping List | P0 | ✅ | ❌ | Medium | High | High | Planned |
| SHOP-002 | Shopping | Auto Shopping List | P0 | ✅ | ✅ | High | High | High | Planned |
| SHOP-003 | Shopping | Shared Shopping List | P1 | ❌ | ❌ | Medium | High | High | Planned |
| NUTRI-001 | Nutrition | Nutrition Summary | P1 | ❌ | ❌ | Medium | High | High | Planned |
| NUTRI-002 | Nutrition | Daily Calories | P1 | ❌ | ❌ | Medium | Medium | High | Planned |
| NUTRI-003 | Nutrition | Nutrition Goals | P1 | ❌ | ✅ | High | High | High | Planned |
| NOTIF-001 | Notification | Expiration Reminder | P1 | ❌ | ❌ | Medium | High | High | Planned |
| NOTIF-002 | Notification | Meal Reminder | P2 | ❌ | ❌ | Low | Medium | Medium | Planned |
| OCR-001 | OCR | Receipt OCR | P2 | ❌ | ✅ | Epic | Medium | High | Planned |
| OCR-002 | OCR | Pantry Image Recognition | P2 | ❌ | ✅ | Epic | High | High | Planned |
| REPORT-001 | Analytics | Pantry Analytics | P2 | ❌ | ❌ | Medium | Medium | Medium | Planned |
| REPORT-002 | Analytics | Spending Analytics | P2 | ❌ | ❌ | Medium | High | High | Planned |
| ADMIN-001 | Admin | User Management | P2 | ❌ | ❌ | Medium | Medium | Low | Planned |
| ADMIN-002 | Admin | AI Prompt Management | P2 | ❌ | ❌ | High | High | Low | Planned |

---

# Module Summary

| Module | Total Features | MVP |
|---------|---------------:|----:|
| Authentication | 3 | 3 |
| User | 2 | 2 |
| Family | 1 | 0 |
| Pantry | 6 | 5 |
| Recipe | 3 | 3 |
| AI | 5 | 3 |
| Meal Planner | 3 | 2 |
| Shopping | 3 | 2 |
| Nutrition | 3 | 0 |
| Notification | 2 | 0 |
| OCR | 2 | 0 |
| Analytics | 2 | 0 |
| Admin | 2 | 0 |

---

# MVP Feature Summary

Modul yang termasuk MVP:

- Authentication
- User Management
- Pantry Management
- Recipe
- AI Recommendation
- AI Chat
- Meal Planner
- Shopping List

---

# Future Modules

Fitur yang akan dikembangkan setelah MVP:

- OCR
- Computer Vision
- Smart Kitchen
- Grocery Integration
- Public API
- AI Nutrition Coach
- Budget Planner
- Enterprise Dashboard

---

# Acceptance Criteria

Dokumen dianggap selesai apabila:

- Seluruh fitur memiliki Feature ID.
- Tidak ada fitur duplikat.
- Setiap fitur memiliki prioritas.
- Setiap fitur memiliki status.
- Setiap fitur memiliki kompleksitas.
- Fitur MVP telah ditandai dengan jelas.
- Seluruh fitur sesuai Product Scope.

---

# Related Documents

- Product Scope
- MVP Definition
- Product Roadmap
- Success Metrics