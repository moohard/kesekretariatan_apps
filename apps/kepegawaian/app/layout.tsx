import type { Metadata } from "next"
import { Inter } from "next/font/google"
import "@sikerma/ui/src/styles/globals.css"
import { AuthProvider } from "@sikerma/auth"

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "SIKERMA Kepegawaian - Sistem Informasi Kesekretariatan Mahkamah Agung",
  description: "Kepegawaian SIKERMA - Sistem Informasi Kesekretariatan Mahkamah Agung",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="id">
      <body className={inter.className}>
        <AuthProvider>{children}</AuthProvider>
      </body>
    </html>
  )
}
