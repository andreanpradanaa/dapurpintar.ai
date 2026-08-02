---
name: frontend-nextjs
description: Use when implementing or reviewing the DapurPintar AI frontend. Covers the Next.js application, the M3 design system, API integration against /api/v1, session auth, and accessibility. Trigger on Next.js, React, component, UI, design system, styling, or API client work.
---

# Frontend Next.js Skill

## Purpose

Build the DapurPintar AI client against the approved API contract and design system so it is consistent, accessible, and resilient to backend degradations.

## Responsibilities

- Implement UI in the Next.js app using the design system components.
- Consume `/api/v1` through a typed API client.
- Handle session authentication and shared layouts.
- Show accurate loading, empty, and error states.

## Inputs

- `docs/api/openapi.yaml` for the request/response shapes.
- `docs/api/m6-error-catalog.md` for user-facing error handling.
- `docs/api/` auth and session contract for login flows.
- `docs/architecture/` persona docs for Sarah and Daniel.
- `.opencode/rules/frontend.md` for conventions.

## Outputs

- Component code using the design system.
- Typed API client and error handling.
- Routes, layouts, and state for the milestone features.

## Dependencies

- The API contract is authoritative; the client must not invent shapes.
- The M3 design system tokens drive styling.

## Status

Active - planned for the M10 frontend milestone.
