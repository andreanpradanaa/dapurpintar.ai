package recommendation

import "context"

type Store interface {
	CreateRecommendation(ctx context.Context, profileID string, contextRef []byte, purpose string) (*Recommendation, error)
	GetRecommendationByID(ctx context.Context, id string) (*Recommendation, error)
	ListRecommendations(ctx context.Context, profileID, cursor string, limit int32, status, purpose *string) ([]Recommendation, error)
	UpdateRecommendationStatus(ctx context.Context, id string, status Status) (*Recommendation, error)

	CreateRecommendationOption(ctx context.Context, recID string, recipeID *string, position int32, rationale string) (*RecommendationOption, error)
	ListRecommendationOptions(ctx context.Context, recID string) ([]RecommendationOption, error)
	GetRecommendationOptionByID(ctx context.Context, id string) (*RecommendationOption, error)
	UpdateRecommendationOptionStatus(ctx context.Context, id string, status OptionStatus) (*RecommendationOption, error)
}
