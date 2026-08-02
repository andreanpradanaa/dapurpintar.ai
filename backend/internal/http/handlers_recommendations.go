package http

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recommendation"
)

type recView struct {
	ID                  string  `json:"id"`
	Status              string  `json:"status"`
	Purpose             string  `json:"purpose"`
	Rationale           string  `json:"rationale"`
	ConfidenceStatement *string `json:"confidence_statement"`
	CreatedAt           string  `json:"created_at"`
}

type recDetailView struct {
	ID                  string          `json:"id"`
	Status              string          `json:"status"`
	Purpose             string          `json:"purpose"`
	Rationale           string          `json:"rationale"`
	ConfidenceStatement *string         `json:"confidence_statement"`
	Options             []recOptionView `json:"options"`
	CreatedAt           string          `json:"created_at"`
}

type recOptionView struct {
	ID        string  `json:"id"`
	Position  int32   `json:"position"`
	RecipeID  *string `json:"recipe_id"`
	Rationale string  `json:"rationale"`
	Status    string  `json:"status"`
}

func toRecView(r *recommendation.Recommendation) recView {
	return recView{
		ID:                  r.ID,
		Status:              string(r.Status),
		Purpose:             r.Purpose,
		Rationale:           r.Rationale,
		ConfidenceStatement: r.ConfidenceStatement,
		CreatedAt:           r.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func toRecDetailView(r *recommendation.Recommendation, opts []recommendation.RecommendationOption) recDetailView {
	optViews := make([]recOptionView, len(opts))
	for i, o := range opts {
		optViews[i] = recOptionView{
			ID:        o.ID,
			Position:  o.Position,
			RecipeID:  o.RecipeID,
			Rationale: o.Rationale,
			Status:    string(o.Status),
		}
	}
	return recDetailView{
		ID:                  r.ID,
		Status:              string(r.Status),
		Purpose:             r.Purpose,
		Rationale:           r.Rationale,
		ConfidenceStatement: r.ConfidenceStatement,
		Options:             optViews,
		CreatedAt:           r.CreatedAt.UTC().Format(time.RFC3339),
	}
}

// listRecommendations handles GET /api/v1/recommendations.
func (h *Handler) listRecommendations(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	status := queryPointer(c.Query("status"))
	purpose := queryPointer(c.Query("purpose"))
	recs, page, err := h.rec.List(c.Context(), profile.ID, c.Query("cursor"), parseLimit(c.Query("limit", "20")), status, purpose)
	if err != nil {
		return response.Error(c, err)
	}
	views := make([]recView, len(recs))
	for i, r := range recs {
		views[i] = toRecView(&r)
	}
	return response.OK(c, map[string]any{"data": views, "page": page})
}

// requestRecommendation handles POST /api/v1/recommendations.
func (h *Handler) requestRecommendation(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	var req struct {
		Purpose          string `json:"purpose"`
		MaxPrepMinutes   *int32 `json:"max_prep_minutes"`
		UseExpiringFirst bool   `json:"use_expiring_first"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	rec, err := h.rec.Request(c.Context(), profile.ID, req.Purpose, req.MaxPrepMinutes, req.UseExpiringFirst)
	if err != nil {
		return response.Error(c, err)
	}
	c.Status(202)
	return response.Created(c, toRecView(rec))
}

// getRecommendation handles GET /api/v1/recommendations/:recId.
func (h *Handler) getRecommendation(c *fiber.Ctx) error {
	rec, opts, err := h.rec.Get(c.Context(), c.Params("recId"))
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toRecDetailView(rec, opts))
}

// presentRecommendation handles POST /api/v1/recommendations/:recId/present.
func (h *Handler) presentRecommendation(c *fiber.Ctx) error {
	rec, err := h.rec.Present(c.Context(), c.Params("recId"))
	if err != nil {
		return response.Error(c, err)
	}
	_, opts, _ := h.rec.Get(c.Context(), rec.ID)
	return response.OK(c, toRecDetailView(rec, opts))
}

// acceptOption handles POST /api/v1/recommendations/:recId/options/:optId/accept.
func (h *Handler) acceptOption(c *fiber.Ctx) error {
	rec, err := h.rec.AcceptOption(c.Context(), c.Params("recId"), c.Params("optId"))
	if err != nil {
		return response.Error(c, err)
	}
	_, opts, _ := h.rec.Get(c.Context(), rec.ID)
	return response.OK(c, toRecDetailView(rec, opts))
}

// rejectRecommendation handles POST /api/v1/recommendations/:recId/reject.
func (h *Handler) rejectRecommendation(c *fiber.Ctx) error {
	rec, err := h.rec.Reject(c.Context(), c.Params("recId"))
	if err != nil {
		return response.Error(c, err)
	}
	_, opts, _ := h.rec.Get(c.Context(), rec.ID)
	return response.OK(c, toRecDetailView(rec, opts))
}

// supersedeRecommendation handles POST /api/v1/recommendations/:recId/supersede.
func (h *Handler) supersedeRecommendation(c *fiber.Ctx) error {
	rec, err := h.rec.Supersede(c.Context(), c.Params("recId"))
	if err != nil {
		return response.Error(c, err)
	}
	_, opts, _ := h.rec.Get(c.Context(), rec.ID)
	return response.OK(c, toRecDetailView(rec, opts))
}
