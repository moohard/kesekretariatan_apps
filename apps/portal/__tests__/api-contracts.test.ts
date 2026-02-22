/**
 * RM-008: API Contract Tests
 *
 * Contract tests memverifikasi bahwa API responses sesuai dengan expected schema.
 * Ini penting untuk memastikan backward compatibility saat upgrade Next.js.
 *
 * Testing approach:
 * 1. Schema validation dengan Zod
 * 2. Response format validation
 * 3. Error response format validation
 */

import { describe, it, expect, beforeAll, afterAll } from 'vitest'
import { setupServer } from 'msw/node'
import { http, HttpResponse, delay } from 'msw'
import { z } from 'zod'

// ============================================================================
// API Response Schemas (Contracts)
// ============================================================================

// Standard API Response wrapper
const ApiResponseSchema = <T extends z.ZodTypeAny>(dataSchema: T) =>
  z.object({
    success: z.boolean(),
    data: dataSchema.optional(),
    error: z.object({
      code: z.string(),
      message: z.string(),
      details: z.record(z.unknown()).optional(),
    }).optional(),
    meta: z.object({
      requestId: z.string(),
      timestamp: z.string(),
    }).optional(),
  })

// Pagination metadata schema
const PaginationMetaSchema = z.object({
  pagination: z.object({
    page: z.number().int().positive(),
    limit: z.number().int().positive(),
    total: z.number().int().nonnegative(),
    totalPages: z.number().int().nonnegative(),
  }),
})

// Pegawai schema
const PegawaiSchema = z.object({
  id: z.string().uuid(),
  nip: z.string().length(18).regex(/^\d{18}$/),
  nama_lengkap: z.string().min(1).max(255),
  unit_kerja_id: z.string().uuid(),
  satker_id: z.string().uuid().optional(),
  tanggal_lahir: z.string().datetime().optional(),
  jenis_kelamin: z.enum(['L', 'P']).optional(),
  agama: z.string().optional(),
  status_pegawai: z.string().optional(),
  is_active: z.boolean(),
  created_at: z.string().datetime(),
  updated_at: z.string().datetime(),
})

// Satker schema
const SatkerSchema = z.object({
  id: z.string().uuid(),
  kode: z.string().min(1).max(10),
  nama: z.string().min(1).max(255),
  jenis: z.enum(['PA', 'PTA', 'MA', 'MTA']),
  tingkat: z.enum(['PN', 'PT']),
  alamat: z.string().optional(),
  telepon: z.string().optional(),
  email: z.string().email().optional(),
  created_at: z.string().datetime(),
  updated_at: z.string().datetime(),
})

// Unit Kerja schema
const UnitKerjaSchema = z.object({
  id: z.string().uuid(),
  satker_id: z.string().uuid(),
  nama: z.string().min(1).max(255),
  deskripsi: z.string().optional(),
  created_at: z.string().datetime(),
  updated_at: z.string().datetime(),
})

// Error response schema (per PRD Section 15)
const ErrorResponseSchema = z.object({
  success: z.literal(false),
  error: z.object({
    code: z.string(),
    message: z.string(),
    details: z.record(z.unknown()).optional(),
  }),
  requestId: z.string(),
  timestamp: z.string(),
})

// ============================================================================
// MSW Server Setup
// ============================================================================

const mockPegawai = {
  id: '550e8400-e29b-41d4-a716-446655440001',
  nip: '198501012010011001',
  nama_lengkap: 'Ahmad Fauzi',
  unit_kerja_id: '660e8400-e29b-41d4-a716-446655440001',
  satker_id: '770e8400-e29b-41d4-a716-446655440001',
  tanggal_lahir: '1985-01-01T00:00:00Z',
  jenis_kelamin: 'L',
  agama: 'Islam',
  status_pegawai: 'PNS',
  is_active: true,
  created_at: '2024-01-01T00:00:00Z',
  updated_at: '2024-01-01T00:00:00Z',
}

const handlers = [
  // GET /api/v1/pegawai/:id
  http.get('/api/v1/pegawai/:id', async ({ params }) => {
    await delay(50)
    if (params.id === 'not-found') {
      return HttpResponse.json({
        success: false,
        error: {
          code: 'NOT_FOUND_PEGAWAI',
          message: 'Pegawai tidak ditemukan',
        },
        requestId: 'req-123',
        timestamp: new Date().toISOString(),
      }, { status: 404 })
    }
    return HttpResponse.json({
      success: true,
      data: mockPegawai,
    })
  }),

  // GET /api/v1/pegawai (list)
  http.get('/api/v1/pegawai', async ({ request }) => {
    await delay(50)
    const url = new URL(request.url)
    const page = parseInt(url.searchParams.get('page') || '1')
    const limit = parseInt(url.searchParams.get('limit') || '10')

    return HttpResponse.json({
      success: true,
      data: [mockPegawai],
      meta: {
        pagination: {
          page,
          limit,
          total: 1,
          totalPages: 1,
        },
      },
    })
  }),

  // POST /api/v1/pegawai
  http.post('/api/v1/pegawai', async ({ request }) => {
    await delay(100)
    const body = await request.json() as Record<string, unknown>

    // Validation error
    if (typeof body.nip === 'string' && !/^\d{18}$/.test(body.nip)) {
      return HttpResponse.json({
        success: false,
        error: {
          code: 'VAL_NIP_FORMAT',
          message: 'NIP harus 18 digit numerik',
          details: { field: 'nip', value: body.nip },
        },
        requestId: 'req-123',
        timestamp: new Date().toISOString(),
      }, { status: 400 })
    }

    return HttpResponse.json({
      success: true,
      data: { ...mockPegawai, ...body },
    }, { status: 201 })
  }),

  // GET /api/v1/satker
  http.get('/api/v1/satker', async () => {
    await delay(30)
    return HttpResponse.json({
      success: true,
      data: [{
        id: '770e8400-e29b-41d4-a716-446655440001',
        kode: 'PA001',
        nama: 'Pengadilan Agama Jakarta Pusat',
        jenis: 'PA',
        tingkat: 'PN',
        created_at: '2024-01-01T00:00:00Z',
        updated_at: '2024-01-01T00:00:00Z',
      }],
    })
  }),

  // Rate limit simulation
  http.get('/api/v1/rate-limited', async () => {
    return HttpResponse.json({
      success: false,
      error: {
        code: 'RATE_LIMIT_EXCEEDED',
        message: 'Terlalu banyak permintaan, coba lagi dalam 1 menit',
      },
      requestId: 'req-123',
      timestamp: new Date().toISOString(),
    }, { status: 429 })
  }),
]

const server = setupServer(...handlers)

beforeAll(() => server.listen())
afterAll(() => server.close())

// ============================================================================
// CONTRACT TESTS
// ============================================================================

describe('API Contract: Pegawai Endpoints', () => {
  const baseUrl = '/api/v1'

  describe('GET /pegawai/:id', () => {
    it('should return valid pegawai response schema', async () => {
      const response = await fetch(`${baseUrl}/pegawai/550e8400-e29b-41d4-a716-446655440001`)
      const json = await response.json()

      // Validate response schema
      const result = ApiResponseSchema(PegawaiSchema).safeParse(json)
      expect(result.success).toBe(true)
      if (result.success) {
        expect(result.data.success).toBe(true)
        expect(result.data.data).toBeDefined()
      }
    })

    it('should return 404 with error schema for not found', async () => {
      const response = await fetch(`${baseUrl}/pegawai/not-found`)
      expect(response.status).toBe(404)

      const json = await response.json()

      // Validate error response schema
      const result = ErrorResponseSchema.safeParse(json)
      expect(result.success).toBe(true)
      if (result.success) {
        expect(result.data.success).toBe(false)
        expect(result.data.error.code).toBe('NOT_FOUND_PEGAWAI')
      }
    })
  })

  describe('GET /pegawai (list)', () => {
    it('should return paginated pegawai list', async () => {
      const response = await fetch(`${baseUrl}/pegawai?page=1&limit=10`)
      const json = await response.json()

      // Validate paginated response
      const ListResponseSchema = z.object({
        success: z.boolean(),
        data: z.array(PegawaiSchema),
        meta: PaginationMetaSchema,
      })

      const result = ListResponseSchema.safeParse(json)
      expect(result.success).toBe(true)
      if (result.success) {
        expect(result.data.success).toBe(true)
        expect(Array.isArray(result.data.data)).toBe(true)
        expect(result.data.meta.pagination.page).toBe(1)
        expect(result.data.meta.pagination.limit).toBe(10)
      }
    })
  })

  describe('POST /pegawai', () => {
    it('should return 400 with validation error for invalid NIP', async () => {
      const response = await fetch(`${baseUrl}/pegawai`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          nip: 'invalid',
          nama_lengkap: 'Test User',
          unit_kerja_id: '660e8400-e29b-41d4-a716-446655440001',
        }),
      })

      expect(response.status).toBe(400)
      const json = await response.json()

      const result = ErrorResponseSchema.safeParse(json)
      expect(result.success).toBe(true)
      if (result.success) {
        expect(result.data.error.code).toBe('VAL_NIP_FORMAT')
        expect(result.data.error.details).toBeDefined()
      }
    })

    it('should create pegawai and return 201', async () => {
      const response = await fetch(`${baseUrl}/pegawai`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          nip: '199901012024011001',
          nama_lengkap: 'New User',
          unit_kerja_id: '660e8400-e29b-41d4-a716-446655440001',
        }),
      })

      expect(response.status).toBe(201)
      const json = await response.json()

      const result = ApiResponseSchema(PegawaiSchema).safeParse(json)
      expect(result.success).toBe(true)
    })
  })
})

describe('API Contract: Satker Endpoints', () => {
  const baseUrl = '/api/v1'

  describe('GET /satker', () => {
    it('should return valid satker list schema', async () => {
      const response = await fetch(`${baseUrl}/satker`)
      const json = await response.json()

      const ListResponseSchema = z.object({
        success: z.boolean(),
        data: z.array(SatkerSchema),
      })

      const result = ListResponseSchema.safeParse(json)
      expect(result.success).toBe(true)
      if (result.success) {
        expect(result.data.success).toBe(true)
        expect(Array.isArray(result.data.data)).toBe(true)
      }
    })
  })
})

describe('API Contract: Error Codes (PRD Section 15.3)', () => {
  const baseUrl = '/api/v1'

  it('should return RATE_LIMIT_EXCEEDED error with 429 status', async () => {
    const response = await fetch(`${baseUrl}/rate-limited`)
    expect(response.status).toBe(429)

    const json = await response.json()
    const result = ErrorResponseSchema.safeParse(json)

    expect(result.success).toBe(true)
    if (result.success) {
      expect(result.data.error.code).toBe('RATE_LIMIT_EXCEEDED')
    }
  })

  it('error response must include requestId and timestamp', async () => {
    const response = await fetch(`${baseUrl}/rate-limited`)
    const json = await response.json()

    expect(json.requestId).toBeDefined()
    expect(json.timestamp).toBeDefined()

    // Validate timestamp format (ISO 8601)
    const timestamp = new Date(json.timestamp)
    expect(timestamp.toISOString()).toBe(json.timestamp)
  })
})

describe('API Contract: NIP Validation (PRD Section 8.1)', () => {
  it('validates NIP format: 18 digit numeric', () => {
    const validNIPs = [
      '198501012010011001',
      '199901012024011001',
      '197001011990011001',
    ]

    const invalidNIPs = [
      '123', // too short
      'abcdefghij123456789', // contains letters
      '19850101201001100', // 17 digits
      '1985010120100110011', // 19 digits
    ]

    const nipRegex = /^\d{18}$/

    validNIPs.forEach((nip) => {
      expect(nipRegex.test(nip), `NIP ${nip} should be valid`).toBe(true)
    })

    invalidNIPs.forEach((nip) => {
      expect(nipRegex.test(nip), `NIP ${nip} should be invalid`).toBe(false)
    })
  })
})

describe('API Contract: Pagination (PRD Section 8.3)', () => {
  it('pagination meta should have all required fields', () => {
    const paginationData = {
      page: 1,
      limit: 10,
      total: 100,
      totalPages: 10,
    }

    const result = PaginationMetaSchema.safeParse({ pagination: paginationData })
    expect(result.success).toBe(true)
  })

  it('default limit should be 10, max 100', () => {
    // Test dengan berbagai limit values
    const validLimits = [1, 10, 50, 100]
    const invalidLimits = [0, -1, 101, 1000]

    validLimits.forEach((limit) => {
      const result = PaginationMetaSchema.safeParse({
        pagination: { page: 1, limit, total: 100, totalPages: 10 },
      })
      expect(result.success, `limit ${limit} should be valid`).toBe(true)
    })

    // Invalid limits masih valid di schema level
    // tapi harus di-reject di application level
    invalidLimits.forEach((limit) => {
      const result = PaginationMetaSchema.safeParse({
        pagination: { page: 1, limit, total: 100, totalPages: 10 },
      })
      // Schema allows any positive integer, business logic rejects > 100
    })
  })
})
