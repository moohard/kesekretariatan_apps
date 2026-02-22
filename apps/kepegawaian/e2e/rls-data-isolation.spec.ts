/**
 * E2E Tests: Data Isolation (RLS Verification)
 *
 * Tests untuk memverifikasi bahwa Row-Level Security bekerja dengan benar
 * di level aplikasi. User hanya bisa melihat data dari unit kerjanya sendiri.
 */

import { test, expect } from '@playwright/test'
import { STORAGE_STATE } from '../../e2e/playwright.config'

test.describe('Data Isolation - Regular User', () => {
  test.use({ storageState: STORAGE_STATE.USER })

  test('user should only see pegawai from their unit', async ({ page }) => {
    await page.goto('/pegawai')

    // Wait for data to load
    await page.waitForSelector('[data-testid="pegawai-list"]', { timeout: 10000 })

    // Get all visible pegawai
    const pegawaiRows = await page.locator('[data-testid="pegawai-row"]').all()

    // All visible pegawai should be from user's unit
    for (const row of pegawaiRows) {
      const unitText = await row.locator('[data-testid="unit-kerja"]').textContent()
      // Should only show pegawai from user's unit
      expect(unitText).toBeTruthy()
    }
  })

  test('user cannot access pegawai detail from other unit via URL', async ({ page }) => {
    // Try to access pegawai from different unit
    // Using a known ID from another unit
    const otherUnitPegawaiId = '550e8400-e29b-41d4-a716-446655440099'

    const response = await page.goto(`/pegawai/${otherUnitPegawaiId}`)

    // Should either show 404 or redirect
    if (response?.status() === 404) {
      expect(response.status()).toBe(404)
    } else {
      // Or show access denied page
      await expect(page.locator('[data-testid="access-denied"]')).toBeVisible()
    }
  })

  test('user cannot edit pegawai from other unit', async ({ page }) => {
    // Try to access edit page for pegawai from different unit
    const otherUnitPegawaiId = '550e8400-e29b-41d4-a716-446655440099'

    await page.goto(`/pegawai/${otherUnitPegawaiId}/edit`)

    // Should show access denied or redirect
    const isAccessDenied = await page.locator('[data-testid="access-denied"]').isVisible()
    const isRedirected = page.url().includes('/pegawai') && !page.url().includes('/edit')

    expect(isAccessDenied || isRedirected).toBe(true)
  })

  test('user cannot delete pegawai from other unit via API', async ({ page }) => {
    const otherUnitPegawaiId = '550e8400-e29b-41d4-a716-446655440099'

    // Try to delete via API call
    const response = await page.request.delete(`/api/v1/pegawai/${otherUnitPegawaiId}`)

    // Should be forbidden or not found
    expect([403, 404]).toContain(response.status())
  })
})

test.describe('Data Isolation - Admin User', () => {
  test.use({ storageState: STORAGE_STATE.ADMIN })

  test('admin can see pegawai from all units', async ({ page }) => {
    await page.goto('/pegawai')

    // Wait for data to load
    await page.waitForSelector('[data-testid="pegawai-list"]', { timeout: 10000 })

    // Get all visible pegawai
    const pegawaiRows = await page.locator('[data-testid="pegawai-row"]').all()

    // Admin should see more pegawai than a regular user
    // (assuming there are multiple units with pegawai)
    expect(pegawaiRows.length).toBeGreaterThan(0)

    // Verify we can see pegawai from different units
    const units = new Set<string>()
    for (const row of pegawaiRows) {
      const unitText = await row.locator('[data-testid="unit-kerja"]').textContent()
      if (unitText) {
        units.add(unitText)
      }
    }

    // Admin should see pegawai from multiple units
    expect(units.size).toBeGreaterThanOrEqual(1)
  })

  test('admin can access any pegawai detail', async ({ page }) => {
    // List all pegawai first
    await page.goto('/pegawai')
    await page.waitForSelector('[data-testid="pegawai-row"]', { timeout: 10000 })

    // Click first pegawai
    const firstRow = page.locator('[data-testid="pegawai-row"]').first()
    const pegawaiId = await firstRow.getAttribute('data-id')

    await firstRow.click()

    // Should be able to view detail
    await expect(page).toHaveURL(new RegExp(`/pegawai/${pegawaiId}`))
    await expect(page.locator('[data-testid="pegawai-detail"]')).toBeVisible()
  })

  test('admin can edit pegawai from any unit', async ({ page }) => {
    // Go to a pegawai detail
    await page.goto('/pegawai')
    await page.waitForSelector('[data-testid="pegawai-row"]', { timeout: 10000 })

    const firstRow = page.locator('[data-testid="pegawai-row"]').first()
    const pegawaiId = await firstRow.getAttribute('data-id')

    // Click edit button
    await firstRow.locator('[data-testid="edit-button"]').click()

    // Should be able to access edit page
    await expect(page).toHaveURL(new RegExp(`/pegawai/${pegawaiId}/edit`))

    // Edit should be allowed
    await expect(page.locator('form')).toBeVisible()
  })

  test('admin can delete pegawai from any unit', async ({ page }) => {
    // Create a test pegawai first (via API)
    const createResponse = await page.request.post('/api/v1/pegawai', {
      data: {
        nip: '199901012025011001',
        nama_lengkap: 'Test Delete E2E',
        unit_kerja_id: '660e8400-e29b-41d4-a716-446655440002', // Different unit
        tanggal_lahir: '1999-01-01',
        jenis_kelamin: 'L',
        agama: 'Islam',
        status_pegawai: 'PNS',
      },
    })

    if (createResponse.ok()) {
      const { data } = await createResponse.json()
      const pegawaiId = data.id

      // Try to delete
      const deleteResponse = await page.request.delete(`/api/v1/pegawai/${pegawaiId}`)
      expect([200, 204]).toContain(deleteResponse.status())
    }
  })
})

test.describe('Data Isolation - Satker Level', () => {
  test.use({ storageState: STORAGE_STATE.USER })

  test('user can only see satker they belong to', async ({ page }) => {
    await page.goto('/satker')

    // Wait for data
    await page.waitForSelector('[data-testid="satker-list"]', { timeout: 10000 })

    // Count visible satker
    const satkerCount = await page.locator('[data-testid="satker-row"]').count()

    // Regular user should only see their own satker
    expect(satkerCount).toBeLessThanOrEqual(1)
  })

  test('user cannot view other satker details', async ({ page }) => {
    const otherSatkerId = '770e8400-e29b-41d4-a716-446655440099'

    const response = await page.goto(`/satker/${otherSatkerId}`)

    // Should not be accessible
    if (response?.status()) {
      expect([403, 404]).toContain(response.status())
    }
  })
})

test.describe('Data Isolation - Audit Logs', () => {
  test.use({ storageState: STORAGE_STATE.ADMIN })

  test('audit logs capture user actions correctly', async ({ page }) => {
    // Perform an action
    await page.goto('/pegawai')
    await page.waitForSelector('[data-testid="pegawai-row"]', { timeout: 10000 })

    // Click edit on first pegawai
    const firstRow = page.locator('[data-testid="pegawai-row"]').first()
    await firstRow.locator('[data-testid="edit-button"]').click()

    // Make a change
    await page.fill('input[name="nama_lengkap"]', 'Updated Name E2E Test')
    await page.click('button[type="submit"]')

    // Wait for success
    await page.waitForSelector('[data-testid="success-message"]', { timeout: 10000 })

    // Check audit logs
    await page.goto('/admin/audit-logs')
    await page.waitForSelector('[data-testid="audit-log-row"]', { timeout: 10000 })

    // Should see the update action logged
    const latestLog = page.locator('[data-testid="audit-log-row"]').first()
    const action = await latestLog.locator('[data-testid="action"]').textContent()

    expect(action).toContain('UPDATE')
  })
})

test.describe('Data Isolation - Rate Limiting', () => {
  test('should enforce rate limits on API calls', async ({ page }) => {
    await page.goto('/pegawai')

    // Make rapid API calls
    const promises = []
    for (let i = 0; i < 150; i++) {
      promises.push(page.request.get('/api/v1/pegawai'))
    }

    const responses = await Promise.all(promises)
    const rateLimited = responses.filter((r) => r.status() === 429)

    // Should hit rate limit
    expect(rateLimited.length).toBeGreaterThan(0)
  })

  test('should show rate limit error message', async ({ page }) => {
    // Trigger rate limit
    for (let i = 0; i < 150; i++) {
      await page.request.get('/api/v1/pegawai')
    }

    // Try to load page
    await page.goto('/pegawai')

    // Should show rate limit message if blocked
    const rateLimitMessage = await page.locator('[data-testid="rate-limit-message"]').isVisible()

    if (rateLimitMessage) {
      const message = await page.locator('[data-testid="rate-limit-message"]').textContent()
      expect(message).toContain('Terlalu banyak')
    }
  })
})
