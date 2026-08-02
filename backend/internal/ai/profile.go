package ai

import (
	"fmt"
	"time"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// Profile describes one approved model configuration (M4-DEC-010). The
// Gateway exposes a fixed default profile with versioned alternatives;
// model changes are a governed, regression-evaluated change, never implicit
// drift between environments.
type Profile struct {
	// Name is the model identifier for the configured provider.
	Name string
	// Provider is the provider adapter key (for example "openai").
	Provider string
	// Capabilities lists the capability set this profile supports.
	Capabilities []string
	// ContextBudget is the maximum context budget in tokens.
	ContextBudget int
	// MaxTokens is an upper bound on generated tokens per completion.
	MaxTokens int
	// Temperature is the sampling temperature; -1 leaves the provider default.
	Temperature float64
	// Seed, when set, requests deterministic sampling where supported.
	Seed int64
	// Timeout is the bounded per-request provider operation deadline.
	Timeout time.Duration
}

// DefaultProfile returns the MVP default OpenAI profile. The concrete model
// name and capability targets are validated during M8 with DP-SPK-003 and
// recorded in the deployment configuration; these defaults are safe for local
// development only.
func DefaultProfile() Profile {
	return Profile{
		Name:          "gpt-4o-mini",
		Provider:      "openai",
		Capabilities:  []string{"text", "structured-output"},
		ContextBudget: 128_000,
		MaxTokens:     2048,
		Temperature:   0.2,
		Seed:          0,
		Timeout:       30 * time.Second,
	}
}

// Validate ensures the profile is complete enough to invoke a provider.
func (p Profile) Validate() error {
	switch {
	case p.Name == "":
		return apperr.New(apperr.CodeAIUnavailable, "AI model profile has no model name.")
	case p.Provider == "":
		return apperr.New(apperr.CodeAIUnavailable, "AI model profile has no provider.")
	case p.ContextBudget <= 0:
		return apperr.New(apperr.CodeAIUnavailable, "AI model profile has an invalid context budget.")
	case p.MaxTokens <= 0:
		return apperr.New(apperr.CodeAIUnavailable, "AI model profile has an invalid token limit.")
	case p.Timeout <= 0:
		return apperr.New(apperr.CodeAIUnavailable, "AI model profile has an invalid timeout.")
	default:
		return nil
	}
}

// String returns a stable, redacted profile identifier for telemetry.
func (p Profile) String() string {
	return fmt.Sprintf("%s:%s", p.Provider, p.Name)
}
