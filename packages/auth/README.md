# @sikerma/auth - Authentication Utilities

Authentication and session management utilities for SIKERMA application.

## Installation

```bash
pnpm install @sikerma/auth
```

## Usage

```tsx
import { authClient } from '@sikerma/auth'

// Initialize auth
await authClient.init()

// Check if user is authenticated
if (authClient.session) {
  console.log('User:', authClient.session.user)
}
```

## Features

- Keycloak integration
- Session management
- Role-based access control (RBAC)
- Auth state persistence

## Development

```bash
pnpm type-check
pnpm build
```
