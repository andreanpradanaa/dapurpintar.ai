package store

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/pantry"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/database"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
)

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; skipping Postgres integration test")
	}
	ctx := context.Background()
	pool, err := database.New(ctx, url, logger.New("test"))
	if err != nil {
		t.Fatalf("database.New: %v", err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func testStore(t *testing.T) (*Postgres, *pgxpool.Pool) {
	t.Helper()
	pool := testPool(t)
	return New(pool), pool
}

func testSeedProfile(t *testing.T, pool *pgxpool.Pool) (profileID string) {
	t.Helper()
	accountID := fmt.Sprintf("00000000-0000-0000-0000-%012d", time.Now().UnixNano()%1e12)
	profileID = fmt.Sprintf("00000000-0000-0000-9999-%012d", time.Now().UnixNano()%1e12)
	ctx := context.Background()
	email := fmt.Sprintf("seed-%d@test.example", time.Now().UnixNano())
	_, err := pool.Exec(ctx, `insert into accounts (id, email, password_hash, status) values ($1, $2, 'test', 'active')`, accountID, email)
	if err != nil {
		t.Fatalf("seed account: %v", err)
	}
	_, err = pool.Exec(ctx, `insert into user_profiles (id, account_id, display_name, status) values ($1, $2, 'Test', 'created')`, profileID, accountID)
	if err != nil {
		t.Fatalf("seed profile: %v", err)
	}
	return profileID
}

func TestPostgres_PantryLifecycle(t *testing.T) {
	st, pool := testStore(t)
	profileID := testSeedProfile(t, pool)
	ctx := context.Background()

	p, err := st.CreatePantry(ctx, profileID)
	if err != nil {
		t.Fatalf("CreatePantry: %v", err)
	}

	got, err := st.GetPantryByProfileID(ctx, profileID)
	if err != nil {
		t.Fatalf("GetPantryByProfileID: %v", err)
	}
	if got.ID != p.ID {
		t.Fatalf("pantry id = %q, want %q", got.ID, p.ID)
	}
}

func TestPostgres_PantryItemCRUD(t *testing.T) {
	st, pool := testStore(t)
	profileID := testSeedProfile(t, pool)
	ctx := context.Background()

	p, err := st.CreatePantry(ctx, profileID)
	if err != nil {
		t.Fatalf("CreatePantry: %v", err)
	}

	expiry := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
	created, err := st.CreatePantryItem(ctx, p.ID, "bawang merah", "bumbu", 3.0, "siung", &expiry)
	if err != nil {
		t.Fatalf("CreatePantryItem: %v", err)
	}
	if created.Quantity != 3.0 || created.IngredientName != "bawang merah" {
		t.Fatalf("unexpected item: %+v", created)
	}

	got, err := st.GetPantryItemByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetPantryItemByID: %v", err)
	}
	if got.ID != created.ID {
		t.Fatalf("got id = %q, want %q", got.ID, created.ID)
	}

	qty := 5.0
	updated, err := st.UpdatePantryItem(ctx, created.ID, &qty, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("UpdatePantryItem: %v", err)
	}
	if updated.Quantity != 5.0 {
		t.Fatalf("updated quantity = %f, want 5.0", updated.Quantity)
	}

	_, err = st.RemovePantryItem(ctx, created.ID)
	if err != nil {
		t.Fatalf("RemovePantryItem: %v", err)
	}

	_, err = st.GetPantryItemByID(ctx, created.ID)
	if err == nil {
		t.Fatal("expected not found after removal")
	}
}

func TestPostgres_ListPantryItems(t *testing.T) {
	st, pool := testStore(t)
	profileID := testSeedProfile(t, pool)
	ctx := context.Background()

	p, err := st.CreatePantry(ctx, profileID)
	if err != nil {
		t.Fatalf("CreatePantry: %v", err)
	}

	expiry1 := time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC)
	expiry2 := time.Date(2026, 12, 25, 0, 0, 0, 0, time.UTC)
	st.CreatePantryItem(ctx, p.ID, "beras", "pokok", 2.0, "kg", &expiry1)
	st.CreatePantryItem(ctx, p.ID, "telur", "protein", 12.0, "butir", &expiry2)

	items, err := st.ListPantryItems(ctx, profileID, "", 20, nil, nil, nil)
	if err != nil {
		t.Fatalf("ListPantryItems: %v", err)
	}
	if len(items) < 2 {
		t.Fatalf("expected >=2 items, got %d", len(items))
	}

	expiring, err := st.ListExpiringItems(ctx, profileID, "", 20, time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("ListExpiringItems: %v", err)
	}
	if len(expiring) < 1 {
		t.Fatalf("expected >=1 expiring items, got %d", len(expiring))
	}
}

func TestPantryServiceEndToEnd(t *testing.T) {
	st, pool := testStore(t)
	profileID := testSeedProfile(t, pool)
	ctx := context.Background()
	svc := pantry.NewService(st)

	expiry := time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC)
	item, err := svc.AddItem(ctx, profileID, pantry.AddItemInput{
		IngredientName: "minyak goreng",
		Category:       "bumbu",
		Quantity:       1.5,
		Unit:           "liter",
		ExpiryDate:     &expiry,
	})
	if err != nil {
		t.Fatalf("AddItem: %v", err)
	}
	if item.PantryID == "" {
		t.Fatal("item has no pantry_id")
	}

	sum, err := svc.GetSummary(ctx, profileID)
	if err != nil {
		t.Fatalf("GetSummary: %v", err)
	}
	if sum.TotalItems != 1 {
		t.Fatalf("total_items = %d, want 1", sum.TotalItems)
	}

	got, err := svc.GetItem(ctx, profileID, item.ID)
	if err != nil {
		t.Fatalf("GetItem: %v", err)
	}
	if got.IngredientName != "minyak goreng" {
		t.Fatalf("got ingredient = %q", got.IngredientName)
	}

	newQty := 2.0
	updated, err := svc.UpdateItem(ctx, profileID, item.ID, pantry.UpdateItemInput{
		Quantity: &newQty,
	})
	if err != nil {
		t.Fatalf("UpdateItem: %v", err)
	}
	if updated.Quantity != 2.0 {
		t.Fatalf("updated quantity = %f", updated.Quantity)
	}

	if err := svc.RemoveItem(ctx, profileID, item.ID); err != nil {
		t.Fatalf("RemoveItem: %v", err)
	}

	_, err = svc.GetItem(ctx, profileID, item.ID)
	if err == nil {
		t.Fatal("expected not found after removal")
	}

	items, page, err := svc.ListItems(ctx, profileID, "", 10, nil, nil, nil)
	if err != nil {
		t.Fatalf("ListItems: %v", err)
	}
	_ = items
	if page.HasMore {
		t.Fatalf("expected HasMore=false on empty pantry")
	}
}
