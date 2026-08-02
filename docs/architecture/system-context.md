# DapurPintar AI System Context

## Executive Summary

The C4 Level 1 System Context Diagram describes DapurPintar AI in its business environment. It shows the people and external organizations or systems that interact with the product, the purpose of those relationships, and the boundary of responsibility owned by DapurPintar AI.

The diagram intentionally treats DapurPintar AI as one system. Its internal modular monolith, bounded contexts, domain services, storage, API, AI gateway, and observability concerns are not represented at this level. Those concerns belong to lower-level architecture views.

The MVP system exists to help users manage kitchen context and make better cooking decisions through pantry management, recipe discovery, AI assistance, meal planning, and shopping-list support. OpenAI is the only current external system explicitly required by the MVP architecture. Other external participants are future relationships supported by the product roadmap and are marked accordingly.

## System Scope

### Inside DapurPintar AI

DapurPintar AI owns the integrated business experience for personal and household kitchen management, including:

- Account participation and personal profile context.
- User preferences and cooking constraints.
- Pantry visibility, ingredient quantities, and expiry context.
- Recipe discovery, recipe detail, cooking guidance, and favorites.
- AI-assisted kitchen decision support and recommendation acceptance.
- Daily and weekly meal planning.
- Manual and generated shopping lists.
- Business rules that preserve user control, privacy, recommendation trust, and domain ownership.

The system may use external assistance, but it owns the product meaning of a recommendation and does not treat AI output as authoritative pantry, recipe, meal-plan, shopping, or nutrition truth.

### Outside DapurPintar AI

The following remain outside the system boundary:

- The internal operations and policies of an external AI provider.
- Future identity or account-verification providers.
- Grocery marketplace inventory, offers, ordering, and fulfillment.
- Notification delivery outside the product's reminder meaning.
- Payment processing and external subscription settlement.
- Household members' private decisions outside the shared product scope.
- Nutrition professionals' independent advice and professional responsibility.
- Commercial operators' business operations outside product participation and entitlements.
- Food sales, food delivery, hardware, smart-home devices, social feeds, POS, accounting, and ERP inventory behavior.

## Primary Actors

### User (MVP)

The User is the primary actor. A User registers, manages a profile, records pantry ingredients, explores recipes, requests AI assistance, accepts or rejects recommendations, plans meals, and reviews or completes shopping intentions.

The User remains the final decision-maker. DapurPintar AI may propose and prioritize, but it must not silently commit a meal, alter pantry truth, or confirm a purchase.

### Household Member (Future)

A Household Member participates in shared pantry, meal-plan, and shopping decisions when Family and Household Collaboration is introduced. Household participation and visibility are future business capabilities and are not required for the MVP.

### Grocery Partner (Future)

A Grocery Partner participates in future grocery or supermarket integration. The partner may receive user-approved purchase intent or provide partner information, but grocery inventory, offers, ordering, and fulfillment remain outside DapurPintar AI.

### Nutrition Professional (Future)

A Nutrition Professional may contribute specialized nutrition guidance or support future nutrition-oriented experiences. The professional's independent advice and health responsibility remain outside DapurPintar AI.

### Commercial Operator (Future)

A Commercial Operator manages future product participation, premium capabilities, partnerships, or enterprise relationships. Commercial operations may govern access or entitlements but do not own kitchen decisions or recommendation suitability.

## External Systems

### External AI Provider (MVP)

The External AI Provider supplies generative assistance used by DapurPintar AI for recommendation and AI-assisted kitchen experiences. DapurPintar AI owns context selection, safety, relevance, validation, user presentation, and acceptance meaning. The provider does not own user kitchen data or product decisions.

OpenAI is the initial provider identified by the architecture and ADRs. The System Context view names the business role, External AI Provider, rather than coupling the Level 1 model to a provider implementation.

### Authentication Provider (Future)

An Authentication Provider may support future identity or account-verification scenarios. Identity and access responsibility remains a DapurPintar AI business boundary; an external provider does not own user profile, pantry, or kitchen permissions.

### Grocery Marketplace (Future)

A Grocery Marketplace may support future grocery or supermarket integration. It may exchange approved purchase intent or partner information. It remains responsible for its own product availability, pricing, ordering, and fulfillment.

### Notification Provider (Future)

A Notification Provider may deliver future reminders related to ingredient expiry or planned meals. DapurPintar AI owns why a reminder matters and the user's product-level preference; the provider owns external delivery.

### Payment Provider (Future)

A Payment Provider may support future premium subscriptions and recurring commercial participation. It owns payment processing and settlement; DapurPintar AI owns product participation and entitlement meaning.

## Relationships

| From | To | Business interaction | Status |
|---|---|---|---|
| User | DapurPintar AI | Registers, manages personal context, records pantry items, discovers recipes, requests assistance, plans meals, and manages shopping intent | MVP |
| Household Member | DapurPintar AI | Participates in shared pantry, meal-plan, and shopping decisions | Future |
| Nutrition Professional | DapurPintar AI | Provides or supports specialized nutrition guidance | Future |
| Grocery Partner | DapurPintar AI | Participates in approved grocery integration and purchase-intent exchange | Future |
| Commercial Operator | DapurPintar AI | Manages future commercial participation, premium capabilities, or partnerships | Future |
| DapurPintar AI | External AI Provider | Requests contextual AI assistance and receives generated recommendations or conversation assistance | MVP |
| DapurPintar AI | Authentication Provider | Exchanges identity or account-verification information | Future |
| DapurPintar AI | Grocery Marketplace | Shares approved purchase intent or receives partner information | Future |
| DapurPintar AI | Notification Provider | Requests delivery of future expiry or meal reminders | Future |
| DapurPintar AI | Payment Provider | Requests future subscription payment processing and receives payment outcome information | Future |

### Business relationship notes

- User context, pantry context, recipe knowledge, meal intent, and shopping intent are owned by DapurPintar AI even when they influence an external interaction.
- A recommendation from the External AI Provider is not a business commitment until the User accepts or confirms it through DapurPintar AI.
- Completing a Shopping Item does not automatically add an ingredient to the Pantry; the User must explicitly record pantry availability.
- A Grocery Marketplace is not a food-delivery or marketplace capability owned by DapurPintar AI.
- Future external systems are optional extensions and must not become dependencies of the MVP user journey.

## System Boundary

The DapurPintar AI system boundary includes the business capability to understand a user's kitchen context, combine it with appropriate recipe and planning context, provide personalized AI-assisted decision support, and help the user turn decisions into meal and shopping intentions.

The boundary includes responsibility for:

- The meaning and lifecycle of user-owned kitchen context.
- The meaning, rationale, safety, and acceptance state of recommendations.
- The distinction between suggestions, planned meals, and purchase intentions.
- Privacy, access scope, and user control over personal and future household context.
- Product rules that reduce food waste and prevent avoidable duplicate purchases.
- Business relationships with future partners, without owning the partners' operations.

The boundary excludes the internal behavior of external providers and partners. DapurPintar AI may request, receive, or share approved business information, but it does not own external identity verification, AI model behavior, grocery fulfillment, notification delivery, or payment settlement.

## Assumptions

- The User is the primary MVP actor and uses a smartphone-oriented product experience.
- DapurPintar AI is the system of responsibility for personal kitchen decisions and related user commitments.
- OpenAI, represented here as an External AI Provider, is the initial external AI relationship for MVP capabilities.
- External authentication, grocery, notification, and payment relationships are future capabilities supported by the roadmap, not MVP prerequisites.
- Household collaboration and nutrition guidance are future bounded-context extensions.
- External systems cannot silently change the meaning of pantry availability, recipe suitability, meal intent, or shopping intent.
- The system remains a single business system at C4 Level 1 even though internal architecture may evolve.

## Risks

- Users may attribute external AI output to DapurPintar AI and lose trust when recommendations are unsupported or irrelevant.
- Dependence on an External AI Provider may introduce cost, availability, policy, or quality changes.
- Future partner relationships may blur the boundary between DapurPintar AI and grocery, nutrition, notification, or commercial operations.
- Household sharing may expose private preferences or create unclear decision authority.
- External identity and payment providers may create privacy, compliance, or regional expansion concerns.
- Showing future systems in the context diagram may be mistaken for an MVP commitment; all future relationships must remain explicitly optional.
- Marketplace or food-delivery expectations could expand the system beyond the approved kitchen-management scope.
- The system context may be incorrectly expanded to include internal containers, databases, or infrastructure; those belong to lower-level C4 views.

## Diagram References

- Mermaid source: `docs/architecture/diagrams/system-context.mmd`
- Editable draw.io source: `docs/architecture/diagrams/system-context.drawio`
