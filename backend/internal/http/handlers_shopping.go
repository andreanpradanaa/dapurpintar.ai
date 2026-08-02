package http

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/shopping"
)

type shoppingListView struct {
	ID         string              `json:"id"`
	Title      string              `json:"title"`
	Status     string              `json:"status"`
	ItemCounts shopping.ItemCounts `json:"item_counts"`
	CreatedAt  string              `json:"created_at"`
}

type shoppingItemView struct {
	ID             string  `json:"id"`
	ShoppingListID string  `json:"shopping_list_id"`
	IngredientName string  `json:"ingredient_name"`
	Quantity       float64 `json:"quantity"`
	Unit           string  `json:"unit"`
	Source         string  `json:"source"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"created_at"`
}

func toShoppingListView(l *shopping.ShoppingList, counts shopping.ItemCounts) shoppingListView {
	return shoppingListView{
		ID:         l.ID,
		Title:      l.Title,
		Status:     string(l.Status),
		ItemCounts: counts,
		CreatedAt:  l.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func toShoppingItemView(it *shopping.ShoppingItem) shoppingItemView {
	return shoppingItemView{
		ID:             it.ID,
		ShoppingListID: it.ShoppingListID,
		IngredientName: it.IngredientName,
		Quantity:       it.Quantity,
		Unit:           it.Unit,
		Source:         it.Source,
		Status:         string(it.Status),
		CreatedAt:      it.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func (h *Handler) countsFor(ctx *fiber.Ctx, listID, profileID string) shopping.ItemCounts {
	items, _, err := h.shopping.ListItems(ctx.Context(), profileID, listID, "", 200)
	if err != nil {
		return shopping.ItemCounts{}
	}
	return h.shopping.ItemCounts(items)
}

// listShoppingLists handles GET /api/v1/shopping-lists.
func (h *Handler) listShoppingLists(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	lists, page, err := h.shopping.ListLists(c.Context(), profile.ID, c.Query("cursor"), parseLimit(c.Query("limit", "20")))
	if err != nil {
		return response.Error(c, err)
	}
	views := make([]shoppingListView, len(lists))
	for i, l := range lists {
		items, _, _ := h.shopping.ListItems(c.Context(), profile.ID, l.ID, "", 200)
		views[i] = toShoppingListView(&l, h.shopping.ItemCounts(items))
	}
	return response.OK(c, map[string]any{"data": views, "page": page})
}

// createShoppingList handles POST /api/v1/shopping-lists.
func (h *Handler) createShoppingList(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	var req struct {
		Title string `json:"title"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	list, err := h.shopping.CreateList(c.Context(), profile.ID, req.Title)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Created(c, toShoppingListView(list, shopping.ItemCounts{}))
}

// generateShoppingList handles POST /api/v1/shopping-lists/generate.
func (h *Handler) generateShoppingList(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	var req struct {
		MealPlanID         *string `json:"meal_plan_id"`
		RecommendationID   *string `json:"recommendation_id"`
		IncludePantryCheck bool    `json:"include_pantry_check"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	list, err := h.shopping.GenerateList(c.Context(), profile.ID, req.MealPlanID, req.RecommendationID)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Created(c, toShoppingListView(list, shopping.ItemCounts{}))
}

// getShoppingList handles GET /api/v1/shopping-lists/:listId.
func (h *Handler) getShoppingList(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	list, err := h.shopping.GetList(c.Context(), c.Params("listId"))
	if err != nil {
		return response.Error(c, err)
	}
	counts := h.countsFor(c, list.ID, profile.ID)
	return response.OK(c, toShoppingListView(list, counts))
}

// updateShoppingList handles PATCH /api/v1/shopping-lists/:listId.
func (h *Handler) updateShoppingList(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	var req struct {
		Title *string `json:"title"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	list, err := h.shopping.UpdateList(c.Context(), c.Params("listId"), req.Title)
	if err != nil {
		return response.Error(c, err)
	}
	counts := h.countsFor(c, list.ID, profile.ID)
	return response.OK(c, toShoppingListView(list, counts))
}

// activateShoppingList handles POST /api/v1/shopping-lists/:listId/activate.
func (h *Handler) activateShoppingList(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	list, err := h.shopping.ActivateList(c.Context(), c.Params("listId"))
	if err != nil {
		return response.Error(c, err)
	}
	counts := h.countsFor(c, list.ID, profile.ID)
	return response.OK(c, toShoppingListView(list, counts))
}

// completeShoppingList handles POST /api/v1/shopping-lists/:listId/complete.
func (h *Handler) completeShoppingList(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	list, err := h.shopping.CompleteList(c.Context(), c.Params("listId"))
	if err != nil {
		return response.Error(c, err)
	}
	counts := h.countsFor(c, list.ID, profile.ID)
	return response.OK(c, toShoppingListView(list, counts))
}

// cancelShoppingList handles POST /api/v1/shopping-lists/:listId/cancel.
func (h *Handler) cancelShoppingList(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	list, err := h.shopping.CancelList(c.Context(), c.Params("listId"))
	if err != nil {
		return response.Error(c, err)
	}
	counts := h.countsFor(c, list.ID, profile.ID)
	return response.OK(c, toShoppingListView(list, counts))
}

// archiveShoppingList handles POST /api/v1/shopping-lists/:listId/archive.
func (h *Handler) archiveShoppingList(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	list, err := h.shopping.ArchiveList(c.Context(), c.Params("listId"))
	if err != nil {
		return response.Error(c, err)
	}
	counts := h.countsFor(c, list.ID, profile.ID)
	return response.OK(c, toShoppingListView(list, counts))
}

// listShoppingItems handles GET /api/v1/shopping-lists/:listId/items.
func (h *Handler) listShoppingItems(c *fiber.Ctx) error {
	profile, err := h.profileFor(c)
	if err != nil {
		return response.Error(c, err)
	}
	items, page, err := h.shopping.ListItems(c.Context(), profile.ID, c.Params("listId"), c.Query("cursor"), parseLimit(c.Query("limit", "20")))
	if err != nil {
		return response.Error(c, err)
	}
	views := make([]shoppingItemView, len(items))
	for i, it := range items {
		views[i] = toShoppingItemView(&it)
	}
	return response.OK(c, map[string]any{"data": views, "page": page})
}

// addShoppingItem handles POST /api/v1/shopping-lists/:listId/items.
func (h *Handler) addShoppingItem(c *fiber.Ctx) error {
	var req struct {
		IngredientName string  `json:"ingredient_name"`
		Quantity       float64 `json:"quantity"`
		Unit           string  `json:"unit"`
		Source         string  `json:"source"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	item, err := h.shopping.AddItem(c.Context(), c.Params("listId"), req.IngredientName, req.Quantity, req.Unit, req.Source)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Created(c, toShoppingItemView(item))
}

// updateShoppingItem handles PATCH /api/v1/shopping-lists/:listId/items/:itemId.
func (h *Handler) updateShoppingItem(c *fiber.Ctx) error {
	var req struct {
		IngredientName *string  `json:"ingredient_name"`
		Quantity       *float64 `json:"quantity"`
		Unit           *string  `json:"unit"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, payloadError(err))
	}
	item, err := h.shopping.UpdateItem(c.Context(), c.Params("itemId"), req.IngredientName, req.Quantity, req.Unit)
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toShoppingItemView(item))
}

// removeShoppingItem handles DELETE /api/v1/shopping-lists/:listId/items/:itemId.
func (h *Handler) removeShoppingItem(c *fiber.Ctx) error {
	if err := h.shopping.RemoveItem(c.Context(), c.Params("itemId")); err != nil {
		return response.Error(c, err)
	}
	return response.NoContent(c)
}

// completeShoppingItem handles POST /api/v1/shopping-lists/:listId/items/:itemId/complete.
func (h *Handler) completeShoppingItem(c *fiber.Ctx) error {
	item, err := h.shopping.CompleteItem(c.Context(), c.Params("itemId"))
	if err != nil {
		return response.Error(c, err)
	}
	return response.OK(c, toShoppingItemView(item))
}
