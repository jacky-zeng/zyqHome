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
    >
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
    </PageBackground>

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
</template>

<style scoped>
.home-page {
  height: 100vh;
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
</style>
