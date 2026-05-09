import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { MenuItem } from '@/types'
import { getActiveMenusApi, getAllMenusApi, createMenuApi, updateMenuApi, deleteMenuApi, sortMenusApi } from '@/api/menu'
import type { SortRequest } from '@/types'

export const useMenuStore = defineStore('menu', () => {
  const menus = ref<MenuItem[]>([])
  const allMenus = ref<MenuItem[]>([])

  async function fetchActiveMenus() {
    const res = await getActiveMenusApi()
    if (res.data.code === 0) {
      menus.value = res.data.data
    }
  }

  async function fetchAllMenus() {
    const res = await getAllMenusApi()
    if (res.data.code === 0) {
      allMenus.value = res.data.data
    }
    return allMenus.value
  }

  async function create(data: Partial<MenuItem>) {
    const res = await createMenuApi(data)
    return res.data
  }

  async function update(id: number, data: Partial<MenuItem>) {
    const res = await updateMenuApi(id, data)
    return res.data
  }

  async function remove(id: number) {
    const res = await deleteMenuApi(id)
    return res.data
  }

  async function sort(data: SortRequest) {
    const res = await sortMenusApi(data)
    return res.data
  }

  return { menus, allMenus, fetchActiveMenus, fetchAllMenus, create, update, remove, sort }
})
