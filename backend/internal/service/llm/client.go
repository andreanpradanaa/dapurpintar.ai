package llm

import (
	"context"
	"time"

	"github.com/dapurpintar/backend/internal/model"
)

// Client is the interface for any LLM provider. The generator service
// depends on this interface, not a concrete client, so we can swap
// OpenAI / Anthropic / a mock without changing the orchestrator.
type Client interface {
	// GenerateRecipe returns a fresh recipe based on the user's
	// ingredients, dietary preferences, language, and a set of
	// reference recipes for style guidance (RAG).
	GenerateRecipe(ctx context.Context, req GenerateRequest) (*model.Recipe, error)
	// Name returns the provider name (for health checks and logging).
	Name() string
}

type GenerateRequest struct {
	UserIngredients []string
	Dietary         []string
	Language        string // "en" | "id"
	References      []*model.Recipe
	Timeout         time.Duration
}
