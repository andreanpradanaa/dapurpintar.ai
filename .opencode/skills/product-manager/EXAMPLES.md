# Product Manager Examples

## Story with full behavior

```text
As Sarah, I want to accept a recommended meal so that it becomes my plan
without me re-entering ingredients.

AC-1: Accepting converts the proposal into a confirmed plan (testable).
AC-2: Acceptance records the meal and date server-side.
AC-3: If the AI proposal is invalid, acceptance is rejected with a clear error.
AC-4: Accepting does not auto-purchase or auto-add pantry items.
```

## Scope decision record

```text
Deferred: multi-household sharing.
Why: serves neither Sarah nor Daniel in MVP; adds ownership complexity
conflicting with single-profile scope. Revisit after M16.
```

## Gate check

```text
M8 AI Foundation gate: do Sarah's and Daniel's AI surfaces meet the
accuracy and safety criteria in the evaluation rubric? If not, gate stays open.
```
