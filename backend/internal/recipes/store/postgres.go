package store

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/gen/sqlc"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recipes"
)

type Postgres struct {
	db *sqlc.Queries
}

func New(conn sqlc.DBTX) *Postgres {
	return &Postgres{db: sqlc.New(conn)}
}

func (p *Postgres) ListPublicRecipes(ctx context.Context, cursor string, limit int32, query *string, maxPrep *int32, sortOrder *string) ([]recipes.RecipeSummary, error) {
	rows, err := p.db.ListPublicRecipes(ctx, sqlc.ListPublicRecipesParams{
		Column1:   cursor,
		Limit:     limit,
		Q:         query,
		MaxPrep:   int32PtrToPgtype(maxPrep),
		SortOrder: sortOrder,
	})
	if err != nil {
		return nil, mapRecipeErr(err)
	}
	out := make([]recipes.RecipeSummary, len(rows))
	for i, r := range rows {
		out[i] = toRecipeSummary(r)
	}
	return out, nil
}

func (p *Postgres) GetRecipeByID(ctx context.Context, id string) (*recipes.Recipe, error) {
	row, err := p.db.GetRecipeByID(ctx, sqlc.GetRecipeByIDParams{
		ID:             id,
		IncludePrivate: pgtype.Bool{Bool: true, Valid: true},
	})
	if err != nil {
		return nil, mapRecipeErr(err)
	}
	return toRecipe(row), nil
}

func (p *Postgres) GetActiveFavorite(ctx context.Context, profileID, recipeID string) (*recipes.Favorite, error) {
	row, err := p.db.GetActiveFavorite(ctx, sqlc.GetActiveFavoriteParams{
		UserProfileID: profileID,
		RecipeID:      recipeID,
	})
	if err != nil {
		return nil, mapRecipeErr(err)
	}
	return &recipes.Favorite{ID: row.ID, FavoritedAt: row.CreatedAt}, nil
}

func (p *Postgres) CreateFavorite(ctx context.Context, profileID, recipeID string) (*recipes.Favorite, error) {
	row, err := p.db.CreateFavorite(ctx, sqlc.CreateFavoriteParams{
		UserProfileID: profileID,
		RecipeID:      recipeID,
	})
	if err != nil {
		if err.Error() == "no rows in result set" {
			return p.GetActiveFavorite(ctx, profileID, recipeID)
		}
		return nil, mapRecipeErr(err)
	}
	return &recipes.Favorite{ID: row.ID, FavoritedAt: row.CreatedAt}, nil
}

func (p *Postgres) RemoveFavorite(ctx context.Context, profileID, recipeID string) error {
	_, err := p.db.RemoveFavorite(ctx, sqlc.RemoveFavoriteParams{
		UserProfileID: profileID,
		RecipeID:      recipeID,
	})
	if err != nil && err.Error() != "no rows in result set" {
		return mapRecipeErr(err)
	}
	return nil
}

func (p *Postgres) ListFavorites(ctx context.Context, profileID, cursor string, limit int32) ([]recipes.Favorite, error) {
	rows, err := p.db.ListFavorites(ctx, sqlc.ListFavoritesParams{
		UserProfileID: profileID,
		Column2:       cursor,
		Limit:         limit,
	})
	if err != nil {
		return nil, mapRecipeErr(err)
	}
	out := make([]recipes.Favorite, len(rows))
	for i, row := range rows {
		out[i] = recipes.Favorite{
			ID: row.ID,
			RecipeSummary: recipes.RecipeSummary{
				ID:              row.RID,
				Title:           row.RTitle,
				Summary:         row.RSummary,
				Servings:        row.RServings,
				PrepTimeMinutes: int4Ptr(row.RPrepTimeMinutes),
				CookTimeMinutes: int4Ptr(row.RCookTimeMinutes),
			},
			FavoritedAt: row.CreatedAt,
		}
	}
	return out, nil
}

func toRecipeSummary(r sqlc.Recipe) recipes.RecipeSummary {
	return recipes.RecipeSummary{
		ID:              r.ID,
		Title:           r.Title,
		Summary:         r.Summary,
		Servings:        r.Servings,
		PrepTimeMinutes: int4Ptr(r.PrepTimeMinutes),
		CookTimeMinutes: int4Ptr(r.CookTimeMinutes),
	}
}

func toRecipe(r sqlc.Recipe) *recipes.Recipe {
	recipe := &recipes.Recipe{
		ID:              r.ID,
		Title:           r.Title,
		Summary:         r.Summary,
		Servings:        r.Servings,
		PrepTimeMinutes: int4Ptr(r.PrepTimeMinutes),
		CookTimeMinutes: int4Ptr(r.CookTimeMinutes),
		IsPublic:        r.IsPublic,
		Status:          r.Status,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}
	if r.Ingredients != nil {
		json.Unmarshal(r.Ingredients, &recipe.Ingredients)
	}
	if r.Instructions != nil {
		json.Unmarshal(r.Instructions, &recipe.Instructions)
	}
	return recipe
}

func int4Ptr(v pgtype.Int4) *int32 {
	if !v.Valid {
		return nil
	}
	c := v.Int32
	return &c
}

func int32PtrToPgtype(v *int32) pgtype.Int4 {
	if v == nil {
		return pgtype.Int4{Valid: false}
	}
	return pgtype.Int4{Int32: *v, Valid: true}
}

func mapRecipeErr(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == "no rows in result set" || err == pgx.ErrNoRows {
		return recipes.ErrNotFound
	}
	return err
}
