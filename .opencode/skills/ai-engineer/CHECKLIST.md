# AI Engineer Checklist

## Before starting AI work

- [ ] Read `docs/architecture/ai-architecture.md`.
- [ ] Confirm the AI decisions M4-DEC-010/011/012/016 status in the Decision Register.
- [ ] Confirm the Recommendation lifecycle and invariants.
- [ ] Confirm the retention policy for prompts and conversations (M4-DEC-013).

## While implementing

- [ ] Business modules depend on the Gateway port, not a provider SDK.
- [ ] Prompts and policies are versioned.
- [ ] Structured output validates against a schema.
- [ ] Timeout, retry, quota, and cost controls are bounded.
- [ ] Provider failure degrades safely; core operations stay usable.
- [ ] No raw provider prompts or payloads are stored in business data.
- [ ] AI output is never silently converted into a commitment.

## Before finishing

- [ ] Gateway and adapter tests pass.
- [ ] Safety and injection test cases pass.
- [ ] Evaluation scenarios run against the harness.
- [ ] Quota/cost alerts are wired.
- [ ] Documentation and decision records are updated.
