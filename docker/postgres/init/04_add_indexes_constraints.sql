-- ============================================
-- 04_add_indexes_constraints.sql
-- ============================================
-- Script untuk menambahkan missing indexes dan constraints
-- Berdasarkan PRD_REMEDIATION.md RM-011 dan RM-012
-- ============================================

-- ============================================
-- MASTER DATABASE (db_master)
-- ============================================
\c db_master;

-- ============================================
-- RM-011: Missing Indexes
-- ============================================

-- Composite index untuk audit_logs (resource_type + resource_id)
-- Mempercepat query audit berdasarkan resource
CREATE INDEX IF NOT EXISTS idx_audit_logs_resource_composite ON audit_logs(resource, resource_id);

-- Full-text search configuration untuk Indonesian language
CREATE EXTENSION IF NOT EXISTS unaccent;

-- ============================================
-- KEPEGAWAIAN DATABASE (db_kepegawaian)
-- ============================================
\c db_kepegawaian;

-- ============================================
-- RM-011: Missing Indexes
-- ============================================

-- Index untuk pegawai.is_active (filtering active/inactive)
CREATE INDEX IF NOT EXISTS idx_pegawai_is_active ON pegawai(is_active);

-- Index untuk pegawai.unit_kerja_id (filtering by unit)
CREATE INDEX IF NOT EXISTS idx_pegawai_unit_kerja ON pegawai(unit_kerja_id);

-- Full-text search index untuk nama pegawai
-- Menggunakan GIN index dengan Indonesian text search configuration
CREATE INDEX IF NOT EXISTS idx_pegawai_nama_fts ON pegawai
    USING GIN(to_tsvector('indonesian', nama));

-- Composite index untuk pencarian pegawai dengan filter umum
CREATE INDEX IF NOT EXISTS idx_pegawai_search ON pegawai(satker_id, status_pegawai, is_active);

-- ============================================
-- RM-012: Data Validation Constraints
-- ============================================

-- Constraint: NIP harus 18 digit angka
ALTER TABLE pegawai
DROP CONSTRAINT IF EXISTS chk_nip_format;

ALTER TABLE pegawai
ADD CONSTRAINT chk_nip_format CHECK (nip ~ '^\d{18}$');

-- Constraint: NIK harus 16 digit angka (jika diisi)
ALTER TABLE pegawai
DROP CONSTRAINT IF EXISTS chk_nik_format;

ALTER TABLE pegawai
ADD CONSTRAINT chk_nik_format CHECK (nik IS NULL OR nik ~ '^\d{16}$');

-- Constraint: Tanggal lahir tidak boleh di masa depan
ALTER TABLE pegawai
DROP CONSTRAINT IF EXISTS chk_tanggal_lahir;

ALTER TABLE pegawai
ADD CONSTRAINT chk_tanggal_lahir CHECK (tanggal_lahir IS NULL OR tanggal_lahir <= CURRENT_DATE);

-- Constraint: Email format validation (jika diisi)
ALTER TABLE pegawai
DROP CONSTRAINT IF EXISTS chk_email_format;

ALTER TABLE pegawai
ADD CONSTRAINT chk_email_format CHECK (
    email IS NULL OR
    email = '' OR
    email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
);

-- Constraint: Jenis kelamin harus L atau P
ALTER TABLE pegawai
DROP CONSTRAINT IF EXISTS chk_jenis_kelamin;

ALTER TABLE pegawai
ADD CONSTRAINT chk_jenis_kelamin CHECK (jenis_kelamin IN ('L', 'P'));

-- Constraint: Status pegawai harus valid
ALTER TABLE pegawai
DROP CONSTRAINT IF EXISTS chk_status_pegawai;

ALTER TABLE pegawai
ADD CONSTRAINT chk_status_pegawai CHECK (
    status_pegawai IN ('aktif', 'pensiun', 'mutasi', 'cps', 'clt', 'meninggal', 'keluar')
);

-- ============================================
-- Constraints untuk RIWAYAT_PANGKAT
-- ============================================

-- Constraint: TMT pangkat tidak boleh di masa depan
ALTER TABLE riwayat_pangkat
DROP CONSTRAINT IF EXISTS chk_tmt_pangkat;

ALTER TABLE riwayat_pangkat
ADD CONSTRAINT chk_tmt_pangkat CHECK (tmt <= CURRENT_DATE);

-- Constraint: Tanggal SK tidak boleh di masa depan
ALTER TABLE riwayat_pangkat
DROP CONSTRAINT IF EXISTS chk_tanggal_sk_pangkat;

ALTER TABLE riwayat_pangkat
ADD CONSTRAINT chk_tanggal_sk_pangkat CHECK (tanggal_sk <= CURRENT_DATE);

-- Constraint: Gaji pokok tidak boleh negatif
ALTER TABLE riwayat_pangkat
DROP CONSTRAINT IF EXISTS chk_gaji_pokok;

ALTER TABLE riwayat_pangkat
ADD CONSTRAINT chk_gaji_pokok CHECK (gaji_pokok >= 0);

-- ============================================
-- Constraints untuk RIWAYAT_JABATAN
-- ============================================

-- Constraint: TMT jabatan tidak boleh di masa depan
ALTER TABLE riwayat_jabatan
DROP CONSTRAINT IF EXISTS chk_tmt_jabatan;

ALTER TABLE riwayat_jabatan
ADD CONSTRAINT chk_tmt_jabatan CHECK (tmt <= CURRENT_DATE);

-- Constraint: Tanggal SK tidak boleh di masa depan
ALTER TABLE riwayat_jabatan
DROP CONSTRAINT IF EXISTS chk_tanggal_sk_jabatan;

ALTER TABLE riwayat_jabatan
ADD CONSTRAINT chk_tanggal_sk_jabatan CHECK (tanggal_sk <= CURRENT_DATE);

-- ============================================
-- Constraints untuk KELUARGA
-- ============================================

-- Constraint: Status keluarga harus valid
ALTER TABLE keluarga
DROP CONSTRAINT IF EXISTS chk_status_keluarga;

ALTER TABLE keluarga
ADD CONSTRAINT chk_status_keluarga CHECK (
    status_keluarga IN ('Suami', 'Istri', 'Anak', 'Ayah', 'Ibu', 'Saudara')
);

-- Constraint: NIK harus 16 digit (jika diisi)
ALTER TABLE keluarga
DROP CONSTRAINT IF EXISTS chk_keluarga_nik_format;

ALTER TABLE keluarga
ADD CONSTRAINT chk_keluarga_nik_format CHECK (nik IS NULL OR nik ~ '^\d{16}$');

-- Constraint: Tanggal lahir tidak boleh di masa depan
ALTER TABLE keluarga
DROP CONSTRAINT IF EXISTS chk_keluarga_tanggal_lahir;

ALTER TABLE keluarga
ADD CONSTRAINT chk_keluarga_tanggal_lahir CHECK (tanggal_lahir IS NULL OR tanggal_lahir <= CURRENT_DATE);

-- Constraint: Jenis kelamin harus L atau P (jika diisi)
ALTER TABLE keluarga
DROP CONSTRAINT IF EXISTS chk_keluarga_jenis_kelamin;

ALTER TABLE keluarga
ADD CONSTRAINT chk_keluarga_jenis_kelamin CHECK (jenis_kelamin IS NULL OR jenis_kelamin IN ('L', 'P'));

-- ============================================
-- Constraints untuk HUKDIS
-- ============================================

-- Constraint: Tanggal SK tidak boleh di masa depan
ALTER TABLE hukdis
DROP CONSTRAINT IF EXISTS chk_hukdis_tanggal_sk;

ALTER TABLE hukdis
ADD CONSTRAINT chk_hukdis_tanggal_sk CHECK (tanggal_sk <= CURRENT_DATE);

-- Constraint: Tanggal mulai tidak boleh di masa depan
ALTER TABLE hukdis
DROP CONSTRAINT IF EXISTS chk_hukdis_tanggal_mulai;

ALTER TABLE hukdis
ADD CONSTRAINT chk_hukdis_tanggal_mulai CHECK (tanggal_mulai <= CURRENT_DATE);

-- Constraint: Tanggal selesai harus setelah tanggal mulai (jika diisi)
ALTER TABLE hukdis
DROP CONSTRAINT IF EXISTS chk_hukdis_tanggal_selesai;

ALTER TABLE hukdis
ADD CONSTRAINT chk_hukdis_tanggal_selesai CHECK (
    tanggal_selesai IS NULL OR tanggal_selesai >= tanggal_mulai
);

-- ============================================
-- Constraints untuk DIKLAT
-- ============================================

-- Constraint: Jam JPL tidak boleh negatif
ALTER TABLE diklat
DROP CONSTRAINT IF EXISTS chk_diklat_jam_jpl;

ALTER TABLE diklat
ADD CONSTRAINT chk_diklat_jam_jpl CHECK (jam_jpl IS NULL OR jam_jpl >= 0);

-- Constraint: Tanggal mulai tidak boleh di masa depan (jika diisi)
ALTER TABLE diklat
DROP CONSTRAINT IF EXISTS chk_diklat_tanggal_mulai;

ALTER TABLE diklat
ADD CONSTRAINT chk_diklat_tanggal_mulai CHECK (tanggal_mulai IS NULL OR tanggal_mulai <= CURRENT_DATE);

-- Constraint: Tanggal selesai harus setelah tanggal mulai (jika keduanya diisi)
ALTER TABLE diklat
DROP CONSTRAINT IF EXISTS chk_diklat_tanggal_range;

ALTER TABLE diklat
ADD CONSTRAINT chk_diklat_tanggal_range CHECK (
    tanggal_selesai IS NULL OR
    tanggal_mulai IS NULL OR
    tanggal_selesai >= tanggal_mulai
);

-- ============================================
-- Comments untuk dokumentasi
-- ============================================

COMMENT ON CONSTRAINT chk_nip_format ON pegawai IS 'NIP harus 18 digit angka';
COMMENT ON CONSTRAINT chk_nik_format ON pegawai IS 'NIK harus 16 digit angka jika diisi';
COMMENT ON CONSTRAINT chk_tanggal_lahir ON pegawai IS 'Tanggal lahir tidak boleh di masa depan';
COMMENT ON CONSTRAINT chk_email_format ON pegawai IS 'Email harus dalam format yang valid';
COMMENT ON CONSTRAINT chk_status_keluarga ON keluarga IS 'Status keluarga harus salah satu dari: Suami, Istri, Anak, Ayah, Ibu, Saudara';
