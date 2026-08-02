package store

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/gen/sqlc"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recommendation"
)

type Postgres struct {
	db *sqlc.Queries
}

func New(conn sqlc.DBTX) *Postgres {
	return &Postgres{db: sqlc.New(conn)}
}

func (p *Postgres) CreateRecommendation(ctx context.Context, profileID string, contextRef []byte, purpose string) (*recommendation.Recommendation, error) {
	row, err := p.db.CreateRecommendation(ctx, sqlc.CreateRecommendationParams{
		UserProfileID:    profileID,
		ContextReference: contextRef,
		Purpose:          purpose,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toRecommendation(row), nil
}

func (p *Postgres) GetRecommendationByID(ctx context.Context, id string) (*recommendation.Recommendation, error) {
	row, err := p.db.GetRecommendationByID(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toRecommendation(row), nil
}

func (p *Postgres) ListRecommendations(ctx context.Context, profileID, cursor string, limit int32, status, purpose *string) ([]recommendation.Recommendation, error) {
	rows, err := p.db.ListRecommendations(ctx, sqlc.ListRecommendationsParams{
		UserProfileID: profileID,
		Column2:       cursor,
		Limit:         limit,
		Status:        status,
		Purpose:       purpose,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	out := make([]recommendation.Recommendation, len(rows))
	for i, r := range rows {
		out[i] = *toRecommendation(r)
	}
	return out, nil
}

func (p *Postgres) UpdateRecommendationStatus(ctx context.Context, id string, status recommendation.Status) (*recommendation.Recommendation, error) {
	row, err := p.db.UpdateRecommendationStatus(ctx, sqlc.UpdateRecommendationStatusParams{
		ID:     id,
		Status: string(status),
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toRecommendation(row), nil
}

func (p *Postgres) CreateRecommendationOption(ctx context.Context, recID string, recipeID *string, position int32, rationale string) (*recommendation.RecommendationOption, error) {
	row, err := p.db.CreateRecommendationOption(ctx, sqlc.CreateRecommendationOptionParams{
		RecommendationID: recID,
		RecipeID:         recipeID,
		Position:         position,
		Rationale:        rationale,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toOption(row), nil
}

func (p *Postgres) ListRecommendationOptions(ctx context.Context, recID string) ([]recommendation.RecommendationOption, error) {
	rows, err := p.db.ListRecommendationOptions(ctx, recID)
	if err != nil {
		return nil, mapErr(err)
	}
	out := make([]recommendation.RecommendationOption, len(rows))
	for i, r := range rows {
		out[i] = *toOption(r)
	}
	return out, nil
}

func (p *Postgres) GetRecommendationOptionByID(ctx context.Context, id string) (*recommendation.RecommendationOption, error) {
	row, err := p.db.GetRecommendationOptionByID(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toOption(row), nil
}

func (p *Postgres) UpdateRecommendationOptionStatus(ctx context.Context, id string, status recommendation.OptionStatus) (*recommendation.RecommendationOption, error) {
	row, err := p.db.UpdateRecommendationOptionStatus(ctx, sqlc.UpdateRecommendationOptionStatusParams{
		ID:     id,
		Status: string(status),
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toOption(row), nil
}

func toRecommendation(r sqlc.KitchenRecommendation) *recommendation.Recommendation {
	return &recommendation.Recommendation{
		ID:                  r.ID,
		UserProfileID:       r.UserProfileID,
		ContextReference:    r.ContextReference,
		Purpose:             r.Purpose,
		Rationale:           r.Rationale,
		ConfidenceStatement: r.ConfidenceStatement,
		Status:              recommendation.Status(r.Status),
		CreatedAt:           r.CreatedAt,
		UpdatedAt:           r.UpdatedAt,
	}
}

func toOption(r sqlc.RecommendationOption) *recommendation.RecommendationOption {
	return &recommendation.RecommendationOption{
		ID:               r.ID,
		RecommendationID: r.RecommendationID,
		RecipeID:         r.RecipeID,
		Position:         r.Position,
		Rationale:        r.Rationale,
		Status:           recommendation.OptionStatus(r.Status),
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
	}
}

func mapErr(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == "no rows in result set" || err == pgx.ErrNoRows {
		return recommendation.ErrNotFound
	}
	return err
}

func (p *Postgres) CreateConversation(ctx context.Context, recID string) (*recommendation.Conversation, error) {
	row, err := p.db.CreateConversation(ctx, recID)
	if err != nil {
		return nil, mapErr(err)
	}
	return toConversation(row), nil
}

func (p *Postgres) GetConversationByRecommendation(ctx context.Context, recID string) (*recommendation.Conversation, error) {
	row, err := p.db.GetConversationByRecommendation(ctx, recID)
	if err != nil {
		return nil, mapErr(err)
	}
	return toConversation(row), nil
}

func (p *Postgres) UpdateConversationSnapshot(ctx context.Context, recID string, snapshot []byte) (*recommendation.Conversation, error) {
	row, err := p.db.UpdateConversationSnapshot(ctx, sqlc.UpdateConversationSnapshotParams{
		RecommendationID: recID,
		ContextSnapshot:  snapshot,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toConversation(row), nil
}

func (p *Postgres) CloseConversation(ctx context.Context, recID string) (*recommendation.Conversation, error) {
	row, err := p.db.CloseConversation(ctx, recID)
	if err != nil {
		return nil, mapErr(err)
	}
	return toConversation(row), nil
}

func toConversation(r sqlc.RecommendationConversation) *recommendation.Conversation {
	conv := &recommendation.Conversation{
		ID:               r.ID,
		RecommendationID: r.RecommendationID,
		Status:           r.Status,
		ExpiresAt:        r.ExpiresAt,
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
	}
	if r.ContextSnapshot != nil {
		var msgs []recommendation.ConversationMessage
		if err := json.Unmarshal(r.ContextSnapshot, &msgs); err == nil {
			conv.Messages = msgs
		}
	}
	return conv
}
