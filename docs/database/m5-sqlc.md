# DapurPintar AI M5 Database Design - SQLC Query Contract

## Document Control

| Item | Value |
|---|---|
| Milestone | M5 - Database Design |
| Deliverable | M5-003 SQLC Query Contract |
| Status | Draft - Awaiting Cross-Functional Review |
| Parent documents | `docs/database/m5-schema.md`, `docs/architecture/ADR-008-use-sqlc-for-database-access.md` |
| Scope | Query boundaries, naming, generated-code expectations, and ownership-preserving queries |

## Purpose

This document defines how SQLC is used to access the M5-001 schema without violating bounded-context ownership. It establishes query boundaries, naming conventions, soft-delete defaults, timezone handling, and the generated-code contract that backend repositories in M7 will consume.

## SQLC Baseline

- **Tool:** SQLC for type-safe SQL generation from reviewed SQL.
- **Config:** `sqlc.yaml` at repository root, generating code under `internal/gen/sqlc` (or the equivalent location approved in M7).
- **Queries location:** SQL query files are colocated or organized by bounded context under a queries directory (e.g., `internal/database/queries/`), with one file or section per owning context.
- **Engine:** PostgreSQL.
- **Generated code:** Type-safe structs and query methods generated at build or CI time; generated code is not hand-edited.

## Query Boundaries

Queries follow the owning bounded context. A query that reads or writes another context's authoritative table is prohibited unless it is a documented cross-context reference or a read-model aggregation approved in M6.

- Identity and Access owns `accounts` queries.
- User Context and Preferences owns `user_profiles` and `preference_sets` queries.
- Pantry Management owns `pantries` and `pantry_items` queries.
- Culinary Knowledge and Recipe Experience owns `recipes` and `recipe_favorites` queries.
- Meal Planning owns `meal_plans` and `planned_meals` queries.
- Shopping Optimization owns `shopping_lists` and `shopping_items` queries.
- AI-Assisted Kitchen Decision Support owns `kitchen_recommendations`, `recommendation_options`, and `recommendation_conversations` queries.

Cross-context joins are limited to reads that preserve source ownership (for example, a Recommendation detail view joining referenced recipe identity). A cross-context read must never become a write path into another context's table.

## Naming Conventions

- Query names describe the business operation, not the SQL mechanism. Example: `GetPantryItemByID`, `ListActivePantryItemsForProfile`, `ListRecipesPublic`, `CreateMealPlan`, `CompleteShoppingItem`, `MarkRecommendationAccepted`.
- File naming follows the owning context. Example: `pantry_queries.sql`, `meal_plan_queries.sql`.
- Parameters and result columns use snake_case in SQL and generated Go fields in camelCase by SQLC convention.
- Every query file begins with a comment block identifying the owning bounded context and the business purpose.

## Soft-Delete Defaults

- Queries that return business views must filter `deleted_at is null` by default.
- Soft-deleted rows are never returned as active; the filter is part of the query, not applied in application code after retrieval.
- Administrative or audit queries may include deleted rows explicitly and must be named to indicate that (e.g., `ListAccountsIncludingDeleted`).

## Timezone Handling in Queries

- Date-bounded queries receive the user timezone as a parameter and compute the local date boundary before execution.
- `meal_date` and `expiry_date` comparisons are evaluated against the user's local date boundary.
- Queries must not fall back to server-UTC defaults for user-visible date logic.

## Query Catalog by Context

The catalog below lists required queries by bounded context. The exact SQL is written during M7 implementation and reviewed against this contract. This is the authoritative list of query boundaries.

### Identity and Access

- `GetAccountByID`
- `GetAccountByEmail`
- `CreateAccount`
- `UpdateAccountStatus`
- `MarkAccountEmailVerified`
- `ListAccountsIncludingDeleted` (administrative only)

### User Context and Preferences

- `GetUserProfileByAccountID`
- `GetUserProfileByID`
- `CreateUserProfile`
- `UpdateUserProfile`
- `GetActivePreferenceSetForProfile`
- `CreatePreferenceSet`
- `RetirePreferenceSet`

### Pantry Management

- `GetPantryByProfileID`
- `GetPantryByID`
- `CreatePantry`
- `ListActivePantryItemsForProfile`
- `GetPantryItemByID`
- `CreatePantryItem`
- `UpdatePantryItem`
- `RemovePantryItem`
- `ListExpiringSoonForProfile`
- `UpdatePantryItemStatus`

### Culinary Knowledge and Recipe Experience

- `GetRecipeByID`
- `ListRecipesPublic` (with pagination and filters)
- `SearchRecipes`
- `ListActiveFavoritesForProfile`
- `AddRecipeFavorite`
- `RemoveRecipeFavorite`

### Meal Planning

- `ListMealPlansForProfile`
- `GetMealPlanByID`
- `CreateMealPlan`
- `UpdateMealPlan`
- `CancelMealPlan`
- `CompleteMealPlan`
- `ListPlannedMealsForPlan`
- `GetPlannedMealByID`
- `PlanMeal`
- `UpdatePlannedMeal`
- `RemovePlannedMeal`

### Shopping Optimization

- `ListShoppingListsForProfile`
- `GetShoppingListByID`
- `CreateShoppingList`
- `UpdateShoppingList`
- `GenerateShoppingList`
- `ActivateShoppingList`
- `CompleteShoppingList`
- `CancelShoppingList`
- `ArchiveShoppingList`
- `ListShoppingItemsForList`
- `GetShoppingItemByID`
- `AddShoppingItem`
- `UpdateShoppingItem`
- `RemoveShoppingItem`
- `CompleteShoppingItem`

### AI-Assisted Kitchen Decision Support

- `ListRecommendationsForProfile`
- `GetRecommendationByID`
- `CreateRecommendation`
- `UpdateRecommendationStatus`
- `PresentRecommendation`
- `AcceptRecommendation`
- `RejectRecommendation`
- `SupersedeRecommendation`
- `ListRecommendationOptions`
- `GetRecommendationOptionByID`
- `CreateRecommendationOption`
- `UpdateRecommendationOptionStatus`
- `GetConversationForRecommendation`
- `CreateConversation`
- `UpdateConversationStatus`
- `ListConversationsPastRetention` (retention cleanup)

## Generated Code Expectations

- Generated query methods return typed structs or slices; no `map[string]any` for business queries.
- Transaction handling follows the aggregate boundary: a repository method that mutates one aggregate runs in a single transaction; cross-aggregate workflows are coordinated in the Application Layer, not inside a generated query.
- Query methods used by repositories must not leak SQL primitives into the domain layer; the repository interface in the domain language maps generated results to domain types.
- Generated code is regenerated whenever a query or schema changes; CI verifies generated output matches committed output.

## Integrity and Review

- Every query is reviewed for ownership preservation and soft-delete correctness.
- Date-bounded queries are reviewed for timezone parameterization.
- Cross-context queries require explicit approval and must be named as read-model aggregations.
- SQLC config and generated output are part of code review.

## Exit Criteria

M5-003 is complete when:

- Query boundaries match the bounded contexts.
- Naming and file conventions are defined.
- Soft-delete and timezone defaults are explicit in the query contract.
- The query catalog covers the M5-001 schema for MVP features.
- Generated-code expectations are defined and reviewable.

## Related Documents

- `docs/database/m5-schema.md`
- `docs/architecture/adr/ADR-008-use-sqlc-for-database-access.md`
- `docs/architecture/tactical-ddd.md`
- `docs/architecture/api-design.md`
- `docs/architecture/m4-m5-blocking-decisions.md`
