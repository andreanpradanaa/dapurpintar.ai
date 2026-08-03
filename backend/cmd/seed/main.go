package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/dapurpintar/backend/internal/config"
	"github.com/dapurpintar/backend/internal/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

// seed is a standalone command to load data/recipes.json into the
// database. Useful for prod deploys where you want to control when
// the seed happens. The server's auto-seed on boot handles dev.
//
// Usage:  go run ./cmd/seed
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer pool.Close()

	pgRepo := repo.NewPostgresRepo(pool)

	count, err := pgRepo.Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Info("library already populated, skipping", "count", count)
		return nil
	}

	if err := repo.SeedFromJSON(ctx, pgRepo, log, "data/recipes.json"); err != nil {
		return err
	}
	return nil
}
