package http

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/mealplan"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
)

type mealPlanView struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	PeriodStart string `json:"period_start"`
	PeriodEnd   string `json:"period_end"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

type plannedMealView struct {
	ID                     string  `json:"id"`
	MealPlanID             string  `json:"meal_plan_id"`
	MealDate               string  `json:"meal_date"`
	MealOccasion           string  `json:"meal_occasion"`
	RecipeID               *string `json:"recipe_id"`
	RecommendationOptionID *string `json:"recommendation_option_id"`
	Status                 string  `json:"status"`
}

func toMealPlanView(p *mealplan.MealPlan) mealPlanView {
	return mealPlanView{
		ID:          p.ID,
		Title:       p.Title,
		PeriodStart: p.PeriodStart.Format("2006-01-02"),
		PeriodEnd:   p.PeriodEnd.Format("2006-01-02"),
		Status:      string(p.Status),
		CreatedAt:   p.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func toPlannedMealView(m *mealplan.PlannedMeal) plannedMealView {
	return plannedMealView{
		ID:                     m.ID,
		MealPlanID:             m.MealPlanID,
		MealDate:               m.MealDate.Format("2006-01-02"),
		MealOccasion:           m.MealOccasion,
		RecipeID:               m.RecipeID,
		RecommendationOptionID: m.RecommendationOptionID,
		Status:                 string(m.Status),
	}
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func dateOrNil(s *string) *time.Time {
	if s == nil || *s == "" {
		return nil
	}
	t, err := parseDate(*s)
	if err != nil {
		return nil
	}
	return &t
}

// listMealPlans handles GET /api/v1/meal-plans.
func (h *Handler) listMealPlans(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	plans, page, err := h.mealPlan.ListMealPlans(c.Context(), profile.ID, c.Query("cursor"), parseLimit(c.Query("limit", "20")))
	if err != nil {
		return response.Error(c, err)
	}
	views := make([]mealPlanView, len(plans))
	for i, p := range plans {
		views[i] = toMealPlanView(&p)
	}
	return response.OK(c, map[string]any{"data": views, "page": page})
}

// createMealPlan handles POST /api/v1/meal-plans.
func (h *Handler) createMealPlan(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	var req struct {
		PeriodStart string `json:"period_start"`
		PeriodEnd   string `json:"period_end"`
		Title       string `json:"title"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	ps, err := parseDate(req.PeriodStart)
	if err != nil {
		return response.Error(c, apperr.New(apperr.CodeFieldInvalid, "Period start is invalid.").
			WithDetails(apperr.Detail{Field: "period_start", Code: string(apperr.CodeFieldInvalid), Message: "Period start must be a valid date."}))
	}
	pe, err := parseDate(req.PeriodEnd)
	if err != nil {
		return response.Error(c, apperr.New(apperr.CodeFieldInvalid, "Period end is invalid.").
			WithDetails(apperr.Detail{Field: "period_end", Code: string(apperr.CodeFieldInvalid), Message: "Period end must be a valid date."}))
	}
	plan, err := h.mealPlan.CreateMealPlan(c.Context(), profile.ID, ps, pe, req.Title)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Created(c, toMealPlanView(plan))
}

// getMealPlan handles GET /api/v1/meal-plans/:planId.
func (h *Handler) getMealPlan(c *fiber.Ctx) error {
	plan, err := h.mealPlan.GetMealPlan(c.Context(), c.Params("planId"))
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toMealPlanView(plan))
}

// updateMealPlan handles PATCH /api/v1/meal-plans/:planId.
func (h *Handler) updateMealPlan(c *fiber.Ctx) error {
	var req struct {
		Title       *string `json:"title"`
		PeriodStart *string `json:"period_start"`
		PeriodEnd   *string `json:"period_end"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	plan, err := h.mealPlan.UpdateMealPlan(c.Context(), c.Params("planId"), req.Title, dateOrNil(req.PeriodStart), dateOrNil(req.PeriodEnd))
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toMealPlanView(plan))
}

// cancelMealPlan handles POST /api/v1/meal-plans/:planId/cancel.
func (h *Handler) cancelMealPlan(c *fiber.Ctx) error {
	plan, err := h.mealPlan.CancelMealPlan(c.Context(), c.Params("planId"))
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toMealPlanView(plan))
}

// completeMealPlan handles POST /api/v1/meal-plans/:planId/complete.
func (h *Handler) completeMealPlan(c *fiber.Ctx) error {
	plan, err := h.mealPlan.CompleteMealPlan(c.Context(), c.Params("planId"))
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toMealPlanView(plan))
}

// listPlannedMeals handles GET /api/v1/meal-plans/:planId/meals.
func (h *Handler) listPlannedMeals(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	meals, page, err := h.mealPlan.ListPlannedMeals(c.Context(), profile.ID, c.Params("planId"), c.Query("cursor"), parseLimit(c.Query("limit", "20")))
	if err != nil {
		return response.Error(c, err)
	}
	views := make([]plannedMealView, len(meals))
	for i, m := range meals {
		views[i] = toPlannedMealView(&m)
	}
	return response.OK(c, map[string]any{"data": views, "page": page})
}

// planMeal handles POST /api/v1/meal-plans/:planId/meals.
func (h *Handler) planMeal(c *fiber.Ctx) error {
	var req struct {
		MealDate               string  `json:"meal_date"`
		MealOccasion           string  `json:"meal_occasion"`
		RecipeID               *string `json:"recipe_id"`
		RecommendationOptionID *string `json:"recommendation_option_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	md, err := parseDate(req.MealDate)
	if err != nil {
		return response.Error(c, apperr.New(apperr.CodeFieldInvalid, "Meal date is invalid.").
			WithDetails(apperr.Detail{Field: "meal_date", Code: string(apperr.CodeFieldInvalid), Message: "Meal date must be a valid date."}))
	}
	meal, err := h.mealPlan.PlanMeal(c.Context(), c.Params("planId"), md, req.MealOccasion, req.RecipeID, req.RecommendationOptionID)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Created(c, toPlannedMealView(meal))
}

// updatePlannedMeal handles PATCH /api/v1/meal-plans/:planId/meals/:plannedMealId.
func (h *Handler) updatePlannedMeal(c *fiber.Ctx) error {
	var req struct {
		MealOccasion *string `json:"meal_occasion"`
		RecipeID     *string `json:"recipe_id"`
		Status       *string `json:"status"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	meal, err := h.mealPlan.UpdatePlannedMeal(c.Context(), c.Params("plannedMealId"), req.MealOccasion, req.RecipeID, req.Status)
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toPlannedMealView(meal))
}

// removePlannedMeal handles DELETE /api/v1/meal-plans/:planId/meals/:plannedMealId.
func (h *Handler) removePlannedMeal(c *fiber.Ctx) error {
	if err := h.mealPlan.RemovePlannedMeal(c.Context(), c.Params("plannedMealId")); err != nil {
		return response.Error(c, err)
	}
	return response.NoContent(c)
}
