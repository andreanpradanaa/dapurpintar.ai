package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/dapurpintar/backend/internal/config"
	"github.com/dapurpintar/backend/internal/handler"
	"github.com/dapurpintar/backend/internal/middleware"
	"github.com/dapurpintar/backend/internal/repo"
	"github.com/dapurpintar/backend/internal/router"
	"github.com/dapurpintar/backend/internal/service"
	"github.com/dapurpintar/backend/internal/service/llm"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	log := newLogger(cfg.LogLevel)
	log.Info("starting dapur pintar backend",
		"env", cfg.Env,
		"port", cfg.Port,
		"model", cfg.OpenAIModel,
	)

	// Connect to Postgres
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := connectDB(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("connect db: %w", err)
	}
	defer pool.Close()
	log.Info("connected to postgres")

	// Auto-migrate
	if err := runMigrations(ctx, pool, log); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	log.Info("migrations applied")

	// Auto-seed if empty
	pgRepo := repo.NewPostgresRepo(pool)
	count, err := pgRepo.Count(ctx)
	if err != nil {
		return fmt.Errorf("count recipes: %w", err)
	}
	if count == 0 {
		log.Info("recipe library empty, seeding from data/recipes.json")
		if err := repo.SeedFromJSON(ctx, pgRepo, log, "data/recipes.json"); err != nil {
			return fmt.Errorf("seed: %w", err)
		}
	} else {
		log.Info("recipe library ready", "count", count)
	}

	// Build LLM client
	llmClient := llm.NewOpenAIClient(cfg.OpenAIKey, cfg.OpenAIModel, cfg.OpenAIBaseURL, log)
	log.Info("llm client ready",
		"provider", llmClient.Name(),
		"model", cfg.OpenAIModel,
		"base_url", cfg.OpenAIBaseURL,
	)

	// Build services
	gen := service.NewGenerator(pgRepo, llmClient, log)

	// Build handlers
	healthH := handler.NewHealthHandler(pgRepo, llmClient.Name())
	recipesH := handler.NewRecipesHandler(gen, pgRepo, log)

	// Build fiber app
	app := fiber.New(fiber.Config{
		AppName:               "dapur-pintar-backend",
		DisableStartupMessage: true,
		ReadTimeout:           15 * time.Second,
		WriteTimeout:          60 * time.Second, // LLM calls can take time
		IdleTimeout:           60 * time.Second,
		ErrorHandler:          errorHandler(log),
	})

	// Middleware
	app.Use(requestid.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: false}))
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} ${latency} ${method} ${path} ${locals:requestid}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "UTC",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(cfg.CORSOrigins, ","),
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization",
		AllowCredentials: false,
		MaxAge:           86400,
	}))
	app.Use(middleware.Recover(log))

	// Routes
	router.Register(app, router.Deps{
		Health:  healthH,
		Recipes: recipesH,
		Log:     log,
	})

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	errCh := make(chan error, 1)
	go func() {
		addr := ":" + cfg.Port
		log.Info("listening", "addr", addr)
		errCh <- app.Listen(addr)
	}()

	select {
	case sig := <-stop:
		log.Info("shutdown signal received", "signal", sig.String())
	case err := <-errCh:
		if err != nil && !errors.Is(err, fiber.ErrServiceUnavailable) {
			return err
		}
	}

	// Graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Error("shutdown error", "err", err)
	}
	log.Info("shutdown complete")
	return nil
}

func connectDB(ctx context.Context, url string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute

	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(pingCtx, cfg)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

func newLogger(level string) *slog.Logger {
	var lvl slog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn", "warning":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}))
}

func errorHandler(log *slog.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var fe *fiber.Error
		if errors.As(err, &fe) {
			code = fe.Code
		}
		reqID, _ := c.Locals("requestid").(string)
		log.Error("unhandled error",
			"status", code,
			"err", err.Error(),
			"path", c.Path(),
			"method", c.Method(),
			"requestId", reqID,
		)
		return c.Status(code).JSON(fiber.Map{
			"error":     "internal_error",
			"message":   err.Error(), // expose for dev; switch to generic in prod
			"requestId": reqID,
		})
	}
}
