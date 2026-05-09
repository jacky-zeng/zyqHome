import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import type { User } from '@/types'
import { getMeApi, loginApi } from '@/api/auth'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || '')
  const user = ref<User | null>(null)

  const isLoggedIn = computed(() => !!token.value)

  async function login(username: string, password: string) {
    const res = await loginApi(username, password)
    if (res.data.code === 0) {
      token.value = res.data.data.token
      user.value = res.data.data.user
      localStorage.setItem('token', res.data.data.token)
    }
    return res.data
  }

  async function fetchUser() {
    if (!token.value) return
    try {
      const res = await getMeApi()
      if (res.data.code === 0) {
        user.value = res.data.data
      }
    } catch {
      logout()
    }
  }

  function logout() {
    token.value = ''
    user.value = null
    localStorage.removeItem('token')
  }

  return { token, user, isLoggedIn, login, fetchUser, logout }
})
