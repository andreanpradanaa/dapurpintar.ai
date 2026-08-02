package pantry

import (
	"context"
	"strings"
	"time"
	"unicode/utf8"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

type Summary struct {
	TotalItems        int `json:"total_items"`
	ExpiringSoonCount int `json:"expiring_soon_count"`
	RunningLowCount   int `json:"running_low_count"`
}

type PageInfo struct {
	Cursor  string `json:"cursor"`
	Limit   int32  `json:"limit"`
	HasMore bool   `json:"has_more"`
}

func (s *Service) GetSummary(ctx context.Context, profileID string) (*Summary, error) {
	pantry, err := s.store.GetPantryByProfileID(ctx, profileID)
	if err != nil {
		return &Summary{}, nil
	}

	items, err := s.store.ListPantryItems(ctx, profileID, "", 200, nil, nil, ptr("id_asc"))
	if err != nil {
		return nil, apperr.Internal(err)
	}

	total := 0
	exp := 0
	low := 0
	for _, it := range items {
		total++
		if it.Status == ItemExpiringSoon {
			exp++
		}
		if it.Status == ItemRunningLow {
			low++
		}
	}
	_ = pantry
	return &Summary{TotalItems: total, ExpiringSoonCount: exp, RunningLowCount: low}, nil
}

const maxPageLimit = 100

func (s *Service) ListItems(ctx context.Context, profileID, cursor string, limit int32, category, status, sortOrder *string) ([]PantryItem, *PageInfo, error) {
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	if cursor == "" {
		cursor = zeroUUID
	}

	items, err := s.store.ListPantryItems(ctx, profileID, cursor, limit+1, category, status, sortOrder)
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
	return items, &PageInfo{Cursor: nextCursor, Limit: limit, HasMore: hasMore}, nil
}

type AddItemInput struct {
	IngredientName string
	Category       string
	Quantity       float64
	Unit           string
	ExpiryDate     *time.Time
}

func (s *Service) AddItem(ctx context.Context, profileID string, in AddItemInput) (*PantryItem, error) {
	if err := validateIngredientName(in.IngredientName); err != nil {
		return nil, err
	}
	if err := validateCategory(in.Category); err != nil {
		return nil, err
	}
	if in.Quantity < 0 {
		return nil, apperr.New(apperr.CodePantryQtyNegative, "Quantity cannot be negative.").
			WithDetails(apperr.Detail{Field: "quantity", Code: string(apperr.CodePantryQtyNegative), Message: "Quantity must be zero or greater."})
	}

	unit := in.Unit
	if unit == "" {
		unit = "unit"
	}

	pantry, err := s.ensurePantry(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return s.store.CreatePantryItem(ctx, pantry.ID, in.IngredientName, in.Category, in.Quantity, unit, in.ExpiryDate)
}

func (s *Service) GetItem(ctx context.Context, profileID, itemID string) (*PantryItem, error) {
	item, err := s.store.GetPantryItemByID(ctx, itemID)
	if err != nil {
		return nil, err
	}
	if err := s.authorizeItem(ctx, item, profileID); err != nil {
		return nil, err
	}
	return item, nil
}

type UpdateItemInput struct {
	Quantity   *float64
	Unit       *string
	Category   *string
	ExpiryDate *time.Time
}

func (s *Service) UpdateItem(ctx context.Context, profileID, itemID string, in UpdateItemInput) (*PantryItem, error) {
	if _, err := s.GetItem(ctx, profileID, itemID); err != nil {
		return nil, err
	}
	if in.Quantity != nil && *in.Quantity < 0 {
		return nil, apperr.New(apperr.CodePantryQtyNegative, "Quantity cannot be negative.").
			WithDetails(apperr.Detail{Field: "quantity", Code: string(apperr.CodePantryQtyNegative), Message: "Quantity must be zero or greater."})
	}
	if in.Category != nil {
		v := strings.TrimSpace(*in.Category)
		if v == "" || utf8.RuneCountInString(v) > 100 {
			return nil, apperr.New(apperr.CodeFieldInvalid, "Category is invalid.").
				WithDetails(apperr.Detail{Field: "category", Code: string(apperr.CodeFieldInvalid), Message: "Category must be 1-100 characters."})
		}
	}

	return s.store.UpdatePantryItem(ctx, itemID, in.Quantity, in.Unit, in.Category, in.ExpiryDate, nil)
}

func (s *Service) RemoveItem(ctx context.Context, profileID, itemID string) error {
	if _, err := s.GetItem(ctx, profileID, itemID); err != nil {
		return err
	}
	_, err := s.store.RemovePantryItem(ctx, itemID)
	return err
}

func (s *Service) ListExpiringItems(ctx context.Context, profileID, cursor string, limit int32, beforeDays int) ([]PantryItem, *PageInfo, error) {
	if beforeDays <= 0 {
		beforeDays = 7
	}
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	if cursor == "" {
		cursor = zeroUUID
	}

	beforeDate := time.Now().AddDate(0, 0, beforeDays)

	items, err := s.store.ListExpiringItems(ctx, profileID, cursor, limit+1, beforeDate)
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
	return items, &PageInfo{Cursor: nextCursor, Limit: limit, HasMore: hasMore}, nil
}

func (s *Service) RefreshStatuses(ctx context.Context, profileID string) error {
	return s.store.RefreshPantryItemStatuses(ctx, profileID, 3)
}

func (s *Service) ensurePantry(ctx context.Context, profileID string) (*Pantry, error) {
	pantry, err := s.store.GetPantryByProfileID(ctx, profileID)
	if err == nil {
		return pantry, nil
	}
	return s.store.CreatePantry(ctx, profileID)
}

func (s *Service) authorizeItem(ctx context.Context, item *PantryItem, profileID string) error {
	pantry, err := s.store.GetPantryByProfileID(ctx, profileID)
	if err != nil {
		return apperr.New(apperr.CodePantryItemNotFound, "Pantry item not found.").
			WithDetails(apperr.Detail{Field: "pantry_item_id", Code: string(apperr.CodePantryItemNotFound), Message: "The pantry item was not found in your pantry."})
	}
	if item.PantryID != pantry.ID {
		return apperr.New(apperr.CodePantryItemNotFound, "Pantry item not found.").
			WithDetails(apperr.Detail{Field: "pantry_item_id", Code: string(apperr.CodePantryItemNotFound), Message: "The pantry item was not found in your pantry."})
	}
	return nil
}

const zeroUUID = "00000000-0000-0000-0000-000000000000"

func ptr[T any](v T) *T { return &v }

func validateIngredientName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > 200 {
		return apperr.New(apperr.CodeFieldInvalid, "Ingredient name is invalid.").
			WithDetails(apperr.Detail{Field: "ingredient_name", Code: string(apperr.CodeFieldInvalid), Message: "Ingredient name must be 1-200 characters."})
	}
	return nil
}

func validateCategory(category string) error {
	category = strings.TrimSpace(category)
	if category == "" || utf8.RuneCountInString(category) > 100 {
		return apperr.New(apperr.CodeFieldInvalid, "Category is invalid.").
			WithDetails(apperr.Detail{Field: "category", Code: string(apperr.CodeFieldInvalid), Message: "Category must be 1-100 characters."})
	}
	return nil
}
