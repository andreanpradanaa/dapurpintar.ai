package ai

import (
	"encoding/json"
	"fmt"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// OutputSchema describes a structured-output contract for a purpose. It is
// versioned (SchemaRev) so that prompt/model/policy changes can be
// regression-evaluated before promotion (M4-DEC-011, M4-DEC-012).
type OutputSchema struct {
	// Revision is the schema revision identifier echoed in requests/results.
	Revision string
	// JSONSchema is the JSON Schema used for provider structured output and
	// for server-side validation.
	JSONSchema map[string]any
}

// ValidateOutput checks that a provider result conforms to the schema and that
// required product fields are present. This is the schema-validation layer;
// business and safety validation are owned by the calling domain.
//
// The current validator is intentionally strict and dependency-free: it checks
// JSON well-formedness and that the object shape matches the schema's top-level
// object. Richer structural validation can replace this without changing the
// Gateway contract.
func ValidateOutput(schema OutputSchema, content []byte) error {
	if len(content) == 0 {
		return apperr.New(apperr.CodeAIUnavailable, "The AI provider returned no content.")
	}

	var v any
	if err := json.Unmarshal(content, &v); err != nil {
		return apperr.Wrap(apperr.CodeAIUnavailable, "The AI provider returned malformed output.", err)
	}

	obj, ok := v.(map[string]any)
	if !ok {
		return apperr.New(apperr.CodeAIUnavailable, "The AI provider returned unexpected output.")
	}

	props, _ := schema.JSONSchema["properties"].(map[string]any)
	for name := range props {
		required := isRequired(schema.JSONSchema, name)
		if !required {
			continue
		}
		if _, ok := obj[name]; !ok {
			return apperr.New(apperr.CodeAIUnavailable,
				fmt.Sprintf("The AI provider returned incomplete output (missing %q).", name))
		}
	}

	return nil
}

func isRequired(schema map[string]any, name string) bool {
	req, ok := schema["required"].([]any)
	if !ok {
		return false
	}
	for _, r := range req {
		if s, ok := r.(string); ok && s == name {
			return true
		}
	}
	return false
}

// KitchenRecommendationSchema returns the structured-output contract for the
// Kitchen Recommendation purpose. It captures the product-level fields an
// accepted proposal must carry; business validation of those fields is owned by
// the Recommendation domain.
func KitchenRecommendationSchema() OutputSchema {
	return OutputSchema{
		Revision: "kitchen-recommendation-v1",
		JSONSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"summary": map[string]any{"type": "string"},
				"options": map[string]any{
					"type":  "array",
					"items": map[string]any{"type": "object"},
				},
				"limitations": map[string]any{
					"type":  "array",
					"items": map[string]any{"type": "string"},
				},
				"confident": map[string]any{"type": "boolean"},
			},
			"required": []any{"summary", "options", "limitations", "confident"},
		},
	}
}
