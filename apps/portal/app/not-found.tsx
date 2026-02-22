import { notFound } from "next/navigation"

export default function NotFound() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-50 to-slate-100">
      <div className="text-center">
        <h1 className="text-9xl font-bold text-primary mb-4">404</h1>
        <h2 className="text-2xl font-semibold text-slate-900 mb-4">
          Halaman Tidak Ditemukan
        </h2>
        <p className="text-slate-600 mb-8">
          Maaf, halaman yang Anda cari tidak tersedia.
        </p>
        <a
          href="/"
          className="inline-flex items-center px-6 py-3 bg-primary text-primary-foreground rounded-md hover:bg-primary/90 transition-colors"
        >
          Kembali ke Beranda
        </a>
      </div>
    </div>
  )
}