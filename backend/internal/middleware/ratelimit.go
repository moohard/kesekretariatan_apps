package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/google/uuid"
)

// RateLimitConfig untuk konfigurasi rate limiting
type RateLimitConfig struct {
	GlobalMax      int
	GlobalWindow   time.Duration
	LoginMax       int
	LoginWindow    time.Duration
	UploadMax      int
	UploadWindow   time.Duration
}

// DefaultRateLimitConfig mengembalikan konfigurasi default
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		GlobalMax:    100,             // 100 requests
		GlobalWindow: 1 * time.Minute, // per minute
		LoginMax:     5,               // 5 attempts
		LoginWindow:  15 * time.Minute, // per 15 minutes
		UploadMax:    10,              // 10 uploads
		UploadWindow: 1 * time.Minute, // per minute
	}
}

// GlobalRateLimiter membatasi request per IP secara global
// Default: 100 requests per menit per IP
func GlobalRateLimiter(config RateLimitConfig) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        config.GlobalMax,
		Expiration: config.GlobalWindow,
		KeyGenerator: func(c fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			requestID := GetRequestID(c)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "RATE_LIMIT_EXCEEDED",
					"message": "Terlalu banyak permintaan, coba lagi dalam 1 menit",
				},
				"request_id": requestID,
				"timestamp":  time.Now().Format(time.RFC3339),
			})
		},
		SkipFailedRequests: false,
		SkipSuccessfulRequests: false,
	})
}

// LoginRateLimiter membatasi percobaan login untuk mencegah brute force
// Default: 5 attempts per 15 menit per IP
func LoginRateLimiter(config RateLimitConfig) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        config.LoginMax,
		Expiration: config.LoginWindow,
		KeyGenerator: func(c fiber.Ctx) string {
			// Combine IP with "login" prefix untuk isolasi
			return "login_" + c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			requestID := GetRequestID(c)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "LOGIN_RATE_LIMIT_EXCEEDED",
					"message": "Terlalu banyak percobaan login, coba lagi dalam 15 menit",
				},
				"request_id": requestID,
				"timestamp":  time.Now().Format(time.RFC3339),
			})
		},
	})
}

// UploadRateLimiter membatasi upload file per user
// Default: 10 uploads per menit per user
func UploadRateLimiter(config RateLimitConfig) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        config.UploadMax,
		Expiration: config.UploadWindow,
		KeyGenerator: func(c fiber.Ctx) string {
			// Use userID if authenticated, otherwise use IP
			userID := GetUserID(c)
			if userID != "" {
				return "upload_" + userID
			}
			return "upload_" + c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			requestID := GetRequestID(c)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "UPLOAD_RATE_LIMIT_EXCEEDED",
					"message": "Terlalu banyak upload, coba lagi dalam 1 menit",
				},
				"request_id": requestID,
				"timestamp":  time.Now().Format(time.RFC3339),
			})
		},
	})
}

// APIKeyRateLimiter rate limiter khusus untuk API key
// Lebih longgar dibanding rate limiter global
func APIKeyRateLimiter(max int, window time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: window,
		KeyGenerator: func(c fiber.Ctx) string {
			// Use API key if present, otherwise use IP
			apiKey := c.Get("X-API-Key")
			if apiKey != "" {
				return "apikey_" + apiKey
			}
			return "noapikey_" + c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			requestID := GetRequestID(c)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "API_RATE_LIMIT_EXCEEDED",
					"message": "Rate limit API exceeded",
				},
				"request_id": requestID,
				"timestamp":  time.Now().Format(time.RFC3339),
			})
		},
	})
}

// CustomRateLimiter untuk kasus khusus dengan konfigurasi fleksibel
func CustomRateLimiter(max int, window time.Duration, keyPrefix string) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: window,
		KeyGenerator: func(c fiber.Ctx) string {
			userID := GetUserID(c)
			if userID != "" {
				return keyPrefix + "_" + userID
			}
			return keyPrefix + "_" + c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			requestID := GetRequestID(c)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "CUSTOM_RATE_LIMIT_EXCEEDED",
					"message": "Terlalu banyak permintaan",
				},
				"request_id": requestID,
				"timestamp":  time.Now().Format(time.RFC3339),
			})
		},
	})
}

// GenerateKey generates a unique key for rate limiting
func GenerateKey(prefix string) string {
	return prefix + "_" + uuid.New().String()
}
