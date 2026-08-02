package pantry

import (
	"context"
	"time"
)

type Store interface {
	GetPantryByProfileID(ctx context.Context, profileID string) (*Pantry, error)
	CreatePantry(ctx context.Context, profileID string) (*Pantry, error)

	GetPantryItemByID(ctx context.Context, id string) (*PantryItem, error)
	CreatePantryItem(ctx context.Context, pantryID, name, category string, quantity float64, unit string, expiryDate *time.Time) (*PantryItem, error)
	UpdatePantryItem(ctx context.Context, id string, quantity *float64, unit, category *string, expiryDate *time.Time, status *string) (*PantryItem, error)
	RemovePantryItem(ctx context.Context, id string) (*PantryItem, error)
	UpdatePantryItemStatus(ctx context.Context, id string, status ItemStatus) (*PantryItem, error)

	ListPantryItems(ctx context.Context, profileID, cursor string, limit int32, category, status, sortOrder *string) ([]PantryItem, error)
	ListExpiringItems(ctx context.Context, profileID, cursor string, limit int32, beforeDate time.Time) ([]PantryItem, error)
	RefreshPantryItemStatuses(ctx context.Context, profileID string, expiryDays int32) error
}
