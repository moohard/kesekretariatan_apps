# TASK 002: Setup Shared Packages

**Sprint:** 1 - Infrastruktur & Fondasi
**Priority:** P0 (Critical Path)
**Estimasi:** 1.5 hari kerja
**FR:** FR-002, FR-003, FR-004

---

## Deskripsi
Membuat 3 shared packages yang akan dipakai oleh semua apps:
1. `@sikerma/ui` - Shared UI components (shadcn/ui + custom)
2. `@sikerma/auth` - Authentication (Better Auth + Keycloak OIDC)
3. `@sikerma/shared` - Shared utilities (API client, types, utils)

---

## File Yang Perlu Dibuat

### 1. Package: @sikerma/ui

```
packages/ui/
├── package.json
├── tsconfig.json
├── components.json
├── tailwind.config.js
├── postcss.config.js
├── src/
│   ├── index.ts
│   ├── lib/
│   │   └── utils.ts
│   ├── components/
│   │   ├── ui/                 # shadcn/ui components
│   │   │   ├── button.tsx
│   │   │   ├── card.tsx
│   │   │   ├── dialog.tsx
│   │   │   ├── table.tsx
│   │   │   ├── input.tsx
│   │   │   ├── label.tsx
│   │   │   ├── select.tsx
│   │   │   ├── badge.tsx
│   │   │   ├── dropdown-menu.tsx
│   │   │   ├── separator.tsx
│   │   │   └── ...
│   │   ├── layout/
│   │   │   ├── sidebar.tsx
│   │   │   ├── app-header.tsx
│   │   │   └── page-header.tsx
│   │   ├── data-display/
│   │   │   ├── data-table.tsx
│   │   │   ├── breadcrumb.tsx
│   │   │   └── status-badge.tsx
│   │   ├── feedback/
│   │   │   ├── delete-confirm.tsx
│   │   │   └── file-upload.tsx
│   │   └── forms/
│   │       └── form-dialog.tsx
│   └── styles/
│       └── globals.css
└── README.md
```

#### `packages/ui/package.json`
```json
{
  "name": "@sikerma/ui",
  "version": "1.0.0",
  "private": true,
  "type": "module",
  "main": "./src/index.ts",
  "types": "./src/index.ts",
  "scripts": {
    "dev": "tsc --watch",
    "build": "tsc",
    "shadcn:add": "pnpm dlx shadcn@latest add"
  },
  "dependencies": {
    "react": "^18.3.1",
    "react-dom": "^18.3.1",
    "class-variance-authority": "^0.7.0",
    "clsx": "^2.1.1",
    "tailwind-merge": "^2.5.4",
    "lucide-react": "^0.454.0",
    "date-fns": "^3.6.0"
  },
  "peerDependencies": {
    "react": "^18.3.1",
    "react-dom": "^18.3.1"
  },
  "devDependencies": {
    "@types/react": "^18.3.12",
    "@types/react-dom": "^18.3.1",
    "typescript": "^5.6.3",
    "tailwindcss": "^3.4.14",
    "postcss": "^8.4.47",
    "autoprefixer": "^10.4.20"
  }
}
```

#### `packages/ui/tailwind.config.js`
```js
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{ts,tsx}",
    "../../apps/**/*.{ts,tsx}"
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: "#f0f9ff",
          100: "#e0f2fe",
          500: "#0ea5e9",
          600: "#0284c7",
          700: "#0369a1",
          900: "#0c4a6e"
        }
      }
    },
  },
  plugins: [],
}
```

#### `packages/ui/src/index.ts`
```typescript
// Layout Components
export { Sidebar } from "./components/layout/sidebar";
export { AppHeader } from "./components/layout/app-header";
export { PageHeader } from "./components/layout/page-header";

// Data Display
export { DataTable } from "./components/data-display/data-table";
export { Breadcrumb } from "./components/data-display/breadcrumb";
export { StatusBadge } from "./components/data-display/status-badge";

// Forms
export { FormDialog } from "./components/forms/form-dialog";

// Feedback
export { DeleteConfirm } from "./components/feedback/delete-confirm";
export { FileUpload } from "./components/feedback/file-upload";

// shadcn/ui re-exports
export { Button } from "./components/ui/button";
export { Card } from "./components/ui/card";
export { Dialog } from "./components/ui/dialog";
export { Table } from "./components/ui/table";
export { Input } from "./components/ui/input";
export { Label } from "./components/ui/label";
export { Select } from "./components/ui/select";
export { Badge } from "./components/ui/badge";
export { DropdownMenu } from "./components/ui/dropdown-menu";
export { Separator } from "./components/ui/separator";

// Utils
export { cn } from "./lib/utils";
```

---

### 2. Package: @sikerma/auth

```
packages/auth/
├── package.json
├── tsconfig.json
├── src/
│   ├── index.ts
│   ├── config/
│   │   └── better-auth.ts
│   ├── lib/
│   │   ├── keycloak-oidc.ts
│   │   └── auth-utils.ts
│   ├── hooks/
│   │   ├── use-session.ts
│   │   ├── use-user.ts
│   │   ├── use-roles.ts
│   │   └── use-permissions.ts
│   ├── middleware/
│   │   └── auth-middleware.ts
│   └── types/
│       └── auth-types.ts
└── README.md
```

#### `packages/auth/package.json`
```json
{
  "name": "@sikerma/auth",
  "version": "1.0.0",
  "private": true,
  "type": "module",
  "main": "./src/index.ts",
  "types": "./src/index.ts",
  "scripts": {
    "dev": "tsc --watch",
    "build": "tsc"
  },
  "dependencies": {
    "better-auth": "^1.0.0",
    "react": "^18.3.1"
  },
  "peerDependencies": {
    "react": "^18.3.1"
  },
  "devDependencies": {
    "@types/react": "^18.3.12",
    "typescript": "^5.6.3"
  }
}
```

#### `packages/auth/src/config/better-auth.ts`
```typescript
import { BetterAuthOptions } from "better-auth";

export const authConfig: BetterAuthOptions = {
  database: {
    provider: "pg",
    url: process.env.DATABASE_URL,
  },
  emailAndPassword: {
    enabled: false,
  },
  socialProviders: {
    keycloak: {
      enabled: true,
      issuer: process.env.KEYCLOAK_URL,
      clientId: process.env.KEYCLOAK_CLIENT_ID,
      clientSecret: process.env.KEYCLOAK_CLIENT_SECRET,
    },
  },
  plugins: [
    // Custom plugins for SIKERMA
  ],
};
```

#### `packages/auth/src/hooks/use-session.ts`
```typescript
import { useAuth } from "better-auth/react";

export function useSession() {
  const { user, status } = useAuth();
  return { user, status, isAuthenticated: status === "authenticated" };
}
```

#### `packages/auth/src/hooks/use-roles.ts`
```typescript
import { useAuth } from "better-auth/react";
import { useMemo } from "react";

export function useRoles(app: string) {
  const { user } = useAuth();

  const roles = useMemo(() => {
    if (!user) return [];
    return user.roles?.filter(r => r.startsWith(`${app}:`)) || [];
  }, [user, app]);

  const hasRole = (role: string) => roles.includes(`${app}:${role}`);

  return { roles, hasRole };
}
```

---

### 3. Package: @sikerma/shared

```
packages/shared/
├── package.json
├── tsconfig.json
├── src/
│   ├── index.ts
│   ├── api/
│   │   ├── client.ts
│   │   └── types.ts
│   ├── types/
│   │   ├── master-data.ts
│   │   ├── kepegawaian.ts
│   │   └── portal.ts
│   ├── utils/
│   │   ├── formatters.ts
│   │   ├── validators.ts
│   │   └── helpers.ts
│   └── constants/
│       ├── routes.ts
│       ├── permissions.ts
│       └── statuses.ts
└── README.md
```

#### `packages/shared/package.json`
```json
{
  "name": "@sikerma/shared",
  "version": "1.0.0",
  "private": true,
  "type": "module",
  "main": "./src/index.ts",
  "types": "./src/index.ts",
  "scripts": {
    "dev": "tsc --watch",
    "build": "tsc"
  },
  "dependencies": {
    "zod": "^3.23.8",
    "axios": "^1.7.7"
  },
  "devDependencies": {
    "typescript": "^5.6.3"
  }
}
```

#### `packages/shared/src/api/client.ts`
```typescript
import axios from "axios";

const apiClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:3003/api/v1",
  headers: {
    "Content-Type": "application/json",
  },
});

// Request interceptor - attach auth token
apiClient.interceptors.request.use(
  (config) => {
    // Get token from Better Auth or cookies
    // TODO: Implement token retrieval
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor - handle errors
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Redirect to login
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

export default apiClient;
```

#### `packages/shared/src/utils/formatters.ts`
```typescript
/**
 * Format NIP (Nomor Induk Pegawai)
 * Format: XX.XXXX.XXXXX.XXXX.XX
 */
export function formatNIP(nip: string): string {
  if (!nip) return "";
  const cleaned = nip.replace(/\D/g, "");
  if (cleaned.length !== 18) return nip;

  return `${cleaned.substring(0, 2)}.${cleaned.substring(2, 6)}.${cleaned.substring(6, 11)}.${cleaned.substring(11, 15)}.${cleaned.substring(15, 18)}`;
}

/**
 * Format tanggal (YYYY-MM-DD) ke format Indonesia
 */
export function formatTanggal(tanggal: string | Date): string {
  if (!tanggal) return "-";
  const date = new Date(tanggal);
  return date.toLocaleDateString("id-ID", {
    day: "numeric",
    month: "long",
    year: "numeric",
  });
}
```

#### `packages/shared/src/utils/validators.ts`
```typescript
import { z } from "zod";

// Validator NIP (18 digit)
export const nipValidator = z.string()
  .min(18, "NIP harus 18 digit")
  .max(18, "NIP harus 18 digit")
  .regex(/^\d+$/, "NIP hanya boleh berisi angka")
  .refine((val) => val.length === 18, "NIP harus 18 digit");

// Validator Email
export const emailValidator = z.string()
  .email("Format email tidak valid");

// Validator Telepon
export const teleponValidator = z.string()
  .regex(/^(\+62|62|0)8[1-9][0-9]{7,10}$/, "Format telepon tidak valid");
```

---

## Perintah Setup

```bash
# Di root project
cd /media/moohard/windows/laragon/www/kesekretariatan-apps

# Install shadcn/ui di package ui
cd packages/ui
pnpm dlx shadcn@latest init

# Add components yang diperlukan
pnpm dlx shadcn@latest add button
pnpm dlx shadcn@latest add card
pnpm dlx shadcn@latest add dialog
pnpm dlx shadcn@latest add table
pnpm dlx shadcn@latest add input
pnpm dlx shadcn@latest add label
pnpm dlx shadcn@latest add select
pnpm dlx shadcn@latest add badge
pnpm dlx shadcn@latest add dropdown-menu
pnpm dlx shadcn@latest add separator

# Build semua packages
cd ../..
pnpm build
```

---

## Acceptance Criteria

- [ ] `@sikerma/ui` package terbuat dengan shadcn/ui components
- [ ] `@sikerma/auth` package terbuat dengan Better Auth + Keycloak config
- [ ] `@sikerma/shared` package terbuat dengan API client + types + utils
- [ ] Semua packages bisa di-import dari apps (test import)
- [ ] `pnpm build` berjalan tanpa error
- [ ] Tailwind config shared berfungsi di semua apps
- [ ] TypeScript types ter-export dengan benar

---

## Catatan Penting

1. **shadcn/ui** adalah headless component library, artinya:
   - Components disalin ke project (bukan dari node_modules)
   - Bisa di-custom sepenuhnya
   - Tidak ada dependency runtime

2. **Better Auth** vs NextAuth.js:
   - Better Auth lebih modern dan performant
   - Mendukung Keycloak OIDC dengan baik
   - TypeScript-first

3. **Shared Package Versioning**:
   - Semua packages pakai version 1.0.0 di Fase 1
   - Breaking changes harus di-handle dengan hati-hati
   - Selalu test semua apps setelah update shared package

---

## Next Task
Setelah task ini selesai, lanjut ke:
- **TASK 003:** Setup 3 Next.js Apps (Portal, Master Data, Kepegawaian)
