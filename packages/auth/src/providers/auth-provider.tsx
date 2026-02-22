"use client"

import React, { ReactNode, useEffect } from "react"
import useAuthStore from "../hooks/use-auth"
import type { KeycloakConfig } from "../types"

interface AuthProviderProps {
  children: ReactNode
  config?: Partial<KeycloakConfig>
}

export function AuthProvider({ children, config }: AuthProviderProps) {
  const setConfig = useAuthStore((state: { setConfig: (config: KeycloakConfig) => void }) => state.setConfig)

  useEffect(() => {
    if (config) {
      setConfig({
        url: config.url || process.env.NEXT_PUBLIC_KEYCLOAK_URL || "http://localhost:8081",
        realm: config.realm || process.env.NEXT_PUBLIC_KEYCLOAK_REALM || "pengadilan-agama",
        clientId: config.clientId || process.env.NEXT_PUBLIC_KEYCLOAK_CLIENT_ID || "portal-client",
        clientSecret: config.clientSecret,
      })
    }
  }, [config, setConfig])

  return <>{children}</>
}

// Role-based access control component
interface RoleGuardProps {
  children: ReactNode
  roles?: string[]
  requireAll?: boolean // Default: false (require any)
  fallback?: ReactNode
}

export function RoleGuard({ children, roles, requireAll = false, fallback = null }: RoleGuardProps) {
  const user = useAuthStore((state) => state.user)
  const isLoading = useAuthStore((state) => state.isLoading)

  if (isLoading) {
    return <div className="flex items-center justify-center min-h-screen">
      <div className="h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent" />
    </div>
  }

  // Admin bypasses role checks
  if (user?.roles.includes("admin")) {
    return <>{children}</>
  }

  // If no roles specified, show children
  if (!roles || roles.length === 0) {
    return <>{children}</>
  }

  // Check if user has required roles
  const hasAccess = requireAll
    ? roles.every((role) => user?.roles.includes(role))
    : roles.some((role) => user?.roles.includes(role))

  return hasAccess ? <>{children}</> : <>{fallback}</>
}

// Permission check hook
export function usePermissions() {
  const user = useAuthStore((state) => state.user)

  const hasRole = (role: string) => {
    return user?.roles.includes(role) || user?.roles.includes("admin") || false
  }

  const hasAnyRole = (roles: string[]) => {
    if (user?.roles.includes("admin")) return true
    return roles.some((role) => user?.roles.includes(role))
  }

  const hasAllRoles = (roles: string[]) => {
    if (user?.roles.includes("admin")) return true
    return roles.every((role) => user?.roles.includes(role))
  }

  const hasPermission = (resource: string, action: string): boolean => {
    // Admin bypasses permission checks
    if (user?.roles.includes("admin")) return true

    // Simplified permission check based on roles
    // In production, this should check against database RBAC
    const roleHierarchy = {
      admin: 100,
      supervisor: 80,
      officer: 60,
      staff: 40,
      user: 20,
    }

    const maxLevel = user?.roles.reduce((max, role) => {
      return Math.max(max, roleHierarchy[role as keyof typeof roleHierarchy] || 0)
    }, 0) || 0

    // Define required levels for resources and actions
    const requiredLevels: Record<string, Record<string, number>> = {
      master_data: { read: 20, create: 60, update: 80, delete: 100 },
      kepegawaian: { read: 20, create: 40, update: 60, delete: 80 },
      rbac: { read: 100, create: 100, update: 100, delete: 100 },
      audit: { read: 100, delete: 100 },
    }

    const requiredLevel = requiredLevels[resource]?.[action] || 100

    return maxLevel >= requiredLevel
  }

  return {
    user,
    hasRole,
    hasAnyRole,
    hasAllRoles,
    hasPermission,
    isAdmin: user?.roles.includes("admin") || false,
    isAuthenticated: !!user,
  }
}