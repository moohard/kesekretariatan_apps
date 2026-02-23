package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/sikerma/backend/internal/handlers"
	"github.com/sikerma/backend/internal/middleware"
)

// Setup mengonfigurasi semua routes aplikasi
func Setup(app *fiber.App, h *handlers.Handlers) {
	api := app.Group("/api/v1")

	// Health check (public)
	api.Get("/health", h.HealthCheck)

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/login", h.Login)
	auth.Post("/logout", h.Logout)

	// Authenticated routes
	authenticated := api.Group("", h.AuthMiddleware.Authenticate())

	// User profile
	authenticated.Get("/auth/me", h.GetCurrentUser)

	// ==================== MASTER DATA ====================
	masterData := authenticated.Group("/master-data")
	masterData.Use(middleware.RequirePermission("master_data.read"))

	// Satker
	satker := masterData.Group("/satker")
	satker.Get("", h.ListSatker)
	satker.Get("/dropdown", h.GetDropdownSatker)
	satker.Get("/:id", h.GetSatker)
	satker.Post("", middleware.RequirePermission("master_data.create"), h.CreateSatker)
	satker.Put("/:id", middleware.RequirePermission("master_data.update"), h.UpdateSatker)
	satker.Delete("/:id", middleware.RequirePermission("master_data.delete"), h.DeleteSatker)

	// Jabatan
	jabatan := masterData.Group("/jabatan")
	jabatan.Get("", h.ListJabatan)
	jabatan.Get("/dropdown", h.GetDropdownJabatan)

	// Golongan
	golongan := masterData.Group("/golongan")
	golongan.Get("", h.ListGolongan)
	golongan.Get("/dropdown", h.GetDropdownGolongan)

	// Unit Kerja
	unitKerja := masterData.Group("/unit-kerja")
	unitKerja.Get("", h.ListUnitKerja)
	unitKerja.Get("/dropdown", h.GetDropdownUnitKerja)

	// Eselon
	eselon := masterData.Group("/eselon")
	eselon.Get("", h.ListEselon)
	eselon.Get("/dropdown", h.GetDropdownEselon)

	// ==================== KEGAWAAN ====================
	kepegawaian := authenticated.Group("/kepegawaian")
	kepegawaian.Use(middleware.RequirePermission("kepegawaian.read"))

	// Pegawai
	pegawai := kepegawaian.Group("/pegawai")
	pegawai.Get("", h.ListPegawai)
	pegawai.Get("/:id", h.GetPegawai)
	pegawai.Post("", middleware.RequirePermission("kepegawaian.create"), h.CreatePegawai)
	pegawai.Put("/:id", middleware.RequirePermission("kepegawaian.update"), h.UpdatePegawai)
	pegawai.Delete("/:id", middleware.RequirePermission("kepegawaian.delete"), h.DeletePegawai)

	// Statistik
	kepegawaian.Get("/statistik", h.GetStatistikKepegawaian)

	// ==================== RBAC ====================
	rbac := authenticated.Group("/rbac")
	rbac.Use(middleware.RequirePermission("rbac.read"))

	roles := rbac.Group("/roles")
	roles.Get("", h.ListRoles)
	roles.Post("", middleware.RequirePermission("rbac.create"), h.CreateRole)

	// ==================== AUDIT LOGS ====================
	audit := authenticated.Group("/audit-logs")
	audit.Use(middleware.RequirePermission("audit.read"))
	audit.Get("", h.ListAuditLogs)
}