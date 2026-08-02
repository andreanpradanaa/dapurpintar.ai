# DapurPintar AI M3 Usability Validation Plan

## Document Control

| Item | Value |
|---|---|
| Milestone | M3 - UX/UI Design |
| Status | Draft - Awaiting Product Review |
| Scope | Validation of MVP UX flows, screen hierarchy, visual direction, and interaction language |
| Parent documents | `docs/ux/ux-ui-design.md`, `docs/ux/mvp-wireframes.md`, `docs/ux/design-system.md`, `docs/ux/high-fidelity-screen-spec.md` |
| Related diagram | `docs/ux/diagrams/usability-validation-journey.mmd` |

## Purpose

This plan defines how DapurPintar AI will validate the MVP experience before frontend implementation. The goal is not to prove that users like the visual style. The goal is to learn whether people can understand the product, trust the recommendation boundary, and complete the core kitchen decision with reasonable effort.

The validation focuses on the primary journey:

```text
Set context -> Understand recommendation -> Accept one option
-> Add to plan or create shopping list explicitly
```

## Research Questions

### Core value

- Do users understand that DapurPintar AI helps them decide what to cook, not only browse recipes?
- Can users identify the next useful action on Today without exploring every destination?
- Does the Context Strip explain why a recommendation appeared?

### Trust and AI boundaries

- Do users understand the difference between a suggestion, an accepted option, a planned meal, and a shopping item?
- Do users know what the AI used as context?
- Do users understand that accepting a recommendation does not automatically create a Meal Plan or Shopping List?
- Can users recover when AI is unavailable without assuming their pantry data was lost?

### Workflow usability

- Can users add enough pantry context without feeling forced into a long setup?
- Can users select one Recommendation Option confidently?
- Can users plan a meal and generate a shopping list as two intentional steps?
- Can users complete a shopping item without expecting Pantry to change automatically?

### Visual and content clarity

- Does the kitchen ledger visual language feel calm, practical, and trustworthy?
- Is the primary action visually clear on mobile and desktop?
- Are time, quantity, expiry, and missing ingredients readable at a glance?
- Is Indonesian-first copy understandable without technical or AI jargon?

## Participant Profile

### Primary participants

Recruit participants who match the product's primary target users:

- Adults aged approximately 25-45 in Indonesia.
- Smartphone-first users.
- People who cook at home at least two times per week.
- People who have recently decided what to cook, planned meals, or made a grocery list.

### Segment coverage

| Segment | Minimum representation | Main focus |
|---|---:|---|
| Young families | 3 participants | Time pressure, household context, food waste |
| Busy professionals | 3 participants | Speed, simple recipes, shopping clarity |
| Health-conscious users | 2 participants | Trust and avoiding unsupported nutrition claims |

The health-conscious group validates comprehension and boundaries only. Nutrition features are not tested as MVP functionality.

### Exclusions

- Do not recruit people who designed the prototype.
- Do not use production personal data.
- Do not ask participants to share passwords, private health information, or real pantry records beyond what they freely choose to describe.

## Prototype Scope

The prototype should contain these screens:

- Registration and login.
- Profile setup.
- Today.
- Pantry and add/edit item.
- Discover and Recipe Detail.
- Recommendation Detail.
- Recommendation Conversation.
- Planner.
- Shopping List.
- AI Degraded State.

The prototype does not need working backend integration. It must faithfully represent:

- Loading, empty, error, permission, and degraded-AI states.
- Recommendation Option acceptance.
- Explicit Add to Plan and Generate Shopping List actions.
- Public versus authenticated recipe context.
- Mobile and desktop layouts.

## Test Method

### Moderated task session

- Duration: 30-45 minutes.
- Device: participant's smartphone first; desktop follow-up where available.
- Format: remote or in-person moderated session.
- Think-aloud: use lightly; do not interrupt the first attempt at a task.
- Prototype: one realistic account context with seeded pantry items and two recommendation options.
- Recording: only with explicit consent; redact names and private details from notes.

### Session structure

1. Explain that the prototype, not the participant, is being tested.
2. Ask about the participant's normal cooking decision and shopping habit.
3. Run the task scenarios without teaching the interface first.
4. Ask comprehension and trust questions after each critical task.
5. Compare mobile and desktop observations where possible.
6. Ask for final confidence, friction, and missing information.

## Task Scenarios

### Task 1: First decision

**Prompt:**

> Kamu pulang kerja dan punya waktu sekitar 25 menit. Buka DapurPintar AI dan cari tahu apa yang bisa kamu masak hari ini.

**Observe:**

- Where the participant looks first.
- Whether Today communicates a useful next action.
- Whether context is understood without explanation.
- Whether the participant can distinguish recommendation from recipe browsing.

**Success:** Participant reaches a useful recommendation or intentionally chooses Browse Recipes within two minutes.

### Task 2: Add pantry context

**Prompt:**

> Tambahkan telur dan bayam ke pantry. Bayam sebaiknya digunakan dalam waktu dekat.

**Observe:**

- Whether Add Item is discoverable.
- Whether quantity and expiry fields are understandable.
- Whether the user knows what was saved.
- Whether expiry attention is perceived as guidance rather than an error.

**Success:** Participant saves both items with correct meaning and can find them again.

### Task 3: Evaluate a recommendation

**Prompt:**

> Lihat rekomendasi yang diberikan. Pilih opsi yang paling cocok untuk waktumu dan jelaskan kenapa kamu memilihnya.

**Observe:**

- Whether the Context Strip is read.
- Whether available and missing ingredients are understood.
- Whether rationale and limitations support trust.
- Whether the participant knows how to ask a bounded question.

**Success:** Participant selects one Recommendation Option and can state at least one reason for the choice.

### Task 4: Make an explicit plan

**Prompt:**

> Jadwalkan opsi yang kamu pilih untuk makan malam besok.

**Observe:**

- Whether Add to Plan is distinct from Accept Option.
- Whether the date or meal occasion is confirmed.
- Whether the participant believes the plan changed Pantry or Shopping automatically.

**Success:** Participant creates the intended Planned Meal and can explain what changed.

### Task 5: Generate a shopping list

**Prompt:**

> Buat daftar belanja dari rencana minggu ini, lalu periksa item sebelum mengaktifkannya.

**Observe:**

- Whether the source of generated items is clear.
- Whether the user reviews instead of blindly activating.
- Whether removing or adjusting an item is discoverable.

**Success:** Participant generates, reviews, and activates the intended list without assuming that completion changes Pantry.

### Task 6: Recover from AI failure

**Prompt:**

> Anggap rekomendasi AI tidak tersedia. Lanjutkan aktivitas memasak tanpa kehilangan data pantry.

**Observe:**

- Whether the failure message is understandable.
- Whether the participant chooses Browse Recipes or Pantry.
- Whether the participant believes their data was lost.

**Success:** Participant reaches a non-AI alternative and retains confidence in existing data.

## Measurement Plan

### Task metrics

| Metric | Target or interpretation |
|---|---|
| Task completion | At least 6 of 8 participants complete core tasks without moderator rescue |
| Time to choose menu | Median under 2 minutes after context is available |
| Recommendation comprehension | At least 6 of 8 can explain why an option fits |
| Commitment distinction | At least 7 of 8 distinguish acceptance from planning and shopping |
| Pantry correction | At least 7 of 8 save an item with correct quantity and expiry meaning |
| AI recovery | At least 6 of 8 reach a non-AI alternative without assuming data loss |
| Critical error count | Zero errors that create or appear to create an unintended commitment |
| Confidence | Median post-task confidence of at least 4 on a 5-point scale |

These targets are design-review thresholds, not product success claims. They should be adjusted if participant recruitment or prototype fidelity makes a metric invalid.

### Qualitative signals

Capture quotes and observations for:

- Trust or distrust of the recommendation rationale.
- Confusion around `Accept option`, `Add to plan`, and `Make a shopping list`.
- Perceived effort of pantry setup.
- Interpretation of expiry attention.
- Emotional response to an empty or AI-degraded state.
- Whether the visual style feels useful or merely decorative.

## Comprehension Questions

Ask after the relevant task, without leading:

- What did the recommendation use to make this suggestion?
- What changed when you accepted this option?
- What would you expect to happen after adding it to the plan?
- Which ingredients will be added to the shopping list, and why?
- What happens to the pantry when you complete a shopping item?
- What would you do if the AI suggestion could not be created?

Correct answers should align with the domain boundaries, not with implementation details.

## Heuristic Review

Before participant testing, review the prototype against these heuristics:

- Visibility of current context and system status.
- Match between interface language and everyday cooking language.
- User control and easy recovery.
- Consistency of action labels across screens.
- Prevention of accidental commitments.
- Recognition rather than memory.
- Support for both novice and time-pressured users.
- Minimal, actionable error messages.
- Accessibility of focus, contrast, labels, and touch targets.
- Respect for AI uncertainty and user decision authority.

## Decision Rules

### Keep

Keep an interaction when participants complete it, understand its consequence, and it supports the core decision without unnecessary friction.

### Revise

Revise when:

- Two or more participants make the same non-critical mistake.
- Participants need repeated explanation of a label or state.
- The visual hierarchy hides the primary action.
- Users understand the action but not its consequence.
- Mobile behavior causes avoidable scrolling or control loss.

### Block

Block high-fidelity implementation or frontend build when:

- Participants believe AI acceptance automatically changes a plan, shopping list, or pantry.
- Participants cannot identify what context informed a recommendation.
- A primary task requires moderator rescue for most participants.
- An error or empty state causes users to believe their data was deleted.
- Public recipe views expose private context in the prototype.

## Research Output

After each validation round, produce:

- Participant summary without identifying personal details.
- Task completion and timing table.
- Critical confusion and trust observations.
- Screens or components requiring revision.
- Decisions kept, revised, or blocked.
- Updated open questions and next experiment.

The output should update the parent UX documents and create product issues for changes that affect API, domain, or scope decisions.

## Ethical and Privacy Rules

- Obtain informed consent before recording or quoting participants.
- Use synthetic accounts and pantry data.
- Do not request passwords, private medical information, or provider credentials.
- Remove identifying details from notes and screenshots.
- Do not use participant content as AI evaluation data without separate consent and policy.
- Allow participants to stop at any time.

## Exit Criteria

M3 usability validation is ready to close when:

- At least one round of representative usability testing is complete.
- Critical commitment and AI-boundary misunderstandings are resolved.
- Primary mobile journey meets the agreed usability thresholds or has an explicit remediation plan.
- High-fidelity screens and components are updated from evidence.
- Product, Design, and Architecture agree on changes that affect scope or contracts.
- The resulting UX is ready to guide M10 frontend implementation.

## Related Documents

- `docs/ux/ux-ui-design.md`
- `docs/ux/mvp-wireframes.md`
- `docs/ux/design-system.md`
- `docs/ux/high-fidelity-screen-spec.md`
- `docs/ux/diagrams/usability-validation-journey.mmd`
- `docs/product/user-personas.md`
- `docs/product/success-metrics.md`
- `docs/architecture/api-design.md`
- `docs/architecture/ai-architecture.md`
