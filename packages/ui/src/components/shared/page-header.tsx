"use client"

import * as React from "react"
import { cn } from "../../lib/utils"
import { Breadcrumb, type BreadcrumbItem } from "./breadcrumb"

// ============================================
// Types
// ============================================
export interface PageHeaderProps extends React.HTMLAttributes<HTMLDivElement> {
  title: string
  description?: string
  actions?: React.ReactNode
  breadcrumb?: BreadcrumbItem[]
  backHref?: string
  backLabel?: string
  meta?: React.ReactNode
  size?: "sm" | "md" | "lg"
}

// ============================================
// Components
// ============================================

/**
 * PageHeader - Header halaman dengan title, description, breadcrumb, dan actions
 *
 * Fitur:
 * - Title dan description
 * - Breadcrumb navigation
 * - Action buttons area
 * - Back navigation
 * - Meta information
 * - Responsive layout
 */
export function PageHeader({
  title,
  description,
  actions,
  breadcrumb,
  backHref,
  backLabel = "Kembali",
  meta,
  size = "md",
  className,
  ...props
}: PageHeaderProps) {
  const sizeClasses = {
    sm: "py-4",
    md: "py-6",
    lg: "py-8",
  }

  const titleSizeClasses = {
    sm: "text-xl",
    md: "text-2xl",
    lg: "text-3xl",
  }

  return (
    <div className={cn(sizeClasses[size], className)} {...props}>
      {/* Breadcrumb */}
      {breadcrumb && breadcrumb.length > 0 && (
        <div className="mb-2">
          <Breadcrumb items={breadcrumb} />
        </div>
      )}

      {/* Main Content */}
      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
        {/* Left: Title & Description */}
        <div className="flex-1 min-w-0">
          {/* Back Link */}
          {backHref && (
            <a
              href={backHref}
              className="inline-flex items-center text-sm text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 mb-2"
            >
              <svg
                className="w-4 h-4 mr-1"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M15 19l-7-7 7-7"
                />
              </svg>
              {backLabel}
            </a>
          )}

          {/* Title */}
          <h1 className={cn(titleSizeClasses[size], "font-semibold text-gray-900 dark:text-gray-100")}>
            {title}
          </h1>

          {/* Description */}
          {description && (
            <p className="mt-1 text-sm text-gray-500 dark:text-gray-400 max-w-2xl">
              {description}
            </p>
          )}

          {/* Meta */}
          {meta && <div className="mt-2">{meta}</div>}
        </div>

        {/* Right: Actions */}
        {actions && (
          <div className="flex items-center gap-2 shrink-0">
            {actions}
          </div>
        )}
      </div>
    </div>
  )
}

export default PageHeader
