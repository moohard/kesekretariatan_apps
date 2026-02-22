"use client"

import { create } from "zustand"
import type { AuthState, AuthActions, KeycloakConfig } from "../types"
import * as keycloakLib from "../lib/keycloak"

interface AuthStore extends AuthState, AuthActions {
  config: KeycloakConfig
  setConfig: (config: KeycloakConfig) => void
}

const defaultConfig: KeycloakConfig = {
  url: process.env.NEXT_PUBLIC_KEYCLOAK_URL || "http://localhost:8081",
  realm: process.env.NEXT_PUBLIC_KEYCLOAK_REALM || "pengadilan-agama",
  clientId: process.env.NEXT_PUBLIC_KEYCLOAK_CLIENT_ID || "portal-client",
}

const useAuthStore = create<AuthStore>((set, get) => ({
  // Initial state
  isAuthenticated: false,
  user: null,
  token: null,
  isLoading: true,
  config: defaultConfig,

  // Actions
  setConfig: (config: KeycloakConfig) => set({ config }),

  login: async () => {
    set({ isLoading: true })
    try {
      const user = await keycloakLib.login(get().config)
      const token = keycloakLib.getToken(get().config)
      set({
        isAuthenticated: true,
        user,
        token: token || null,
        isLoading: false,
      })
    } catch (error) {
      console.error("Login error:", error)
      set({
        isAuthenticated: false,
        user: null,
        token: null,
        isLoading: false,
      })
      throw error
    }
  },

  logout: async (redirectUri?: string) => {
    set({ isLoading: true })
    try {
      await keycloakLib.logout(get().config, redirectUri)
      set({
        isAuthenticated: false,
        user: null,
        token: null,
        isLoading: false,
      })
    } catch (error) {
      console.error("Logout error:", error)
      set({ isLoading: false })
      throw error
    }
  },

  refreshToken: async () => {
    try {
      const token = await keycloakLib.refreshToken(get().config)
      set({ token: token || null })
    } catch (error) {
      console.error("Token refresh error:", error)
      throw error
    }
  },

  checkAuth: async () => {
    set({ isLoading: true })
    try {
      const keycloak = keycloakLib.getKeycloakInstance(get().config)
      const authenticated = await keycloak.init({
        onLoad: "check-sso",
        checkLoginIframe: false,
        pkceMethod: "S256",
        silentCheckSsoRedirectUri: typeof window !== 'undefined' ? window.location.origin + "/silent-check-sso.html" : undefined,
      })

      if (authenticated) {
        const user = keycloakLib.getUserInfo(get().config)
        const token = keycloakLib.getToken(get().config)
        set({
          isAuthenticated: true,
          user,
          token: token || null,
          isLoading: false,
        })
      } else {
        set({
          isAuthenticated: false,
          user: null,
          token: null,
          isLoading: false,
        })
      }
    } catch (error) {
      console.error("Auth check error:", error)
      set({
        isAuthenticated: false,
        user: null,
        token: null,
        isLoading: false,
      })
    }
  },
}))

export default useAuthStore