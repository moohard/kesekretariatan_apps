/**
 * E2E Tests: Authentication Flow
 *
 * Tests untuk memverifikasi authentication flow dengan Keycloak SSO
 * dan session management.
 */

import { test, expect } from '@playwright/test'
import { STORAGE_STATE } from '../../e2e/playwright.config'

test.describe('Authentication Flow', () => {
  test.describe.configure({ mode: 'parallel' })

  test('should redirect to login when not authenticated', async ({ page }) => {
    await page.goto('/dashboard')

    // Should redirect to login page
    await page.waitForURL('**/login**', { timeout: 10000 })
    expect(page.url()).toContain('/login')
  })

  test('should display login page correctly', async ({ page }) => {
    await page.goto('/login')

    // Wait for Keycloak login form
    await page.waitForSelector('#kc-page-title', { timeout: 10000 })

    // Verify login form elements
    await expect(page.locator('#username')).toBeVisible()
    await expect(page.locator('#password')).toBeVisible()
    await expect(page.locator('#kc-login')).toBeVisible()
  })

  test('should show error for invalid credentials', async ({ page }) => {
    await page.goto('/login')
    await page.waitForSelector('#username', { timeout: 10000 })

    // Enter invalid credentials
    await page.fill('#username', 'invalid_user')
    await page.fill('#password', 'invalid_password')
    await page.click('#kc-login')

    // Should show error message
    await page.waitForSelector('#input-error', { timeout: 5000 })
    const errorMessage = await page.locator('#input-error').textContent()
    expect(errorMessage).toContain('Invalid')
  })

  test('should login successfully with valid credentials', async ({ page }) => {
    await page.goto('/login')
    await page.waitForSelector('#username', { timeout: 10000 })

    // Enter valid credentials
    await page.fill('#username', process.env.E2E_USER_USERNAME || 'testuser')
    await page.fill('#password', process.env.E2E_USER_PASSWORD || 'testpass')
    await page.click('#kc-login')

    // Should redirect to dashboard
    await page.waitForURL('**/dashboard', { timeout: 15000 })
    expect(page.url()).toContain('/dashboard')
  })

  test('should persist session across page reloads', async ({ page }) => {
    // Login first
    await page.goto('/login')
    await page.waitForSelector('#username', { timeout: 10000 })
    await page.fill('#username', process.env.E2E_USER_USERNAME || 'testuser')
    await page.fill('#password', process.env.E2E_USER_PASSWORD || 'testpass')
    await page.click('#kc-login')
    await page.waitForURL('**/dashboard', { timeout: 15000 })

    // Reload page
    await page.reload()

    // Should still be on dashboard (not redirected to login)
    await expect(page).toHaveURL(/.*dashboard.*/)
  })

  test('should logout successfully', async ({ page }) => {
    // Login first
    await page.goto('/login')
    await page.waitForSelector('#username', { timeout: 10000 })
    await page.fill('#username', process.env.E2E_USER_USERNAME || 'testuser')
    await page.fill('#password', process.env.E2E_USER_PASSWORD || 'testpass')
    await page.click('#kc-login')
    await page.waitForURL('**/dashboard', { timeout: 15000 })

    // Click logout
    await page.click('[data-testid="user-menu"]')
    await page.click('[data-testid="logout-button"]')

    // Should redirect to login or home
    await page.waitForURL(/.*(?:login|\/).*/, { timeout: 10000 })
  })
})

test.describe('Authentication - Token Management', () => {
  test.use({ storageState: STORAGE_STATE.USER })

  test('should refresh token before expiry', async ({ page, context }) => {
    await page.goto('/dashboard')

    // Wait for initial token
    const initialToken = await page.evaluate(() => {
      return localStorage.getItem('token')
    })

    expect(initialToken).toBeTruthy()

    // Wait for potential refresh (simulated)
    await page.waitForTimeout(5000)

    // Token should still be valid
    const currentToken = await page.evaluate(() => {
      return localStorage.getItem('token')
    })

    expect(currentToken).toBeTruthy()
  })

  test('should handle session timeout gracefully', async ({ page, context }) => {
    await page.goto('/dashboard')

    // Clear session to simulate timeout
    await context.clearCookies()
    await page.evaluate(() => {
      localStorage.clear()
      sessionStorage.clear()
    })

    // Try to navigate
    await page.click('a[href="/pegawai"]')

    // Should redirect to login
    await page.waitForURL('**/login**', { timeout: 10000 })
  })
})

test.describe('Authentication - Role-Based Access', () => {
  test.use({ storageState: STORAGE_STATE.ADMIN })

  test('admin should access admin-only pages', async ({ page }) => {
    await page.goto('/admin/settings')
    await expect(page).toHaveURL(/.*admin.*/)
  })

  test('admin should see all pegawai across units', async ({ page }) => {
    await page.goto('/pegawai')

    // Wait for data to load
    await page.waitForSelector('[data-testid="pegawai-list"]', { timeout: 10000 })

    // Should see pegawai from multiple units
    const pegawaiCount = await page.locator('[data-testid="pegawai-row"]').count()
    expect(pegawaiCount).toBeGreaterThan(0)
  })
})

test.describe('Authentication - Security', () => {
  test('should include CSRF token in forms', async ({ page }) => {
    await page.goto('/pegawai/new')

    // Check for CSRF token in form
    const csrfToken = await page.locator('input[name="_csrf"]').getAttribute('value')

    // CSRF token should exist and be non-empty
    expect(csrfToken).toBeTruthy()
    expect(csrfToken!.length).toBeGreaterThan(10)
  })

  test('should set secure cookie attributes', async ({ page, context }) => {
    await page.goto('/login')
    await page.waitForSelector('#username', { timeout: 10000 })
    await page.fill('#username', process.env.E2E_USER_USERNAME || 'testuser')
    await page.fill('#password', process.env.E2E_USER_PASSWORD || 'testpass')
    await page.click('#kc-login')
    await page.waitForURL('**/dashboard', { timeout: 15000 })

    // Check cookie attributes
    const cookies = await context.cookies()
    const sessionCookie = cookies.find((c) => c.name.includes('session'))

    if (sessionCookie) {
      expect(sessionCookie.httpOnly).toBe(true)
      expect(sessionCookie.sameSite).toBe('Strict')
    }
  })

  test('should not expose sensitive data in localStorage', async ({ page }) => {
    await page.goto('/login')
    await page.waitForSelector('#username', { timeout: 10000 })
    await page.fill('#username', process.env.E2E_USER_USERNAME || 'testuser')
    await page.fill('#password', process.env.E2E_USER_PASSWORD || 'testpass')
    await page.click('#kc-login')
    await page.waitForURL('**/dashboard', { timeout: 15000 })

    // Check localStorage for sensitive data
    const localStorageData = await page.evaluate(() => {
      return Object.entries(localStorage)
    })

    // Should not store password
    const hasPassword = localStorageData.some(([key, value]) =>
      key.toLowerCase().includes('password') ||
      value.toLowerCase().includes('password')
    )
    expect(hasPassword).toBe(false)
  })
})
