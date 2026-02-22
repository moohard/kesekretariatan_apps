"use client"

import * as React from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import * as z from "zod"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog"
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "../ui/alert-dialog"
import { Button } from "../ui/button"
import { Input } from "../ui/input"
import { Label } from "../ui/label"
import { Loader2 } from "lucide-react"
import type { FormDataField } from "../../types"

interface FormDialogProps<T extends z.ZodType> {
  open: boolean
  onOpenChange: (open: boolean) => void
  title: string
  description?: string
  fields: FormDataField[]
  schema: T
  defaultValues?: z.infer<T>
  onSubmit: (data: z.infer<T>) => void | Promise<void>
  isLoading?: boolean
  submitLabel?: string
  cancelLabel?: string
}

function FormDialog<T extends z.ZodType>({
  open,
  onOpenChange,
  title,
  description,
  fields,
  schema,
  defaultValues,
  onSubmit,
  isLoading = false,
  submitLabel = "Simpan",
  cancelLabel = "Batal",
}: FormDialogProps<T>) {
  const form = useForm<z.infer<T>>({
    resolver: zodResolver(schema),
    defaultValues,
  })

  const handleSubmit = async (data: z.infer<T>) => {
    await onSubmit(data)
  }

  const renderField = (field: FormDataField) => {
    switch (field.type) {
      case "textarea":
        return (
          <div key={field.name} className="space-y-2">
            <Label htmlFor={field.name}>
              {field.label}
              {field.required && <span className="text-destructive ml-1">*</span>}
            </Label>
            <textarea
              id={field.name}
              placeholder={field.placeholder}
              {...form.register(field.name as any)}
              className="flex min-h-[80px] w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            />
            {form.formState.errors[field.name as keyof typeof form.formState.errors] && (
              <p className="text-sm text-destructive">
                {form.formState.errors[field.name as keyof typeof form.formState.errors]?.message as string}
              </p>
            )}
          </div>
        )

      case "select":
        return (
          <div key={field.name} className="space-y-2">
            <Label htmlFor={field.name}>
              {field.label}
              {field.required && <span className="text-destructive ml-1">*</span>}
            </Label>
            <select
              id={field.name}
              {...form.register(field.name as any)}
              className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            >
              <option value="">Pilih...</option>
              {field.options?.map((option) => (
                <option key={String(option.value)} value={String(option.value)}>
                  {option.label}
                </option>
              ))}
            </select>
            {form.formState.errors[field.name as keyof typeof form.formState.errors] && (
              <p className="text-sm text-destructive">
                {form.formState.errors[field.name as keyof typeof form.formState.errors]?.message as string}
              </p>
            )}
          </div>
        )

      case "date":
        return (
          <div key={field.name} className="space-y-2">
            <Label htmlFor={field.name}>
              {field.label}
              {field.required && <span className="text-destructive ml-1">*</span>}
            </Label>
            <Input
              id={field.name}
              type="date"
              placeholder={field.placeholder}
              {...form.register(field.name as any)}
            />
            {form.formState.errors[field.name as keyof typeof form.formState.errors] && (
              <p className="text-sm text-destructive">
                {form.formState.errors[field.name as keyof typeof form.formState.errors]?.message as string}
              </p>
            )}
          </div>
        )

      case "file":
        return (
          <div key={field.name} className="space-y-2">
            <Label htmlFor={field.name}>
              {field.label}
              {field.required && <span className="text-destructive ml-1">*</span>}
            </Label>
            <Input
              id={field.name}
              type="file"
              placeholder={field.placeholder}
              {...form.register(field.name as any)}
            />
            {form.formState.errors[field.name as keyof typeof form.formState.errors] && (
              <p className="text-sm text-destructive">
                {form.formState.errors[field.name as keyof typeof form.formState.errors]?.message as string}
              </p>
            )}
          </div>
        )

      default:
        return (
          <div key={field.name} className="space-y-2">
            <Label htmlFor={field.name}>
              {field.label}
              {field.required && <span className="text-destructive ml-1">*</span>}
            </Label>
            <Input
              id={field.name}
              type={field.type}
              placeholder={field.placeholder}
              {...form.register(field.name as any)}
            />
            {form.formState.errors[field.name as keyof typeof form.formState.errors] && (
              <p className="text-sm text-destructive">
                {form.formState.errors[field.name as keyof typeof form.formState.errors]?.message as string}
              </p>
            )}
          </div>
        )
    }
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[600px]">
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
          {description && <DialogDescription>{description}</DialogDescription>}
        </DialogHeader>

        <form onSubmit={form.handleSubmit(handleSubmit)} className="space-y-4">
          <div className="grid gap-4 py-4 max-h-[60vh] overflow-y-auto">
            {fields.map(renderField)}
          </div>

          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={() => onOpenChange(false)}
              disabled={isLoading}
            >
              {cancelLabel}
            </Button>
            <Button type="submit" disabled={isLoading}>
              {isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
              {submitLabel}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}

// Delete Confirm Component
interface DeleteConfirmProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  title?: string
  description?: string
  onConfirm: () => void | Promise<void>
  isLoading?: boolean
}

function DeleteConfirm({
  open,
  onOpenChange,
  title = "Hapus Data",
  description = "Apakah Anda yakin ingin menghapus data ini? Tindakan ini tidak dapat dibatalkan.",
  onConfirm,
  isLoading = false,
}: DeleteConfirmProps) {
  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>{title}</AlertDialogTitle>
          <AlertDialogDescription>{description}</AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogAction asChild onClick={(e) => e.preventDefault()}>
            <Button
              variant="outline"
              onClick={() => onOpenChange(false)}
              disabled={isLoading}
            >
              Batal
            </Button>
          </AlertDialogAction>
          <AlertDialogAction asChild>
            <Button
              variant="destructive"
              onClick={onConfirm}
              disabled={isLoading}
            >
              {isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
              Hapus
            </Button>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}

export { FormDialog, DeleteConfirm }