package utils

import (
	"regexp"
	"strings"
)

// ============================================
// PII Masking Utilities
// ============================================

// PII fields yang perlu di-mask
var sensitiveFields = map[string]bool{
	"nik":            true,
	"nip":            true,
	"password":       true,
	"alamat":         true,
	"telepon":        true,
	"phone":          true,
	"gaji_pokok":     true,
	"gaji":           true,
	"no_rekening":    true,
	"npwp":           true,
	"no_ktp":         true,
	"tempat_lahir":   true,
	"tanggal_lahir":  true,
	"ibu_kandung":    true,
}

// MaskValue melakukan masking pada nilai field
// Format: "****1234" (menampilkan 4 karakter terakhir)
func MaskValue(value string) string {
	if len(value) <= 4 {
		return "****"
	}
	return "****" + value[len(value)-4:]
}

// MaskEmail melakukan masking pada email
// Format: "j***@example.com"
func MaskEmail(email string) string {
	if email == "" {
		return ""
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return MaskValue(email)
	}

	localPart := parts[0]
	domain := parts[1]

	if len(localPart) <= 1 {
		return "*" + "@" + domain
	}

	return string(localPart[0]) + "***@" + domain
}

// MaskPhone melakukan masking pada nomor telepon
// Format: "****7890" (menampilkan 4 digit terakhir)
func MaskPhone(phone string) string {
	// Remove non-digit characters
	re := regexp.MustCompile(`[^0-9]`)
	digits := re.ReplaceAllString(phone, "")

	if len(digits) <= 4 {
		return "****"
	}

	return "****" + digits[len(digits)-4:]
}

// MaskNIK melakukan masking pada NIK (16 digit)
// Format: "****************1234"
func MaskNIK(nik string) string {
	if len(nik) != 16 {
		return MaskValue(nik)
	}
	return "************" + nik[12:]
}

// MaskNIP melakukan masking pada NIP (18 digit)
// Format: "******************1234"
func MaskNIP(nip string) string {
	if len(nip) != 18 {
		return MaskValue(nip)
	}
	return "**************" + nip[14:]
}

// ShouldMaskField mengecek apakah field perlu di-mask
func ShouldMaskField(fieldName string) bool {
	lowerField := strings.ToLower(fieldName)
	_, exists := sensitiveFields[lowerField]
	return exists
}

// MaskSensitiveData melakukan masking pada data sensitif dalam map
func MaskSensitiveData(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}

	masked := make(map[string]interface{})
	for key, value := range data {
		if ShouldMaskField(key) {
			switch v := value.(type) {
			case string:
				switch key {
				case "nik", "no_ktp":
					masked[key] = MaskNIK(v)
				case "nip":
					masked[key] = MaskNIP(v)
				case "email":
					masked[key] = MaskEmail(v)
				case "telepon", "phone":
					masked[key] = MaskPhone(v)
				default:
					masked[key] = MaskValue(v)
				}
			default:
				masked[key] = "****"
			}
		} else {
			// Recursively mask nested maps
			if nested, ok := value.(map[string]interface{}); ok {
				masked[key] = MaskSensitiveData(nested)
			} else {
				masked[key] = value
			}
		}
	}

	return masked
}

// MaskPII melakukan masking pada data berdasarkan field list
// Field list bisa digunakan untuk custom masking
func MaskPII(data map[string]interface{}, fieldsToMask []string) map[string]interface{} {
	if data == nil {
		return nil
	}

	masked := make(map[string]interface{})
	fieldsMap := make(map[string]bool)

	for _, field := range fieldsToMask {
		fieldsMap[strings.ToLower(field)] = true
	}

	for key, value := range data {
		lowerKey := strings.ToLower(key)
		shouldMask := fieldsMap[lowerKey] || sensitiveFields[lowerKey]

		if shouldMask {
			switch v := value.(type) {
			case string:
				switch lowerKey {
				case "nik", "no_ktp":
					masked[key] = MaskNIK(v)
				case "nip":
					masked[key] = MaskNIP(v)
				case "email":
					masked[key] = MaskEmail(v)
				case "telepon", "phone":
					masked[key] = MaskPhone(v)
				default:
					masked[key] = MaskValue(v)
				}
			default:
				masked[key] = "****"
			}
		} else {
			if nested, ok := value.(map[string]interface{}); ok {
				masked[key] = MaskPII(nested, fieldsToMask)
			} else {
				masked[key] = value
			}
		}
	}

	return masked
}
