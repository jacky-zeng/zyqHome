import axios from 'axios'
import type { AxiosInstance, InternalAxiosRequestConfig } from 'axios'
import type { ApiResponse } from '@/types'

const request: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '',
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Request interceptor - inject token
request.interceptors.request.use((config: InternalAxiosRequestConfig) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Response interceptor - handle auth errors
request.interceptors.response.use(
  (response) => {
    const data: ApiResponse = response.data
    if (data.code === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return response
  },
  (error) => {
    return Promise.reject(error)
  },
)

export default request
