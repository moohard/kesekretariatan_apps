# @sikerma/shared - Shared Utilities

Shared utilities, types, and constants for SIKERMA application.

## Installation

```bash
pnpm install @sikerma/shared
```

## Usage

```tsx
import { formatDate, validateEmail } from '@sikerma/shared'

const dateStr = formatDate(new Date())
const isValid = validateEmail('test@example.com')
```

## Utilities

- `formatDate` - Format date to readable string
- `validateEmail` - Validate email format
- `generateId` - Generate unique ID
- `debounce` - Debounce function
- `throttle` - Throttle function

## Types

```ts
interface User {
  id: string
  name: string
  email: string
  roles: string[]
}

interface ApiResponse<T> {
  data: T
  message: string
  status: number
}
```

## Development

```bash
pnpm type-check
pnpm build
```
