# SPRINT 4: Portal + Cetak SK + Polish

**Timeline:** 6-8 hari kerja
**Status:** Blocked (Dependency: Sprint 1, Sprint 3)
**Dependency:** Sprint 1 (Auth), Sprint 3 (Data Pegawai untuk Widget)

---

## Overview
Sprint 4 adalah sprint final Fase 1 yang fokus pada:
1. **Portal** sebagai launcher & dashboard utama
2. **Admin UI** untuk user management, RBAC, audit log
3. **Cetak SK** dengan template management & Gotenberg integration
4. **Polish & Bug Fixing** untuk seluruh aplikasi

---

## Goals Sprint 4

| ID | Goal | Metric | Target |
|----|------|--------|--------|
| G-01 | Portal launcher berfungsi | Tile app muncul sesuai hak akses user | 100% |
| G-02 | Dashboard widgets informatif | 4+ widget menampilkan data real | ≥ 4 |
| G-03 | Admin user management | List & manage users dari Keycloak | Fully functional |
| G-04 | Admin RBAC UI | Assign role & permission dari UI | Fully functional |
| G-05 | Audit log viewer | Lihat & filter audit logs | Fully functional |
| G-06 | Cetak SK berfungsi | Minimal 3 template generate PDF | ≥ 3 |
| G-07 | Template management | Upload & kelola template .docx | Fully functional |
| G-08 | End-to-end testing | Semua flow utama ter-test | 100% |
| G-09 | Bug fixing & polish | Zero critical bugs | 0 |

---

## Task Breakdown

### Task 016: Portal Launcher & Dashboard
**Estimasi:** 2 hari
**FR:** FR-101, FR-102, FR-103, FR-104

**Deliverables:**

#### 1. Login Page (`/login`)
```tsx
// apps/portal/app/(auth)/login/page.tsx
export default function LoginPage() {
  const router = useRouter();

  const handleLogin = () => {
    // Redirect ke Keycloak
    const keycloakUrl = `${process.env.NEXT_PUBLIC_KEYCLOAK_URL}/realms/pengadilan-agama/protocol/openid-connect/auth`;
    const params = new URLSearchParams({
      client_id: 'portal-client',
      redirect_uri: `${window.location.origin}/auth/callback`,
      response_type: 'code',
      scope: 'openid profile email',
      state: crypto.randomUUID(),
    });

    window.location.href = `${keycloakUrl}?${params}`;
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-primary-50 to-primary-100">
      <Card className="w-full max-w-md">
        <CardHeader className="space-y-4 text-center">
          <div className="flex justify-center">
            <div className="bg-primary-600 text-white rounded-full p-4">
              <Building2 className="h-12 w-12" />
            </div>
          </div>
          <CardTitle className="text-3xl font-bold">SIKERMA</CardTitle>
          <CardDescription>
            Sistem Informasi Kesekretariatan Mahkamah Agung
          </CardDescription>
        </CardHeader>
        <CardContent>
          <Button
            className="w-full"
            size="lg"
            onClick={handleLogin}
          >
            <LogIn className="mr-2 h-5 w-5" />
            Login dengan SSO Pengadilan Agama
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
```

#### 2. Callback Handler (`/auth/callback`)
```tsx
// apps/portal/app/auth/callback/route.ts
export async function GET(request: Request) {
  const { searchParams } = new URL(request.url);
  const code = searchParams.get('code');
  const state = searchParams.get('state');

  if (!code) {
    return NextResponse.redirect(new URL('/login?error=invalid_code', request.url));
  }

  try {
    // Exchange code untuk token
    const tokenResponse = await fetch(
      `${process.env.KEYCLOAK_URL}/realms/pengadilan-agama/protocol/openid-connect/token`,
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: new URLSearchParams({
          grant_type: 'authorization_code',
          client_id: 'portal-client',
          client_secret: process.env.KEYCLOAK_CLIENT_SECRET_PORTAL!,
          code,
          redirect_uri: `${process.env.NEXT_PUBLIC_APP_URL}/auth/callback`,
        }),
      }
    );

    const tokens = await tokenResponse.json();

    // Get user info
    const userInfoResponse = await fetch(
      `${process.env.KEYCLOAK_URL}/realms/pengadilan-agama/protocol/openid-connect/userinfo`,
      {
        headers: { Authorization: `Bearer ${tokens.access_token}` },
      }
    );

    const userInfo = await userInfoResponse.json();

    // Set session (Better Auth)
    await setSession({
      user: {
        id: userInfo.sub,
        name: userInfo.name,
        email: userInfo.email,
        username: userInfo.preferred_username,
      },
      tokens,
    });

    // Redirect ke dashboard
    return NextResponse.redirect(new URL('/', request.url));
  } catch (error) {
    return NextResponse.redirect(new URL('/login?error=auth_failed', request.url));
  }
}
```

#### 3. Dashboard / Launcher (`/`)
```tsx
// apps/portal/app/(app)/page.tsx
export default function PortalDashboard() {
  const { user } = useSession();
  const { data: userData } = useQuery({
    queryKey: ['auth', 'me'],
    queryFn: () => apiClient.get('/auth/me').then(res => res.data)
  });

  // Apps yang bisa diakses user (berdasarkan roles)
  const accessibleApps = useMemo(() => {
    const apps = [];

    if (userData?.permissions?.some(p => p.startsWith('pegawai'))) {
      apps.push({
        name: "Kepegawaian",
        description: "Kelola data pegawai dan riwayatnya",
        icon: Users,
        color: "bg-blue-500",
        href: "http://localhost:3002",
      });
    }

    if (userData?.permissions?.some(p => p.startsWith('master'))) {
      apps.push({
        name: "Master Data",
        description: "Kelola data referensi",
        icon: Database,
        color: "bg-green-500",
        href: "http://localhost:3001",
      });
    }

    if (userData?.roles?.includes('portal:admin')) {
      apps.push({
        name: "Admin Portal",
        description: "Kelola user, roles, dan audit log",
        icon: Shield,
        color: "bg-purple-500",
        href: "/admin",
      });
    }

    return apps;
  }, [userData]);

  const { data: dashboardStats } = useQuery({
    queryKey: ['dashboard', 'summary'],
    queryFn: () => apiClient.get('/dashboard/summary').then(res => res.data)
  });

  return (
    <div className="p-6 space-y-6">
      {/* Welcome Banner */}
      <Card>
        <CardHeader>
          <div className="flex items-center gap-4">
            <Avatar className="h-16 w-16">
              <AvatarImage src={user?.image} alt={user?.name} />
              <AvatarFallback>{user?.name?.charAt(0)}</AvatarFallback>
            </Avatar>
            <div>
              <p className="text-sm text-muted-foreground">Selamat datang,</p>
              <h1 className="text-2xl font-bold">{user?.name}</h1>
              <p className="text-sm text-muted-foreground">{user?.email}</p>
            </div>
          </div>
        </CardHeader>
      </Card>

      {/* Quick Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle className="text-sm font-medium">Total Pegawai</CardTitle>
            <Users className="h-5 w-5 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{dashboardStats?.total_pegawai}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle className="text-sm font-medium">PNS</CardTitle>
            <BadgeCheck className="h-5 w-5 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{dashboardStats?.pegawai_pns}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle className="text-sm font-medium">PPPK</CardTitle>
            <FileText className="h-5 w-5 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{dashboardStats?.pegawai_pppk}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0">
            <CardTitle className="text-sm font-medium">Aktivitas Terakhir</CardTitle>
            <History className="h-5 w-5 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-3xl font-bold">{dashboardStats?.aktivitas_hari_ini}</div>
          </CardContent>
        </Card>
      </div>

      {/* App Launcher */}
      <Card>
        <CardHeader>
          <CardTitle>Aplikasi</CardTitle>
          <CardDescription>Klik untuk membuka aplikasi</CardDescription>
        </CardHeader>
        <CardContent>
          {accessibleApps.length === 0 ? (
            <div className="text-center py-12">
              <AlertCircle className="h-12 w-12 mx-auto text-muted-foreground mb-4" />
              <p className="text-muted-foreground">Anda belum memiliki akses ke aplikasi apapun</p>
              <p className="text-sm text-muted-foreground mt-2">
                Hubungi administrator untuk mendapatkan akses
              </p>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {accessibleApps.map((app) => (
                <Link key={app.name} href={app.href} target={app.href.startsWith('http') ? '_blank' : '_self'}>
                  <Card className="hover:shadow-lg transition-shadow cursor-pointer">
                    <CardHeader>
                      <div className="flex items-center justify-between">
                        <div className={cn("p-3 rounded-lg", app.color)}>
                          <app.icon className="h-6 w-6 text-white" />
                        </div>
                        <ChevronRight className="h-5 w-5 text-muted-foreground" />
                      </div>
                      <CardTitle className="mt-4">{app.name}</CardTitle>
                      <CardDescription>{app.description}</CardDescription>
                    </CardHeader>
                  </Card>
                </Link>
              ))}
            </div>
          )}
        </CardContent>
      </Card>

      {/* Recent Activity */}
      <Card>
        <CardHeader>
          <CardTitle>Aktivitas Terakhir</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            {dashboardStats?.recent_activity?.map((activity: any) => (
              <div key={activity.id} className="flex items-start gap-3">
                <div className="mt-1">
                  <activity.icon className="h-5 w-5 text-muted-foreground" />
                </div>
                <div className="flex-1">
                  <p className="font-medium">{activity.description}</p>
                  <p className="text-sm text-muted-foreground">
                    {activity.user_name} • {formatTanggal(activity.timestamp)}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
```

#### 4. Profil Saya (`/profil`)
```tsx
// apps/portal/app/(app)/profil/page.tsx
export default function ProfilPage() {
  const { user } = useSession();
  const { data: userData, refetch } = useQuery({
    queryKey: ['auth', 'me'],
    queryFn: () => apiClient.get('/auth/me').then(res => res.data)
  });

  const { mutate: updateProfil, isPending } = useMutation({
    mutationFn: (data: any) => apiClient.put('/auth/me', data),
    onSuccess: () => {
      refetch();
      toast.success("Profil berhasil diperbarui");
    }
  });

  const form = useForm({
    defaultValues: {
      name: user?.name || '',
      email: user?.email || '',
      phone: userData?.phone || '',
      avatar: user?.image || '',
    }
  });

  return (
    <div className="p-6 max-w-2xl mx-auto">
      <PageHeader
        title="Profil Saya"
        description="Kelola informasi profil Anda"
      />

      <Card>
        <CardHeader>
          <CardTitle>Informasi Profil</CardTitle>
        </CardHeader>
        <CardContent>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(updateProfil)} className="space-y-6">
              <FormField
                control={form.control}
                name="avatar"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Foto Profil</FormLabel>
                    <FormControl>
                      <FileUpload
                        accept="image/*"
                        maxSize={2 * 1024 * 1024} // 2MB
                        onUpload={(file) => {
                          // Upload to server
                          const formData = new FormData();
                          formData.append('file', file);
                          fetch('/api/v1/upload/avatar', {
                            method: 'POST',
                            body: formData,
                          }).then(res => res.json()).then(data => {
                            field.onChange(data.file_url);
                          });
                        }}
                        preview={field.value}
                      />
                    </FormControl>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Nama Lengkap</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="email"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Email</FormLabel>
                    <FormControl>
                      <Input type="email" {...field} disabled />
                    </FormControl>
                    <FormDescription>
                      Email terhubung dengan akun SSO, tidak dapat diubah
                    </FormDescription>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="phone"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Nomor Telepon</FormLabel>
                    <FormControl>
                      <Input type="tel" {...field} />
                    </FormControl>
                  </FormItem>
                )}
              />

              <Button type="submit" disabled={isPending}>
                Simpan Perubahan
              </Button>
            </form>
          </Form>
        </CardContent>
      </Card>
    </div>
  );
}
```

---

### Task 017: Portal Admin UI
**Estimasi:** 2-3 hari
**FR:** FR-105, FR-106, FR-107, FR-108

**Deliverables:**

#### 1. Admin Layout
```tsx
// apps/portal/app/(app)/admin/layout.tsx
export default function AdminLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const menu = [
    { title: "Dashboard", href: "/admin", icon: LayoutDashboard },
    { title: "User Management", href: "/admin/users", icon: Users },
    { title: "Hak Akses", href: "/admin/hak-akses", icon: Shield },
    { title: "Kelola Roles", href: "/admin/hak-akses/roles", icon: Key },
    { title: "Audit Log", href: "/admin/audit-log", icon: History },
  ];

  return (
    <div className="flex h-screen">
      <Sidebar menuItems={menu} />
      <div className="flex-1 flex flex-col overflow-hidden">
        <AppHeader />
        <main className="flex-1 overflow-y-auto p-6">
          {children}
        </main>
      </div>
    </div>
  );
}
```

#### 2. User Management (`/admin/users`)
```tsx
// apps/portal/app/(app)/admin/users/page.tsx
export default function UserManagementPage() {
  const { data: users, refetch } = useQuery({
    queryKey: ['admin', 'users'],
    queryFn: () => apiClient.get('/admin/users').then(res => res.data)
  });

  const { mutate: createUser } = useMutation({
    mutationFn: (data: any) => apiClient.post('/admin/users', data),
    onSuccess: () => refetch()
  });

  const { mutate: updateUser } = useMutation({
    mutationFn: ({ id, data }: { id: string; data: any }) =>
      apiClient.put(`/admin/users/${id}`, data),
    onSuccess: () => refetch()
  });

  const { mutate: assignClientAccess } = useMutation({
    mutationFn: ({ userId, clientId, grant }: { userId: string; clientId: string; grant: boolean }) =>
      apiClient.put(`/admin/users/${userId}/client-access`, { client_id: clientId, grant }),
    onSuccess: () => refetch()
  });

  return (
    <div className="p-6">
      <PageHeader
        title="Manajemen Pengguna"
        description="Kelola akun pengguna SIKERMA"
        actions={[
          { label: "Tambah User", onClick: () => setShowDialog(true) }
        ]}
      />

      <DataTable
        columns={[
          { accessorKey: 'username', header: 'Username' },
          { accessorKey: 'email', header: 'Email' },
          { accessorKey: 'firstName', header: 'Nama Depan' },
          { accessorKey: 'lastName', header: 'Nama Belakang' },
          {
            accessorKey: 'enabled',
            header: 'Status',
            cell: ({ getValue }) => (
              <StatusBadge
                status={getValue() ? 'active' : 'inactive'}
                variant={getValue() ? 'success' : 'secondary'}
              />
            )
          },
          {
            accessorKey: 'clientRoles',
            header: 'Akses Aplikasi',
            cell: ({ row }) => (
              <div className="flex flex-wrap gap-1">
                {row.clientRoles?.map((role: string) => (
                  <Badge key={role} variant="outline">
                    {role}
                  </Badge>
                ))}
              </div>
            )
          },
        ]}
        data={users?.data || []}
        actions={[
          {
            label: "Edit",
            onClick: (row) => {/* Edit user */}
          },
          {
            label: "Kelola Akses",
            onClick: (row) => setShowAccessDialog(row)
          },
          {
            label: "Reset Password",
            onClick: (row) => {/* Reset password via Keycloak */}
          },
        ]}
      />
    </div>
  );
}
```

#### 3. Hak Akses - Assign Role (`/admin/hak-akses`)
```tsx
// apps/portal/app/(app)/admin/hak-akses/page.tsx
export default function HakAksesPage() {
  const { data: userRoles, refetch } = useQuery({
    queryKey: ['admin', 'user-roles'],
    queryFn: () => apiClient.get('/admin/user-roles').then(res => res.data)
  });

  const { data: allUsers } = useQuery({
    queryKey: ['admin', 'users', 'all'],
    queryFn: () => apiClient.get('/admin/users').then(res => res.data)
  });

  const { data: allRoles } = useQuery({
    queryKey: ['admin', 'roles', 'all'],
    queryFn: () => apiClient.get('/admin/roles').then(res => res.data)
  });

  const { mutate: assignRole } = useMutation({
    mutationFn: (data: { user_id: string; role_id: string }) =>
      apiClient.post('/admin/user-roles', data),
    onSuccess: () => refetch()
  });

  return (
    <div className="p-6">
      <PageHeader
        title="Hak Akses Pengguna"
        description="Assign role ke pengguna per aplikasi"
      />

      <Card>
        <CardHeader>
          <CardTitle>Assign Role Baru</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-2 gap-4">
            <Select onValueChange={(value) => setSelectedUser(value)}>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Pengguna" />
              </SelectTrigger>
              <SelectContent>
                {allUsers?.data?.map((user: any) => (
                  <SelectItem key={user.id} value={user.id}>
                    {user.username} - {user.email}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>

            <Select onValueChange={(value) => setSelectedRole(value)}>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Role" />
              </SelectTrigger>
              <SelectContent>
                {allRoles?.data?.map((role: any) => (
                  <SelectItem key={role.id} value={role.id}>
                    {role.app}: {role.role_name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <Button
            className="mt-4"
            onClick={() => assignRole({ user_id: selectedUser, role_id: selectedRole })}
          >
            Assign Role
          </Button>
        </CardContent>
      </Card>

      {/* List User Roles */}
      <Card className="mt-6">
        <CardHeader>
          <CardTitle>Daftar Hak Akses</CardTitle>
        </CardHeader>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Pengguna</TableHead>
                <TableHead>Role</TableHead>
                <TableHead>Aplikasi</TableHead>
                <TableHead>Aksi</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {userRoles?.data?.map((ur: any) => (
                <TableRow key={ur.id}>
                  <TableCell>{ur.user_name}</TableCell>
                  <TableCell>
                    <Badge variant="outline">{ur.role_name}</Badge>
                  </TableCell>
                  <TableCell>
                    <Badge variant={ur.app === 'portal' ? 'default' : 'secondary'}>
                      {ur.app}
                    </Badge>
                  </TableCell>
                  <TableCell>
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => deleteRole(ur.id)}
                    >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </div>
  );
}
```

#### 4. Kelola Roles (`/admin/hak-akses/roles`)
```tsx
// apps/portal/app/(app)/admin/hak-akses/roles/page.tsx
export default function KelolaRolesPage() {
  const { data: roles, refetch } = useQuery({
    queryKey: ['admin', 'roles'],
    queryFn: () => apiClient.get('/admin/roles').then(res => res.data)
  });

  const { data: permissions } = useQuery({
    queryKey: ['admin', 'permissions'],
    queryFn: () => apiClient.get('/admin/permissions').then(res => res.data)
  });

  const { mutate: createRole } = useMutation({
    mutationFn: (data: any) => apiClient.post('/admin/roles', data),
    onSuccess: () => refetch()
  });

  const { mutate: updateRolePermissions } = useMutation({
    mutationFn: ({ roleId, permissionIds }: { roleId: string; permissionIds: string[] }) =>
      apiClient.post(`/admin/roles/${roleId}/permissions`, { permission_ids: permissionIds }),
    onSuccess: () => refetch()
  });

  return (
    <div className="p-6">
      <PageHeader
        title="Kelola Roles"
        description="Buat dan kelola role untuk setiap aplikasi"
        actions={[
          { label: "Tambah Role", onClick: () => setShowDialog(true) }
        ]}
      />

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {roles?.data?.map((role: any) => (
          <Card key={role.id}>
            <CardHeader>
              <div className="flex items-start justify-between">
                <div>
                  <CardTitle>{role.role_name}</CardTitle>
                  <CardDescription>
                    {role.app.toUpperCase()} • {role.description}
                  </CardDescription>
                </div>
                {!role.is_system && (
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="sm">
                        <MoreVertical className="h-4 w-4" />
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                      <DropdownMenuItem onClick={() => setShowEditDialog(role)}>
                        Edit
                      </DropdownMenuItem>
                      <DropdownMenuItem onClick={() => setShowPermissionDialog(role)}>
                        Kelola Permissions
                      </DropdownMenuItem>
                      <DropdownMenuSeparator />
                      <DropdownMenuItem className="text-red-600" onClick={() => deleteRole(role.id)}>
                        Hapus
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                )}
              </div>
            </CardHeader>
            <CardContent>
              <div className="space-y-2">
                <p className="text-sm font-medium">Permissions:</p>
                <div className="flex flex-wrap gap-1">
                  {role.permissions?.map((perm: any) => (
                    <Badge key={perm.id} variant="outline">
                      {perm.permission_name}
                    </Badge>
                  ))}
                </div>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
```

#### 5. Audit Log Viewer (`/admin/audit-log`)
```tsx
// apps/portal/app/(app)/admin/audit-log/page.tsx
export default function AuditLogPage() {
  const [filters, setFilters] = useState({
    search: '',
    app: '',
    action: '',
    date_from: '',
    date_to: '',
  });

  const { data, isLoading } = useQuery({
    queryKey: ['admin', 'audit-log', filters],
    queryFn: () => apiClient.get('/audit-logs', { params: filters }).then(res => res.data)
  });

  return (
    <div className="p-6">
      <PageHeader
        title="Audit Log"
        description="Pantau semua aktivitas pengguna"
      />

      {/* Filters */}
      <Card className="mb-6">
        <CardContent className="pt-6">
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4">
            <Input
              placeholder="Cari..."
              value={filters.search}
              onChange={(e) => setFilters({ ...filters, search: e.target.value })}
            />
            <Select
              value={filters.app}
              onValueChange={(value) => setFilters({ ...filters, app: value })}
            >
              <SelectTrigger>
                <SelectValue placeholder="Filter App" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="">Semua Aplikasi</SelectItem>
                <SelectItem value="portal">Portal</SelectItem>
                <SelectItem value="master">Master Data</SelectItem>
                <SelectItem value="kepegawaian">Kepegawaian</SelectItem>
              </SelectContent>
            </Select>
            <Select
              value={filters.action}
              onValueChange={(value) => setFilters({ ...filters, action: value })}
            >
              <SelectTrigger>
                <SelectValue placeholder="Filter Action" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="">Semua Action</SelectItem>
                <SelectItem value="CREATE">Create</SelectItem>
                <SelectItem value="UPDATE">Update</SelectItem>
                <SelectItem value="DELETE">Delete</SelectItem>
              </SelectContent>
            </Select>
            <Input
              type="date"
              value={filters.date_from}
              onChange={(e) => setFilters({ ...filters, date_from: e.target.value })}
              placeholder="Dari Tanggal"
            />
            <Input
              type="date"
              value={filters.date_to}
              onChange={(e) => setFilters({ ...filters, date_to: e.target.value })}
              placeholder="Sampai Tanggal"
            />
          </div>
        </CardContent>
      </Card>

      {/* Table */}
      <DataTable
        columns={[
          {
            accessorKey: 'timestamp',
            header: 'Waktu',
            cell: ({ getValue }) => formatTanggal(getValue())
          },
          { accessorKey: 'user_name', header: 'Pengguna' },
          {
            accessorKey: 'app',
            header: 'Aplikasi',
            cell: ({ getValue }) => (
              <Badge variant="outline">{getValue()}</Badge>
            )
          },
          {
            accessorKey: 'action',
            header: 'Action',
            cell: ({ getValue }) => {
              const action = getValue();
              return (
                <Badge variant={
                  action === 'CREATE' ? 'success' :
                  action === 'UPDATE' ? 'warning' :
                  'destructive'
                }>
                  {action}
                </Badge>
              );
            }
          },
          { accessorKey: 'resource', header: 'Resource' },
          {
            accessorKey: 'resource_id',
            header: 'ID',
            cell: ({ getValue }) => (
              <code className="text-xs">{getValue()}</code>
            )
          },
          {
            accessorKey: 'details',
            header: 'Detail',
            cell: ({ getValue }) => (
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setShowDetailDialog(getValue())}
              >
                Lihat
              </Button>
            )
          },
        ]}
        data={data?.data || []}
        pagination={data?.pagination}
      />
    </div>
  );
}
```

---

### Task 018: Cetak SK & Template Management
**Estimasi:** 2 hari
**FR:** FR-311, FR-312

**Deliverables:**

#### 1. Template Management Backend
```go
// internal/handlers/kepegawaian/template_handler.go
type TemplateHandler struct {
  db *gorm.DB
}

func (h *TemplateHandler) List(c *fiber.Ctx) error {
  var templates []TemplateDokumen
  h.db.Where("is_active = ?", true).Find(&templates)
  return c.JSON(templates)
}

func (h *TemplateHandler) Upload(c *fiber.Ctx) error {
  file, err := c.FormFile("file")
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Validate: .docx only
  if !strings.HasSuffix(file.Filename, ".docx") {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format file harus .docx"})
  }

  // Save file
  filename := fmt.Sprintf("template_%d.docx", time.Now().Unix())
  filepath := "uploads/templates/" + filename
  c.SaveFile(file, filepath)

  var placeholders []string
  if placeholdersJSON := c.FormValue("placeholders"); placeholdersJSON != "" {
    json.Unmarshal([]byte(placeholdersJSON), &placeholders)
  }

  template := TemplateDokumen{
    Nama:         c.FormValue("nama"),
    Kode:         c.FormValue("kode"),
    Filepath:     filepath,
    Placeholders: placeholders,
    IsActive:     true,
  }

  h.db.Create(&template)
  return c.Status(fiber.StatusCreated).JSON(template)
}

func (h *TemplateHandler) Update(c *fiber.Ctx) error {
  id := c.Params("id")
  var data UpdateTemplateRequest
  c.BodyParser(&data)

  updates := make(map[string]interface{})
  if data.Nama != "" { updates["nama"] = data.Nama }
  if data.Kode != "" { updates["kode"] = data.Kode }
  if data.Placeholders != nil { updates["placeholders"] = data.Placeholders }
  if data.IsActive != nil { updates["is_active"] = data.IsActive }

  h.db.Model(&TemplateDokumen{}).Where("id = ?", id).Updates(updates)
  return c.JSON(fiber.Map{"message": "updated successfully"})
}
```

#### 2. Cetak SK Backend (Gotenberg Integration)
```go
// internal/handlers/kepegawaian/cetak_handler.go
type CetakHandler struct {
  db         *gorm.DB
  gotenbergURL string
}

func (h *CetakHandler) CetakSK(c *fiber.Ctx) error {
  var req CetakSKRequest
  if err := c.BodyParser(&req); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
  }

  // Get template
  var template TemplateDokumen
  if err := h.db.Where("id = ? AND is_active = ?", req.TemplateID, true).First(&template).Error; err != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Template tidak ditemukan"})
  }

  // Get pegawai data
  var pegawai Pegawai
  h.db.Preload("Golongan").
    Preload("Jabatan").
    Preload("UnitKerja").
    Where("nip = ?", req.NIP).
    First(&pegawai)

  // Read template file
  templateBytes, err := os.ReadFile(template.Filepath)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  // Replace placeholders dengan data pegawai
  docContent := string(templateBytes)
  docContent = strings.ReplaceAll(docContent, "{{NIP}}", pegawai.NIP)
  docContent = strings.ReplaceAll(docContent, "{{NAMA_LENGKAP}}", pegawai.Nama)
  docContent = strings.ReplaceAll(docContent, "{{JABATAN}}", pegawai.Jabatan.Nama)
  docContent = strings.ReplaceAll(docContent, "{{GOLONGAN}}", pegawai.Golongan.Kode)
  docContent = strings.ReplaceAll(docContent, "{{TANGGAL_SEKARANG}}", time.Now().Format("02 Januari 2006"))

  // Convert DOCX to PDF via Gotenberg
  pdfBytes, err := h.convertToPDF(docContent)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  // Save PDF
  filename := fmt.Sprintf("SK_%s_%d.pdf", pegawai.NIP, time.Now().Unix())
  pdfPath := "uploads/cetak/" + filename
  os.WriteFile(pdfPath, pdfBytes, 0644)

  // Insert ke tabel cetak_history
  h.db.Create(&CetakHistory{
    NIP:        pegawai.NIP,
    TemplateID: req.TemplateID,
    FilePath:   pdfPath,
  })

  // Return PDF untuk download
  c.Set("Content-Type", "application/pdf")
  c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
  return c.Send(pdfBytes)
}

func (h *CetakHandler) convertToPDF(docContent string) ([]byte, error) {
  // Create temp file
  tmpFile, _ := os.CreateTemp("", "*.docx")
  defer os.Remove(tmpFile.Name())
  tmpFile.WriteString(docContent)
  tmpFile.Close()

  // Send to Gotenberg
  formData := &bytes.Buffer{}
  writer := multipart.NewWriter(formData)
  fileWriter, _ := writer.CreateFormFile("files", filepath.Base(tmpFile.Name()))
  file, _ := os.Open(tmpFile.Name())
  io.Copy(fileWriter, file)
  writer.Close()

  resp, err := http.Post(
    h.gotenbergURL+"/forms/chromium/convert/doc",
    writer.FormDataContentType(),
    formData,
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  return io.ReadAll(resp.Body)
}
```

#### 3. Template Management Frontend (`/admin/template`)
```tsx
// apps/kepegawaian/app/(app)/admin/template/page.tsx
export default function TemplateManagementPage() {
  const { data: templates, refetch } = useQuery({
    queryKey: ['template'],
    queryFn: () => apiClient.get('/dokumen/templates').then(res => res.data)
  });

  const { mutate: uploadTemplate, isPending } = useMutation({
    mutationFn: (formData: FormData) =>
      apiClient.post('/dokumen/templates', formData, {
        headers: { 'Content-Type': 'multipart/form-data' }
      }),
    onSuccess: () => refetch()
  });

  return (
    <div className="p-6">
      <PageHeader
        title="Manajemen Template"
        description="Upload dan kelola template dokumen .docx"
        actions={[
          { label: "Upload Template", onClick: () => setShowDialog(true) }
        ]}
      />

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {templates?.map((template: any) => (
          <Card key={template.id}>
            <CardHeader>
              <div className="flex items-start justify-between">
                <div>
                  <CardTitle>{template.nama}</CardTitle>
                  <CardDescription>{template.kode}</CardDescription>
                </div>
                <Toggle
                  defaultChecked={template.is_active}
                  onPressedChange={(checked) =>
                    updateTemplate(template.id, { is_active: checked })
                  }
                />
              </div>
            </CardHeader>
            <CardContent>
              <div className="space-y-2">
                <div className="flex items-center gap-2 text-sm">
                  <FileText className="h-4 w-4 text-muted-foreground" />
                  <span>{template.filename}</span>
                </div>
                <div className="text-sm">
                  <p className="text-muted-foreground">Placeholders:</p>
                  <div className="flex flex-wrap gap-1 mt-1">
                    {template.placeholders?.map((ph: string) => (
                      <Badge key={ph} variant="outline">{ph}</Badge>
                    ))}
                  </div>
                </div>
              </div>
            </CardContent>
            <CardFooter>
              <Button variant="outline" size="sm" className="w-full">
                <Download className="mr-2 h-4 w-4" />
                Download
              </Button>
            </CardFooter>
          </Card>
        ))}
      </div>
    </div>
  );
}
```

#### 4. Cetak SK Frontend (`/cetak-sk`)
```tsx
// apps/kepegawaian/app/(app)/cetak-sk/page.tsx
export default function CetakSKPage() {
  const { data: templates } = useQuery({
    queryKey: ['template', 'active'],
    queryFn: () => apiClient.get('/dokumen/templates?active=true').then(res => res.data)
  });

  const { data: pegawai } = useQuery({
    queryKey: ['pegawai', 'dropdown'],
    queryFn: () => apiClient.get('/pegawai/dropdown').then(res => res.data)
  });

  const { mutate: cetakSK, isPending } = useMutation({
    mutationFn: (data: { template_id: string; nip: string }) =>
      apiClient.post('/dokumen/cetak', data, {
        responseType: 'blob' // Download PDF
      }),
    onSuccess: (response) => {
      // Trigger download
      const url = window.URL.createObjectURL(new Blob([response.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', `SK_${Date.now()}.pdf`);
      document.body.appendChild(link);
      link.click();
      link.remove();
    }
  });

  return (
    <div className="p-6 max-w-2xl mx-auto">
      <PageHeader
        title="Cetak SK"
        description="Generate Surat Keputusan dari template"
      />

      <Card>
        <CardHeader>
          <CardTitle>Pilih Template & Pegawai</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <Label>Template SK</Label>
            <Select value={selectedTemplate} onValueChange={setSelectedTemplate}>
              <SelectTrigger>
                <SelectValue placeholder="Pilih template SK" />
              </SelectTrigger>
              <SelectContent>
                {templates?.map((template: any) => (
                  <SelectItem key={template.id} value={template.id}>
                    {template.nama}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div>
            <Label>Pegawai</Label>
            <Select value={selectedPegawai} onValueChange={setSelectedPegawai}>
              <SelectTrigger>
                <SelectValue placeholder="Pilih pegawai" />
              </SelectTrigger>
              <SelectContent>
                {pegawai?.map((p: any) => (
                  <SelectItem key={p.nip} value={p.nip}>
                    {formatNIP(p.nip)} - {p.nama}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <Separator />

          {/* Preview */}
          {selectedTemplate && selectedPegawai && (
            <div className="space-y-2">
              <Label>Preview Data</Label>
              <Card>
                <CardContent className="pt-4">
                  <div className="space-y-2 text-sm">
                    <div className="flex justify-between">
                      <span className="text-muted-foreground">NIP:</span>
                      <span>{formatNIP(selectedPegawai)}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="text-muted-foreground">Nama:</span>
                      <span>{pegawai?.find(p => p.nip === selectedPegawai)?.nama}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="text-muted-foreground">Jabatan:</span>
                      <span>{pegawai?.find(p => p.nip === selectedPegawai)?.jabatan_nama}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="text-muted-foreground">Golongan:</span>
                      <span>{pegawai?.find(p => p.nip === selectedPegawai)?.golongan_kode}</span>
                    </div>
                  </div>
                </CardContent>
              </Card>
            </div>
          )}

          <Button
            className="w-full"
            size="lg"
            onClick={() => cetakSK({ template_id: selectedTemplate, nip: selectedPegawai })}
            disabled={!selectedTemplate || !selectedPegawai || isPending}
          >
            {isPending ? (
              <>
                <Loader2 className="mr-2 h-5 w-5 animate-spin" />
                Generating PDF...
              </>
            ) : (
              <>
                <FileDown className="mr-2 h-5 w-5" />
                Generate & Download PDF
              </>
            )}
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
```

---

## Definition of Done - Sprint 4

Sprint 4 dianggap **DONE** jika:

### Portal
- ✅ Login page redirect ke Keycloak dan kembali dengan session valid
- ✅ Dashboard menampilkan tile app sesuai hak akses user
- ✅ Dashboard widgets (≥4) menampilkan data real
- ✅ Profil Saya dapat edit foto, telepon, dll

### Admin UI
- ✅ User Management dapat list & edit user dari Keycloak
- ✅ Hak Akses dapat assign role ke user
- ✅ Kelola Roles dapat CRUD role & assign permissions
- ✅ Audit Log dapat lihat & filter semua aktivitas

### Cetak SK
- ✅ Template management dapat upload .docx & definisi placeholders
- ✅ Cetak SK dapat generate PDF dari template + data pegawai
- ✅ Minimal 3 template SK berfungsi (SK Pangkat, SK Jabatan, SK CPNS)

### Polish
- ✅ End-to-end testing semua flow utama (login → akses app → CRUD)
- ✅ Zero critical bugs
- ✅ Documentation lengkap (setup guide, user manual)
- ✅ Code review passed

---

## Risks & Mitigations

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Gotenberg template conversion gagal | Medium | High | Test dengan template sederhana dulu, siapkan fallback (HTML → PDF) |
| Keycloak admin API rate limiting | Low | Medium | Implementasi caching & retry logic |
| Complex RBAC UI sulit digunakan | Medium | Medium | User testing dengan admin IT, simplify UI |
| PDF styling tidak sesuai template | Medium | Medium | Siapkan template .docx reference yang sudah tested |

---

## Success Metrics - Fase 1 Overall

| Metric | Target | Actual |
|--------|--------|--------|
| Sprint 1 DONE | 100% | ___ |
| Sprint 2 DONE | 100% | ___ |
| Sprint 3 DONE | 100% | ___ |
| Sprint 4 DONE | 100% | ___ |
| **Overall Fase 1** | **DONE** | ___ |

### Final Checklist Fase 1
- ✅ 29 pegawai PA Penajam ter-migrasi ke sistem
- ✅ Admin bisa login 1x (SSO) dan akses ketiga app
- ✅ Semua data referensi bisa dikelola dari Master Data app
- ✅ Data pegawai bisa dikelola lengkap (biodata + 4 riwayat)
- ✅ Minimal 3 template SK bisa generate PDF
- ✅ RBAC berfungsi (role-based menu & button visibility)
- ✅ Audit trail mencatat semua operasi CRUD
- ✅ Portal menampilkan dashboard ringkasan
