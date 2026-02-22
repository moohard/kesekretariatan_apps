# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**SIKERMA** (Sistem Informasi Kesekretariatan Mahkamah Agung) is a monorepo for building a comprehensive secretary事务 system for Indonesian Religious Courts. The project follows a phased delivery approach, with Fase 1 focused on establishing the foundational infrastructure.

### Architecture

- **Monorepo:** Turborepo + pnpm workspace
- **Frontend:** 3 Next.js 14 (App Router) applications (Micro-frontend pattern)
- **Backend:** Go Fiber v3 monolith API server
- **Auth:** Keycloak 26.0 (SSO via OIDC)
- **Database:** PostgreSQL 17 (separate databases per module)
- **PDF Generation:** Gotenberg 8

### Apps

| App | Port | Module | Status |
|-----|------|--------|--------|
| Portal | 3000 | Launcher & Dashboard | Fase 1 |
| Master Data | 3001 | Master Data CRUD | Fase 1 |
| Kepegawaian | 3002 | Employee Management | Fase 1 |

### Shared Packages

| Package | Description |
|---------|-------------|
| `@sikerma/ui` | Reusable UI components (shadcn/ui + custom) |
| `@sikerma/auth` | Authentication using Better Auth + Keycloak OIDC |
| `@sikerma/shared` | Shared utilities, API client, types |

## Development Commands

### Root Level
```bash
pnpm install      # Install all dependencies
pnpm dev          # Start all apps in dev mode
pnpm build        # Build all apps and packages
pnpm lint         # Run ESLint on all projects
pnpm type-check   # Run TypeScript type checking
pnpm clean        # Clean all build directories
```

### App-Specific Commands
```bash
pnpm --filter=portal run dev
pnpm --filter=master-data run build
pnpm --filter=kepegawaian run lint
```

### Package-Specific Commands
```bash
pnpm --filter=@sikerma/ui run build
pnpm --filter=@sikerma/shared run type-check
```

## Environment Setup

Copy `.env.example` to `.env` and configure:

```bash
cp .env.example .env
```

Key environment variables include:
- Database URLs for `db_master` and `db_kepegawaian`
- Keycloak configuration (URL, realm, client credentials)
- Backend API URL
- Port configurations for all apps

## Docker Services

```bash
# Start databases and Keycloak
docker-compose up -d

# Services
# PostgreSQL (Master): localhost:5435
# Keycloak: localhost:8081
```

## Project Structure

```
/apps
├── portal/          # Next.js App: Launcher & Dashboard
├── master-data/     # Next.js App: Master Data CRUD
└── kepegawaian/     # Next.js App: Employee Management

/packages
├── ui/              # @sikerma/ui
├── auth/            # @sikerma/auth
└── shared/          # @sikerma/shared

/docs
├── PRD_FASE_1/      # Fase 1 planning & task breakdown
├── blueprint_arch.md
└── overview_aplikasi.md

/docker
├── postgres/init/   # Database migration scripts
└── docker-compose.yml
```

## Implementation Phases

### Fase 1: Fondasi (Current Focus)
**Timeline:** ~6-8 minggu (4 sprints)

| Sprint | Focus | Duration | Status |
|--------|-------|----------|--------|
| Sprint 1 | Infrastruktur & Fondasi | 8-10 hari | Not Started |
| Sprint 2 | Master Data (CRUD Lengkap) | 6-8 hari | Blocked |
| Sprint 3 | Kepegawaian Dasar | 8-10 hari | Blocked |
| Sprint 4 | Portal + Cetak SK + Polish | 6-8 hari | Blocked |

See `docs/PRD_FASE_1/` for detailed task breakdown.

## Tech Stack

| Layer | Technology |
|-------|------------|
| Frontend Framework | Next.js 14 (App Router) |
| UI Library | shadcn/ui + Tailwind CSS |
| State/Fetch | TanStack Query |
| Form | React Hook Form + Zod |
| Backend | Go Fiber v3 |
| Auth | Keycloak 26 + Better Auth |
| Database | PostgreSQL 17 |

## Key Principles

1. **Micro-frontend pattern:** Each app is a separate Next.js application
2. **Single Sign-On:** Keycloak manages SSO across all apps
3. **Database per module:** Separate databases for isolation
4. **Shared packages:** Reusable UI components, auth, and utilities
5. **Phased delivery:** Fase 1 foundation enables future enhancements

## Reference Documentation

- `docs/PRD_FASE_1/` - Complete Fase 1 planning and task breakdown
- `docs/PRD_FASE_1/INDEX.md` - Task navigation and index
- `docs/overview_aplikasi.md` - Application architecture overview
- `docs/blueprint_arch.md` - Technical architecture blueprint
- `context7 and omnisearch` - Troubleshoot or Reference tech stack
