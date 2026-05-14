import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { UserIcon } from '@/types'
import { getUserIconsApi, createUserIconApi, deleteUserIconApi } from '@/api/userIcon'

export const useUserIconStore = defineStore('userIcon', () => {
  const icons = ref<UserIcon[]>([])

  async function fetchMyIcons() {
    try {
      const res = await getUserIconsApi()
      if (res.data.code === 0) {
        icons.value = res.data.data || []
      }
    } catch {
      icons.value = []
    }
  }

  async function create(data: { title: string; image_url: string; link_url: string }) {
    const res = await createUserIconApi(data)
    if (res.data.code === 0) {
      icons.value.push(res.data.data)
    }
    return res.data
  }

  async function remove(id: number) {
    const res = await deleteUserIconApi(id)
    if (res.data.code === 0) {
      icons.value = icons.value.filter(i => i.id !== id)
    }
    return res.data
  }

  return { icons, fetchMyIcons, create, remove }
})
