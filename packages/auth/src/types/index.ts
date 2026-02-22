export interface AuthState {
  isAuthenticated: boolean
  user: UserInfo | null
  token: string | null
  isLoading: boolean
}

export interface AuthActions {
  login: () => Promise<void>
  logout: (redirectUri?: string) => Promise<void>
  refreshToken: () => Promise<void>
  checkAuth: () => Promise<void>
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

export interface KeycloakConfig {
  url: string
  realm: string
  clientId: string
  clientSecret?: string
}