-- ============================================================================
-- MIGRATION: Fix Pegawai Schema
-- Version: 05
-- Date: 2026-02-23
-- Description: Menambahkan field yang kurang di tabel pegawai berdasarkan Gap Analysis
-- ============================================================================

\c db_kepegawaian;

-- ============================================================================
-- 1. TAMBAH FIELD YANG KURANG DI TABEL PEGAWAI
-- ============================================================================

-- Field eselon (P0 - Kritis)
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS eselon_id UUID;

-- Field TMT (P0 - Kritis)
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS tmt_cpns DATE;
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS tmt_pns DATE;
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS tmt_pangkat_terakhir DATE;
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS tmt_jabatan_terakhir DATE;

-- Field status kerja (P0 - Kritis)
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS status_kerja VARCHAR(20) DEFAULT 'aktif';

-- Field dokumen kepegawaian (P2)
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS nip_lama VARCHAR(9);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS karpeg_no VARCHAR(50);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS karpeg_file VARCHAR(500);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS taspen_no VARCHAR(50);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS npwp VARCHAR(20);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS bpjs_kesehatan VARCHAR(30);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS bpjs_ketenagakerjaan VARCHAR(30);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS kk_no VARCHAR(30);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS kk_file VARCHAR(500);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS ktp_no VARCHAR(30);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS ktp_file VARCHAR(500);
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS alamat_domisili TEXT;

-- Field audit trail (P1)
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS created_by UUID;
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS updated_by UUID;
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ;
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS deleted_by UUID;

-- Field integrasi (P3)
ALTER TABLE pegawai ADD COLUMN IF NOT EXISTS sikep_id VARCHAR(50);

-- ============================================================================
-- 2. PERBAIKI CONSTRAINT STATUS_PEGAWAI
-- ============================================================================

-- Hapus constraint lama jika ada
ALTER TABLE pegawai DROP CONSTRAINT IF EXISTS pegawai_status_pegawai_check;

-- Tambah constraint baru dengan enum yang benar
ALTER TABLE pegawai ADD CONSTRAINT pegawai_status_pegawai_check
    CHECK (status_pegawai IN ('PNS', 'CPNS', 'PPPK', 'HONORER'));

-- Tambah constraint untuk status_kerja
ALTER TABLE pegawai ADD CONSTRAINT pegawai_status_kerja_check
    CHECK (status_kerja IN ('aktif', 'cuti', 'pensiun', 'mutasi_keluar', 'mutasi_masuk', 'meninggal', 'pemberhentian'));

-- ============================================================================
-- 3. TAMBAH INDEX UNTUK FIELD BARU
-- ============================================================================

CREATE INDEX IF NOT EXISTS idx_pegawai_eselon ON pegawai(eselon_id);
CREATE INDEX IF NOT EXISTS idx_pegawai_status_kerja ON pegawai(status_kerja);
CREATE INDEX IF NOT EXISTS idx_pegawai_sikep ON pegawai(sikep_id);
CREATE INDEX IF NOT EXISTS idx_pegawai_deleted ON pegawai(deleted_at) WHERE deleted_at IS NOT NULL;

-- ============================================================================
-- 4. UPDATE DATA EXISTING (MIGRASI DATA)
-- ============================================================================

-- Update status_pegawai berdasarkan golongan
-- Golongan standar (I/a - IV/e) = PNS
-- Golongan IX = HONORER
-- Golongan I, V tanpa angka = juga non-PNS, tapi perlu cek manual

-- Catatan: Query ini perlu disesuaikan dengan data aktual
-- UPDATE pegawai SET status_pegawai = 'PNS' WHERE status_pegawai = 'aktif' AND golongan_id IS NOT NULL;

-- Set default status_kerja untuk data existing
UPDATE pegawai SET status_kerja = 'aktif' WHERE status_kerja IS NULL;

-- ============================================================================
-- 5. TAMBAH FIELD DI TABEL RIWAYAT
-- ============================================================================

-- Riwayat Pangkat
ALTER TABLE riwayat_pangkat ADD COLUMN IF NOT EXISTS jenis_kenaikan VARCHAR(50)
    CHECK (jenis_kenaikan IN ('reguler', 'pilihan', 'penyesuaian_ijazah', 'lainnya'));
ALTER TABLE riwayat_pangkat ADD COLUMN IF NOT EXISTS masa_kerja_tahun INT DEFAULT 0;
ALTER TABLE riwayat_pangkat ADD COLUMN IF NOT EXISTS masa_kerja_bulan INT DEFAULT 0;
ALTER TABLE riwayat_pangkat ADD COLUMN IF NOT EXISTS created_by UUID;

-- Riwayat Jabatan
ALTER TABLE riwayat_jabatan ADD COLUMN IF NOT EXISTS jenis_jabatan VARCHAR(20)
    CHECK (jenis_jabatan IN ('struktural', 'fungsional_tertentu', 'fungsional_umum', 'pelaksana'));
ALTER TABLE riwayat_jabatan ADD COLUMN IF NOT EXISTS created_by UUID;

-- Riwayat Pendidikan
ALTER TABLE riwayat_pendidikan ADD COLUMN IF NOT EXISTS created_by UUID;

-- Keluarga
ALTER TABLE keluarga ADD COLUMN IF NOT EXISTS created_by UUID;

-- ============================================================================
-- 6. KOMENTAR UNTUK DOKUMENTASI
-- ============================================================================

COMMENT ON COLUMN pegawai.eselon_id IS 'Referensi ke tabel eselon di db_master untuk pegawai struktural';
COMMENT ON COLUMN pegawai.tmt_cpns IS 'Terhitung Mulai Tanggal sebagai CPNS';
COMMENT ON COLUMN pegawai.tmt_pns IS 'Terhitung Mulai Tanggal sebagai PNS';
COMMENT ON COLUMN pegawai.tmt_pangkat_terakhir IS 'TMT pangkat/golongan terakhir';
COMMENT ON COLUMN pegawai.tmt_jabatan_terakhir IS 'TMT jabatan terakhir';
COMMENT ON COLUMN pegawai.status_kerja IS 'Status kerja: aktif, cuti, pensiun, mutasi_keluar, mutasi_masuk, meninggal, pemberhentian';
COMMENT ON COLUMN pegawai.status_pegawai IS 'Status kepegawaian: PNS, CPNS, PPPK, HONORER';
COMMENT ON COLUMN pegawai.sikep_id IS 'ID pegawai di sistem SIKEP Mahkamah Agung';

COMMENT ON COLUMN riwayat_pangkat.jenis_kenaikan IS 'Jenis kenaikan pangkat: reguler, pilihan, penyesuaian_ijazah, lainnya';
COMMENT ON COLUMN riwayat_jabatan.jenis_jabatan IS 'Jenis jabatan: struktural, fungsional_tertentu, fungsional_umum, pelaksana';

-- ============================================================================
-- SELESAI
-- ============================================================================
