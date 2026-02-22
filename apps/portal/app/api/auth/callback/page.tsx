"use client"

import { useEffect } from "react"
import { useAuthStore } from "@sikerma/auth"
import { Loader2 } from "lucide-react"

export default function AuthCallbackPage() {
  const { checkAuth, isAuthenticated } = useAuthStore()

  useEffect(() => {
    checkAuth()
  }, [checkAuth])

  useEffect(() => {
    if (isAuthenticated) {
      window.location.href = "/"
    }
  }, [isAuthenticated])

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-50 to-slate-100">
      <div className="text-center">
        <Loader2 className="h-8 w-8 animate-spin text-primary mx-auto mb-4" />
        <p className="text-slate-600">Memproses login...</p>
      </div>
    </div>
  )
}