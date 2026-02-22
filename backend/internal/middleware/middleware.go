package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	appErrors "github.com/sikerma/backend/internal/errors"
)

// RequestID middleware untuk menambahkan request ID unik
func RequestID() fiber.Handler {
	return func(c fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Locals("requestID", requestID)
		c.Set("X-Request-ID", requestID)
		return c.Next()
	}
}

// ErrorHandler untuk error handling global dengan format standar
// Format response sesuai PRD Section 15
func ErrorHandler(c fiber.Ctx, err error) error {
	requestID := GetRequestID(c)

	// Log error
	fmt.Printf("[ERROR] RequestID: %s | Error: %v\n", requestID, err)

	// Cek jika error adalah Fiber error
	if e, ok := err.(*fiber.Error); ok {
		// Map Fiber error code ke application error code
		var appCode string
		switch e.Code {
		case fiber.StatusBadRequest:
			appCode = appErrors.ValInvalidFormat
		case fiber.StatusUnauthorized:
			appCode = appErrors.AuthInvalidToken
		case fiber.StatusForbidden:
			appCode = appErrors.AuthzForbidden
		case fiber.StatusNotFound:
			appCode = appErrors.NotFoundResource
		case fiber.StatusTooManyRequests:
			appCode = appErrors.RateLimitExceeded
		default:
			appCode = appErrors.SysInternalError
		}

		return c.Status(e.Code).JSON(fiber.Map{
			"success": false,
			"error": fiber.Map{
				"code":    appCode,
				"message": e.Message,
			},
			"requestId": requestID,
			"timestamp": time.Now().Format(time.RFC3339),
		})
	}

	// Default internal server error
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    appErrors.SysInternalError,
			"message": "Terjadi kesalahan internal sistem",
		},
		"requestId": requestID,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// GetUserID dari context (setelah auth middleware)
func GetUserID(c fiber.Ctx) string {
	if userID, ok := c.Locals("userID").(string); ok {
		return userID
	}
	return ""
}

// GetUserRole dari context (setelah auth middleware)
func GetUserRole(c fiber.Ctx) string {
	if role, ok := c.Locals("userRole").(string); ok {
		return role
	}
	return ""
}

// GetRequestID dari context
func GetRequestID(c fiber.Ctx) string {
	if requestID, ok := c.Locals("requestID").(string); ok {
		return requestID
	}
	return ""
}

// GetUnitKerjaID mengambil unit kerja ID dari context (setelah auth & RLS middleware)
func GetUnitKerjaID(c fiber.Ctx) string {
	if unitKerjaID, ok := c.Locals("unitKerjaID").(string); ok {
		return unitKerjaID
	}
	return ""
}

// GetSatkerID mengambil satker ID dari context (setelah auth & RLS middleware)
func GetSatkerID(c fiber.Ctx) string {
	if satkerID, ok := c.Locals("satkerID").(string); ok {
		return satkerID
	}
	return ""
}