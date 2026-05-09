import request from './request'
import type { ApiResponse, UserBehavior } from '@/types'

export function trackBehaviorApi(data: UserBehavior) {
  return request.post<ApiResponse<null>>('/api/behavior/track', data)
}
