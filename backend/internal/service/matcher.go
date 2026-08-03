package service

import (
	"context"
	"strings"

	"github.com/dapurpintar/backend/internal/model"
)

// Scored pairs a recipe with its relevance score.
type Scored struct {
	Recipe *model.Recipe
	Score  int
}

// Score returns a relevance score for a recipe against the user's
// ingredients and dietary preferences. Mirrors the scoring algorithm
// that previously lived in frontend/lib/generate.ts.
func Score(recipe *model.Recipe, userIngredients []string, dietary []string) int {
	if recipe == nil {
		return 0
	}
	lower := make([]string, 0, len(userIngredients))
	for _, i := range userIngredients {
		t := strings.ToLower(strings.TrimSpace(i))
		if t != "" {
			lower = append(lower, t)
		}
	}
	if len(lower) == 0 {
		return 0
	}

	// Build a haystack of all recipe text for substring matching.
	haystack := strings.ToLower(strings.Join(
		append([]string{recipe.Title, recipe.TitleID, recipe.Description, recipe.DescriptionID, recipe.Cuisine},
			append(recipe.Tags, ingredientNames(recipe.Ingredients)...)...),
		" ",
	))

	score := 0
	for _, ing := range lower {
		if strings.Contains(haystack, ing) {
			score += 3
		}
		for _, w := range strings.Fields(ing) {
			if len(w) >= 3 && strings.Contains(haystack, w) {
				score++
			}
		}
	}

	// Dietary alignment
	if len(dietary) > 0 {
		matches := 0
		for _, d := range dietary {
			for _, rd := range recipe.Dietary {
				if strings.EqualFold(string(rd), d) {
					matches++
					break
				}
			}
		}
		score += matches * 2
		if matches == len(dietary) {
			score += 5
		}
	}

	// Popularity bias
	score += int(recipe.Rating * 10)

	return score
}

// Rank scores and sorts recipes by relevance (highest first).
func Rank(ctx context.Context, recipes []*model.Recipe, ingredients []string, dietary []string) []Scored {
	out := make([]Scored, 0, len(recipes))
	for _, r := range recipes {
		out = append(out, Scored{
			Recipe: r,
			Score:  Score(r, ingredients, dietary),
		})
	}
	// Sort descending by score (insertion sort — small N, ~32 recipes)
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j].Score > out[j-1].Score; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out
}

func NormalizeIngredients(input []string) []string {
	out := make([]string, 0, len(input))
	for _, i := range input {
		t := strings.ToLower(strings.TrimSpace(i))
		if t != "" {
			out = append(out, t)
		}
	}
	return out
}

func ingredientNames(ings []model.RecipeIngredient) []string {
	out := make([]string, 0, len(ings)*2)
	for _, i := range ings {
		out = append(out, i.Name, i.NameID)
	}
	return out
}
