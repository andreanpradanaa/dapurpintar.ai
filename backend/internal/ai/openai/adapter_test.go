package openai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/ai"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

func newTestAdapter(t *testing.T, handler http.HandlerFunc) *Adapter {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	adapter, err := New(Config{
		APIKey:     "test-key",
		BaseURL:    server.URL,
		MaxRetries: 0,
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	return adapter
}

func recommendationRequest() ai.Request {
	return ai.Request{
		Purpose: ai.PurposeKitchenRecommendation,
		Profile: ai.DefaultProfile(),
		Messages: []ai.Message{
			{Role: ai.RoleSystem, Content: "You are a helpful kitchen assistant."},
			{Role: ai.RoleUser, Content: "Suggest meals from my pantry."},
		},
		PromptRev: "kitchen-recommendation-v1",
		SafetyRev: "safety-v1",
		SchemaRev: "kitchen-recommendation-v1",
	}
}

func validCompletionJSON() string {
	return `{
		"id": "chatcmpl-test",
		"object": "chat.completion",
		"created": 1710000000,
		"model": "gpt-4o-mini",
		"choices": [{
			"index": 0,
			"message": {
				"role": "assistant",
				"content": "{\"summary\":\"Use up tomatoes and spinach\",\"options\":[],\"limitations\":[\"Pantry confidence is low\"],\"confident\":false}"
			},
			"finish_reason": "stop"
		}],
		"usage": {
			"prompt_tokens": 120,
			"completion_tokens": 40,
			"total_tokens": 160
		}
	}`
}

func TestComplete_Success(t *testing.T) {
	adapter := newTestAdapter(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(validCompletionJSON()))
	})

	result, err := adapter.Complete(context.Background(), recommendationRequest())
	if err != nil {
		t.Fatalf("Complete() error = %v", err)
	}
	if result.Provider != "openai" {
		t.Errorf("Provider = %q, want openai", result.Provider)
	}
	if result.Model != "gpt-4o-mini" {
		t.Errorf("Model = %q, want gpt-4o-mini", result.Model)
	}
	if result.Usage.TotalTokens != 160 {
		t.Errorf("TotalTokens = %d, want 160", result.Usage.TotalTokens)
	}
	if result.SchemaRev != "kitchen-recommendation-v1" {
		t.Errorf("SchemaRev = %q", result.SchemaRev)
	}

	var out map[string]any
	if err := json.Unmarshal(result.Content, &out); err != nil {
		t.Fatalf("content not valid JSON: %v", err)
	}
	if out["summary"] == "" {
		t.Error("content missing summary field")
	}
}

func TestComplete_SendsStructuredOutputRequest(t *testing.T) {
	var gotBody map[string]any
	adapter := newTestAdapter(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(validCompletionJSON()))
	})

	if _, err := adapter.Complete(context.Background(), recommendationRequest()); err != nil {
		t.Fatalf("Complete() error = %v", err)
	}

	rf, ok := gotBody["response_format"].(map[string]any)
	if !ok {
		t.Fatal("response_format missing from request body")
	}
	if rf["type"] != "json_schema" {
		t.Errorf("response_format.type = %v, want json_schema", rf["type"])
	}
	schema, ok := rf["json_schema"].(map[string]any)
	if !ok {
		t.Fatal("response_format.json_schema missing")
	}
	if schema["name"] != "kitchen-recommendation-v1" {
		t.Errorf("schema name = %v, want kitchen-recommendation-v1", schema["name"])
	}
}

func TestComplete_NoContentReturnsError(t *testing.T) {
	adapter := newTestAdapter(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id": "x", "object": "chat.completion", "created": 1, "model": "m",
			"choices": [{"index": 0, "message": {"role": "assistant", "content": ""}, "finish_reason": "stop"}],
			"usage": {"prompt_tokens": 1, "completion_tokens": 1, "total_tokens": 2}
		}`))
	})

	_, err := adapter.Complete(context.Background(), recommendationRequest())
	if err == nil {
		t.Fatal("Complete() expected error for empty content")
	}
	appErr, ok := apperr.As(err)
	if !ok || appErr.Code != apperr.CodeAIUnavailable {
		t.Errorf("error code = %v, want AI_UNAVAILABLE", appErr.Code)
	}
}

func TestComplete_ContentFilterReturnsSafetyError(t *testing.T) {
	adapter := newTestAdapter(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id": "x", "object": "chat.completion", "created": 1, "model": "m",
			"choices": [{"index": 0, "message": {"role": "assistant", "content": "blocked"}, "finish_reason": "content_filter"}],
			"usage": {"prompt_tokens": 1, "completion_tokens": 1, "total_tokens": 2}
		}`))
	})

	_, err := adapter.Complete(context.Background(), recommendationRequest())
	if err == nil {
		t.Fatal("Complete() expected error for content filter")
	}
	appErr, ok := apperr.As(err)
	if !ok || appErr.Code != apperr.CodeAISafetyRejected {
		t.Errorf("error code = %v, want AI_SAFETY_REJECTED", appErr.Code)
	}
}

func TestComplete_InvalidOutputRejected(t *testing.T) {
	adapter := newTestAdapter(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id": "x", "object": "chat.completion", "created": 1, "model": "m",
			"choices": [{"index": 0, "message": {"role": "assistant", "content": "{\"options\":[]}"}, "finish_reason": "stop"}],
			"usage": {"prompt_tokens": 1, "completion_tokens": 1, "total_tokens": 2}
		}`))
	})

	_, err := adapter.Complete(context.Background(), recommendationRequest())
	if err == nil {
		t.Fatal("Complete() expected schema validation error")
	}
	appErr, ok := apperr.As(err)
	if !ok || appErr.Code != apperr.CodeAIUnavailable {
		t.Errorf("error code = %v, want AI_UNAVAILABLE", appErr.Code)
	}
}

func TestComplete_ProviderQuotaError(t *testing.T) {
	adapter := newTestAdapter(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"error":{"message":"rate limited","type":"rate_limit_error","param":null,"code":"rate_limit_exceeded"}}`))
	})

	_, err := adapter.Complete(context.Background(), recommendationRequest())
	if err == nil {
		t.Fatal("Complete() expected quota error")
	}
	appErr, ok := apperr.As(err)
	if !ok || appErr.Code != apperr.CodeAIQuotaExceeded {
		t.Errorf("error code = %v, want AI_QUOTA_EXCEEDED", appErr.Code)
	}
}

func TestComplete_ProviderServerError(t *testing.T) {
	adapter := newTestAdapter(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"message":"boom","type":"server_error","param":null,"code":"server_error"}}`))
	})

	_, err := adapter.Complete(context.Background(), recommendationRequest())
	if err == nil {
		t.Fatal("Complete() expected provider error")
	}
	appErr, ok := apperr.As(err)
	if !ok || appErr.Code != apperr.CodeAIUnavailable {
		t.Errorf("error code = %v, want AI_UNAVAILABLE", appErr.Code)
	}
}

func TestComplete_UnsupportedPurpose(t *testing.T) {
	adapter := newTestAdapter(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("provider should not be called for unsupported purpose")
	})

	req := recommendationRequest()
	req.Purpose = ai.Purpose("unsupported-purpose")

	_, err := adapter.Complete(context.Background(), req)
	if err == nil {
		t.Fatal("Complete() expected unsupported purpose error")
	}
	appErr, ok := apperr.As(err)
	if !ok || appErr.Code != apperr.CodeAIUnavailable {
		t.Errorf("error code = %v, want AI_UNAVAILABLE", appErr.Code)
	}
}

func TestComplete_MissingAPIKey(t *testing.T) {
	if _, err := New(Config{}); err == nil {
		t.Fatal("New() expected error when API key is empty")
	}
}
