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

// ==================== PEGAWAI ====================

// PegawaiRepository mengelola operasi database untuk Pegawai
type PegawaiRepository struct {
	db *pgxpool.Pool
}

// NewPegawaiRepository membuat instance PegawaiRepository baru
func NewPegawaiRepository(db *pgxpool.Pool) *PegawaiRepository {
	return &PegawaiRepository{db: db}
}

// List mengambil daftar pegawai dengan pagination dan filter
func (r *PegawaiRepository) List(ctx context.Context, page, limit int, search, satkerID, jabatanID, golonganID, statusPegawai string) ([]models.Pegawai, int64, error) {
	offset := (page - 1) * limit

	query := `SELECT p.id, p.nip, p.nama, p.gelar_depan, p.gelar_belakang,
			  p.tempat_lahir, p.tanggal_lahir, p.jenis_kelamin,
			  p.agama_id, p.status_kawin_id, p.nik, p.email, p.telepon,
			  p.alamat, p.foto, p.satker_id, p.jabatan_id, p.unit_kerja_id,
			  p.golongan_id, p.status_pegawai, p.tmt_jabatan, p.is_pns, p.is_active,
			  p.created_at, p.updated_at
			  FROM pegawai p
			  WHERE p.is_active = true`
	args := []interface{}{}
	argCount := 1

	if search != "" {
		query += fmt.Sprintf(" AND (p.nip ILIKE $%d OR p.nama ILIKE $%d)", argCount, argCount+1)
		args = append(args, "%"+search+"%", "%"+search+"%")
		argCount += 2
	}

	if satkerID != "" {
		query += fmt.Sprintf(" AND p.satker_id = $%d", argCount)
		args = append(args, uuid.MustParse(satkerID))
		argCount++
	}

	if jabatanID != "" {
		query += fmt.Sprintf(" AND p.jabatan_id = $%d", argCount)
		args = append(args, uuid.MustParse(jabatanID))
		argCount++
	}

	if golonganID != "" {
		query += fmt.Sprintf(" AND p.golongan_id = $%d", argCount)
		args = append(args, uuid.MustParse(golonganID))
		argCount++
	}

	if statusPegawai != "" {
		query += fmt.Sprintf(" AND p.status_pegawai = $%d", argCount)
		args = append(args, statusPegawai)
		argCount++
	}

	// Get total count
	countQuery := "SELECT COUNT(*) FROM pegawai p WHERE p.is_active = true"
	countArgs := []interface{}{}
	countArgCount := 1

	if search != "" {
		countQuery += fmt.Sprintf(" AND (p.nip ILIKE $%d OR p.nama ILIKE $%d)", countArgCount, countArgCount+1)
		countArgs = append(countArgs, "%"+search+"%", "%"+search+"%")
		countArgCount += 2
	}

	if satkerID != "" {
		countQuery += fmt.Sprintf(" AND p.satker_id = $%d", countArgCount)
		countArgs = append(countArgs, uuid.MustParse(satkerID))
		countArgCount++
	}

	if jabatanID != "" {
		countQuery += fmt.Sprintf(" AND p.jabatan_id = $%d", countArgCount)
		countArgs = append(countArgs, uuid.MustParse(jabatanID))
		countArgCount++
	}

	if golonganID != "" {
		countQuery += fmt.Sprintf(" AND p.golongan_id = $%d", countArgCount)
		countArgs = append(countArgs, uuid.MustParse(golonganID))
		countArgCount++
	}

	if statusPegawai != "" {
		countQuery += fmt.Sprintf(" AND p.status_pegawai = $%d", countArgCount)
		countArgs = append(countArgs, statusPegawai)
		countArgCount++
	}

	var total int64
	err := r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count pegawai: %w", err)
	}

	// Get data
	query += fmt.Sprintf(" ORDER BY p.nama LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query pegawai: %w", err)
	}
	defer rows.Close()

	pegawais := []models.Pegawai{}
	for rows.Next() {
		var pegawai models.Pegawai
		err := rows.Scan(
			&pegawai.ID, &pegawai.NIP, &pegawai.Nama, &pegawai.GelarDepan, &pegawai.GelarBelakang,
			&pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.JenisKelamin,
			&pegawai.AgamaID, &pegawai.StatusKawinID, &pegawai.NIK, &pegawai.Email, &pegawai.Telepon,
			&pegawai.Alamat, &pegawai.Foto, &pegawai.SatkerID, &pegawai.JabatanID, &pegawai.UnitKerjaID,
			&pegawai.GolonganID, &pegawai.StatusPegawai, &pegawai.TMTJabatan, &pegawai.IsPNS, &pegawai.IsActive,
			&pegawai.CreatedAt, &pegawai.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan pegawai: %w", err)
		}
		pegawais = append(pegawais, pegawai)
	}

	return pegawais, total, nil
}

// GetByID mengambil detail pegawai dengan relasi
func (r *PegawaiRepository) GetByID(ctx context.Context, id string) (*models.Pegawai, error) {
	query := `SELECT p.id, p.nip, p.nama, p.gelar_depan, p.gelar_belakang,
			  p.tempat_lahir, p.tanggal_lahir, p.jenis_kelamin,
			  p.agama_id, p.status_kawin_id, p.nik, p.email, p.telepon,
			  p.alamat, p.foto, p.satker_id, p.jabatan_id, p.unit_kerja_id,
			  p.golongan_id, p.status_pegawai, p.tmt_jabatan, p.is_pns, p.is_active,
			  p.created_at, p.updated_at
			  FROM pegawai p
			  WHERE p.id = $1`

	var pegawai models.Pegawai
	err := r.db.QueryRow(ctx, query, uuid.MustParse(id)).Scan(
		&pegawai.ID, &pegawai.NIP, &pegawai.Nama, &pegawai.GelarDepan, &pegawai.GelarBelakang,
		&pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.JenisKelamin,
		&pegawai.AgamaID, &pegawai.StatusKawinID, &pegawai.NIK, &pegawai.Email, &pegawai.Telepon,
		&pegawai.Alamat, &pegawai.Foto, &pegawai.SatkerID, &pegawai.JabatanID, &pegawai.UnitKerjaID,
		&pegawai.GolonganID, &pegawai.StatusPegawai, &pegawai.TMTJabatan, &pegawai.IsPNS, &pegawai.IsActive,
		&pegawai.CreatedAt, &pegawai.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("pegawai not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get pegawai: %w", err)
	}

	return &pegawai, nil
}

// Create membuat pegawai baru
func (r *PegawaiRepository) Create(ctx context.Context, input CreatePegawaiInput) (*models.Pegawai, error) {
	id := uuid.New()

	query := `INSERT INTO pegawai (id, nip, nama, gelar_depan, gelar_belakang,
			  tempat_lahir, tanggal_lahir, jenis_kelamin, agama_id, status_kawin_id,
			  nik, email, telepon, alamat, satker_id, jabatan_id, unit_kerja_id,
			  golongan_id, status_pegawai, tmt_jabatan, is_pns, is_active)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
			  RETURNING created_at, updated_at`

	err := r.db.QueryRow(ctx, query,
		id, input.NIP, input.Nama, input.GelarDepan, input.GelarBelakang,
		input.TempatLahir, input.TanggalLahir, input.JenisKelamin,
		input.AgamaID, input.StatusKawinID, input.NIK, input.Email,
		input.Telepon, input.Alamat, input.SatkerID, input.JabatanID,
		input.UnitKerjaID, input.GolonganID, input.StatusPegawai,
		input.TMTJabatan, input.IsPNS, true,
	).Scan(&input.CreatedAt, &input.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create pegawai: %w", err)
	}

	pegawai := &models.Pegawai{
		ID:            id,
		NIP:           input.NIP,
		Nama:          input.Nama,
		GelarDepan:    input.GelarDepan,
		GelarBelakang: input.GelarBelakang,
		TempatLahir:   input.TempatLahir,
		TanggalLahir:  input.TanggalLahir,
		JenisKelamin:  input.JenisKelamin,
		AgamaID:       input.AgamaID,
		StatusKawinID: input.StatusKawinID,
		NIK:           input.NIK,
		Email:         input.Email,
		Telepon:       input.Telepon,
		Alamat:        input.Alamat,
		SatkerID:      input.SatkerID,
		JabatanID:     input.JabatanID,
		UnitKerjaID:   input.UnitKerjaID,
		GolonganID:    input.GolonganID,
		StatusPegawai: input.StatusPegawai,
		TMTJabatan:    input.TMTJabatan,
		IsPNS:         input.IsPNS,
		IsActive:      true,
		CreatedAt:     input.CreatedAt,
		UpdatedAt:     input.UpdatedAt,
	}

	return pegawai, nil
}

// Update mengupdate pegawai
func (r *PegawaiRepository) Update(ctx context.Context, id string, input UpdatePegawaiInput) (*models.Pegawai, error) {
	query := `UPDATE pegawai
			  SET nama = $2, gelar_depan = $3, gelar_belakang = $4,
				  email = $5, telepon = $6, alamat = $7,
				  satker_id = $8, jabatan_id = $9, unit_kerja_id = $10,
				  golongan_id = $11, status_pegawai = $12, tmt_jabatan = $13,
				  is_pns = $14, is_active = $15, updated_at = NOW()
			  WHERE id = $1
			  RETURNING nip, tempat_lahir, tanggal_lahir, jenis_kelamin,
			  agama_id, status_kawin_id, nik, foto, created_at, updated_at`

	var pegawai models.Pegawai
	err := r.db.QueryRow(ctx, query,
		uuid.MustParse(id), input.Nama, input.GelarDepan, input.GelarBelakang,
		input.Email, input.Telepon, input.Alamat, input.SatkerID,
		input.JabatanID, input.UnitKerjaID, input.GolonganID,
		input.StatusPegawai, input.TMTJabatan, input.IsPNS, input.IsActive,
	).Scan(
		&pegawai.NIP, &pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.JenisKelamin,
		&pegawai.AgamaID, &pegawai.StatusKawinID, &pegawai.NIK, &pegawai.Foto,
		&pegawai.CreatedAt, &pegawai.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("pegawai not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update pegawai: %w", err)
	}

	pegawai.ID = uuid.MustParse(id)
	pegawai.Nama = input.Nama
	pegawai.GelarDepan = input.GelarDepan
	pegawai.GelarBelakang = input.GelarBelakang
	pegawai.Email = input.Email
	pegawai.Telepon = input.Telepon
	pegawai.Alamat = input.Alamat
	pegawai.SatkerID = input.SatkerID
	pegawai.JabatanID = input.JabatanID
	pegawai.UnitKerjaID = input.UnitKerjaID
	pegawai.GolonganID = input.GolonganID
	pegawai.StatusPegawai = input.StatusPegawai
	pegawai.TMTJabatan = input.TMTJabatan
	pegawai.IsPNS = input.IsPNS
	pegawai.IsActive = input.IsActive

	return &pegawai, nil
}

// Delete menghapus pegawai (soft delete)
func (r *PegawaiRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE pegawai SET is_active = false, updated_at = NOW() WHERE id = $1`

	result, err := r.db.Exec(ctx, query, uuid.MustParse(id))
	if err != nil {
		return fmt.Errorf("failed to delete pegawai: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("pegawai not found")
	}

	return nil
}

// GetStatistik mengambil statistik kepegawaian
func (r *PegawaiRepository) GetStatistik(ctx context.Context) (map[string]interface{}, error) {
	statistik := make(map[string]interface{})

	// Total pegawai aktif
	var totalPegawai int64
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM pegawai WHERE is_active = true").Scan(&totalPegawai)
	if err != nil {
		return nil, fmt.Errorf("failed to count total pegawai: %w", err)
	}
	statistik["total_pegawai"] = totalPegawai

	// Pegawai per status
	statusQuery := `SELECT status_pegawai, COUNT(*) FROM pegawai WHERE is_active = true GROUP BY status_pegawai`
	rows, err := r.db.Query(ctx, statusQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query pegawai by status: %w", err)
	}
	defer rows.Close()

	statusData := make(map[string]int64)
	for rows.Next() {
		var status string
		var count int64
		err := rows.Scan(&status, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan status: %w", err)
		}
		statusData[status] = count
	}
	statistik["per_status"] = statusData

	// Pegawai PNS vs Non-PNS
	var totalPNS, totalNonPNS int64
	r.db.QueryRow(ctx, "SELECT COUNT(*) FROM pegawai WHERE is_pns = true AND is_active = true").Scan(&totalPNS)
	r.db.QueryRow(ctx, "SELECT COUNT(*) FROM pegawai WHERE is_pns = false AND is_active = true").Scan(&totalNonPNS)
	statistik["pns"] = totalPNS
	statistik["non_pns"] = totalNonPNS

	// Pegawai per golongan
	golonganQuery := `SELECT g.nama, COUNT(*) FROM pegawai p
					  JOIN golongan g ON p.golongan_id = g.id
					  WHERE p.is_active = true AND p.golongan_id IS NOT NULL
					  GROUP BY g.nama ORDER BY g.angka DESC`
	rows, err = r.db.Query(ctx, golonganQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query pegawai by golongan: %w", err)
	}
	defer rows.Close()

	golonganData := make(map[string]int64)
	for rows.Next() {
		var nama string
		var count int64
		err := rows.Scan(&nama, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan golongan: %w", err)
		}
		golonganData[nama] = count
	}
	statistik["per_golongan"] = golonganData

	return statistik, nil
}

// ==================== INPUT TYPES ====================

// CreatePegawaiInput input untuk membuat pegawai
type CreatePegawaiInput struct {
	NIP            string    `json:"nip"`
	Nama           string    `json:"nama"`
	GelarDepan     string    `json:"gelar_depan"`
	GelarBelakang  string    `json:"gelar_belakang"`
	TempatLahir    string    `json:"tempat_lahir"`
	TanggalLahir   time.Time `json:"tanggal_lahir"`
	JenisKelamin   string    `json:"jenis_kelamin"`
	AgamaID        uuid.UUID `json:"agama_id"`
	StatusKawinID  uuid.UUID `json:"status_kawin_id"`
	NIK            string    `json:"nik"`
	Email          string    `json:"email"`
	Telepon        string    `json:"telepon"`
	Alamat         string    `json:"alamat"`
	SatkerID       uuid.UUID `json:"satker_id"`
	JabatanID      *uuid.UUID `json:"jabatan_id,omitempty"`
	UnitKerjaID    *uuid.UUID `json:"unit_kerja_id,omitempty"`
	GolonganID     *uuid.UUID `json:"golongan_id,omitempty"`
	StatusPegawai  string    `json:"status_pegawai"`
	TMTJabatan     *time.Time `json:"tmt_jabatan,omitempty"`
	IsPNS          bool      `json:"is_pns"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

// UpdatePegawaiInput input untuk update pegawai
type UpdatePegawaiInput struct {
	Nama           string     `json:"nama"`
	GelarDepan     string     `json:"gelar_depan"`
	GelarBelakang  string     `json:"gelar_belakang"`
	Email          string     `json:"email"`
	Telepon        string     `json:"telepon"`
	Alamat         string     `json:"alamat"`
	SatkerID       uuid.UUID  `json:"satker_id"`
	JabatanID      *uuid.UUID `json:"jabatan_id,omitempty"`
	UnitKerjaID    *uuid.UUID `json:"unit_kerja_id,omitempty"`
	GolonganID     *uuid.UUID `json:"golongan_id,omitempty"`
	StatusPegawai  string     `json:"status_pegawai"`
	TMTJabatan     *time.Time `json:"tmt_jabatan,omitempty"`
	IsPNS          bool       `json:"is_pns"`
	IsActive       bool       `json:"is_active"`
}