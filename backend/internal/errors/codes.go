package errors

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

// ============================================
// Error Codes (sesuai PRD Section 15.3)
// ============================================

// Validation Errors (400)
const (
	ValNIPFormat       = "VAL_NIP_FORMAT"
	ValNIPDuplicate    = "VAL_NIP_DUPLICATE"
	ValNIKFormat       = "VAL_NIK_FORMAT"
	ValRequiredField   = "VAL_REQUIRED_FIELD"
	ValInvalidFormat   = "VAL_INVALID_FORMAT"
	ValFileSize        = "VAL_FILE_SIZE"
	ValFileType        = "VAL_FILE_TYPE"
	ValInvalidDate     = "VAL_INVALID_DATE"
	ValInvalidEmail    = "VAL_INVALID_EMAIL"
	ValInvalidPhone    = "VAL_INVALID_PHONE"
	ValOutOfRange      = "VAL_OUT_OF_RANGE"
	ValDuplicateEntry  = "VAL_DUPLICATE_ENTRY"
)

// Authentication Errors (401)
const (
	AuthInvalidToken   = "AUTH_INVALID_TOKEN"
	AuthTokenExpired   = "AUTH_TOKEN_EXPIRED"
	AuthLoginFailed    = "AUTH_LOGIN_FAILED"
	AuthSessionExpired = "AUTH_SESSION_EXPIRED"
	AuthInvalidCredentials = "AUTH_INVALID_CREDENTIALS"
)

// Authorization Errors (403)
const (
	AuthzForbidden        = "AUTHZ_FORBIDDEN"
	AuthzRoleInsufficient = "AUTHZ_ROLE_INSUFFICIENT"
	AuthzUnitAccessDenied = "AUTHZ_UNIT_ACCESS_DENIED"
)

// Not Found Errors (404)
const (
	NotFoundPegawai    = "NOT_FOUND_PEGAWAI"
	NotFoundSatker     = "NOT_FOUND_SATKER"
	NotFoundJabatan    = "NOT_FOUND_JABATAN"
	NotFoundGolongan   = "NOT_FOUND_GOLONGAN"
	NotFoundUnitKerja  = "NOT_FOUND_UNIT_KERJA"
	NotFoundRiwayat    = "NOT_FOUND_RIWAYAT"
	NotFoundDocument   = "NOT_FOUND_DOCUMENT"
	NotFoundUser       = "NOT_FOUND_USER"
	NotFoundResource   = "NOT_FOUND_RESOURCE"
)

// Conflict Errors (409)
const (
	ConflictNIPExists     = "CONFLICT_NIP_EXISTS"
	ConflictNIKExists     = "CONFLICT_NIK_EXISTS"
	ConflictEmailExists   = "CONFLICT_EMAIL_EXISTS"
	ConflictRelation      = "CONFLICT_RELATION"
	ConflictState         = "CONFLICT_STATE"
)

// Rate Limiting Errors (429)
const (
	RateLimitExceeded       = "RATE_LIMIT_EXCEEDED"
	LoginRateLimitExceeded  = "LOGIN_RATE_LIMIT_EXCEEDED"
	UploadRateLimitExceeded = "UPLOAD_RATE_LIMIT_EXCEEDED"
	APIRateLimitExceeded    = "API_RATE_LIMIT_EXCEEDED"
)

// System Errors (500)
const (
	SysDatabaseError   = "SYS_DATABASE_ERROR"
	SysInternalError   = "SYS_INTERNAL_ERROR"
	SysExternalService = "SYS_EXTERNAL_SERVICE"
	SysFileUpload      = "SYS_FILE_UPLOAD"
	SysPDFGeneration   = "SYS_PDF_GENERATION"
)

// Service Unavailable (503)
const (
	SvcUnavailable     = "SVC_UNAVAILABLE"
	SvcDatabaseDown    = "SVC_DATABASE_DOWN"
	SvcAuthDown        = "SVC_AUTH_DOWN"
)

// ============================================
// Error Messages (Indonesian)
// ============================================

var errorMessages = map[string]string{
	// Validation
	ValNIPFormat:       "Format NIP tidak valid. NIP harus 18 digit angka",
	ValNIPDuplicate:    "NIP sudah terdaftar dalam sistem",
	ValNIKFormat:       "Format NIK tidak valid. NIK harus 16 digit angka",
	ValRequiredField:   "Field wajib tidak boleh kosong",
	ValInvalidFormat:   "Format data tidak valid",
	ValFileSize:        "Ukuran file melebihi batas maksimal",
	ValFileType:        "Tipe file tidak didukung",
	ValInvalidDate:     "Format tanggal tidak valid",
	ValInvalidEmail:    "Format email tidak valid",
	ValInvalidPhone:    "Format nomor telepon tidak valid",
	ValOutOfRange:      "Nilai di luar jangkauan yang diizinkan",
	ValDuplicateEntry:  "Data sudah ada dalam sistem",

	// Authentication
	AuthInvalidToken:       "Token autentikasi tidak valid",
	AuthTokenExpired:       "Token autentikasi sudah kadaluarsa",
	AuthLoginFailed:        "Login gagal. Periksa kembali kredensial Anda",
	AuthSessionExpired:     "Sesi sudah kadaluarsa. Silakan login kembali",
	AuthInvalidCredentials: "Kredensial tidak valid",

	// Authorization
	AuthzForbidden:        "Anda tidak memiliki akses ke resource ini",
	AuthzRoleInsufficient: "Role Anda tidak memiliki permission yang diperlukan",
	AuthzUnitAccessDenied: "Akses ke unit kerja ini ditolak",

	// Not Found
	NotFoundPegawai:   "Data pegawai tidak ditemukan",
	NotFoundSatker:    "Data Satker tidak ditemukan",
	NotFoundJabatan:   "Data jabatan tidak ditemukan",
	NotFoundGolongan:  "Data golongan tidak ditemukan",
	NotFoundUnitKerja: "Data unit kerja tidak ditemukan",
	NotFoundRiwayat:   "Data riwayat tidak ditemukan",
	NotFoundDocument:  "Dokumen tidak ditemukan",
	NotFoundUser:      "User tidak ditemukan",
	NotFoundResource:  "Resource tidak ditemukan",

	// Conflict
	ConflictNIPExists:   "NIP sudah digunakan oleh pegawai lain",
	ConflictNIKExists:   "NIK sudah digunakan oleh pegawai lain",
	ConflictEmailExists: "Email sudah digunakan oleh pegawai lain",
	ConflictRelation:    "Relasi data tidak dapat dihapus karena masih digunakan",
	ConflictState:       "Status data tidak dapat diubah",

	// Rate Limiting
	RateLimitExceeded:       "Terlalu banyak permintaan, coba lagi dalam 1 menit",
	LoginRateLimitExceeded:  "Terlalu banyak percobaan login, coba lagi dalam 15 menit",
	UploadRateLimitExceeded: "Terlalu banyak upload, coba lagi dalam 1 menit",
	APIRateLimitExceeded:    "Rate limit API exceeded",

	// System
	SysDatabaseError:   "Terjadi kesalahan pada database",
	SysInternalError:   "Terjadi kesalahan internal sistem",
	SysExternalService: "Layanan eksternal tidak tersedia",
	SysFileUpload:      "Gagal mengunggah file",
	SysPDFGeneration:   "Gagal membuat dokumen PDF",

	// Service Unavailable
	SvcUnavailable:  "Layanan tidak tersedia",
	SvcDatabaseDown: "Database tidak tersedia",
	SvcAuthDown:     "Layanan autentikasi tidak tersedia",
}

// ============================================
// Error Response Types
// ============================================

// ErrorResponse adalah format standar untuk error response
type ErrorResponse struct {
	Success   bool                   `json:"success"`
	Error     ErrorDetail            `json:"error"`
	RequestID string                 `json:"requestId"`
	Timestamp string                 `json:"timestamp"`
}

// ErrorDetail adalah detail error
type ErrorDetail struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ============================================
// Error Functions
// ============================================

// NewError membuat error baru dengan code
func NewError(code string, details ...map[string]interface{}) *ErrorResponse {
	message, ok := errorMessages[code]
	if !ok {
		message = "Terjadi kesalahan"
	}

	var detailMap map[string]interface{}
	if len(details) > 0 {
		detailMap = details[0]
	}

	return &ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Code:    code,
			Message: message,
			Details: detailMap,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// GetMessage mendapatkan pesan error dari code
func GetMessage(code string) string {
	if message, ok := errorMessages[code]; ok {
		return message
	}
	return "Terjadi kesalahan"
}

// ToFiberResponse mengkonversi ErrorResponse ke Fiber response
func (e *ErrorResponse) ToFiberResponse(c fiber.Ctx, statusCode int) error {
	e.RequestID = c.Locals("requestId").(string)
	return c.Status(statusCode).JSON(e)
}

// ============================================
// Helper Functions untuk HTTP Status
// ============================================

// GetStatusCode mendapatkan HTTP status code dari error code
func GetStatusCode(code string) int {
	switch {
	case code[:3] == "VAL":
		return fiber.StatusBadRequest
	case code[:4] == "AUTH":
		return fiber.StatusUnauthorized
	case code[:5] == "AUTHZ":
		return fiber.StatusForbidden
	case code[:7] == "NOT_FND" || code[:3] == "NOT":
		return fiber.StatusNotFound
	case code[:8] == "CONFLICT" || code[:3] == "CON":
		return fiber.StatusConflict
	case code[:4] == "RATE" || code[:5] == "LOGIN" || code[:6] == "UPLOAD" || code[:3] == "API":
		return fiber.StatusTooManyRequests
	case code[:3] == "SYS":
		return fiber.StatusInternalServerError
	case code[:3] == "SVC":
		return fiber.StatusServiceUnavailable
	default:
		return fiber.StatusInternalServerError
	}
}

// BadRequest membuat error 400
func BadRequest(code string, details ...map[string]interface{}) *ErrorResponse {
	return NewError(code, details...)
}

// Unauthorized membuat error 401
func Unauthorized(code string, details ...map[string]interface{}) *ErrorResponse {
	return NewError(code, details...)
}

// Forbidden membuat error 403
func Forbidden(code string, details ...map[string]interface{}) *ErrorResponse {
	return NewError(code, details...)
}

// NotFound membuat error 404
func NotFound(code string, details ...map[string]interface{}) *ErrorResponse {
	return NewError(code, details...)
}

// Conflict membuat error 409
func Conflict(code string, details ...map[string]interface{}) *ErrorResponse {
	return NewError(code, details...)
}

// TooManyRequests membuat error 429
func TooManyRequests(code string, details ...map[string]interface{}) *ErrorResponse {
	return NewError(code, details...)
}

// InternalError membuat error 500
func InternalError(code string, details ...map[string]interface{}) *ErrorResponse {
	return NewError(code, details...)
}

// ServiceUnavailable membuat error 503
func ServiceUnavailable(code string, details ...map[string]interface{}) *ErrorResponse {
	return NewError(code, details...)
}
