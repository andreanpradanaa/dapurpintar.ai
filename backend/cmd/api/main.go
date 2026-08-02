package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/ai/openai"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/config"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/http"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/identity"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/identity/store"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/mealplan"
	mealplanStore "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/mealplan/store"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/pantry"
	pantryStore "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/pantry/store"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/cache"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/database"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/telemetry"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recipes"
	recipesStore "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/recipes/store"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.AppEnv)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Telemetry is initialized first so spans are correlated from startup.
	shutdownTelemetry, err := telemetry.Setup(ctx, cfg.AppName, cfg.OTLPEndpoint, cfg.OTLPDisable, log)
	if err != nil {
		log.Error("telemetry setup failed", "error", err)
		os.Exit(1)
	}
	defer shutdownTelemetry(context.Background())

	// PostgreSQL is the system of record; a missing or unreachable database is
	// a startup error.
	pool, err := database.New(ctx, cfg.DatabaseURL, log)
	if err != nil {
		log.Error("database setup failed", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := database.Migrate(ctx, cfg.DatabaseURL, "migrations", log); err != nil {
		log.Error("database migration failed", "error", err)
		os.Exit(1)
	}

	// Redis is supporting infrastructure; its absence is non-fatal at startup
	// but security-critical features must fail closed without it.
	_, _ = cache.New(ctx, cfg.RedisURL, log)

	tokens, err := auth.NewTokenManager(cfg.JWTSecret, cfg.AppName, cfg.AppName, cfg.AccessTokenTTL, cfg.RefreshTokenTTL)
	if err != nil {
		log.Error("token manager setup failed", "error", err)
		os.Exit(1)
	}

	identityStore := store.New(pool)
	identityService := identity.NewService(identityStore, tokens)

	pantryStore := pantryStore.New(pool)
	pantryService := pantry.NewService(pantryStore)

	recipesStr := recipesStore.New(pool)
	rsvc := recipes.NewService(recipesStr)

	mps := mealplanStore.New(pool)
	mpsvc := mealplan.NewService(mps)

	// The AI Gateway is an optional dependency (docs/architecture/ai-architecture.md).
	// When AI is not configured, core non-AI operations remain fully usable and
	// AI features fail closed with a bounded unavailable error.
	if cfg.AIProvider != "" {
		adapter, err := openai.New(openai.Config{
			APIKey:     cfg.AIAPIKey,
			Timeout:    cfg.AITimeout,
			MaxRetries: cfg.AIMaxRetries,
			Log:        log,
		})
		if err != nil {
			log.Error("ai gateway setup failed", "error", err)
			os.Exit(1)
		}
		log.Info("ai gateway configured", "provider", cfg.AIProvider, "model", cfg.AIModel)
		_ = adapter
	}

	server := http.New(&cfg, log, tokens, identityService, pantryService, rsvc, mpsvc)

	go func() {
		log.Info("server listening", "port", cfg.AppPort)
		if err := server.Listen(); err != nil && !errors.Is(err, os.ErrClosed) {
			log.Error("server error", "error", err)
			stop()
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown error", "error", err)
	}
	log.Info("server stopped")
}
