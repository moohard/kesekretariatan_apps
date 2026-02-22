package testutil

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

// TestFixture holds test data fixtures for RLS testing
type TestFixture struct {
	t         *testing.T
	db        *TestDB
	UserIDs   []string
	UnitIDs   []string
	SatkerIDs []string
}

// NewTestFixture creates a new test fixture
func NewTestFixture(t *testing.T, db *TestDB) *TestFixture {
	return &TestFixture{
		t:  t,
		db: db,
	}
}

// SetupRLSData creates test data for RLS testing
// Creates: 2 Satkers, 2 Unit Kerja per Satker, 2 Pegawai per Unit Kerja
func (f *TestFixture) SetupRLSData(ctx context.Context) {
	now := time.Now()

	// Create Satkers
	satker1 := uuid.New().String()
	satker2 := uuid.New().String()
	f.SatkerIDs = []string{satker1, satker2}

	_, err := f.db.Pool.Exec(ctx, `
		INSERT INTO satker (id, kode, nama, jenis, tingkat, created_at, updated_at)
		VALUES
			($1, 'PA001', 'Pengadilan Agama Jakarta Pusat', 'PA', 'PN', $3, $3),
			($2, 'PA002', 'Pengadilan Agama Jakarta Selatan', 'PA', 'PN', $3, $3)
	`, satker1, satker2, now)
	if err != nil {
		f.t.Fatalf("failed to create satkers: %v", err)
	}

	// Create Unit Kerja per Satker
	unit1 := uuid.New().String() // Unit Satker 1
	unit2 := uuid.New().String() // Unit Satker 1
	unit3 := uuid.New().String() // Unit Satker 2
	unit4 := uuid.New().String() // Unit Satker 2
	f.UnitIDs = []string{unit1, unit2, unit3, unit4}

	_, err = f.db.Pool.Exec(ctx, `
		INSERT INTO unit_kerja (id, satker_id, nama, created_at, updated_at)
		VALUES
			($1, $5, 'Kesekretariatan', $9, $9),
			($2, $5, 'Kepaniteraan', $9, $9),
			($3, $6, 'Kesekretariatan', $9, $9),
			($4, $6, 'Kepaniteraan', $9, $9)
	`, unit1, unit2, unit3, unit4, satker1, satker2, now, now, now)
	if err != nil {
		f.t.Fatalf("failed to create unit_kerja: %v", err)
	}

	// Create Pegawai per Unit Kerja (2 per unit = 8 total)
	users := []struct {
		id        string
		unitID    string
		nip       string
		nama      string
	}{
		{uuid.New().String(), unit1, "198501012010011001", "Ahmad Fauzi"},
		{uuid.New().String(), unit1, "198702152012012002", "Budi Santoso"},
		{uuid.New().String(), unit2, "199001032013011003", "Citra Dewi"},
		{uuid.New().String(), unit2, "198803102014012004", "Diana Putri"},
		{uuid.New().String(), unit3, "198704122011011005", "Eko Prasetyo"},
		{uuid.New().String(), unit3, "198905202012012006", "Fitri Handayani"},
		{uuid.New().String(), unit4, "199106252013011007", "Gunawan Wibowo"},
		{uuid.New().String(), unit4, "198807302014012008", "Hesti Rahayu"},
	}

	for _, u := range users {
		f.UserIDs = append(f.UserIDs, u.id)
		_, err := f.db.Pool.Exec(ctx, `
			INSERT INTO pegawai (id, unit_kerja_id, nip, nama_lengkap, tanggal_lahir, jenis_kelamin, agama, status_pegawai, is_active, created_at, updated_at)
			VALUES ($1, $2, $3, $4, '1985-01-01', 'L', 'Islam', 'PNS', true, $5, $5)
		`, u.id, u.unitID, u.nip, u.nama, now)
		if err != nil {
			f.t.Fatalf("failed to create pegawai %s: %v", u.nama, err)
		}
	}
}

// Cleanup removes all test data
func (f *TestFixture) Cleanup(ctx context.Context) {
	// Delete in reverse order of dependencies
	f.db.Pool.Exec(ctx, "DELETE FROM pegawai WHERE id = ANY($1)", f.UserIDs)
	f.db.Pool.Exec(ctx, "DELETE FROM unit_kerja WHERE id = ANY($1)", f.UnitIDs)
	f.db.Pool.Exec(ctx, "DELETE FROM satker WHERE id = ANY($1)", f.SatkerIDs)
}

// AssertRowCount asserts the number of rows accessible with current RLS context
func (f *TestFixture) AssertRowCount(ctx context.Context, table string, expected int, msgAndArgs ...any) {
	count, err := f.db.CountRows(ctx, table, "")
	if err != nil {
		f.t.Fatalf("failed to count rows: %v", err)
	}

	if count != expected {
		msg := fmt.Sprintf("expected %d rows in %s, got %d", expected, table, count)
		if len(msgAndArgs) > 0 {
			msg = fmt.Sprintf("%s: %v", msg, msgAndArgs[0])
		}
		f.t.Error(msg)
	}
}
