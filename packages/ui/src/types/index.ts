export interface User {
  id: string;
  username: string;
  email: string;
  name: string;
  roles: string[];
}

export interface DropdownItem<T = any> {
  value: T;
  label: string;
}

export interface TableColumn<T = any> {
  id: string;
  header: string;
  accessor: keyof T | ((data: T) => React.ReactNode);
  sortable?: boolean;
  filterable?: boolean;
  width?: string | number;
}

export interface PaginationMeta {
  page: number;
  limit: number;
  total: number;
  totalPages: number;
}

export interface FilterOption {
  id: string;
  label: string;
  value: any;
}

export interface NavigationItem {
  label: string;
  href: string;
  icon?: React.ReactNode;
  children?: NavigationItem[];
  roles?: string[]; // Hanya tampil untuk role ini
}

export type StatusType = "success" | "warning" | "error" | "info" | "neutral";

export interface ApiResponse<T = any> {
  success: boolean;
  message?: string;
  data?: T;
  error?: string;
  request_id?: string;
}

export interface FormDataField {
  name: string;
  label: string;
  type: "text" | "email" | "password" | "number" | "select" | "date" | "textarea" | "file";
  placeholder?: string;
  required?: boolean;
  options?: DropdownItem[];
  validation?: any;
}

export interface AuditLog {
  id: string;
  user_id?: string;
  username?: string;
  action: string;
  resource: string;
  resource_id?: string;
  ip_address?: string;
  user_agent?: string;
  changes?: Record<string, any>;
  status: string;
  error_message?: string;
  created_at: string;
}