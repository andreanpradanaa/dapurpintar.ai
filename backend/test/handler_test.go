package test

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dapurpintar/backend/internal/handler"
	"github.com/dapurpintar/backend/internal/model"
	"github.com/dapurpintar/backend/internal/repo"
	"github.com/dapurpintar/backend/internal/router"
	"github.com/dapurpintar/backend/internal/service"
	"github.com/dapurpintar/backend/internal/service/llm"
	"github.com/gofiber/fiber/v2"
)

// stubLLM returns a deterministic recipe so handler tests don't hit OpenAI.
type stubLLM struct{}

func (stubLLM) Name() string { return "stub" }

func (stubLLM) GenerateRecipe(_ context.Context, _ llm.GenerateRequest) (*model.Recipe, error) {
	return &model.Recipe{
		ID:    "gen_test",
		Slug:  "test-recipe",
		Title: "Test Recipe",
	}, nil
}

// memoryRepo is a tiny in-memory RecipeRepo for tests.
type memoryRepo struct {
	byID   map[string]*model.Recipe
	bySlug map[string]*model.Recipe
}

func newMemoryRepo() *memoryRepo {
	return &memoryRepo{
		byID:   map[string]*model.Recipe{},
		bySlug: map[string]*model.Recipe{},
	}
}

func (m *memoryRepo) List(_ context.Context) ([]*model.Recipe, error) {
	out := make([]*model.Recipe, 0, len(m.byID))
	for _, r := range m.byID {
		out = append(out, r)
	}
	return out, nil
}

func (m *memoryRepo) GetBySlug(_ context.Context, slug string) (*model.Recipe, error) {
	r, ok := m.bySlug[slug]
	if !ok {
		return nil, repo.ErrNotFound
	}
	return r, nil
}

func (m *memoryRepo) Count(_ context.Context) (int, error) { return len(m.byID), nil }

func (m *memoryRepo) BulkInsert(_ context.Context, recipes []*model.Recipe) error {
	for _, r := range recipes {
		m.byID[r.ID] = r
		m.bySlug[r.Slug] = r
	}
	return nil
}

func (m *memoryRepo) seedSample() {
	_ = m.BulkInsert(context.Background(), []*model.Recipe{
		{
			ID: "r-001", Slug: "nasi-goreng-ayam",
			Title: "Nasi Goreng Ayam", TitleID: "Nasi Goreng Ayam",
			Cuisine: "Indonesian",
			Ingredients: []model.RecipeIngredient{
				{Name: "Rice", NameID: "Nasi", Category: model.CategoryGrain},
				{Name: "Chicken", NameID: "Ayam", Category: model.CategoryProtein},
			},
			Dietary: []model.DietaryTag{model.DietaryHalal},
			Rating:  4.9,
		},
	})
}

func newDiscardLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func setupApp(t *testing.T) *fiber.App {
	t.Helper()
	mr := newMemoryRepo()
	mr.seedSample()
	log := newDiscardLogger()
	gen := service.NewGenerator(mr, stubLLM{}, log)
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	recipesH := handler.NewRecipesHandler(gen, mr, log)
	healthH := handler.NewHealthHandler(mr, "stub")
	router.Register(app, router.Deps{Health: healthH, Recipes: recipesH, Log: log})
	return app
}

// -------- tests --------

func TestHealth(t *testing.T) {
	app := setupApp(t)
	req := httptest.NewRequest("GET", "/api/v1/health", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestGenerate_OK(t *testing.T) {
	app := setupApp(t)
	body := `{"ingredients":["chicken","rice"],"language":"en"}`
	req := httptest.NewRequest("POST", "/api/v1/recipes/generate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		raw, _ := io.ReadAll(resp.Body)
		t.Fatalf("expected 200, got %d: %s", resp.StatusCode, string(raw))
	}
	var got struct {
		Recipe struct {
			Title string `json:"title"`
		} `json:"recipe"`
		MatchScore int `json:"matchScore"`
		Sources    []struct {
			Title string `json:"title"`
		} `json:"sources"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if got.Recipe.Title == "" {
		t.Fatal("expected recipe title")
	}
	if got.MatchScore <= 0 {
		t.Fatal("expected positive match score")
	}
	if len(got.Sources) == 0 {
		t.Fatal("expected at least one source")
	}
}

func TestGenerate_EmptyIngredients(t *testing.T) {
	app := setupApp(t)
	body := `{"ingredients":[]}`
	req := httptest.NewRequest("POST", "/api/v1/recipes/generate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestGetBySlug_NotFound(t *testing.T) {
	app := setupApp(t)
	req := httptest.NewRequest("GET", "/api/v1/recipes/does-not-exist", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 404 {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func TestGetBySlug_OK(t *testing.T) {
	app := setupApp(t)
	req := httptest.NewRequest("GET", "/api/v1/recipes/nasi-goreng-ayam", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestMain(m *testing.M) {
	_ = os.Chdir(filepath.Join(".."))
	os.Exit(m.Run())
}
