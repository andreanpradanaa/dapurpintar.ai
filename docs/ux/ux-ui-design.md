# DapurPintar AI UX/UI Design Foundation

## Document Control

| Item | Value |
|---|---|
| Milestone | M3 - UX/UI Design |
| Status | Draft - Awaiting Product and Design Review |
| Scope | MVP information architecture, primary journeys, interface direction, and interaction rules |
| Primary user | Sarah, Young Working Mom |
| Secondary user | Daniel, Busy Professional |
| Related diagram | `docs/ux/diagrams/m3-primary-user-flow.mmd` |

## Design Thesis

DapurPintar AI should feel like a calm kitchen board that turns scattered ingredients into one confident next action. The interface is not a recipe catalogue with an AI chat bolted on. Its primary job is to help a person answer:

> "Dengan bahan dan waktu yang saya punya, apa yang paling masuk akal untuk dimasak sekarang?"

The MVP experience therefore prioritizes a short path from context to decision:

```text
Set context -> See what matters -> Get a grounded suggestion
-> Choose an option -> Plan it or shop for it
```

The interface must preserve the distinction between an AI suggestion and a user commitment. Accepting a recommendation does not silently create a Meal Plan, Shopping List, or Pantry change.

## Product and UX Goals

- Help a user choose a menu in under two minutes.
- Make pantry state easy to scan, add, adjust, and correct.
- Make AI reasoning understandable without exposing prompts or technical model details.
- Turn one accepted cooking option into an explicit next action.
- Make weekly and daily planning feel lighter than a blank calendar.
- Generate a useful shopping list without hiding or silently changing the user's choices.
- Make errors, missing context, AI unavailability, and empty states actionable.
- Work well on a phone first while remaining clear on desktop.

## Non-Goals

M3 does not design or imply:

- Family collaboration or shared household permissions.
- Nutrition dashboards or medical guidance.
- Barcode, OCR, image recognition, voice, grocery integrations, or notifications.
- A public recipe marketplace or social feed.
- Autonomous AI actions that change user commitments.
- Frontend authorization as a substitute for backend authorization.

## Primary User Focus

### Sarah - primary

Sarah has a full-time job, a child, and limited decision energy. She cooks often, shops weekly, and wants healthy food without managing several disconnected tools.

UX implications:

- Put the next useful action before feature exploration.
- Make time, ingredients, and household preference visible inputs.
- Prefer concise recommendation cards over long lists.
- Let her correct pantry data quickly instead of requiring perfect setup.
- Keep planning and shopping as explicit follow-up actions.

### Daniel - secondary

Daniel needs speed, simple recipes, and immediate shopping clarity. He is comfortable with AI but does not want a long conversation before receiving a useful option.

UX implications:

- Support a fast path with minimal setup.
- Show preparation time and missing ingredients early.
- Keep AI clarification optional unless context is insufficient.
- Make “Cook this” and “Add to plan” distinct actions.

### Maya - future-informed

Maya's nutrition needs are useful as a constraint on future design language, but nutrition is not an MVP surface. M3 must not introduce nutrition claims or controls that imply the P1 Nutrition domain exists.

## Information Architecture

### Application shell

The authenticated shell uses five primary destinations:

| Destination | User question | MVP scope |
|---|---|---|
| Today | What should I do next? | Composed view of current recommendations, plan, and action prompts |
| Pantry | What do I have and what needs attention? | Pantry summary, items, categories, quantity, expiry |
| Discover | What can I cook? | Recipe search, recipe detail, favorites |
| Planner | What will I cook this week? | Daily and weekly Meal Plans |
| Shopping | What do I need to buy? | Shopping Lists, generated items, completion |

Profile and preferences are reached from the account menu and onboarding prompts. AI conversation is entered through a Recommendation, not presented as an unbounded primary destination.

### Public information architecture

- Landing page.
- Product explanation.
- General recipe discovery and recipe detail where public access is enabled.
- Registration.
- Login.

Public screens must not reveal personal pantry, profile, favorites, recommendations, or activity.

### Screen inventory

| Screen | Primary action | Data or API boundary |
|---|---|---|
| Landing | Start using DapurPintar AI | Public |
| Register | Create account | Accounts |
| Login | Start participation | Accounts |
| Profile setup | Declare cooking context | Profile and preferences |
| Today | Choose the next kitchen action | Composed authenticated reads |
| Pantry | Add or correct an ingredient | Pantry |
| Pantry item editor | Adjust quantity or expiry | Pantry Item |
| Discover | Search for a recipe | Recipes |
| Recipe detail | Understand and favorite a recipe | Recipe and Favorites |
| Recommendation detail | Evaluate AI guidance | Recommendation |
| Recommendation conversation | Clarify one recommendation | Recommendation Conversation |
| Meal Planner | Plan or revise meals | Meal Plans and Planned Meals |
| Shopping List | Review and complete purchases | Shopping Lists and Items |
| Preferences | Update constraints and goals | Profile preferences |

## Primary User Journeys

### First-use journey

1. User lands on the product and understands the concrete benefit.
2. User registers or logs in.
3. Product asks for only the minimum useful context: cooking goals, constraints, household size or serving context, and time preference.
4. User reaches Today with a clear invitation to add pantry context or request a recommendation.
5. Product explains that better pantry context improves relevance but does not block exploration.

### Decide what to cook

1. User opens Today or requests a Recommendation.
2. Product shows the context used: available ingredients, expiry attention, time, and preferences where relevant.
3. AI returns a small set of Recommendation Options with recipe reference, preparation time, relevant ingredients, limitations, and rationale.
4. User can accept one option, reject the recommendation, ask a bounded clarification, or view the recipe.
5. Acceptance changes Recommendation state only.
6. User separately chooses `Add to plan` or `Create shopping list`.

### Manage pantry

1. User opens Pantry and sees useful summary information before the full item list.
2. User adds an item with name, quantity, category, and optional expiry date.
3. Product confirms the saved item in plain language.
4. User can adjust or remove an item without losing orientation.
5. Expiry attention is visible as a prioritization signal, not as an alarmist warning.

### Plan a week

1. User opens Planner and selects a daily or weekly period.
2. Empty slots invite a meal choice rather than presenting a blank spreadsheet.
3. User selects a recipe or accepted Recommendation Option.
4. User confirms the date or meal occasion.
5. Planner shows the resulting commitment and offers an explicit path to generate a Shopping List.

### Build and use a shopping list

1. User starts from Shopping or an explicit planning action.
2. Product explains the source of generated items: planned meals, pantry gaps, or user input.
3. User reviews, removes, or adjusts items before activating the list.
4. User marks items complete while shopping.
5. Completing an item does not silently add it to Pantry; the UI must not imply that it does.

## Interaction Rules

### AI recommendation card

Every recommendation card should make these facts scannable:

- What to cook.
- Why it fits the current context.
- Preparation time and serving context.
- Ingredients available and ingredients missing.
- Important limitations or uncertainty.
- The current recommendation state.

Primary actions:

- `Use this idea` or `Accept option`.
- `Add to plan` only after acceptance or through an explicit planning path.
- `Make a shopping list` as a separate commitment.
- `Ask about this idea` for recommendation-scoped conversation.

Avoid anthropomorphic claims, unsupported nutrition claims, and “perfect” or “guaranteed” language.

### Forms

- Use progressive disclosure for pantry details and preferences.
- Preserve entered values when validation fails.
- Show units and expiry meaning in the user's language.
- Validate at the boundary and return a field-level correction.
- Never use a client-provided owner or account identifier in visible form controls.

### Destructive actions

Removing a Pantry Item, Favorite, Planned Meal, or Shopping Item requires a clear action label and an undo or recovery path where the domain lifecycle permits it. Destructive confirmation must state what changes and what does not change.

### Navigation and back behavior

- Back navigation preserves search, filters, selected planning period, and scroll position where practical.
- A Recommendation Conversation returns to its owning Recommendation.
- A recipe opened from a Recommendation retains a clear return path.
- A completed command returns the user to the updated business view, not a generic success page.

## Visual Direction

### Design language

The interface uses the visual vocabulary of an annotated kitchen board: quiet paper surfaces, dark pantry ink, small utility labels, and one bright action color that marks the next decision. It should feel precise and warm without becoming rustic decoration.

Signature element: **the Context Strip**. Recommendation and planning surfaces carry a compact horizontal strip showing the active context, for example `3 bahan tersedia · 1 segera dipakai · 25 menit`. This makes AI grounding visible and gives the product a recognizable interaction pattern.

### Color tokens

| Token | Value | Use |
|---|---|---|
| `ink-900` | `#17243A` | Primary text, navigation, dark surfaces |
| `paper-050` | `#F4F1E8` | App canvas and warm neutral surfaces |
| `card-000` | `#FFFEFA` | Cards and focused content |
| `chili-500` | `#F45B3C` | Primary action, urgent expiry attention |
| `herb-400` | `#A7D46F` | Available, accepted, healthy progress |
| `turmeric-400` | `#F3C969` | Context attention and pending state |
| `steel-400` | `#8793A5` | Secondary text and inactive controls |

Color must never be the only carrier of status. Pair status color with text, icon, or pattern.

### Typography

- **Display:** `Fraunces`, used for page titles and short recommendation headlines.
- **Body:** `Manrope`, used for navigation, controls, forms, and explanatory copy.
- **Utility:** `IBM Plex Mono`, used sparingly for quantities, time, dates, and context metadata.

Typography should remain readable in Indonesian, support dynamic text sizing, and avoid all-caps paragraphs.

### Layout

- Mobile-first single-column flow for decision tasks.
- Desktop uses a quiet two-zone layout: persistent navigation plus focused work surface.
- Maximum reading width for recommendation rationale and recipe instructions.
- Dense utility information belongs in the Context Strip or metadata row, not in paragraph blocks.
- Use restrained corner radii and strong alignment rather than decorative cards everywhere.

### Motion

Motion is limited to state confirmation and orientation:

- Recommendation generation reveals the Context Strip before the result.
- Saving a Pantry Item uses a brief confirmation transition.
- Planner drag or move feedback is optional and must have a non-drag alternative.
- Respect `prefers-reduced-motion` and never hide essential feedback in animation.

## Component Foundations

M3 should establish these reusable component families before screen-specific decoration:

- App shell and responsive navigation.
- Context Strip.
- Recommendation Card and Option Card.
- Pantry Item Row and expiry indicator.
- Recipe Card and Recipe Detail sections.
- Meal Slot and Planning Period selector.
- Shopping Item Row with completion state.
- Form field, validation message, and inline save state.
- Empty State, Loading State, Error State, and Degraded AI State.
- Confirmation, undo, and permission feedback.

Component names should describe user-facing meaning, not backend module or database terminology.

## Required States

Every MVP screen must define:

- Loading state.
- First-use empty state.
- Returning-user empty state.
- Validation error state.
- Authorization or unavailable-resource state.
- Network or dependency failure state.
- AI unavailable or unable-to-complete state where AI is involved.
- Success confirmation.
- Mobile and desktop layout behavior.
- Keyboard focus and screen-reader behavior.

### Degraded AI state

When AI cannot complete a request:

- Explain that the suggestion could not be completed.
- Preserve the user's pantry, recipe, plan, and shopping data.
- Offer a non-AI path such as Browse Recipes or View Pantry.
- Do not show provider names, prompts, retry internals, or raw errors.
- Allow a bounded retry when the operation is safe to repeat.

## Accessibility and Localization

- Meet WCAG 2.2 AA intent for contrast, focus, keyboard access, labels, and error messaging.
- Use semantic headings and landmarks.
- Ensure touch targets are comfortable on mobile.
- Provide text alternatives for status icons and expiry indicators.
- Do not rely on color alone.
- Support Indonesian copy first and allow longer localized strings.
- Use locale-aware dates, numbers, quantities, and currency when those values appear.
- Keep cooking terminology plain and test comprehension with target users.

## UX Measurement Plan

M3 prototypes and later implementation should measure:

- Time from authenticated entry to a useful recommendation.
- Time to choose a menu, with the MVP target under two minutes.
- Completion rate for profile setup and first pantry item.
- Recommendation presentation-to-option-acceptance rate.
- Accepted option to explicit Meal Plan conversion.
- Time to create or review a Shopping List.
- Error recovery and abandonment at each primary journey step.
- Qualitative trust: whether users understand why a recommendation appeared and what acceptance changes.

These measures complement product metrics such as recommendation acceptance, meal plans created, shopping lists generated, and weekly AI-assisted meals planned.

## Architecture and API Constraints

- Protected screens require authenticated backend state; route guards are not authorization.
- UI actions map to the approved `/api/v1` resources and commands.
- Recommendation acceptance must identify a specific Recommendation Option.
- Recommendation Conversation remains nested under its owning Recommendation in the MVP.
- Accepted AI guidance does not automatically create a Meal Plan or Shopping List.
- Shopping completion does not automatically update Pantry.
- Public recipe screens must not expose private context.
- Loading, error, and degraded states must reflect safe API error categories without leaking internals.

## Open Questions

- Should Today be the default authenticated landing view or should first-time users remain in onboarding until a minimum context is complete?
- Which pantry fields are essential for the first add flow, and which belong behind “more details”?
- Should a recommendation show one option by default or a compact set of alternatives?
- What is the clearest Indonesian wording for Recommendation Option acceptance?
- Should weekly planning use a calendar grid or a day-stack layout on desktop?
- Which recipe content is public in the MVP and which requires login?
- Which empty states should invite pantry setup versus recipe discovery?

## Exit Criteria

M3 foundation is ready for design review when:

- Primary MVP journeys are documented and aligned with personas and product scope.
- Information architecture covers all MVP screens without introducing post-MVP ownership.
- Recommendation, planning, shopping, and pantry interactions preserve explicit user decisions.
- A distinctive visual direction and reusable component foundation are defined.
- Loading, empty, error, permission, degraded-AI, mobile, desktop, and accessibility states are specified.
- UX measurements map to approved product success metrics.
- Open questions are resolved through product or user validation before high-fidelity implementation.

## Related Documents

- `docs/product/product-vision.md`
- `docs/product/product-scope.md`
- `docs/product/feature-inventory.md`
- `docs/product/user-personas.md`
- `docs/product/target-users.md`
- `docs/product/value-proposition.md`
- `docs/product/success-metrics.md`
- `docs/architecture/m2-signoff.md`
- `docs/architecture/api-design.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/ai-architecture.md`
- `docs/ux/diagrams/m3-primary-user-flow.mmd`
