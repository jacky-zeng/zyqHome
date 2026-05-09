<script setup lang="ts">
import { ref } from 'vue'
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

const searchEngines: Record<string, string> = {
  google: 'https://www.google.com/search?q=',
  baidu: 'https://www.baidu.com/s?wd=',
  bing: 'https://www.bing.com/search?q=',
  github: 'https://github.com/search?q=',
}

function doSearch() {
  if (!keyword.value.trim()) return
  const engine = searchEngines[props.defaultSearch] || searchEngines.google
  track({ action: 'search', target: keyword.value })
  window.open(engine + encodeURIComponent(keyword.value), '_blank')
}

function onKeydown(e: KeyboardEvent) {
  if (e.key === 'Enter') doSearch()
}
</script>

<template>
  <div class="search-bar">
    <div class="search-container">
      <input
        v-model="keyword"
        :placeholder="placeholder"
        class="search-input"
        @keydown="onKeydown"
      />
      <button class="search-btn" @click="doSearch">
        <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor">
          <path d="M15.5 14h-.79l-.28-.27A6.471 6.471 0 0 0 16 9.5 6.5 6.5 0 1 0 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/>
        </svg>
      </button>
    </div>
  </div>
</template>

<style scoped>
.search-bar {
  display: flex;
  justify-content: center;
  padding: 20px 0;
}

.search-container {
  position: relative;
  width: 100%;
  max-width: 500px;
}

.search-input {
  width: 100%;
  height: 48px;
  padding: 0 50px 0 20px;
  border: none;
  border-radius: 24px;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  font-size: 16px;
  color: #333;
  outline: none;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  transition: box-shadow 0.3s;
}

.search-input:focus {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
}

.search-btn {
  position: absolute;
  right: 6px;
  top: 50%;
  transform: translateY(-50%);
  width: 36px;
  height: 36px;
  border: none;
  border-radius: 50%;
  background: #409eff;
  color: #fff;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background 0.3s;
}

.search-btn:hover {
  background: #66b1ff;
}
</style>
