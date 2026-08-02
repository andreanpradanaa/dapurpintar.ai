package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/gen/sqlc"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/pantry"
)

type Postgres struct {
	db *sqlc.Queries
}

func New(conn sqlc.DBTX) *Postgres {
	return &Postgres{db: sqlc.New(conn)}
}

func (p *Postgres) GetPantryByProfileID(ctx context.Context, profileID string) (*pantry.Pantry, error) {
	row, err := p.db.GetPantryByProfileID(ctx, profileID)
	if err != nil {
		return nil, mapPantryErr(err)
	}
	return toPantry(row), nil
}

func (p *Postgres) CreatePantry(ctx context.Context, profileID string) (*pantry.Pantry, error) {
	row, err := p.db.CreatePantry(ctx, profileID)
	if err != nil {
		return nil, mapPantryErr(err)
	}
	return toPantry(row), nil
}

func (p *Postgres) GetPantryItemByID(ctx context.Context, id string) (*pantry.PantryItem, error) {
	row, err := p.db.GetPantryItemByID(ctx, id)
	if err != nil {
		return nil, mapPantryErr(err)
	}
	return toPantryItem(row), nil
}

func (p *Postgres) CreatePantryItem(ctx context.Context, pantryID, name, category string, quantity float64, unit string, expiryDate *time.Time) (*pantry.PantryItem, error) {
	row, err := p.db.CreatePantryItem(ctx, sqlc.CreatePantryItemParams{
		PantryID:       pantryID,
		IngredientName: name,
		Category:       category,
		Quantity:       floatToNumeric(quantity),
		Unit:           unit,
		ExpiryDate:     expiryDate,
	})
	if err != nil {
		return nil, mapPantryErr(err)
	}
	return toPantryItem(row), nil
}

func (p *Postgres) UpdatePantryItem(ctx context.Context, id string, quantity *float64, unit, category *string, expiryDate *time.Time, status *string) (*pantry.PantryItem, error) {
	params := sqlc.UpdatePantryItemParams{
		ID:       id,
		Quantity: pgtype.Numeric{Valid: false},
	}
	if quantity != nil {
		params.Quantity = floatToNumeric(*quantity)
	}
	params.Unit = unit
	params.Category = category
	params.ExpiryDate = expiryDate
	params.Status = status

	row, err := p.db.UpdatePantryItem(ctx, params)
	if err != nil {
		return nil, mapPantryErr(err)
	}
	return toPantryItem(row), nil
}

func (p *Postgres) RemovePantryItem(ctx context.Context, id string) (*pantry.PantryItem, error) {
	row, err := p.db.RemovePantryItem(ctx, id)
	if err != nil {
		return nil, mapPantryErr(err)
	}
	return toPantryItem(row), nil
}

func (p *Postgres) UpdatePantryItemStatus(ctx context.Context, id string, status pantry.ItemStatus) (*pantry.PantryItem, error) {
	row, err := p.db.UpdatePantryItemStatus(ctx, sqlc.UpdatePantryItemStatusParams{
		ID:     id,
		Status: string(status),
	})
	if err != nil {
		return nil, mapPantryErr(err)
	}
	return toPantryItem(row), nil
}

func (p *Postgres) ListPantryItems(ctx context.Context, profileID, cursor string, limit int32, category, status, sortOrder *string) ([]pantry.PantryItem, error) {
	rows, err := p.db.ListPantryItems(ctx, sqlc.ListPantryItemsParams{
		Column1:       cursor,
		Limit:         limit,
		UserProfileID: profileID,
		Category:      category,
		Status:        status,
		SortOrder:     sortOrder,
	})
	if err != nil {
		return nil, mapPantryErr(err)
	}
	items := make([]pantry.PantryItem, len(rows))
	for i, r := range rows {
		items[i] = *toPantryItem(r)
	}
	return items, nil
}

func (p *Postgres) ListExpiringItems(ctx context.Context, profileID, cursor string, limit int32, beforeDate time.Time) ([]pantry.PantryItem, error) {
	rows, err := p.db.ListExpiringItems(ctx, sqlc.ListExpiringItemsParams{
		Column1:       cursor,
		Limit:         limit,
		ExpiryDate:    &beforeDate,
		UserProfileID: profileID,
	})
	if err != nil {
		return nil, mapPantryErr(err)
	}
	items := make([]pantry.PantryItem, len(rows))
	for i, r := range rows {
		items[i] = *toPantryItem(r)
	}
	return items, nil
}

func toPantry(r sqlc.Pantry) *pantry.Pantry {
	return &pantry.Pantry{
		ID:            r.ID,
		UserProfileID: r.UserProfileID,
		Status:        r.Status,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}
}

func toPantryItem(r sqlc.PantryItem) *pantry.PantryItem {
	return &pantry.PantryItem{
		ID:             r.ID,
		PantryID:       r.PantryID,
		IngredientName: r.IngredientName,
		Category:       r.Category,
		Quantity:       numericToFloat(r.Quantity),
		Unit:           r.Unit,
		ExpiryDate:     r.ExpiryDate,
		Status:         pantry.ItemStatus(r.Status),
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

func mapPantryErr(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == "no rows in result set" || err == pgx.ErrNoRows {
		return pantry.ErrNotFound
	}
	return err
}
