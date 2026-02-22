"use client"

import { useEffect } from "react"
import { Card } from "@sikerma/ui"

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string }
  reset: () => void
}) {
  useEffect(() => {
    console.error(error)
  }, [error])

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-50 to-slate-100">
      <Card className="p-8 max-w-md w-full">
        <h2 className="text-2xl font-bold text-slate-900 mb-4">
          Terjadi Kesalahan
        </h2>
        <p className="text-slate-600 mb-6">
          Maaf, terjadi kesalahan saat memproses permintaan Anda.
        </p>
        <button
          onClick={reset}
          className="w-full px-6 py-3 bg-primary text-primary-foreground rounded-md hover:bg-primary/90 transition-colors"
        >
          Coba Lagi
        </button>
      </Card>
    </div>
  )
}