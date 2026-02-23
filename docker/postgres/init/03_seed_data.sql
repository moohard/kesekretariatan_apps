-- ============================================
-- 03_seed_data.sql
-- ============================================
-- Script untuk seed data awal
--
-- ARCHITECTURAL NOTE:
-- Cross-database queries tidak didukung di PostgreSQL.
-- Seed data pegawai memerlukan data referensi dari db_master.
-- Untuk verification, hanya db_master yang di-seed.
-- Seed data db_kepegawaian perlu dihandle terpisah dengan:
--   1. postgres_fdw extension, atau
--   2. Application-level data sync
-- ============================================

\c db_master;

-- ============================================
-- SEED DATA REFERENSI
-- ============================================

-- Golongan
INSERT INTO golongan (kode, nama, ruang, angka, min_pangkat, max_pangkat) VALUES
('I/a', 'Juru Muda', 'I', 1, 1, 5),
('I/b', 'Juru Muda Tingkat I', 'I', 2, 6, 10),
('I/c', 'Juru', 'I', 3, 11, 15),
('I/d', 'Juru Tingkat I', 'I', 4, 16, 20),
('II/a', 'Pengatur Muda', 'II', 5, 21, 24),
('II/b', 'Pengatur Muda Tingkat I', 'II', 6, 25, 28),
('II/c', 'Pengatur', 'II', 7, 29, 32),
('II/d', 'Pengatur Tingkat I', 'II', 8, 33, 36),
('III/a', 'Penata Muda', 'III', 9, 37, 40),
('III/b', 'Penata Muda Tingkat I', 'III', 10, 41, 44),
('III/c', 'Penata', 'III', 11, 45, 48),
('III/d', 'Penata Tingkat I', 'III', 12, 49, 52),
('IV/a', 'Pembina', 'IV', 13, 53, 56),
('IV/b', 'Pembina Tingkat I', 'IV', 14, 57, 60),
('IV/c', 'Pembina Utama Muda', 'IV', 15, 61, 64),
('IV/d', 'Pembina Utama Madya', 'IV', 16, 65, 68),
('IV/e', 'Pembina Utama', 'IV', 17, 69, 72);

-- Eselon
INSERT INTO eselon (kode, nama, tunjangan) VALUES
('I', 'Eselon I', 5000000),
('II', 'Eselon II', 4000000),
('III', 'Eselon III', 3000000),
('IV', 'Eselon IV', 2000000),
('V', 'Eselon V', 1000000),
('NON-ESELON', 'Non-Eselon', 0);

-- Agama
INSERT INTO ref_agama (kode, nama) VALUES
('ISLAM', 'Islam'),
('KATHOLIK', 'Katholik'),
('PROTESTAN', 'Kristen Protestan'),
('HINDU', 'Hindu'),
('BUDDHA', 'Buddha'),
('KONGHUCU', 'Konghucu');

-- Status Kawin
INSERT INTO ref_status_kawin (kode, nama) VALUES
('BK', 'Belum Kawin'),
('K0', 'Kawin (Tidak Punya Anak)'),
('K1', 'Kawin (Punya 1 Anak)'),
('K2', 'Kawin (Punya 2 Anak)'),
('K3', 'Kawin (Punya 3 Anak)'),
('K4', 'Kawin (Punya 4 Anak)'),
('K5', 'Kawin (Punya 5 Anak)'),
('J/D', 'Janda/Duda');

-- Pendidikan
INSERT INTO ref_pendidikan (kode, nama, tingkat) VALUES
('SD', 'Sekolah Dasar', 'SD'),
('SMP', 'Sekolah Menengah Pertama', 'SMP'),
('SMA', 'Sekolah Menengah Atas', 'SMA'),
('D3', 'Diploma 3', 'D3'),
('S1', 'Sarjana 1', 'S1'),
('S2', 'Sarjana 2 (Magister)', 'S2'),
('S3', 'Sarjana 3 (Doktor)', 'S3');

-- Jenis Hukdis
INSERT INTO ref_jenis_hukdis (kode, nama) VALUES
('RINGAN', 'Hukuman Disiplin Ringan'),
('SEDANG', 'Hukuman Disiplin Sedang'),
('BERAT', 'Hukuman Disiplin Berat');

-- Jenis Diklat
INSERT INTO ref_jenis_diklat (kode, nama) VALUES
('STRUKTURAL', 'Diklat Struktural'),
('TEKNIS', 'Diklat Teknis'),
('FUNGSIONAL', 'Diklat Fungsional'),
('LAINNYA', 'Diklat Lainnya');

-- ============================================
-- SEED DATA RBAC
-- ============================================

-- App Roles
INSERT INTO app_roles (nama, deskripsi) VALUES
('admin', 'Administrator dengan akses penuh'),
('supervisor', 'Supervisor dengan akses terbatas untuk approval'),
('officer', 'Officer dengan akses operasional'),
('staff', 'Staff dengan akses dasar'),
('user', 'User dengan akses read-only');

-- App Permissions
INSERT INTO app_permissions (nama, resource, action, deskripsi) VALUES
-- Master Data Permissions
('master_data.read', 'master_data', 'read', 'Akses baca data master'),
('master_data.create', 'master_data', 'create', 'Akses tambah data master'),
('master_data.update', 'master_data', 'update', 'Akses ubah data master'),
('master_data.delete', 'master_data', 'delete', 'Akses hapus data master'),
-- Kepegawaian Permissions
('kepegawaian.read', 'kepegawaian', 'read', 'Akses baca data kepegawaian'),
('kepegawaian.create', 'kepegawaian', 'create', 'Akses tambah pegawai'),
('kepegawaian.update', 'kepegawaian', 'update', 'Akses ubah data pegawai'),
('kepegawaian.delete', 'kepegawaian', 'delete', 'Akses hapus pegawai'),
-- RBAC Permissions
('rbac.read', 'rbac', 'read', 'Akses baca konfigurasi RBAC'),
('rbac.create', 'rbac', 'create', 'Akses tambah role/permission'),
('rbac.update', 'rbac', 'update', 'Akses ubah role/permission'),
('rbac.delete', 'rbac', 'delete', 'Akses hapus role/permission'),
-- Audit Permissions
('audit.read', 'audit', 'read', 'Akses baca audit logs');

-- Role Permissions untuk Admin
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM app_roles r, app_permissions p
WHERE r.nama = 'admin';

-- Role Permissions untuk Supervisor
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM app_roles r, app_permissions p
WHERE r.nama = 'supervisor' AND p.action IN ('read', 'update');

-- Role Permissions untuk Officer
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM app_roles r, app_permissions p
WHERE r.nama = 'officer' AND p.resource IN ('master_data', 'kepegawaian') AND p.action = 'read';

-- Role Permissions untuk Staff
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM app_roles r, app_permissions p
WHERE r.nama = 'staff' AND p.action = 'read';

-- Role Permissions untuk User
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM app_roles r, app_permissions p
WHERE r.nama = 'user' AND p.action = 'read';

-- ============================================
-- SEED DATA UNIT KERJA
-- ============================================

INSERT INTO unit_kerja (kode, nama, singkatan) VALUES
('01', 'Kepaniteraan Agama', 'KPA'),
('02', 'Kepaniteraan Perdata', 'KPP'),
('03', 'Kepaniteraan Pidana', 'KPN'),
('04', 'Kepaniteraan Jinayat', 'KPJ'),
('05', 'Kesekretariatan', 'Sekretariat'),
('06', 'Sub Bagian Kepegawaian', 'Sub.Kepegawaian'),
('07', 'Sub Bagian Keuangan', 'Sub.Keuangan'),
('08', 'Sub Bagian Perlengkapan', 'Sub.Perlengkapan'),
('09', 'Sub Bagian Umum', 'Sub.Umum'),
('10', 'Pengadilan Agama', 'PA');

-- ============================================
-- SEED DATA SATKER (CONTOH)
-- ============================================

INSERT INTO satker (kode, nama, parent_id, level, alamat, telepon, email) VALUES
('PA-0001', 'Pengadilan Agama Jakarta Pusat', NULL, 1, 'Jl. Kramat Raya No. 123, Jakarta Pusat', '021-1234567', 'pa-jakpus@mahkamahagung.go.id'),
('PA-0002', 'Pengadilan Agama Jakarta Selatan', NULL, 1, 'Jl. Fatmawati No. 45, Jakarta Selatan', '021-7654321', 'pa-jaksel@mahkamahagung.go.id'),
('PA-0003', 'Pengadilan Agama Jakarta Barat', NULL, 1, 'Jl. Daan Mogot No. 78, Jakarta Barat', '021-9876543', 'pa-jakbar@mahkamahagung.go.id'),
('PA-0004', 'Pengadilan Agama Jakarta Timur', NULL, 1, 'Jl. Jatinegara Timur No. 56, Jakarta Timur', '021-3456789', 'pa-jaktim@mahkamahagung.go.id'),
('PA-0005', 'Pengadilan Agama Jakarta Utara', NULL, 1, 'Jl. Yos Sudarso No. 90, Jakarta Utara', '021-5678901', 'pa-jakut@mahkamahagung.go.id');

-- ============================================
-- SEED DATA JABATAN (CONTOH)
-- ============================================

INSERT INTO jabatan (kode, nama, eselon_id, kelas) VALUES
('KETUA', 'Ketua Pengadilan Agama', (SELECT id FROM eselon WHERE kode = 'II'), '1A'),
('WAKIL-KETUA', 'Wakil Ketua Pengadilan Agama', (SELECT id FROM eselon WHERE kode = 'III'), '1B'),
('SEKRETARIS', 'Sekretaris', (SELECT id FROM eselon WHERE kode = 'III'), '1B'),
('KPA-KEPALA', 'Kepala Kepaniteraan Agama', (SELECT id FROM eselon WHERE kode = 'IV'), '2A'),
('KPP-KEPALA', 'Kepala Kepaniteraan Perdata', (SELECT id FROM eselon WHERE kode = 'IV'), '2A'),
('KPN-KEPALA', 'Kepala Kepaniteraan Pidana', (SELECT id FROM eselon WHERE kode = 'IV'), '2A'),
('KPJ-KEPALA', 'Kepala Kepaniteraan Jinayat', (SELECT id FROM eselon WHERE kode = 'IV'), '2A'),
('SUBBAG-KEPALA', 'Kepala Sub Bagian', (SELECT id FROM eselon WHERE kode = 'IV'), '2A'),
('HAKIM', 'Hakim', (SELECT id FROM eselon WHERE kode = 'III'), '1B'),
('PANITERA', 'Panitera', (SELECT id FROM eselon WHERE kode = 'IV'), '2A'),
('JURUSITA', 'Jurusita', (SELECT id FROM eselon WHERE kode = 'V'), '2B'),
('STAF', 'Staf', (SELECT id FROM eselon WHERE kode = 'NON-ESELON'), '2B');

-- ============================================
-- SEED DATA db_kepegawaian (SIMPLIFIED)
-- ============================================
-- NOTE: Cross-database queries not supported
-- Only creating minimal template data

\c db_kepegawaian;

-- Template Dokumen (doesn't require cross-db references)
INSERT INTO template_dokumen (kode, nama, tipe, konten_html) VALUES
('SK-PANGKAT', 'Template SK Kenaikan Pangkat', 'sk_pangkat', '<html><body><h1>SURAT KEPUTUSAN</h1><p>Tentang Kenaikan Pangkat</p><p>No: {{nomor_sk}}</p><p>Diberikan kepada: {{nama_pegawai}}</p><p>NIP: {{nip}}</p><p>Golongan: {{golongan}}</p><p>TMT: {{tmt}}</p></body></html>'),
('SK-JABATAN', 'Template SK Pengangkatan Jabatan', 'sk_jabatan', '<html><body><h1>SURAT KEPUTUSAN</h1><p>Tentang Pengangkatan Jabatan</p><p>No: {{nomor_sk}}</p><p>Diberikan kepada: {{nama_pegawai}}</p><p>NIP: {{nip}}</p><p>Jabatan: {{jabatan}}</p><p>TMT: {{tmt}}</p></body></html>'),
('SURAT-TUGAS', 'Template Surat Tugas', 'surat_tugas', '<html><body><h1>SURAT TUGAS</h1><p>No: {{nomor_surat}}</p><p>Diberikan kepada: {{nama_pegawai}}</p><p>NIP: {{nip}}</p><p>Untuk: {{keperluan}}</p><p>Tempat: {{tempat}}</p><p>Waktu: {{waktu}}</p></body></html>');

-- ============================================
-- SEED SELESAI
-- ============================================

\c db_master;
\echo "Seed data berhasil dibuat!"
\echo "Total Golongan: " || (SELECT COUNT(*) FROM golongan)
\echo "Total Agama: " || (SELECT COUNT(*) FROM ref_agama)
\echo "Total Status Kawin: " || (SELECT COUNT(*) FROM ref_status_kawin)
\echo "Total App Roles: " || (SELECT COUNT(*) FROM app_roles)
\echo "Total App Permissions: " || (SELECT COUNT(*) FROM app_permissions)
\echo "Total Satker: " || (SELECT COUNT(*) FROM satker)
\echo "Total Jabatan: " || (SELECT COUNT(*) FROM jabatan)
\echo "Total Unit Kerja: " || (SELECT COUNT(*) FROM unit_kerja)

\c db_kepegawaian;
\echo "Total Template Dokumen: " || (SELECT COUNT(*) FROM template_dokumen)
