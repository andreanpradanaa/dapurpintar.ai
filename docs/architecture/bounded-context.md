# DapurPintar AI Bounded Contexts

## Executive Summary

DapurPintar AI is a household kitchen-management product whose differentiating capability is personalized, trustworthy assistance with everyday cooking decisions. Bounded contexts are used to protect the meaning and ownership of business concepts as the product grows from an MVP into a broader AI Kitchen Ecosystem.

The proposed Core Bounded Context is **AI-Assisted Kitchen Decision Support**. It uses business meaning from the Pantry, Culinary Knowledge, User Context, Meal Planning, and Shopping Optimization contexts to help users decide what to cook and how to act. It does not own the truth of those surrounding contexts. Identity and Access, Notifications and Reminders, and future SaaS Commercial Operations are supporting or generic contexts that provide capabilities required by the product but are not the primary source of differentiation.

These boundaries are strategic DDD hypotheses. They should be validated and refined through Event Storming, product discovery, user interviews, and MVP feedback. The boundaries describe language, responsibility, ownership, and relationships only.

## Context Identification

### Core Bounded Context

#### AI-Assisted Kitchen Decision Support

The context where DapurPintar AI creates its distinctive value: transforming authorized kitchen and personal context into relevant, practical, and trustworthy cooking guidance.

MVP status: **Core**.

### Supporting Bounded Contexts

#### User Context and Preferences

Defines the personal context that makes kitchen guidance relevant, including preferences, constraints, time considerations, and intended goals.

MVP status: **Supporting, MVP**.

#### Pantry Management

Defines the user's view of available ingredients and their use-by context. It provides the kitchen facts needed to avoid duplicate purchases and reduce waste.

MVP status: **Supporting, MVP**.

#### Culinary Knowledge and Recipe Experience

Defines recipe discovery, recipe meaning, cooking guidance, and the user's relationship with preferred recipes.

MVP status: **Supporting, MVP**.

#### Meal Planning

Defines intended meals over daily and weekly horizons. It turns individual decisions into a repeatable cooking habit.

MVP status: **Supporting, MVP**.

#### Shopping Optimization

Defines the user's intended purchases and the practical organization of ingredients needed to carry out meal decisions.

MVP status: **Supporting, MVP**.

#### Family and Household Collaboration

Defines shared kitchen responsibility and participation across household members, including shared pantry, meals, and shopping decisions.

MVP status: **Supporting, post-MVP**.

#### Nutrition Guidance

Defines nutrition-oriented information, goals, and constraints that may influence meal and recipe decisions.

MVP status: **Supporting, post-MVP**.

### Generic Bounded Contexts

#### Identity and Access

Defines trusted account identity and access scope. It protects the use of personal and future household context but does not define kitchen decisions.

MVP status: **Generic, MVP**.

#### Notifications and Reminders

Defines the delivery of timely prompts related to business situations such as ingredient expiry or planned meals. The reason for a reminder remains owned by the originating business context.

MVP status: **Generic, post-MVP**.

#### SaaS Commercial Operations

Defines future plans, entitlements, usage limits, premium participation, and commercial relationships. It must not redefine what a suitable recipe or meal is.

MVP status: **Generic, future**.

## Context Responsibilities

| Bounded Context | Responsibility | Primary business question |
|---|---|---|
| AI-Assisted Kitchen Decision Support | Produce relevant and trustworthy cooking guidance from approved context | "What should this person cook or do next?" |
| User Context and Preferences | Define the user's relevant preferences and constraints | "What matters to this person?" |
| Pantry Management | Define what ingredients are available and at risk of expiry | "What is available in the kitchen?" |
| Culinary Knowledge and Recipe Experience | Define available recipes and cooking guidance | "What can be cooked and how?" |
| Meal Planning | Define intended meals by day and week | "What does the person intend to cook?" |
| Shopping Optimization | Define what the person intends to acquire | "What should be bought to carry out the plan?" |
| Family and Household Collaboration | Define shared participation and household responsibility | "Who can participate in this kitchen decision?" |
| Nutrition Guidance | Define nutrition-related information and goals | "What health-related considerations apply?" |
| Identity and Access | Define who is acting and what personal or household scope is trusted | "Who is this person and what may they access?" |
| Notifications and Reminders | Deliver timely prompts for relevant business situations | "When should the person be reminded?" |
| SaaS Commercial Operations | Define commercial access and product participation | "What level of product participation applies?" |

## Context Ownership

Ownership means that the context is the final business authority for the meaning and rules of its concepts. A context may use another context's information without owning or redefining it.

| Context | Owns | Does not own |
|---|---|---|
| AI-Assisted Kitchen Decision Support | Recommendation intent, relevance, explanation, and user decision support | Pantry truth, recipe truth, meal commitments, shopping commitments, identity, or nutrition policy |
| User Context and Preferences | Personal preferences and cooking constraints | Authentication, household membership, or generated recommendations |
| Pantry Management | Ingredient availability and expiry understanding | Recipe suitability, meal commitments, or shopping decisions |
| Culinary Knowledge and Recipe Experience | Recipe and cooking guidance meaning | Personal suitability decisions or pantry availability |
| Meal Planning | Planned meal intent and schedule | Recommendation rationale or shopping ownership |
| Shopping Optimization | Intended purchase organization and shopping decisions | Pantry truth or meal-plan ownership |
| Family and Household Collaboration | Shared participation and responsibility | The private meaning of each member's personal preferences |
| Nutrition Guidance | Nutrition information, goals, and nutrition constraints | General recipe quality or unrestricted medical advice |
| Identity and Access | Trusted identity and access scope | Personal preference or kitchen meaning |
| Notifications and Reminders | Reminder delivery and timing preference | The business reason that makes a reminder necessary |
| SaaS Commercial Operations | Plans, entitlements, usage policy, and commercial participation | Kitchen recommendations and domain facts |

## Context Dependencies

Dependencies are expressed in business language:

- AI-Assisted Kitchen Decision Support depends upstream on User Context and Preferences, Pantry Management, and Culinary Knowledge and Recipe Experience for relevant context.
- Meal Planning and Shopping Optimization depend downstream on cooking guidance and intent from AI-Assisted Kitchen Decision Support.
- Shopping Optimization also depends on Meal Planning because planned meals create purchase needs.
- AI-Assisted Kitchen Decision Support depends on Nutrition Guidance only when nutrition becomes an accepted product capability.
- Pantry Management, Meal Planning, and Shopping Optimization depend on Family and Household Collaboration when the product introduces shared household ownership.
- All personal contexts depend on Identity and Access for trusted scope.
- Notifications and Reminders depends on meaningful situations originating in Pantry Management, Meal Planning, and AI-Assisted Kitchen Decision Support.
- AI-Assisted Kitchen Decision Support may depend on SaaS Commercial Operations for participation rules or usage limits, without allowing commercial rules to redefine recommendation meaning.

## Upstream and Downstream Relationships

### Upstream contexts

An upstream context supplies business meaning or constraints that downstream contexts need. The upstream context controls the meaning of what it supplies.

- **User Context and Preferences -> AI-Assisted Kitchen Decision Support:** personal relevance and constraints.
- **Pantry Management -> AI-Assisted Kitchen Decision Support:** available and expiring ingredients.
- **Culinary Knowledge and Recipe Experience -> AI-Assisted Kitchen Decision Support:** recipe possibilities and cooking guidance.
- **Nutrition Guidance -> AI-Assisted Kitchen Decision Support:** future nutrition goals and constraints.
- **Identity and Access -> all protected contexts:** trusted identity and access scope.

### Downstream contexts

A downstream context uses upstream meaning to perform its own business responsibility.

- **AI-Assisted Kitchen Decision Support -> Meal Planning:** suggested meals become planning options.
- **AI-Assisted Kitchen Decision Support -> Shopping Optimization:** cooking decisions become purchase intent.
- **Meal Planning -> Shopping Optimization:** planned meals inform what must be acquired.
- **Pantry Management, Meal Planning, and AI-Assisted Kitchen Decision Support -> Notifications and Reminders:** meaningful situations become timely prompts.
- **SaaS Commercial Operations -> AI-Assisted Kitchen Decision Support:** future participation rules constrain access without owning the decision itself.

### Relationship patterns

The preferred strategic relationship is a **customer-supplier** relationship: the downstream context states the business information it needs, while the upstream context remains responsible for its meaning and quality. Where a context must protect its language from an external or differently modeled context, an **anti-corruption layer** is required conceptually.

No context should become a global shared model. Shared meaning must be limited to stable concepts that genuinely belong to more than one context.

## Shared Concepts

Shared concepts are intentionally narrow. They are vocabulary references, not shared ownership of every detail.

| Shared concept | Contexts that use it | Ownership rule |
|---|---|---|
| Person | Identity and Access, User Context and Preferences, Family and Household Collaboration | Identity establishes who the person is; other contexts add their own meaning. |
| Household | Family and Household Collaboration, Pantry Management, Meal Planning, Shopping Optimization | Collaboration defines shared participation; kitchen contexts define their own household use. |
| Ingredient | Pantry Management, Culinary Knowledge and Recipe Experience, Shopping Optimization, Nutrition Guidance | Each context defines the ingredient meaning needed for its responsibility. |
| Recipe | Culinary Knowledge and Recipe Experience, AI-Assisted Kitchen Decision Support, Meal Planning, Nutrition Guidance | Culinary Knowledge owns recipe meaning; other contexts use suitable representations for their decisions. |
| Meal | Meal Planning, AI-Assisted Kitchen Decision Support, Nutrition Guidance, Shopping Optimization | Meal Planning owns planned-meal meaning; other contexts use meal information for their own purpose. |
| Preference and Constraint | User Context and Preferences, AI-Assisted Kitchen Decision Support, Nutrition Guidance | User Context owns declared personal context; other contexts interpret it within their boundaries. |
| Recommendation | AI-Assisted Kitchen Decision Support, Meal Planning, Shopping Optimization | AI Assistance owns recommendation intent; downstream contexts own whether a recommendation becomes a plan or purchase decision. |

Shared concepts must remain small and stable. A context should not import another context's full language merely because the concepts have similar names.

## Context Integration Strategy

Integration should preserve business meaning and context autonomy.

1. **Use explicit business agreements:** Each relationship should identify what meaning is supplied, why it is needed, and which context owns it.
2. **Prefer purpose-specific language:** A recipe used for recommendation may not have the same meaning as a recipe used for nutrition guidance.
3. **Keep the Core Context focused:** AI Assistance coordinates decisions but does not become the owner of every kitchen concept.
4. **Translate at boundaries:** Information should be translated into the consuming context's language when concepts, rules, or trust levels differ.
5. **Protect privacy and scope:** Only the minimum personal or household context required for a business decision should cross a boundary.
6. **Keep future contexts optional:** Family, Nutrition, Notifications, and Commercial Operations must be attachable without changing the meaning of the MVP's Core Context.
7. **Validate relationships with the business:** Context relationships should be reviewed when product behavior, ownership, or commercial assumptions change.

## Anti-Corruption Layer Considerations

An anti-corruption layer is a business-language protection mechanism. It prevents a foreign model, vocabulary, or policy from distorting a context's own language.

### AI provider boundary

External AI concepts such as model-specific instructions, generated text, confidence signals, or provider-specific limitations must not define the product meaning of a recommendation. AI-Assisted Kitchen Decision Support translates external assistance into a product recommendation that can be assessed, explained, accepted, changed, or rejected by the user.

### Nutrition boundary

Nutrition Guidance should protect its specialized health and nutrition meaning from being reduced to generic recipe advice. General AI recommendations may request nutrition context, but they must not silently create nutrition policy or medical claims.

### Family boundary

Family and Household Collaboration should translate shared participation into the language understood by Pantry Management, Meal Planning, and Shopping Optimization. It should not force personal preferences to become shared household facts without an explicit business rule.

### Commercial boundary

SaaS Commercial Operations may determine whether a capability is available to a customer, but it must not change what a pantry item, meal, recipe, or recommendation means. Commercial vocabulary should remain outside kitchen decision language.

### External partner boundary

Future grocery, nutrition, or smart-kitchen partners will have their own business vocabulary. Their concepts should be translated before they influence household kitchen decisions. Marketplace and food-delivery behavior remains outside the current product domain.

## Future Extraction Strategy

Bounded contexts are business boundaries first. The MVP may host them together, but their ownership and language should remain distinct.

Future extraction should be considered only when one or more of these conditions are demonstrated:

- A context has materially different scaling, availability, or security needs.
- A context has a distinct business owner or delivery cadence.
- A context needs independent commercial or partner relationships.
- A context's change rate creates unacceptable risk for other contexts.
- The cost of separation is justified by measured product or operational value.

Likely candidates for future independent operation are AI-Assisted Kitchen Decision Support, Notifications and Reminders, SaaS Commercial Operations, or external partner capabilities. This is not a commitment to separate them. The decision must preserve context ownership, business agreements, and an understandable customer journey.

The Core Context should not be extracted merely because it is named "AI." Extraction is justified by business boundaries and measured needs, not technology fashion.

## Risks

- The Core Context could become an oversized "AI" context that owns recommendations, recipes, pantry, meal plans, and shopping indiscriminately.
- Similar words such as ingredient, recipe, meal, and recommendation could hide different meanings across contexts.
- Pantry, Meal Planning, and Shopping may disagree about whether a suggestion is merely proposed or has become a user commitment.
- User preferences, household preferences, and nutrition goals may be confused without explicit ownership.
- Identity and Access may be treated as a technical concern instead of a business trust boundary.
- Future commercial entitlements may leak into kitchen rules and distort product behavior.
- An external AI provider may introduce vocabulary or policies that corrupt the product's trustworthy-assistance language.
- Premature extraction could create organizational and operational complexity before MVP assumptions are validated.
- Domain boundaries may need to change after beta testing, new personas, or regional expansion.

## Assumptions

- AI-assisted decision support is the primary differentiator of DapurPintar AI.
- Pantry, recipe, planning, shopping, and user context are separate business responsibilities even when they are experienced as one product.
- Users will understand the difference between a recommendation, a plan, and a purchase intention.
- Identity and access scope can protect personal context and later household context.
- Family collaboration and nutrition require explicit business rules and should not be inferred from MVP behavior.
- The MVP will validate the Core Context before future contexts become first-class capabilities.
- Context boundaries can evolve through Event Storming and product learning without being treated as permanent.
- The business will favor clear ownership and trustworthy AI over maximum feature breadth.

## Scope Boundary

This document defines strategic bounded contexts and their business relationships only. It intentionally does not define database tables, Go packages, REST APIs, events, entity models, deployment boundaries, or implementation mechanisms.
