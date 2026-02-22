package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/sikerma/backend/internal/config"
	"github.com/sikerma/backend/internal/database"
	"github.com/sikerma/backend/internal/handlers"
	customMiddleware "github.com/sikerma/backend/internal/middleware"
	"github.com/sikerma/backend/internal/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database connections
	dbMaster, dbKepegawaian, err := database.InitConnections(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database connections: %v", err)
	}
	defer database.Close(dbMaster, dbKepegawaian)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:       "SIKERMA Backend API",
		ServerHeader:  "SIKERMA",
		EnablePrintRoutes: os.Getenv("LOG_LEVEL") == "debug",
		ErrorHandler:  customMiddleware.ErrorHandler,
		BodyLimit:     10 * 1024 * 1024, // 10MB
		ReadTimeout:   30 * time.Second,
		WriteTimeout:  30 * time.Second,
		IdleTimeout:   60 * time.Second,
	})

	// Middleware
	app.Use(recover.New())

	// Security Headers (Helmet)
	app.Use(helmet.New(helmet.Config{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		HSTSSeconds:        31536000,
		HSTSIncludeSubdomains: true,
		HSTSPreloadEnabled: true,
		ReferrerPolicy:     "strict-origin-when-cross-origin",
		PermissionsPolicy:  "geolocation=(), microphone=(), camera=(), payment=(), usb=()",
	}))

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.Origins,
		AllowCredentials: cfg.CORS.Credentials,
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Request-ID,X-CSRF-Token",
		ExposeHeaders:    "X-Request-ID,X-CSRF-Token",
	}))

	// CSRF Protection
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-CSRF-Token",
		CookieName:     "csrf_",
		CookieSecure:   cfg.Environment == "production",
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   customMiddleware.GenerateCSRFToken,
		ContextKey:     "csrf",
	}))

	// Global Rate Limiting (100 req/min per IP)
	rateLimitConfig := customMiddleware.DefaultRateLimitConfig()
	app.Use(customMiddleware.GlobalRateLimiter(rateLimitConfig))

	// Compression
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Request Logging
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	// Custom Middleware
	app.Use(customMiddleware.RequestID())
	app.Use(customMiddleware.AuditTrail(dbMaster))

	// Health check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "healthy",
			"time":    time.Now().Format(time.RFC3339),
			"version": "1.0.0",
		})
	})

	// Initialize handlers
	h := handlers.New(dbMaster, dbKepegawaian, cfg)

	// Setup routes
	routes.Setup(app, h)

	// Graceful shutdown
	go gracefulShutdown(app, cfg)

	// Start server
	addr := cfg.Host + ":" + cfg.Port
	log.Printf("Server starting on %s", addr)
	log.Printf("Environment: %s", cfg.Environment)
	log.Printf("Database Master: %s", cfg.DBMaster.DSN())
	log.Printf("Database Kepegawaian: %s", cfg.DBKepegawaian.DSN())
	log.Printf("Keycloak URL: %s", cfg.KeycloakURL)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func gracefulShutdown(app *fiber.App, cfg *config.Config) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}