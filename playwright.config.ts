import { defineConfig, devices } from '@playwright/test'

/**
 * Playwright Configuration untuk E2E Testing
 * Mendukung parallel testing dengan multiple browsers
 */
export default defineConfig({
  testDir: './e2e',
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 1 : undefined,
  reporter: [
    ['html', { open: 'never' }],
    ['json', { outputFile: 'e2e-results/results.json' }],
    ['list'],
  ],
  use: {
    baseURL: process.env.E2E_BASE_URL || 'http://localhost:3000',
    trace: 'on-first-retry',
    screenshot: 'only-on-failure',
    video: 'retain-on-failure',
  },
  projects: [
    // Setup project untuk authentication state
    {
      name: 'setup',
      testMatch: /.*\.setup\.ts/,
    },
    // Portal App Tests
    {
      name: 'portal-chromium',
      use: { ...devices['Desktop Chrome'] },
      dependencies: ['setup'],
      testDir: './apps/portal/e2e',
    },
    {
      name: 'portal-firefox',
      use: { ...devices['Desktop Firefox'] },
      dependencies: ['setup'],
      testDir: './apps/portal/e2e',
    },
    {
      name: 'portal-webkit',
      use: { ...devices['Desktop Safari'] },
      dependencies: ['setup'],
      testDir: './apps/portal/e2e',
    },
    // Master Data App Tests
    {
      name: 'master-data-chromium',
      use: { ...devices['Desktop Chrome'] },
      dependencies: ['setup'],
      testDir: './apps/master-data/e2e',
    },
    // Kepegawaian App Tests
    {
      name: 'kepegawaian-chromium',
      use: { ...devices['Desktop Chrome'] },
      dependencies: ['setup'],
      testDir: './apps/kepegawaian/e2e',
    },
    // Mobile viewport tests
    {
      name: 'mobile-chrome',
      use: { ...devices['Pixel 5'] },
      testDir: './e2e/mobile',
    },
  ],
  // Run local dev server before starting tests
  webServer: process.env.CI
    ? undefined
    : [
        {
          command: 'pnpm --filter=portal dev',
          url: 'http://localhost:3000',
          reuseExistingServer: !process.env.CI,
          timeout: 120 * 1000,
        },
        {
          command: 'pnpm --filter=master-data dev',
          url: 'http://localhost:3001',
          reuseExistingServer: !process.env.CI,
          timeout: 120 * 1000,
        },
        {
          command: 'pnpm --filter=kepegawaian dev',
          url: 'http://localhost:3002',
          reuseExistingServer: !process.env.CI,
          timeout: 120 * 1000,
        },
      ],
})
