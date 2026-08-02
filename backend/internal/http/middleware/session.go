package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/response"
)

// Authenticated requires a valid access token and stores the authenticated
// subject on the request. The access token is read from the dp_session cookie
// (M4-DEC-004) or the Authorization bearer header. Authorization of a specific
// resource is a separate server-side decision made by the owning bounded
// context.
func Authenticated(tokens *auth.TokenManager, sessionCookieName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := accessTokenFrom(c, sessionCookieName)
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

// accessTokenFrom extracts the access token from the session cookie or the
// Authorization bearer header. Cookie-first matches the session-transport
// decision in the M6 contract.
func accessTokenFrom(c *fiber.Ctx, sessionCookieName string) (string, bool) {
	if cookie := c.Cookies(sessionCookieName); cookie != "" {
		return cookie, true
	}

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

// SessionCookies manages the access and refresh session cookies (M4-DEC-004).
// Cookies are HttpOnly, SameSite=Lax, and Secure outside development.
type SessionCookies struct {
	AccessCookieName  string
	RefreshCookieName string
	Secure            bool
}

// SetAccess stores the access token cookie.
func (s SessionCookies) SetAccess(c *fiber.Ctx, token string, expires time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     s.AccessCookieName,
		Value:    token,
		Path:     "/",
		Expires:  expires,
		HTTPOnly: true,
		Secure:   s.Secure,
		SameSite: fiber.CookieSameSiteLaxMode,
	})
}

// SetRefresh stores the refresh token cookie.
func (s SessionCookies) SetRefresh(c *fiber.Ctx, token string, expires time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     s.RefreshCookieName,
		Value:    token,
		Path:     "/api/v1/accounts",
		Expires:  expires,
		HTTPOnly: true,
		Secure:   s.Secure,
		SameSite: fiber.CookieSameSiteLaxMode,
	})
}

// ClearAccess expires the access cookie.
func (s SessionCookies) ClearAccess(c *fiber.Ctx) {
	c.ClearCookie(s.AccessCookieName)
}

// ClearRefresh expires the refresh cookie.
func (s SessionCookies) ClearRefresh(c *fiber.Ctx) {
	c.ClearCookie(s.RefreshCookieName)
}
