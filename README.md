# SIKERMA Monorepo

Monorepo untuk aplikasi SIKERMA menggunakan Turborepo dengan pnpm workspace.

## 📁 Struktur Monorepo

```
/sikerma/
├── apps/                          # Next.js Applications
│   ├── portal/                   # Portal Launcher & Dashboard
│   ├── master-data/              # Master Data Management
│   └── kepegawaian/              # Kepegawaian Management
├── packages/                      # Shared Packages
│   ├── ui/                       # @sikerma/ui - UI Components
│   ├── auth/                     # @sikerma/auth - Auth Utilities
│   └── shared/                   # @sikerma/shared - Shared Utilities
├── docker/
│   └── postgres/
│       └── init/                 # Database migration scripts
├── turbo.json                     # Turborepo Configuration
├── package.json                   # Root Package.json
├── pnpm-workspace.yaml            # Workspace Definition
├── .gitignore
├── .env.example
└── README.md                     # File ini
```

## 🚀 Quick Start

### Prerequisites

- Node.js >= 18.0.0
- pnpm >= 8.0.0
- Docker & Docker Compose

### Installation

```bash
# Install dependencies
pnpm install

# Start Docker services (PostgreSQL, Keycloak)
docker-compose up -d

# Copy environment variables
cp .env.example .env
```

### Development

```bash
# Run all apps in dev mode
pnpm dev

# Build all apps
pnpm build

# Type check
pnpm type-check

# Lint
pnpm lint
```

### Package-specific Commands

```bash
# Build specific package
pnpm --filter=@sikerma/ui run build

# Run specific app
pnpm --filter=portal run dev
```

## 📦 Shared Packages

| Package | Description |
|---------|-------------|
| `@sikerma/ui` | Reusable UI components (Button, Card, Table, etc.) |
| `@sikerma/auth` | Authentication & session management |
| `@sikerma/shared` | Shared utilities & types |

## 🔧 Available Apps

| App | Port | Description |
|-----|------|-------------|
| Portal | 3000 | Dashboard & Launcher |
| Master Data | 3001 | Master Data CRUD |
| Kepegawaian | 3002 | Employee Management |
| Backend API | 3003 | Go Fiber API |

## 🐳 Docker Services

| Service | Port | Description |
|---------|------|-------------|
| PostgreSQL (Master) | 5435 | Database for master data |
| PostgreSQL (Kepegawaian) | 5436 | Database for kepegawaian |
| Keycloak | 8081 | Identity & Access Management |
| Gotenberg | 3100 | PDF Generation Service |

## 📝 Scripts

| Command | Description |
|---------|-------------|
| `pnpm dev` | Start all apps in dev mode |
| `pnpm build` | Build all apps and packages |
| `pnpm lint` | Run ESLint |
| `pnpm type-check` | Run TypeScript type checking |
| `pnpm clean` | Clean all build directories |

## 🎯 Next Steps

1. Setup shared packages (@sikerma/ui, @sikerma/auth, @sikerma/shared)
2. Setup Next.js apps (portal, master-data, kepegawaian)
3. Configure Keycloak realm and clients
4. Setup database migrations
5. Implement backend API (Go Fiber)

## 📚 References

- [Turborepo Docs](https://turbo.build/repo/docs)
- [pnpm Workspaces](https://pnpm.io/workspaces)
- [Next.js Docs](https://nextjs.org/docs)
- [Go Fiber Docs](https://docs.gofiber.io/)
