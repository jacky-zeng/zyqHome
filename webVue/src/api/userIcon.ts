import request from './request'
import type { ApiResponse, UserIcon } from '@/types'

export function getUserIconsApi() {
  return request.get<ApiResponse<UserIcon[]>>('/api/user/icons')
}

export function createUserIconApi(data: { image_url: string; link_url: string }) {
  return request.post<ApiResponse<UserIcon>>('/api/user/icons', data)
}

export function deleteUserIconApi(id: number) {
  return request.delete<ApiResponse<null>>(`/api/user/icons/${id}`)
}

export function adminGetUserIconsApi(userId: number) {
  return request.get<ApiResponse<UserIcon[]>>(`/api/admin/users/${userId}/icons`)
}
