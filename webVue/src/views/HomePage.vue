<script setup lang="ts">
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useConfigStore } from '@/stores/configStore'
import { useMenuStore } from '@/stores/menuStore'
import { useIconStore } from '@/stores/iconStore'
import { useTrackBehavior } from '@/composables/useTrackBehavior'
import PageBackground from '@/components/homepage/PageBackground.vue'
import SearchBar from '@/components/homepage/SearchBar.vue'
import CenterIcons from '@/components/homepage/CenterIcons.vue'
import RightMenu from '@/components/homepage/RightMenu.vue'

const configStore = useConfigStore()
const menuStore = useMenuStore()
const iconStore = useIconStore()
const { track } = useTrackBehavior()

const { publicConfig } = storeToRefs(configStore)
const { menus } = storeToRefs(menuStore)
const { icons } = storeToRefs(iconStore)

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
  <PageBackground
    :image-url="publicConfig?.background_image"
    :bg-color="publicConfig?.background_color"
  >
    <div class="home-container">
      <!-- Title -->
      <div class="site-title">
        <h1>{{ publicConfig?.site_title || 'ZyqHome' }}</h1>
        <p class="site-desc">{{ publicConfig?.site_description || '我的导航' }}</p>
      </div>

      <!-- Search -->
      <SearchBar
        :placeholder="publicConfig?.search_placeholder"
        :default-search="publicConfig?.default_search"
      />

      <!-- Icons -->
      <CenterIcons
        v-if="publicConfig?.show_center_icons"
        :icons="icons"
        :columns="publicConfig?.icon_columns || 5"
      />
    </div>

    <!-- Right Menu -->
    <RightMenu
      v-if="publicConfig?.show_right_menu"
      :menus="menus"
    />
  </PageBackground>
</template>

<style scoped>
.home-container {
  position: relative;
  z-index: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 40px 200px 40px 40px;
}

.site-title {
  text-align: center;
  margin-bottom: 12px;
}

.site-title h1 {
  font-size: 42px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.95);
  text-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
}

.site-desc {
  margin-top: 8px;
  font-size: 16px;
  color: rgba(255, 255, 255, 0.8);
  text-shadow: 0 1px 6px rgba(0, 0, 0, 0.15);
}
</style>
