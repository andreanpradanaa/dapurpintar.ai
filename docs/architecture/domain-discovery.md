# DapurPintar AI Domain Discovery

## Executive Summary

DapurPintar AI operates in the household kitchen-management business domain. Its purpose is not merely to publish recipes, but to help people make better daily cooking decisions: what to cook, how to use available ingredients, how to plan meals, and what to buy. The product combines pantry awareness, culinary knowledge, meal planning, shopping support, and personalized AI assistance into one kitchen companion.

The proposed Core Domain is **AI-Assisted Kitchen Decision Support**. This is the differentiating business capability that turns a user's real context into relevant and trustworthy cooking recommendations. The surrounding domains provide the facts and constraints that make those recommendations useful. Generic domains provide capabilities that are necessary but not differentiating.

This is a strategic domain map, not a final implementation boundary. The classifications and relationships should be validated through Event Storming, product discovery, user research, and MVP feedback. No entities, database structures, API contracts, or code are defined here.

## Business Domain

The overall business domain is **Personal and Household Kitchen Management**.

It covers the decisions and activities that happen before, during, and around cooking:

- Understanding what food and ingredients are available.
- Finding suitable recipes and cooking inspiration.
- Choosing and organizing daily or weekly meals.
- Identifying what must be purchased.
- Adapting decisions to preferences, time, household needs, nutrition goals, and waste reduction.
- Learning from user choices and feedback to improve future assistance.

The business value is created when users spend less time deciding, buy more efficiently, waste fewer ingredients, and cook meals that fit their real circumstances. The initial market is Indonesia, with a future direction toward Southeast Asia, broader household use, partner ecosystems, and SaaS commercialization.

## Core Domain

### AI-Assisted Kitchen Decision Support

The Core Domain is the capability that differentiates DapurPintar AI from generic recipe applications and disconnected tools. It interprets authorized user and kitchen context, applies product intent and constraints, and produces personalized, practical, and trustworthy recommendations.

The Core Domain is responsible for:

- Turning pantry availability, preferences, constraints, and plans into cooking guidance.
- Recommending recipes that are relevant to the user's actual situation.
- Helping users decide what to cook and how to use ingredients efficiently.
- Supporting meal-plan decisions and shopping decisions through connected context.
- Presenting AI as decision support rather than as an authoritative source of user data.
- Capturing acceptance and feedback signals that help the product improve.
- Preserving trust through explainable intent, safety, privacy, and validation of recommendations.

The MVP expression of this domain is AI chat, recipe recommendation, pantry recommendation, and AI-assisted shopping-list generation. AI-assisted meal planning is a strategic extension that should use the same core decision-support language.

The Core Domain is not the OpenAI provider, prompt tooling, or a generic chat feature. Those are enabling capabilities. The business advantage lies in the product's context-aware decisions and the integrated kitchen experience.

## Supporting Domains

Supporting Domains provide important business capabilities that enable the Core Domain but are not, individually, the primary differentiator.

### Pantry Management

Maintains the user's understanding of available ingredients, quantities, categories, and expiration context. It provides the factual kitchen state needed to recommend recipes, reduce duplicate purchases, and prioritize ingredients that may be wasted.

MVP importance: **Core supporting capability**.

### Culinary Knowledge and Recipe Experience

Provides recipe discovery, recipe detail, cooking instructions, and the user's ability to retain preferred recipes. It supplies the culinary content and user interaction context that the AI can use when forming recommendations.

MVP importance: **Core supporting capability**.

### Meal Planning

Organizes intended meals over daily and weekly horizons. It transforms isolated recommendations into a repeatable household habit and is directly connected to the north star metric of weekly AI-assisted meals planned.

MVP importance: **Core supporting capability**.

### Shopping Optimization

Represents the ingredients the user intends to acquire and supports automatic list generation from cooking intentions and pantry context. It connects planning decisions to practical household action.

MVP importance: **Core supporting capability**.

### User Context and Preferences

Captures the personal context that makes recommendations relevant, such as preferences, constraints, available time, and intended goals. It should remain focused on product value and should not become an unrestricted profile-data domain.

MVP importance: **Supporting capability**.

### Family and Household Collaboration

Coordinates shared kitchen context across household members, including shared pantry, meal planning, and shopping. It is strategically important to the family value proposition but is explicitly post-MVP.

MVP importance: **Future supporting domain**.

### Nutrition Guidance

Provides nutrition information, summaries, goals, and future personalized guidance. It can constrain or enrich the Core Domain, but its specialized knowledge and health-related risk justify a distinct future domain rather than folding it into general AI assistance.

MVP importance: **Future supporting domain**.

## Generic Domains

Generic Domains are necessary across the product but are not sources of differentiated kitchen value. They should be adopted or standardized where practical rather than over-designed.

### Identity and Access

Provides account identity, authentication, access control, and the basis for protecting user-owned kitchen context. It is essential to trust and privacy but does not differentiate the kitchen experience.

### Notifications and Reminders

Delivers time-based or event-based reminders such as expiry and meal reminders. The business rules that decide why a reminder matters belong to the relevant supporting domain; delivery is generic.

### Observability and Operational Insights

Provides the shared ability to understand product health, AI usage, latency, errors, and key business signals. It supports operational excellence and learning but is not a kitchen business capability.

### SaaS Commercial Operations

Future subscription, entitlement, usage, premium, and partner-commercial capabilities support commercialization. They are important to the business model but should remain separate from kitchen decision logic and are not part of the MVP runtime scope.

### External Provider Integration

Provider access, including AI provider access, is a generic integration concern. OpenAI is an initial provider, while the product's differentiated capability remains the Core Domain's contextual decision support.

## Domain Responsibilities

| Domain | Strategic responsibility | MVP position |
|---|---|---|
| AI-Assisted Kitchen Decision Support | Convert kitchen context into personalized, trustworthy cooking decisions | Core |
| Pantry Management | Maintain usable knowledge of available and expiring ingredients | Supporting, MVP |
| Culinary Knowledge and Recipe Experience | Provide recipes, cooking guidance, discovery, and saved preferences | Supporting, MVP |
| Meal Planning | Turn cooking decisions into daily and weekly intent | Supporting, MVP |
| Shopping Optimization | Turn pantry gaps and meal intent into practical shopping action | Supporting, MVP |
| User Context and Preferences | Supply personal constraints and preferences for relevance | Supporting, MVP |
| Family and Household Collaboration | Enable shared kitchen decisions and responsibilities | Supporting, post-MVP |
| Nutrition Guidance | Add health and nutrition-oriented decision context | Supporting, post-MVP |
| Identity and Access | Establish trusted identity and protect user context | Generic, MVP |
| Notifications and Reminders | Deliver timely user prompts based on domain signals | Generic/supporting, post-MVP |
| Observability and Operational Insights | Make product and business behavior measurable and diagnosable | Generic, cross-cutting |
| SaaS Commercial Operations | Support plans, entitlements, usage, and recurring revenue | Generic, future |
| External Provider Integration | Connect to replaceable third-party capabilities | Generic, cross-cutting |

## Business Capabilities

### Discover and Understand

- Capture user preferences and cooking constraints.
- Maintain visibility of pantry ingredients and expiry context.
- Discover recipes and cooking instructions.

### Decide and Plan

- Recommend what to cook based on real context.
- Organize daily and weekly meal plans.
- Adapt recommendations to time, preferences, available ingredients, and future nutrition goals.

### Act and Coordinate

- Generate and manage shopping lists.
- Help users use ingredients before they are wasted.
- Enable future household collaboration around shared kitchen decisions.

### Learn and Improve

- Observe recommendation acceptance and user feedback.
- Improve personalization without compromising privacy.
- Measure weekly AI-assisted meals planned and related product outcomes.

### Trust and Operate

- Protect identity, kitchen data, and AI context.
- Provide transparent and safe AI assistance.
- Monitor availability, quality, cost, and operational behavior.
- Enable future SaaS plans and commercial entitlements.

## Domain Relationships

The relationships below describe business dependencies and information flow, not technical calls or database joins.

| Relationship | Meaning |
|---|---|
| User Context -> Core Decision Support | Preferences and constraints establish what is relevant to a person. |
| Pantry Management -> Core Decision Support | Available and expiring ingredients ground recommendations in reality. |
| Culinary Knowledge -> Core Decision Support | Recipes and cooking knowledge provide possible actions. |
| Meal Planning <-> Core Decision Support | The Core Domain proposes plans and learns from accepted or changed plans. |
| Shopping Optimization <-> Core Decision Support | Recommendations can produce shopping intent, while shopping constraints can shape recommendations. |
| Core Decision Support -> Pantry Management | Accepted decisions may motivate users to use or update pantry information; ownership of pantry truth remains with Pantry Management. |
| Core Decision Support -> Meal Planning | Personalized decisions can become planned meals; the plan remains owned by Meal Planning. |
| Core Decision Support -> Shopping Optimization | Meal and pantry context can produce a practical list; list behavior remains owned by Shopping Optimization. |
| Identity and Access -> All protected domains | Trusted identity and access scope protect personal and future household context. |
| Nutrition Guidance -> Core Decision Support | Future nutrition constraints enrich recommendations without making general AI the owner of nutrition policy. |
| Family Collaboration -> Pantry, Meal Planning, Shopping | Future shared ownership enables household-level workflows. |
| Notifications <- Supporting Domains | Domain signals may result in reminders, while delivery remains generic. |
| Operational Insights <- All domains | Business and operational signals are observed without taking ownership of domain behavior. |
| External Provider Integration <- Core Decision Support | Provider capabilities are replaceable means for delivering decisions, not business concepts. |

## High-Level Domain Landscape

The landscape has three strategic layers:

1. **Core:** AI-Assisted Kitchen Decision Support, where DapurPintar AI creates differentiated value.
2. **Supporting:** Pantry, Culinary Knowledge, Meal Planning, Shopping Optimization, and User Context in the MVP; Family Collaboration and Nutrition Guidance later.
3. **Generic:** Identity and Access, Notifications, Observability, SaaS Commercial Operations, and External Provider Integration.

The Core Domain should consume context through explicit business relationships and should not absorb ownership of pantry truth, recipe content, meal plans, shopping lists, identity, or provider behavior. This separation protects the product's language as it grows and prevents AI from becoming an unbounded global domain.

## Future Evolution

### MVP to validated product

The initial landscape should prioritize the Core Domain and its MVP supporting capabilities. User research and beta analytics should validate whether AI recommendations are more valuable than generic recipe search, whether users maintain a digital pantry, and whether weekly meal planning becomes a habit.

### Personalization and household collaboration

As the product moves toward AI personalization and Family Workspace, User Context and Family Collaboration may become more explicit supporting domains. Shared household context should be modeled as a business concept only after roles, ownership, and collaboration rules are validated.

### Nutrition and responsible guidance

Nutrition Guidance may grow into a specialized domain with its own trusted sources, constraints, and risk policies. It should not be treated as interchangeable with general recipe recommendation, particularly where health claims or personalized goals are involved.

### Ecosystem and commercialization

Grocery integrations, public APIs, premium plans, partner capabilities, and enterprise use cases may introduce new domains and external relationships. The Core Domain should remain focused on kitchen decisions while partner and commercial concerns evolve around it.

### Regional and platform expansion

Expansion from Indonesia to Malaysia, Singapore, and wider Southeast Asia may introduce localization, language, cultural food knowledge, and regional business concerns. These should be evaluated as separate supporting or generic capabilities rather than embedded as assumptions in every domain.

## Risks and Assumptions

### Assumptions

- Users value context-aware AI assistance more than a standalone recipe catalogue.
- Users are willing to maintain enough pantry information for recommendations to be useful.
- Daily and weekly meal planning can become a recurring behavior.
- Pantry, recipe, plan, and shopping context can be combined without violating privacy expectations.
- AI recommendations can be made trustworthy through grounding, validation, transparency, and user feedback.
- The initial domain landscape can support the MVP scale and future SaaS growth without premature service decomposition.

### Risks

- The Core Domain may be defined too broadly as "AI" and lose a precise business purpose.
- Overlapping ownership between Pantry, Meal Planning, Shopping, and AI could create contradictory decisions.
- The product may mistake generated suggestions for authoritative facts or user commitments.
- Nutrition and health-related expectations may introduce safety and trust requirements beyond the MVP.
- Family collaboration may expose unclear ownership and authorization rules when introduced.
- External AI-provider behavior, cost, latency, or policy changes may affect the perceived value of the Core Domain.
- Generic capabilities such as billing, notifications, or analytics may be pulled into the Core Domain and create unnecessary coupling.
- Product assumptions may change after beta testing, requiring domain boundaries to be refined.

### Discovery Actions

- Conduct Event Storming with Product, Engineering, AI, QA, and domain stakeholders.
- Validate domain language with the primary personas: young families and busy professionals.
- Identify ownership, decision points, policies, and handoffs before designing detailed models.
- Record unresolved boundary questions as domain decisions rather than prematurely encoding them in technical structure.
- Revisit this map at MVP beta, AI personalization, Family Workspace, and Nutrition milestones.

## Scope Boundary

This document intentionally excludes entities, value objects, aggregates, database tables, REST endpoints, deployment units, and Go package structure. Those decisions belong to later System Design artifacts and must remain consistent with this strategic domain landscape.
