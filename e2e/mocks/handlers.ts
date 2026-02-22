/**
 * MSW (Mock Service Worker) Handlers
 * Untuk mocking API calls dalam unit/integration tests
 */

import { http, HttpResponse, delay } from 'msw'
import { API_ENDPOINTS } from '../e2e/playwright.config'

// Mock data
const mockPegawai = [
  {
    id: '1',
    nip: '198501012010011001',
    nama_lengkap: 'Ahmad Fauzi',
    unit_kerja_id: 'unit-1',
    satker_id: 'satker-1',
    is_active: true,
    created_at: '2024-01-01T00:00:00Z',
    updated_at: '2024-01-01T00:00:00Z',
  },
  {
    id: '2',
    nip: '198702152012012002',
    nama_lengkap: 'Budi Santoso',
    unit_kerja_id: 'unit-1',
    satker_id: 'satker-1',
    is_active: true,
    created_at: '2024-01-01T00:00:00Z',
    updated_at: '2024-01-01T00:00:00Z',
  },
  {
    id: '3',
    nip: '199001032013011003',
    nama_lengkap: 'Citra Dewi',
    unit_kerja_id: 'unit-2',
    satker_id: 'satker-2',
    is_active: true,
    created_at: '2024-01-01T00:00:00Z',
    updated_at: '2024-01-01T00:00:00Z',
  },
]

const mockSatker = [
  {
    id: 'satker-1',
    kode: 'PA001',
    nama: 'Pengadilan Agama Jakarta Pusat',
    jenis: 'PA',
    tingkat: 'PN',
  },
  {
    id: 'satker-2',
    kode: 'PA002',
    nama: 'Pengadilan Agama Jakarta Selatan',
    jenis: 'PA',
    tingkat: 'PN',
  },
]

// Success response wrapper
function successResponse<T>(data: T, meta?: Record<string, unknown>) {
  return HttpResponse.json({
    success: true,
    data,
    ...(meta && { meta }),
  })
}

// Error response wrapper
function errorResponse(code: string, message: string, status = 400) {
  return HttpResponse.json(
    {
      success: false,
      error: {
        code,
        message,
      },
      requestId: 'test-request-id',
      timestamp: new Date().toISOString(),
    },
    { status }
  )
}

// Paginated response wrapper
function paginatedResponse<T>(items: T[], page = 1, limit = 10) {
  const start = (page - 1) * limit
  const end = start + limit
  const paginatedItems = items.slice(start, end)

  return successResponse(paginatedItems, {
    pagination: {
      page,
      limit,
      total: items.length,
      totalPages: Math.ceil(items.length / limit),
    },
  })
}

// API Handlers
export const handlers = [
  // Health check
  http.get(API_ENDPOINTS.HEALTH, () => {
    return HttpResponse.json({ status: 'ok' })
  }),

  // Auth endpoints
  http.post(`${API_ENDPOINTS.AUTH}/login`, async () => {
    await delay(100) // Simulate network delay
    return successResponse({
      accessToken: 'mock-access-token',
      refreshToken: 'mock-refresh-token',
      expiresIn: 900,
    })
  }),

  http.post(`${API_ENDPOINTS.AUTH}/refresh`, async () => {
    await delay(50)
    return successResponse({
      accessToken: 'mock-new-access-token',
      refreshToken: 'mock-new-refresh-token',
      expiresIn: 900,
    })
  }),

  http.post(`${API_ENDPOINTS.AUTH}/logout`, () => {
    return successResponse({ message: 'Logged out successfully' })
  }),

  // Pegawai endpoints
  http.get(API_ENDPOINTS.PEGAWAI, async ({ request }) => {
    await delay(50)
    const url = new URL(request.url)
    const page = parseInt(url.searchParams.get('page') || '1')
    const limit = parseInt(url.searchParams.get('limit') || '10')
    const unitKerjaId = url.searchParams.get('unit_kerja_id')

    let filtered = mockPegawai
    if (unitKerjaId) {
      filtered = mockPegawai.filter((p) => p.unit_kerja_id === unitKerjaId)
    }

    return paginatedResponse(filtered, page, limit)
  }),

  http.get(`${API_ENDPOINTS.PEGAWAI}/:id`, async ({ params }) => {
    await delay(30)
    const pegawai = mockPegawai.find((p) => p.id === params.id)
    if (!pegawai) {
      return errorResponse('NOT_FOUND_PEGAWAI', 'Pegawai tidak ditemukan', 404)
    }
    return successResponse(pegawai)
  }),

  http.post(API_ENDPOINTS.PEGAWAI, async ({ request }) => {
    await delay(100)
    const body = (await request.json()) as Record<string, unknown>

    // Validation: Check NIP format
    if (typeof body.nip !== 'string' || !/^\d{18}$/.test(body.nip)) {
      return errorResponse('VAL_NIP_FORMAT', 'NIP harus 18 digit numerik')
    }

    // Check duplicate NIP
    if (mockPegawai.some((p) => p.nip === body.nip)) {
      return errorResponse('VAL_NIP_DUPLICATE', 'NIP sudah terdaftar')
    }

    const newPegawai = {
      id: `${mockPegawai.length + 1}`,
      ...body,
      is_active: true,
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    }

    return successResponse(newPegawai)
  }),

  http.put(`${API_ENDPOINTS.PEGAWAI}/:id`, async ({ params, request }) => {
    await delay(100)
    const pegawai = mockPegawai.find((p) => p.id === params.id)
    if (!pegawai) {
      return errorResponse('NOT_FOUND_PEGAWAI', 'Pegawai tidak ditemukan', 404)
    }

    const body = (await request.json()) as Record<string, unknown>
    const updated = { ...pegawai, ...body, updated_at: new Date().toISOString() }

    return successResponse(updated)
  }),

  http.delete(`${API_ENDPOINTS.PEGAWAI}/:id`, async ({ params }) => {
    await delay(50)
    const index = mockPegawai.findIndex((p) => p.id === params.id)
    if (index === -1) {
      return errorResponse('NOT_FOUND_PEGAWAI', 'Pegawai tidak ditemukan', 404)
    }

    return successResponse({ message: 'Pegawai berhasil dihapus' })
  }),

  // Satker endpoints
  http.get(API_ENDPOINTS.SATKER, async () => {
    await delay(30)
    return successResponse(mockSatker)
  }),

  http.get(`${API_ENDPOINTS.SATKER}/:id`, async ({ params }) => {
    await delay(20)
    const satker = mockSatker.find((s) => s.id === params.id)
    if (!satker) {
      return errorResponse('NOT_FOUND_SATKER', 'Satker tidak ditemukan', 404)
    }
    return successResponse(satker)
  }),

  // Rate limit simulation
  http.all('/api/rate-limited/*', async ({ request }) => {
    const authHeader = request.headers.get('Authorization')
    if (!authHeader) {
      return errorResponse('AUTH_INVALID_TOKEN', 'Unauthorized', 401)
    }
    return successResponse({ message: 'OK' })
  }),
]

// Export for use in tests
export { mockPegawai, mockSatker }
