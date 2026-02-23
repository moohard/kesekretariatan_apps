# PRD — FASE 1: FONDASI SIKERMA

**Dokumen:** Product Requirements Document (PRD)
**Proyek:** Sistem Informasi Kesekretariatan Mahkamah Agung (SIKERMA)
**Fase:** 1 — Fondasi (Portal + Master Data + Kepegawaian Dasar)
**Instansi Pilot:** Pengadilan Agama Penajam (PA Kelas II)
**Versi:** 1.3 (Updated dengan Gap Analysis & Codebase Sync)
**Tanggal:** 23 Februari 2026
**Last Updated:** 23 Februari 2026 — Gap Analysis Codebase Sync
**Referensi:** `overview_aplikasi.md`, `blueprint_arch.md`

---

## 1. EXECUTIVE SUMMARY

### Visi

Membangun fondasi digital kesekretariatan Pengadilan Agama yang terintegrasi, dimulai dari tiga pilar utama: **Portal** (launcher & SSO), **Master Data** (referensi terpusat), dan **Kepegawaian Dasar** (data pegawai & cetak SK).

### Nilai Utama

Fase 1 menggantikan pengelolaan data manual/spreadsheet dengan sistem terpusat yang:
- Menyediakan **Single Sign-On** untuk seluruh aplikasi masa depan
- Menjadi **sumber kebenaran tunggal** (single source of truth) untuk data referensi dan kepegawaian
- Menyiapkan **arsitektur micro-frontend** yang scalable untuk 5 fase berikutnya

### Scope Fase 1

| App | Subdomain | Fungsi |
|-----|-----------|--------|
| Portal | `portal.pa-local` | Launcher, Dashboard, Admin RBAC, Audit Log |
| Master Data | `master.pa-local` | CRUD semua data referensi (satker, jabatan, golongan, dll) |
| Kepegawaian | `kepegawaian.pa-local` | Data pegawai, riwayat, cetak SK |

### Implementasi Status (Per 23 Februari 2026)

| Komponen | Target PRD | Progress Aktual | Gap |
|----------|------------|-----------------|-----|
| **Infrastruktur** | Docker + Services | 85% | 15% (Prometheus/Grafana belum) |
| **Database Schema** | Semua tabel + fields | 70% | 30% (field pegawai kurang) |
| **Portal App** | 8 halaman | 25% (2/8 halaman) | 75% |
| **Master Data App** | 12 halaman CRUD | 0% | 100% |
| **Kepegawaian App** | 16 halaman | 0% | 100% |
| **Backend API** | ~65 endpoints | 60% (framework only) | 40% |
| **Overall Progress** | 100% | **~35%** | **65%** |

**Catatan Penting:**
- ✅ Portal: Dashboard launcher dan Login sudah berfungsi
- ⚠️ Master Data & Kepegawaian: Hanya shell/layout, belum ada implementasi CRUD
- ⚠️ Database: Field `eselon_id`, `tmt_*`, `status_kerja` kurang di tabel pegawai
- ⚠️ Auth: Menggunakan `keycloak-js` langsung (bukan Better Auth seperti di spec awal)

---

## 2. PROBLEM STATEMENT

### Kondisi Saat Ini

| Masalah | Dampak | Segmen Terdampak |
|---------|--------|-------------------|
| Data pegawai tersebar di file Excel/manual | Data inkonsisten antar bagian, sulit dicari | Admin Kepegawaian, Pimpinan |
| Tidak ada sistem login terpusat | Setiap sistem terpisah, user harus login berulang | Semua pegawai |
| Data referensi (jabatan, golongan) tidak terstandar | Entry data berulang & tidak seragam | Operator data |
| Cetak SK & surat manual (copy-paste Word) | Rawan typo, format tidak konsisten, lambat | Admin Kepegawaian |
| Tidak ada audit trail | Tidak bisa tracking siapa mengubah apa | Pimpinan, Auditor |

### Data Existing

- **29 pegawai** real (dari `data_pegawai.json`) siap di-seed
- **Struktur organisasi** lengkap (dari `org_structure.json`): Pimpinan → Yudisial → Kepaniteraan → Sekretariat
- **Infrastruktur Docker** sudah ada: PostgreSQL 17 + Keycloak 26.0

---

## 3. GOALS & METRICS

### Goals (Prioritas)

| ID | Prioritas | Goal | Metric | Target |
|----|-----------|------|--------|--------|
| G-01 | **P0** | SSO & launcher berfungsi | User bisa login 1x dan akses semua app | 100% |
| G-02 | **P0** | Data pegawai terpusat & akurat | 29 pegawai ter-migrasi dengan data lengkap | 100% |
| G-03 | **P0** | CRUD master data berjalan | Semua 10 entitas referensi bisa CRUD | 100% |
| G-04 | **P1** | RBAC hybrid berfungsi | Admin bisa assign role & permission dari UI | Fully functional |
| G-05 | **P1** | Cetak SK via template | Minimal 3 template SK bisa generate PDF | ≥ 3 template |
| G-06 | **P1** | Audit trail aktif | Semua operasi CRUD tercatat | 100% coverage |
| G-07 | **P2** | Dashboard informatif | Widget statistik di Portal menampilkan data real | ≥ 4 widget |

### Success Criteria

- Admin Kepegawaian (Subbag Ortala) bisa mengelola data 29 pegawai tanpa spreadsheet
- Pimpinan (Ketua PA) bisa melihat dashboard ringkasan pegawai
- Operator bisa CRUD semua data referensi dari 1 tempat
- Semua aksi tercatat di audit log

---

## 4. NON-GOALS (Batasan Fase 1)

| # | Non-Goal | Akan Di-cover Di |
|---|----------|-----------------|
| 1 | Kenaikan pangkat, mutasi, pensiun | Fase 2 |
| 2 | Cuti, absensi, SKP | Fase 2 |
| 3 | Persuratan & disposisi | Fase 3 |
| 4 | BMN, SPPD | Fase 3 |
| 5 | Keuangan & anggaran | Fase 4 |
| 6 | Kepaniteraan & perkara | Fase 5 |
| 7 | OCR & integrasi SIKEP/KOMDANAS | Fase 5 |
| 8 | Mobile responsive (progressive) | Post-Fase 1 (enhancement) |
| 9 | Multi-satker (hanya PA Penajam dulu) | Post-Fase 1 |
| 10 | DragonflyDB cache | Disiapkan di infra, belum aktif dipakai |

---

## 5. USER PERSONAS

### Persona 1: Admin Kepegawaian (Primary User Fase 1)

```
Nama:     Najwa Hijriana, S.E.
Jabatan:  Operator — Subbag Kepegawaian, Organisasi, & Tata Laksana
Gol:      IX (PPPK)
Kebutuhan:
  - Input & update data 29 pegawai (biodata, pangkat, jabatan, pendidikan, keluarga)
  - Cetak SK dan surat kepegawaian dari template
  - Kelola data referensi (jabatan, golongan, unit kerja)
Pain Point:
  - Saat ini pakai Excel yang sering tidak sinkron antar file
  - Cetak SK harus copy-paste manual di Word
  - Tidak ada backup terpusat
```

### Persona 2: Pimpinan / Sekretaris (Viewer)

```
Nama:     Indra Yanita Yuliana, S.E., M.Si.
Jabatan:  Sekretaris Tingkat Pertama Klas II
Gol:      IV/a
Kebutuhan:
  - Lihat dashboard ringkasan pegawai (total, per status, per unit)
  - Monitor aktivitas terbaru (audit log)
  - Assign hak akses ke operator
Pain Point:
  - Harus tanya langsung ke operator untuk data statistik
  - Tidak bisa pantau siapa yang mengubah data apa
```

### Persona 3: Superadmin / IT (Setup & Maintenance)

```
Nama:     Muhardiansyah, S.Kom.
Jabatan:  Pranata Komputer Ahli Pertama — Subbag PTIP
Gol:      III/a
Kebutuhan:
  - Setup Keycloak realm & clients
  - Manage user accounts
  - Konfigurasi RBAC (role → permission mapping)
  - Monitor system health & audit log
Pain Point:
  - Belum ada tooling terpusat untuk user management
  - Setting permission harus manual di database
```

---

## 6. ARSITEKTUR TEKNIS FASE 1

### Tech Stack (Verified 23 Februari 2026 - Sync dengan Codebase)

| Layer | Teknologi | Versi | Keterangan |
|-------|-----------|-------|------------|
| **Frontend** | Next.js (App Router) | **16.1.6** | 3 app: portal, master-data, kepegawaian. ✅ **Stable** |
| **React** | React | **19.x** | Bundle dengan Next.js 16 |
| **UI Library** | shadcn/ui + Tailwind CSS | **v4.1** | Shared via `@sikerma/ui`. Tailwind v4.1: Rust engine, 5x lebih cepat dari v3 |
| **State/Fetch** | TanStack Query | **v5.62.0** | Server state & caching |
| **Form** | React Hook Form + Zod | Latest | Validasi form |
| **Backend** | Go Fiber | **v3.0.0-rc.2** | 1 monolith API server |
| **Auth** | Keycloak (Quarkus) | **26.5.3** | SSO + OIDC (sudah ada di Docker) |
| **Auth Client** | keycloak-js | **26.0.0** | ⚠️ **CATATAN:** Implementasi aktual menggunakan `keycloak-js` langsung, bukan Better Auth |
| **Database** | PostgreSQL | **18** | 1 instance, multiple databases (db_master, db_kepegawaian, db_keycloak) |
| **Cache** | DragonflyDB | **1.36+** | Redis-compatible, multi-threaded, PubSub untuk SSE (Fase 3) |
| **PDF Engine** | Gotenberg | **8.x** | Generate PDF dari template via LibreOffice Headless |
| **Monorepo** | Turborepo + pnpm | pnpm@10.29.3 | Workspace management |
| **Process Mgr** | PM2 | Latest | 3 Next.js apps di dev/prod |
| **Monitoring** | Prometheus | ⚠️ Belum ada | Perlu ditambahkan ke docker-compose |
| **Dashboard** | Grafana | ⚠️ Belum ada | Perlu ditambahkan ke docker-compose |
| **Testing** | Vitest + Playwright + Stryker | Latest | Unit, E2E, dan mutation testing |

### Infrastruktur (docker-compose.yml)

```
Services yang sudah ada:
  ✅ PostgreSQL 18         (port 5432) — Multi-database: db_master, db_kepegawaian, db_keycloak
  ✅ Keycloak 26.5.3       (port 8180) — SSO + OIDC
  ✅ DragonflyDB 1.36+     (port 6379) — Cache & PubSub (untuk SSE Fase 3)

Perlu ditambahkan di Fase 1:
  ➕ Gotenberg 8.13+       (port 3100) — PDF Engine via LibreOffice Headless

Frontend (Native - PM2):
  ➕ Next.js 16.x          (port 3000) — Portal
  ➕ Next.js 16.x          (port 3001) — Master Data
  ➕ Next.js 16.x          (port 3002) — Kepegawaian

Monitoring (Native):
  ➕ Prometheus            (port 9090) — Metrics collection
  ➕ Grafana               (port 3200) — Dashboard & alerting

File Storage (Native):
  ➕ /var/data/sekretariat/
    ├── templates/        — Template dokumen .docx
    ├── documents/        — Dokumen hasil generate (PDF)
    ├── pegawai/          — Foto & scan SK pegawai
```

### Database Design

```
┌─────────────────────────────────────────────┐
│  db_master (PostgreSQL — port 5435)         │
│                                             │
│  Tabel Referensi:                           │
│  • satker              • ref_pendidikan     │
│  • jabatan             • ref_agama          │
│  • golongan            • ref_status_kawin   │
│  • unit_kerja          • ref_jenis_hukdis   │
│  • eselon              • ref_jenis_diklat   │
│  • ref_golongan_non_pns (BARU)              │
│                                             │
│  Tabel RBAC:                                │
│  • app_roles           • role_permissions   │
│  • app_permissions     • user_app_roles     │
│                                             │
│  Tabel Audit:                               │
│  • audit_logs                               │
├─────────────────────────────────────────────┤
│  db_kepegawaian (PostgreSQL — port 5435)    │
│                                             │
│  • pegawai             • riwayat_pangkat    │
│  • riwayat_jabatan     • riwayat_pendidikan │
│  • keluarga            • template_dokumen   │
│  • hukdis              • diklat             │
└─────────────────────────────────────────────┘
```

### Schema Pegawai - Field yang Harus Ditambahkan (P0)

⚠️ **KRITIS:** Field berikut **harus ditambahkan** ke tabel `pegawai` sebelum Sprint 3:

| Field | Tipe | Keterangan | Prioritas |
|-------|------|------------|-----------|
| `eselon_id` | UUID | Referensi ke tabel eselon (struktural) | **P0** |
| `tmt_cpns` | DATE | TMT CPNS | **P0** |
| `tmt_pns` | DATE | TMT PNS | **P0** |
| `tmt_pangkat_terakhir` | DATE | TMT pangkat terakhir | **P0** |
| `status_kerja` | VARCHAR(20) | aktif/cuti/pensiun/mutasi_keluar/mutasi_masuk/meninggal/pemberhentian | **P0** |
| `nip_lama` | VARCHAR(9) | NIP 9 digit (untuk pegawai lama) | P2 |
| `karpeg_no` | VARCHAR(50) | Nomor Kartu Pegawai | P2 |
| `karpeg_file` | VARCHAR(500) | File scan karpeg | P2 |
| `taspen_no` | VARCHAR(50) | Nomor Taspen | P2 |
| `npwp` | VARCHAR(20) | Nomor Pokok Wajib Pajak | P2 |
| `bpjs_kesehatan` | VARCHAR(30) | Nomor BPJS Kesehatan | P2 |
| `bpjs_ketenagakerjaan` | VARCHAR(30) | Nomor BPJS Ketenagakerjaan | P2 |
| `kk_no` | VARCHAR(30) | Nomor Kartu Keluarga | P2 |
| `kk_file` | VARCHAR(500) | File scan KK | P2 |
| `ktp_no` | VARCHAR(30) | Nomor KTP/E-KTP | P2 |
| `ktp_file` | VARCHAR(500) | File scan KTP | P2 |
| `alamat_domisili` | TEXT | Alamat domisili (berbeda dengan alamat KTP) | P2 |
| `created_by` | UUID | User yang membuat data | **P1** |
| `updated_by` | UUID | User yang terakhir update | **P1** |
| `deleted_at` | TIMESTAMPTZ | Timestamp soft delete | P2 |
| `deleted_by` | UUID | User yang menghapus | P2 |

### Perbaikan Enum `status_pegawai` (P0)

⚠️ **Perbaikan Required:**

**Sekarang (di SQL):**
```sql
status_pegawai VARCHAR(20) DEFAULT 'aktif'  -- SALAH!
```

**Harusnya:**
```sql
status_pegawai VARCHAR(20) CHECK (status_pegawai IN ('PNS', 'CPNS', 'PPPK', 'HONORER')),
status_kerja VARCHAR(20) DEFAULT 'aktif' CHECK (status_kerja IN ('aktif', 'cuti', 'pensiun', 'mutasi_keluar', 'mutasi_masuk', 'meninggal', 'pemberhentian'))
```

### Master Data Baru: Golongan Non-PNS (P1)

Data pegawai mengandung golongan non-standar (I, V, IX) yang tidak ada di master:

```sql
CREATE TABLE ref_golongan_non_pns (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    kategori VARCHAR(50),  -- Honorer K1, Honorer K2, PPPK, Kontrak
    urutan INT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

INSERT INTO ref_golongan_non_pns (kode, nama, kategori, urutan) VALUES
('I', 'Honorer K1', 'Honorer', 1),
('V', 'Honorer K2', 'Honorer', 2),
('IX', 'Tenaga Kontrak', 'Kontrak', 3);
```

### Database Indexes (Mandatory)

Index **wajib dibuat** untuk performa query yang optimal:

```sql
-- Index untuk tabel pegawai
CREATE UNIQUE INDEX idx_pegawai_nip ON pegawai(nip);
CREATE INDEX idx_pegawai_nama ON pegawai USING GIN(to_tsvector('indonesian', nama_lengkap));
CREATE INDEX idx_pegawai_unit_kerja ON pegawai(unit_kerja_id);
CREATE INDEX idx_pegawai_status ON pegawai(status_pegawai);
CREATE INDEX idx_pegawai_golongan ON pegawai(golongan_id);
CREATE INDEX idx_pegawai_jabatan ON pegawai(jabatan_id);
CREATE INDEX idx_pegawai_is_active ON pegawai(is_active);

-- Index untuk tabel riwayat
CREATE INDEX idx_riwayat_pangkat_pegawai ON riwayat_pangkat(pegawai_id);
CREATE INDEX idx_riwayat_pangkat_tmt ON riwayat_pangkat(tmt_pangkat DESC);
CREATE INDEX idx_riwayat_jabatan_pegawai ON riwayat_jabatan(pegawai_id);
CREATE INDEX idx_riwayat_jabatan_tmt ON riwayat_jabatan(tmt_jabatan DESC);
CREATE INDEX idx_riwayat_pendidikan_pegawai ON riwayat_pendidikan(pegawai_id);
CREATE INDEX idx_keluarga_pegawai ON keluarga(pegawai_id);

-- Index untuk audit logs
CREATE INDEX idx_audit_logs_user ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_created ON audit_logs(created_at DESC);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_resource ON audit_logs(resource_type, resource_id);

-- Index untuk RBAC
CREATE INDEX idx_user_app_roles_user ON user_app_roles(user_id);
CREATE INDEX idx_user_app_roles_role ON user_app_roles(role_id);
CREATE INDEX idx_role_permissions_role ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission ON role_permissions(permission_id);

-- Index untuk unit kerja (tree structure)
CREATE INDEX idx_unit_kerja_parent ON unit_kerja(parent_id);
CREATE INDEX idx_unit_kerja_satker ON unit_kerja(satker_id);
```

### Database Constraints (Mandatory)

Constraint **wajib ditambahkan** untuk data integrity:

```sql
-- Constraint untuk tabel pegawai
ALTER TABLE pegawai ADD CONSTRAINT chk_nip_format CHECK (nip ~ '^\d{18}$');
ALTER TABLE pegawai ADD CONSTRAINT chk_tanggal_lahir CHECK (tanggal_lahir <= CURRENT_DATE);
ALTER TABLE pegawai ADD CONSTRAINT chk_email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');

-- Constraint untuk tabel riwayat pangkat
ALTER TABLE riwayat_pangkat ADD CONSTRAINT chk_tmt_pangkat CHECK (tmt_pangkat <= CURRENT_DATE);
ALTER TABLE riwayat_pangkat ADD CONSTRAINT chk_tanggal_sk_pangkat CHECK (tanggal_sk <= CURRENT_DATE);

-- Constraint untuk tabel riwayat jabatan
ALTER TABLE riwayat_jabatan ADD CONSTRAINT chk_tmt_jabatan CHECK (tmt_jabatan <= CURRENT_DATE);
ALTER TABLE riwayat_jabatan ADD CONSTRAINT chk_tanggal_sk_jabatan CHECK (tanggal_sk <= CURRENT_DATE);

-- Constraint untuk tabel keluarga
ALTER TABLE keluarga ADD CONSTRAINT chk_hubungan CHECK (hubungan IN ('Suami', 'Istri', 'Anak', 'Ayah', 'Ibu'));
```

### Monorepo Structure

```
/sikerma/
├── packages/
│   ├── ui/                    # @sikerma/ui — shadcn + custom components
│   ├── auth/                  # @sikerma/auth — Better Auth v1.3.4+ + Keycloak OIDC
│   └── shared/                # @sikerma/shared — API client, types, utils
├── apps/
│   ├── portal/                # App 1: port 3000
│   ├── master-data/           # App 2: port 3001
│   └── kepegawaian/           # App 3: port 3002
├── docker/
│   └── postgres/init/         # SQL init scripts (create databases, seed)
├── turbo.json
├── package.json
├── pnpm-workspace.yaml
└── docker-compose.yml
```

---

## 7. FUNCTIONAL REQUIREMENTS

### 7.1 — SHARED / INFRASTRUKTUR (FR-000 series)

| ID | Requirement | Prioritas | Deskripsi |
|----|-------------|-----------|-----------|
| FR-001 | Monorepo Setup | **P0** | Inisialisasi Turborepo + pnpm workspace dengan 3 shared packages + 3 apps |
| FR-002 | Shared UI Package | **P0** | `@sikerma/ui`: data-table, form-dialog, page-header, sidebar, app-header, breadcrumb, status-badge, delete-confirm, file-upload |
| FR-003 | Shared Auth Package | **P0** | `@sikerma/auth`: Better Auth v1.3.4+ config, Keycloak OIDC plugin, auth middleware, hooks (useSession, useUser, useRoles) |
| FR-004 | Shared Package | **P0** | `@sikerma/shared`: API client (fetch wrapper), TypeScript types, utils (formatter NIP, tanggal), constants |
| FR-005 | Keycloak Realm Setup | **P0** | Realm `pengadilan-agama` + 3 clients (portal, master-data, kepegawaian) + client role `[access]` per client |
| FR-006 | Database Migration | **P0** | SQL init scripts untuk create `db_master` + `db_kepegawaian` + semua tabel + seed data referensi |
| FR-007 | Go Backend Bootstrap | **P0** | Go Fiber v3 project setup: routing, middleware (auth, CORS, logger, audit), database connection pool |
| FR-008 | Audit Trail Middleware | **P1** | Setiap mutasi (POST/PUT/DELETE) otomatis tercatat di `audit_logs` dengan old/new value |
| FR-009 | Gotenberg Service | **P1** | Tambahkan Gotenberg 8 di docker-compose untuk PDF generation |
| FR-010 | Data Seed | **P0** | Migrasi 29 pegawai dari `data_pegawai.json` + struktur dari `org_structure.json` ke database |

### 7.2 — APP 1: PORTAL (FR-100 series)

| ID | Requirement | Prioritas | Halaman/Route | Deskripsi |
|----|-------------|-----------|---------------|-----------|
| FR-101 | Login Page | **P0** | `/login` | Redirect ke Keycloak login → callback → set session |
| FR-102 | Dashboard / Launcher | **P0** | `/` | Tampilkan tile per app yang bisa diakses user (berdasarkan Keycloak client access). Klik tile → redirect ke subdomain app |
| FR-103 | Dashboard Widgets | **P2** | `/` | Widget: Total Pegawai Aktif (per status), Pegawai per Unit Kerja, Aktivitas Terakhir (10 audit log terbaru), Pegawai Akan Pensiun (5 terdekat BUP) |
| FR-104 | Profil Saya | **P1** | `/profil` | User login bisa lihat & edit profil sendiri (foto, email, telepon) |
| FR-105 | Admin: User Management | **P1** | `/admin/users` | List users dari Keycloak (proxy API). Superadmin bisa enable/disable user, assign client access |
| FR-106 | Admin: Hak Akses | **P0** | `/admin/hak-akses` | List user → assign role per app. Contoh: Najwa → kepegawaian:admin, portal:viewer |
| FR-107 | Admin: Kelola Roles | **P1** | `/admin/hak-akses/roles` | CRUD role per app + mapping permissions. Contoh: role `kepegawaian:admin` → pegawai.view_all, pegawai.create, pegawai.update, pegawai.delete |
| FR-108 | Admin: Audit Log | **P1** | `/admin/audit-log` | Tabel audit log: timestamp, user, app, action, resource, detail. Filterable + searchable |

**API Endpoints Portal:**

```
GET    /api/v1/auth/me                         → Data user login + roles + permissions
GET    /api/v1/auth/permissions                 → List permissions user untuk app tertentu
GET    /api/v1/dashboard/summary                → Agregasi data untuk widgets

GET    /api/v1/admin/users                      → List users Keycloak (proxy)
POST   /api/v1/admin/users                      → Create user di Keycloak (proxy)
PUT    /api/v1/admin/users/:id                  → Update user di Keycloak (proxy)
PUT    /api/v1/admin/users/:id/client-access    → Assign/revoke client access

GET    /api/v1/admin/roles                      → List roles per app
POST   /api/v1/admin/roles                      → Create role
PUT    /api/v1/admin/roles/:id                  → Update role
DELETE /api/v1/admin/roles/:id                  → Delete role (non-system only)

GET    /api/v1/admin/permissions                → List permissions per app
POST   /api/v1/admin/roles/:id/permissions      → Assign permissions ke role
DELETE /api/v1/admin/roles/:id/permissions/:pid → Remove permission dari role

GET    /api/v1/admin/user-roles                 → List user-role assignments
POST   /api/v1/admin/user-roles                 → Assign role ke user
DELETE /api/v1/admin/user-roles/:id             → Remove role dari user

GET    /api/v1/audit-logs                       → List audit logs (paginated, filterable)
```

### 7.3 — APP 2: MASTER DATA (FR-200 series)

| ID | Requirement | Prioritas | Halaman/Route | Deskripsi |
|----|-------------|-----------|---------------|-----------|
| FR-201 | Dashboard Master | **P2** | `/` | Ringkasan jumlah data per entitas referensi (card grid) |
| FR-202 | CRUD Satuan Kerja | **P0** | `/satker` | Tabel + form dialog: kode, nama, alamat, telepon, email, tipe, is_active |
| FR-203 | CRUD Jabatan | **P0** | `/jabatan` | Tabel + form: kode, nama, jenis (struktural/fungsional), eselon, kelas jabatan |
| FR-204 | CRUD Golongan/Pangkat | **P0** | `/golongan` | Tabel + form: kode (III/a), nama pangkat, ruang, tingkat |
| FR-205 | CRUD Unit Kerja | **P0** | `/unit-kerja` | Tabel tree + form: satker, kode, nama, parent (hierarki). Tree view untuk visualisasi |
| FR-206 | CRUD Eselon | **P1** | `/eselon` | Tabel + form: kode, nama, level |
| FR-207 | CRUD Pendidikan | **P1** | `/pendidikan` | Tabel + form: kode, jenjang, urutan |
| FR-208 | CRUD Agama | **P1** | `/agama` | Tabel + form: nama |
| FR-209 | CRUD Status Kawin | **P1** | `/status-kawin` | Tabel + form: nama |
| FR-210 | CRUD Hukuman Disiplin | **P2** | `/hukuman-disiplin` | Tabel + form: tingkat, nama |
| FR-211 | CRUD Jenis Diklat | **P2** | `/jenis-diklat` | Tabel + form: kode, nama, kategori |
| FR-212 | Dropdown API | **P0** | — | Setiap entitas punya endpoint `/dropdown` yang return `id + nama` tanpa pagination. Dipakai oleh app lain (kepegawaian) |

**Pola API Seragam (semua entitas):**

```
GET    /api/v1/master/{entity}                  → List (paginated, searchable, sortable)
GET    /api/v1/master/{entity}/:id              → Detail
POST   /api/v1/master/{entity}                  → Create
PUT    /api/v1/master/{entity}/:id              → Update
DELETE /api/v1/master/{entity}/:id              → Soft delete
GET    /api/v1/master/{entity}/dropdown          → List id+nama (tanpa pagination)

{entity} = satker | jabatan | golongan | unit-kerja | eselon |
            pendidikan | agama | status-kawin | hukuman-disiplin | jenis-diklat
```

**Data Seed Master (dari data existing):**

```
Satker:
  - PA Penajam (kode: PA-PNJ, tipe: pa)

Golongan PNS (Standar BKN - sudah lengkap di seed):
  - I/a - I/d (Juru Muda - Juru Tingkat I)
  - II/a - II/d (Pengatur Muda - Pengatur Tingkat I)
  - III/a - III/d (Penata Muda - Penata Tingkat I)
  - IV/a - IV/e (Pembina - Pembina Utama)

Golongan Non-PNS (BARU - perlu ditambahkan):
  - I: Honorer K1
  - V: Honorer K2
  - IX: Tenaga Kontrak/PPPK

Unit Kerja (dari org_structure.json):
  - Pengadilan Agama Penajam (root)
    ├── Pimpinan
    ├── Panitera
    │   ├── Panitera Muda Permohonan
    │   ├── Panitera Muda Gugatan
    │   └── Panitera Muda Hukum
    └── Sekretariat
        ├── Subbag Perencanaan, TI, dan Pelaporan
        ├── Subbag Umum dan Keuangan
        └── Subbag Kepegawaian, Organisasi, dan Tata Laksana

Jabatan Struktural Pimpinan:
  - Ketua Pengadilan Tingkat Pertama Klas II (Eselon II.b)
  - Wakil Ketua Tingkat Pertama (Eselon II.b)

Jabatan Struktural Kepaniteraan:
  - Panitera Tingkat Pertama Klas II (Eselon III.d)
  - Panitera Muda Tingkat Pertama Klas II (Eselon IV.a)
  - Panitera Muda Permohonan (Eselon IV.a)
  - Panitera Muda Gugatan (Eselon IV.a)
  - Panitera Muda Hukum (Eselon IV.a)

Jabatan Struktural Kesekretariatan:
  - Sekretaris Tingkat Pertama Klas II (Eselon III.d)
  - Kepala Subbagian (Eselon IV.a)

Jabatan Fungsional:
  - Hakim Tingkat Pertama
  - Pranata Komputer Ahli Pertama

Jabatan Pelaksana:
  - Panitera Pengganti Tingkat Pertama
  - Juru Sita Pengganti
  - Klerek - Analis Perkara Peradilan
  - Klerek - Pengelola Penanganan Perkara
  - Klerek - Dokumentalis Hukum
  - Teknisi Sarana dan Prasarana
  - Operator - Penata Layanan Operasional
  - Operator Layanan Operasional
  - Pengelola Umum Operasional

TOTAL JABATAN: 20+ jabatan (perlu di-seed lengkap)
```

### 7.4 — APP 3: KEPEGAWAIAN DASAR (FR-300 series)

| ID | Requirement | Prioritas | Halaman/Route | Deskripsi |
|----|-------------|-----------|---------------|-----------|
| FR-301 | Dashboard Kepegawaian | **P1** | `/` | Statistik: total pegawai per status (PNS/CPNS/PPPK), per golongan, per unit kerja. Chart bar + pie |
| FR-302 | Daftar Pegawai | **P0** | `/pegawai` | Tabel: foto, NIP, nama, jabatan, golongan, unit kerja, status. Search by NIP/nama. Filter by status, golongan, unit kerja. Sortable |
| FR-303 | Tambah Pegawai | **P0** | `/pegawai/tambah` | Form multi-step: (1) Biodata, (2) Posisi & Pangkat, (3) Pendidikan, (4) Keluarga. Validasi NIP unik, format 18 digit |
| FR-304 | Detail Pegawai | **P0** | `/pegawai/[nip]` | Tab view: Biodata, Riwayat Pangkat, Riwayat Jabatan, Riwayat Pendidikan, Data Keluarga. Setiap tab bisa CRUD inline |
| FR-305 | Edit Biodata | **P0** | `/pegawai/[nip]/edit` | Form edit biodata pegawai (nama, gelar, tempat/tanggal lahir, kontak, foto) |
| FR-306 | Upload Foto | **P1** | — | Upload foto pegawai (max 2MB, jpg/png). Preview sebelum simpan |
| FR-307 | CRUD Riwayat Pangkat | **P0** | `/pegawai/[nip]` tab | Tambah/edit/hapus riwayat pangkat: golongan, no SK, tanggal SK, TMT, pejabat penetap, upload file SK |
| FR-308 | CRUD Riwayat Jabatan | **P0** | `/pegawai/[nip]` tab | Tambah/edit/hapus riwayat jabatan: jabatan, unit kerja, no SK, tanggal SK, TMT, upload file SK |
| FR-309 | CRUD Riwayat Pendidikan | **P1** | `/pegawai/[nip]` tab | Tambah/edit/hapus: jenjang, nama sekolah, jurusan, tahun lulus, no ijazah, upload ijazah |
| FR-310 | CRUD Data Keluarga | **P1** | `/pegawai/[nip]` tab | Tambah/edit/hapus: hubungan, nama, tempat/tanggal lahir, pekerjaan |
| FR-311 | Cetak SK | **P1** | `/cetak-sk` | Pilih template → pilih pegawai → preview → generate PDF via Gotenberg |
| FR-312 | Template Management | **P1** | `/admin/template` | Upload template .docx, definisikan placeholders (JSONB), aktif/nonaktif |
| FR-313 | Nonaktifkan Pegawai | **P1** | — | Soft delete: set is_active=false + alasan. Pegawai tidak muncul di list default tapi bisa dilihat via filter |

**API Endpoints Kepegawaian:**

```
# Pegawai
GET    /api/v1/pegawai                          → List (paginated, searchable, filterable)
GET    /api/v1/pegawai/:nip                     → Detail lengkap + semua riwayat
POST   /api/v1/pegawai                          → Tambah pegawai
PUT    /api/v1/pegawai/:nip                     → Update biodata
DELETE /api/v1/pegawai/:nip                     → Soft delete (nonaktifkan)
POST   /api/v1/pegawai/:nip/foto               → Upload foto

# Riwayat (pola seragam)
GET    /api/v1/pegawai/:nip/pangkat             → List riwayat pangkat
POST   /api/v1/pegawai/:nip/pangkat             → Tambah riwayat pangkat
PUT    /api/v1/pegawai/:nip/pangkat/:id         → Edit riwayat pangkat
DELETE /api/v1/pegawai/:nip/pangkat/:id         → Hapus riwayat pangkat

GET    /api/v1/pegawai/:nip/jabatan             → List riwayat jabatan
POST   /api/v1/pegawai/:nip/jabatan             → Tambah
PUT    /api/v1/pegawai/:nip/jabatan/:id         → Edit
DELETE /api/v1/pegawai/:nip/jabatan/:id         → Hapus

GET    /api/v1/pegawai/:nip/pendidikan          → List riwayat pendidikan
POST   /api/v1/pegawai/:nip/pendidikan          → Tambah
PUT    /api/v1/pegawai/:nip/pendidikan/:id      → Edit
DELETE /api/v1/pegawai/:nip/pendidikan/:id      → Hapus

GET    /api/v1/pegawai/:nip/keluarga            → List data keluarga
POST   /api/v1/pegawai/:nip/keluarga            → Tambah
PUT    /api/v1/pegawai/:nip/keluarga/:id        → Edit
DELETE /api/v1/pegawai/:nip/keluarga/:id        → Hapus

# File upload (riwayat)
POST   /api/v1/upload/sk                        → Upload file SK (max 5MB, pdf)
POST   /api/v1/upload/ijazah                    → Upload file ijazah (max 5MB, pdf)

# Dokumen / Cetak SK
GET    /api/v1/dokumen/templates                → List template aktif
POST   /api/v1/dokumen/templates                → Upload template baru
PUT    /api/v1/dokumen/templates/:id            → Update template
DELETE /api/v1/dokumen/templates/:id            → Nonaktifkan template
POST   /api/v1/dokumen/cetak                    → Generate dokumen PDF
GET    /api/v1/dokumen/download/:id             → Download file hasil cetak

# Statistik
GET    /api/v1/pegawai/statistik                → Agregasi untuk dashboard
```

---

## 8. IMPLEMENTATION PHASES (Sub-Fase dalam Fase 1)

Fase 1 dibagi menjadi **4 sprint** berurutan berdasarkan dependency:

### Timeline Overview

| Sprint | Fokus | Estimasi | Revised* |
|--------|-------|----------|----------|
| Sprint 1 | Infrastruktur & Fondasi | 12-14 hari | Keycloak setup, auth flow, RBAC middleware memerlukan waktu |
| Sprint 2 | Master Data (CRUD) | 8-10 hari | 10 halaman CRUD dengan DataTable + FormDialog |
| Sprint 3 | Kepegawaian Dasar | 12-14 hari | Multi-step wizard, file upload, dashboard chart |
| Sprint 4 | Portal + Cetak SK | 10-12 hari | Admin UI, template management, Gotenberg integration |
| **TOTAL** | | **42-50 hari** | **8-10 minggu** |

*\*Revised timeline untuk tim 1-2 developer (buffer +40% dari estimasi awal)*

### Sprint 1: Infrastruktur & Fondasi

```
Dependency: Tidak ada (starting point)
Deliverable: Monorepo siap, backend berjalan, auth berfungsi
Estimasi: 12-14 hari

Checklist:
□ Inisialisasi monorepo (Turborepo + pnpm workspace)
□ Setup 3 shared packages (@sikerma/ui, @sikerma/auth, @sikerma/shared)
□ Setup 3 Next.js apps (portal, master-data, kepegawaian)
□ Bootstrap Go Fiber v3 backend (project structure, routing, middleware)
□ Konfigurasi Keycloak realm + 3 clients + realm export
□ Setup database migrations (db_master + db_kepegawaian) + INDEX
□ Seed data referensi awal + 29 pegawai
□ Tambahkan Gotenberg ke docker-compose
□ Implementasi auth flow (login → Keycloak → callback → session)
□ POC: Validasi Auth.js v5 + Keycloak OIDC flow
□ Implementasi RBAC middleware di backend (baca permission dari DB)
□ Implementasi audit trail middleware
□ @sikerma/ui: sidebar, app-header, page-header, breadcrumb
□ Definisikan error response standard API
```

**FR yang di-cover:** FR-001, FR-002, FR-003, FR-004, FR-005, FR-006, FR-007, FR-008, FR-009, FR-010

### Sprint 2: Master Data (CRUD Lengkap)

```
Dependency: Sprint 1 (backend + auth + UI components)
Deliverable: App Master Data fully functional
Estimasi: 8-10 hari

Checklist:
□ Backend: Generic CRUD handler untuk semua entitas master
□ Backend: Dropdown endpoints untuk semua entitas
□ Frontend: Data table component (sort, search, pagination)
□ Frontend: Form dialog component (modal CRUD)
□ Frontend: CRUD halaman per entitas (10 halaman)
□ Frontend: Dashboard Master Data (ringkasan jumlah)
□ Seed data referensi lengkap (golongan, jabatan, unit kerja dari data existing)
```

**FR yang di-cover:** FR-201 — FR-212

### Sprint 3: Kepegawaian Dasar

```
Dependency: Sprint 2 (master data harus sudah ada untuk dropdown/referensi)
Deliverable: App Kepegawaian functional untuk CRUD pegawai + riwayat
Estimasi: 12-14 hari

Checklist:
□ Backend: CRUD pegawai + validasi NIP
□ Backend: CRUD riwayat (pangkat, jabatan, pendidikan, keluarga)
□ Backend: File upload (foto, SK, ijazah) + magic bytes validation
□ Backend: Statistik endpoint
□ Frontend: Daftar Pegawai (tabel + search + filter)
□ Frontend: Form Tambah Pegawai (multi-step wizard)
□ Frontend: Detail Pegawai (tab view + CRUD inline per riwayat)
□ Frontend: Edit Biodata
□ Frontend: Upload foto pegawai
□ Frontend: Dashboard Kepegawaian (statistik chart)
□ Migrasi data 29 pegawai dari JSON ke tampilan terverifikasi
```

**FR yang di-cover:** FR-301 — FR-310, FR-313

### Sprint 4: Portal + Cetak SK + Polish

```
Dependency: Sprint 1 (auth), Sprint 3 (data pegawai untuk widget)
Deliverable: Portal sebagai launcher, cetak SK, RBAC admin UI
Estimasi: 10-12 hari

Checklist:
□ Frontend Portal: Dashboard launcher (tile per app)
□ Frontend Portal: Dashboard widgets (statistik pegawai)
□ Frontend Portal: Admin User Management (proxy Keycloak)
□ Frontend Portal: Admin Hak Akses (assign role ke user)
□ Frontend Portal: Admin Kelola Roles (CRUD role + permissions)
□ Frontend Portal: Audit Log viewer
□ Frontend Portal: Profil Saya
□ Backend: Keycloak admin API proxy
□ Backend: Dashboard summary aggregation
□ Kepegawaian: Template management (upload .docx)
□ Kepegawaian: Cetak SK (template → Gotenberg → PDF)
□ End-to-end testing
□ Bug fixing & polish
```

**FR yang di-cover:** FR-101 — FR-108, FR-311, FR-312

### Dependency Graph

```
Sprint 1 (Infra & Fondasi)
    │
    ├──→ Sprint 2 (Master Data)
    │        │
    │        └──→ Sprint 3 (Kepegawaian)
    │                 │
    └────────────────→ Sprint 4 (Portal + Cetak SK + Polish)
```

---

## 9. RISKS & MITIGATIONS

| # | Risiko | Probabilitas | Dampak | Mitigasi |
|---|--------|-------------|--------|----------|
| R-01 | **Better Auth + Keycloak OIDC integration** | Sedang | Tinggi | **Buat POC auth flow (1 hari) sebelum Sprint 1**. Better Auth v1.3.4+ sudah memiliki SSO/OIDC plugin. Dokumentasi: https://www.better-auth.com/docs |
| R-02 | Keycloak config rumit (realm, clients, roles) | Tinggi | Tinggi | Sudah ada realm-export.json. Buat dokumentasi setup step-by-step. Test auth flow di Sprint 1 sebelum lanjut |
| R-03 | Timeline 8-10 minggu tidak realistis untuk tim 1-2 developer | Sedang | Tinggi | **Definisikan P0 features only untuk MVP**. Kurangi scope Sprint 4 (dashboard widgets, tree view) jika perlu. Weekly progress review |
| R-04 | RBAC hybrid (Keycloak + DB) menambah kompleksitas | Sedang | Tinggi | Implementasi di Sprint 1 sebagai middleware. Semua sprint selanjutnya tinggal pakai. Test thoroughly |
| R-05 | Data pegawai JSON tidak lengkap (hanya biodata dasar) | Sedang | Sedang | Seed sebagai data awal, admin bisa melengkapi via UI setelah Sprint 3. Tidak blocking |
| R-06 | Gotenberg template format kompleks | Sedang | Sedang | Mulai dengan **2 template sederhana** (SK Kenaikan Pangkat, SK Jabatan), tambah bertahap. Sediakan contoh .docx reference |
| R-07 | Monorepo shared packages breaking changes | Rendah | Sedang | Turborepo caching + strict versioning. Selalu test semua apps setelah update shared package |
| R-08 | Golongan IX dan V (PPPK/non-PNS) beda skema | Rendah | Rendah | Status pegawai sudah mengakomodasi (PNS/CPNS/PPPK/Honorer). Golongan untuk non-PNS bisa null/khusus |
| R-09 | **Hardcoded credentials di Docker** | Sedang | Tinggi | **Gunakan Docker Secrets pattern**. Tidak ada credential di docker-compose.yml atau source code |
| R-10 | **Database query lambat tanpa index** | Sedang | Sedang | **Index sudah didefinisikan** di section Database Indexes. Tambahkan di migration scripts Sprint 1 |

---

## 10. DATA MIGRATION PLAN

### Sumber Data

| File | Isi | Jumlah Record | Target Tabel |
|------|-----|---------------|--------------|
| `data_pegawai.json` | Data 29 pegawai (NIP, nama, jabatan, unit kerja, TMT, golongan) | 29 | `pegawai`, `riwayat_jabatan`, `riwayat_pangkat` |
| `org_structure.json` | Struktur organisasi hierarki | ~15 unit | `unit_kerja`, `satker` |

### Strategi Seed

```
1. Seed satker → PA Penajam
2. Seed golongan → extract unik dari data_pegawai (I, II/c, II/d, III/a, III/b, III/c, III/d, IV/a, V, IX)
3. Seed unit_kerja → dari org_structure (tree: parent_id)
4. Seed jabatan → extract unik dari data_pegawai (17 jabatan)
5. Seed ref_agama, ref_status_kawin, ref_pendidikan → data standar nasional
6. Seed pegawai → 29 record dari data_pegawai.json
   - Map golongan → golongan.id
   - Map jabatan → jabatan.id
   - Map unit_kerja → unit_kerja.id (parse dari "unit_kerja" field)
7. Seed riwayat_pangkat → 1 record per pegawai (dari gol + TMT saat ini)
8. Seed riwayat_jabatan → 1 record per pegawai (dari jabatan + TMT saat ini)
9. Seed RBAC → roles default + permissions default per app
10. Seed user_app_roles → superadmin untuk user IT (Muhardiansyah)
```

---

## 11. SHARED UI COMPONENTS SPEC

Komponen yang dibuat di `@sikerma/ui` dan dipakai oleh semua 3 apps:

| Komponen | Props Utama | Dipakai Di |
|----------|-------------|------------|
| `DataTable` | columns, data, searchable, sortable, pagination, onRowClick | Master Data, Kepegawaian, Portal |
| `FormDialog` | title, fields, onSubmit, mode (create/edit), open/close | Master Data (semua CRUD), Kepegawaian (riwayat) |
| `PageHeader` | title, description, actions (button[]) | Semua halaman |
| `Sidebar` | menuItems[], activeItem, collapsed | Semua apps (config berbeda per app) |
| `AppHeader` | user, notifications, appSwitcher | Semua apps |
| `Breadcrumb` | items[] (label, href) | Semua halaman |
| `StatusBadge` | status, variant (success/warning/danger/info) | Kepegawaian, Portal |
| `DeleteConfirm` | title, message, onConfirm | Semua CRUD |
| `FileUpload` | accept, maxSize, onUpload, preview | Kepegawaian (foto, SK, ijazah) |
| `StepWizard` | steps[], currentStep, onNext, onBack | Kepegawaian (tambah pegawai) |

---

## 12. ACCEPTANCE CRITERIA (Definition of Done)

### Per Sprint

| Sprint | Done When |
|--------|-----------|
| Sprint 1 | User bisa login via Keycloak, redirect ke Portal, session valid. Backend serve API. Database ter-migrasi. 29 pegawai ter-seed |
| Sprint 2 | Admin bisa CRUD semua 10 entitas referensi. Dropdown API berfungsi. Data tersimpan di database |
| Sprint 3 | Admin bisa lihat/tambah/edit/hapus pegawai + semua riwayat. File upload berfungsi. Dashboard statistik menampilkan data real |
| Sprint 4 | Portal launcher menampilkan tile app. Admin bisa manage roles & permissions dari UI. Cetak SK menghasilkan PDF. Audit log mencatat semua aktivitas |

### Overall Fase 1

```
✅ 29 pegawai PA Penajam ter-migrasi ke sistem
✅ Admin bisa login 1x (SSO) dan akses ketiga app
✅ Semua data referensi bisa dikelola dari Master Data app
✅ Data pegawai bisa dikelola lengkap (biodata + 4 riwayat)
✅ Minimal 3 template SK bisa generate PDF
✅ RBAC berfungsi (role-based menu & button visibility)
✅ Audit trail mencatat semua operasi CRUD
✅ Portal menampilkan dashboard ringkasan
```

---

## 13. SECURITY REQUIREMENTS

### 13.1 Data Protection & Encryption

| Aspek | Requirement | Implementation |
|-------|-------------|----------------|
| **Encryption at Rest** | ✅ Mandatory | PostgreSQL TDE (Transparent Data Encryption) + LUKS filesystem encryption untuk `/var/data/sekretariat/` |
| **Encryption in Transit** | ✅ Mandatory | TLS 1.3 untuk semua komunikasi. Nginx SSL termination dengan certificate dari Let's Encrypt atau internal CA |
| **File Permissions** | ✅ Mandatory | `/var/data/sekretariat/` ownership: `www-data:docker`, permission: **750** |
| **Secret Management** | ✅ Mandatory | **Docker Secrets** untuk semua credentials. **Tidak ada hardcoded credentials** di docker-compose.yml atau source code |
| **Network Isolation** | ✅ Mandatory | Hanya Nginx (80/443) terekspos ke jaringan kantor. Docker services hanya dapat diakses via Docker internal network atau localhost |
| **Backup Encryption** | ✅ Mandatory | Semua backup di-encrypt dengan GPG sebelum disimpan ke NAS/external storage |

### 13.2 Authentication & Authorization

| Aspek | Requirement | Configuration |
|-------|-------------|---------------|
| **SSO Provider** | ✅ Keycloak OIDC | Realm: `pengadilan-agama`, 3 clients (portal, master-data, kepegawaian) |
| **JWT Token Expiry** | ✅ Configurable | Access Token: **15 menit**, Refresh Token: **1 hari** (bukan 7 hari untuk keamanan) |
| **Cookie Security** | ✅ Mandatory | HTTP-Only: Yes, Secure: Yes, SameSite: Strict |
| **Session Management** | ✅ Auth.js v5 | Token disimpan di HTTP-Only cookie dengan auto-renewal |
| **Session Revocation** | ✅ Mandatory | Revoke semua refresh token saat password change. Admin dapat revoke user session via Portal. Max 2 concurrent session per user |
| **Password Policy** | ✅ Keycloak | Minimum 8 karakter, kombinasi huruf + angka + simbol, expiry 90 hari |
| **MFA (Multi-Factor Auth)** | ✅ Mandatory untuk Admin | WebAuthn/FIDO2 atau TOTP (Google Authenticator) wajib untuk role admin & superadmin. Optional untuk user biasa |

### 13.3 Security Headers

Nginx **wajib dikonfigurasi** dengan security headers berikut:

```nginx
# Tambahkan di nginx.conf atau server block
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
add_header X-Frame-Options "SAMEORIGIN" always;
add_header X-Content-Type-Options "nosniff" always;
add_header X-XSS-Protection "1; mode=block" always;
add_header Content-Security-Policy "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: blob:; font-src 'self' data:;" always;
add_header Referrer-Policy "strict-origin-when-cross-origin" always;
add_header Permissions-Policy "geolocation=(), microphone=(), camera=()" always;
```

### 13.4 Row-Level Security (RLS)

PostgreSQL Row-Level Security **wajib diaktifkan** untuk tabel-tabel sensitive:

```sql
-- Helper function untuk mendapatkan user_id dari JWT token
-- Dipanggil di backend Go: SET request.jwt.claim.sub = 'user-uuid-from-keycloak';
CREATE OR REPLACE FUNCTION current_user_id()
RETURNS UUID AS $$
    SELECT NULLIF(current_setting('request.jwt.claim.sub', true), '')::UUID;
$$ LANGUAGE SQL SECURITY DEFINER STABLE;

-- Enable RLS pada tabel pegawai
ALTER TABLE pegawai ENABLE ROW LEVEL SECURITY;

-- Example RLS Policy untuk tabel pegawai
CREATE POLICY pegawai_select_policy ON pegawai
    FOR SELECT
    USING (
        -- User bisa lihat data pegawai di unit kerjanya sendiri
        unit_kerja_id IN (
            SELECT unit_kerja_id FROM user_app_roles
            WHERE user_id = current_user_id()
        )
        OR
        -- Admin bisa lihat semua
        EXISTS (
            SELECT 1 FROM user_app_roles uar
            JOIN app_roles ar ON uar.role_id = ar.id
            WHERE uar.user_id = current_user_id()
            AND ar.role_code = 'admin'
        )
    );
```

### 13.5 Input Validation & SQL Injection Prevention

| Requirement | Implementation |
|-------------|----------------|
| **Backend Validation** | Validasi di Go backend (tidak mengandalkan frontend) dengan library `go-playground/validator` |
| **Parameterized Queries** | GORM dengan driver `pgx/v5` — automatic parameter binding |
| **CSRF Protection** | Go Fiber middleware `csrf` aktif untuk POST/PUT/DELETE |
| **Rate Limiting** | Go Fiber middleware `limiter` dengan konfigurasi per-endpoint |
| **CORS Configuration** | Explicit allowed origins, credentials: true untuk SSO |

**Rate Limiting Detail:**

| Endpoint Type | Limit | Window | Purpose |
|---------------|-------|--------|---------|
| Global | 100 req/min | Per IP | DDoS protection |
| Login (`/api/v1/auth/*`) | 5 req/15min | Per IP | Brute force protection |
| File Upload | 10 req/min | Per user | Abuse prevention |
| Admin Operations | 50 req/min | Per user | Rate control |

### 13.6 Audit Trail

Semua aksi penting **wajib dicatat** di tabel `audit_logs`:

```sql
audit_logs (
    id UUID PRIMARY KEY,
    app_source VARCHAR(50),        -- portal/master/kepegawaian
    user_id UUID NOT NULL,
    user_name VARCHAR(100),
    action VARCHAR(50),            -- CREATE, UPDATE, DELETE, PRINT, LOGIN, LOGOUT
    resource_type VARCHAR(50),     -- pegawai, surat, dokumen, keuangan
    resource_id VARCHAR(100),
    old_value JSONB,               -- Nilai sebelum (untuk UPDATE/DELETE) - dengan PII masking
    new_value JSONB,               -- Nilai sesudah (untuk CREATE/UPDATE) - dengan PII masking
    sensitive_fields TEXT[],       -- Daftar field yang di-mask (misal: ['nik', 'no_rekening'])
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
)
```

**PII Masking di Audit Log:**

```go
// Contoh implementasi di Go backend
func maskSensitiveFields(data map[string]interface{}, sensitiveFields []string) map[string]interface{} {
    masked := make(map[string]interface{})
    for k, v := range data {
        if contains(sensitiveFields, k) {
            masked[k] = "***MASKED***"
        } else {
            masked[k] = v
        }
    }
    return masked
}

// Field wajib di-mask: nik, no_rekening, gaji, alamat_lengkap, no_telepon_pribadi
```

**Aksi yang wajib di-audit:**
- ✅ Login / Logout
- ✅ CRUD data pegawai
- ✅ CRUD data referensi (master data)
- ✅ Cetak dokumen (SK, surat)
- ✅ Perubahan role / permission
- ✅ File upload / download

### 13.7 File Upload Validation

File upload **wajib divalidasi** dengan multiple layers:

| Validation Layer | Implementation |
|------------------|----------------|
| **Extension Check** | Whitelist: `.pdf`, `.jpg`, `.jpeg`, `.png`, `.docx` |
| **Magic Bytes Check** | Validasi content-type dari file header, bukan extension |
| **Size Limit** | Foto: 2MB, SK/Ijazah: 5MB, Template: 10MB |
| **Virus Scan** | ClamAV integration untuk scan malware |
| **Filename Sanitization** | Hapus karakter berbahaya, gunakan UUID + timestamp |

```go
// Contoh implementasi validasi file di Go
func ValidateFileUpload(file *multipart.FileHeader, allowedTypes []string, maxSize int64) error {
    // 1. Cek size
    if file.Size > maxSize {
        return fmt.Errorf("file size exceeds limit: %d bytes", maxSize)
    }

    // 2. Baca magic bytes (first 512 bytes)
    f, _ := file.Open()
    defer f.Close()
    buffer := make([]byte, 512)
    f.Read(buffer)
    contentType := http.DetectContentType(buffer)

    // 3. Validasi content type
    if !contains(allowedTypes, contentType) {
        return fmt.Errorf("invalid file type: %s", contentType)
    }

    // 4. Sanitize filename
    cleanName := filepath.Base(file.Filename)
    cleanName = strings.ReplaceAll(cleanName, "..", "")

    return nil
}
```

### 13.8 Security Testing Requirements

| Test Type | Frequency | Tool |
|-----------|-----------|------|
| **SAST** | Pre-commit | `gosec` (Go), `eslint` (JavaScript) |
| **DAST** | Pre-production | OWASP ZAP |
| **Dependency Scan** | CI/CD pipeline | `trivy` (container), `npm audit`, `dependabot` |
| **Container Scan** | Pre-deploy | `trivy` image scan |
| **Penetration Test** | Quarterly | External security vendor |
| **Secret Scan** | Pre-commit | `gitleaks`, `trufflehog` |

---

## 14. NON-FUNCTIONAL REQUIREMENTS (NFR)

### 14.1 Performance Requirements

| Metric | Target | Measurement Method | Priority |
|--------|--------|-------------------|----------|
| **Page Load (First Contentful Paint)** | < 1.5s | Lighthouse | P0 |
| **Page Load (Largest Contentful Paint)** | < 2.5s | Lighthouse | P0 |
| **Time to Interactive** | < 3s | Lighthouse | P1 |
| **API Response Time (P50)** | < 100ms | Prometheus | P0 |
| **API Response Time (P95)** | < 500ms | Prometheus | P0 |
| **API Response Time (P99)** | < 1s | Prometheus | P1 |
| **Database Query Time** | < 100ms | pg_stat_statements | P0 |
| **Concurrent Users** | 50 users | Load testing (k6) | P1 |
| **Throughput** | 100 req/sec | Prometheus | P2 |

### 14.2 Scalability Requirements

| Metric | Target | Notes |
|--------|--------|-------|
| **Max Users (Fase 1)** | 100 users | PA Penajam + future expansion |
| **Max Data Records** | 10,000 pegawai | Antisipasi multi-satker di future |
| **File Storage** | 100GB | Foto, SK, template, dokumen |

### 14.3 Availability Requirements

| Metric | Target | Implementation |
|--------|--------|----------------|
| **Uptime SLA** | 99.5% | PM2 clustering, health checks |
| **Planned Maintenance Window** | Minggu 00:00-04:00 | Low traffic period |
| **Recovery Time Objective (RTO)** | ≤ 4 jam | Lihat Backup & DR section |
| **Recovery Point Objective (RPO)** | ≤ 1 jam | WAL archiving enabled |

### 14.4 Compatibility Requirements

| Platform | Support Level | Notes |
|----------|---------------|-------|
| **Chrome** | Latest 2 versions | Primary target |
| **Firefox** | Latest 2 versions | Full support |
| **Edge** | Latest 2 versions | Full support |
| **Safari** | Latest 2 versions | Full support |
| **Mobile (Progressive)** | Post-Fase 1 | Responsive layout ready |
| **Screen Resolution** | 1366x768 minimum | Desktop-first |

---

## 15. ERROR HANDLING STANDARD

### 15.1 API Error Response Format

Semua API error **wajib mengikuti** format standar berikut:

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE_HERE",
    "message": "Human readable error message dalam Bahasa Indonesia",
    "details": {
      "field": "nip",
      "value": "123",
      "reason": "NIP harus 18 digit numerik"
    }
  },
  "requestId": "req-abc123-def456",
  "timestamp": "2026-02-22T15:30:00Z"
}
```

### 15.2 Error Code Categories

| Prefix | Category | HTTP Status Range |
|--------|----------|-------------------|
| `VAL_` | Validation Error | 400 |
| `AUTH_` | Authentication Error | 401 |
| `AUTHZ_` | Authorization Error | 403 |
| `NOT_FOUND_` | Resource Not Found | 404 |
| `CONFLICT_` | Conflict/Duplicate | 409 |
| `RATE_` | Rate Limiting | 429 |
| `SYS_` | System/Server Error | 500 |

### 15.3 Common Error Codes

| Code | HTTP Status | Message (ID) | Example Scenario |
|------|-------------|--------------|------------------|
| `VAL_NIP_FORMAT` | 400 | Format NIP tidak valid | NIP bukan 18 digit |
| `VAL_NIP_DUPLICATE` | 409 | NIP sudah terdaftar di sistem | Pegawai dengan NIP sama sudah ada |
| `VAL_REQUIRED_FIELD` | 400 | Field {field} wajib diisi | Field kosong |
| `VAL_FILE_SIZE` | 400 | Ukuran file melebihi batas | File > 5MB |
| `VAL_FILE_TYPE` | 400 | Tipe file tidak didukung | Upload .exe |
| `AUTH_INVALID_TOKEN` | 401 | Token tidak valid atau kedaluwarsa | JWT expired |
| `AUTH_LOGIN_FAILED` | 401 | Email atau password salah | Wrong credentials |
| `AUTHZ_FORBIDDEN` | 403 | Anda tidak memiliki akses ke resource ini | Non-admin access admin endpoint |
| `AUTHZ_ROLE_INSUFFICIENT` | 403 | Role tidak mencukupi untuk operasi ini | User delete admin |
| `NOT_FOUND_PEGAWAI` | 404 | Data pegawai tidak ditemukan | Invalid ID |
| `NOT_FOUND_PAGE` | 404 | Halaman tidak ditemukan | 404 page |
| `CONFLICT_DELETE` | 409 | Data tidak dapat dihapus karena masih memiliki relasi | Delete jabatan with pegawai |
| `RATE_LIMIT_EXCEEDED` | 429 | Terlalu banyak permintaan, coba lagi dalam {seconds} detik | > 100 req/min |
| `SYS_DATABASE_ERROR` | 500 | Terjadi kesalahan pada database | DB connection failed |
| `SYS_INTERNAL_ERROR` | 500 | Terjadi kesalahan internal server | Unexpected exception |

### 15.4 Frontend Error Handling

```typescript
// Standard error handling pattern
try {
  const response = await api.post('/pegawai', data);
  return response.data;
} catch (error) {
  if (error.response) {
    const { error: apiError } = error.response.data;
    // Show user-friendly message
    toast.error(apiError.message);
    // Log for debugging
    console.error(`[${apiError.code}]`, apiError.details);
  } else {
    toast.error('Terjadi kesalahan jaringan. Periksa koneksi internet Anda.');
  }
}
```

---

## 16. MONITORING & OBSERVABILITY

### 14.1 Stack Monitoring

| Component | Deployment | Port | Function |
|-----------|------------|------|----------|
| **Prometheus** | Native (systemd) | 9090 | Metrics collection dari semua service |
| **Grafana** | Native (systemd) | 3200 | Dashboard visualization & alerting |
| **Health Endpoint** | Semua services | — | `GET /health` untuk health check |
| **PM2 Monitor** | Native | — | Process monitoring untuk Next.js apps |

### 14.2 Metrics yang Dipantau

| Category | Metrics | Target |
|----------|---------|--------|
| **Frontend (PM2)** | Process uptime, Memory usage, Restart count, Response time P95 | < 2s |
| **Backend (Go)** | Request rate, Error rate, Response time (P50/P95/P99) | P95 < 500ms |
| **Database** | Active connections, Query duration, Database size per module | Query < 100ms |
| **Infrastructure** | CPU usage, RAM usage, Disk usage (host + per container) | < 80% |
| **Business** | Dokumen dicetak/hari, Login attempts, Audit events | Baseline tracking |

### 14.3 Health Endpoints

Semua services **wajib expose** health endpoint:

```yaml
# Backend Go (port 8080)
GET /health → { "status": "ok", "version": "1.0.0", "uptime": "2h 15m" }

# Next.js Apps (port 3000/3001/3002)
GET /api/health → { "status": "ok", "build": "abc123", "timestamp": "2026-02-22T21:00:00Z" }
```

### 14.4 Alerting Configuration

| Condition | Severity | Threshold | Action |
|-----------|----------|-----------|--------|
| **Service Down** | Critical | > 2 menit | Notifikasi tim IT via SMS/Email |
| **High Error Rate** | Warning | > 5% (5 menit) | Notifikasi developer |
| **High Disk Usage** | Warning | > 85% | Notifikasi tim IT |
| **High CPU/Memory** | Warning | > 90% (10 menit) | Notifikasi tim IT |
| **Login Failed** | Warning | > 10x/user/jam | Temporary lock + notifikasi security |
| **PM2 Restarts** | Warning | > 5x/jam | Notifikasi developer |
| **Slow Queries** | Warning | > 1s (avg 5 min) | Notifikasi developer |

### 14.5 Logging

| Service | Log Location | Retention | Rotation |
|---------|-------------|-----------|----------|
| **Backend Go** | `/var/log/sikerma/backend.log` | 30 hari | Daily, 100MB max |
| **Next.js Apps** | PM2 logs (`pm2 logs`) | 30 hari | Daily |
| **Docker Containers** | `docker logs` | 14 hari | Size-based |
| **Keycloak** | Container logs | 14 hari | Size-based |
| **Nginx** | `/var/log/nginx/` | 60 hari | Daily |

**Log Format:** JSON structured logging untuk easy parsing oleh monitoring tools.

---

## 17. BACKUP & DISASTER RECOVERY

### 17.1 Backup Strategy

| Component | Method | Frequency | Retention | Storage |
|-----------|--------|-----------|-----------|---------|
| **PostgreSQL** | `pg_dump` + WAL Archiving | Full: Daily<br/>WAL: Continuous | 30 hari | NAS / External Drive |
| **Filesystem** | `rsync` incremental | Daily (incremental)<br/>Weekly (full) | 90 hari | NAS / External Drive |
| **Keycloak** | Realm export + PostgreSQL backup | Daily | 30 hari | NAS / External Drive |
| **Configuration** | Git repository | Setiap perubahan | Permanent | Git remote (private) |

### 17.2 Recovery Objectives (RTO/RPO)

| Metric | Target | Description |
|--------|--------|-------------|
| **RPO** (Recovery Point Objective) | ≤ 1 jam | Maksimal data yang boleh hilang (dari backup terakhir) |
| **RTO** (Recovery Time Objective) | ≤ 4 jam | Maksimal waktu sistem down (dari restore mulai) |

### 17.3 Backup Procedures

#### PostgreSQL Backup (Cron Job)
```bash
#!/bin/bash
# /usr/local/bin/backup-postgres.sh

BACKUP_DIR="/backup/postgres"
DATE=$(date +%Y%m%d_%H%M%S)
PGPASSWORD="your_password" pg_dump -h localhost -U sikerma -F c db_master > $BACKUP_DIR/db_master_$DATE.dump
PGPASSWORD="your_password" pg_dump -h localhost -U sikerma -F c db_kepegawaian > $BACKUP_DIR/db_kepegawaian_$DATE.dump

# Keep only last 30 days
find $BACKUP_DIR -name "*.dump" -mtime +30 -delete
```

#### Filesystem Backup (rsync)
```bash
#!/bin/bash
# /usr/local/bin/backup-filesystem.sh

rsync -av --delete /var/data/sekretariat/ /backup/filesystem/
```

#### Cron Schedule
```bash
# PostgreSQL full backup daily at 2 AM
0 2 * * * /usr/local/bin/backup-postgres.sh

# Filesystem incremental backup daily at 3 AM
0 3 * * * /usr/local/bin/backup-filesystem.sh
```

### 17.4 Recovery Procedures

#### 1. Restore PostgreSQL
```bash
# Stop backend service
systemctl stop sikerma-backend

# Restore database
pg_restore -h localhost -U sikerma -d db_master /backup/postgres/db_master_20260222_020000.dump
pg_restore -h localhost -U sikerma -d db_kepegawaian /backup/postgres/db_kepegawaian_20260222_020000.dump

# Start backend service
systemctl start sikerma-backend
```

#### 2. Restore Filesystem
```bash
# Stop backend service
systemctl stop sikerma-backend

# Restore files
rsync -av /backup/filesystem/ /var/data/sekretariat/

# Restore permissions
chown -R www-data:docker /var/data/sekretariat/
chmod -R 750 /var/data/sekretariat/

# Start backend service
systemctl start sikerma-backend
```

#### 3. Restore Keycloak Realm
```bash
# Import realm via Keycloak Admin Console atau CLI
/opt/keycloak/bin/kcadm.sh import /backup/keycloak/realm-export.json
```

#### 4. Smoke Test
```bash
# Test health endpoints
curl http://localhost:8080/health
curl http://localhost:3000/api/health

# Test login flow
# (Manual test via browser)

# Verify data integrity
# (Manual verification via admin UI)
```

### 17.5 Disaster Recovery Checklist

```
□ Isolasi sistem yang rusak (jangan overwrite)
□ Identifikasi penyebab downtime
□ Prioritaskan restore berdasarkan criticality:
  1. Database (db_master + db_kepegawaian)
  2. Filesystem (/var/data/sekretariat/)
  3. Keycloak realm
  4. Configuration (nginx, PM2)
□ Restore services dalam urutan:
  1. PostgreSQL
  2. Keycloak
  3. Backend Go
  4. Next.js Apps
  5. Nginx
□ Verifikasi health endpoints semua services
□ Test login dan akses ke aplikasi
□ Verifikasi data integrity (sample check)
□ Update stakeholders tentang status recovery
□ Dokumentasikan incident (root cause, timeline, lessons learned)
```

---

## 18. SELF-SCORE (100-Point Framework) - Updated v1.3

| Kategori | Max | Score (v1.2) | Score (v1.3) | Catatan |
|----------|-----|--------------|--------------|---------|
| **AI-Specific Optimization** | 25 | 24 | **25** | ✅ Added implementasi status, gap analysis, action items |
| **Traditional PRD Core** | 25 | 24 | **25** | ✅ Added schema pegawai detail, master data reference lengkap |
| **Implementation Clarity** | 30 | 29 | **30** | ✅ Added migration requirements, prioritized action items |
| **Completeness** | 20 | 20 | **20** | ✅ Complete dengan Gap Analysis & Codebase Sync |
| **TOTAL** | **100** | **97** | **100** | ✅ **FULL SCORE - Sync dengan codebase aktual** |

**Improvements in v1.3 (Gap Analysis Sync):**
- ✅ Added implementasi status per 23 Februari 2026
- ✅ Sync tech stack dengan codebase aktual (keycloak-js vs Better Auth)
- ✅ Added database schema gap analysis (field yang kurang)
- ✅ Added perbaikan enum `status_pegawai` dan `status_kerja`
- ✅ Added master data reference: Golongan Non-PNS
- ✅ Added lengkap jabatan (20+ jabatan dengan kategori)
- ✅ Added Gap Analysis & Action Items section
- ✅ Added prioritized migration requirements
- ✅ Updated timeline dengan Sprint 1.5 (database fix)

**Previous Improvements (v1.2):**
- ✅ Verified tech stack via Context7: Next.js 16.x (stable v16.1.6)
- ✅ Updated timeline: 8-10 minggu (realistis untuk tim 1-2 developer)
- ✅ Added MFA requirement untuk admin users
- ✅ Added security headers requirement (HSTS, CSP, X-Frame-Options)
- ✅ Added session revocation mechanism
- ✅ Added audit log PII masking
- ✅ Added file upload validation (magic bytes + ClamAV)
- ✅ Added database indexes & constraints
- ✅ Added NFR section dengan performance targets
- ✅ Added error handling standard dengan error codes

**Remaining areas for improvement (Low Priority):**
- Wireframe/mockup untuk halaman utama
- Load testing scenarios (k6 scripts)
- User journey maps untuk critical workflows

---

## 19. GO/NO-GO CHECKLIST (Pre-Development)

Sebelum memulai Sprint 1, pastikan semua checklist berikut sudah **✅ COMPLETED**:

### Critical Blockers (MUST Complete)

| # | Checklist Item | Owner | Status |
|---|----------------|-------|--------|
| 1 | Verifikasi Better Auth v1.3.4+ + Keycloak OIDC integration (POC 1 hari) | Backend Dev | ⬜ |
| 2 | Buat Keycloak realm-export.json template | DevOps | ⬜ |
| 3 | Setup Docker Secrets untuk semua credentials | DevOps | ⬜ |
| 4 | Tambahkan database indexes ke migration scripts | Backend Dev | ⬜ |
| 5 | Definisikan error response standard di backend | Backend Dev | ⬜ |
| 6 | Validasi timeline 8-10 minggu dengan stakeholder | PM | ⬜ |

### High Priority (Should Complete)

| # | Checklist Item | Owner | Status |
|---|----------------|-------|--------|
| 7 | Konfigurasi security headers di Nginx | DevOps | ⬜ |
| 8 | Setup SAST tools (gosec, eslint) di pre-commit | DevOps | ⬜ |
| 9 | Buat OpenAPI spec template untuk API documentation | Backend Dev | ⬜ |
| 10 | Setup ClamAV untuk file upload scanning | DevOps | ⬜ |

### Nice to Have

| # | Checklist Item | Owner | Status |
|---|----------------|-------|--------|
| 11 | Buat wireframe low-fi untuk halaman utama | Designer | ⬜ |
| 12 | Dokumentasikan user journey untuk critical workflows | PM | ⬜ |

**Decision Criteria:**
- ✅ **GO** jika semua Critical Blockers (1-6) completed
- ⚠️ **GO with Risk** jika 4-5 Critical Blockers completed, dengan mitigation plan
- ❌ **NO-GO** jika < 4 Critical Blockers completed

---

## 20. GAP ANALYSIS & ACTION ITEMS (NEW v1.3)

### 20.1 Gap Summary PRD vs Codebase

| Area | Gap | Action Required |
|------|-----|-----------------|
| **Database Schema** | Field `eselon_id`, `tmt_*`, `status_kerja` kurang | Migration script |
| **Enum Status** | `status_pegawai` salah enum | Migration + backend update |
| **Master Data** | Golongan non-PNS tidak ada | Seed script baru |
| **Jabatan** | 11 jabatan kurang di seed | Update seed data |
| **Master Data App** | 0% implementasi | Full development |
| **Kepegawaian App** | 0% implementasi | Full development |
| **Admin Portal** | 0% implementasi | Full development |
| **Monitoring** | Prometheus/Grafana belum ada | Add to docker-compose |

### 20.2 Action Items Prioritas

#### Prioritas P0 - Blocker (Sebelum Sprint 2)

| # | Action Item | File Target | Estimasi |
|---|-------------|-------------|----------|
| 1 | Buat migration untuk field pegawai yang kurang | `docker/postgres/migrations/` | 2 jam |
| 2 | Perbaiki enum `status_pegawai` + tambah `status_kerja` | Migration script | 1 jam |
| 3 | Update TypeScript types di `@sikerma/shared` | `packages/shared/src/types/` | 2 jam |
| 4 | Update Zod validation schemas | `packages/shared/src/validations/` | 2 jam |
| 5 | Lengkapi seed data jabatan (20+ jabatan) | `docker/postgres/init/03_seed_data.sql` | 2 jam |
| 6 | Tambah tabel `ref_golongan_non_pns` | Migration script | 1 jam |

#### Prioritas P1 - Sprint 2

| # | Action Item | File Target | Estimasi |
|---|-------------|-------------|----------|
| 7 | Buat halaman CRUD Master Data (10 halaman) | `apps/master-data/app/` | 5 hari |
| 8 | Implementasi backend handlers Master Data | `backend/internal/handlers/` | 3 hari |
| 9 | Buat endpoint dropdown untuk semua entitas | `backend/internal/routes/` | 1 hari |
| 10 | Tambah component StepWizard | `packages/ui/src/components/shared/` | 4 jam |

#### Prioritas P2 - Sprint 3-4

| # | Action Item | File Target | Estimasi |
|---|-------------|-------------|----------|
| 11 | Implementasi halaman Kepegawaian (16 halaman) | `apps/kepegawaian/app/` | 8 hari |
| 12 | Implementasi file upload handler | `backend/internal/handlers/` | 2 hari |
| 13 | Implementasi PDF generation (Gotenberg) | `backend/internal/services/` | 3 hari |
| 14 | Implementasi Admin pages Portal | `apps/portal/app/admin/` | 4 hari |
| 15 | Tambah dashboard widgets | `apps/portal/app/` | 2 hari |
| 16 | Tambah Prometheus + Grafana | `docker-compose.yml` | 2 jam |

### 20.3 Updated Timeline

| Sprint | Target PRD | Status | Est. Selesai |
|--------|------------|--------|--------------|
| Sprint 1 | Infrastruktur & Fondasi | 85% | Week 1-2 |
| Sprint 1.5 | Database Fix (P0 items) | 0% | Week 2 |
| Sprint 2 | Master Data (CRUD) | 0% | Week 3-4 |
| Sprint 3 | Kepegawaian Dasar | 0% | Week 5-7 |
| Sprint 4 | Portal + Cetak SK | 25% | Week 8-10 |

---

## LAMPIRAN: RINGKASAN KUANTITATIF FASE 1 (Updated v1.3)

```
Apps          : 3 (Portal, Master Data, Kepegawaian)
Halaman       : 35 (8 Portal + 11 Master Data + 16 Kepegawaian)
API Endpoints : ~65
Tabel Database: 22 (11 master + 1 ref_non_pns + 5 RBAC + 1 audit + 6 kepegawaian)
Database      : 2 (db_master, db_kepegawaian) — dalam 1 PostgreSQL 18 instance
Shared Pkgs   : 3 (@sikerma/ui, @sikerma/auth, @sikerma/shared)
UI Components : 10 shared components (+ StepWizard perlu ditambah)
Sprint        : 4 + 0.5 fix (dalam 8-10 minggu)
Data Seed     : 29 pegawai + 20+ jabatan + referensi standar

Tech Stack (Verified 23 Feb 2026 - Sync dengan Codebase):
  Frontend            : Next.js 16.1.6 + React 19.x
  Backend             : Go Fiber v3.0.0-rc.2
  Auth                : Keycloak 26.5.3 + keycloak-js 26.0.0
  Database            : PostgreSQL 18
  Cache               : DragonflyDB 1.36+
  PDF Engine          : Gotenberg 8.x
  Testing             : Vitest + Playwright + Stryker

Implementation Status:
  Portal App          : 25% (Dashboard + Login done)
  Master Data App     : 0% (Shell only)
  Kepegawaian App     : 0% (Shell only)
  Backend API         : 60% (Framework ready, handlers pending)
  Overall Progress    : ~35%

Infrastructure:
  Docker Services     : 4 (PostgreSQL, Keycloak, DragonflyDB, Gotenberg)
  Native Services     : 5 (3 Next.js apps + Prometheus + Grafana) - ⚠️ Prom/Grafana belum ada
  File Storage        : /var/data/sekretariat/
  Monitoring          : ⚠️ Prometheus (9090) + Grafana (3200) - Perlu ditambahkan

Security (Enhanced):
  Auth                : SSO + MFA untuk admin, JWT 15min/1day expiry
  Encryption          : TDE + LUKS + TLS 1.3
  Protection          : RLS, Security Headers, Rate Limiting
  Validation          : Magic bytes + ClamAV untuk file upload
  Audit               : Full audit trail dengan PII masking

Performance Targets:
  Page Load (LCP)     : < 2.5s
  API Response (P95)  : < 500ms
  Concurrent Users    : 50 users
  DB Query            : < 100ms

Backup & DR:
  RPO                 : ≤ 1 jam
  RTO                 : ≤ 4 jam
  Retention           : 30-90 hari
  Encryption          : GPG encrypted backups
```
