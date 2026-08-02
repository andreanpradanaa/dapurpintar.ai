package ai

import (
	"testing"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

func TestDefaultProfile_Validate(t *testing.T) {
	p := DefaultProfile()
	if err := p.Validate(); err != nil {
		t.Fatalf("DefaultProfile().Validate() error = %v", err)
	}
	if p.Name == "" || p.Provider == "" || p.Timeout <= 0 {
		t.Error("default profile has unset required fields")
	}
}

func TestProfile_ValidateRejectsIncomplete(t *testing.T) {
	cases := []struct {
		name    string
		mutate  func(*Profile)
		wantErr bool
	}{
		{"empty name", func(p *Profile) { p.Name = "" }, true},
		{"empty provider", func(p *Profile) { p.Provider = "" }, true},
		{"zero context budget", func(p *Profile) { p.ContextBudget = 0 }, true},
		{"zero max tokens", func(p *Profile) { p.MaxTokens = 0 }, true},
		{"zero timeout", func(p *Profile) { p.Timeout = 0 }, true},
		{"valid custom", func(p *Profile) { p.Name = "gpt-4o" }, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := DefaultProfile()
			tc.mutate(&p)
			err := p.Validate()
			if tc.wantErr && err == nil {
				t.Fatal("Validate() expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("Validate() unexpected error: %v", err)
			}
			if tc.wantErr {
				appErr, ok := apperr.As(err)
				if !ok || appErr.Code != apperr.CodeAIUnavailable {
					t.Errorf("error code = %v, want AI_UNAVAILABLE", appErr.Code)
				}
			}
		})
	}
}

func TestProfile_String(t *testing.T) {
	p := DefaultProfile()
	got := p.String()
	if got != "openai:gpt-4o-mini" {
		t.Errorf("String() = %q, want openai:gpt-4o-mini", got)
	}
}

func TestValidateOutput_EmptyContent(t *testing.T) {
	schema := KitchenRecommendationSchema()
	if err := ValidateOutput(schema, nil); err == nil {
		t.Fatal("ValidateOutput() expected error for empty content")
	}
}

func TestValidateOutput_MalformedJSON(t *testing.T) {
	schema := KitchenRecommendationSchema()
	if err := ValidateOutput(schema, []byte("not json")); err == nil {
		t.Fatal("ValidateOutput() expected error for malformed content")
	}
}

func TestValidateOutput_MissingRequiredField(t *testing.T) {
	schema := KitchenRecommendationSchema()
	content := []byte(`{"summary":"ok","options":[],"limitations":[]}`)
	if err := ValidateOutput(schema, content); err == nil {
		t.Fatal("ValidateOutput() expected error for missing confident field")
	}
}

func TestValidateOutput_Valid(t *testing.T) {
	schema := KitchenRecommendationSchema()
	content := []byte(`{"summary":"ok","options":[],"limitations":["none"],"confident":true}`)
	if err := ValidateOutput(schema, content); err != nil {
		t.Fatalf("ValidateOutput() error = %v", err)
	}
}

func TestValidateOutput_NonObject(t *testing.T) {
	schema := KitchenRecommendationSchema()
	if err := ValidateOutput(schema, []byte(`[1,2,3]`)); err == nil {
		t.Fatal("ValidateOutput() expected error for non-object content")
	}
}
