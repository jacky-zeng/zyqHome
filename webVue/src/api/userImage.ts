import request from './request'
import type { ApiResponse, ImageItem, ImageListResponse } from '@/types'

export function getUserImageListApi(category: string, page: number, pageSize: number) {
  const params: Record<string, string | number> = { page, page_size: pageSize }
  if (category) params.category = category
  return request.get<ApiResponse<ImageListResponse>>('/api/user/images', { params })
}

export function getUserImageCategoriesApi() {
  return request.get<ApiResponse<string[]>>('/api/user/images/categories')
}

export function uploadUserImageApi(file: File, category: string) {
  const formData = new FormData()
  formData.append('file', file)
  if (category) formData.append('category', category)
  return request.post<ApiResponse<ImageItem>>('/api/user/images', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
}

export function deleteUserImageApi(id: number) {
  return request.delete<ApiResponse<null>>(`/api/user/images/${id}`)
}
