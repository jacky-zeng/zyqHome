<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { useAuthStore } from '@/stores/authStore'

const props = defineProps<{ visible: boolean }>()
const emit = defineEmits<{ 'update:visible': [value: boolean] }>()

const authStore = useAuthStore()

const activeTab = ref<'login' | 'register'>('login')
const loading = ref(false)

const loginForm = reactive({
  username: '',
  password: '',
})

const registerForm = reactive({
  email: '',
  password: '',
  confirmPassword: '',
})

function resetForms() {
  loginForm.username = ''
  loginForm.password = ''
  registerForm.email = ''
  registerForm.password = ''
  registerForm.confirmPassword = ''
}

function close() {
  emit('update:visible', false)
  resetForms()
}

async function handleLogin() {
  if (!loginForm.username || !loginForm.password) {
    ElMessage.warning('请输入邮箱/用户名和密码')
    return
  }

  loading.value = true
  try {
    const result = await authStore.login(loginForm.username, loginForm.password)
    if (result.code === 0) {
      ElMessage.success('登录成功')
      close()
    } else {
      ElMessage.error(result.message || '登录失败')
    }
  } catch {
    ElMessage.error('登录失败，请检查网络连接')
  } finally {
    loading.value = false
  }
}

async function handleRegister() {
  if (!registerForm.email) {
    ElMessage.warning('请输入邮箱')
    return
  }
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(registerForm.email)) {
    ElMessage.warning('请输入有效的邮箱地址')
    return
  }
  if (!registerForm.password) {
    ElMessage.warning('请输入密码')
    return
  }
  if (registerForm.password.length < 6) {
    ElMessage.warning('密码长度不能少于6位')
    return
  }
  if (registerForm.password !== registerForm.confirmPassword) {
    ElMessage.warning('两次输入的密码不一致')
    return
  }

  loading.value = true
  try {
    const result = await authStore.register({
      email: registerForm.email,
      password: registerForm.password,
    })
    if (result.code === 0) {
      ElMessage.success('注册成功，欢迎加入')
      close()
    } else {
      ElMessage.error(result.message || '注册失败')
    }
  } catch {
    ElMessage.error('注册失败，请检查网络连接')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <el-dialog
    :model-value="visible"
    @update:model-value="emit('update:visible', $event)"
    :close-on-click-modal="false"
    width="420px"
    top="15vh"
    destroy-on-close
  >
    <el-tabs v-model="activeTab" stretch>
      <el-tab-pane label="会员登录" name="login">
        <el-form @submit.prevent="handleLogin" style="padding: 10px 0">
          <el-form-item>
            <el-input
              v-model="loginForm.username"
              placeholder="邮箱/用户名"
              size="large"
              :prefix-icon="'User'"
            />
          </el-form-item>
          <el-form-item>
            <el-input
              v-model="loginForm.password"
              type="password"
              placeholder="密码"
              size="large"
              :prefix-icon="'Lock'"
              @keyup.enter="handleLogin"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              size="large"
              :loading="loading"
              style="width: 100%"
              @click="handleLogin"
            >
              登 录
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="会员注册" name="register">
        <el-form @submit.prevent="handleRegister" style="padding: 10px 0">
          <el-form-item>
            <el-input
              v-model="registerForm.email"
              placeholder="邮箱"
              size="large"
              :prefix-icon="'Message'"
            />
          </el-form-item>
          <el-form-item>
            <el-input
              v-model="registerForm.password"
              type="password"
              placeholder="密码（至少6位）"
              size="large"
              :prefix-icon="'Lock'"
            />
          </el-form-item>
          <el-form-item>
            <el-input
              v-model="registerForm.confirmPassword"
              type="password"
              placeholder="确认密码"
              size="large"
              :prefix-icon="'Lock'"
              @keyup.enter="handleRegister"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              size="large"
              :loading="loading"
              style="width: 100%"
              @click="handleRegister"
            >
              注 册
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </el-dialog>
</template>
