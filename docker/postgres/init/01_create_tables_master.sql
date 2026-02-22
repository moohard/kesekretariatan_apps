-- ============================================
-- 01_create_tables_master.sql
-- ============================================
-- Script untuk membuat tabel di database db_master
-- ============================================

\c db_master;

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- TABEL REFERENSI
-- ============================================

-- Tabel Satker (Satuan Kerja)
CREATE TABLE IF NOT EXISTS satker (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    parent_id UUID REFERENCES satker(id) ON DELETE SET NULL,
    level INTEGER NOT NULL DEFAULT 1,
    alamat TEXT,
    telepon VARCHAR(50),
    email VARCHAR(100),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_by UUID,
    updated_by UUID
);

-- Index untuk satker
CREATE INDEX idx_satker_kode ON satker(kode);
CREATE INDEX idx_satker_parent ON satker(parent_id);
CREATE INDEX idx_satker_level ON satker(level);

-- Tabel Jabatan
CREATE TABLE IF NOT EXISTS jabatan (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    eselon_id UUID REFERENCES eselon(id) ON DELETE SET NULL,
    kelas VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk jabatan
CREATE INDEX idx_jabatan_kode ON jabatan(kode);
CREATE INDEX idx_jabatan_eselon ON jabatan(eselon_id);

-- Tabel Golongan
CREATE TABLE IF NOT EXISTS golongan (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    ruang VARCHAR(10) NOT NULL,
    angka INTEGER NOT NULL,
    min_pangkat INTEGER DEFAULT 1,
    max_pangkat INTEGER DEFAULT 27,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk golongan
CREATE INDEX idx_golongan_kode ON golongan(kode);
CREATE INDEX idx_golongan_angka ON golongan(angka);

-- Tabel Unit Kerja
CREATE TABLE IF NOT EXISTS unit_kerja (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    singkatan VARCHAR(50),
    parent_id UUID REFERENCES unit_kerja(id) ON DELETE SET NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk unit_kerja
CREATE INDEX idx_unit_kerja_kode ON unit_kerja(kode);
CREATE INDEX idx_unit_kerja_parent ON unit_kerja(parent_id);

-- Tabel Eselon
CREATE TABLE IF NOT EXISTS eselon (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    tunjangan DECIMAL(15, 2) DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk eselon
CREATE INDEX idx_eselon_kode ON eselon(kode);

-- Tabel Referensi Pendidikan
CREATE TABLE IF NOT EXISTS ref_pendidikan (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    tingkat VARCHAR(50) NOT NULL, -- SD, SMP, SMA, D3, S1, S2, S3
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabel Referensi Agama
CREATE TABLE IF NOT EXISTS ref_agama (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabel Referensi Status Kawin
CREATE TABLE IF NOT EXISTS ref_status_kawin (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabel Referensi Jenis Hukdis
CREATE TABLE IF NOT EXISTS ref_jenis_hukdis (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabel Referensi Jenis Diklat
CREATE TABLE IF NOT EXISTS ref_jenis_diklat (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(10) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- ============================================
-- TABEL RBAC
-- ============================================

-- Tabel App Roles
CREATE TABLE IF NOT EXISTS app_roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nama VARCHAR(100) UNIQUE NOT NULL,
    deskripsi TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabel App Permissions
CREATE TABLE IF NOT EXISTS app_permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nama VARCHAR(100) UNIQUE NOT NULL,
    resource VARCHAR(100) NOT NULL, -- master_data, kepegawaian, rbac, audit
    action VARCHAR(50) NOT NULL,    -- read, create, update, delete
    deskripsi TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk permissions
CREATE INDEX idx_permissions_resource ON app_permissions(resource);
CREATE UNIQUE INDEX idx_permissions_unique ON app_permissions(resource, action);

-- Tabel Role Permissions
CREATE TABLE IF NOT EXISTS role_permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id UUID NOT NULL REFERENCES app_roles(id) ON DELETE CASCADE,
    permission_id UUID NOT NULL REFERENCES app_permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(role_id, permission_id)
);

-- Index untuk role_permissions
CREATE INDEX idx_role_permissions_role ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission ON role_permissions(permission_id);

-- Tabel User App Roles
CREATE TABLE IF NOT EXISTS user_app_roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id VARCHAR(255) NOT NULL, -- Keycloak user ID
    role_id UUID NOT NULL REFERENCES app_roles(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, role_id)
);

-- Index untuk user_app_roles
CREATE INDEX idx_user_app_roles_user ON user_app_roles(user_id);
CREATE INDEX idx_user_app_roles_role ON user_app_roles(role_id);

-- ============================================
-- TABEL AUDIT
-- ============================================

-- Tabel Audit Logs
CREATE TABLE IF NOT EXISTS audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id VARCHAR(255), -- Keycloak user ID
    username VARCHAR(100),
    action VARCHAR(50) NOT NULL, -- create, read, update, delete
    resource VARCHAR(100) NOT NULL, -- nama tabel atau resource
    resource_id UUID,
    ip_address INET,
    user_agent TEXT,
    changes JSONB, -- perubahan data (before/after)
    status VARCHAR(20) NOT NULL DEFAULT 'success', -- success, failed
    error_message TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk audit_logs
CREATE INDEX idx_audit_logs_user ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_resource ON audit_logs(resource);
CREATE INDEX idx_audit_logs_created ON audit_logs(created_at DESC);

-- ============================================
-- TRIGGER untuk updated_at
-- ============================================

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply trigger ke tabel yang memiliki updated_at
CREATE TRIGGER update_satker_updated_at BEFORE UPDATE ON satker FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_jabatan_updated_at BEFORE UPDATE ON jabatan FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_golongan_updated_at BEFORE UPDATE ON golongan FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_unit_kerja_updated_at BEFORE UPDATE ON unit_kerja FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_eselon_updated_at BEFORE UPDATE ON eselon FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_ref_pendidikan_updated_at BEFORE UPDATE ON ref_pendidikan FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_ref_agama_updated_at BEFORE UPDATE ON ref_agama FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_ref_status_kawin_updated_at BEFORE UPDATE ON ref_status_kawin FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_ref_jenis_hukdis_updated_at BEFORE UPDATE ON ref_jenis_hukdis FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_ref_jenis_diklat_updated_at BEFORE UPDATE ON ref_jenis_diklat FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_app_roles_updated_at BEFORE UPDATE ON app_roles FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_app_permissions_updated_at BEFORE UPDATE ON app_permissions FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();