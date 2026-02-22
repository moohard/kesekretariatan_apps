-- ============================================
-- 02_create_tables_kepegawaian.sql
-- ============================================
-- Script untuk membuat tabel di database db_kepegawaian
-- ============================================

\c db_kepegawaian;

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- TABEL PEGAWAI
-- ============================================

CREATE TABLE IF NOT EXISTS pegawai (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nip VARCHAR(18) UNIQUE NOT NULL, -- NIP 18 digit
    nama VARCHAR(255) NOT NULL,
    gelar_depan VARCHAR(50),
    gelar_belakang VARCHAR(50),
    tempat_lahir VARCHAR(100),
    tanggal_lahir DATE,
    jenis_kelamin VARCHAR(10) NOT NULL, -- L/P
    agama_id UUID NOT NULL,
    status_kawin_id UUID NOT NULL,
    nik VARCHAR(16), -- NIK 16 digit
    email VARCHAR(100),
    telepon VARCHAR(20),
    alamat TEXT,
    foto VARCHAR(255), -- path ke file foto
    satker_id UUID NOT NULL,
    jabatan_id UUID,
    unit_kerja_id UUID,
    golongan_id UUID,
    status_pegawai VARCHAR(20) DEFAULT 'aktif', -- aktif, pensiun, mutasi
    tmt_jabatan DATE,
    is_pns BOOLEAN DEFAULT true,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    -- Foreign keys ke db_master
    CONSTRAINT fk_pegawai_agama FOREIGN KEY (agama_id) REFERENCES db_master.ref_agama(id),
    CONSTRAINT fk_pegawai_status_kawin FOREIGN KEY (status_kawin_id) REFERENCES db_master.ref_status_kawin(id),
    CONSTRAINT fk_pegawai_satker FOREIGN KEY (satker_id) REFERENCES db_master.satker(id),
    CONSTRAINT fk_pegawai_jabatan FOREIGN KEY (jabatan_id) REFERENCES db_master.jabatan(id),
    CONSTRAINT fk_pegawai_unit_kerja FOREIGN KEY (unit_kerja_id) REFERENCES db_master.unit_kerja(id),
    CONSTRAINT fk_pegawai_golongan FOREIGN KEY (golongan_id) REFERENCES db_master.golongan(id)
);

-- Index untuk pegawai
CREATE INDEX idx_pegawai_nip ON pegawai(nip);
CREATE INDEX idx_pegawai_nama ON pegawai(nama);
CREATE INDEX idx_pegawai_satker ON pegawai(satker_id);
CREATE INDEX idx_pegawai_jabatan ON pegawai(jabatan_id);
CREATE INDEX idx_pegawai_golongan ON pegawai(golongan_id);
CREATE INDEX idx_pegawai_status ON pegawai(status_pegawai);
CREATE INDEX idx_pegawai_nik ON pegawai(nik);
CREATE INDEX idx_pegawai_email ON pegawai(email);

-- ============================================
-- TABEL RIWAYAT PANGKAT
-- ============================================

CREATE TABLE IF NOT EXISTS riwayat_pangkat (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pegawai_id UUID NOT NULL REFERENCES pegawai(id) ON DELETE CASCADE,
    golongan_id UUID NOT NULL REFERENCES db_master.golongan(id),
    pangkat VARCHAR(100) NOT NULL, -- misal: Pengatur Muda, Penata Tingkat I, dll
    tmt DATE NOT NULL, -- Terhitung Mulai Tanggal
    nomor_sk VARCHAR(100) NOT NULL,
    tanggal_sk DATE NOT NULL,
    pejabat VARCHAR(255) NOT NULL,
    file_sk VARCHAR(255), -- path ke file SK
    gaji_pokok DECIMAL(15, 2) DEFAULT 0,
    is_terakhir BOOLEAN DEFAULT false, -- flag untuk pangkat terakhir
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk riwayat_pangkat
CREATE INDEX idx_riwayat_pangkat_pegawai ON riwayat_pangkat(pegawai_id);
CREATE INDEX idx_riwayat_pangkat_golongan ON riwayat_pangkat(golongan_id);
CREATE INDEX idx_riwayat_pangkat_tmt ON riwayat_pangkat(tmt DESC);
CREATE INDEX idx_riwayat_pangkat_terakhir ON riwayat_pangkat(pegawai_id, is_terakhir) WHERE is_terakhir = true;

-- ============================================
-- TABEL RIWAYAT JABATAN
-- ============================================

CREATE TABLE IF NOT EXISTS riwayat_jabatan (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pegawai_id UUID NOT NULL REFERENCES pegawai(id) ON DELETE CASCADE,
    jabatan_id UUID REFERENCES db_master.jabatan(id),
    unit_kerja_id UUID REFERENCES db_master.unit_kerja(id),
    satker_id UUID REFERENCES db_master.satker(id),
    nama_jabatan VARCHAR(255) NOT NULL, -- nama jabatan aktual
    tmt DATE NOT NULL,
    nomor_sk VARCHAR(100) NOT NULL,
    tanggal_sk DATE NOT NULL,
    pejabat VARCHAR(255) NOT NULL,
    file_sk VARCHAR(255),
    is_terakhir BOOLEAN DEFAULT false, -- flag untuk jabatan terakhir
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk riwayat_jabatan
CREATE INDEX idx_riwayat_jabatan_pegawai ON riwayat_jabatan(pegawai_id);
CREATE INDEX idx_riwayat_jabatan_jabatan ON riwayat_jabatan(jabatan_id);
CREATE INDEX idx_riwayat_jabatan_tmt ON riwayat_jabatan(tmt DESC);
CREATE INDEX idx_riwayat_jabatan_terakhir ON riwayat_jabatan(pegawai_id, is_terakhir) WHERE is_terakhir = true;

-- ============================================
-- TABEL RIWAYAT PENDIDIKAN
-- ============================================

CREATE TABLE IF NOT EXISTS riwayat_pendidikan (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pegawai_id UUID NOT NULL REFERENCES pegawai(id) ON DELETE CASCADE,
    pendidikan_id UUID NOT NULL REFERENCES db_master.ref_pendidikan(id),
    nama_institusi VARCHAR(255) NOT NULL,
    jurusan VARCHAR(255),
    tahun_masuk INTEGER,
    tahun_lulus INTEGER,
    nomor_ijazah VARCHAR(100),
    tanggal_ijazah DATE,
    file_ijazah VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk riwayat_pendidikan
CREATE INDEX idx_riwayat_pendidikan_pegawai ON riwayat_pendidikan(pegawai_id);
CREATE INDEX idx_riwayat_pendidikan_pendidikan ON riwayat_pendidikan(pendidikan_id);
CREATE INDEX idx_riwayat_pendidikan_tahun ON riwayat_pendidikan(tahun_lulus DESC);

-- ============================================
-- TABEL KELUARGA
-- ============================================

CREATE TABLE IF NOT EXISTS keluarga (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pegawai_id UUID NOT NULL REFERENCES pegawai(id) ON DELETE CASCADE,
    status_keluarga VARCHAR(20) NOT NULL, -- suami, istri, anak
    nama VARCHAR(255) NOT NULL,
    tempat_lahir VARCHAR(100),
    tanggal_lahir DATE,
    jenis_kelamin VARCHAR(10),
    nik VARCHAR(16),
    pendidikan VARCHAR(100),
    pekerjaan VARCHAR(100),
    is_tanggungan BOOLEAN DEFAULT false, -- untuk tunjangan keluarga
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk keluarga
CREATE INDEX idx_keluarga_pegawai ON keluarga(pegawai_id);
CREATE INDEX idx_keluarga_status ON keluarga(pegawai_id, status_keluarga);
CREATE INDEX idx_keluarga_nik ON keluarga(nik);
CREATE INDEX idx_keluarga_tanggungan ON keluarga(pegawai_id, is_tanggungan) WHERE is_tanggungan = true;

-- ============================================
-- TABEL TEMPLATE DOKUMEN
-- ============================================

CREATE TABLE IF NOT EXISTS template_dokumen (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kode VARCHAR(50) UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    tipe VARCHAR(50) NOT NULL, -- sk_pangkat, sk_jabatan, dll
    konten_html TEXT NOT NULL, -- template HTML untuk PDF
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk template_dokumen
CREATE INDEX idx_template_kode ON template_dokumen(kode);
CREATE INDEX idx_template_tipe ON template_dokumen(tipe);

-- ============================================
-- TABEL HUKDIS (Hukuman Disiplin)
-- ============================================

CREATE TABLE IF NOT EXISTS hukdis (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pegawai_id UUID NOT NULL REFERENCES pegawai(id) ON DELETE CASCADE,
    jenis_hukdis_id UUID NOT NULL REFERENCES db_master.ref_jenis_hukdis(id),
    nomor_sk VARCHAR(100) NOT NULL,
    tanggal_sk DATE NOT NULL,
    tanggal_mulai DATE NOT NULL,
    tanggal_selesai DATE,
    alasan TEXT,
    pejabat VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk hukdis
CREATE INDEX idx_hukdis_pegawai ON hukdis(pegawai_id);
CREATE INDEX idx_hukdis_jenis ON hukdis(jenis_hukdis_id);
CREATE INDEX idx_hukdis_tanggal ON hukdis(tanggal_mulai DESC);

-- ============================================
-- TABEL DIKLAT
-- ============================================

CREATE TABLE IF NOT EXISTS diklat (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pegawai_id UUID NOT NULL REFERENCES pegawai(id) ON DELETE CASCADE,
    jenis_diklat_id UUID NOT NULL REFERENCES db_master.ref_jenis_diklat(id),
    nama_diklat VARCHAR(255) NOT NULL,
    penyelenggara VARCHAR(255),
    tempat VARCHAR(255),
    tanggal_mulai DATE,
    tanggal_selesai DATE,
    jam_jpl INTEGER, -- Jam Pelatihan
    nomor_sertifikat VARCHAR(100),
    tanggal_sertifikat DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index untuk diklat
CREATE INDEX idx_diklat_pegawai ON diklat(pegawai_id);
CREATE INDEX idx_diklat_jenis ON diklat(jenis_diklat_id);
CREATE INDEX idx_diklat_tanggal ON diklat(tanggal_mulai DESC);

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
CREATE TRIGGER update_pegawai_updated_at BEFORE UPDATE ON pegawai FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_riwayat_pangkat_updated_at BEFORE UPDATE ON riwayat_pangkat FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_riwayat_jabatan_updated_at BEFORE UPDATE ON riwayat_jabatan FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_riwayat_pendidikan_updated_at BEFORE UPDATE ON riwayat_pendidikan FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_keluarga_updated_at BEFORE UPDATE ON keluarga FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_template_dokumen_updated_at BEFORE UPDATE ON template_dokumen FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_hukdis_updated_at BEFORE UPDATE ON hukdis FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_diklat_updated_at BEFORE UPDATE ON diklat FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();