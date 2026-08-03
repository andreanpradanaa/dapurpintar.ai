package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

// Recover catches panics and returns a uniform 500 response.
func Recover(log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic recovered",
					"panic", r,
					"path", c.Path(),
					"method", c.Method(),
				)
				err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "internal_error",
					"message": "Something went wrong",
				})
			}
		}()
		return c.Next()
	}
}
