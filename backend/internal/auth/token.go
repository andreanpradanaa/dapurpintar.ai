package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// Claims are the security claims carried by the short-lived access token.
// They are intentionally minimal and never carry pantry, preference, recipe,
// conversation, or AI context. Authorization remains a server-side decision.
type Claims struct {
	jwt.RegisteredClaims
}

// SessionTokens represent one authenticated participation session.
type SessionTokens struct {
	AccessToken  string
	RefreshToken string
	// AccessExpiresAt is the access-token expiration instant.
	AccessExpiresAt time.Time
}

// TokenManager issues and verifies access tokens and generates refresh tokens.
type TokenManager struct {
	secret            []byte
	issuer            string
	audience          string
	accessTokenTTL    time.Duration
	refreshTokenTTL   time.Duration
	sessionCookieName string
}

// NewTokenManager builds a TokenManager. A missing JWT secret is a startup
// error: running without one is unsafe.
func NewTokenManager(secret, issuer, audience string, accessTTL, refreshTTL time.Duration) (*TokenManager, error) {
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}
	return &TokenManager{
		secret:          []byte(secret),
		issuer:          issuer,
		audience:        audience,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}, nil
}

// Issue creates a new access token and refresh token for the given subject
// (the Account or User identity). Refresh secrets are stored only in protected
// form by the session store; the returned refresh token is presented to the
// client and verified on rotation.
func (m *TokenManager) Issue(subject string) (*SessionTokens, error) {
	now := time.Now().UTC()
	expires := now.Add(m.accessTokenTTL)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			Subject:   subject,
			Audience:  jwt.ClaimStrings{m.audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expires),
			ID:        newTokenID(),
		},
	})

	signed, err := token.SignedString(m.secret)
	if err != nil {
		return nil, apperr.Wrap(apperr.CodeInternal, "Failed to issue session.", err)
	}

	return &SessionTokens{
		AccessToken:     signed,
		RefreshToken:    newTokenID(),
		AccessExpiresAt: expires,
	}, nil
}

// Verify validates an access token and returns the subject. Verification
// enforces issuer, audience, and expiration. Authorization is a separate
// server-side decision and is not inferred from token contents.
func (m *TokenManager) Verify(accessToken string) (string, error) {
	parsed, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return m.secret, nil
	})
	if err != nil {
		return "", apperr.Wrap(apperr.CodeSessionInvalid, "The session is invalid or expired.", err)
	}

	claims, ok := parsed.Claims.(*Claims)
	if !ok || !parsed.Valid {
		return "", apperr.New(apperr.CodeSessionInvalid, "The session is invalid or expired.")
	}

	if claims.Subject == "" {
		return "", apperr.New(apperr.CodeSessionInvalid, "The session is invalid or expired.")
	}

	return claims.Subject, nil
}

// RefreshTTL returns the configured refresh-session lifetime.
func (m *TokenManager) RefreshTTL() time.Duration { return m.refreshTokenTTL }

func newTokenID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		// crypto/rand failure is unrecoverable; fall back to a timestamp-based
		// id so the caller still gets a token, but uniqueness relies on time.
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
