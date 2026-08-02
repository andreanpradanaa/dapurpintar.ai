package mealplan

import (
	"context"
	"time"

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

func (s *Service) CreateMealPlan(ctx context.Context, profileID string, periodStart, periodEnd time.Time, title string) (*MealPlan, error) {
	if !periodEnd.After(periodStart) {
		return nil, apperr.New(apperr.CodeMealPlanPeriodInv, "Period end must be after period start.").
			WithDetails(apperr.Detail{Field: "period_end", Code: string(apperr.CodeMealPlanPeriodInv), Message: "Period end must be after period start."})
	}
	return s.store.CreateMealPlan(ctx, profileID, periodStart, periodEnd, title)
}

func (s *Service) GetMealPlan(ctx context.Context, id string) (*MealPlan, error) {
	return s.store.GetMealPlanByID(ctx, id)
}

func (s *Service) ListMealPlans(ctx context.Context, profileID, cursor string, limit int32) ([]MealPlan, *PageInfo, error) {
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	items, err := s.store.ListMealPlans(ctx, profileID, cursor, limit+1)
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

func (s *Service) UpdateMealPlan(ctx context.Context, id string, title *string, periodStart, periodEnd *time.Time) (*MealPlan, error) {
	if periodStart != nil && periodEnd != nil && !periodEnd.After(*periodStart) {
		return nil, apperr.New(apperr.CodeMealPlanPeriodInv, "Period end must be after period start.").
			WithDetails(apperr.Detail{Field: "period_end", Code: string(apperr.CodeMealPlanPeriodInv), Message: "Period end must be after period start."})
	}
	return s.store.UpdateMealPlan(ctx, id, title, periodStart, periodEnd)
}

func (s *Service) CancelMealPlan(ctx context.Context, id string) (*MealPlan, error) {
	plan, err := s.store.GetMealPlanByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if plan.Status == PlanCompleted {
		return nil, apperr.New(apperr.CodeMealSlotConflict, "A completed meal plan cannot be cancelled.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeMealSlotConflict), Message: "Completed meal plans cannot be cancelled."})
	}
	return s.store.UpdateMealPlanStatus(ctx, id, PlanCancelled)
}

func (s *Service) CompleteMealPlan(ctx context.Context, id string) (*MealPlan, error) {
	plan, err := s.store.GetMealPlanByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if plan.Status == PlanCancelled {
		return nil, apperr.New(apperr.CodeMealSlotConflict, "A cancelled meal plan cannot be completed.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeMealSlotConflict), Message: "Cancelled meal plans cannot be completed."})
	}
	return s.store.UpdateMealPlanStatus(ctx, id, PlanCompleted)
}

func (s *Service) PlanMeal(ctx context.Context, planID string, mealDate time.Time, occasion string, recipeID, optionID *string) (*PlannedMeal, error) {
	if !validOccasion(occasion) {
		return nil, apperr.New(apperr.CodeFieldInvalid, "Meal occasion is invalid.").
			WithDetails(apperr.Detail{Field: "meal_occasion", Code: string(apperr.CodeFieldInvalid), Message: "Meal occasion must be breakfast, lunch, dinner, or snack."})
	}

	plan, err := s.store.GetMealPlanByID(ctx, planID)
	if err != nil {
		return nil, err
	}
	if mealDate.Before(plan.PeriodStart) || mealDate.After(plan.PeriodEnd) {
		return nil, apperr.New(apperr.CodeMealPlanPeriodInv, "Meal date is outside the plan period.").
			WithDetails(apperr.Detail{Field: "meal_date", Code: string(apperr.CodeMealPlanPeriodInv), Message: "Meal date is outside the plan period."})
	}

	existing, err := s.store.GetPlannedMealBySlot(ctx, planID, mealDate, occasion)
	if err == nil && existing != nil {
		return nil, apperr.New(apperr.CodeMealSlotConflict, "This meal slot is already taken.").
			WithDetails(apperr.Detail{Field: "meal_date", Code: string(apperr.CodeMealSlotConflict), Message: "A meal is already planned for this date and occasion."})
	}

	return s.store.CreatePlannedMeal(ctx, planID, mealDate, occasion, recipeID, optionID)
}

func (s *Service) ListPlannedMeals(ctx context.Context, profileID, planID, cursor string, limit int32) ([]PlannedMeal, *PageInfo, error) {
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	items, err := s.store.ListPlannedMeals(ctx, profileID, planID, cursor, limit+1)
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

func (s *Service) GetPlannedMeal(ctx context.Context, id string) (*PlannedMeal, error) {
	return s.store.GetPlannedMealByID(ctx, id)
}

func (s *Service) UpdatePlannedMeal(ctx context.Context, id string, occasion, recipeID *string, status *string) (*PlannedMeal, error) {
	if occasion != nil && !validOccasion(*occasion) {
		return nil, apperr.New(apperr.CodeFieldInvalid, "Meal occasion is invalid.").
			WithDetails(apperr.Detail{Field: "meal_occasion", Code: string(apperr.CodeFieldInvalid), Message: "Meal occasion must be breakfast, lunch, dinner, or snack."})
	}
	return s.store.UpdatePlannedMeal(ctx, id, occasion, recipeID, status)
}

func (s *Service) RemovePlannedMeal(ctx context.Context, id string) error {
	_, err := s.store.RemovePlannedMeal(ctx, id)
	return err
}

func validOccasion(o string) bool {
	switch o {
	case "breakfast", "lunch", "dinner", "snack":
		return true
	}
	return false
}
