import request from './request'
import type { ApiResponse, MenuItem, SortRequest } from '@/types'

export function getActiveMenusApi() {
  return request.get<ApiResponse<MenuItem[]>>('/api/menus')
}

export function getAllMenusApi() {
  return request.get<ApiResponse<MenuItem[]>>('/api/admin/menus')
}

export function createMenuApi(data: Partial<MenuItem>) {
  return request.post<ApiResponse<MenuItem>>('/api/admin/menus', data)
}

export function updateMenuApi(id: number, data: Partial<MenuItem>) {
  return request.put<ApiResponse<MenuItem>>(`/api/admin/menus/${id}`, data)
}

export function deleteMenuApi(id: number) {
  return request.delete<ApiResponse<null>>(`/api/admin/menus/${id}`)
}

export function sortMenusApi(data: SortRequest) {
  return request.put<ApiResponse<null>>('/api/admin/menus/sort', data)
}
