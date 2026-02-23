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
func (r *PegawaiRepository) List(ctx context.Context, page, limit int, search, satkerID, jabatanID, golonganID, statusPegawai, statusKerja string) ([]models.Pegawai, int64, error) {
	offset := (page - 1) * limit

	// Query menggunakan kolom baru dari schema
	query := `SELECT p.id, p.nip, p.nip_lama, p.nama_lengkap, p.gelar_depan, p.gelar_belakang,
			  p.tempat_lahir, p.tanggal_lahir, p.jenis_kelamin,
			  p.agama_id, p.status_kawin_id, p.nik, p.email, p.telepon,
			  p.alamat, p.alamat_domisili, p.foto, p.satker_id, p.jabatan_id, p.unit_kerja_id,
			  p.golongan_id, p.eselon_id, p.status_pegawai, p.status_kerja,
			  p.tmt_cpns, p.tmt_pns, p.tmt_jabatan, p.tmt_pangkat_terakhir, p.tmt_jabatan_terakhir,
			  p.karpeg_no, p.karpeg_file, p.taspen_no, p.npwp,
			  p.bpjs_kesehatan, p.bpjs_ketenagakerjaan, p.kk_no, p.kk_file, p.ktp_no, p.ktp_file,
			  p.sikep_id, p.is_active, p.created_at, p.updated_at, p.created_by, p.updated_by, p.deleted_at, p.deleted_by
			  FROM pegawai p
			  WHERE p.is_active = true`
	args := []interface{}{}
	argCount := 1

	if search != "" {
		query += fmt.Sprintf(" AND (p.nip ILIKE $%d OR p.nama_lengkap ILIKE $%d)", argCount, argCount+1)
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

	if statusKerja != "" {
		query += fmt.Sprintf(" AND p.status_kerja = $%d", argCount)
		args = append(args, statusKerja)
		argCount++
	}

	// Get total count
	countQuery := "SELECT COUNT(*) FROM pegawai p WHERE p.is_active = true"
	countArgs := []interface{}{}
	countArgCount := 1

	if search != "" {
		countQuery += fmt.Sprintf(" AND (p.nip ILIKE $%d OR p.nama_lengkap ILIKE $%d)", countArgCount, countArgCount+1)
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

	if statusKerja != "" {
		countQuery += fmt.Sprintf(" AND p.status_kerja = $%d", countArgCount)
		countArgs = append(countArgs, statusKerja)
		countArgCount++
	}

	var total int64
	err := r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count pegawai: %w", err)
	}

	// Get data
	query += fmt.Sprintf(" ORDER BY p.nama_lengkap LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query pegawai: %w", err)
	}
	defer rows.Close()

	pegawais := []models.Pegawai{}
	for rows.Next() {
		var pegawai models.Pegawai
		var statusPegawai models.StatusPegawai
		var statusKerja models.StatusKerja

		err := rows.Scan(
			&pegawai.ID, &pegawai.NIP, &pegawai.NIPLama, &pegawai.NamaLengkap, &pegawai.GelarDepan, &pegawai.GelarBelakang,
			&pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.JenisKelamin,
			&pegawai.AgamaID, &pegawai.StatusKawinID, &pegawai.NIK, &pegawai.Email, &pegawai.Telepon,
			&pegawai.Alamat, &pegawai.AlamatDomisili, &pegawai.Foto, &pegawai.SatkerID, &pegawai.JabatanID, &pegawai.UnitKerjaID,
			&pegawai.GolonganID, &pegawai.EselonID, &statusPegawai, &statusKerja,
			&pegawai.TMTCpns, &pegawai.TMTPns, &pegawai.TMTJabatan, &pegawai.TMTPangkatTerakhir, &pegawai.TMTJabatanTerakhir,
			&pegawai.KarpegNo, &pegawai.KarpegFile, &pegawai.TaspenNo, &pegawai.NPWP,
			&pegawai.BPJSSehatan, &pegawai.BPJSKetenagakerjaan, &pegawai.KKNo, &pegawai.KKFile, &pegawai.KTPNo, &pegawai.KTPFile,
			&pegawai.SikepID, &pegawai.IsActive, &pegawai.CreatedAt, &pegawai.UpdatedAt, &pegawai.CreatedBy, &pegawai.UpdatedBy, &pegawai.DeletedAt, &pegawai.DeletedBy,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan pegawai: %w", err)
		}
		pegawai.StatusPegawai = statusPegawai
		pegawai.StatusKerja = statusKerja
		pegawais = append(pegawais, pegawai)
	}

	return pegawais, total, nil
}

// GetByID mengambil detail pegawai dengan relasi
func (r *PegawaiRepository) GetByID(ctx context.Context, id string) (*models.Pegawai, error) {
	query := `SELECT p.id, p.nip, p.nip_lama, p.nama_lengkap, p.gelar_depan, p.gelar_belakang,
			  p.tempat_lahir, p.tanggal_lahir, p.jenis_kelamin,
			  p.agama_id, p.status_kawin_id, p.nik, p.email, p.telepon,
			  p.alamat, p.alamat_domisili, p.foto, p.satker_id, p.jabatan_id, p.unit_kerja_id,
			  p.golongan_id, p.eselon_id, p.status_pegawai, p.status_kerja,
			  p.tmt_cpns, p.tmt_pns, p.tmt_jabatan, p.tmt_pangkat_terakhir, p.tmt_jabatan_terakhir,
			  p.karpeg_no, p.karpeg_file, p.taspen_no, p.npwp,
			  p.bpjs_kesehatan, p.bpjs_ketenagakerjaan, p.kk_no, p.kk_file, p.ktp_no, p.ktp_file,
			  p.sikep_id, p.is_active, p.created_at, p.updated_at, p.created_by, p.updated_by, p.deleted_at, p.deleted_by
			  FROM pegawai p
			  WHERE p.id = $1`

	var pegawai models.Pegawai
	var statusPegawai models.StatusPegawai
	var statusKerja models.StatusKerja

	err := r.db.QueryRow(ctx, query, uuid.MustParse(id)).Scan(
		&pegawai.ID, &pegawai.NIP, &pegawai.NIPLama, &pegawai.NamaLengkap, &pegawai.GelarDepan, &pegawai.GelarBelakang,
		&pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.JenisKelamin,
		&pegawai.AgamaID, &pegawai.StatusKawinID, &pegawai.NIK, &pegawai.Email, &pegawai.Telepon,
		&pegawai.Alamat, &pegawai.AlamatDomisili, &pegawai.Foto, &pegawai.SatkerID, &pegawai.JabatanID, &pegawai.UnitKerjaID,
		&pegawai.GolonganID, &pegawai.EselonID, &statusPegawai, &statusKerja,
		&pegawai.TMTCpns, &pegawai.TMTPns, &pegawai.TMTJabatan, &pegawai.TMTPangkatTerakhir, &pegawai.TMTJabatanTerakhir,
		&pegawai.KarpegNo, &pegawai.KarpegFile, &pegawai.TaspenNo, &pegawai.NPWP,
		&pegawai.BPJSSehatan, &pegawai.BPJSKetenagakerjaan, &pegawai.KKNo, &pegawai.KKFile, &pegawai.KTPNo, &pegawai.KTPFile,
		&pegawai.SikepID, &pegawai.IsActive, &pegawai.CreatedAt, &pegawai.UpdatedAt, &pegawai.CreatedBy, &pegawai.UpdatedBy, &pegawai.DeletedAt, &pegawai.DeletedBy,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("pegawai not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get pegawai: %w", err)
	}

	pegawai.StatusPegawai = statusPegawai
	pegawai.StatusKerja = statusKerja

	return &pegawai, nil
}

// GetByNIP mengambil detail pegawai berdasarkan NIP
func (r *PegawaiRepository) GetByNIP(ctx context.Context, nip string) (*models.Pegawai, error) {
	query := `SELECT p.id, p.nip, p.nip_lama, p.nama_lengkap, p.gelar_depan, p.gelar_belakang,
			  p.tempat_lahir, p.tanggal_lahir, p.jenis_kelamin,
			  p.agama_id, p.status_kawin_id, p.nik, p.email, p.telepon,
			  p.alamat, p.alamat_domisili, p.foto, p.satker_id, p.jabatan_id, p.unit_kerja_id,
			  p.golongan_id, p.eselon_id, p.status_pegawai, p.status_kerja,
			  p.tmt_cpns, p.tmt_pns, p.tmt_jabatan, p.tmt_pangkat_terakhir, p.tmt_jabatan_terakhir,
			  p.karpeg_no, p.karpeg_file, p.taspen_no, p.npwp,
			  p.bpjs_kesehatan, p.bpjs_ketenagakerjaan, p.kk_no, p.kk_file, p.ktp_no, p.ktp_file,
			  p.sikep_id, p.is_active, p.created_at, p.updated_at, p.created_by, p.updated_by, p.deleted_at, p.deleted_by
			  FROM pegawai p
			  WHERE p.nip = $1 AND p.is_active = true`

	var pegawai models.Pegawai
	var statusPegawai models.StatusPegawai
	var statusKerja models.StatusKerja

	err := r.db.QueryRow(ctx, query, nip).Scan(
		&pegawai.ID, &pegawai.NIP, &pegawai.NIPLama, &pegawai.NamaLengkap, &pegawai.GelarDepan, &pegawai.GelarBelakang,
		&pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.JenisKelamin,
		&pegawai.AgamaID, &pegawai.StatusKawinID, &pegawai.NIK, &pegawai.Email, &pegawai.Telepon,
		&pegawai.Alamat, &pegawai.AlamatDomisili, &pegawai.Foto, &pegawai.SatkerID, &pegawai.JabatanID, &pegawai.UnitKerjaID,
		&pegawai.GolonganID, &pegawai.EselonID, &statusPegawai, &statusKerja,
		&pegawai.TMTCpns, &pegawai.TMTPns, &pegawai.TMTJabatan, &pegawai.TMTPangkatTerakhir, &pegawai.TMTJabatanTerakhir,
		&pegawai.KarpegNo, &pegawai.KarpegFile, &pegawai.TaspenNo, &pegawai.NPWP,
		&pegawai.BPJSSehatan, &pegawai.BPJSKetenagakerjaan, &pegawai.KKNo, &pegawai.KKFile, &pegawai.KTPNo, &pegawai.KTPFile,
		&pegawai.SikepID, &pegawai.IsActive, &pegawai.CreatedAt, &pegawai.UpdatedAt, &pegawai.CreatedBy, &pegawai.UpdatedBy, &pegawai.DeletedAt, &pegawai.DeletedBy,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("pegawai not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get pegawai: %w", err)
	}

	pegawai.StatusPegawai = statusPegawai
	pegawai.StatusKerja = statusKerja

	return &pegawai, nil
}

// Create membuat pegawai baru
func (r *PegawaiRepository) Create(ctx context.Context, input CreatePegawaiInput) (*models.Pegawai, error) {
	id := uuid.New()

	query := `INSERT INTO pegawai (
		id, nip, nip_lama, nama_lengkap, gelar_depan, gelar_belakang,
		tempat_lahir, tanggal_lahir, jenis_kelamin, agama_id, status_kawin_id,
		nik, email, telepon, alamat, alamat_domisili, satker_id, jabatan_id, unit_kerja_id,
		golongan_id, eselon_id, status_pegawai, status_kerja,
		tmt_cpns, tmt_pns, tmt_jabatan, tmt_pangkat_terakhir, tmt_jabatan_terakhir,
		karpeg_no, taspen_no, npwp, bpjs_kesehatan, bpjs_ketenagakerjaan, kk_no, ktp_no, sikep_id,
		is_active, created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41
	) RETURNING created_at, updated_at`

	now := time.Now()

	err := r.db.QueryRow(ctx, query,
		id, input.NIP, input.NIPLama, input.NamaLengkap, input.GelarDepan, input.GelarBelakang,
		input.TempatLahir, input.TanggalLahir, input.JenisKelamin,
		input.AgamaID, input.StatusKawinID, input.NIK, input.Email,
		input.Telepon, input.Alamat, input.AlamatDomisili, input.SatkerID, input.JabatanID, input.UnitKerjaID,
		input.GolonganID, input.EselonID, input.StatusPegawai, input.StatusKerja,
		input.TMTCpns, input.TMTPns, input.TMTJabatan, input.TMTPangkatTerakhir, input.TMTJabatanTerakhir,
		input.KarpegNo, input.TaspenNo, input.NPWP, input.BPJSSehatan, input.BPJSKetenagakerjaan, input.KKNo, input.KTPNo, input.SikepID,
		true, now, now,
	).Scan(&now, &now) // dummy scan untuk createdAt, updatedAt

	if err != nil {
		return nil, fmt.Errorf("failed to create pegawai: %w", err)
	}

	pegawai := &models.Pegawai{
		ID:                 id,
		NIP:                input.NIP,
		NIPLama:            input.NIPLama,
		NamaLengkap:       input.NamaLengkap,
		GelarDepan:         input.GelarDepan,
		GelarBelakang:      input.GelarBelakang,
		TempatLahir:        input.TempatLahir,
		TanggalLahir:       input.TanggalLahir,
		JenisKelamin:        input.JenisKelamin,
		AgamaID:             input.AgamaID,
		StatusKawinID:       input.StatusKawinID,
		NIK:                 input.NIK,
		Email:               input.Email,
		Telepon:             input.Telepon,
		Alamat:              input.Alamat,
		AlamatDomisili:      input.AlamatDomisili,
		SatkerID:            input.SatkerID,
		JabatanID:           input.JabatanID,
		UnitKerjaID:         input.UnitKerjaID,
		GolonganID:          input.GolonganID,
		EselonID:            input.EselonID,
		StatusPegawai:       input.StatusPegawai,
		StatusKerja:         input.StatusKerja,
		TMTCpns:             input.TMTCpns,
		TMTPns:              input.TMTPns,
		TMTJabatan:          input.TMTJabatan,
		TMTPangkatTerakhir: input.TMTPangkatTerakhir,
		TMTJabatanTerakhir:  input.TMTJabatanTerakhir,
		KarpegNo:            input.KarpegNo,
		TaspenNo:            input.TaspenNo,
		NPWP:                input.NPWP,
		BPJSSehatan:         input.BPJSSehatan,
		BPJSKetenagakerjaan: input.BPJSKetenagakerjaan,
		KKNo:                input.KKNo,
		KTPNo:               input.KTPNo,
		SikepID:             input.SikepID,
		IsActive:            true,
		CreatedAt:           now,
		UpdatedAt:           now,
	}

	return pegawai, nil
}

// Update mengupdate pegawai
func (r *PegawaiRepository) Update(ctx context.Context, id string, input UpdatePegawaiInput) (*models.Pegawai, error) {
	query := `UPDATE pegawai
			  SET nama_lengkap = $2, gelar_depan = $3, gelar_belakang = $4,
				  email = $5, telepon = $6, alamat = $7, alamat_domisili = $8,
				  satker_id = $9, jabatan_id = $10, unit_kerja_id = $11,
				  golongan_id = $12, eselon_id = $13, status_pegawai = $14, status_kerja = $15,
				  tmt_jabatan = $16, tmt_pangkat_terakhir = $17, updated_at = NOW()
			  WHERE id = $1
			  RETURNING nip, nip_lama, tempat_lahir, tanggal_lahir, jenis_kelamin,
			  agama_id, status_kawin_id, nik, foto, created_at, updated_at,
			  karpeg_no, karpeg_file, taspen_no, npwp, bpjs_kesehatan, bpjs_ketenagakerjaan, kk_no, kk_file, ktp_no, ktp_file, sikep_id,
			  tmt_cpns, tmt_pns, tmt_jabatan_terakhir, is_active, created_by, updated_by, deleted_at, deleted_by`

	var pegawai models.Pegawai
	var statusPegawai models.StatusPegawai
	var statusKerja models.StatusKerja

	err := r.db.QueryRow(ctx, query,
		uuid.MustParse(id), input.NamaLengkap, input.GelarDepan, input.GelarBelakang,
		input.Email, input.Telepon, input.Alamat, input.AlamatDomisili, input.SatkerID,
		input.JabatanID, input.UnitKerjaID, input.GolonganID,
		input.EselonID, input.StatusPegawai, input.StatusKerja, input.TMTJabatan, input.TMTPangkatTerakhir,
	).Scan(
		&pegawai.NIP, &pegawai.NIPLama, &pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.JenisKelamin,
		&pegawai.AgamaID, &pegawai.StatusKawinID, &pegawai.NIK, &pegawai.Foto,
		&pegawai.CreatedAt, &pegawai.UpdatedAt,
		&pegawai.KarpegNo, &pegawai.KarpegFile, &pegawai.TaspenNo, &pegawai.NPWP,
		&pegawai.BPJSSehatan, &pegawai.BPJSKetenagakerjaan, &pegawai.KKNo, &pegawai.KKFile, &pegawai.KTPNo, &pegawai.KTPFile, &pegawai.SikepID,
		&pegawai.TMTCpns, &pegawai.TMTPns, &pegawai.TMTJabatanTerakhir, &pegawai.IsActive, &pegawai.CreatedBy, &pegawai.UpdatedBy, &pegawai.DeletedAt, &pegawai.DeletedBy,
		&statusPegawai, &statusKerja,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("pegawai not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update pegawai: %w", err)
	}

	pegawai.ID = uuid.MustParse(id)
	pegawai.NamaLengkap = input.NamaLengkap
	pegawai.GelarDepan = input.GelarDepan
	pegawai.GelarBelakang = input.GelarBelakang
	pegawai.Email = input.Email
	pegawai.Telepon = input.Telepon
	pegawai.Alamat = input.Alamat
	pegawai.AlamatDomisili = input.AlamatDomisili
	pegawai.SatkerID = input.SatkerID
	pegawai.JabatanID = input.JabatanID
	pegawai.UnitKerjaID = input.UnitKerjaID
	pegawai.GolonganID = input.GolonganID
	pegawai.EselonID = input.EselonID
	pegawai.StatusPegawai = statusPegawai
	pegawai.StatusKerja = statusKerja
	pegawai.TMTJabatan = input.TMTJabatan
	pegawai.TMTPangkatTerakhir = input.TMTPangkatTerakhir

	return &pegawai, nil
}

// Delete menghapus pegawai (soft delete)
func (r *PegawaiRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE pegawai SET is_active = false, deleted_at = NOW(), updated_at = NOW() WHERE id = $1`

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

	// Pegawai per status pegawai (PNS, CPNS, PPPK, HONORER)
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
	statistik["per_status_pegawai"] = statusData

	// Pegawai per status kerja
	kerjaQuery := `SELECT status_kerja, COUNT(*) FROM pegawai WHERE is_active = true GROUP BY status_kerja`
	rows, err = r.db.Query(ctx, kerjaQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query pegawai by status kerja: %w", err)
	}
	defer rows.Close()

	kerjaData := make(map[string]int64)
	for rows.Next() {
		var status string
		var count int64
		err := rows.Scan(&status, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan status kerja: %w", err)
		}
		kerjaData[status] = count
	}
	statistik["per_status_kerja"] = kerjaData

	// Pegawai PNS vs Non-PNS (berdasarkan status_pegawai)
	var totalPNS, totalNonPNS int64
	r.db.QueryRow(ctx, "SELECT COUNT(*) FROM pegawai WHERE status_pegawai IN ('PNS', 'CPNS') AND is_active = true").Scan(&totalPNS)
	r.db.QueryRow(ctx, "SELECT COUNT(*) FROM pegawai WHERE status_pegawai IN ('PPPK', 'HONORER') AND is_active = true").Scan(&totalNonPNS)
	statistik["pns"] = totalPNS
	statistik["non_pns"] = totalNonPNS

	// Pegawai per golongan
	golonganQuery := `SELECT g.nama, COUNT(*) FROM pegawai p
					  JOIN ref_golongan g ON p.golongan_id = g.id
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
	NIP                  string                `json:"nip"`
	NIPLama              *string               `json:"nip_lama,omitempty"`
	NamaLengkap          string                `json:"nama_lengkap"`
	GelarDepan           *string               `json:"gelar_depan,omitempty"`
	GelarBelakang        *string               `json:"gelar_belakang,omitempty"`
	TempatLahir          string                `json:"tempat_lahir"`
	TanggalLahir         time.Time             `json:"tanggal_lahir"`
	JenisKelamin         string                `json:"jenis_kelamin"`
	AgamaID              uuid.UUID             `json:"agama_id"`
	StatusKawinID        uuid.UUID             `json:"status_kawin_id"`
	NIK                  *string               `json:"nik,omitempty"`
	Email                *string               `json:"email,omitempty"`
	Telepon              *string               `json:"telepon,omitempty"`
	Alamat               *string               `json:"alamat,omitempty"`
	AlamatDomisili       *string               `json:"alamat_domisili,omitempty"`
	SatkerID             uuid.UUID             `json:"satker_id"`
	JabatanID            *uuid.UUID            `json:"jabatan_id,omitempty"`
	UnitKerjaID          *uuid.UUID            `json:"unit_kerja_id,omitempty"`
	GolonganID           *uuid.UUID            `json:"golongan_id,omitempty"`
	EselonID             *uuid.UUID            `json:"eselon_id,omitempty"`
	StatusPegawai        models.StatusPegawai  `json:"status_pegawai"`
	StatusKerja          models.StatusKerja    `json:"status_kerja"`
	TMTCpns              *time.Time            `json:"tmt_cpns,omitempty"`
	TMTPns               *time.Time            `json:"tmt_pns,omitempty"`
	TMTJabatan           *time.Time            `json:"tmt_jabatan,omitempty"`
	TMTPangkatTerakhir   *time.Time            `json:"tmt_pangkat_terakhir,omitempty"`
	TMTJabatanTerakhir   *time.Time            `json:"tmt_jabatan_terakhir,omitempty"`
	KarpegNo             *string               `json:"karpeg_no,omitempty"`
	TaspenNo             *string               `json:"taspen_no,omitempty"`
	NPWP                 *string               `json:"npwp,omitempty"`
	BPJSSehatan          *string               `json:"bpjs_kesehatan,omitempty"`
	BPJSKetenagakerjaan  *string               `json:"bpjs_ketenagakerjaan,omitempty"`
	KKNo                 *string               `json:"kk_no,omitempty"`
	KTPNo                *string               `json:"ktp_no,omitempty"`
	SikepID              *string               `json:"sikep_id,omitempty"`
}

// UpdatePegawaiInput input untuk update pegawai
type UpdatePegawaiInput struct {
	NamaLengkap        string               `json:"nama_lengkap"`
	GelarDepan         *string              `json:"gelar_depan,omitempty"`
	GelarBelakang      *string              `json:"gelar_belakang,omitempty"`
	Email              *string              `json:"email,omitempty"`
	Telepon            *string              `json:"telepon,omitempty"`
	Alamat             *string              `json:"alamat,omitempty"`
	AlamatDomisili     *string              `json:"alamat_domisili,omitempty"`
	SatkerID           uuid.UUID            `json:"satker_id"`
	JabatanID          *uuid.UUID           `json:"jabatan_id,omitempty"`
	UnitKerjaID        *uuid.UUID           `json:"unit_kerja_id,omitempty"`
	GolonganID         *uuid.UUID           `json:"golongan_id,omitempty"`
	EselonID           *uuid.UUID           `json:"eselon_id,omitempty"`
	StatusPegawai      models.StatusPegawai `json:"status_pegawai"`
	StatusKerja        models.StatusKerja   `json:"status_kerja"`
	TMTJabatan         *time.Time           `json:"tmt_jabatan,omitempty"`
	TMTPangkatTerakhir *time.Time           `json:"tmt_pangkat_terakhir,omitempty"`
}
