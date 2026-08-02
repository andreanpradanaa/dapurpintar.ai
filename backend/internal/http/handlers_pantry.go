package http

import (
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/identity"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/pantry"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recommendation"
)

type pantryItemView struct {
	ID             string  `json:"id"`
	PantryID       string  `json:"pantry_id"`
	IngredientName string  `json:"ingredient_name"`
	Category       string  `json:"category"`
	Quantity       float64 `json:"quantity"`
	Unit           string  `json:"unit"`
	ExpiryDate     *string `json:"expiry_date"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"created_at"`
}

func toPantryItemView(it *pantry.PantryItem) pantryItemView {
	v := pantryItemView{
		ID:             it.ID,
		PantryID:       it.PantryID,
		IngredientName: it.IngredientName,
		Category:       it.Category,
		Quantity:       it.Quantity,
		Unit:           it.Unit,
		Status:         string(it.Status),
		CreatedAt:      it.CreatedAt.UTC().Format(time.RFC3339),
	}
	if it.ExpiryDate != nil {
		s := it.ExpiryDate.Format("2006-01-02")
		v.ExpiryDate = &s
	}
	return v
}

// getPantrySummary handles GET /api/v1/pantry.
func (h *Handler) getPantrySummary(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	sum, err := h.pantry.GetSummary(c.Context(), profile.ID)
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, sum)
}

// listPantryItems handles GET /api/v1/pantry/items.
func (h *Handler) listPantryItems(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}

	cursor := c.Query("cursor")
	limit := parseLimit(c.Query("limit", "20"))
	category := queryPointer(c.Query("category"))
	status := queryPointer(c.Query("status"))
	sortOrder := queryPointer(c.Query("sort", "expiry_date"))

	items, page, err := h.pantry.ListItems(c.Context(), profile.ID, cursor, limit, category, status, sortOrder)
	if err != nil {
		return response.Error(c, err)
	}

	views := make([]pantryItemView, len(items))
	for i, it := range items {
		views[i] = toPantryItemView(&it)
	}
	return response.OK(c, map[string]any{
		"data": views,
		"page": page,
	})
}

// addPantryItem handles POST /api/v1/pantry/items.
func (h *Handler) addPantryItem(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}

	var req struct {
		IngredientName string  `json:"ingredient_name"`
		Category       string  `json:"category"`
		Quantity       float64 `json:"quantity"`
		Unit           string  `json:"unit"`
		ExpiryDate     *string `json:"expiry_date"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}

	var expiry *time.Time
	if req.ExpiryDate != nil && *req.ExpiryDate != "" {
		t, parseErr := time.Parse("2006-01-02", *req.ExpiryDate)
		if parseErr != nil {
			return response.Error(c, apperr.New(apperr.CodePantryExpiryInvalid, "Expiry date is invalid.").
				WithDetails(apperr.Detail{Field: "expiry_date", Code: string(apperr.CodePantryExpiryInvalid), Message: "Expiry date must be a valid date in YYYY-MM-DD format."}))
		}
		expiry = &t
	}

	item, err := h.pantry.AddItem(c.Context(), profile.ID, pantry.AddItemInput{
		IngredientName: req.IngredientName,
		Category:       req.Category,
		Quantity:       req.Quantity,
		Unit:           req.Unit,
		ExpiryDate:     expiry,
	})
	if err != nil {
		return response.Error(c, err)
	}
	return response.Created(c, toPantryItemView(item))
}

// getPantryItem handles GET /api/v1/pantry/items/:itemId.
func (h *Handler) getPantryItem(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	item, err := h.pantry.GetItem(c.Context(), profile.ID, c.Params("itemId"))
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toPantryItemView(item))
}

// updatePantryItem handles PATCH /api/v1/pantry/items/:itemId.
func (h *Handler) updatePantryItem(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}

	var req struct {
		Quantity   *float64 `json:"quantity"`
		Unit       *string  `json:"unit"`
		Category   *string  `json:"category"`
		ExpiryDate *string  `json:"expiry_date"`
		Status     *string  `json:"status"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}

	if req.Unit != nil {
		unit := *req.Unit
		if utf8.RuneCountInString(unit) > 20 {
			return response.Error(c, apperr.New(apperr.CodeFieldInvalid, "Unit is too long.").
				WithDetails(apperr.Detail{Field: "unit", Code: string(apperr.CodeFieldInvalid), Message: "Unit must be 20 characters or fewer."}))
		}
	}

	var expiry *time.Time
	if req.ExpiryDate != nil {
		if *req.ExpiryDate != "" {
			t, parseErr := time.Parse("2006-01-02", *req.ExpiryDate)
			if parseErr != nil {
				return response.Error(c, apperr.New(apperr.CodePantryExpiryInvalid, "Expiry date is invalid.").
					WithDetails(apperr.Detail{Field: "expiry_date", Code: string(apperr.CodePantryExpiryInvalid), Message: "Expiry date must be a valid date in YYYY-MM-DD format."}))
			}
			expiry = &t
		}
	}

	var cat *string
	if req.Category != nil && *req.Category != "" {
		cat = req.Category
	}

	item, err := h.pantry.UpdateItem(c.Context(), profile.ID, c.Params("itemId"), pantry.UpdateItemInput{
		Quantity:   req.Quantity,
		Unit:       req.Unit,
		Category:   cat,
		ExpiryDate: expiry,
	})
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toPantryItemView(item))
}

// removePantryItem handles DELETE /api/v1/pantry/items/:itemId.
func (h *Handler) removePantryItem(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	if err := h.pantry.RemoveItem(c.Context(), profile.ID, c.Params("itemId")); err != nil {
		return response.Error(c, err)
	}
	return response.NoContent(c)
}

// listExpiringItems handles GET /api/v1/pantry/expiry.
func (h *Handler) listExpiringItems(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}

	cursor := c.Query("cursor")
	limit := parseLimit(c.Query("limit", "20"))
	days := 7
	if d := c.Query("days"); d != "" {
		if parsed, parseErr := strconv.Atoi(d); parseErr == nil && parsed > 0 {
			days = parsed
		}
	}

	items, page, err := h.pantry.ListExpiringItems(c.Context(), profile.ID, cursor, limit, days)
	if err != nil {
		return response.Error(c, err)
	}

	views := make([]pantryItemView, len(items))
	for i, it := range items {
		views[i] = toPantryItemView(&it)
	}
	return response.OK(c, map[string]any{
		"data": views,
		"page": page,
	})
}

func (h *Handler) profileFor(c *fiber.Ctx) (*identity.UserProfile, error) {
	accountID := middlewareSubject(c)
	profile, err := h.identity.GetProfileByAccountID(c.Context(), accountID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func parseLimit(raw string) int32 {
	if raw == "" {
		return 20
	}
	v, err := strconv.Atoi(raw)
	if err != nil || v <= 0 {
		return 20
	}
	if v > 100 {
		return 100
	}
	return int32(v)
}

func queryPointer(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

// analyzePantry handles POST /api/v1/ai/pantry-analysis.
func (h *Handler) analyzePantry(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}

	items, _, listErr := h.pantry.ListItems(c.Context(), profile.ID, "", 200, nil, nil, nil)
	if listErr != nil {
		items = nil
	}

	snapshots := make([]recommendation.PantryItemSnapshot, len(items))
	for i, it := range items {
		var expiry string
		if it.ExpiryDate != nil {
			expiry = it.ExpiryDate.Format("2006-01-02")
		}
		snapshots[i] = recommendation.PantryItemSnapshot{
			ID:             it.ID,
			IngredientName: it.IngredientName,
			Category:       it.Category,
			Quantity:       it.Quantity,
			Unit:           it.Unit,
			ExpiryDate:     expiry,
		}
	}

	analysis, err := h.rec.AnalyzePantry(c.Context(), profile.ID, snapshots)
	if err != nil {
		return response.Error(c, err)
	}

	return response.OK(c, map[string]any{
		"data": analysis,
	})
}

// refreshPantryStatuses handles POST /api/v1/pantry/refresh-statuses.
func (h *Handler) refreshPantryStatuses(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	if err := h.pantry.RefreshStatuses(c.Context(), profile.ID); err != nil {
		return response.Error(c, err)
	}
	return response.NoContent(c)
}

var categoryHints = map[string]string{
	"beras": "pokok", "tepung": "pokok", "gula": "pokok", "garam": "pokok", "minyak": "pokok",
	"bawang": "bumbu", "cabai": "bumbu", "jahe": "bumbu", "kunyit": "bumbu", "lengkuas": "bumbu",
	"serai": "bumbu", "daun": "bumbu", "merica": "bumbu", "ketumbar": "bumbu",
	"ayam": "protein", "daging": "protein", "ikan": "protein", "telur": "protein",
	"tahu": "protein", "tempe": "protein", "udang": "protein", "cumi": "protein",
	"sayur": "sayur", "bayam": "sayur", "kangkung": "sayur", "wortel": "sayur",
	"kol": "sayur", "sawi": "sayur", "tomat": "sayur", "brokoli": "sayur",
	"susu": "susu", "keju": "susu", "yogurt": "susu", "mentega": "susu",
	"kecap": "saus", "saus": "saus", "sambal": "saus",
	"buah": "buah", "apel": "buah", "pisang": "buah", "jeruk": "buah", "mangga": "buah",
	"roti": "roti", "mi": "mi", "pasta": "mi",
}

// suggestCategory handles GET /api/v1/pantry/suggest-category.
func (h *Handler) suggestCategory(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		return response.OK(c, map[string]string{"category": ""})
	}
	lower := strings.ToLower(name)
	for keyword, cat := range categoryHints {
		if strings.Contains(lower, keyword) {
			return response.OK(c, map[string]string{"category": cat})
		}
	}
	return response.OK(c, map[string]string{"category": "other"})
}
