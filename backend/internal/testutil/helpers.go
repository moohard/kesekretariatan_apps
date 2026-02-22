// Package testutil provides testing utilities for SIKERMA backend
package testutil

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestDB holds the test database connection and container
type TestDB struct {
	Pool      *pgxpool.Pool
	Container *postgres.PostgresContainer
	ConnStr   string
}

// SetupTestDB creates a test database container for integration tests
// Usage: defer testdb.Cleanup(t)
func SetupTestDB(t *testing.T, migrationsPath string) *TestDB {
	ctx := context.Background()

	// Create PostgreSQL container with migrations
	container, err := postgres.Run(ctx,
		"postgres:18-alpine",
		postgres.WithInitScripts(migrationsPath),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %s", err)
	}

	// Create connection pool
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		t.Fatalf("failed to create pool: %s", err)
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		t.Fatalf("failed to ping database: %s", err)
	}

	return &TestDB{
		Pool:      pool,
		Container: container,
		ConnStr:   connStr,
	}
}

// Cleanup stops the container and closes the pool
func (tdb *TestDB) Cleanup(t *testing.T) {
	if tdb.Pool != nil {
		tdb.Pool.Close()
	}
	if tdb.Container != nil {
		if err := testcontainers.TerminateContainer(tdb.Container); err != nil {
			t.Logf("failed to terminate container: %s", err)
		}
	}
}

// SetTestUser sets the current user context for RLS tests
// This simulates setting the JWT claims in PostgreSQL session
func (tdb *TestDB) SetTestUser(ctx context.Context, userID, unitKerjaID, role string) error {
	_, err := tdb.Pool.Exec(ctx, `
		SET LOCAL request.jwt.claim.user_id = $1;
		SET LOCAL request.jwt.claim.unit_kerja_id = $2;
		SET LOCAL request.jwt.claim.role = $3;
	`, userID, unitKerjaID, role)
	return err
}

// ClearTestUser clears the test user context
func (tdb *TestDB) ClearTestUser(ctx context.Context) error {
	_, err := tdb.Pool.Exec(ctx, `
		RESET request.jwt.claim.user_id;
		RESET request.jwt.claim.unit_kerja_id;
		RESET request.jwt.claim.role;
	`)
	return err
}

// CountRows counts rows in a table with optional where clause
func (tdb *TestDB) CountRows(ctx context.Context, table string, where string, args ...any) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	if where != "" {
		query += " WHERE " + where
	}

	var count int
	err := tdb.Pool.QueryRow(ctx, query, args...).Scan(&count)
	return count, err
}

// GetEnvOrDefault gets environment variable or returns default value
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
