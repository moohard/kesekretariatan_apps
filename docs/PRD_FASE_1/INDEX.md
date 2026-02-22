# INDEX TASK - FASE 1 SIKERMA

Dokumen ini adalah index/peta navigasi untuk semua file task Fase 1.

---

## 📁 STRUKTUR FOLDER

```
docs/PRD_FASE_1/
│
├── 📄 README.md                           # Ringkasan & overview Fase 1
├── 📄 INDEX.md                            # File ini - Peta navigasi
├── 📄 RINGKASAN_TASK.md                   # Ringkasan eksekutif (1 halaman)
├── 📄 ACCEPTANCE_CRITERIA.md             # Definition of Done lengkap
│
├── 🏃 SPRINTS (Planning per Sprint)
│   ├── 📄 SPRINT_1_INFRASTRUKTUR.md      # Sprint 1: Infrastruktur & Fondasi
│   ├── 📄 SPRINT_2_MASTER_DATA.md        # Sprint 2: Master Data (CRUD Lengkap)
│   ├── 📄 SPRINT_3_KEPEGAWAIAN.md        # Sprint 3: Kepegawaian Dasar
│   └── 📄 SPRINT_4_PORTAL_CETAK_SK.md    # Sprint 4: Portal + Cetak SK + Polish
│
└── 📝 TASKS (Detail per Task)
    │
    ├── 🔧 SPRINT 1: Infrastruktur & Fondasi
    │   ├── 📄 TASK_001_SETUP_MONOREPO.md          # Setup Monorepo (Turborepo + pnpm)
    │   ├── 📄 TASK_002_SETUP_SHARED_PKGS.md       # Setup 3 Shared Packages
    │   ├── 📄 TASK_003_SETUP_NEXTJS_APPS.md       # Setup 3 Next.js Apps
    │   ├── 📄 TASK_004_SETUP_GO_BACKEND.md        # Bootstrap Go Fiber Backend
    │   ├── 📄 TASK_005_SETUP_KEYCLOAK.md          # Setup Keycloak Realm & Clients
    │   ├── 📄 TASK_006_SETUP_DATABASE.md          # Setup Database Migration & Seed
    │   ├── 📄 TASK_007_SETUP_GOTENBERG.md         # Setup Gotenberg Service
    │   ├── 📄 TASK_008_SETUP_AUTH_FLOW.md         # Implementasi Auth Flow
    │   ├── 📄 TASK_009_SETUP_RBAC.md              # Implementasi RBAC Middleware
    │   ├── 📄 TASK_010_SETUP_AUDIT.md             # Implementasi Audit Trail Middleware
    │   └── 📄 TASK_011_SETUP_UI_COMPONENTS.md     # Setup 10 Shared UI Components
    │
    ├── 📊 SPRINT 2: Master Data
    │   ├── 📄 TASK_012_MASTER_CRUD_API.md         # Generic CRUD Backend (10 entitas)
    │   └── 📄 TASK_013_MASTER_FRONTEND.md         # Master Data Frontend (10 halaman)
    │
    ├── 👥 SPRINT 3: Kepegawaian
    │   ├── 📄 TASK_014_KEPEGAWAIAN_CRUD.md        # Kepegawaian CRUD Backend
    │   └── 📄 TASK_015_KEPEGAWAIAN_FRONTEND.md    # Kepegawaian Frontend (5 halaman)
    │
    └── 🚀 SPRINT 4: Portal + Cetak SK
        ├── 📄 TASK_016_PORTAL_LAUNCHER.md         # Portal Launcher & Dashboard
        ├── 📄 TASK_017_PORTAL_ADMIN.md            # Portal Admin UI (User, RBAC, Audit)
        └── 📄 TASK_018_CETAK_SK.md                # Cetak SK & Template Management
```

---

## 📖 CARA MENGGUNAKAN DOKUMEN INI

### Untuk Project Manager / Product Owner:
1. Baca `README.md` untuk overview Fase 1
2. Baca `RINGKASAN_TASK.md` untuk ringkasan eksekutif
3. Baca file sprint (`SPRINT_X_*.md`) untuk planning per sprint
4. Track progress menggunakan checklist di setiap sprint file

### Untuk Developer:
1. Baca `ACCEPTANCE_CRITERIA.md` untuk memahami Definition of Done
2. Baca file sprint yang sedang dikerjakan (`SPRINT_X_*.md`)
3. Baca file task detail (`TASK_XXX_*.md`) untuk implementasi spesifik
4. Implementasi sesuai acceptance criteria di task file

### Untuk QA / Tester:
1. Baca `ACCEPTANCE_CRITERIA.md` untuk test cases
2. Baca file sprint untuk memahami deliverables per sprint
3. Gunakan checklist di `ACCEPTANCE_CRITERIA.md` untuk verifikasi

---

## 📋 QUICK LINKS PER SPRINT

### SPRINT 1: Infrastruktur & Fondasi (8-10 hari)
**Status:** ⏸️ Not Started
**Priority:** 🔴 Critical Path (Blocking)

**Planning:** [`SPRINT_1_INFRASTRUKTUR.md`](./SPRINT_1_INFRASTRUKTUR.md)

**Tasks:**
1. [`TASK_001_SETUP_MONOREPO.md`](./TASK_001_SETUP_MONOREPO.md) - Setup Monorepo
2. [`TASK_002_SETUP_SHARED_PKGS.md`](./TASK_002_SETUP_SHARED_PKGS.md) - Setup Shared Packages
3. [`TASK_003_SETUP_NEXTJS_APPS.md`](./TASK_003_SETUP_NEXTJS_APPS.md) - Setup 3 Next.js Apps
4. [`TASK_004_SETUP_GO_BACKEND.md`](./TASK_004_SETUP_GO_BACKEND.md) - Bootstrap Go Backend
5. [`TASK_005_SETUP_KEYCLOAK.md`](./TASK_005_SETUP_KEYCLOAK.md) - Setup Keycloak
6. [`TASK_006_SETUP_DATABASE.md`](./TASK_006_SETUP_DATABASE.md) - Setup Database
7. [`TASK_007_SETUP_GOTENBERG.md`](./TASK_007_SETUP_GOTENBERG.md) - Setup Gotenberg
8. [`TASK_008_SETUP_AUTH_FLOW.md`](./TASK_008_SETUP_AUTH_FLOW.md) - Setup Auth Flow
9. [`TASK_009_SETUP_RBAC.md`](./TASK_009_SETUP_RBAC.md) - Setup RBAC Middleware
10. [`TASK_010_SETUP_AUDIT.md`](./TASK_010_SETUP_AUDIT.md) - Setup Audit Trail
11. [`TASK_011_SETUP_UI_COMPONENTS.md`](./TASK_011_SETUP_UI_COMPONENTS.md) - Setup UI Components

**Deliverables:**
- ✅ Monorepo siap
- ✅ Backend API server
- ✅ Keycloak SSO
- ✅ Database ter-migrasi
- ✅ Auth flow berfungsi
- ✅ 10 UI components

---

### SPRINT 2: Master Data (6-8 hari)
**Status:** ⏸️ Blocked (Dependency: Sprint 1)
**Priority:** 🔴 High

**Planning:** [`SPRINT_2_MASTER_DATA.md`](./SPRINT_2_MASTER_DATA.md)

**Tasks:**
12. [`TASK_012_MASTER_CRUD_API.md`](./TASK_012_MASTER_CRUD_API.md) - Generic CRUD Backend
13. [`TASK_013_MASTER_FRONTEND.md`](./TASK_013_MASTER_FRONTEND.md) - Master Data Frontend

**Deliverables:**
- ✅ 60 API endpoints (6 × 10 entities)
- ✅ 10 halaman CRUD frontend
- ✅ Dropdown API untuk semua entitas
- ✅ Data referensi ter-seed

**Entities (10):**
1. Satuan Kerja
2. Jabatan
3. Golongan/Pangkat
4. Unit Kerja (Tree)
5. Eselon
6. Pendidikan
7. Agama
8. Status Kawin
9. Hukuman Disiplin
10. Jenis Diklat

---

### SPRINT 3: Kepegawaian (8-10 hari)
**Status:** ⏸️ Blocked (Dependency: Sprint 2)
**Priority:** 🔴 High

**Planning:** [`SPRINT_3_KEPEGAWAIAN.md`](./SPRINT_3_KEPEGAWAIAN.md)

**Tasks:**
14. [`TASK_014_KEPEGAWAIAN_CRUD.md`](./TASK_014_KEPEGAWAIAN_CRUD.md) - CRUD Backend
15. [`TASK_015_KEPEGAWAIAN_FRONTEND.md`](./TASK_015_KEPEGAWAIAN_FRONTEND.md) - Frontend

**Deliverables:**
- ✅ CRUD Pegawai (validasi NIP)
- ✅ Multi-step wizard (4 step)
- ✅ Detail dengan tab view (5 tabs)
- ✅ CRUD 4 riwayat (Pangkat, Jabatan, Pendidikan, Keluarga)
- ✅ File upload (Foto, SK, Ijazah)
- ✅ Dashboard statistik + chart
- ✅ 29 pegawai ter-verifikasi

**Pages:**
- `/` - Dashboard
- `/pegawai` - Daftar pegawai
- `/pegawai/tambah` - Form wizard
- `/pegawai/[nip]` - Detail (5 tabs)
- `/pegawai/[nip]/edit` - Edit

---

### SPRINT 4: Portal + Cetak SK (6-8 hari)
**Status:** ⏸️ Blocked (Dependency: Sprint 1, 3)
**Priority:** 🔴 High

**Planning:** [`SPRINT_4_PORTAL_CETAK_SK.md`](./SPRINT_4_PORTAL_CETAK_SK.md)

**Tasks:**
16. [`TASK_016_PORTAL_LAUNCHER.md`](./TASK_016_PORTAL_LAUNCHER.md) - Portal Launcher
17. [`TASK_017_PORTAL_ADMIN.md`](./TASK_017_PORTAL_ADMIN.md) - Admin UI
18. [`TASK_018_CETAK_SK.md`](./TASK_018_CETAK_SK.md) - Cetak SK

**Deliverables:**
- ✅ Login → Keycloak → session
- ✅ Dashboard launcher (tile app)
- ✅ 4+ dashboard widgets
- ✅ Admin User Management
- ✅ Admin Hak Akses (RBAC UI)
- ✅ Admin Kelola Roles
- ✅ Audit Log viewer
- ✅ Template management
- ✅ Cetak SK (≥3 template → PDF)
- ✅ End-to-end testing

**Pages:**
- `/login` - Login
- `/` - Dashboard
- `/profil` - Profil Saya
- `/admin` - Admin dashboard
- `/admin/users` - User Management
- `/admin/hak-akses` - Assign Role
- `/admin/hak-akses/roles` - Kelola Roles
- `/admin/audit-log` - Audit Log
- `/cetak-sk` - Cetak SK
- `/admin/template` - Template Management

---

## 🎯 ACCEPTANCE CRITERIA

File lengkap: [`ACCEPTANCE_CRITERIA.md`](./ACCEPTANCE_CRITERIA.md)

**Quick Checklist:**

### Sprint 1
- [ ] `pnpm install` sukses
- [ ] Backend listen di port 3003
- [ ] Keycloak realm + 3 clients ready
- [ ] 2 DB + 21 tabel ter-migrasi
- [ ] Login → Keycloak → session berfungsi
- [ ] 10 shared components terbuat

### Sprint 2
- [ ] 60 API endpoints working
- [ ] 10 halaman CRUD frontend
- [ ] Dropdown API untuk semua entitas
- [ ] Data referensi ter-seed

### Sprint 3
- [ ] CRUD pegawai (validasi NIP)
- [ ] Multi-step wizard 4 step
- [ ] 4 jenis riwayat CRUD
- [ ] File upload berfungsi
- [ ] Dashboard statistik + chart
- [ ] 29 pegawai ter-verifikasi

### Sprint 4
- [ ] Portal launcher berfungsi
- [ ] 4+ dashboard widgets
- [ ] Admin UI (User, RBAC, Audit)
- [ ] Cetak SK (≥3 template)
- [ ] Zero P0/P1 bugs
- [ ] End-to-end testing pass

### Overall Fase 1
- [ ] 29 pegawai ter-migrasi
- [ ] SSO login 1x → akses 3 app
- [ ] Master Data CRUD lengkap
- [ ] Kepegawaian CRUD lengkap
- [ ] Minimal 3 template SK → PDF
- [ ] RBAC berfungsi
- [ ] Audit trail mencatat semua
- [ ] Portal dashboard informatif

---

## 📊 PROGRESS TRACKING

### Sprint Completion Status

| Sprint | Tasks | Estimasi | Status | Progress |
|--------|-------|----------|--------|----------|
| 1 | 11 tasks | 8-10 hari | ⏸️ Not Started | 0% |
| 2 | 2 tasks | 6-8 hari | ⏸️ Blocked | 0% |
| 3 | 2 tasks | 8-10 hari | ⏸️ Blocked | 0% |
| 4 | 3 tasks | 6-8 hari | ⏸️ Blocked | 0% |
| **TOTAL** | **18 tasks** | **28-36 hari** | | **0%** |

### Task Completion Status

```bash
# Quick status check (manual update)
Sprint 1: [░░░░░░░░░░] 0/11
Sprint 2: [░░░░░░░░░░] 0/2
Sprint 3: [░░░░░░░░░░] 0/2
Sprint 4: [░░░░░░░░░░] 0/3

Overall: [░░░░░░░░░░] 0/18 (0%)
```

---

## 🔗 REFERENSI EKSTERNAL

### Dokumen PRD & Planning
- **PRD Fase 1:** [`../PRD_FASE_1.md`](../PRD_FASE_1.md)
- **Blueprint Architecture:** [`../blueprint_arch.md`](../blueprint_arch.md)
- **Overview Aplikasi:** [`../overview_aplikasi.md`](../overview_aplikasi.md)

### Source Data (Seed)
- **Data Pegawai (29 orang):** `data_pegawai.json`
- **Struktur Organisasi:** `org_structure.json`
- **Keycloak Realm Export:** `realm-export.json`

### External Documentation
- **Next.js 14:** https://nextjs.org/docs
- **Go Fiber v3:** https://docs.gofiber.io/
- **Keycloak 26:** https://www.keycloak.org/documentation
- **Gotenberg 8:** https://gotenberg.dev/
- **shadcn/ui:** https://ui.shadcn.com/
- **Better Auth:** https://www.better-auth.com/

---

## 📝 VERSION HISTORY

| Version | Date | Changes | Author |
|---------|------|---------|--------|
| 1.0 | 22 Feb 2026 | Initial creation | Muhardiansyah |
| | | - 18 tasks breakdown | |
| | | - 4 sprint planning | |
| | | - Complete index | |

---

## ✅ CHECKLIST: SEBELUM MULAI FASE 1

Sebelum memulai implementasi, pastikan:

### Environment Setup
- [ ] Node.js ≥ 18.x terinstall
- [ ] pnpm ≥ 8.x terinstall
- [ ] Go 1.21+ terinstall
- [ ] Docker + Docker Compose terinstall
- [ ] Git terinstall & configured
- [ ] Code editor (VS Code) siap dengan extensions:
  - [ ] Go extension
  - [ ] ESLint
  - [ ] Prettier
  - [ ] Tailwind CSS IntelliSense

### Repository
- [ ] Repository sudah di-clone
- [ ] Branch `fase-1` sudah dibuat
- [ ] `pnpm install` sudah dijalankan
- [ ] Docker containers sudah di-start (`docker-compose up -d`)

### Documentation
- [ ] PRD Fase 1 sudah dibaca
- [ ] Blueprint architecture sudah dibaca
- [ ] Task files sudah direview
- [ ] Acceptance criteria sudah dipahami

### Team Alignment
- [ ] Kickoff meeting sudah dilakukan
- [ ] Task assignment sudah jelas
- [ ] Communication channel sudah diset (Slack/Teams)
- [ ] Daily standup schedule sudah diset

---

**Dokumen ini adalah single source of truth untuk Task Fase 1 SIKERMA.**
**Update status progress secara berkala.**

**Last Updated:** 22 Februari 2026
**Maintained by:** Tim Pengembang SIKERMA
