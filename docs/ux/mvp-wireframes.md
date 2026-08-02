# DapurPintar AI MVP Wireframe Specification

## Document Control

| Item | Value |
|---|---|
| Milestone | M3 - UX/UI Design |
| Status | Draft - Awaiting Design Review |
| Scope | Low-fidelity screen structure and interaction behavior for MVP journeys |
| Parent document | `docs/ux/ux-ui-design.md` |
| Related diagram | `docs/ux/diagrams/mvp-screen-states.mmd` |

## Purpose

This document turns the M3 UX/UI foundation into low-fidelity screen specifications. It defines what each MVP screen helps the user do, the order of information, the primary action, responsive behavior, and important states.

These are wireframe contracts, not final visual mockups. The visual direction, typography, colors, and Context Strip defined in `ux-ui-design.md` remain the source for later high-fidelity design.

## Global Shell

### Desktop structure

```text
+----------------------------------------------------------------+
| mark        Today  Pantry  Discover  Planner  Shopping   user  |
+----------------------------------------------------------------+
|                                                                |
|                         focused work surface                  |
|                                                                |
+----------------------------------------------------------------+
```

- Persistent left or top navigation is quiet and never competes with the next action.
- The active destination uses `ink-900` and a small `chili-500` marker.
- The user menu contains Profile, Preferences, and Log out.
- Notifications are not an MVP destination; inline status belongs to the relevant screen.

### Mobile structure

```text
+--------------------------+
| menu     DapurPintar  user|
+--------------------------+
|                          |
|       focused content    |
|                          |
+--------------------------+
| Today Pantry Plan Shop   |
+--------------------------+
```

- Use a four-item bottom navigation: Today, Pantry, Planner, Shopping.
- Discover is reached from Today, Pantry recommendations, and the menu to avoid a crowded bottom bar.
- Preserve the current destination and scroll position when opening detail screens.

## Screen 01: Registration and Login

### Job

Start a trusted personal cooking workspace without introducing unnecessary setup before the user understands the product.

### Structure

```text
        [DapurPintar]
  Decide dinner with what you have.

  Email                         [             ]
  Password                      [             ]
  [ Create account ]

  Already have an account? Log in
  Small privacy and AI context note
```

### Behavior

- Registration and login are separate modes with a stable switch.
- Password guidance is visible before submission, not only after failure.
- Generic authentication errors avoid account enumeration.
- Successful registration or login continues to profile setup or Today according to account state.
- The screen never displays raw API, provider, or credential errors.

### States

- Initial.
- Submitting.
- Invalid field values.
- Authentication rejected.
- Network unavailable.
- Account restricted or pending verification.

## Screen 02: Profile Setup

### Job

Collect the minimum context needed to make the first recommendation useful.

### Structure

```text
  Make the first suggestion fit you
  Step 1 of 2                         [progress]

  What matters today?
  [ Quick meals ] [ Save money ] [ Use pantry ]

  Cooking time
  [ 15 min ] [ 30 min ] [ Flexible ]

  Dietary or cooking constraints (optional)
  [                                      ]

  [ Continue ]       Skip for now
```

### Behavior

- Explain why each preference helps without implying medical advice.
- Optional context can be skipped and completed later.
- Do not present nutrition goals in the MVP setup.
- Completion returns to Today with a clear invitation to add pantry items.

## Screen 03: Today

### Job

Help the user choose the next useful kitchen action, not summarize every module.

### Structure

```text
  Good evening, Sarah                 [Add context]
  What are you cooking today?

  [ 3 ingredients available · 1 use soon · 25 min ]
  ┌─────────────────────────────────────────────────┐
  │ Try: Tumis sayur telur                          │
  │ Uses 3 pantry items  ·  20 min                  │
  │ “Fits your time and uses spinach first.”        │
  │ [Use this idea]  [Ask about it]                │
  └─────────────────────────────────────────────────┘

  Continue where you left off
  [ Tuesday dinner ] [ Shopping list: 4 items ]

  [ Add pantry item ]       Browse recipes
```

### Behavior

- Today is a composed view, not a new business aggregate.
- Show a useful empty state when no recommendation exists: add pantry context, browse recipes, or request a recommendation.
- Context Strip appears before recommendation rationale.
- `Use this idea` accepts a specific Recommendation Option; it does not create a plan.
- Existing plan and shopping prompts link to their owning screens.

### Responsive behavior

- Mobile: recommendation first, then continuation cards, then secondary actions.
- Desktop: recommendation occupies the main column; context and continuation occupy a narrow side column.
- Avoid equal-weight card grids that make every action look urgent.

## Screen 04: Pantry

### Job

Let the user understand and correct available ingredients quickly.

### Structure

```text
  Pantry                              [ + Add item ]
  12 items · 3 categories · 1 use soon

  [ Search ingredients... ] [Category v]

  USE SOON
  ┌─────────────────────────────────────────────────┐
  │ Spinach       1 bunch       Use by 12 Aug       │
  │ [Edit]                                     ...  │
  └─────────────────────────────────────────────────┘

  ALL ITEMS
  Eggs           6             Protein
  Rice           2 kg           Staples
```

### Add/edit drawer

```text
  Add pantry item                         [Close]
  Ingredient name [                         ]
  Quantity         [       ] [unit v]
  Category         [       ]
  Expiry           [       ]  optional
  [Save item]
```

### Behavior

- Add flow asks for the minimum fields first.
- Expiry attention is a prioritization signal and includes exact dates.
- Edit preserves the list position and confirms the changed value.
- Delete or remove explains that the item will no longer be considered available.
- Search and filters preserve their state when opening an item.

## Screen 05: Discover and Recipe Detail

### Job

Help a user find and understand a cooking option, whether or not AI is available.

### Discover structure

```text
  Discover recipes
  [ What do you want to cook?                 ] [Search]
  [15 min] [Uses pantry] [Easy]

  Recommended for your pantry
  [ recipe card ] [ recipe card ]

  All recipes
  [ recipe card with time, difficulty, favorite ]
```

### Recipe detail structure

```text
  < Back to Discover                         [Favorite]
  Tumis sayur telur
  20 min · Easy · 2 servings

  Why it may fit
  3 ingredients from your pantry · 1 missing

  Ingredients       Instructions
  [available]       1. ...
  [missing]         2. ...

  [Ask for a recommendation] [Add to plan]
```

### Behavior

- Public recipe views show general recipe content only.
- Personalized suitability and pantry context require authentication.
- `Add to plan` opens explicit date or meal-occasion confirmation.
- Favorite uses an idempotent action and gives immediate feedback.

## Screen 06: Recommendation Detail and Conversation

### Job

Help the user evaluate one AI Recommendation with visible grounding and clear control.

### Structure

```text
  Recommendation                         [state]
  Based on your pantry and preferences
  [ 3 ingredients available · 25 min ]

  Option A: Tumis sayur telur
  Fits because: quick, uses spinach, matches your goal
  Available: spinach, eggs      Missing: garlic
  [Accept option]

  Option B: Nasi goreng sayur
  ...

  [Ask about this recommendation]
```

### Conversation structure

```text
  Tumis sayur telur                         [Back]
  AI suggestion · recommendation context

  Why is this a good fit?                  (user)
  It uses spinach first and takes 20 min.  (assistant)

  [Ask a question about this idea... ] [Send]
```

### Behavior

- Conversation is nested under the Recommendation and cannot become global chat.
- The user can ask for clarification, substitution, or practical detail only within approved capability boundaries.
- Assistant responses show limitations when context is missing.
- AI failure offers Browse Recipes or retry and preserves the recommendation state.
- Acceptance is always tied to one `optionId`.

## Screen 07: Planner

### Job

Turn a cooking intention into a visible daily or weekly commitment.

### Mobile structure

```text
  Planner                         [Week v]
  < 12 - 18 Aug >

  MON 12
  Dinner  [ + Choose a meal ]

  TUE 13
  Dinner  [ Tumis sayur telur ]

  WED 14
  Dinner  [ + Choose a meal ]
```

### Desktop structure

```text
  Planner          < This week >       [Daily] [Weekly]
  +---------+---------+---------+---------+---------+
  | Mon     | Tue     | Wed     | Thu     | Fri     |
  | Dinner  | Dinner  | Dinner  | Dinner  | Dinner  |
  | + add   | recipe  | + add   | + add   | recipe  |
  +---------+---------+---------+---------+---------+
  [Generate shopping list]
```

### Behavior

- Empty slots provide a clear action and do not look like missing data.
- Selecting an accepted Recommendation Option is explicit and confirms the date or occasion.
- Moving or removing a meal requires a non-drag alternative.
- Generate Shopping List explains which planned meals and pantry context will be used.
- A plan never changes Pantry automatically.

## Screen 08: Shopping List

### Job

Let the user review purchase intent before and during shopping.

### Structure

```text
  Shopping list                         [New list]
  Week of 12 Aug · 6 items · Draft
  Generated from: weekly plan + pantry

  REVIEW BEFORE SHOPPING
  [ ] Garlic          2 bulbs          ...
  [ ] Eggs            6                ...
  [ ] Soy sauce       1 bottle         ...

  [Activate list]

  COMPLETED
  [x] Rice            2 kg
```

### Behavior

- Generated items are purchase intent and remain editable before activation.
- Source context is visible but not editable from the list.
- Completion is immediate but reversible only if the domain lifecycle permits it.
- Completing a Shopping Item does not update Pantry and the screen must not suggest otherwise.
- Empty state offers manual item addition or explicit generation from a plan.

## Shared State Specifications

### Loading

- Preserve screen structure and show content-shaped placeholders.
- Keep navigation available unless the entire authenticated session is unavailable.
- Do not imply that an AI result exists before it is validated.

### Empty

Every empty state answers three questions:

- What is empty?
- Why is it useful?
- What is the next action?

Example: `Your pantry is empty. Add one ingredient and we can make the first suggestion more relevant.`

### Error

- Use plain language and identify the action that failed.
- Preserve user input where safe.
- Offer retry, correction, or a non-dependent alternative.
- Do not expose status internals, SQL, prompts, provider names, or stack traces.

### Permission and unavailable resource

- Do not reveal whether another user's resource exists.
- Return the user to a safe parent view with a clear next action.
- Never render private data from cached or previously visited state after authorization fails.

### AI degraded

- Keep pantry, recipe, planner, and shopping features usable.
- Explain that AI assistance is temporarily unavailable.
- Offer Browse Recipes, View Pantry, or retry.
- Keep the failed request bounded and prevent accidental duplicate commands.

## Responsive Rules

| Area | Mobile | Desktop |
|---|---|---|
| Navigation | Bottom bar plus menu | Persistent navigation rail or top bar |
| Recommendation | Single option stack | Main option with compact alternatives |
| Pantry | List and drawer editor | Two-column list and editor panel |
| Planner | Day stack | Week grid with day detail |
| Shopping | Full-width checklist | Checklist with source summary side panel |
| Conversation | Full-screen nested view | Focused panel with recommendation context |

Minimum touch targets, visible focus, readable line lengths, and reduced-motion behavior apply in both modes.

## API Interaction Map

| UI action | API capability | Resulting business state |
|---|---|---|
| Save profile | `PATCH /api/v1/profile` | User Profile updated |
| Add pantry item | `POST /api/v1/pantry/items` | Pantry Item added |
| Search recipes | `GET /api/v1/recipes` | Read-only recipe results |
| Favorite recipe | `PUT /api/v1/favorites/recipes/{recipeId}` | Favorite active |
| Request recommendation | `POST /api/v1/recommendations` | Recommendation requested/created |
| Accept option | `POST /api/v1/recommendations/{id}/options/{optionId}/accept` | Option accepted only |
| Add meal | `POST /api/v1/meal-plans/{id}/meals` | Planned Meal added |
| Generate shopping list | `POST /api/v1/shopping-lists/generate` | Shopping List generated for review |
| Complete shopping item | `POST /api/v1/shopping-lists/{id}/items/{itemId}/complete` | Shopping Item completed |

Frontend navigation must not invent a write action that is absent from the API contract.

## Review Checklist

- Does every screen have one primary action?
- Can a first-time user recover from an empty pantry without understanding the full product?
- Is AI grounding visible without exposing internal implementation?
- Are acceptance, planning, shopping, and pantry updates distinct actions?
- Are mobile and desktop layouts both usable without hiding important controls?
- Can keyboard and assistive-technology users complete the primary journey?
- Do error and degraded states preserve user trust and data?
- Does each screen map to an approved API capability?

## Exit Criteria

This wireframe specification is ready for high-fidelity design when:

- Product and Design approve the primary screen hierarchy.
- The primary user journey is usable on mobile and desktop.
- All MVP API interactions have a corresponding user action.
- State behavior is defined for loading, empty, error, permission, and AI degradation.
- No screen implies an out-of-scope domain or automatic cross-context mutation.

## Related Documents

- `docs/ux/ux-ui-design.md`
- `docs/ux/diagrams/mvp-screen-states.mmd`
- `docs/product/product-scope.md`
- `docs/product/feature-inventory.md`
- `docs/product/user-personas.md`
- `docs/architecture/api-design.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/ai-architecture.md`
