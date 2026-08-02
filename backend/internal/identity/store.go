package identity

import (
	"context"
	"time"
)

// Store is the persistence port for the Identity and Access and User Context
// and Preferences bounded contexts. Handlers and use cases depend on this
// interface, never on the SQLC adapter directly.
type Store interface {
	// Account queries.
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*Account, error)
	CreateAccount(ctx context.Context, email, passwordHash string, status AccountStatus, timezone string) (*Account, error)
	UpdateAccountStatus(ctx context.Context, id string, status AccountStatus) (*Account, error)

	// User profile queries.
	GetUserProfileByAccountID(ctx context.Context, accountID string) (*UserProfile, error)
	CreateUserProfile(ctx context.Context, accountID, displayName string, timezone *string) (*UserProfile, error)
	UpdateUserProfile(ctx context.Context, id string, displayName, timezone *string) (*UserProfile, error)

	// Preference queries.
	GetActivePreferenceSetForProfile(ctx context.Context, profileID string) (*PreferenceSet, error)
	CreatePreferenceSet(ctx context.Context, profileID string, preferences []byte, validFrom *string) (*PreferenceSet, error)

	// Session queries (M4-DEC-003 durable session authority).
	CreateSession(ctx context.Context, accountID, refreshSecretHash, familyID string, expiresAt time.Time) (*Session, error)
	GetSessionBySecretHash(ctx context.Context, refreshSecretHash string) (*Session, error)
	RevokeSession(ctx context.Context, id string) (*Session, error)
	RevokeSessionFamily(ctx context.Context, familyID string) error
	MarkSessionReplacedBy(ctx context.Context, id, replacementID string) (*Session, error)
}
