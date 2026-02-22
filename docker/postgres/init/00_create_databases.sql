-- ============================================
-- 00_create_databases.sql
-- ============================================
-- Script untuk membuat database master dan kepegawaian
-- ============================================

-- Membuat database db_master
-- Database ini menyimpan:
--   - Tabel referensi (satker, jabatan, golongan, dll)
--   - Tabel RBAC (roles, permissions, user roles)
--   - Tabel audit logs
CREATE DATABASE db_master
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
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
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE db_master TO postgres;
GRANT ALL PRIVILEGES ON DATABASE db_kepegawaian TO postgres;