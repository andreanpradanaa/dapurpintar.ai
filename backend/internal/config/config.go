package config

import (
	"os"
	"time"
)

// Config holds runtime configuration loaded from environment variables.
// Secrets (JWT secret, AI keys) are supplied through deployment secret
// management and are never logged or written to source control.
type Config struct {
	AppName string
	AppEnv  string
	AppPort string

	DatabaseURL string

	RedisURL string

	JWTSecret string

	AccessTokenTTL    time.Duration
	RefreshTokenTTL   time.Duration
	SessionCookieName string

	OTLPEndpoint string
	OTLPDisable  bool

	ShutdownTimeout time.Duration
}

// Load reads configuration from the environment. Required values must be
// present or the application refuses to start rather than running with an
// unsafe or ambiguous baseline.
func Load() Config {
	return Config{
		AppName:           getEnv("APP_NAME", "dapurpintar"),
		AppEnv:            getEnv("APP_ENV", "development"),
		AppPort:           getEnv("APP_PORT", "8080"),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		RedisURL:          getEnv("REDIS_URL", ""),
		JWTSecret:         getEnv("JWT_SECRET", ""),
		AccessTokenTTL:    getDurationEnv("ACCESS_TOKEN_TTL", 15*time.Minute),
		RefreshTokenTTL:   getDurationEnv("REFRESH_TOKEN_TTL", 30*24*time.Hour),
		SessionCookieName: getEnv("SESSION_COOKIE_NAME", "dp_session"),
		OTLPEndpoint:      getEnv("OTLP_ENDPOINT", "http://localhost:4318"),
		OTLPDisable:       getBoolEnv("OTLP_DISABLE", true),
		ShutdownTimeout:   getDurationEnv("SHUTDOWN_TIMEOUT", 10*time.Second),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}

func getBoolEnv(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		return v == "true" || v == "1"
	}
	return fallback
}
