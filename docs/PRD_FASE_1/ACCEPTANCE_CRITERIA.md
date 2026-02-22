# Acceptance Criteria - FASE 1

Dokumen ini berisi Definition of Done untuk setiap sprint dan overall Fase 1.

---

## Definition of Done (General)

Sebuah task/sprint dianggap **DONE** jika:

- ✅ Semua acceptance criteria terpenuhi
- ✅ Kode telah di-review dan approved
- ✅ Tidak ada critical bugs (P0/P1)
- ✅ Documentation lengkap (inline comments, README jika perlu)
- ✅ Test (jika ada) passing
- ✅ Build/deploy berhasil

---

## SPRINT 1: Infrastruktur & Fondasi

### Goal 1: Monorepo siap pakai
**Metric:** `pnpm install` sukses, `pnpm dev` bisa jalan
**Target:** 100%

**Acceptance Criteria:**
- [ ] `turbo.json` terkonfigurasi dengan benar
- [ ] `package.json` root memiliki workspace definition yang benar
- [ ] `pnpm-workspace.yaml` mendefinisikan apps/ dan packages/
- [ ] `.env.example` berisi semua variabel environment yang diperlukan
- [ ] Struktur folder apps/ dan packages/ terbuat
- [ ] `pnpm install` berjalan tanpa error
- [ ] `pnpm dev` dapat dijalankan (minimal tidak crash, meskipun apps belum ada)

---

### Goal 2: Backend Go Fiber berfungsi
**Metric:** Server listen di port 3003, API dapat diakses
**Target:** 100%

**Acceptance Criteria:**
- [ ] Go Fiber server dapat dijalankan (`go run cmd/server/main.go`)
- [ ] Server listen di port 3003 (atau dari env `BACKEND_PORT`)
- [ ] Route dasar `/api/v1/health` return `{ "status": "ok" }`
- [ ] Middleware CORS berfungsi (allow origin dari 3 apps)
- [ ] Middleware logger mencatat setiap request
- [ ] Struktur project modular (cmd/, internal/, pkg/)
- [ ] `go.mod` dan `go.sum` ter-generate dengan benar

---

### Goal 3: Keycloak terkonfigurasi
**Metric:** Realm + 3 clients + roles siap
**Target:** 100%

**Acceptance Criteria:**
- [ ] Keycloak container running di port 8081
- [ ] Realm `pengadilan-agama` terbuat dan dapat diakses
- [ ] 3 clients terdaftar:
  - [ ] `portal-client` dengan redirect URI `http://localhost:3000/*`
  - [ ] `master-data-client` dengan redirect URI `http://localhost:3001/*`
  - [ ] `kepegawaian-client` dengan redirect URI `http://localhost:3002/*`
- [ ] Client roles `[access]` terbuat untuk setiap client
- [ ] Test user dapat login via Keycloak admin console
- [ ] Realm export JSON tersedia untuk backup/recovery
- [ ] Keycloak admin console dapat diakses di `http://localhost:8081`

---

### Goal 4: Database ter-migrasi
**Metric:** 2 DB + 21 tabel terbuat + seed data
**Target:** 100%

**Acceptance Criteria:**
- [ ] PostgreSQL container running di port 5435
- [ ] Database `db_master` terbuat
- [ ] Database `db_kepegawaian` terbuat
- [ ] Semua tabel ter-create dengan schema yang benar:
  - [ ] **Master Data (11 tabel):**
    - [ ] `satker`
    - [ ] `jabatan`
    - [ ] `golongan`
    - [ ] `unit_kerja`
    - [ ] `eselon`
    - [ ] `ref_pendidikan`
    - [ ] `ref_agama`
    - [ ] `ref_status_kawin`
    - [ ] `ref_jenis_hukdis`
    - [ ] `ref_jenis_diklat`
    - [ ] (tabel lainnya)
  - [ ] **RBAC (4 tabel):**
    - [ ] `app_roles`
    - [ ] `app_permissions`
    - [ ] `role_permissions`
    - [ ] `user_app_roles`
  - [ ] **Audit (1 tabel):**
    - [ ] `audit_logs`
  - [ ] **Kepegawaian (6 tabel):**
    - [ ] `pegawai`
    - [ ] `riwayat_pangkat`
    - [ ] `riwayat_jabatan`
    - [ ] `riwayat_pendidikan`
    - [ ] `keluarga`
    - [ ] `template_dokumen`
- [ ] Foreign keys ter-set dengan benar
- [ ] Indexes ter-create untuk kolom yang sering di-query
- [ ] Data referensi ter-seed:
  - [ ] Satker: PA Penajam
  - [ ] Golongan: I, II/c, II/d, III/a, III/b, III/c, III/d, IV/a, V, IX
  - [ ] Jabatan: 17 jabatan dari data existing
  - [ ] Unit Kerja: Tree structure dari org_structure.json
  - [ ] Ref data lainnya (agama, status kawin, dll)
- [ ] 29 pegawai ter-seed dari `data_pegawai.json`
- [ ] RBAC default roles & permissions ter-seed
- [ ] Test query ke database berhasil

---

### Goal 5: Auth flow berfungsi
**Metric:** Login via Keycloak → redirect → session valid
**Target:** 100%

**Acceptance Criteria:**
- [ ] Better Auth terkonfigurasi dengan benar
- [ ] Login page menampilkan tombol "Login dengan SSO"
- [ ] Klik tombol login redirect ke Keycloak dengan parameter yang benar:
  - [ ] `client_id` sesuai app
  - [ ] `redirect_uri` ke `{app-url}/auth/callback`
  - [ ] `response_type=code`
  - [ ] `scope=openid profile email`
- [ ] User dapat login di Keycloak (test user)
- [ ] Setelah login, Keycloak redirect ke callback URL dengan `code` parameter
- [ ] Callback handler berhasil exchange `code` ke `access_token` + `refresh_token`
- [ ] Session tersimpan (Better Auth) dan valid
- [ ] User redirect ke dashboard setelah login sukses
- [ ] Session tetap valid setelah refresh page
- [ ] Logout berfungsi (clear session + redirect ke login)

---

### Goal 6: Shared UI components
**Metric:** 10 komponen reusable terbuat
**Target:** 100%

**Acceptance Criteria:**
- [ ] Package `@sikerma/ui` terbuat dan ter-build
- [ ] 10 komponen ter-export dari package:
  - [ ] `Sidebar` - navigasi samping dengan collapsible
  - [ ] `AppHeader` - header atas dengan user menu & app switcher
  - [ ] `PageHeader` - header halaman dengan title, description, actions
  - [ ] `Breadcrumb` - navigasi breadcrumb
  - [ ] `DataTable` - tabel dengan pagination, search, sort
  - [ ] `StatusBadge` - badge status dengan variant (success/warning/danger/info)
  - [ ] `FormDialog` - dialog modal untuk form create/edit
  - [ ] `DeleteConfirm` - dialog konfirmasi delete
  - [ ] `FileUpload` - upload file dengan preview & validation
  - [ ] `StepWizard` - multi-step wizard dengan progress indicator
- [ ] Semua komponen dapat di-import dari apps
- [ ] Props types ter-define dengan benar (TypeScript)
- [ ] Styling konsisten (pakai tailwind + shadcn classes)
- [ ] Responsive design (mobile-friendly)
- [ ] Accessibility (ARIA labels, keyboard navigation)

---

## SPRINT 2: Master Data (CRUD Lengkap)

### Goal 1: Generic CRUD backend
**Metric:** 1 pattern handler untuk semua entitas
**Target:** 100%

**Acceptance Criteria:**
- [ ] Generic handler dapat handle semua 10 entitas tanpa duplikasi kode
- [ ] Endpoint `GET /api/v1/master/{entity}` berfungsi (list dengan pagination)
- [ ] Endpoint `GET /api/v1/master/{entity}/:id` berfungsi (detail)
- [ ] Endpoint `POST /api/v1/master/{entity}` berfungsi (create)
- [ ] Endpoint `PUT /api/v1/master/{entity}/:id` berfungsi (update)
- [ ] Endpoint `DELETE /api/v1/master/{entity}/:id` berfungsi (soft delete)
- [ ] Endpoint `GET /api/v1/master/{entity}/dropdown` berfungsi (id + nama saja)
- [ ] Pagination bekerja (page, limit, total, total_pages)
- [ ] Search berfungsi (query parameter `search`, filter by `nama`)
- [ ] Sort berfungsi (query parameter `sort`, `order`)
- [ ] Soft delete (set `is_active = false`) bukan hard delete
- [ ] Validation berfungsi (required fields, format validation)
- [ ] Error handling jelas dan konsisten (status code + message)
- [ ] Audit trail ter-trigger untuk setiap operasi CRUD

---

### Goal 2: CRUD frontend 10 entitas
**Metric:** Semua entitas bisa Create, Read, Update, Delete
**Target:** 100%

**Acceptance Criteria:**
- [ ] 10 halaman CRUD terbuat:
  - [ ] `/master/satker`
  - [ ] `/master/jabatan`
  - [ ] `/master/golongan`
  - [ ] `/master/unit-kerja`
  - [ ] `/master/eselon`
  - [ ] `/master/pendidikan`
  - [ ] `/master/agama`
  - [ ] `/master/status-kawin`
  - [ ] `/master/hukuman-disiplin`
  - [ ] `/master/jenis-diklat`
- [ ] Setiap halaman memiliki:
  - [ ] Page header dengan title & actions
  - [ ] Data table dengan kolom yang relevan
  - [ ] Tombol "Tambah" untuk create
  - [ ] Aksi "Edit" pada setiap row
  - [ ] Aksi "Hapus" pada setiap row dengan konfirmasi
- [ ] Form dialog (create/edit) berfungsi:
  - [ ] Field validation (required, format)
  - [ ] Error message ditampilkan dengan jelas
  - [ ] Success message setelah submit
- [ ] Search berfungsi di semua halaman
- [ ] Pagination berfungsi (next/prev, page number)
- [ ] Sort berfungsi (klik header kolom)
- [ ] Unit Kerja menampilkan tree view untuk hierarki

---

### Goal 3: Dropdown API berfungsi
**Metric:** Setiap entitas punya endpoint `/dropdown`
**Target:** 100%

**Acceptance Criteria:**
- [ ] Semua 10 entitas memiliki endpoint `/dropdown`
- [ ] Response format konsisten: array of `{ id, nama }`
- [ ] Hanya return data aktif (`is_active = true`)
- [ ] Tidak ada pagination (return semua)
- [ ] Sorted by `nama ASC`
- [ ] Dapat diakses tanpa authentication (public endpoint)
- [ ] Response time < 100ms
- [ ] App Kepegawaian dapat consume dropdown API

---

### Goal 4: Data referensi ter-seed
**Metric:** Semua data dari existing ter-migrate
**Target:** 100%

**Acceptance Criteria:**
- [ ] Satker ter-seed: PA Penajam (kode: PA-PNJ)
- [ ] Golongan ter-seed (dari data pegawai): I, II/c, II/d, III/a, III/b, III/c, III/d, IV/a, V, IX
- [ ] Jabatan ter-seed (dari data pegawai): 17 jabatan
- [ ] Unit Kerja ter-seed dengan tree structure yang benar:
  - [ ] Root: Pengadilan Agama Penajam
  - [ ] Level 1: Pimpinan, Panitera, Sekretariat
  - [ ] Level 2: Panitera Muda (3), Subbag (3)
- [ ] Ref Agama ter-seed: Islam, Kristen, Katolik, Hindu, Budha, Konghucu
- [ ] Ref Status Kawin ter-seed: Kawin, Belum Kawin, Cerai Hidup, Cerai Mati
- [ ] Ref Pendidikan ter-seed: SD, SMP, SMA, D3, S1, S2, S3
- [ ] Data dapat dilihat di UI (frontend)
- [ ] Data dapat di-edit via UI

---

### Goal 5: UI konsisten & responsive
**Metric:** Pakai shared components, UX smooth
**Target:** Fully functional

**Acceptance Criteria:**
- [ ] Semua halaman menggunakan shared components dari `@sikerma/ui`
- [ ] Layout konsisten (Sidebar + AppHeader + Page content)
- [ ] Color scheme konsisten (primary color biru)
- [ ] Typography konsisten (font family, sizes)
- [ ] Spacing konsisten (padding, margin menggunakan tailwind classes)
- [ ] Responsive design:
  - [ ] Desktop (≥1024px): Sidebar expanded, full layout
  - [ ] Tablet (768px - 1023px): Sidebar collapsed, content adapted
  - [ ] Mobile (<768px): Mobile-friendly (stacked layout)
- [ ] Loading states ditampilkan saat fetching data
- [ ] Error states ditampilkan saat gagal
- [ ] Success messages ditampilkan setelah aksi berhasil
- [ ] Tidak ada console errors
- [ ] Tidak ada TypeScript errors
- [ ] Performance: Page load < 2s, interaction responsive

---

## SPRINT 3: Kepegawaian Dasar

### Goal 1: CRUD Pegawai lengkap
**Metric:** Tambah, lihat, edit, nonaktifkan pegawai
**Target:** 100%

**Acceptance Criteria:**
- [ ] Backend endpoint `GET /api/v1/pegawai` berfungsi (list dengan filter)
- [ ] Backend endpoint `GET /api/v1/pegawai/:nip` berfungsi (detail lengkap + semua riwayat)
- [ ] Backend endpoint `POST /api/v1/pegawai` berfungsi (create)
- [ ] Backend endpoint `PUT /api/v1/pegawai/:nip` berfungsi (update biodata)
- [ ] Backend endpoint `DELETE /api/v1/pegawai/:nip` berfungsi (soft delete)
- [ ] Backend endpoint `POST /api/v1/pegawai/:nip/foto` berfungsi (upload foto)
- [ ] Validasi NIP:
  - [ ] Harus 18 digit
  - [ ] Hanya angka
  - [ ] Unique (tidak boleh duplikat)
- [ ] Filter berfungsi:
  - [ ] Search by NIP atau nama
  - [ ] Filter by status pegawai (PNS/CPNS/PPPK/Honorer)
  - [ ] Filter by golongan
  - [ ] Filter by unit kerja
- [ ] Frontend halaman `/pegawai` menampilkan daftar pegawai dengan filter UI
- [ ] Frontend halaman `/pegawai/[nip]` menampilkan detail pegawai lengkap
- [ ] Frontend halaman `/pegawai/[nip]/edit` menampilkan form edit biodata
- [ ] Nonaktifkan pegawai set `is_active = false` + simpan alasan
- [ ] Pegawai nonaktif tidak muncul di list default, tapi bisa dilihat via filter

---

### Goal 2: Multi-step form wizard
**Metric:** Form tambah pegawai 4 step berfungsi
**Target:** 100%

**Acceptance Criteria:**
- [ ] Halaman `/pegawai/tambah` menampilkan step wizard
- [ ] Step indicator menampilkan 4 step:
  - [ ] Step 1: Biodata (NIP, nama, tempat/tanggal lahir, dll)
  - [ ] Step 2: Posisi & Pangkat (jabatan, golongan, unit kerja, TMT)
  - [ ] Step 3: Pendidikan (jenjang, sekolah, jurusan, tahun lulus)
  - [ ] Step 4: Keluarga (istri/suami, anak)
- [ ] Navigation berfungsi:
  - [ ] Tombol "Selanjutnya" pindah ke step berikutnya
  - [ ] Tombol "Sebelumnya" pindah ke step sebelumnya
  - [ ] Step indicator clickable untuk pindah step
- [ ] Form validation per step:
  - [ ] Step 1: NIP required & valid, nama required
  - [ ] Step 2: Jabatan, golongan, unit kerja required
  - [ ] Step 3 & 4: Optional (bisa skip)
- [ ] Data tersimpan sementara saat pindah step (tidak hilang)
- [ ] Tombol "Simpan Pegawai" di step terakhir submit semua data
- [ ] Success redirect ke halaman detail pegawai setelah simpan
- [ ] Progress bar menampilkan step saat ini
- [ ] UX smooth (animasi transition antar step)

---

### Goal 3: CRUD 4 jenis riwayat
**Metric:** Pangkat, jabatan, pendidikan, keluarga
**Target:** 100%

**Acceptance Criteria:**
- [ ] Backend CRUD untuk Riwayat Pangkat:
  - [ ] `GET /api/v1/pegawai/:nip/pangkat` (list)
  - [ ] `POST /api/v1/pegawai/:nip/pangkat` (create + upload file SK)
  - [ ] `PUT /api/v1/pegawai/:nip/pangkat/:id` (update)
  - [ ] `DELETE /api/v1/pegawai/:nip/pangkat/:id` (delete)
- [ ] Backend CRUD untuk Riwayat Jabatan (pattern sama)
- [ ] Backend CRUD untuk Riwayat Pendidikan (pattern sama)
- [ ] Backend CRUD untuk Data Keluarga (pattern sama)
- [ ] File upload untuk SK:
  - [ ] Max 5MB
  - [ ] Format PDF only
  - [ ] File tersimpan di server
  - [ ] URL file tersimpan di database
- [ ] Frontend tab "Riwayat Pangkat" menampilkan tabel riwayat
- [ ] Frontend tab "Riwayat Jabatan" menampilkan tabel riwayat
- [ ] Frontend tab "Riwayat Pendidikan" menampilkan tabel riwayat
- [ ] Frontend tab "Data Keluarga" menampilkan tabel keluarga
- [ ] Setiap tab memiliki tombol "Tambah" untuk create
- [ ] Setiap row memiliki aksi "Edit" dan "Hapus"
- [ ] Form dialog create/edit muncul di modal
- [ ] File SK dapat di-download dari tabel

---

### Goal 4: File upload berfungsi
**Metric:** Foto, SK, ijazah dapat diupload
**Target:** 100%

**Acceptance Criteria:**
- [ ] Upload foto pegawai:
  - [ ] Max 2MB
  - [ ] Format JPG/PNG only
  - [ ] Preview sebelum upload
  - [ ] Crop/resize option (optional)
  - [ ] File tersimpan di `uploads/foto/`
  - [ ] URL tersimpan di `pegawai.foto_url`
- [ ] Upload file SK:
  - [ ] Max 5MB
  - [ ] Format PDF only
  - [ ] File tersimpan di `uploads/sk/`
  - [ ] URL tersimpan di `riwayat_pangkat.file_sk_url` atau `riwayat_jabatan.file_sk_url`
- [ ] Upload file ijazah:
  - [ ] Max 5MB
  - [ ] Format PDF only
  - [ ] File tersimpan di `uploads/ijazah/`
  - [ ] URL tersimpan di `riwayat_pendidikan.file_ijazah_url`
- [ ] Error handling:
  - [ ] File terlalu besar → error message jelas
  - [ ] Format tidak sesuai → error message jelas
  - [ ] Upload gagal → error message jelas
- [ ] Success message setelah upload berhasil
- [ ] Loading indicator saat upload
- [ ] File dapat di-download dari UI

---

### Goal 5: Dashboard statistik
**Metric:** Chart & statistik menampilkan data real
**Target:** 100%

**Acceptance Criteria:**
- [ ] Dashboard menampilkan statistik cards:
  - [ ] Total Pegawai (semua status)
  - [ ] Total PNS
  - [ ] Total CPNS
  - [ ] Total PPPK
- [ ] Chart "Pegawai per Status":
  - [ ] Bar chart atau pie chart
  - [ ] Data real dari database
  - [ ] Legend jelas
  - [ ] Responsive
- [ ] Chart "Pegawai per Golongan":
  - [ ] Bar chart
  - [ ] Data real dari database
  - [ ] Sorted by jumlah descending
- [ ] Chart "Pegawai per Unit Kerja":
  - [ ] Bar chart atau horizontal bar
  - [ ] Data real dari database
- [ ] Data refresh otomatis saat ada perubahan
- [ ] Loading state saat fetching data
- [ ] Error state jika gagal load data
- [ ] Performance: Chart render < 1s

---

### Goal 6: 29 pegawai ter-verifikasi
**Metric:** Data dari JSON ter-migrate dan terlihat di UI
**Target:** 100%

**Acceptance Criteria:**
- [ ] Semua 29 pegawai dari `data_pegawai.json` ter-seed ke database
- [ ] Data dapat dilihat di halaman `/pegawai`
- [ ] NIP sesuai dengan JSON (18 digit)
- [ ] Nama sesuai dengan JSON
- [ ] Jabatan sesuai dengan JSON
- [ ] Golongan sesuai dengan JSON
- [ ] Unit Kerja sesuai dengan JSON
- [ ] TMT (TMT pangkat & jabatan) sesuai dengan JSON
- [ ] Data dapat di-edit via UI
- [ ] Data dapat di-delete (soft delete) via UI
- [ ] Verifikasi manual: Buka UI, cek 5-10 sample pegawai, pastikan datanya benar

---

## SPRINT 4: Portal + Cetak SK + Polish

### Goal 1: Portal launcher berfungsi
**Metric:** Tile app muncul sesuai hak akses user
**Target:** 100%

**Acceptance Criteria:**
- [ ] Login page menampilkan tombol "Login dengan SSO Pengadilan Agama"
- [ ] Klik tombol login redirect ke Keycloak
- [ ] User dapat login dengan credential Keycloak
- [ ] Setelah login, redirect ke dashboard Portal
- [ ] Dashboard menampilkan tile aplikasi:
  - [ ] Tile "Kepegawaian" muncul jika user punya permission `pegawai.*`
  - [ ] Tile "Master Data" muncul jika user punya permission `master.*`
  - [ ] Tile "Admin Portal" muncul jika user punya role `portal:admin`
- [ ] Klik tile redirect ke aplikasi yang sesuai (di subdomain/port berbeda)
- [ ] Session valid di semua aplikasi (SSO berfungsi)
- [ ] Logout dari satu app, logout dari semua app (single logout)

---

### Goal 2: Dashboard widgets informatif
**Metric:** 4+ widget menampilkan data real
**Target:** ≥ 4

**Acceptance Criteria:**
- [ ] Widget "Total Pegawai":
  - [ ] Menampilkan jumlah total pegawai aktif
  - [ ] Data real dari database
- [ ] Widget "Pegawai per Status":
  - [ ] Breakdown PNS, CPNS, PPPK
  - [ ] Data real dari database
- [ ] Widget "Aktivitas Terakhir":
  - [ ] Menampilkan 10 aktivitas terbaru dari audit log
  - [ ] Format: `{user} {action} {resource} pada {timestamp}`
  - [ ] Data real dari `audit_logs`
- [ ] Widget "Pegawai Akan Pensiun":
  - [ ] Menampilkan 5 pegawai dengan tanggal pensiun terdekat
  - [ ] Hitung BUP (60 tahun) dari `tanggal_lahir`
  - [ ] Data real dari database
- [ ] Widget "Pegawai per Unit Kerja":
  - [ ] Chart bar jumlah pegawai per unit kerja
  - [ ] Data real dari database
- [ ] Data auto-refresh setiap 30 detik (optional)
- [ ] Loading state saat fetching data

---

### Goal 3: Admin user management
**Metric:** List & manage users dari Keycloak
**Target:** Fully functional

**Acceptance Criteria:**
- [ ] Halaman `/admin/users` menampilkan daftar users dari Keycloak
- [ ] Kolom tabel: Username, Email, Nama Depan, Nama Belakang, Status, Akses Aplikasi
- [ ] Status: Active/Inactive (dari `enabled` field Keycloak)
- [ ] Akses Aplikasi: Badge per client role
- [ ] Tombol "Tambah User" membuka form create user
- [ ] Form create user:
  - [ ] Field: Username, Email, Password, Nama Depan, Nama Belakang
  - [ ] Validation required fields
  - [ ] Password strength indicator
  - [ ] Submit create user di Keycloak via API proxy
- [ ] Tombol "Edit" membuka form edit user
- [ ] Tombol "Kelola Akses" membuka dialog assign/revoke client access:
  - [ ] Checklist clients (Portal, Master Data, Kepegawaian)
  - [ ] Grant/revoke access
  - [ ] Submit update ke Keycloak
- [ ] Tombol "Reset Password" membuka dialog reset password
- [ ] Search berfungsi (filter by username/email)
- [ ] Filter berfungsi (filter by status)

---

### Goal 4: Admin RBAC UI
**Metric:** Assign role & permission dari UI
**Target:** Fully functional

**Acceptance Criteria:**
- [ ] Halaman `/admin/hak-akses` menampilkan:
  - [ ] Form assign role ke user (dropdown user + dropdown role)
  - [ ] Tombol "Assign Role"
  - [ ] Tabel daftar user-role assignments (User, Role, App)
  - [ ] Tombol "Hapus" untuk remove assignment
- [ ] Halaman `/admin/hak-akses/roles` menampilkan:
  - [ ] Card per role dengan info: Role Name, App, Description, Permissions
  - [ ] Tombol "Tambah Role" membuka form create role
  - [ ] Form create role:
    - [ ] Field: App, Role Name, Description
    - [ ] Submit create role di database
  - [ ] Tombol "Edit" membuka form edit role (non-system role only)
  - [ ] Tombol "Kelola Permissions" membuka dialog assign permissions:
    - [ ] Checklist permissions per app
    - [ ] Submit assign permissions ke role
  - [ ] Tombol "Hapus" untuk delete role (non-system role only)
- [ ] Validation:
  - [ ] Role name unique per app
  - [ ] System roles tidak bisa di-edit/hapus
- [ ] Success message setelah aksi berhasil
- [ ] Error message jika gagal

---

### Goal 5: Audit log viewer
**Metric:** Lihat & filter audit logs
**Target:** Fully functional

**Acceptance Criteria:**
- [ ] Halaman `/admin/audit-log` menampilkan tabel audit logs
- [ ] Kolom tabel: Waktu, Pengguna, Aplikasi, Action, Resource, ID, Detail
- [ ] Action badge dengan warna:
  - [ ] CREATE = hijau
  - [ ] UPDATE = kuning
  - [ ] DELETE = merah
- [ ] Filter UI:
  - [ ] Search (text filter)
  - [ ] Filter by App (dropdown)
  - [ ] Filter by Action (dropdown)
  - [ ] Filter by Date Range (date picker from-to)
- [ ] Tombol "Lihat Detail" membuka dialog dengan:
  - [ ] Old value (sebelum perubahan)
  - [ ] New value (setelah perubahan)
  - [ ] Diff highlight (opsional)
- [ ] Pagination berfungsi
- [ ] Sort berfungsi (klik header kolom)
- [ ] Export to CSV/Excel (opsional)
- [ ] Data real-time (auto-refresh atau manual refresh button)

---

### Goal 6: Cetak SK berfungsi
**Metric:** Minimal 3 template generate PDF
**Target:** ≥ 3

**Acceptance Criteria:**
- [ ] Gotenberg service running di port 3100
- [ ] Health check Gotenberg `/health` return 200
- [ ] Template SK Pangkat:
  - [ ] File .docx ter-upload
  - [ ] Placeholders: {{NIP}}, {{NAMA_LENGKAP}}, {{JABATAN}}, {{GOLONGAN}}, {{TANGGAL_SEKARANG}}
  - [ ] Generate PDF berhasil
  - [ ] PDF sesuai format yang diharapkan
- [ ] Template SK Jabatan:
  - [ ] File .docx ter-upload
  - [ ] Placeholders sama
  - [ ] Generate PDF berhasil
  - [ ] PDF sesuai format
- [ ] Template SK CPNS:
  - [ ] File .docx ter-upload
  - [ ] Placeholders sama
  - [ ] Generate PDF berhasil
  - [ ] PDF sesuai format
- [ ] Halaman `/cetak-sk`:
  - [ ] Dropdown pilih template
  - [ ] Dropdown pilih pegawai
  - [ ] Preview data pegawai sebelum cetak
  - [ ] Tombol "Generate & Download PDF"
  - [ ] Loading indicator saat generate
  - [ ] Auto-download PDF setelah generate selesai
- [ ] File PDF tersimpan di server (`uploads/cetak/`)
- [ ] History cetak tersimpan di database
- [ ] Error handling jika generate gagal

---

### Goal 7: Template management
**Metric:** Upload & kelola template .docx
**Target:** Fully functional

**Acceptance Criteria:**
- [ ] Halaman `/admin/template` menampilkan card per template
- [ ] Card menampilkan: Nama Template, Kode, Filename, Placeholders, Status (Active/Inactive)
- [ ] Toggle button untuk activate/deactivate template
- [ ] Tombol "Upload Template" membuka dialog:
  - [ ] File upload (.docx only)
  - [ ] Field: Nama Template, Kode
  - [ ] Field: Placeholders (JSON array, bisa multi-select)
  - [ ] Submit upload template
- [ ] Tombol "Download" untuk download file template .docx
- [ ] Tombol "Edit" untuk edit metadata template (non-system)
- [ ] Tombol "Hapus" untuk delete template (non-system)
- [ ] Validation:
  - [ ] File harus .docx
  - [ ] Nama template required
  - [ ] Kode template required & unique
- [ ] Success message setelah upload/edit/delete
- [ ] Error message jika gagal

---

### Goal 8: End-to-end testing
**Metric:** Semua flow utama ter-test
**Target:** 100%

**Acceptance Criteria:**
- [ ] Test Flow 1: Login → Akses App → CRUD:
  - [ ] User login via SSO
  - [ ] Redirect ke Portal dashboard
  - [ ] Klik tile "Master Data"
  - [ ] Create entitas baru (misal: Satker)
  - [ ] Edit entitas
  - [ ] Delete entitas
  - [ ] Verifikasi audit log tercatat
- [ ] Test Flow 2: Tambah Pegawai:
  - [ ] Login sebagai admin kepegawaian
  - [ ] Klik "Tambah Pegawai"
  - [ ] Isi step 1 (Biodata)
  - [ ] Isi step 2 (Posisi & Pangkat)
  - [ ] Skip step 3 & 4 atau isi opsional
  - [ ] Submit
  - [ ] Verifikasi pegawai muncul di list
  - [ ] Verifikasi detail pegawai benar
- [ ] Test Flow 3: Cetak SK:
  - [ ] Login sebagai admin kepegawaian
  - [ ] Buka halaman "Cetak SK"
  - [ ] Pilih template SK Pangkat
  - [ ] Pilih pegawai
  - [ ] Klik "Generate PDF"
  - [ ] Verifikasi PDF ter-download
  - [ ] Verifikasi isi PDF benar (data pegawai sesuai)
- [ ] Test Flow 4: Admin RBAC:
  - [ ] Login sebagai superadmin
  - [ ] Buka "Hak Akses"
  - [ ] Assign role ke user
  - [ ] Logout
  - [ ] Login sebagai user yang baru di-assign
  - [ ] Verifikasi akses sesuai role
- [ ] Test Flow 5: Audit Trail:
  - [ ] Lakukan beberapa operasi CRUD
  - [ ] Buka "Audit Log"
  - [ ] Verifikasi semua operasi tercatat dengan benar
  - [ ] Verifikasi old value & new value benar
- [ ] Semua test pass tanpa error
- [ ] Document hasil testing

---

### Goal 9: Bug fixing & polish
**Metric:** Zero critical bugs
**Target:** 0

**Acceptance Criteria:**
- [ ] Tidak ada P0 bugs (critical: crash, data loss, security)
- [ ] Tidak ada P1 bugs (high: major feature broken, data inconsistency)
- [ ] P2 bugs (medium: minor feature broken, UX issue) < 5
- [ ] P3 bugs (low: cosmetic, typo) < 10
- [ ] Code review passed (tidak ada major issues)
- [ ] Performance acceptable:
  - [ ] Page load < 3s
  - [ ] API response < 500ms
  - [ ] No memory leaks
- [ ] Documentation lengkap:
  - [ ] Setup guide (how to install & run)
  - [ ] API documentation
  - [ ] User manual (how to use fitur)
- [ ] Error handling robust:
  - [ ] User-friendly error messages
  - [ ] Graceful degradation
  - [ ] Logging yang memadai
- [ ] Accessibility:
  - [ ] Keyboard navigation works
  - [ ] Screen reader compatible (basic)
  - [ ] Color contrast memadai

---

## OVERALL FASE 1: Final Checklist

Fase 1 dianggap **COMPLETE** jika SEMUA checklist berikut terpenuhi:

### Infrastructure
- [ ] Monorepo (Turborepo + pnpm) berfungsi
- [ ] 3 Shared packages (@sikerma/ui, @sikerma/auth, @sikerma/shared) terbuat
- [ ] 3 Next.js apps (Portal, Master Data, Kepegawaian) dapat dijalankan
- [ ] Go Fiber backend berjalan di port 3003
- [ ] Keycloak realm + 3 clients terkonfigurasi
- [ ] PostgreSQL 2 database + 21 tabel ter-migrasi
- [ ] Gotenberg service running untuk PDF generation

### Authentication & Authorization
- [ ] SSO login via Keycloak berfungsi
- [ ] Session management (Better Auth) berfungsi
- [ ] RBAC middleware berfungsi (permission check)
- [ ] Audit trail middleware berfungsi (catat semua operasi)

### Master Data App
- [ ] 10 entitas referensi bisa CRUD (Satker, Jabatan, Golongan, Unit Kerja, dll)
- [ ] Generic CRUD backend handler berfungsi
- [ ] Dropdown API berfungsi untuk semua entitas
- [ ] Data referensi ter-seed lengkap
- [ ] Unit Kerja tree view berfungsi

### Kepegawaian App
- [ ] CRUD pegawai berfungsi (validasi NIP 18 digit)
- [ ] Multi-step wizard tambah pegawai (4 step) berfungsi
- [ ] Detail pegawai dengan tab view (5 tabs) berfungsi
- [ ] CRUD 4 jenis riwayat berfungsi (Pangkat, Jabatan, Pendidikan, Keluarga)
- [ ] File upload berfungsi (Foto max 2MB, SK/Ijazah max 5MB PDF)
- [ ] Dashboard statistik dengan chart berfungsi
- [ ] 29 pegawai ter-verifikasi di UI

### Portal App
- [ ] Login page → Keycloak → callback → session berfungsi
- [ ] Dashboard launcher menampilkan tile app sesuai hak akses
- [ ] Dashboard widgets (≥4) menampilkan data real
- [ ] Profil Saya dapat edit
- [ ] Admin User Management berfungsi
- [ ] Admin Hak Akses (assign role) berfungsi
- [ ] Admin Kelola Roles berfungsi
- [ ] Audit Log viewer berfungsi

### Cetak SK
- [ ] Template management (upload .docx) berfungsi
- [ ] Minimal 3 template SK berfungsi (Pangkat, Jabatan, CPNS)
- [ ] Generate PDF via Gotenberg berfungsi
- [ ] Auto-download PDF berfungsi

### Quality
- [ ] Zero P0/P1 bugs
- [ ] End-to-end testing semua flow utama pass
- [ ] Documentation lengkap (setup guide, user manual)
- [ ] Code review passed
- [ ] Performance acceptable (page load < 3s, API < 500ms)

---

## Sign-off

Fase 1 dinyatakan **COMPLETE** setelah:

- [ ] Semua acceptance criteria terpenuhi
- [ ] Product Owner approval
- [ ] Technical Lead approval
- [ ] QA Lead approval
- [ ] Documentation finalisasi

**Tanggal Completion:** __________

**Approved by:**
- Product Owner: ___________________
- Technical Lead: ___________________
- QA Lead: ___________________
