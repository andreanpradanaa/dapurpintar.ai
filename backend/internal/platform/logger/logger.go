package logger

import (
	"log/slog"
	"os"
)

// New creates a structured slog logger. Development uses human-friendly text
// output; production uses JSON for machine parsing. All logging follows the
// redaction policy: credentials, tokens, provider secrets, and private kitchen
// context are never logged.
func New(env string) *slog.Logger {
	level := slog.LevelInfo
	if env == "development" {
		level = slog.LevelDebug
	}

	var handler slog.Handler
	handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	if env == "development" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}

	return slog.New(handler)
}
