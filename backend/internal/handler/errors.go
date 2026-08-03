package handler

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/dapurpintar/backend/internal/repo"
	"github.com/gofiber/fiber/v2"
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

func SendError(c *fiber.Ctx, log *slog.Logger, status int, code, message string, fields map[string]string) error {
	reqID, _ := c.Locals("requestid").(string)
	if log != nil && status >= 500 {
		log.Error("request error",
			"status", status,
			"code", code,
			"message", message,
			"path", c.Path(),
			"method", c.Method(),
			"requestId", reqID,
		)
	}
	return c.Status(status).JSON(ErrorResponse{
		Error: code, Message: message, RequestID: reqID, Fields: fields,
	})
}

func SendInternal(c *fiber.Ctx, log *slog.Logger, err error) error {
	reqID, _ := c.Locals("requestid").(string)
	if log != nil {
		log.Error("internal error",
			"err", err.Error(),
			"path", c.Path(),
			"method", c.Method(),
			"requestId", reqID,
		)
	}
	return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
		Error: CodeInternal, Message: err.Error(), RequestID: reqID,
	})
}

// FromError maps known errors to status + code.
func FromError(c *fiber.Ctx, log *slog.Logger, err error) error {
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return SendError(c, log, fiber.StatusNotFound, CodeNotFound, "Recipe not found", nil)
	case errors.Is(err, repo.ErrEmptyLibrary) || strings.Contains(err.Error(), "empty"):
		return SendError(c, log, fiber.StatusServiceUnavailable, CodeEmptyLibrary,
			"The recipe library is empty. Run seed.", nil)
	}
	return SendInternal(c, log, err)
}
