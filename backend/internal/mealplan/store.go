package mealplan

import (
	"context"
	"time"
)

type Store interface {
	CreateMealPlan(ctx context.Context, profileID string, periodStart, periodEnd time.Time, title string) (*MealPlan, error)
	GetMealPlanByID(ctx context.Context, id string) (*MealPlan, error)
	ListMealPlans(ctx context.Context, profileID, cursor string, limit int32) ([]MealPlan, error)
	UpdateMealPlan(ctx context.Context, id string, title *string, periodStart, periodEnd *time.Time) (*MealPlan, error)
	UpdateMealPlanStatus(ctx context.Context, id string, status PlanStatus) (*MealPlan, error)

	CreatePlannedMeal(ctx context.Context, planID string, mealDate time.Time, occasion string, recipeID, optionID *string) (*PlannedMeal, error)
	ListPlannedMeals(ctx context.Context, profileID, planID, cursor string, limit int32) ([]PlannedMeal, error)
	GetPlannedMealByID(ctx context.Context, id string) (*PlannedMeal, error)
	UpdatePlannedMeal(ctx context.Context, id string, occasion, recipeID *string, status *string) (*PlannedMeal, error)
	RemovePlannedMeal(ctx context.Context, id string) (*PlannedMeal, error)
	GetPlannedMealBySlot(ctx context.Context, planID string, mealDate time.Time, occasion string) (*PlannedMeal, error)
}
