# DapurPintar AI Tactical DDD Model

## Executive Summary

DapurPintar AI helps individuals and households make better kitchen decisions using pantry context, culinary knowledge, meal planning, shopping intent, and personal preferences. The tactical domain model translates the strategic bounded contexts into business objects that have clear ownership, lifecycle, and consistency rules.

The MVP model is centered on seven aggregate roots: **Account**, **User Profile**, **Pantry**, **Recipe**, **Meal Plan**, **Shopping List**, and **Kitchen Recommendation**. The Kitchen Recommendation aggregate is the primary expression of the Core Domain. It coordinates a recommendation decision without taking ownership of pantry truth, recipe truth, meal-plan commitments, or shopping commitments.

The model is intentionally conservative. Aggregates are boundaries for business consistency, not a mirror of screens or storage. Entities have identity and lifecycle within an aggregate. Value objects express meaningful concepts whose validity depends on their values. Domain services are reserved for business decisions that genuinely span aggregate responsibilities. The model remains subject to validation through Event Storming, product discovery, and MVP feedback.

## Tactical DDD Principles

- **Business language first:** Names must come from the product and user vocabulary: pantry, meal plan, recipe, shopping list, preference, and recommendation.
- **One owner for each rule:** A business rule belongs to the aggregate or bounded context that has the authority to decide it.
- **Aggregate boundaries protect consistency:** A single aggregate is the boundary for immediate business validity.
- **Small aggregates:** Include only the information required to make the aggregate's decisions, not every related concept.
- **Reference other aggregates by meaning:** A recommendation may use pantry and recipe context without owning those concepts.
- **Entities have identity:** An entity remains distinguishable through changes during its lifecycle.
- **Value objects express meaning:** A value object is defined by its values and is replaceable when those values change.
- **AI assists, the domain decides:** Generated suggestions are proposals until accepted or confirmed through product rules.
- **Explicit transitions:** State changes such as proposed, accepted, planned, completed, or expired must have business meaning.
- **No shared global model:** Similar words may have different meanings in different bounded contexts.
- **Protect privacy:** Personal and household context crosses boundaries only when needed for a legitimate business decision.
- **MVP before completeness:** Model the approved MVP behavior first; future nutrition, household, notification, and commercial behavior remains extensible but not assumed.

## Aggregate Design

### Aggregate selection

An aggregate is selected when a group of business concepts must change together to remain valid. The following questions guide the boundary:

- What decision must always be valid immediately after a change?
- Which concept has authority over that decision?
- Can the related information change independently without violating the rule?
- Would including it make the aggregate responsible for another bounded context's language?

### MVP aggregate overview

| Bounded context | Aggregate root | Business purpose | Consistency focus |
|---|---|---|---|
| Identity and Access | Account | Establish trusted account identity and access state | Account identity and access status |
| User Context and Preferences | User Profile | Maintain personal cooking context | Preference and constraint validity |
| Pantry Management | Pantry | Maintain the user's kitchen availability | Item quantity and expiry validity |
| Culinary Knowledge and Recipe Experience | Recipe | Define a usable cooking option | Recipe completeness and cooking suitability |
| Meal Planning | Meal Plan | Define intended meals over a planning period | Meal schedule and meal commitment validity |
| Shopping Optimization | Shopping List | Define intended purchases for kitchen activity | Shopping item status and list coherence |
| AI-Assisted Kitchen Decision Support | Kitchen Recommendation | Record and govern a contextual cooking decision | Recommendation provenance, relevance, and acceptance state |

### Aggregate boundary decisions

- A Pantry owns its pantry items because quantity and expiry rules must remain consistent within the pantry context.
- A Meal Plan owns its planned meals because dates, meal slots, and planning status must agree.
- A Shopping List owns its shopping items because list status and completion are one user intention.
- A Kitchen Recommendation owns the recommendation decision and its lifecycle, but only references information from other aggregates.
- A Recipe owns recipe meaning and cooking guidance; a recommendation may use a recipe without changing the recipe.
- A User Profile owns declared preferences and constraints; it does not own the recommendation made from them.
- An Account owns trusted identity; it does not own personal kitchen preferences.

## Aggregate Roots

### Account

The Account is the authoritative business entry point for trusted identity.

Responsibilities:

- Establish account identity.
- Control whether the account is active or restricted.
- Support the identity needed to access personal kitchen context.

It does not contain pantry, recipe, meal, shopping, or recommendation decisions.

### User Profile

The User Profile is the authoritative source for personal cooking context.

Responsibilities:

- Maintain the user's declared preferences.
- Maintain relevant cooking constraints and goals.
- Provide a coherent personal context for decision support.

It does not decide which recipe is best or which meal is accepted.

### Pantry

The Pantry is the authoritative source for what the user considers available in the kitchen.

Responsibilities:

- Add, adjust, and remove pantry items.
- Maintain quantity and ingredient category meaning.
- Track expiry context.
- Identify ingredients that require attention or are at risk of waste.

It does not own recipe suitability, planned meals, or purchase decisions.

### Recipe

The Recipe is the authoritative source for a cooking option and its instructions.

Responsibilities:

- Define recipe identity and cooking meaning.
- Provide ingredients, instructions, servings, and expected preparation effort.
- Support discovery and user preference for a recipe.

It does not determine personal suitability; that is the responsibility of Kitchen Recommendation.

### Meal Plan

The Meal Plan is the authoritative source for intended meals over a defined planning period.

Responsibilities:

- Organize daily and weekly meal intent.
- Maintain the relationship between a planned meal and its position in the planning period.
- Allow a user to accept, change, or remove a planned meal according to business rules.

It does not own the recommendation rationale or shopping-list completion.

### Shopping List

The Shopping List is the authoritative source for intended kitchen purchases.

Responsibilities:

- Represent ingredients the user intends to acquire.
- Combine needs from meal intent and pantry gaps without changing pantry truth.
- Track whether a shopping item remains open, completed, or removed.

It does not own meal-plan decisions or supplier relationships.

### Kitchen Recommendation

The Kitchen Recommendation is the Core Domain aggregate root. It represents a contextual decision-support result rather than a permanent fact about the kitchen.

Responsibilities:

- State why a recommendation is relevant to the user's context.
- Present one or more practical cooking options.
- Identify the context used to form the recommendation.
- Track whether a recommendation is proposed, accepted, rejected, or superseded.
- Preserve the distinction between generated guidance and user-confirmed decisions.

It does not own the underlying pantry, recipe, user preference, meal-plan, or shopping data.

## Entities

Entities have identity and a meaningful lifecycle within their owning aggregate. These are business entities, not storage records.

| Aggregate | Entity | Business meaning |
|---|---|---|
| Account | Account | The trusted identity whose access state can change over time. |
| User Profile | Preference Set | A coherent set of personal cooking preferences that can be revised while remaining part of the profile's history. |
| Pantry | Pantry Item | A particular ingredient position in a pantry with a quantity and lifecycle. |
| Recipe | Recipe Favorite | A user's enduring preference for a recipe, including its changing status. |
| Meal Plan | Planned Meal | A meal intention assigned to a meaningful day or meal occasion. |
| Shopping List | Shopping Item | A particular ingredient purchase intention with a completion lifecycle. |
| Kitchen Recommendation | Recommendation Option | A candidate cooking choice with its own acceptance or rejection state within a recommendation. |
| Kitchen Recommendation | Recommendation Conversation | A bounded exchange that gives context to a recommendation request and response. |

An entity should not be introduced merely because a screen or external response contains a separate field. It needs identity, lifecycle, or a business rule that distinguishes it from a value.

## Value Objects

Value objects have no independent identity. They are valid because of their values and are replaced when those values change.

| Value object | Meaning | Likely contexts |
|---|---|---|
| Ingredient Reference | The business meaning of an ingredient in a specific context | Pantry, Recipe, Shopping, Nutrition |
| Quantity | An amount paired with an appropriate unit | Pantry, Recipe, Shopping |
| Expiry Date | The date after which an ingredient needs attention | Pantry |
| Ingredient Category | The classification used to organize kitchen ingredients | Pantry |
| Serving Size | The intended number of portions | Recipe, Meal Plan |
| Preparation Time | The time a cooking option requires | Recipe, Recommendation |
| Meal Date | The date associated with a planned meal | Meal Plan |
| Planning Period | The daily or weekly period governed by a meal plan | Meal Plan |
| Preference | A stated taste, cuisine, or cooking preference | User Profile, Recommendation |
| Cooking Constraint | A condition such as time, equipment, budget, or avoidance | User Profile, Recommendation |
| Nutrition Goal | A desired nutrition direction | User Profile, Nutrition, future Recommendation |
| Recommendation Rationale | The business explanation for relevance | Recommendation |
| Recommendation Status | Proposed, accepted, rejected, or superseded meaning | Recommendation |
| Shopping Status | Open, completed, or removed meaning | Shopping List |
| Confidence Statement | A bounded expression of recommendation certainty or limitation | Recommendation |

The same word may be represented by different value objects in different contexts. An Ingredient Reference in Pantry Management does not automatically have the same meaning as an ingredient requirement in Recipe Management.

## Domain Services

Domain services are stateless business capabilities used only when a decision does not naturally belong to one aggregate.

### Kitchen Context Assessment

Combines approved context from the User Profile, Pantry, Recipe, Meal Plan, and Shopping List to identify the relevant situation for a cooking decision. It does not own any of those contexts.

### Recommendation Suitability

Assesses whether a cooking option is suitable against the user's stated preferences, constraints, available ingredients, time, and other accepted context. It protects the difference between relevance and certainty.

### Meal Planning Guidance

Turns accepted cooking guidance into candidate daily or weekly meal intent while respecting the Meal Plan aggregate's scheduling rules.

### Shopping Need Assessment

Compares meal intent and available pantry context to identify purchase needs. The result is used by Shopping List without changing Pantry or Meal Plan ownership.

### Ingredient Waste Prioritization

Identifies which available ingredients should receive attention based on expiry and intended use. Pantry remains the authority for ingredient state.

Domain services must not become a generic utility layer or a second owner for aggregate rules.

## Repository Interfaces

Repository interfaces are conceptual business access contracts used by the domain and application language. They describe what a context needs to know, not how information is stored or retrieved.

| Interface | Business responsibility | Owning context |
|---|---|---|
| Account Records | Find and maintain trusted account identity and access state | Identity and Access |
| User Context Records | Find and maintain declared preferences and constraints | User Context and Preferences |
| Pantry Records | Find and maintain the authoritative pantry state | Pantry Management |
| Recipe Knowledge | Find and maintain recipes and cooking guidance | Culinary Knowledge and Recipe Experience |
| Meal Plan Records | Find and maintain planned meal intent | Meal Planning |
| Shopping List Records | Find and maintain purchase intent and shopping status | Shopping Optimization |
| Kitchen Recommendation Records | Find and maintain recommendation decisions and their lifecycle | AI-Assisted Kitchen Decision Support |

Each interface belongs to the language of its owning bounded context. A consuming context should ask for the business information it needs rather than access another aggregate's internal representation.

## Domain Events (Identified Only)

The following events identify meaningful business occurrences. Their payloads, delivery rules, and technical representation are intentionally not defined here.

### MVP events

- **AccountRegistered:** A new account has been established.
- **AccountAccessChanged:** An account's ability to participate has changed.
- **UserPreferencesChanged:** Personal cooking context has been changed.
- **PantryItemAdded:** An ingredient has entered the user's pantry context.
- **PantryItemAdjusted:** The quantity or expiry context of a pantry item has changed.
- **PantryItemRemoved:** An ingredient has left the pantry context.
- **IngredientApproachingExpiry:** A pantry ingredient requires attention.
- **RecipeFavorited:** A user has expressed an enduring preference for a recipe.
- **KitchenRecommendationRequested:** A user or product flow has asked for decision support.
- **KitchenRecommendationCreated:** A contextual recommendation has been produced.
- **RecommendationAccepted:** A user has accepted a recommendation for use.
- **RecommendationRejected:** A user has rejected a recommendation.
- **MealPlanCreated:** A daily or weekly planning intention has been established.
- **MealPlanned:** A meal has been assigned to a planning period.
- **PlannedMealChanged:** A planned meal has been changed or removed.
- **ShoppingListGenerated:** Purchase intent has been created from kitchen needs.
- **ShoppingItemCompleted:** A shopping intention has been fulfilled by the user.

### Future events

- **HouseholdCollaborationStarted:** A shared kitchen context has been established.
- **HouseholdMemberParticipationChanged:** A household member's participation has changed.
- **NutritionGoalChanged:** A nutrition goal has been established or revised.
- **ReminderRequested:** A business situation requires a timely user prompt.
- **CommercialParticipationChanged:** A customer's product participation or entitlement has changed.

These events are identified to clarify business vocabulary and relationships only. They are not an event architecture or integration design.

## Aggregate Invariants

### Account invariants

- An account must have a valid identity before it can participate in protected kitchen decisions.
- An inactive or restricted account cannot act as an active participant.
- Account access state must not change the meaning of pantry, recipe, meal, or shopping information.

### User Profile invariants

- Preferences and constraints must be internally coherent and understandable.
- A preference must be attributable to the profile that declared it.
- A recommendation cannot be treated as a user preference without explicit user acceptance or confirmation.

### Pantry invariants

- A pantry item must identify an ingredient and a meaningful quantity.
- Quantity cannot be negative.
- Expiry information must be valid when supplied.
- Removing an item must not silently create or change a shopping commitment.
- Pantry state is the authority for availability; recommendations cannot rewrite it.

### Recipe invariants

- A recipe must contain enough cooking meaning to be usable by the target user.
- Ingredient requirements, preparation effort, and instructions must remain coherent.
- A favorite recipe must refer to a recipe that exists in the recipe context.
- A recipe does not become personally suitable merely because it is available.

### Meal Plan invariants

- A planned meal must belong to a defined planning period and meaningful date or occasion.
- A planning period cannot contain conflicting meals where the product rules disallow conflict.
- A meal recommendation remains a proposal until it is intentionally planned.
- Changing a plan must not silently change the recipe or pantry truth.

### Shopping List invariants

- A shopping item must represent a meaningful ingredient need.
- A completed shopping item cannot remain simultaneously open.
- Generated purchase intent must be distinguishable from manually added purchase intent.
- A shopping list cannot claim that an ingredient is already in the pantry unless Pantry Management confirms it.

### Kitchen Recommendation invariants

- A recommendation must identify the user context for which it was made.
- A recommendation must have an explainable relevance basis or an explicit limitation.
- Generated guidance must remain distinguishable from accepted user decisions.
- Only the user or an explicit product rule can move a recommendation into an accepted state.
- A recommendation cannot claim authoritative pantry, recipe, meal-plan, shopping, or nutrition facts it does not own.

## Aggregate Lifecycles

Aggregate lifecycles describe meaningful business states and permitted transitions. They are intentionally expressed at a high level and do not prescribe implementation mechanics.

### Kitchen Recommendation lifecycle

`Requested -> Created -> Presented -> Accepted`

Alternative terminal paths:

- `Presented -> Rejected`
- `Presented -> Superseded`
- `Created -> Unable to Complete` when the available context is insufficient or the assistance cannot be trusted

Rules:

- A recommendation must be created from an identifiable user context.
- Only a presented recommendation can be accepted or rejected by the user.
- An accepted recommendation may guide a Meal Plan or Shopping List but does not automatically become either one.
- A superseded recommendation remains historical context and cannot be presented as current guidance.

### Meal Plan lifecycle

`Draft -> Planned -> In Progress -> Completed`

Alternative transitions:

- `Draft -> Cancelled`
- `Planned -> Revised`
- `Planned -> Cancelled`
- `Revised -> Planned`

Rules:

- A meal plan must have a defined planning period before it can become planned.
- A planned meal must occupy a valid date or meal occasion.
- Revision preserves the distinction between the previous intention and the current intention.
- Completion indicates that the planning period or meal intention has been fulfilled according to product meaning; it does not claim that every ingredient was purchased.

### Shopping List lifecycle

`Draft -> Generated -> Reviewed -> Active -> Completed`

Alternative transitions:

- `Draft -> Cancelled`
- `Generated -> Revised`
- `Reviewed -> Revised`
- `Active -> Archived`

Rules:

- Generated shopping intent must be reviewable before becoming active purchase intent.
- An active list may contain open and completed shopping items, but an individual item cannot be both at once.
- Completion of a list means the user's shopping intention is closed; it does not change Pantry truth without an explicit pantry action.

### Pantry lifecycle

`Available -> Running Low -> Expiring Soon -> Consumed`

Alternative transitions:

- `Available -> Removed`
- `Running Low -> Available` when quantity is replenished
- `Expiring Soon -> Consumed`
- `Expiring Soon -> Removed`

Rules:

- The lifecycle describes the business attention needed for an ingredient, not an assertion that the ingredient is physically present.
- Pantry Management remains the authority for every transition.
- Recommendations may prioritize a state but may not move it without a pantry decision.

### Account and User Profile lifecycle

Account: `Pending -> Active -> Restricted -> Closed`

User Profile: `Created -> Incomplete -> Ready -> Updated`

The Account lifecycle controls trusted participation. The User Profile lifecycle controls whether sufficient personal context exists for personalization; it does not imply that an AI recommendation is available or correct.

## Domain Policies

Domain policies are business rules that guide decisions across bounded contexts. They coordinate meaning without taking ownership away from the aggregate or context that owns the underlying facts.

### Context-aware recommendation policy

Recommendations should prefer options that satisfy known personal preferences and constraints, use available pantry ingredients where practical, and match the user's available time. Missing context must reduce certainty rather than be silently invented.

### Waste reduction policy

When two suitable cooking options are otherwise comparable, the product should prioritize ingredients that are approaching expiry or otherwise at higher risk of waste. Pantry Management remains the authority for expiry and availability.

### User decision policy

AI guidance is a proposal until the user accepts or confirms it. No recommendation may silently become a planned meal, purchase intention, pantry change, or nutrition commitment.

### Planning continuity policy

Meal planning should favor a practical sequence of meals over isolated suggestions. Accepted guidance may inform a plan, but Meal Planning owns whether the user commits to that plan.

### Shopping efficiency policy

Shopping intent should account for available pantry quantities and planned meals to avoid unnecessary duplicate purchases. Shopping Optimization owns the purchase intention, while Pantry Management owns availability.

### Trust and explanation policy

Every significant recommendation should be explainable in terms of the context that influenced it. Where the product lacks sufficient information, the recommendation should communicate a limitation or request clarification.

### Privacy boundary policy

Personal preferences, household context, and conversation context may be used only for a relevant business purpose and only within the access scope accepted by the user.

### Health responsibility policy

Nutrition-related constraints must come from an accepted Nutrition Guidance capability when that capability exists. General kitchen assistance must not invent medical claims or silently act as a health authority.

### Commercial neutrality policy

Commercial participation may limit access, usage, or availability of a capability, but it must not change the meaning of pantry availability, recipe suitability, meal intent, or shopping need.

## Specification Candidates

The following business specifications are candidates for formal, composable rules during implementation. They are named here to guide discovery and Event Storming; their expression and composition are not defined yet.

### Pantry specifications

- Ingredient is available for cooking.
- Ingredient is approaching expiry.
- Ingredient quantity is sufficient for a recipe requirement.
- Ingredient belongs to a permitted category.

### User context specifications

- Recipe matches declared preferences.
- Recommendation respects a cooking constraint.
- Recommendation fits the user's available preparation time.
- Context is sufficient for a trustworthy recommendation.

### Recipe specifications

- Recipe is complete enough to cook.
- Recipe can use the available ingredient context.
- Recipe meets a stated serving or preparation requirement.

### Meal planning specifications

- Planned meal fits the planning period.
- Planned meal does not violate a scheduling rule.
- Meal plan contains a useful balance of accepted meals.

### Shopping specifications

- Shopping need is not already covered by pantry availability.
- Shopping item is required by an accepted meal intention.
- Shopping list is ready for user review.

### Recommendation specifications

- Recommendation is relevant to the current user context.
- Recommendation has a known rationale.
- Recommendation is safe to present as decision support.
- Recommendation remains a proposal until user acceptance.

Specifications should remain business predicates that can be discussed with Product and domain stakeholders. They should not become a substitute for aggregate invariants or a general-purpose rule container.

## Business Rules

- Recommendations should prioritize ingredients that are available and approaching expiry when the resulting meal is otherwise suitable.
- Recommendations should respect declared preferences and constraints rather than treating generic popularity as personalization.
- A recommendation should consider preparation time and the user's available time when that context is known.
- A meal plan should reduce decision fatigue by turning accepted guidance into a clear daily or weekly intention.
- A shopping list should account for existing pantry availability before proposing a purchase.
- An automatically generated shopping item must remain reviewable by the user before it becomes a confirmed intention.
- AI assistance may propose, explain, or prioritize; it must not silently commit a meal, alter pantry truth, or confirm a purchase.
- User acceptance is a business signal, not proof that a recommendation is objectively correct.
- Personal context should be used only for a relevant decision and should remain protected from unnecessary sharing.
- Nutrition-related guidance must not be presented as medical advice without an explicitly validated future capability.
- Household sharing must not expose personal preferences or private kitchen context without an explicit sharing rule.
- Future commercial participation may restrict access or usage but must not alter core kitchen meaning.

## Transaction Boundaries

Transaction boundaries are business consistency boundaries. They describe which decision must be valid together, without prescribing a technical transaction mechanism.

| Boundary | Must be consistent together | Independent outcomes |
|---|---|---|
| Account | Identity and access state | Profile preferences and kitchen activity |
| User Profile | Preference set and cooking constraints | Recommendations made using the profile |
| Pantry | Pantry item changes, quantity, and expiry context | Meal plans and shopping intentions |
| Recipe | Recipe meaning and cooking guidance | Personal suitability and meal planning |
| Meal Plan | Planning period, planned meals, and schedule validity | Pantry updates and shopping completion |
| Shopping List | Shopping items and their statuses | Pantry availability and supplier fulfillment |
| Kitchen Recommendation | Recommendation options, rationale, and acceptance state | User confirmation in a meal plan or shopping list |

Cross-aggregate workflows should preserve each aggregate's independent consistency. For example, accepting a recommendation does not make a meal planned and a shopping list completed in the same business boundary; it creates a valid starting point for those separate decisions.

## Domain Validation Rules

### Identity and personal context

- The acting person must be known and allowed to participate.
- Personal context must be attributable to the correct person or, in the future, the correct household scope.
- Shared use must be explicit; a personal preference is private by default.

### Pantry and ingredients

- Ingredient names must be understandable and consistently classified within the pantry context.
- Amounts and units must be meaningful for the ingredient.
- Expiry dates must not contradict the product's time interpretation.
- A pantry item with no usable amount cannot be treated as available for recommendation.

### Recipes and recommendations

- A recipe option must be complete enough for a user to act on it.
- Recommendation rationale must refer to known context, not invented user facts.
- A recommendation must declare when key context is missing or uncertain.
- Generated output must be rejected or revised when it conflicts with known constraints.

### Plans and shopping

- A planned meal must have a clear time and intended recipe or cooking option.
- A shopping item must have a clear ingredient need and status.
- Shopping generation must be reviewable and must not duplicate known available pantry quantities without a reason.
- Completed actions must remain distinguishable from suggestions.

### Future nutrition and household behavior

- Nutrition constraints must come from an accepted nutrition context and not be invented by general assistance.
- Household decisions must respect participation and authority rules.
- Commercial restrictions must be validated separately from kitchen suitability.

## Future Tactical Considerations

The following candidates may require refinement as the product moves beyond the MVP. They are deliberately not added to the current aggregate model.

### Ingredient Catalog

If ingredient naming, categorization, substitutions, nutrition references, and regional food terminology become complex, an **Ingredient Catalog** aggregate or bounded context may be needed. Until that complexity is proven, Pantry Management and Recipe Experience may retain their purpose-specific ingredient meanings.

### Conversation as an independent aggregate

The current Recommendation Conversation is part of Kitchen Recommendation because the MVP treats conversation as context for decision support. If AI Chat grows into a durable user activity with its own continuity, privacy, retention, feedback, and search behavior, Conversation may become an independent aggregate or bounded context.

### Recommendation history

If recommendation acceptance, rejection, comparison, and learning become a major product capability, Recommendation History may need a separate business responsibility. It should not be extracted merely to store more history; it requires its own lifecycle and user value.

### Household context

When shared pantry, meal plans, and shopping become first-class features, a Household or Household Collaboration aggregate may be needed to define participation, authority, and visibility. Personal User Profile preferences should remain distinct from shared household decisions.

### Nutrition context

If nutrition goals and health guidance become central, Nutrition Guidance may require its own aggregates and policies. Its boundaries should be established with appropriate domain expertise rather than inferred from general recommendation behavior.

### Commercial participation

When premium plans and usage limits become operationally significant, Commercial Participation may require a dedicated aggregate. It should govern access and entitlement decisions without entering the semantic ownership of kitchen concepts.

## Risks

- Aggregate boundaries may become too large if AI is allowed to own all context.
- Aggregates may be split too finely around UI actions rather than business consistency.
- Similar concepts may be duplicated without clear ownership across contexts.
- Recommendation acceptance may be mistaken for factual correctness.
- Pantry quantity and expiry semantics may be too ambiguous for reliable recommendations.
- Meal planning and shopping generation may create conflicting interpretations of user intent.
- Future household and nutrition capabilities may require new invariants that cannot be inferred from the MVP.
- Business rules may be encoded before user research validates them.
- Identified events may be treated as a technical commitment before the business meaning is stable.

## Assumptions

- The Core Domain is AI Kitchen Intelligence expressed through contextual kitchen decision support.
- Users remain the final decision-makers for recommendations, plans, and purchases.
- Pantry, Recipe, Meal Plan, Shopping List, User Profile, and Kitchen Recommendation need separate ownership boundaries.
- PostgreSQL-backed durability and repository patterns do not change the business aggregate boundaries.
- The MVP needs daily and weekly meal planning, pantry management, recipes, favorites, shopping lists, and AI recommendation/chat behavior.
- Family collaboration and nutrition are future contexts and may introduce additional aggregates later.
- Domain events will be refined after business scenarios and Event Storming, not treated as finalized integration contracts now.
- Tactical boundaries may change as product assumptions are validated through beta usage and feedback.

## Scope Boundary

This document defines a tactical business model only. It intentionally excludes database tables, SQL schema, Go packages, REST endpoints, infrastructure, event payloads, persistence mechanisms, and deployment design.
