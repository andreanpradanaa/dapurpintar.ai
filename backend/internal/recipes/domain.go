package recipes

import (
	"encoding/json"
	"errors"
	"time"
)

var ErrNotFound = errors.New("recipe not found")

type Recipe struct {
	ID              string
	Title           string
	Summary         string
	Servings        int32
	PrepTimeMinutes *int32
	CookTimeMinutes *int32
	Ingredients     []RecipeIngredient
	Instructions    []string
	IsPublic        bool
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type RecipeIngredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type RecipeSummary struct {
	ID              string
	Title           string
	Summary         string
	Servings        int32
	PrepTimeMinutes *int32
	CookTimeMinutes *int32
}

type Favorite struct {
	ID string
	RecipeSummary
	FavoritedAt time.Time
}

func (r *Recipe) UnmarshalIngredients(data []byte) error { return json.Unmarshal(data, &r.Ingredients) }
func (r *Recipe) UnmarshalInstructions(data []byte) error {
	return json.Unmarshal(data, &r.Instructions)
}
