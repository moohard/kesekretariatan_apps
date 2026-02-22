"use client"

import * as React from "react"
import { cva, type VariantProps } from "class-variance-authority"
import {
  Upload,
  X,
  File,
  FileText,
  FileImage,
  FileVideo,
  FileAudio,
  FileArchive,
  AlertCircle,
  CheckCircle2,
  Loader2,
} from "lucide-react"
import { cn } from "../../lib/utils"
import { Button } from "../ui/button"

// ============================================
// Types
// ============================================
export interface FileUploadProps
  extends Omit<React.HTMLAttributes<HTMLDivElement>, "onUpload">,
    VariantProps<typeof uploadVariants> {
  onUpload: (file: File) => void | Promise<void>
  onFilesSelected?: (files: File[]) => void
  accept?: string[]
  maxSize?: number // in bytes
  maxFiles?: number
  multiple?: boolean
  preview?: boolean
  disabled?: boolean
  dragLabel?: string
  browseLabel?: string
  showFileSize?: boolean
  value?: File | File[] | null
  onRemove?: (index?: number) => void
  uploadProgress?: number
  uploadError?: string
  uploadSuccess?: boolean
}

export interface UploadedFile {
  id: string
  file: File
  preview?: string
  progress?: number
  error?: string
  uploaded?: boolean
}

// ============================================
// Variants
// ============================================
const uploadVariants = cva(
  "relative flex flex-col items-center justify-center rounded-lg border-2 border-dashed transition-colors cursor-pointer",
  {
    variants: {
      variant: {
        default: "border-gray-300 bg-gray-50 hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-900/50 dark:hover:bg-gray-900",
        error: "border-red-300 bg-red-50 dark:border-red-800 dark:bg-red-900/20",
        success: "border-green-300 bg-green-50 dark:border-green-800 dark:bg-green-900/20",
      },
      size: {
        sm: "p-4 min-h-[100px]",
        md: "p-8 min-h-[150px]",
        lg: "p-12 min-h-[200px]",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "md",
    },
  }
)

// ============================================
// Helper Functions
// ============================================
function formatFileSize(bytes: number): string {
  if (bytes === 0) return "0 Bytes"
  const k = 1024
  const sizes = ["Bytes", "KB", "MB", "GB"]
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i]
}

function getFileIcon(file: File): React.ElementType {
  const type = file.type.split("/")[0]
  const subtype = file.type.split("/")[1]

  switch (type) {
    case "image":
      return FileImage
    case "video":
      return FileVideo
    case "audio":
      return FileAudio
    case "application":
      if (subtype?.includes("zip") || subtype?.includes("rar") || subtype?.includes("7z")) {
        return FileArchive
      }
      return FileText
    default:
      return File
  }
}

function validateFile(file: File, accept?: string[], maxSize?: number): string | null {
  if (maxSize && file.size > maxSize) {
    return `Ukuran file melebihi ${formatFileSize(maxSize)}`
  }

  if (accept && accept.length > 0) {
    const fileType = file.type
    const fileExtension = "." + file.name.split(".").pop()?.toLowerCase()

    const isAccepted = accept.some((pattern) => {
      if (pattern.startsWith(".")) {
        return fileExtension === pattern.toLowerCase()
      }
      if (pattern.endsWith("/*")) {
        return fileType.startsWith(pattern.replace("/*", ""))
      }
      return fileType === pattern
    })

    if (!isAccepted) {
      return `Tipe file tidak didukung. Format yang didukung: ${accept.join(", ")}`
    }
  }

  return null
}

// ============================================
// Components
// ============================================

/**
 * FileUpload - Komponen upload file dengan drag & drop dan preview
 *
 * Fitur:
 * - Drag and drop support
 * - File type dan size validation
 * - Preview untuk gambar
 * - Progress indicator
 * - Multiple file support
 * - Error handling
 * - Dark mode support
 */
export function FileUpload({
  onUpload,
  onFilesSelected,
  accept = [],
  maxSize = 10 * 1024 * 1024, // 10MB default
  maxFiles = 1,
  multiple = false,
  preview = true,
  disabled = false,
  dragLabel = "Seret file ke sini atau",
  browseLabel = "pilih dari komputer",
  showFileSize = true,
  value,
  onRemove,
  uploadProgress,
  uploadError,
  uploadSuccess,
  variant,
  size = "md",
  className,
  ...props
}: FileUploadProps) {
  const [isDragging, setIsDragging] = React.useState(false)
  const [files, setFiles] = React.useState<UploadedFile[]>([])
  const [errors, setErrors] = React.useState<string[]>([])
  const inputRef = React.useRef<HTMLInputElement>(null)

  const resolvedVariant = uploadError ? "error" : uploadSuccess ? "success" : variant

  const handleFiles = React.useCallback(
    (fileList: FileList) => {
      const newFiles: UploadedFile[] = []
      const newErrors: string[] = []

      const filesArray = Array.from(fileList)

      // Check max files
      if (!multiple && filesArray.length > 1) {
        newErrors.push("Hanya dapat mengupload 1 file")
        setErrors(newErrors)
        return
      }

      if (maxFiles && files.length + filesArray.length > maxFiles) {
        newErrors.push(`Maksimal ${maxFiles} file`)
        setErrors(newErrors)
        return
      }

      filesArray.forEach((file) => {
        const error = validateFile(file, accept, maxSize)
        if (error) {
          newErrors.push(`${file.name}: ${error}`)
        } else {
          const uploadedFile: UploadedFile = {
            id: Math.random().toString(36).substr(2, 9),
            file,
            uploaded: false,
          }

          // Create preview for images
          if (preview && file.type.startsWith("image/")) {
            const reader = new FileReader()
            reader.onload = (e) => {
              uploadedFile.preview = e.target?.result as string
            }
            reader.readAsDataURL(file)
          }

          newFiles.push(uploadedFile)
        }
      })

      setErrors(newErrors)
      if (newFiles.length > 0) {
        const updatedFiles = [...files, ...newFiles]
        setFiles(updatedFiles)
        onFilesSelected?.(updatedFiles.map((f) => f.file))

        // Auto upload jika single file
        if (!multiple && newFiles.length === 1) {
          onUpload(newFiles[0].file)
        }
      }
    },
    [files, accept, maxSize, maxFiles, multiple, preview, onUpload, onFilesSelected]
  )

  const handleDragEnter = (e: React.DragEvent) => {
    e.preventDefault()
    e.stopPropagation()
    if (!disabled) setIsDragging(true)
  }

  const handleDragLeave = (e: React.DragEvent) => {
    e.preventDefault()
    e.stopPropagation()
    setIsDragging(false)
  }

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault()
    e.stopPropagation()
  }

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault()
    e.stopPropagation()
    setIsDragging(false)

    if (disabled) return
    if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
      handleFiles(e.dataTransfer.files)
    }
  }

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      handleFiles(e.target.files)
    }
  }

  const handleRemove = (index: number) => {
    const newFiles = files.filter((_, i) => i !== index)
    setFiles(newFiles)
    onRemove?.(index)
    onFilesSelected?.(newFiles.map((f) => f.file))
  }

  const handleClick = () => {
    if (!disabled) {
      inputRef.current?.click()
    }
  }

  return (
    <div className={cn("w-full", className)} {...props}>
      {/* Drop Zone */}
      <div
        className={cn(
          uploadVariants({ variant: resolvedVariant, size }),
          isDragging && "border-primary bg-primary/5",
          disabled && "opacity-50 cursor-not-allowed"
        )}
        onDragEnter={handleDragEnter}
        onDragLeave={handleDragLeave}
        onDragOver={handleDragOver}
        onDrop={handleDrop}
        onClick={handleClick}
      >
        <input
          ref={inputRef}
          type="file"
          className="hidden"
          accept={accept.join(",")}
          multiple={multiple}
          onChange={handleInputChange}
          disabled={disabled}
        />

        <div className="flex flex-col items-center gap-2 text-center">
          <div className="p-3 bg-gray-100 dark:bg-gray-800 rounded-full">
            <Upload className="h-6 w-6 text-gray-500 dark:text-gray-400" />
          </div>

          <div className="text-sm text-gray-600 dark:text-gray-400">
            {dragLabel}{" "}
            <span className="text-primary font-medium cursor-pointer hover:underline">
              {browseLabel}
            </span>
          </div>

          {showFileSize && maxSize && (
            <div className="text-xs text-gray-500 dark:text-gray-500">
              Maksimal {formatFileSize(maxSize)}
            </div>
          )}

          {accept.length > 0 && (
            <div className="text-xs text-gray-500 dark:text-gray-500">
              Format: {accept.join(", ")}
            </div>
          )}
        </div>
      </div>

      {/* Errors */}
      {errors.length > 0 && (
        <div className="mt-2 space-y-1">
          {errors.map((error, index) => (
            <div
              key={index}
              className="flex items-center gap-2 text-sm text-red-600 dark:text-red-400"
            >
              <AlertCircle className="h-4 w-4" />
              {error}
            </div>
          ))}
        </div>
      )}

      {/* Upload Error */}
      {uploadError && (
        <div className="mt-2 flex items-center gap-2 text-sm text-red-600 dark:text-red-400">
          <AlertCircle className="h-4 w-4" />
          {uploadError}
        </div>
      )}

      {/* File Preview List */}
      {files.length > 0 && (
        <div className="mt-4 space-y-2">
          {files.map((file, index) => {
            const Icon = getFileIcon(file.file)

            return (
              <div
                key={file.id}
                className="flex items-center gap-3 p-3 bg-gray-50 dark:bg-gray-800 rounded-lg"
              >
                {/* Preview or Icon */}
                <div className="shrink-0">
                  {file.preview ? (
                    <img
                      src={file.preview}
                      alt={file.file.name}
                      className="h-10 w-10 rounded object-cover"
                    />
                  ) : (
                    <div className="h-10 w-10 rounded bg-gray-200 dark:bg-gray-700 flex items-center justify-center">
                      <Icon className="h-5 w-5 text-gray-500 dark:text-gray-400" />
                    </div>
                  )}
                </div>

                {/* File Info */}
                <div className="flex-1 min-w-0">
                  <div className="text-sm font-medium text-gray-900 dark:text-gray-100 truncate">
                    {file.file.name}
                  </div>
                  <div className="text-xs text-gray-500">
                    {formatFileSize(file.file.size)}
                  </div>
                </div>

                {/* Progress or Status */}
                {uploadProgress !== undefined && uploadProgress > 0 && (
                  <div className="flex items-center gap-2">
                    {uploadProgress < 100 ? (
                      <>
                        <Loader2 className="h-4 w-4 animate-spin text-primary" />
                        <span className="text-xs text-gray-500">{uploadProgress}%</span>
                      </>
                    ) : (
                      <CheckCircle2 className="h-4 w-4 text-green-500" />
                    )}
                  </div>
                )}

                {/* Remove Button */}
                {onRemove && (
                  <button
                    onClick={(e) => {
                      e.stopPropagation()
                      handleRemove(index)
                    }}
                    className="p-1 hover:bg-gray-200 dark:hover:bg-gray-700 rounded"
                  >
                    <X className="h-4 w-4 text-gray-500" />
                  </button>
                )}
              </div>
            )
          })}
        </div>
      )}
    </div>
  )
}

export default FileUpload
