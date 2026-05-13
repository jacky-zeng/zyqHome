import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { ImageItem } from '@/types'
import { getUserImageListApi, getUserImageCategoriesApi, uploadUserImageApi, deleteUserImageApi } from '@/api/userImage'

export const useUserImageStore = defineStore('userImage', () => {
  const images = ref<ImageItem[]>([])
  const total = ref(0)
  const page = ref(1)
  const categories = ref<string[]>([])

  async function fetchList(category: string, pageNum: number, pageSize: number = 50) {
    const res = await getUserImageListApi(category, pageNum, pageSize)
    if (res.data.code === 0) {
      images.value = res.data.data.items
      total.value = res.data.data.total
      page.value = res.data.data.page
    }
    return res.data
  }

  async function fetchCategories() {
    const res = await getUserImageCategoriesApi()
    if (res.data.code === 0) {
      categories.value = res.data.data
    }
    return res.data
  }

  async function upload(file: File, category: string) {
    const res = await uploadUserImageApi(file, category)
    return res.data
  }

  async function remove(id: number) {
    const res = await deleteUserImageApi(id)
    return res.data
  }

  // type-compatibility stub (never called in user mode, but needed for ImagePicker)
  async function update(_id: number, _data: { category: string; filename?: string }) {
    return { code: -1, message: '无权限', data: null }
  }

  return { images, total, page, categories, fetchList, fetchCategories, upload, update, remove }
})
