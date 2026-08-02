package http

import (
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/config"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
)

func newTestServer(t *testing.T) *Server {
	t.Helper()
	cfg := config.Config{
		AppName:         "dapurpintar-test",
		AppEnv:          "test",
		AppPort:         "0",
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 24 * time.Hour,
	}
	log := logger.New("test")
	tokens, err := auth.NewTokenManager("test-secret", "dapurpintar-test", "dapurpintar-test", cfg.AccessTokenTTL, cfg.RefreshTokenTTL)
	if err != nil {
		t.Fatalf("NewTokenManager: %v", err)
	}
	return New(&cfg, log, tokens)
}

func TestHealth(t *testing.T) {
	s := newTestServer(t)
	req, _ := s.app.Test(createTestRequest(fiber.MethodGet, "/health", nil))
	if req.StatusCode != fiber.StatusOK {
		t.Fatalf("expected 200, got %d", req.StatusCode)
	}
}

func TestProtectedRouteRequiresSession(t *testing.T) {
	s := newTestServer(t)
	req, _ := s.app.Test(createTestRequest(fiber.MethodGet, "/api/v1/accounts/me", nil))
	if req.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", req.StatusCode)
	}
}

func TestProtectedRouteWithValidToken(t *testing.T) {
	s := newTestServer(t)

	session, err := s.tokens.Issue("account-123")
	if err != nil {
		t.Fatalf("Issue: %v", err)
	}

	req := createTestRequest(fiber.MethodGet, "/api/v1/accounts/me", nil)
	req.Header.Set("Authorization", "Bearer "+session.AccessToken)

	resp, _ := s.app.Test(req)
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}
