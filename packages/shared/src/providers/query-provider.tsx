"use client"

import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { ReactQueryDevtools } from "@tanstack/react-query-devtools"
import { useState, ReactNode } from "react"

/**
 * QueryProvider - Provider untuk TanStack Query (React Query v5)
 *
 * Konfigurasi default:
 * - staleTime: 1 menit (data dianggap fresh selama 1 menit)
 * - gcTime: 5 menit (cache disimpan selama 5 menit)
 * - refetchOnWindowFocus: false (tidak auto refetch saat window fokus)
 * - retry: 1 (hanya retry 1 kali untuk failed queries)
 */
export function QueryProvider({ children }: { children: ReactNode }) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            // Data dianggap fresh selama 1 menit
            staleTime: 60 * 1000,
            // Cache disimpan selama 5 menit
            gcTime: 5 * 60 * 1000,
            // Tidak auto refetch saat window fokus
            refetchOnWindowFocus: false,
            // Retry 1 kali untuk failed queries
            retry: 1,
            // Retry delay dengan exponential backoff
            retryDelay: (attemptIndex) => Math.min(1000 * 2 ** attemptIndex, 30000),
          },
          mutations: {
            // Retry 1 kali untuk failed mutations
            retry: 1,
          },
        },
      })
  )

  return (
    <QueryClientProvider client={queryClient}>
      {children}
      {process.env.NODE_ENV === "development" && <ReactQueryDevtools initialIsOpen={false} />}
    </QueryClientProvider>
  )
}

/**
 * Hook untuk mendapatkan QueryClient instance
 * Berguna untuk manual cache manipulation
 */
export function getQueryClient() {
  return new QueryClient({
    defaultOptions: {
      queries: {
        staleTime: 60 * 1000,
        gcTime: 5 * 60 * 1000,
        refetchOnWindowFocus: false,
        retry: 1,
      },
    },
  })
}

export default QueryProvider
