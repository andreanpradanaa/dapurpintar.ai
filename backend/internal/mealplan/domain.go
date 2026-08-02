package mealplan

import (
	"errors"
	"time"
)

type PlanStatus string

const (
	PlanDraft      PlanStatus = "draft"
	PlanPlanned    PlanStatus = "planned"
	PlanInProgress PlanStatus = "in_progress"
	PlanCompleted  PlanStatus = "completed"
	PlanCancelled  PlanStatus = "cancelled"
	PlanRevised    PlanStatus = "revised"
)

type MealStatus string

const (
	MealProposed  MealStatus = "proposed"
	MealPlanned   MealStatus = "planned"
	MealRevised   MealStatus = "revised"
	MealRemoved   MealStatus = "removed"
	MealCompleted MealStatus = "completed"
)

var ErrNotFound = errors.New("meal plan not found")

type MealPlan struct {
	ID            string
	UserProfileID string
	PeriodStart   time.Time
	PeriodEnd     time.Time
	Title         string
	Status        PlanStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PlannedMeal struct {
	ID                     string
	MealPlanID             string
	MealDate               time.Time
	MealOccasion           string
	RecipeID               *string
	RecommendationOptionID *string
	Status                 MealStatus
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
