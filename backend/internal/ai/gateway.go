// Package ai defines the AI Gateway boundary for DapurPintar AI.
//
// The Gateway is the single application boundary for external AI assistance
// (docs/architecture/ai-architecture.md, ADR-010). Business modules depend on
// the provider-neutral Gateway contract and product vocabulary, never on a
// provider SDK. Each request is a minimized, authorized payload; each result is
// a validated, provider-independent outcome.
package ai

import (
	"context"
	"time"
)

// Purpose identifies a product capability requested from the AI Gateway.
type Purpose string

const (
	// PurposeKitchenRecommendation is a contextual Kitchen Recommendation
	// proposal (POST /api/v1/recommendations).
	PurposeKitchenRecommendation Purpose = "kitchen-recommendation"

	// PurposePantryAnalysis is a Pantry use-first analysis
	// (POST /api/v1/ai/pantry-analysis).
	PurposePantryAnalysis Purpose = "pantry-analysis"
)

// Role is a message author in product terms.
type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

// Message is a single product-level conversation or instruction turn.
type Message struct {
	Role    Role
	Content string
}

// Request is a minimized, provider-neutral AI request. It carries only the
// authorized context needed for the purpose and the policy revision identifiers
// required for reproducibility. Credentials and unrelated personal data are
// never part of a request.
type Request struct {
	Purpose Purpose
	// Messages carries the bounded instruction and conversation turns. The
	// system message is assembled from the versioned prompt and safety policy;
	// raw provider prompts and payloads are never stored.
	Messages []Message
	// Profile selects the approved model profile (M4-DEC-010).
	Profile Profile
	// Revisions identify the prompt, safety policy, and output schema versions
	// used, enabling reproducibility and regression evaluation.
	PromptRev string
	SafetyRev string
	SchemaRev string
}

// Usage carries token and cost metadata for observability and quota control
// (M4-DEC-016).
type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// Result is a validated, provider-independent AI outcome. Content is the
// structured output that already passed transport and schema validation; it is
// never a raw provider payload.
type Result struct {
	// Content is the validated structured output as raw JSON.
	Content []byte
	Model   string
	Usage   Usage
	// Latency is the total provider operation latency.
	Latency time.Duration
	// Provider is the adapter identifier (for example "openai").
	Provider string
	// Revisions echo the request revisions for observability.
	PromptRev string
	SafetyRev string
	SchemaRev string
}

// Gateway is the single application boundary for external AI assistance.
//
// Implementations are responsible for provider invocation, bounded timeout and
// retry, provider error translation, and transport/schema validation. A
// provider failure or invalid output is returned as a safe error (M6 AI codes);
// it is never surfaced as a product commitment.
type Gateway interface {
	// Complete performs a bounded AI operation for the given purpose.
	Complete(ctx context.Context, req Request) (*Result, error)
}
