package store

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/mealplan"
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

func testSeedMealProfile(t *testing.T, pool *pgxpool.Pool, profileID, accountID string) {
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

func TestMealPlan_CreateAndList(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	testSeedMealProfile(t, pool, "41111111-1111-1111-1111-111111111101", "51111111-1111-1111-1111-111111111101")

	ps, _ := time.Parse("2006-01-02", "2026-08-10")
	pe, _ := time.Parse("2006-01-02", "2026-08-16")
	plan, err := st.CreateMealPlan(ctx, "41111111-1111-1111-1111-111111111101", ps, pe, "Menu Mingguan")
	if err != nil {
		t.Fatalf("CreateMealPlan: %v", err)
	}
	if plan.Status != mealplan.PlanDraft {
		t.Fatalf("status = %q, want draft", plan.Status)
	}

	got, err := st.GetMealPlanByID(ctx, plan.ID)
	if err != nil {
		t.Fatalf("GetMealPlanByID: %v", err)
	}
	if got.Title != "Menu Mingguan" {
		t.Fatalf("title = %q", got.Title)
	}

	plans, err := st.ListMealPlans(ctx, "41111111-1111-1111-1111-111111111101", "", 20)
	if err != nil {
		t.Fatalf("ListMealPlans: %v", err)
	}
	if len(plans) < 1 {
		t.Fatal("expected >=1 plan")
	}
}

func TestMealPlan_Lifecycle(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	testSeedMealProfile(t, pool, "41111111-1111-1111-1111-111111111102", "51111111-1111-1111-1111-111111111102")

	ps, _ := time.Parse("2006-01-02", "2026-09-01")
	pe, _ := time.Parse("2006-01-02", "2026-09-07")
	plan, _ := st.CreateMealPlan(ctx, "41111111-1111-1111-1111-111111111102", ps, pe, "Rencana")

	cancelled, err := st.UpdateMealPlanStatus(ctx, plan.ID, mealplan.PlanCancelled)
	if err != nil {
		t.Fatalf("cancel: %v", err)
	}
	if cancelled.Status != mealplan.PlanCancelled {
		t.Fatalf("status = %q, want cancelled", cancelled.Status)
	}
}

func TestPlannedMeal_CRUDAndSlotConflict(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	testSeedMealProfile(t, pool, "41111111-1111-1111-1111-111111111103", "51111111-1111-1111-1111-111111111103")

	ps, _ := time.Parse("2006-01-02", "2026-10-01")
	pe, _ := time.Parse("2006-01-02", "2026-10-07")
	plan, _ := st.CreateMealPlan(ctx, "41111111-1111-1111-1111-111111111103", ps, pe, "Oktober")

	md, _ := time.Parse("2006-01-02", "2026-10-03")
	meal, err := st.CreatePlannedMeal(ctx, plan.ID, md, "dinner", nil, nil)
	if err != nil {
		t.Fatalf("CreatePlannedMeal: %v", err)
	}
	if meal.MealOccasion != "dinner" {
		t.Fatalf("occasion = %q", meal.MealOccasion)
	}

	dup, err := st.GetPlannedMealBySlot(ctx, plan.ID, md, "dinner")
	if err != nil {
		t.Fatalf("GetPlannedMealBySlot: %v", err)
	}
	if dup.ID != meal.ID {
		t.Fatalf("slot lookup returned %q", dup.ID)
	}

	meals, err := st.ListPlannedMeals(ctx, "41111111-1111-1111-1111-111111111103", plan.ID, "", 20)
	if err != nil {
		t.Fatalf("ListPlannedMeals: %v", err)
	}
	if len(meals) < 1 {
		t.Fatal("expected >=1 meal")
	}

	occasion := "lunch"
	updated, err := st.UpdatePlannedMeal(ctx, meal.ID, &occasion, nil, nil)
	if err != nil {
		t.Fatalf("UpdatePlannedMeal: %v", err)
	}
	if updated.MealOccasion != "lunch" {
		t.Fatalf("occasion = %q, want lunch", updated.MealOccasion)
	}

	_, err = st.RemovePlannedMeal(ctx, meal.ID)
	if err != nil {
		t.Fatalf("RemovePlannedMeal: %v", err)
	}
}

func TestMealPlan_ServiceEndToEnd(t *testing.T) {
	st, pool := testStore(t)
	ctx := context.Background()
	svc := mealplan.NewService(st)
	testSeedMealProfile(t, pool, "41111111-1111-1111-1111-111111111104", "51111111-1111-1111-1111-111111111104")

	ps, _ := time.Parse("2006-01-02", "2026-11-01")
	pe, _ := time.Parse("2006-01-02", "2026-11-07")
	plan, err := svc.CreateMealPlan(ctx, "41111111-1111-1111-1111-111111111104", ps, pe, "November")
	if err != nil {
		t.Fatalf("CreateMealPlan: %v", err)
	}

	md, _ := time.Parse("2006-01-02", "2026-11-03")
	meal, err := svc.PlanMeal(ctx, plan.ID, md, "breakfast", nil, nil)
	if err != nil {
		t.Fatalf("PlanMeal: %v", err)
	}

	_, err = svc.PlanMeal(ctx, plan.ID, md, "breakfast", nil, nil)
	if err == nil {
		t.Fatal("expected slot conflict")
	}

	_, err = svc.UpdatePlannedMeal(ctx, meal.ID, nil, nil, ptr("completed"))
	if err != nil {
		t.Fatalf("UpdatePlannedMeal: %v", err)
	}

	updated, err := svc.GetPlannedMeal(ctx, meal.ID)
	if err != nil {
		t.Fatalf("GetPlannedMeal: %v", err)
	}
	if updated.RecipeID != nil {
		t.Fatalf("unexpected recipe_id: %v", updated.RecipeID)
	}
	_ = updated

	_, err = svc.CompleteMealPlan(ctx, plan.ID)
	if err != nil {
		t.Fatalf("CompleteMealPlan: %v", err)
	}

	_, err = svc.CancelMealPlan(ctx, plan.ID)
	if err == nil {
		t.Fatal("expected state conflict when cancelling completed plan")
	}
}

func ptr[T any](v T) *T { return &v }
