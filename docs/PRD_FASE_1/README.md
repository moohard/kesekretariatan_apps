# FASE 1: Task Breakdown

Dokumen ini memecah PRD Fase 1 menjadi task-task implementasi yang dapat dilacak per sprint.

## Struktur Folder

```
PRD_FASE_1/
├── README.md                           # Ringkasan & overview
├── SPRINT_1_INFRASTRUKTUR.md          # Sprint 1: Infrastruktur & Fondasi
├── SPRINT_2_MASTER_DATA.md            # Sprint 2: Master Data (CRUD Lengkap)
├── SPRINT_3_KEPEGAWAIAN.md            # Sprint 3: Kepegawaian Dasar
├── SPRINT_4_PORTAL_CETAK_SK.md        # Sprint 4: Portal + Cetak SK + Polish
├── TASK_001_SETUP_MONOREPO.md         # Task 1: Inisialisasi Monorepo
├── TASK_002_SETUP_SHARED_PKGS.md      # Task 2: Setup Shared Packages
├── TASK_003_SETUP_NEXTJS_APPS.md      # Task 3: Setup 3 Next.js Apps
├── TASK_004_SETUP_GO_BACKEND.md       # Task 4: Bootstrap Go Fiber Backend
├── TASK_005_SETUP_KEYCLOAK.md         # Task 5: Konfigurasi Keycloak Realm
├── TASK_006_SETUP_DATABASE.md         # Task 6: Database Migration & Seed
├── TASK_007_SETUP_GOTENBERG.md        # Task 7: Setup Gotenberg Service
├── TASK_008_SETUP_AUTH_FLOW.md        # Task 8: Implementasi Auth Flow
├── TASK_009_SETUP_RBAC.md             # Task 9: Implementasi RBAC Middleware
├── TASK_010_SETUP_AUDIT.md            # Task 10: Implementasi Audit Trail
├── TASK_011_SETUP_UI_COMPONENTS.md    # Task 11: Shared UI Components
├── TASK_012_MASTER_CRUD_API.md        # Task 12: Generic CRUD Backend Master Data
├── TASK_013_MASTER_FRONTEND.md        # Task 13: Master Data Frontend
├── TASK_014_KEPEGAWAIAN_CRUD.md       # Task 14: Kepegawaian CRUD Backend
├── TASK_015_KEPEGAWAIAN_FRONTEND.md   # Task 15: Kepegawaian Frontend
├── TASK_016_PORTAL_LAUNCHER.md        # Task 16: Portal Launcher & Dashboard
├── TASK_017_PORTAL_ADMIN.md           # Task 17: Portal Admin UI
├── TASK_018_CETAK_SK.md               # Task 18: Cetak SK & Template Management
└── ACCEPTANCE_CRITERIA.md             # Definition of Done & Acceptance Criteria
```

## Sprint Overview

### Sprint 1: Infrastruktur & Fondasi (Task 1-11)
**Dependency:** Tidak ada (starting point)
**Deliverable:** Monorepo siap, backend berjalan, auth berfungsi
**Estimasi:** 8-10 hari kerja

### Sprint 2: Master Data (CRUD Lengkap) (Task 12-13)
**Dependency:** Sprint 1
**Deliverable:** App Master Data fully functional
**Estimasi:** 6-8 hari kerja

### Sprint 3: Kepegawaian Dasar (Task 14-15)
**Dependency:** Sprint 2
**Deliverable:** App Kepegawaian functional untuk CRUD pegawai + riwayat
**Estimasi:** 8-10 hari kerja

### Sprint 4: Portal + Cetak SK + Polish (Task 16-18)
**Dependency:** Sprint 1, Sprint 3
**Deliverable:** Portal sebagai launcher, cetak SK, RBAC admin UI
**Estimasi:** 6-8 hari kerja

## Total Estimasi Fase 1
**Timeline:** 28-36 hari kerja (~6-8 minggu)

## Prioritas Task
- **P0 (Critical Path):** Task yang blocking untuk sprint berikutnya
- **P1 (High Priority):** Task penting tapi tidak blocking
- **P2 (Nice to Have):** Task enhancement/tambahan

## Cara Menggunakan
1. Mulai dari Sprint 1, kerjakan task berurutan sesuai dependency
2. Setiap task memiliki acceptance criteria yang jelas
3. Task file berisi detail implementasi, file-file yang perlu dibuat/modified
4. Setelah selesai sprint, review `ACCEPTANCE_CRITERIA.md` untuk verifikasi

## Checklist Per Sprint

### Sprint 1
- [ ] Task 001: Setup Monorepo
- [ ] Task 002: Setup Shared Packages
- [ ] Task 003: Setup 3 Next.js Apps
- [ ] Task 004: Bootstrap Go Fiber Backend
- [ ] Task 005: Setup Keycloak Realm
- [ ] Task 006: Setup Database Migration
- [ ] Task 007: Setup Gotenberg
- [ ] Task 008: Setup Auth Flow
- [ ] Task 009: Setup RBAC Middleware
- [ ] Task 010: Setup Audit Trail
- [ ] Task 011: Setup Shared UI Components

### Sprint 2
- [ ] Task 012: Generic CRUD Backend Master Data
- [ ] Task 013: Master Data Frontend (10 entitas CRUD)

### Sprint 3
- [ ] Task 014: Kepegawaian CRUD Backend
- [ ] Task 015: Kepegawaian Frontend

### Sprint 4
- [ ] Task 016: Portal Launcher & Dashboard
- [ ] Task 017: Portal Admin UI (User Management, Roles, Audit)
- [ ] Task 018: Cetak SK & Template Management
