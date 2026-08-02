package http

import (
	"context"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/config"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/http/middleware"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/identity"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/pantry"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
)

// Server bundles the application's HTTP dependencies.
type Server struct {
	app              *fiber.App
	cfg              *config.Config
	log              *slog.Logger
	tokens           *auth.TokenManager
	handler          *Handler
	sessionCookies   middleware.SessionCookies
	refreshCookieTTL time.Duration
}

// Handler carries application dependencies into route handlers.
type Handler struct {
	cfg              *config.Config
	log              *slog.Logger
	tokens           *auth.TokenManager
	identity         *identity.Service
	pantry           *pantry.Service
	sessionCookies   middleware.SessionCookies
	refreshCookieTTL time.Duration
}

// New builds the Fiber application with global middleware and routes.
func New(cfg *config.Config, log *slog.Logger, tokens *auth.TokenManager, identityService *identity.Service, pantryService *pantry.Service) *Server {
	app := fiber.New(fiber.Config{
		AppName:               cfg.AppName,
		DisableStartupMessage: true,
		ReadTimeout:           10 * time.Second,
		WriteTimeout:          30 * time.Second,
		IdleTimeout:           60 * time.Second,
	})

	app.Use(middleware.RequestID())
	app.Use(middleware.Recover())

	sessionCookies := middleware.SessionCookies{
		AccessCookieName:  cfg.SessionCookieName,
		RefreshCookieName: "dp_refresh",
		Secure:            cfg.AppEnv == "production",
	}

	s := &Server{
		app:    app,
		cfg:    cfg,
		log:    log,
		tokens: tokens,
		handler: &Handler{
			cfg:              cfg,
			log:              log,
			tokens:           tokens,
			identity:         identityService,
			pantry:           pantryService,
			sessionCookies:   sessionCookies,
			refreshCookieTTL: tokens.RefreshTTL(),
		},
		sessionCookies:   sessionCookies,
		refreshCookieTTL: tokens.RefreshTTL(),
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
	api.Post("/accounts/refresh", s.handler.refresh)

	authed := api.Group("", middleware.Authenticated(s.tokens, s.cfg.SessionCookieName))
	authed.Get("/accounts/me", s.handler.me)
	authed.Post("/accounts/logout", s.handler.logout)
	authed.Get("/profile", s.handler.getProfile)
	authed.Patch("/profile", s.handler.updateProfile)
	authed.Patch("/profile/preferences", s.handler.updatePreferences)

	authed.Get("/pantry", s.handler.getPantrySummary)
	authed.Get("/pantry/items", s.handler.listPantryItems)
	authed.Post("/pantry/items", s.handler.addPantryItem)
	authed.Get("/pantry/items/:itemId", s.handler.getPantryItem)
	authed.Patch("/pantry/items/:itemId", s.handler.updatePantryItem)
	authed.Delete("/pantry/items/:itemId", s.handler.removePantryItem)
	authed.Get("/pantry/expiry", s.handler.listExpiringItems)
}

// health reports service liveness without exposing internals.
func (h *Handler) health(c *fiber.Ctx) error {
	return response.OK(c, map[string]any{"status": "ok"})
}

// middlewareSubject returns the authenticated subject from the request.
func middlewareSubject(c *fiber.Ctx) string {
	return middleware.Subject(c)
}

// payloadError maps a body parsing failure to the stable M6 code.
func payloadError(err error) *apperr.Error {
	return apperr.New(apperr.CodePayloadMalformed, "The request body is invalid.").WithDetails(
		apperr.Detail{Field: "body", Code: string(apperr.CodePayloadMalformed), Message: "The request body must be valid JSON."})
}
