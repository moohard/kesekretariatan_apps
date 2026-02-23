// ============================================
// DOMAIN TYPES
// ============================================

// UUID type alias
export type UUID = string

// Timestamp type
export type Timestamp = string | Date

// Pagination
export interface PaginationMeta {
  page: number
  limit: number
  total: number
  total_pages: number
}

// ============================================
// MASTER DATA TYPES
// ============================================

export interface Satker {
  id: UUID
  kode: string
  nama: string
  parent_id: UUID | null
  level: number
  alamat: string | null
  telepon: string | null
  email: string | null
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
  created_by: UUID | null
  updated_by: UUID | null
}

export interface Jabatan {
  id: UUID
  kode: string
  nama: string
  eselon_id: UUID | null
  kelas: string | null
  jenis: "struktural" | "fungsional_tertentu" | "fungsional_umum" | "pelaksana" | null
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
  // Relations
  eselon?: Eselon
}

export interface Golongan {
  id: UUID
  kode: string
  nama: string
  ruang: string
  angka: number
  min_pangkat: number
  max_pangkat: number
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
}

// Golongan untuk pegawai non-PNS (Honorer, Kontrak, PPPK)
export interface GolonganNonPNS {
  id: UUID
  kode: string
  nama: string
  kategori: "Honorer" | "Kontrak"
  urutan: number
  keterangan: string | null
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
}

// Combined view untuk dropdown golongan (PNS + Non-PNS)
export interface GolonganAll {
  id: UUID
  kode: string
  nama: string
  ruang: string | null
  angka: number
  kategori: "PNS" | "Honorer" | "Kontrak"
  min_pangkat: number | null
  max_pangkat: number | null
}

export interface UnitKerja {
  id: UUID
  kode: string
  nama: string
  singkatan: string | null
  parent_id: UUID | null
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
}

export interface Eselon {
  id: UUID
  kode: string
  nama: string
  tunjangan: number
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
}

export interface RefPendidikan {
  id: UUID
  kode: string
  nama: string
  tingkat: string
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
}

export interface RefAgama {
  id: UUID
  kode: string
  nama: string
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
}

export interface RefStatusKawin {
  id: UUID
  kode: string
  nama: string
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
}

// ============================================
// KEGAWAAN TYPES
// ============================================

// Status pegawai berdasarkan kategori kepegawaian
export type StatusPegawai = "PNS" | "CPNS" | "PPPK" | "HONORER"

// Status kerja pegawai
export type StatusKerja = "aktif" | "cuti" | "pensiun" | "mutasi_keluar" | "mutasi_masuk" | "meninggal" | "pemberhentian"

export interface Pegawai {
  id: UUID
  nip: string
  nip_lama: string | null  // NIP 9 digit untuk pegawai lama
  nama_lengkap: string
  gelar_depan: string | null
  gelar_belakang: string | null
  tempat_lahir: string
  tanggal_lahir: Timestamp
  jenis_kelamin: "L" | "P"
  agama_id: UUID
  status_kawin_id: UUID
  nik: string | null
  email: string | null
  telepon: string | null
  alamat: string | null
  alamat_domisili: string | null  // Alamat domisili (berbeda dengan KTP)
  foto: string | null
  satker_id: UUID
  jabatan_id: UUID | null
  unit_kerja_id: UUID | null
  golongan_id: UUID | null
  eselon_id: UUID | null  // Untuk pegawai struktural
  status_pegawai: StatusPegawai
  status_kerja: StatusKerja

  // TMT (Terhitung Mulai Tanggal)
  tmt_cpns: Timestamp | null
  tmt_pns: Timestamp | null
  tmt_jabatan: Timestamp | null
  tmt_pangkat_terakhir: Timestamp | null
  tmt_jabatan_terakhir: Timestamp | null

  // Dokumen kepegawaian
  karpeg_no: string | null
  karpeg_file: string | null
  taspen_no: string | null
  npwp: string | null
  bpjs_kesehatan: string | null
  bpjs_ketenagakerjaan: string | null
  kk_no: string | null
  kk_file: string | null
  ktp_no: string | null
  ktp_file: string | null

  // Integrasi
  sikep_id: string | null  // ID di sistem SIKEP MA

  // Audit
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
  created_by: UUID | null
  updated_by: UUID | null
  deleted_at: Timestamp | null
  deleted_by: UUID | null

  // Relations
  satker?: Satker
  jabatan?: Jabatan
  unit_kerja?: UnitKerja
  golongan?: Golongan
  eselon?: Eselon
  agama?: RefAgama
  status_kawin?: RefStatusKawin
}

// Jenis kenaikan pangkat
export type JenisKenaikanPangkat = "reguler" | "pilihan" | "penyesuaian_ijazah" | "lainnya"

// Jenis jabatan
export type JenisJabatan = "struktural" | "fungsional_tertentu" | "fungsional_umum" | "pelaksana"

export interface RiwayatPangkat {
  id: UUID
  pegawai_id: UUID
  golongan_id: UUID
  pangkat: string
  tmt: Timestamp
  nomor_sk: string
  tanggal_sk: Timestamp
  pejabat: string
  file_sk: string | null
  gaji_pokok: number
  is_terakhir: boolean

  // Field baru
  jenis_kenaikan: JenisKenaikanPangkat | null
  masa_kerja_tahun: number
  masa_kerja_bulan: number

  created_at: Timestamp
  updated_at: Timestamp
  created_by: UUID | null

  // Relations
  golongan?: Golongan
}

export interface RiwayatJabatan {
  id: UUID
  pegawai_id: UUID
  jabatan_id: UUID | null
  unit_kerja_id: UUID | null
  satker_id: UUID | null
  nama_jabatan: string
  tmt: Timestamp
  nomor_sk: string
  tanggal_sk: Timestamp
  pejabat: string
  file_sk: string | null
  is_terakhir: boolean

  // Field baru
  jenis_jabatan: JenisJabatan | null

  created_at: Timestamp
  updated_at: Timestamp
  created_by: UUID | null

  // Relations
  jabatan?: Jabatan
  unit_kerja?: UnitKerja
  satker?: Satker
}

export interface RiwayatPendidikan {
  id: UUID
  pegawai_id: UUID
  pendidikan_id: UUID
  nama_institusi: string
  jurusan: string | null
  tahun_masuk: number
  tahun_lulus: number
  nomor_ijazah: string | null
  tanggal_ijazah: Timestamp | null
  file_ijazah: string | null
  created_at: Timestamp
  updated_at: Timestamp
  created_by: UUID | null
  // Relations
  pendidikan?: RefPendidikan
}

// Status hubungan keluarga
export type StatusKeluarga = "Suami" | "Istri" | "Anak" | "Ayah" | "Ibu"

export interface Keluarga {
  id: UUID
  pegawai_id: UUID
  hubungan: StatusKeluarga
  nama: string
  tempat_lahir: string | null
  tanggal_lahir: Timestamp | null
  jenis_kelamin: "L" | "P" | null
  nik: string | null
  pendidikan: string | null
  pekerjaan: string | null
  is_tanggungan: boolean
  created_at: Timestamp
  updated_at: Timestamp
  created_by: UUID | null
}

// ============================================
// RBAC TYPES
// ============================================

export interface AppRole {
  id: UUID
  nama: string
  deskripsi: string
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
}

export interface AppPermission {
  id: UUID
  nama: string
  resource: string
  action: string
  deskripsi: string
  created_at: Timestamp
  updated_at: Timestamp
}

// ============================================
// AUDIT TYPES
// ============================================

export interface AuditLog {
  id: UUID
  user_id: string | null
  username: string | null
  action: string
  resource: string
  resource_id: UUID | null
  ip_address: string | null
  user_agent: string | null
  changes: Record<string, any> | null
  status: string
  error_message: string | null
  created_at: Timestamp
}

// ============================================
// API RESPONSE TYPES
// ============================================

export interface ApiResponse<T = any> {
  success: boolean
  message?: string
  data?: T
  error?: string
  request_id?: string
}

export interface PaginatedResponse<T> extends ApiResponse<T[]> {
  pagination: PaginationMeta
}

export interface ApiError {
  error: true
  message: string
  code: number
  request_id?: string
}

// ============================================
// USER TYPES
// ============================================

export interface UserInfo {
  id: string
  username: string
  email: string
  name: string
  roles: string[]
  given_name?: string
  family_name?: string
}

// ============================================
// FORM TYPES
// ============================================

export type FormDataField = {
  name: string
  label: string
  type: "text" | "email" | "password" | "number" | "select" | "date" | "textarea" | "file"
  placeholder?: string
  required?: boolean
  options?: DropdownOption[]
  validation?: any
}

export type DropdownOption<T = any> = {
  value: T
  label: string
}

// ============================================
// STATISTIK TYPES
// ============================================

export interface StatistikKepegawaian {
  total_pegawai: number
  per_status_pegawai: Record<StatusPegawai, number>
  per_status_kerja: Record<StatusKerja, number>
  pns: number  // PNS + CPNS
  non_pns: number  // PPPK + HONORER
  per_golongan: Record<string, number>
  per_unit_kerja: Record<string, number>
  per_jenis_jabatan: Record<string, number>
}

export interface StatistikPerGolongan {
  golongan: string
  total: number
  pns: number
  non_pns: number
}

export interface StatistikPerJabatan {
  jabatan: string
  total: number
}

export interface StatistikPerUnitKerja {
  unit_kerja: string
  total: number
}

// Pegawai yang akan mencapai BUP (Batas Usia Pensiun)
export interface PegawaiAkanPensiun {
  id: UUID
  nip: string
  nama_lengkap: string
  jabatan: string
  golongan: string
  tanggal_lahir: Timestamp
  usia: number
  tanggal_pensiun: Timestamp  // Tanggal perkiraan pensiun
  hari_menuju_pensiun: number
}

// Dashboard summary untuk Portal
export interface DashboardSummary {
  total_pegawai: number
  pegawai_aktif: number
  pegawai_pns: number
  pegawai_cpns: number
  pegawai_pppk: number
  pegawai_honorer: number
  per_unit_kerja: StatistikPerUnitKerja[]
  per_golongan: StatistikPerGolongan[]
  aktivitas_terakhir: AuditLog[]
  akan_pensiun: PegawaiAkanPensiun[]
}