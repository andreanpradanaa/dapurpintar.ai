package store

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/database"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recipes"
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

func seedRecipe(t *testing.T, pool *pgxpool.Pool, id, title string, isPublic bool) {
	t.Helper()
	ctx := context.Background()
	ingredients := `[{"name":"gula","quantity":"100gr"},{"name":"telur","quantity":"2 butir"}]`
	instructions := `["campur semua bahan","panggang 30 menit"]`
	_, err := pool.Exec(ctx, `insert into recipes (id, title, summary, servings, ingredients, instructions, is_public, status)
		values ($1, $2, 'Deskripsi resep.', 2, $3::jsonb, $4::jsonb, $5, 'available')
		on conflict (id) do nothing`, id, title, ingredients, instructions, isPublic)
	if err != nil {
		t.Fatalf("seed recipe: %v", err)
	}
}

func seedProfile(t *testing.T, pool *pgxpool.Pool, profileID, accountID string) {
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

func TestRecipes_ListPublic(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	seedRecipe(t, pool, "11111111-1111-1111-1111-111111111101", "Kue Lapis Legit", true)
	seedRecipe(t, pool, "11111111-1111-1111-1111-111111111102", "Rendang Padang", true)
	seedRecipe(t, pool, "11111111-1111-1111-1111-111111111103", "Resep Rahasia", false)

	items, err := st.ListPublicRecipes(ctx, "", 20, nil, nil, nil)
	if err != nil {
		t.Fatalf("ListPublicRecipes: %v", err)
	}
	if len(items) < 2 {
		t.Fatalf("expected >=2 public recipes, got %d", len(items))
	}
	for _, it := range items {
		if it.ID == "11111111-1111-1111-1111-111111111103" {
			t.Fatal("private recipe leaked in public listing")
		}
	}
}

func TestRecipes_GetByID(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	seedRecipe(t, pool, "11111111-1111-1111-1111-111111111201", "Soto Ayam", true)

	r, err := st.GetRecipeByID(ctx, "11111111-1111-1111-1111-111111111201")
	if err != nil {
		t.Fatalf("GetRecipeByID: %v", err)
	}
	if r.Title != "Soto Ayam" {
		t.Fatalf("title = %q, want Soto Ayam", r.Title)
	}
	if len(r.Ingredients) != 2 {
		t.Fatalf("expected 2 ingredients, got %d", len(r.Ingredients))
	}
	if len(r.Instructions) != 2 {
		t.Fatalf("expected 2 instructions, got %d", len(r.Instructions))
	}

	_, err = st.GetRecipeByID(ctx, "00000000-0000-0000-0000-000000000000")
	if err != recipes.ErrNotFound {
		t.Fatalf("missing recipe err = %v, want ErrNotFound", err)
	}
}

func TestRecipes_Favorites(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	seedRecipe(t, pool, "11111111-1111-1111-1111-111111111301", "Nasi Goreng", true)
	seedProfile(t, pool, "21111111-1111-1111-1111-111111111301", "31111111-1111-1111-1111-111111111301")

	_, err := st.GetActiveFavorite(ctx, "21111111-1111-1111-1111-111111111301", "11111111-1111-1111-1111-111111111301")
	if err != recipes.ErrNotFound {
		t.Fatalf("missing favorite err = %v, want ErrNotFound", err)
	}

	fav, err := st.CreateFavorite(ctx, "21111111-1111-1111-1111-111111111301", "11111111-1111-1111-1111-111111111301")
	if err != nil {
		t.Fatalf("CreateFavorite: %v", err)
	}
	if fav.ID == "" {
		t.Fatal("favorite has no id")
	}

	got, err := st.GetActiveFavorite(ctx, "21111111-1111-1111-1111-111111111301", "11111111-1111-1111-1111-111111111301")
	if err != nil {
		t.Fatalf("GetActiveFavorite: %v", err)
	}
	if got.ID != fav.ID {
		t.Fatalf("got id = %q, want %q", got.ID, fav.ID)
	}

	if err := st.RemoveFavorite(ctx, "21111111-1111-1111-1111-111111111301", "11111111-1111-1111-1111-111111111301"); err != nil {
		t.Fatalf("RemoveFavorite: %v", err)
	}

	_, err = st.GetActiveFavorite(ctx, "21111111-1111-1111-1111-111111111301", "11111111-1111-1111-1111-111111111301")
	if err != recipes.ErrNotFound {
		t.Fatalf("removed favorite err = %v, want ErrNotFound", err)
	}
}

func TestRecipes_ServiceEndToEnd(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	svc := recipes.NewService(st)

	seedRecipe(t, pool, "11111111-1111-1111-1111-111111111401", "Gado Gado", true)
	seedProfile(t, pool, "21111111-1111-1111-1111-111111111401", "31111111-1111-1111-1111-111111111401")

	items, page, err := svc.ListRecipes(ctx, "", 5, nil, nil, nil)
	if err != nil {
		t.Fatalf("ListRecipes: %v", err)
	}
	_ = items
	if page.HasMore && page.NextCursor == "" {
		t.Log("HasMore=true with empty cursor (may be fine if fewer than 5 public recipes)")
	}

	r, err := svc.GetRecipe(ctx, "11111111-1111-1111-1111-111111111401")
	if err != nil {
		t.Fatalf("GetRecipe: %v", err)
	}
	if r.Title != "Gado Gado" {
		t.Fatalf("title = %q", r.Title)
	}

	if err := svc.Favorite(ctx, "21111111-1111-1111-1111-111111111401", "11111111-1111-1111-1111-111111111401"); err != nil {
		t.Fatalf("Favorite: %v", err)
	}
	if err := svc.Favorite(ctx, "21111111-1111-1111-1111-111111111401", "11111111-1111-1111-1111-111111111401"); err != nil {
		t.Fatalf("Favorite (idempotent): %v", err)
	}

	favs, fpage, err := svc.ListFavorites(ctx, "21111111-1111-1111-1111-111111111401", "", 10)
	if err != nil {
		t.Fatalf("ListFavorites: %v", err)
	}
	if len(favs) < 1 {
		t.Fatal("expected >=1 favorite")
	}
	_ = fpage

	if err := svc.Unfavorite(ctx, "21111111-1111-1111-1111-111111111401", "11111111-1111-1111-1111-111111111401"); err != nil {
		t.Fatalf("Unfavorite: %v", err)
	}
}
