package handler

import (
	"errors"

	"github.com/dapurpintar/backend/internal/repo"
	"github.com/dapurpintar/backend/internal/service/llm"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type ErrorResponse struct {
	Error     string            `json:"error"`
	Message   string            `json:"message"`
	RequestID string            `json:"requestId,omitempty"`
	Fields    map[string]string `json:"fields,omitempty"`
}

const (
	CodeValidation      = "validation_error"
	CodeNotFound        = "not_found"
	CodeMethod          = "method_not_allowed"
	CodeInternal        = "internal_error"
	CodeUpstream        = "upstream_error"
	CodeUpstreamTimeout = "upstream_timeout"
	CodeEmptyLibrary    = "empty_library"
)

func SendError(c *fiber.Ctx, log *zerolog.Logger, status int, code, message string, fields map[string]string) error {
	reqID, _ := c.Locals("requestid").(string)
	if log != nil && status >= 500 {
		log.Error().
			Int("status", status).
			Str("code", code).
			Str("message", message).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("requestId", reqID).
			Msg("request error")
	}
	return c.Status(status).JSON(ErrorResponse{
		Error: code, Message: message, RequestID: reqID, Fields: fields,
	})
}

func SendInternal(c *fiber.Ctx, log *zerolog.Logger, err error) error {
	reqID, _ := c.Locals("requestid").(string)
	if log != nil {
		log.Error().
			Str("err", err.Error()).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("requestId", reqID).
			Msg("internal error")
	}
	return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
		Error: CodeInternal, Message: err.Error(), RequestID: reqID,
	})
}

// FromError maps known errors to status + code.
func FromError(c *fiber.Ctx, log *zerolog.Logger, err error) error {
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return SendError(c, log, fiber.StatusNotFound, CodeNotFound, "Recipe not found", nil)
	case errors.Is(err, repo.ErrEmptyLibrary):
		return SendError(c, log, fiber.StatusServiceUnavailable, CodeEmptyLibrary,
			"The recipe library is empty. Run seed.", nil)
	case errors.Is(err, llm.ErrTimeout):
		return SendError(c, log, fiber.StatusGatewayTimeout, CodeUpstreamTimeout,
			"LLM took too long to respond. Please try again.", nil)
	case errors.Is(err, llm.ErrRateLimit):
		return SendError(c, log, fiber.StatusTooManyRequests, CodeUpstream,
			"LLM rate limit reached. Please wait a moment and try again.", nil)
	case errors.Is(err, llm.ErrBadOutput):
		return SendError(c, log, fiber.StatusBadGateway, CodeUpstream,
			"LLM returned malformed output. Please try again.", nil)
	}
	return SendInternal(c, log, err)
}
