# @sikerma/ui

Paket komponen UI shared untuk aplikasi SIKERMA. Berisi komponen reusable berbasis shadcn/ui dan Tailwind CSS.

## Komponen

### UI Components (shadcn/ui)
- Button
- Input
- Label
- Dialog
- AlertDialog
- Table
- Pagination

### Shared Components
- DataTable - Tabel dengan pagination, sorting, dan filtering
- FormDialog - Modal form dengan React Hook Form dan Zod
- DeleteConfirm - Konfirmasi hapus dengan AlertDialog

## Dependencies

```json
{
  "dependencies": {
    "@radix-ui/react-alert-dialog": "^1.1.2",
    "@radix-ui/react-dialog": "^1.1.2",
    "@radix-ui/react-slot": "^1.1.0",
    "class-variance-authority": "^0.7.1",
    "clsx": "^2.1.1",
    "lucide-react": "^0.468.0",
    "tailwind-merge": "^2.5.4",
    "tailwindcss-animate": "^1.0.7",
    "@hookform/resolvers": "^3.9.1",
    "react-hook-form": "^7.53.2",
    "zod": "^3.23.8"
  }
}
```

## Penggunaan

### DataTable

```tsx
import { DataTable } from "@sikerma/ui"

const columns: TableColumn<Pegawai>[] = [
  { id: "nip", header: "NIP", accessor: "nip" },
  { id: "nama", header: "Nama", accessor: "nama" },
  { id: "jabatan", header: "Jabatan", accessor: (row) => row.jabatan?.nama },
]

<DataTable
  data={pegawai}
  columns={columns}
  loading={isLoading}
  pagination={pagination}
  onPageChange={setPage}
  searchable
  onSearch={setSearch}
  onRowClick={(row) => router.push(`/pegawai/${row.id}`)}
/>
```

### FormDialog

```tsx
import { FormDialog } from "@sikerma/ui"
import { z } from "zod"

const schema = z.object({
  nama: z.string().min(1, "Nama wajib diisi"),
  nip: z.string().length(18, "NIP harus 18 digit"),
})

const fields: FormDataField[] = [
  { name: "nama", label: "Nama", type: "text", required: true },
  { name: "nip", label: "NIP", type: "text", required: true },
]

<FormDialog
  open={isOpen}
  onOpenChange={setIsOpen}
  title="Tambah Pegawai"
  fields={fields}
  schema={schema}
  onSubmit={handleSubmit}
  isLoading={isLoading}
/>
```

## Utilities

### cn()
Menggabungkan className dengan Tailwind CSS:

```tsx
import { cn } from "@sikerma/ui"

<div className={cn("base-class", condition && "conditional-class")} />
```

### formatDate()
Format tanggal ke format Indonesia:

```tsx
import { formatDate } from "@sikerma/ui"

const formatted = formatDate(new Date()) // "22 Februari 2026"
```

### formatCurrency()
Format mata uang Rupiah:

```tsx
import { formatCurrency } from "@sikerma/ui"

const formatted = formatCurrency(1500000) // "Rp 1.500.000"
```

## Setup di App Next.js

1. Import Tailwind CSS dari package:

```tsx
// app/layout.tsx
import "@sikerma/ui/src/styles/globals.css"

export default function RootLayout({ children }) {
  return (
    <html lang="id">
      <body>{children}</body>
    </html>
  )
}
```

2. Update `tailwind.config.ts` di app:

```ts
import type { Config } from "tailwindcss"

const config: Config = {
  content: [
    "./src/**/*.{ts,tsx}",
    "../../packages/ui/src/**/*.{ts,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("tailwindcss-animate")],
}

export default config
```