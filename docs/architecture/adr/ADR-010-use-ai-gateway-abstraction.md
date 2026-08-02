# ADR-010 Use AI Gateway Abstraction for OpenAI Integration

- Status: Accepted
- Date: 2026-08-02

## Context

AI is central to the product's value proposition: recipe recommendation, pantry analysis, AI chat, and future meal planning and nutrition assistance. OpenAI is an external dependency with cost, latency, quota, availability, policy, and prompt-injection risks. The architecture vision requires AI to be grounded in authorized product context while keeping the domain independent from provider details.

## Decision

Introduce an AI Gateway / Provider Adapter abstraction owned by the AI Assistant application boundary. The gateway will mediate model and prompt selection, context minimization, request validation, structured response validation, timeout and retry policy, usage and cost metadata, safety controls, and provider-specific errors. OpenAI is the initial provider implementation; core domains will depend on the gateway contract rather than OpenAI APIs or SDK types.

## Consequences

- OpenAI credentials, SDK details, model configuration, and provider failures remain outside the domain and core use cases.
- A second provider or model can be evaluated or introduced without rewriting product modules.
- AI calls become a consistent place to enforce privacy, prompt-injection defenses, output validation, quotas, and observability.
- The gateway adds an abstraction and requires contract tests, AI evaluations, prompt/version management, and careful context assembly.
- Provider-independent behavior cannot hide genuine model differences; capability and quality limits must remain explicit.
- AI remains decision support, not the authoritative source of pantry, recipe, meal-plan, or shopping data; user confirmation and domain validation remain necessary.

## Alternatives Considered

- **Direct OpenAI calls from each feature:** Faster for a prototype, but duplicates policy and error handling and creates deep vendor coupling.
- **OpenAI-only domain integration:** Simple while one provider is used, but conflicts with long-term maintainability and the documented vendor-risk mitigation.
- **Build a full multi-provider platform immediately:** More flexibility, but premature for MVP; the gateway creates the seam without adding unused provider complexity.
