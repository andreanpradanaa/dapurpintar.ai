package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	// Best-effort load of .env. We don't fail if the file is missing —
	// env vars may be set directly in production / Docker.
	_ = godotenv.Load()
}

type Config struct {
	Port          string
	Env           string
	LogLevel      string
	DatabaseURL   string
	OpenAIKey     string
	OpenAIModel   string
	OpenAIBaseURL string
	OpenAITimeout time.Duration
	CORSOrigins   []string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:          getEnv("PORT", "8080"),
		Env:           getEnv("ENV", "development"),
		LogLevel:      getEnv("LOG_LEVEL", "info"),
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		OpenAIKey:     getEnv("OPENAI_API_KEY", ""),
		OpenAIModel:   getEnv("OPENAI_MODEL", "gpt-4o-mini"),
		OpenAIBaseURL: getEnv("OPENAI_BASE_URL", "https://api.openai.com/v1"),
		OpenAITimeout: getEnvDuration("OPENAI_TIMEOUT", 30*time.Second),
		CORSOrigins:   getEnvList("CORS_ORIGINS", []string{"http://localhost:3000"}),
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.OpenAIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY is required")
	}
	return cfg, nil
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func getEnvDuration(key string, def time.Duration) time.Duration {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return def
	}
	return d
}

func getEnvList(key string, def []string) []string {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	if len(out) == 0 {
		return def
	}
	return out
}

func getEnvInt(key string, def int) int {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
