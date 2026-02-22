#!/bin/bash
# Script untuk generate Docker Secrets
# Jalankan sekali untuk setup development environment
# JANGAN commit file secrets yang di-generate ke version control!

SECRETS_DIR="$(dirname "$0")"

# Function untuk generate random password
generate_password() {
    openssl rand -base64 32 | tr -d '\n'
}

echo "🔐 Generating Docker Secrets untuk SIKERMA Development..."
echo ""

# Database credentials
echo "📊 Database Secrets..."
echo -n "sikerma_admin" > "$SECRETS_DIR/db_user.txt"
generate_password > "$SECRETS_DIR/db_password.txt"
echo "   ✓ db_user.txt, db_password.txt"

# Keycloak database credentials
echo -n "keycloak" > "$SECRETS_DIR/keycloak_db_user.txt"
generate_password > "$SECRETS_DIR/keycloak_db_password.txt"
echo "   ✓ keycloak_db_user.txt, keycloak_db_password.txt"

# Keycloak admin credentials
echo -n "sikerma_admin" > "$SECRETS_DIR/keycloak_admin_user.txt"
generate_password > "$SECRETS_DIR/keycloak_admin_password.txt"
echo "   ✓ keycloak_admin_user.txt, keycloak_admin_password.txt"

# Keycloak client secrets
generate_password > "$SECRETS_DIR/keycloak_portal_secret.txt"
generate_password > "$SECRETS_DIR/keycloak_master_data_secret.txt"
generate_password > "$SECRETS_DIR/keycloak_kepegawaian_secret.txt"
generate_password > "$SECRETS_DIR/keycloak_backend_api_secret.txt"
echo "   ✓ keycloak_*_secret.txt (4 files)"

# DragonflyDB password (untuk future implementation)
generate_password > "$SECRETS_DIR/dragonfly_password.txt"
echo "   ✓ dragonfly_password.txt"

echo ""
echo "✅ Semua secrets berhasil di-generate!"
echo ""
echo "📝 Langkah selanjutnya:"
echo "   1. Review file secrets di: $SECRETS_DIR/"
echo "   2. Update .env dengan secrets yang sesuai"
echo "   3. Jalankan: docker-compose up -d"
echo ""
echo "⚠️  PENTING: File secrets SUDAH di-ignore di .gitignore"
echo "   Pastikan TIDAK ADA file secrets yang di-commit ke git!"
