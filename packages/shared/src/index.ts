// ============================================
// @sikerma/shared - Shared Utilities Package
// ============================================
// Package untuk utilities, types, constants, dan API client
// yang digunakan bersama oleh semua aplikasi SIKERMA

// ============================================
// API Client
// ============================================
export {
  ApiClient,
  ApiError,
  apiClient,
  useApiClient,
} from "./api/client"

export type {
  ApiClientOptions,
  RequestOptions,
} from "./api/client"

// ============================================
// Types
// ============================================
export type {
  // Basic types
  UUID,
  Timestamp,
  PaginationMeta,

  // Master Data types
  Satker,
  Jabatan,
  Golongan,
  UnitKerja,
  Eselon,
  RefPendidikan,
  RefAgama,
  RefStatusKawin,

  // Kepegawaian types
  Pegawai,
  RiwayatPangkat,
  RiwayatJabatan,
  RiwayatPendidikan,
  Keluarga,

  // RBAC types
  AppRole,
  AppPermission,

  // Audit types
  AuditLog,

  // API Response types
  ApiResponse,
  PaginatedResponse,
  ApiError as ApiErrorType,

  // User types
  UserInfo,

  // Form types
  FormDataField,
  DropdownOption,

  // Statistik types
  StatistikKepegawaian,
  StatistikPerGolongan,
  StatistikPerJabatan,
} from "./types"

// ============================================
// Utilities
// ============================================
// Date utilities
export {
  formatDate,
  formatDateTime,
  formatTime,
  formatRelativeTime,
  formatAge,
  isValidDate,
} from "./utils"

// Number utilities
export {
  formatNumber,
  formatCurrency,
  formatDecimal,
} from "./utils"

// String utilities
export {
  truncate,
  slugify,
  capitalize,
  getInitials,
  generateId,
} from "./utils"

// NIP/NIK utilities
export {
  validateNIP,
  formatNIP,
  parseNIP,
  validateNIK,
  formatNIK,
} from "./utils"

// File utilities
export {
  getFileExtension,
  formatFileSize,
  isImageFile,
  isPdfFile,
} from "./utils"

// Array utilities
export {
  groupBy,
  uniqueBy,
  sortBy,
} from "./utils"

// Color utilities
export { stringToColor } from "./utils"

// LocalStorage utilities
export {
  getLocalStorage,
  setLocalStorage,
  removeLocalStorage,
  clearLocalStorage,
} from "./utils"

// ============================================
// Constants
// ============================================
export {
  // API Configuration
  API_BASE_URL,
  API_ENDPOINTS,
  HTTP_STATUS_CODES,

  // App Configuration
  APP_PORTS,

  // Pagination
  PAGINATION_DEFAULTS,

  // File Upload
  FILE_UPLOAD,

  // Domain Constants
  PEGAWAI_STATUS,
  JENIS_KELAMIN,
  STATUS_KELUARGA,
} from "./constants"

// ============================================
// Validation Schemas (Zod)
// ============================================
export {
  // Master Data Schemas
  SatkerSchema,
  JabatanSchema,
  GolonganSchema,
  UnitKerjaSchema,
  EselonSchema,

  // Kepegawaian Schemas
  PegawaiSchema,
  CreatePegawaiSchema,
  UpdatePegawaiSchema,
  RiwayatPangkatSchema,
  RiwayatJabatanSchema,
  RiwayatPendidikanSchema,
  KeluargaSchema,

  // Pagination Schemas
  PaginationParamsSchema,
  PaginationResponseSchema,

  // API Response Schemas
  ApiResponseSchema,
  ApiErrorSchema,

  // Filter Schemas
  PegawaiFilterSchema,
} from "./validations/schemas"

// ============================================
// Providers (TanStack Query)
// ============================================
export { QueryProvider, getQueryClient } from "./providers"
