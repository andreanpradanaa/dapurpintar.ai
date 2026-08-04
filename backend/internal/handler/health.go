package handler

import (
	"github.com/dapurpintar/backend/internal/repo"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	repo    repo.RecipeRepo
	llmName string
}

func NewHealthHandler(r repo.RecipeRepo, llmName string) *HealthHandler {
	return &HealthHandler{repo: r, llmName: llmName}
}

type healthResponse struct {
	Status  string `json:"status"`
	DB      string `json:"db"`
	LLM     string `json:"llm"`
	Recipes int    `json:"recipes"`
}

func (h *HealthHandler) Get(c *fiber.Ctx) error {
	status := fiber.StatusOK
	db := "up"
	count := 0

	n, err := h.repo.Count(c.UserContext())
	if err != nil {
		db = "down"
		status = fiber.StatusServiceUnavailable
	} else {
		count = n
	}

	return c.Status(status).JSON(healthResponse{
		Status:  "ok",
		DB:      db,
		LLM:     h.llmName,
		Recipes: count,
	})
}
