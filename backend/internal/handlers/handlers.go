package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sikerma/backend/internal/config"
	"github.com/sikerma/backend/internal/middleware"
	"github.com/sikerma/backend/internal/repositories"
)

// Handlers mengelola semua handlers aplikasi
type Handlers struct {
	dbMaster      *pgxpool.Pool
	dbKepegawaian *pgxpool.Pool
	cfg           *config.Config
	authMiddleware *middleware.AuthMiddleware

	// Repositories
	satkerRepo       *repositories.SatkerRepository
	jabatanRepo      *repositories.JabatanRepository
	golonganRepo     *repositories.GolonganRepository
	unitKerjaRepo    *repositories.UnitKerjaRepository
	eselonRepo       *repositories.EselonRepository
	pegawaiRepo      *repositories.PegawaiRepository
	riwayatRepo      *repositories.RiwayatRepository
	roleRepo         *repositories.RoleRepository
	auditRepo        *repositories.AuditRepository
}

// New membuat instance Handlers baru
func New(dbMaster, dbKepegawaian *pgxpool.Pool, cfg *config.Config) *Handlers {
	authMid := middleware.NewAuthMiddleware(cfg.Keycloak.JWKSURL, cfg.Keycloak.Realm)

	return &Handlers{
		dbMaster:      dbMaster,
		dbKepegawaian: dbKepegawaian,
		cfg:           cfg,
		authMiddleware: authMid,

		// Initialize repositories
		satkerRepo:    repositories.NewSatkerRepository(dbMaster),
		jabatanRepo:   repositories.NewJabatanRepository(dbMaster),
		golonganRepo:  repositories.NewGolonganRepository(dbMaster),
		unitKerjaRepo: repositories.NewUnitKerjaRepository(dbMaster),
		eselonRepo:    repositories.NewEselonRepository(dbMaster),
		pegawaiRepo:   repositories.NewPegawaiRepository(dbKepegawaian),
		riwayatRepo:   repositories.NewRiwayatRepository(dbKepegawaian, dbMaster),
		roleRepo:      repositories.NewRoleRepository(dbMaster),
		auditRepo:     repositories.NewAuditRepository(dbMaster),
	}
}

// ==================== HEALTH CHECK ====================

// HealthCheck mengecek kesehatan sistem
func (h *Handlers) HealthCheck(c fiber.Ctx) error {
	ctx := c.Context()

	// Cek database connection
	masterStatus := "ok"
	if err := h.dbMaster.Ping(ctx); err != nil {
		masterStatus = fmt.Sprintf("error: %v", err)
	}

	kepegawaianStatus := "ok"
	if err := h.dbKepegawaian.Ping(ctx); err != nil {
		kepegawaianStatus = fmt.Sprintf("error: %v", err)
	}

	return c.JSON(fiber.Map{
		"status": "healthy",
		"services": fiber.Map{
			"database_master":     masterStatus,
			"database_kepegawaian": kepegawaianStatus,
			"keycloak":            h.cfg.Keycloak.URL,
		},
		"version": "1.0.0",
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== AUTH ====================

// Login menangani login request
// Untuk sekarang, ini hanya proxy ke Keycloak
func (h *Handlers) Login(c fiber.Ctx) error {
	// TODO: Implement Keycloak direct login grant
	return c.Status(501).JSON(fiber.Map{
		"error": true,
		"message": "Login through Keycloak frontend flow only",
		"code": 501,
		"request_id": middleware.GetRequestID(c),
	})
}

// Logout menangani logout request
func (h *Handlers) Logout(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Logged out successfully",
		"request_id": middleware.GetRequestID(c),
	})
}

// GetCurrentUser mengambil info user saat ini
func (h *Handlers) GetCurrentUser(c fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)
	username := c.Locals("username").(string)
	email := c.Locals("email").(string)
	name := c.Locals("name").(string)
	roles := c.Locals("userRoles").([]string)

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":       userID,
			"username": username,
			"email":    email,
			"name":     name,
			"role":     userRole,
			"roles":    roles,
		},
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== MASTER DATA - SATKER ====================

// ListSatker mengambil daftar satker dengan pagination
func (h *Handlers) ListSatker(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 20)
	search := fiber.Query[string](c, "search", "")
	isActive := c.Query("is_active")

	satkers, total, err := h.satkerRepo.List(c.Context(), page, limit, search, isActive)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    satkers,
		"pagination": fiber.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
		"request_id": middleware.GetRequestID(c),
	})
}

// GetSatker mengambil detail satker
func (h *Handlers) GetSatker(c fiber.Ctx) error {
	id := c.Params("id")

	satker, err := h.satkerRepo.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    satker,
		"request_id": middleware.GetRequestID(c),
	})
}

// CreateSatker membuat satker baru
func (h *Handlers) CreateSatker(c fiber.Ctx) error {
	var input repositories.CreateSatkerInput
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"message": "Invalid request body",
			"code": 400,
			"request_id": middleware.GetRequestID(c),
		})
	}

	satker, err := h.satkerRepo.Create(c.Context(), input, middleware.GetUserID(c))
	if err != nil {
		return err
	}

	// Audit log
	go h.auditRepo.Log(context.Background(), repositories.AuditLogInput{
		UserID:     middleware.GetUserID(c),
		Action:     "create",
		Resource:   "satker",
		ResourceID: &satker.ID,
		Changes:    fiber.Map{"input": input},
		Status:     "success",
	})

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Satker created successfully",
		"data":    satker,
		"request_id": middleware.GetRequestID(c),
	})
}

// UpdateSatker mengupdate satker
func (h *Handlers) UpdateSatker(c fiber.Ctx) error {
	id := c.Params("id")

	var input repositories.UpdateSatkerInput
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"message": "Invalid request body",
			"code": 400,
			"request_id": middleware.GetRequestID(c),
		})
	}

	satker, err := h.satkerRepo.Update(c.Context(), id, input, middleware.GetUserID(c))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Satker updated successfully",
		"data":    satker,
		"request_id": middleware.GetRequestID(c),
	})
}

// DeleteSatker menghapus satker
func (h *Handlers) DeleteSatker(c fiber.Ctx) error {
	id := c.Params("id")

	if err := h.satkerRepo.Delete(c.Context(), id); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Satker deleted successfully",
		"request_id": middleware.GetRequestID(c),
	})
}

// GetDropdownSatker mengambil dropdown data untuk satker
func (h *Handlers) GetDropdownSatker(c fiber.Ctx) error {
	data, err := h.satkerRepo.GetDropdown(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== MASTER DATA - JABATAN ====================

// ListJabatan mengambil daftar jabatan dengan pagination
func (h *Handlers) ListJabatan(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 20)
	search := fiber.Query[string](c, "search", "")

	jabatans, total, err := h.jabatanRepo.List(c.Context(), page, limit, search)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    jabatans,
		"pagination": fiber.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
		"request_id": middleware.GetRequestID(c),
	})
}

// GetDropdownJabatan mengambil dropdown data untuk jabatan
func (h *Handlers) GetDropdownJabatan(c fiber.Ctx) error {
	data, err := h.jabatanRepo.GetDropdown(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== MASTER DATA - GOLONGAN ====================

// ListGolongan mengambil daftar golongan dengan pagination
func (h *Handlers) ListGolongan(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 20)
	search := fiber.Query[string](c, "search", "")

	golongans, total, err := h.golonganRepo.List(c.Context(), page, limit, search)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    golongans,
		"pagination": fiber.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
		"request_id": middleware.GetRequestID(c),
	})
}

// GetDropdownGolongan mengambil dropdown data untuk golongan
func (h *Handlers) GetDropdownGolongan(c fiber.Ctx) error {
	data, err := h.golonganRepo.GetDropdown(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== MASTER DATA - UNIT KERJA ====================

// ListUnitKerja mengambil daftar unit kerja dengan pagination
func (h *Handlers) ListUnitKerja(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 20)
	search := fiber.Query[string](c, "search", "")

	unitKerjas, total, err := h.unitKerjaRepo.List(c.Context(), page, limit, search)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    unitKerjas,
		"pagination": fiber.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
		"request_id": middleware.GetRequestID(c),
	})
}

// GetDropdownUnitKerja mengambil dropdown data untuk unit kerja
func (h *Handlers) GetDropdownUnitKerja(c fiber.Ctx) error {
	data, err := h.unitKerjaRepo.GetDropdown(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== MASTER DATA - ESELON ====================

// ListEselon mengambil daftar eselon dengan pagination
func (h *Handlers) ListEselon(c fiber.Ctx) error {
	eselons, err := h.eselonRepo.List(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    eselons,
		"request_id": middleware.GetRequestID(c),
	})
}

// GetDropdownEselon mengambil dropdown data untuk eselon
func (h *Handlers) GetDropdownEselon(c fiber.Ctx) error {
	data, err := h.eselonRepo.GetDropdown(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== KEGAWAAN - PEGAWAI ====================

// ListPegawai mengambil daftar pegawai dengan pagination dan filter
func (h *Handlers) ListPegawai(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 20)
	search := fiber.Query[string](c, "search", "")
	satkerID := fiber.Query[string](c, "satker_id", "")
	jabatanID := fiber.Query[string](c, "jabatan_id", "")
	golonganID := fiber.Query[string](c, "golongan_id", "")
	statusPegawai := fiber.Query[string](c, "status_pegawai", "")

	pegawais, total, err := h.pegawaiRepo.List(c.Context(), page, limit, search, satkerID, jabatanID, golonganID, statusPegawai)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    pegawais,
		"pagination": fiber.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
		"request_id": middleware.GetRequestID(c),
	})
}

// GetPegawai mengambil detail pegawai
func (h *Handlers) GetPegawai(c fiber.Ctx) error {
	id := c.Params("id")

	pegawai, err := h.pegawaiRepo.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    pegawai,
		"request_id": middleware.GetRequestID(c),
	})
}

// CreatePegawai membuat pegawai baru
func (h *Handlers) CreatePegawai(c fiber.Ctx) error {
	var input repositories.CreatePegawaiInput
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"message": "Invalid request body",
			"code": 400,
			"request_id": middleware.GetRequestID(c),
		})
	}

	pegawai, err := h.pegawaiRepo.Create(c.Context(), input)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Pegawai created successfully",
		"data":    pegawai,
		"request_id": middleware.GetRequestID(c),
	})
}

// UpdatePegawai mengupdate pegawai
func (h *Handlers) UpdatePegawai(c fiber.Ctx) error {
	id := c.Params("id")

	var input repositories.UpdatePegawaiInput
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"message": "Invalid request body",
			"code": 400,
			"request_id": middleware.GetRequestID(c),
		})
	}

	pegawai, err := h.pegawaiRepo.Update(c.Context(), id, input)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Pegawai updated successfully",
		"data":    pegawai,
		"request_id": middleware.GetRequestID(c),
	})
}

// DeletePegawai menghapus pegawai
func (h *Handlers) DeletePegawai(c fiber.Ctx) error {
	id := c.Params("id")

	if err := h.pegawaiRepo.Delete(c.Context(), id); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Pegawai deleted successfully",
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== STATISTIK ====================

// GetStatistikKepegawaian mengambil statistik kepegawaian
func (h *Handlers) GetStatistikKepegawaian(c fiber.Ctx) error {
	statistik, err := h.pegawaiRepo.GetStatistik(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    statistik,
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== RBAC ====================

// ListRoles mengambil daftar roles
func (h *Handlers) ListRoles(c fiber.Ctx) error {
	roles, err := h.roleRepo.List(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    roles,
		"request_id": middleware.GetRequestID(c),
	})
}

// CreateRole membuat role baru
func (h *Handlers) CreateRole(c fiber.Ctx) error {
	var input repositories.CreateRoleInput
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"message": "Invalid request body",
			"code": 400,
			"request_id": middleware.GetRequestID(c),
		})
	}

	role, err := h.roleRepo.Create(c.Context(), input)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Role created successfully",
		"data":    role,
		"request_id": middleware.GetRequestID(c),
	})
}

// ==================== AUDIT LOGS ====================

// ListAuditLogs mengambil daftar audit logs
func (h *Handlers) ListAuditLogs(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 50)
	action := fiber.Query[string](c, "action", "")
	resource := fiber.Query[string](c, "resource", "")
	userID := fiber.Query[string](c, "user_id", "")

	logs, total, err := h.auditRepo.List(c.Context(), page, limit, action, resource, userID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    logs,
		"pagination": fiber.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
		"request_id": middleware.GetRequestID(c),
	})
}
