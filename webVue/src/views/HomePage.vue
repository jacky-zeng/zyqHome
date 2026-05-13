<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useConfigStore } from '@/stores/configStore'
import { useMenuStore } from '@/stores/menuStore'
import { useIconStore } from '@/stores/iconStore'
import { useTrackBehavior } from '@/composables/useTrackBehavior'
import { useClock } from '@/composables/useClock'
import { useLunarCalendar } from '@/composables/useLunarCalendar'
import { useAuthStore } from '@/stores/authStore'
import { useUserIconStore } from '@/stores/userIconStore'
import PageBackground from '@/components/homepage/PageBackground.vue'
import SearchBar from '@/components/homepage/SearchBar.vue'
import CenterIcons from '@/components/homepage/CenterIcons.vue'
import RightMenu from '@/components/homepage/RightMenu.vue'
import AuthDialog from '@/components/auth/AuthDialog.vue'
import UserCenterIcons from '@/components/homepage/UserCenterIcons.vue'

const configStore = useConfigStore()
const menuStore = useMenuStore()
const iconStore = useIconStore()
const authStore = useAuthStore()
const { track } = useTrackBehavior()
const { now } = useClock()
const { getLunarDateString } = useLunarCalendar()

const { publicConfig } = storeToRefs(configStore)
const { menus } = storeToRefs(menuStore)
const { icons } = storeToRefs(iconStore)

const activeMenuId = ref<number | undefined>(undefined)
const authDialogVisible = ref(false)
const showUserMenu = ref(false)
const userIconStore = useUserIconStore()
const showUserCenter = ref(false)

// 从 localStorage 恢复上次选中的菜单
const savedMenuId = localStorage.getItem('zyqhome_selected_menu')
if (savedMenuId) {
  const id = parseInt(savedMenuId)
  if (id === -1) {
    showUserCenter.value = true
  } else if (!isNaN(id) && id > 0) {
    activeMenuId.value = id
  }
}

function saveMenuState() {
  if (showUserCenter.value) {
    localStorage.setItem('zyqhome_selected_menu', '-1')
  } else if (activeMenuId.value !== undefined) {
    localStorage.setItem('zyqhome_selected_menu', String(activeMenuId.value))
  } else {
    localStorage.removeItem('zyqhome_selected_menu')
  }
}

// 已登录会员时，在菜单列表头部插入"我的"菜单
const displayMenus = computed(() => {
  const items = [...menus.value]
  if (authStore.isLoggedIn && authStore.user?.role === 'member') {
    items.unshift({
      id: -1,
      title: '我的',
      url: '',
      icon: 'User',
      parent_id: 0,
      sort_order: 0,
      is_active: true,
      target: '',
      created_at: '',
      updated_at: '',
    })
  }
  return items
})

// 登出时关闭用户中心并清除保存的菜单状态
watch(() => authStore.isLoggedIn, (loggedIn) => {
  if (!loggedIn) {
    showUserCenter.value = false
    activeMenuId.value = undefined
    localStorage.removeItem('zyqhome_selected_menu')
  }
})

// 网页缩放
const zoomLevel = ref(100)
const savedZoom = localStorage.getItem('zyqhome_page_zoom')
if (savedZoom) {
  const v = parseInt(savedZoom)
  if (!isNaN(v) && v >= 20 && v <= 180) {
    zoomLevel.value = v
  }
}

function zoomIn() {
  zoomLevel.value = Math.min(180, zoomLevel.value + 10)
  localStorage.setItem('zyqhome_page_zoom', String(zoomLevel.value))
}

function zoomOut() {
  zoomLevel.value = Math.max(20, zoomLevel.value - 10)
  localStorage.setItem('zyqhome_page_zoom', String(zoomLevel.value))
}

function resetZoom() {
  zoomLevel.value = 100
  localStorage.setItem('zyqhome_page_zoom', '100')
}

async function onMenuSelect(menuId: number) {
  if (menuId === -1) {
    // "我的"菜单切换
    showUserCenter.value = !showUserCenter.value
    if (showUserCenter.value) {
      activeMenuId.value = undefined
      try {
        await userIconStore.fetchMyIcons()
      } catch {
        // ignore fetch errors
      }
    }
  } else {
    showUserCenter.value = false
    if (activeMenuId.value === menuId) {
      activeMenuId.value = undefined
      await iconStore.fetchActiveIcons()
    } else {
      activeMenuId.value = menuId
      await iconStore.fetchActiveIcons(menuId)
    }
  }
  saveMenuState()
}

function toggleAuth() {
  if (authStore.isLoggedIn) {
    showUserMenu.value = !showUserMenu.value
  } else {
    authDialogVisible.value = true
  }
}

function logout() {
  showUserMenu.value = false
  authStore.logout()
}

// 点击弹出层外部关闭
function onDocumentClick(e: MouseEvent) {
  if (showUserMenu.value) {
    const target = e.target as HTMLElement
    if (!target.closest('.login-btn') && !target.closest('.user-menu-dropdown')) {
      showUserMenu.value = false
    }
  }
}

onMounted(() => {
  document.addEventListener('click', onDocumentClick)
})

onUnmounted(() => {
  document.removeEventListener('click', onDocumentClick)
})

onMounted(async () => {
  await Promise.all([
    configStore.fetchPublicConfig(),
    menuStore.fetchActiveMenus(),
    iconStore.fetchActiveIcons(),
  ])

  // 恢复上次选中的菜单状态
  if (showUserCenter.value && authStore.isLoggedIn) {
    await userIconStore.fetchMyIcons()
  } else if (activeMenuId.value !== undefined) {
    await iconStore.fetchActiveIcons(activeMenuId.value)
  }

  track({ action: 'page_view' })
})
</script>

<template>
  <div class="home-page">
    <PageBackground
      :image-url="publicConfig?.background_image"
      :bg-color="publicConfig?.background_color"
    />

    <!-- 内容层：缩放除背景外的所有元素 -->
    <div class="content-layer" :style="zoomLevel !== 100 ? { zoom: zoomLevel / 100 } : {}">
      <div class="home-container">
        <!-- 时钟 -->
        <div class="clock-section">
          <div class="clock-time">{{ now.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit', second: '2-digit' }) }}</div>
          <div class="clock-date">{{ getLunarDateString(now) }}</div>
        </div>

        <!-- 搜索栏 -->
        <SearchBar
          :placeholder="publicConfig?.search_placeholder"
          :default-search="publicConfig?.default_search"
        />

        <!-- 图标网格 -->
        <CenterIcons
          v-if="!showUserCenter"
          :icons="icons"
          :columns="publicConfig?.icon_columns || 12"
        />
        <!-- 会员自定义图标 -->
        <UserCenterIcons v-if="showUserCenter" />
      </div>
    <!-- 侧边菜单 — 放在 PageBackground 外面，避免 scoped CSS 覆盖 position: fixed -->
    <RightMenu
      :menus="displayMenus"
      :active-menu-id="showUserCenter ? -1 : activeMenuId"
      @select="onMenuSelect"
    />

    <!-- 登录/会员按钮 -->
    <div
      class="login-btn"
      :class="{ 'is-loggedin': authStore.isLoggedIn }"
      :title="authStore.isLoggedIn ? '已登录' : '会员登录/注册'"
      @click="toggleAuth"
    >
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
        <circle cx="12" cy="7" r="4"></circle>
      </svg>
    </div>
    <!-- 已登录会员弹出层 -->
    <div v-if="showUserMenu" class="user-menu-dropdown">
      <div class="user-menu-header">
        <div class="user-avatar">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
            <circle cx="12" cy="7" r="4"></circle>
          </svg>
        </div>
        <div class="user-info">
          <div class="user-nickname">{{ authStore.user?.nickname || authStore.user?.username || '已登录' }}</div>
          <div class="user-role">会员</div>
        </div>
      </div>
      <div class="user-menu-body">
        <div class="user-menu-item logout-item" @click="logout">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
            <polyline points="16 17 21 12 16 7"></polyline>
            <line x1="21" y1="12" x2="9" y2="12"></line>
          </svg>
          <span>退出登录</span>
        </div>
      </div>
    </div>

    <!-- 底部文字 -->
    <div class="footer-text">
      {{ publicConfig?.footer_text || '只羡忘羡不羡仙，说是天天就天天' }}
    </div>

    </div>

    <!-- 网页缩放控制（不受缩放影响） -->
    <div class="zoom-control">
      <button class="zoom-btn" @click="zoomIn" title="放大">＋</button>
      <div class="zoom-value" @click="resetZoom" title="点击重置为 100%">{{ zoomLevel }}%</div>
      <button class="zoom-btn" @click="zoomOut" title="缩小">－</button>
    </div>

    <!-- 会员注册/登录弹窗 -->
    <AuthDialog v-model:visible="authDialogVisible" />
  </div>
</template>

<style scoped>
.home-page {
  height: 100vh;
  overflow: hidden;
  position: relative;
}

/* 内容缩放层：覆盖在背景之上 */
.content-layer {
  position: absolute;
  inset: 0;
  z-index: 1;
  overflow: hidden;
}

.home-container {
  position: relative;
  z-index: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 80px;
}

/* 时钟 */
.clock-section {
  text-align: center;
  position: relative;
  z-index: 10;
}

.clock-time {
  font-size: 80px;
  font-weight: 300;
  color: #fff;
  letter-spacing: 8px;
  text-shadow: 0 4px 30px rgba(0, 0, 0, 0.3);
  font-variant-numeric: tabular-nums;
}

.clock-date {
  margin-top: 10px;
  font-size: 16px;
  color: rgba(255, 255, 255, 0.7);
  letter-spacing: 2px;
}

/* 登录按钮 */
.login-btn {
  position: fixed;
  top: 20px;
  left: 20px;
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.12);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.3s ease, background 0.3s ease;
  z-index: 100;
  border: 1px solid rgba(255, 255, 255, 0.15);
}

.login-btn.is-loggedin {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
  border: none;
}

.login-btn:hover {
  transform: scale(1.1);
}

.login-btn svg {
  width: 22px;
  height: 22px;
  color: #fff;
}

/* 用户菜单弹出层 */
.user-menu-dropdown {
  position: fixed;
  top: 76px;
  left: 20px;
  width: 200px;
  background: rgba(30, 30, 40, 0.92);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.12);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
  z-index: 200;
  overflow: hidden;
  animation: userMenuFadeIn 0.15s ease;
}

@keyframes userMenuFadeIn {
  from {
    opacity: 0;
    transform: translateY(-8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.user-menu-header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.user-avatar svg {
  width: 20px;
  height: 20px;
  color: #fff;
}

.user-info {
  flex: 1;
  min-width: 0;
}

.user-nickname {
  font-size: 14px;
  font-weight: 500;
  color: #fff;
  line-height: 1.3;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-role {
  font-size: 11px;
  color: rgba(255, 255, 255, 0.4);
  margin-top: 2px;
}

.user-menu-body {
  padding: 6px;
}

.user-menu-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.15s;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.8);
}

.user-menu-item:hover {
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
}

.user-menu-item svg {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

.logout-item:hover {
  background: rgba(255, 80, 80, 0.15);
  color: #ff6b6b;
}

/* 底部文字 */
.footer-text {
  position: fixed;
  bottom: 30px;
  left: 50%;
  transform: translateX(-50%);
  color: rgba(255, 255, 255, 0.3);
  font-size: 12px;
  z-index: 10;
  text-align: center;
}

/* 网页缩放控制 */
.zoom-control {
  position: fixed;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
  z-index: 200;
  background: rgba(255, 255, 255, 0.08);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border-radius: 14px;
  padding: 8px 6px;
  border: 1px solid rgba(255, 255, 255, 0.12);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  user-select: none;
}

.zoom-btn {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  border: none;
  background: transparent;
  color: rgba(255, 255, 255, 0.7);
  font-size: 18px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  line-height: 1;
}

.zoom-btn:hover {
  background: rgba(255, 255, 255, 0.15);
  color: #fff;
}

.zoom-btn:active {
  transform: scale(0.9);
}

.zoom-value {
  color: rgba(255, 255, 255, 0.85);
  font-size: 11px;
  font-weight: 600;
  text-align: center;
  cursor: pointer;
  padding: 4px 0;
  letter-spacing: 0.5px;
  transition: color 0.2s;
  font-variant-numeric: tabular-nums;
}

.zoom-value:hover {
  color: #fff;
}
</style>
