# AI Rules

> AI-specific rules for DapurPintar AI. These complement `.opencode/AI_RULES.md` and the `ai-engineer` skill.

## Gateway boundary

- OpenAI (and any provider) is reached only through the AI Gateway port.
- Business modules depend on the port, never on a provider SDK.
- Provider credentials are deployment secrets, never code or logs.

## AI is decision support

- AI output is a proposal until the user accepts or confirms it.
- AI never silently commits business state: no auto-planned meals, no auto-pantry changes, no auto-purchase confirmation.
- Generated guidance stays distinguishable from user-confirmed decisions.

## Versioning and validation

- Prompts, safety policy, and structured-output schemas are versioned.
- A change to a live behavior is a new version, not an in-place edit.
- Structured output validates against a schema; invalid or unsafe output is rejected.

## Reliability and cost

- Provider failures are safe dependency failures; core operations stay usable.
- Timeouts and retries are bounded; no cost spikes or duplicate commands.
- Quota and cost budgets are enforced and alerted (M4-DEC-016).
- AI-intensive endpoints are rate-limited.

## Privacy and retention

- AI never receives auth credentials or raw session data.
- Conversation context is retained only for the active Recommendation (M4-DEC-013).
- Raw provider prompts or payloads are never stored in business tables.

## Evaluation

- AI quality is measured with the approved dataset and rubric (M4-DEC-012).
- Acceptance, rejection, and rationale are preserved as durable records.
- The model must declare low confidence rather than invent facts.
