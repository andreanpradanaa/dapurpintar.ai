package auth

import (
	"strings"
	"testing"
	"time"
)

func TestHashAndVerifyPassword(t *testing.T) {
	encoded, err := HashPassword("s3cure-password", DefaultPasswordConfig)
	if err != nil {
		t.Fatalf("HashPassword: %v", err)
	}
	if !strings.HasPrefix(encoded, "$argon2id$") {
		t.Fatalf("expected argon2id prefix, got %s", encoded)
	}

	ok, err := VerifyPassword("s3cure-password", encoded)
	if err != nil {
		t.Fatalf("VerifyPassword: %v", err)
	}
	if !ok {
		t.Fatal("expected correct password to verify")
	}

	ok, err = VerifyPassword("wrong-password", encoded)
	if err != nil {
		t.Fatalf("VerifyPassword wrong: %v", err)
	}
	if ok {
		t.Fatal("expected wrong password not to verify")
	}
}

func TestIssueAndVerifyToken(t *testing.T) {
	m, err := NewTokenManager("test-secret", "dapurpintar", "dapurpintar", 15*time.Minute, 24*time.Hour)
	if err != nil {
		t.Fatalf("NewTokenManager: %v", err)
	}

	session, err := m.Issue("account-123")
	if err != nil {
		t.Fatalf("Issue: %v", err)
	}
	if session.AccessToken == "" || session.RefreshToken == "" {
		t.Fatal("expected access and refresh tokens")
	}

	subject, err := m.Verify(session.AccessToken)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if subject != "account-123" {
		t.Fatalf("expected subject account-123, got %s", subject)
	}
}

func TestVerifyRejectsWrongSecret(t *testing.T) {
	m1, _ := NewTokenManager("secret-one", "dapurpintar", "dapurpintar", 15*time.Minute, 24*time.Hour)
	m2, _ := NewTokenManager("secret-two", "dapurpintar", "dapurpintar", 15*time.Minute, 24*time.Hour)

	session, _ := m1.Issue("account-123")
	if _, err := m2.Verify(session.AccessToken); err == nil {
		t.Fatal("expected verification to fail with a different secret")
	}
}

func TestNewTokenManagerRequiresSecret(t *testing.T) {
	if _, err := NewTokenManager("", "dapurpintar", "dapurpintar", 15*time.Minute, 24*time.Hour); err == nil {
		t.Fatal("expected empty secret to be rejected")
	}
}
