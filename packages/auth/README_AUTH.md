# @sikerma/auth

Paket authentication untuk SIKERMA dengan Keycloak OIDC.

## Fitur

- Keycloak OIDC integration
- Authentication state management (Zustand)
- Role-based access control (RBAC)
- Permission checking
- Token refresh
- Protected routes

## Konfigurasi

Environment variables:

```env
NEXT_PUBLIC_KEYCLOAK_URL=http://localhost:8081
NEXT_PUBLIC_KEYCLOAK_REALM=pengadilan-agama
NEXT_PUBLIC_KEYCLOAK_CLIENT_ID=portal-client
```

## Penggunaan

### 1. Setup AuthProvider di root layout

```tsx
// app/layout.tsx
import { AuthProvider } from "@sikerma/auth"

export default function RootLayout({ children }) {
  return (
    <html lang="id">
      <body>
        <AuthProvider>{children}</AuthProvider>
      </body>
    </html>
  )
}
```

### 2. Check authentication di page

```tsx
// app/page.tsx
import { useEffect } from "react"
import { useAuthStore } from "@sikerma/auth"
import { useRouter } from "next/navigation"

export default function Dashboard() {
  const { checkAuth, isAuthenticated, isLoading } = useAuthStore()
  const router = useRouter()

  useEffect(() => {
    checkAuth()
  }, [])

  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      router.push("/login")
    }
  }, [isLoading, isAuthenticated, router])

  if (isLoading) {
    return <div>Loading...</div>
  }

  return <div>Dashboard</div>
}
```

### 3. Login dan Logout

```tsx
import { useAuthStore } from "@sikerma/auth"

function LoginButton() {
  const { login, logout, isAuthenticated, user } = useAuthStore()

  return (
    <div>
      {isAuthenticated ? (
        <div>
          <p>Selamat datang, {user?.name}</p>
          <button onClick={() => logout()}>Logout</button>
        </div>
      ) : (
        <button onClick={() => login()}>Login</button>
      )}
    </div>
  )
}
```

### 4. Role Guard (Proteksi berdasarkan role)

```tsx
import { RoleGuard } from "@sikerma/auth"

function AdminPanel() {
  return (
    <RoleGuard roles={["admin"]}>
      <div>Panel Admin</div>
    </RoleGuard>
  )
}
```

### 5. Permission Checking

```tsx
import { usePermissions } from "@sikerma/auth"

function CreateButton() {
  const { hasPermission } = usePermissions()

  if (!hasPermission("master_data", "create")) {
    return null
  }

  return <button>Create</button>
}
```

### 6. API Client dengan Auth Token

```tsx
import { useAuthStore } from "@sikerma/auth"

function useApi() {
  const { token } = useAuthStore()

  const fetchWithAuth = async (url: string, options: RequestInit = {}) => {
    const response = await fetch(url, {
      ...options,
      headers: {
        ...options.headers,
        Authorization: `Bearer ${token}`,
      },
    })

    return response.json()
  }

  return { fetchWithAuth }
}
```

## Middleware untuk Next.js

```ts
// middleware.ts
import { NextResponse } from "next/server"
import type { NextRequest } from "next/server"

export function middleware(request: NextRequest) {
  const token = request.cookies.get("kc_token")?.value

  // Protect admin routes
  if (request.nextUrl.pathname.startsWith("/admin")) {
    if (!token) {
      return NextResponse.redirect(new URL("/login", request.url))
    }
  }

  return NextResponse.next()
}

export const config = {
  matcher: ["/admin/:path*", "/dashboard/:path*"],
}
```