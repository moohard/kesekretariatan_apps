/**
 * Shared Playwright Constants
 */

// Storage state paths for authenticated tests
export const STORAGE_STATE = {
  ADMIN: '.auth/admin.json',
  USER: '.auth/regular-user.json',
}

// Test data constants
export const TEST_USER = {
  ADMIN: {
    username: 'admin',
    role: 'admin',
    unitKerja: 'all',
  },
  USER: {
    username: 'user',
    role: 'user',
    unitKerja: 'unit-1',
  },
}

// API endpoints for testing
export const API_ENDPOINTS = {
  PEGAWAI: '/api/v1/pegawai',
  SATKER: '/api/v1/satker',
  UNIT_KERJA: '/api/v1/unit-kerja',
  JABATAN: '/api/v1/jabatan',
  AUTH: '/api/v1/auth',
  HEALTH: '/health',
}
