// Uniform API response
export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

// Menu item
export interface MenuItem {
  id: number
  title: string
  url: string
  icon: string
  parent_id?: number
  sort_order: number
  is_active: boolean
  target: string
  children?: MenuItem[]
  created_at: string
  updated_at: string
}

// Center icon
export interface CenterIcon {
  id: number
  title: string
  url: string
  icon: string
  color: string
  sort_order: number
  is_active: boolean
  menu_id: number
  created_at: string
  updated_at: string
}

// Public site configuration
export interface PublicConfig {
  site_title: string
  site_description: string
  background_image: string
  background_color: string
  default_search: string
  search_placeholder: string
  footer_text: string
  icon_columns: number
}

// User
export interface User {
  id: number
  username: string
  nickname: string
  email?: string
  avatar: string
  role?: string
  status?: number
  last_login_ip: string
  last_login_at: string
  created_at: string
  updated_at: string
}

// User behavior
export interface UserBehavior {
  action: 'page_view' | 'icon_click' | 'menu_click' | 'search'
  target?: string
}

// Sort request
export interface SortItem {
  id: number
  sort_order: number
}

export interface SortRequest {
  items: SortItem[]
}

// Image
export interface ImageItem {
  id: number
  filename: string
  url: string
  category: string
  created_at: string
  updated_at: string
}

export interface ImageListResponse {
  items: ImageItem[]
  total: number
  page: number
}

// User custom icon
export interface UserIcon {
  id: number
  user_id: number
  title: string
  image_url: string
  link_url: string
  sort_order: number
  created_at: string
  updated_at: string
}
