# DapurPintar AI Public API Design

## Executive Summary

This document defines the public API contract for DapurPintar AI. It establishes the resource vocabulary, ownership boundaries, URI patterns, business operations, authentication categories, consistency expectations, error categories, and version lifecycle required for the Web Application and future approved consumers.

The API is a versioned REST/JSON interface over the Backend Modular Monolith. It exposes business capabilities, not database tables or internal component structure. The API Layer translates requests to Application use cases; the Application Layer orchestrates Domain behavior; the Domain Layer remains the authority for business rules, aggregates, policies, invariants, and domain events.

AI operations are exposed as product capabilities through DapurPintar AI. Clients never communicate directly with the AI Gateway or External AI Provider. The API returns recommendations as decision support, and the user remains the final decision-maker for meal and shopping commitments.

## API Design Principles

### REST principles

- Resources represent business concepts owned by bounded contexts.
- HTTP methods express retrieval, creation, replacement or partial change, and removal intent.
- URI patterns use nouns for resources and explicit action sub-resources when a domain command is clearer than generic CRUD.
- Responses represent business outcomes and read models, not internal aggregates, repository types, prompts, or provider payloads.
- The API is stateless from the contract perspective; each request carries the context required to authorize the business operation.
- Read operations must not change business state.

### Resource ownership

Every resource has one owning bounded context. Clients may use references to related resources, but they cannot mutate another context's resource through an unrelated URI.

- Identity and Access owns Accounts.
- User Context and Preferences owns User Profiles.
- Pantry Management owns Pantry and Pantry Items.
- Culinary Knowledge and Recipe Experience owns Recipes and Favorites.
- Meal Planning owns Meal Plans and Planned Meals.
- Shopping Optimization owns Shopping Lists and Shopping Items.
- AI-Assisted Kitchen Decision Support owns Kitchen Recommendations and Recommendation Conversations.

### Versioning strategy

Use URI versioning beginning with `/api/v1`. The version identifies the public contract, not an internal application version. Breaking changes require a new major URI version. Additive, backward-compatible changes may remain within the current version under the API compatibility policy.

### Idempotency

- `GET` operations are safe and repeatable.
- `PUT` and `DELETE` operations are designed to be idempotent when the resource semantics allow it.
- `POST` operations represent commands or creation and are not assumed to be idempotent by default.
- Creation and generation commands that may be retried by a client should support an idempotency requirement at the contract level, using a client-supplied request identity or equivalent business deduplication rule.
- Repeating an acceptance, completion, or removal command must not create contradictory business state.

Retry-sensitive `POST` commands use the `Idempotency-Key` header. The server scopes the key to the authenticated User and endpoint, retains it for at least 24 hours, and returns the original result for a matching retry. Reusing a key with a different request payload is a validation conflict.

### Pagination

Collection resources use cursor pagination with `cursor` and `limit` query parameters. The default limit is 20 and the server-enforced maximum is 100. Responses provide `page.next_cursor` and `page.has_more`. Pagination applies to recipes, pantry items, favorites, meal plans, shopping lists, recommendations, and conversations where collections can grow.

Pagination must be stable enough for a user to continue reviewing a collection without silently skipping or duplicating business items. A resource documents its default ordering and supports only documented `sort` and `order` values.

### Filtering

Filters are resource-specific and must use business vocabulary. Examples include pantry category, expiry attention, meal-plan period, shopping status, recommendation status, and recipe discovery criteria. Query collections use `sort` and `order` for documented sorting. Unknown or unsupported filters or sorts are validation errors rather than silently ignored conditions.

### Sorting

Sorting is explicit and limited to documented business fields. Examples include expiry priority for pantry items, recipe relevance or preparation time, meal date, shopping status, and recommendation creation time. A default order must be documented for every sortable collection.

### Error model

Errors have stable business categories and a safe human-readable explanation. An error response should provide a machine-readable error code, a user-safe message, relevant validation details when applicable, and a request reference for support and observability. It must not expose SQL, stack traces, internal component names, prompts, credentials, or provider payloads.

### Consistency rules

- A successful command response represents a business operation accepted by the Backend Modular Monolith and validated by the relevant Aggregate.
- PostgreSQL-backed business state is authoritative.
- Redis-backed state must not be exposed as authoritative business truth.
- Read Models may combine context for a question but must preserve source ownership.
- Cross-context actions are not implicit cascades. Accepting a Recommendation does not automatically create a Meal Plan; completing a Shopping Item does not automatically create a Pantry Item.
- AI output remains a proposal until the relevant business acceptance operation is completed.

## Authentication Boundary

The categories below describe contract access, not authentication implementation.

### Public endpoints

- `POST /api/v1/accounts` for account registration.
- `POST /api/v1/accounts/login` for account participation.
- `GET /api/v1/recipes` for public general recipe discovery.
- `GET /api/v1/recipes/{recipeId}` for public general recipe detail.

Public recipe views must not expose private preferences, pantry context, recommendations, or user activity.

### Authenticated endpoints

All endpoints involving a user's identity, profile, preferences, pantry, favorites, meal plans, shopping lists, recommendations, or conversations require an authenticated User context. This includes:

- Account participation changes and logout.
- User Profile and preference management.
- Pantry and Pantry Item operations.
- Favorites.
- Meal Plans and Planned Meals.
- Shopping Lists and Shopping Items.
- Kitchen Recommendations and Recommendation Conversations.
- AI Chat, Pantry Analysis, and personalized recommendation operations.

Authorization is evaluated against the resource owner and future household scope. A valid identity alone does not grant access to another user's business context.

### Future roles

The following role boundaries are future and must not be treated as MVP access:

- Household Member for shared household resources.
- Nutrition Professional for approved Nutrition Guidance capabilities.
- Grocery Partner for approved purchase-intent or partner capabilities.
- Commercial Operator for SaaS Commercial Operations, premium, or enterprise administration.

Future roles require explicit authorization policies and must not bypass the owning bounded context.

## Resource Catalog

| Resource | Purpose | Owner Bounded Context | Supported operations |
|---|---|---|---|
| Accounts | Establish trusted account participation | Identity and Access | Register, login, logout, view current participation, future access administration |
| User Profiles | Maintain personal cooking context | User Context and Preferences | View current profile, update profile, update preferences and constraints |
| Pantry | Represent the user's available kitchen context | Pantry Management | View pantry summary, view availability, review expiry attention |
| Pantry Items | Manage individual ingredients and their lifecycle | Pantry Management | List, add, view, adjust, remove |
| Recipes | Discover and understand cooking options | Culinary Knowledge and Recipe Experience | Search, list, view detail, view cooking guidance |
| Favorites | Express a user's preference for recipes | Culinary Knowledge and Recipe Experience | List, add, remove |
| Meal Plans | Organize daily and weekly cooking intention | Meal Planning | List, create, view, revise, cancel, complete |
| Planned Meals | Assign meals to dates or meal occasions | Meal Planning | List, plan, revise, remove, complete |
| Shopping Lists | Organize intended kitchen purchases | Shopping Optimization | List, create, generate, view, review, revise, complete |
| Shopping Items | Track individual purchase intentions | Shopping Optimization | List, add, revise, complete, remove |
| Kitchen Recommendations | Provide contextual AI-assisted cooking decisions | AI-Assisted Kitchen Decision Support | Request, list, view, present, accept, reject, supersede |
| Recommendation Conversations | Preserve recommendation-scoped AI conversation context | AI-Assisted Kitchen Decision Support | Start, view, continue, close, future independent lifecycle |

## Endpoint Catalog

The endpoint catalog defines URI patterns and business purpose only. Request and response payloads are intentionally excluded.

### Accounts

| Method | URI Pattern | Business purpose |
|---|---|---|
| POST | `/api/v1/accounts` | Register a new account. |
| POST | `/api/v1/accounts/login` | Begin authenticated account participation. |
| POST | `/api/v1/accounts/logout` | End the current account participation session. |
| GET | `/api/v1/accounts/me` | View the current account participation context. |

`POST /api/v1/accounts`, `POST /api/v1/accounts/login`, and general recipe discovery are public. The remaining account operation is authenticated.

### User Profiles

| Method | URI Pattern | Business purpose |
|---|---|---|
| GET | `/api/v1/profile` | View the current User Profile. |
| PATCH | `/api/v1/profile` | Update personal profile information. |
| PATCH | `/api/v1/profile/preferences` | Change personal cooking preferences and constraints. |

All User Profile operations are authenticated and scoped to the current User.

### Pantry and Pantry Items

| Method | URI Pattern | Business purpose |
|---|---|---|
| GET | `/api/v1/pantry` | View the current Pantry summary and availability context. |
| GET | `/api/v1/pantry/items` | List Pantry Items with documented filtering and sorting. |
| POST | `/api/v1/pantry/items` | Add a Pantry Item. |
| GET | `/api/v1/pantry/items/{itemId}` | View one Pantry Item. |
| PATCH | `/api/v1/pantry/items/{itemId}` | Adjust Pantry Item quantity, category, or expiry context. |
| DELETE | `/api/v1/pantry/items/{itemId}` | Remove a Pantry Item from available pantry context. |
| GET | `/api/v1/pantry/expiry` | Review ingredients approaching expiry. |

All Pantry operations are authenticated and owned by the current User or future authorized Household scope.

### Recipes

| Method | URI Pattern | Business purpose |
|---|---|---|
| GET | `/api/v1/recipes` | Discover and search recipes. |
| GET | `/api/v1/recipes/{recipeId}` | View recipe detail and cooking guidance. |

Recipe discovery may be public for general content. Personalized suitability remains an authenticated AI Recommendation capability.

### Favorites

| Method | URI Pattern | Business purpose |
|---|---|---|
| GET | `/api/v1/favorites` | List the current User's favorite recipes. |
| PUT | `/api/v1/favorites/recipes/{recipeId}` | Mark a Recipe as favorited in an idempotent manner. |
| DELETE | `/api/v1/favorites/recipes/{recipeId}` | Remove a Recipe from the current User's favorites. |

Favorites are authenticated and owned by the User's Recipe preference context.

### Meal Plans and Planned Meals

| Method | URI Pattern | Business purpose |
|---|---|---|
| GET | `/api/v1/meal-plans` | List daily and weekly Meal Plans. |
| POST | `/api/v1/meal-plans` | Create a Meal Plan for a defined planning period. |
| GET | `/api/v1/meal-plans/{planId}` | View one Meal Plan. |
| PATCH | `/api/v1/meal-plans/{planId}` | Revise Meal Plan intent without performing a lifecycle command. |
| POST | `/api/v1/meal-plans/{planId}/cancel` | Cancel a Meal Plan. |
| POST | `/api/v1/meal-plans/{planId}/complete` | Complete a Meal Plan. |
| GET | `/api/v1/meal-plans/{planId}/meals` | List Planned Meals in a Meal Plan. |
| POST | `/api/v1/meal-plans/{planId}/meals` | Plan a meal in a date or meal occasion. |
| PATCH | `/api/v1/meal-plans/{planId}/meals/{plannedMealId}` | Change a Planned Meal. |
| DELETE | `/api/v1/meal-plans/{planId}/meals/{plannedMealId}` | Remove a Planned Meal. |

All Meal Plan operations are authenticated. A Recommendation may guide planning, but the user must make the planning decision.

### Shopping Lists and Shopping Items

| Method | URI Pattern | Business purpose |
|---|---|---|
| GET | `/api/v1/shopping-lists` | List the current User's Shopping Lists. |
| POST | `/api/v1/shopping-lists` | Create a Shopping List. |
| POST | `/api/v1/shopping-lists/generate` | Generate purchase intent from meal, pantry, or explicit user context. |
| GET | `/api/v1/shopping-lists/{listId}` | View one Shopping List. |
| PATCH | `/api/v1/shopping-lists/{listId}` | Revise Shopping List details without performing a lifecycle command. |
| POST | `/api/v1/shopping-lists/{listId}/activate` | Activate a reviewed Shopping List. |
| POST | `/api/v1/shopping-lists/{listId}/complete` | Complete a Shopping List. |
| POST | `/api/v1/shopping-lists/{listId}/cancel` | Cancel a Shopping List. |
| POST | `/api/v1/shopping-lists/{listId}/archive` | Archive a completed or cancelled Shopping List. |
| GET | `/api/v1/shopping-lists/{listId}/items` | List Shopping Items. |
| POST | `/api/v1/shopping-lists/{listId}/items` | Add a Shopping Item. |
| PATCH | `/api/v1/shopping-lists/{listId}/items/{itemId}` | Revise Shopping Item details without changing completion state. |
| DELETE | `/api/v1/shopping-lists/{listId}/items/{itemId}` | Remove a Shopping Item. |
| POST | `/api/v1/shopping-lists/{listId}/items/{itemId}/complete` | Complete one Shopping Item. |

All Shopping operations are authenticated. Completing an item does not automatically update the Pantry.

### Kitchen Recommendations

| Method | URI Pattern | Business purpose |
|---|---|---|
| GET | `/api/v1/recommendations` | List the current User's recommendations with status and purpose filters. |
| POST | `/api/v1/recommendations` | Request a contextual Kitchen Recommendation. |
| GET | `/api/v1/recommendations/{recommendationId}` | View one Kitchen Recommendation and its options. |
| POST | `/api/v1/recommendations/{recommendationId}/present` | Mark a Recommendation as available for user consideration where presentation is a business state. |
| POST | `/api/v1/recommendations/{recommendationId}/options/{optionId}/accept` | Accept one Recommendation Option for possible later planning or shopping action. |
| POST | `/api/v1/recommendations/{recommendationId}/reject` | Reject a Recommendation for the current decision. |
| POST | `/api/v1/recommendations/{recommendationId}/supersede` | Supersede current guidance when a replacement or changed context applies. |

All Recommendation operations are authenticated and preserve user decision authority.

### Recommendation Conversations

| Method | URI Pattern | Business purpose |
|---|---|---|
| POST | `/api/v1/recommendations/{recommendationId}/conversation` | Start the one recommendation-scoped AI conversation owned by a Recommendation. |
| GET | `/api/v1/recommendations/{recommendationId}/conversation` | View the recommendation-scoped conversation context. |
| POST | `/api/v1/recommendations/{recommendationId}/conversation/messages` | Continue the recommendation-scoped conversation. |
| POST | `/api/v1/recommendations/{recommendationId}/conversation/close` | Close the conversation context. |

Conversation operations are authenticated. In the MVP, the Conversation is a child of Kitchen Recommendation and has no independent collection or lifecycle outside its owning Recommendation. It does not expose provider prompts, provider payloads, or credentials.

### AI-assisted operations

| Method | URI Pattern | Contract status | Business purpose |
|---|---|---|---|
| POST | `/api/v1/ai/pantry-analysis` | MVP | Analyze the authorized Pantry context and return decision support. |
| POST | `/api/v1/ai/meal-suggestions` | Future | Return candidate meal guidance without creating a Meal Plan. |

AI-assisted operations are authenticated and rate-limited. Future operations are not part of the MVP implementation contract until their product scope is promoted.

## Command APIs

Command APIs express business intent and may cause a Domain Event after validation and successful business handling.

### Account and profile commands

- Register Account.
- Login and Logout Account Participation.
- Change User Preferences.
- Change Cooking Constraints or Goals.

### Pantry commands

- Add Pantry Item.
- Adjust Pantry Item.
- Remove Pantry Item.
- Request Use-First Recommendation from expiry context.

### Recipe and preference commands

- Favorite Recipe.
- Remove Favorite Recipe.

### Planning commands

- Create Meal Plan.
- Plan Meal.
- Change Planned Meal.
- Remove Planned Meal.
- Complete Meal Plan when the business definition is finalized.

### Shopping commands

- Generate Shopping List.
- Add Shopping Item.
- Revise Shopping Item.
- Complete Shopping Item.
- Complete or Archive Shopping List.

### Recommendation commands

- Request Kitchen Recommendation.
- Present Recommendation.
- Accept Recommendation Option.
- Reject Recommendation.
- Supersede Recommendation.

Command APIs must preserve aggregate boundaries. A command on one resource must not silently mutate another bounded context's authoritative resource.

## Query APIs

Query APIs answer business questions and return read-oriented views.

- Account participation view.
- User Profile and preference view.
- Pantry summary, Pantry Item list, and expiry view.
- Recipe discovery and detail views.
- Favorite Recipe list.
- Meal Plan and Planned Meal views.
- Shopping List and Shopping Item views.
- Recommendation list, detail, and presentation view.
- Recommendation Conversation view.
- Future Home Dashboard and household views.

Query APIs support documented pagination, filtering, and sorting. A query view may combine information from multiple contexts, but source ownership remains unchanged.

## AI APIs

AI APIs expose product capabilities rather than prompt or provider details.

### AI Chat

`POST /api/v1/recommendations/{recommendationId}/conversation` starts a recommendation-scoped conversation. `POST /api/v1/recommendations/{recommendationId}/conversation/messages` continues it. The API returns product-level assistance and conversation state without exposing prompts or provider payloads.

### Pantry Analysis

`POST /api/v1/ai/pantry-analysis` requests analysis of the authenticated User's authorized pantry context. The result is decision support and may identify use-first opportunities, but Pantry Management remains authoritative for availability and expiry.

### Recommendation

`POST /api/v1/recommendations` requests a Kitchen Recommendation using approved User Profile, Pantry, Recipe, Meal Plan, and Shopping context. The result follows the Kitchen Recommendation lifecycle. A specific Recommendation Option must be accepted before it becomes a basis for a commitment.

### Meal Suggestion

`POST /api/v1/ai/meal-suggestions` is a future AI-assisted planning surface for candidate daily or weekly meal guidance. It must not create a Meal Plan without an explicit planning command and is not an MVP dependency unless the product scope is promoted.

All AI APIs are authenticated, rate-limited, privacy-scoped, and subject to AI safety and provider-failure handling.

## Error Strategy

### Response conventions

Successful responses use a consistent envelope:

```json
{
  "data": {},
  "page": null,
  "request_id": "req_01..."
}
```

Collection responses populate `data` with the collection and `page` with `next_cursor` and `has_more`. Non-collection responses set `page` to `null`.

Error responses use this stable shape:

```json
{
  "error": {
    "code": "RECOMMENDATION_OPTION_NOT_ACCEPTABLE",
    "message": "The recommendation option cannot be accepted in its current state.",
    "details": [],
    "request_id": "req_01..."
  }
}
```

`code` is stable and machine-readable. `message` is safe for the user. `details` is present for field or query validation and never contains secrets or internal implementation details. The `request_id` is returned in the response body and `X-Request-ID` header.

### HTTP status categories

- **2xx:** The requested read or business operation completed successfully.
- **400:** The request is malformed or violates contract-level validation.
- **401:** The request lacks valid authenticated participation where required.
- **403:** The identity is known but not authorized for the resource or business scope.
- **404:** The requested resource or business view is not available within the authorized scope.
- **409:** The requested operation conflicts with current aggregate state or a business invariant.
- **422:** The request is understandable but violates a domain validation or business rule.
- **429:** The caller has exceeded a documented usage or abuse limit.
- **5xx:** DapurPintar AI cannot complete the operation because of an internal or dependency failure.

### Business errors

Business errors should name the violated business meaning, such as invalid Meal Plan slot, unavailable Recommendation state, conflicting Shopping Item status, or Pantry ownership mismatch. They must not expose implementation details.

### Validation errors

Validation errors identify the business input or query condition that must be corrected. Unknown filters, unsupported sorts, invalid resource identifiers, and invalid state transitions are validation or business errors according to the contract definition.

### Authorization errors

Authorization errors do not reveal whether another user's private resource exists. The API must preserve User and future Household scope without leaking ownership information.

### AI provider failures

AI provider failures are represented as safe dependency failures. The API must not expose provider payloads, credentials, prompts, or internal retry behavior. Where possible, core non-AI operations remain usable and the response explains that AI assistance could not be completed.

## Versioning Strategy

The public URI begins with `/api/v1`.

- Additive response fields, new optional filters, and new resources may be introduced compatibly within `v1`.
- Removing or renaming a resource, changing ownership semantics, changing required behavior, or changing error meaning requires a new major URI version.
- Deprecated resources remain documented during a transition period and include a replacement direction.
- Future `/v2` or later versions must not silently change the meaning of existing user commitments or aggregate ownership.

## API Lifecycle

### Stable

An API is Stable when its resource ownership, business purpose, URI pattern, authorization behavior, and error categories are approved for MVP or public use.

MVP Stable resources include Accounts, User Profiles, Pantry, Pantry Items, Recipes, Favorites, Meal Plans, Planned Meals, Shopping Lists, Shopping Items, Kitchen Recommendations, and Recommendation Conversations.

### Deprecated

An API is Deprecated when it remains available for compatibility but should not be used for new consumers. Deprecation requires a documented replacement, migration period, and preserved ownership semantics.

### Future

An API is Future when its business capability is supported by the roadmap but not part of the MVP contract. Future surfaces include household sharing, nutrition, notifications, grocery integration, subscription, commercial operations, advanced analytics, and AI meal suggestions.

## Security Considerations

- API ownership is enforced by the Backend Modular Monolith and the owning bounded context.
- Authentication establishes identity; authorization establishes access to the User or future Household scope.
- User Profiles, Pantry, Favorites, Meal Plans, Shopping Lists, Recommendations, and Conversations are private by default.
- Account and AI-intensive operations are rate-limited according to documented business protection policies.
- All externally supplied identifiers, commands, filters, and AI-related user input are validated at the API boundary and again by business rules.
- API responses must minimize personal, household, pantry, conversation, and recommendation data.
- AI APIs must defend against prompt injection, unsafe generated content, unsupported claims, data leakage, and cost exhaustion.
- The API never exposes provider credentials, prompts, raw provider payloads, database details, or internal component failures.
- Future Household Member, Nutrition Professional, Grocery Partner, and Commercial Operator access requires explicit role and scope policies.

## Risks

- Resource names may drift from bounded-context language and create ownership ambiguity.
- Action endpoints may become hidden workflow engines instead of explicit business commands.
- Clients may infer automatic cross-context mutations that the domain intentionally forbids.
- Public recipe discovery may accidentally expose personalized or private context.
- AI responses may be treated as authoritative facts instead of decision support.
- API retries may duplicate recommendation, shopping-list, or other business commands without a clear idempotency policy.
- Read models may become de facto write models or expose stale/conflicting context.
- Future roles and partner APIs may be introduced before authorization and ownership rules are validated.
- Error details may leak private-resource existence, prompts, provider information, or internal architecture.
- Contract versioning may lag behind business-language changes discovered through MVP validation.

## Assumptions

- `/api/v1` is the initial public API contract version.
- The Web Application is the primary MVP consumer.
- The Backend Modular Monolith owns API orchestration and business rules.
- The Domain Layer owns aggregate invariants and resource meaning.
- PostgreSQL-backed state is authoritative; Redis is not a source of truth.
- AI clients communicate through DapurPintar AI and never directly with the AI Gateway or External AI Provider.
- Meal Plan, Shopping List, and Recommendation commitments require explicit business actions.
- Future roles and resources remain outside the MVP contract until product and domain policy are approved.
- No payload schemas are defined by this document; they belong to a later API specification deliverable.

## Diagram Reference

The API flow is maintained in `docs/architecture/diagrams/api-flow.mmd`.
