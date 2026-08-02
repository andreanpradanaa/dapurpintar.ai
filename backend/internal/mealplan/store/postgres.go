package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/gen/sqlc"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/mealplan"
)

type Postgres struct {
	db *sqlc.Queries
}

func New(conn sqlc.DBTX) *Postgres {
	return &Postgres{db: sqlc.New(conn)}
}

func (p *Postgres) CreateMealPlan(ctx context.Context, profileID string, periodStart, periodEnd time.Time, title string) (*mealplan.MealPlan, error) {
	row, err := p.db.CreateMealPlan(ctx, sqlc.CreateMealPlanParams{
		UserProfileID: profileID,
		PeriodStart:   periodStart,
		PeriodEnd:     periodEnd,
		Title:         title,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toMealPlan(row), nil
}

func (p *Postgres) GetMealPlanByID(ctx context.Context, id string) (*mealplan.MealPlan, error) {
	row, err := p.db.GetMealPlanByID(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toMealPlan(row), nil
}

func (p *Postgres) ListMealPlans(ctx context.Context, profileID, cursor string, limit int32) ([]mealplan.MealPlan, error) {
	rows, err := p.db.ListMealPlans(ctx, sqlc.ListMealPlansParams{
		UserProfileID: profileID,
		Column2:       cursor,
		Limit:         limit,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	out := make([]mealplan.MealPlan, len(rows))
	for i, r := range rows {
		out[i] = *toMealPlan(r)
	}
	return out, nil
}

func (p *Postgres) UpdateMealPlan(ctx context.Context, id string, title *string, periodStart, periodEnd *time.Time) (*mealplan.MealPlan, error) {
	row, err := p.db.UpdateMealPlan(ctx, sqlc.UpdateMealPlanParams{
		ID:          id,
		Title:       title,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toMealPlan(row), nil
}

func (p *Postgres) UpdateMealPlanStatus(ctx context.Context, id string, status mealplan.PlanStatus) (*mealplan.MealPlan, error) {
	row, err := p.db.UpdateMealPlanStatus(ctx, sqlc.UpdateMealPlanStatusParams{
		ID:     id,
		Status: string(status),
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toMealPlan(row), nil
}

func (p *Postgres) CreatePlannedMeal(ctx context.Context, planID string, mealDate time.Time, occasion string, recipeID, optionID *string) (*mealplan.PlannedMeal, error) {
	row, err := p.db.CreatePlannedMeal(ctx, sqlc.CreatePlannedMealParams{
		MealPlanID:             planID,
		MealDate:               mealDate,
		MealOccasion:           occasion,
		RecipeID:               recipeID,
		RecommendationOptionID: optionID,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toPlannedMeal(row), nil
}

func (p *Postgres) ListPlannedMeals(ctx context.Context, profileID, planID, cursor string, limit int32) ([]mealplan.PlannedMeal, error) {
	rows, err := p.db.ListPlannedMeals(ctx, sqlc.ListPlannedMealsParams{
		UserProfileID: profileID,
		MealPlanID:    planID,
		Column3:       cursor,
		Limit:         limit,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	out := make([]mealplan.PlannedMeal, len(rows))
	for i, r := range rows {
		out[i] = *toPlannedMeal(r)
	}
	return out, nil
}

func (p *Postgres) GetPlannedMealByID(ctx context.Context, id string) (*mealplan.PlannedMeal, error) {
	row, err := p.db.GetPlannedMealByID(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toPlannedMeal(row), nil
}

func (p *Postgres) UpdatePlannedMeal(ctx context.Context, id string, occasion, recipeID *string, status *string) (*mealplan.PlannedMeal, error) {
	row, err := p.db.UpdatePlannedMeal(ctx, sqlc.UpdatePlannedMealParams{
		ID:       id,
		Occasion: occasion,
		RecipeID: recipeID,
		Status:   status,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toPlannedMeal(row), nil
}

func (p *Postgres) RemovePlannedMeal(ctx context.Context, id string) (*mealplan.PlannedMeal, error) {
	row, err := p.db.RemovePlannedMeal(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toPlannedMeal(row), nil
}

func (p *Postgres) GetPlannedMealBySlot(ctx context.Context, planID string, mealDate time.Time, occasion string) (*mealplan.PlannedMeal, error) {
	row, err := p.db.GetPlannedMealBySlot(ctx, sqlc.GetPlannedMealBySlotParams{
		MealPlanID:   planID,
		MealDate:     mealDate,
		MealOccasion: occasion,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toPlannedMeal(row), nil
}

func toMealPlan(r sqlc.MealPlan) *mealplan.MealPlan {
	return &mealplan.MealPlan{
		ID:            r.ID,
		UserProfileID: r.UserProfileID,
		PeriodStart:   r.PeriodStart,
		PeriodEnd:     r.PeriodEnd,
		Title:         r.Title,
		Status:        mealplan.PlanStatus(r.Status),
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}
}

func toPlannedMeal(r sqlc.PlannedMeal) *mealplan.PlannedMeal {
	return &mealplan.PlannedMeal{
		ID:                     r.ID,
		MealPlanID:             r.MealPlanID,
		MealDate:               r.MealDate,
		MealOccasion:           r.MealOccasion,
		RecipeID:               r.RecipeID,
		RecommendationOptionID: r.RecommendationOptionID,
		Status:                 mealplan.MealStatus(r.Status),
		CreatedAt:              r.CreatedAt,
		UpdatedAt:              r.UpdatedAt,
	}
}

func mapErr(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == "no rows in result set" || err == pgx.ErrNoRows {
		return mealplan.ErrNotFound
	}
	return err
}
