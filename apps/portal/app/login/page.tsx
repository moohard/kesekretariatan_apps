import { redirect } from "next/navigation"

export default function LoginPage() {
  const keycloakUrl = process.env.NEXT_PUBLIC_KEYCLOAK_URL || "http://localhost:8081"
  const realm = process.env.NEXT_PUBLIC_KEYCLOAK_REALM || "pengadilan-agama"
  const clientId = process.env.NEXT_PUBLIC_KEYCLOAK_CLIENT_ID || "portal-client"
  const redirectUri = `${process.env.NEXT_PUBLIC_APP_URL || "http://localhost:3000"}/api/auth/callback`

  const authUrl = `${keycloakUrl}/realms/${realm}/protocol/openid-connect/auth?client_id=${clientId}&redirect_uri=${encodeURIComponent(redirectUri)}&response_type=code&scope=openid profile email`

  redirect(authUrl)
}
