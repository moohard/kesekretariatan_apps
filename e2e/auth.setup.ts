/**
 * Playwright Auth Setup
 * Menyiapkan authentication state untuk E2E tests
 */

import { test as setup, expect } from '@playwright/test'
import { STORAGE_STATE } from './playwright.config'

const AUTH_FILE = '.auth/user.json'

setup('authenticate as admin', async ({ page }) => {
  // Navigate to login page
  await page.goto('/login')

  // Wait for Keycloak login form
  await page.waitForSelector('#username', { timeout: 10000 })

  // Fill login credentials
  await page.fill('#username', process.env.E2E_ADMIN_USERNAME || 'admin')
  await page.fill('#password', process.env.E2E_ADMIN_PASSWORD || 'admin123')

  // Submit login
  await page.click('#kc-login')

  // Wait for redirect back to app
  await page.waitForURL('**/dashboard', { timeout: 15000 })

  // Verify authenticated state
  await expect(page.locator('[data-testid="user-menu"]')).toBeVisible()

  // Save authentication state
  await page.context().storageState({ path: AUTH_FILE })
})

setup('authenticate as regular user', async ({ page }) => {
  const AUTH_USER_FILE = '.auth/regular-user.json'

  await page.goto('/login')
  await page.waitForSelector('#username', { timeout: 10000 })

  await page.fill('#username', process.env.E2E_USER_USERNAME || 'user')
  await page.fill('#password', process.env.E2E_USER_PASSWORD || 'user123')
  await page.click('#kc-login')

  await page.waitForURL('**/dashboard', { timeout: 15000 })
  await expect(page.locator('[data-testid="user-menu"]')).toBeVisible()

  await page.context().storageState({ path: AUTH_USER_FILE })
})
