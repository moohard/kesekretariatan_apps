#!/bin/bash
# ============================================
# SIKERMA Restore Script
# ============================================
# Usage: ./restore-database.sh <encrypted_backup_file> <database_name>
# Example: ./restore-database.sh /var/backup/sikerma/db_master_20240101_120000.sql.gz.gpg db_master
# ============================================

set -e

# ============================================
# Configuration
# ============================================
ENCRYPTED_FILE="${1}"
DATABASE="${2}"

if [ -z "${ENCRYPTED_FILE}" ] || [ -z "${DATABASE}" ]; then
    echo "Usage: $0 <encrypted_backup_file> <database_name>"
    echo "Example: $0 /var/backup/sikerma/db_master_20240101_120000.sql.gz.gpg db_master"
    exit 1
fi

# Database credentials
DB_HOST="${DB_MASTER_HOST:-localhost}"
DB_PORT="${DB_MASTER_PORT:-5435}"
DB_USER="${DB_MASTER_USER:-postgres}"
SECRETS_DIR="../docker/secrets"

# ============================================
# Functions
# ============================================

log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

error() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] ERROR: $1" >&2
    exit 1
}

get_db_password() {
    local secret_file="${SECRETS_DIR}/db_password.txt"
    if [ -f "$secret_file" ]; then
        cat "$secret_file"
    else
        echo "${DB_MASTER_PASSWORD:-postgres}"
    fi
}

# ============================================
# Restore Process
# ============================================

log "Starting restore for database: ${DATABASE}"
log "Backup file: ${ENCRYPTED_FILE}"

# Check if file exists
if [ ! -f "${ENCRYPTED_FILE}" ]; then
    error "Backup file not found: ${ENCRYPTED_FILE}"
fi

# Verify checksum if available
CHECKSUM_FILE="${ENCRYPTED_FILE}.sha256"
if [ -f "${CHECKSUM_FILE}" ]; then
    log "Verifying checksum..."
    if ! sha256sum -c "${CHECKSUM_FILE}" &> /dev/null; then
        error "Checksum verification failed! File may be corrupted."
    fi
    log "Checksum verified."
fi

# Set password
export PGPASSWORD=$(get_db_password)

# Create temporary file for decrypted backup
TEMP_DIR=$(mktemp -d)
COMPRESSED_FILE="${TEMP_DIR}/backup.sql.gz"
SQL_FILE="${TEMP_DIR}/backup.sql"

cleanup() {
    rm -rf "${TEMP_DIR}"
}
trap cleanup EXIT

log "Decrypting backup..."

# Decrypt
gpg --decrypt --output "${COMPRESSED_FILE}" "${ENCRYPTED_FILE}"

if [ $? -ne 0 ]; then
    error "Decryption failed!"
fi

log "Decompressing backup..."

# Decompress
gunzip "${COMPRESSED_FILE}"

log "Restoring database..."

# Restore
psql \
    -h "${DB_HOST}" \
    -p "${DB_PORT}" \
    -U "${DB_USER}" \
    -d "${DATABASE}" \
    -f "${SQL_FILE}"

if [ $? -ne 0 ]; then
    error "Database restore failed!"
fi

# Unset password
unset PGPASSWORD

log "Restore completed successfully!"
