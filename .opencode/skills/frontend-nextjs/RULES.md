# Frontend Next.js Rules

## API integration

- All data flows through the typed API client generated from `docs/api/openapi.yaml`.
- The client never invents fields, endpoints, or error codes not in the contract.
- Backend errors map to the M6 envelope; stable codes drive user-facing messages.

## Design system

- Use the M3 design system tokens and components; no ad-hoc styling for standard elements.
- Spacing, color, and type come from the token set.
- Reusable patterns are promoted to components, not copied.

## State and resilience

- Server-render the shell; hydrate with explicit loading states.
- Empty and error states are designed, not afterthoughts.
- AI-dependent surfaces degrade gracefully when the recommendation is unavailable.
- Never cache sensitive personal context in client-side storage.

## Accessibility

- Keyboard, focus, and screen-reader behavior is verified on key flows.
- Color contrast and touch targets follow the design system.
- Form errors are announced and linked to the field.

## Security

- Session cookie auth is used; tokens are not stored in client storage.
- No credentials or provider keys are sent from the client.
- All mutations go through the API, never direct provider access.

## Reviews

- Components are reviewed against the design system and acceptance criteria.
- Client changes that assume backend shapes must be reconciled with the OpenAPI contract.
