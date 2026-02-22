#!/bin/bash
# ============================================
# SIKERMA Backup Script with GPG Encryption
# ============================================
# Usage: ./backup-database.sh [database_name]
# Example: ./backup-database.sh db_master
#
# Prerequisites:
# - GPG installed
# - Recipient public key imported
# - PostgreSQL client tools installed
# ============================================

set -e

# ============================================
# Configuration
# ============================================
BACKUP_DIR="/var/backup/sikerma"
DATE=$(date +%Y%m%d_%H%M%S)
DATABASE="${1:-db_master}"
GPG_RECIPIENT="${GPG_RECIPIENT:-admin@sikerma.go.id}"
RETENTION_DAYS="${RETENTION_DAYS:-30}"

# Database credentials from environment
DB_HOST="${DB_MASTER_HOST:-localhost}"
DB_PORT="${DB_MASTER_PORT:-5435}"
DB_USER="${DB_MASTER_USER:-postgres}"
DB_NAME="${DATABASE}"

# Docker compose path (for getting credentials from secrets)
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

check_prerequisites() {
    # Check GPG
    if ! command -v gpg &> /dev/null; then
        error "GPG is not installed. Please install gnupg."
    fi

    # Check pg_dump
    if ! command -v pg_dump &> /dev/null; then
        error "pg_dump is not installed. Please install postgresql-client."
    fi

    # Create backup directory
    mkdir -p "${BACKUP_DIR}"
}

cleanup_old_backups() {
    log "Cleaning up backups older than ${RETENTION_DAYS} days..."
    find "${BACKUP_DIR}" -name "*.sql.gz.gpg" -type f -mtime +${RETENTION_DAYS} -delete
    find "${BACKUP_DIR}" -name "*.sql.gpg" -type f -mtime +${RETENTION_DAYS} -delete
    log "Cleanup completed."
}

# ============================================
# Main Backup Process
# ============================================

main() {
    log "Starting backup for database: ${DB_NAME}"

    check_prerequisites

    # Set password
    export PGPASSWORD=$(get_db_password)

    # Backup filename
    BACKUP_FILE="${BACKUP_DIR}/${DB_NAME}_${DATE}.sql"
    COMPRESSED_FILE="${BACKUP_FILE}.gz"
    ENCRYPTED_FILE="${COMPRESSED_FILE}.gpg"

    log "Creating database dump..."

    # Create SQL dump
    pg_dump \
        -h "${DB_HOST}" \
        -p "${DB_PORT}" \
        -U "${DB_USER}" \
        -d "${DB_NAME}" \
        -F p \
        --no-owner \
        --no-privileges \
        > "${BACKUP_FILE}"

    if [ $? -ne 0 ]; then
        rm -f "${BACKUP_FILE}"
        error "Database dump failed!"
    fi

    log "Compressing backup..."

    # Compress
    gzip "${BACKUP_FILE}"

    log "Encrypting backup with GPG..."

    # Encrypt with GPG
    gpg \
        --trust-model always \
        --encrypt \
        --recipient "${GPG_RECIPIENT}" \
        --output "${ENCRYPTED_FILE}" \
        "${COMPRESSED_FILE}"

    if [ $? -ne 0 ]; then
        rm -f "${COMPRESSED_FILE}"
        error "Encryption failed!"
    fi

    # Remove unencrypted compressed file
    rm -f "${COMPRESSED_FILE}"

    # Calculate checksum
    CHECKSUM=$(sha256sum "${ENCRYPTED_FILE}" | awk '{print $1}')
    echo "${CHECKSUM}  ${ENCRYPTED_FILE}" > "${ENCRYPTED_FILE}.sha256"

    # Get file size
    FILE_SIZE=$(du -h "${ENCRYPTED_FILE}" | cut -f1)

    log "Backup completed successfully!"
    log "  File: ${ENCRYPTED_FILE}"
    log "  Size: ${FILE_SIZE}"
    log "  SHA256: ${CHECKSUM}"

    # Cleanup old backups
    cleanup_old_backups

    # Unset password
    unset PGPASSWORD
}

# ============================================
# Run
# ============================================

main "$@"
