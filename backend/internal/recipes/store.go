package recipes

import "context"

type Store interface {
	ListPublicRecipes(ctx context.Context, cursor string, limit int32, query *string, maxPrep *int32, sortOrder *string) ([]RecipeSummary, error)
	GetRecipeByID(ctx context.Context, id string) (*Recipe, error)
	GetActiveFavorite(ctx context.Context, profileID, recipeID string) (*Favorite, error)
	CreateFavorite(ctx context.Context, profileID, recipeID string) (*Favorite, error)
	RemoveFavorite(ctx context.Context, profileID, recipeID string) error
	ListFavorites(ctx context.Context, profileID, cursor string, limit int32) ([]Favorite, error)
}
