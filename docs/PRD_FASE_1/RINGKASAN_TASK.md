# RINGKASAN TASK - FASE 1 SIKERMA

Dokumen ini adalah ringkasan eksekutif dari seluruh task Fase 1.

---

## Total Task: 18 Tasks
## Total Sprint: 4 Sprints
## Estimasi Timeline: 28-36 hari kerja (~6-8 minggu)

---

## BREAKDOWN PER SPRINT

### SPRINT 1: Infrastruktur & Fondasi (8-10 hari)
**Critical Path:** ✅ Ya - Blocking sprint berikutnya

| # | Task | Estimasi | Status | FR |
|---|------|----------|--------|-----|
| 1 | Setup Monorepo (Turborepo + pnpm) | 1 hari | ◻ | FR-001 |
| 2 | Setup Shared Packages (@sikerma/ui, @sikerma/auth, @sikerma/shared) | 1.5 hari | ◻ | FR-002, FR-003, FR-004 |
| 3 | Setup 3 Next.js Apps | 1 hari | ◻ | FR-001 |
| 4 | Bootstrap Go Fiber Backend | 1.5 hari | ◻ | FR-007 |
| 5 | Setup Keycloak Realm & Clients | 1 hari | ◻ | FR-005 |
| 6 | Setup Database Migration & Seed | 1.5 hari | ◻ | FR-006, FR-010 |
| 7 | Setup Gotenberg Service | 0.5 hari | ◻ | FR-009 |
| 8 | Implementasi Auth Flow | 1 hari | ◻ | FR-003, FR-101 |
| 9 | Implementasi RBAC Middleware | 1 hari | ◻ | FR-008 |
| 10 | Implementasi Audit Trail Middleware | 1 hari | ◻ | FR-008 |
| 11 | Setup Shared UI Components | 1 hari | ◻ | FR-002 |

**Deliverables Sprint 1:**
- ✅ Monorepo siap dengan 3 apps + 3 packages
- ✅ Backend API server berjalan
- ✅ Keycloak SSO terkonfigurasi
- ✅ Database ter-migrasi dengan 29 pegawai ter-seed
- ✅ Auth flow (login → Keycloak → session) berfungsi
- ✅ 10 shared UI components siap pakai

---

### SPRINT 2: Master Data (CRUD Lengkap) (6-8 hari)
**Dependency:** Sprint 1

| # | Task | Estimasi | Status | FR |
|---|------|----------|--------|-----|
| 12 | Generic CRUD Backend Master Data | 2 hari | ◻ | FR-201 - FR-212 |
| 13 | Master Data Frontend (10 entitas CRUD) | 4-6 hari | ◻ | FR-201 - FR-212 |

**Deliverables Sprint 2:**
- ✅ Generic CRUD handler untuk 10 entitas
- ✅ 60 API endpoints (6 per entity × 10 entities)
- ✅ 10 halaman frontend CRUD
- ✅ Dropdown API untuk semua entitas
- ✅ Data referensi ter-seed lengkap

**10 Entitas Master Data:**
1. Satuan Kerja (Satker)
2. Jabatan
3. Golongan/Pangkat
4. Unit Kerja (Tree View)
5. Eselon
6. Pendidikan
7. Agama
8. Status Kawin
9. Hukuman Disiplin
10. Jenis Diklat

---

### SPRINT 3: Kepegawaian Dasar (8-10 hari)
**Dependency:** Sprint 2 (butuh dropdown dari Master Data)

| # | Task | Estimasi | Status | FR |
|---|------|----------|--------|-----|
| 14 | Kepegawaian CRUD Backend | 3-4 hari | ◻ | FR-301 - FR-310, FR-313 |
| 15 | Kepegawaian Frontend | 5-6 hari | ◻ | FR-301 - FR-310, FR-313 |

**Deliverables Sprint 3:**
- ✅ CRUD Pegawai lengkap (validasi NIP 18 digit)
- ✅ Multi-step wizard tambah pegawai (4 step)
- ✅ Detail pegawai dengan tab view (5 tabs)
- ✅ CRUD 4 jenis riwayat (Pangkat, Jabatan, Pendidikan, Keluarga)
- ✅ File upload (Foto, SK, Ijazah)
- ✅ Dashboard statistik dengan chart
- ✅ 29 pegawai ter-verifikasi di UI

**Halaman Kepegawaian:**
- `/` - Dashboard statistik
- `/pegawai` - Daftar pegawai (filter & search)
- `/pegawai/tambah` - Form wizard tambah pegawai
- `/pegawai/[nip]` - Detail pegawai (5 tabs)
- `/pegawai/[nip]/edit` - Edit biodata

---

### SPRINT 4: Portal + Cetak SK + Polish (6-8 hari)
**Dependency:** Sprint 1 (Auth), Sprint 3 (Data untuk widget)

| # | Task | Estimasi | Status | FR |
|---|------|----------|--------|-----|
| 16 | Portal Launcher & Dashboard | 2 hari | ◻ | FR-101, FR-102, FR-103, FR-104 |
| 17 | Portal Admin UI | 2-3 hari | ◻ | FR-105, FR-106, FR-107, FR-108 |
| 18 | Cetak SK & Template Management | 2 hari | ◻ | FR-311, FR-312 |

**Deliverables Sprint 4:**
- ✅ Login page → Keycloak → callback → session
- ✅ Dashboard launcher dengan tile app
- ✅ 4+ dashboard widgets informatif
- ✅ Admin User Management (Keycloak proxy)
- ✅ Admin Hak Akses (assign role & permission)
- ✅ Admin Kelola Roles (CRUD roles)
- ✅ Audit Log viewer (filter & search)
- ✅ Template management (upload .docx)
- ✅ Cetak SK (minimal 3 template → PDF)
- ✅ End-to-end testing & bug fixing

**Halaman Portal:**
- `/login` - Login page
- `/` - Dashboard launcher
- `/profil` - Profil Saya
- `/admin` - Admin dashboard
- `/admin/users` - User Management
- `/admin/hak-akses` - Assign Role
- `/admin/hak-akses/roles` - Kelola Roles
- `/admin/audit-log` - Audit Log Viewer

---

## PRIORITY MATRIX

### P0 (Critical Path - Must Have)
- ✅ Task 1-11 (Sprint 1) - Infrastruktur dasar
- ✅ Task 12 (Sprint 2) - Generic CRUD backend
- ✅ Task 14 (Sprint 3) - CRUD Pegawai backend
- ✅ Task 16 (Sprint 4) - Portal launcher & auth

### P1 (High Priority - Should Have)
- ✅ Task 13 (Sprint 2) - Master Data frontend
- ✅ Task 15 (Sprint 3) - Kepegawaian frontend
- ✅ Task 17 (Sprint 4) - Admin UI
- ✅ Task 18 (Sprint 4) - Cetak SK

### P2 (Nice to Have - Could Have)
- ⭕ Dashboard widgets (bisa dikurangi dari 4 ke 2)
- ⭕ Chart di dashboard kepegawaian (bisa pakai table saja)
- ⭕ Tree view Unit Kerja (bisa pakai table hierarchical)
- ⭕ Multi-step wizard (bisa single form panjang)

---

## TECHNICAL DEPENDENCIES

```
Sprint 1 (Infrastruktur)
    │
    ├──→ Sprint 2 (Master Data)
    │        │
    │        └──→ Sprint 3 (Kepegawaian)
    │                 │
    │                 └──→ (data untuk widget)
    │
    └────────────────→ Sprint 4 (Portal + Cetak SK)
```

**Critical Dependencies:**
1. Sprint 2 butuh Sprint 1 (backend API, auth, UI components)
2. Sprint 3 butuh Sprint 2 (dropdown API untuk referensi)
3. Sprint 4 butuh Sprint 1 (auth) + Sprint 3 (data pegawai untuk widget)

---

## RESOURCE REQUIREMENTS

### Developer Skills Needed:
- ✅ **Frontend (Next.js 14):** 1-2 developers
  - App Router, Server Components, Client Components
  - TypeScript, Tailwind CSS, shadcn/ui
  - TanStack Query, React Hook Form, Zod

- ✅ **Backend (Go Fiber):** 1 developer
  - Go 1.21+, Fiber v3, GORM
  - PostgreSQL 17, SQL migrations
  - OIDC/OAuth2 (Keycloak integration)

- ✅ **DevOps/Infra:** 0.5 developer
  - Docker, docker-compose
  - Keycloak configuration
  - Gotenberg PDF service

### Estimated Team Size:
- **Minimum:** 2 developers (full-stack + backend)
- **Optimal:** 3 developers (frontend, backend, full-stack)
- **With DevOps:** 4 developers

---

## RISK ASSESSMENT

| Risk | Probability | Impact | Severity | Mitigation |
|------|-------------|--------|----------|------------|
| Keycloak config rumit | High | High | 🔴 Critical | Test auth flow di awal Sprint 1, buat realm-export.json |
| Port conflict (3000-3003) | Medium | Medium | 🟡 Medium | Check port availability, siapkan alternative ports |
| Data seed tidak lengkap | Low | Medium | 🟡 Medium | Admin bisa lengkapi via UI, tidak blocking |
| Gotenberg template fail | Medium | High | 🔴 Critical | Test dengan template sederhana dulu, siapkan fallback |
| RBAC complexity | Medium | High | 🔴 Critical | Implementasi di Sprint 1 sebagai middleware, test thoroughly |
| Timeline overrun | Medium | Medium | 🟡 Medium | Fokus P0 dulu, P2 bisa di-cut jika perlu |

---

## SUCCESS METRICS

### Sprint Completion Rate
- Target: 100% task completion per sprint
- Warning threshold: < 80% completion

### Code Quality
- Zero P0/P1 bugs
- TypeScript errors: 0
- Console errors: 0
- Code review pass rate: 100%

### Performance
- Page load time: < 3s
- API response time: < 500ms
- Database query time: < 100ms

### User Acceptance
- 29 pegawai ter-verifikasi di UI
- Minimal 3 template SK generate PDF berhasil
- End-to-end flow testing pass 100%

---

## DELIVERABLES SUMMARY

### Apps (3)
1. **Portal** (port 3000) - Launcher + Dashboard + Admin
2. **Master Data** (port 3001) - CRUD 10 entitas referensi
3. **Kepegawaian** (port 3002) - CRUD pegawai + 4 riwayat

### Backend API
- **Go Fiber Server** (port 3003)
- **~65 API endpoints** total
- **Authentication:** Keycloak OIDC + Better Auth
- **Database:** PostgreSQL 17 (2 databases, 21 tables)

### Shared Packages (3)
1. **@sikerma/ui** - 10 shared components
2. **@sikerma/auth** - Auth utilities & hooks
3. **@sikerma/shared** - API client, types, utils

### Infrastructure
- **Keycloak** (port 8081) - SSO & Identity Provider
- **PostgreSQL** (port 5435) - Main database
- **Gotenberg** (port 3100) - PDF generation service

### Documentation
- Setup guide
- API documentation
- User manual
- Acceptance criteria (this file)

---

## NEXT STEPS

### Before Starting:
1. ✅ Review PRD Fase 1 (`docs/PRD_FASE_1.md`)
2. ✅ Review technical architecture (`docs/blueprint_arch.md`)
3. ✅ Review app overview (`docs/overview_aplikasi.md`)
4. ✅ Setup development environment (Node.js, Go, Docker)
5. ✅ Clone repository & install dependencies

### Week 1 (Start Sprint 1):
1. ✅ Run `pnpm install` di root
2. ✅ Setup Turborepo workspace
3. ✅ Create 3 shared packages
4. ✅ Initialize 3 Next.js apps
5. ✅ Setup Go Fiber backend structure
6. ✅ Configure Keycloak realm

### Daily Standup Checklist:
- [ ] Task kemarin: DONE atau BLOCKED?
- [ ] Task hari ini: Apa yang akan dikerjakan?
- [ ] Blockers: Ada yang menghambat?
- [ ] Code review: Ada PR yang perlu direview?

---

## CONTACT & SUPPORT

**Technical Lead:** Muhardiansyah, S.Kom. (IT - Subbag PTIP)
**Product Owner:** Indra Yanita Yuliana, S.E., M.Si. (Sekretaris)
**Primary User:** Najwa Hijriana, S.E. (Admin Kepegawaian)

**Documentation:**
- PRD Fase 1: `docs/PRD_FASE_1.md`
- Task Details: `docs/PRD_FASE_1/TASK_*.md`
- Sprint Plans: `docs/PRD_FASE_1/SPRINT_*.md`
- Acceptance Criteria: `docs/PRD_FASE_1/ACCEPTANCE_CRITERIA.md`

---

## VERSION HISTORY

| Version | Date | Changes | Author |
|---------|------|---------|--------|
| 1.0 | 22 Feb 2026 | Initial creation | Muhardiansyah |
| | | - Breakdown 18 tasks | |
| | | - 4 sprint planning | |
| | | - Acceptance criteria | |

---

**Status Fase 1:** 🟡 Not Started
**Target Completion:** 28-36 hari kerja dari start date
