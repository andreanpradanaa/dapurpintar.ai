package ai

import (
	"time"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// RevisionStatus tracks whether a policy revision is safe to use in requests.
//
// M4-DEC-011 requires prompts, safety policy, and output schemas to be governed,
// versioned artifacts whose changes are reviewable and regression-evaluated
// before promotion. Only a Promoted revision is served to live requests; a
// Pending revision exists for review and evaluation.
type RevisionStatus string

const (
	// RevisionPending is a new revision awaiting review and evaluation.
	RevisionPending RevisionStatus = "pending"
	// RevisionPromoted is the active revision served to live requests.
	RevisionPromoted RevisionStatus = "promoted"
)

// PolicyBundle is an immutable, versioned set of the artifacts that govern one
// AI purpose: the prompt text, the safety policy, and the output schema. All
// three revisions are recorded so a request is fully reproducible.
type PolicyBundle struct {
	// Purpose is the capability this bundle governs.
	Purpose Purpose
	// PromptRevision and SafetyRevision identify the prompt and safety policy
	// versions.
	PromptRevision string
	SafetyRevision string
	// SchemaRevision identifies the structured-output schema version.
	SchemaRevision string
	// PromptText is the versioned system prompt template. It is a controlled
	// product artifact, not an arbitrary handler string.
	PromptText string
	// SafetyText is the versioned safety policy injected into the system
	// message (M4-DEC-011).
	SafetyText string
	// JSONSchema is the structured-output schema versioned with this bundle.
	JSONSchema map[string]any
	// Status is the promotion status of this revision.
	Status RevisionStatus
	// CreatedAt records when the revision was registered.
	CreatedAt time.Time
}

// Validate ensures a bundle is complete enough to build a request.
func (b PolicyBundle) Validate() error {
	switch {
	case b.Purpose == "":
		return apperr.New(apperr.CodeAIUnavailable, "AI policy bundle has no purpose.")
	case b.PromptRevision == "" || b.SafetyRevision == "" || b.SchemaRevision == "":
		return apperr.New(apperr.CodeAIUnavailable, "AI policy bundle is missing a revision identifier.")
	case b.PromptText == "":
		return apperr.New(apperr.CodeAIUnavailable, "AI policy bundle has no prompt text.")
	case b.SafetyText == "":
		return apperr.New(apperr.CodeAIUnavailable, "AI policy bundle has no safety policy.")
	case b.Status == "":
		return apperr.New(apperr.CodeAIUnavailable, "AI policy bundle has no promotion status.")
	default:
		return nil
	}
}

// Revisions returns the three revision identifiers as a stable set.
func (b PolicyBundle) Revisions() [3]string {
	return [3]string{b.PromptRevision, b.SafetyRevision, b.SchemaRevision}
}
