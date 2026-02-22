package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sikerma/backend/internal/repositories"
	"github.com/sikerma/backend/internal/utils"
)

// ============================================
// Audit Middleware Configuration
// ============================================

// AuditConfig konfigurasi untuk audit middleware
type AuditConfig struct {
	// Skip paths yang tidak perlu diaudit
	SkipPaths []string
	// Methods yang perlu diaudit (default: POST, PUT, DELETE, PATCH)
	Methods []string
	// Enable PII masking
	EnablePIIMasking bool
	// Async logging (tidak blocking request)
	AsyncLogging bool
}

// DefaultAuditConfig mengembalikan konfigurasi default
func DefaultAuditConfig() AuditConfig {
	return AuditConfig{
		SkipPaths: []string{
			"/health",
			"/metrics",
			"/favicon.ico",
		},
		Methods: []string{
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		},
		EnablePIIMasking: true,
		AsyncLogging:     true,
	}
}

// ============================================
// Audit Middleware
// ============================================

// AuditTrail middleware untuk mencatat audit trail ke database
// dengan PII masking untuk data sensitif
func AuditTrail(db *pgxpool.Pool, config ...AuditConfig) fiber.Handler {
	cfg := DefaultAuditConfig()
	if len(config) > 0 {
		cfg = config[0]
	}

	auditRepo := repositories.NewAuditRepository(db)

	return func(c fiber.Ctx) error {
		startTime := time.Now()

		// Skip jika path tidak perlu diaudit
		path := c.Path()
		for _, skipPath := range cfg.SkipPaths {
			if strings.HasPrefix(path, skipPath) {
				return c.Next()
			}
		}

		// Capture request body untuk POST/PUT/PATCH
		var requestBody map[string]interface{}
		method := c.Method()
		shouldCaptureBody := false
		for _, m := range cfg.Methods {
			if method == m {
				shouldCaptureBody = true
				break
			}
		}

		if shouldCaptureBody && len(c.Body()) > 0 {
			// Read and restore body
			bodyBytes := c.Body()
			if err := json.Unmarshal(bodyBytes, &requestBody); err == nil && cfg.EnablePIIMasking {
				requestBody = utils.MaskSensitiveData(requestBody)
			}
		}

		// Process request
		err := c.Next()

		// Capture response info
		duration := time.Since(startTime)
		statusCode := c.Response().StatusCode()

		// Determine action based on method
		action := getActionFromMethod(method)
		if action == "" {
			return err
		}

		// Build audit log
		requestID := GetRequestID(c)
		userID := GetUserID(c)
		username := getUsername(c)

		// Determine resource from path
		resource := getResourceFromPath(path)

		// Get resource ID if available
		var resourceID *uuid.UUID
		if id := c.Params("id"); id != "" {
			if parsedID, parseErr := uuid.Parse(id); parseErr == nil {
				resourceID = &parsedID
			}
		}

		// Build changes
		changes := map[string]interface{}{
			"request_id":   requestID,
			"method":       method,
			"path":         path,
			"status_code":  statusCode,
			"duration_ms":  duration.Milliseconds(),
			"request_body": requestBody,
		}

		// Get IP address
		ipAddress := c.IP()

		// Get user agent
		userAgent := c.Get(fiber.HeaderUserAgent)

		// Determine status
		status := "success"
		var errorMessage *string
		if err != nil || statusCode >= 400 {
			status = "failed"
			if err != nil {
				errMsg := err.Error()
				errorMessage = &errMsg
			}
		}

		// Create audit input
		auditInput := repositories.AuditLogInput{
			UserID:       userID,
			Username:     username,
			Action:       action,
			Resource:     resource,
			ResourceID:   resourceID,
			IPAddress:    &ipAddress,
			UserAgent:    &userAgent,
			Changes:      changes,
			Status:       status,
			ErrorMessage: errorMessage,
		}

		// Log to console for debugging
		fmt.Printf("[AUDIT] RequestID: %s | User: %s | Action: %s | Resource: %s | Status: %s | Duration: %v\n",
			requestID, username, action, resource, status, duration)

		// Save to database
		if cfg.AsyncLogging {
			// Async logging - tidak blocking request
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if logErr := auditRepo.Log(ctx, auditInput); logErr != nil {
					fmt.Printf("[AUDIT ERROR] Failed to save audit log: %v\n", logErr)
				}
			}()
		} else {
			// Sync logging
			if logErr := auditRepo.Log(c.Context(), auditInput); logErr != nil {
				fmt.Printf("[AUDIT ERROR] Failed to save audit log: %v\n", logErr)
			}
		}

		return err
	}
}

// ============================================
// Helper Functions
// ============================================

// getActionFromMethod mengembalikan action berdasarkan HTTP method
func getActionFromMethod(method string) string {
	switch method {
	case fiber.MethodPost:
		return "create"
	case fiber.MethodPut, fiber.MethodPatch:
		return "update"
	case fiber.MethodDelete:
		return "delete"
	default:
		return ""
	}
}

// getResourceFromPath mengekstrak resource dari path
func getResourceFromPath(path string) string {
	// Remove leading slash and split
	parts := strings.Split(strings.Trim(path, "/"), "/")

	// Get first part as resource
	if len(parts) > 0 {
		resource := parts[0]

		// Handle API prefix
		if resource == "api" && len(parts) > 1 {
			resource = parts[1]
		}

		// Handle version prefix
		if strings.HasPrefix(resource, "v") && len(parts) > 2 {
			resource = parts[2]
		}

		return resource
	}

	return "unknown"
}

// getUsername mendapatkan username dari context
func getUsername(c fiber.Ctx) string {
	if username, ok := c.Locals("username").(string); ok {
		return username
	}
	return "anonymous"
}

// ============================================
// Response Body Capture (Optional)
// ============================================

// ResponseWriter wrapper untuk capture response body
type responseWriter struct {
	fiber.Ctx
	body *bytes.Buffer
}

// Write menangkap response body
func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.Ctx.Write(b)
}

// CaptureResponseBody middleware untuk menangkap response body
// Gunakan dengan hati-hati karena bisa mempengaruhi performance
func CaptureResponseBody() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Create wrapper
		wrapper := &responseWriter{
			Ctx:  c,
			body: &bytes.Buffer{},
		}

		// Process request
		err := c.Next()

		// Store response body in locals
		c.Locals("responseBody", wrapper.body.String())

		return err
	}
}

// ReadBody reads and restores request body
func ReadBody(c fiber.Ctx) ([]byte, error) {
	body := c.Body()
	if len(body) == 0 {
		return nil, nil
	}

	// Create a copy for reading
	bodyCopy := make([]byte, len(body))
	copy(bodyCopy, body)

	return bodyCopy, nil
}

// ReadBodyAsMap reads request body as map with PII masking
func ReadBodyAsMap(c fiber.Ctx) (map[string]interface{}, error) {
	body, err := ReadBody(c)
	if err != nil || len(body) == 0 {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return utils.MaskSensitiveData(data), nil
}

// ReadBodyAsReader returns body as io.Reader
func ReadBodyAsReader(c fiber.Ctx) io.Reader {
	return bytes.NewReader(c.Body())
}
