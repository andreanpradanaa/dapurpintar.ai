package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
)

// RequestID assigns a correlation identifier to every request and echoes it in
// the X-Request-ID response header and the response envelope.
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Get("X-Request-ID")
		if id == "" {
			id = newRequestID()
		}
		c.Locals("request_id", id)
		c.Set("X-Request-ID", id)
		return c.Next()
	}
}

func newRequestID() string {
	return fmt.Sprintf("req_%d", time.Now().UnixNano())
}

// Recover converts panics into a safe internal error response without exposing
// stack traces or internal details.
func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				_ = response.Error(c, nil)
			}
		}()
		return c.Next()
	}
}
