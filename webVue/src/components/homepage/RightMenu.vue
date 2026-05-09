<script setup lang="ts">
import type { MenuItem } from '@/types'
import { useTrackBehavior } from '@/composables/useTrackBehavior'

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

function getIconText(title: string): string {
  return title.charAt(0)
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
      {{ getIconText(item.title) }}
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
  gap: 20px;
  z-index: 100;
  margin-top: 40px;
}

.sidebar-item {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
  font-size: 18px;
  font-weight: 600;
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
</style>
