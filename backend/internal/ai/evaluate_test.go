package ai

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// fakeGateway is a stub Gateway returning canned structured content.
type fakeGateway struct {
	content []byte
	err     error
}

func (f *fakeGateway) Complete(_ context.Context, _ Request) (*Result, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &Result{Content: f.content, Model: "fake", Provider: "fake"}, nil
}

// groundedGateway generates a grounded Kitchen Recommendation response from the
// authorized context in the request, satisfying each scenario's Expected terms
// and honoring thin-context scenarios. It never invents terms.
type groundedGateway struct{}

func (g *groundedGateway) Complete(_ context.Context, req Request) (*Result, error) {
	var full strings.Builder
	for _, m := range req.Messages {
		full.WriteString(m.Content)
		full.WriteString("\n")
	}
	text := full.String()
	emptyPantry := strings.Contains(text, "(no items recorded)")
	thin := emptyPantry || strings.Contains(text, "for the week")

	summary := "Use the pantry staples for a simple dinner."
	ingredients := []string{"egg", "rice"}
	if strings.Contains(text, "expire") {
		summary = "Use spinach and tomatoes tonight."
		ingredients = []string{"spinach", "tomato"}
	}
	if strings.Contains(text, "vegetarian") {
		summary = "Make a chickpea and lentil stew."
		ingredients = []string{"chickpea", "lentil"}
	}
	if thin {
		summary = "With only eggs, a simple omelette is possible."
		ingredients = []string{"egg"}
	}

	options := []any{map[string]any{"title": summary, "ingredients": ingredients}}
	limitations := []string{}
	confident := true
	if emptyPantry {
		// Empty pantry: no options, an honest limitation, no invented facts.
		summary = "Your pantry has no recorded items, so I cannot recommend a meal from inventory."
		options = nil
		limitations = []string{"No pantry items are recorded; recommendations require a pantry inventory."}
		confident = false
	} else if thin {
		limitations = []string{"Pantry context is thin; recommendations are limited."}
		confident = false
	}

	out := kitchenOutput{Summary: summary, Options: options, Limitations: limitations, Confident: confident}
	content, _ := json.Marshal(out)
	return &Result{Content: content, Model: "fake", Provider: "fake"}, nil
}

func kitchenOutputJSON(out kitchenOutput) []byte {
	b, _ := json.Marshal(out)
	return b
}

func TestEvaluate_AllPass(t *testing.T) {
	eval := NewEvaluator(&groundedGateway{}, DefaultRubric())
	registry := SeedRegistry()

	report, err := eval.Evaluate(context.Background(), registry, PurposeKitchenRecommendation, "", SeedScenarios())
	if err != nil {
		t.Fatalf("Evaluate() error = %v", err)
	}
	if report.Pass {
		return
	}
	for _, res := range report.Results {
		t.Logf("scenario %s pass=%v overall=%.2f", res.ScenarioID, res.Pass, res.Overall)
		for _, s := range res.Scores {
			t.Logf("  %s value=%.2f pass=%v notes=%v", s.Dimension, s.Value, s.Pass, s.Notes)
		}
	}
	t.Fatal("Evaluate() expected gate to pass")
}

func TestEvaluate_InventedTermFailsGate(t *testing.T) {
	// rec-001 forbids salmon; a response that invents it must fail.
	gateway := &fakeGateway{content: kitchenOutputJSON(kitchenOutput{
		Summary:     "Grilled salmon with spinach.",
		Options:     []any{map[string]any{"title": "Salmon dish", "ingredients": []string{"salmon", "spinach"}}},
		Limitations: []string{},
		Confident:   true,
	})}
	eval := NewEvaluator(gateway, DefaultRubric())
	registry := SeedRegistry()

	report, err := eval.Evaluate(context.Background(), registry, PurposeKitchenRecommendation, "", SeedScenarios())
	if err != nil {
		t.Fatalf("Evaluate() error = %v", err)
	}
	if report.Pass {
		t.Fatal("Evaluate() expected gate to fail when output invents prohibited terms")
	}
}

func TestEvaluate_ThinContextMustDeclareLimitation(t *testing.T) {
	// rec-002 and rec-004 are thin-context; high confidence without
	// limitations must fail the safety dimension.
	gateway := &fakeGateway{content: kitchenOutputJSON(kitchenOutput{
		Summary:     "Omelette from eggs.",
		Options:     []any{map[string]any{"title": "Omelette", "ingredients": []string{"egg"}}},
		Limitations: []string{},
		Confident:   true,
	})}
	eval := NewEvaluator(gateway, DefaultRubric())
	registry := SeedRegistry()

	report, err := eval.Evaluate(context.Background(), registry, PurposeKitchenRecommendation, "", SeedScenarios())
	if err != nil {
		t.Fatalf("Evaluate() error = %v", err)
	}
	if report.Pass {
		t.Fatal("Evaluate() expected gate to fail on thin-context confidence")
	}

	var safetyOk bool
	for _, res := range report.Results {
		for _, s := range res.Scores {
			if s.Dimension == DimensionSafety && !s.Pass {
				safetyOk = true
			}
		}
	}
	if !safetyOk {
		t.Error("expected at least one safety dimension failure")
	}
}

func TestEvaluate_MissingGroundingFailsAccuracy(t *testing.T) {
	// rec-003 expects chickpea/lentil grounding and forbids chicken; missing
	// grounding reduces accuracy below the minimum.
	gateway := &fakeGateway{content: kitchenOutputJSON(kitchenOutput{
		Summary:     "Plain rice.",
		Options:     []any{map[string]any{"title": "Rice", "ingredients": []string{"rice"}}},
		Limitations: []string{"Pantry context is thin."},
		Confident:   false,
	})}
	eval := NewEvaluator(gateway, DefaultRubric())
	registry := SeedRegistry()

	report, err := eval.Evaluate(context.Background(), registry, PurposeKitchenRecommendation, "", SeedScenarios())
	if err != nil {
		t.Fatalf("Evaluate() error = %v", err)
	}
	if report.Pass {
		t.Fatal("Evaluate() expected gate to fail when grounding is missing")
	}
}

func TestEvaluate_MalformedOutputFailsConformance(t *testing.T) {
	gateway := &fakeGateway{content: []byte(`not json`)}
	eval := NewEvaluator(gateway, DefaultRubric())
	registry := SeedRegistry()

	report, err := eval.Evaluate(context.Background(), registry, PurposeKitchenRecommendation, "", SeedScenarios())
	if err != nil {
		t.Fatalf("Evaluate() error = %v", err)
	}
	if report.Pass {
		t.Fatal("Evaluate() expected gate to fail on malformed output")
	}
	for _, res := range report.Results {
		for _, s := range res.Scores {
			if s.Dimension == DimensionConformance && s.Value != 0 {
				t.Errorf("conformance value = %.2f, want 0", s.Value)
			}
		}
	}
}

func TestEvaluate_GatewayErrorPropagates(t *testing.T) {
	gateway := &fakeGateway{err: apperr.New(apperr.CodeAIUnavailable, "provider down")}
	eval := NewEvaluator(gateway, DefaultRubric())
	registry := SeedRegistry()

	if _, err := eval.Evaluate(context.Background(), registry, PurposeKitchenRecommendation, "", SeedScenarios()); err == nil {
		t.Fatal("Evaluate() expected error to propagate")
	}
}

func TestEvaluate_UnknownPurpose(t *testing.T) {
	gateway := &fakeGateway{}
	eval := NewEvaluator(gateway, DefaultRubric())
	registry := SeedRegistry()

	if _, err := eval.Evaluate(context.Background(), registry, Purpose("missing"), "", SeedScenarios()); err == nil {
		t.Fatal("Evaluate() expected error for unknown purpose")
	}
}
