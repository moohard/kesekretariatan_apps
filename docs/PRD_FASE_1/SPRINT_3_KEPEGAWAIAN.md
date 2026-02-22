# SPRINT 3: Kepegawaian Dasar

**Timeline:** 8-10 hari kerja
**Status:** Blocked (Dependency: Sprint 2)
**Dependency:** Sprint 2 (Master Data harus ready untuk dropdown/referensi)

---

## Overview
Sprint 3 fokus pada pembuatan aplikasi Kepegawaian dengan fitur CRUD lengkap untuk data pegawai dan 4 jenis riwayat (pangkat, jabatan, pendidikan, keluarga). App ini adalah **core functionality** dari Fase 1.

---

## Goals Sprint 3

| ID | Goal | Metric | Target |
|----|------|--------|--------|
| G-01 | CRUD Pegawai lengkap | Tambah, lihat, edit, nonaktifkan pegawai | 100% |
| G-02 | Multi-step form wizard | Form tambah pegawai 4 step berfungsi | 100% |
| G-03 | CRUD 4 jenis riwayat | Pangkat, jabatan, pendidikan, keluarga | 100% |
| G-04 | File upload berfungsi | Foto, SK, ijazah dapat diupload | 100% |
| G-05 | Dashboard statistik | Chart & statistik menampilkan data real | 100% |
| G-06 | 29 pegawai ter-verifikasi | Data dari JSON ter-migrate dan terlihat di UI | 100% |

---

## Task Breakdown

### Task 014: Kepegawaian CRUD Backend
**Estimasi:** 3-4 hari
**FR:** FR-301 - FR-310, FR-313

**Deliverables:**

#### 1. Pegawai Handler
```go
// internal/handlers/kepegawaian/pegawai_handler.go
type PegawaiHandler struct {
  db *gorm.DB
}

func (h *PegawaiHandler) List(c *fiber.Ctx) error {
  page := c.QueryInt("page", 1)
  limit := c.QueryInt("limit", 10)
  search := c.Query("search")
  status := c.Query("status")
  golonganID := c.Query("golongan_id")
  unitKerjaID := c.Query("unit_kerja_id")

  query := h.db.Model(&Pegawai{}).
    Preload("Golongan").
    Preload("Jabatan").
    Preload("UnitKerja").
    Where("is_active = ?", true)

  // Search by NIP or nama
  if search != "" {
    query = query.Where("nip ILIKE ? OR nama ILIKE ?", "%"+search+"%", "%"+search+"%")
  }

  // Filter
  if status != "" {
    query = query.Where("status_pegawai = ?", status)
  }
  if golonganID != "" {
    query = query.Where("golongan_id = ?", golonganID)
  }
  if unitKerjaID != "" {
    query = query.Where("unit_kerja_id = ?", unitKerjaID)
  }

  // Pagination
  offset := (page - 1) * limit
  var total int64
  query.Count(&total)

  var pegawai []Pegawai
  query.Offset(offset).Limit(limit).
    Order("nama ASC").
    Find(&pegawai)

  return c.JSON(fiber.Map{
    "data": pegawai,
    "pagination": fiber.Map{
      "page": page,
      "limit": limit,
      "total": total,
      "total_pages": (total + int64(limit) - 1) / int64(limit),
    },
  })
}

func (h *PegawaiHandler) Detail(c *fiber.Ctx) error {
  nip := c.Params("nip")

  var pegawai Pegawai
  if err := h.db.Preload("Golongan").
    Preload("Jabatan").
    Preload("UnitKerja").
    Preload("Agama").
    Preload("RiwayatPangkat").
    Preload("RiwayatJabatan").
    Preload("RiwayatPendidikan").
    Preload("Keluarga").
    Where("nip = ?", nip).
    First(&pegawai).Error; err != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
  }

  return c.JSON(pegawai)
}

func (h *PegawaiHandler) Create(c *fiber.Ctx) error {
  var data CreatePegawaiRequest
  if err := c.BodyParser(&data); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Validate NIP (18 digit, unique)
  if len(data.NIP) != 18 {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "NIP harus 18 digit"})
  }
  var count int64
  h.db.Model(&Pegawai{}).Where("nip = ?", data.NIP).Count(&count)
  if count > 0 {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "NIP sudah terdaftar"})
  }

  // Create pegawai
  pegawai := Pegawai{
    NIP:              data.NIP,
    Nama:             data.Nama,
    GelarDepan:       data.GelarDepan,
    GelarBelakang:    data.GelarBelakang,
    TempatLahir:      data.TempatLahir,
    TanggalLahir:     data.TanggalLahir,
    JenisKelamin:     data.JenisKelamin,
    AgamaID:          data.AgamaID,
    StatusPegawai:    data.StatusPegawai,
    GolonganID:       data.GolonganID,
    JabatanID:        data.JabatanID,
    UnitKerjaID:      data.UnitKerjaID,
    NoHP:             data.NoHP,
    Email:            data.Email,
    FotoURL:          data.FotoURL,
    IsActive:         true,
  }

  if err := h.db.Create(&pegawai).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  // Auto-create riwayat pangkat & jabatan pertama
  if data.TMTPangkat != nil {
    h.db.Create(&RiwayatPangkat{
      NIP:            data.NIP,
      GolonganID:     data.GolonganID,
      NoSK:           data.NoSKPangkat,
      TanggalSK:      data.TanggalSKPangkat,
      TMT:            *data.TMTPangkat,
      PejabatPenetap: data.PejabatPenetap,
    })
  }

  if data.TMTJabatan != nil {
    h.db.Create(&RiwayatJabatan{
      NIP:            data.NIP,
      JabatanID:      data.JabatanID,
      UnitKerjaID:    data.UnitKerjaID,
      NoSK:           data.NoSKJabatan,
      TanggalSK:      data.TanggalSKJabatan,
      TMT:            *data.TMTJabatan,
    })
  }

  return c.Status(fiber.StatusCreated).JSON(pegawai)
}

func (h *PegawaiHandler) Update(c *fiber.Ctx) error {
  nip := c.Params("nip")
  var data UpdatePegawaiRequest
  if err := c.BodyParser(&data); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  updates := make(map[string]interface{})
  if data.Nama != "" { updates["nama"] = data.Nama }
  if data.GelarDepan != nil { updates["gelar_depan"] = data.GelarDepan }
  if data.GelarBelakang != nil { updates["gelar_belakang"] = data.GelarBelakang }
  if data.TempatLahir != nil { updates["tempat_lahir"] = data.TempatLahir }
  if data.TanggalLahir != nil { updates["tanggal_lahir"] = data.TanggalLahir }
  if data.JenisKelamin != nil { updates["jenis_kelamin"] = data.JenisKelamin }
  if data.AgamaID != nil { updates["agama_id"] = data.AgamaID }
  if data.StatusPegawai != nil { updates["status_pegawai"] = data.StatusPegawai }
  if data.GolonganID != nil { updates["golongan_id"] = data.GolonganID }
  if data.JabatanID != nil { updates["jabatan_id"] = data.JabatanID }
  if data.UnitKerjaID != nil { updates["unit_kerja_id"] = data.UnitKerjaID }
  if data.NoHP != nil { updates["no_hp"] = data.NoHP }
  if data.Email != nil { updates["email"] = data.Email }
  if data.FotoURL != nil { updates["foto_url"] = data.FotoURL }

  if err := h.db.Model(&Pegawai{}).Where("nip = ?", nip).Updates(updates).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  return c.JSON(fiber.Map{"message": "updated successfully"})
}

func (h *PegawaiHandler) SoftDelete(c *fiber.Ctx) error {
  nip := c.Params("nip")
  alasan := c.Query("alasan", "Nonaktifkan")

  // Set is_active = false
  if err := h.db.Model(&Pegawai{}).
    Where("nip = ?", nip).
    Updates(map[string]interface{}{
      "is_active": false,
      "alasan_nonaktif": alasan,
    }).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  return c.JSON(fiber.Map{"message": "pegawai nonaktif"})
}

func (h *PegawaiHandler) UploadFoto(c *fiber.Ctx) error {
  nip := c.Params("nip")

  // Get uploaded file
  file, err := c.FormFile("file")
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Validate file (max 2MB, jpg/png)
  if file.Size > 2*1024*1024 {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File maksimal 2MB"})
  }
  if !strings.HasSuffix(file.Filename, ".jpg") &&
     !strings.HasSuffix(file.Filename, ".jpeg") &&
     !strings.HasSuffix(file.Filename, ".png") {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format file harus JPG/PNG"})
  }

  // Save file
  filename := fmt.Sprintf("%s_%d%s", nip, time.Now().Unix(), filepath.Ext(file.Filename))
  filepath := fmt.Sprintf("uploads/foto/%s", filename)
  if err := c.SaveFile(file, filepath); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  // Update foto_url di database
  fotoURL := fmt.Sprintf("/uploads/foto/%s", filename)
  h.db.Model(&Pegawai{}).Where("nip = ?", nip).Update("foto_url", fotoURL)

  return c.JSON(fiber.Map{"foto_url": fotoURL})
}
```

#### 2. Riwayat Handlers (Pattern yang sama untuk 4 jenis riwayat)

**Riwayat Pangkat:**
```go
// internal/handlers/kepegawaian/riwayat_pangkat_handler.go
func (h *RiwayatPangkatHandler) List(c *fiber.Ctx) error {
  nip := c.Params("nip")
  var riwayat []RiwayatPangkat
  h.db.Preload("Golongan").
    Where("nip = ?", nip).
    Order("tmt DESC").
    Find(&riwayat)
  return c.JSON(riwayat)
}

func (h *RiwayatPangkatHandler) Create(c *fiber.Ctx) error {
  nip := c.Params("nip")
  var data CreateRiwayatPangkatRequest
  c.BodyParser(&data)

  // Upload file SK jika ada
  var fileSKURL string
  if file, err := c.FormFile("file_sk"); err == nil {
    filename := fmt.Sprintf("sk_pangkat_%s_%d.pdf", nip, time.Now().Unix())
    filepath := "uploads/sk/" + filename
    c.SaveFile(file, filepath)
    fileSKURL = "/uploads/sk/" + filename
  }

  riwayat := RiwayatPangkat{
    NIP:            nip,
    GolonganID:     data.GolonganID,
    NoSK:           data.NoSK,
    TanggalSK:      data.TanggalSK,
    TMT:            data.TMT,
    PejabatPenetap: data.PejabatPenetap,
    FileSKURL:      fileSKURL,
  }

  h.db.Create(&riwayat)
  return c.Status(fiber.StatusCreated).JSON(riwayat)
}

// Update & Delete similar pattern
```

**Riwayat Jabatan, Pendidikan, Keluarga** mengikuti pattern yang sama.

#### 3. File Upload Handler
```go
// internal/handlers/kepegawaian/file_handler.go
func UploadSK(c *fiber.Ctx) error {
  file, err := c.FormFile("file")
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Validate: max 5MB, PDF only
  if file.Size > 5*1024*1024 {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File maksimal 5MB"})
  }
  if !strings.HasSuffix(file.Filename, ".pdf") {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format file harus PDF"})
  }

  filename := fmt.Sprintf("sk_%d%s", time.Now().Unix(), filepath.Ext(file.Filename))
  filepath := "uploads/sk/" + filename
  c.SaveFile(file, filepath)

  return c.JSON(fiber.Map{
    "file_url": "/uploads/sk/" + filename,
    "filename": file.Filename,
  })
}

func UploadIjazah(c *fiber.Ctx) error {
  // Similar to UploadSK, different directory
  // ...
}
```

#### 4. Statistik Handler
```go
// internal/handlers/kepegawaian/statistik_handler.go
func GetStatistik(c *fiber.Ctx) error {
  var stats StatistikResponse

  // Total pegawai per status
  h.db.Model(&Pegawai{}).
    Where("is_active = ?", true).
    Select("status_pegawai, COUNT(*) as total").
    Group("status_pegawai").
    Scan(&stats.PerStatus)

  // Total pegawai per golongan
  h.db.Model(&Pegawai{}).
    Where("is_active = ?", true).
    Select("golongan_id, COUNT(*) as total").
    Group("golongan_id").
    Scan(&stats.PerGolongan)

  // Total pegawai per unit kerja
  h.db.Model(&Pegawai{}).
    Where("is_active = ?", true).
    Select("unit_kerja_id, COUNT(*) as total").
    Group("unit_kerja_id").
    Scan(&stats.PerUnitKerja)

  // Pegawai akan pensiun (5 terdekat)
  h.db.Model(&Pegawai{}).
    Where("is_active = ? AND tanggal_lahir IS NOT NULL", true).
    Order("tanggal_lahir DESC").
    Limit(5).
    Scan(&stats.AkanPensiun)

  return c.JSON(stats)
}
```

---

### Task 015: Kepegawaian Frontend
**Estimasi:** 5-6 hari
**FR:** FR-301 - FR-310, FR-313

**Deliverables:**

#### 1. Dashboard Page (`/`)
```tsx
// apps/kepegawaian/app/(app)/page.tsx
export default function KepegawaianDashboard() {
  const { data: stats } = useQuery({
    queryKey: ['pegawai', 'statistik'],
    queryFn: () => apiClient.get('/pegawai/statistik').then(res => res.data)
  });

  // Chart: Pegawai per Status
  const statusChart = {
    labels: stats?.perStatus.map(s => s.status_pegawai),
    datasets: [{
      label: 'Jumlah',
      data: stats?.perStatus.map(s => s.total),
      backgroundColor: ['#3b82f6', '#10b981', '#f59e0b', '#ef4444']
    }]
  };

  // Chart: Pegawai per Golongan
  const golonganChart = {
    labels: stats?.perGolongan.map(g => g.golongan_nama),
    datasets: [{
      label: 'Jumlah',
      data: stats?.perGolongan.map(g => g.total),
      backgroundColor: '#8b5cf6'
    }]
  };

  return (
    <div className="p-6 space-y-6">
      <PageHeader
        title="Kepegawaian"
        description="Kelola data pegawai Pengadilan Agama Penajam"
        actions={[
          {
            label: "Tambah Pegawai",
            onClick: () => router.push('/pegawai/tambah')
          }
        ]}
      />

      {/* Statistik Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between">
            <CardTitle>Total Pegawai</CardTitle>
            <Users className="h-5 w-5 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{stats?.total}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between">
            <CardTitle>PNS</CardTitle>
            <BadgeCheck className="h-5 w-5 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{stats?.pns}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between">
            <CardTitle>CPNS</CardTitle>
            <Clock className="h-5 w-5 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{stats?.cpns}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between">
            <CardTitle>PPPK</CardTitle>
            <FileText className="h-5 w-5 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{stats?.pppk}</div>
          </CardContent>
        </Card>
      </div>

      {/* Charts */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <CardTitle>Pegawai per Status</CardTitle>
          </CardHeader>
          <CardContent>
            <BarChart data={statusChart} />
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Pegawai per Golongan</CardTitle>
          </CardHeader>
          <CardContent>
            <BarChart data={golonganChart} />
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
```

#### 2. Daftar Pegawai Page (`/pegawai`)
```tsx
// apps/kepegawaian/app/(app)/pegawai/page.tsx
const columns = [
  {
    accessorKey: 'foto',
    header: 'Foto',
    cell: ({ getValue }) => (
      <Avatar className="h-10 w-10">
        <AvatarImage src={getValue()} alt="Foto" />
        <AvatarFallback>
          <User className="h-5 w-5" />
        </AvatarFallback>
      </Avatar>
    )
  },
  { accessorKey: 'nip', header: 'NIP', cell: ({ getValue }) => formatNIP(getValue()) },
  { accessorKey: 'nama', header: 'Nama' },
  { accessorKey: 'jabatan_nama', header: 'Jabatan' },
  { accessorKey: 'golongan_kode', header: 'Golongan' },
  { accessorKey: 'unit_kerja_nama', header: 'Unit Kerja' },
  {
    accessorKey: 'status_pegawai',
    header: 'Status',
    cell: ({ getValue }) => (
      <StatusBadge
        status={getValue()}
        variant={
          getValue() === 'PNS' ? 'success' :
          getValue() === 'CPNS' ? 'warning' :
          getValue() === 'PPPK' ? 'info' : 'secondary'
        }
      />
    )
  }
];

export default function DaftarPegawaiPage() {
  const router = useRouter();
  const [filters, setFilters] = useState({
    search: '',
    status: '',
    golongan_id: '',
    unit_kerja_id: ''
  });

  const { data, isLoading } = useQuery({
    queryKey: ['pegawai', filters],
    queryFn: () => apiClient.get('/pegawai', { params: filters }).then(res => res.data)
  });

  return (
    <div className="p-6">
      <PageHeader
        title="Daftar Pegawai"
        description="Kelola data pegawai dan riwayatnya"
        actions={[
          { label: "Tambah Pegawai", onClick: () => router.push('/pegawai/tambah') }
        ]}
      />

      {/* Filters */}
      <div className="mb-4 grid grid-cols-1 md:grid-cols-4 gap-4">
        <Input
          placeholder="Cari NIP atau nama..."
          value={filters.search}
          onChange={(e) => setFilters({ ...filters, search: e.target.value })}
        />
        <Select
          value={filters.status}
          onValueChange={(value) => setFilters({ ...filters, status: value })}
        >
          <SelectTrigger>
            <SelectValue placeholder="Filter Status" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="">Semua Status</SelectItem>
            <SelectItem value="PNS">PNS</SelectItem>
            <SelectItem value="CPNS">CPNS</SelectItem>
            <SelectItem value="PPPK">PPPK</SelectItem>
            <SelectItem value="Honorer">Honorer</SelectItem>
          </SelectContent>
        </Select>
        <Select
          value={filters.golongan_id}
          onValueChange={(value) => setFilters({ ...filters, golongan_id: value })}
        >
          <SelectTrigger>
            <SelectValue placeholder="Filter Golongan" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="">Semua Golongan</SelectItem>
            {/* Dynamic dari API dropdown */}
          </SelectContent>
        </Select>
        <Select
          value={filters.unit_kerja_id}
          onValueChange={(value) => setFilters({ ...filters, unit_kerja_id: value })}
        >
          <SelectTrigger>
            <SelectValue placeholder="Filter Unit Kerja" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="">Semua Unit</SelectItem>
            {/* Dynamic dari API dropdown */}
          </SelectContent>
        </Select>
      </div>

      <DataTable
        columns={columns}
        data={data?.data || []}
        pagination={data?.pagination}
        searchable={false} // Already have search filter above
        sortable
        onRowClick={(row) => router.push(`/pegawai/${row.nip}`)}
        actions={[
          {
            label: "Lihat Detail",
            onClick: (row) => router.push(`/pegawai/${row.nip}`)
          },
          {
            label: "Edit",
            onClick: (row) => router.push(`/pegawai/${row.nip}/edit`)
          },
          {
            label: "Nonaktifkan",
            onClick: (row) => {
              DeleteConfirm({
                title: "Nonaktifkan Pegawai",
                message: `Yakin ingin menonaktifkan ${row.nama}?`,
                onConfirm: () => apiClient.delete(`/pegawai/${row.nip}`)
              });
            }
          }
        ]}
      />
    </div>
  );
}
```

#### 3. Tambah Pegawai Page - Multi-Step Wizard (`/pegawai/tambah`)
```tsx
// apps/kepegawaian/app/(app)/pegawai/tambah/page.tsx
export default function TambahPegawaiPage() {
  const router = useRouter();
  const [currentStep, setCurrentStep] = useState(0);

  const steps = [
    { label: "Biodata", component: BiodataStep },
    { label: "Posisi & Pangkat", component: PosisiPangkatStep },
    { label: "Pendidikan", component: PendidikanStep },
    { label: "Keluarga", component: KeluargaStep },
  ];

  const { mutate: createPegawai, isPending } = useMutation({
    mutationFn: (data: any) => apiClient.post('/pegawai', data),
    onSuccess: () => {
      router.push('/pegawai');
    }
  });

  const CurrentStepComponent = steps[currentStep].component;

  return (
    <div className="p-6 max-w-4xl mx-auto">
      <PageHeader
        title="Tambah Pegawai Baru"
        description="Isi data pegawai secara bertahap"
      />

      {/* Step Progress */}
      <div className="mb-8">
        <StepWizard
          steps={steps.map(s => s.label)}
          currentStep={currentStep}
          onStepClick={setCurrentStep}
        />
      </div>

      {/* Step Form */}
      <CurrentStepComponent
        onNext={(data) => {
          // Store data in context/state
          if (currentStep === steps.length - 1) {
            createPegawai(data); // Submit all data
          } else {
            setCurrentStep(currentStep + 1);
          }
        }}
        onBack={() => setCurrentStep(currentStep - 1)}
        isFinalStep={currentStep === steps.length - 1}
        isPending={isPending}
      />
    </div>
  );
}

// Step 1: Biodata
function BiodataStep({ onNext, onBack, isFinalStep, isPending }: StepProps) {
  const form = useForm({
    defaultValues: {
      nip: '',
      nama: '',
      gelar_depan: '',
      gelar_belakang: '',
      tempat_lahir: '',
      tanggal_lahir: '',
      jenis_kelamin: '',
      agama_id: '',
      no_hp: '',
      email: '',
    }
  });

  const { data: agamaDropdown } = useQuery({
    queryKey: ['master', 'agama', 'dropdown'],
    queryFn: () => apiClient.get('/master/agama/dropdown').then(res => res.data)
  });

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onNext)} className="space-y-6">
        <FormField
          control={form.control}
          name="nip"
          rules={{
            required: "NIP wajib diisi",
            minLength: { value: 18, message: "NIP harus 18 digit" },
            maxLength: { value: 18, message: "NIP harus 18 digit" },
            pattern: { value: /^\d+$/, message: "NIP hanya boleh angka" }
          }}
          render={({ field }) => (
            <FormItem>
              <FormLabel>NIP <span className="text-red-500">*</span></FormLabel>
              <FormControl>
                <Input placeholder="Contoh: 199011222010121001" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="nama"
          rules={{ required: "Nama wajib diisi" }}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Nama Lengkap <span className="text-red-500">*</span></FormLabel>
              <FormControl>
                <Input placeholder="Nama lengkap sesuai KTP" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        {/* ... form fields lainnya */}

        <div className="flex justify-between">
          {currentStep > 0 && (
            <Button type="button" variant="outline" onClick={onBack}>
              <ArrowLeft className="mr-2 h-4 w-4" /> Sebelumnya
            </Button>
          )}
          <Button type="submit" disabled={isPending}>
            {isFinalStep ? "Simpan Pegawai" : "Selanjutnya"}
            {!isFinalStep && <ArrowRight className="ml-2 h-4 w-4" />}
          </Button>
        </div>
      </form>
    </Form>
  );
}
```

#### 4. Detail Pegawai Page - Tab View (`/pegawai/[nip]`)
```tsx
// apps/kepegawaian/app/(app)/pegawai/[nip]/page.tsx
export default function DetailPegawaiPage() {
  const params = useParams();
  const nip = params.nip as string;

  const { data: pegawai } = useQuery({
    queryKey: ['pegawai', nip],
    queryFn: () => apiClient.get(`/pegawai/${nip}`).then(res => res.data)
  });

  return (
    <div className="p-6">
      <Breadcrumb
        items={[
          { label: "Kepegawaian", href: "/pegawai" },
          { label: pegawai?.nama || "Detail Pegawai" }
        ]}
      />

      {/* Header Card */}
      <Card className="mb-6">
        <CardContent className="pt-6">
          <div className="flex items-start gap-6">
            <Avatar className="h-24 w-24">
              <AvatarImage src={pegawai?.foto_url} alt={pegawai?.nama} />
              <AvatarFallback className="text-2xl">
                {pegawai?.nama?.charAt(0)}
              </AvatarFallback>
            </Avatar>

            <div className="flex-1">
              <div className="flex items-center gap-2">
                <h1 className="text-2xl font-bold">{pegawai?.nama}</h1>
                <StatusBadge
                  status={pegawai?.status_pegawai || ''}
                  variant={
                    pegawai?.status_pegawai === 'PNS' ? 'success' :
                    pegawai?.status_pegawai === 'CPNS' ? 'warning' :
                    pegawai?.status_pegawai === 'PPPK' ? 'info' : 'secondary'
                  }
                />
              </div>

              <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mt-4 text-sm">
                <div>
                  <div className="text-muted-foreground">NIP</div>
                  <div className="font-medium">{formatNIP(pegawai?.nip || '')}</div>
                </div>
                <div>
                  <div className="text-muted-foreground">Jabatan</div>
                  <div className="font-medium">{pegawai?.jabatan?.nama}</div>
                </div>
                <div>
                  <div className="text-muted-foreground">Golongan</div>
                  <div className="font-medium">{pegawai?.golongan?.kode}</div>
                </div>
                <div>
                  <div className="text-muted-foreground">Unit Kerja</div>
                  <div className="font-medium">{pegawai?.unit_kerja?.nama}</div>
                </div>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Tabs */}
      <Tabs defaultValue="biodata" className="space-y-4">
        <TabsList>
          <TabsTrigger value="biodata">Biodata</TabsTrigger>
          <TabsTrigger value="pangkat">Riwayat Pangkat</TabsTrigger>
          <TabsTrigger value="jabatan">Riwayat Jabatan</TabsTrigger>
          <TabsTrigger value="pendidikan">Riwayat Pendidikan</TabsTrigger>
          <TabsTrigger value="keluarga">Data Keluarga</TabsTrigger>
        </TabsList>

        <TabsContent value="biodata">
          <BiodataTab pegawai={pegawai} />
        </TabsContent>

        <TabsContent value="pangkat">
          <RiwayatPangkatTab nip={nip} />
        </TabsContent>

        <TabsContent value="jabatan">
          <RiwayatJabatanTab nip={nip} />
        </TabsContent>

        <TabsContent value="pendidikan">
          <RiwayatPendidikanTab nip={nip} />
        </TabsContent>

        <TabsContent value="keluarga">
          <KeluargaTab nip={nip} />
        </TabsContent>
      </Tabs>
    </div>
  );
}

// Riwayat Pangkat Tab
function RiwayatPangkatTab({ nip }: { nip: string }) {
  const { data: riwayat, refetch } = useQuery({
    queryKey: ['pegawai', nip, 'pangkat'],
    queryFn: () => apiClient.get(`/pegawai/${nip}/pangkat`).then(res => res.data)
  });

  const { mutate: createRiwayat, isPending } = useMutation({
    mutationFn: (data: FormData) => {
      return apiClient.post(`/pegawai/${nip}/pangkat`, data, {
        headers: { 'Content-Type': 'multipart/form-data' }
      });
    },
    onSuccess: () => refetch()
  });

  return (
    <div className="space-y-4">
      <div className="flex justify-end">
        <Button onClick={() => setShowDialog(true)}>
          <Plus className="mr-2 h-4 w-4" /> Tambah Riwayat
        </Button>
      </div>

      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Golongan</TableHead>
            <TableHead>No SK</TableHead>
            <TableHead>Tanggal SK</TableHead>
            <TableHead>TMT</TableHead>
            <TableHead>Pejabat Penetap</TableHead>
            <TableHead>Aksi</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {riwayat?.map((item: any) => (
            <TableRow key={item.id}>
              <TableCell>{item.golongan?.kode}</TableCell>
              <TableCell>{item.no_sk}</TableCell>
              <TableCell>{formatTanggal(item.tanggal_sk)}</TableCell>
              <TableCell>{formatTanggal(item.tmt)}</TableCell>
              <TableCell>{item.pejabat_penetap}</TableCell>
              <TableCell>
                <Button variant="ghost" size="sm">
                  <Download className="h-4 w-4" />
                </Button>
                <Button variant="ghost" size="sm">
                  <Trash2 className="h-4 w-4" />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
```

---

## Definition of Done - Sprint 3

Sprint 3 dianggap **DONE** jika:

- ✅ Backend CRUD pegawai berfungsi (validasi NIP 18 digit, unique)
- ✅ Backend CRUD 4 jenis riwayat berfungsi
- ✅ File upload (foto max 2MB, SK/ijazah max 5MB PDF) berfungsi
- ✅ Frontend daftar pegawai dengan filter & search berfungsi
- ✅ Multi-step wizard tambah pegawai berfungsi (4 step)
- ✅ Detail pegawai dengan tab view (5 tabs) berfungsi
- ✅ CRUD inline di tab riwayat berfungsi
- ✅ Dashboard statistik dengan chart menampilkan data real
- ✅ 29 pegawai dari JSON terlihat di UI dan datanya benar
- ✅ Tidak ada console errors atau TypeScript errors

---

## Risks & Mitigations

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Multi-step form state management kompleks | Medium | Medium | Gunakan React Context atau Zustand untuk state management |
| File upload dengan FormData di TanStack Query tricky | Medium | Medium | Test upload flow di awal sprint, siapkan helper function |
| Chart library (recharts) learning curve | Low | Low | Pakai komponen chart yang sudah ada di shadcn atau buat custom simple |
| Data 29 pegawai tidak lengkap di JSON awal | Low | Medium | Admin bisa lengkapi manual via UI, tidak blocking |

---

## Success Metrics

| Metric | Target | Actual |
|--------|--------|--------|
| Pegawai CRUD endpoints working | 5/5 (GET list, GET detail, POST, PUT, DELETE) | ___ |
| Riwayat CRUD endpoints working | 16/16 (4 riwayat × 4 operations) | ___ |
| File upload endpoints working | 3/3 (foto, SK, ijazah) | ___ |
| Frontend pages complete | 4/4 (dashboard, daftar, tambah, detail) | ___ |
| 29 pegawai visible in UI | 29/29 | ___ |
| No critical bugs | 0 | ___ |
