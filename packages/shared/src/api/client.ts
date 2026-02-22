import { API_BASE_URL } from "../constants"

// API Client Options
export interface ApiClientOptions {
  baseURL?: string
  timeout?: number
  headers?: Record<string, string>
}

// Request Options
export interface RequestOptions extends Omit<RequestInit, 'body'> {
  params?: Record<string, unknown>
  signal?: AbortSignal
  timeout?: number
  body?: unknown
}

// API Error
export class ApiError extends Error {
  code: number
  request_id?: string
  details?: any

  constructor(message: string, code: number, request_id?: string, details?: any) {
    super(message)
    this.name = "ApiError"
    this.code = code
    this.request_id = request_id
    this.details = details
  }
}

// API Client Class
export class ApiClient {
  private baseURL: string
  private timeout: number
  private defaultHeaders: Record<string, string>

  constructor(options: ApiClientOptions = {}) {
    this.baseURL = options.baseURL || API_BASE_URL
    this.timeout = options.timeout || 30000
    this.defaultHeaders = {
      "Content-Type": "application/json",
      ...options.headers,
    }
  }

  private buildURL(endpoint: string, params?: Record<string, any>): string {
    const url = new URL(endpoint, this.baseURL)

    if (params) {
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined && value !== null) {
          url.searchParams.append(key, String(value))
        }
      })
    }

    return url.toString()
  }

  private async request<T>(
    endpoint: string,
    options: RequestOptions = {}
  ): Promise<T> {
    const {
      params,
      headers = {},
      body,
      signal,
      timeout = this.timeout,
      ...fetchOptions
    } = options

    const url = this.buildURL(endpoint, params)
    const timeoutSignal = AbortSignal.timeout(timeout)
    const combinedSignal = signal ? AbortSignal.any([signal, timeoutSignal]) : timeoutSignal

    const response = await fetch(url, {
      ...fetchOptions,
      signal: combinedSignal,
      headers: {
        ...this.defaultHeaders,
        ...headers,
      },
      body: body ? JSON.stringify(body) : undefined,
    })

    const data = await response.json() as { message?: string; request_id?: string; details?: unknown }

    if (!response.ok) {
      throw new ApiError(
        data.message || response.statusText,
        response.status,
        data.request_id,
        data.details
      )
    }

    return data as T
  }

  async get<T>(endpoint: string, options?: RequestOptions): Promise<T> {
    return this.request<T>(endpoint, { ...options, method: "GET" })
  }

  async post<T>(endpoint: string, body?: any, options?: RequestOptions): Promise<T> {
    return this.request<T>(endpoint, { ...options, method: "POST", body })
  }

  async put<T>(endpoint: string, body?: any, options?: RequestOptions): Promise<T> {
    return this.request<T>(endpoint, { ...options, method: "PUT", body })
  }

  async patch<T>(endpoint: string, body?: any, options?: RequestOptions): Promise<T> {
    return this.request<T>(endpoint, { ...options, method: "PATCH", body })
  }

  async delete<T>(endpoint: string, options?: RequestOptions): Promise<T> {
    return this.request<T>(endpoint, { ...options, method: "DELETE" })
  }

  async upload<T>(endpoint: string, file: File, options?: RequestOptions): Promise<T> {
    const formData = new FormData()
    formData.append("file", file)

    return this.request<T>(endpoint, {
      ...options,
      method: "POST",
      headers: {
        ...options?.headers,
      }, // Remove Content-Type to let browser set it with boundary
      body: formData as any, // FormData is not RequestInit body type
    })
  }

  // Set authentication token
  setAuthToken(token: string) {
    this.defaultHeaders.Authorization = `Bearer ${token}`
  }

  // Clear authentication token
  clearAuthToken() {
    delete this.defaultHeaders.Authorization
  }
}

// Default API client instance
export const apiClient = new ApiClient()

// Hook for API client with auth token
export function useApiClient() {
  return apiClient
}