// Package store implements the identity.Store port with the generated SQLC
// queries against PostgreSQL. It is the only package that touches the identity
// SQL for the M9 DP-FEAT-001 slice.
package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/gen/sqlc"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/identity"
)

// Postgres implements identity.Store over a pgx connection.
type Postgres struct {
	db *sqlc.Queries
}

// New builds the identity store adapter. The pool must already be open.
func New(conn sqlc.DBTX) *Postgres {
	return &Postgres{db: sqlc.New(conn)}
}

func (p *Postgres) GetAccountByID(ctx context.Context, id string) (*identity.Account, error) {
	row, err := p.db.GetAccountByID(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toAccount(row), nil
}

func (p *Postgres) GetAccountByEmail(ctx context.Context, email string) (*identity.Account, error) {
	row, err := p.db.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, mapErr(err)
	}
	return toAccount(row), nil
}

func (p *Postgres) CreateAccount(ctx context.Context, email, passwordHash string, status identity.AccountStatus, timezone string) (*identity.Account, error) {
	row, err := p.db.CreateAccount(ctx, sqlc.CreateAccountParams{
		Email:        email,
		PasswordHash: passwordHash,
		Status:       string(status),
		Timezone:     timezone,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toAccount(row), nil
}

func (p *Postgres) UpdateAccountStatus(ctx context.Context, id string, status identity.AccountStatus) (*identity.Account, error) {
	row, err := p.db.UpdateAccountStatus(ctx, sqlc.UpdateAccountStatusParams{ID: id, Status: string(status)})
	if err != nil {
		return nil, mapErr(err)
	}
	return toAccount(row), nil
}

func (p *Postgres) GetUserProfileByAccountID(ctx context.Context, accountID string) (*identity.UserProfile, error) {
	row, err := p.db.GetUserProfileByAccountID(ctx, accountID)
	if err != nil {
		return nil, mapErr(err)
	}
	return toUserProfile(row), nil
}

func (p *Postgres) CreateUserProfile(ctx context.Context, accountID, displayName string, timezone *string) (*identity.UserProfile, error) {
	row, err := p.db.CreateUserProfile(ctx, sqlc.CreateUserProfileParams{
		AccountID:   accountID,
		DisplayName: displayName,
		Status:      string(identity.ProfileCreated),
		Timezone:    timezone,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toUserProfile(row), nil
}

func (p *Postgres) UpdateUserProfile(ctx context.Context, id string, displayName, timezone *string) (*identity.UserProfile, error) {
	var nameValue string
	if displayName != nil {
		nameValue = *displayName
	}
	row, err := p.db.UpdateUserProfile(ctx, sqlc.UpdateUserProfileParams{
		ID:          id,
		DisplayName: nameValue,
		Timezone:    timezone,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toUserProfile(row), nil
}

func (p *Postgres) GetActivePreferenceSetForProfile(ctx context.Context, profileID string) (*identity.PreferenceSet, error) {
	row, err := p.db.GetActivePreferenceSetForProfile(ctx, profileID)
	if err != nil {
		return nil, mapErr(err)
	}
	return toPreferenceSet(row), nil
}

func (p *Postgres) CreatePreferenceSet(ctx context.Context, profileID string, preferences []byte, validFrom *string) (*identity.PreferenceSet, error) {
	validFromValue := ""
	if validFrom != nil {
		validFromValue = *validFrom
	}
	row, err := p.db.CreatePreferenceSet(ctx, sqlc.CreatePreferenceSetParams{
		UserProfileID: profileID,
		Preferences:   preferences,
		ValidFrom:     validFromValue,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toPreferenceSet(row), nil
}

func (p *Postgres) CreateSession(ctx context.Context, accountID, refreshSecretHash, familyID string, expiresAt time.Time) (*identity.Session, error) {
	row, err := p.db.CreateAuthSession(ctx, sqlc.CreateAuthSessionParams{
		AccountID:         accountID,
		RefreshSecretHash: refreshSecretHash,
		FamilyID:          familyID,
		UserAgentHash:     nil,
		IpHash:            nil,
		ExpiresAt:         expiresAt,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toSession(row), nil
}

func (p *Postgres) GetSessionBySecretHash(ctx context.Context, refreshSecretHash string) (*identity.Session, error) {
	row, err := p.db.GetAuthSessionBySecretHash(ctx, refreshSecretHash)
	if err != nil {
		return nil, mapErr(err)
	}
	return toSession(row), nil
}

func (p *Postgres) RevokeSession(ctx context.Context, id string) (*identity.Session, error) {
	row, err := p.db.RevokeAuthSession(ctx, id)
	if err != nil {
		return nil, mapErr(err)
	}
	return toSession(row), nil
}

func (p *Postgres) RevokeSessionFamily(ctx context.Context, familyID string) error {
	if err := p.db.RevokeSessionFamily(ctx, familyID); err != nil {
		return mapErr(err)
	}
	return nil
}

func (p *Postgres) MarkSessionReplacedBy(ctx context.Context, id, replacementID string) (*identity.Session, error) {
	row, err := p.db.MarkAuthSessionReplacedBy(ctx, sqlc.MarkAuthSessionReplacedByParams{
		ID:         id,
		ReplacedBy: &replacementID,
	})
	if err != nil {
		return nil, mapErr(err)
	}
	return toSession(row), nil
}

func toAccount(r sqlc.Account) *identity.Account {
	return &identity.Account{
		ID:              r.ID,
		Email:           r.Email,
		PasswordHash:    r.PasswordHash,
		Status:          identity.AccountStatus(r.Status),
		Timezone:        r.Timezone,
		EmailVerifiedAt: r.EmailVerifiedAt,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}
}

func toUserProfile(r sqlc.UserProfile) *identity.UserProfile {
	return &identity.UserProfile{
		ID:          r.ID,
		AccountID:   r.AccountID,
		DisplayName: r.DisplayName,
		Status:      identity.ProfileStatus(r.Status),
		Timezone:    r.Timezone,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func toPreferenceSet(r sqlc.PreferenceSet) *identity.PreferenceSet {
	return &identity.PreferenceSet{
		ID:            r.ID,
		UserProfileID: r.UserProfileID,
		Status:        r.Status,
		Preferences:   r.Preferences,
		ValidFrom:     r.ValidFrom,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}
}

func toSession(r sqlc.AuthSession) *identity.Session {
	return &identity.Session{
		ID:                r.ID,
		AccountID:         r.AccountID,
		RefreshSecretHash: r.RefreshSecretHash,
		FamilyID:          r.FamilyID,
		ExpiresAt:         r.ExpiresAt,
		RevokedAt:         r.RevokedAt,
		ReplacedBy:        r.ReplacedBy,
	}
}
func mapErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return identity.ErrNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// 23505 = unique_violation. Registration duplicates the unique email.
		if pgErr.Code == "23505" {
			return identity.ErrEmailInUse
		}
	}
	return err
}
