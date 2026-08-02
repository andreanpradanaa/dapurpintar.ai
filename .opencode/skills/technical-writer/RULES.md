# Technical Writer Rules

## Structure

- One idea per document; use headings that mirror the document's purpose.
- Put a short status and purpose block at the top of each document.
- Use tables for facts, prose for rationale, code blocks for examples.

## Accuracy

- Documents describe the approved state, not an aspirational one.
- Code examples match the actual API and backend layout.
- If a document contradicts a decision record, the decision record wins and the doc must be fixed.

## Cross-referencing

- Link decisions by their register IDs (M4-DEC-XXX).
- Link artifacts: OpenAPI, schema, backend package, roadmap.
- Reference files by repository-relative paths so links survive moves.

## Consistency

- Terminology is stable across docs (expiry, not expiration; Recommendation, not suggestion).
- Milestone names match the milestone list.
- Statuses match the review protocol (In Review, Approved, Complete).

## Review

- Flag broken references and stale content in review.
- Prefer small, focused edits over rewrites.
- Update dependent documents in the same change as the source change.
