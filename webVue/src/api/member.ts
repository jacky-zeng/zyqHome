import request from './request'
import type { ApiResponse, User } from '@/types'

export interface MemberListResponse {
  items: User[]
  total: number
  page: number
  page_size: number
}

export function getMemberListApi(params: { page: number; page_size: number }) {
  return request.get<ApiResponse<MemberListResponse>>('/api/admin/members', { params })
}

export function getMemberDetailApi(id: number) {
  return request.get<ApiResponse<User>>(`/api/admin/members/${id}`)
}

export function updateMemberStatusApi(id: number, status: number) {
  return request.put<ApiResponse<null>>(`/api/admin/members/${id}/status`, { status })
}
