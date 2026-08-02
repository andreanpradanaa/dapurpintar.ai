# DapurPintar AI Design System Foundation

## Document Control

| Item | Value |
|---|---|
| Milestone | M3 - UX/UI Design |
| Status | Draft - Awaiting Design Review |
| Scope | MVP visual tokens, components, states, and accessibility rules |
| Parent documents | `docs/ux/ux-ui-design.md`, `docs/ux/mvp-wireframes.md` |
| Related diagram | `docs/ux/diagrams/design-system-component-map.mmd` |

## Purpose

This document defines the reusable visual and interaction foundation for DapurPintar AI. It is intentionally implementation-neutral: tokens and component behavior can be translated into Tailwind, CSS variables, or another approved frontend system later.

The system should feel like an annotated kitchen board: calm paper surfaces, dark pantry ink, precise utility metadata, and one bright action color that marks the next decision. Reuse should create recognition without flattening every screen into the same card grid.

## Design Principles

- **Context before decoration:** show the ingredients, time, and reason before visual polish.
- **One primary action:** every surface has one obvious next step.
- **Suggestion is not commitment:** AI states and user commitments use different visual treatments.
- **Quiet density:** use compact metadata for practical kitchen information, not ornamental cards.
- **Warm precision:** combine human, conversational copy with exact quantities, times, and dates.
- **Accessible by construction:** focus, contrast, text alternatives, and reduced motion are part of each component.
- **Responsive by priority:** mobile keeps the decision path; desktop adds orientation and comparison space.

## Color Tokens

### Base palette

| Token | Value | Meaning |
|---|---|---|
| `ink-950` | `#101A2B` | Highest contrast text and dark surfaces |
| `ink-900` | `#17243A` | Primary text and navigation |
| `ink-700` | `#3D4B63` | Secondary text |
| `steel-400` | `#8793A5` | Placeholder and inactive metadata |
| `steel-200` | `#D5DAE1` | Borders and dividers |
| `paper-050` | `#F4F1E8` | Application canvas |
| `card-000` | `#FFFEFA` | Content surfaces |
| `white-000` | `#FFFFFF` | High-contrast controls |

### Semantic palette

| Token | Value | Meaning |
|---|---|---|
| `action-primary` | `#F45B3C` | Main action and urgent attention |
| `action-primary-dark` | `#A52F1C` | Hover, pressed, and strong border |
| `context-positive` | `#A7D46F` | Available, accepted, complete |
| `context-positive-dark` | `#5D7D36` | Positive text and icon contrast |
| `context-attention` | `#F3C969` | Pending, use soon, needs review |
| `context-attention-dark` | `#765816` | Attention text |
| `feedback-error` | `#D64545` | Error and blocked action |
| `feedback-info` | `#4C82C3` | Informational state |

Color is never the only status signal. Every semantic state combines color with text, icon, shape, or pattern.

## Typography

| Role | Typeface | Weight | Use |
|---|---|---:|---|
| Display | Fraunces | 600 | Page titles and short recommendation headlines |
| Body | Manrope | 400, 500, 600 | Navigation, controls, forms, explanations |
| Utility | IBM Plex Mono | 400, 500 | Quantities, time, dates, context metadata |

### Type scale

| Token | Size | Line height | Use |
|---|---:|---:|---|
| `display-lg` | 40 px | 1.1 | Desktop page thesis |
| `display-md` | 32 px | 1.15 | Mobile page title |
| `heading-lg` | 24 px | 1.2 | Section and recommendation title |
| `heading-md` | 20 px | 1.25 | Card and panel title |
| `body-lg` | 18 px | 1.45 | Introductory explanation |
| `body-md` | 16 px | 1.5 | Default interface copy |
| `body-sm` | 14 px | 1.45 | Secondary copy and form help |
| `utility-sm` | 12 px | 1.35 | Metadata, labels, status |

Use sentence case. Avoid all-caps paragraphs and do not use utility typography for essential instructions.

## Spacing, Shape, and Elevation

### Spacing scale

Use a four-point base scale:

```text
space-1  = 4 px
space-2  = 8 px
space-3  = 12 px
space-4  = 16 px
space-5  = 20 px
space-6  = 24 px
space-8  = 32 px
space-10 = 40 px
space-12 = 48 px
```

### Shape

- `radius-sm`: 6 px for inputs and compact controls.
- `radius-md`: 12 px for cards and drawers.
- `radius-lg`: 20 px for prominent recommendation surfaces.
- Pills are reserved for filters and short status labels, not full navigation.
- Use thin borders and alignment before adding shadows.

### Elevation

- `surface-0`: canvas, no shadow.
- `surface-1`: card, one subtle border or low shadow.
- `surface-2`: drawer and modal, stronger separation.
- `surface-3`: temporary toast or command feedback, highest local elevation.

## Layout Primitives

### App shell

- Desktop: navigation rail or top navigation plus focused content region.
- Mobile: compact header plus four-item bottom navigation.
- Content max width: 1200 px for operational screens.
- Reading max width: 720 px for instructions, rationale, and conversation.
- Page gutters: 20 px mobile, 32 px tablet, 48 px desktop.

### Stack and cluster

- **Stack:** vertical rhythm for forms, explanations, and screen sections.
- **Cluster:** horizontal grouping for metadata, filters, and action pairs.
- **Split:** main decision surface plus secondary context on desktop; collapses to stack on mobile.
- **Grid:** only for comparable items such as recipes or planner days; never use a grid when one action is clearly primary.

## Component Inventory

### App shell components

| Component | Responsibility | Primary states |
|---|---|---|
| `AppHeader` | Identity, page context, account menu | Default, compact, authenticated |
| `PrimaryNavigation` | Move between MVP destinations | Active, hover, focus, collapsed |
| `MobileNavigation` | Four-item mobile navigation | Active, unread-free, focus |
| `PageFrame` | Consistent title, actions, and content width | Default, loading |

### Decision components

| Component | Responsibility | Primary states |
|---|---|---|
| `ContextStrip` | Show ingredients, expiry, time, or preference context | Complete, partial, missing |
| `RecommendationCard` | Present one recommendation purpose and rationale | Requested, ready, accepted, rejected, unavailable |
| `RecommendationOption` | Allow evaluation of one cooking option | Proposed, selected, accepted, disabled |
| `ReasonLine` | Explain why a result is relevant | Positive, neutral, limitation |
| `CommitmentAction` | Separate accepted guidance from plan or shopping action | Available, pending, complete |

### Kitchen data components

| Component | Responsibility | Primary states |
|---|---|---|
| `PantryItemRow` | Display and edit ingredient state | Available, use soon, consumed, removed |
| `ExpiryIndicator` | Show expiry attention with text and date | Normal, soon, overdue, unknown |
| `RecipeCard` | Give enough recipe information to choose | Default, favorite, loading |
| `RecipeDetail` | Present instructions, ingredients, and context | Ready, missing context, error |
| `MealSlot` | Represent a daily or weekly planning choice | Empty, proposed, planned, completed |
| `ShoppingItemRow` | Review and complete a purchase intention | Open, completed, removed |

### Feedback and form components

| Component | Responsibility | Primary states |
|---|---|---|
| `TextField` | Capture bounded text input | Default, focus, error, disabled |
| `SelectField` | Choose category, unit, or period | Default, open, selected, error |
| `DateField` | Capture expiry or planning date | Default, invalid, unavailable |
| `PrimaryButton` | Execute the screen's main action | Default, hover, pressed, loading, disabled |
| `SecondaryButton` | Execute a supporting action | Default, focus, disabled |
| `DestructiveButton` | Remove or cancel an item | Default, confirm, loading |
| `EmptyState` | Explain missing data and invite action | First-use, filtered, completed |
| `ErrorState` | Explain recoverable failure | Retry, alternative, blocked |
| `Toast` | Confirm a completed local action | Success, info, error |
| `Drawer` | Edit or inspect focused content | Closed, open, saving |

## Component Contracts

### Context Strip

The Context Strip is the signature component of the product. It must:

- Show only context relevant to the current task.
- Use short labels and exact values.
- Distinguish authoritative facts from AI interpretation through copy.
- Show a limitation when context is incomplete.
- Wrap cleanly on mobile rather than truncating essential values.

Example:

```text
3 ingredients available  ·  1 use soon  ·  25 min
```

### Recommendation Card

Required content:

- Recommendation purpose.
- Option title.
- Preparation time or practical constraint.
- Available and missing ingredients where relevant.
- Short rationale.
- Limitation or uncertainty when relevant.
- One primary option action.

It must not show model name, prompt text, confidence score without product meaning, or provider error detail.

### Button hierarchy

- One `PrimaryButton` per focused surface wherever possible.
- `SecondaryButton` for navigation, clarification, or comparison.
- `DestructiveButton` only for removal, cancellation, or irreversible-looking action.
- Button labels name the result: `Save item`, `Accept option`, `Add to plan`, `Activate list`.
- Avoid vague labels such as `Submit`, `Continue` without context, or `Do it`.

### EmptyState

Every empty state contains:

1. A specific statement of what is empty.
2. Why the user may care.
3. One primary action.
4. One safe alternative when available.

### ErrorState

Every recoverable error contains:

- The failed action in plain language.
- What remains safe or unchanged.
- Retry, correction, or alternative action.
- No SQL, stack trace, prompt, provider, or internal component names.

## State Matrix

| State | Visual treatment | Copy behavior | Allowed action |
|---|---|---|---|
| Loading | Content-shaped placeholder | No false result claim | Navigate or cancel where safe |
| Empty | Calm illustration or text block | Explain value and next action | Add, explore, or generate |
| Validation error | Field border plus message | Name exact correction | Edit and retry |
| Permission denied | Neutral feedback surface | Do not reveal resource existence | Return to safe parent |
| Network error | Error surface with retry | Preserve local input | Retry or use offline-safe path |
| AI unavailable | Amber contextual state | Explain AI is unavailable | Browse, view pantry, retry |
| Accepted option | Positive state plus next-step action | Acceptance does not imply planning | Add to plan or shop explicitly |
| Completed | Positive confirmation | State exactly what changed | Continue to parent workflow |

## Accessibility Contract

- All interactive elements have visible keyboard focus.
- Focus order follows the visual reading order.
- Drawers and dialogs trap focus while open and return focus to the trigger when closed.
- Status colors meet contrast requirements and are paired with text or icons.
- Form fields have programmatic labels, hints, and error associations.
- Loading and save results are announced without interrupting unrelated work.
- Recommendation rationale and limitations are available to screen readers as text.
- Touch targets are at least 44 by 44 CSS pixels.
- Motion respects `prefers-reduced-motion`.
- The design remains usable at 200% text zoom.

## Responsive Contract

### Mobile

- Preserve the primary action and Context Strip above secondary content.
- Stack cards and form fields vertically.
- Use drawers or full-screen panels for editing.
- Keep bottom navigation visible unless a focused modal is active.

### Desktop

- Use split layouts for decision plus context.
- Support comparison of recommendation options without making them visually equal to commitment actions.
- Keep forms and recipe instructions within readable line lengths.
- Use hover as enhancement only; focus and click behavior must remain complete.

## Content Rules

- Use Indonesian-first copy for MVP.
- Use short active verbs: `Add item`, `Save changes`, `Accept option`, `Add to plan`.
- Explain system behavior, not implementation: `AI could not create a suggestion` instead of `Provider timeout`.
- Keep the same action name from button to confirmation.
- Treat uncertainty as useful information, not as a hidden failure.
- Never imply that AI is a person, authority, or owner of user data.

## Token and Component Governance

- New colors require a semantic use, contrast check, and design review.
- Screen-specific values should reference tokens rather than introduce arbitrary colors or spacing.
- Components are added when the behavior repeats across at least two MVP screens or represents a critical accessibility pattern.
- Changes to Recommendation, Context Strip, Button, Form, or State components require review against all consuming screens.
- Visual changes must not alter API ownership, authorization, or business commitment semantics.

## Exit Criteria

The design system foundation is ready for high-fidelity implementation when:

- Product and Design approve tokens and visual direction.
- Core component contracts cover all M3 wireframe screens.
- State, accessibility, responsive, and content rules are testable.
- Recommendation and commitment components preserve AI decision-support boundaries.
- No screen needs an unapproved visual or interaction primitive to proceed.

## Related Documents

- `docs/ux/ux-ui-design.md`
- `docs/ux/mvp-wireframes.md`
- `docs/ux/diagrams/design-system-component-map.mmd`
- `docs/architecture/api-design.md`
- `docs/architecture/authentication-authorization.md`
- `docs/architecture/ai-architecture.md`
