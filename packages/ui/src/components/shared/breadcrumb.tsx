"use client"

import * as React from "react"
import { ChevronRight, Home } from "lucide-react"
import { cn } from "../../lib/utils"

// ============================================
// Types
// ============================================
export interface BreadcrumbItem {
  label: string
  href?: string
  icon?: React.ComponentType<{ className?: string }>
  current?: boolean
}

export interface BreadcrumbProps extends React.HTMLAttributes<HTMLElement> {
  items: BreadcrumbItem[]
  homeHref?: string
  showHome?: boolean
  maxItems?: number
  separator?: React.ReactNode
}

// ============================================
// Components
// ============================================

/**
 * Breadcrumb - Navigasi breadcrumb untuk menunjukkan hierarki halaman
 *
 * Fitur:
 * - Home icon optional
 * - Truncation untuk banyak items
 * - Current page indicator
 * - Custom separator
 * - Responsive design
 */
export function Breadcrumb({
  items,
  homeHref = "/",
  showHome = true,
  maxItems = 5,
  separator,
  className,
  ...props
}: BreadcrumbProps) {
  const defaultSeparator = <ChevronRight className="h-4 w-4" />

  // Determine if we need to truncate
  const shouldTruncate = items.length > maxItems
  const visibleItems = shouldTruncate
    ? [items[0], ...items.slice(-(maxItems - 1))]
    : items

  return (
    <nav aria-label="Breadcrumb" className={cn("flex items-center text-sm", className)} {...props}>
      <ol className="flex items-center flex-wrap gap-1">
        {/* Home */}
        {showHome && (
          <>
            <li>
              <a
                href={homeHref}
                className="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 transition-colors"
                aria-label="Beranda"
              >
                <Home className="h-4 w-4" />
              </a>
            </li>
            {(items.length > 0) && (
              <li className="text-gray-400 dark:text-gray-600" aria-hidden="true">
                {separator || defaultSeparator}
              </li>
            )}
          </>
        )}

        {/* Items */}
        {visibleItems.map((item, index) => {
          const isLast = index === visibleItems.length - 1
          const isCurrent = item.current ?? isLast
          const Icon = item.icon

          // Show ellipsis if truncated
          if (shouldTruncate && index === 1) {
            return (
              <React.Fragment key={`ellipsis-${index}`}>
                <li className="text-gray-400 dark:text-gray-600 px-1">...</li>
                <li className="text-gray-400 dark:text-gray-600" aria-hidden="true">
                  {separator || defaultSeparator}
                </li>
              </React.Fragment>
            )
          }

          return (
            <React.Fragment key={`breadcrumb-${index}`}>
              <li>
                {isCurrent || !item.href ? (
                  <span
                    className={cn(
                      "flex items-center gap-1.5",
                      isCurrent
                        ? "font-medium text-gray-900 dark:text-gray-100"
                        : "text-gray-500 dark:text-gray-400"
                    )}
                    aria-current={isCurrent ? "page" : undefined}
                  >
                    {Icon && <Icon className="h-4 w-4" />}
                    {item.label}
                  </span>
                ) : (
                  <a
                    href={item.href}
                    className="flex items-center gap-1.5 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 transition-colors"
                  >
                    {Icon && <Icon className="h-4 w-4" />}
                    {item.label}
                  </a>
                )}
              </li>

              {/* Separator */}
              {!isLast && (
                <li className="text-gray-400 dark:text-gray-600" aria-hidden="true">
                  {separator || defaultSeparator}
                </li>
              )}
            </React.Fragment>
          )
        })}
      </ol>
    </nav>
  )
}

export default Breadcrumb
