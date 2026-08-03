package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

// runMigrations applies every .up.sql file in /migrations to the
// database, in lexical order. Idempotent — each file uses CREATE
// TABLE IF NOT EXISTS / CREATE INDEX IF NOT EXISTS so re-running is
// safe. Path resolution works in dev, in a built binary, and in
// Docker where the working directory is /app.
func runMigrations(ctx context.Context, pool *pgxpool.Pool, log *zerolog.Logger) error {
	dirs := []string{}
	if cwd, err := os.Getwd(); err == nil {
		dirs = append(dirs,
			filepath.Join(cwd, "migrations"),
			filepath.Join(cwd, "..", "migrations"),
			filepath.Join(cwd, "backend", "migrations"),
		)
	}

	var dir string
	for _, d := range dirs {
		if info, err := os.Stat(d); err == nil && info.IsDir() {
			dir = d
			break
		}
	}
	if dir == "" {
		return fmt.Errorf("migrations dir not found (tried %v)", dirs)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read migrations dir %s: %w", dir, err)
	}

	ups := make([]string, 0, len(entries))
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".up.sql") {
			ups = append(ups, e.Name())
		}
	}
	sort.Strings(ups)

	for _, name := range ups {
		raw, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return fmt.Errorf("read %s: %w", name, err)
		}
		log.Info().Str("file", name).Msg("applying migration")
		if _, err := pool.Exec(ctx, string(raw)); err != nil {
			return fmt.Errorf("apply %s: %w", name, err)
		}
	}
	return nil
}
