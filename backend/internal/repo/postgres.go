package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dapurpintar/backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{pool: pool}
}

const recipeColumns = `id, slug, title, title_id, description, description_id,
	image, gradient, cuisine, difficulty, prep_time, cook_time, servings,
	ingredients, steps, nutrition, tags, dietary, rating, reviews, created_at`

func (r *PostgresRepo) List(ctx context.Context) ([]*model.Recipe, error) {
	rows, err := r.pool.Query(ctx, `SELECT `+recipeColumns+` FROM recipes ORDER BY created_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("query recipes: %w", err)
	}
	defer rows.Close()

	out := make([]*model.Recipe, 0, 32)
	for rows.Next() {
		r, err := scanRecipe(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

func (r *PostgresRepo) GetBySlug(ctx context.Context, slug string) (*model.Recipe, error) {
	row := r.pool.QueryRow(ctx, `SELECT `+recipeColumns+` FROM recipes WHERE slug = $1`, slug)
	rec, err := scanRecipe(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return rec, nil
}

func (r *PostgresRepo) Count(ctx context.Context) (int, error) {
	var n int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM recipes`).Scan(&n)
	return n, err
}

func (r *PostgresRepo) BulkInsert(ctx context.Context, recipes []*model.Recipe) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, rec := range recipes {
		ingredients, _ := json.Marshal(rec.Ingredients)
		steps, _ := json.Marshal(rec.Steps)
		nutrition, _ := json.Marshal(rec.Nutrition)

		_, err := tx.Exec(ctx, `
			INSERT INTO recipes (
				id, slug, title, title_id, description, description_id,
				image, gradient, cuisine, difficulty, prep_time, cook_time, servings,
				ingredients, steps, nutrition, tags, dietary, rating, reviews, created_at
			) VALUES (
				$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21
			)
			ON CONFLICT (id) DO UPDATE SET
				slug = EXCLUDED.slug,
				title = EXCLUDED.title,
				title_id = EXCLUDED.title_id,
				description = EXCLUDED.description,
				description_id = EXCLUDED.description_id,
				image = EXCLUDED.image,
				gradient = EXCLUDED.gradient,
				cuisine = EXCLUDED.cuisine,
				difficulty = EXCLUDED.difficulty,
				prep_time = EXCLUDED.prep_time,
				cook_time = EXCLUDED.cook_time,
				servings = EXCLUDED.servings,
				ingredients = EXCLUDED.ingredients,
				steps = EXCLUDED.steps,
				nutrition = EXCLUDED.nutrition,
				tags = EXCLUDED.tags,
				dietary = EXCLUDED.dietary,
				rating = EXCLUDED.rating,
				created_at = EXCLUDED.created_at
		`,
			rec.ID, rec.Slug, rec.Title, rec.TitleID, rec.Description, rec.DescriptionID,
			nullIfEmpty(rec.Image), rec.Gradient, rec.Cuisine, string(rec.Difficulty), rec.PrepTime, rec.CookTime, rec.Servings,
			ingredients, steps, nutrition, rec.Tags, dietaryToStrings(rec.Dietary), rec.Rating, rec.Reviews, rec.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("insert %s: %w", rec.ID, err)
		}
	}
	return tx.Commit(ctx)
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanRecipe(r rowScanner) (*model.Recipe, error) {
	var rec model.Recipe
	var image *string
	var ingredients, steps, nutrition []byte
	var dietary []string
	var createdAt time.Time

	err := r.Scan(
		&rec.ID, &rec.Slug, &rec.Title, &rec.TitleID, &rec.Description, &rec.DescriptionID,
		&image, &rec.Gradient, &rec.Cuisine, &rec.Difficulty, &rec.PrepTime, &rec.CookTime, &rec.Servings,
		&ingredients, &steps, &nutrition, &rec.Tags, &dietary, &rec.Rating, &rec.Reviews, &createdAt,
	)
	if err != nil {
		return nil, err
	}

	if image != nil {
		rec.Image = *image
	}
	rec.CreatedAt = createdAt.UTC().Format(time.RFC3339)
	if err := json.Unmarshal(ingredients, &rec.Ingredients); err != nil {
		return nil, fmt.Errorf("unmarshal ingredients: %w", err)
	}
	if err := json.Unmarshal(steps, &rec.Steps); err != nil {
		return nil, fmt.Errorf("unmarshal steps: %w", err)
	}
	if err := json.Unmarshal(nutrition, &rec.Nutrition); err != nil {
		return nil, fmt.Errorf("unmarshal nutrition: %w", err)
	}
	rec.Dietary = stringsToDietary(dietary)
	return &rec, nil
}

func nullIfEmpty(s string) any {
	if s == "" {
		return nil
	}
	return s
}

func dietaryToStrings(d []model.DietaryTag) []string {
	out := make([]string, 0, len(d))
	for _, t := range d {
		out = append(out, string(t))
	}
	return out
}

func stringsToDietary(s []string) []model.DietaryTag {
	out := make([]model.DietaryTag, 0, len(s))
	for _, t := range s {
		out = append(out, model.DietaryTag(t))
	}
	return out
}
