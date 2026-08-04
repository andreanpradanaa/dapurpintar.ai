package main

import (
	"context"
	"errors"
	"fmt"
	"io"
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
	"github.com/rs/zerolog"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	log := newLogger(cfg.LogLevel, cfg.Env)
	log.Info().
		Str("env", cfg.Env).
		Str("port", cfg.Port).
		Str("model", cfg.OpenAIModel).
		Msg("starting dapur pintar backend")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := connectDB(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("connect db: %w", err)
	}
	defer pool.Close()
	log.Info().Msg("connected to postgres")

	if err := runMigrations(ctx, pool, log); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	log.Info().Msg("migrations applied")

	pgRepo := repo.NewPostgresRepo(pool)
	count, err := pgRepo.Count(ctx)
	if err != nil {
		return fmt.Errorf("count recipes: %w", err)
	}
	if count == 0 {
		log.Info().Msg("recipe library empty, seeding from data/recipes.json")
		if err := repo.SeedFromJSON(ctx, pgRepo, log, "data/recipes.json"); err != nil {
			return fmt.Errorf("seed: %w", err)
		}
	} else {
		log.Info().Int("count", count).Msg("recipe library ready")
	}

	llmClient := llm.NewOpenAIClient(cfg.OpenAIKey, cfg.OpenAIModel, cfg.OpenAIBaseURL, log)
	log.Info().
		Str("provider", llmClient.Name()).
		Str("model", cfg.OpenAIModel).
		Str("base_url", cfg.OpenAIBaseURL).
		Msg("llm client ready")

	gen := service.NewGenerator(pgRepo, llmClient, log)

	healthH := handler.NewHealthHandler(pgRepo, llmClient.Name())
	recipesH := handler.NewRecipesHandler(gen, pgRepo, log)

	app := fiber.New(fiber.Config{
		AppName:               "dapur-pintar-backend",
		DisableStartupMessage: true,
		ReadTimeout:           15 * time.Second,
		WriteTimeout:          60 * time.Second,
		IdleTimeout:           60 * time.Second,
		ErrorHandler:          errorHandler(log),
	})

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

	router.Register(app, router.Deps{
		Health:  healthH,
		Recipes: recipesH,
		Log:     log,
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	errCh := make(chan error, 1)
	go func() {
		addr := ":" + cfg.Port
		log.Info().Str("addr", addr).Msg("listening")
		errCh <- app.Listen(addr)
	}()

	select {
	case sig := <-stop:
		log.Info().Str("signal", sig.String()).Msg("shutdown signal received")
	case err := <-errCh:
		if err != nil && !errors.Is(err, fiber.ErrServiceUnavailable) {
			return err
		}
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("shutdown error")
	}
	log.Info().Msg("shutdown complete")
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

func newLogger(level, env string) *zerolog.Logger {
	lvl, err := zerolog.ParseLevel(strings.ToLower(level))
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	var output io.Writer = os.Stdout
	if env == "development" {
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	}
	logger := zerolog.New(output).Level(lvl).With().Timestamp().Logger()
	return &logger
}

func errorHandler(log *zerolog.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var fe *fiber.Error
		if errors.As(err, &fe) {
			code = fe.Code
		}
		reqID, _ := c.Locals("requestid").(string)
		log.Error().
			Int("status", code).
			Str("err", err.Error()).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("requestId", reqID).
			Msg("unhandled error")
		return c.Status(code).JSON(fiber.Map{
			"error":     "internal_error",
			"message":   err.Error(),
			"requestId": reqID,
		})
	}
}
