<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import { useAuthStore } from '@/stores/authStore'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const menuItems = [
  { path: '/admin', label: '仪表盘', icon: 'DataBoard' },
  { path: '/admin/menus', label: '菜单管理', icon: 'Menu' },
  { path: '/admin/icons', label: '图标管理', icon: 'Grid' },
  { path: '/admin/images', label: '图片管理', icon: 'Picture' },
  { path: '/admin/config', label: '站点配置', icon: 'Setting' },
]

function handleLogout() {
  ElMessageBox.confirm('确定要退出登录吗？', '提示').then(() => {
    authStore.logout()
    router.push('/login')
  }).catch(() => {})
}
</script>

<template>
  <div class="admin-layout">
    <aside class="admin-sidebar">
      <div class="sidebar-header">
        <h3>后台管理</h3>
      </div>
      <el-menu
        :default-active="route.path"
        router
        class="sidebar-menu"
      >
        <el-menu-item
          v-for="item in menuItems"
          :key="item.path"
          :index="item.path"
        >
          <el-icon>
            <component :is="item.icon" />
          </el-icon>
          <span>{{ item.label }}</span>
        </el-menu-item>
      </el-menu>
      <div class="sidebar-footer">
        <el-button text @click="handleLogout">
          <el-icon><SwitchButton /></el-icon>
          退出登录
        </el-button>
      </div>
    </aside>
    <main class="admin-main">
      <router-view />
    </main>
  </div>
</template>

<style scoped>
.admin-layout {
  display: flex;
  min-height: 100vh;
}

.admin-sidebar {
  width: 220px;
  background: #304156;
  color: #fff;
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
}

.sidebar-header {
  padding: 20px;
  text-align: center;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.sidebar-header h3 {
  font-size: 18px;
  font-weight: 600;
}

.sidebar-menu {
  flex: 1;
  border-right: none;
}

.sidebar-footer {
  padding: 16px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  text-align: center;
}

.sidebar-footer .el-button {
  color: rgba(255, 255, 255, 0.65);
}

.admin-main {
  flex: 1;
  padding: 24px;
  background: #f0f2f5;
  overflow-y: auto;
}
</style>
