package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/dapurpintar/backend/internal/model"
)

// SeedFromJSON loads recipes from a JSON file and bulk-inserts them
// into the given repo. Idempotent (ON CONFLICT). Used by both the
// server's auto-seed on boot and the standalone seed command.
//
// The path is resolved relative to the working directory, with
// sensible fallbacks so the same code works in dev (go run), in a
// built binary, and in a Docker container where CWD is /app.
func SeedFromJSON(ctx context.Context, r RecipeRepo, log *slog.Logger, explicitPath string) error {
	paths := []string{}
	if explicitPath != "" {
		paths = append(paths, explicitPath)
	}
	if cwd, err := os.Getwd(); err == nil {
		paths = append(paths,
			filepath.Join(cwd, "data", "recipes.json"),
			filepath.Join(cwd, "..", "data", "recipes.json"),
			filepath.Join(cwd, "backend", "data", "recipes.json"),
		)
	}

	var raw []byte
	var lastErr error
	for _, p := range paths {
		b, err := os.ReadFile(p)
		if err == nil {
			raw = b
			log.Info("loaded recipes.json", "path", p)
			break
		}
		lastErr = err
	}
	if raw == nil {
		return fmt.Errorf("read recipes.json (tried %d paths): %w", len(paths), lastErr)
	}

	var recipes []*model.Recipe
	if err := json.Unmarshal(raw, &recipes); err != nil {
		return fmt.Errorf("parse recipes.json: %w", err)
	}

	log.Info("seeding recipes", "count", len(recipes))
	if err := r.BulkInsert(ctx, recipes); err != nil {
		return err
	}
	log.Info("seed complete", "count", len(recipes))
	return nil
}
