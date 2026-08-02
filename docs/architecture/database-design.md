# DapurPintar AI Database Design

## Executive Summary

This document defines the conceptual and logical database model for DapurPintar AI. It translates the approved bounded contexts and tactical DDD model into durable business ownership and relationships without prescribing SQL, migrations, queries, repository code, or PostgreSQL-specific implementation.

PostgreSQL is the authoritative system of record for durable business data. The model follows aggregate ownership: each aggregate root owns the lifecycle and consistency of its child entities, while references between bounded contexts preserve ownership boundaries. Redis remains a supporting store for cache, session management, rate limiting, and temporary state only.

The MVP model supports account participation, user profile and preferences, pantry management, recipes and favorites, daily and weekly meal plans, shopping lists, kitchen recommendations, and recommendation-scoped AI conversation context. Future household collaboration, nutrition, grocery partnerships, notifications, subscriptions, ingredient catalog capabilities, and advanced analytics can extend the model without becoming MVP dependencies.

## Design Principles

### System of Record

PostgreSQL is the source of truth for durable business facts owned by DapurPintar AI. A business decision must be recoverable from authoritative data owned by the relevant bounded context.

Redis is never authoritative. A cache miss, session-state loss, rate-limit reset, or temporary-state loss must not remove or redefine a Pantry Item, Recipe, Meal Plan, Shopping Item, Recommendation, Account, or User Profile.

### Aggregate Ownership

Each Aggregate Root owns the meaning, lifecycle, and invariants of its child entities:

- Account owns account participation and access state.
- User Profile owns personal preferences and cooking constraints.
- Pantry owns Pantry Items, quantity, and expiry context.
- Recipe owns recipe meaning and Recipe Favorites.
- Meal Plan owns Planned Meals and planning validity.
- Shopping List owns Shopping Items and shopping status.
- Kitchen Recommendation owns Recommendation Options, recommendation state, and MVP conversation context.

An aggregate may refer to another aggregate's business identity, but it must not persist another aggregate's child entities as if it owned them.

### Bounded Context Ownership

Persistence follows bounded-context authority:

- Identity and Access owns Account.
- User Context and Preferences owns User Profile and Preference Set.
- Pantry Management owns Pantry and Pantry Item.
- Culinary Knowledge and Recipe Experience owns Recipe and Recipe Favorite.
- Meal Planning owns Meal Plan and Planned Meal.
- Shopping Optimization owns Shopping List and Shopping Item.
- AI-Assisted Kitchen Decision Support owns Kitchen Recommendation, Recommendation Option, and Recommendation Conversation.

The database model must not create a shared global meaning for Ingredient, Recipe, Meal, Recommendation, or Preference. Context-specific references preserve the consuming context's language.

### Transaction Boundaries

An aggregate is the default boundary for immediate business consistency. Changes that must be valid together belong to one aggregate. Cross-aggregate workflows are coordinated by the Application Layer and preserve the independent consistency of each aggregate.

For example, accepting a Recommendation may guide a Meal Plan, but it does not make the Meal Plan part of the Recommendation aggregate. Completing a Shopping Item does not automatically make a Pantry Item available.

### Normalization Philosophy

The logical model favors clear ownership, minimal duplication of authoritative facts, and normalized relationships between business concepts. Repeated information may be represented as a context-specific snapshot when historical meaning or decision reproducibility requires it, but snapshots must be clearly distinguished from current source-of-truth data.

Read-oriented views may combine and reshape business information for a specific question. They must not become a second source of truth or replace aggregate ownership.

### Future Extensibility

Future additions should extend the model through new bounded contexts or aggregates rather than expanding existing aggregates without business justification. Candidate extensions include Household Collaboration, Nutrition Guidance, Ingredient Catalog, Notifications and Reminders, SaaS Commercial Operations, grocery partner relationships, and Recommendation Feedback.

## Database Strategy

### PostgreSQL as authoritative storage

PostgreSQL stores durable business information for the Backend Modular Monolith. It supports relationships and transaction boundaries required by Account, User Profile, Pantry, Recipe, Meal Plan, Shopping List, and Kitchen Recommendation ownership.

The database design does not define schema syntax, indexes, implementation constraints, or migrations. Those decisions belong to later database implementation work and must preserve the ownership described here.

### Redis is not a source of truth

Redis may support:

- Safe, appropriately scoped cache information.
- Short-lived session management state.
- Rate limiting and abuse controls.
- Temporary coordination or bounded transient state.

Redis must not own durable business facts, aggregate lifecycle, event history, or authoritative read models.

### Aggregate persistence

Each Aggregate Root has a durable representation that owns its child-entity lifecycle. Child entities are persisted as part of the aggregate's business ownership, while cross-context references point to the identity or business meaning supplied by another context.

The Backend Modular Monolith is the only logical application boundary that manages these durable business records. Read Models and AI assistance consume approved data through the domain/application ownership boundaries.

### Ownership boundaries

Persistence responsibility must follow the owning bounded context. A context may use another context's published business meaning, but it must not directly change another context's authoritative state. This prevents Pantry, Meal Planning, Shopping, and Recommendation data from becoming a coupled shared model.

## Entity Overview

The following entities are the conceptual MVP model. The list describes business identity and ownership, not columns or schema structures.

### Account

- **Purpose:** Establish trusted identity and participation in the product.
- **Owner Bounded Context:** Identity and Access.
- **Aggregate Ownership:** Aggregate Root.
- **Lifecycle:** Pending, Active, Restricted, Closed.
- **Relationships:** Has one User Profile; provides trusted scope for personal aggregates.

### User Profile

- **Purpose:** Maintain personal cooking context, preferences, constraints, and goals.
- **Owner Bounded Context:** User Context and Preferences.
- **Aggregate Ownership:** Aggregate Root.
- **Lifecycle:** Created, Incomplete, Ready, Updated.
- **Relationships:** Belongs to one Account; contextualizes Pantry, Meal Plan, Shopping List, and Kitchen Recommendation decisions.

### Preference Set

- **Purpose:** Represent a coherent set of personal cooking preferences that can change over time.
- **Owner Bounded Context:** User Context and Preferences.
- **Aggregate Ownership:** Child Entity of User Profile.
- **Lifecycle:** Declared, Active, Revised, Retired.
- **Relationships:** Belongs to one User Profile; may be used by Kitchen Recommendation without transferring ownership.

### Pantry

- **Purpose:** Represent the ingredients a user considers available in the kitchen.
- **Owner Bounded Context:** Pantry Management.
- **Aggregate Ownership:** Aggregate Root.
- **Lifecycle:** Active, Archived when the owning personal or future household context closes.
- **Relationships:** Belongs to one User Profile or future Household scope; owns Pantry Items; informs Recommendations and Shopping Lists.

### Pantry Item

- **Purpose:** Represent one ingredient position with quantity, category, and expiry context.
- **Owner Bounded Context:** Pantry Management.
- **Aggregate Ownership:** Child Entity of Pantry.
- **Lifecycle:** Available, Running Low, Expiring Soon, Consumed, Removed.
- **Relationships:** Belongs to one Pantry; may inform Recipe suitability, Recommendations, and Shopping Need Assessment without being owned by them.

### Recipe

- **Purpose:** Represent a usable cooking option with recipe meaning and cooking guidance.
- **Owner Bounded Context:** Culinary Knowledge and Recipe Experience.
- **Aggregate Ownership:** Aggregate Root.
- **Lifecycle:** Available, Revised, Retired. Exact catalog lifecycle remains a business decision for detailed design.
- **Relationships:** May have many Recipe Favorites; may be referenced by Recommendation Options and Planned Meals.

### Recipe Favorite

- **Purpose:** Represent a user's expressed preference for a Recipe.
- **Owner Bounded Context:** Culinary Knowledge and Recipe Experience.
- **Aggregate Ownership:** Child Entity of Recipe, with a relationship to User Profile.
- **Lifecycle:** Active, Removed.
- **Relationships:** Refers to one Recipe and one User Profile; may provide a preference signal to Recommendation.

### Meal Plan

- **Purpose:** Represent a daily or weekly future cooking intention.
- **Owner Bounded Context:** Meal Planning.
- **Aggregate Ownership:** Aggregate Root.
- **Lifecycle:** Draft, Planned, In Progress, Completed, Cancelled, Revised.
- **Relationships:** Belongs to one User Profile or future Household scope; owns Planned Meals; may inform Shopping Lists.

### Planned Meal

- **Purpose:** Represent a meal assigned to a date or meal occasion within a Meal Plan.
- **Owner Bounded Context:** Meal Planning.
- **Aggregate Ownership:** Child Entity of Meal Plan.
- **Lifecycle:** Proposed, Planned, Revised, Removed, Completed.
- **Relationships:** Belongs to one Meal Plan; may reference one Recipe or accepted Recommendation Option; may create shopping needs.

### Shopping List

- **Purpose:** Represent intended purchases required for kitchen activity.
- **Owner Bounded Context:** Shopping Optimization.
- **Aggregate Ownership:** Aggregate Root.
- **Lifecycle:** Draft, Generated, Reviewed, Active, Completed, Cancelled, Archived, Revised.
- **Relationships:** Belongs to one User Profile or future Household scope; owns Shopping Items; may be informed by Meal Plans, Pantry availability, and Recommendations.

### Shopping Item

- **Purpose:** Represent one ingredient purchase intention inside a Shopping List.
- **Owner Bounded Context:** Shopping Optimization.
- **Aggregate Ownership:** Child Entity of Shopping List.
- **Lifecycle:** Open, Completed, Removed.
- **Relationships:** Belongs to one Shopping List; references an Ingredient meaning and may be informed by Pantry Items or Planned Meals.

### Kitchen Recommendation

- **Purpose:** Represent a contextual AI-assisted cooking decision and its acceptance state.
- **Owner Bounded Context:** AI-Assisted Kitchen Decision Support.
- **Aggregate Ownership:** Aggregate Root.
- **Lifecycle:** Requested, Created, Presented, Accepted, Rejected, Superseded, Unable to Complete.
- **Relationships:** Belongs to one User Profile; references Pantry, Recipe, Meal Plan, and Shopping context; owns Recommendation Options and MVP conversation context.

### Recommendation Option

- **Purpose:** Represent one possible cooking choice within a Kitchen Recommendation.
- **Owner Bounded Context:** AI-Assisted Kitchen Decision Support.
- **Aggregate Ownership:** Child Entity of Kitchen Recommendation.
- **Lifecycle:** Proposed, Selected, Rejected, Superseded.
- **Relationships:** Belongs to one Kitchen Recommendation; may reference one Recipe; may guide a Planned Meal or Shopping List after explicit user action.

### Recommendation Conversation

- **Purpose:** Preserve the bounded conversation context that explains a Recommendation request and response in the MVP.
- **Owner Bounded Context:** AI-Assisted Kitchen Decision Support.
- **Aggregate Ownership:** Child Entity of Kitchen Recommendation in the MVP.
- **Lifecycle:** Open, Completed, Archived or Deleted according to product retention policy; exact lifecycle requires validation.
- **Relationships:** Belongs to one Kitchen Recommendation; may contain multiple user-assistance turns; does not own User Profile, Pantry, Recipe, Meal Plan, or Shopping data.

The future tactical model may promote Recommendation Conversation to an independent aggregate if AI Chat develops durable continuity, independent privacy rules, retention needs, feedback behavior, or search value.

## Aggregate Persistence

### Aggregate Root

An Aggregate Root is the durable business entry point for its aggregate. It owns the aggregate's identity, lifecycle, and consistency boundary. Other aggregates reference the root's business identity rather than reaching into child entities.

### Child Entities

Child entities are persisted within the ownership boundary of their Aggregate Root. Their lifecycle is controlled by the root and they are not independently authoritative in another context.

Examples:

- Pantry owns Pantry Items.
- Meal Plan owns Planned Meals.
- Shopping List owns Shopping Items.
- Kitchen Recommendation owns Recommendation Options and MVP Recommendation Conversation.

### Value Objects

Value Objects are represented as part of the business meaning of their owning entity or root. They do not receive independent aggregate ownership merely because they may be represented separately in a logical schema.

Examples include Ingredient Name, Ingredient Quantity, Expiry Date, Meal Slot, Meal Date, Planning Period, Cooking Time, Serving Size, Preference, Cooking Constraint, Recommendation Score, and lifecycle statuses.

## Relationship Model

### One-to-One

One-to-one relationships represent a unique business association:

- One Account has one User Profile in the MVP.
- One User Profile has one personal Pantry in the MVP.
- One Kitchen Recommendation has one bounded Recommendation Conversation context in the MVP, subject to future validation.

These relationships may change for Household Collaboration, multiple profiles, or independent AI Conversation behavior in future phases.

### One-to-Many

One-to-many relationships represent a root or context governing multiple child or related business concepts:

- One Pantry owns many Pantry Items.
- One Recipe may have many Recipe Favorites.
- One Meal Plan owns many Planned Meals.
- One Shopping List owns many Shopping Items.
- One Kitchen Recommendation owns many Recommendation Options.
- One User Profile may have many Meal Plans, Shopping Lists, and Kitchen Recommendations over time.

The many side does not gain ownership of the aggregate root's rules.

### Many-to-Many

Many-to-many relationships represent business associations that require an explicit relationship concept rather than direct shared ownership:

- Users may favorite many Recipes, and Recipes may be favorited by many Users; Recipe Favorite represents the association.
- Meal Plans may require many ingredients, and ingredients may appear across many meal plans; Planned Meal and Shopping Need Assessment preserve the business meaning without making Pantry own plans.
- Recommendations may consider many Pantry Items and Recipes, while those Pantry Items and Recipes may participate in many Recommendations; Recommendation references context but does not own it.
- Future Household Collaboration may connect many Household Members to shared Pantry, Meal Plan, and Shopping scopes through explicit participation meaning.

## Cross Context References

Cross-context references preserve ownership by storing or communicating the business identity and purpose needed by the consuming context, not the full internal representation of the source context.

- User Profile is referenced as the personal scope for Pantry, Meal Plan, Shopping List, and Kitchen Recommendation.
- Pantry Items are referenced by Recommendation and Shopping Need Assessment as availability context; Pantry Management remains authoritative.
- Recipes are referenced by Recommendation Options and Planned Meals; Culinary Knowledge remains authoritative for recipe meaning.
- Accepted Recommendation Options may guide Planned Meals; Meal Planning owns the commitment.
- Meal Plans may inform Shopping Lists; Shopping Optimization owns purchase intent.
- Shopping completion does not directly create Pantry Items; a user action must establish pantry availability.
- Future Household scope may replace or supplement User Profile scope only after household ownership rules are validated.

Cross-context references must not create direct ownership, shared aggregate transactions, or a global Ingredient or Recipe model.

## Data Lifecycle

### User

Account lifecycle: Pending -> Active -> Restricted -> Closed.

User Profile lifecycle: Created -> Incomplete -> Ready -> Updated.

Account closure or user deletion requests must apply the approved privacy and retention policy across owned personal context. Exact retention periods are outside this conceptual design.

### Pantry

The Pantry remains active while its owning personal or future household scope is active. Pantry Items move through Available, Running Low, Expiring Soon, Consumed, or Removed according to business actions and policies. Pantry history is represented through business events and audit information rather than a separate authoritative ownership model in this document.

### Recipe

Recipes are available for discovery while active in the Culinary Knowledge context. Recipe revisions must preserve the meaning required by existing favorites, recommendations, and planned meals. Retirement must not rewrite historical recommendation or planning meaning.

### Meal Plan

Meal Plan lifecycle: Draft -> Planned -> In Progress -> Completed, with Revised and Cancelled alternatives. Planned Meals may be changed or removed without changing the meaning of the source Recipe or Recommendation.

### Shopping List

Shopping List lifecycle: Draft -> Generated -> Reviewed -> Active -> Completed, with Cancelled, Revised, and Archived alternatives. Shopping Items may be individually Open, Completed, or Removed. Completion does not assert that Pantry state changed.

### Recommendation

Kitchen Recommendation lifecycle: Requested -> Created -> Presented -> Accepted, with Rejected, Superseded, and Unable to Complete alternatives. An Accepted Recommendation cannot return to Draft or Proposed. Acceptance only authorizes later explicit planning or shopping actions.

### AI Conversation

In the MVP, Recommendation Conversation is retained within Kitchen Recommendation context. Its business lifecycle follows the associated recommendation and product retention policy. If AI Chat becomes a durable independent capability, it may become an independent aggregate with its own lifecycle, privacy, and retention rules.

## Auditing Strategy

Auditing is a business accountability concern and must remain separate from aggregate ownership.

### Created

Owned aggregates and child entities should retain the business creation moment and responsible actor or product decision context. Creation history supports accountability and recommendation interpretation.

### Updated

Meaningful changes should retain the last business update moment and responsible actor or product decision context. Updates must not erase the prior meaning needed to understand accepted plans, completed shopping actions, or recommendation decisions.

### Soft Delete

Soft deletion is a conceptual lifecycle choice for user-owned information that may need history, recovery, audit, or privacy processing. Candidate uses include removed Pantry Items, removed Planned Meals, removed Shopping Items, removed Recipe Favorites, restricted Accounts, and retired Recommendations.

Soft deletion must not make an entity appear active in business views and must not be used to avoid legitimate privacy deletion obligations.

### Versioning

Business versioning is applicable where historical meaning matters:

- Recipe revisions may need to preserve the version considered by a past Recommendation or Planned Meal.
- User preference changes may need historical interpretation for Recommendation context.
- Recommendations should preserve the context and rationale used when they were created.
- AI conversation and prompt history may need a business reference to the applicable assistance context, without exposing provider implementation.

This section does not prescribe version columns, locking, or implementation strategy.

## AI Data Strategy

### Prompt history

The system may retain a business reference to the prompt purpose, prompt revision, context policy, and evaluation context used for AI assistance. Prompt history must not expose provider secrets, credentials, or unnecessary personal data.

### Conversation

Recommendation Conversation captures the bounded business context needed to understand a request and its response in the MVP. It should retain only information justified by product value, privacy, evaluation, and retention policy. It is not a source of truth for Pantry, Recipe, Meal Plan, Shopping List, or User Profile.

### Recommendation history

Kitchen Recommendation is the durable business record of recommendation decisions. It should preserve recommendation identity, context references, options, rationale or limitations, lifecycle state, user acceptance or rejection, and relevant business timestamps. It must distinguish generated guidance from user-confirmed decisions.

AI data should support recommendation quality, acceptance measurement, safety review, and future personalization without making external provider output authoritative.

## Future Database Evolution

Future additions may include:

- **Household Collaboration:** Household scope, membership participation, shared Pantry, shared Meal Plan, and shared Shopping List ownership.
- **Nutrition Guidance:** Nutrition goals, trusted guidance, nutrition summaries, and health-related policies.
- **Ingredient Catalog:** Shared ingredient identity, regional naming, substitutions, nutrition references, and categorization if current context-specific meanings become insufficient.
- **Notifications and Reminders:** Reminder intent and delivery history while the originating domain remains responsible for why the reminder exists.
- **SaaS Commercial Operations:** Plans, entitlements, usage, premium participation, and subscription relationships.
- **Grocery Partnerships:** Partner purchase intent, partner offers, and approved grocery relationships without making DapurPintar AI own fulfillment.
- **Recommendation Feedback:** Feedback, issue reports, acceptance context, and recommendation history if learning becomes a distinct business capability.
- **Advanced Read Models:** Pantry analytics, spending analytics, home dashboards, and future operational or product reporting.
- **AI Conversation Aggregate:** Independent conversation ownership if chat develops durable value beyond recommendation context.

Each addition must preserve aggregate ownership, bounded-context language, privacy, and the PostgreSQL system-of-record principle. New data stores should be introduced only for a validated business or scale requirement.

## Risks

- A shared logical model may accidentally become a global data model that erases bounded-context ownership.
- Cross-context references may be mistaken for shared transactions or direct write authority.
- The model may over-normalize business concepts and make user-facing decisions difficult to reconstruct.
- The model may under-normalize snapshots and duplicate authoritative data without clear historical purpose.
- AI conversation and recommendation history may retain more personal or prompt data than the product needs.
- Recipe revisions may invalidate past recommendation or meal-plan meaning if historical context is not preserved.
- Household, nutrition, commercial, and grocery extensions may introduce ownership and privacy rules that cannot be inferred from the MVP.
- Redis may be misused as a durable source of truth for sessions, recommendations, or business events.
- Soft deletion or retention policy may conflict with privacy deletion, audit, or regulatory requirements.
- Future analytics and read models may become accidental business authorities.

## Scope Boundary

This document intentionally excludes:

- Goose migration files.
- SQL DDL, SQLC queries, repository code, and Go structs.
- PostgreSQL-specific data types, indexes, implementation constraints, and syntax.
- Database deployment, backups, replication, scaling topology, and infrastructure operations.
- REST API payloads or persistence models derived directly from API contracts.
- Detailed column definitions, foreign-key implementation, and constraint implementation.

The conceptual and logical model must be refined during detailed database design while remaining consistent with the approved DDD aggregates, bounded contexts, event catalog, and Container/Component architecture.

## Diagram Reference

The implementation-independent ERD is maintained in `docs/architecture/diagrams/erd.mmd`.
