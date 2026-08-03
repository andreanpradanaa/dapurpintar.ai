package handler

import (
	"log/slog"

	"github.com/dapurpintar/backend/internal/repo"
	"github.com/dapurpintar/backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RecipesHandler struct {
	gen      *service.Generator
	repo     repo.RecipeRepo
	log      *slog.Logger
	validate *validator.Validate
}

func NewRecipesHandler(gen *service.Generator, r repo.RecipeRepo, log *slog.Logger) *RecipesHandler {
	return &RecipesHandler{
		gen:      gen,
		repo:     r,
		log:      log,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

type generateBody struct {
	Ingredients []string `json:"ingredients" validate:"required,min=1,dive,min=1"`
	Dietary     []string `json:"dietary"`
	Language    string   `json:"language" validate:"omitempty,oneof=en id"`
}

func (h *RecipesHandler) Generate(c *fiber.Ctx) error {
	var body generateBody
	if err := c.BodyParser(&body); err != nil {
		return SendError(c, h.log, fiber.StatusBadRequest, CodeValidation,
			"Invalid JSON body", nil)
	}
	if err := h.validate.Struct(body); err != nil {
		fields := map[string]string{}
		for _, e := range err.(validator.ValidationErrors) {
			fields[e.Field()] = e.Tag()
		}
		return SendError(c, h.log, fiber.StatusBadRequest, CodeValidation,
			"At least one ingredient is required", fields)
	}

	resp, err := h.gen.Generate(c.UserContext(), service.GenerateRequest{
		Ingredients: body.Ingredients,
		Dietary:     body.Dietary,
		Language:    body.Language,
	})
	if err != nil {
		return FromError(c, h.log, err)
	}
	return c.JSON(resp)
}

func (h *RecipesHandler) GetBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return SendError(c, h.log, fiber.StatusBadRequest, CodeValidation,
			"Slug is required", nil)
	}
	r, err := h.repo.GetBySlug(c.UserContext(), slug)
	if err != nil {
		return FromError(c, h.log, err)
	}
	return c.JSON(r)
}
