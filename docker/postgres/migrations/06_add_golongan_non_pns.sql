-- ============================================================================
-- MIGRATION: Add Golongan Non-PNS
-- Version: 06
-- Date: 2026-02-23
-- Description: Menambahkan tabel dan data untuk golongan pegawai non-PNS
-- ============================================================================

\c db_master;

-- ============================================================================
-- 1. BUAT TABEL REF_GOLONGAN_NON_PNS
-- ============================================================================

CREATE TABLE IF NOT EXISTS ref_golongan_non_pns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    kategori VARCHAR(50) NOT NULL,
    urutan INT NOT NULL,
    keterangan TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Index
CREATE INDEX idx_golongan_non_pns_kode ON ref_golongan_non_pns(kode);
CREATE INDEX idx_golongan_non_pns_kategori ON ref_golongan_non_pns(kategori);

-- Trigger untuk updated_at
CREATE OR REPLACE FUNCTION update_golongan_non_pns_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_golongan_non_pns_updated_at
    BEFORE UPDATE ON ref_golongan_non_pns
    FOR EACH ROW
    EXECUTE FUNCTION update_golongan_non_pns_updated_at();

-- ============================================================================
-- 2. SEED DATA GOLONGAN NON-PNS
-- ============================================================================

INSERT INTO ref_golongan_non_pns (kode, nama, kategori, urutan, keterangan) VALUES
-- Honorer K1 (Golongan I)
('I', 'Honorer K1', 'Honorer', 1, 'Tenaga Honorer Kategori 1 - Sudah memenuhi syarat untuk diangkat menjadi PNS'),

-- Honorer K2 (Golongan V)
('V', 'Honorer K2', 'Honorer', 2, 'Tenaga Honorer Kategori 2 - Belum memenuhi syarat untuk diangkat menjadi PNS'),

-- Tenaga Kontrak/PPPK (Golongan IX)
('IX', 'Tenaga Kontrak', 'Kontrak', 3, 'Tenaga Kontrak atau PPPK (Pegawai Pemerintah dengan Perjanjian Kerja)'),

-- Variasi lain yang mungkin ada
('TK', 'Tenaga Kontrak', 'Kontrak', 4, 'Tenaga Kontrak umum'),
('TH', 'Tenaga Harian', 'Honorer', 5, 'Tenaga Harian Lepas')
ON CONFLICT (kode) DO UPDATE SET
    nama = EXCLUDED.nama,
    kategori = EXCLUDED.kategori,
    urutan = EXCLUDED.urutan,
    keterangan = EXCLUDED.keterangan;

-- ============================================================================
-- 3. BUAT VIEW UNTUK MENGGABUNGKAN GOLONGAN PNS DAN NON-PNS
-- ============================================================================

CREATE OR REPLACE VIEW v_golongan_all AS
SELECT
    id,
    kode,
    nama,
    ruang,
    angka,
    'PNS' as kategori,
    min_pangkat,
    max_pangkat
FROM golongan
WHERE is_active = true

UNION ALL

SELECT
    id,
    kode,
    nama,
    NULL as ruang,
    urutan as angka,
    kategori,
    NULL as min_pangkat,
    NULL as max_pangkat
FROM ref_golongan_non_pns
WHERE is_active = true
ORDER BY kategori, angka;

-- ============================================================================
-- 4. KOMENTAR UNTUK DOKUMENTASI
-- ============================================================================

COMMENT ON TABLE ref_golongan_non_pns IS 'Referensi golongan untuk pegawai non-PNS (Honorer, Kontrak, PPPK)';
COMMENT ON COLUMN ref_golongan_non_pns.kode IS 'Kode golongan (I, V, IX, TK, TH)';
COMMENT ON COLUMN ref_golongan_non_pns.kategori IS 'Kategori: Honorer, Kontrak';
COMMENT ON VIEW v_golongan_all IS 'View yang menggabungkan golongan PNS dan Non-PNS untuk keperluan dropdown';

-- ============================================================================
-- SELESAI
-- ============================================================================
