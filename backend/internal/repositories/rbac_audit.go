package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sikerma/backend/internal/models"
)

// ==================== RIWAYAT ====================

// RiwayatRepository mengelola operasi database untuk riwayat
type RiwayatRepository struct {
	dbKepegawaian *pgxpool.Pool
	dbMaster      *pgxpool.Pool
}

// NewRiwayatRepository membuat instance RiwayatRepository baru
func NewRiwayatRepository(dbKepegawaian, dbMaster *pgxpool.Pool) *RiwayatRepository {
	return &RiwayatRepository{
		dbKepegawaian: dbKepegawaian,
		dbMaster:      dbMaster,
	}
}

// ==================== RBAC ====================

// RoleRepository mengelola operasi database untuk Role
type RoleRepository struct {
	db *pgxpool.Pool
}

// NewRoleRepository membuat instance RoleRepository baru
func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{db: db}
}

// List mengambil daftar roles
func (r *RoleRepository) List(ctx context.Context) ([]models.AppRole, error) {
	query := `SELECT id, nama, deskripsi, is_active, created_at, updated_at
			  FROM app_roles
			  WHERE is_active = true
			  ORDER BY nama`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query roles: %w", err)
	}
	defer rows.Close()

	roles := []models.AppRole{}
	for rows.Next() {
		var role models.AppRole
		err := rows.Scan(
			&role.ID, &role.Nama, &role.Deskripsi, &role.IsActive,
			&role.CreatedAt, &role.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan role: %w", err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// Create membuat role baru
func (r *RoleRepository) Create(ctx context.Context, input CreateRoleInput) (*models.AppRole, error) {
	id := uuid.New()

	query := `INSERT INTO app_roles (id, nama, deskripsi, is_active)
			  VALUES ($1, $2, $3, true)
			  RETURNING created_at, updated_at`

	err := r.db.QueryRow(ctx, query, id, input.Nama, input.Deskripsi).Scan(&input.CreatedAt, &input.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	role := &models.AppRole{
		ID:        id,
		Nama:      input.Nama,
		Deskripsi: input.Deskripsi,
		IsActive:  true,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}

	return role, nil
}

// ==================== AUDIT ====================

// AuditRepository mengelola operasi database untuk Audit Log
type AuditRepository struct {
	db *pgxpool.Pool
}

// NewAuditRepository membuat instance AuditRepository baru
func NewAuditRepository(db *pgxpool.Pool) *AuditRepository {
	return &AuditRepository{db: db}
}

// Log menyimpan audit log baru
func (r *AuditRepository) Log(ctx context.Context, input AuditLogInput) error {
	id := uuid.New()
	var userID *string
	if input.UserID != "" {
		userID = &input.UserID
	}

	query := `INSERT INTO audit_logs (id, user_id, username, action, resource, resource_id,
			  ip_address, user_agent, changes, status, error_message)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.Exec(ctx, query,
		id, userID, input.Username, input.Action, input.Resource,
		input.ResourceID, input.IPAddress, input.UserAgent,
		input.Changes, input.Status, input.ErrorMessage,
	)

	return err
}

// List mengambil daftar audit logs dengan filter
func (r *AuditRepository) List(ctx context.Context, page, limit int, action, resource, userID string) ([]models.AuditLog, int64, error) {
	offset := (page - 1) * limit

	query := `SELECT id, user_id, username, action, resource, resource_id,
			  ip_address, user_agent, changes, status, error_message, created_at
			  FROM audit_logs
			  WHERE 1=1`
	args := []interface{}{}
	argCount := 1

	if action != "" {
		query += fmt.Sprintf(" AND action = $%d", argCount)
		args = append(args, action)
		argCount++
	}

	if resource != "" {
		query += fmt.Sprintf(" AND resource = $%d", argCount)
		args = append(args, resource)
		argCount++
	}

	if userID != "" {
		query += fmt.Sprintf(" AND user_id = $%d", argCount)
		args = append(args, userID)
		argCount++
	}

	// Get total count
	countQuery := "SELECT COUNT(*) FROM audit_logs WHERE 1=1"
	countArgs := []interface{}{}
	countArgCount := 1

	if action != "" {
		countQuery += fmt.Sprintf(" AND action = $%d", countArgCount)
		countArgs = append(countArgs, action)
		countArgCount++
	}

	if resource != "" {
		countQuery += fmt.Sprintf(" AND resource = $%d", countArgCount)
		countArgs = append(countArgs, resource)
		countArgCount++
	}

	if userID != "" {
		countQuery += fmt.Sprintf(" AND user_id = $%d", countArgCount)
		countArgs = append(countArgs, userID)
		countArgCount++
	}

	var total int64
	err := r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count audit logs: %w", err)
	}

	// Get data
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query audit logs: %w", err)
	}
	defer rows.Close()

	logs := []models.AuditLog{}
	for rows.Next() {
		var log models.AuditLog
		err := rows.Scan(
			&log.ID, &log.UserID, &log.Username, &log.Action, &log.Resource,
			&log.ResourceID, &log.IPAddress, &log.UserAgent,
			&log.Changes, &log.Status, &log.ErrorMessage, &log.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan audit log: %w", err)
		}
		logs = append(logs, log)
	}

	return logs, total, nil
}

// ==================== INPUT TYPES ====================

// CreateRoleInput input untuk membuat role
type CreateRoleInput struct {
	Nama      string    `json:"nama"`
	Deskripsi string    `json:"deskripsi"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// AuditLogInput input untuk audit log
type AuditLogInput struct {
	UserID       string                 `json:"user_id"`
	Username     string                 `json:"username"`
	Action       string                 `json:"action"`
	Resource     string                 `json:"resource"`
	ResourceID   *uuid.UUID             `json:"resource_id,omitempty"`
	IPAddress    *string                `json:"ip_address,omitempty"`
	UserAgent    *string                `json:"user_agent,omitempty"`
	Changes      map[string]interface{} `json:"changes,omitempty"`
	Status       string                 `json:"status"`
	ErrorMessage *string                `json:"error_message,omitempty"`
}