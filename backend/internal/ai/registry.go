package ai

import (
	"fmt"
	"sort"
	"sync"
	"time"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// Registry is the versioned store of policy bundles per purpose (M4-DEC-011).
// It owns the promotion gate: only a Promoted revision is served to live
// requests, while Pending revisions are available for review and evaluation.
// It is safe for concurrent use.
type Registry struct {
	mu      sync.RWMutex
	bundles map[Purpose][]PolicyBundle // ordered oldest first
}

// NewRegistry creates an empty policy registry.
func NewRegistry() *Registry {
	return &Registry{bundles: make(map[Purpose][]PolicyBundle)}
}

// Register adds a new immutable revision for a purpose. A purpose's revisions
// must use distinct schema revisions; re-registering the same schema revision
// is rejected to prevent silent overwrite. The new revision is Pending unless
// the caller marks it Promoted explicitly.
func (r *Registry) Register(b PolicyBundle) error {
	if err := b.Validate(); err != nil {
		return err
	}
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now().UTC()
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, existing := range r.bundles[b.Purpose] {
		if existing.SchemaRevision == b.SchemaRevision {
			return apperr.New(apperr.CodeAIUnavailable,
				fmt.Sprintf("AI policy revision %q already registered for purpose %q.",
					b.SchemaRevision, b.Purpose))
		}
	}

	r.bundles[b.Purpose] = append(r.bundles[b.Purpose], b)
	sort.Slice(r.bundles[b.Purpose], func(i, j int) bool {
		return r.bundles[b.Purpose][i].CreatedAt.Before(r.bundles[b.Purpose][j].CreatedAt)
	})
	return nil
}

// Promote marks a pending revision as the promoted revision for its purpose.
// Any previously promoted revision for the purpose is demoted so there is
// exactly one active revision.
func (r *Registry) Promote(purpose Purpose, schemaRevision string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	entries, ok := r.bundles[purpose]
	if !ok {
		return apperr.New(apperr.CodeAIUnavailable,
			fmt.Sprintf("AI policy for purpose %q is not registered.", purpose))
	}

	found := false
	for i := range entries {
		if entries[i].SchemaRevision == schemaRevision {
			found = true
			entries[i].Status = RevisionPromoted
		} else {
			entries[i].Status = RevisionPending
		}
	}
	if !found {
		return apperr.New(apperr.CodeAIUnavailable,
			fmt.Sprintf("AI policy revision %q not found for purpose %q.", schemaRevision, purpose))
	}
	return nil
}

// Resolve returns the bundle for a purpose. When schemaRevision is empty, the
// promoted revision is returned; otherwise the matching revision is returned
// regardless of promotion status (for evaluation and rollback analysis).
func (r *Registry) Resolve(purpose Purpose, schemaRevision string) (PolicyBundle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entries, ok := r.bundles[purpose]
	if !ok {
		return PolicyBundle{}, apperr.New(apperr.CodeAIUnavailable,
			fmt.Sprintf("AI policy for purpose %q is not registered.", purpose))
	}

	if schemaRevision == "" {
		for i := len(entries) - 1; i >= 0; i-- {
			if entries[i].Status == RevisionPromoted {
				return entries[i], nil
			}
		}
		return PolicyBundle{}, apperr.New(apperr.CodeAIUnavailable,
			fmt.Sprintf("AI policy for purpose %q has no promoted revision.", purpose))
	}

	for _, b := range entries {
		if b.SchemaRevision == schemaRevision {
			return b, nil
		}
	}
	return PolicyBundle{}, apperr.New(apperr.CodeAIUnavailable,
		fmt.Sprintf("AI policy revision %q not found for purpose %q.", schemaRevision, purpose))
}
