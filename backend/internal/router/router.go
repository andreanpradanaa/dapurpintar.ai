package router

import (
	"log/slog"

	"github.com/dapurpintar/backend/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type Deps struct {
	Health  *handler.HealthHandler
	Recipes *handler.RecipesHandler
	Log     *slog.Logger
}

func Register(app *fiber.App, d Deps) {
	api := app.Group("/api/v1")

	api.Get("/health", d.Health.Get)

	recipes := api.Group("/recipes")
	recipes.Post("/generate", d.Recipes.Generate)
	recipes.Get("/:slug", d.Recipes.GetBySlug)
}
