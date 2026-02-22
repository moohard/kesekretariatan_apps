import { z } from "zod"

// Define enum values locally to avoid type issues
const PEGAWAI_STATUS_VALUES = ["aktif", "pensiun", "mutasi"] as const
const JENIS_KELAMIN_VALUES = ["L", "P"] as const
const STATUS_KELUARGA_VALUES = ["suami", "istri", "anak"] as const

// ============================================
// MASTER DATA SCHEMAS
// ============================================

export const SatkerSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1, "Kode wajib diisi"),
  nama: z.string().min(1, "Nama wajib diisi"),
  parent_id: z.string().uuid().nullable().optional(),
  level: z.number().int().min(1),
  alamat: z.string().optional(),
  telepon: z.string().optional(),
  email: z.string().email().optional(),
  is_active: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
})

export const JabatanSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1),
  nama: z.string().min(1),
  eselon_id: z.string().uuid().nullable().optional(),
  kelas: z.string().optional(),
  is_active: z.boolean().default(true),
})

export const GolonganSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1),
  nama: z.string().min(1),
  ruang: z.string().min(1),
  angka: z.number().int().min(1).max(27),
})

export const UnitKerjaSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1),
  nama: z.string().min(1),
  singkatan: z.string().optional(),
  parent_id: z.string().uuid().nullable().optional(),
})

export const EselonSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1),
  nama: z.string().min(1),
  tunjangan: z.number().min(0),
})

// ============================================
// KEGAWAAN SCHEMAS
// ============================================

export const PegawaiSchema = z.object({
  id: z.string().uuid(),
  nip: z.string().length(18, "NIP harus 18 digit"),
  nama: z.string().min(1, "Nama wajib diisi"),
  gelar_depan: z.string().optional(),
  gelar_belakang: z.string().optional(),
  tempat_lahir: z.string().optional(),
  tanggal_lahir: z.string().or(z.date()).optional(),
  jenis_kelamin: z.enum(["L", "P"], { required_error: "Jenis kelamin wajib dipilih" }),
  agama_id: z.string().uuid(),
  status_kawin_id: z.string().uuid(),
  nik: z.string().length(16, "NIK harus 16 digit").optional(),
  email: z.string().email("Email tidak valid").optional(),
  telepon: z.string().optional(),
  alamat: z.string().optional(),
  foto: z.string().url().optional(),
  satker_id: z.string().uuid(),
  jabatan_id: z.string().uuid().optional(),
  unit_kerja_id: z.string().uuid().optional(),
  golongan_id: z.string().uuid().optional(),
  status_pegawai: z.enum(PEGAWAI_STATUS_VALUES).default("aktif"),
  tmt_jabatan: z.string().or(z.date()).optional(),
  is_pns: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
})

export const CreatePegawaiSchema = PegawaiSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
  foto: true,
})

export const UpdatePegawaiSchema = PegawaiSchema.partial().omit({
  id: true,
  nip: true,
  created_at: true,
})

export const RiwayatPangkatSchema = z.object({
  id: z.string().uuid(),
  pegawai_id: z.string().uuid(),
  golongan_id: z.string().uuid(),
  pangkat: z.string().min(1),
  tmt: z.string().or(z.date()),
  nomor_sk: z.string().min(1),
  tanggal_sk: z.string().or(z.date()),
  pejabat: z.string().min(1),
  file_sk: z.string().optional(),
  gaji_pokok: z.number().min(0),
  is_terakhir: z.boolean().default(false),
})

export const RiwayatJabatanSchema = z.object({
  id: z.string().uuid(),
  pegawai_id: z.string().uuid(),
  jabatan_id: z.string().uuid().optional(),
  unit_kerja_id: z.string().uuid().optional(),
  satker_id: z.string().uuid().optional(),
  nama_jabatan: z.string().min(1),
  tmt: z.string().or(z.date()),
  nomor_sk: z.string().min(1),
  tanggal_sk: z.string().or(z.date()),
  pejabat: z.string().min(1),
  file_sk: z.string().optional(),
  is_terakhir: z.boolean().default(false),
})

export const RiwayatPendidikanSchema = z.object({
  id: z.string().uuid(),
  pegawai_id: z.string().uuid(),
  pendidikan_id: z.string().uuid(),
  nama_institusi: z.string().min(1),
  jurusan: z.string().optional(),
  tahun_masuk: z.number().int().min(1900).max(new Date().getFullYear()),
  tahun_lulus: z.number().int().min(1900).max(new Date().getFullYear() + 10),
  nomor_ijazah: z.string().optional(),
  tanggal_ijazah: z.string().or(z.date()).optional(),
  file_ijazah: z.string().optional(),
})

export const KeluargaSchema = z.object({
  id: z.string().uuid(),
  pegawai_id: z.string().uuid(),
  status_keluarga: z.enum(STATUS_KELUARGA_VALUES),
  nama: z.string().min(1),
  tempat_lahir: z.string().optional(),
  tanggal_lahir: z.string().or(z.date()).optional(),
  jenis_kelamin: z.enum(["L", "P"]).optional(),
  nik: z.string().length(16).optional(),
  pendidikan: z.string().optional(),
  pekerjaan: z.string().optional(),
  is_tanggungan: z.boolean().default(false),
})

// ============================================
// PAGINATION SCHEMAS
// ============================================

export const PaginationParamsSchema = z.object({
  page: z.coerce.number().int().min(1).default(1),
  limit: z.coerce.number().int().min(1).max(100).default(20),
  search: z.string().optional(),
  sort_by: z.string().optional(),
  sort_order: z.enum(["asc", "desc"]).default("desc"),
})

export const PaginationResponseSchema = <T extends z.ZodTypeAny>(itemSchema: T) =>
  z.object({
    success: z.boolean(),
    data: z.array(itemSchema),
    pagination: z.object({
      page: z.number(),
      limit: z.number(),
      total: z.number(),
      total_pages: z.number(),
    }),
    request_id: z.string().optional(),
  })

// ============================================
// API RESPONSE SCHEMAS
// ============================================

export const ApiResponseSchema = <T extends z.ZodTypeAny>(dataSchema?: T) =>
  z.object({
    success: z.boolean(),
    message: z.string().optional(),
    data: dataSchema || z.unknown().optional(),
    error: z.string().optional(),
    request_id: z.string().optional(),
  })

export const ApiErrorSchema = z.object({
  error: z.literal(true),
  message: z.string(),
  code: z.number(),
  request_id: z.string().optional(),
})

// ============================================
// FILTER SCHEMAS
// ============================================

export const PegawaiFilterSchema = z.object({
  satker_id: z.string().uuid().optional(),
  jabatan_id: z.string().uuid().optional(),
  golongan_id: z.string().uuid().optional(),
  status_pegawai: z.enum(PEGAWAI_STATUS_VALUES).optional(),
})