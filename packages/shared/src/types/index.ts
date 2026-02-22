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
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
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

export interface Pegawai {
  id: UUID
  nip: string
  nama: string
  gelar_depan: string
  gelar_belakang: string
  tempat_lahir: string
  tanggal_lahir: Timestamp
  jenis_kelamin: "L" | "P"
  agama_id: UUID
  status_kawin_id: UUID
  nik: string | null
  email: string | null
  telepon: string | null
  alamat: string | null
  foto: string | null
  satker_id: UUID
  jabatan_id: UUID | null
  unit_kerja_id: UUID | null
  golongan_id: UUID | null
  status_pegawai: "aktif" | "pensiun" | "mutasi"
  tmt_jabatan: Timestamp | null
  is_pns: boolean
  is_active: boolean
  created_at: Timestamp
  updated_at: Timestamp
  // Relations
  satker?: Satker
  jabatan?: Jabatan
  unit_kerja?: UnitKerja
  golongan?: Golongan
  agama?: RefAgama
  status_kawin?: RefStatusKawin
}

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
  created_at: Timestamp
  updated_at: Timestamp
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
  created_at: Timestamp
  updated_at: Timestamp
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
  pendidikan?: RefPendidikan
}

export interface Keluarga {
  id: UUID
  pegawai_id: UUID
  status_keluarga: "suami" | "istri" | "anak"
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
  per_status: Record<string, number>
  pns: number
  non_pns: number
  per_golongan: Record<string, number>
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