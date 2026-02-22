# File Storage Configuration

## Directory Structure

```
/var/data/sikerma/
├── uploads/       # File uploads dari user
├── documents/     # Dokumen generated (SK, surat, dll)
├── templates/     # Template dokumen
└── temp/          # Temporary files
```

## Permissions

Recommended permissions for file storage:

```bash
# Set ownership (ganti dengan user yang menjalankan aplikasi)
sudo chown -R www-data:www-data /var/data/sikerma/

# Set permissions
sudo chmod -R 750 /var/data/sikerma/

# Set SELinux context (jika menggunakan SELinux)
sudo chcon -R -t httpd_sys_rw_content_t /var/data/sikerma/
```

## Directory Permissions

| Directory | Permissions | Description |
|-----------|-------------|-------------|
| uploads/  | 750         | User uploads (foto, dokumen) |
| documents/| 750         | Generated documents |
| templates/| 750         | Document templates |
| temp/     | 750         | Temporary files |

## Environment Variables

Tambahkan ke `.env`:

```env
# File Storage Configuration
FILE_STORAGE_PATH=/var/data/sikerma
UPLOAD_MAX_SIZE=10485760  # 10MB in bytes
UPLOAD_ALLOWED_TYPES=jpg,jpeg,png,pdf,doc,docx,xls,xlsx
```

## Backup Strategy

File storage harus di-include dalam backup rutin:

```bash
# Backup script
BACKUP_DATE=$(date +%Y%m%d_%H%M%S)
tar -czvf /backup/sikerma_files_${BACKUP_DATE}.tar.gz /var/data/sikerma/
```

## Docker Volume Mount

File storage dimount sebagai volume di docker-compose.yml:

```yaml
volumes:
  sikerma_file_storage:
    driver: local
    driver_opts:
      type: none
      o: bind
      device: ./docker/data/sikerma
```

## Security Considerations

1. **Virus Scanning**: Scan file yang di-upload dengan ClamAV
2. **File Type Validation**: Validasi file type berdasarkan magic number, bukan extension
3. **Size Limit**: Batasi ukuran file (default 10MB)
4. **Access Control**: File hanya bisa diakses melalui API dengan authentication
5. **Encryption**: File sensitif harus di-encrypt at rest

## Cleanup Policy

Temporary files harus dibersihkan secara berkala:

```bash
# Cron job untuk cleanup temp files (setiap hari pukul 02:00)
0 2 * * * find /var/data/sikerma/temp -type f -mtime +1 -delete
```
