-- ============================================================================
-- RM-013: Row-Level Security Migration
--
-- File ini mengimplementasikan RLS policies untuk data isolation
-- antar unit kerja dan satker.
--
-- Execution Order:
-- 1. Create helper functions
-- 2. Enable RLS on tables
-- 3. Create policies for different roles
-- ============================================================================

-- ============================================================================
-- PART 1: Helper Functions
-- ============================================================================

-- Function: current_user_id()
-- Returns the current user's ID from JWT claims
-- Returns NULL if not authenticated
CREATE OR REPLACE FUNCTION current_user_id()
RETURNS UUID
LANGUAGE SQL
SECURITY DEFINER
STABLE
AS $$
    SELECT NULLIF(
        current_setting('request.jwt.claim.user_id', true),
        ''
    )::UUID;
$$;

-- Function: current_unit_kerja_id()
-- Returns the current user's unit kerja ID from JWT claims
CREATE OR REPLACE FUNCTION current_unit_kerja_id()
RETURNS UUID
LANGUAGE SQL
SECURITY DEFINER
STABLE
AS $$
    SELECT NULLIF(
        current_setting('request.jwt.claim.unit_kerja_id', true),
        ''
    )::UUID;
$$;

-- Function: current_satker_id()
-- Returns the current user's satker ID from JWT claims
CREATE OR REPLACE FUNCTION current_satker_id()
RETURNS UUID
LANGUAGE SQL
SECURITY DEFINER
STABLE
AS $$
    SELECT NULLIF(
        current_setting('request.jwt.claim.satker_id', true),
        ''
    )::UUID;
$$;

-- Function: current_user_role()
-- Returns the current user's role from JWT claims
CREATE OR REPLACE FUNCTION current_user_role()
RETURNS TEXT
LANGUAGE SQL
SECURITY DEFINER
STABLE
AS $$
    SELECT COALESCE(
        NULLIF(
            current_setting('request.jwt.claim.role', true),
            ''
        ),
        'anonymous'
    );
$$;

-- Function: is_admin()
-- Returns true if current user has admin role
CREATE OR REPLACE FUNCTION is_admin()
RETURNS BOOLEAN
LANGUAGE SQL
SECURITY DEFINER
STABLE
AS $$
    SELECT current_user_role() IN ('admin', 'superadmin');
$$;

-- Function: is_user_in_unit(unit_id UUID)
-- Returns true if current user belongs to the given unit
CREATE OR REPLACE FUNCTION is_user_in_unit(unit_id UUID)
RETURNS BOOLEAN
LANGUAGE SQL
SECURITY DEFINER
STABLE
AS $$
    SELECT
        is_admin() OR
        current_unit_kerja_id() = unit_id;
$$;

-- ============================================================================
-- PART 2: Enable RLS on Tables
-- ============================================================================

-- Enable RLS on pegawai table
ALTER TABLE pegawai ENABLE ROW LEVEL SECURITY;

-- Enable RLS on unit_kerja table
ALTER TABLE unit_kerja ENABLE ROW LEVEL SECURITY;

-- Enable RLS on satker table
ALTER TABLE satker ENABLE ROW LEVEL SECURITY;

-- Enable RLS on audit_logs table (optional, for audit isolation)
-- ALTER TABLE audit_logs ENABLE ROW LEVEL SECURITY;

-- ============================================================================
-- PART 3: Pegawai Policies
-- ============================================================================

-- Policy: pegawai_select_policy
-- Users can only see pegawai in their own unit kerja
-- Admins can see all pegawai
CREATE POLICY pegawai_select_policy ON pegawai
    FOR SELECT
    USING (
        is_admin() OR
        unit_kerja_id IN (
            SELECT id FROM unit_kerja
            WHERE id = current_unit_kerja_id()
            OR satker_id = current_satker_id()
        )
    );

-- Policy: pegawai_insert_policy
-- Users can only insert pegawai into their own unit kerja
-- Admins can insert anywhere
CREATE POLICY pegawai_insert_policy ON pegawai
    FOR INSERT
    WITH CHECK (
        is_admin() OR
        unit_kerja_id = current_unit_kerja_id()
    );

-- Policy: pegawai_update_policy
-- Users can only update pegawai in their own unit kerja
-- Admins can update any pegawai
CREATE POLICY pegawai_update_policy ON pegawai
    FOR UPDATE
    USING (
        is_admin() OR
        unit_kerja_id = current_unit_kerja_id()
    )
    WITH CHECK (
        is_admin() OR
        unit_kerja_id = current_unit_kerja_id()
    );

-- Policy: pegawai_delete_policy
-- Users can only delete pegawai in their own unit kerja
-- Admins can delete any pegawai
CREATE POLICY pegawai_delete_policy ON pegawai
    FOR DELETE
    USING (
        is_admin() OR
        unit_kerja_id = current_unit_kerja_id()
    );

-- ============================================================================
-- PART 4: Unit Kerja Policies
-- ============================================================================

-- Policy: unit_kerja_select_policy
-- Users can see unit kerja within their satker
-- Admins can see all unit kerja
CREATE POLICY unit_kerja_select_policy ON unit_kerja
    FOR SELECT
    USING (
        is_admin() OR
        satker_id = current_satker_id() OR
        id = current_unit_kerja_id()
    );

-- Policy: unit_kerja_insert_policy
-- Only admins can create new unit kerja
CREATE POLICY unit_kerja_insert_policy ON unit_kerja
    FOR INSERT
    WITH CHECK (is_admin());

-- Policy: unit_kerja_update_policy
-- Only admins can update unit kerja
CREATE POLICY unit_kerja_update_policy ON unit_kerja
    FOR UPDATE
    USING (is_admin())
    WITH CHECK (is_admin());

-- Policy: unit_kerja_delete_policy
-- Only admins can delete unit kerja
CREATE POLICY unit_kerja_delete_policy ON unit_kerja
    FOR DELETE
    USING (is_admin());

-- ============================================================================
-- PART 5: Satker Policies
-- ============================================================================

-- Policy: satker_select_policy
-- Users can only see their own satker
-- Admins can see all satker
CREATE POLICY satker_select_policy ON satker
    FOR SELECT
    USING (
        is_admin() OR
        id = current_satker_id()
    );

-- Policy: satker_insert_policy
-- Only superadmin can create new satker
CREATE POLICY satker_insert_policy ON satker
    FOR INSERT
    WITH CHECK (current_user_role() = 'superadmin');

-- Policy: satker_update_policy
-- Only superadmin can update satker
CREATE POLICY satker_update_policy ON satker
    FOR UPDATE
    USING (current_user_role() = 'superadmin')
    WITH CHECK (current_user_role() = 'superadmin');

-- Policy: satker_delete_policy
-- Only superadmin can delete satker
CREATE POLICY satker_delete_policy ON satker
    FOR DELETE
    USING (current_user_role() = 'superadmin');

-- ============================================================================
-- PART 6: Audit Log Policies (Optional)
-- ============================================================================

-- Uncomment if audit log isolation is needed

-- Policy: audit_logs_select_policy
-- Users can see audit logs for resources they can access
-- CREATE POLICY audit_logs_select_policy ON audit_logs
--     FOR SELECT
--     USING (
--         is_admin() OR
--         user_id = current_user_id() OR
--         resource_type = 'pegawai' AND
--         resource_id IN (
--             SELECT id FROM pegawai
--             WHERE unit_kerja_id = current_unit_kerja_id()
--         )
--     );

-- ============================================================================
-- PART 7: Additional Indexes for RLS Performance
-- ============================================================================

-- Index for RLS policy lookups
CREATE INDEX IF NOT EXISTS idx_pegawai_unit_kerja_rls ON pegawai(unit_kerja_id);
CREATE INDEX IF NOT EXISTS idx_unit_kerja_satker_rls ON unit_kerja(satker_id);

-- ============================================================================
-- PART 8: Grant Permissions
-- ============================================================================

-- Grant execute on functions to authenticated users
GRANT EXECUTE ON FUNCTION current_user_id() TO authenticated;
GRANT EXECUTE ON FUNCTION current_unit_kerja_id() TO authenticated;
GRANT EXECUTE ON FUNCTION current_satker_id() TO authenticated;
GRANT EXECUTE ON FUNCTION current_user_role() TO authenticated;
GRANT EXECUTE ON FUNCTION is_admin() TO authenticated;
GRANT EXECUTE ON FUNCTION is_user_in_unit(UUID) TO authenticated;

-- ============================================================================
-- VERIFICATION QUERIES (Run after migration)
-- ============================================================================

-- Verify RLS is enabled:
-- SELECT tablename, rowsecurity FROM pg_tables
-- WHERE schemaname = 'public' AND tablename IN ('pegawai', 'unit_kerja', 'satker');

-- Verify policies exist:
-- SELECT schemaname, tablename, policyname, cmd, qual, with_check
-- FROM pg_policies WHERE schemaname = 'public';

-- Test with different users:
-- SET request.jwt.claim.user_id = '<user-uuid>';
-- SET request.jwt.claim.unit_kerja_id = '<unit-uuid>';
-- SET request.jwt.claim.role = 'user';
-- SELECT * FROM pegawai; -- Should only see pegawai in user's unit
