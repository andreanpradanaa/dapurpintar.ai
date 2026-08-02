package shopping

import "context"

type Store interface {
	CreateShoppingList(ctx context.Context, profileID, title string, status ListStatus) (*ShoppingList, error)
	GetShoppingListByID(ctx context.Context, id string) (*ShoppingList, error)
	ListShoppingLists(ctx context.Context, profileID, cursor string, limit int32) ([]ShoppingList, error)
	UpdateShoppingList(ctx context.Context, id string, title *string, status *string) (*ShoppingList, error)
	UpdateShoppingListStatus(ctx context.Context, id string, status ListStatus) (*ShoppingList, error)

	CreateShoppingItem(ctx context.Context, listID, name string, quantity float64, unit, source string) (*ShoppingItem, error)
	ListShoppingItems(ctx context.Context, profileID, listID, cursor string, limit int32) ([]ShoppingItem, error)
	GetShoppingItemByID(ctx context.Context, id string) (*ShoppingItem, error)
	UpdateShoppingItem(ctx context.Context, id string, name *string, quantity *float64, unit *string) (*ShoppingItem, error)
	UpdateShoppingItemStatus(ctx context.Context, id string, status ItemStatus) (*ShoppingItem, error)
	RemoveShoppingItem(ctx context.Context, id string) (*ShoppingItem, error)
}
