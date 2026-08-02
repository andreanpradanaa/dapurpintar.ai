package cache

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

// New opens a Redis client. Redis is supporting infrastructure only and is
// never an authorization source of truth. A Redis failure must fail closed for
// security-critical decisions; callers guard accordingly.
func New(ctx context.Context, redisURL string, log *slog.Logger) (*redis.Client, error) {
	if redisURL == "" {
		log.Warn("REDIS_URL is empty; redis-backed features will be unavailable")
		return nil, nil
	}

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}

	client := redis.NewClient(opts)

	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := client.Ping(pingCtx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	log.Info("connected to redis")
	return client, nil
}
