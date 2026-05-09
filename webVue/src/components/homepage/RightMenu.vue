<script setup lang="ts">
import { ref } from 'vue'
import type { MenuItem } from '@/types'
import { useTrackBehavior } from '@/composables/useTrackBehavior'

const props = defineProps<{
  menus: MenuItem[]
}>()

const { track } = useTrackBehavior()
const isCollapsed = ref(false)
const expandedKeys = ref<number[]>([])

function toggleCollapse() {
  isCollapsed.value = !isCollapsed.value
}

function toggleExpand(id: number) {
  const idx = expandedKeys.value.indexOf(id)
  if (idx > -1) {
    expandedKeys.value.splice(idx, 1)
  } else {
    expandedKeys.value.push(id)
  }
}

function onMenuClick(item: MenuItem) {
  if (item.children && item.children.length > 0) {
    toggleExpand(item.id)
    return
  }
  if (item.url) {
    track({ action: 'menu_click', target: item.url })
    window.open(item.url, item.target || '_blank')
  }
}
</script>

<template>
  <div class="right-menu" :class="{ collapsed: isCollapsed }">
    <button class="toggle-btn" @click="toggleCollapse" :title="isCollapsed ? '展开菜单' : '收起菜单'">
      <svg v-if="isCollapsed" viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
        <path d="M7 10l5 5 5-5z"/>
      </svg>
      <svg v-else viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
        <path d="M17 10l-5 5-5-5z"/>
      </svg>
    </button>

    <div class="menu-content" v-show="!isCollapsed">
      <div
        v-for="item in menus"
        :key="item.id"
        class="menu-item"
        :class="{ 'has-children': item.children && item.children.length > 0, expanded: expandedKeys.includes(item.id) }"
        @click="onMenuClick(item)"
      >
        <span class="menu-title">{{ item.title }}</span>
        <span v-if="item.children && item.children.length > 0" class="menu-arrow">
          <svg viewBox="0 0 24 24" width="14" height="14" fill="currentColor">
            <path d="M10 17l5-5-5-5v10z"/>
          </svg>
        </span>
      </div>

      <!-- children -->
      <template v-for="item in menus" :key="'child-'+item.id">
        <div v-if="item.children && expandedKeys.includes(item.id)" class="sub-menu">
          <div
            v-for="child in item.children"
            :key="child.id"
            class="menu-item sub-item"
            @click="onMenuClick(child)"
          >
            <span class="menu-title">{{ child.title }}</span>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<style scoped>
.right-menu {
  position: fixed;
  right: 0;
  top: 0;
  bottom: 0;
  width: 180px;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(12px);
  box-shadow: -2px 0 12px rgba(0, 0, 0, 0.08);
  z-index: 100;
  padding: 60px 12px 20px;
  overflow-y: auto;
  transition: width 0.3s;
}

.right-menu.collapsed {
  width: 44px;
  padding: 60px 4px 20px;
}

.toggle-btn {
  position: absolute;
  top: 16px;
  right: 12px;
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 8px;
  background: rgba(64, 158, 255, 0.1);
  color: #409eff;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background 0.2s;
}

.toggle-btn:hover {
  background: rgba(64, 158, 255, 0.2);
}

.menu-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.2s;
  margin-bottom: 2px;
  user-select: none;
}

.menu-item:hover {
  background: rgba(64, 158, 255, 0.1);
}

.menu-item.has-children.expanded {
  background: rgba(64, 158, 255, 0.08);
}

.menu-title {
  font-size: 14px;
  color: #333;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.menu-arrow {
  margin-left: 4px;
  display: flex;
  transition: transform 0.2s;
}

.menu-item.expanded .menu-arrow {
  transform: rotate(90deg);
}

.sub-menu {
  padding-left: 16px;
  border-left: 2px solid rgba(64, 158, 255, 0.2);
  margin-left: 4px;
  margin-bottom: 4px;
}

.sub-item {
  padding: 8px 12px;
  font-size: 13px;
}
</style>
