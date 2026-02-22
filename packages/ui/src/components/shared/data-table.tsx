"use client"

import * as React from "react"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../ui/table"
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "../ui/pagination"
import { Input } from "../ui/input"
import { Button } from "../ui/button"
import { Search, ChevronDown, ChevronUp, MoreVertical } from "lucide-react"
import type { TableColumn, PaginationMeta, DropdownItem } from "../../types"

interface DataTableProps<T> {
  data: T[]
  columns: TableColumn<T>[]
  loading?: boolean
  pagination?: PaginationMeta
  onPageChange?: (page: number) => void
  onRowClick?: (row: T) => void
  onSort?: (columnId: string, direction: "asc" | "desc") => void
  sortColumn?: string
  sortDirection?: "asc" | "desc"
  onSearch?: (search: string) => void
  searchPlaceholder?: string
  searchable?: boolean
  actions?: (row: T) => React.ReactNode
  emptyMessage?: string
  className?: string
}

function DataTable<T>({
  data,
  columns,
  loading = false,
  pagination,
  onPageChange,
  onRowClick,
  onSort,
  sortColumn,
  sortDirection,
  onSearch,
  searchPlaceholder = "Cari...",
  searchable = true,
  actions,
  emptyMessage = "Tidak ada data ditemukan",
  className,
}: DataTableProps<T>) {
  const [search, setSearch] = React.useState("")
  const [localSortColumn, setLocalSortColumn] = React.useState<string | undefined>(sortColumn)
  const [localSortDirection, setLocalSortDirection] = React.useState<"asc" | "desc" | undefined>(sortDirection)

  React.useEffect(() => {
    setLocalSortColumn(sortColumn)
    setLocalSortDirection(sortDirection)
  }, [sortColumn, sortDirection])

  const handleSearch = (value: string) => {
    setSearch(value)
    onSearch?.(value)
  }

  const handleSort = (columnId: string) => {
    let newDirection: "asc" | "desc" = "asc"
    if (localSortColumn === columnId) {
      newDirection = localSortDirection === "asc" ? "desc" : "asc"
    }

    setLocalSortColumn(columnId)
    setLocalSortDirection(newDirection)
    onSort?.(columnId, newDirection)
  }

  const getCellValue = (row: T, column: TableColumn<T>): React.ReactNode => {
    const value = column.accessor
    if (typeof value === "function") {
      return value(row)
    }
    const key = value as string
    return row[key as keyof T] as React.ReactNode
  }

  const renderSortIcon = (columnId: string) => {
    if (localSortColumn !== columnId) return null
    if (localSortDirection === "asc") {
      return <ChevronUp className="ml-2 h-4 w-4" />
    }
    return <ChevronDown className="ml-2 h-4 w-4" />
  }

  return (
    <div className={className}>
      {searchable && onSearch && (
        <div className="mb-4">
          <div className="relative">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              type="text"
              placeholder={searchPlaceholder}
              value={search}
              onChange={(e) => handleSearch(e.target.value)}
              className="pl-10"
            />
          </div>
        </div>
      )}

      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              {columns.map((column) => (
                <TableHead
                  key={column.id}
                  style={{ width: column.width }}
                  className={
                    column.sortable
                      ? "cursor-pointer hover:bg-muted/50"
                      : ""
                  }
                  onClick={() => column.sortable && handleSort(column.id)}
                >
                  <div className="flex items-center">
                    {column.header}
                    {renderSortIcon(column.id)}
                  </div>
                </TableHead>
              ))}
              {actions && <TableHead className="w-[50px]" />}
            </TableRow>
          </TableHeader>
          <TableBody>
            {loading ? (
              <TableRow>
                <TableCell colSpan={columns.length + (actions ? 1 : 0)}>
                  <div className="flex items-center justify-center py-8">
                    <div className="h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent" />
                  </div>
                </TableCell>
              </TableRow>
            ) : data.length === 0 ? (
              <TableRow>
                <TableCell colSpan={columns.length + (actions ? 1 : 0)}>
                  <div className="text-center py-8 text-muted-foreground">
                    {emptyMessage}
                  </div>
                </TableCell>
              </TableRow>
            ) : (
              data.map((row, index) => (
                <TableRow
                  key={index}
                  onClick={() => onRowClick?.(row)}
                  className={onRowClick ? "cursor-pointer" : ""}
                >
                  {columns.map((column) => (
                    <TableCell key={column.id}>
                      {getCellValue(row, column)}
                    </TableCell>
                  ))}
                  {actions && (
                    <TableCell>
                      {actions(row)}
                    </TableCell>
                  )}
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </div>

      {pagination && onPageChange && pagination.totalPages > 1 && (
        <div className="mt-4">
          <Pagination>
            <PaginationContent>
              <PaginationItem>
                <PaginationPrevious
                  onClick={() => onPageChange(pagination.page - 1)}
                  className={
                    pagination.page === 1 ? "pointer-events-none opacity-50" : "cursor-pointer"
                  }
                />
              </PaginationItem>
              {Array.from({ length: Math.min(5, pagination.totalPages) }).map((_, i) => {
                let page = i + 1
                if (pagination.totalPages > 5) {
                  if (pagination.page > 3) {
                    page = pagination.page - 3 + i
                  }
                  if (page > pagination.totalPages) {
                    page = pagination.totalPages
                  }
                }
                return (
                  <PaginationItem key={page}>
                    <PaginationLink
                      onClick={() => onPageChange(page)}
                      isActive={page === pagination.page}
                      className="cursor-pointer"
                    >
                      {page}
                    </PaginationLink>
                  </PaginationItem>
                )
              })}
              <PaginationItem>
                <PaginationNext
                  onClick={() => onPageChange(pagination.page + 1)}
                  className={
                    pagination.page === pagination.totalPages ? "pointer-events-none opacity-50" : "cursor-pointer"
                  }
                />
              </PaginationItem>
            </PaginationContent>
          </Pagination>
        </div>
      )}
    </div>
  )
}

export { DataTable }