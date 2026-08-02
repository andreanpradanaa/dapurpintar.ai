package apperr

import (
	"errors"
	"net/http"
	"testing"
)

func TestHTTPStatusMapping(t *testing.T) {
	cases := []struct {
		code   Code
		status int
	}{
		{CodeCredentialsInvalid, http.StatusUnauthorized},
		{CodeSessionExpired, http.StatusUnauthorized},
		{CodeAccountRestricted, http.StatusForbidden},
		{CodeScopeForbidden, http.StatusForbidden},
		{CodeFieldInvalid, http.StatusBadRequest},
		{CodeQueryUnsupported, http.StatusBadRequest},
		{CodePaginationInvalid, http.StatusBadRequest},
		{CodeEmailInUse, http.StatusConflict},
		{CodeMealSlotConflict, http.StatusConflict},
		{CodeRecommendationStateInvalid, http.StatusConflict},
		{CodeAIUnavailable, http.StatusServiceUnavailable},
		{CodeRateLimitExceeded, http.StatusTooManyRequests},
		{CodeAIQuotaExceeded, http.StatusTooManyRequests},
		{CodePantryExpiryInvalid, http.StatusUnprocessableEntity},
	}

	for _, c := range cases {
		if got := c.code.HTTPStatus(); got != c.status {
			t.Errorf("code %s: expected status %d, got %d", c.code, c.status, got)
		}
	}
}

func TestNewAndWrap(t *testing.T) {
	err := New(CodePantryItemNotFound, "The pantry item is not available.")
	if err.Code != CodePantryItemNotFound {
		t.Fatalf("expected code %s, got %s", CodePantryItemNotFound, err.Code)
	}
	if !IsNotFound(err) {
		t.Fatal("expected IsNotFound to be true")
	}

	cause := errors.New("root cause")
	wrapped := Wrap(CodeInternal, "An unexpected error occurred.", cause)
	if !errors.Is(wrapped, cause) {
		t.Fatal("expected wrapped error to unwrap to cause")
	}

	if got, ok := As(wrapped); !ok || got.Code != CodeInternal {
		t.Fatal("expected As to extract CodeInternal")
	}
}

func TestWithDetails(t *testing.T) {
	err := New(CodeFieldInvalid, "Validation failed.").WithDetails(
		Detail{Field: "quantity", Code: "PANTRY_QUANTITY_NEGATIVE", Message: "Quantity cannot be negative."},
	)
	if len(err.Details) != 1 {
		t.Fatalf("expected 1 detail, got %d", len(err.Details))
	}
	if err.Details[0].Field != "quantity" {
		t.Fatalf("expected field quantity, got %s", err.Details[0].Field)
	}
}

func TestInternalErrorDoesNotLeak(t *testing.T) {
	err := Internal(errors.New("sql: connection failed"))
	appErr, ok := As(err)
	if !ok {
		t.Fatal("expected As to succeed")
	}
	if appErr.Message == "sql: connection failed" {
		t.Fatal("internal error message must not leak internals")
	}
}
