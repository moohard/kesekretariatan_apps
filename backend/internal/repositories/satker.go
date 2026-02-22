package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sikerma/backend/internal/models"
)

// SatkerRepository mengelola operasi database untuk Satker
type SatkerRepository struct {
	db *pgxpool.Pool
}

// NewSatkerRepository membuat instance SatkerRepository baru
func NewSatkerRepository(db *pgxpool.Pool) *SatkerRepository {
	return &SatkerRepository{db: db}
}

// List mengambil daftar satker dengan pagination
func (r *SatkerRepository) List(ctx context.Context, page, limit int, search, isActive string) ([]models.Satker, int64, error) {
	offset := (page - 1) * limit

	query := `SELECT id, kode, nama, parent_id, level, alamat, telepon, email, is_active, created_at, updated_at, created_by, updated_by
			  FROM satker
			  WHERE 1=1`
	args := []interface{}{}
	argCount := 1

	if search != "" {
		query += fmt.Sprintf(" AND (kode ILIKE $%d OR nama ILIKE $%d)", argCount, argCount+1)
		args = append(args, "%"+search+"%", "%"+search+"%")
		argCount += 2
	}

	if isActive != "" {
		query += fmt.Sprintf(" AND is_active = $%d", argCount)
		args = append(args, isActive == "true")
		argCount++
	}

	// Get total count
	countQuery := "SELECT COUNT(*) FROM satker" + query[20:] // Remove SELECT columns
	var total int64
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count satker: %w", err)
	}

	// Get data
	query += fmt.Sprintf(" ORDER BY kode LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query satker: %w", err)
	}
	defer rows.Close()

	satkers := []models.Satker{}
	for rows.Next() {
		var satker models.Satker
		err := rows.Scan(
			&satker.ID, &satker.Kode, &satker.Nama, &satker.ParentID, &satker.Level,
			&satker.Alamat, &satker.Telepon, &satker.Email, &satker.IsActive,
			&satker.CreatedAt, &satker.UpdatedAt, &satker.CreatedBy, &satker.UpdatedBy,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan satker: %w", err)
		}
		satkers = append(satkers, satker)
	}

	return satkers, total, nil
}

// GetByID mengambil satker berdasarkan ID
func (r *SatkerRepository) GetByID(ctx context.Context, id string) (*models.Satker, error) {
	query := `SELECT id, kode, nama, parent_id, level, alamat, telepon, email, is_active, created_at, updated_at, created_by, updated_by
			  FROM satker WHERE id = $1`

	var satker models.Satker
	err := r.db.QueryRow(ctx, query, uuid.MustParse(id)).Scan(
		&satker.ID, &satker.Kode, &satker.Nama, &satker.ParentID, &satker.Level,
		&satker.Alamat, &satker.Telepon, &satker.Email, &satker.IsActive,
		&satker.CreatedAt, &satker.UpdatedAt, &satker.CreatedBy, &satker.UpdatedBy,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("satker not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get satker: %w", err)
	}

	return &satker, nil
}

// Create membuat satker baru
func (r *SatkerRepository) Create(ctx context.Context, input CreateSatkerInput, userID string) (*models.Satker, error) {
	id := uuid.New()

	query := `INSERT INTO satker (id, kode, nama, parent_id, level, alamat, telepon, email, is_active, created_by, updated_by)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $10)
			  RETURNING created_at, updated_at`

	err := r.db.QueryRow(ctx, query,
		id, input.Kode, input.Nama, input.ParentID, input.Level,
		input.Alamat, input.Telepon, input.Email, true,
	).Scan(&input.CreatedAt, &input.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create satker: %w", err)
	}

	satker := &models.Satker{
		ID:        id,
		Kode:      input.Kode,
		Nama:      input.Nama,
		ParentID:  input.ParentID,
		Level:     input.Level,
		Alamat:    input.Alamat,
		Telepon:   input.Telepon,
		Email:     input.Email,
		IsActive:  true,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}

	return satker, nil
}

// Update mengupdate satker
func (r *SatkerRepository) Update(ctx context.Context, id string, input UpdateSatkerInput, userID string) (*models.Satker, error) {
	query := `UPDATE satker
			  SET kode = $2, nama = $3, parent_id = $4, level = $5,
				  alamat = $6, telepon = $7, email = $8, is_active = $9,
				  updated_at = NOW()
			  WHERE id = $1
			  RETURNING created_at, updated_at, created_by, updated_by`

	var satker models.Satker
	err := r.db.QueryRow(ctx, query,
		uuid.MustParse(id), input.Kode, input.Nama, input.ParentID, input.Level,
		input.Alamat, input.Telepon, input.Email, input.IsActive,
	).Scan(&satker.CreatedAt, &satker.UpdatedAt, &satker.CreatedBy, &satker.UpdatedBy)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("satker not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update satker: %w", err)
	}

	satker.ID = uuid.MustParse(id)
	satker.Kode = input.Kode
	satker.Nama = input.Nama
	satker.ParentID = input.ParentID
	satker.Level = input.Level
	satker.Alamat = input.Alamat
	satker.Telepon = input.Telepon
	satker.Email = input.Email
	satker.IsActive = input.IsActive

	return &satker, nil
}

// Delete menghapus satker
func (r *SatkerRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM satker WHERE id = $1`

	result, err := r.db.Exec(ctx, query, uuid.MustParse(id))
	if err != nil {
		return fmt.Errorf("failed to delete satker: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("satker not found")
	}

	return nil
}

// GetDropdown mengambil data dropdown
func (r *SatkerRepository) GetDropdown(ctx context.Context) ([]DropdownItem, error) {
	query := `SELECT id, kode, nama FROM satker WHERE is_active = true ORDER BY kode`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query dropdown: %w", err)
	}
	defer rows.Close()

	items := []DropdownItem{}
	for rows.Next() {
		var item DropdownItem
		err := rows.Scan(&item.Value, &item.Label, &item.Label)
		if err != nil {
			return nil, fmt.Errorf("failed to scan dropdown: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

// ==================== INPUT TYPES ====================

// CreateSatkerInput input untuk membuat satker
type CreateSatkerInput struct {
	Kode      string     `json:"kode"`
	Nama      string     `json:"nama"`
	ParentID  *uuid.UUID `json:"parent_id,omitempty"`
	Level     int        `json:"level"`
	Alamat    string     `json:"alamat"`
	Telepon   string     `json:"telepon"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}

// UpdateSatkerInput input untuk update satker
type UpdateSatkerInput struct {
	Kode     string     `json:"kode"`
	Nama     string     `json:"nama"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"`
	Level    int        `json:"level"`
	Alamat   string     `json:"alamat"`
	Telepon  string     `json:"telepon"`
	Email    string     `json:"email"`
	IsActive bool       `json:"is_active"`
}

// DropdownItem item untuk dropdown
type DropdownItem struct {
	Value uuid.UUID `json:"value"`
	Label string    `json:"label"`
}