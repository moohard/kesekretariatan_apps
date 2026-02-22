package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// SecurityHeadersConfig untuk konfigurasi security headers
type SecurityHeadersConfig struct {
	// XSSProtection enables X-XSS-Protection header
	XSSProtection string
	// ContentTypeNosniff enables X-Content-Type-Options header
	ContentTypeNosniff string
	// XFrameOptions enables X-Frame-Options header
	XFrameOptions string
	// HSTSEnabled enables Strict-Transport-Security header
	HSTSEnabled bool
	// HSTSMaxAge sets max-age for HSTS
	HSTSMaxAge int
	// HSTSIncludeSubdomains includes subdomains in HSTS
	HSTSIncludeSubdomains bool
	// HSTSPreload enables preload for HSTS
	HSTSPreload bool
	// ContentSecurityPolicy sets CSP header
	ContentSecurityPolicy string
	// ReferrerPolicy sets Referrer-Policy header
	ReferrerPolicy string
	// PermissionsPolicy sets Permissions-Policy header
	PermissionsPolicy string
	// IsProduction flag for production-specific settings
	IsProduction bool
}

// DefaultSecurityHeadersConfig mengembalikan konfigurasi default
func DefaultSecurityHeadersConfig(isProduction bool) SecurityHeadersConfig {
	return SecurityHeadersConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		HSTSEnabled:        isProduction,
		HSTSMaxAge:         31536000, // 1 year
		HSTSIncludeSubdomains: true,
		HSTSPreload:        true,
		ContentSecurityPolicy: "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline' 'unsafe-eval'; " +
			"style-src 'self' 'unsafe-inline'; " +
			"img-src 'self' data: blob: https:; " +
			"font-src 'self' data:; " +
			"connect-src 'self' https:; " +
			"frame-ancestors 'self';",
		ReferrerPolicy:  "strict-origin-when-cross-origin",
		PermissionsPolicy: "geolocation=(), microphone=(), camera=(), payment=(), usb=()",
		IsProduction:    isProduction,
	}
}

// SecurityHeaders middleware menambahkan security headers
func SecurityHeaders(config SecurityHeadersConfig) fiber.Handler {
	return func(c fiber.Ctx) error {
		// X-XSS-Protection
		if config.XSSProtection != "" {
			c.Set("X-XSS-Protection", config.XSSProtection)
		}

		// X-Content-Type-Options
		if config.ContentTypeNosniff != "" {
			c.Set("X-Content-Type-Options", config.ContentTypeNosniff)
		}

		// X-Frame-Options
		if config.XFrameOptions != "" {
			c.Set("X-Frame-Options", config.XFrameOptions)
		}

		// Strict-Transport-Security (HSTS) - only in production
		if config.HSTSEnabled && config.IsProduction {
			hstsValue := "max-age=" + itoa(config.HSTSMaxAge)
			if config.HSTSIncludeSubdomains {
				hstsValue += "; includeSubDomains"
			}
			if config.HSTSPreload {
				hstsValue += "; preload"
			}
			c.Set("Strict-Transport-Security", hstsValue)
		}

		// Content-Security-Policy
		if config.ContentSecurityPolicy != "" {
			c.Set("Content-Security-Policy", config.ContentSecurityPolicy)
		}

		// Referrer-Policy
		if config.ReferrerPolicy != "" {
			c.Set("Referrer-Policy", config.ReferrerPolicy)
		}

		// Permissions-Policy
		if config.PermissionsPolicy != "" {
			c.Set("Permissions-Policy", config.PermissionsPolicy)
		}

		// Additional security headers
		c.Set("X-Permitted-Cross-Domain-Policies", "none")
		c.Set("Cross-Origin-Embedder-Policy", "require-corp")
		c.Set("Cross-Origin-Opener-Policy", "same-origin")
		c.Set("Cross-Origin-Resource-Policy", "same-origin")

		// Cache control for API responses
		if !config.IsProduction {
			c.Set("Cache-Control", "no-store, no-cache, must-revalidate")
		}

		return c.Next()
	}
}

// CSRFConfig untuk konfigurasi CSRF protection
type CSRFConfig struct {
	// KeyLookup where to look for CSRF token ("header:X-CSRF-Token" or "form:_csrf")
	KeyLookup string
	// CookieName name of the CSRF cookie
	CookieName string
	// CookieSecure sets Secure flag on cookie
	CookieSecure bool
	// CookieHTTPOnly sets HttpOnly flag on cookie
	CookieHTTPOnly bool
	// CookieSameSite sets SameSite attribute
	CookieSameSite string
	// Expiration time for CSRF token
	Expiration time.Duration
	// KeyGenerator function to generate CSRF key
	KeyGenerator func() string
	// ContextKey key to store CSRF token in context
	ContextKey string
	// ErrorHandler custom error handler
	ErrorHandler fiber.ErrorHandler
	// Extractor custom token extractor
	Extractor func(fiber.Ctx) (string, error)
}

// DefaultCSRFConfig mengembalikan konfigurasi CSRF default
func DefaultCSRFConfig(isProduction bool) CSRFConfig {
	return CSRFConfig{
		KeyLookup:      "header:X-CSRF-Token",
		CookieName:     "csrf_",
		CookieSecure:   isProduction,
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		ContextKey:     "csrf",
		KeyGenerator:   generateCSRFToken,
		ErrorHandler:   defaultCSRFErrorHandler,
	}
}

// CSRFProtection middleware untuk perlindungan CSRF
func CSRFProtection(config CSRFConfig) fiber.Handler {
	// Simple CSRF implementation
	// For production, consider using the fiber/csrf middleware
	return func(c fiber.Ctx) error {
		// Skip CSRF for safe methods (GET, HEAD, OPTIONS, TRACE)
		method := c.Method()
		if method == fiber.MethodGet || method == fiber.MethodHead ||
			method == fiber.MethodOptions || method == fiber.MethodTrace {
			return c.Next()
		}

		// Get token from header
		token := c.Get("X-CSRF-Token")
		if token == "" {
			return config.ErrorHandler(c, fiber.NewError(fiber.StatusForbidden,
				"CSRF token missing"))
		}

		// Get token from cookie
		cookieToken := c.Cookies(config.CookieName)
		if cookieToken == "" {
			return config.ErrorHandler(c, fiber.NewError(fiber.StatusForbidden,
				"CSRF cookie missing"))
		}

		// Compare tokens (in production, use constant-time comparison)
		if token != cookieToken {
			return config.ErrorHandler(c, fiber.NewError(fiber.StatusForbidden,
				"CSRF token invalid"))
		}

		return c.Next()
	}
}

// CSRFTokenHandler generates and sets CSRF token
func CSRFTokenHandler(config CSRFConfig) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Generate new token
		token := config.KeyGenerator()

		// Set cookie
		cookie := &fiber.Cookie{
			Name:     config.CookieName,
			Value:    token,
			Expires:  time.Now().Add(config.Expiration),
			HTTPOnly: config.CookieHTTPOnly,
			Secure:   config.CookieSecure,
			SameSite: parseSameSite(config.CookieSameSite),
		}
		c.Cookie(cookie)

		// Store in context
		c.Locals(config.ContextKey, token)

		// Set header for client to read
		c.Set("X-CSRF-Token", token)

		return c.Next()
	}
}

// Helper functions

// GenerateCSRFToken generates a new CSRF token (exported for use in main.go)
func GenerateCSRFToken() string {
	return uuid.New().String()
}

func generateCSRFToken() string {
	return GenerateCSRFToken()
}

func defaultCSRFErrorHandler(c fiber.Ctx, err error) error {
	requestID := GetRequestID(c)
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    "CSRF_TOKEN_INVALID",
			"message": err.Error(),
		},
		"request_id": requestID,
		"timestamp":  time.Now().Format(time.RFC3339),
	})
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}

	neg := false
	if i < 0 {
		neg = true
		i = -i
	}

	var digits []byte
	for i > 0 {
		digits = append([]byte{byte('0' + i%10)}, digits...)
		i /= 10
	}

	if neg {
		digits = append([]byte{'-'}, digits...)
	}

	return string(digits)
}

func parseSameSite(s string) string {
	switch s {
	case "Strict", "Lax", "None":
		return s
	default:
		return "Strict"
	}
}
