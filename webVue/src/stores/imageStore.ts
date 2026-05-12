import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { ImageItem } from '@/types'
import { getImageListApi, getImageCategoriesApi, uploadImageApi, updateImageApi, cropImageApi, deleteImageApi } from '@/api/image'

export const useImageStore = defineStore('image', () => {
  const images = ref<ImageItem[]>([])
  const total = ref(0)
  const page = ref(1)
  const categories = ref<string[]>([])

  async function fetchList(category: string, pageNum: number, pageSize: number = 20) {
    const res = await getImageListApi(category, pageNum, pageSize)
    if (res.data.code === 0) {
      images.value = res.data.data.items
      total.value = res.data.data.total
      page.value = res.data.data.page
    }
    return res.data
  }

  async function fetchCategories() {
    const res = await getImageCategoriesApi()
    if (res.data.code === 0) {
      categories.value = res.data.data
    }
    return res.data
  }

  async function upload(file: File, category: string) {
    const res = await uploadImageApi(file, category)
    return res.data
  }

  async function update(id: number, data: { category: string; filename?: string }) {
    const res = await updateImageApi(id, data)
    return res.data
  }

  async function crop(id: number, file: File, data: { category: string; filename?: string }) {
    const res = await cropImageApi(id, file, data)
    return res.data
  }

  async function remove(id: number) {
    const res = await deleteImageApi(id)
    return res.data
  }

  return { images, total, page, categories, fetchList, fetchCategories, upload, update, crop, remove }
})
