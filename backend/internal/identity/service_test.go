package identity

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/auth"
)

type stubStore struct {
	Store
	accounts              map[string]*Account
	profiles              map[string]*UserProfile
	sessionsByHash        map[string]*Session
	createAccount         func(ctx context.Context, email, hash string, status AccountStatus, tz string) (*Account, error)
	getActivePreferenceFn func(ctx context.Context, profileID string) (*PreferenceSet, error)
}

func newStubStore() *stubStore {
	return &stubStore{
		accounts:       make(map[string]*Account),
		profiles:       make(map[string]*UserProfile),
		sessionsByHash: make(map[string]*Session),
	}
}

func (s *stubStore) GetAccountByEmail(ctx context.Context, email string) (*Account, error) {
	a, ok := s.accounts[email]
	if !ok {
		return nil, ErrNotFound
	}
	return a, nil
}

func (s *stubStore) CreateAccount(ctx context.Context, email, hash string, status AccountStatus, tz string) (*Account, error) {
	if s.createAccount != nil {
		return s.createAccount(ctx, email, hash, status, tz)
	}
	a := &Account{ID: "acc-1", Email: email, PasswordHash: hash, Status: status, Timezone: tz, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	s.accounts[email] = a
	return a, nil
}

func (s *stubStore) CreateUserProfile(ctx context.Context, accountID, displayName string, timezone *string) (*UserProfile, error) {
	p := &UserProfile{ID: "prof-1", AccountID: accountID, DisplayName: displayName, Status: ProfileCreated, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	s.profiles[accountID] = p
	return p, nil
}

func (s *stubStore) CreatePreferenceSet(ctx context.Context, profileID string, preferences []byte, validFrom *string) (*PreferenceSet, error) {
	return &PreferenceSet{ID: "pref-1", Preferences: []byte("{}")}, nil
}

func (s *stubStore) CreateSession(ctx context.Context, accountID, hash, familyID string, expiresAt time.Time) (*Session, error) {
	ss := &Session{ID: "sess-1", AccountID: accountID, RefreshSecretHash: hash, FamilyID: familyID, ExpiresAt: expiresAt}
	s.sessionsByHash[hash] = ss
	return ss, nil
}

func (s *stubStore) GetSessionBySecretHash(ctx context.Context, hash string) (*Session, error) {
	ss, ok := s.sessionsByHash[hash]
	if !ok {
		return nil, ErrNotFound
	}
	return ss, nil
}

func (s *stubStore) RevokeSession(ctx context.Context, id string) (*Session, error) {
	return &Session{ID: id, RevokedAt: timePtr(time.Now())}, nil
}

func (s *stubStore) MarkSessionReplacedBy(ctx context.Context, id, replacementID string) (*Session, error) {
	return &Session{ID: id, ReplacedBy: &replacementID}, nil
}

func (s *stubStore) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	for _, a := range s.accounts {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, ErrNotFound
}

func (s *stubStore) GetUserProfileByAccountID(ctx context.Context, accountID string) (*UserProfile, error) {
	p, ok := s.profiles[accountID]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (s *stubStore) RevokeSessionFamily(ctx context.Context, familyID string) error { return nil }

func (s *stubStore) GetActivePreferenceSetForProfile(ctx context.Context, profileID string) (*PreferenceSet, error) {
	if s.getActivePreferenceFn != nil {
		return s.getActivePreferenceFn(ctx, profileID)
	}
	return &PreferenceSet{Preferences: []byte("{}")}, nil
}

func timePtr(t time.Time) *time.Time { return &t }

func testTokens(t *testing.T) *auth.TokenManager {
	t.Helper()
	tm, err := auth.NewTokenManager("test-secret-abc", "dp-test", "dp-test", 15*time.Minute, 24*time.Hour)
	if err != nil {
		t.Fatalf("NewTokenManager: %v", err)
	}
	return tm
}

func TestRegister_Success(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	acc, prof, err := svc.Register(context.Background(), RegisterInput{
		Email: "budi@example.com", Password: "password123", DisplayName: "Budi", Timezone: "Asia/Jakarta",
	})
	if err != nil {
		t.Fatalf("Register: %v", err)
	}
	if acc.Email != "budi@example.com" {
		t.Fatalf("email = %q", acc.Email)
	}
	if prof == nil {
		t.Fatal("profile is nil")
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	st := newStubStore()
	st.createAccount = func(ctx context.Context, email, hash string, status AccountStatus, tz string) (*Account, error) {
		return nil, ErrEmailInUse
	}
	svc := NewService(st, testTokens(t))
	_, _, err := svc.Register(context.Background(), RegisterInput{
		Email: "budi@example.com", Password: "password123", DisplayName: "Budi",
	})
	if !errors.Is(err, ErrEmailInUse) {
		t.Fatalf("expected ErrEmailInUse, got %v", err)
	}
}

func TestRegister_WeakPassword(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	_, _, err := svc.Register(context.Background(), RegisterInput{
		Email: "a@b.com", Password: "short", DisplayName: "A",
	})
	if err == nil {
		t.Fatal("expected password error")
	}
}

func TestRegister_InvalidEmail(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	_, _, err := svc.Register(context.Background(), RegisterInput{
		Email: "not-an-email", Password: "password123", DisplayName: "A",
	})
	if err == nil {
		t.Fatal("expected email error")
	}
}

func TestRegister_LongDisplayName(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	long := ""
	for i := 0; i < 150; i++ {
		long += "a"
	}
	_, _, err := svc.Register(context.Background(), RegisterInput{
		Email: "a@b.com", Password: "password123", DisplayName: long,
	})
	if err == nil {
		t.Fatal("expected display name length error")
	}
}

func TestLogin_Success(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	_, _, _ = svc.Register(context.Background(), RegisterInput{
		Email: "budi@example.com", Password: "password123", DisplayName: "Budi",
	})
	res, err := svc.Login(context.Background(), LoginInput{Email: "budi@example.com", Password: "password123"})
	if err != nil {
		t.Fatalf("Login: %v", err)
	}
	if res.AccessToken == "" || res.RefreshToken == "" {
		t.Fatal("tokens are empty")
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	_, _, _ = svc.Register(context.Background(), RegisterInput{
		Email: "budi@example.com", Password: "password123", DisplayName: "Budi",
	})
	_, err := svc.Login(context.Background(), LoginInput{Email: "budi@example.com", Password: "wrongpassword"})
	if err == nil {
		t.Fatal("expected credentials error")
	}
}

func TestLogin_NonexistentEmail(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	_, err := svc.Login(context.Background(), LoginInput{Email: "none@example.com", Password: "password123"})
	if err == nil {
		t.Fatal("expected error for non-existent email")
	}
}

func TestLogout_Success(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	_, _, _ = svc.Register(context.Background(), RegisterInput{
		Email: "budi@example.com", Password: "password123", DisplayName: "Budi",
	})
	res, _ := svc.Login(context.Background(), LoginInput{Email: "budi@example.com", Password: "password123"})
	err := svc.Logout(context.Background(), res.RefreshToken)
	if err != nil {
		t.Fatalf("Logout: %v", err)
	}
}

func TestLogout_NoToken(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	err := svc.Logout(context.Background(), "")
	// log out with no token should not fail — no session to revoke
	if err != nil {
		t.Fatalf("unexpected logout error: %v", err)
	}
}

func TestRefresh_Success(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	_, _, _ = svc.Register(context.Background(), RegisterInput{
		Email: "budi@example.com", Password: "password123", DisplayName: "Budi",
	})
	res, _ := svc.Login(context.Background(), LoginInput{Email: "budi@example.com", Password: "password123"})
	rotated, err := svc.Refresh(context.Background(), res.RefreshToken)
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if rotated.AccessToken == res.AccessToken {
		t.Fatal("access token should change on refresh")
	}
}

func TestRefresh_RevokedReuse(t *testing.T) {
	st := newStubStore()
	svc := NewService(st, testTokens(t))
	_, _, _ = svc.Register(context.Background(), RegisterInput{
		Email: "budi@example.com", Password: "password123", DisplayName: "Budi",
	})
	res, _ := svc.Login(context.Background(), LoginInput{Email: "budi@example.com", Password: "password123"})

	// Manually mark the session as revoked to simulate reuse
	hash := hashRefreshSecret(res.RefreshToken)
	st.sessionsByHash[hash] = &Session{ID: "old", RevokedAt: timePtr(time.Now()), FamilyID: "family-1"}

	_, err := svc.Refresh(context.Background(), res.RefreshToken)
	if err == nil {
		t.Fatal("expected reuse detection error")
	}
}

func TestRefresh_EmptyToken(t *testing.T) {
	svc := NewService(newStubStore(), testTokens(t))
	_, err := svc.Refresh(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty token")
	}
}

func TestUpdatePreferences_Success(t *testing.T) {
	st := newStubStore()
	svc := NewService(st, testTokens(t))
	_, _, _ = svc.Register(context.Background(), RegisterInput{
		Email: "budi@example.com", Password: "password123", DisplayName: "Budi",
	})
	prefs, err := svc.UpdatePreferences(context.Background(), "acc-1", UpdatePreferencesInput{
		Preferences: []byte(`{"diet":"halal"}`),
	})
	if err != nil {
		t.Fatalf("UpdatePreferences: %v", err)
	}
	if prefs == nil {
		t.Fatal("preferences is nil")
	}
}

func TestActivePreferences_NotFound(t *testing.T) {
	st := newStubStore()
	st.getActivePreferenceFn = func(ctx context.Context, profileID string) (*PreferenceSet, error) {
		return nil, ErrNotFound
	}
	svc := NewService(st, testTokens(t))
	prefs, err := svc.ActivePreferences(context.Background(), "prof-1")
	if err != nil {
		t.Fatalf("ActivePreferences: %v", err)
	}
	if string(prefs.Preferences) != "{}" {
		t.Fatalf("expected empty preferences for not-found, got %s", prefs.Preferences)
	}
}
