-- ============================================================================
-- MIGRATION: Seed Jabatan Lengkap
-- Version: 07
-- Date: 2026-02-23
-- Description: Melengkapi data jabatan yang kurang berdasarkan data pegawai aktual
-- ============================================================================

\c db_master;

-- ============================================================================
-- 1. SEED JABATAN YANG KURANG
-- ============================================================================

-- Pastikan kolom jenis ada (jika belum)
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'jabatan' AND column_name = 'jenis'
    ) THEN
        ALTER TABLE jabatan ADD COLUMN jenis VARCHAR(30)
            CHECK (jenis IN ('struktural', 'fungsional_tertentu', 'fungsional_umum', 'pelaksana'));
    END IF;
END $$;

-- Insert jabatan yang kurang
INSERT INTO jabatan (id, kode, nama, eselon_id, kelas, jenis, is_active, created_at, updated_at) VALUES

-- ============================================================================
-- JABATAN STRUKTURAL KEPAJITERAAAN (Lanjutan)
-- ============================================================================

-- Panitera Muda per bidang (Eselon IV.a)
(
    gen_random_uuid(),
    'PM-PERMOHONAN',
    'Panitera Muda Permohonan',
    (SELECT id FROM eselon WHERE kode = 'IV' LIMIT 1),
    '2A',
    'struktural',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'PM-GUGATAN',
    'Panitera Muda Gugatan',
    (SELECT id FROM eselon WHERE kode = 'IV' LIMIT 1),
    '2A',
    'struktural',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'PM-HUKUM',
    'Panitera Muda Hukum',
    (SELECT id FROM eselon WHERE kode = 'IV' LIMIT 1),
    '2A',
    'struktural',
    true,
    NOW(),
    NOW()
),

-- ============================================================================
-- JABATAN PELAKSANA KEPAJITERAAAN
-- ============================================================================

(
    gen_random_uuid(),
    'PP-TP',
    'Panitera Pengganti Tingkat Pertama',
    NULL,
    '2B',
    'pelaksana',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'JS-PG',
    'Juru Sita Pengganti',
    NULL,
    '2D',
    'pelaksana',
    true,
    NOW(),
    NOW()
),

-- Klerek / Staf Kepaniteraan
(
    gen_random_uuid(),
    'KLEREK-AP',
    'Klerek - Analis Perkara Peradilan',
    NULL,
    '3C',
    'pelaksana',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'KLEREK-PP',
    'Klerek - Pengelola Penanganan Perkara',
    NULL,
    '3C',
    'pelaksana',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'KLEREK-DH',
    'Klerek - Dokumentalis Hukum',
    NULL,
    '3C',
    'pelaksana',
    true,
    NOW(),
    NOW()
),

-- ============================================================================
-- JABATAN FUNGSIONAL TEKNIK
-- ============================================================================

(
    gen_random_uuid(),
    'FUNG-PKOM-A1',
    'Pranata Komputer Ahli Pertama',
    NULL,
    '3A',
    'fungsional_tertentu',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'FUNG-PKOM-AM',
    'Pranata Komputer Ahli Muda',
    NULL,
    '3A',
    'fungsional_tertentu',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'FUNG-PKOM-AK',
    'Pranata Komputer Ahli Madya',
    NULL,
    '2A',
    'fungsional_tertentu',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'TEKNIS-SARPRAS',
    'Teknisi Sarana dan Prasarana',
    NULL,
    '3A',
    'fungsional_umum',
    true,
    NOW(),
    NOW()
),

-- ============================================================================
-- JABATAN PELAKSANA KESEKRETARIATAN
-- ============================================================================

(
    gen_random_uuid(),
    'OP-PLO-9',
    'Operator - Penata Layanan Operasional',
    NULL,
    '9',
    'pelaksana',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'OP-LO-5',
    'Operator Layanan Operasional',
    NULL,
    '5',
    'pelaksana',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'PUO-1',
    'Pengelola Umum Operasional',
    NULL,
    '1',
    'pelaksana',
    true,
    NOW(),
    NOW()
),

-- ============================================================================
-- JABATAN LAIN-LAIN
-- ============================================================================

(
    gen_random_uuid(),
    'ARSIPARIS',
    'Arsiparis',
    NULL,
    '4A',
    'fungsional_umum',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'BENDAHARA',
    'Bendahara Pengeluaran',
    NULL,
    '3A',
    'fungsional_umum',
    true,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'PEMBUKU',
    'Pembukuan',
    NULL,
    '4A',
    'pelaksana',
    true,
    NOW(),
    NOW()
)

ON CONFLICT (kode) DO UPDATE SET
    nama = EXCLUDED.nama,
    eselon_id = EXCLUDED.eselon_id,
    kelas = EXCLUDED.kelas,
    jenis = EXCLUDED.jenis,
    updated_at = NOW();

-- ============================================================================
-- 2. UPDATE JABATAN YANG SUDAH ADA (TAMBAH JENIS)
-- ============================================================================

UPDATE jabatan SET jenis = 'struktural' WHERE eselon_id IS NOT NULL AND jenis IS NULL;
UPDATE jabatan SET jenis = 'pelaksana' WHERE eselon_id IS NULL AND jenis IS NULL AND nama ILIKE '%operator%';
UPDATE jabatan SET jenis = 'pelaksana' WHERE eselon_id IS NULL AND jenis IS NULL AND nama ILIKE '%klerek%';
UPDATE jabatan SET jenis = 'fungsional_tertentu' WHERE eselon_id IS NULL AND jenis IS NULL AND (nama ILIKE '%hakim%' OR nama ILIKE '%pranata%');

-- ============================================================================
-- 3. VERIFIKASI DATA
-- ============================================================================

-- Tampilkan ringkasan jabatan per jenis
SELECT
    COALESCE(jenis, 'tanpa_jenis') as jenis_jabatan,
    COUNT(*) as jumlah
FROM jabatan
WHERE is_active = true
GROUP BY jenis
ORDER BY jumlah DESC;

-- ============================================================================
-- SELESAI
-- ============================================================================
