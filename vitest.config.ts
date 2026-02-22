import { defineConfig } from 'vitest/config'
import react from '@vitejs/plugin-react'
import path from 'path'
import { fileURLToPath } from 'url'

const __dirname = path.dirname(fileURLToPath(import.meta.url))

export default defineConfig({
  plugins: [react()],
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: [path.resolve(__dirname, './vitest.setup.tsx')],
    include: ['apps/**/__tests__/**/*.{test,spec}.{ts,tsx}'],
    exclude: [
      '**/node_modules/**',
      '**/dist/**',
      '**/.next/**',
      '**/coverage/**',
      '**/e2e/**',
      '**/backend/**'
    ],
    coverage: {
      provider: 'v8',
      reporter: ['text', 'json', 'html'],
    },
    testTimeout: 10000,
    hookTimeout: 10000,
  },
  resolve: {
    alias: {
      '@sikerma/auth': path.resolve(__dirname, './packages/auth/src'),
      '@sikerma/shared': path.resolve(__dirname, './packages/shared/src'),
      '@sikerma/ui': path.resolve(__dirname, './packages/ui/src'),
    },
  },
})
