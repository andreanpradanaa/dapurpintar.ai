package http

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recipes"
)

type recipeSummaryView struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Summary         string `json:"summary"`
	Servings        int32  `json:"servings"`
	PrepTimeMinutes *int32 `json:"prep_time_minutes"`
	CookTimeMinutes *int32 `json:"cook_time_minutes"`
}

type recipeDetailView struct {
	ID              string                     `json:"id"`
	Title           string                     `json:"title"`
	Summary         string                     `json:"summary"`
	Servings        int32                      `json:"servings"`
	PrepTimeMinutes *int32                     `json:"prep_time_minutes"`
	CookTimeMinutes *int32                     `json:"cook_time_minutes"`
	Ingredients     []recipes.RecipeIngredient `json:"ingredients"`
	Instructions    []string                   `json:"instructions"`
	IsPublic        bool                       `json:"is_public"`
}

type favoriteView struct {
	Recipe    recipeSummaryView `json:"recipe"`
	CreatedAt string            `json:"created_at"`
}

func toRecipeSummaryView(r *recipes.RecipeSummary) recipeSummaryView {
	return recipeSummaryView{
		ID:              r.ID,
		Title:           r.Title,
		Summary:         r.Summary,
		Servings:        r.Servings,
		PrepTimeMinutes: r.PrepTimeMinutes,
		CookTimeMinutes: r.CookTimeMinutes,
	}
}

func toRecipeDetailView(r *recipes.Recipe) recipeDetailView {
	return recipeDetailView{
		ID:              r.ID,
		Title:           r.Title,
		Summary:         r.Summary,
		Servings:        r.Servings,
		PrepTimeMinutes: r.PrepTimeMinutes,
		CookTimeMinutes: r.CookTimeMinutes,
		Ingredients:     r.Ingredients,
		Instructions:    r.Instructions,
		IsPublic:        r.IsPublic,
	}
}

func toFavoriteView(f *recipes.Favorite) favoriteView {
	return favoriteView{
		Recipe:    toRecipeSummaryView(&f.RecipeSummary),
		CreatedAt: f.FavoritedAt.UTC().Format(time.RFC3339),
	}
}

// listRecipes handles GET /api/v1/recipes (public).
func (h *Handler) listRecipes(c *fiber.Ctx) error {
	cursor := c.Query("cursor")
	limit := parseLimit(c.Query("limit", "20"))
	query := queryPointer(c.Query("q"))
	sortOrder := queryPointer(c.Query("sort", "relevance"))

	var maxPrep *int32
	if mp := c.Query("max_prep_minutes"); mp != "" {
		if v, err := parseI32(mp); err == nil && v >= 0 {
			maxPrep = &v
		}
	}

	items, page, err := h.recipes.ListRecipes(c.Context(), cursor, limit, query, maxPrep, sortOrder)
	if err != nil {
		return response.Error(c, err)
	}

	views := make([]recipeSummaryView, len(items))
	for i, it := range items {
		views[i] = toRecipeSummaryView(&it)
	}
	return response.OK(c, map[string]any{
		"data": views,
		"page": page,
	})
}

// getRecipe handles GET /api/v1/recipes/:recipeId (public).
func (h *Handler) getRecipe(c *fiber.Ctx) error {
	recipe, err := h.recipes.GetRecipe(c.Context(), c.Params("recipeId"))
	if err != nil {
		return response.Error(c, err)
	}
	if !recipe.IsPublic {
		return response.Error(c, errRecipeNotPublic)
	}
	return response.OK(c, toRecipeDetailView(recipe))
}

// listFavorites handles GET /api/v1/favorites.
func (h *Handler) listFavorites(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	cursor := c.Query("cursor")
	limit := parseLimit(c.Query("limit", "20"))

	items, page, err := h.recipes.ListFavorites(c.Context(), profile.ID, cursor, limit)
	if err != nil {
		return response.Error(c, err)
	}

	views := make([]favoriteView, len(items))
	for i, f := range items {
		views[i] = toFavoriteView(&f)
	}
	return response.OK(c, map[string]any{
		"data": views,
		"page": page,
	})
}

// favoriteRecipe handles PUT /api/v1/favorites/recipes/:recipeId.
func (h *Handler) favoriteRecipe(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	if err := h.recipes.Favorite(c.Context(), profile.ID, c.Params("recipeId")); err != nil {
		return response.Error(c, err)
	}
	return response.NoContent(c)
}

// unfavoriteRecipe handles DELETE /api/v1/favorites/recipes/:recipeId.
func (h *Handler) unfavoriteRecipe(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	if err := h.recipes.Unfavorite(c.Context(), profile.ID, c.Params("recipeId")); err != nil {
		return response.Error(c, err)
	}
	return response.NoContent(c)
}

func parseI32(s string) (int32, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}
