#!/bin/bash
# ============================================
# SIKERMA Backup All Databases Script
# ============================================
# Backup semua database dengan GPG encryption
# ============================================

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

# Backup databases
log "Starting full backup process..."

# Backup master database
log "Backing up db_master..."
"${SCRIPT_DIR}/backup-database.sh" "db_master"

# Backup kepegawaian database
log "Backing up db_kepegawaian..."
DB_MASTER_HOST="${DB_KEPEGAWAIAN_HOST:-localhost}" \
DB_MASTER_PORT="${DB_KEPEGAWAIAN_PORT:-5435}" \
DB_MASTER_USER="${DB_KEPEGAWAIAN_USER:-postgres}" \
DB_MASTER_PASSWORD="${DB_KEPEGAWAIAN_PASSWORD:-postgres}" \
"${SCRIPT_DIR}/backup-database.sh" "db_kepegawaian"

# Backup file storage
log "Backing up file storage..."
BACKUP_DIR="/var/backup/sikerma"
DATE=$(date +%Y%m%d_%H%M%S)
STORAGE_BACKUP="${BACKUP_DIR}/file_storage_${DATE}.tar.gz.gpg"
GPG_RECIPIENT="${GPG_RECIPIENT:-admin@sikerma.go.id}"

if [ -d "/var/data/sikerma" ]; then
    tar -czf - /var/data/sikerma 2>/dev/null | \
        gpg --trust-model always --encrypt --recipient "${GPG_RECIPIENT}" \
        --output "${STORAGE_BACKUP}"

    log "File storage backup completed: ${STORAGE_BACKUP}"
fi

# Cleanup old backups (30 days retention)
log "Cleaning up old backups..."
find "${BACKUP_DIR}" -name "*.gpg" -type f -mtime +30 -delete
find "${BACKUP_DIR}" -name "*.sha256" -type f -mtime +30 -delete

log "Full backup process completed!"
