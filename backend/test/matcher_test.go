package test

import (
	"testing"

	"github.com/dapurpintar/backend/internal/model"
	"github.com/dapurpintar/backend/internal/service"
)

func mkRecipe() *model.Recipe {
	return &model.Recipe{
		ID:    "r-001",
		Slug:  "test",
		Title: "Nasi Goreng",
		Tags:  []string{"rice", "wok"},
		Ingredients: []model.RecipeIngredient{
			{Name: "Cooked rice", NameID: "Nasi putih", Amount: "2 cups", Category: model.CategoryGrain},
			{Name: "Chicken", NameID: "Ayam", Amount: "150g", Category: model.CategoryProtein},
		},
		Difficulty: model.DifficultyEasy,
		Rating:     4.5,
	}
}

func TestScore_ExactMatch(t *testing.T) {
	r := mkRecipe()
	s := service.Score(r, []string{"chicken", "rice"}, nil)
	if s <= 0 {
		t.Fatalf("expected positive score, got %d", s)
	}
}

func TestScore_DietaryMatch(t *testing.T) {
	r := mkRecipe()
	r.Dietary = []model.DietaryTag{model.DietaryHalal}
	s := service.Score(r, []string{"rice"}, []string{"halal"})
	if s <= 5 {
		t.Fatalf("expected dietary bonus, got %d", s)
	}
}

func TestScore_NoIngredients(t *testing.T) {
	r := mkRecipe()
	if service.Score(r, nil, nil) != 0 {
		t.Fatal("expected zero score with no ingredients")
	}
}

func TestRank_OrderIsDescending(t *testing.T) {
	a := mkRecipe()
	a.Title = "Nasi Goreng"
	a.Rating = 4.9
	b := mkRecipe()
	b.ID = "r-002"
	b.Title = "Soto Ayam"
	b.Rating = 4.0
	c := mkRecipe()
	c.ID = "r-003"
	c.Title = "Mie Goreng"
	c.Rating = 3.5

	ranked := service.Rank(nil, []*model.Recipe{a, b, c}, []string{"nasi", "goreng"}, nil)
	if len(ranked) != 3 {
		t.Fatalf("expected 3 results, got %d", len(ranked))
	}
	if ranked[0].Recipe.ID != "r-001" {
		t.Fatalf("expected r-001 first, got %s", ranked[0].Recipe.ID)
	}
	if ranked[0].Score < ranked[1].Score {
		t.Fatal("not in descending order")
	}
}

func TestNormalizeIngredients(t *testing.T) {
	in := []string{"  Chicken ", "", " GARLIC  ", "rice"}
	out := service.NormalizeIngredients(in)
	if len(out) != 3 {
		t.Fatalf("expected 3 normalized, got %d", len(out))
	}
	if out[0] != "chicken" {
		t.Fatalf("expected lowercase, got %q", out[0])
	}
}
