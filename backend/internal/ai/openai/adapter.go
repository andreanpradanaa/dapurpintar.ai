// Package openai implements the AI Gateway provider adapter for OpenAI
// (docs/architecture/ai-architecture.md, ADR-010). It is the only package that
// touches the OpenAI SDK: credentials, request mapping, transport behavior, and
// provider error classification. It must never be called directly by HTTP
// handlers or domain code; business modules depend on the ai.Gateway contract.
package openai

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/shared"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/ai"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// Adapter is the OpenAI provider adapter for the AI Gateway. It owns bounded
// timeout and retry behavior, structured output, schema validation, and
// provider error translation.
type Adapter struct {
	client  openai.Client
	timeout time.Duration
	log     *slog.Logger
}

// Config carries the OpenAI adapter settings. APIKey is a deployment secret
// supplied through environment or secret management; it is never logged.
type Config struct {
	// APIKey is the OpenAI API key.
	APIKey string
	// Timeout bounds the per-request provider operation (M4-DEC-016).
	Timeout time.Duration
	// MaxRetries bounds automatic retries for retryable transient failures.
	MaxRetries int
	// BaseURL overrides the provider base URL (test/sandbox use). Empty uses
	// the OpenAI default.
	BaseURL string
	// HTTPClient supplies the transport. Nil uses a default bounded client.
	HTTPClient *http.Client
	// Log is the application logger. Nil disables adapter logging.
	Log *slog.Logger
}

// New builds an OpenAI adapter implementing ai.Gateway.
func New(cfg Config) (*Adapter, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("openai adapter: api key required")
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 30 * time.Second
	}
	if cfg.MaxRetries < 0 {
		cfg.MaxRetries = 0
	}

	opts := []option.RequestOption{
		option.WithAPIKey(cfg.APIKey),
		option.WithMaxRetries(cfg.MaxRetries),
	}
	if cfg.BaseURL != "" {
		opts = append(opts, option.WithBaseURL(cfg.BaseURL))
	}
	if cfg.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(cfg.HTTPClient))
	} else {
		opts = append(opts, option.WithHTTPClient(&http.Client{
			Timeout: cfg.Timeout,
		}))
	}

	return &Adapter{
		client:  openai.NewClient(opts...),
		timeout: cfg.Timeout,
		log:     cfg.Log,
	}, nil
}

// Complete implements ai.Gateway. It builds a structured-output chat request,
// invokes the provider under a bounded timeout, validates the result schema,
// and translates failures into safe M6 errors.
func (a *Adapter) Complete(ctx context.Context, req ai.Request) (*ai.Result, error) {
	if err := req.Profile.Validate(); err != nil {
		return nil, err
	}
	if req.Purpose == "" {
		return nil, apperr.New(apperr.CodeAIUnavailable, "AI request purpose is required.")
	}
	if len(req.Messages) == 0 {
		return nil, apperr.New(apperr.CodeAIUnavailable, "AI request has no messages.")
	}

	schema, err := schemaFor(req.Purpose)
	if err != nil {
		return nil, err
	}

	params := openai.ChatCompletionNewParams{
		Model:    openai.ChatModel(req.Profile.Name),
		Messages: buildMessages(req.Messages),
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &shared.ResponseFormatJSONSchemaParam{
				JSONSchema: shared.ResponseFormatJSONSchemaJSONSchemaParam{
					Name:   schema.Revision,
					Strict: param.NewOpt(true),
					Schema: schema.JSONSchema,
				},
			},
		},
	}
	if req.Profile.Temperature >= 0 {
		params.Temperature = param.NewOpt(req.Profile.Temperature)
	}
	if req.Profile.MaxTokens > 0 {
		params.MaxCompletionTokens = param.NewOpt(int64(req.Profile.MaxTokens))
	}

	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	start := time.Now()
	completion, err := a.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return nil, translateError(err)
	}

	if len(completion.Choices) == 0 || completion.Choices[0].Message.Content == "" {
		return nil, apperr.New(apperr.CodeAIUnavailable, "The AI provider returned no content.")
	}
	if completion.Choices[0].FinishReason == "content_filter" {
		return nil, apperr.New(apperr.CodeAISafetyRejected, "The AI provider declined the request on safety grounds.")
	}

	content := []byte(completion.Choices[0].Message.Content)
	if err := ai.ValidateOutput(schema, content); err != nil {
		return nil, err
	}

	result := &ai.Result{
		Content:   content,
		Model:     completion.Model,
		Latency:   time.Since(start),
		Provider:  "openai",
		PromptRev: req.PromptRev,
		SafetyRev: req.SafetyRev,
		SchemaRev: schema.Revision,
	}
	if completion.Usage.TotalTokens > 0 || completion.Usage.PromptTokens > 0 {
		result.Usage = ai.Usage{
			PromptTokens:     int(completion.Usage.PromptTokens),
			CompletionTokens: int(completion.Usage.CompletionTokens),
			TotalTokens:      int(completion.Usage.TotalTokens),
		}
	}

	if a.log != nil {
		a.log.Info("ai gateway completion",
			"purpose", string(req.Purpose),
			"model", result.Model,
			"provider", result.Provider,
			"schema_rev", result.SchemaRev,
			"prompt_rev", result.PromptRev,
			"latency_ms", result.Latency.Milliseconds(),
			"tokens", result.Usage.TotalTokens,
		)
	}

	return result, nil
}

// buildMessages maps product-level messages to OpenAI message parameters.
func buildMessages(messages []ai.Message) []openai.ChatCompletionMessageParamUnion {
	out := make([]openai.ChatCompletionMessageParamUnion, 0, len(messages))
	for _, m := range messages {
		switch m.Role {
		case ai.RoleSystem:
			out = append(out, openai.SystemMessage(m.Content))
		case ai.RoleAssistant:
			out = append(out, openai.AssistantMessage(m.Content))
		default:
			out = append(out, openai.UserMessage(m.Content))
		}
	}
	return out
}

// translateError classifies a provider error into a stable M6 AI error code.
func translateError(err error) *apperr.Error {
	var apiErr *openai.Error
	if errors.As(err, &apiErr) {
		switch {
		case apiErr.StatusCode == http.StatusTooManyRequests:
			return apperr.Wrap(apperr.CodeAIQuotaExceeded, "The AI provider is over quota. Please try again later.", err)
		case apiErr.StatusCode == http.StatusBadRequest:
			return apperr.Wrap(apperr.CodeAISafetyRejected, "The AI provider rejected the request.", err)
		case apiErr.StatusCode >= http.StatusInternalServerError:
			return apperr.Wrap(apperr.CodeAIUnavailable, "The AI provider is temporarily unavailable.", err)
		}
		return apperr.Wrap(apperr.CodeAIUnavailable, "The AI provider request failed.", err)
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return apperr.Wrap(apperr.CodeAIUnavailable, "The AI provider request timed out.", err)
	}
	return apperr.Wrap(apperr.CodeAIUnavailable, "The AI provider request failed.", err)
}

// schemaFor returns the structured-output schema for a purpose.
func schemaFor(purpose ai.Purpose) (ai.OutputSchema, error) {
	switch purpose {
	case ai.PurposeKitchenRecommendation:
		return ai.KitchenRecommendationSchema(), nil
	default:
		return ai.OutputSchema{}, apperr.New(apperr.CodeAIUnavailable,
			fmt.Sprintf("Unsupported AI purpose %q.", purpose))
	}
}
