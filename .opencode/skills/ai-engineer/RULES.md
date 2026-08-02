# AI Engineer Rules

## Gateway boundary

- OpenAI and any provider are accessed only through the AI Gateway.
- Business modules depend on the Gateway port, never on a provider SDK.
- Provider credentials are deployment secrets and never appear in code or logs.

## AI as decision support

- AI output is a proposal until the user accepts or confirms it.
- AI never silently mutates business state: no auto-planned meals, no auto-pantry changes, no auto-purchase confirmation.
- Generated guidance stays distinguishable from user-confirmed decisions.

## Prompts and policies

- Prompts and policies are versioned; a change is a new version, not an edit to a live prompt.
- Structured output must validate against a schema; invalid or unsafe output is rejected, not silently accepted.
- Prompt history never stores raw provider prompts or payloads in business tables (M4-DEC-013).

## Safety and reliability

- Provider failures are safe dependency failures; core non-AI operations remain usable.
- Timeouts and retries are bounded; retry loops never create cost spikes or duplicate business commands.
- Prompt injection, unsafe content, unsupported claims, and data leakage are defended at the gateway.
- Missing context reduces certainty; the system must declare limitations rather than invent facts.

## Cost and quota

- AI quota and cost budgets are enforced and alerted (M4-DEC-016).
- AI-intensive endpoints are rate-limited at the API boundary.

## Evaluation

- AI quality is measured with the approved evaluation dataset and rubric (M4-DEC-012).
- Acceptance, rejection, and rationale are preserved as durable business records for measurement.
- Evaluation records are separated from business data and subject to privacy review.

## Privacy

- AI never receives authentication credentials or raw session data.
- Personal context is minimized and used only for a relevant business decision.
- Conversation context is retained only for the active Recommendation per the retention policy.
