# Docker Secrets Management

Dokumentasi ini menjelaskan cara mengelola secrets untuk development environment SIKERMA.

## Quick Start

```bash
# Generate semua secrets
./docker/secrets/generate-secrets.sh

# Start services dengan secrets
docker-compose up -d
```

## Struktur Secrets

```
docker/secrets/
├── .gitkeep
├── generate-secrets.sh       # Script untuk generate secrets
├── db_user.txt               # PostgreSQL main DB username
├── db_password.txt           # PostgreSQL main DB password
├── keycloak_db_user.txt      # Keycloak DB username
├── keycloak_db_password.txt  # Keycloak DB password
├── keycloak_admin_user.txt   # Keycloak admin username
├── keycloak_admin_password.txt # Keycloak admin password
├── keycloak_portal_secret.txt       # Portal client secret
├── keycloak_master_data_secret.txt  # Master Data client secret
├── keycloak_kepegawaian_secret.txt  # Kepegawaian client secret
├── keycloak_backend_api_secret.txt  # Backend API client secret
└── dragonfly_password.txt    # DragonflyDB password
```

## Cara Kerja Docker Secrets

Docker Compose membaca secrets dari file lokal dan menyajikannya ke container melalui filesystem (`/run/secrets/`).

### PostgreSQL Example

```yaml
services:
  db:
    environment:
      POSTGRES_USER_FILE: /run/secrets/db_user
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
    secrets:
      - db_user
      - db_password

secrets:
  db_user:
    file: ./docker/secrets/db_user.txt
  db_password:
    file: ./docker/secrets/db_password.txt
```

## Security Best Practices

### ✅ DO
- Jalankan `generate-secrets.sh` sekali saat setup awal
- Backup secrets ke secure storage (misal: 1Password, Vault)
- Rotate secrets secara berkala di production
- Pastikan `docker/secrets/*.txt` ada di `.gitignore`

### ❌ DON'T
- Jangan commit file secrets ke git
- Jangan hardcode credentials di docker-compose.yml
- Jangan share secrets melalui chat/email
- Jangan gunakan password yang mudah ditebak

## Konfigurasi Keycloak Client Secrets

Setelah Keycloak berjalan, update client secrets melalui Admin Console:

1. Buka http://localhost:8081
2. Login dengan credentials dari `keycloak_admin_user.txt` dan `keycloak_admin_password.txt`
3. Pilih realm "pengadilan-agama"
4. Go to Clients → [client-name] → Credentials
5. Regenerate secret dan copy ke:
   - File: `docker/secrets/keycloak_*_secret.txt`
   - .env: `KEYCLOAK_CLIENT_SECRET_*`

## Troubleshooting

### Error: "secret not found"

```bash
# Pastikan secrets sudah di-generate
ls -la docker/secrets/

# Generate jika belum ada
./docker/secrets/generate-secrets.sh
```

### Error: "permission denied"

```bash
# Pastikan script executable
chmod +x docker/secrets/generate-secrets.sh
```

### Database connection failed

```bash
# Cek credentials di secrets
cat docker/secrets/db_user.txt
cat docker/secrets/db_password.txt

# Test connection manual
docker exec -it sikerma_db psql -U $(cat docker/secrets/db_user.txt) -d postgres
```

## Production Deployment

Untuk production, gunakan Docker Swarm secrets atau Kubernetes secrets:

### Docker Swarm

```yaml
secrets:
  db_password:
    external: true
```

```bash
echo "strong_password" | docker secret create db_password -
```

### Kubernetes

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: sikerma-secrets
type: Opaque
data:
  db_password: <base64-encoded>
```

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-02-22 | Initial implementation - Docker Secrets + Version Alignment |
