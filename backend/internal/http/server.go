package http

import (
	"context"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/config"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/http/middleware"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
)

// Server bundles the application's HTTP dependencies.
type Server struct {
	app     *fiber.App
	cfg     *config.Config
	log     *slog.Logger
	tokens  *auth.TokenManager
	handler *Handler
}

// Handler carries application dependencies into route handlers.
type Handler struct {
	cfg    *config.Config
	log    *slog.Logger
	tokens *auth.TokenManager
}

// New builds the Fiber application with global middleware and routes.
func New(cfg *config.Config, log *slog.Logger, tokens *auth.TokenManager) *Server {
	app := fiber.New(fiber.Config{
		AppName:               cfg.AppName,
		DisableStartupMessage: true,
		ReadTimeout:           10 * time.Second,
		WriteTimeout:          30 * time.Second,
		IdleTimeout:           60 * time.Second,
	})

	app.Use(middleware.RequestID())
	app.Use(middleware.Recover())

	s := &Server{
		app:    app,
		cfg:    cfg,
		log:    log,
		tokens: tokens,
		handler: &Handler{
			cfg:    cfg,
			log:    log,
			tokens: tokens,
		},
	}
	s.registerRoutes()

	return s
}

// App exposes the underlying Fiber app (used for tests).
func (s *Server) App() *fiber.App { return s.app }

// Listen starts the HTTP server.
func (s *Server) Listen() error {
	return s.app.Listen(":" + s.cfg.AppPort)
}

// Shutdown gracefully shuts the server down within the configured timeout.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func (s *Server) registerRoutes() {
	s.app.Get("/health", s.handler.health)

	api := s.app.Group("/api/v1")

	api.Post("/accounts", s.handler.register)
	api.Post("/accounts/login", s.handler.login)

	authed := api.Group("", middleware.Authenticated(s.tokens))
	authed.Get("/accounts/me", s.handler.me)
	authed.Post("/accounts/logout", s.handler.logout)
}

// health reports service liveness without exposing internals.
func (h *Handler) health(c *fiber.Ctx) error {
	return response.OK(c, map[string]any{"status": "ok"})
}

// register is a placeholder for the Identity and Access registration use case.
func (h *Handler) register(c *fiber.Ctx) error {
	return response.Error(c, nil)
}

// login is a placeholder for the Identity and Access login use case.
func (h *Handler) login(c *fiber.Ctx) error {
	return response.Error(c, nil)
}

// me returns the current account participation context.
func (h *Handler) me(c *fiber.Ctx) error {
	return response.OK(c, map[string]any{"subject": middleware.Subject(c)})
}

// logout revokes the current session.
func (h *Handler) logout(c *fiber.Ctx) error {
	return response.NoContent(c)
}
