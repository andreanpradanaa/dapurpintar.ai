package store

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/database"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recommendation"
)

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
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

func testSeedRecProfile(t *testing.T, pool *pgxpool.Pool, profileID, accountID string) {
	t.Helper()
	ctx := context.Background()
	pool.Exec(ctx, `insert into accounts (id, email, password_hash, status) values ($1, $2, 'test', 'active') on conflict (id) do nothing`, accountID, accountID+"@test.example")
	pool.Exec(ctx, `insert into user_profiles (id, account_id, display_name, status) values ($1, $2, 'Test', 'created') on conflict (id) do nothing`, profileID, accountID)
}

func TestRecommendation_CRUD(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	testSeedRecProfile(t, pool, "81111111-1111-1111-1111-111111111101", "91111111-1111-1111-1111-111111111101")

	rec, err := st.CreateRecommendation(ctx, "81111111-1111-1111-1111-111111111101", []byte("{}"), "kitchen-recommendation")
	if err != nil {
		t.Fatalf("CreateRecommendation: %v", err)
	}
	if rec.Status != recommendation.StatusRequested {
		t.Fatalf("status = %q", rec.Status)
	}

	got, err := st.GetRecommendationByID(ctx, rec.ID)
	if err != nil {
		t.Fatalf("GetRecommendationByID: %v", err)
	}
	if got.ID != rec.ID {
		t.Fatalf("id mismatch")
	}

	presented, err := st.UpdateRecommendationStatus(ctx, rec.ID, recommendation.StatusPresented)
	if err != nil {
		t.Fatalf("present: %v", err)
	}
	if presented.Status != recommendation.StatusPresented {
		t.Fatalf("status = %q", presented.Status)
	}

	recs, err := st.ListRecommendations(ctx, "81111111-1111-1111-1111-111111111101", "", 20, nil, nil)
	if err != nil {
		t.Fatalf("ListRecommendations: %v", err)
	}
	if len(recs) < 1 {
		t.Fatal("expected >=1 recommendation")
	}
}

func TestRecommendation_OptionsAndService(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	svc := recommendation.NewService(st)
	testSeedRecProfile(t, pool, "81111111-1111-1111-1111-111111111102", "91111111-1111-1111-1111-111111111102")

	rec, _ := svc.Request(ctx, "81111111-1111-1111-1111-111111111102", "kitchen-recommendation", nil, true)

	opt, err := st.CreateRecommendationOption(ctx, rec.ID, nil, 1, "Coba ayam goreng")
	if err != nil {
		t.Fatalf("CreateOption: %v", err)
	}

	opts, err := st.ListRecommendationOptions(ctx, rec.ID)
	if err != nil {
		t.Fatalf("ListOptions: %v", err)
	}
	if len(opts) < 1 {
		t.Fatal("expected >=1 option")
	}

	_, _, err = svc.Get(ctx, rec.ID)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	_, err = st.UpdateRecommendationStatus(ctx, rec.ID, recommendation.StatusCreated)
	if err != nil {
		t.Fatalf("set created: %v", err)
	}

	presented, err := svc.Present(ctx, rec.ID)
	if err != nil {
		t.Fatalf("Present: %v", err)
	}
	if presented.Status != recommendation.StatusPresented {
		t.Fatalf("status = %q", presented.Status)
	}

	_, err = svc.AcceptOption(ctx, rec.ID, opt.ID)
	if err != nil {
		t.Fatalf("AcceptOption: %v", err)
	}
}

func TestRecommendation_StateGuards(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	svc := recommendation.NewService(st)
	testSeedRecProfile(t, pool, "81111111-1111-1111-1111-111111111103", "91111111-1111-1111-1111-111111111103")

	rec, _ := svc.Request(ctx, "81111111-1111-1111-1111-111111111103", "test", nil, true)

	_, err := svc.Present(ctx, rec.ID)
	if err == nil {
		t.Fatal("expected state guard: cannot present requested")
	}

	_, err = st.UpdateRecommendationStatus(ctx, rec.ID, recommendation.StatusCreated)
	_ = err
	st.CreateRecommendationOption(ctx, rec.ID, nil, 1, "opsi 1")

	_, err = svc.Reject(ctx, rec.ID)
	if err != nil {
		t.Fatalf("Reject from created: %v", err)
	}

	_, err = svc.Supersede(ctx, rec.ID)
	if err == nil {
		t.Fatal("expected state guard: cannot supersede rejected")
	}
}
