# Backend Go Examples

## Handler returning the M6 envelope

```go
func (h *Handler) getPantryItem(c *fiber.Ctx) error {
    itemID, err := uuid.Parse(c.Params("itemId"))
    if err != nil {
        return response.Error(c, apperr.New(apperr.CodeIDInvalid, "Invalid pantry item id."))
    }

    subject := middleware.Subject(c)
    item, err := h.usecases.GetPantryItem(c.Context(), subject, itemID)
    if err != nil {
        return response.Error(c, err)
    }

    return response.OK(c, item)
}
```

## Application use case preserving ownership

```go
func (u *UseCases) GetPantryItem(ctx context.Context, subject string, itemID uuid.UUID) (*View, error) {
    profile, err := u.profiles.FindByAccount(ctx, subject)
    if err != nil {
        return nil, err
    }

    // Ownership is resolved server-side; the item is fetched through the
    // owning Pantry Management repository, never by a client-supplied id.
    item, err := u.pantry.FindItem(ctx, profile.ID, itemID)
    if err != nil {
        return nil, apperr.Wrap(apperr.CodePantryItemNotFound, "The pantry item is not available.", err)
    }

    return toView(item), nil
}
```

## Soft-delete filter in SQLC SQL

```sql
-- name: ListActivePantryItemsForProfile :many
SELECT id, ingredient_name, category, quantity, unit, expiry_date, status
FROM pantry_items
WHERE pantry_id = $1 AND deleted_at IS NULL
ORDER BY expiry_date ASC NULLS LAST
LIMIT $2;
```

## Test structure

```go
func TestGetPantryItem_NotFoundForOtherOwner(t *testing.T) {
    // Repository stub returns a not-found error for an item owned by
    // another profile; the use case must return PANTRY_ITEM_NOT_FOUND.
}
```
