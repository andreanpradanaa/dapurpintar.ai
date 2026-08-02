package store

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/database"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/shopping"
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

func testSeedShoppingProfile(t *testing.T, pool *pgxpool.Pool, profileID, accountID string) {
	t.Helper()
	ctx := context.Background()
	_, err := pool.Exec(ctx, `insert into accounts (id, email, password_hash, status) values ($1, $2, 'test', 'active') on conflict (id) do nothing`, accountID, accountID+"@test.example")
	if err == nil {
		_, err = pool.Exec(ctx, `insert into user_profiles (id, account_id, display_name, status) values ($1, $2, 'Test', 'created') on conflict (id) do nothing`, profileID, accountID)
	}
	if err != nil {
		t.Fatalf("seed profile: %v", err)
	}
}

func TestShopping_CreateAndList(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	testSeedShoppingProfile(t, pool, "61111111-1111-1111-1111-111111111101", "71111111-1111-1111-1111-111111111101")

	list, err := st.CreateShoppingList(ctx, "61111111-1111-1111-1111-111111111101", "Belanja Mingguan", shopping.ListDraft)
	if err != nil {
		t.Fatalf("CreateShoppingList: %v", err)
	}
	if list.Status != shopping.ListDraft {
		t.Fatalf("status = %q", list.Status)
	}

	got, err := st.GetShoppingListByID(ctx, list.ID)
	if err != nil {
		t.Fatalf("GetShoppingListByID: %v", err)
	}
	if got.Title != "Belanja Mingguan" {
		t.Fatalf("title = %q", got.Title)
	}

	lists, err := st.ListShoppingLists(ctx, "61111111-1111-1111-1111-111111111101", "", 20)
	if err != nil {
		t.Fatalf("ListShoppingLists: %v", err)
	}
	if len(lists) < 1 {
		t.Fatal("expected >=1 list")
	}
}

func TestShopping_Lifecycle(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	testSeedShoppingProfile(t, pool, "61111111-1111-1111-1111-111111111102", "71111111-1111-1111-1111-111111111102")

	list, _ := st.CreateShoppingList(ctx, "61111111-1111-1111-1111-111111111102", "Test", shopping.ListDraft)

	active, err := st.UpdateShoppingListStatus(ctx, list.ID, shopping.ListActive)
	if err != nil {
		t.Fatalf("activate: %v", err)
	}
	if active.Status != shopping.ListActive {
		t.Fatalf("status = %q", active.Status)
	}

	completed, err := st.UpdateShoppingListStatus(ctx, list.ID, shopping.ListCompleted)
	if err != nil {
		t.Fatalf("complete: %v", err)
	}
	if completed.Status != shopping.ListCompleted {
		t.Fatalf("status = %q", completed.Status)
	}

	archived, err := st.UpdateShoppingListStatus(ctx, list.ID, shopping.ListArchived)
	if err != nil {
		t.Fatalf("archive: %v", err)
	}
	if archived.Status != shopping.ListArchived {
		t.Fatalf("status = %q", archived.Status)
	}
}

func TestShopping_Items(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	testSeedShoppingProfile(t, pool, "61111111-1111-1111-1111-111111111103", "71111111-1111-1111-1111-111111111103")

	list, _ := st.CreateShoppingList(ctx, "61111111-1111-1111-1111-111111111103", "Dapur", shopping.ListDraft)

	item, err := st.CreateShoppingItem(ctx, list.ID, "Beras", 2.5, "kg", "manual")
	if err != nil {
		t.Fatalf("CreateShoppingItem: %v", err)
	}
	if item.Quantity != 2.5 {
		t.Fatalf("quantity = %f", item.Quantity)
	}

	items, err := st.ListShoppingItems(ctx, "61111111-1111-1111-1111-111111111103", list.ID, "", 20)
	if err != nil {
		t.Fatalf("ListShoppingItems: %v", err)
	}
	if len(items) < 1 {
		t.Fatal("expected >=1 item")
	}

	name := "Beras Merah"
	updated, err := st.UpdateShoppingItem(ctx, item.ID, &name, nil, nil)
	if err != nil {
		t.Fatalf("UpdateShoppingItem: %v", err)
	}
	if updated.IngredientName != "Beras Merah" {
		t.Fatalf("name = %q", updated.IngredientName)
	}

	completed, err := st.UpdateShoppingItemStatus(ctx, item.ID, shopping.ItemCompleted)
	if err != nil {
		t.Fatalf("complete item: %v", err)
	}
	if completed.Status != shopping.ItemCompleted {
		t.Fatalf("status = %q", completed.Status)
	}

	_, err = st.RemoveShoppingItem(ctx, item.ID)
	if err != nil {
		t.Fatalf("RemoveShoppingItem: %v", err)
	}
}

func TestShopping_ServiceEndToEnd(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	svc := shopping.NewService(st)
	testSeedShoppingProfile(t, pool, "61111111-1111-1111-1111-111111111104", "71111111-1111-1111-1111-111111111104")

	list, err := svc.CreateList(ctx, "61111111-1111-1111-1111-111111111104", "Belanja")
	if err != nil {
		t.Fatalf("CreateList: %v", err)
	}

	_, err = svc.AddItem(ctx, list.ID, "Gula", 1.0, "kg", "manual")
	if err != nil {
		t.Fatalf("AddItem: %v", err)
	}

	lists, _, err := svc.ListLists(ctx, "61111111-1111-1111-1111-111111111104", "", 10)
	if err != nil {
		t.Fatalf("ListLists: %v", err)
	}
	if len(lists) < 1 {
		t.Fatal("expected >=1 list")
	}

	_, err = svc.ActivateList(ctx, list.ID)
	if err != nil {
		t.Fatalf("ActivateList: %v", err)
	}

	_, err = svc.CompleteList(ctx, list.ID)
	if err != nil {
		t.Fatalf("CompleteList: %v", err)
	}

	_, err = svc.ArchiveList(ctx, list.ID)
	if err != nil {
		t.Fatalf("ArchiveList: %v", err)
	}
}
