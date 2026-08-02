# DapurPintar AI High-Fidelity Screen Specification

## Document Control

| Item | Value |
|---|---|
| Milestone | M3 - UX/UI Design |
| Status | Draft - Awaiting Design Review |
| Scope | High-fidelity visual composition for the core MVP screens |
| Parent documents | `docs/ux/ux-ui-design.md`, `docs/ux/mvp-wireframes.md`, `docs/ux/design-system.md` |
| Related diagram | `docs/ux/diagrams/high-fidelity-screen-composition.mmd` |

## Visual Direction

The high-fidelity interface should look like a **modern kitchen ledger**: warm paper canvas, dark ink navigation, bright action marks, and compact evidence that explains why a suggestion is useful. It should feel more like a trusted working surface than a glossy food marketplace.

The signature visual move is the **Context Strip**. It sits directly above an AI recommendation or planning decision and behaves like a small annotated label:

```text
3 bahan tersedia  ·  1 segera dipakai  ·  25 menit
```

The strip makes the product's intelligence visible without exposing prompts or model language.

## Shared Composition Rules

- Canvas: `paper-050`.
- Primary content surface: `card-000` with a thin `steel-200` border.
- Primary text: `ink-900`.
- Primary action: `action-primary` with white label text.
- AI suggestion surfaces: `context-attention` tint, never the same color as a confirmed commitment.
- Accepted or completed surfaces: `context-positive` tint with `context-positive-dark` text.
- Page title uses Fraunces; all controls use Manrope.
- Quantity, time, and dates use IBM Plex Mono.
- Use one dominant surface per screen; secondary cards remain visually quieter.
- Avoid food photography as the main source of hierarchy. Content and context should lead.

## Screen A: Today

### Purpose

Make the next cooking decision feel immediate and grounded.

### Desktop composition

```text
  [nav rail]  Good evening, Sarah                         [profile]
              What are you cooking today?

              [3 available · 1 use soon · 25 min]
  +--------------------------------------+  +---------------------+
  | TRY THIS                             |  | YOUR WEEK            |
  | Tumis sayur telur                    |  | Tue 13               |
  | 20 min  ·  2 servings                |  | Dinner               |
  |                                      |  | + Choose a meal     |
  | Uses spinach first and fits your     |  |                     |
  | available time.                      |  | Pantry attention     |
  |                                      |  | 1 item use soon      |
  | [Use this idea]  [Ask about it]      |  | [Open pantry]        |
  +--------------------------------------+  +---------------------+
              Add one ingredient      Browse recipes
```

### Visual behavior

- The recommendation surface uses a thin `context-attention` top rule, not a full yellow card.
- The `ContextStrip` is a single line on desktop and wraps to two lines on mobile.
- “Try this” is a utility eyebrow, not a marketing headline.
- The primary action is `Use this idea`; secondary action is outlined.
- The right column never competes with the recommendation card in contrast or size.

### Mobile composition

```text
  Good evening, Sarah                 [user]
  What are you cooking today?

  [3 available · 1 use soon · 25 min]
  TRY THIS
  Tumis sayur telur
  20 min · 2 servings
  Uses spinach first and fits your time.
  [Use this idea]
  [Ask about it]

  Your week                     [Open planner]
  Tue 13 · Dinner · Choose a meal

  [ + Add pantry item ]
  Browse recipes
```

## Screen B: Recommendation Detail

### Purpose

Give the user enough evidence to accept one option with confidence.

### Composition

```text
  [Back to Today]                         Recommendation [Created]
  What fits your kitchen right now?

  [3 available · 1 use soon · 25 min]

  OPTION 01                              OPTION 02
  Tumis sayur telur                      Nasi goreng sayur
  20 min · Easy                          25 min · Easy
  [available] spinach, eggs               [available] rice, eggs
  [missing] garlic                        [missing] carrot

  Why this fits                           Why this fits
  Uses spinach first and stays            Uses pantry staples and is
  within your time preference.            easy to portion.

  [Accept option 01]                      [Choose option 02]

  [Ask about this recommendation]
```

### Visual behavior

- The selected option receives a `context-positive` edge only after the user selects it.
- The accepted state shows `Option accepted` and then offers two distinct actions: `Add to plan` and `Make a shopping list`.
- Do not use a confidence percentage. Use reason and limitation text instead.
- The conversation entry remains secondary to accepting or rejecting an option.

## Screen C: Pantry

### Purpose

Make inventory correction quick enough to become a habit.

### Composition

```text
  Pantry                                      [ + Add item ]
  Know what you have. Cook with less waste.

  [12 items] [3 categories] [1 use soon]
  [Search ingredients                         ] [Category v]

  USE SOON                                      [See all]
  +---------------------------------------------------------+
  | spinach              1 bunch             12 Aug         |
  | Best used in the next suggestion                         |
  +---------------------------------------------------------+

  ALL ITEMS
  Eggs                  6 pcs                 [Edit]
  Rice                  2 kg                  [Edit]
  Garlic                2 bulbs               [Edit]
```

### Visual behavior

- Use `context-attention` only for the use-soon section and its indicator.
- The item name is primary; quantity and date are utility metadata.
- The "Best used..." line is a product explanation, not an AI claim.
- The add drawer uses a white surface over the paper canvas with a clear save action.
- Empty Pantry uses a single illustrated ingredient mark and one primary action.

## Screen D: Recipe Detail

### Purpose

Provide a reliable cooking option independent of AI availability.

### Composition

```text
  [Back] Discover                         [Favorite outline]

  Tumis sayur telur
  A quick dinner for two
  20 min       Easy       2 servings

  [3 pantry ingredients available · 1 missing]

  Ingredients                              Instructions
  [available] spinach                       1. Prepare ...
  [available] eggs                          2. Heat ...
  [missing] garlic                          3. Add ...

  [Ask for a recommendation]        [Add to plan]
```

### Visual behavior

- Public recipe detail omits the personalized Context Strip.
- Authenticated detail may show authorized pantry relevance.
- Missing ingredients use `steel-400` plus a text label, never a red error state.
- Instructions use `body-lg` and a narrow reading width.
- `Favorite` is an icon plus accessible text label, not icon-only in the first interaction.

## Screen E: Planner

### Purpose

Turn a suggestion into a visible commitment without making the calendar feel like a spreadsheet.

### Desktop composition

```text
  Planner                    [<] 12 - 18 Aug [>]     [Week v]
  Plan the meals you actually want to cook.

  MON          TUE          WED          THU          FRI
  12           13           14           15           16
  +--------+   +--------+   +--------+   +--------+   +--------+
  | Dinner |   | Dinner |   | Dinner |   | Dinner |   | Dinner |
  | + add  |   | recipe |   | + add  |   | + add  |   | recipe |
  +--------+   +--------+   +--------+   +--------+   +--------+

  [Generate shopping list from this week]
```

### Visual behavior

- Empty slots are outlined opportunities, not gray disabled blocks.
- Planned meals use `card-000`; completed meals use a subtle `context-positive` edge.
- The week selector is a compact utility control, not a hero element.
- Drag and drop may be added later; every move must have a menu or edit alternative.
- Shopping generation is a bottom action because it is downstream from planning.

## Screen F: Shopping List

### Purpose

Make purchase intent reviewable before the user leaves for the store.

### Composition

```text
  Shopping list                              [New list]
  Week of 12 Aug · Draft
  Generated from weekly plan + pantry

  REVIEW BEFORE SHOPPING
  +---------------------------------------------------------+
  | [ ] garlic                 2 bulbs                    ...|
  | [ ] eggs                   6 pcs                      ...|
  | [ ] soy sauce              1 bottle                   ...|
  +---------------------------------------------------------+
  [Activate list]

  COMPLETED
  [x] rice                    2 kg
```

### Visual behavior

- Draft and Active use distinct labels; do not communicate lifecycle through color only.
- The `Activate list` action is primary while the list is Draft.
- Completion uses a calm strike-through and `context-positive`, never a celebratory animation that slows shopping.
- Source context is shown as a small note and remains read-only.
- A completed Shopping Item has no "Add to pantry" prompt.

## Screen G: AI Degraded State

### Purpose

Preserve user momentum when an external AI dependency is unavailable.

### Composition

```text
  We could not make a suggestion right now.
  Your pantry and recipes are still available.

  [Try again]
  [Browse recipes]
  [Open pantry]

  No recommendation was created.
```

### Visual behavior

- Use a muted `feedback-info` or `context-attention` surface, not a red alarm.
- Explain the product consequence: no recommendation was created.
- Never mention provider names, timeout values, prompts, or retry internals.
- Preserve the previous screen's navigation and user-entered data.

## Motion and Transitions

- Page entry: no global animation; content appears stable.
- Recommendation result: Context Strip fades in before the recommendation body.
- Save action: button transitions to a compact confirmation state for 1.5 seconds.
- Option acceptance: selected edge and status text update; no confetti.
- Shopping completion: row settles into the Completed section.
- Reduced motion: replace all movement with immediate state changes.

## Visual QA Checklist

- Does the primary action remain visually dominant at 320 px width?
- Is the Context Strip readable when it wraps?
- Can users distinguish a suggestion from an accepted option and a planned meal?
- Are status colors paired with text or icons?
- Does every screen have a visible keyboard focus state?
- Are public recipe screens free of private context?
- Does AI degradation offer a useful non-AI path?
- Are empty and error states as intentional as ready states?
- Do cards remain content-led rather than image-led?

## Exit Criteria

This high-fidelity specification is ready for prototype or implementation when:

- Product and Design approve the visual direction and screen hierarchy.
- The primary Today-to-Plan-to-Shop journey is visually coherent.
- All high-priority component states are represented.
- Responsive and accessibility checks pass at design review.
- No screen implies automatic mutation or authority beyond the approved domain and API contract.

## Related Documents

- `docs/ux/ux-ui-design.md`
- `docs/ux/mvp-wireframes.md`
- `docs/ux/design-system.md`
- `docs/ux/diagrams/high-fidelity-screen-composition.mmd`
- `docs/architecture/api-design.md`
- `docs/architecture/ai-architecture.md`
