// ============================================
// @sikerma/auth - Authentication Package
// ============================================
// Package untuk mengelola autentikasi dengan Keycloak OIDC
// Menggunakan keycloak-js dan Zustand untuk state management

// ============================================
// Keycloak Library Functions
// ============================================
export {
  getKeycloakInstance,
  login,
  logout,
  getToken,
  refreshToken,
  isAuthenticated,
  getUserInfo,
  hasRole,
  hasAnyRole,
  hasAllRoles,
} from "./lib/keycloak"

// ============================================
// Types
// ============================================
export type {
  AuthState,
  AuthActions,
  UserInfo,
  KeycloakConfig,
} from "./types"

// Re-export UserInfo and KeycloakConfig from keycloak module as well
export type { UserInfo as KeycloakUserInfo, KeycloakConfig as KeycloakConfigType } from "./lib/keycloak"

// ============================================
// Hooks
// ============================================
export { default as useAuthStore } from "./hooks/use-auth"

// ============================================
// Providers & Components
// ============================================
export { AuthProvider, RoleGuard, usePermissions } from "./providers/auth-provider"
