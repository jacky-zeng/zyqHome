import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { CenterIcon } from '@/types'
import type { SortRequest } from '@/types'
import { getActiveIconsApi, getAllIconsApi, createIconApi, updateIconApi, deleteIconApi, sortIconsApi } from '@/api/icon'

export const useIconStore = defineStore('icon', () => {
  const icons = ref<CenterIcon[]>([])
  const allIcons = ref<CenterIcon[]>([])

  async function fetchActiveIcons() {
    const res = await getActiveIconsApi()
    if (res.data.code === 0) {
      icons.value = res.data.data
    }
  }

  async function fetchAllIcons() {
    const res = await getAllIconsApi()
    if (res.data.code === 0) {
      allIcons.value = res.data.data
    }
    return allIcons.value
  }

  async function create(data: Partial<CenterIcon>) {
    const res = await createIconApi(data)
    return res.data
  }

  async function update(id: number, data: Partial<CenterIcon>) {
    const res = await updateIconApi(id, data)
    return res.data
  }

  async function remove(id: number) {
    const res = await deleteIconApi(id)
    return res.data
  }

  async function sort(data: SortRequest) {
    const res = await sortIconsApi(data)
    return res.data
  }

  return { icons, allIcons, fetchActiveIcons, fetchAllIcons, create, update, remove, sort }
})
