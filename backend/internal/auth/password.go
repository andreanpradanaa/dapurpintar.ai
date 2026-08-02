package auth

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// PasswordConfig defines the Argon2id parameters for password hashing.
type PasswordConfig struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
	SaltLen uint32
}

// DefaultPasswordConfig is the MVP Argon2id baseline. It is memory-hard and
// adaptive; parameters may be tuned without changing the format.
var DefaultPasswordConfig = PasswordConfig{
	Time:    3,
	Memory:  64 * 1024,
	Threads: 2,
	KeyLen:  32,
	SaltLen: 16,
}

// HashPassword hashes a plaintext password using Argon2id. The encoded value
// carries the parameters so verification does not depend on current defaults.
func HashPassword(password string, cfg PasswordConfig) (string, error) {
	if password == "" {
		return "", errors.New("password is required")
	}
	salt := make([]byte, cfg.SaltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, cfg.Time, cfg.Memory, cfg.Threads, cfg.KeyLen)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, cfg.Memory, cfg.Time, cfg.Threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	), nil
}

// VerifyPassword compares a password against an encoded Argon2id hash using a
// timing-safe operation. It returns true when the password matches.
func VerifyPassword(password, encoded string) (bool, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return false, errors.New("invalid password hash format")
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return false, err
	}

	var memory uint32
	var time uint32
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d", &memory, &time); err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}
	expected, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	actual := argon2.IDKey([]byte(password), salt, time, memory, 2, uint32(len(expected)))

	if subtle.ConstantTimeCompare(actual, expected) == 1 {
		return true, nil
	}
	return false, nil
}

// ctxKey is the context key type for the authenticated subject.
type ctxKey string

const subjectKey ctxKey = "auth.subject"

// WithSubject stores the authenticated subject (Account/User identity) in the
// context. Authorization is still evaluated server-side per resource.
func WithSubject(ctx context.Context, subject string) context.Context {
	return context.WithValue(ctx, subjectKey, subject)
}

// SubjectFrom returns the authenticated subject stored in the context.
func SubjectFrom(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(subjectKey).(string)
	return v, ok
}
