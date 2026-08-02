package ai

import (
	"strings"
	"testing"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

func basePolicy(purpose Purpose, rev string) PolicyBundle {
	return PolicyBundle{
		Purpose:        purpose,
		PromptRevision: "prompt-" + rev,
		SafetyRevision: "safety-" + rev,
		SchemaRevision: rev,
		PromptText:     "Recommend meals from the pantry.",
		SafetyText:     "Never invent pantry facts.",
		JSONSchema:     KitchenRecommendationSchema().JSONSchema,
		Status:         RevisionPending,
	}
}

func TestSeedRegistry_PromotesInitialKitchenPolicy(t *testing.T) {
	r := SeedRegistry()
	b, err := r.Resolve(PurposeKitchenRecommendation, "")
	if err != nil {
		t.Fatalf("Resolve(active) error = %v", err)
	}
	if b.Status != RevisionPromoted {
		t.Errorf("status = %v, want promoted", b.Status)
	}
	if b.SchemaRevision != "kitchen-recommendation-v1" {
		t.Errorf("schema revision = %q", b.SchemaRevision)
	}
}

func TestRegistry_RegisterAndResolveActive(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(basePolicy(PurposeKitchenRecommendation, "v1"))
	_ = r.Promote(PurposeKitchenRecommendation, "v1")

	_ = r.Register(basePolicy(PurposeKitchenRecommendation, "v2"))

	active, err := r.Resolve(PurposeKitchenRecommendation, "")
	if err != nil {
		t.Fatalf("Resolve(active) error = %v", err)
	}
	if active.SchemaRevision != "v1" {
		t.Errorf("active = %q, want v1 (unpromoted v2 must not leak)", active.SchemaRevision)
	}

	// Promoting v2 flips the active revision.
	if err := r.Promote(PurposeKitchenRecommendation, "v2"); err != nil {
		t.Fatalf("Promote(v2) error = %v", err)
	}
	active, _ = r.Resolve(PurposeKitchenRecommendation, "")
	if active.SchemaRevision != "v2" {
		t.Errorf("active = %q, want v2", active.SchemaRevision)
	}

	// The older revision remains available for evaluation and rollback.
	pinned, err := r.Resolve(PurposeKitchenRecommendation, "v1")
	if err != nil {
		t.Fatalf("Resolve(v1) error = %v", err)
	}
	if pinned.SchemaRevision != "v1" {
		t.Errorf("pinned = %q, want v1", pinned.SchemaRevision)
	}
}

func TestRegistry_RejectsDuplicateRevision(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(basePolicy(PurposeKitchenRecommendation, "v1"))
	err := r.Register(basePolicy(PurposeKitchenRecommendation, "v1"))
	if err == nil {
		t.Fatal("Register() expected error for duplicate schema revision")
	}
}

func TestRegistry_RejectsInvalidBundle(t *testing.T) {
	r := NewRegistry()
	err := r.Register(PolicyBundle{Purpose: PurposeKitchenRecommendation})
	if err == nil {
		t.Fatal("Register() expected validation error")
	}
	appErr, ok := apperr.As(err)
	if !ok || appErr.Code != apperr.CodeAIUnavailable {
		t.Errorf("error code = %v, want AI_UNAVAILABLE", appErr.Code)
	}
}

func TestRegistry_UnknownPurpose(t *testing.T) {
	r := NewRegistry()
	if _, err := r.Resolve(Purpose("missing"), ""); err == nil {
		t.Fatal("Resolve() expected error for unknown purpose")
	}
}

func TestRegistry_PromoteUnknownRevision(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(basePolicy(PurposeKitchenRecommendation, "v1"))
	if err := r.Promote(PurposeKitchenRecommendation, "v9"); err == nil {
		t.Fatal("Promote() expected error for unknown revision")
	}
}

func TestBuildSystemMessage_InjectsContext(t *testing.T) {
	b := SeedKitchenRecommendationPolicy()
	msg := BuildSystemMessage(b, "pantry: tomatoes, spinach")
	if msg.Role != RoleSystem {
		t.Errorf("role = %q, want system", msg.Role)
	}
	if !strings.Contains(msg.Content, "pantry: tomatoes, spinach") {
		t.Error("system message missing injected context")
	}
	if !strings.Contains(msg.Content, b.SafetyText) {
		t.Error("system message missing safety policy")
	}
	if !strings.Contains(msg.Content, b.PromptText) {
		t.Error("system message missing prompt text")
	}
}

func TestBuildSystemMessage_EmptyContextSkipped(t *testing.T) {
	b := SeedKitchenRecommendationPolicy()
	msg := BuildSystemMessage(b, "   ")
	if strings.Contains(msg.Content, "Authorized pantry context") {
		t.Error("system message should not include context section when context is empty")
	}
}

func TestBuildKitchenRecommendationRequest(t *testing.T) {
	b := SeedKitchenRecommendationPolicy()
	req := BuildKitchenRecommendationRequest(b, "pantry: eggs", "use up eggs", DefaultProfile())
	if req.Purpose != PurposeKitchenRecommendation {
		t.Errorf("purpose = %q", req.Purpose)
	}
	if len(req.Messages) != 2 {
		t.Fatalf("messages = %d, want 2", len(req.Messages))
	}
	if req.Messages[1].Role != RoleUser || req.Messages[1].Content != "use up eggs" {
		t.Errorf("user message = %+v", req.Messages[1])
	}
	if req.PromptRev != "kitchen-recommendation-v1" || req.SafetyRev != "safety-v1" {
		t.Errorf("revisions = prompt:%q safety:%q", req.PromptRev, req.SafetyRev)
	}
	if req.SchemaRev != "kitchen-recommendation-v1" {
		t.Errorf("schema rev = %q", req.SchemaRev)
	}
}

func TestBuildKitchenRecommendationRequest_NoUserIntent(t *testing.T) {
	b := SeedKitchenRecommendationPolicy()
	req := BuildKitchenRecommendationRequest(b, "pantry: eggs", "", DefaultProfile())
	if len(req.Messages) != 1 {
		t.Fatalf("messages = %d, want 1 when no user intent", len(req.Messages))
	}
}

func TestPolicyBundle_Validate(t *testing.T) {
	b := SeedKitchenRecommendationPolicy()
	if err := b.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}
