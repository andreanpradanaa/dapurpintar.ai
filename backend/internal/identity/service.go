package identity

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// Service implements the DP-FEAT-001 use cases: registration, login, logout,
// refresh, current account, profile, and preferences.
type Service struct {
	store  Store
	tokens *auth.TokenManager
}

// NewService builds the identity use-case service.
func NewService(store Store, tokens *auth.TokenManager) *Service {
	return &Service{store: store, tokens: tokens}
}

// SessionResult is the outcome of a successful login or refresh.
type SessionResult struct {
	Account         *Account
	Profile         *UserProfile
	AccessToken     string
	RefreshToken    string
	AccessExpiresAt time.Time
}

// RegisterInput is the validated registration intent.
type RegisterInput struct {
	Email       string
	Password    string
	DisplayName string
	Timezone    string
}

// Register creates an account, its profile, and an empty preference set. Email
// verification stays optional for local development (M4-DEC-001); account
// status is active on success.
func (s *Service) Register(ctx context.Context, in RegisterInput) (*Account, *UserProfile, error) {
	email, err := normalizeEmail(in.Email)
	if err != nil {
		return nil, nil, err
	}
	if err := validatePassword(in.Password); err != nil {
		return nil, nil, err
	}

	displayName := strings.TrimSpace(in.DisplayName)
	if displayName == "" {
		displayName = emailLocalPart(email)
	}
	if utf8.RuneCountInString(displayName) > 120 {
		return nil, nil, apperr.New(apperr.CodeFieldInvalid, "Display name is too long.").
			WithDetails(apperr.Detail{Field: "display_name", Code: string(apperr.CodeFieldInvalid), Message: "Display name must be 120 characters or fewer."})
	}
	timezone := in.Timezone
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	// The account is created as active because email verification is optional
	// in the MVP (M4-DEC-001); it becomes required before public launch.
	passwordHash, err := auth.HashPassword(in.Password, auth.DefaultPasswordConfig)
	if err != nil {
		return nil, nil, apperr.Internal(err)
	}

	account, err := s.store.CreateAccount(ctx, email, passwordHash, AccountActive, timezone)
	if err != nil {
		if errors.Is(err, ErrEmailInUse) {
			return nil, nil, ErrEmailInUse
		}
		return nil, nil, apperr.Internal(err)
	}

	profile, err := s.store.CreateUserProfile(ctx, account.ID, displayName, nil)
	if err != nil {
		return nil, nil, apperr.Internal(err)
	}
	if _, err := s.store.CreatePreferenceSet(ctx, profile.ID, []byte("{}"), nil); err != nil {
		return nil, nil, apperr.Internal(err)
	}

	return account, profile, nil
}

// LoginInput is the validated login intent.
type LoginInput struct {
	Email    string
	Password string
}

// Login verifies credentials and opens a refresh-session lineage. The returned
// refresh token is presented to the client and stored only in hashed form by
// the Store. Account status is enforced: restricted or closed accounts cannot
// authenticate.
func (s *Service) Login(ctx context.Context, in LoginInput) (*SessionResult, error) {
	email, err := normalizeEmail(in.Email)
	if err != nil {
		return nil, err
	}

	account, err := s.store.GetAccountByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			// Generic failure: do not reveal whether the email exists (M6
			// AUTH_CREDENTIALS_INVALID, no enumeration).
			return nil, apperr.New(apperr.CodeCredentialsInvalid, "The email or password is incorrect.")
		}
		return nil, apperr.Internal(err)
	}
	if err := validateAccountActive(account); err != nil {
		return nil, err
	}

	ok, err := auth.VerifyPassword(in.Password, account.PasswordHash)
	if err != nil || !ok {
		return nil, apperr.New(apperr.CodeCredentialsInvalid, "The email or password is incorrect.")
	}

	return s.openSession(ctx, account)
}

// Logout revokes the current refresh session identified by the refresh token
// hash presented on the request.
func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	hash := hashRefreshSecret(refreshToken)
	session, err := s.store.GetSessionBySecretHash(ctx, hash)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}
		return apperr.Internal(err)
	}
	if _, err := s.store.RevokeSession(ctx, session.ID); err != nil {
		return apperr.Internal(err)
	}
	return nil
}

// Refresh rotates a valid refresh session (M4-DEC-003). A revoked session that
// is reused is treated as suspected theft and revokes the whole lineage.
func (s *Service) Refresh(ctx context.Context, refreshToken string) (*SessionResult, error) {
	if refreshToken == "" {
		return nil, apperr.New(apperr.CodeSessionInvalid, "A valid refresh session is required.")
	}

	hash := hashRefreshSecret(refreshToken)
	session, err := s.store.GetSessionBySecretHash(ctx, hash)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, apperr.New(apperr.CodeSessionInvalid, "The refresh session is invalid.")
		}
		return nil, apperr.Internal(err)
	}

	if session.RevokedAt != nil {
		// Reuse of a rotated (revoked) credential is suspected theft: revoke
		// the entire lineage and refuse.
		if err := s.store.RevokeSessionFamily(ctx, session.FamilyID); err != nil {
			return nil, apperr.Internal(err)
		}
		return nil, apperr.New(apperr.CodeRefreshReused, "The refresh session has already been used.")
	}

	if time.Now().UTC().After(session.ExpiresAt) {
		return nil, apperr.New(apperr.CodeSessionExpired, "The refresh session has expired.")
	}

	account, err := s.store.GetAccountByID(ctx, session.AccountID)
	if err != nil {
		return nil, apperr.Internal(err)
	}
	if err := validateAccountActive(account); err != nil {
		return nil, err
	}

	return s.rotateSession(ctx, account, session)
}

// GetAccountByID returns an account for a subject.
func (s *Service) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	account, err := s.store.GetAccountByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, apperr.New(apperr.CodeSessionInvalid, "The account does not exist.")
		}
		return nil, apperr.Internal(err)
	}
	return account, nil
}

// GetProfileByAccountID returns the profile for an account.
func (s *Service) GetProfileByAccountID(ctx context.Context, accountID string) (*UserProfile, error) {
	profile, err := s.store.GetUserProfileByAccountID(ctx, accountID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, apperr.New(apperr.CodeIDInvalid, "No profile exists for this account.")
		}
		return nil, apperr.Internal(err)
	}
	return profile, nil
}

// ActivePreferences returns the current preference set for a profile.
func (s *Service) ActivePreferences(ctx context.Context, profileID string) (*PreferenceSet, error) {
	prefs, err := s.store.GetActivePreferenceSetForProfile(ctx, profileID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return &PreferenceSet{Preferences: []byte("{}")}, nil
		}
		return nil, apperr.Internal(err)
	}
	return prefs, nil
}

// UpdateProfileInput is the validated profile update intent.
type UpdateProfileInput struct {
	DisplayName *string
	Timezone    *string
}

// UpdateProfile applies a profile update.
func (s *Service) UpdateProfile(ctx context.Context, accountID string, in UpdateProfileInput) (*UserProfile, error) {
	profile, err := s.GetProfileByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	if in.DisplayName != nil {
		name := strings.TrimSpace(*in.DisplayName)
		if name == "" {
			return nil, apperr.New(apperr.CodeFieldInvalid, "Display name cannot be empty.").
				WithDetails(apperr.Detail{Field: "display_name", Code: string(apperr.CodeFieldInvalid), Message: "Display name must not be empty."})
		}
		if utf8.RuneCountInString(name) > 120 {
			return nil, apperr.New(apperr.CodeFieldInvalid, "Display name is too long.").
				WithDetails(apperr.Detail{Field: "display_name", Code: string(apperr.CodeFieldInvalid), Message: "Display name must be 120 characters or fewer."})
		}
	}

	var displayName, timezone *string
	if in.DisplayName != nil {
		v := strings.TrimSpace(*in.DisplayName)
		displayName = &v
	}
	if in.Timezone != nil {
		v := *in.Timezone
		if strings.TrimSpace(v) == "" {
			return nil, apperr.New(apperr.CodeFieldInvalid, "Timezone cannot be empty.").
				WithDetails(apperr.Detail{Field: "timezone", Code: string(apperr.CodeFieldInvalid), Message: "Timezone must not be empty."})
		}
		timezone = &v
	}

	updated, err := s.store.UpdateUserProfile(ctx, profile.ID, displayName, timezone)
	if err != nil {
		return nil, apperr.Internal(err)
	}
	return updated, nil
}

// UpdatePreferencesInput is the validated preference update intent.
type UpdatePreferencesInput struct {
	Preferences []byte
	ValidFrom   *string
}

// UpdatePreferences retires the current preference set and creates a new active
// one, preserving the versioned history.
func (s *Service) UpdatePreferences(ctx context.Context, accountID string, in UpdatePreferencesInput) (*PreferenceSet, error) {
	profile, err := s.GetProfileByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	preferences := in.Preferences
	if len(preferences) == 0 {
		preferences = []byte("{}")
	}

	created, err := s.store.CreatePreferenceSet(ctx, profile.ID, preferences, in.ValidFrom)
	if err != nil {
		return nil, apperr.Internal(err)
	}
	return created, nil
}

// openSession creates a new refresh session and issues tokens.
func (s *Service) openSession(ctx context.Context, account *Account) (*SessionResult, error) {
	tokens, err := s.tokens.Issue(account.ID)
	if err != nil {
		return nil, err
	}

	familyID, err := newSessionFamilyID()
	if err != nil {
		return nil, apperr.Internal(err)
	}
	if _, err := s.store.CreateSession(ctx, account.ID, hashRefreshSecret(tokens.RefreshToken), familyID, time.Now().UTC().Add(s.tokens.RefreshTTL())); err != nil {
		return nil, apperr.Internal(err)
	}

	profile, _ := s.store.GetUserProfileByAccountID(ctx, account.ID)
	return &SessionResult{
		Account:         account,
		Profile:         profile,
		AccessToken:     tokens.AccessToken,
		RefreshToken:    tokens.RefreshToken,
		AccessExpiresAt: tokens.AccessExpiresAt,
	}, nil
}

// rotateSession rotates a valid refresh session: the old session is marked
// replaced, a new session is created in the same lineage, and fresh tokens are
// issued.
func (s *Service) rotateSession(ctx context.Context, account *Account, current *Session) (*SessionResult, error) {
	tokens, err := s.tokens.Issue(account.ID)
	if err != nil {
		return nil, err
	}

	replacement, err := s.store.CreateSession(ctx, account.ID, hashRefreshSecret(tokens.RefreshToken), current.FamilyID, time.Now().UTC().Add(s.tokens.RefreshTTL()))
	if err != nil {
		return nil, apperr.Internal(err)
	}
	if _, err := s.store.MarkSessionReplacedBy(ctx, current.ID, replacement.ID); err != nil {
		return nil, apperr.Internal(err)
	}

	profile, _ := s.store.GetUserProfileByAccountID(ctx, account.ID)
	return &SessionResult{
		Account:         account,
		Profile:         profile,
		AccessToken:     tokens.AccessToken,
		RefreshToken:    tokens.RefreshToken,
		AccessExpiresAt: tokens.AccessExpiresAt,
	}, nil
}

func validateAccountActive(account *Account) error {
	switch account.Status {
	case AccountActive:
		return nil
	case AccountRestricted:
		return apperr.New(apperr.CodeAccountRestricted, "This account is restricted.")
	case AccountClosed:
		return apperr.New(apperr.CodeAccountNotActive, "This account is closed.")
	default:
		return apperr.New(apperr.CodeAccountNotActive, "This account is not active.")
	}
}

func normalizeEmail(email string) (string, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" {
		return "", apperr.New(apperr.CodeEmailInvalid, "An email address is required.").
			WithDetails(apperr.Detail{Field: "email", Code: string(apperr.CodeEmailInvalid), Message: "Email must not be empty."})
	}
	if len(email) > 254 {
		return "", apperr.New(apperr.CodeEmailInvalid, "The email address is invalid.").
			WithDetails(apperr.Detail{Field: "email", Code: string(apperr.CodeEmailInvalid), Message: "Email must be 254 characters or fewer."})
	}
	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Address != email {
		return "", apperr.New(apperr.CodeEmailInvalid, "The email address is invalid.").
			WithDetails(apperr.Detail{Field: "email", Code: string(apperr.CodeEmailInvalid), Message: "Email must be a valid address."})
	}
	return email, nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return apperr.New(apperr.CodePasswordWeak, "The password is too weak.").
			WithDetails(apperr.Detail{Field: "password", Code: string(apperr.CodePasswordWeak), Message: "Password must be at least 8 characters."})
	}
	if len(password) > 128 {
		return apperr.New(apperr.CodePasswordWeak, "The password is too long.").
			WithDetails(apperr.Detail{Field: "password", Code: string(apperr.CodePasswordWeak), Message: "Password must be 128 characters or fewer."})
	}
	return nil
}

func emailLocalPart(email string) string {
	idx := strings.IndexByte(email, '@')
	if idx < 0 {
		return email
	}
	return email[:idx]
}

func hashRefreshSecret(secret string) string {
	sum := sha256.Sum256([]byte(secret))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func newSessionFamilyID() (string, error) {
	b := make([]byte, 16)
	if _, err := cryptoRandRead(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

// cryptoRandRead is a thin wrapper so tests can inject a deterministic source.
var cryptoRandRead = rand.Read
