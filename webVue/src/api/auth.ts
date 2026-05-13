import request from './request'
import type { ApiResponse, User } from '@/types'

interface LoginResult {
  token: string
  user: User
}

export function loginApi(username: string, password: string) {
  return request.post<ApiResponse<LoginResult>>('/api/auth/login', { username, password })
}

export function registerApi(data: { email: string; password: string }) {
  return request.post<ApiResponse<LoginResult>>('/api/auth/register', data)
}

export function getMeApi() {
  return request.get<ApiResponse<User>>('/api/admin/auth/me')
}
