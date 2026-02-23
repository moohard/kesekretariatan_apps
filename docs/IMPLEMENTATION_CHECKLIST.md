# IMPLEMENTATION CHECKLIST - SIKERMA FASE 1

**Version:** 1.1
**Date:** 24 Februari 2026
**Reference:** PRD_FASE_1.md v1.3

---

## DAFTAR ISTILAH

| Istilah | Arti                            |
| ------- | ------------------------------- |
| ✅      | Selesai / Complete              |
| 🔄      | Dalam Progress                  |
| ⬜      | Belum Dimulai                   |
| ❌      | Blocked                         |
| **P0**  | Prioritas Tinggi - Blocker      |
| **P1**  | Prioritas Sedang - Penting      |
| **P2**  | Prioritas Rendah - Nice to Have |

---

## RINGKASAN PROGRESS

| Sprint     | Target                  | Progress | Status |
| ---------- | ----------------------- | -------- | ------ |
| Sprint 1   | Infrastruktur & Fondasi | 95%      | 🔄     |
| Sprint 1.5 | Database Fix            | 100%     | ✅     |
| Sprint 2   | Master Data (CRUD)      | 0%       | ⬜     |
| Sprint 3   | Kepegawaian Dasar       | 0%       | ⬜     |
| Sprint 4   | Portal + Cetak SK       | 25%      | 🔄     |
| **TOTAL**  | **100%**                | **~45%** | 🔄     |

---

## SPRINT 1: INFRASTRUKTUR & FONDASI

**Durasi:** 12-14 hari
**Status:** 95% Selesai

### 1.1 Monorepo Setup ⬜→✅

| #     | Task                                    | Prioritas | Status | Owner  | Est. | Notes            |
| ----- | --------------------------------------- | --------- | ------ | ------ | ---- | ---------------- |
| 1.1.1 | Inisialisasi Turborepo + pnpm workspace | P0        | ✅     | DevOps | 2h   | Done             |
| 1.1.2 | Setup turbo.json dengan caching         | P0        | ✅     | DevOps | 1h   | Done             |
| 1.1.3 | Konfigurasi pnpm-workspace.yaml         | P0        | ✅     | DevOps | 0.5h | Done             |
| 1.1.4 | Setup ESLint + Prettier                 | P1        | ✅     | DevOps | 1h   | Fixed 2026-02-24 |
| 1.1.5 | Setup Husky pre-commit hooks            | P1        | 🔄     | DevOps | 1h   | In Progress      |

### 1.2 Docker Services ⬜→✅

| #     | Task                             | Prioritas | Status | Owner  | Est. | Notes                     |
| ----- | -------------------------------- | --------- | ------ | ------ | ---- | ------------------------- |
| 1.2.1 | PostgreSQL 18 multi-database     | P0        | ✅     | DevOps | 2h   | db_master, db_kepegawaian |
| 1.2.2 | Keycloak 26.5.3 setup            | P0        | ✅     | DevOps | 4h   | Realm export ready        |
| 1.2.3 | DragonflyDB cache                | P1        | ✅     | DevOps | 1h   |                           |
| 1.2.4 | Gotenberg 8 PDF engine           | P1        | ✅     | DevOps | 1h   |                           |
| 1.2.5 | Docker Secrets untuk credentials | P0        | ✅     | DevOps | 2h   | docker/secrets/ ready     |
| 1.2.6 | Prometheus monitoring            | P2        | ⬜     | DevOps | 2h   | Deferred to post-Fase 1   |
| 1.2.7 | Grafana dashboard                | P2        | ⬜     | DevOps | 2h   | Deferred to post-Fase 1   |

### 1.3 Shared Packages ⬜→✅

| #      | Task                                 | Prioritas | Status | Owner    | Est. | Notes              |
| ------ | ------------------------------------ | --------- | ------ | -------- | ---- | ------------------ |
| 1.3.1  | @sikerma/ui setup (shadcn/ui)        | P0        | ✅     | Frontend | 4h   |                    |
| 1.3.2  | @sikerma/ui: DataTable component     | P0        | ✅     | Frontend | 3h   |                    |
| 1.3.3  | @sikerma/ui: FormDialog component    | P0        | ✅     | Frontend | 2h   |                    |
| 1.3.4  | @sikerma/ui: Sidebar component       | P0        | ✅     | Frontend | 2h   |                    |
| 1.3.5  | @sikerma/ui: PageHeader component    | P0        | ✅     | Frontend | 1h   |                    |
| 1.3.6  | @sikerma/ui: AppHeader component     | P0        | ✅     | Frontend | 2h   |                    |
| 1.3.7  | @sikerma/ui: Breadcrumb component    | P1        | ✅     | Frontend | 1h   |                    |
| 1.3.8  | @sikerma/ui: StatusBadge component   | P1        | ✅     | Frontend | 1h   |                    |
| 1.3.9  | @sikerma/ui: DeleteConfirm component | P1        | ✅     | Frontend | 1h   |                    |
| 1.3.10 | @sikerma/ui: FileUpload component    | P1        | ✅     | Frontend | 2h   |                    |
| 1.3.11 | @sikerma/ui: StepWizard component    | P1        | 🔄     | Frontend | 3h   | In Progress        |
| 1.3.12 | @sikerma/auth setup (Keycloak)       | P0        | ✅     | Frontend | 4h   | keycloak-js        |
| 1.3.13 | @sikerma/auth: useSession hook       | P0        | ✅     | Frontend | 2h   |                    |
| 1.3.14 | @sikerma/auth: useRoles hook         | P0        | ✅     | Frontend | 2h   |                    |
| 1.3.15 | @sikerma/auth: RoleGuard component   | P1        | ✅     | Frontend | 2h   |                    |
| 1.3.16 | @sikerma/shared: API client          | P0        | ✅     | Frontend | 3h   |                    |
| 1.3.17 | @sikerma/shared: TypeScript types    | P0        | ✅     | Frontend | 4h   | Updated 2026-02-24 |
| 1.3.18 | @sikerma/shared: Zod schemas         | P0        | ✅     | Frontend | 3h   | Updated 2026-02-24 |
| 1.3.19 | @sikerma/shared: Constants           | P1        | ✅     | Frontend | 1h   | Updated 2026-02-24 |
| 1.3.20 | @sikerma/shared: Query Provider      | P0        | ✅     | Frontend | 1h   |                    |

### 1.4 Frontend Apps (Skeleton) ⬜→✅

| #     | Task                              | Prioritas | Status | Owner    | Est. | Notes      |
| ----- | --------------------------------- | --------- | ------ | -------- | ---- | ---------- |
| 1.4.1 | Portal app setup (port 3000)      | P0        | ✅     | Frontend | 2h   |            |
| 1.4.2 | Portal: Layout with AuthProvider  | P0        | ✅     | Frontend | 2h   |            |
| 1.4.3 | Portal: Login page                | P0        | ✅     | Frontend | 3h   |            |
| 1.4.4 | Portal: Dashboard launcher        | P0        | ✅     | Frontend | 4h   |            |
| 1.4.5 | Master Data app setup (port 3001) | P0        | ✅     | Frontend | 1h   | Shell only |
| 1.4.6 | Kepegawaian app setup (port 3002) | P0        | ✅     | Frontend | 1h   | Shell only |

### 1.5 Backend (Go Fiber) ⬜→✅

| #      | Task                        | Prioritas | Status | Owner   | Est. | Notes                |
| ------ | --------------------------- | --------- | ------ | ------- | ---- | -------------------- |
| 1.5.1  | Project structure setup     | P0        | ✅     | Backend | 2h   |                      |
| 1.5.2  | Database connection pool    | P0        | ✅     | Backend | 3h   |                      |
| 1.5.3  | CORS middleware             | P0        | ✅     | Backend | 1h   | Fixed for Fiber v3   |
| 1.5.4  | Auth middleware (JWT)       | P0        | ✅     | Backend | 4h   |                      |
| 1.5.5  | Rate limiting middleware    | P0        | ✅     | Backend | 2h   |                      |
| 1.5.6  | CSRF middleware             | P1        | ✅     | Backend | 1h   | Fixed for Fiber v3   |
| 1.5.7  | Audit trail middleware      | P1        | ✅     | Backend | 3h   |                      |
| 1.5.8  | Security headers middleware | P0        | ✅     | Backend | 1h   | Fixed for Fiber v3   |
| 1.5.9  | Graceful shutdown           | P1        | ✅     | Backend | 1h   |                      |
| 1.5.10 | Health endpoint             | P0        | ✅     | Backend | 0.5h |                      |
| 1.5.11 | Error response standard     | P0        | ✅     | Backend | 2h   | ErrorResponse struct |

### 1.6 Database Migrations ⬜→✅

| #     | Task                         | Prioritas | Status | Owner   | Est. | Notes |
| ----- | ---------------------------- | --------- | ------ | ------- | ---- | ----- |
| 1.6.1 | Create db_master schema      | P0        | ✅     | Backend | 4h   |       |
| 1.6.2 | Create db_kepegawaian schema | P0        | ✅     | Backend | 4h   |       |
| 1.6.3 | Add database indexes         | P0        | ✅     | Backend | 2h   |       |
| 1.6.4 | Add database constraints     | P0        | ✅     | Backend | 1h   |       |
| 1.6.5 | Seed master data referensi   | P0        | ✅     | Backend | 3h   |       |
| 1.6.6 | Seed RBAC data               | P0        | ✅     | Backend | 2h   |       |
| 1.6.7 | Seed sample pegawai (29)     | P1        | ✅     | Backend | 2h   |       |

---

## SPRINT 1.5: DATABASE FIX (P0 ITEMS)

**Durasi:** 1-2 hari
**Status:** 100% Selesai ✅
**Dependency:** Harus selesai sebelum Sprint 2

### 1.5.1 Schema Fixes ⬜→✅

| #       | Task                                        | Prioritas | Status | Owner   | Est. | Notes              |
| ------- | ------------------------------------------- | --------- | ------ | ------- | ---- | ------------------ |
| 1.5.1.1 | Tambah field eselon_id ke pegawai           | P0        | ✅     | Backend | 0.5h | Migration 05 ready |
| 1.5.1.2 | Tambah field tmt_cpns, tmt_pns              | P0        | ✅     | Backend | 0.5h | Migration 05 ready |
| 1.5.1.3 | Tambah field tmt_pangkat_terakhir           | P0        | ✅     | Backend | 0.5h | Migration 05 ready |
| 1.5.1.4 | Tambah field status_kerja                   | P0        | ✅     | Backend | 0.5h | Migration 05 ready |
| 1.5.1.5 | Perbaiki constraint status_pegawai          | P0        | ✅     | Backend | 0.5h | Migration 05 ready |
| 1.5.1.6 | Tambah field audit (created_by, updated_by) | P1        | ✅     | Backend | 0.5h | Migration 05 ready |

### 1.5.2 New Master Data ⬜→✅

| #       | Task                                | Prioritas | Status | Owner   | Est. | Notes              |
| ------- | ----------------------------------- | --------- | ------ | ------- | ---- | ------------------ |
| 1.5.2.1 | Buat tabel ref_golongan_non_pns     | P1        | ✅     | Backend | 1h   | Migration 06 ready |
| 1.5.2.2 | Seed golongan non-PNS (I, V, IX)    | P1        | ✅     | Backend | 0.5h | Migration 06 ready |
| 1.5.2.3 | Lengkapi seed jabatan (20+ jabatan) | P0        | ✅     | Backend | 1h   | Migration 07 ready |

### 1.5.3 Type Updates ⬜→✅

| #       | Task                                       | Prioritas | Status | Owner    | Est. | Notes           |
| ------- | ------------------------------------------ | --------- | ------ | -------- | ---- | --------------- |
| 1.5.3.1 | Update TypeScript types di @sikerma/shared | P0        | ✅     | Frontend | 2h   | Done 2026-02-24 |
| 1.5.3.2 | Update Zod validation schemas              | P0        | ✅     | Frontend | 2h   | Done 2026-02-24 |
| 1.5.3.3 | Update backend Go models                   | P0        | ✅     | Backend  | 2h   | Done 2026-02-24 |

---

## SPRINT 2: MASTER DATA (CRUD LENGKAP)

**Durasi:** 8-10 hari
**Status:** 0% Selesai
**Dependency:** Sprint 1 + Sprint 1.5 ✅

### 2.1 Backend - Master Data CRUD ⬜

| #      | Task                               | Prioritas | Status | Owner   | Est. | Notes            |
| ------ | ---------------------------------- | --------- | ------ | ------- | ---- | ---------------- |
| 2.1.1  | Generic CRUD handler               | P0        | ⬜     | Backend | 4h   | Reusable pattern |
| 2.1.2  | Satker CRUD handler                | P0        | ⬜     | Backend | 2h   |                  |
| 2.1.3  | Jabatan CRUD handler               | P0        | ⬜     | Backend | 2h   |                  |
| 2.1.4  | Golongan CRUD handler              | P0        | ⬜     | Backend | 2h   |                  |
| 2.1.5  | Unit Kerja CRUD handler            | P0        | ⬜     | Backend | 2h   | Tree support     |
| 2.1.6  | Eselon CRUD handler                | P1        | ⬜     | Backend | 1h   |                  |
| 2.1.7  | Pendidikan CRUD handler            | P1        | ⬜     | Backend | 1h   |                  |
| 2.1.8  | Agama CRUD handler                 | P1        | ⬜     | Backend | 1h   |                  |
| 2.1.9  | Status Kawin CRUD handler          | P1        | ⬜     | Backend | 1h   |                  |
| 2.1.10 | Hukdis CRUD handler                | P2        | ⬜     | Backend | 1h   |                  |
| 2.1.11 | Jenis Diklat CRUD handler          | P2        | ⬜     | Backend | 1h   |                  |
| 2.1.12 | Dropdown endpoints (semua entitas) | P0        | ⬜     | Backend | 2h   |                  |

### 2.2 Frontend - Master Data Pages ⬜

| #      | Task                           | Prioritas | Status | Owner    | Est. | Notes             |
| ------ | ------------------------------ | --------- | ------ | -------- | ---- | ----------------- |
| 2.2.1  | Dashboard Master Data          | P2        | ⬜     | Frontend | 3h   | Card grid         |
| 2.2.2  | Satker page: list + search     | P0        | ⬜     | Frontend | 4h   | DataTable         |
| 2.2.3  | Satker page: form dialog       | P0        | ⬜     | Frontend | 3h   | FormDialog        |
| 2.2.4  | Jabatan page: list + search    | P0        | ⬜     | Frontend | 4h   |                   |
| 2.2.5  | Jabatan page: form dialog      | P0        | ⬜     | Frontend | 3h   |                   |
| 2.2.6  | Golongan page: list + search   | P0        | ⬜     | Frontend | 3h   |                   |
| 2.2.7  | Golongan page: form dialog     | P0        | ⬜     | Frontend | 2h   |                   |
| 2.2.8  | Unit Kerja page: tree view     | P0        | ⬜     | Frontend | 6h   | Special component |
| 2.2.9  | Unit Kerja page: form dialog   | P0        | ⬜     | Frontend | 3h   |                   |
| 2.2.10 | Eselon page: list + form       | P1        | ⬜     | Frontend | 3h   |                   |
| 2.2.11 | Pendidikan page: list + form   | P1        | ⬜     | Frontend | 3h   |                   |
| 2.2.12 | Agama page: list + form        | P1        | ⬜     | Frontend | 2h   |                   |
| 2.2.13 | Status Kawin page: list + form | P1        | ⬜     | Frontend | 2h   |                   |
| 2.2.14 | Hukdis page: list + form       | P2        | ⬜     | Frontend | 3h   |                   |
| 2.2.15 | Jenis Diklat page: list + form | P2        | ⬜     | Frontend | 3h   |                   |

### 2.3 Testing - Master Data ⬜

| #     | Task                     | Prioritas | Status | Owner    | Est. | Notes |
| ----- | ------------------------ | --------- | ------ | -------- | ---- | ----- |
| 2.3.1 | Backend unit tests       | P1        | ⬜     | Backend  | 4h   |       |
| 2.3.2 | Frontend component tests | P2        | ⬜     | Frontend | 4h   |       |
| 2.3.3 | E2E tests                | P2        | ⬜     | QA       | 4h   |       |

---

## SPRINT 3: KEPEGAWAIAN DASAR

**Durasi:** 12-14 hari
**Status:** 0% Selesai
**Dependency:** Sprint 2

### 3.1 Backend - Pegawai CRUD ⬜

| #      | Task                          | Prioritas | Status | Owner   | Est. | Notes |
| ------ | ----------------------------- | --------- | ------ | ------- | ---- | ----- |
| 3.1.1  | Pegawai CRUD handler          | P0        | ⬜     | Backend | 4h   |       |
| 3.1.2  | Pegawai search/filter         | P0        | ⬜     | Backend | 3h   |       |
| 3.1.3  | NIP validation (18 digit)     | P0        | ⬜     | Backend | 2h   |       |
| 3.1.4  | Riwayat Pangkat CRUD          | P0        | ⬜     | Backend | 3h   |       |
| 3.1.5  | Riwayat Jabatan CRUD          | P0        | ⬜     | Backend | 3h   |       |
| 3.1.6  | Riwayat Pendidikan CRUD       | P1        | ⬜     | Backend | 2h   |       |
| 3.1.7  | Keluarga CRUD                 | P1        | ⬜     | Backend | 2h   |       |
| 3.1.8  | File upload handler           | P0        | ⬜     | Backend | 4h   |       |
| 3.1.9  | File validation (magic bytes) | P0        | ⬜     | Backend | 2h   |       |
| 3.1.10 | Statistik endpoint            | P1        | ⬜     | Backend | 3h   |       |

### 3.2 Frontend - Kepegawaian Pages ⬜

| #      | Task                             | Prioritas | Status | Owner    | Est. | Notes               |
| ------ | -------------------------------- | --------- | ------ | -------- | ---- | ------------------- |
| 3.2.1  | Dashboard Kepegawaian            | P1        | ⬜     | Frontend | 4h   | Charts              |
| 3.2.2  | Daftar Pegawai page              | P0        | ⬜     | Frontend | 6h   | DataTable + filters |
| 3.2.3  | Pegawai Detail page (tabs)       | P0        | ⬜     | Frontend | 6h   | Tab view            |
| 3.2.4  | Tambah Pegawai form (multi-step) | P0        | ⬜     | Frontend | 8h   | StepWizard          |
| 3.2.5  | Edit Biodata page                | P0        | ⬜     | Frontend | 4h   |                     |
| 3.2.6  | Upload Foto dialog               | P1        | ⬜     | Frontend | 3h   |                     |
| 3.2.7  | Riwayat Pangkat tab              | P0        | ⬜     | Frontend | 4h   | Inline CRUD         |
| 3.2.8  | Riwayat Jabatan tab              | P0        | ⬜     | Frontend | 4h   | Inline CRUD         |
| 3.2.9  | Riwayat Pendidikan tab           | P1        | ⬜     | Frontend | 3h   | Inline CRUD         |
| 3.2.10 | Keluarga tab                     | P1        | ⬜     | Frontend | 3h   | Inline CRUD         |
| 3.2.11 | Nonaktifkan Pegawai dialog       | P1        | ⬜     | Frontend | 2h   |                     |

### 3.3 Testing - Kepegawaian ⬜

| #     | Task                     | Prioritas | Status | Owner    | Est. | Notes |
| ----- | ------------------------ | --------- | ------ | -------- | ---- | ----- |
| 3.3.1 | Backend unit tests       | P1        | ⬜     | Backend  | 6h   |       |
| 3.3.2 | Frontend component tests | P2        | ⬜     | Frontend | 6h   |       |
| 3.3.3 | E2E tests                | P1        | ⬜     | QA       | 6h   |       |

---

## SPRINT 4: PORTAL + CETAK SK + POLISH

**Durasi:** 10-12 hari
**Status:** 25% Selesai
**Dependency:** Sprint 1 + Sprint 3

### 4.1 Portal - Admin Pages ⬜

| #     | Task                     | Prioritas | Status | Owner    | Est. | Notes           |
| ----- | ------------------------ | --------- | ------ | -------- | ---- | --------------- |
| 4.1.1 | Dashboard widgets        | P2        | ⬜     | Frontend | 6h   | Statistics      |
| 4.1.2 | Profil Saya page         | P1        | ⬜     | Frontend | 4h   |                 |
| 4.1.3 | Admin: User Management   | P1        | ⬜     | Frontend | 6h   | Keycloak proxy  |
| 4.1.4 | Admin: Hak Akses page    | P0        | ⬜     | Frontend | 6h   | Role assignment |
| 4.1.5 | Admin: Kelola Roles page | P1        | ⬜     | Frontend | 6h   | CRUD roles      |
| 4.1.6 | Admin: Audit Log page    | P1        | ⬜     | Frontend | 4h   |                 |

### 4.2 Backend - Admin & Cetak ⬜

| #     | Task                        | Prioritas | Status | Owner   | Est. | Notes |
| ----- | --------------------------- | --------- | ------ | ------- | ---- | ----- |
| 4.2.1 | Keycloak admin API proxy    | P1        | ⬜     | Backend | 4h   |       |
| 4.2.2 | Dashboard summary endpoint  | P2        | ⬜     | Backend | 3h   |       |
| 4.2.3 | Template dokumen CRUD       | P1        | ⬜     | Backend | 4h   |       |
| 4.2.4 | PDF generation (Gotenberg)  | P1        | ⬜     | Backend | 6h   |       |
| 4.2.5 | Template placeholder parser | P1        | ⬜     | Backend | 4h   |       |

### 4.3 Kepegawaian - Cetak SK ⬜

| #     | Task                     | Prioritas | Status | Owner    | Est. | Notes          |
| ----- | ------------------------ | --------- | ------ | -------- | ---- | -------------- |
| 4.3.1 | Template Management page | P1        | ⬜     | Frontend | 4h   |                |
| 4.3.2 | Cetak SK page            | P1        | ⬜     | Frontend | 6h   | Template → PDF |
| 4.3.3 | Preview PDF component    | P1        | ⬜     | Frontend | 4h   |                |
| 4.3.4 | Download PDF handler     | P1        | ⬜     | Frontend | 2h   |                |

### 4.4 Testing & Polish ⬜

| #     | Task                     | Prioritas | Status | Owner    | Est. | Notes |
| ----- | ------------------------ | --------- | ------ | -------- | ---- | ----- |
| 4.4.1 | E2E testing all flows    | P0        | ⬜     | QA       | 8h   |       |
| 4.4.2 | Bug fixing               | P0        | ⬜     | All      | 8h   |       |
| 4.4.3 | Performance optimization | P1        | ⬜     | All      | 4h   |       |
| 4.4.4 | Accessibility review     | P2        | ⬜     | Frontend | 4h   |       |
| 4.4.5 | Documentation update     | P2        | ⬜     | All      | 4h   |       |

---

## DEFERRED / POST-FASE 1

| #   | Task                       | Priority | Notes          |
| --- | -------------------------- | -------- | -------------- |
| D1  | Prometheus + Grafana setup | P2       | Monitoring     |
| D2  | ClamAV integration         | P2       | File scanning  |
| D3  | Mobile responsive          | P3       | Enhancement    |
| D4  | Multi-satker support       | P3       | Future feature |
| D5  | SIKEP integration          | P3       | Future feature |

---

## TEST RESULTS (2026-02-24)

| Category     | Test Type             | Result  | Details               |
| ------------ | --------------------- | ------- | --------------------- |
| **Frontend** | TypeScript Type-Check | ✅ PASS | 6 packages            |
| **Frontend** | ESLint                | ✅ PASS | 0 errors, 36 warnings |
| **Frontend** | Vitest Unit Tests     | ✅ PASS | 29 tests              |
| **Backend**  | Go Build              | ✅ PASS | Compiles successfully |
| **Database** | Migrations            | ✅ PASS | 7 scripts ready       |

---

## NOTES

### Testing Strategy

1. **Unit Tests**: Vitest untuk frontend, Go testing untuk backend
2. **Component Tests**: React Testing Library
3. **E2E Tests**: Playwright
4. **Mutation Tests**: Stryker (optional)

### Deployment Strategy

1. Development: Docker Compose local
2. Staging: PM2 + Nginx
3. Production: PM2 cluster + Nginx + SSL

### Key Dependencies

- Sprint 2 requires Sprint 1.5 (database fix) ✅ DONE
- Sprint 3 requires Sprint 2 (master data ready)
- Sprint 4 can start parallel with Sprint 3 (after Sprint 1)

### Recent Commits

| Commit  | Description                                            | Date       |
| ------- | ------------------------------------------------------ | ---------- |
| a5e2ef0 | docs: Update TEST_SUMMARY.md with current test results | 2026-02-24 |
| ac3fc08 | feat: Update database schema, models, and shared types | 2026-02-24 |
| 8b6bda2 | fix: Resolve compilation and linting issues            | 2026-02-24 |

---

**Last Updated:** 24 Februari 2026
**Next Review:** Weekly sprint review
