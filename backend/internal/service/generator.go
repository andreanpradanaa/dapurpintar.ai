package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dapurpintar/backend/internal/model"
	"github.com/dapurpintar/backend/internal/repo"
	"github.com/dapurpintar/backend/internal/service/llm"
	"github.com/rs/zerolog"
)

type Generator struct {
	repo  repo.RecipeRepo
	llm   llm.Client
	log   *zerolog.Logger
	limit int
}

func NewGenerator(r repo.RecipeRepo, l llm.Client, log *zerolog.Logger) *Generator {
	return &Generator{repo: r, llm: l, log: log, limit: 3}
}

type GenerateRequest struct {
	Ingredients []string `json:"ingredients"`
	Dietary     []string `json:"dietary"`
	Language    string   `json:"language"`
}

type GenerateResponse struct {
	Recipe       *model.Recipe  `json:"recipe"`
	MatchScore   int            `json:"matchScore"`
	Sources      []RecipeRef    `json:"sources"`
	Alternatives []RecipeRef    `json:"alternatives"`
}

type RecipeRef struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Title  string `json:"title"`
	Score  int    `json:"score"`
}

// Generate is the RAG orchestrator: it scores the curated library,
// takes the top references, and asks the LLM to compose a fresh
// recipe based on them. The library stays in Postgres as the source
// of style truth.
func (g *Generator) Generate(ctx context.Context, req GenerateRequest) (*GenerateResponse, error) {
	if len(req.Ingredients) == 0 {
		return nil, ErrEmptyIngredients
	}

	recipes, err := g.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list recipes: %w", err)
	}
	if len(recipes) == 0 {
		return nil, errors.New("recipe library is empty")
	}

	ranked := Rank(ctx, recipes, req.Ingredients, req.Dietary)

	// Take top N as RAG references
	refCount := g.limit
	if len(ranked) < refCount {
		refCount = len(ranked)
	}
	references := make([]*model.Recipe, 0, refCount)
	for _, s := range ranked[:refCount] {
		references = append(references, s.Recipe)
	}

	// Build alternatives from the next few (without score 0)
	altCount := 3
	if len(ranked)-refCount < altCount {
		altCount = len(ranked) - refCount
	}
	if altCount < 0 {
		altCount = 0
	}
	alternatives := make([]RecipeRef, 0, altCount)
	for _, s := range ranked[refCount : refCount+altCount] {
		if s.Score <= 0 {
			continue
		}
		alternatives = append(alternatives, RecipeRef{
			ID: s.Recipe.ID, Slug: s.Recipe.Slug, Title: s.Recipe.Title, Score: s.Score,
		})
	}

	// Sources are the references (used as LLM context)
	sources := make([]RecipeRef, 0, len(references))
	for _, r := range references {
		sources = append(sources, RecipeRef{
			ID: r.ID, Slug: r.Slug, Title: r.Title,
		})
	}

	// Call LLM
	llmReq := llm.GenerateRequest{
		UserIngredients: req.Ingredients,
		Dietary:         req.Dietary,
		Language:        langOrDefault(req.Language),
		References:      references,
		Timeout:         120 * time.Second,
	}
	g.log.Info().
		Str("provider", g.llm.Name()).
		Int("references", len(references)).
		Int("ingredients", len(req.Ingredients)).
		Msg("calling llm")
	recipe, err := g.llm.GenerateRecipe(ctx, llmReq)
	if err != nil {
		return nil, fmt.Errorf("llm: %w", err)
	}

	// Match score = top reference's score
	matchScore := 0
	if len(ranked) > 0 {
		matchScore = ranked[0].Score
	}

	return &GenerateResponse{
		Recipe:       recipe,
		MatchScore:   matchScore,
		Sources:      sources,
		Alternatives: alternatives,
	}, nil
}

var ErrEmptyIngredients = errors.New("at least one ingredient is required")

func langOrDefault(l string) string {
	if l == "id" {
		return "id"
	}
	return "en"
}
