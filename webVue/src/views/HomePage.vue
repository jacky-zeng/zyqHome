<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useConfigStore } from '@/stores/configStore'
import { useMenuStore } from '@/stores/menuStore'
import { useIconStore } from '@/stores/iconStore'
import { useTrackBehavior } from '@/composables/useTrackBehavior'
import { useClock } from '@/composables/useClock'
import { useLunarCalendar } from '@/composables/useLunarCalendar'
import PageBackground from '@/components/homepage/PageBackground.vue'
import SearchBar from '@/components/homepage/SearchBar.vue'
import CenterIcons from '@/components/homepage/CenterIcons.vue'
import RightMenu from '@/components/homepage/RightMenu.vue'

const router = useRouter()
const configStore = useConfigStore()
const menuStore = useMenuStore()
const iconStore = useIconStore()
const { track } = useTrackBehavior()
const { now } = useClock()
const { getLunarDateString } = useLunarCalendar()

const { publicConfig } = storeToRefs(configStore)
const { menus } = storeToRefs(menuStore)
const { icons } = storeToRefs(iconStore)

const activeMenuId = ref<number | undefined>(undefined)

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
  // 点击同一个菜单取消选中，回到默认视图
  if (activeMenuId.value === menuId) {
    activeMenuId.value = undefined
    await iconStore.fetchActiveIcons()
  } else {
    activeMenuId.value = menuId
    await iconStore.fetchActiveIcons(menuId)
  }
}

onMounted(async () => {
  await Promise.all([
    configStore.fetchPublicConfig(),
    menuStore.fetchActiveMenus(),
    iconStore.fetchActiveIcons(),
  ])
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
          :icons="icons"
          :columns="publicConfig?.icon_columns || 12"
        />
      </div>
    <!-- 侧边菜单 — 放在 PageBackground 外面，避免 scoped CSS 覆盖 position: fixed -->
    <RightMenu
      :menus="menus"
      :active-menu-id="activeMenuId"
      @select="onMenuSelect"
    />

    <!-- 登录按钮 -->
    <div class="login-btn" @click="router.push('/login')" title="管理后台">
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
        <circle cx="12" cy="7" r="4"></circle>
      </svg>
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
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: transform 0.3s ease;
  z-index: 100;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.login-btn:hover {
  transform: scale(1.1);
}

.login-btn svg {
  width: 22px;
  height: 22px;
  color: #fff;
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
