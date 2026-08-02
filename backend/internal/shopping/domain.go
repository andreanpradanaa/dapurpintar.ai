package shopping

import (
	"errors"
	"time"
)

type ListStatus string

const (
	ListDraft     ListStatus = "draft"
	ListGenerated ListStatus = "generated"
	ListReviewed  ListStatus = "reviewed"
	ListActive    ListStatus = "active"
	ListCompleted ListStatus = "completed"
	ListCancelled ListStatus = "cancelled"
	ListArchived  ListStatus = "archived"
	ListRevised   ListStatus = "revised"
)

type ItemStatus string

const (
	ItemOpen      ItemStatus = "open"
	ItemCompleted ItemStatus = "completed"
	ItemRemoved   ItemStatus = "removed"
)

var ErrNotFound = errors.New("shopping list not found")

type ShoppingList struct {
	ID            string
	UserProfileID string
	MealPlanID    *string
	RecommendID   *string
	Title         string
	Status        ListStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ShoppingItem struct {
	ID             string
	ShoppingListID string
	IngredientName string
	Quantity       float64
	Unit           string
	Source         string
	PantryItemID   *string
	Status         ItemStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ItemCounts struct {
	Open      int `json:"open"`
	Completed int `json:"completed"`
}
