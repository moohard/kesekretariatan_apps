"use client"

import Link from "next/link"
import { Suspense } from "react"
import { ArrowRight, BarChart3, FileText, Users, Settings, LogOut } from "lucide-react"
import { Button } from "@sikerma/ui"
import { useAuthStore, usePermissions } from "@sikerma/auth"

export default function DashboardPage() {
  const { isAuthenticated, user, logout } = useAuthStore()
  const { isAdmin } = usePermissions()

  const modules = [
    {
      name: "Master Data",
      description: "Kelola data referensi dan master data",
      icon: FileText,
      href: "http://localhost:3001",
      roles: ["admin", "supervisor", "officer", "staff"],
    },
    {
      name: "Kepegawaian",
      description: "Kelola data pegawai dan riwayat",
      icon: Users,
      href: "http://localhost:3002",
      roles: ["admin", "supervisor", "officer", "staff"],
    },
    {
      name: "Statistik",
      description: "Lihat statistik dan laporan",
      icon: BarChart3,
      href: "/statistik",
      roles: ["admin", "supervisor", "officer"],
    },
    {
      name: "Admin",
      description: "Pengaturan sistem dan hak akses",
      icon: Settings,
      href: "/admin",
      roles: ["admin"],
    },
  ]

  const filteredModules = isAdmin
    ? modules
    : modules.filter((module) =>
        module.roles.some((role) => user?.roles.includes(role) || role === "user")
      )

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
      {/* Header */}
      <header className="bg-white border-b border-slate-200 sticky top-0 z-10">
        <div className="container mx-auto px-4 py-4 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold text-slate-900">
              SIKERMA Portal
            </h1>
            <p className="text-sm text-slate-600">
              Sistem Informasi Kesekretariatan Mahkamah Agung
            </p>
          </div>

          {isAuthenticated && (
            <div className="flex items-center gap-4">
              <div className="text-right">
                <p className="text-sm font-medium text-slate-900">
                  {user?.name}
                </p>
                <p className="text-xs text-slate-600">{user?.email}</p>
              </div>
              <Button
                variant="outline"
                size="sm"
                onClick={() => logout()}
              >
                <LogOut className="h-4 w-4 mr-2" />
                Logout
              </Button>
            </div>
          )}
        </div>
      </header>

      {/* Main Content */}
      <main className="container mx-auto px-4 py-12">
        {!isAuthenticated ? (
          <div className="max-w-md mx-auto text-center">
            <div className="bg-white rounded-lg shadow-md p-8">
              <h2 className="text-2xl font-bold text-slate-900 mb-4">
                Selamat Datang
              </h2>
              <p className="text-slate-600 mb-6">
                Silakan login untuk mengakses sistem SIKERMA
              </p>
              <Button onClick={() => window.location.href = "/login"} className="w-full">
                Login
              </Button>
            </div>
          </div>
        ) : (
          <div>
            {/* Welcome Section */}
            <div className="bg-white rounded-lg shadow-md p-8 mb-8">
              <h2 className="text-3xl font-bold text-slate-900 mb-2">
                Selamat Datang, {user?.name}!
              </h2>
              <p className="text-slate-600">
                Pilih modul yang ingin Anda akses
              </p>
            </div>

            {/* Modules Grid */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {filteredModules.map((module) => (
                <Link
                  key={module.name}
                  href={module.href}
                  className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow group"
                >
                  <div className="flex items-start gap-4">
                    <div className="p-3 bg-primary/10 rounded-lg group-hover:bg-primary/20 transition-colors">
                      <module.icon className="h-6 w-6 text-primary" />
                    </div>
                    <div className="flex-1">
                      <h3 className="text-lg font-semibold text-slate-900 mb-1">
                        {module.name}
                      </h3>
                      <p className="text-sm text-slate-600">
                        {module.description}
                      </p>
                    </div>
                  </div>
                  <div className="mt-4 flex items-center text-primary text-sm font-medium group-hover:translate-x-1 transition-transform">
                    Buka Modul
                    <ArrowRight className="h-4 w-4 ml-1" />
                  </div>
                </Link>
              ))}
            </div>
          </div>
        )}
      </main>

      {/* Footer */}
      <footer className="bg-white border-t border-slate-200 mt-auto">
        <div className="container mx-auto px-4 py-6">
          <Suspense fallback={<p className="text-center text-sm text-slate-600">&copy; Mahkamah Agung Republik Indonesia. SIKERMA v1.0.0</p>}>
            <FooterContent />
          </Suspense>
        </div>
      </footer>
    </div>
  )
}

function FooterContent() {
  return (
    <p className="text-center text-sm text-slate-600">
      &copy; {new Date().getFullYear()} Mahkamah Agung Republik Indonesia.
      SIKERMA v1.0.0
    </p>
  )
}