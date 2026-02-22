// API Configuration
export const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3003/api/v1"

// API Endpoints
export const API_ENDPOINTS = {
  // Health
  HEALTH: "/health",

  // Auth
  AUTH_LOGIN: "/auth/login",
  AUTH_LOGOUT: "/auth/logout",
  AUTH_ME: "/auth/me",

  // Master Data
  MASTER_DATA_SATKER: "/master-data/satker",
  MASTER_DATA_SATKER_DROPDOWN: "/master-data/satker/dropdown",
  MASTER_DATA_JABATAN: "/master-data/jabatan",
  MASTER_DATA_JABATAN_DROPDOWN: "/master-data/jabatan/dropdown",
  MASTER_DATA_GOLONGAN: "/master-data/golongan",
  MASTER_DATA_GOLONGAN_DROPDOWN: "/master-data/golongan/dropdown",
  MASTER_DATA_UNIT_KERJA: "/master-data/unit-kerja",
  MASTER_DATA_UNIT_KERJA_DROPDOWN: "/master-data/unit-kerja/dropdown",
  MASTER_DATA_ESELON: "/master-data/eselon",
  MASTER_DATA_ESELON_DROPDOWN: "/master-data/eselon/dropdown",

  // Kepegawaian
  PEGAWAI: "/kepegawaian/pegawai",
  PEGAWAI_DETAIL: (id: string) => `/kepegawaian/pegawai/${id}`,
  PEGAWAI_RIWAYAT_PANGKAT: (id: string) => `/kepegawaian/pegawai/${id}/riwayat-pangkat`,
  PEGAWAI_RIWAYAT_JABATAN: (id: string) => `/kepegawaian/pegawai/${id}/riwayat-jabatan`,
  PEGAWAI_RIWAYAT_PENDIDIKAN: (id: string) => `/kepegawaian/pegawai/${id}/riwayat-pendidikan`,
  PEGAWAI_KELUARGA: (id: string) => `/kepegawaian/pegawai/${id}/keluarga`,
  PEGAWAI_FOTO: (id: string) => `/kepegawaian/pegawai/${id}/upload-foto`,
  PEGAWAI_SK: (id: string, tipe: string) => `/kepegawaian/pegawai/${id}/upload-sk/${tipe}`,
  KEPEGAWAIAN_STATISTIK: "/kepegawaian/statistik",
  KEPEGAWAIAN_STATISTIK_PANGKAT: "/kepegawaian/statistik/pangkat",
  KEPEGAWAIAN_STATISTIK_JABATAN: "/kepegawaian/statistik/jabatan",

  // RBAC
  RBAC_ROLES: "/rbac/roles",
  RBAC_PERMISSIONS: "/rbac/permissions",

  // Audit Logs
  AUDIT_LOGS: "/audit-logs",

  // PDF
  PDF_GENERATE: "/pdf/generate",
  PDF_TEMPLATES: "/pdf/templates",
} as const

// HTTP Status Codes
export const HTTP_STATUS_CODES = {
  OK: 200,
  CREATED: 201,
  NO_CONTENT: 204,
  BAD_REQUEST: 400,
  UNAUTHORIZED: 401,
  FORBIDDEN: 403,
  NOT_FOUND: 404,
  CONFLICT: 409,
  UNPROCESSABLE_ENTITY: 422,
  INTERNAL_SERVER_ERROR: 500,
  SERVICE_UNAVAILABLE: 503,
} as const

// App Ports
export const APP_PORTS = {
  PORTAL: process.env.NEXT_PUBLIC_PORTAL_PORT || "3000",
  MASTER_DATA: process.env.NEXT_PUBLIC_MASTER_DATA_PORT || "3001",
  KEGAWAIAN: process.env.NEXT_PUBLIC_KEPEGAWAIAN_PORT || "3002",
  BACKEND: process.env.BACKEND_PORT || "3003",
  GOTENBERG: process.env.NEXT_PUBLIC_GOTENBERG_PORT || "3100",
} as const

// Pagination defaults
export const PAGINATION_DEFAULTS = {
  PAGE: 1,
  LIMIT: 20,
  LIMIT_OPTIONS: [10, 20, 50, 100],
} as const

// File upload limits
export const FILE_UPLOAD = {
  MAX_SIZE: 5 * 1024 * 1024, // 5MB
  ALLOWED_TYPES: [
    "image/jpeg",
    "image/png",
    "application/pdf",
  ],
  ALLOWED_EXTENSIONS: [".jpg", ".jpeg", ".png", ".pdf"],
} as const

// Pegawai status
export const PEGAWAI_STATUS = {
  AKTIF: "aktif",
  PENSIUN: "pensiun",
  MUTASI: "mutasi",
} as const

// Jenis kelamin
export const JENIS_KELAMIN = {
  L: "Laki-laki",
  P: "Perempuan",
} as const

// Status keluarga
export const STATUS_KELUARGA = {
  SUAMI: "suami",
  ISTRI: "istri",
  ANAK: "anak",
} as const