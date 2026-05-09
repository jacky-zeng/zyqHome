<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useTrackBehavior } from '@/composables/useTrackBehavior'

const props = withDefaults(defineProps<{
  placeholder?: string
  defaultSearch?: string
}>(), {
  placeholder: '搜索...',
  defaultSearch: 'google',
})

const { track } = useTrackBehavior()
const keyword = ref('')
const dropdownOpen = ref(false)
const triggerRef = ref<HTMLElement | null>(null)
const dropdownRef = ref<HTMLElement | null>(null)

interface SearchEngine {
  key: string
  name: string
  url: string
  icon: string
  color: string
}

const searchEngines: SearchEngine[] = [
  { key: 'google', name: 'Google', url: 'https://www.google.com/search?q=', icon: 'G', color: '#4285f4' },
  { key: 'baidu', name: '百度', url: 'https://www.baidu.com/s?wd=', icon: '百', color: '#2932e1' },
  { key: 'bing', name: 'Bing', url: 'https://www.bing.com/search?q=', icon: 'B', color: '#008373' },
  { key: 'github', name: 'GitHub', url: 'https://github.com/search?q=', icon: 'H', color: '#333' },
  { key: 'zhihu', name: '知乎', url: 'https://www.zhihu.com/search?type=content&q=', icon: '知', color: '#0066ff' },
  { key: 'bilibili', name: 'B站', url: 'https://search.bilibili.com/all?keyword=', icon: 'B', color: '#fb7299' },
]

const STORAGE_KEY = 'search_engine'

function getSavedEngine(): string {
  try {
    const saved = localStorage.getItem(STORAGE_KEY)
    if (saved && searchEngines.some(e => e.key === saved)) return saved
  } catch { /* localStorage blocked */ }
  return ''
}

const currentEngineKey = ref(getSavedEngine() || props.defaultSearch || 'google')

const currentEngine = computed(() =>
  searchEngines.find(e => e.key === currentEngineKey.value) || searchEngines[0]
)

function selectEngine(key: string) {
  currentEngineKey.value = key
  dropdownOpen.value = false
  try {
    localStorage.setItem(STORAGE_KEY, key)
  } catch { /* localStorage blocked */ }
}

function toggleDropdown() {
  dropdownOpen.value = !dropdownOpen.value
}

function doSearch() {
  if (!keyword.value.trim()) return
  track({ action: 'search', target: keyword.value })
  window.open(currentEngine.value.url + encodeURIComponent(keyword.value), '_blank')
}

function onKeydown(e: KeyboardEvent) {
  if (e.key === 'Enter') doSearch()
}

function handleClickOutside(e: MouseEvent) {
  if (
    dropdownRef.value &&
    !dropdownRef.value.contains(e.target as Node) &&
    triggerRef.value &&
    !triggerRef.value.contains(e.target as Node)
  ) {
    dropdownOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
  <div class="search-bar-wrapper">
    <div class="search-bar">
      <!-- 搜索引擎选择器 -->
      <div class="engine-selector">
        <button
          ref="triggerRef"
          class="engine-trigger"
          @click.stop="toggleDropdown"
          title="切换搜索引擎"
        >
          <span
            class="engine-icon"
            :style="{ background: currentEngine.color }"
          >{{ currentEngine.icon }}</span>
          <svg class="chevron" :class="{ open: dropdownOpen }" width="10" height="10" viewBox="0 0 24 24" fill="currentColor">
            <path d="M7 10l5 5 5-5z"/>
          </svg>
        </button>

        <!-- 下拉菜单 -->
        <div v-if="dropdownOpen" ref="dropdownRef" class="engine-dropdown">
          <div
            v-for="engine in searchEngines"
            :key="engine.key"
            class="engine-item"
            :class="{ active: engine.key === currentEngineKey }"
            @click.stop="selectEngine(engine.key)"
          >
            <span
              class="engine-item-icon"
              :style="{ background: engine.color }"
            >{{ engine.icon }}</span>
            <span class="engine-item-name">{{ engine.name }}</span>
            <svg
              v-if="engine.key === currentEngineKey"
              class="check-icon"
              width="14"
              height="14"
              viewBox="0 0 24 24"
              fill="currentColor"
            >
              <path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/>
            </svg>
          </div>
        </div>
      </div>

      <input
        v-model="keyword"
        :placeholder="placeholder"
        class="search-input"
        @keydown="onKeydown"
      />
      <svg class="search-btn" width="18" height="18" viewBox="0 0 24 24" fill="currentColor" @click="doSearch">
        <path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/>
      </svg>
    </div>
  </div>
</template>

<style scoped>
.search-bar-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 30px;
  position: relative;
  z-index: 20;
}

.search-bar {
  width: 550px;
  height: 44px;
  background: rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 22px;
  display: flex;
  align-items: center;
  padding: 0 20px;
  transition: all 0.3s ease;
  position: relative;
}

.search-bar:focus-within {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.3);
}

/* 搜索引擎选择器 */
.engine-selector {
  position: relative;
  flex-shrink: 0;
  margin-right: 12px;
}

.engine-trigger {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 4px 6px;
  border-radius: 8px;
  border: none;
  background: rgba(255, 255, 255, 0.08);
  cursor: pointer;
  transition: background 0.2s;
}

.engine-trigger:hover {
  background: rgba(255, 255, 255, 0.15);
}

.engine-icon {
  width: 22px;
  height: 22px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 11px;
  font-weight: 700;
  color: #fff;
  flex-shrink: 0;
}

.chevron {
  color: rgba(255, 255, 255, 0.5);
  transition: transform 0.2s;
  flex-shrink: 0;
}

.chevron.open {
  transform: rotate(180deg);
}

/* 下拉菜单 */
.engine-dropdown {
  position: absolute;
  top: calc(100% + 8px);
  left: 0;
  min-width: 170px;
  background: rgba(30, 30, 40, 0.95);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.12);
  border-radius: 12px;
  padding: 6px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
  z-index: 100;
}

.engine-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 10px;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.15s;
  color: #fff;
}

.engine-item:hover {
  background: rgba(255, 255, 255, 0.1);
}

.engine-item.active {
  background: rgba(255, 255, 255, 0.08);
}

.engine-item-icon {
  width: 24px;
  height: 24px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 700;
  color: #fff;
  flex-shrink: 0;
}

.engine-item-name {
  flex: 1;
  font-size: 13px;
}

.check-icon {
  color: rgba(255, 255, 255, 0.5);
  flex-shrink: 0;
}

.search-input {
  flex: 1;
  background: transparent;
  border: none;
  outline: none;
  color: #fff;
  font-size: 14px;
}

.search-input::placeholder {
  color: rgba(255, 255, 255, 0.5);
}

.search-btn {
  color: rgba(255, 255, 255, 0.5);
  cursor: pointer;
  flex-shrink: 0;
  transition: color 0.2s;
}

.search-btn:hover {
  color: rgba(255, 255, 255, 1);
}
</style>
