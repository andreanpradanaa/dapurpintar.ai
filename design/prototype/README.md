# DapurPintar AI M3 Prototype

This is a static, browser-only prototype for M3 UX validation. It does not call the API, persist data, or represent production frontend architecture.

## Run

Open `index.html` in a browser, or serve the repository with any static file server.

## Validation Paths

- Today -> accept a recommendation option -> add it to Planner or Shopping.
- Pantry -> add an ingredient.
- Planner -> add a meal and generate a Shopping List.
- Shopping -> complete an item without changing Pantry.
- Today -> simulate AI unavailable -> browse a non-AI alternative.

The prototype intentionally keeps AI acceptance, meal planning, shopping, and pantry changes as separate actions.
