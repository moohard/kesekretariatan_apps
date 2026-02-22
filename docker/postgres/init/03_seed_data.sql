-- ============================================
-- 03_seed_data.sql
-- ============================================
-- Script untuk seed data awal
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
-- SEED DATA PEGAWAI (SAMPLE - BUKAN 29 PEGAWAI ASLI)
-- ============================================

\c db_kepegawaian;

INSERT INTO pegawai (nip, nama, gelar_depan, gelar_belakang, tempat_lahir, tanggal_lahir, jenis_kelamin, agama_id, status_kawin_id, nik, email, telepon, alamat, satker_id, jabatan_id, unit_kerja_id, golongan_id, status_pegawai, tmt_jabatan, is_pns) VALUES
('197601012005011001', 'Ahmad Fauzi', '', 'S.H., M.H.', 'Jakarta', '1976-01-01', 'L', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'K2'), '3171234567890001', 'ahmad.fauzi@mahkamahagung.go.id', '081234567890', 'Jl. Sudirman No. 10, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'KETUA'), (SELECT id FROM db_master.unit_kerja WHERE kode = '10'), (SELECT id FROM db_master.golongan WHERE kode = 'IV/c'), 'aktif', '2020-01-01', true),
('198203022010011002', 'Budi Santoso', '', 'S.H.I.', 'Bandung', '1982-03-02', 'L', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'K1'), '3271234567890002', 'budi.santoso@mahkamahagung.go.id', '081234567891', 'Jl. Gatot Subroto No. 20, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'WAKIL-KETUA'), (SELECT id FROM db_master.unit_kerja WHERE kode = '10'), (SELECT id FROM db_master.golongan WHERE kode = 'IV/b'), 'aktif', '2021-01-01', true),
('198505122012021003', 'Citra Dewi', 'Dra.', 'M.Hum.', 'Surabaya', '1985-05-12', 'P', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'BK'), '3571234567890003', 'citra.dewi@mahkamahagung.go.id', '081234567892', 'Jl. Diponegoro No. 30, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'SEKRETARIS'), (SELECT id FROM db_master.unit_kerja WHERE kode = '05'), (SELECT id FROM db_master.golongan WHERE kode = 'IV/a'), 'aktif', '2022-01-01', true),
('199010082015032004', 'Dedi Kurniawan', '', 'S.H.', 'Yogyakarta', '1990-10-08', 'L', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'BK'), '3471234567890004', 'dedi.kurniawan@mahkamahagung.go.id', '081234567893', 'Jl. Ahmad Yani No. 40, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'KPA-KEPALA'), (SELECT id FROM db_master.unit_kerja WHERE kode = '01'), (SELECT id FROM db_master.golongan WHERE kode = 'III/d'), 'aktif', '2023-01-01', true),
('199307152017031005', 'Eka Pratama', '', 'S.H.I., M.H.I.', 'Semarang', '1993-07-15', 'L', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'BK'), '3371234567890005', 'eka.pratama@mahkamahagung.go.id', '081234567894', 'Jl. Sudirman No. 50, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'KPP-KEPALA'), (SELECT id FROM db_master.unit_kerja WHERE kode = '02'), (SELECT id FROM db_master.golongan WHERE kode = 'III/c'), 'aktif', '2023-01-01', true),
('199402202018021006', 'Fitriani', 'Dra.', 'M.A.', 'Medan', '1994-02-20', 'P', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'K0'), '3171234567890006', 'fitriani@mahkamahagung.go.id', '081234567895', 'Jl. Rasuna Said No. 60, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'KPN-KEPALA'), (SELECT id FROM db_master.unit_kerja WHERE kode = '03'), (SELECT id FROM db_master.golongan WHERE kode = 'III/c'), 'aktif', '2024-01-01', true),
('199509252019031007', 'Gunawan', '', 'S.H.I.', 'Makassar', '1995-09-25', 'L', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'BK'), '3171234567890007', 'gunawan@mahkamahagung.go.id', '081234567896', 'Jl. Gatot Subroto No. 70, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'KPJ-KEPALA'), (SELECT id FROM db_master.unit_kerja WHERE kode = '04'), (SELECT id FROM db_master.golongan WHERE kode = 'III/c'), 'aktif', '2024-01-01', true),
('199603102020022008', 'Hendra Wijaya', '', 'S.Kom.', 'Bandar Lampung', '1996-03-10', 'L', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'BK'), '3271234567890008', 'hendra.wijaya@mahkamahagung.go.id', '081234567897', 'Jl. Sudirman No. 80, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'SUBBAG-KEPALA'), (SELECT id FROM db_master.unit_kerja WHERE kode = '06'), (SELECT id FROM db_master.golongan WHERE kode = 'III/b'), 'aktif', '2024-01-01', true),
('199706152021012009', 'Indah Permata', '', 'S.E.', 'Palembang', '1997-06-15', 'P', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'BK'), '3171234567890009', 'indah.permata@mahkamahagung.go.id', '081234567898', 'Jl. Fatmawati No. 90, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'HAKIM'), (SELECT id FROM db_master.unit_kerja WHERE kode = '10'), (SELECT id FROM db_master.golongan WHERE kode = 'III/a'), 'aktif', '2024-01-01', true),
('199812202022031010', 'Joko Susilo', '', 'S.H.I.', 'Surakarta', '1998-12-20', 'L', (SELECT id FROM db_master.ref_agama WHERE kode = 'ISLAM'), (SELECT id FROM db_master.ref_status_kawin WHERE kode = 'BK'), '3471234567890010', 'joko.susilo@mahkamahagung.go.id', '081234567899', 'Jl. Ahmad Yani No. 100, Jakarta', (SELECT id FROM db_master.satker WHERE kode = 'PA-0001'), (SELECT id FROM db_master.jabatan WHERE kode = 'PANITERA'), (SELECT id FROM db_master.unit_kerja WHERE kode = '01'), (SELECT id FROM db_master.golongan WHERE kode = 'III/a'), 'aktif', '2024-01-01', true);

-- ============================================
-- SEED DATA RIWAYAT PANGKAT (SAMPLE)
-- ============================================

INSERT INTO riwayat_pangkat (pegawai_id, golongan_id, pangkat, tmt, nomor_sk, tanggal_sk, pejabat, gaji_pokok, is_terakhir)
SELECT p.id, g.id, CASE g.angka
    WHEN 13 THEN 'Pembina'
    WHEN 14 THEN 'Pembina Tingkat I'
    WHEN 15 THEN 'Pembina Utama Muda'
    WHEN 16 THEN 'Pembina Utama Madya'
    WHEN 17 THEN 'Pembina Utama'
    WHEN 9 THEN 'Penata Muda'
    WHEN 10 THEN 'Penata Muda Tingkat I'
    WHEN 11 THEN 'Penata'
    WHEN 12 THEN 'Penata Tingkat I'
    ELSE 'Pegawai'
END, p.tmt_jabatan, 'SK/' || p.nip || '/2024', '2024-01-01', 'Ketua MA', 4000000, true
FROM pegawai p
JOIN db_master.golongan g ON p.golongan_id = g.id;

-- ============================================
-- SEED DATA RIWAYAT JABATAN (SAMPLE)
-- ============================================

INSERT INTO riwayat_jabatan (pegawai_id, jabatan_id, unit_kerja_id, satker_id, nama_jabatan, tmt, nomor_sk, tanggal_sk, pejabat, is_terakhir)
SELECT p.id, p.jabatan_id, p.unit_kerja_id, p.satker_id, j.nama, p.tmt_jabatan, 'SKJ/' || p.nip || '/2024', '2024-01-01', 'Ketua MA', true
FROM pegawai p
LEFT JOIN db_master.jabatan j ON p.jabatan_id = j.id
WHERE p.jabatan_id IS NOT NULL;

-- ============================================
-- SEED DATA TEMPLATE DOKUMEN
-- ============================================

INSERT INTO template_dokumen (kode, nama, tipe, konten_html) VALUES
('SK-PANGKAT', 'Template SK Kenaikan Pangkat', 'sk_pangkat', '<html><body><h1>SURAT KEPUTUSAN</h1><p>Tentang Kenaikan Pangkat</p><p>No: {{nomor_sk}}</p><p>Diberikan kepada: {{nama_pegawai}}</p><p>NIP: {{nip}}</p><p>Golongan: {{golongan}}</p><p>TMT: {{tmt}}</p></body></html>'),
('SK-JABATAN', 'Template SK Pengangkatan Jabatan', 'sk_jabatan', '<html><body><h1>SURAT KEPUTUSAN</h1><p>Tentang Pengangkatan Jabatan</p><p>No: {{nomor_sk}}</p><p>Diberikan kepada: {{nama_pegawai}}</p><p>NIP: {{nip}}</p><p>Jabatan: {{jabatan}}</p><p>TMT: {{tmt}}</p></body></html>'),
('SURAT-TUGAS', 'Template Surat Tugas', 'surat_tugas', '<html><body><h1>SURAT TUGAS</h1><p>No: {{nomor_surat}}</p><p>Diberikan kepada: {{nama_pegawai}}</p><p>NIP: {{nip}}</p><p>Untuk: {{keperluan}}</p><p>Tempat: {{tempat}}</p><p>Waktu: {{waktu}}</p></body></html>');

-- ============================================
-- SEED SELESAI
-- ============================================

\echo "Seed data berhasil dibuat!"
\echo "Total Golongan: " || (SELECT COUNT(*) FROM db_master.golongan)
\echo "Total Agama: " || (SELECT COUNT(*) FROM db_master.ref_agama)
\echo "Total Status Kawin: " || (SELECT COUNT(*) FROM db_master.ref_status_kawin)
\echo "Total App Roles: " || (SELECT COUNT(*) FROM db_master.app_roles)
\echo "Total App Permissions: " || (SELECT COUNT(*) FROM db_master.app_permissions)
\echo "Total Satker: " || (SELECT COUNT(*) FROM db_master.satker)
\echo "Total Jabatan: " || (SELECT COUNT(*) FROM db_master.jabatan)
\echo "Total Unit Kerja: " || (SELECT COUNT(*) FROM db_master.unit_kerja)
\echo "Total Pegawai: " || (SELECT COUNT(*) FROM pegawai)
\echo "Total Template Dokumen: " || (SELECT COUNT(*) FROM template_dokumen)