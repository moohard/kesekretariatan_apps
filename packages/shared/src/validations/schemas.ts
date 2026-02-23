import { z } from "zod"

// ============================================
// ENUM VALUES
// ============================================

// Status kepegawaian (kategori)
const STATUS_PEGAWAI_VALUES = ["PNS", "CPNS", "PPPK", "HONORER"] as const

// Status kerja
const STATUS_KERJA_VALUES = [
  "aktif",
  "cuti",
  "pensiun",
  "mutasi_keluar",
  "mutasi_masuk",
  "meninggal",
  "pemberhentian",
] as const

// Jenis kelamin
const JENIS_KELAMIN_VALUES = ["L", "P"] as const

// Status keluarga (hubungan)
const STATUS_KELUARGA_VALUES = ["Suami", "Istri", "Anak", "Ayah", "Ibu"] as const

// Jenis jabatan
const JENIS_JABATAN_VALUES = [
  "struktural",
  "fungsional_tertentu",
  "fungsional_umum",
  "pelaksana",
] as const

// Jenis kenaikan pangkat
const JENIS_KENAIKAN_PANGKAT_VALUES = [
  "reguler",
  "pilihan",
  "penyesuaian_ijazah",
  "lainnya",
] as const

// Kategori golongan non-PNS
const KATEGORI_GOLONGAN_NON_PNS_VALUES = ["Honorer", "Kontrak"] as const

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
  email: z.string().email("Format email tidak valid").optional(),
  is_active: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
  created_by: z.string().uuid().nullable().optional(),
  updated_by: z.string().uuid().nullable().optional(),
})

export const JabatanSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1, "Kode jabatan wajib diisi"),
  nama: z.string().min(1, "Nama jabatan wajib diisi"),
  eselon_id: z.string().uuid().nullable().optional(),
  kelas: z.string().optional(),
  jenis: z.enum(JENIS_JABATAN_VALUES).nullable().optional(),
  is_active: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
})

export const GolonganSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1, "Kode golongan wajib diisi"),
  nama: z.string().min(1, "Nama pangkat wajib diisi"),
  ruang: z.string().min(1, "Ruang wajib diisi"),
  angka: z.number().int().min(1).max(17),
  min_pangkat: z.number().int().optional(),
  max_pangkat: z.number().int().optional(),
  is_active: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
})

export const GolonganNonPNSSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1, "Kode golongan wajib diisi"),
  nama: z.string().min(1, "Nama golongan wajib diisi"),
  kategori: z.enum(KATEGORI_GOLONGAN_NON_PNS_VALUES),
  urutan: z.number().int().min(1),
  keterangan: z.string().nullable().optional(),
  is_active: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
})

// Combined view untuk dropdown golongan (PNS + Non-PNS)
export const GolonganAllSchema = z.object({
  id: z.string().uuid(),
  kode: z.string(),
  nama: z.string(),
  ruang: z.string().nullable(),
  angka: z.number().int(),
  kategori: z.enum(["PNS", "Honorer", "Kontrak"]),
  min_pangkat: z.number().int().nullable(),
  max_pangkat: z.number().int().nullable(),
})

export const UnitKerjaSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1, "Kode unit kerja wajib diisi"),
  nama: z.string().min(1, "Nama unit kerja wajib diisi"),
  singkatan: z.string().max(20).optional(),
  parent_id: z.string().uuid().nullable().optional(),
  satker_id: z.string().uuid().nullable().optional(),
  is_active: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
})

export const EselonSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1, "Kode eselon wajib diisi"),
  nama: z.string().min(1, "Nama eselon wajib diisi"),
  level: z.number().int().min(1).max(5).optional(),
  tunjangan: z.number().min(0).optional(),
  is_active: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
})

export const RefPendidikanSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1),
  nama: z.string().min(1),
  tingkat: z.string().optional(),
  urutan: z.number().int().optional(),
  is_active: z.boolean().default(true),
})

export const RefAgamaSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1),
  nama: z.string().min(1),
  is_active: z.boolean().default(true),
})

export const RefStatusKawinSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1),
  nama: z.string().min(1),
  is_active: z.boolean().default(true),
})

// ============================================
// KEPEGAWAIAN SCHEMAS
// ============================================

// NIP validation helper (18 digit)
const NIP_SCHEMA = z
  .string()
  .length(18, "NIP harus 18 digit")
  .regex(/^\d{18}$/, "NIP harus berupa angka")

// NIK validation helper (16 digit)
const NIK_SCHEMA = z
  .string()
  .length(16, "NIK harus 16 digit")
  .regex(/^\d{16}$/, "NIK harus berupa angka")
  .optional()
  .nullable()

// NPWP validation helper (16 digit)
const NPWP_SCHEMA = z
  .string()
  .regex(/^\d{15,16}$/, "NPWP harus 15 atau 16 digit")
  .optional()
  .nullable()

export const PegawaiSchema = z.object({
  id: z.string().uuid(),
  nip: NIP_SCHEMA,
  nip_lama: z.string().max(9).optional().nullable(),

  // Biodata
  nama_lengkap: z.string().min(1, "Nama lengkap wajib diisi"),
  gelar_depan: z.string().max(20).optional().nullable(),
  gelar_belakang: z.string().max(20).optional().nullable(),
  tempat_lahir: z.string().min(1, "Tempat lahir wajib diisi"),
  tanggal_lahir: z.string().or(z.date()),
  jenis_kelamin: z.enum(JENIS_KELAMIN_VALUES, {
    required_error: "Jenis kelamin wajib dipilih",
  }),
  agama_id: z.string().uuid("Agama wajib dipilih"),
  status_kawin_id: z.string().uuid("Status perkawinan wajib dipilih"),
  nik: NIK_SCHEMA,
  email: z.string().email("Format email tidak valid").optional().nullable(),
  telepon: z.string().max(20).optional().nullable(),
  alamat: z.string().optional().nullable(),
  alamat_domisili: z.string().optional().nullable(),
  foto: z.string().optional().nullable(),

  // Posisi
  satker_id: z.string().uuid("Satuan kerja wajib dipilih"),
  jabatan_id: z.string().uuid().optional().nullable(),
  unit_kerja_id: z.string().uuid().optional().nullable(),
  golongan_id: z.string().uuid().optional().nullable(),
  eselon_id: z.string().uuid().optional().nullable(),

  // Status
  status_pegawai: z.enum(STATUS_PEGAWAI_VALUES).default("PNS"),
  status_kerja: z.enum(STATUS_KERJA_VALUES).default("aktif"),

  // TMT
  tmt_cpns: z.string().or(z.date()).optional().nullable(),
  tmt_pns: z.string().or(z.date()).optional().nullable(),
  tmt_jabatan: z.string().or(z.date()).optional().nullable(),
  tmt_pangkat_terakhir: z.string().or(z.date()).optional().nullable(),
  tmt_jabatan_terakhir: z.string().or(z.date()).optional().nullable(),

  // Dokumen kepegawaian
  karpeg_no: z.string().max(50).optional().nullable(),
  karpeg_file: z.string().max(500).optional().nullable(),
  taspen_no: z.string().max(50).optional().nullable(),
  npwp: NPWP_SCHEMA,
  bpjs_kesehatan: z.string().max(30).optional().nullable(),
  bpjs_ketenagakerjaan: z.string().max(30).optional().nullable(),
  kk_no: z.string().max(30).optional().nullable(),
  kk_file: z.string().max(500).optional().nullable(),
  ktp_no: NIK_SCHEMA,
  ktp_file: z.string().max(500).optional().nullable(),

  // Integrasi
  sikep_id: z.string().max(50).optional().nullable(),

  // Audit
  is_active: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
  created_by: z.string().uuid().nullable().optional(),
  updated_by: z.string().uuid().nullable().optional(),
  deleted_at: z.string().or(z.date()).nullable().optional(),
  deleted_by: z.string().uuid().nullable().optional(),
})

// Schema untuk create pegawai (tanpa id dan audit fields)
export const CreatePegawaiSchema = PegawaiSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
  updated_by: true,
  deleted_at: true,
  deleted_by: true,
  foto: true,
}).partial({
  // Field yang boleh kosong saat create
  nip_lama: true,
  gelar_depan: true,
  gelar_belakang: true,
  alamat_domisili: true,
  karpeg_no: true,
  karpeg_file: true,
  taspen_no: true,
  npwp: true,
  bpjs_kesehatan: true,
  bpjs_ketenagakerjaan: true,
  kk_no: true,
  kk_file: true,
  ktp_no: true,
  ktp_file: true,
  sikep_id: true,
  eselon_id: true,
  tmt_cpns: true,
  tmt_pns: true,
  tmt_jabatan: true,
  tmt_pangkat_terakhir: true,
  tmt_jabatan_terakhir: true,
})

// Schema untuk update pegawai
export const UpdatePegawaiSchema = PegawaiSchema.partial().omit({
  id: true,
  nip: true, // NIP tidak bisa diubah
  created_at: true,
  updated_at: true,
  created_by: true,
  deleted_at: true,
  deleted_by: true,
})

// Schema untuk filter pegawai
export const PegawaiFilterSchema = z.object({
  page: z.coerce.number().int().min(1).default(1),
  limit: z.coerce.number().int().min(1).max(100).default(20),
  search: z.string().optional(),
  sort_by: z.string().optional(),
  sort_order: z.enum(["asc", "desc"]).default("asc"),
  satker_id: z.string().uuid().optional(),
  jabatan_id: z.string().uuid().optional(),
  unit_kerja_id: z.string().uuid().optional(),
  golongan_id: z.string().uuid().optional(),
  eselon_id: z.string().uuid().optional(),
  status_pegawai: z.enum(STATUS_PEGAWAI_VALUES).optional(),
  status_kerja: z.enum(STATUS_KERJA_VALUES).optional(),
  jenis_kelamin: z.enum(JENIS_KELAMIN_VALUES).optional(),
  is_active: z.coerce.boolean().optional(),
})

// ============================================
// RIWAYAT SCHEMAS
// ============================================

export const RiwayatPangkatSchema = z.object({
  id: z.string().uuid(),
  pegawai_id: z.string().uuid(),
  golongan_id: z.string().uuid("Golongan wajib dipilih"),
  pangkat: z.string().min(1, "Nama pangkat wajib diisi"),
  tmt: z.string().or(z.date()),
  nomor_sk: z.string().min(1, "Nomor SK wajib diisi"),
  tanggal_sk: z.string().or(z.date()),
  pejabat: z.string().min(1, "Pejabat penetap wajib diisi"),
  file_sk: z.string().max(500).optional().nullable(),
  gaji_pokok: z.number().min(0).default(0),
  is_terakhir: z.boolean().default(false),

  // Field baru
  jenis_kenaikan: z.enum(JENIS_KENAIKAN_PANGKAT_VALUES).optional().nullable(),
  masa_kerja_tahun: z.number().int().min(0).default(0),
  masa_kerja_bulan: z.number().int().min(0).max(11).default(0),

  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
  created_by: z.string().uuid().nullable().optional(),
})

export const CreateRiwayatPangkatSchema = RiwayatPangkatSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
})

export const UpdateRiwayatPangkatSchema = RiwayatPangkatSchema.partial().omit({
  id: true,
  pegawai_id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
})

export const RiwayatJabatanSchema = z.object({
  id: z.string().uuid(),
  pegawai_id: z.string().uuid(),
  jabatan_id: z.string().uuid().optional().nullable(),
  unit_kerja_id: z.string().uuid().optional().nullable(),
  satker_id: z.string().uuid().optional().nullable(),
  nama_jabatan: z.string().min(1, "Nama jabatan wajib diisi"),
  tmt: z.string().or(z.date()),
  nomor_sk: z.string().min(1, "Nomor SK wajib diisi"),
  tanggal_sk: z.string().or(z.date()),
  pejabat: z.string().min(1, "Pejabat penetap wajib diisi"),
  file_sk: z.string().max(500).optional().nullable(),
  is_terakhir: z.boolean().default(false),

  // Field baru
  jenis_jabatan: z.enum(JENIS_JABATAN_VALUES).optional().nullable(),

  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
  created_by: z.string().uuid().nullable().optional(),
})

export const CreateRiwayatJabatanSchema = RiwayatJabatanSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
})

export const UpdateRiwayatJabatanSchema = RiwayatJabatanSchema.partial().omit({
  id: true,
  pegawai_id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
})

export const RiwayatPendidikanSchema = z.object({
  id: z.string().uuid(),
  pegawai_id: z.string().uuid(),
  pendidikan_id: z.string().uuid("Jenjang pendidikan wajib dipilih"),
  nama_institusi: z.string().min(1, "Nama institusi wajib diisi"),
  jurusan: z.string().optional().nullable(),
  tahun_masuk: z.number().int().min(1900).max(new Date().getFullYear()),
  tahun_lulus: z.number().int().min(1900).max(new Date().getFullYear() + 10),
  nomor_ijazah: z.string().optional().nullable(),
  tanggal_ijazah: z.string().or(z.date()).optional().nullable(),
  file_ijazah: z.string().max(500).optional().nullable(),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
  created_by: z.string().uuid().nullable().optional(),
})

export const CreateRiwayatPendidikanSchema = RiwayatPendidikanSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
})

export const UpdateRiwayatPendidikanSchema = RiwayatPendidikanSchema.partial().omit({
  id: true,
  pegawai_id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
})

export const KeluargaSchema = z.object({
  id: z.string().uuid(),
  pegawai_id: z.string().uuid(),
  hubungan: z.enum(STATUS_KELUARGA_VALUES, {
    required_error: "Hubungan keluarga wajib dipilih",
  }),
  nama: z.string().min(1, "Nama wajib diisi"),
  tempat_lahir: z.string().optional().nullable(),
  tanggal_lahir: z.string().or(z.date()).optional().nullable(),
  jenis_kelamin: z.enum(JENIS_KELAMIN_VALUES).optional().nullable(),
  nik: NIK_SCHEMA,
  pendidikan: z.string().optional().nullable(),
  pekerjaan: z.string().optional().nullable(),
  is_tanggungan: z.boolean().default(true),
  created_at: z.string().or(z.date()).optional(),
  updated_at: z.string().or(z.date()).optional(),
  created_by: z.string().uuid().nullable().optional(),
})

export const CreateKeluargaSchema = KeluargaSchema.omit({
  id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
})

export const UpdateKeluargaSchema = KeluargaSchema.partial().omit({
  id: true,
  pegawai_id: true,
  created_at: true,
  updated_at: true,
  created_by: true,
})

// ============================================
// PAGINATION & RESPONSE SCHEMAS
// ============================================

export const PaginationParamsSchema = z.object({
  page: z.coerce.number().int().min(1).default(1),
  limit: z.coerce.number().int().min(1).max(100).default(20),
  search: z.string().optional(),
  sort_by: z.string().optional(),
  sort_order: z.enum(["asc", "desc"]).default("asc"),
})

export const PaginationMetaSchema = z.object({
  page: z.number(),
  limit: z.number(),
  total: z.number(),
  total_pages: z.number(),
})

export const PaginationResponseSchema = <T extends z.ZodTypeAny>(itemSchema: T) =>
  z.object({
    success: z.boolean(),
    data: z.array(itemSchema),
    pagination: PaginationMetaSchema,
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
    error: z
      .object({
        code: z.string(),
        message: z.string(),
        details: z.record(z.unknown()).optional(),
      })
      .optional(),
    request_id: z.string().optional(),
    timestamp: z.string().optional(),
  })

export const ApiErrorSchema = z.object({
  success: z.literal(false),
  error: z.object({
    code: z.string(),
    message: z.string(),
    details: z.record(z.unknown()).optional(),
  }),
  request_id: z.string().optional(),
  timestamp: z.string().optional(),
})

// ============================================
// DROPDOWN OPTION SCHEMA
// ============================================

export const DropdownOptionSchema = z.object({
  value: z.string(),
  label: z.string(),
})

// ============================================
// TYPE EXPORTS FROM SCHEMAS
// ============================================

// Extract literal union types from const arrays
export type StatusPegawai = (typeof STATUS_PEGAWAI_VALUES)[number]
export type StatusKerja = (typeof STATUS_KERJA_VALUES)[number]
export type JenisJabatan = (typeof JENIS_JABATAN_VALUES)[number]
export type JenisKenaikanPangkat = (typeof JENIS_KENAIKAN_PANGKAT_VALUES)[number]
export type StatusKeluarga = (typeof STATUS_KELUARGA_VALUES)[number]
export type KategoriGolonganNonPNS = (typeof KATEGORI_GOLONGAN_NON_PNS_VALUES)[number]
export type JenisKelamin = (typeof JENIS_KELAMIN_VALUES)[number]
