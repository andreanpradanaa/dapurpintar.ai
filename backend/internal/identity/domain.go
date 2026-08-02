// Package identity implements the Identity and Access bounded context
// (accounts, authentication sessions) and the User Context and Preferences
// bounded context (user profiles, preference sets) for the M9 MVP feature
// slice DP-FEAT-001. It follows Clean Architecture: the domain types and use
// cases live here, persistence is behind the Store port, and handlers in
// internal/http translate the M6 API contract.
package identity

import (
	"errors"
	"time"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// AccountStatus is the account lifecycle from M5-001.
type AccountStatus string

const (
	AccountPending    AccountStatus = "pending"
	AccountActive     AccountStatus = "active"
	AccountRestricted AccountStatus = "restricted"
	AccountClosed     AccountStatus = "closed"
)

// ProfileStatus is the user profile lifecycle from M5-001.
type ProfileStatus string

const (
	ProfileCreated    ProfileStatus = "created"
	ProfileIncomplete ProfileStatus = "incomplete"
	ProfileReady      ProfileStatus = "ready"
	ProfileUpdated    ProfileStatus = "updated"
)

// Account is the Identity and Access aggregate root.
type Account struct {
	ID              string
	Email           string
	PasswordHash    string
	Status          AccountStatus
	Timezone        string
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// UserProfile is the User Context and Preferences aggregate root.
type UserProfile struct {
	ID          string
	AccountID   string
	DisplayName string
	Status      ProfileStatus
	Timezone    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// PreferenceSet is a versioned declaration of preferences and constraints.
type PreferenceSet struct {
	ID            string
	UserProfileID string
	Status        string
	Preferences   []byte
	ValidFrom     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Session is a durable refresh-session record (M4-DEC-003). The refresh secret
// itself is never stored; only its hash and the metadata required for
// revocation, expiry, rotation, and reuse detection are persisted.
type Session struct {
	ID                string
	AccountID         string
	RefreshSecretHash string
	FamilyID          string
	ExpiresAt         time.Time
	RevokedAt         *time.Time
	ReplacedBy        *string
}

// ErrEmailInUse wraps the M6 ACCOUNT_EMAIL_IN_USE code for registration
// conflicts. It is produced by the Store when an active account with the same
// email already exists.
var ErrEmailInUse = apperr.New(apperr.CodeEmailInUse, "An account with this email already exists.")

// ErrNotFound marks a missing aggregate in the Store. The use case maps it to
// the appropriate M6 code for the operation.
var ErrNotFound = errors.New("identity resource not found")
