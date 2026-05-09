import request from './request'
import type { ApiResponse, CenterIcon, SortRequest } from '@/types'

export function getActiveIconsApi(menuId?: number) {
  const params = menuId ? { menu_id: menuId } : {}
  return request.get<ApiResponse<CenterIcon[]>>('/api/icons', { params })
}

export function getAllIconsApi() {
  return request.get<ApiResponse<CenterIcon[]>>('/api/admin/icons')
}

export function createIconApi(data: Partial<CenterIcon>) {
  return request.post<ApiResponse<CenterIcon>>('/api/admin/icons', data)
}

export function updateIconApi(id: number, data: Partial<CenterIcon>) {
  return request.put<ApiResponse<CenterIcon>>(`/api/admin/icons/${id}`, data)
}

export function deleteIconApi(id: number) {
  return request.delete<ApiResponse<null>>(`/api/admin/icons/${id}`)
}

export function sortIconsApi(data: SortRequest) {
  return request.put<ApiResponse<null>>('/api/admin/icons/sort', data)
}
