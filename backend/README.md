# SIKERMA Backend API

Backend API untuk SIKERMA (Sistem Informasi Kesekretariatan Mahkamah Agung) menggunakan Go Fiber v3.

## Tech Stack

- **Framework**: Go Fiber v3
- **Database**: PostgreSQL 17 (multi-database: db_master, db_kepegawaian)
- **Driver**: pgx/v5
- **Authentication**: Keycloak JWT (OIDC)
- **Logging**: logrus + zap
- **PDF Generation**: Gotenberg 8

## Project Structure

```
backend/
├── cmd/
│   └── main.go              # Entry point aplikasi
├── internal/
│   ├── config/              # Konfigurasi aplikasi
│   ├── database/            # Database connections & pools
│   ├── handlers/            # HTTP handlers per module
│   ├── middleware/          # Custom middleware (auth, RBAC, audit)
│   ├── models/              # Data models & DTOs
│   ├── repositories/        # Database repositories
│   ├── services/            # Business logic layer
│   └── utils/               # Utility functions
├── pkg/                     # Reusable packages
├── docker/                  # Dockerfile untuk deployment
├── go.mod
├── go.sum
└── .env.example
```

## Setup

### Prerequisites

- Go 1.23.0 atau lebih baru
- PostgreSQL 17 berjalan pada port 5435
- Keycloak berjalan pada port 8081
- Gotenberg berjalan pada port 3100

### Installation

```bash
# Install dependencies
go mod download

# Copy environment variables
cp .env.example .env

# Edit .env sesuai konfigurasi
nano .env
```

### Running the Application

```bash
# Development mode
go run cmd/main.go

# Build binary
go build -o sikerma-api cmd/main.go

# Run binary
./sikerma-api
```

### Docker

```bash
# Build image
docker build -t sikerma-backend:latest .

# Run container
docker run -p 3003:3003 --env-file .env sikerma-backend:latest
```

## API Endpoints

Base URL: `http://localhost:3003/api/v1`

### Authentication
- `POST /auth/login` - Login dengan Keycloak
- `POST /auth/logout` - Logout
- `POST /auth/refresh` - Refresh access token
- `GET /auth/me` - Get current user info

### Master Data
- `GET /master-data/satker` - List satker
- `POST /master-data/satker` - Create satker
- `GET /master-data/satker/:id` - Get satker detail
- `PUT /master-data/satker/:id` - Update satker
- `DELETE /master-data/satker/:id` - Delete satker

... (similiar untuk 9 entitas lain: jabatan, golongan, unit_kerja, eselon, ref_pendidikan, ref_agama, ref_status_kawin, ref_jenis_hukdis, ref_jenis_diklat)

### Kepegawaian
- `GET /pegawai` - List pegawai dengan pagination
- `POST /pegawai` - Create pegawai baru
- `GET /pegawai/:id` - Detail pegawai
- `PUT /pegawai/:id` - Update pegawai
- `DELETE /pegawai/:id` - Delete pegawai
- `GET /pegawai/:id/riwayat-pangkat` - Riwayat pangkat
- `POST /pegawai/:id/riwayat-pangkat` - Tambah riwayat pangkat
- ... (riwayat lain: jabatan, pendidikan, keluarga)
- `POST /pegawai/:id/upload-foto` - Upload foto pegawai
- `POST /pegawai/:id/upload-sk/:tipe` - Upload SK

### Statistik
- `GET /kepegawaian/statistik` - Statistik kepegawaian
- `GET /kepegawaian/statistik/pangkat` - Statistik per pangkat
- `GET /kepegawaian/statistik/jabatan` - Statistik per jabatan

### RBAC
- `GET /rbac/roles` - List roles
- `POST /rbac/roles` - Create role
- `GET /rbac/permissions` - List permissions
- `POST /rbac/permissions` - Create permission
- `POST /rbac/roles/:id/permissions` - Assign permission ke role
- `POST /rbac/users/:id/roles` - Assign role ke user

### Audit Log
- `GET /audit-logs` - List audit logs dengan filter

### PDF Generation
- `POST /pdf/generate` - Generate PDF dari template
- `GET /pdf/templates` - List available templates

## Configuration

Environment variables di `.env`:

```env
# Server
SERVER_PORT=3003
SERVER_HOST=0.0.0.0

# Database Master
DB_MASTER_HOST=localhost
DB_MASTER_PORT=5435
DB_MASTER_USER=postgres
DB_MASTER_PASSWORD=postgres
DB_MASTER_NAME=db_master

# Database Kepegawaian
DB_KEPEGAWAIAN_HOST=localhost
DB_KEPEGAWAIAN_PORT=5435
DB_KEPEGAWAIAN_USER=postgres
DB_KEPEGAWAIAN_PASSWORD=postgres
DB_KEPEGAWAIAN_NAME=db_kepegawaian

# Keycloak
KEYCLOAK_URL=http://localhost:8081
KEYCLOAK_REALM=pengadilan-agama
KEYCLOAK_JWKS_URL=http://localhost:8081/realms/pengadilan-agama/protocol/openid-connect/certs

# Gotenberg
GOTENBERG_URL=http://localhost:3100

# JWT
JWT_SECRET=your-jwt-secret
JWT_EXPIRATION=24h

# Logging
LOG_LEVEL=debug
LOG_FORMAT=json
```

## Database Schema

### db_master
- Tabel referensi: satker, jabatan, golongan, unit_kerja, eselon, ref_pendidikan, ref_agama, ref_status_kawin, ref_jenis_hukdis, ref_jenis_diklat
- Tabel RBAC: app_roles, app_permissions, role_permissions, user_app_roles
- Tabel audit: audit_logs

### db_kepegawaian
- pegawai, riwayat_pangkat, riwayat_jabatan, riwayat_pendidikan, keluarga, template_dokumen

## Authentication Flow

1. Frontend redirect ke Keycloak untuk login
2. Keycloak mengembalikan JWT access token
3. Frontend mengirim JWT ke backend untuk setiap request
4. Backend verifikasi JWT signature dan claims
5. RBAC middleware cek permissions

## Development

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/handlers/auth
```

### Code Style

```bash
# Format code
go fmt ./...

# Run linter (jika terinstall golangci-lint)
golangci-lint run
```

## Deployment

### PM2

```bash
# Start with PM2
pm2 start sikerma-api --name sikerma-backend

# View logs
pm2 logs sikerma-backend

# Restart
pm2 restart sikerma-backend
```

### Docker Compose (Production)

```bash
# Build and run
docker-compose -f docker-compose.prod.yml up -d
```

## Troubleshooting

### Database Connection Error

```bash
# Check PostgreSQL is running
docker ps | grep postgres

# Check port
netstat -tulpn | grep 5435
```

### Keycloak Connection Error

```bash
# Check Keycloak is running
docker ps | grep keycloak

# Check JWKS endpoint
curl http://localhost:8081/realms/pengadilan-agama/protocol/openid-connect/certs
```

## License

Copyright 2026 Mahkamah Agung Republik Indonesia