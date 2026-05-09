import request from './request'
import type { ApiResponse, ImageItem, ImageListResponse } from '@/types'

export function getImageListApi(category: string, page: number, pageSize: number) {
  const params: Record<string, string | number> = { page, page_size: pageSize }
  if (category) params.category = category
  return request.get<ApiResponse<ImageListResponse>>('/api/admin/images', { params })
}

export function getImageCategoriesApi() {
  return request.get<ApiResponse<string[]>>('/api/admin/images/categories')
}

export function uploadImageApi(file: File, category: string) {
  const formData = new FormData()
  formData.append('file', file)
  if (category) formData.append('category', category)
  return request.post<ApiResponse<ImageItem>>('/api/admin/images', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
}

export function updateImageApi(id: number, data: { category: string; filename?: string }) {
  return request.put<ApiResponse<ImageItem>>(`/api/admin/images/${id}`, data)
}

export function deleteImageApi(id: number) {
  return request.delete<ApiResponse<null>>(`/api/admin/images/${id}`)
}
