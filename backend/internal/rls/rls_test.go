// Package rls_test tests Row-Level Security policies for SIKERMA
package rls_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sikerma/backend/internal/testutil"
)

// ============================================================================
// TEST SETUP
// ============================================================================

// RLSTestSuite holds test resources for RLS testing
type RLSTestSuite struct {
	DB      *testutil.TestDB
	Fixture *testutil.TestFixture
}

// SetupRLSTest initializes test resources for RLS tests
func SetupRLSTest(t *testing.T) *RLSTestSuite {
	ctx := context.Background()

	// Start test database container
	db := testutil.SetupTestDB(t, "file://../../docker/postgres/init")

	// Create test fixture
	fixture := testutil.NewTestFixture(t, db)
	fixture.SetupRLSData(ctx)

	return &RLSTestSuite{
		DB:      db,
		Fixture: fixture,
	}
}

// Cleanup releases test resources
func (s *RLSTestSuite) Cleanup(t *testing.T) {
	ctx := context.Background()
	s.Fixture.Cleanup(ctx)
	s.DB.Cleanup(t)
}

// ============================================================================
// TEST 1: current_user_id() Helper Function
// ============================================================================

func TestCurrentUserIDFunction(t *testing.T) {
	suite := SetupRLSTest(t)
	defer suite.Cleanup(t)

	ctx := context.Background()
	testUserID := uuid.New().String()

	t.Run("returns NULL when no JWT claim is set", func(t *testing.T) {
		var result *string
		err := suite.DB.Pool.QueryRow(ctx, "SELECT current_user_id()").Scan(&result)
		require.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("returns user ID when JWT claim is set", func(t *testing.T) {
		// Set JWT claim
		_, err := suite.DB.Pool.Exec(ctx,
			"SET LOCAL request.jwt.claim.user_id = $1", testUserID)
		require.NoError(t, err)

		var result string
		err = suite.DB.Pool.QueryRow(ctx, "SELECT current_user_id()::text").Scan(&result)
		require.NoError(t, err)
		assert.Equal(t, testUserID, result)
	})

	t.Run("returns NULL for invalid UUID format", func(t *testing.T) {
		_, err := suite.DB.Pool.Exec(ctx,
			"SET LOCAL request.jwt.claim.user_id = 'invalid-uuid'")
		require.NoError(t, err)

		var result *string
		err = suite.DB.Pool.QueryRow(ctx, "SELECT current_user_id()").Scan(&result)
		require.NoError(t, err)
		assert.Nil(t, result)
	})
}

// ============================================================================
// TEST 2: RLS Policy - Unit Isolation
// ============================================================================

func TestRLSPolicyUnitIsolation(t *testing.T) {
	suite := SetupRLSTest(t)
	defer suite.Cleanup(t)

	ctx := context.Background()

	// Get test data
	unit1ID := suite.Fixture.UnitIDs[0]
	unit2ID := suite.Fixture.UnitIDs[1]
	userInUnit1 := suite.Fixture.UserIDs[0] // User di unit 1

	t.Run("user can only see pegawai in their unit kerja", func(t *testing.T) {
		// Set user context (user di unit 1)
		err := suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Count rows visible (should only see users in unit 1)
		count, err := suite.DB.CountRows(ctx, "pegawai", "unit_kerja_id = $1", unit1ID)
		require.NoError(t, err)
		assert.Equal(t, 2, count, "User should see 2 pegawai in their unit")

		// Verify can't see pegawai from other unit
		countOther, err := suite.DB.CountRows(ctx, "pegawai", "unit_kerja_id = $1", unit2ID)
		require.NoError(t, err)
		assert.Equal(t, 0, countOther, "User should NOT see pegawai from other unit")
	})

	t.Run("user cannot access specific pegawai from other unit", func(t *testing.T) {
		// Set user context (user di unit 1)
		err := suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Try to select pegawai from other unit
		userInUnit2 := suite.Fixture.UserIDs[2] // User di unit 2

		var exists bool
		err = suite.DB.Pool.QueryRow(ctx,
			"SELECT EXISTS(SELECT 1 FROM pegawai WHERE id = $1)", userInUnit2).
			Scan(&exists)
		require.NoError(t, err)
		assert.False(t, exists, "User should not be able to see pegawai from other unit")
	})
}

// ============================================================================
// TEST 3: RLS Policy - Admin Bypass
// ============================================================================

func TestRLSPolicyAdminBypass(t *testing.T) {
	suite := SetupRLSTest(t)
	defer suite.Cleanup(t)

	ctx := context.Background()

	adminID := uuid.New().String()

	t.Run("admin can see all pegawai across units", func(t *testing.T) {
		// Set admin context
		err := suite.DB.SetTestUser(ctx, adminID, "", "admin")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Count all rows visible (admin should see all)
		count, err := suite.DB.CountRows(ctx, "pegawai", "")
		require.NoError(t, err)

		expectedTotal := len(suite.Fixture.UserIDs)
		assert.Equal(t, expectedTotal, count, "Admin should see all pegawai")
	})

	t.Run("admin can update pegawai in any unit", func(t *testing.T) {
		// Set admin context
		err := suite.DB.SetTestUser(ctx, adminID, "", "admin")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Update pegawai from unit 2 (not admin's unit)
		targetUser := suite.Fixture.UserIDs[3]
		newName := "Updated by Admin"

		result, err := suite.DB.Pool.Exec(ctx,
			"UPDATE pegawai SET nama_lengkap = $1, updated_at = $2 WHERE id = $3",
			newName, time.Now(), targetUser)
		require.NoError(t, err)

		rowsAffected := result.RowsAffected()
		assert.Equal(t, int64(1), rowsAffected, "Admin should be able to update any pegawai")
	})

	t.Run("admin can delete pegawai from any unit", func(t *testing.T) {
		// Create a new pegawai to delete
		newPegawaiID := uuid.New().String()
		unitID := suite.Fixture.UnitIDs[3]

		_, err := suite.DB.Pool.Exec(ctx, `
			INSERT INTO pegawai (id, unit_kerja_id, nip, nama_lengkap, tanggal_lahir, jenis_kelamin, agama, status_pegawai, is_active, created_at, updated_at)
			VALUES ($1, $2, '199912312024011099', 'To Be Deleted', '1999-12-31', 'L', 'Islam', 'PNS', true, $3, $3)
		`, newPegawaiID, unitID, time.Now())
		require.NoError(t, err)

		// Set admin context
		err = suite.DB.SetTestUser(ctx, adminID, "", "admin")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Delete pegawai
		result, err := suite.DB.Pool.Exec(ctx,
			"DELETE FROM pegawai WHERE id = $1", newPegawaiID)
		require.NoError(t, err)

		rowsAffected := result.RowsAffected()
		assert.Equal(t, int64(1), rowsAffected, "Admin should be able to delete any pegawai")
	})
}

// ============================================================================
// TEST 4: RLS Policy - Cross-Unit Access Prevention
// ============================================================================

func TestRLSCrossUnitAccessPrevention(t *testing.T) {
	suite := SetupRLSTest(t)
	defer suite.Cleanup(t)

	ctx := context.Background()

	unit1ID := suite.Fixture.UnitIDs[0]
	userInUnit1 := suite.Fixture.UserIDs[0]

	t.Run("cannot insert pegawai into other unit", func(t *testing.T) {
		// Set user context (user di unit 1)
		err := suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Try to insert pegawai into unit 2
		unit2ID := suite.Fixture.UnitIDs[2]
		newPegawaiID := uuid.New().String()

		_, err = suite.DB.Pool.Exec(ctx, `
			INSERT INTO pegawai (id, unit_kerja_id, nip, nama_lengkap, tanggal_lahir, jenis_kelamin, agama, status_pegawai, is_active, created_at, updated_at)
			VALUES ($1, $2, '199912312024011098', 'Cross Unit Insert', '1999-12-31', 'L', 'Islam', 'PNS', true, $3, $3)
		`, newPegawaiID, unit2ID, time.Now())

		// Should fail due to RLS
		assert.Error(t, err, "User should not be able to insert into other unit")
	})

	t.Run("cannot update pegawai in other unit", func(t *testing.T) {
		// Set user context (user di unit 1)
		err := suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Try to update pegawai in unit 2
		userInUnit2 := suite.Fixture.UserIDs[2]

		result, err := suite.DB.Pool.Exec(ctx,
			"UPDATE pegawai SET nama_lengkap = 'Hacked' WHERE id = $1",
			userInUnit2)
		require.NoError(t, err)

		// Should affect 0 rows due to RLS
		assert.Equal(t, int64(0), result.RowsAffected(),
			"User should not be able to update pegawai in other unit")
	})

	t.Run("cannot delete pegawai in other unit", func(t *testing.T) {
		// Set user context (user di unit 1)
		err := suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Try to delete pegawai in unit 2
		userInUnit2 := suite.Fixture.UserIDs[2]

		result, err := suite.DB.Pool.Exec(ctx,
			"DELETE FROM pegawai WHERE id = $1", userInUnit2)
		require.NoError(t, err)

		// Should affect 0 rows due to RLS
		assert.Equal(t, int64(0), result.RowsAffected(),
			"User should not be able to delete pegawai in other unit")
	})
}

// ============================================================================
// TEST 5: RLS Policy - Satker-Level Isolation
// ============================================================================

func TestRLSPolicySatkerIsolation(t *testing.T) {
	suite := SetupRLSTest(t)
	defer suite.Cleanup(t)

	ctx := context.Background()

	_ = suite.Fixture.SatkerIDs[0] // satker1ID reserved for future tests
	unitInSatker1 := suite.Fixture.UnitIDs[0]
	userInSatker1 := suite.Fixture.UserIDs[0]

	t.Run("user can only see satker data they belong to", func(t *testing.T) {
		// Set user context
		err := suite.DB.SetTestUser(ctx, userInSatker1, unitInSatker1, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Count visible satker (should only see 1)
		count, err := suite.DB.CountRows(ctx, "satker", "")
		require.NoError(t, err)
		assert.LessOrEqual(t, count, 1, "User should see at most 1 satker")
	})

	t.Run("user cannot access other satker data", func(t *testing.T) {
		// Set user context
		err := suite.DB.SetTestUser(ctx, userInSatker1, unitInSatker1, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Try to access other satker
		satker2ID := suite.Fixture.SatkerIDs[1]

		var exists bool
		err = suite.DB.Pool.QueryRow(ctx,
			"SELECT EXISTS(SELECT 1 FROM satker WHERE id = $1)", satker2ID).
			Scan(&exists)
		require.NoError(t, err)
		assert.False(t, exists, "User should not see other satker")
	})
}

// ============================================================================
// TEST 6: RLS Policy - Bulk Operations
// ============================================================================

func TestRLSPolicyBulkOperations(t *testing.T) {
	suite := SetupRLSTest(t)
	defer suite.Cleanup(t)

	ctx := context.Background()

	unit1ID := suite.Fixture.UnitIDs[0]
	userInUnit1 := suite.Fixture.UserIDs[0]

	t.Run("bulk update only affects rows in user's unit", func(t *testing.T) {
		// Set user context
		err := suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Bulk update all pegawai (should only affect user's unit)
		result, err := suite.DB.Pool.Exec(ctx,
			"UPDATE pegawai SET status_pegawai = 'PPPK'")
		require.NoError(t, err)

		// Should only affect 2 rows (pegawai in unit 1)
		assert.Equal(t, int64(2), result.RowsAffected(),
			"Bulk update should only affect rows in user's unit")
	})

	t.Run("bulk delete only affects rows in user's unit", func(t *testing.T) {
		// Create additional pegawai for bulk delete test
		for i := 0; i < 2; i++ {
			_, err := suite.DB.Pool.Exec(ctx, `
				INSERT INTO pegawai (id, unit_kerja_id, nip, nama_lengkap, tanggal_lahir, jenis_kelamin, agama, status_pegawai, is_active, created_at, updated_at)
				VALUES ($1, $2, $3, $4, '1999-12-31', 'L', 'Islam', 'PNS', false, $5, $5)
			`, uuid.New().String(), unit1ID, fmt.Sprintf("19991231202401110%d", i),
				fmt.Sprintf("Bulk Delete %d", i), time.Now())
			require.NoError(t, err)
		}

		// Set user context
		err := suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Bulk delete inactive pegawai
		result, err := suite.DB.Pool.Exec(ctx,
			"DELETE FROM pegawai WHERE is_active = false")
		require.NoError(t, err)

		// Should only delete from user's unit
		assert.Equal(t, int64(2), result.RowsAffected(),
			"Bulk delete should only affect rows in user's unit")
	})
}

// ============================================================================
// TEST 7: RLS Policy - Audit Trail
// ============================================================================

func TestRLSPolicyAuditTrail(t *testing.T) {
	suite := SetupRLSTest(t)
	defer suite.Cleanup(t)

	ctx := context.Background()

	unit1ID := suite.Fixture.UnitIDs[0]
	userInUnit1 := suite.Fixture.UserIDs[0]

	t.Run("audit logs respect RLS for reads", func(t *testing.T) {
		// First, create some audit entries
		_, err := suite.DB.Pool.Exec(ctx, `
			INSERT INTO audit_logs (id, app_source, user_id, user_name, action, resource_type, resource_id, created_at)
			VALUES
				($1, 'test', $2, 'User1', 'UPDATE', 'pegawai', $2, $4),
				($3, 'test', $5, 'User2', 'UPDATE', 'pegawai', $5, $4)
		`, uuid.New().String(), userInUnit1, uuid.New().String(),
			suite.Fixture.UserIDs[2], time.Now())
		require.NoError(t, err)

		// Set user context
		err = suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)
		defer suite.DB.ClearTestUser(ctx)

		// Count audit logs visible
		count, err := suite.DB.CountRows(ctx, "audit_logs", "")
		require.NoError(t, err)

		// Should only see logs related to accessible resources
		assert.LessOrEqual(t, count, 1, "User should only see audit logs for accessible resources")
	})
}

// ============================================================================
// TEST 8: RLS Policy - Edge Cases
// ============================================================================

func TestRLSPolicyEdgeCases(t *testing.T) {
	suite := SetupRLSTest(t)
	defer suite.Cleanup(t)

	ctx := context.Background()

	t.Run("handles null user_id gracefully", func(t *testing.T) {
		// Don't set any user context (null user_id)
		count, err := suite.DB.CountRows(ctx, "pegawai", "")
		require.NoError(t, err)
		assert.Equal(t, 0, count, "Without user context, should see 0 rows")
	})

	t.Run("handles session reset correctly", func(t *testing.T) {
		unit1ID := suite.Fixture.UnitIDs[0]
		userInUnit1 := suite.Fixture.UserIDs[0]

		// Set and clear user context
		err := suite.DB.SetTestUser(ctx, userInUnit1, unit1ID, "user")
		require.NoError(t, err)

		err = suite.DB.ClearTestUser(ctx)
		require.NoError(t, err)

		// Should see 0 rows after clearing context
		count, err := suite.DB.CountRows(ctx, "pegawai", "")
		require.NoError(t, err)
		assert.Equal(t, 0, count, "After clearing context, should see 0 rows")
	})

	t.Run("RLS works within transactions", func(t *testing.T) {
		unit1ID := suite.Fixture.UnitIDs[0]
		userInUnit1 := suite.Fixture.UserIDs[0]

		// Start transaction
		tx, err := suite.DB.Pool.Begin(ctx)
		require.NoError(t, err)
		defer tx.Rollback(ctx)

		// Set user context in transaction
		_, err = tx.Exec(ctx, "SET LOCAL request.jwt.claim.user_id = $1", userInUnit1)
		require.NoError(t, err)

		// Query should respect RLS
		var count int
		err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM pegawai WHERE unit_kerja_id = $1", unit1ID).
			Scan(&count)
		require.NoError(t, err)
		assert.Equal(t, 2, count, "RLS should work within transactions")
	})
}
