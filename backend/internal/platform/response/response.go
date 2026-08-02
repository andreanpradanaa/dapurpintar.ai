package response

import (
	"github.com/gofiber/fiber/v2"

	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

// Page is the pagination object in the success envelope.
type Page struct {
	NextCursor *string `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
}

// Envelope is the success response shape from the M6 API contract.
type Envelope struct {
	Data      any    `json:"data"`
	Page      *Page  `json:"page"`
	RequestID string `json:"request_id"`
}

// ErrDetail is a field-level validation detail in the error envelope.
type ErrDetail struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorBody is the error response shape from the M6 API contract.
type ErrorBody struct {
	Error ErrorEnvelope `json:"error"`
}

// ErrorEnvelope carries the stable error code, user-safe message, and details.
type ErrorEnvelope struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Details   []ErrDetail `json:"details"`
	RequestID string      `json:"request_id"`
}

// requestID resolves the correlation identifier from the fiber context.
func requestID(c *fiber.Ctx) string {
	if id := c.GetRespHeader("X-Request-ID"); id != "" {
		return id
	}
	return c.Locals("request_id").(string)
}

// OK sends a success envelope with 200.
func OK(c *fiber.Ctx, data any) error {
	return c.JSON(Envelope{Data: data, Page: nil, RequestID: requestID(c)})
}

// Created sends a success envelope with 201.
func Created(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(Envelope{Data: data, Page: nil, RequestID: requestID(c)})
}

// NoContent sends 204 with no body.
func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// Collection sends a paginated collection envelope.
func Collection(c *fiber.Ctx, data any, page *Page) error {
	return c.JSON(Envelope{Data: data, Page: page, RequestID: requestID(c)})
}

// Error sends the stable error envelope for the given error.
// Unknown errors become a generic internal error without leaking internals.
func Error(c *fiber.Ctx, err error) error {
	appErr, ok := apperr.AsError(err)
	if !ok {
		appErr = apperr.New(apperr.CodeInternal, "An unexpected error occurred.")
	}

	details := make([]ErrDetail, 0, len(appErr.Details))
	for _, d := range appErr.Details {
		details = append(details, ErrDetail{Field: d.Field, Code: d.Code, Message: d.Message})
	}

	return c.Status(appErr.Code.HTTPStatus()).JSON(ErrorBody{
		Error: ErrorEnvelope{
			Code:      string(appErr.Code),
			Message:   appErr.Message,
			Details:   details,
			RequestID: requestID(c),
		},
	})
}
