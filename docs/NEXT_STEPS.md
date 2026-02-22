# Next Steps - SIKERMA Project

**Last Updated:** 2026-02-23
**Status:** Post-Checkpoint (Commits: d40a29a → 8c469ae)

---

## Executive Summary

Dokumen ini berisi rencana kerja selanjutnya berdasarkan:
- PRD_REMEDIATION.md (Critical issues dari review)
- Status implementasi TDD untuk RM-008 & RM-013
- Pre-existing issues yang perlu diperbaiki

---

## Phase 1: Test Infrastructure Stabilization (Priority: P0)

### 1.1 Fix Vitest Configuration
**File:** `vitest.config.ts`, `apps/portal/vitest.config.ts`
**Issue:** Test runner masih men-scan `node_modules` (1128 failed suites dari zod tests)

**Tasks:**
- [ ] Update exclude pattern untuk semua vitest config
- [ ] Tambahkan `**/node_modules/**` di exclude array
- [ ] Verifikasi hanya test files di `apps/**/__tests__/` yang di-scan

**Commands:**
```bash
pnpm test 2>&1 | head -50  # Verify tests run correctly
```

### 1.2 Verify Test Results
**Expected:** 29 tests passing (18 Next.js + 11 API contracts)

**Tasks:**
- [ ] Jalankan `pnpm test` dan verifikasi output
- [ ] Document any failing tests
- [ ] Fix failing tests jika ada

---

## Phase 2: RM-008 Next.js 16 Upgrade (Priority: P0 - Critical)

**Reference:** `docs/PRD_REMEDIATION.md` - RM-008
**Current Version:** Next.js 14.2.18
**Target Version:** Next.js 16.x (latest stable)

### 2.1 Breaking Changes Implementation

| Breaking Change | Status | File Reference |
|-----------------|--------|----------------|
| useSearchParams Suspense boundary | ⏳ Pending | `apps/portal/__tests__/next16-breaking-changes.test.tsx` |
| viewport export from metadata | ⏳ Pending | - |
| async server components | ⏳ Pending | - |
| PPR (Partial Prerendering) | ⏳ Pending | - |
| React 19 compatibility | ⏳ Pending | - |
| Image component changes | ⏳ Pending | - |
| Router handler changes | ⏳ Pending | - |

### 2.2 Implementation Steps

```
Step 1: Update package.json dependencies
├── next: ^14.2.18 → ^16.x.x
├── react: ^18 → ^19
├── react-dom: ^18 → ^19
└── Update related @types packages

Step 2: Fix breaking changes di semua apps
├── apps/portal/
├── apps/master-data/
└── apps/kepegawaian/

Step 3: Update shared packages
├── packages/ui/
├── packages/auth/
└── packages/shared/

Step 4: Run full test suite
└── pnpm test && pnpm build
```

### 2.3 Rollback Plan
```bash
# Jika upgrade gagal, rollback ke:
git checkout d40a29a -- package.json pnpm-lock.yaml
pnpm install
```

---

## Phase 3: RM-013 RLS Implementation (Priority: P0 - Critical)

**Reference:** `docs/PRD_REMEDIATION.md` - RM-013
**Files Created:**
- `docker/postgres/migrations/03_enable_rls.sql` ✅
- `backend/internal/middleware/rls.go` ✅
- `backend/internal/rls/rls_test.go` ✅

### 3.1 Remaining Tasks

| Task | Status | Description |
|------|--------|-------------|
| RLS migration SQL | ✅ Done | Policy functions & policies created |
| RLS middleware Go | ✅ Done | JWT claims → PostgreSQL session |
| RLS unit tests | ✅ Done | testcontainers-based tests |
| Auth integration | ⏳ Pending | Connect dengan Keycloak JWT |
| E2E data isolation tests | ⏳ Pending | Playwright tests |
| Migration deployment | ⏳ Pending | Apply to dev database |

### 3.2 Auth Integration Steps

```go
// backend/internal/middleware/rls.go
// TODO: Integrate dengan Keycloak JWT claims

// Flow:
// 1. Keycloak JWT contains: user_id, unit_kerja_id, satker_id, role
// 2. Middleware extracts claims dari JWT
// 3. Set PostgreSQL session variables
// 4. RLS policies automatically filter data
```

### 3.3 Database Migration Commands

```bash
# Apply RLS migration to development database
psql -h localhost -p 5435 -U postgres -d db_master -f docker/postgres/migrations/03_enable_rls.sql
psql -h localhost -p 5435 -U postgres -d db_kepegawaian -f docker/postgres/migrations/03_enable_rls.sql
```

---

## Phase 4: Pre-existing Issues (Priority: P1)

### 4.1 Go Build Errors

**File:** `backend/internal/routes/routes.go`

```
Issue: Pre-existing build errors yang perlu diinvestigasi
Action: Read file, identify errors, fix
```

**Tasks:**
- [ ] Run `go build ./...` di backend directory
- [ ] Document semua build errors
- [ ] Fix errors satu per satu
- [ ] Verify build success

### 4.2 TypeScript Errors

```bash
# Check TypeScript errors
pnpm type-check
```

**Tasks:**
- [ ] Run type-check command
- [ ] Document semua errors
- [ ] Fix berdasarkan priority

---

## Phase 5: Testing & Coverage (Priority: P1)

### 5.1 Unit Tests
```bash
pnpm test              # Run all unit tests
pnpm test:coverage     # Run with coverage report
```

### 5.2 E2E Tests
```bash
pnpm test:e2e          # Run Playwright E2E tests
```

### 5.3 Mutation Testing
```bash
pnpm test:mutation     # Run Stryker mutation testing
```

### 5.4 Coverage Targets
| Type | Target | Current |
|------|--------|---------|
| Unit Tests | 80% | TBD |
| Integration | 70% | TBD |
| E2E | Critical paths | TBD |
| Mutation | 80% kill rate | TBD |

---

## Phase 6: Documentation Updates (Priority: P2)

### 6.1 Update PRD Status
- [ ] Update `docs/PRD_REMEDIATION.md` dengan completion status
- [ ] Mark RM-008 dan RM-013 sebagai completed setelah selesai

### 6.2 Update Technical Docs
- [ ] Update `docs/blueprint_arch.md` dengan RLS architecture
- [ ] Add testing strategy documentation

---

## Quick Reference Commands

```bash
# Development
pnpm dev                    # Start all apps
pnpm --filter=portal dev    # Start specific app

# Testing
pnpm test                   # Unit tests
pnpm test:watch             # Watch mode
pnpm test:coverage          # With coverage
pnpm test:e2e               # E2E tests
pnpm test:mutation          # Mutation tests

# Building
pnpm build                  # Build all
pnpm --filter=@sikerma/ui build  # Build specific package

# Type Checking
pnpm type-check             # TypeScript check

# Git
git status                  # Check status
git log --oneline -5        # Recent commits
```

---

## Commit History Reference

| Commit | Description |
|--------|-------------|
| `8c469ae` | chore: Update .gitignore to exclude plugin state directories |
| `8c60d6e` | chore: Add project configuration and tooling setup |
| `d40a29a` | feat: Initial commit with TDD implementation for RM-008 & RM-013 |

---

## Risk Mitigation

| Risk | Mitigation |
|------|------------|
| Next.js 16 breaking changes | Comprehensive test suite, rollback plan |
| RLS policy errors | Test with testcontainers, gradual rollout |
| Go build failures | Fix incrementally, maintain backward compatibility |
| Test flakiness | Isolate tests, use proper mocks |

---

## Next Session Checklist

- [ ] Run `pnpm test` dan verify 29 tests passing
- [ ] Check `pnpm type-check` results
- [ ] Run `cd backend && go build ./...`
- [ ] Review PRD_REMEDIATION.md untuk detail requirements
- [ ] Start dengan highest priority item (P0)

---

*Document generated from session context on 2026-02-23*
