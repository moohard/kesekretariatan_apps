import Keycloak from "keycloak-js"

// Keycloak configuration types
export interface KeycloakConfig {
  url: string
  realm: string
  clientId: string
  clientSecret?: string
}

export interface UserInfo {
  id: string
  username: string
  email: string
  name: string
  roles: string[]
  given_name?: string
  family_name?: string
}

// Initialize Keycloak instance
let keycloakInstance: Keycloak | null = null

export const getKeycloakInstance = (config: KeycloakConfig): Keycloak => {
  if (!keycloakInstance) {
    keycloakInstance = new Keycloak({
      url: config.url,
      realm: config.realm,
      clientId: config.clientId,
    })
  }
  return keycloakInstance
}

// Login to Keycloak
export const login = async (config: KeycloakConfig): Promise<UserInfo> => {
  const keycloak = getKeycloakInstance(config)

  try {
    const authenticated = await keycloak.init({
      onLoad: "login-required",
      checkLoginIframe: false,
      pkceMethod: "S256",
    })

    if (!authenticated) {
      await keycloak.login()
    }

    const tokenParsed = keycloak.tokenParsed
    if (!tokenParsed) {
      throw new Error("Failed to parse token")
    }

    const userInfo: UserInfo = {
      id: tokenParsed.sub || "",
      username: tokenParsed.preferred_username || "",
      email: tokenParsed.email || "",
      name: tokenParsed.name || "",
      roles: tokenParsed.realm_access?.roles || [],
      given_name: tokenParsed.given_name,
      family_name: tokenParsed.family_name,
    }

    return userInfo
  } catch (error) {
    console.error("Keycloak login error:", error)
    throw error
  }
}

// Logout from Keycloak
export const logout = async (config: KeycloakConfig, redirectUri?: string): Promise<void> => {
  const keycloak = getKeycloakInstance(config)
  await keycloak.logout({ redirectUri: redirectUri || window.location.origin })
}

// Get current token
export const getToken = (config: KeycloakConfig): string | undefined => {
  const keycloak = getKeycloakInstance(config)
  return keycloak.token
}

// Refresh token
export const refreshToken = async (config: KeycloakConfig): Promise<string | undefined> => {
  const keycloak = getKeycloakInstance(config)
  try {
    const refreshed = await keycloak.updateToken(60) // 60 seconds
    if (refreshed) {
      return keycloak.token
    }
    return keycloak.token
  } catch (error) {
    console.error("Token refresh error:", error)
    throw error
  }
}

// Check if user is authenticated
export const isAuthenticated = (config: KeycloakConfig): boolean => {
  const keycloak = getKeycloakInstance(config)
  return keycloak.authenticated || false
}

// Get user info from token
export const getUserInfo = (config: KeycloakConfig): UserInfo | null => {
  const keycloak = getKeycloakInstance(config)

  if (!keycloak.authenticated || !keycloak.tokenParsed) {
    return null
  }

  const tokenParsed = keycloak.tokenParsed
  return {
    id: tokenParsed.sub || "",
    username: tokenParsed.preferred_username || "",
    email: tokenParsed.email || "",
    name: tokenParsed.name || "",
    roles: tokenParsed.realm_access?.roles || [],
    given_name: tokenParsed.given_name,
    family_name: tokenParsed.family_name,
  }
}

// Check if user has specific role
export const hasRole = (config: KeycloakConfig, role: string): boolean => {
  const userInfo = getUserInfo(config)
  if (!userInfo) return false
  return userInfo.roles.includes(role) || userInfo.roles.includes("admin")
}

// Check if user has any of the specified roles
export const hasAnyRole = (config: KeycloakConfig, roles: string[]): boolean => {
  const userInfo = getUserInfo(config)
  if (!userInfo) return false

  // Admin bypasses role checks
  if (userInfo.roles.includes("admin")) return true

  return roles.some(role => userInfo.roles.includes(role))
}

// Check if user has all specified roles
export const hasAllRoles = (config: KeycloakConfig, roles: string[]): boolean => {
  const userInfo = getUserInfo(config)
  if (!userInfo) return false

  // Admin bypasses role checks
  if (userInfo.roles.includes("admin")) return true

  return roles.every(role => userInfo.roles.includes(role))
}