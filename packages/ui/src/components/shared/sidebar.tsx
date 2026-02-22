"use client"

import * as React from "react"
import { cva, type VariantProps } from "class-variance-authority"
import {
  ChevronLeft,
  ChevronRight,
  Home,
  type LucideIcon,
} from "lucide-react"
import { cn } from "../../lib/utils"
import { Button } from "../ui/button"

// ============================================
// Types
// ============================================
export interface MenuItem {
  id: string
  label: string
  href?: string
  icon?: LucideIcon
  badge?: string | number
  children?: MenuItem[]
  disabled?: boolean
}

export interface SidebarProps
  extends React.HTMLAttributes<HTMLDivElement>,
    VariantProps<typeof sidebarVariants> {
  menuItems: MenuItem[]
  activeItem?: string
  collapsed?: boolean
  onCollapseChange?: (collapsed: boolean) => void
  logo?: React.ReactNode
  footer?: React.ReactNode
  showHomeButton?: boolean
  homeHref?: string
}

// ============================================
// Variants
// ============================================
const sidebarVariants = cva(
  "flex flex-col h-full bg-white dark:bg-gray-900 border-r border-gray-200 dark:border-gray-800 transition-all duration-300",
  {
    variants: {
      variant: {
        default: "",
        dark: "bg-gray-900 text-white dark:bg-gray-950",
        light: "bg-gray-50 dark:bg-gray-100",
      },
      size: {
        default: "",
        sm: "text-sm",
        lg: "text-base",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  }
)

const menuItemVariants = cva(
  "flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-colors w-full",
  {
    variants: {
      active: {
        true: "bg-primary text-primary-foreground",
        false: "text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-800",
      },
      disabled: {
        true: "opacity-50 cursor-not-allowed",
        false: "cursor-pointer",
      },
      collapsed: {
        true: "justify-center px-2",
        false: "",
      },
    },
    defaultVariants: {
      active: false,
      disabled: false,
      collapsed: false,
    },
  }
)

// ============================================
// Components
// ============================================

/**
 * Sidebar - Komponen navigasi samping untuk aplikasi SIKERMA
 *
 * Fitur:
 * - Collapsible dengan animasi
 * - Nested menu support
 * - Badge support untuk notifikasi
 * - Dark mode support
 * - Keyboard accessible
 */
export function Sidebar({
  menuItems,
  activeItem,
  collapsed = false,
  onCollapseChange,
  logo,
  footer,
  showHomeButton = true,
  homeHref = "/",
  variant,
  size,
  className,
  ...props
}: SidebarProps) {
  const [internalCollapsed, setInternalCollapsed] = React.useState(collapsed)
  const [expandedItems, setExpandedItems] = React.useState<string[]>([])

  const isCollapsed = onCollapseChange ? collapsed : internalCollapsed

  const handleCollapseToggle = () => {
    if (onCollapseChange) {
      onCollapseChange(!collapsed)
    } else {
      setInternalCollapsed(!internalCollapsed)
    }
  }

  const handleExpandToggle = (itemId: string) => {
    setExpandedItems((prev) =>
      prev.includes(itemId)
        ? prev.filter((id) => id !== itemId)
        : [...prev, itemId]
    )
  }

  const renderMenuItem = (item: MenuItem, level = 0) => {
    const isActive = activeItem === item.id
    const isExpanded = expandedItems.includes(item.id)
    const hasChildren = item.children && item.children.length > 0
    const Icon = item.icon

    return (
      <div key={item.id}>
        <button
          onClick={() => {
            if (hasChildren && !isCollapsed) {
              handleExpandToggle(item.id)
            } else if (item.href && !item.disabled) {
              window.location.href = item.href
            }
          }}
          disabled={item.disabled}
          className={cn(
            menuItemVariants({
              active: isActive,
              disabled: item.disabled,
              collapsed: isCollapsed,
            }),
            level > 0 && !isCollapsed && "ml-6"
          )}
          title={isCollapsed ? item.label : undefined}
        >
          {Icon && <Icon className="h-5 w-5 shrink-0" />}
          {!isCollapsed && (
            <>
              <span className="flex-1 text-left truncate">{item.label}</span>
              {item.badge && (
                <span className="bg-primary/10 text-primary px-2 py-0.5 rounded-full text-xs font-semibold">
                  {item.badge}
                </span>
              )}
              {hasChildren && (
                <ChevronRight
                  className={cn(
                    "h-4 w-4 transition-transform",
                    isExpanded && "rotate-90"
                  )}
                />
              )}
            </>
          )}
        </button>

        {/* Render children if expanded */}
        {hasChildren && isExpanded && !isCollapsed && (
          <div className="mt-1 space-y-1">
            {item.children!.map((child) => renderMenuItem(child, level + 1))}
          </div>
        )}
      </div>
    )
  }

  return (
    <aside
      className={cn(
        sidebarVariants({ variant, size }),
        isCollapsed ? "w-16" : "w-64",
        className
      )}
      {...props}
    >
      {/* Header / Logo */}
      <div className="flex items-center justify-between p-4 border-b border-gray-200 dark:border-gray-800">
        {!isCollapsed && logo && <div className="flex-1">{logo}</div>}
        {isCollapsed && logo && <div className="mx-auto">{logo}</div>}
      </div>

      {/* Home Button */}
      {showHomeButton && (
        <div className="p-2">
          <Button
            variant="ghost"
            size={isCollapsed ? "icon" : "default"}
            className={cn("w-full", !isCollapsed && "justify-start")}
            onClick={() => (window.location.href = homeHref)}
          >
            <Home className="h-5 w-5" />
            {!isCollapsed && <span className="ml-2">Beranda</span>}
          </Button>
        </div>
      )}

      {/* Navigation */}
      <nav className="flex-1 overflow-y-auto p-2 space-y-1">
        {menuItems.map((item) => renderMenuItem(item))}
      </nav>

      {/* Footer */}
      {footer && <div className="border-t border-gray-200 dark:border-gray-800 p-4">{footer}</div>}

      {/* Collapse Toggle */}
      <div className="border-t border-gray-200 dark:border-gray-800 p-2">
        <Button
          variant="ghost"
          size="icon"
          className="w-full"
          onClick={handleCollapseToggle}
        >
          {isCollapsed ? (
            <ChevronRight className="h-5 w-5" />
          ) : (
            <ChevronLeft className="h-5 w-5" />
          )}
        </Button>
      </div>
    </aside>
  )
}

export default Sidebar
