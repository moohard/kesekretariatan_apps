import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

/**
 * Menggabungkan className dengan Tailwind CSS
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

/**
 * Format tanggal ke format Indonesia
 */
export function formatDate(date: Date | string): string {
  const d = typeof date === "string" ? new Date(date) : date;
  return d.toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "long",
    year: "numeric"
  });
}

/**
 * Format tanggal dan waktu
 */
export function formatDateTime(date: Date | string): string {
  const d = typeof date === "string" ? new Date(date) : date;
  return d.toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "long",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit"
  });
}

/**
 * Format angka dengan pemisah ribuan
 */
export function formatNumber(num: number): string {
  return new Intl.NumberFormat("id-ID").format(num);
}

/**
 * Format mata uang (Rupiah)
 */
export function formatCurrency(amount: number): string {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0
  }).format(amount);
}

/**
 * Truncate teks panjang
 */
export function truncate(text: string, length: number = 50): string {
  if (text.length <= length) return text;
  return text.slice(0, length) + "...";
}

/**
 * Generate initial dari nama
 */
export function getInitials(name: string): string {
  return name
    .split(" ")
    .map((n) => n[0])
    .join("")
    .toUpperCase()
    .slice(0, 2);
}