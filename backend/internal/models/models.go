package models

import (
	"time"
	"github.com/google/uuid"
)

// ==================== ENUMS ====================

// StatusPegawai - Kategori kepegawaian
type StatusPegawai string

const (
	StatusPegawaiPNS     StatusPegawai = "PNS"
	StatusPegawaiCPNS    StatusPegawai = "CPNS"
	StatusPegawaiPPPK    StatusPegawai = "PPPK"
	StatusPegawaiHonorer StatusPegawai = "HONORER"
)

// StatusKerja - Status kerja pegawai
type StatusKerja string

const (
	StatusKerjaAktif          StatusKerja = "aktif"
	StatusKerjaCuti           StatusKerja = "cuti"
	StatusKerjaPensiun        StatusKerja = "pensiun"
	StatusKerjaMutasiKeluar   StatusKerja = "mutasi_keluar"
	StatusKerjaMutasiMasuk    StatusKerja = "mutasi_masuk"
	StatusKerjaMeninggal      StatusKerja = "meninggal"
	StatusKerjaPemberhentian  StatusKerja = "pemberhentian"
)

// JenisJabatan - Jenis jabatan
type JenisJabatan string

const (
	JenisJabatanStruktural         JenisJabatan = "struktural"
	JenisJabatanFungsionalTertentu JenisJabatan = "fungsional_tertentu"
	JenisJabatanFungsionalUmum     JenisJabatan = "fungsional_umum"
	JenisJabatanPelaksana           JenisJabatan = "pelaksana"
)

// JenisKenaikanPangkat - Jenis kenaikan pangkat
type JenisKenaikanPangkat string

const (
	JenisKenaikanReguler           JenisKenaikanPangkat = "reguler"
	JenisKenaikanPilihan           JenisKenaikanPangkat = "pilihan"
	JenisKenaikanPenyesuaianIjazah JenisKenaikanPangkat = "penyesuaian_ijazah"
	JenisKenaikanLainnya           JenisKenaikanPangkat = "lainnya"
)

// StatusKeluarga - Hubungan keluarga
type StatusKeluarga string

const (
	StatusKeluargaSuami StatusKeluarga = "Suami"
	StatusKeluargaIstri StatusKeluarga = "Istri"
	StatusKeluargaAnak  StatusKeluarga = "Anak"
	StatusKeluargaAyah  StatusKeluarga = "Ayah"
	StatusKeluargaIbu   StatusKeluarga = "Ibu"
)

// KategoriGolonganNonPNS - Kategori golongan non-PNS
type KategoriGolonganNonPNS string

const (
	KategoriGolonganHonorer KategoriGolonganNonPNS = "Honorer"
	KategoriGolonganKontrak KategoriGolonganNonPNS = "Kontrak"
)

// ==================== MASTER DATA MODELS ====================

// Satker (Satuan Kerja)
type Satker struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Kode        string     `json:"kode" db:"kode"`
	Nama        string     `json:"nama" db:"nama"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty" db:"parent_id"`
	Level       int        `json:"level" db:"level"`
	Alamat      string     `json:"alamat,omitempty" db:"alamat"`
	Telepon     string     `json:"telepon,omitempty" db:"telepon"`
	Email       string     `json:"email,omitempty" db:"email"`
	IsActive    bool       `json:"is_active" db:"is_active"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy   *uuid.UUID `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy   *uuid.UUID `json:"updated_by,omitempty" db:"updated_by"`
}

// Jabatan
type Jabatan struct {
	ID        uuid.UUID     `json:"id" db:"id"`
	Kode      string        `json:"kode" db:"kode"`
	Nama      string        `json:"nama" db:"nama"`
	EselonID  *uuid.UUID    `json:"eselon_id,omitempty" db:"eselon_id"`
	Kelas     string        `json:"kelas,omitempty" db:"kelas"`
	Jenis     *JenisJabatan `json:"jenis,omitempty" db:"jenis"`
	IsActive  bool          `json:"is_active" db:"is_active"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`

	// Relations
	Eselon *Eselon `json:"eselon,omitempty"`
}

// Golongan (PNS)
type Golongan struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Kode       string    `json:"kode" db:"kode"`
	Nama       string    `json:"nama" db:"nama"`
	Ruang      string    `json:"ruang" db:"ruang"`
	Angka      int       `json:"angka" db:"angka"`
	MinPangkat int       `json:"min_pangkat,omitempty" db:"min_pangkat"`
	MaxPangkat int       `json:"max_pangkat,omitempty" db:"max_pangkat"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// GolonganNonPNS - Golongan untuk pegawai non-PNS (Honorer, Kontrak, PPPK)
type GolonganNonPNS struct {
	ID         uuid.UUID              `json:"id" db:"id"`
	Kode       string                 `json:"kode" db:"kode"`
	Nama       string                 `json:"nama" db:"nama"`
	Kategori   KategoriGolonganNonPNS `json:"kategori" db:"kategori"`
	Urutan     int                    `json:"urutan" db:"urutan"`
	Keterangan *string                `json:"keterangan,omitempty" db:"keterangan"`
	IsActive   bool                   `json:"is_active" db:"is_active"`
	CreatedAt  time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at" db:"updated_at"`
}

// UnitKerja
type UnitKerja struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	Kode       string     `json:"kode" db:"kode"`
	Nama       string     `json:"nama" db:"nama"`
	Singkatan  string     `json:"singkatan,omitempty" db:"singkatan"`
	ParentID   *uuid.UUID `json:"parent_id,omitempty" db:"parent_id"`
	SatkerID   *uuid.UUID `json:"satker_id,omitempty" db:"satker_id"`
	IsActive   bool       `json:"is_active" db:"is_active"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
}

// Eselon
type Eselon struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Kode      string    `json:"kode" db:"kode"`
	Nama      string    `json:"nama" db:"nama"`
	Level     int       `json:"level,omitempty" db:"level"`
	Tunjangan float64   `json:"tunjangan,omitempty" db:"tunjangan"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// RefPendidikan
type RefPendidikan struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Kode      string    `json:"kode" db:"kode"`
	Nama      string    `json:"nama" db:"nama"`
	Tingkat   string    `json:"tingkat,omitempty" db:"tingkat"`
	Urutan    int       `json:"urutan,omitempty" db:"urutan"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// RefAgama
type RefAgama struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Kode      string    `json:"kode" db:"kode"`
	Nama      string    `json:"nama" db:"nama"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// RefStatusKawin
type RefStatusKawin struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Kode      string    `json:"kode" db:"kode"`
	Nama      string    `json:"nama" db:"nama"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// RefJenisHukdis
type RefJenisHukdis struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Kode      string    `json:"kode" db:"kode"`
	Nama      string    `json:"nama" db:"nama"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// RefJenisDiklat
type RefJenisDiklat struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Kode      string    `json:"kode" db:"kode"`
	Nama      string    `json:"nama" db:"nama"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ==================== KEPEGAWAIAN MODELS ====================

// Pegawai - Model lengkap dengan field baru
type Pegawai struct {
	ID     uuid.UUID `json:"id" db:"id"`
	NIP    string    `json:"nip" db:"nip"`
	NIPLama *string  `json:"nip_lama,omitempty" db:"nip_lama"`

	// Biodata
	NamaLengkap    string     `json:"nama_lengkap" db:"nama_lengkap"`
	GelarDepan     *string    `json:"gelar_depan,omitempty" db:"gelar_depan"`
	GelarBelakang  *string    `json:"gelar_belakang,omitempty" db:"gelar_belakang"`
	TempatLahir    string     `json:"tempat_lahir" db:"tempat_lahir"`
	TanggalLahir   time.Time  `json:"tanggal_lahir" db:"tanggal_lahir"`
	JenisKelamin   string     `json:"jenis_kelamin" db:"jenis_kelamin"`
	AgamaID        uuid.UUID  `json:"agama_id" db:"agama_id"`
	StatusKawinID  uuid.UUID  `json:"status_kawin_id" db:"status_kawin_id"`
	NIK            *string    `json:"nik,omitempty" db:"nik"`
	Email          *string    `json:"email,omitempty" db:"email"`
	Telepon        *string    `json:"telepon,omitempty" db:"telepon"`
	Alamat         *string    `json:"alamat,omitempty" db:"alamat"`
	AlamatDomisili *string    `json:"alamat_domisili,omitempty" db:"alamat_domisili"`
	Foto           *string    `json:"foto,omitempty" db:"foto"`

	// Posisi
	SatkerID     uuid.UUID  `json:"satker_id" db:"satker_id"`
	JabatanID    *uuid.UUID `json:"jabatan_id,omitempty" db:"jabatan_id"`
	UnitKerjaID  *uuid.UUID `json:"unit_kerja_id,omitempty" db:"unit_kerja_id"`
	GolonganID   *uuid.UUID `json:"golongan_id,omitempty" db:"golongan_id"`
	EselonID     *uuid.UUID `json:"eselon_id,omitempty" db:"eselon_id"`

	// Status
	StatusPegawai StatusPegawai `json:"status_pegawai" db:"status_pegawai"` // PNS, CPNS, PPPK, HONORER
	StatusKerja   StatusKerja   `json:"status_kerja" db:"status_kerja"`     // aktif, cuti, pensiun, dll

	// TMT (Terhitung Mulai Tanggal)
	TMTCpns             *time.Time `json:"tmt_cpns,omitempty" db:"tmt_cpns"`
	TMTPns              *time.Time `json:"tmt_pns,omitempty" db:"tmt_pns"`
	TMTJabatan          *time.Time `json:"tmt_jabatan,omitempty" db:"tmt_jabatan"`
	TMTPangkatTerakhir  *time.Time `json:"tmt_pangkat_terakhir,omitempty" db:"tmt_pangkat_terakhir"`
	TMTJabatanTerakhir  *time.Time `json:"tmt_jabatan_terakhir,omitempty" db:"tmt_jabatan_terakhir"`

	// Dokumen Kepegawaian
	KarpegNo            *string `json:"karpeg_no,omitempty" db:"karpeg_no"`
	KarpegFile          *string `json:"karpeg_file,omitempty" db:"karpeg_file"`
	TaspenNo            *string `json:"taspen_no,omitempty" db:"taspen_no"`
	NPWP                *string `json:"npwp,omitempty" db:"npwp"`
	BPJSSehatan         *string `json:"bpjs_kesehatan,omitempty" db:"bpjs_kesehatan"`
	BPJSKetenagakerjaan *string `json:"bpjs_ketenagakerjaan,omitempty" db:"bpjs_ketenagakerjaan"`
	KKNo                *string `json:"kk_no,omitempty" db:"kk_no"`
	KKFile              *string `json:"kk_file,omitempty" db:"kk_file"`
	KTPNo               *string `json:"ktp_no,omitempty" db:"ktp_no"`
	KTPFile             *string `json:"ktp_file,omitempty" db:"ktp_file"`

	// Integrasi
	SikepID *string `json:"sikep_id,omitempty" db:"sikep_id"`

	// Audit
	IsActive  bool       `json:"is_active" db:"is_active"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	CreatedBy *uuid.UUID `json:"created_by,omitempty" db:"created_by"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty" db:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by,omitempty" db:"deleted_by"`

	// Relations
	Satker      *Satker         `json:"satker,omitempty"`
	Jabatan     *Jabatan        `json:"jabatan,omitempty"`
	UnitKerja   *UnitKerja      `json:"unit_kerja,omitempty"`
	Golongan    *Golongan       `json:"golongan,omitempty"`
	Eselon      *Eselon         `json:"eselon,omitempty"`
	Agama       *RefAgama       `json:"agama,omitempty"`
	StatusKawin *RefStatusKawin `json:"status_kawin,omitempty"`
}

// RiwayatPangkat - dengan field baru
type RiwayatPangkat struct {
	ID            uuid.UUID            `json:"id" db:"id"`
	PegawaiID     uuid.UUID            `json:"pegawai_id" db:"pegawai_id"`
	GolonganID    uuid.UUID            `json:"golongan_id" db:"golongan_id"`
	Pangkat       string               `json:"pangkat" db:"pangkat"`
	TMT           time.Time            `json:"tmt" db:"tmt"`
	NomorSK       string               `json:"nomor_sk" db:"nomor_sk"`
	TanggalSK     time.Time            `json:"tanggal_sk" db:"tanggal_sk"`
	Pejabat       string               `json:"pejabat" db:"pejabat"`
	FileSK        *string              `json:"file_sk,omitempty" db:"file_sk"`
	GajiPokok     float64              `json:"gaji_pokok" db:"gaji_pokok"`
	IsTerakhir    bool                 `json:"is_terakhir" db:"is_terakhir"`
	JenisKenaikan *JenisKenaikanPangkat `json:"jenis_kenaikan,omitempty" db:"jenis_kenaikan"`
	MasaKerjaTahun int                  `json:"masa_kerja_tahun" db:"masa_kerja_tahun"`
	MasaKerjaBulan int                  `json:"masa_kerja_bulan" db:"masa_kerja_bulan"`
	CreatedAt     time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at" db:"updated_at"`
	CreatedBy     *uuid.UUID           `json:"created_by,omitempty" db:"created_by"`

	// Relations
	Golongan *Golongan `json:"golongan,omitempty"`
}

// RiwayatJabatan - dengan field baru
type RiwayatJabatan struct {
	ID           uuid.UUID     `json:"id" db:"id"`
	PegawaiID    uuid.UUID     `json:"pegawai_id" db:"pegawai_id"`
	JabatanID    *uuid.UUID    `json:"jabatan_id,omitempty" db:"jabatan_id"`
	UnitKerjaID  *uuid.UUID    `json:"unit_kerja_id,omitempty" db:"unit_kerja_id"`
	SatkerID     *uuid.UUID    `json:"satker_id,omitempty" db:"satker_id"`
	NamaJabatan  string        `json:"nama_jabatan" db:"nama_jabatan"`
	TMT          time.Time     `json:"tmt" db:"tmt"`
	NomorSK      string        `json:"nomor_sk" db:"nomor_sk"`
	TanggalSK    time.Time     `json:"tanggal_sk" db:"tanggal_sk"`
	Pejabat      string        `json:"pejabat" db:"pejabat"`
	FileSK       *string       `json:"file_sk,omitempty" db:"file_sk"`
	IsTerakhir   bool          `json:"is_terakhir" db:"is_terakhir"`
	JenisJabatan *JenisJabatan `json:"jenis_jabatan,omitempty" db:"jenis_jabatan"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" db:"updated_at"`
	CreatedBy    *uuid.UUID    `json:"created_by,omitempty" db:"created_by"`

	// Relations
	Jabatan   *Jabatan   `json:"jabatan,omitempty"`
	UnitKerja *UnitKerja `json:"unit_kerja,omitempty"`
	Satker    *Satker    `json:"satker,omitempty"`
}

// RiwayatPendidikan
type RiwayatPendidikan struct {
	ID            uuid.UUID     `json:"id" db:"id"`
	PegawaiID     uuid.UUID     `json:"pegawai_id" db:"pegawai_id"`
	PendidikanID  uuid.UUID     `json:"pendidikan_id" db:"pendidikan_id"`
	NamaInstitusi string        `json:"nama_institusi" db:"nama_institusi"`
	Jurusan       *string       `json:"jurusan,omitempty" db:"jurusan"`
	TahunMasuk    int           `json:"tahun_masuk" db:"tahun_masuk"`
	TahunLulus    int           `json:"tahun_lulus" db:"tahun_lulus"`
	NomorIjazah   *string       `json:"nomor_ijazah,omitempty" db:"nomor_ijazah"`
	TanggalIjazah *time.Time    `json:"tanggal_ijazah,omitempty" db:"tanggal_ijazah"`
	FileIjazah    *string       `json:"file_ijazah,omitempty" db:"file_ijazah"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`
	CreatedBy     *uuid.UUID    `json:"created_by,omitempty" db:"created_by"`

	// Relations
	Pendidikan *RefPendidikan `json:"pendidikan,omitempty"`
}

// Keluarga
type Keluarga struct {
	ID            uuid.UUID      `json:"id" db:"id"`
	PegawaiID     uuid.UUID      `json:"pegawai_id" db:"pegawai_id"`
	Hubungan      StatusKeluarga `json:"hubungan" db:"hubungan"`
	Nama          string         `json:"nama" db:"nama"`
	TempatLahir   *string        `json:"tempat_lahir,omitempty" db:"tempat_lahir"`
	TanggalLahir  *time.Time     `json:"tanggal_lahir,omitempty" db:"tanggal_lahir"`
	JenisKelamin  *string        `json:"jenis_kelamin,omitempty" db:"jenis_kelamin"`
	NIK           *string        `json:"nik,omitempty" db:"nik"`
	Pendidikan    *string        `json:"pendidikan,omitempty" db:"pendidikan"`
	Pekerjaan     *string        `json:"pekerjaan,omitempty" db:"pekerjaan"`
	IsTanggungan  bool           `json:"is_tanggungan" db:"is_tanggungan"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
	CreatedBy     *uuid.UUID     `json:"created_by,omitempty" db:"created_by"`
}

// TemplateDokumen
type TemplateDokumen struct {
	ID         uuid.UUID              `json:"id" db:"id"`
	Kode       string                 `json:"kode" db:"kode"`
	Nama       string                 `json:"nama" db:"nama"`
	Tipe       string                 `json:"tipe" db:"tipe"`
	KontenHTML string                 `json:"konten_html" db:"konten_html"`
	Placeholders map[string]string     `json:"placeholders,omitempty" db:"placeholders"`
	IsActive   bool                   `json:"is_active" db:"is_active"`
	CreatedAt  time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at" db:"updated_at"`
}

// ==================== RBAC MODELS ====================

// AppRole
type AppRole struct {
	ID          uuid.UUID `json:"id" db:"id"`
	RoleCode    string    `json:"role_code" db:"role_code"`
	Nama        string    `json:"nama" db:"nama"`
	AppSource   string    `json:"app_source" db:"app_source"` // portal, master-data, kepegawaian
	Deskripsi   string    `json:"deskripsi,omitempty" db:"deskripsi"`
	IsSystem    bool      `json:"is_system" db:"is_system"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// AppPermission
type AppPermission struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Kode        string    `json:"kode" db:"kode"`
	Nama        string    `json:"nama" db:"nama"`
	Resource    string    `json:"resource" db:"resource"`
	Action      string    `json:"action" db:"action"`
	AppSource   string    `json:"app_source" db:"app_source"`
	Deskripsi   string    `json:"deskripsi,omitempty" db:"deskripsi"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// RolePermission
type RolePermission struct {
	ID           uuid.UUID `json:"id" db:"id"`
	RoleID       uuid.UUID `json:"role_id" db:"role_id"`
	PermissionID uuid.UUID `json:"permission_id" db:"permission_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// UserAppRole
type UserAppRole struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	UserID     string     `json:"user_id" db:"user_id"` // Keycloak user ID
	RoleID     uuid.UUID  `json:"role_id" db:"role_id"`
	UnitKerjaID *uuid.UUID `json:"unit_kerja_id,omitempty" db:"unit_kerja_id"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
}

// ==================== AUDIT MODELS ====================

// AuditLog
type AuditLog struct {
	ID            uuid.UUID              `json:"id" db:"id"`
	UserID        *string                `json:"user_id,omitempty" db:"user_id"`
	Username      *string                `json:"username,omitempty" db:"username"`
	Action         string                 `json:"action" db:"action"`
	Resource       string                 `json:"resource" db:"resource"`
	ResourceID     *uuid.UUID             `json:"resource_id,omitempty" db:"resource_id"`
	IPAddress      *string                `json:"ip_address,omitempty" db:"ip_address"`
	UserAgent      *string                `json:"user_agent,omitempty" db:"user_agent"`
	Changes        map[string]interface{} `json:"changes,omitempty" db:"changes"`
	Status         string                 `json:"status" db:"status"`
	ErrorMessage   *string                `json:"error_message,omitempty" db:"error_message"`
	CreatedAt      time.Time              `json:"created_at" db:"created_at"`
}

// ==================== REQUEST/RESPONSE DTOs ====================

// PaginationRequest
type PaginationRequest struct {
	Page       int                    `json:"page" query:"page"`
	PerPage    int                    `json:"per_page" query:"per_page"`
	SortBy     string                 `json:"sort_by" query:"sort_by"`
	SortOrder  string                 `json:"sort_order" query:"sort_order"`
	Search     string                 `json:"search" query:"search"`
	Filters    map[string]interface{} `json:"filters" query:"filters"`
}

// PaginationResponse
type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	TotalPages int         `json:"total_pages"`
}

// APIResponse
type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     *APIError   `json:"error,omitempty"`
	RequestID string      `json:"request_id,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
}

// APIError
type APIError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// LoginRequest
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse
type LoginResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    int64   `json:"expires_in"`
	User         UserDTO `json:"user"`
}

// UserDTO
type UserDTO struct {
	ID          string          `json:"id"`
	Username    string          `json:"username"`
	Email       string          `json:"email"`
	NamaLengkap string          `json:"nama_lengkap"`
	Roles       []string        `json:"roles"`
	Permissions map[string]bool `json:"permissions,omitempty"`
}

// DropdownOption - Generic dropdown response
type DropdownOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// DashboardSummary - Portal dashboard
type DashboardSummary struct {
	TotalPegawai      int                    `json:"total_pegawai"`
	PegawaiAktif      int                    `json:"pegawai_aktif"`
	PegawaiPNS        int                    `json:"pegawai_pns"`
	PegawaiCPNS       int                    `json:"pegawai_cpns"`
	PegawaiPPPK       int                    `json:"pegawai_pppk"`
	PegawaiHonorer    int                    `json:"pegawai_honorer"`
	PerUnitKerja      []StatistikPerUnitKerja `json:"per_unit_kerja"`
	PerGolongan       []StatistikPerGolongan  `json:"per_golongan"`
	AktivitasTerakhir []AuditLog              `json:"aktivitas_terakhir"`
	AkanPensiun       []PegawaiAkanPensiun    `json:"akan_pensiun"`
}

// StatistikPerUnitKerja
type StatistikPerUnitKerja struct {
	UnitKerja string `json:"unit_kerja"`
	Total     int    `json:"total"`
}

// StatistikPerGolongan
type StatistikPerGolongan struct {
	Golongan string `json:"golongan"`
	Total    int    `json:"total"`
}

// PegawaiAkanPensiun
type PegawaiAkanPensiun struct {
	ID                uuid.UUID  `json:"id"`
	NIP               string     `json:"nip"`
	NamaLengkap       string     `json:"nama_lengkap"`
	Jabatan           string     `json:"jabatan"`
	Golongan          string     `json:"golongan"`
	TanggalLahir      time.Time  `json:"tanggal_lahir"`
	Usia              int        `json:"usia"`
	TanggalPensiun    time.Time  `json:"tanggal_pensiun"`
	HariMenujuPensiun int        `json:"hari_menuju_pensiun"`
}
