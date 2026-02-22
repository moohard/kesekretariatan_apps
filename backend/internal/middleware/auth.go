package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// AuthMiddleware untuk verifikasi JWT dari Keycloak
type AuthMiddleware struct {
	jwksURL string
	issuer   string
	audience string
}

// NewAuthMiddleware membuat auth middleware baru
func NewAuthMiddleware(jwksURL, realm string) *AuthMiddleware {
	return &AuthMiddleware{
		jwksURL: jwksURL,
		issuer:   fmt.Sprintf("http://localhost:8081/realms/%s", realm),
		audience: "account",
	}
}

// Authenticate verifikasi JWT token dari Authorization header
func (am *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Skip auth untuk public endpoints
		publicPaths := []string{
			"/health",
			"/api/v1/auth/login",
			"/api/v1/public",
		}

		for _, path := range publicPaths {
			if strings.HasPrefix(c.Path(), path) {
				return c.Next()
			}
		}

		// Get token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": true,
				"message": "Authorization header is missing",
				"code": 401,
				"request_id": GetRequestID(c),
			})
		}

		// Parse Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"error": true,
				"message": "Invalid authorization header format",
				"code": 401,
				"request_id": GetRequestID(c),
			})
		}

		tokenString := parts[1]

		// Fetch JWKS from Keycloak
		jwks, err := jwk.Fetch(c.Context(), am.jwksURL)
		if err != nil {
			fmt.Printf("[ERROR] Failed to fetch JWKS: %v\n", err)
			return c.Status(401).JSON(fiber.Map{
				"error": true,
				"message": "Failed to verify token",
				"code": 401,
				"request_id": GetRequestID(c),
			})
		}

		// Parse and verify JWT
		parsedToken, err := jwt.ParseString(tokenString,
			jwt.WithKeySet(jwks),
		)
		if err != nil {
			fmt.Printf("[ERROR] Failed to parse JWT: %v\n", err)
			return c.Status(401).JSON(fiber.Map{
				"error": true,
				"message": "Invalid token",
				"code": 401,
				"request_id": GetRequestID(c),
			})
		}

		// Extract claims
		sub := parsedToken.Subject()
		if sub == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": true,
				"message": "Invalid token: missing subject",
				"code": 401,
				"request_id": GetRequestID(c),
			})
		}

		// Extract roles from realm_access claim
		roles := make([]string, 0)
		if realmAccessValue, ok := parsedToken.Get("realm_access"); ok {
			if realmAccess, ok := realmAccessValue.(map[string]interface{}); ok {
				if rolesList, ok := realmAccess["roles"].([]interface{}); ok {
					for _, role := range rolesList {
						if roleStr, ok := role.(string); ok {
							roles = append(roles, roleStr)
						}
					}
				}
			}
		}

		// Extract preferred_username
		preferredUsername := ""
		if usernameValue, ok := parsedToken.Get("preferred_username"); ok {
			if username, ok := usernameValue.(string); ok {
				preferredUsername = username
			}
		}

		// Extract email
		email := ""
		if emailValue, ok := parsedToken.Get("email"); ok {
			if emailClaim, ok := emailValue.(string); ok {
				email = emailClaim
			}
		}

		// Extract name
		name := ""
		if nameValue, ok := parsedToken.Get("name"); ok {
			if nameClaim, ok := nameValue.(string); ok {
				name = nameClaim
			}
		}

		// Set user info to context
		c.Locals("userID", sub)
		c.Locals("userRole", am.getHighestRole(roles))
		c.Locals("userRoles", roles)
		c.Locals("username", preferredUsername)
		c.Locals("email", email)
		c.Locals("name", name)

		return c.Next()
	}
}

// getHighestRole mengambil role tertinggi berdasarkan prioritas
func (am *AuthMiddleware) getHighestRole(roles []string) string {
	rolePriority := map[string]int{
		"admin":      100,
		"supervisor": 80,
		"officer":    60,
		"staff":      40,
		"user":       20,
	}

	highestRole := "user"
	highestPriority := 0

	for _, role := range roles {
		if priority, ok := rolePriority[role]; ok && priority > highestPriority {
			highestPriority = priority
			highestRole = role
		}
	}

	return highestRole
}

// RequireRole middleware untuk memeriksa apakah user memiliki role yang diperlukan
func RequireRole(allowedRoles []string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userRole := GetUserRole(c)

		// Admin bypasses role checks
		if userRole == "admin" {
			return c.Next()
		}

		for _, role := range allowedRoles {
			if userRole == role {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{
			"error": true,
			"message": "Insufficient permissions",
			"code": 403,
			"request_id": GetRequestID(c),
		})
	}
}

// RequirePermission middleware untuk memeriksa permissions spesifik
// Ini akan mengambil permission dari database RBAC
func RequirePermission(permission string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := GetUserID(c)
		userRole := GetUserRole(c)

		// Admin bypasses permission checks
		if userRole == "admin" {
			return c.Next()
		}

		// TODO: Implement RBAC permission check from database
		// Untuk saat ini, kita allow semua request
		fmt.Printf("[INFO] Checking permission %s for user %s (role: %s)\n", permission, userID, userRole)

		return c.Next()
	}
}
