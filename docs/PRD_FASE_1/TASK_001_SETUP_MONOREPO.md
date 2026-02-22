# TASK 001: Setup Monorepo (Turborepo + pnpm)

**Sprint:** 1 - Infrastruktur & Fondasi
**Priority:** P0 (Critical Path)
**Estimasi:** 1 hari kerja
**FR:** FR-001

---

## Deskripsi
Inisialisasi monorepo menggunakan Turborepo dengan pnpm workspace untuk mengelola 3 apps dan 3 shared packages.

---

## File Yang Perlu Dibuat

### Root Level
```
/sikerma/
├── turbo.json                    # Konfigurasi Turborepo
├── package.json                  # Root package.json
├── pnpm-workspace.yaml           # Workspace definition
├── .gitignore
├── .env.example
└── README.md
```

### Detail File

#### 1. `turbo.json`
```json
{
  "$schema": "https://turbo.build/schema.json",
  "ui": "tui",
  "tasks": {
    "build": {
      "dependsOn": ["^build"],
      "inputs": ["$TURBO_DEFAULT$", ".env*"],
      "outputs": [".next/**", "!.next/cache/**"]
    },
    "lint": {
      "dependsOn": ["^lint"]
    },
    "dev": {
      "cache": false,
      "persistent": true
    },
    "type-check": {
      "dependsOn": ["^type-check"]
    }
  }
}
```

#### 2. `package.json` (root)
```json
{
  "name": "sikerma",
  "version": "1.0.0",
  "private": true,
  "workspaces": [
    "apps/*",
    "packages/*"
  ],
  "scripts": {
    "build": "turbo build",
    "dev": "turbo dev",
    "lint": "turbo lint",
    "type-check": "turbo type-check",
    "clean": "turbo clean"
  },
  "devDependencies": {
    "turbo": "^latest",
    "typescript": "^latest"
  },
  "engines": {
    "node": ">=18.0.0",
    "pnpm": ">=8.0.0"
  }
}
```

#### 3. `pnpm-workspace.yaml`
```yaml
packages:
  - "apps/*"
  - "packages/*"
```

#### 4. `.gitignore`
```
node_modules/
.next/
apps/*/dist/
apps/*/build/
*.env
*.env.local
*.log
.vscode/
.idea/
.DS_Store
```

#### 5. `.env.example`
```
# Database
DATABASE_URL=postgresql://postgres:postgres@localhost:5435/db_master
DATABASE_KEPEGAWAIAN_URL=postgresql://postgres:postgres@localhost:5435/db_kepegawaian

# Keycloak
KEYCLOAK_URL=http://localhost:8081
KEYCLOAK_REALM=pengadilan-agama
KEYCLOAK_CLIENT_ID_PORTAL=portal-client
KEYCLOAK_CLIENT_SECRET_PORTAL=your-secret-here
KEYCLOAK_CLIENT_ID_MASTER=master-data-client
KEYCLOAK_CLIENT_SECRET_MASTER=your-secret-here
KEYCLOAK_CLIENT_ID_KEPEGAWAIAN=kepegawaian-client
KEYCLOAK_CLIENT_SECRET_KEPEGAWAIAN=your-secret-here

# Backend API
NEXT_PUBLIC_API_URL=http://localhost:3003/api/v1
API_URL=http://localhost:3003/api/v1

# Gotenberg
GOTENBERG_URL=http://localhost:3100

# PM2
PORTAL_PORT=3000
MASTER_DATA_PORT=3001
KEPEGAWAIAN_PORT=3002
BACKEND_PORT=3003
```

---

## Struktur Monorepo

```
/sikerma/
├── apps/
│   ├── portal/                    # Next.js App 1
│   ├── master-data/               # Next.js App 2
│   └── kepegawaian/               # Next.js App 3
├── packages/
│   ├── ui/                        # @sikerma/ui
│   ├── auth/                      # @sikerma/auth
│   └── shared/                    # @sikerma/shared
├── docker/
│   ├── postgres/
│   │   └── init/
│   │       ├── 00_create_databases.sql
│   │       ├── 01_create_tables_master.sql
│   │       ├── 02_create_tables_kepegawaian.sql
│   │       └── 03_seed_data.sql
│   └── gotenberg/
│       └── docker-compose.yml (append)
├── turbo.json
├── package.json
├── pnpm-workspace.yaml
├── .gitignore
├── .env.example
└── docker-compose.yml
```

---

## Perintah Setup

```bash
# Install pnpm (jika belum)
npm install -g pnpm

# Install turbo (jika belum)
npm install -g turbo

# Clone/Init project
cd /media/moohard/windows/laragon/www/kesekretariatan-apps

# Install dependencies
pnpm install

# Verify workspace
pnpm --filter=@sikerma/ui run build
```

---

## Acceptance Criteria

- [ ] Turborepo terinstall dan berfungsi
- [ ] pnpm workspace terkonfigurasi dengan benar
- [ ] File-file root (turbo.json, package.json, pnpm-workspace.yaml) terbuat
- [ ] `.env.example` berisi semua variabel yang diperlukan
- [ ] Struktur folder apps/ dan packages/ terbuat
- [ ] `pnpm install` berjalan tanpa error
- [ ] `pnpm dev` dapat dijalankan (akan error karena apps belum ada, tapi tidak crash)

---

## Catatan Penting

1. **Node.js Version:** Minimal 18.x (cek dengan `node -v`)
2. **pnpm Version:** Minimal 8.x (cek dengan `pnpm -v`)
3. **Port Conflict:** Pastikan port 3000-3003 tidak dipakai aplikasi lain
4. **Database:** Port 5435 sudah digunakan PostgreSQL di docker-compose, pastikan tidak conflict

---

## Next Task
Setelah task ini selesai, lanjut ke:
- **TASK 002:** Setup Shared Packages (@sikerma/ui, @sikerma/auth, @sikerma/shared)
