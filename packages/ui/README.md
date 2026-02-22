# @sikerma/ui - Shared UI Components

Collection of reusable UI components for SIKERMA application.

## Installation

```bash
pnpm install @sikerma/ui
```

## Usage

```tsx
import { Button, Card } from '@sikerma/ui'

export default function Example() {
  return (
    <Card>
      <Button>Hello World</Button>
    </Card>
  )
}
```

## Development

```bash
# Run type checking
pnpm type-check

# Build
pnpm build
```

## Components

- `Button` - Reusable button component
- `Card` - Card container component
- `Input` - Form input component
- `Select` - Form select component
- `Table` - Data table component
- `Modal` - Modal dialog component
- `Toast` - Toast notification component
- `Badge` - Badge/label component
- `Avatar` - Avatar component
- `Skeleton` - Loading skeleton component
