import { format, formatDistanceToNow, parseISO } from "date-fns"
import { id } from "date-fns/locale"

// ============================================
// DATE UTILITIES
// ============================================

export function formatDate(date: Date | string | null | undefined): string {
  if (!date) return "-"

  const d = typeof date === "string" ? parseISO(date) : date
  return format(d, "dd MMMM yyyy", { locale: id })
}

export function formatDateTime(date: Date | string | null | undefined): string {
  if (!date) return "-"

  const d = typeof date === "string" ? parseISO(date) : date
  return format(d, "dd MMMM yyyy HH:mm", { locale: id })
}

export function formatTime(date: Date | string | null | undefined): string {
  if (!date) return "-"

  const d = typeof date === "string" ? parseISO(date) : date
  return format(d, "HH:mm")
}

export function formatRelativeTime(date: Date | string | null | undefined): string {
  if (!date) return "-"

  const d = typeof date === "string" ? parseISO(date) : date
  return formatDistanceToNow(d, { addSuffix: true, locale: id })
}

export function formatAge(tanggalLahir: Date | string | null): string {
  if (!tanggalLahir) return "-"

  const birth = typeof tanggalLahir === "string" ? parseISO(tanggalLahir) : tanggalLahir
  const today = new Date()
  let years = today.getFullYear() - birth.getFullYear()
  let months = today.getMonth() - birth.getMonth()
  let days = today.getDate() - birth.getDate()

  if (days < 0) {
    months--
    const lastMonth = new Date(today.getFullYear(), today.getMonth(), 0)
    days += lastMonth.getDate()
  }

  if (months < 0) {
    years--
    months += 12
  }

  if (years === 0) {
    return `${months} bulan ${days} hari`
  }

  return `${years} tahun ${months} bulan`
}

export function isValidDate(date: any): date is Date {
  return date instanceof Date && !isNaN(date.getTime())
}

// ============================================
// NUMBER UTILITIES
// ============================================

export function formatNumber(num: number | string | null | undefined): string {
  if (num === null || num === undefined) return "0"

  return new Intl.NumberFormat("id-ID").format(Number(num))
}

export function formatCurrency(amount: number | string | null | undefined): string {
  if (amount === null || amount === undefined) return "Rp 0"

  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(Number(amount))
}

export function formatDecimal(num: number | string | null | undefined, decimals: number = 2): string {
  if (num === null || num === undefined) return `0.${"0".repeat(decimals)}`

  return new Intl.NumberFormat("id-ID", {
    minimumFractionDigits: decimals,
    maximumFractionDigits: decimals,
  }).format(Number(num))
}

// ============================================
// STRING UTILITIES
// ============================================

export function truncate(text: string | null | undefined, length: number = 50): string {
  if (!text) return ""
  if (text.length <= length) return text
  return text.slice(0, length) + "..."
}

export function slugify(text: string): string {
  return text
    .toLowerCase()
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/(^-|-$)+/g, "")
}

export function capitalize(text: string): string {
  return text.charAt(0).toUpperCase() + text.slice(1).toLowerCase()
}

export function getInitials(name: string): string {
  return name
    .split(" ")
    .map((n) => n[0])
    .join("")
    .toUpperCase()
    .slice(0, 2)
}

export function generateId(prefix: string = ""): string {
  return `${prefix}${Date.now().toString(36)}-${Math.random().toString(36).substring(2)}`
}

// ============================================
// NIP VALIDATION
// ============================================

export function validateNIP(nip: string): { valid: boolean; message?: string } {
  if (!nip) {
    return { valid: false, message: "NIP wajib diisi" }
  }

  if (nip.length !== 18) {
    return { valid: false, message: "NIP harus 18 digit" }
  }

  if (!/^\d+$/.test(nip)) {
    return { valid: false, message: "NIP harus berupa angka" }
  }

  return { valid: true }
}

export function formatNIP(nip: string): string {
  if (!nip) return ""

  // Format NIP: 19800101 202001 1 001 000001
  // [Tgl Lahir 8 digit] [TMT CPNS 6 digit] [Gender 1 digit] [Nomor Urut 3 digit] [Kode Kerja 6 digit]
  return nip.replace(/(\d{8})(\d{6})(\d{1})(\d{3})(\d{6})/, "$1 $2 $3 $4 $5")
}

export function parseNIP(nip: string): {
  tanggalLahir: string
  tahunLahir: number
  bulanLahir: number
  tanggalLahirNum: number
  tmtCPNS: string
  gender: "L" | "P"
  nomorUrut: string
} | null {
  if (nip.length !== 18 || !/^\d+$/.test(nip)) return null

  const tanggalLahir = nip.substring(0, 8)
  const tahunLahir = parseInt(nip.substring(0, 4))
  const bulanLahir = parseInt(nip.substring(4, 6))
  const tanggalLahirNum = parseInt(nip.substring(6, 8))
  const tmtCPNS = nip.substring(8, 14)
  const genderDigit = nip.substring(14, 15)
  const gender = parseInt(genderDigit) % 2 === 0 ? "P" : "L"
  const nomorUrut = nip.substring(15, 18)

  return {
    tanggalLahir,
    tahunLahir,
    bulanLahir,
    tanggalLahirNum,
    tmtCPNS,
    gender,
    nomorUrut,
  }
}

// ============================================
// NIK VALIDATION
// ============================================

export function validateNIK(nik: string): { valid: boolean; message?: string } {
  if (!nik) {
    return { valid: false, message: "NIK wajib diisi" }
  }

  if (nik.length !== 16) {
    return { valid: false, message: "NIK harus 16 digit" }
  }

  if (!/^\d+$/.test(nik)) {
    return { valid: false, message: "NIK harus berupa angka" }
  }

  return { valid: true }
}

export function formatNIK(nik: string): string {
  if (!nik) return ""

  // Format NIK: 3171 1234 5678 0001
  return nik.replace(/(\d{4})(\d{4})(\d{4})(\d{4})/, "$1 $2 $3 $4")
}

// ============================================
// FILE UTILITIES
// ============================================

export function getFileExtension(filename: string): string {
  return filename.slice(((filename.lastIndexOf(".") - 1) >>> 0) + 2)
}

export function formatFileSize(bytes: number): string {
  if (bytes === 0) return "0 Bytes"

  const k = 1024
  const sizes = ["Bytes", "KB", "MB", "GB"]
  const i = Math.floor(Math.log(bytes) / Math.log(k))

  return Math.round(bytes / Math.pow(k, i) * 100) / 100 + " " + sizes[i]
}

export function isImageFile(filename: string): boolean {
  const imageExtensions = ["jpg", "jpeg", "png", "gif", "bmp", "webp"]
  const ext = getFileExtension(filename).toLowerCase()
  return imageExtensions.includes(ext)
}

export function isPdfFile(filename: string): boolean {
  return getFileExtension(filename).toLowerCase() === "pdf"
}

// ============================================
// ARRAY UTILITIES
// ============================================

export function groupBy<T>(array: T[], key: keyof T): Record<string, T[]> {
  return array.reduce((result, item) => {
    const groupKey = String(item[key])
    if (!result[groupKey]) {
      result[groupKey] = []
    }
    result[groupKey].push(item)
    return result
  }, {} as Record<string, T[]>)
}

export function uniqueBy<T>(array: T[], key: keyof T): T[] {
  return Array.from(new Map(array.map(item => [item[key], item])).values())
}

export function sortBy<T>(array: T[], key: keyof T, order: "asc" | "desc" = "asc"): T[] {
  return [...array].sort((a, b) => {
    const valueA = a[key]
    const valueB = b[key]

    if (valueA < valueB) return order === "asc" ? -1 : 1
    if (valueA > valueB) return order === "asc" ? 1 : -1
    return 0
  })
}

// ============================================
// COLOR UTILITIES
// ============================================

export function stringToColor(str: string): string {
  let hash = 0
  for (let i = 0; i < str.length; i++) {
    hash = str.charCodeAt(i) + ((hash << 5) - hash)
  }

  const hue = Math.abs(hash % 360)
  return `hsl(${hue}, 70%, 50%)`
}

// ============================================
// LOCAL STORAGE UTILITIES
// ============================================

export function getLocalStorage(key: string): any {
  if (typeof window === "undefined") return null

  try {
    const item = window.localStorage.getItem(key)
    return item ? JSON.parse(item) : null
  } catch (error) {
    console.error("Error reading from localStorage:", error)
    return null
  }
}

export function setLocalStorage(key: string, value: any): void {
  if (typeof window === "undefined") return

  try {
    window.localStorage.setItem(key, JSON.stringify(value))
  } catch (error) {
    console.error("Error writing to localStorage:", error)
  }
}

export function removeLocalStorage(key: string): void {
  if (typeof window === "undefined") return

  try {
    window.localStorage.removeItem(key)
  } catch (error) {
    console.error("Error removing from localStorage:", error)
  }
}

export function clearLocalStorage(): void {
  if (typeof window === "undefined") return

  try {
    window.localStorage.clear()
  } catch (error) {
    console.error("Error clearing localStorage:", error)
  }
}