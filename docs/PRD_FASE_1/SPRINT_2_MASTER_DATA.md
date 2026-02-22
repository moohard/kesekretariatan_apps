# SPRINT 2: Master Data (CRUD Lengkap)

**Timeline:** 6-8 hari kerja
**Status:** Blocked (Dependency: Sprint 1)
**Dependency:** Sprint 1 (Backend + Auth + UI Components)

---

## Overview
Sprint 2 fokus pada pembuatan aplikasi Master Data dengan CRUD lengkap untuk 10 entitas referensi. App ini akan menjadi **single source of truth** untuk semua data referensi yang dipakai oleh app lain.

---

## Goals Sprint 2

| ID | Goal | Metric | Target |
|----|------|--------|--------|
| G-01 | Generic CRUD backend | 1 pattern handler untuk semua entitas | 100% |
| G-02 | CRUD frontend 10 entitas | Semua entitas bisa Create, Read, Update, Delete | 100% |
| G-03 | Dropdown API berfungsi | Setiap entitas punya endpoint `/dropdown` | 100% |
| G-04 | Data referensi ter-seed | Semua data dari existing ter-migrate | 100% |
| G-05 | UI konsisten & responsive | Pakai shared components, UX smooth | Fully functional |

---

## Task Breakdown

### Task 012: Generic CRUD Backend Master Data
**Estimasi:** 2 hari
**FR:** FR-201 - FR-212

**Deliverables:**

Backend Handler Pattern (Reusable):
```go
// internal/handlers/master/generic_handler.go
type MasterEntity interface {
  GetTableName() string
  GetPrimaryKey() string
  Validate(data map[string]interface{}) error
}

type GenericHandler struct {
  entity MasterEntity
}

func (h *GenericHandler) List(c *fiber.Ctx) error {
  // Pagination, search, sort
  page := c.QueryInt("page", 1)
  limit := c.QueryInt("limit", 10)
  search := c.Query("search")
  sortBy := c.Query("sort", "created_at")
  sortOrder := c.Query("order", "desc")

  // Query builder dengan filter search
  query := db.Model(h.entity)
  if search != "" {
    query = query.Where("nama ILIKE ?", "%"+search+"%")
  }

  // Pagination
  offset := (page - 1) * limit
  var total int64
  query.Count(&total)
  query.Offset(offset).Limit(limit).Order(sortBy + " " + sortOrder)

  var results []map[string]interface{}
  query.Find(&results)

  return c.JSON(fiber.Map{
    "data": results,
    "pagination": fiber.Map{
      "page": page,
      "limit": limit,
      "total": total,
      "total_pages": (total + int64(limit) - 1) / int64(limit),
    },
  })
}

func (h *GenericHandler) Get(c *fiber.Ctx) error {
  id := c.Params("id")
  var result map[string]interface{}
  if err := db.Model(h.entity).Where("id = ?", id).First(&result).Error; err != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
  }
  return c.JSON(result)
}

func (h *GenericHandler) Create(c *fiber.Ctx) error {
  var data map[string]interface{}
  if err := c.BodyParser(&data); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Validate
  if err := h.entity.Validate(data); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Insert
  if err := db.Model(h.entity).Create(&data).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  return c.Status(fiber.StatusCreated).JSON(data)
}

func (h *GenericHandler) Update(c *fiber.Ctx) error {
  id := c.Params("id")
  var data map[string]interface{}
  if err := c.BodyParser(&data); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Validate
  if err := h.entity.Validate(data); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Update (exclude ID)
  delete(data, "id")
  if err := db.Model(h.entity).Where("id = ?", id).Updates(data).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  return c.JSON(fiber.Map{"message": "updated successfully"})
}

func (h *GenericHandler) Delete(c *fiber.Ctx) error {
  id := c.Params("id")

  // Soft delete - set is_active = false
  if err := db.Model(h.entity).Where("id = ?", id).Update("is_active", false).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  return c.JSON(fiber.Map{"message": "deleted successfully"})
}

func (h *GenericHandler) Dropdown(c *fiber.Ctx) error {
  var results []map[string]interface{}
  db.Model(h.entity).
    Where("is_active = ?", true).
    Select("id, nama").
    Order("nama ASC").
    Find(&results)

  return c.JSON(results)
}
```

**Routes Setup:**
```go
// internal/routes/master_routes.go
func SetupMasterRoutes(router fiber.Router, rbacMiddleware fiber.Handler) {
  entities := []struct {
    name    string
    handler *GenericHandler
  }{
    {"satker", NewGenericHandler(&Satker{})},
    {"jabatan", NewGenericHandler(&Jabatan{})},
    {"golongan", NewGenericHandler(&Golongan{})},
    {"unit-kerja", NewGenericHandler(&UnitKerja{})},
    {"eselon", NewGenericHandler(&Eselon{})},
    {"pendidikan", NewGenericHandler(&RefPendidikan{})},
    {"agama", NewGenericHandler(&RefAgama{})},
    {"status-kawin", NewGenericHandler(&RefStatusKawin{})},
    {"hukuman-disiplin", NewGenericHandler(&RefJenisHukdis{})},
    {"jenis-diklat", NewGenericHandler(&RefJenisDiklat{})},
  }

  for _, entity := range entities {
    entityRouter := router.Group("/" + entity.name)

    entityRouter.Get("/", rbacMiddleware(entity.name+".view"), entity.handler.List)
    entityRouter.Get("/:id", rbacMiddleware(entity.name+".view"), entity.handler.Get)
    entityRouter.Post("/", rbacMiddleware(entity.name+".create"), entity.handler.Create)
    entityRouter.Put("/:id", rbacMiddleware(entity.name+".update"), entity.handler.Update)
    entityRouter.Delete("/:id", rbacMiddleware(entity.name+".delete"), entity.handler.Delete)
    entityRouter.Get("/dropdown", entity.handler.Dropdown) // Public dropdown
  }
}
```

**Acceptance Criteria:**
- ✅ Generic handler dapat handle semua 10 entitas
- ✅ Pagination, search, sort berfungsi di semua list endpoint
- ✅ Soft delete (set is_active = false) berfungsi
- ✅ Dropdown endpoint return id + nama saja
- ✅ Validation berfungsi (required fields, format)
- ✅ Error handling jelas dan konsisten

---

### Task 013: Master Data Frontend
**Estimasi:** 4-6 hari
**FR:** FR-201 - FR-212

**Deliverables:**

#### 1. Layout & Navigation
```
apps/master-data/
├── app/
│   └── (app)/
│       ├── layout.tsx              # App layout dengan Sidebar + AppHeader
│       ├── page.tsx                # Dashboard ringkasan
│       └── master/
│           ├── layout.tsx          # Master layout dengan breadcrumb
│           ├── page.tsx            # Redirect atau overview
│           ├── satker/
│           │   └── page.tsx        # CRUD Satker
│           ├── jabatan/
│           │   └── page.tsx        # CRUD Jabatan
│           ├── golongan/
│           │   └── page.tsx        # CRUD Golongan
│           ├── unit-kerja/
│           │   └── page.tsx        # CRUD Unit Kerja (tree view)
│           ├── eselon/
│           │   └── page.tsx
│           ├── pendidikan/
│           │   └── page.tsx
│           ├── agama/
│           │   └── page.tsx
│           ├── status-kawin/
│           │   └── page.tsx
│           ├── hukuman-disiplin/
│           │   └── page.tsx
│           └── jenis-diklat/
│               └── page.tsx
```

#### 2. Dashboard Page (`/`)
```tsx
// apps/master-data/app/(app)/page.tsx
export default function MasterDashboard() {
  const { data: summary } = useQuery({
    queryKey: ['master-summary'],
    queryFn: () => apiClient.get('/master/summary').then(res => res.data)
  });

  const cards = [
    { title: "Satuan Kerja", value: summary?.satker || 0, icon: Building2 },
    { title: "Jabatan", value: summary?.jabatan || 0, icon: Briefcase },
    { title: "Golongan", value: summary?.golongan || 0, icon: Award },
    { title: "Unit Kerja", value: summary?.unit_kerja || 0, icon: Users },
    // ... 6 cards lainnya
  ];

  return (
    <div className="p-6 space-y-6">
      <PageHeader
        title="Master Data"
        description="Kelola data referensi untuk seluruh aplikasi SIKERMA"
      />

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        {cards.map((card) => (
          <Card key={card.title}>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">
                {card.title}
              </CardTitle>
              <card.icon className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{card.value}</div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
```

#### 3. Generic CRUD Page Template
```tsx
// apps/master-data/components/master-crud-page.tsx
interface CRUDPageProps {
  entity: string;
  entityLabel: string;
  columns: ColumnDef<any>[];
  formFields: FormField[];
}

export function MasterCRUDPage({
  entity,
  entityLabel,
  columns,
  formFields,
}: CRUDPageProps) {
  const [dialogOpen, setDialogOpen] = useState(false);
  const [editingItem, setEditingItem] = useState<any>(null);

  const { data, isLoading, refetch } = useQuery({
    queryKey: ['master', entity],
    queryFn: () => apiClient.get(`/master/${entity}`).then(res => res.data)
  });

  const { mutate: createItem } = useMutation({
    mutationFn: (data: any) => apiClient.post(`/master/${entity}`, data),
    onSuccess: () => {
      refetch();
      setDialogOpen(false);
    }
  });

  const { mutate: updateItem } = useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      apiClient.put(`/master/${entity}/${id}`, data),
    onSuccess: () => {
      refetch();
      setDialogOpen(false);
      setEditingItem(null);
    }
  });

  const { mutate: deleteItem } = useMutation({
    mutationFn: (id: string) => apiClient.delete(`/master/${entity}/${id}`),
    onSuccess: () => refetch()
  });

  return (
    <div className="p-6 space-y-6">
      <PageHeader
        title={entityLabel}
        description={`Kelola data ${entityLabel.toLowerCase()}`}
        actions={[
          {
            label: "Tambah",
            onClick: () => {
              setEditingItem(null);
              setDialogOpen(true);
            }
          }
        ]}
      />

      <DataTable
        columns={columns}
        data={data?.data || []}
        pagination={data?.pagination}
        searchable
        sortable
        actions={[
          {
            label: "Edit",
            onClick: (row) => {
              setEditingItem(row);
              setDialogOpen(true);
            }
          },
          {
            label: "Hapus",
            onClick: (row) => {
              DeleteConfirm({
                title: "Hapus Data",
                message: `Yakin ingin menghapus ${entityLabel.toLowerCase()} ini?`,
                onConfirm: () => deleteItem(row.id)
              });
            }
          }
        ]}
      />

      <FormDialog
        open={dialogOpen}
        onOpenChange={setDialogOpen}
        title={editingItem ? `Edit ${entityLabel}` : `Tambah ${entityLabel}`}
        fields={formFields}
        onSubmit={(values) => {
          if (editingItem) {
            updateItem({ id: editingItem.id, data: values });
          } else {
            createItem(values);
          }
        }}
        initialValues={editingItem || {}}
      />
    </div>
  );
}
```

#### 4. Per-Entity Pages (10 halaman)

Setiap entity punya halaman dengan konfigurasi spesifik:

**Contoh: Satker Page**
```tsx
// apps/master-data/app/(app)/master/satker/page.tsx
const columns = [
  { accessorKey: 'kode', header: 'Kode' },
  { accessorKey: 'nama', header: 'Nama Satker' },
  { accessorKey: 'alamat', header: 'Alamat' },
  { accessorKey: 'telepon', header: 'Telepon' },
  {
    accessorKey: 'is_active',
    header: 'Status',
    cell: ({ getValue }) => (
      <StatusBadge
        status={getValue() ? 'active' : 'inactive'}
        variant={getValue() ? 'success' : 'secondary'}
      />
    )
  }
];

const formFields = [
  { name: 'kode', label: 'Kode', type: 'text', required: true },
  { name: 'nama', label: 'Nama Satker', type: 'text', required: true },
  { name: 'alamat', label: 'Alamat', type: 'textarea' },
  { name: 'telepon', label: 'Telepon', type: 'text' },
  { name: 'email', label: 'Email', type: 'email' },
  {
    name: 'tipe',
    label: 'Tipe',
    type: 'select',
    options: [
      { value: 'pa', label: 'Pengadilan Agama' },
      { value: 'pn', label: 'Pengadilan Negeri' },
      { value: 'pt', label: 'Pengadilan Tinggi' }
    ]
  }
];

export default function SatkerPage() {
  return (
    <MasterCRUDPage
      entity="satker"
      entityLabel="Satuan Kerja"
      columns={columns}
      formFields={formFields}
    />
  );
}
```

**Contoh: Unit Kerja (Tree View)**
```tsx
// apps/master-data/app/(app)/master/unit-kerja/page.tsx
// Special page dengan tree view untuk hierarki
export default function UnitKerjaPage() {
  const { data } = useQuery({
    queryKey: ['master', 'unit-kerja', 'tree'],
    queryFn: () => apiClient.get('/master/unit-kerja/tree').then(res => res.data)
  });

  return (
    <div className="p-6">
      <PageHeader
        title="Unit Kerja"
        description="Struktur organisasi hierarki"
        actions={[{ label: "Tambah Root", onClick: () => {/*...*/} }]}
      />

      <Tree
        data={data}
        renderItem={(node) => (
          <div className="flex items-center justify-between p-2 hover:bg-muted">
            <span>{node.nama}</span>
            <div>
              <Button variant="ghost" size="sm" onClick={() => {/* Edit */}}>
                <Pencil className="h-4 w-4" />
              </Button>
              <Button variant="ghost" size="sm" onClick={() => {/* Add Child */}}>
                <Plus className="h-4 w-4" />
              </Button>
            </div>
          </div>
        )}
      />
    </div>
  );
}
```

#### 5. Sidebar Menu Configuration
```tsx
// apps/master-data/components/sidebar-menu.tsx
export const sidebarMenu = [
  {
    title: "Dashboard",
    href: "/",
    icon: LayoutDashboard,
  },
  {
    title: "Master Data",
    icon: Database,
    items: [
      { title: "Satuan Kerja", href: "/master/satker", icon: Building2 },
      { title: "Jabatan", href: "/master/jabatan", icon: Briefcase },
      { title: "Golongan/Pangkat", href: "/master/golongan", icon: Award },
      { title: "Unit Kerja", href: "/master/unit-kerja", icon: Users },
      { title: "Eselon", href: "/master/eselon", icon: GitBranch },
      { title: "Pendidikan", href: "/master/pendidikan", icon: GraduationCap },
      { title: "Agama", href: "/master/agama", icon: Book },
      { title: "Status Kawin", href: "/master/status-kawin", icon: Heart },
      { title: "Hukuman Disiplin", href: "/master/hukuman-disiplin", icon: AlertTriangle },
      { title: "Jenis Diklat", href: "/master/jenis-diklat", icon: Calendar },
    ],
  },
];
```

---

## Definition of Done - Sprint 2

Sprint 2 dianggap **DONE** jika:

- ✅ Backend generic CRUD handler berfungsi untuk semua 10 entitas
- ✅ Semua endpoint API (GET, POST, PUT, DELETE, /dropdown) berfungsi
- ✅ Frontend 10 halaman CRUD terbuat dan berfungsi
- ✅ DataTable dengan pagination, search, sort berfungsi di semua halaman
- ✅ Form dialog (create/edit) berfungsi dengan validasi
- ✅ Soft delete berfungsi (data tidak hilang, hanya is_active=false)
- ✅ Dropdown API dapat diakses oleh app lain (Kepegawaian)
- ✅ Data referensi ter-seed lengkap (satker, golongan, jabatan, unit kerja)
- ✅ UI konsisten pakai shared components (@sikerma/ui)
- ✅ Tidak ada console errors atau TypeScript errors

---

## Entity Specification

### 1. Satker (Satuan Kerja)
**Fields:**
- `id` (UUID, PK)
- `kode` (VARCHAR 20, unique) - contoh: "PA-PNJ"
- `nama` (VARCHAR 255) - contoh: "Pengadilan Agama Penajam"
- `alamat` (TEXT)
- `telepon` (VARCHAR 20)
- `email` (VARCHAR 100)
- `tipe` (VARCHAR 10) - enum: pa, pn, pt, pta, ma
- `is_active` (BOOLEAN)
- `created_at`, `updated_at`

**Dropdown Return:**
```json
[
  { "id": "uuid-1", "nama": "Pengadilan Agama Penajam" },
  { "id": "uuid-2", "nama": "Pengadilan Agama Balikpapan" }
]
```

### 2. Jabatan
**Fields:**
- `id` (UUID)
- `kode` (VARCHAR 20, unique)
- `nama` (VARCHAR 255) - contoh: "Ketua Pengadilan Tingkat Pertama Klas II"
- `jenis` (VARCHAR 20) - enum: struktural, fungsional
- `eselon` (VARCHAR 10) - contoh: "IV/a"
- `kelas_jabatan` (VARCHAR 10) - contoh: "Kelas II"
- `is_active` (BOOLEAN)

### 3. Golongan/Pangkat
**Fields:**
- `id` (UUID)
- `kode` (VARCHAR 10, unique) - contoh: "III/a", "IV/a"
- `nama_pangkat` (VARCHAR 100) - contoh: "Penata Muda"
- `ruang` (VARCHAR 5) - contoh: "III/a"
- `tingkat` (INT) - contoh: 3 (untuk III/a)
- `is_active` (BOOLEAN)

### 4. Unit Kerja
**Fields:**
- `id` (UUID)
- `satker_id` (UUID, FK → satker.id)
- `kode` (VARCHAR 20, unique)
- `nama` (VARCHAR 255)
- `parent_id` (UUID, nullable, FK → unit_kerja.id) - untuk hierarki
- `is_active` (BOOLEAN)

**Tree Structure:**
```
Pengadilan Agama Penajam (root, parent_id=NULL)
├── Pimpinan (parent_id=root)
├── Panitera (parent_id=root)
│   ├── Panitera Muda Permohonan
│   ├── Panitera Muda Gugatan
│   └── Panitera Muda Hukum
└── Sekretariat (parent_id=root)
    ├── Subbag Perencanaan, TI, dan Pelaporan
    ├── Subbag Umum dan Keuangan
    └── Subbag Kepegawaian, Organisasi, dan Tata Laksana
```

### 5-10. Entitas Lain
Lihat `TASK_013_MASTER_FRONTEND.md` untuk detail lengkap.

---

## Risks & Mitigations

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Generic handler terlalu generic, susah customize per entity | Medium | Medium | Siapkan override pattern untuk entity spesifik |
| Tree view Unit Kerja kompleks diimplementasi | Medium | Medium | Gunakan library react-arborist atau react-sortable-tree |
| Dropdown API dipanggil terlalu sering dari app lain | Low | Low | Implementasi caching di frontend (TanStack Query) |
| Data seed tidak lengkap untuk semua entity | Low | Low | Admin bisa tambah manual via UI setelah app jadi |

---

## Success Metrics

| Metric | Target | Actual |
|--------|--------|--------|
| Backend endpoints working | 60/60 (6 per entity × 10) | ___ |
| Frontend pages complete | 10/10 entities | ___ |
| Data seeded correctly | 100% | ___ |
| Dropdown API accessible | 10/10 entities | ___ |
| No critical bugs | 0 | ___ |
