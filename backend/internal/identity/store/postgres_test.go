package store

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/identity"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/database"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
)

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; skipping Postgres integration test")
	}
	ctx := context.Background()
	pool, err := database.New(ctx, url, logger.New("test"))
	if err != nil {
		t.Fatalf("database.New: %v", err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func testStore(t *testing.T) *Postgres {
	t.Helper()
	return New(testPool(t))
}

func TestPostgres_AccountLifecycle(t *testing.T) {
	st := testStore(t)
	ctx := context.Background()

	email := fmt.Sprintf("integration-%d", time.Now().UnixNano()) + "@example.com"
	created, err := st.CreateAccount(ctx, email, "hash-abc", identity.AccountActive, "Asia/Jakarta")
	if err != nil {
		t.Fatalf("CreateAccount: %v", err)
	}
	if created.ID == "" || created.Email != email {
		t.Fatalf("unexpected created account: %+v", created)
	}

	byID, err := st.GetAccountByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetAccountByID: %v", err)
	}
	if byID.Email != email {
		t.Fatalf("byID email = %q, want %q", byID.Email, email)
	}

	byEmail, err := st.GetAccountByEmail(ctx, email)
	if err != nil {
		t.Fatalf("GetAccountByEmail: %v", err)
	}
	if byEmail.ID != created.ID {
		t.Fatalf("byEmail id = %q, want %q", byEmail.ID, created.ID)
	}

	updated, err := st.UpdateAccountStatus(ctx, created.ID, identity.AccountRestricted)
	if err != nil {
		t.Fatalf("UpdateAccountStatus: %v", err)
	}
	if updated.Status != identity.AccountRestricted {
		t.Fatalf("updated status = %q, want restricted", updated.Status)
	}

	dup, err := st.CreateAccount(ctx, email, "hash-xyz", identity.AccountActive, "UTC")
	if err != identity.ErrEmailInUse {
		t.Fatalf("duplicate CreateAccount err = %v, want ErrEmailInUse", err)
	}
	_ = dup
}

func TestPostgres_ProfileAndPreferences(t *testing.T) {
	st := testStore(t)
	ctx := context.Background()

	account, err := st.CreateAccount(ctx, fmt.Sprintf("profile-%d@example.com", time.Now().UnixNano()), "hash", identity.AccountActive, "UTC")
	if err != nil {
		t.Fatalf("CreateAccount: %v", err)
	}

	profile, err := st.CreateUserProfile(ctx, account.ID, "Budi", nil)
	if err != nil {
		t.Fatalf("CreateUserProfile: %v", err)
	}

	got, err := st.GetUserProfileByAccountID(ctx, account.ID)
	if err != nil {
		t.Fatalf("GetUserProfileByAccountID: %v", err)
	}
	if got.ID != profile.ID {
		t.Fatalf("profile id = %q, want %q", got.ID, profile.ID)
	}

	tz := "Asia/Makassar"
	updated, err := st.UpdateUserProfile(ctx, profile.ID, nil, &tz)
	if err != nil {
		t.Fatalf("UpdateUserProfile: %v", err)
	}
	if updated.Timezone == nil || *updated.Timezone != tz {
		t.Fatalf("updated timezone = %v, want %q", updated.Timezone, tz)
	}

	prefs, err := st.CreatePreferenceSet(ctx, profile.ID, []byte(`{"diet":"pescatarian"}`), nil)
	if err != nil {
		t.Fatalf("CreatePreferenceSet: %v", err)
	}

	active, err := st.GetActivePreferenceSetForProfile(ctx, profile.ID)
	if err != nil {
		t.Fatalf("GetActivePreferenceSetForProfile: %v", err)
	}
	if active.ID != prefs.ID {
		t.Fatalf("active id = %q, want %q", active.ID, prefs.ID)
	}

	if _, err := st.GetActivePreferenceSetForProfile(ctx, "00000000-0000-0000-0000-000000000000"); err != identity.ErrNotFound {
		t.Fatalf("missing prefs err = %v, want ErrNotFound", err)
	}
}

func TestPostgres_SessionLifecycle(t *testing.T) {
	st := testStore(t)
	ctx := context.Background()

	sessionFamilyID := fmt.Sprintf("11111111-1111-1111-1111-%012d", time.Now().UnixNano()%1e12)
	secret1 := fmt.Sprintf("secret-hash-1-%d", time.Now().UnixNano())
	secret2 := fmt.Sprintf("secret-hash-2-%d", time.Now().UnixNano())

	account, err := st.CreateAccount(ctx, fmt.Sprintf("session-%d@example.com", time.Now().UnixNano()), "hash", identity.AccountActive, "UTC")
	if err != nil {
		t.Fatalf("CreateAccount: %v", err)
	}

	created, err := st.CreateSession(ctx, account.ID, secret1, sessionFamilyID, time.Now().Add(24*time.Hour))
	if err != nil {
		t.Fatalf("CreateSession: %v", err)
	}

	got, err := st.GetSessionBySecretHash(ctx, secret1)
	if err != nil {
		t.Fatalf("GetSessionBySecretHash: %v", err)
	}
	if got.ID != created.ID || got.FamilyID != sessionFamilyID {
		t.Fatalf("unexpected session: %+v", got)
	}

	replacement, err := st.CreateSession(ctx, account.ID, secret2, sessionFamilyID, time.Now().Add(24*time.Hour))
	if err != nil {
		t.Fatalf("CreateSession(replacement): %v", err)
	}

	marked, err := st.MarkSessionReplacedBy(ctx, created.ID, replacement.ID)
	if err != nil {
		t.Fatalf("MarkSessionReplacedBy: %v", err)
	}
	if marked.ReplacedBy == nil || *marked.ReplacedBy != replacement.ID {
		t.Fatalf("marked.ReplacedBy = %v, want %q", marked.ReplacedBy, replacement.ID)
	}

	revoked, err := st.RevokeSession(ctx, replacement.ID)
	if err != nil {
		t.Fatalf("RevokeSession: %v", err)
	}
	if revoked.RevokedAt == nil {
		t.Fatalf("RevokeSession did not set revoked_at")
	}

	if err := st.RevokeSessionFamily(ctx, sessionFamilyID); err != nil {
		t.Fatalf("RevokeSessionFamily: %v", err)
	}
}

// TestServiceEndToEnd drives the full DP-FEAT-001 use cases over the adapter.
func TestServiceEndToEnd(t *testing.T) {
	st := testStore(t)
	ctx := context.Background()

	tokens, err := auth.NewTokenManager("test-secret", "dapurpintar-test", "dapurpintar-test", 15*time.Minute, 24*time.Hour)
	if err != nil {
		t.Fatalf("NewTokenManager: %v", err)
	}
	svc := identity.NewService(st, tokens)

	email := fmt.Sprintf("e2e-%d", time.Now().UnixNano()) + "@example.com"
	account, profile, err := svc.Register(ctx, identity.RegisterInput{
		Email:       email,
		Password:    "correct-horse-battery",
		DisplayName: "Siti",
		Timezone:    "Asia/Jakarta",
	})
	if err != nil {
		t.Fatalf("Register: %v", err)
	}
	if profile == nil {
		t.Fatal("Register returned nil profile")
	}

	session, err := svc.Login(ctx, identity.LoginInput{Email: email, Password: "correct-horse-battery"})
	if err != nil {
		t.Fatalf("Login: %v", err)
	}
	if session.RefreshToken == "" {
		t.Fatal("Login returned empty refresh token")
	}

	rotated, err := svc.Refresh(ctx, session.RefreshToken)
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if rotated.Account.ID != account.ID {
		t.Fatalf("rotated account = %q, want %q", rotated.Account.ID, account.ID)
	}

	// Rotated session is revoked; reusing it is detected as theft.
	if _, err := svc.Refresh(ctx, session.RefreshToken); err == nil {
		t.Fatal("expected refresh reuse to fail")
	}

	if err := svc.Logout(ctx, rotated.RefreshToken); err != nil {
		t.Fatalf("Logout: %v", err)
	}
}
