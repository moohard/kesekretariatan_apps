# SPRINT 1: Infrastruktur & Fondasi

**Timeline:** 8-10 hari kerja
**Status:** Not Started
**Dependency:** Tidak ada (starting point)

---

## Overview
Sprint 1 adalah fondasi teknis untuk seluruh Fase 1. Semua task di sprint ini **critical path** - jika ada yang gagal, sprint berikutnya tidak bisa dimulai.

---

## Goals Sprint 1

| ID | Goal | Metric | Target |
|----|------|--------|--------|
| G-01 | Monorepo siap pakai | `pnpm install` sukses, `pnpm dev` bisa jalan | 100% |
| G-02 | Backend Go Fiber berfungsi | Server listen di port 3003, API dapat diakses | 100% |
| G-03 | Keycloak terkonfigurasi | Realm + 3 clients + roles siap | 100% |
| G-04 | Database ter-migrasi | 2 DB + 21 tabel terbuat + seed data | 100% |
| G-05 | Auth flow berfungsi | Login via Keycloak → redirect → session valid | 100% |
| G-06 | Shared UI components | 10 komponen reusable terbuat | 100% |

---

## Task Breakdown

### Week 1: Setup & Infrastructure

| Task | Nama | Estimasi | Status | FR |
|------|------|----------|--------|-----|
| 001 | Setup Monorepo (Turborepo + pnpm) | 1 hari | ◻ | FR-001 |
| 002 | Setup Shared Packages (@sikerma/ui, @sikerma/auth, @sikerma/shared) | 1.5 hari | ◻ | FR-002, FR-003, FR-004 |
| 003 | Setup 3 Next.js Apps (Portal, Master Data, Kepegawaian) | 1 hari | ◻ | FR-001 |
| 004 | Bootstrap Go Fiber Backend | 1.5 hari | ◻ | FR-007 |
| 005 | Setup Keycloak Realm & Clients | 1 hari | ◻ | FR-005 |

**Subtotal Week 1:** 6 hari

### Week 2: Database, Auth, & Middleware

| Task | Nama | Estimasi | Status | FR |
|------|------|----------|--------|-----|
| 006 | Setup Database Migration & Seed | 1.5 hari | ◻ | FR-006, FR-010 |
| 007 | Setup Gotenberg Service | 0.5 hari | ◻ | FR-009 |
| 008 | Implementasi Auth Flow (Login → Callback) | 1 hari | ◻ | FR-003, FR-101 |
| 009 | Implementasi RBAC Middleware | 1 hari | ◻ | FR-008 |
| 010 | Implementasi Audit Trail Middleware | 1 hari | ◻ | FR-008 |
| 011 | Setup Shared UI Components | 1 hari | ◻ | FR-002 |

**Subtotal Week 2:** 6 hari

---

## Detailed Task List

### Task 001: Setup Monorepo
**File:** `TASK_001_SETUP_MONOREPO.md`

**Deliverables:**
- `turbo.json` terkonfigurasi
- `package.json` root dengan workspace definition
- `pnpm-workspace.yaml`
- `.env.example` dengan semua variabel
- Struktur folder apps/ dan packages/

**Acceptance Criteria:**
- ✅ `pnpm install` berjalan tanpa error
- ✅ `pnpm dev` dapat dijalankan (minimal tidak crash)

---

### Task 002: Setup Shared Packages
**File:** `TASK_002_SETUP_SHARED_PKGS.md`

**Deliverables:**
- Package `@sikerma/ui` (shadcn/ui + custom components)
- Package `@sikerma/auth` (Better Auth + Keycloak OIDC)
- Package `@sikerma/shared` (API client + types + utils)

**Acceptance Criteria:**
- ✅ Semua packages bisa di-build (`tsc`)
- ✅ Components dapat di-import dari apps
- ✅ Tailwind config shared berfungsi

---

### Task 003: Setup 3 Next.js Apps

**Deliverables:**
- App Portal (port 3000) - Next.js 14 (App Router)
- App Master Data (port 3001) - Next.js 14 (App Router)
- App Kepegawaian (port 3002) - Next.js 14 (App Router)

**File Structure:**
```
apps/
├── portal/
│   ├── app/
│   │   ├── (auth)/
│   │   │   └── login/
│   │   │       └── page.tsx
│   │   ├── (app)/
│   │   │   ├── layout.tsx
│   │   │   └── page.tsx
│   │   └── layout.tsx
│   ├── components/
│   ├── lib/
│   ├── public/
│   ├── next.config.js
│   └── package.json
├── master-data/
│   ├── app/
│   │   ├── (auth)/
│   │   │   └── login/
│   │   │       └── page.tsx
│   │   ├── (app)/
│   │   │   ├── layout.tsx
│   │   │   └── page.tsx
│   │   │   └── master/
│   │   │       ├── satker/
│   │   │       ├── jabatan/
│   │   │       └── ...
│   │   └── layout.tsx
│   ├── components/
│   ├── lib/
│   ├── public/
│   ├── next.config.js
│   └── package.json
└── kepegawaian/
    ├── app/
    │   ├── (auth)/
    │   │   └── login/
    │   │       └── page.tsx
    │   ├── (app)/
    │   │   ├── layout.tsx
    │   │   └── page.tsx
    │   │   └── pegawai/
    │   │       ├── page.tsx
    │   │       └── [nip]/
    │   │           └── page.tsx
    │   └── layout.tsx
    ├── components/
    ├── lib/
    ├── public/
    ├── next.config.js
    └── package.json
```

**Acceptance Criteria:**
- ✅ Ketiga app bisa dijalankan terpisah (`pnpm dev`)
- ✅ Port tidak conflict (3000, 3001, 3002)
- ✅ Next.js App Router terkonfigurasi
- ✅ Import shared packages berfungsi

---

### Task 004: Bootstrap Go Fiber Backend

**Deliverables:**
```
api/
├── main.go
├── go.mod
├── go.sum
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── cors.go
│   │   ├── logger.go
│   │   └── audit.go
│   ├── database/
│   │   └── connection.go
│   ├── handlers/
│   │   ├── auth/
│   │   ├── master/
│   │   ├── kepegawaian/
│   │   └── portal/
│   ├── models/
│   │   ├── master/
│   │   ├── kepegawaian/
│   │   └── portal/
│   └── utils/
│       └── helpers.go
├── pkg/
│   └── keycloak/
│       └── client.go
└── Dockerfile
```

**Acceptance Criteria:**
- ✅ Go Fiber server listen di port 3003
- ✅ Middleware dasar (CORS, logger) berfungsi
- ✅ Database connection pool terbuat
- ✅ Struktur project modular dan scalable

---

### Task 005: Setup Keycloak Realm & Clients

**Deliverables:**
- Realm `pengadilan-agama` terbuat
- 3 Clients: `portal-client`, `master-data-client`, `kepegawaian-client`
- Client roles: `[access]` per client
- Realm export JSON untuk backup

**Keycloak Configuration:**
```json
{
  "realm": "pengadilan-agama",
  "enabled": true,
  "clients": [
    {
      "clientId": "portal-client",
      "name": "SIKERMA Portal",
      "baseUrl": "http://localhost:3000",
      "redirectUris": ["http://localhost:3000/*"],
      "webOrigins": ["http://localhost:3000"],
      "clientAuthenticatorType": "client-secret"
    },
    {
      "clientId": "master-data-client",
      "name": "SIKERMA Master Data",
      "baseUrl": "http://localhost:3001",
      "redirectUris": ["http://localhost:3001/*"],
      "webOrigins": ["http://localhost:3001"],
      "clientAuthenticatorType": "client-secret"
    },
    {
      "clientId": "kepegawaian-client",
      "name": "SIKERMA Kepegawaian",
      "baseUrl": "http://localhost:3002",
      "redirectUris": ["http://localhost:3002/*"],
      "webOrigins": ["http://localhost:3002"],
      "clientAuthenticatorType": "client-secret"
    }
  ]
}
```

**Acceptance Criteria:**
- ✅ Realm `pengadilan-agama` dapat diakses di `http://localhost:8081`
- ✅ 3 clients terdaftar dengan redirect URIs yang benar
- ✅ Admin console Keycloak berfungsi
- ✅ Test user dapat login via Keycloak

---

### Task 006: Setup Database Migration & Seed

**Deliverables:**
```
docker/postgres/init/
├── 00_create_databases.sql
├── 01_create_tables_master.sql
├── 02_create_tables_kepegawaian.sql
├── 03_create_tables_rbac.sql
├── 04_create_tables_audit.sql
├── 05_seed_ref_data.sql
├── 06_seed_29_pegawai.sql
└── 07_seed_rbac_roles.sql
```

**Database Schema:**
```sql
-- db_master
CREATE TABLE satker (id, kode, nama, alamat, telepon, email, tipe, is_active, created_at, updated_at);
CREATE TABLE jabatan (id, kode, nama, jenis, eselon, kelas_jabatan, is_active, created_at, updated_at);
CREATE TABLE golongan (id, kode, nama_pangkat, ruang, tingkat, is_active, created_at, updated_at);
CREATE TABLE unit_kerja (id, satker_id, kode, nama, parent_id, is_active, created_at, updated_at);
-- ... 7 tabel lainnya

-- db_kepegawaian
CREATE TABLE pegawai (
  nip VARCHAR(18) PRIMARY KEY,
  nama VARCHAR(255) NOT NULL,
  gelar_depan VARCHAR(50),
  gelar_belakang VARCHAR(50),
  tempat_lahir VARCHAR(100),
  tanggal_lahir DATE,
  jenis_kelamin VARCHAR(1),
  agama_id INT,
  status_pegawai VARCHAR(20),
  golongan_id INT,
  jabatan_id INT,
  unit_kerja_id INT,
  no_hp VARCHAR(20),
  email VARCHAR(100),
  foto_url VARCHAR(500),
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE riwayat_pangkat (
  id SERIAL PRIMARY KEY,
  nip VARCHAR(18) REFERENCES pegawai(nip),
  golongan_id INT REFERENCES golongan(id),
  no_sk VARCHAR(100),
  tanggal_sk DATE,
  tmt DATE,
  pejabat_penetap VARCHAR(255),
  file_sk_url VARCHAR(500),
  created_at TIMESTAMP
);

-- ... tabel riwayat_jabatan, riwayat_pendidikan, keluarga

-- RBAC Tables
CREATE TABLE app_roles (id, app, role_name, description, is_system, created_at);
CREATE TABLE app_permissions (id, app, permission_name, description, resource, action, created_at);
CREATE TABLE role_permissions (id, role_id, permission_id, created_at);
CREATE TABLE user_app_roles (id, user_id, role_id, created_at);
```

**Acceptance Criteria:**
- ✅ 2 database terbuat: `db_master` dan `db_kepegawaian`
- ✅ 21 tabel ter-create dengan foreign keys yang benar
- ✅ Data referensi ter-seed (satker, golongan, jabatan, unit kerja)
- ✅ 29 pegawai ter-seed dari `data_pegawai.json`
- ✅ RBAC roles & permissions default ter-seed

---

### Task 007: Setup Gotenberg Service

**Deliverables:**
- Service Gotenberg ditambahkan ke `docker-compose.yml`
- Port 3100 ter-expose
- Test endpoint `/forms/chromium/convert/html` berfungsi

**docker-compose.yml Addition:**
```yaml
gotenberg:
  image: gotenberg/gotenberg:8
  container_name: gotenberg
  ports:
    - "3100:3000"
  networks:
    - sikerma-network
```

**Acceptance Criteria:**
- ✅ Gotenberg container running di port 3100
- ✅ Health check `/health` return 200
- ✅ Test convert HTML to PDF berhasil

---

### Task 008: Implementasi Auth Flow

**Flow:**
```
1. User akses /login di salah satu app
2. Redirect ke Keycloak:
   http://localhost:8081/realms/pengadilan-agama/protocol/openid-connect/auth?
     client_id={client-id}&
     redirect_uri={app-url}/auth/callback&
     response_type=code&
     scope=openid
3. User login di Keycloak
4. Keycloak redirect ke {app-url}/auth/callback?code={auth-code}
5. App exchange code ke token via /token endpoint
6. Store session (Better Auth)
7. Redirect ke dashboard
```

**Acceptance Criteria:**
- ✅ Login page menampilkan tombol "Login dengan SSO"
- ✅ Redirect ke Keycloak berfungsi
- ✅ Callback handler berhasil exchange code ke token
- ✅ Session tersimpan dan valid
- ✅ User dapat akses halaman protected

---

### Task 009: Implementasi RBAC Middleware

**Middleware Logic:**
```go
func RBACMiddleware(requiredPermission string) fiber.Handler {
  return func(c *fiber.Ctx) error {
    // 1. Get user from context (set by auth middleware)
    user := c.Locals("user").(User)

    // 2. Parse app & permission dari requiredPermission
    // Format: "pegawai.view_all", "pegawai.create", dll
    parts := strings.Split(requiredPermission, ".")
    app := parts[0] // "pegawai"

    // 3. Get user permissions dari database
    permissions, err := GetUserPermissions(user.ID, app)
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(...)
    }

    // 4. Check if user has required permission
    if !Contains(permissions, requiredPermission) {
      return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
        "error": "forbidden",
        "message": "Anda tidak memiliki akses untuk melakukan aksi ini",
      })
    }

    return c.Next()
  }
}
```

**Usage:**
```go
// In handler
pegawaiRouter.Get("/", rbacMiddleware("pegawai.view_all"), ListPegawai)
pegawaiRouter.Post("/", rbacMiddleware("pegawai.create"), CreatePegawai)
pegawaiRouter.Put("/:nip", rbacMiddleware("pegawai.update"), UpdatePegawai)
pegawaiRouter.Delete("/:nip", rbacMiddleware("pegawai.delete"), DeletePegawai)
```

**Acceptance Criteria:**
- ✅ Middleware dapat di-attach ke routes
- ✅ Permission check berfungsi (allow/deny sesuai role)
- ✅ Error message jelas untuk forbidden access
- ✅ Superadmin bypass semua permission check

---

### Task 010: Implementasi Audit Trail Middleware

**Audit Log Schema:**
```sql
CREATE TABLE audit_logs (
  id SERIAL PRIMARY KEY,
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id VARCHAR(255) NOT NULL,
  user_name VARCHAR(255),
  app VARCHAR(50) NOT NULL,
  action VARCHAR(50) NOT NULL, -- CREATE, UPDATE, DELETE
  resource VARCHAR(100) NOT NULL, -- pegawai, jabatan, dll
  resource_id VARCHAR(255),
  old_value JSONB,
  new_value JSONB,
  ip_address VARCHAR(50),
  user_agent TEXT
);
```

**Middleware Logic:**
```go
func AuditMiddleware() fiber.Handler {
  return func(c *fiber.Ctx) error {
    // Capture response
    c.Next()

    // Only log mutating operations
    if c.Method() != "POST" && c.Method() != "PUT" && c.Method() != "DELETE" {
      return nil
    }

    // Get user
    user := c.Locals("user").(User)

    // Parse action
    action := ""
    switch c.Method() {
    case "POST": action = "CREATE"
    case "PUT": action = "UPDATE"
    case "DELETE": action = "DELETE"
    }

    // Extract resource dari path
    // /api/v1/pegawai/:nip → resource = "pegawai"
    pathParts := strings.Split(c.Path(), "/")
    resource := pathParts[3] // "pegawai"

    // Get request body (old value for update/delete)
    var oldVal map[string]interface{}
    if action == "UPDATE" || action == "DELETE" {
      oldVal = GetResourceFromDB(resource, c.Params("nip"))
    }

    // Get response body (new value)
    newVal := c.Response().Body()

    // Insert audit log
    auditLog := AuditLog{
      UserID:     user.ID,
      UserName:   user.Name,
      App:        getAppFromPath(c.Path()),
      Action:     action,
      Resource:   resource,
      ResourceID: c.Params("nip"),
      OldValue:   oldVal,
      NewValue:   newVal,
      IPAddress:  c.IP(),
      UserAgent:  string(c.Request().Header.UserAgent()),
    }

    db.Create(&auditLog)
  }
}
```

**Acceptance Criteria:**
- ✅ Setiap POST/PUT/DELETE tercatat di `audit_logs`
- ✅ Old value dan new value ter-record dengan benar
- ✅ User ID, timestamp, IP address ter-capture
- ✅ Resource dan action ter-identifikasi dengan benar

---

### Task 011: Setup Shared UI Components

**10 Komponen yang Perlu Dibuat:**

| Component | Lokasi | Props Utama | Dipakai Di |
|-----------|--------|-------------|------------|
| `Sidebar` | `@sikerma/ui/layout/sidebar.tsx` | menuItems[], activeItem, collapsed | Semua apps |
| `AppHeader` | `@sikerma/ui/layout/app-header.tsx` | user, notifications, appSwitcher | Semua apps |
| `PageHeader` | `@sikerma/ui/layout/page-header.tsx` | title, description, actions[] | Semua halaman |
| `Breadcrumb` | `@sikerma/ui/data-display/breadcrumb.tsx` | items[] | Semua halaman |
| `DataTable` | `@sikerma/ui/data-display/data-table.tsx` | columns, data, searchable, sortable, pagination | Master Data, Kepegawaian |
| `StatusBadge` | `@sikerma/ui/data-display/status-badge.tsx` | status, variant | Kepegawaian, Portal |
| `FormDialog` | `@sikerma/ui/forms/form-dialog.tsx` | title, fields, onSubmit, mode | CRUD forms |
| `DeleteConfirm` | `@sikerma/ui/feedback/delete-confirm.tsx` | title, message, onConfirm | Semua delete actions |
| `FileUpload` | `@sikerma/ui/feedback/file-upload.tsx` | accept, maxSize, onUpload | Upload foto/SK |
| `StepWizard` | `@sikerma/ui/forms/step-wizard.tsx` | steps[], currentStep, onNext, onBack | Tambah pegawai multi-step |

**Acceptance Criteria:**
- ✅ Semua 10 komponen terbuat dan ter-export dari `@sikerma/ui`
- ✅ Komponen dapat di-import dan digunakan di apps
- ✅ Props types ter-define dengan benar (TypeScript)
- ✅ Styling konsisten (pakai tailwind + shadcn classes)

---

## Definition of Done - Sprint 1

Sprint 1 dianggap **DONE** jika:

- ✅ User dapat membuka `http://localhost:3000/login`
- ✅ Login redirect ke Keycloak dan kembali ke app dengan session valid
- ✅ Backend API di `http://localhost:3003` dapat diakses
- ✅ Database ter-migrasi dengan 29 pegawai ter-seed
- ✅ Ketiga Next.js apps dapat dijalankan bersamaan tanpa conflict
- ✅ Shared packages dapat di-import dan digunakan
- ✅ RBAC middleware berfungsi (test dengan endpoint protected)
- ✅ Audit trail mencatat operasi CRUD
- ✅ Gotenberg service running dan dapat generate PDF

---

## Risks & Mitigations

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Keycloak config kompleks | High | High | Gunakan realm-export.json, test auth flow di awal sprint |
| Port conflict (3000-3003) | Medium | Medium | Check port availability sebelum start, siapkan alternative ports |
| Go Fiber + PostgreSQL connection issue | Medium | High | Test database connection di Task 004, siapkan retry logic |
| Shared package circular dependency | Low | Medium | Struktur packages dengan jelas, hindari import antar shared packages |

---

## Success Metrics

| Metric | Target | Actual |
|--------|--------|--------|
| Task completion rate | 100% | ___ |
| Build success rate | 100% | ___ |
| No critical bugs | 0 | ___ |
| Code review passed | 100% | ___ |
| Documentation complete | 100% | ___ |
