package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Pools menyimpan connection pools untuk kedua database
type Pools struct {
	Master      *pgxpool.Pool
	Kepegawaian *pgxpool.Pool
}

// InitConnections menginisialisasi kedua database connections
func InitConnections(cfg interface{}) (*pgxpool.Pool, *pgxpool.Pool, error) {
	// Import config package di sini untuk akses konfigurasi
	// Untuk saat ini, kita buat manual karena konfigurasi akan di-pass

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Database Master
	masterDSN := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		getEnv("DB_MASTER_USER", "postgres"),
		getEnv("DB_MASTER_PASSWORD", "postgres"),
		getEnv("DB_MASTER_HOST", "localhost"),
		getEnvAsInt("DB_MASTER_PORT", 5435),
		getEnv("DB_MASTER_NAME", "db_master"))

	masterConfig, err := pgxpool.ParseConfig(masterDSN)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse master database config: %w", err)
	}

	// Set pool configuration
	masterConfig.MaxConns = 25
	masterConfig.MinConns = 5
	masterConfig.MaxConnLifetime = 1 * time.Hour
	masterConfig.MaxConnIdleTime = 30 * time.Minute
	masterConfig.HealthCheckPeriod = 1 * time.Minute

	masterPool, err := pgxpool.NewWithConfig(ctx, masterConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create master database pool: %w", err)
	}

	// Test connection
	if err := masterPool.Ping(ctx); err != nil {
		masterPool.Close()
		return nil, nil, fmt.Errorf("failed to ping master database: %w", err)
	}

	logrus.Info("Master database connected successfully")

	// Database Kepegawaian
	kepegawaianDSN := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		getEnv("DB_KEPEGAWAIAN_USER", "postgres"),
		getEnv("DB_KEPEGAWAIAN_PASSWORD", "postgres"),
		getEnv("DB_KEPEGAWAIAN_HOST", "localhost"),
		getEnvAsInt("DB_KEPEGAWAIAN_PORT", 5435),
		getEnv("DB_KEPEGAWAIAN_NAME", "db_kepegawaian"))

	kepegawaianConfig, err := pgxpool.ParseConfig(kepegawaianDSN)
	if err != nil {
		masterPool.Close()
		return nil, nil, fmt.Errorf("failed to parse kepegawaian database config: %w", err)
	}

	// Set pool configuration
	kepegawaianConfig.MaxConns = 25
	kepegawaianConfig.MinConns = 5
	kepegawaianConfig.MaxConnLifetime = 1 * time.Hour
	kepegawaianConfig.MaxConnIdleTime = 30 * time.Minute
	kepegawaianConfig.HealthCheckPeriod = 1 * time.Minute

	kepegawaianPool, err := pgxpool.NewWithConfig(ctx, kepegawaianConfig)
	if err != nil {
		masterPool.Close()
		return nil, nil, fmt.Errorf("failed to create kepegawaian database pool: %w", err)
	}

	// Test connection
	if err := kepegawaianPool.Ping(ctx); err != nil {
		masterPool.Close()
		kepegawaianPool.Close()
		return nil, nil, fmt.Errorf("failed to ping kepegawaian database: %w", err)
	}

	logrus.Info("Kepegawaian database connected successfully")

	return masterPool, kepegawaianPool, nil
}

// Close menutup semua database connections
func Close(master, kepegawaian *pgxpool.Pool) {
	if master != nil {
		master.Close()
		logrus.Info("Master database connection closed")
	}
	if kepegawaian != nil {
		kepegawaian.Close()
		logrus.Info("Kepegawaian database connection closed")
	}
}

// Helper functions untuk environment variables
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}