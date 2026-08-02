package recipes

import (
	"context"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

const maxPageLimit = 100

type PageInfo struct {
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
}

func (s *Service) ListRecipes(ctx context.Context, cursor string, limit int32, query *string, maxPrep *int32, sortOrder *string) ([]RecipeSummary, *PageInfo, error) {
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	items, err := s.store.ListPublicRecipes(ctx, cursor, limit+1, query, maxPrep, sortOrder)
	if err != nil {
		return nil, nil, apperr.Internal(err)
	}
	hasMore := len(items) > int(limit)
	if hasMore {
		items = items[:limit]
	}
	var nextCursor string
	if len(items) > 0 {
		nextCursor = items[len(items)-1].ID
	}
	return items, &PageInfo{NextCursor: nextCursor, HasMore: hasMore}, nil
}

func (s *Service) GetRecipe(ctx context.Context, id string) (*Recipe, error) {
	recipe, err := s.store.GetRecipeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (s *Service) Favorite(ctx context.Context, profileID, recipeID string) error {
	if _, err := s.store.GetRecipeByID(ctx, recipeID); err != nil {
		return err
	}
	if _, err := s.store.CreateFavorite(ctx, profileID, recipeID); err != nil {
		return err
	}
	return nil
}

func (s *Service) Unfavorite(ctx context.Context, profileID, recipeID string) error {
	return s.store.RemoveFavorite(ctx, profileID, recipeID)
}

func (s *Service) ListFavorites(ctx context.Context, profileID, cursor string, limit int32) ([]Favorite, *PageInfo, error) {
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	items, err := s.store.ListFavorites(ctx, profileID, cursor, limit+1)
	if err != nil {
		return nil, nil, apperr.Internal(err)
	}
	hasMore := len(items) > int(limit)
	if hasMore {
		items = items[:limit]
	}
	var nextCursor string
	if len(items) > 0 {
		nextCursor = items[len(items)-1].ID
	}
	return items, &PageInfo{NextCursor: nextCursor, HasMore: hasMore}, nil
}
