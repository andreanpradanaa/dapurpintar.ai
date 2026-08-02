package config

import (
	"os"
	"strconv"
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

	// AI Gateway configuration (M8, ADR-010). AIProvider is the provider
	// adapter key ("openai") and is empty when AI is not configured.
	AIProvider string
	// AIAPIKey is a deployment secret supplied via secret management; it is
	// never logged or written to source control.
	AIAPIKey string
	// AIModel is the approved model profile name (M4-DEC-010).
	AIModel string
	// AITimeout bounds the per-request provider operation.
	AITimeout time.Duration
	// AIMaxRetries bounds automatic retries for transient failures.
	AIMaxRetries int

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
		AIProvider:        getEnv("AI_PROVIDER", ""),
		AIAPIKey:          getEnv("AI_API_KEY", ""),
		AIModel:           getEnv("AI_MODEL", "gpt-4o-mini"),
		AITimeout:         getDurationEnv("AI_TIMEOUT", 30*time.Second),
		AIMaxRetries:      getIntEnv("AI_MAX_RETRIES", 1),
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

func getIntEnv(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}
