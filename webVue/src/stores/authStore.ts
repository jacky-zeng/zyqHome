import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import type { User } from '@/types'
import { getMeApi, loginApi, registerApi } from '@/api/auth'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || '')
  const savedUser = localStorage.getItem('user')
  const user = ref<User | null>(savedUser ? JSON.parse(savedUser) : null)

  const isLoggedIn = computed(() => !!token.value)

  async function login(username: string, password: string) {
    const res = await loginApi(username, password)
    if (res.data.code === 0) {
      token.value = res.data.data.token
      user.value = res.data.data.user
      localStorage.setItem('token', res.data.data.token)
      localStorage.setItem('user', JSON.stringify(res.data.data.user))
    }
    return res.data
  }

  async function register(data: { email: string; password: string }) {
    const res = await registerApi(data)
    if (res.data.code === 0) {
      token.value = res.data.data.token
      user.value = res.data.data.user
      localStorage.setItem('token', res.data.data.token)
      localStorage.setItem('user', JSON.stringify(res.data.data.user))
    }
    return res.data
  }

  async function fetchUser() {
    if (!token.value) return
    try {
      const res = await getMeApi()
      if (res.data.code === 0) {
        user.value = res.data.data
        localStorage.setItem('user', JSON.stringify(res.data.data))
      }
    } catch {
      logout()
    }
  }

  function logout() {
    token.value = ''
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  return { token, user, isLoggedIn, login, register, fetchUser, logout }
})
