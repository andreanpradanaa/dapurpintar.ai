package shopping

import (
	"context"
	"strings"
	"unicode/utf8"

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

func (s *Service) CreateList(ctx context.Context, profileID, title string) (*ShoppingList, error) {
	title = strings.TrimSpace(title)
	if title == "" || utf8.RuneCountInString(title) > 120 {
		return nil, apperr.New(apperr.CodeFieldInvalid, "Title is invalid.").
			WithDetails(apperr.Detail{Field: "title", Code: string(apperr.CodeFieldInvalid), Message: "Title must be 1-120 characters."})
	}
	return s.store.CreateShoppingList(ctx, profileID, title, ListDraft)
}

func (s *Service) GenerateList(ctx context.Context, profileID string, mealPlanID, recommendationID *string) (*ShoppingList, error) {
	title := "Generated List"
	list, err := s.store.CreateShoppingList(ctx, profileID, title, ListGenerated)
	if err != nil {
		return nil, apperr.Internal(err)
	}
	return list, nil
}

func (s *Service) GetList(ctx context.Context, id string) (*ShoppingList, error) {
	return s.store.GetShoppingListByID(ctx, id)
}

func (s *Service) ListLists(ctx context.Context, profileID, cursor string, limit int32) ([]ShoppingList, *PageInfo, error) {
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	items, err := s.store.ListShoppingLists(ctx, profileID, cursor, limit+1)
	if err != nil {
		return nil, nil, apperr.Internal(err)
	}
	hasMore := len(items) > int(limit)
	if hasMore {
		items = items[:limit]
	}
	var nc string
	if len(items) > 0 {
		nc = items[len(items)-1].ID
	}
	return items, &PageInfo{NextCursor: nc, HasMore: hasMore}, nil
}

func (s *Service) UpdateList(ctx context.Context, id string, title *string) (*ShoppingList, error) {
	if title != nil {
		t := strings.TrimSpace(*title)
		if t == "" || utf8.RuneCountInString(t) > 120 {
			return nil, apperr.New(apperr.CodeFieldInvalid, "Title is invalid.").
				WithDetails(apperr.Detail{Field: "title", Code: string(apperr.CodeFieldInvalid), Message: "Title must be 1-120 characters."})
		}
	}
	return s.store.UpdateShoppingList(ctx, id, title, nil)
}

func (s *Service) ActivateList(ctx context.Context, id string) (*ShoppingList, error) {
	list, err := s.store.GetShoppingListByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if list.Status == ListCompleted || list.Status == ListCancelled || list.Status == ListArchived {
		return nil, apperr.New(apperr.CodeShoppingState, "Cannot activate a completed, cancelled, or archived list.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeShoppingState), Message: "List is in a terminal state."})
	}
	return s.store.UpdateShoppingListStatus(ctx, id, ListActive)
}

func (s *Service) CompleteList(ctx context.Context, id string) (*ShoppingList, error) {
	list, err := s.store.GetShoppingListByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if list.Status == ListCancelled || list.Status == ListArchived {
		return nil, apperr.New(apperr.CodeShoppingState, "Cannot complete a cancelled or archived list.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeShoppingState), Message: "List is in a terminal state."})
	}
	return s.store.UpdateShoppingListStatus(ctx, id, ListCompleted)
}

func (s *Service) CancelList(ctx context.Context, id string) (*ShoppingList, error) {
	list, err := s.store.GetShoppingListByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if list.Status == ListArchived {
		return nil, apperr.New(apperr.CodeShoppingState, "Cannot cancel an archived list.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeShoppingState), Message: "Archived lists cannot be cancelled."})
	}
	return s.store.UpdateShoppingListStatus(ctx, id, ListCancelled)
}

func (s *Service) ArchiveList(ctx context.Context, id string) (*ShoppingList, error) {
	list, err := s.store.GetShoppingListByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if list.Status != ListCompleted && list.Status != ListCancelled {
		return nil, apperr.New(apperr.CodeShoppingState, "Only completed or cancelled lists can be archived.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeShoppingState), Message: "Only terminal lists can be archived."})
	}
	return s.store.UpdateShoppingListStatus(ctx, id, ListArchived)
}

func (s *Service) AddItem(ctx context.Context, listID, name string, quantity float64, unit, source string) (*ShoppingItem, error) {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > 200 {
		return nil, apperr.New(apperr.CodeFieldInvalid, "Ingredient name is invalid.").
			WithDetails(apperr.Detail{Field: "ingredient_name", Code: string(apperr.CodeFieldInvalid), Message: "Ingredient name must be 1-200 characters."})
	}
	if quantity < 0 {
		quantity = 0
	}
	if quantity == 0 {
		quantity = 1
	}
	if unit == "" {
		unit = "unit"
	}
	if source == "" {
		source = "manual"
	}
	return s.store.CreateShoppingItem(ctx, listID, name, quantity, unit, source)
}

func (s *Service) ListItems(ctx context.Context, profileID, listID, cursor string, limit int32) ([]ShoppingItem, *PageInfo, error) {
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	items, err := s.store.ListShoppingItems(ctx, profileID, listID, cursor, limit+1)
	if err != nil {
		return nil, nil, apperr.Internal(err)
	}
	hasMore := len(items) > int(limit)
	if hasMore {
		items = items[:limit]
	}
	var nc string
	if len(items) > 0 {
		nc = items[len(items)-1].ID
	}
	return items, &PageInfo{NextCursor: nc, HasMore: hasMore}, nil
}

func (s *Service) GetItem(ctx context.Context, id string) (*ShoppingItem, error) {
	return s.store.GetShoppingItemByID(ctx, id)
}

func (s *Service) UpdateItem(ctx context.Context, id string, name *string, quantity *float64, unit *string) (*ShoppingItem, error) {
	if name != nil {
		n := strings.TrimSpace(*name)
		if n == "" || utf8.RuneCountInString(n) > 200 {
			return nil, apperr.New(apperr.CodeFieldInvalid, "Ingredient name is invalid.").
				WithDetails(apperr.Detail{Field: "ingredient_name", Code: string(apperr.CodeFieldInvalid), Message: "Ingredient name must be 1-200 characters."})
		}
	}
	if quantity != nil && *quantity < 0 {
		return nil, apperr.New(apperr.CodeFieldInvalid, "Quantity cannot be negative.").
			WithDetails(apperr.Detail{Field: "quantity", Code: string(apperr.CodeFieldInvalid), Message: "Quantity must be zero or greater."})
	}
	return s.store.UpdateShoppingItem(ctx, id, name, quantity, unit)
}

func (s *Service) CompleteItem(ctx context.Context, id string) (*ShoppingItem, error) {
	item, err := s.store.GetShoppingItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if item.Status != ItemOpen {
		return nil, apperr.New(apperr.CodeShoppingState, "Only open items can be completed.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeShoppingState), Message: "Only open items can be marked completed."})
	}
	return s.store.UpdateShoppingItemStatus(ctx, id, ItemCompleted)
}

func (s *Service) RemoveItem(ctx context.Context, id string) error {
	_, err := s.store.RemoveShoppingItem(ctx, id)
	return err
}

func (s *Service) ItemCounts(items []ShoppingItem) ItemCounts {
	var c ItemCounts
	for _, it := range items {
		switch it.Status {
		case ItemOpen:
			c.Open++
		case ItemCompleted:
			c.Completed++
		}
	}
	return c
}
