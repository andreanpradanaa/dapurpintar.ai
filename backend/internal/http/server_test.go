package http

import (
	"context"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/config"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/identity"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
)

// stubStore satisfies identity.Store via embedding for HTTP-layer tests.
type stubStore struct {
	identity.Store
	accounts map[string]*identity.Account
}

func (s *stubStore) GetAccountByID(ctx context.Context, id string) (*identity.Account, error) {
	if a, ok := s.accounts[id]; ok {
		return a, nil
	}
	return nil, identity.ErrNotFound
}

func newTestServer(t *testing.T) *Server {
	t.Helper()
	cfg := config.Config{
		AppName:           "dapurpintar-test",
		AppEnv:            "test",
		AppPort:           "0",
		AccessTokenTTL:    15 * time.Minute,
		RefreshTokenTTL:   24 * time.Hour,
		SessionCookieName: "dp_session",
	}
	log := logger.New("test")
	tokens, err := auth.NewTokenManager("test-secret", "dapurpintar-test", "dapurpintar-test", cfg.AccessTokenTTL, cfg.RefreshTokenTTL)
	if err != nil {
		t.Fatalf("NewTokenManager: %v", err)
	}
	svc := identity.NewService(&stubStore{accounts: map[string]*identity.Account{
		"account-123": {ID: "account-123", Email: "test@example.com"},
	}}, tokens)
	return New(&cfg, log, tokens, svc)
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
