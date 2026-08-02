package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pressly/goose/v3"
)

// Migrate applies all pending migrations from the given directory to the
// database. Pending migrations that are not part of the approved deploy are
// treated as a startup error so the application never runs against an
// out-of-date baseline. Migrations follow the M5-002 versioning and ordering
// rules and target the public schema.
func Migrate(ctx context.Context, databaseURL, dir string, log *slog.Logger) error {
	db, err := goose.OpenDBWithDriver("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("open goose db: %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}

	if err := goose.UpContext(ctx, db, dir); err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}

	log.Info("database migrations applied", "dir", dir)
	return nil
}
