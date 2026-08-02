package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/gen/sqlc"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/shopping"
)

type Postgres struct {
	db *sqlc.Queries
}

func New(conn sqlc.DBTX) *Postgres {
	return &Postgres{db: sqlc.New(conn)}
}

func (p *Postgres) CreateShoppingList(ctx context.Context, profileID, title string, status shopping.ListStatus) (*shopping.ShoppingList, error) {
	row, err := p.db.CreateShoppingList(ctx, sqlc.CreateShoppingListParams{
		UserProfileID: profileID,
		Title:         title,
		Status:        string(status),
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingList(row), nil
}

func (p *Postgres) GetShoppingListByID(ctx context.Context, id string) (*shopping.ShoppingList, error) {
	row, err := p.db.GetShoppingListByID(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingList(row), nil
}

func (p *Postgres) ListShoppingLists(ctx context.Context, profileID, cursor string, limit int32) ([]shopping.ShoppingList, error) {
	rows, err := p.db.ListShoppingLists(ctx, sqlc.ListShoppingListsParams{
		UserProfileID: profileID,
		Column2:       cursor,
		Limit:         limit,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	out := make([]shopping.ShoppingList, len(rows))
	for i, r := range rows {
		out[i] = *toShoppingList(r)
	}
	return out, nil
}

func (p *Postgres) UpdateShoppingList(ctx context.Context, id string, title, status *string) (*shopping.ShoppingList, error) {
	row, err := p.db.UpdateShoppingList(ctx, sqlc.UpdateShoppingListParams{
		ID:     id,
		Title:  title,
		Status: status,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingList(row), nil
}

func (p *Postgres) UpdateShoppingListStatus(ctx context.Context, id string, status shopping.ListStatus) (*shopping.ShoppingList, error) {
	row, err := p.db.UpdateShoppingListStatus(ctx, sqlc.UpdateShoppingListStatusParams{
		ID:     id,
		Status: string(status),
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingList(row), nil
}

func (p *Postgres) CreateShoppingItem(ctx context.Context, listID, name string, quantity float64, unit, source string) (*shopping.ShoppingItem, error) {
	row, err := p.db.CreateShoppingItem(ctx, sqlc.CreateShoppingItemParams{
		ShoppingListID: listID,
		IngredientName: name,
		Quantity:       floatToNumeric(quantity),
		Unit:           unit,
		Source:         source,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingItem(row), nil
}

func (p *Postgres) ListShoppingItems(ctx context.Context, profileID, listID, cursor string, limit int32) ([]shopping.ShoppingItem, error) {
	rows, err := p.db.ListShoppingItems(ctx, sqlc.ListShoppingItemsParams{
		UserProfileID:  profileID,
		ShoppingListID: listID,
		Column3:        cursor,
		Limit:          limit,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	out := make([]shopping.ShoppingItem, len(rows))
	for i, r := range rows {
		out[i] = *toShoppingItem(r)
	}
	return out, nil
}

func (p *Postgres) GetShoppingItemByID(ctx context.Context, id string) (*shopping.ShoppingItem, error) {
	row, err := p.db.GetShoppingItemByID(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingItem(row), nil
}

func (p *Postgres) UpdateShoppingItem(ctx context.Context, id string, name *string, quantity *float64, unit *string) (*shopping.ShoppingItem, error) {
	params := sqlc.UpdateShoppingItemParams{
		ID:       id,
		Quantity: pgtype.Numeric{Valid: false},
	}
	if name != nil {
		params.IngredientName = name
	}
	if quantity != nil {
		params.Quantity = floatToNumeric(*quantity)
	}
	if unit != nil {
		params.Unit = unit
	}
	row, err := p.db.UpdateShoppingItem(ctx, params)
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingItem(row), nil
}

func (p *Postgres) UpdateShoppingItemStatus(ctx context.Context, id string, status shopping.ItemStatus) (*shopping.ShoppingItem, error) {
	row, err := p.db.UpdateShoppingItemStatus(ctx, sqlc.UpdateShoppingItemStatusParams{
		ID:     id,
		Status: string(status),
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingItem(row), nil
}

func (p *Postgres) RemoveShoppingItem(ctx context.Context, id string) (*shopping.ShoppingItem, error) {
	row, err := p.db.RemoveShoppingItem(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toShoppingItem(row), nil
}

func toShoppingList(r sqlc.ShoppingList) *shopping.ShoppingList {
	return &shopping.ShoppingList{
		ID:            r.ID,
		UserProfileID: r.UserProfileID,
		MealPlanID:    r.MealPlanID,
		RecommendID:   r.RecommendationID,
		Title:         r.Title,
		Status:        shopping.ListStatus(r.Status),
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}
}

func toShoppingItem(r sqlc.ShoppingItem) *shopping.ShoppingItem {
	return &shopping.ShoppingItem{
		ID:             r.ID,
		ShoppingListID: r.ShoppingListID,
		IngredientName: r.IngredientName,
		Quantity:       numericToFloat(r.Quantity),
		Unit:           r.Unit,
		Source:         r.Source,
		PantryItemID:   r.PantryItemID,
		Status:         shopping.ItemStatus(r.Status),
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}
}

func floatToNumeric(v float64) pgtype.Numeric {
	var n pgtype.Numeric
	_ = n.Scan(fmt.Sprintf("%v", v))
	return n
}

func numericToFloat(n pgtype.Numeric) float64 {
	f, _ := n.Float64Value()
	return f.Float64
}

func mapErr(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == "no rows in result set" || err == pgx.ErrNoRows {
		return shopping.ErrNotFound
	}
	return err
}
