package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sikerma/backend/internal/models"
)

// ==================== JABATAN ====================

// JabatanRepository mengelola operasi database untuk Jabatan
type JabatanRepository struct {
	db *pgxpool.Pool
}

// NewJabatanRepository membuat instance JabatanRepository baru
func NewJabatanRepository(db *pgxpool.Pool) *JabatanRepository {
	return &JabatanRepository{db: db}
}

// List mengambil daftar jabatan dengan pagination
func (r *JabatanRepository) List(ctx context.Context, page, limit int, search string) ([]models.Jabatan, int64, error) {
	offset := (page - 1) * limit

	query := `SELECT id, kode, nama, eselon_id, kelas, is_active, created_at, updated_at
			  FROM jabatan
			  WHERE is_active = true`
	args := []interface{}{}
	argCount := 1

	if search != "" {
		query += fmt.Sprintf(" AND (kode ILIKE $%d OR nama ILIKE $%d)", argCount, argCount+1)
		args = append(args, "%"+search+"%", "%"+search+"%")
		argCount += 2
	}

	// Get total count
	countQuery := "SELECT COUNT(*) FROM jabatan WHERE is_active = true"
	if search != "" {
		countQuery += fmt.Sprintf(" AND (kode ILIKE $%d OR nama ILIKE $%d)", 1, 2)
	}
	var total int64
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count jabatan: %w", err)
	}

	// Get data
	query += fmt.Sprintf(" ORDER BY kode LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query jabatan: %w", err)
	}
	defer rows.Close()

	jabatans := []models.Jabatan{}
	for rows.Next() {
		var jabatan models.Jabatan
		err := rows.Scan(
			&jabatan.ID, &jabatan.Kode, &jabatan.Nama, &jabatan.EselonID,
			&jabatan.Kelas, &jabatan.IsActive, &jabatan.CreatedAt, &jabatan.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan jabatan: %w", err)
		}
		jabatans = append(jabatans, jabatan)
	}

	return jabatans, total, nil
}

// GetDropdown mengambil data dropdown
func (r *JabatanRepository) GetDropdown(ctx context.Context) ([]DropdownItem, error) {
	query := `SELECT id, kode, nama FROM jabatan WHERE is_active = true ORDER BY kode`

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

// ==================== GOLONGAN ====================

// GolonganRepository mengelola operasi database untuk Golongan
type GolonganRepository struct {
	db *pgxpool.Pool
}

// NewGolonganRepository membuat instance GolonganRepository baru
func NewGolonganRepository(db *pgxpool.Pool) *GolonganRepository {
	return &GolonganRepository{db: db}
}

// List mengambil daftar golongan dengan pagination
func (r *GolonganRepository) List(ctx context.Context, page, limit int, search string) ([]models.Golongan, int64, error) {
	offset := (page - 1) * limit

	query := `SELECT id, kode, nama, ruang, angka, min_pangkat, max_pangkat, is_active, created_at, updated_at
			  FROM golongan
			  WHERE is_active = true`
	args := []interface{}{}
	argCount := 1

	if search != "" {
		query += fmt.Sprintf(" AND (kode ILIKE $%d OR nama ILIKE $%d)", argCount, argCount+1)
		args = append(args, "%"+search+"%", "%"+search+"%")
		argCount += 2
	}

	// Get total count
	countQuery := "SELECT COUNT(*) FROM golongan WHERE is_active = true"
	if search != "" {
		countQuery += fmt.Sprintf(" AND (kode ILIKE $%d OR nama ILIKE $%d)", 1, 2)
	}
	var total int64
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count golongan: %w", err)
	}

	// Get data
	query += fmt.Sprintf(" ORDER BY angka DESC LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query golongan: %w", err)
	}
	defer rows.Close()

	golongans := []models.Golongan{}
	for rows.Next() {
		var golongan models.Golongan
		err := rows.Scan(
			&golongan.ID, &golongan.Kode, &golongan.Nama, &golongan.Ruang,
			&golongan.Angka, &golongan.MinPangkat, &golongan.MaxPangkat,
			&golongan.IsActive, &golongan.CreatedAt, &golongan.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan golongan: %w", err)
		}
		golongans = append(golongans, golongan)
	}

	return golongans, total, nil
}

// GetDropdown mengambil data dropdown
func (r *GolonganRepository) GetDropdown(ctx context.Context) ([]DropdownItem, error) {
	query := `SELECT id, kode, nama FROM golongan WHERE is_active = true ORDER BY angka DESC`

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

// ==================== UNIT KERJA ====================

// UnitKerjaRepository mengelola operasi database untuk UnitKerja
type UnitKerjaRepository struct {
	db *pgxpool.Pool
}

// NewUnitKerjaRepository membuat instance UnitKerjaRepository baru
func NewUnitKerjaRepository(db *pgxpool.Pool) *UnitKerjaRepository {
	return &UnitKerjaRepository{db: db}
}

// List mengambil daftar unit kerja dengan pagination
func (r *UnitKerjaRepository) List(ctx context.Context, page, limit int, search string) ([]models.UnitKerja, int64, error) {
	offset := (page - 1) * limit

	query := `SELECT id, kode, nama, singkatan, parent_id, is_active, created_at, updated_at
			  FROM unit_kerja
			  WHERE is_active = true`
	args := []interface{}{}
	argCount := 1

	if search != "" {
		query += fmt.Sprintf(" AND (kode ILIKE $%d OR nama ILIKE $%d)", argCount, argCount+1)
		args = append(args, "%"+search+"%", "%"+search+"%")
		argCount += 2
	}

	// Get total count
	countQuery := "SELECT COUNT(*) FROM unit_kerja WHERE is_active = true"
	if search != "" {
		countQuery += fmt.Sprintf(" AND (kode ILIKE $%d OR nama ILIKE $%d)", 1, 2)
	}
	var total int64
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count unit_kerja: %w", err)
	}

	// Get data
	query += fmt.Sprintf(" ORDER BY kode LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query unit_kerja: %w", err)
	}
	defer rows.Close()

	unitKerjas := []models.UnitKerja{}
	for rows.Next() {
		var unitKerja models.UnitKerja
		err := rows.Scan(
			&unitKerja.ID, &unitKerja.Kode, &unitKerja.Nama, &unitKerja.Singkatan,
			&unitKerja.ParentID, &unitKerja.IsActive, &unitKerja.CreatedAt, &unitKerja.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan unit_kerja: %w", err)
		}
		unitKerjas = append(unitKerjas, unitKerja)
	}

	return unitKerjas, total, nil
}

// GetDropdown mengambil data dropdown
func (r *UnitKerjaRepository) GetDropdown(ctx context.Context) ([]DropdownItem, error) {
	query := `SELECT id, kode, nama FROM unit_kerja WHERE is_active = true ORDER BY kode`

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

// ==================== ESELON ====================

// EselonRepository mengelola operasi database untuk Eselon
type EselonRepository struct {
	db *pgxpool.Pool
}

// NewEselonRepository membuat instance EselonRepository baru
func NewEselonRepository(db *pgxpool.Pool) *EselonRepository {
	return &EselonRepository{db: db}
}

// List mengambil daftar eselon
func (r *EselonRepository) List(ctx context.Context) ([]models.Eselon, error) {
	query := `SELECT id, kode, nama, tunjangan, is_active, created_at, updated_at
			  FROM eselon
			  WHERE is_active = true
			  ORDER BY kode`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query eselon: %w", err)
	}
	defer rows.Close()

	eselons := []models.Eselon{}
	for rows.Next() {
		var eselon models.Eselon
		err := rows.Scan(
			&eselon.ID, &eselon.Kode, &eselon.Nama, &eselon.Tunjangan,
			&eselon.IsActive, &eselon.CreatedAt, &eselon.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan eselon: %w", err)
		}
		eselons = append(eselons, eselon)
	}

	return eselons, nil
}

// GetDropdown mengambil data dropdown
func (r *EselonRepository) GetDropdown(ctx context.Context) ([]DropdownItem, error) {
	query := `SELECT id, kode, nama FROM eselon WHERE is_active = true ORDER BY kode`

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