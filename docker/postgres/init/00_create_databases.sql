-- ============================================
-- 00_create_databases.sql
-- ============================================
-- Script untuk membuat database master dan kepegawaian
-- Note: OWNER defaults to the user executing this script
-- ============================================

-- Membuat database db_master
-- Database ini menyimpan:
--   - Tabel referensi (satker, jabatan, golongan, dll)
--   - Tabel RBAC (roles, permissions, user roles)
--   - Tabel audit logs
CREATE DATABASE db_master
    WITH
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Membuat database db_kepegawaian
-- Database ini menyimpan:
--   - Tabel pegawai
--   - Tabel riwayat (pangkat, jabatan, pendidikan)
--   - Tabel keluarga
--   - Tabel template dokumen
CREATE DATABASE db_kepegawaian
    WITH
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
