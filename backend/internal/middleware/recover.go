package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// Recover catches panics and returns a uniform 500 response.
func Recover(log *zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error().
					Interface("panic", r).
					Str("path", c.Path()).
					Str("method", c.Method()).
					Msg("panic recovered")
				err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "internal_error",
					"message": "Something went wrong",
				})
			}
		}()
		return c.Next()
	}
}
