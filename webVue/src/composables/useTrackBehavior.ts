import { trackBehaviorApi } from '@/api/behavior'
import type { UserBehavior } from '@/types'

export function useTrackBehavior() {
  function track(data: UserBehavior) {
    trackBehaviorApi(data).catch(() => {
      // silently ignore tracking errors
    })
  }

  return { track }
}
