# @sikerma/shared

Paket shared utilities untuk SIKERMA. Berisi API client, types, validations, dan utility functions.

## Fitur

- API Client dengan error handling
- TypeScript types untuk domain models
- Zod schemas untuk validation
- Utility functions (formatting, NIP validation, dll)
- Constants (API endpoints, status codes)

## Penggunaan

### API Client

```tsx
import { apiClient, ApiError } from "@sikerma/shared"

// Set auth token
apiClient.setAuthToken("your-jwt-token")

// GET request
const data = await apiClient.get("/kepegawaian/pegawai", {
  params: { page: 1, limit: 20 }
})

// POST request
const created = await apiClient.post("/kepegawaian/pegawai", formData)

// Error handling
try {
  const result = await apiClient.get("/endpoint")
} catch (error) {
  if (error instanceof ApiError) {
    console.error(error.code, error.message, error.request_id)
  }
}
```

### Types

```tsx
import type { Pegawai, PaginationMeta, ApiResponse } from "@sikerma/shared"

const pegawai: Pegawai = {
  id: "uuid",
  nip: "198001012005011001",
  nama: "John Doe",
  // ...
}

const response: ApiResponse<Pegawai[]> = {
  success: true,
  data: [pegawai],
  pagination: { page: 1, limit: 20, total: 100, total_pages: 5 }
}
```

### Validations

```tsx
import { PegawaiSchema, CreatePegawaiSchema } from "@sikerma/shared"

// Parse and validate data
const result = PegawaiSchema.parse(data)

// Partial validation
const updateData = CreatePegajiSchema.parse(formData)
```

### Utils

```tsx
import {
  formatDate,
  formatDateTime,
  formatCurrency,
  validateNIP,
  formatNIP,
  getInitials
} from "@sikerma/shared"

// Date formatting
const formatted = formatDate(new Date()) // "22 Februari 2026"
const withTime = formatDateTime(new Date()) // "22 Februari 2026 14:30"

// Currency formatting
const rupiah = formatCurrency(1500000) // "Rp 1.500.000"

// NIP validation
const { valid, message } = validateNIP("198001012005011001")
if (!valid) {
  console.error(message)
}

// NIP formatting
const formattedNIP = formatNIP("198001012005011001")
// "19800101 200501 1 001 000001"

// Get initials
const initials = getInitials("John Doe") // "JD"
```

### Constants

```tsx
import {
  API_BASE_URL,
  API_ENDPOINTS,
  PAGINATION_DEFAULTS,
  FILE_UPLOAD
} from "@sikerma/shared"

// API Base URL
console.log(API_BASE_URL) // "http://localhost:3003/api/v1"

// API Endpoints
const url = API_ENDPOINTS.PEGAWAI_DETAIL("uuid")

// Pagination defaults
const { PAGE, LIMIT } = PAGINATION_DEFAULTS

// File upload limits
console.log(FILE_UPLOAD.MAX_SIZE) // 5242880 (5MB)
console.log(FILE_UPLOAD.ALLOWED_TYPES) // ["image/jpeg", "image/png", "application/pdf"]
```

## API Service Examples

### Pegawai Service

```tsx
import { apiClient, API_ENDPOINTS } from "@sikerma/shared"

export const pegawaiService = {
  list: async (params: { page?: number; limit?: number; search?: string }) => {
    return apiClient.get(API_ENDPOINTS.PEGAWAI, { params })
  },

  get: async (id: string) => {
    return apiClient.get(API_ENDPOINTS.PEGAWAI_DETAIL(id))
  },

  create: async (data: any) => {
    return apiClient.post(API_ENDPOINTS.PEGAWAI, data)
  },

  update: async (id: string, data: any) => {
    return apiClient.put(API_ENDPOINTS.PEGAWAI_DETAIL(id), data)
  },

  delete: async (id: string) => {
    return apiClient.delete(API_ENDPOINTS.PEGAWAI_DETAIL(id))
  },

  uploadFoto: async (id: string, file: File) => {
    return apiClient.upload(API_ENDPOINTS.PEGAWAI_FOTO(id), file)
  }
}
```

### Master Data Service

```tsx
import { apiClient, API_ENDPOINTS } from "@sikerma/shared"

export const masterDataService = {
  getSatkerDropdown: async () => {
    return apiClient.get(API_ENDPOINTS.MASTER_DATA_SATKER_DROPDOWN)
  },

  getJabatanDropdown: async () => {
    return apiClient.get(API_ENDPOINTS.MASTER_DATA_JABATAN_DROPDOWN)
  },

  getGolonganDropdown: async () => {
    return apiClient.get(API_ENDPOINTS.MASTER_DATA_GOLONGAN_DROPDOWN)
  },

  getUnitKerjaDropdown: async () => {
    return apiClient.get(API_ENDPOINTS.MASTER_DATA_UNIT_KERJA_DROPDOWN)
  }
}
```

## Error Handling

```tsx
import { ApiError, HTTP_STATUS_CODES } from "@sikerma/shared"

try {
  const result = await apiClient.get("/endpoint")
} catch (error) {
  if (error instanceof ApiError) {
    switch (error.code) {
      case HTTP_STATUS_CODES.UNAUTHORIZED:
        // Redirect to login
        break
      case HTTP_STATUS_CODES.FORBIDDEN:
        // Show permission denied message
        break
      case HTTP_STATUS_CODES.NOT_FOUND:
        // Show not found message
        break
      case HTTP_STATUS_CODES.INTERNAL_SERVER_ERROR:
        // Show generic error message
        break
    }
  }
}
```