package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
)

// Authenticated requires a valid access token and stores the authenticated
// subject on the request. Authorization of a specific resource is a separate
// server-side decision made by the owning bounded context.
func Authenticated(tokens *auth.TokenManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := accessTokenFrom(c)
		if !ok {
			return response.Error(c, apperr.New(apperr.CodeSessionInvalid, "A valid session is required."))
		}

		subject, err := tokens.Verify(token)
		if err != nil {
			return response.Error(c, err)
		}

		c.Locals("auth_subject", subject)
		return c.Next()
	}
}

// Subject returns the authenticated subject stored by the Authenticated
// middleware, or the empty string when absent.
func Subject(c *fiber.Ctx) string {
	if v, ok := c.Locals("auth_subject").(string); ok {
		return v
	}
	return ""
}

// accessTokenFrom extracts the access token from the Authorization bearer header.
func accessTokenFrom(c *fiber.Ctx) (string, bool) {
	header := c.Get("Authorization")
	if header == "" {
		return "", false
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return "", false
	}

	token := strings.TrimSpace(strings.TrimPrefix(header, prefix))
	if token == "" {
		return "", false
	}
	return token, true
}
