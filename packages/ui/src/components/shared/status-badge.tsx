"use client"

import * as React from "react"
import { cva, type VariantProps } from "class-variance-authority"
import { cn } from "../../lib/utils"

// ============================================
// Types
// ============================================
export interface StatusBadgeProps
  extends React.HTMLAttributes<HTMLSpanElement>,
    VariantProps<typeof badgeVariants> {
  status: string
  variant?: "success" | "warning" | "danger" | "info" | "default" | "secondary"
  size?: "sm" | "md" | "lg"
  dot?: boolean
  icon?: React.ReactNode
  pulse?: boolean
}

// ============================================
// Variants
// ============================================
const badgeVariants = cva(
  "inline-flex items-center gap-1.5 font-medium rounded-full transition-colors",
  {
    variants: {
      variant: {
        default: "bg-gray-100 text-gray-800 dark:bg-gray-800 dark:text-gray-200",
        success: "bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-400",
        warning: "bg-yellow-100 text-yellow-800 dark:bg-yellow-900/30 dark:text-yellow-400",
        danger: "bg-red-100 text-red-800 dark:bg-red-900/30 dark:text-red-400",
        info: "bg-blue-100 text-blue-800 dark:bg-blue-900/30 dark:text-blue-400",
        secondary: "bg-gray-100 text-gray-600 dark:bg-gray-800 dark:text-gray-400",
      },
      size: {
        sm: "px-2 py-0.5 text-xs",
        md: "px-2.5 py-1 text-xs",
        lg: "px-3 py-1.5 text-sm",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "md",
    },
  }
)

const dotVariants = cva("rounded-full", {
  variants: {
    variant: {
      default: "bg-gray-500",
      success: "bg-green-500",
      warning: "bg-yellow-500",
      danger: "bg-red-500",
      info: "bg-blue-500",
      secondary: "bg-gray-400",
    },
    size: {
      sm: "h-1.5 w-1.5",
      md: "h-2 w-2",
      lg: "h-2.5 w-2.5",
    },
  },
  defaultVariants: {
    variant: "default",
    size: "md",
  },
})

// ============================================
// Status Mapping
// ============================================
const statusVariantMap: Record<string, StatusBadgeProps["variant"]> = {
  // Status umum
  aktif: "success",
  active: "success",
  nonaktif: "default",
  inactive: "default",
  pending: "warning",
  menunggu: "warning",
  draft: "secondary",
  archived: "secondary",

  // Status pegawai
  pns: "success",
  pppk: "info",
  honorer: "warning",
  kontrak: "info",
  "cps": "info",
  "clt": "warning",

  // Status permohonan
  diajukan: "info",
  diproses: "warning",
  disetujui: "success",
  ditolak: "danger",
  selesai: "success",
  dibatalkan: "danger",

  // Status dokumen
  valid: "success",
  invalid: "danger",
  expired: "warning",
  verifikasi: "info",
}

// ============================================
// Components
// ============================================

/**
 * StatusBadge - Badge status dengan dot indicator dan variant otomatis
 *
 * Fitur:
 * - Auto-detect variant dari status text
 * - Dot indicator dengan pulse animation
 * - Multiple sizes
 * - Icon support
 * - Dark mode support
 */
export function StatusBadge({
  status,
  variant,
  size = "md",
  dot = true,
  icon,
  pulse = false,
  className,
  ...props
}: StatusBadgeProps) {
  // Auto-detect variant jika tidak diset
  const resolvedVariant = variant || statusVariantMap[status.toLowerCase()] || "default"

  return (
    <span
      className={cn(badgeVariants({ variant: resolvedVariant, size }), className)}
      {...props}
    >
      {dot && (
        <span className="relative flex h-2 w-2">
          {pulse && (
            <span
              className={cn(
                "absolute inline-flex h-full w-full rounded-full opacity-75 animate-ping",
                dotVariants({ variant: resolvedVariant })
              )}
            />
          )}
          <span
            className={cn(
              "relative inline-flex rounded-full",
              dotVariants({ variant: resolvedVariant, size })
            )}
          />
        </span>
      )}
      {icon}
      <span className="capitalize">{status}</span>
    </span>
  )
}

/**
 * Helper function untuk mendapatkan variant dari status
 */
export function getStatusVariant(status: string): StatusBadgeProps["variant"] {
  return statusVariantMap[status.toLowerCase()] || "default"
}

export default StatusBadge
