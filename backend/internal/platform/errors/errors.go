package apperr

import (
	"errors"
	"net/http"
)

// Code is a stable machine-readable error code from the M6-002 catalog.
type Code string

const (
	// Authentication and session.
	CodeCredentialsInvalid Code = "AUTH_CREDENTIALS_INVALID"
	CodeSessionExpired     Code = "AUTH_SESSION_EXPIRED"
	CodeSessionInvalid     Code = "AUTH_SESSION_INVALID"
	CodeRefreshReused      Code = "AUTH_REFRESH_REUSED"
	CodeAccountRestricted  Code = "AUTH_ACCOUNT_RESTRICTED"
	CodeAccountNotActive   Code = "AUTH_ACCOUNT_NOT_ACTIVE"
	CodeScopeForbidden     Code = "AUTH_SCOPE_FORBIDDEN"

	// Validation.
	CodeFieldInvalid      Code = "VALIDATION_FIELD_INVALID"
	CodeQueryUnsupported  Code = "VALIDATION_QUERY_UNSUPPORTED"
	CodePaginationInvalid Code = "VALIDATION_PAGINATION_INVALID"
	CodeIDInvalid         Code = "VALIDATION_ID_INVALID"
	CodePayloadMalformed  Code = "VALIDATION_PAYLOAD_MALFORMED"

	// Registration and account.
	CodeEmailInUse           Code = "ACCOUNT_EMAIL_IN_USE"
	CodeEmailInvalid         Code = "ACCOUNT_EMAIL_INVALID"
	CodePasswordWeak         Code = "ACCOUNT_PASSWORD_WEAK"
	CodeVerificationRequired Code = "ACCOUNT_VERIFICATION_REQUIRED"

	// Pantry.
	CodePantryItemNotFound  Code = "PANTRY_ITEM_NOT_FOUND"
	CodePantryQtyNegative   Code = "PANTRY_QUANTITY_NEGATIVE"
	CodePantryExpiryInvalid Code = "PANTRY_EXPIRY_INVALID"

	// Recipe.
	CodeRecipeNotFound  Code = "RECIPE_NOT_FOUND"
	CodeRecipeNotPublic Code = "RECIPE_NOT_PUBLIC"

	// Meal plans.
	CodeMealPlanNotFound   Code = "MEAL_PLAN_NOT_FOUND"
	CodeMealPlanPeriodInv  Code = "MEAL_PLAN_PERIOD_INVALID"
	CodeMealSlotConflict   Code = "MEAL_SLOT_CONFLICT"
	CodePlannedMealMissing Code = "PLANNED_MEAL_NOT_FOUND"

	// Shopping.
	CodeShoppingListMissing Code = "SHOPPING_LIST_NOT_FOUND"
	CodeShoppingItemMissing Code = "SHOPPING_ITEM_NOT_FOUND"
	CodeShoppingState       Code = "SHOPPING_STATE_CONFLICT"

	// Recommendations.
	CodeRecommendationMissing       Code = "RECOMMENDATION_NOT_FOUND"
	CodeRecommendationOptionMissing Code = "RECOMMENDATION_OPTION_NOT_FOUND"
	CodeRecommendationStateInvalid  Code = "RECOMMENDATION_STATE_INVALID"
	CodeRecommendationOptionNotOK   Code = "RECOMMENDATION_OPTION_NOT_ACCEPTABLE"
	CodeConversationMissing         Code = "CONVERSATION_NOT_FOUND"
	CodeConversationStateInvalid    Code = "CONVERSATION_STATE_INVALID"

	// AI dependency.
	CodeAIUnavailable    Code = "AI_UNAVAILABLE"
	CodeAISafetyRejected Code = "AI_SAFETY_REJECTED"
	CodeAIQuotaExceeded  Code = "AI_QUOTA_EXCEEDED"

	// Abuse and limits.
	CodeRateLimitExceeded Code = "RATE_LIMIT_EXCEEDED"

	// Internal.
	CodeInternal Code = "INTERNAL_ERROR"
)

// HTTPStatus maps a stable code to its HTTP status per the M6-002 catalog.
func (c Code) HTTPStatus() int {
	switch c {
	case CodeCredentialsInvalid, CodeSessionExpired, CodeSessionInvalid, CodeRefreshReused:
		return http.StatusUnauthorized
	case CodeAccountRestricted, CodeAccountNotActive, CodeScopeForbidden, CodeVerificationRequired:
		return http.StatusForbidden
	case CodeFieldInvalid, CodeQueryUnsupported, CodePaginationInvalid, CodeIDInvalid, CodePayloadMalformed:
		return http.StatusBadRequest
	case CodePantryItemNotFound, CodeRecipeNotFound, CodeRecipeNotPublic, CodeMealPlanNotFound,
		CodePlannedMealMissing, CodeShoppingListMissing, CodeShoppingItemMissing, CodeRecommendationMissing,
		CodeRecommendationOptionMissing, CodeConversationMissing:
		return http.StatusNotFound
	case CodeEmailInUse, CodeMealSlotConflict, CodeShoppingState, CodeRecommendationStateInvalid,
		CodeRecommendationOptionNotOK, CodeConversationStateInvalid:
		return http.StatusConflict
	case CodeAIUnavailable:
		return http.StatusServiceUnavailable
	case CodeAIQuotaExceeded, CodeRateLimitExceeded:
		return http.StatusTooManyRequests
	default:
		return http.StatusUnprocessableEntity
	}
}

// Detail is a field-level validation detail in the error envelope.
type Detail struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error is a domain-aware application error with a stable code, a user-safe
// message, and optional validation details.
type Error struct {
	Code    Code
	Message string
	Details []Detail
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *Error) Unwrap() error { return e.Err }

// New creates an application error with a stable code and user-safe message.
func New(code Code, message string) *Error {
	return &Error{Code: code, Message: message}
}

// Wrap creates an application error wrapping an underlying cause.
func Wrap(code Code, message string, err error) *Error {
	return &Error{Code: code, Message: message, Err: err}
}

// WithDetails adds field-level validation details.
func (e *Error) WithDetails(details ...Detail) *Error {
	e.Details = append(e.Details, details...)
	return e
}

// As extracts an *Error from the wrapped error chain, if present.
func As(err error) (*Error, bool) {
	var appErr *Error
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

// IsNotFound reports whether the error represents a 404-style missing resource.
func IsNotFound(err error) bool {
	appErr, ok := As(err)
	if !ok {
		return false
	}
	return appErr.Code.HTTPStatus() == http.StatusNotFound
}

// AsError extracts an *Error from the chain, returning nil when absent.
func AsError(err error) (*Error, bool) {
	return As(err)
}

// Internal returns a generic internal error with no internal detail exposed.
func Internal(err error) *Error {
	return Wrap(CodeInternal, "An unexpected error occurred.", err)
}
