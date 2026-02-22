package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config menyimpan seluruh konfigurasi aplikasi
type Config struct {
	Server       ServerConfig
	DBMaster     DatabaseConfig
	DBKepegawaian DatabaseConfig
	Keycloak     KeycloakConfig
	Gotenberg    GotenbergConfig
	JWT          JWTConfig
	CORS         CORSConfig
	Logger       LoggerConfig
	Environment  string
	// Convenience fields
	Host         string
	Port         string
	KeycloakURL  string
}

// ServerConfig konfigurasi server
type ServerConfig struct {
	Port string
	Host string
}

// DatabaseConfig konfigurasi database
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// DSN mengembalikan connection string PostgreSQL
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		d.User, d.Password, d.Host, d.Port, d.Name)
}

// KeycloakConfig konfigurasi Keycloak
type KeycloakConfig struct {
	URL      string
	Realm    string
	JWKSURL  string
}

// GotenbergConfig konfigurasi Gotenberg untuk PDF generation
type GotenbergConfig struct {
	URL string
}

// JWTConfig konfigurasi JWT
type JWTConfig struct {
	Secret     string
	Expiration string
}

// CORSConfig konfigurasi CORS
type CORSConfig struct {
	Origins     string
	Credentials bool
}

// LoggerConfig konfigurasi logger
type LoggerConfig struct {
	Level  string
	Format string
}

// Load memuat konfigurasi dari environment variables
func Load() *Config {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "3003"),
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
		},
		DBMaster: DatabaseConfig{
			Host:     getEnv("DB_MASTER_HOST", "localhost"),
			Port:     getEnvAsInt("DB_MASTER_PORT", 5435),
			User:     getEnv("DB_MASTER_USER", "postgres"),
			Password: getEnv("DB_MASTER_PASSWORD", "postgres"),
			Name:     getEnv("DB_MASTER_NAME", "db_master"),
		},
		DBKepegawaian: DatabaseConfig{
			Host:     getEnv("DB_KEPEGAWAIAN_HOST", "localhost"),
			Port:     getEnvAsInt("DB_KEPEGAWAIAN_PORT", 5435),
			User:     getEnv("DB_KEPEGAWAIAN_USER", "postgres"),
			Password: getEnv("DB_KEPEGAWAIAN_PASSWORD", "postgres"),
			Name:     getEnv("DB_KEPEGAWAIAN_NAME", "db_kepegawaian"),
		},
		Keycloak: KeycloakConfig{
			URL:     getEnv("KEYCLOAK_URL", "http://localhost:8081"),
			Realm:   getEnv("KEYCLOAK_REALM", "pengadilan-agama"),
			JWKSURL: getEnv("KEYCLOAK_JWKS_URL", "http://localhost:8081/realms/pengadilan-agama/protocol/openid-connect/certs"),
		},
		Gotenberg: GotenbergConfig{
			URL: getEnv("GOTENBERG_URL", "http://localhost:3100"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-jwt-secret-change-in-production"),
			Expiration: getEnv("JWT_EXPIRATION", "24h"),
		},
		CORS: CORSConfig{
			Origins:     getEnv("CORS_ORIGINS", "http://localhost:3000,http://localhost:3001,http://localhost:3002"),
			Credentials: getEnvAsBool("CORS_CREDENTIALS", true),
		},
		Logger: LoggerConfig{
			Level:  getEnv("LOG_LEVEL", "debug"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		Environment: getEnv("ENVIRONMENT", "development"),
	}
	// Set convenience fields
	cfg.Host = cfg.Server.Host
	cfg.Port = cfg.Server.Port
	cfg.KeycloakURL = cfg.Keycloak.URL

	return cfg
}

// getEnv mengambil environment variable dengan default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt mengambil environment variable sebagai integer
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsBool mengambil environment variable sebagai boolean
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}