# TASK 003: Setup 3 Next.js Apps

**Sprint:** 1 - Infrastruktur & Fondasi
**Priority:** P0 (Critical Path)
**Estimasi:** 1 hari kerja
**FR:** FR-001

---

## Deskripsi
Setup 3 aplikasi Next.js 14 (App Router) untuk Portal, Master Data, dan Kepegawaian.

---

## File Yang Perlu Dibuat

### Struktur Setiap App
```
apps/{app-name}/
├── app/
│   ├── (auth)/           # Route group: auth pages (login, callback)
│   │   └── login/
│   │       └── page.tsx
│   ├── (app)/            # Route group: protected pages (requires auth)
│   │   ├── layout.tsx
│   │   └── page.tsx
│   └── layout.tsx        # Root layout
├── components/
│   └── layout/
│       ├── sidebar.tsx
│       └── app-header.tsx
├── lib/
│   └── api.ts           # API client instance
├── public/
│   └── favicon.ico
├── next.config.js
├── package.json
├── tsconfig.json
└── tailwind.config.js   # Extend dari @sikerma/ui
```

---

## Acceptance Criteria

- [ ] 3 apps terbuat: `portal`, `master-data`, `kepegawaian`
- [ ] Setiap app bisa dijalankan terpisah (`pnpm dev`)
- [ ] Port tidak conflict: Portal (3000), Master Data (3001), Kepegawaian (3002)
- [ ] Next.js App Router terkonfigurasi
- [ ] Import shared packages (@sikerma/ui, @sikerma/auth, @sikerma/shared) berfungsi
- [ ] Tailwind CSS berfungsi dengan config dari @sikerma/ui
- [ ] TypeScript berfungsi

---

## Next Task
Setelah task ini selesai, lanjut ke:
- **TASK 004:** Bootstrap Go Fiber Backend
