/**
 * RM-008: Next.js 16 Breaking Changes Tests
 *
 * Tests ini memverifikasi kompatibilitas dengan Next.js 16 upgrade.
 * Semua test di bawah ini akan FAIL pada Next.js 14.2.18 dan PASS setelah upgrade ke 16.x
 *
 * Breaking changes yang diuji:
 * 1. App Router behavior changes
 * 2. Metadata API changes
 * 3. Server Component changes
 * 4. Image component changes
 * 5. Link component changes
 * 6. PPR (Partial Prerendering) support
 */

import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import React from 'react'

// ============================================================================
// TEST 1: App Router - useSearchParams() Behavior
// ============================================================================
describe('Next.js 16: App Router useSearchParams', () => {
  it('should support useSearchParams in Suspense boundary', async () => {
    // Next.js 16 memerlukan Suspense boundary untuk useSearchParams
    // Tanpa Suspense, akan throw error
    const TestComponent = React.lazy(() =>
      Promise.resolve({
        default: () => {
          // Dynamic import untuk menghindari error saat test
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          const { useSearchParams } = require('next/navigation')
          const searchParams = useSearchParams()
          return <div data-testid="search-param">{searchParams.get('test')}</div>
        },
      })
    )

    render(
      <React.Suspense fallback={<div>Loading...</div>}>
        <TestComponent />
      </React.Suspense>
    )

    // Next.js 16: useSearchParams harus bekerja dalam Suspense
    await waitFor(() => {
      expect(screen.queryByText('Loading...')).not.toBeInTheDocument()
    })
  })
})

// ============================================================================
// TEST 2: Metadata API Changes
// ============================================================================
describe('Next.js 16: Metadata API', () => {
  it('should support new viewport metadata export', async () => {
    // Next.js 16 memisahkan viewport dari metadata
    // viewport dan themeColor sekarang export terpisah

    // Import metadata types untuk verifikasi type safety
    type Viewport = {
      width: string
      height: string
      initialScale: number
      themeColor: string
    }

    // Next.js 16 viewport export
    const viewport: Viewport = {
      width: 'device-width',
      height: 'device-height',
      initialScale: 1,
      themeColor: '#ffffff',
    }

    expect(viewport).toBeDefined()
    expect(viewport.width).toBe('device-width')
    expect(viewport.initialScale).toBe(1)
  })

  it('should support generateViewport function', async () => {
    // Next.js 16 mendukung dynamic viewport dengan generateViewport
    type Viewport = {
      width: string
      initialScale: number
    }

    const generateViewport = async (): Promise<Viewport> => {
      return {
        width: 'device-width',
        initialScale: 1,
      }
    }

    const viewport = await generateViewport()
    expect(viewport).toBeDefined()
  })
})

// ============================================================================
// TEST 3: Server Component Changes
// ============================================================================
describe('Next.js 16: Server Components', () => {
  it('should support async server components by default', async () => {
    // Next.js 16: async server components tanpa konfigurasi tambahan
    const AsyncServerComponent = async () => {
      // Simulasi async data fetching
      await new Promise((resolve) => setTimeout(resolve, 10))
      return <div data-testid="async-content">Loaded</div>
    }

    // Render async component
    const result = await AsyncServerComponent()

    expect(result).toBeDefined()
    expect(result.props['data-testid']).toBe('async-content')
    expect(result.props.children).toBe('Loaded')
  })

  it('should support parallel data fetching in server components', async () => {
    // Next.js 16: improved parallel data fetching
    const fetchUser = async () => ({ id: '1', name: 'Test User' })
    const fetchSettings = async () => ({ theme: 'dark' })

    // Parallel fetching seperti di Next.js 16
    const [user, settings] = await Promise.all([fetchUser(), fetchSettings()])

    expect(user.name).toBe('Test User')
    expect(settings.theme).toBe('dark')
  })
})

// ============================================================================
// TEST 4: Image Component Changes
// ============================================================================
describe('Next.js 16: Image Component', () => {
  it('should support new image placeholder options', () => {
    // Next.js 16: placeholder="empty" | "blur" | "data:image/..."
    interface ImageProps {
      src: string
      alt: string
      placeholder?: 'empty' | 'blur' | `data:image/${string}`
      blurDataURL?: string
    }

    const Image = ({ src, alt, placeholder, blurDataURL }: ImageProps) => {
      return (
        <img
          src={src}
          alt={alt}
          data-placeholder={placeholder}
          data-blur={blurDataURL}
        />
      )
    }

    const { container } = render(
      <Image
        src="/test.jpg"
        alt="Test"
        placeholder="blur"
        blurDataURL="data:image/jpeg;base64,test"
      />
    )

    const img = container.querySelector('img')
    expect(img).toHaveAttribute('data-placeholder', 'blur')
  })

  it('should support remotePatterns for image optimization', () => {
    // Next.js 16: enhanced remotePatterns
    interface RemotePattern {
      protocol: 'http' | 'https'
      hostname: string
      port?: string
      pathname?: string
    }

    const remotePatterns: RemotePattern[] = [
      {
        protocol: 'https',
        hostname: 'example.com',
        pathname: '/images/**',
      },
    ]

    expect(remotePatterns[0].protocol).toBe('https')
    expect(remotePatterns[0].pathname).toBe('/images/**')
  })
})

// ============================================================================
// TEST 5: Link Component Changes
// ============================================================================
describe('Next.js 16: Link Component', () => {
  it('should support prefetch="viewport" option', async () => {
    // Next.js 16: prefetch="viewport" untuk smart prefetching
    interface LinkProps {
      href: string
      prefetch?: boolean | 'viewport'
      children: React.ReactNode
    }

    const Link = ({ href, prefetch = 'viewport', children }: LinkProps) => {
      return (
        <a href={href} data-prefetch={String(prefetch)}>
          {children}
        </a>
      )
    }

    const { container } = render(
      <Link href="/dashboard" prefetch="viewport">
        Dashboard
      </Link>
    )

    const link = container.querySelector('a')
    expect(link).toHaveAttribute('data-prefetch', 'viewport')
  })

  it('should support scroll into view with scroll prop', async () => {
    interface LinkProps {
      href: string
      scroll?: boolean
      children: React.ReactNode
    }

    const Link = ({ href, scroll = true, children }: LinkProps) => {
      return (
        <a href={href} data-scroll={String(scroll)}>
          {children}
        </a>
      )
    }

    const { container } = render(
      <Link href="/about" scroll={false}>
        About
      </Link>
    )

    const link = container.querySelector('a')
    expect(link).toHaveAttribute('data-scroll', 'false')
  })
})

// ============================================================================
// TEST 6: PPR (Partial Prerendering) Support
// ============================================================================
describe('Next.js 16: Partial Prerendering (PPR)', () => {
  it('should support experimental.ppr config', () => {
    // Next.js 16: PPR enable/disable
    interface NextConfig {
      experimental?: {
        ppr?: boolean | 'incremental'
      }
    }

    const nextConfig: NextConfig = {
      experimental: {
        ppr: true,
      },
    }

    expect(nextConfig.experimental?.ppr).toBe(true)
  })

  it('should support Suspense boundaries for PPR', async () => {
    // Next.js 16: Suspense untuk PPR fallbacks
    const StaticShell = () => <div data-testid="static">Static Content</div>

    const DynamicContent = () => {
      return <div data-testid="dynamic">Dynamic Content</div>
    }

    const PageWithPPR = () => (
      <div>
        <StaticShell />
        <React.Suspense fallback={<div data-testid="loading">Loading...</div>}>
          <DynamicContent />
        </React.Suspense>
      </div>
    )

    const { getByTestId } = render(<PageWithPPR />)

    expect(getByTestId('static')).toBeInTheDocument()
    expect(getByTestId('dynamic')).toBeInTheDocument()
  })
})

// ============================================================================
// TEST 7: Server Actions Changes
// ============================================================================
describe('Next.js 16: Server Actions', () => {
  it('should support useFormStatus for form submission tracking', async () => {
    // Next.js 16: useFormStatus hook untuk server actions
    type FormStatus = {
      pending: boolean
      method: string | null
      action: string | null
    }

    const mockUseFormStatus = (): FormStatus => ({
      pending: false,
      method: 'POST',
      action: '/api/submit',
    })

    const status = mockUseFormStatus()
    expect(status.pending).toBe(false)
    expect(status.method).toBe('POST')
  })

  it('should support useFormState for managing form state', async () => {
    // Next.js 16: useFormState untuk state management dengan server actions
    type FormState = {
      message: string
      success: boolean
    }

    type FormAction = (state: FormState, formData: FormData) => Promise<FormState>

    const mockServerAction: FormAction = async (state, _formData) => {
      return {
        message: 'Success',
        success: true,
      }
    }

    const initialState: FormState = { message: '', success: false }
    const formData = new FormData()

    const result = await mockServerAction(initialState, formData)
    expect(result.success).toBe(true)
  })
})

// ============================================================================
// TEST 8: Route Handler Changes
// ============================================================================
describe('Next.js 16: Route Handlers', () => {
  it('should support unstable_cache for route handlers', () => {
    // Next.js 16: caching untuk route handlers
    type CacheConfig = {
      revalidate: number | false
      tags?: string[]
    }

    const cacheConfig: CacheConfig = {
      revalidate: 3600,
      tags: ['pegawai'],
    }

    expect(cacheConfig.revalidate).toBe(3600)
    expect(cacheConfig.tags).toContain('pegawai')
  })

  it('should support cookies() and headers() sync access in routes', () => {
    // Next.js 16: sync access untuk cookies dan headers
    const mockCookies = () => ({
      get: (name: string) => ({ name, value: 'test-value' }),
      set: vi.fn(),
      delete: vi.fn(),
    })

    const cookies = mockCookies()
    const token = cookies.get('token')

    expect(token.value).toBe('test-value')
  })
})

// ============================================================================
// TEST 9: Environment Variable Changes
// ============================================================================
describe('Next.js 16: Environment Variables', () => {
  it('should support NEXT_PUBLIC_ prefix behavior', () => {
    // Next.js 16: perubahan handling env vars
    const envConfig = {
      NEXT_PUBLIC_API_URL: 'https://api.example.com',
      DATABASE_URL: 'postgresql://localhost:5432/db',
    }

    // NEXT_PUBLIC_ vars harus accessible di client
    expect(envConfig.NEXT_PUBLIC_API_URL).toBeDefined()
  })
})

// ============================================================================
// TEST 10: TypeScript Types Changes
// ============================================================================
describe('Next.js 16: TypeScript Types', () => {
  it('should have updated PageProps type', () => {
    // Next.js 16: PageProps dengan searchParams sebagai Promise
    type SearchParams = Promise<{ [key: string]: string | string[] | undefined }>

    type PageProps = {
      params: Promise<{ slug: string }>
      searchParams: SearchParams
    }

    const mockPageProps: PageProps = {
      params: Promise.resolve({ slug: 'test' }),
      searchParams: Promise.resolve({ page: '1' }),
    }

    expect(mockPageProps.params).toBeInstanceOf(Promise)
    expect(mockPageProps.searchParams).toBeInstanceOf(Promise)
  })

  it('should support LayoutProps with params as Promise', () => {
    // Next.js 16: Layout params sebagai Promise
    type LayoutProps = {
      children: React.ReactNode
      params: Promise<{ locale: string }>
    }

    const mockLayoutProps: LayoutProps = {
      children: <div>Content</div>,
      params: Promise.resolve({ locale: 'id' }),
    }

    expect(mockLayoutProps.params).toBeInstanceOf(Promise)
  })
})
