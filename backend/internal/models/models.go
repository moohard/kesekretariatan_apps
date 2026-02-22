package models

import (
	"time"
	"github.com/google/uuid"
)

// ==================== MASTER DATA MODELS ====================

// Satker (Satuan Kerja)
type Satker struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	ParentID    *uuid.UUID     `json:"parent_id,omitempty" db:"parent_id"`
	Level       int            `json:"level" db:"level"`
	Alamat      string         `json:"alamat" db:"alamat"`
	Telepon     string         `json:"telepon" db:"telepon"`
	Email       string         `json:"email" db:"email"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
	CreatedBy   *uuid.UUID     `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy   *uuid.UUID     `json:"updated_by,omitempty" db:"updated_by"`
}

// Jabatan
type Jabatan struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	EselonID    *uuid.UUID     `json:"eselon_id,omitempty" db:"eselon_id"`
	Kelas       string         `json:"kelas" db:"kelas"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// Golongan
type Golongan struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	Ruang       string         `json:"ruang" db:"ruang"`
	Angka       int            `json:"angka" db:"angka"`
	MinPangkat  int            `json:"min_pangkat" db:"min_pangkat"`
	MaxPangkat  int            `json:"max_pangkat" db:"max_pangkat"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// UnitKerja
type UnitKerja struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	Singkatan   string         `json:"singkatan" db:"singkatan"`
	ParentID    *uuid.UUID     `json:"parent_id,omitempty" db:"parent_id"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// Eselon
type Eselon struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	Tunjangan   float64        `json:"tunjangan" db:"tunjangan"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// RefPendidikan
type RefPendidikan struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	Tingkat     string         `json:"tingkat" db:"tingkat"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// RefAgama
type RefAgama struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// RefStatusKawin
type RefStatusKawin struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// RefJenisHukdis
type RefJenisHukdis struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// RefJenisDiklat
type RefJenisDiklat struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Kode        string         `json:"kode" db:"kode"`
	Nama        string         `json:"nama" db:"nama"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// ==================== KEGAWAAN MODELS ====================

// Pegawai
type Pegawai struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	NIP             string         `json:"nip" db:"nip"`
	Nama            string         `json:"nama" db:"nama"`
	GelarDepan      string         `json:"gelar_depan" db:"gelar_depan"`
	GelarBelakang   string         `json:"gelar_belakang" db:"gelar_belakang"`
	TempatLahir     string         `json:"tempat_lahir" db:"tempat_lahir"`
	TanggalLahir    time.Time      `json:"tanggal_lahir" db:"tanggal_lahir"`
	JenisKelamin    string         `json:"jenis_kelamin" db:"jenis_kelamin"`
	AgamaID         uuid.UUID      `json:"agama_id" db:"agama_id"`
	StatusKawinID   uuid.UUID      `json:"status_kawin_id" db:"status_kawin_id"`
	NIK             string         `json:"nik" db:"nik"`
	Email           string         `json:"email" db:"email"`
	Telepon         string         `json:"telepon" db:"telepon"`
	Alamat          string         `json:"alamat" db:"alamat"`
	Foto            *string        `json:"foto,omitempty" db:"foto"`
	SatkerID        uuid.UUID      `json:"satker_id" db:"satker_id"`
	JabatanID       *uuid.UUID     `json:"jabatan_id,omitempty" db:"jabatan_id"`
	UnitKerjaID     *uuid.UUID     `json:"unit_kerja_id,omitempty" db:"unit_kerja_id"`
	GolonganID      *uuid.UUID     `json:"golongan_id,omitempty" db:"golongan_id"`
	StatusPegawai   string         `json:"status_pegawai" db:"status_pegawai"` // aktif, pensiun, mutasi
	TMTJabatan      *time.Time     `json:"tmt_jabatan,omitempty" db:"tmt_jabatan"`
	IsPNS           bool           `json:"is_pns" db:"is_pns"`
	IsActive        bool           `json:"is_active" db:"is_active"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`

	// Relasi
	Satker          *Satker        `json:"satker,omitempty"`
	Jabatan         *Jabatan       `json:"jabatan,omitempty"`
	UnitKerja       *UnitKerja     `json:"unit_kerja,omitempty"`
	Golongan        *Golongan      `json:"golongan,omitempty"`
	Agama           *RefAgama      `json:"agama,omitempty"`
	StatusKawin     *RefStatusKawin `json:"status_kawin,omitempty"`
}

// RiwayatPangkat
type RiwayatPangkat struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	PegawaiID       uuid.UUID      `json:"pegawai_id" db:"pegawai_id"`
	GolonganID      uuid.UUID      `json:"golongan_id" db:"golongan_id"`
	Pangkat         string         `json:"pangkat" db:"pangkat"`
	TMT             time.Time      `json:"tmt" db:"tmt"`
	NomorSK         string         `json:"nomor_sk" db:"nomor_sk"`
	TanggalSK       time.Time      `json:"tanggal_sk" db:"tanggal_sk"`
	Pejabat         string         `json:"pejabat" db:"pejabat"`
	FileSK          *string        `json:"file_sk,omitempty" db:"file_sk"`
	GajiPokok       float64        `json:"gaji_pokok" db:"gaji_pokok"`
	IsTerakhir      bool           `json:"is_terakhir" db:"is_terakhir"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`

	// Relasi
	Golongan        *Golongan      `json:"golongan,omitempty"`
}

// RiwayatJabatan
type RiwayatJabatan struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	PegawaiID       uuid.UUID      `json:"pegawai_id" db:"pegawai_id"`
	JabatanID       *uuid.UUID     `json:"jabatan_id,omitempty" db:"jabatan_id"`
	UnitKerjaID     *uuid.UUID     `json:"unit_kerja_id,omitempty" db:"unit_kerja_id"`
	SatkerID        *uuid.UUID     `json:"satker_id,omitempty" db:"satker_id"`
	NamaJabatan     string         `json:"nama_jabatan" db:"nama_jabatan"`
	TMT             time.Time      `json:"tmt" db:"tmt"`
	NomorSK         string         `json:"nomor_sk" db:"nomor_sk"`
	TanggalSK       time.Time      `json:"tanggal_sk" db:"tanggal_sk"`
	Pejabat         string         `json:"pejabat" db:"pejabat"`
	FileSK          *string        `json:"file_sk,omitempty" db:"file_sk"`
	IsTerakhir      bool           `json:"is_terakhir" db:"is_terakhir"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`

	// Relasi
	Jabatan         *Jabatan       `json:"jabatan,omitempty"`
	UnitKerja       *UnitKerja     `json:"unit_kerja,omitempty"`
	Satker          *Satker        `json:"satker,omitempty"`
}

// RiwayatPendidikan
type RiwayatPendidikan struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	PegawaiID       uuid.UUID      `json:"pegawai_id" db:"pegawai_id"`
	PendidikanID    uuid.UUID      `json:"pendidikan_id" db:"pendidikan_id"`
	NamaInstitusi   string         `json:"nama_institusi" db:"nama_institusi"`
	Jurusan         string         `json:"jurusan" db:"jurusan"`
	TahunMasuk      int            `json:"tahun_masuk" db:"tahun_masuk"`
	TahunLulus      int            `json:"tahun_lulus" db:"tahun_lulus"`
	NomorIjazah     string         `json:"nomor_ijazah" db:"nomor_ijazah"`
	TanggalIjazah   *time.Time     `json:"tanggal_ijazah,omitempty" db:"tanggal_ijazah"`
	FileIjazah      *string        `json:"file_ijazah,omitempty" db:"file_ijazah"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`

	// Relasi
	Pendidikan      *RefPendidikan `json:"pendidikan,omitempty"`
}

// Keluarga
type Keluarga struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	PegawaiID       uuid.UUID      `json:"pegawai_id" db:"pegawai_id"`
	StatusKeluarga  string         `json:"status_keluarga" db:"status_keluarga"` // suami, istri, anak
	Nama            string         `json:"nama" db:"nama"`
	TempatLahir     string         `json:"tempat_lahir" db:"tempat_lahir"`
	TanggalLahir    *time.Time     `json:"tanggal_lahir,omitempty" db:"tanggal_lahir"`
	JenisKelamin    string         `json:"jenis_kelamin" db:"jenis_kelamin"`
	NIK             *string        `json:"nik,omitempty" db:"nik"`
	Pendidikan      string         `json:"pendidikan" db:"pendidikan"`
	Pekerjaan       string         `json:"pekerjaan" db:"pekerjaan"`
	IsTanggungan    bool           `json:"is_tanggungan" db:"is_tanggungan"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
}

// TemplateDokumen
type TemplateDokumen struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	Kode            string         `json:"kode" db:"kode"`
	Nama            string         `json:"nama" db:"nama"`
	Tipe            string         `json:"tipe" db:"tipe"` // sk_pangkat, sk_jabatan, dll
	KontenHTML      string         `json:"konten_html" db:"konten_html"`
	IsActive        bool           `json:"is_active" db:"is_active"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
}

// ==================== RBAC MODELS ====================

// AppRole
type AppRole struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Nama        string         `json:"nama" db:"nama"`
	Deskripsi   string         `json:"deskripsi" db:"deskripsi"`
	IsActive    bool           `json:"is_active" db:"is_active"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// AppPermission
type AppPermission struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Nama        string         `json:"nama" db:"nama"`
	Resource    string         `json:"resource" db:"resource"`
	Action      string         `json:"action" db:"action"`
	Deskripsi   string         `json:"deskripsi" db:"deskripsi"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

// RolePermission
type RolePermission struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	RoleID      uuid.UUID      `json:"role_id" db:"role_id"`
	PermissionID uuid.UUID     `json:"permission_id" db:"permission_id"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
}

// UserAppRole
type UserAppRole struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	UserID      string         `json:"user_id" db:"user_id"` // Keycloak user ID
	RoleID      uuid.UUID      `json:"role_id" db:"role_id"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
}

// ==================== AUDIT MODELS ====================

// AuditLog
type AuditLog struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	UserID      *string        `json:"user_id,omitempty" db:"user_id"`
	Username    *string        `json:"username,omitempty" db:"username"`
	Action      string         `json:"action" db:"action"` // create, update, delete, read
	Resource    string         `json:"resource" db:"resource"`
	ResourceID  *string        `json:"resource_id,omitempty" db:"resource_id"`
	IPAddress   *string        `json:"ip_address,omitempty" db:"ip_address"`
	UserAgent   *string        `json:"user_agent,omitempty" db:"user_agent"`
	Changes     map[string]interface{} `json:"changes,omitempty" db:"changes"`
	Status      string         `json:"status" db:"status"` // success, failed
	ErrorMessage *string       `json:"error_message,omitempty" db:"error_message"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
}

// ==================== REQUEST/RESPONSE DTOs ====================

// PaginationRequest
type PaginationRequest struct {
	Page        int    `json:"page"`
	PerPage     int    `json:"per_page"`
	SortBy      string `json:"sort_by"`
	SortOrder   string `json:"sort_order"` // asc, desc
	Search      string `json:"search"`
	Filters     map[string]interface{} `json:"filters"`
}

// PaginationResponse
type PaginationResponse struct {
	Data        interface{} `json:"data"`
	Total       int64       `json:"total"`
	Page        int         `json:"page"`
	PerPage     int         `json:"per_page"`
	TotalPages  int         `json:"total_pages"`
}

// APIResponse
type APIResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Error       string      `json:"error,omitempty"`
	RequestID   string      `json:"request_id,omitempty"`
}

// LoginRequest
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse
type LoginResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	ExpiresIn    int64   `json:"expires_in"`
	User         UserDTO `json:"user"`
}

// UserDTO
type UserDTO struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
	Roles    []string `json:"roles"`
}