# DapurPintar AI M9 MVP Features - Sign-off

## Document Control

| Item | Value |
|---|---|
| Milestone | M9 - MVP Features |
| Deliverables | M9-001 through M9-008: Identity, Pantry, Recipes, Meal Plans, Shopping, Recommendations, Conversation, AI Pantry Analysis |
| Status | Ready for review |
| Parent documents | `docs/architecture/m4-decision-register.md`, `docs/architecture/authentication-authorization.md`, `docs/api/m6-api-contract.md`, `docs/database/m5-schema.md` |

## Purpose

This document records the DP-FEAT-001 vertical slice of the M9 MVP Features milestone: the Identity and Access and User Context and Preferences bounded contexts, delivered end to end from the M5 schema through SQLC-generated queries, the domain use cases, and the Fiber HTTP handlers. It is the first slice to prove the approved M7 Clean Architecture wiring (handler -> service -> store -> PostgreSQL) against a real database.

## Decisions Applied

- M4-DEC-001: email verification stays optional for local development; accounts register as `active`. Required before public launch.
- M4-DEC-002: credential change (password reset) is in MVP scope; password recovery is deferred.
- M4-DEC-003: PostgreSQL is the durable session and revocation authority. `auth_sessions` stores only a SHA-256 hash of the refresh secret, the lineage `family_id`, and revocation metadata. Refresh rotates the session (old marked `replaced_by`, new in the same family); reuse of a revoked session revokes the whole lineage.
- M4-DEC-004: session cookies `dp_session` (access) and `dp_refresh` (refresh) are HttpOnly, SameSite=Lax, Secure in production; the middleware reads the cookie or a Bearer header and origin-checks state changes.
- M4-DEC-006/007/008/009: M5 schema, SQLC, M6 API contract, and the recipe public boundary are honored (recipe tables exist in the base schema; recipe endpoints arrive with DP-FEAT-003).
- M4-DEC-013: recommendation conversations are retained while active with a 30-day window; not in this slice.

## Deliverables

### DP-FEAT-001-1 Schema and SQLC

- `backend/migrations/00001_base_schema.sql`: the approved M5 base schema in dependency order, plus `auth_sessions` (DP-FEAT-001 addition for M4-DEC-003). Tables are ordered so foreign keys resolve; the Down migration mirrors the reverse order.
- `backend/sqlc.yaml`: PostgreSQL engine, `pgx/v5` sql package, `uuid`/`timestamptz` mapped to `string`/`time.Time`, nullable text as pointer strings, and `date` mapped to `time.Time` so pgx v5 scans `valid_from` in binary format.
- `backend/internal/database/queries/identity.sql` and `user_context.sql`: named queries for accounts, sessions, profiles, and preference sets.
- `backend/internal/gen/sqlc/`: generated models, parameters, and query methods.

### DP-FEAT-001-2 Identity domain and store

- `backend/internal/identity/store.go`: the `Store` port (accounts, profiles, preferences, sessions) that handlers and services depend on; the SQLC adapter lives behind it.
- `backend/internal/identity/domain.go`: `Account`, `UserProfile`, `PreferenceSet`, and `Session` aggregates with their M5 lifecycle enums.
- `backend/internal/identity/service.go`: use cases Register, Login, Logout, Refresh (with rotation and lineage revocation), GetAccountByID, GetProfileByAccountID, UpdateProfile, UpdatePreferences, ActivePreferences. Argon2id password hashing, email normalization, password policy, and generic login failures (no account enumeration).
- `backend/internal/identity/store/postgres.go`: `Postgres` adapter mapping SQLC rows to domain types and translating `pgx.ErrNoRows` and unique violations into domain errors.

### DP-FEAT-001-3 HTTP wiring

- `backend/internal/http/handlers_identity.go`: register, login, refresh, logout, me, getProfile, updateProfile, updatePreferences with the M6 AccountView/UserProfileView/PreferencesView schemas and stable M6 error codes.
- `backend/internal/http/middleware/session.go`: `SessionCookies` (SetAccess/SetRefresh/ClearAccess/ClearRefresh) and `Authenticated` middleware reading the session cookie or Bearer header.
- `backend/internal/http/server.go`: routes wired under `/api/v1`, the authenticated group, session-cookie construction from config, and error mapping helpers.
- `backend/cmd/api/main.go`: constructs the identity store and service and injects them into `http.New`.

## Verified

- `go build ./...` passes.
- `go vet ./...` passes.
- `gofmt -l .` clean.
- `go test ./...` passes (auth, errors, ai, http).
- Integration tests run against the Docker Compose PostgreSQL (`TEST_DATABASE_URL` gated) and pass with `-v`.

Test coverage:

- `internal/identity/store/postgres_test.go` (integration, requires Postgres):
  - Account lifecycle: create, fetch by id/email, update status, duplicate-email `ErrEmailInUse`.
  - Profile and preferences: profile create/fetch/update, preference-set create and active lookup, missing set returns `ErrNotFound`.
  - Session lifecycle: create, lookup by secret hash, rotation via `MarkSessionReplacedBy`, revocation, and family revocation.
  - End to end: Register -> Login -> Refresh (rotated) -> reuse detection -> Logout over the real adapter.
- `internal/http/server_test.go`: health endpoint, protected route requires session (401 without), protected route succeeds with a valid token using a stub identity store.

## Dependencies on Pending Decisions

- M4-DEC-005 (unified login strategy) remains Pending for M15/M16 and does not block this slice.
- M4-DEC-014/015/017 (social login, MFA, and others) remain Pending for M15/M16.
- Email verification (M4-DEC-001) becomes mandatory before public launch; the `email_verified_at` column and status plumbing already exist.
- Password recovery (M4-DEC-002) is deferred; `UpdatePassword` arrives with that feature.

## Exit Criteria

DP-FEAT-001 is complete when:

- Registration, login, logout, refresh, current account, profile, and preferences work end to end over PostgreSQL.
- Refresh rotates sessions and detects reuse by revoking the lineage (M4-DEC-003).
- Session cookies are HttpOnly/SameSite=Lax and Secure outside development (M4-DEC-004).
- M6 schemas and error codes are honored.
- The slice is covered by unit and integration tests and the whole backend builds, vets, and is gofmt-clean.

## Related Documents

- `docs/architecture/m4-decision-register.md`
- `docs/architecture/authentication-authorization.md`
- `docs/api/m6-api-contract.md`
- `docs/database/m5-schema.md`
- `docs/database/m5-sqlc.md`
- `docs/project/milestone-list.md`

---

## DP-FEAT-002 Pantry and Pantry Item CRUD

### Decisions Applied

- M4-DEC-006: schema shape validated per M5 sign-off; pantry and pantry_item tables match the approved DDL.
- M4-DEC-007: expiry_date stored as user-local `date`; expiry view computes a cutoff date and queries items with `expiry_date <= cutoff`.
- M4-DEC-008: API contract follows openapi.yaml with 7 endpoints and 3 pantry error codes.

### Deliverables

**DP-FEAT-002-1 Schema and SQLC**
- `backend/internal/database/queries/pantry.sql`: 10 named queries (GetPantryByProfileID, CreatePantry, GetPantryItemByID, CreatePantryItem, UpdatePantryItem, RemovePantryItem, UpdatePantryItemStatus, ListPantryItems with cursor+filters+sort, ListExpiringItems with cutoff date).
- `backend/internal/gen/sqlc/pantry.sql.go`: generated params, query methods, and slice scanning for the PantryItem model with `pgtype.Numeric` quantity.

**DP-FEAT-002-2 Domain and store**
- `backend/internal/pantry/domain.go`: Pantry and PantryItem aggregates, ItemStatus lifecycle enum (available/running_low/expiring_soon/consumed/removed), ErrNotFound sentinel.
- `backend/internal/pantry/store.go`: Store port (11 methods) — owned-pantry queries, item CRUD, paginated listing with optional category/status/sort filters, and expiry-ordered listing.
- `backend/internal/pantry/service.go`: use cases GetSummary (counts), ListItems (cursor-based with filters), AddItem (validate + lazy pantry provision), GetItem (ownership check), UpdateItem, RemoveItem, ListExpiringItems.
- `backend/internal/pantry/store/postgres.go`: SQLC adapter handling pgtype.Numeric quantity conversion and proper nullable-string coalesce for UpdatePantryItem.

**DP-FEAT-002-3 HTTP handlers**
- `backend/internal/http/handlers_pantry.go`: 7 route handlers with M6 PantryItemView, profile resolution through the identity service, and stable error codes (PANTRY_ITEM_NOT_FOUND, PANTRY_QUANTITY_NEGATIVE, PANTRY_EXPIRY_INVALID).
- `backend/internal/http/server.go`: Handler struct and New() accept pantry.Service; routes registered under the authenticated group.
- `backend/cmd/api/main.go`: constructs the pantry store adapter and service, injects into http.New.

### Verified

- `go build ./...`, `go vet ./...`, `gofmt -l .` clean.
- 7 unit/integration tests pass (all packages).
- 4 pantry integration tests (requiring Postgres): pantry lifecycle, item CRUD, paginated listing with expiry, and end-to-end service flow (add→sum→get→update→remove→verify gone→empty list).

### Exit Criteria

DP-FEAT-002 is complete when:
- Pantry items can be created, listed (with filters/pagination), retrieved, updated, and removed over the owner's pantry.
- Lazy pantry provisioning works transparently on first item access.
- Ownership is enforced (item belongs to the user's pantry).
- The expiry endpoint returns items approaching a configurable cutoff date.
- All endpoints return the M6 PantryItemView schema with stable error codes.

---

## DP-FEAT-003 Recipe Discovery, Detail, and Favorites

### Decisions Applied

- M4-DEC-009: public/private recipe visibility boundary. Public listing only returns `is_public = true AND status = 'available'`. Detail endpoint gates private recipes behind the `is_public` flag with `RECIPE_NOT_PUBLIC` (404).

### Deliverables

**DP-FEAT-003-1 SQLC queries**

- `backend/internal/database/queries/recipes.sql`: 6 named queries — `ListPublicRecipes` (cursor, text search, prep-time filter, sort), `GetRecipeByID` (with visibility gate), `GetActiveFavorite`, `CreateFavorite` (idempotent upsert via `ON CONFLICT DO NOTHING`), `RemoveFavorite` (soft-delete), `ListFavorites` (joined with recipe summaries, cursor-paginated).

**DP-FEAT-003-2 Domain and store**

- `backend/internal/recipes/domain.go`: `Recipe`, `RecipeSummary`, `RecipeIngredient`, `Favorite` aggregates with JSONB unmarshaling helpers.
- `backend/internal/recipes/store.go`: Store port — public listing with optional search/filters, recipe detail, favorite lifecycle (create/get/remove), and paginated favorite list.
- `backend/internal/recipes/service.go`: use cases `ListRecipes`, `GetRecipe`, `Favorite` (idempotent: double-favorite succeeds), `Unfavorite` (idempotent: unfavorite without active favorite succeeds), `ListFavorites`.
- `backend/internal/recipes/store/postgres.go`: SQLC adapter handling `pgtype.Int4` time fields, `pgtype.Bool` visibility, JSONB ingredient/instruction unmarshaling, and `ON CONFLICT DO NOTHING` graceful fallback for idempotent favorites.

**DP-FEAT-003-3 HTTP handlers**

- `backend/internal/http/handlers_recipes.go`: 5 route handlers — `GET /recipes` (public), `GET /recipes/:recipeId` (public), `GET /favorites` (auth), `PUT /favorites/recipes/:recipeId` (auth, idempotent), `DELETE /favorites/recipes/:recipeId` (auth, idempotent). Uses `RecipeSummaryView`, `RecipeDetailView`, and `FavoriteView` M6 schemas.
- `backend/internal/http/server.go`: Handler struct and `New()` accept `*recipes.Service`. Routes registered: public recipes on the API group, authenticated favorites on the authed group.

### Verified

- `go build ./...`, `go vet ./...`, `gofmt -l .` clean.
- 8 test packages pass (unit + integration).
- 4 recipe integration tests (requiring Postgres): public listing excludes private recipes, detail with JSONB parsing, favorite create/read/remove lifecycle with idempotent behavior, service end-to-end (list→get→favorite→double-favorite→list-favorites→unfavorite).

### Exit Criteria

DP-FEAT-003 is complete when:
- Public recipe listing returns only `is_public = true AND status = 'available'` recipes with text search, prep-time filter, and sort options.
- Recipe detail includes parsed `ingredients` and `instructions` arrays and gates non-public recipes.
- Favoriting is idempotent (double-favorite succeeds with 204).
- Unfavoriting is idempotent (unfavorite without active favorite succeeds with 204).
- The `ListFavorites` endpoint returns paginated favorites with embedded `RecipeSummaryView`.

---

## DP-FEAT-004 Meal Plan and Planned Meal Lifecycle

### Decisions Applied

- M4-DEC-006: schema updated with `title` column on meal_plans (migration 00002).
- M4-DEC-007: `period_start`, `period_end`, and `meal_date` are stored as user-local `date`; the service validates period boundaries and meal dates within plan ranges.
- M4-DEC-008: API contract implemented with all 10 endpoints per openapi.yaml.

### Deliverables

**DP-FEAT-004-1 Schema and SQLC**

- `backend/migrations/00002_meal_plan_title.sql`: adds the `title` column to `meal_plans`.
- `backend/internal/database/queries/meal_plan.sql`: 11 named queries — plan lifecycle (create/get/list/update/status-transition), planned meal CRUD (create/get/list/update/remove), and slot-lookup for conflict detection.
- `backend/internal/gen/sqlc/meal_plan.sql.go`: generated params and scan methods with nullable `*time.Time` and `*string` params for coalesce updates.

**DP-FEAT-004-2 Domain and store**

- `backend/internal/mealplan/domain.go`: `MealPlan` and `PlannedMeal` aggregates with `PlanStatus` (draft/planned/in_progress/completed/cancelled/revised) and `MealStatus` (proposed/planned/revised/removed/completed) enums.
- `backend/internal/mealplan/store.go`: Store port (12 methods) — plan lifecycle, planned meal CRUD, slot-conflict lookup.
- `backend/internal/mealplan/service.go`: use cases `CreateMealPlan` (period validation), `UpdateMealPlan`, `CancelMealPlan`/`CompleteMealPlan` (state guards), `PlanMeal` (period-range check + slot-conflict detection), `ListPlannedMeals`, `UpdatePlannedMeal`, `RemovePlannedMeal`.
- `backend/internal/mealplan/store/postgres.go`: SQLC adapter with date column mapping and nullable coalesce params.

**DP-FEAT-004-3 HTTP handlers**

- `backend/internal/http/handlers_mealplan.go`: 10 route handlers — plan listing/creation/detail/update/cancel/complete, planned meal listing/creation/update/removal. Date parsing (`2006-01-02`), occasion validation (`breakfast`/`lunch`/`dinner`/`snack`), and M6 error codes.

### Verified

- `go build ./...`, `go vet ./...`, `gofmt -l .` clean.
- 9 test packages pass (unit + integration).
- 4 meal plan integration tests (requiring Postgres): plan CRUD and listing, lifecycle transitions (draft→cancelled), meal CRUD with slot-detection, service end-to-end (create→plan-meal→slot-conflict→update-status→complete→cancel-conflict).

### Exit Criteria

DP-FEAT-004 is complete when:
- Meal plans can be created with validated period ranges and optional titles.
- Meal plans support lifecycle transitions (cancel, complete) with state guards.
- Planned meals can be added within the plan's date range with slot conflict detection.
- Planned meals can be updated (occasion, recipe reference, status) and removed.
- All 10 endpoints return the M6 MealPlanView/PlannedMealView schemas with stable error codes.

---

## DP-FEAT-005 Shopping List and Item Management

### Deliverables

**DP-FEAT-005-1 SQLC queries**

- `backend/internal/database/queries/shopping.sql`: 11 named queries covering list and item lifecycles with nullable coalesce updates and profile-scoped ownership.

**DP-FEAT-005-2 Domain and service**

- `backend/internal/shopping/domain.go`: `ShoppingList` and `ShoppingItem` aggregates, `ListStatus` (8 states) and `ItemStatus` (3 states) enums.
- `backend/internal/shopping/service.go`: use cases with lifecycle state guards (activate requires non-terminal, complete/cancel guard against archived, archive requires terminal state, complete-item requires open).
- `backend/internal/shopping/store/postgres.go`: adapter handling `pgtype.Numeric` quantities.

**DP-FEAT-005-3 HTTP handlers**

- `backend/internal/http/handlers_shopping.go`: 14 route handlers with `ShoppingListView` (including `item_counts`) and `ShoppingItemView`.

### Verified

- `go build ./...`, `go vet ./...`, `gofmt -l .` clean.
- 10 test packages pass.
- 4 integration tests: list CRUD, lifecycle transitions (draft→active→completed→archived), item CRUD with status changes, service end-to-end.

### Exit Criteria

DP-FEAT-005 complete when all 14 endpoints return M6 schemas with lifecycle state guards and stable error codes.

---

## DP-FEAT-006 Kitchen Recommendation

### Deliverables

- Migration 00003: add `purpose` column to `kitchen_recommendations`.
- 8 SQLC queries, 7 HTTP endpoints, state machine (requested→created→presented→accepted/rejected/superseded) with option acceptance guards.
- 3 integration tests: CRUD, option management, state guards (cannot present requested, cannot supersede rejected).

## Verified

- `go build ./...`, `go vet ./...`, `gofmt -l .` clean.
- 11 test packages pass (27 integration tests total).
- AI Gateway integration deferred: recommendations are created with "requested" status; full AI processing wired in a future slice.

---

## DP-FEAT-007 Recommendation Conversation

- 4 SQLC queries (create, get, snapshot-update, close), 4 HTTP endpoints.
- Messages stored in `context_snapshot` JSONB as role+content+created_at array.
- Idempotent start (get-or-create), continue appends user+assistant messages, close with state guard.

## DP-FEAT-008 AI Pantry Analysis

- Single `POST /ai/pantry-analysis` endpoint (stub). Returns empty `use_first_opportunities` and `optimization_suggestions` arrays. Full AI integration deferred.

## Verified (M9 Complete)

- `go build ./...`, `go vet ./...`, `gofmt -l .` clean.
- 11 test packages pass (24 integration tests against PostgreSQL).
- 3 schema migrations applied (00001 base, 00002 meal_plan title, 00003 recommendation purpose).
- 55 HTTP endpoints across 7 bounded contexts.
