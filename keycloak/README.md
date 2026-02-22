# Keycloak Setup untuk SIKERMA

Dokumentasi untuk setup dan konfigurasi Keycloak sebagai SSO provider untuk SIKERMA.

## Konfigurasi

- **URL**: http://localhost:8081
- **Realm**: pengadilan-agama
- **Admin Console**: http://localhost:8081/admin
- **Admin Credentials**: admin / admin

## Realm: pengadilan-agama

Realm ini sudah dikonfigurasi untuk SIKERMA dengan:

### Clients

| Client ID | Nama | Redirect URIs | Secret |
|-----------|------|---------------|--------|
| portal-client | Portal SIKERMA | http://localhost:3000/* | portal-secret-change-in-production |
| master-data-client | Master Data SIKERMA | http://localhost:3001/* | master-data-secret-change-in-production |
| kepegawaian-client | Kepegawaian SIKERMA | http://localhost:3002/* | kepegawaian-secret-change-in-production |
| backend-api | Backend API | - | backend-api-secret-change-in-production |

### Realm Roles

| Role | Deskripsi |
|------|-----------|
| admin | Administrator dengan akses penuh |
| supervisor | Supervisor dengan akses terbatas untuk approval |
| officer | Officer dengan akses operasional |
| staff | Staff dengan akses dasar |
| user | User dengan akses read-only |

### Default Users

| Username | Password | Email | Roles |
|----------|----------|-------|-------|
| admin | admin123 | admin@sikerma.go.id | admin |
| supervisor | supervisor123 | supervisor@sikerma.go.id | supervisor |
| officer | officer123 | officer@sikerma.go.id | officer |
| staff | staff123 | staff@sikerma.go.id | staff |

## Cara Import Realm

### Via Admin Console

1. Buka Admin Console: http://localhost:8081/admin
2. Login dengan admin/admin
3. Klik tombol dropdown di pojok kiri atas (di sebelah "Master")
4. Klik "Create realm"
5. Pilih "Import" (di tab kanan atas)
6. Upload file `keycloak/realm-export.json`
7. Klik "Create"

### Via Docker Compose

Realm akan otomatis di-import saat container Keycloak dimulai jika:
- File `keycloak/realm-export.json` ada
- Container di-restart setelah pertama kali dibuat

Untuk import manual:
```bash
# Stop Keycloak container
docker-compose stop keycloak

# Remove container (volume data tetap ada)
docker-compose rm -f keycloak

# Start Keycloak dengan import realm
docker-compose up -d keycloak
```

## Cara Import Realm via CLI

Jika Anda ingin import realm menggunakan Keycloak CLI (kcadm.sh):

```bash
# Masuk ke container Keycloak
docker exec -it sikerma_keycloak bash

# Login ke Keycloak
/opt/keycloak/bin/kcadm.sh config credentials --server http://localhost:8080 --realm master --user admin --password admin

# Import realm
/opt/keycloak/bin/kcadm.sh create realms -f /opt/keycloak/data/import/realm-export.json
```

## OIDC Configuration

Untuk aplikasi frontend, gunakan konfigurasi berikut:

### Portal (Next.js pada port 3000)

```typescript
const oidcConfig = {
  authority: 'http://localhost:8081/realms/pengadilan-agama',
  clientId: 'portal-client',
  clientSecret: 'portal-secret-change-in-production',
  redirectUri: 'http://localhost:3000/auth/callback',
  postLogoutRedirectUri: 'http://localhost:3000',
  scope: 'openid profile email roles'
};
```

### Master Data (Next.js pada port 3001)

```typescript
const oidcConfig = {
  authority: 'http://localhost:8081/realms/pengadilan-agama',
  clientId: 'master-data-client',
  clientSecret: 'master-data-secret-change-in-production',
  redirectUri: 'http://localhost:3001/auth/callback',
  postLogoutRedirectUri: 'http://localhost:3001',
  scope: 'openid profile email roles'
};
```

### Kepegawaian (Next.js pada port 3002)

```typescript
const oidcConfig = {
  authority: 'http://localhost:8081/realms/pengadilan-agama',
  clientId: 'kepegawaian-client',
  clientSecret: 'kepegawaian-secret-change-in-production',
  redirectUri: 'http://localhost:3002/auth/callback',
  postLogoutRedirectUri: 'http://localhost:3002',
  scope: 'openid profile email roles'
};
```

### Backend API (Go pada port 3003)

Untuk verifikasi JWT token di backend:

```go
const keycloakConfig = {
  URL: "http://localhost:8081",
  Realm: "pengadilan-agama",
  JWKSURL: "http://localhost:8081/realms/pengadilan-agama/protocol/openid-connect/certs"
}
```

## JWKS Endpoint

Untuk verifikasi JWT token signature:

```
http://localhost:8081/realms/pengadilan-agama/protocol/openid-connect/certs
```

## OpenID Connect Discovery

Untuk mendapatkan konfigurasi OIDC endpoint:

```
http://localhost:8081/realms/pengadilan-agama/.well-known/openid-configuration
```

## User Management

### Membuat User Baru

1. Masuk ke Admin Console
2. Pilih realm "pengadilan-agama"
3. Klik "Users" di menu kiri
4. Klik "Add user"
5. Isi form dan assign role

### Assign Role ke User

1. Klik pada user
2. Klik tab "Role mapping"
3. Klik "Assign role"
4. Pilih realm role yang ingin di-assign

### Reset Password

1. Klik pada user
2. Klik tab "Credentials"
3. Set password baru
4. Centang/desimalang "Temporary" sesuai kebutuhan

## Security Notes

### Penting untuk Produksi

1. **Ganti password default user** - Jangan gunakan admin/admin di production
2. **Ganti client secret** - Ubah semua client secret di production
3. **Hapus test users** - Hapus admin, supervisor, officer, staff di production
4. **HTTPS required** - Set `sslRequired` ke "all" di production
5. **Strong password policy** - Konfigurasi password policy di realm settings

### Password Policy Configuration

Di Admin Console:
1. Realm Settings
2. Login Tab
3. User Registration (jika diaktifkan)
4. Password Policy

Contoh password policy:
- Length: minimal 12 karakter
- Digits: minimal 1 digit
- Uppercase: minimal 1 huruf kapital
- SpecialChars: minimal 1 karakter spesial
- NotUsername: password tidak boleh sama dengan username

## Troubleshooting

### Realm tidak terimport

Cek log Keycloak:
```bash
docker logs sikerma_keycloak
```

### Login gagal

Pastikan:
- User enabled: true
- Email verified: true
- Password benar
- User memiliki role yang sesuai

### JWT token verification gagal

Cek:
- JWKS endpoint bisa diakses
- Token tidak expired
- Token signature valid
- Realm dan client ID benar

## Referensi

- Keycloak Documentation: https://www.keycloak.org/documentation
- OpenID Connect: https://openid.net/connect/
- JWT: https://jwt.io/