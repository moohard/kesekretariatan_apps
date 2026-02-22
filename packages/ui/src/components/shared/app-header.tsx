"use client"

import * as React from "react"
import { cva, type VariantProps } from "class-variance-authority"
import {
  Bell,
  ChevronDown,
  LogOut,
  Settings,
  User,
  type LucideIcon,
} from "lucide-react"
import { cn } from "../../lib/utils"
import { Button } from "../ui/button"

// ============================================
// Types
// ============================================
export interface UserInfo {
  id: string
  name: string
  email?: string
  avatar?: string
  role?: string
}

export interface AppSwitcherItem {
  id: string
  name: string
  href: string
  icon?: LucideIcon
  description?: string
}

export interface Notification {
  id: string
  title: string
  message?: string
  time?: string
  read?: boolean
  type?: "info" | "warning" | "error" | "success"
}

export interface AppHeaderProps
  extends React.HTMLAttributes<HTMLElement>,
    VariantProps<typeof headerVariants> {
  user: UserInfo
  notifications?: Notification[]
  apps?: AppSwitcherItem[]
  currentAppId?: string
  onLogout?: () => void
  onNotificationClick?: (notification: Notification) => void
  onProfileClick?: () => void
  onSettingsClick?: () => void
  showNotifications?: boolean
  showAppSwitcher?: boolean
}

// ============================================
// Variants
// ============================================
const headerVariants = cva(
  "sticky top-0 z-50 w-full border-b bg-white dark:bg-gray-900",
  {
    variants: {
      variant: {
        default: "border-gray-200 dark:border-gray-800",
        elevated: "shadow-sm",
        transparent: "bg-transparent border-transparent",
      },
    },
    defaultVariants: {
      variant: "default",
    },
  }
)

// ============================================
// Components
// ============================================

/**
 * AppHeader - Header aplikasi dengan user dropdown, notifikasi, dan app switcher
 *
 * Fitur:
 * - User dropdown dengan profile, settings, logout
 * - Notifikasi badge dengan list
 * - App switcher untuk multi-app navigation
 * - Dark mode support
 */
export function AppHeader({
  user,
  notifications = [],
  apps = [],
  currentAppId,
  onLogout,
  onNotificationClick,
  onProfileClick,
  onSettingsClick,
  showNotifications = true,
  showAppSwitcher = true,
  variant,
  className,
  ...props
}: AppHeaderProps) {
  const [showUserMenu, setShowUserMenu] = React.useState(false)
  const [showNotificationsMenu, setShowNotificationsMenu] = React.useState(false)
  const [showAppMenu, setShowAppMenu] = React.useState(false)

  const userMenuRef = React.useRef<HTMLDivElement>(null)
  const notificationsRef = React.useRef<HTMLDivElement>(null)
  const appMenuRef = React.useRef<HTMLDivElement>(null)

  const unreadCount = notifications.filter((n) => !n.read).length
  const currentApp = apps.find((app) => app.id === currentAppId)

  // Close menus when clicking outside
  React.useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (userMenuRef.current && !userMenuRef.current.contains(event.target as Node)) {
        setShowUserMenu(false)
      }
      if (notificationsRef.current && !notificationsRef.current.contains(event.target as Node)) {
        setShowNotificationsMenu(false)
      }
      if (appMenuRef.current && !appMenuRef.current.contains(event.target as Node)) {
        setShowAppMenu(false)
      }
    }

    document.addEventListener("mousedown", handleClickOutside)
    return () => document.removeEventListener("mousedown", handleClickOutside)
  }, [])

  return (
    <header className={cn(headerVariants({ variant }), className)} {...props}>
      <div className="flex h-16 items-center justify-between px-4">
        {/* Left: App Switcher */}
        {showAppSwitcher && apps.length > 0 && (
          <div ref={appMenuRef} className="relative">
            <button
              onClick={() => setShowAppMenu(!showAppMenu)}
              className="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
            >
              {currentApp?.icon && <currentApp.icon className="h-5 w-5" />}
              <span className="font-medium">{currentApp?.name || "Pilih Aplikasi"}</span>
              <ChevronDown className="h-4 w-4" />
            </button>

            {showAppMenu && (
              <div className="absolute left-0 top-full mt-1 w-64 bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-800 rounded-lg shadow-lg py-2 z-50">
                {apps.map((app) => (
                  <a
                    key={app.id}
                    href={app.href}
                    className={cn(
                      "flex items-center gap-3 px-4 py-3 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors",
                      app.id === currentAppId && "bg-primary/5"
                    )}
                  >
                    {app.icon && <app.icon className="h-5 w-5 text-gray-600 dark:text-gray-400" />}
                    <div>
                      <div className="font-medium text-sm">{app.name}</div>
                      {app.description && (
                        <div className="text-xs text-gray-500">{app.description}</div>
                      )}
                    </div>
                  </a>
                ))}
              </div>
            )}
          </div>
        )}

        {/* Right: Notifications & User */}
        <div className="flex items-center gap-2">
          {/* Notifications */}
          {showNotifications && (
            <div ref={notificationsRef} className="relative">
              <button
                onClick={() => setShowNotificationsMenu(!showNotificationsMenu)}
                className="relative p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
              >
                <Bell className="h-5 w-5" />
                {unreadCount > 0 && (
                  <span className="absolute -top-1 -right-1 bg-red-500 text-white text-xs font-bold rounded-full h-5 w-5 flex items-center justify-center">
                    {unreadCount > 9 ? "9+" : unreadCount}
                  </span>
                )}
              </button>

              {showNotificationsMenu && (
                <div className="absolute right-0 top-full mt-1 w-80 bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-800 rounded-lg shadow-lg z-50">
                  <div className="px-4 py-3 border-b border-gray-200 dark:border-gray-800">
                    <h3 className="font-semibold">Notifikasi</h3>
                  </div>
                  <div className="max-h-96 overflow-y-auto">
                    {notifications.length === 0 ? (
                      <div className="px-4 py-8 text-center text-gray-500">
                        Tidak ada notifikasi
                      </div>
                    ) : (
                      notifications.map((notification) => (
                        <button
                          key={notification.id}
                          onClick={() => onNotificationClick?.(notification)}
                          className={cn(
                            "w-full text-left px-4 py-3 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors border-b border-gray-100 dark:border-gray-800 last:border-0",
                            !notification.read && "bg-blue-50/50 dark:bg-blue-900/10"
                          )}
                        >
                          <div className="font-medium text-sm">{notification.title}</div>
                          {notification.message && (
                            <div className="text-xs text-gray-500 mt-1">{notification.message}</div>
                          )}
                          {notification.time && (
                            <div className="text-xs text-gray-400 mt-1">{notification.time}</div>
                          )}
                        </button>
                      ))
                    )}
                  </div>
                </div>
              )}
            </div>
          )}

          {/* User Menu */}
          <div ref={userMenuRef} className="relative">
            <button
              onClick={() => setShowUserMenu(!showUserMenu)}
              className="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
            >
              {user.avatar ? (
                <img src={user.avatar} alt={user.name} className="h-8 w-8 rounded-full" />
              ) : (
                <div className="h-8 w-8 rounded-full bg-primary text-primary-foreground flex items-center justify-center text-sm font-medium">
                  {user.name.charAt(0).toUpperCase()}
                </div>
              )}
              <div className="hidden md:block text-left">
                <div className="text-sm font-medium">{user.name}</div>
                {user.role && (
                  <div className="text-xs text-gray-500">{user.role}</div>
                )}
              </div>
              <ChevronDown className="h-4 w-4" />
            </button>

            {showUserMenu && (
              <div className="absolute right-0 top-full mt-1 w-56 bg-white dark:bg-gray-900 border border-gray-200 dark:border-gray-800 rounded-lg shadow-lg py-2 z-50">
                <div className="px-4 py-3 border-b border-gray-200 dark:border-gray-800">
                  <div className="font-medium">{user.name}</div>
                  {user.email && (
                    <div className="text-xs text-gray-500">{user.email}</div>
                  )}
                </div>

                <button
                  onClick={() => {
                    onProfileClick?.()
                    setShowUserMenu(false)
                  }}
                  className="w-full flex items-center gap-3 px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
                >
                  <User className="h-4 w-4" />
                  Profil
                </button>

                <button
                  onClick={() => {
                    onSettingsClick?.()
                    setShowUserMenu(false)
                  }}
                  className="w-full flex items-center gap-3 px-4 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
                >
                  <Settings className="h-4 w-4" />
                  Pengaturan
                </button>

                <div className="border-t border-gray-200 dark:border-gray-800 mt-2 pt-2">
                  <button
                    onClick={() => {
                      onLogout?.()
                      setShowUserMenu(false)
                    }}
                    className="w-full flex items-center gap-3 px-4 py-2 text-sm text-red-600 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
                  >
                    <LogOut className="h-4 w-4" />
                    Keluar
                  </button>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </header>
  )
}

export default AppHeader
