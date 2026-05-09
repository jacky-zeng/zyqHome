<script setup lang="ts">
import type { MenuItem } from '@/types'
import { useTrackBehavior } from '@/composables/useTrackBehavior'
import { getIcon } from '@/icons'

const props = defineProps<{
  menus: MenuItem[]
  activeMenuId?: number
}>()

const emit = defineEmits<{
  (e: 'select', menuId: number): void
}>()

const { track } = useTrackBehavior()

function onMenuClick(item: MenuItem) {
  emit('select', item.id)
  if (item.url) {
    track({ action: 'menu_click', target: item.url })
    window.open(item.url, item.target || '_blank')
  }
}
</script>

<template>
  <div v-if="menus.length > 0" class="sidebar">
    <div
      v-for="item in menus"
      :key="item.id"
      class="sidebar-item"
      :class="{ active: activeMenuId === item.id }"
      @click="onMenuClick(item)"
      :title="item.title"
    >
      <div class="sidebar-icon" v-if="item.icon && getIcon(item.icon)">
        <component :is="getIcon(item.icon)" :size="22" />
      </div>
      <span class="sidebar-text">{{ item.title }}</span>
    </div>
  </div>
</template>

<style scoped>
.sidebar {
  position: fixed;
  left: 20px;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  flex-direction: column;
  gap: 8px;
  z-index: 100;
  margin-top: 40px;
}

.sidebar-item {
  width: 56px;
  padding: 8px 0;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
  user-select: none;
}

.sidebar-item:hover {
  transform: translateX(8px);
  background: rgba(255, 255, 255, 0.15);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.sidebar-item.active {
  background: linear-gradient(135deg, #ff6b6b, #ee5a24);
  box-shadow: 0 4px 15px rgba(255, 107, 107, 0.4);
}

.sidebar-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
}

.sidebar-text {
  font-size: 10px;
  line-height: 1.2;
  text-align: center;
  max-width: 52px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
