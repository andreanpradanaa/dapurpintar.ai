# Frontend Rules

> Frontend rules for DapurPintar AI. Complement `.opencode/AI_RULES.md` and the `frontend-nextjs` skill.

## Stack

- Next.js application; component and route conventions follow the M3 design system.
- All data flows through the typed API client generated from `docs/api/openapi.yaml`.
- Never invent fields, endpoints, or error codes outside the contract.

## Design system

- Use the M3 design tokens and components; no ad-hoc styling for standard elements.
- Spacing, color, and type come from the token set.
- Promote repeated patterns to components instead of copying.

## State and resilience

- Server-render the shell; hydrate with explicit loading states.
- Empty and error states are designed, not afterthoughts.
- AI-dependent surfaces degrade gracefully when recommendations are unavailable.
- Never cache sensitive personal context in client-side storage.

## Authentication

- Session cookie auth is used; tokens are not stored in client storage.
- No credentials or provider keys are sent from the client.
- All mutations go through the API, never direct provider access.

## Accessibility

- Keyboard, focus, and screen-reader behavior verified on key flows.
- Color contrast and touch targets follow the design system.
- Form errors are announced and linked to the field.

## Quality

- Client changes that assume backend shapes are reconciled with the OpenAPI contract.
- Components reviewed against the design system and acceptance criteria.
