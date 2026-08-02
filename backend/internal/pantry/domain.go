package pantry

import (
	"errors"
	"time"
)

type ItemStatus string

const (
	ItemAvailable    ItemStatus = "available"
	ItemRunningLow   ItemStatus = "running_low"
	ItemExpiringSoon ItemStatus = "expiring_soon"
	ItemConsumed     ItemStatus = "consumed"
	ItemRemoved      ItemStatus = "removed"
)

var ErrNotFound = errors.New("pantry item not found")

type Pantry struct {
	ID            string
	UserProfileID string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PantryItem struct {
	ID             string
	PantryID       string
	IngredientName string
	Category       string
	Quantity       float64
	Unit           string
	ExpiryDate     *time.Time
	Status         ItemStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
