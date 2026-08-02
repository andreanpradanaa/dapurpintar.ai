# DapurPintar AI Architecture

## Document Control

| Item | Value |
|---|---|
| Status | Proposed for M2 review |
| Scope | MVP AI-assisted kitchen decision support |
| Primary audience | Product, backend, frontend, AI, QA, DevOps, and AI-assisted development |
| Related diagram | `docs/architecture/diagrams/ai-architecture-flow.mmd` |

## Executive Summary

DapurPintar AI uses AI as an application capability for trustworthy kitchen decision support. The AI capability assembles authorized product context, requests assistance from an external provider through an internal AI Gateway, validates the result, and returns a product-level Recommendation. It does not become the authority for Pantry, Recipe, Meal Plan, Shopping List, User Profile, or Account data.

The MVP uses OpenAI as the initial external provider behind a provider-neutral gateway boundary. The gateway centralizes provider credentials, model and prompt selection, context minimization, structured output validation, safety controls, timeout and retry behavior, usage and cost metadata, and provider failure translation. Core domains depend on application contracts and product vocabulary rather than OpenAI APIs, SDK types, or raw provider responses.

Every AI response is treated as a proposal until the user explicitly accepts a Recommendation Option and the owning business context validates a later planning or shopping command. PostgreSQL remains authoritative for durable recommendation and conversation context. Redis may support safe, bounded caching and coordination but never becomes the source of truth.

## Architectural Goals

- Deliver relevant, practical, and understandable kitchen guidance.
- Ground AI requests in authorized User Profile, Pantry, Recipe, Meal Plan, and Shopping context.
- Keep user decision authority explicit for recommendations, meal plans, and shopping commitments.
- Isolate provider credentials, SDKs, model behavior, and provider errors from the domain.
- Validate generated output before it reaches a product response or business aggregate.
- Minimize sensitive context sent to an external provider.
- Make prompts, models, policies, cost, latency, quality, and failures observable and reproducible.
- Support provider replacement or additional capabilities without rewriting core business modules.
- Preserve graceful degradation when AI is unavailable.

## Non-Goals

The MVP does not include:

- A multi-provider routing platform.
- Autonomous meal-plan, shopping-list, or pantry mutations.
- AI as an authoritative source of ingredient, recipe, nutrition, or account truth.
- General-purpose unrestricted chatbot behavior.
- RAG infrastructure, vector search, fine-tuning, voice, image recognition, or computer vision.
- Nutrition coaching or medical advice.
- Independent long-term AI Conversation ownership; Recommendation Conversation remains recommendation-scoped in the MVP.

These capabilities may be introduced later through explicit product, domain, safety, and architecture decisions.

## AI Domain Boundary

### Core capability

AI-Assisted Kitchen Decision Support owns:

- Recommendation intent and purpose.
- Context relevance and limitation explanation.
- Recommendation Options.
- Recommendation lifecycle and acceptance state.
- Recommendation-scoped Conversation context in the MVP.
- Product-level safety and suitability decisions for AI assistance.

It does not own:

- Pantry availability or expiry truth.
- Recipe meaning or cooking instruction truth.
- User identity, profile ownership, or declared preferences.
- Meal commitments or scheduled meals.
- Shopping commitments or purchase completion.
- Nutrition policy or medical authority.

### Supporting context relationships

The AI capability is downstream from the meaning supplied by:

- **Identity and Access:** trusted user and authorization scope.
- **User Context and Preferences:** personal preferences, constraints, and goals.
- **Pantry Management:** available and expiring ingredients.
- **Culinary Knowledge and Recipe Experience:** recipes and cooking guidance.
- **Meal Planning:** existing or proposed planning context where authorized.
- **Shopping Optimization:** existing purchase intent where relevant.

The AI capability supplies guidance to:

- **Meal Planning:** an accepted option may become a planning input after an explicit user command.
- **Shopping Optimization:** an accepted option or plan may inform purchase generation after an explicit user command.

No integration grants the AI capability write authority over an upstream or downstream aggregate.

## AI Use Cases

### Kitchen Recommendation

`POST /api/v1/recommendations` requests a contextual Recommendation. The request identifies a product purpose and user intent. The application assembles only authorized context, sends a bounded request through the gateway, validates the result, and persists a Recommendation with its Options, rationale, limitations, and lifecycle state.

The recommendation may be presented to the user, but it remains a proposal. The user accepts a specific Recommendation Option through the API before it can guide a later Meal Plan or Shopping List operation.

### Recommendation Conversation

The MVP Conversation is a bounded child context of one Kitchen Recommendation. It allows the user to clarify or understand the recommendation without becoming an independent general chat history. Conversation turns must remain within the original recommendation purpose and authorized context.

Conversation messages must not:

- Change Pantry, Recipe, Meal Plan, Shopping List, or Profile state implicitly.
- Override a domain rule or authorization decision.
- Expose system prompts, provider payloads, credentials, or internal policy text.
- Be used as an unbounded memory store without retention and privacy justification.

### Pantry Analysis

`POST /api/v1/ai/pantry-analysis` analyzes the authorized Pantry context for use-first opportunities or other approved decision support. Pantry Management remains authoritative for availability, quantity, and expiry. The result may identify an opportunity but cannot consume, remove, or modify a Pantry Item.

### Future meal suggestions

`POST /api/v1/ai/meal-suggestions` remains a future surface. It may return candidate daily or weekly guidance but must not create a Meal Plan. It becomes an MVP capability only after explicit product scope promotion and safety validation.

## AI Request Lifecycle

The conceptual lifecycle for an AI-assisted Recommendation is:

```text
Requested -> Context Authorized -> Context Assembled -> Provider Requested
          -> Output Received -> Output Validated -> Recommendation Created
          -> Presented -> Accepted or Rejected
```

Failure or interruption may result in `Unable to Complete`. A provider response is never treated as a successful Recommendation until product and schema validation succeeds.

The lifecycle separates technical request progress from business Recommendation state. Technical retries or provider attempts must not create duplicate business Recommendations when the same client operation is retried with the same idempotency key.

## Context Assembly

### Authorization first

Context assembly begins only after authentication and resource authorization. The application derives User ownership from the authenticated session and requests business meaning through approved application ports. Client-provided identifiers are references to resolve, not proof of access.

### Purpose-specific context

The context assembler selects data according to the request purpose. A recommendation request may use:

- Relevant User Profile preferences and constraints.
- Current Pantry availability and expiry signals.
- Approved Recipe candidates and cooking guidance.
- Existing Meal Plan context when relevant.
- Existing Shopping context when relevant.
- Conversation context limited to the owning Recommendation.

The assembler must prefer current authoritative facts and clearly label historical or generated snapshots. It must not send unrelated personal data merely because it is available.

### Context minimization

Before provider submission, the gateway boundary applies:

- Field allowlisting by use case.
- User and household scope checks.
- Removal or masking of identifiers not required for the task.
- Bounded collection sizes and message lengths.
- Explicit source labels for facts, preferences, and generated proposals.
- A clear instruction that provider output is not authoritative product state.

The assembled context is an internal request representation. It is not the domain model and must not leak provider-specific structure into domain aggregates.

## AI Gateway Boundary

The AI Gateway is the single application boundary for external AI assistance. It provides a provider-neutral capability contract for:

- Request purpose and capability selection.
- Prompt and policy revision selection.
- Model selection under approved configuration.
- Structured input construction.
- Provider invocation.
- Timeout, retry, and cancellation behavior.
- Provider error translation.
- Output schema validation.
- Safety and policy result handling.
- Usage, latency, and cost metadata.

The gateway must not own Pantry, Recipe, Meal Plan, Shopping List, Profile, or Account rules. It receives a minimized request and returns a validated provider-independent result or a safe failure.

### Provider adapter

OpenAI is the initial provider adapter. The adapter owns:

- OpenAI credentials and SDK details.
- Provider request and response mapping.
- Provider-specific model and capability constraints.
- Provider-specific error classification.
- Provider transport behavior.

The adapter must not be called directly by HTTP handlers, domain aggregates, or unrelated modules. A future provider can be added only when capability, quality, safety, cost, and operational differences are explicitly evaluated.

## Prompt and Policy Management

Prompts are controlled product and safety artifacts, not arbitrary strings owned by feature handlers.

Each AI request records the metadata needed for reproducibility:

- Use-case or capability purpose.
- Prompt revision identifier.
- Safety and context-policy revision.
- Model identifier and provider.
- Input and output schema revision.
- Request and completion timestamps.
- Usage and cost metadata where available.

Stored metadata must not contain provider credentials or unnecessary raw personal data. Prompt text and conversation retention require a documented privacy and evaluation purpose. Prompt changes must be reviewable and evaluated against representative scenarios before promotion.

## Structured Output and Validation

AI features that influence product decisions use structured output contracts. Validation happens in layers:

1. **Transport validation:** provider response exists and is within bounded size and time limits.
2. **Schema validation:** required fields, types, limits, and enumerations are valid.
3. **Source validation:** referenced Recipe, Pantry, Meal Plan, or Shopping facts are recognized and authorized.
4. **Business validation:** the owning domain accepts the proposed state and preserves aggregate invariants.
5. **Safety validation:** unsupported claims, unsafe instructions, privacy leakage, prompt-injection output, and prohibited content are rejected or safely degraded.
6. **Presentation mapping:** only product-level fields are returned to the client.

Invalid output is not silently repaired into a business fact. The system may retry a bounded provider request, request a safer fallback, or return `Unable to Complete` with a safe explanation.

## Safety and Trust

### Prompt injection

User-authored recipe notes, conversation messages, imported text, and external content are untrusted input. They may contain instructions that attempt to override product policy, disclose private data, or manipulate tool behavior. The system must separate instructions from data, preserve fixed policy boundaries, and never treat untrusted content as system authority.

### Unsupported claims

The AI must not invent pantry availability, recipe ingredients, expiry status, nutrition facts, or user preferences. Missing context must be stated as a limitation or trigger a clarification request. Nutrition-related guidance remains outside the MVP authority boundary.

### User control

The product must distinguish:

- Generated suggestion.
- Presented recommendation.
- User-accepted Recommendation Option.
- Explicit Meal Plan or Shopping List commitment.

AI output cannot silently cross these states or perform downstream mutations.

### Sensitive data

The gateway applies data minimization and provider-appropriate privacy controls. Credentials, access tokens, refresh secrets, raw account identifiers, and unrelated personal data must never be sent to the provider. Conversation and prompt data follow retention, deletion, and access policies.

## Resilience and Failure Handling

AI is an optional dependency for core account, pantry, recipe, meal-plan, and shopping CRUD operations. Provider failure must not make ordinary non-AI product data inaccessible.

The gateway applies:

- Bounded request timeout.
- Limited retries only for retryable transient failures.
- Exponential backoff with a total operation deadline.
- Cancellation when the client request is cancelled or the deadline expires.
- Provider quota and concurrency controls.
- Safe fallback or `Unable to Complete` result.
- Idempotency for retry-sensitive business commands.

Retries must not repeat non-idempotent business persistence. Provider calls may be retried before Recommendation creation, while a persisted Recommendation is returned or reused according to the request idempotency policy.

## Caching and Asynchronous Work

The MVP favors synchronous, bounded AI operations for user-facing recommendations and pantry analysis. A future asynchronous workflow may be introduced for long-running analysis or batch evaluation, but it must expose a durable operation state and preserve the same authorization and ownership rules.

Safe caching is allowed only when:

- The result is scoped to the authorized User or is genuinely public.
- The cache key includes the relevant user, purpose, context revision, and policy revision.
- Sensitive conversation or profile data is not shared across users.
- Cache expiry and invalidation are explicit.
- Cache loss does not change authoritative business state.

The system must not cache a personalized response under a global recipe or prompt key.

## Persistence and Data Lifecycle

PostgreSQL may retain the durable business information needed to understand an AI decision:

- Kitchen Recommendation identity, purpose, lifecycle, options, rationale, and limitations.
- References or bounded snapshots of the context used for reproducibility.
- User acceptance or rejection and relevant timestamps.
- Recommendation-scoped Conversation context under its parent Recommendation.
- Prompt, policy, model, provider, and schema metadata required by evaluation policy.

The design must distinguish:

- Current authoritative domain facts.
- Historical context snapshots used to explain a past recommendation.
- Generated AI proposals.
- User-confirmed business decisions.

Retention, deletion, export, and access rules must be defined before storing raw prompts, full conversation text, or sensitive context. Redis may hold transient request coordination and safe cache data only.

## Evaluation and Quality

AI quality is evaluated as a product capability, not only as provider availability. Evaluation should measure:

- Recommendation relevance and practical usefulness.
- Acceptance rate of Recommendation Options.
- Unsupported or hallucinated claims.
- Pantry and preference grounding accuracy.
- Safety-policy violations.
- Clarification quality when context is missing.
- Latency, provider error rate, and cost per successful operation.
- User feedback and downstream weekly AI-assisted meal planning behavior.

Evaluation datasets must be privacy-safe, representative of Indonesian MVP usage, and separated from production private data unless explicit consent and policy permit reuse. Changes to prompts, models, policies, or context assembly require regression evaluation before release.

## Observability Requirements

AI operations emit structured telemetry connected to the originating API request and business operation:

- Capability and purpose.
- Recommendation or operation identifier where available.
- Provider and model metadata.
- Prompt, policy, and schema revision identifiers.
- Latency and timeout outcome.
- Retry count and provider error category.
- Token or usage and cost signals where available.
- Validation, safety, and fallback outcome.
- Recommendation presentation and acceptance outcome.

Telemetry must redact prompts, conversation content, credentials, raw provider payloads, and unnecessary personal context. Detailed observability architecture is specified in M2-014; this document defines the AI-specific signals and privacy boundary it must preserve.

## Future Evolution

Future extensions may include additional providers, retrieval, image or barcode understanding, voice, nutrition guidance, household context, asynchronous analysis, and independent AI Chat. Each extension must preserve:

- Provider isolation through the AI Gateway.
- Explicit context authorization and minimization.
- Structured output and safety validation.
- Product ownership of recommendation meaning.
- User confirmation before business commitments.
- Measurable quality, cost, and operational behavior.

Recommendation Conversation may become an independent aggregate only if durable continuity, independent privacy, retention, feedback, or search requirements justify the change.

## Risks and Assumptions

### Risks

- Provider behavior may change and reduce quality without an evaluation gate.
- Prompt injection or malicious content may attempt to override policy or disclose private data.
- Context assembly may accidentally include unrelated or stale personal information.
- Hallucinated recipe, pantry, expiry, or nutrition claims may damage user trust.
- Retries may increase cost or duplicate business actions if idempotency is incomplete.
- AI latency, quota limits, or provider outages may degrade the core cooking experience.
- Stored prompts and conversations may become a privacy liability without retention discipline.
- Provider abstraction may hide meaningful capability differences if contracts are too generic.

### Assumptions

- OpenAI is the initial provider and is accessed only through the AI Gateway.
- PostgreSQL is authoritative for durable Recommendation and Conversation context.
- Redis is never authoritative for AI or kitchen business state.
- AI requests are authenticated and authorized except for explicitly public recipe discovery, which does not invoke personalized AI context.
- AI output is proposal data until the user and owning domain accept a later business action.
- RAG, additional providers, computer vision, voice, and nutrition coaching are future capabilities.
- `/api/v1` remains the API contract prefix.

## Exit Criteria

M2-013 is ready for review when:

- The AI domain boundary and ownership rules are explicit.
- AI use cases align with the API and Database Design.
- Context authorization, minimization, and provider isolation are defined.
- Prompt, model, policy, schema, and evaluation metadata are accounted for.
- Structured output, safety validation, and failure behavior are defined.
- AI cannot silently mutate Pantry, Recipe, Meal Plan, Shopping List, or Profile state.
- Privacy, retention, cost, latency, and quality risks are identified.
- The AI architecture flow diagram reflects the approved boundaries.

## Related Documents

- `docs/architecture/architecture-vision.md`
- `docs/architecture/api-design.md`
- `docs/architecture/database-design.md`
- `docs/architecture/bounded-context.md`
- `docs/domain/event-storming.md`
- `docs/architecture/adr/ADR-010-use-ai-gateway-abstraction.md`
- `docs/architecture/diagrams/ai-architecture-flow.mmd`
