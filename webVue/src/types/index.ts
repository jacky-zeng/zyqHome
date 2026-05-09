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
  show_center_icons: boolean
  show_right_menu: boolean
  icon_columns: number
}

// User
export interface User {
  id: number
  username: string
  nickname: string
  avatar: string
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
