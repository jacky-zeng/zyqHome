import request from './request'
import type { ApiResponse, PublicConfig } from '@/types'

export function getPublicConfigApi() {
  return request.get<ApiResponse<PublicConfig>>('/api/config')
}

export function getAllConfigsApi() {
  return request.get<ApiResponse<Record<string, string>>>('/api/admin/config')
}

export function updateConfigsApi(data: Record<string, string>) {
  return request.put<ApiResponse<null>>('/api/admin/config', data)
}

export function uploadImageApi(file: File) {
  const formData = new FormData()
  formData.append('file', file)
  return request.post<ApiResponse<{ url: string }>>('/api/admin/upload', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
}
