import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { PublicConfig } from '@/types'
import { getPublicConfigApi, getAllConfigsApi, updateConfigsApi, uploadImageApi } from '@/api/config'

export const useConfigStore = defineStore('config', () => {
  const publicConfig = ref<PublicConfig | null>(null)

  async function fetchPublicConfig() {
    const res = await getPublicConfigApi()
    if (res.data.code === 0) {
      publicConfig.value = res.data.data
    }
  }

  async function fetchAllConfigs() {
    const res = await getAllConfigsApi()
    return res.data
  }

  async function updateConfigs(data: Record<string, string>) {
    const res = await updateConfigsApi(data)
    return res.data
  }

  async function uploadImage(file: File) {
    const res = await uploadImageApi(file)
    return res.data
  }

  return { publicConfig, fetchPublicConfig, fetchAllConfigs, updateConfigs, uploadImage }
})
