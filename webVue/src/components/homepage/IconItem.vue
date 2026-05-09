<script setup lang="ts">
import { useTrackBehavior } from '@/composables/useTrackBehavior'

const props = defineProps<{
  title: string
  url: string
  icon: string
  color: string
}>()

const { track } = useTrackBehavior()

function onClick() {
  track({ action: 'icon_click', target: props.url })
  window.open(props.url, '_blank')
}
</script>

<template>
  <div class="icon-item" @click="onClick">
    <div class="icon-bg" :style="{ background: color }">
      <img v-if="icon.startsWith('http') || icon.startsWith('/')" :src="icon" :alt="title" class="icon-img" />
      <span v-else class="icon-text">{{ title.charAt(0) }}</span>
    </div>
    <span class="icon-title">{{ title }}</span>
  </div>
</template>

<style scoped>
.icon-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 12px 8px;
  border-radius: 12px;
  transition: transform 0.2s, background 0.2s;
}

.icon-item:hover {
  transform: translateY(-4px);
  background: rgba(255, 255, 255, 0.5);
}

.icon-bg {
  width: 56px;
  height: 56px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.icon-img {
  width: 32px;
  height: 32px;
  object-fit: contain;
}

.icon-text {
  color: #fff;
  font-size: 22px;
  font-weight: 600;
}

.icon-title {
  font-size: 13px;
  color: #444;
  text-align: center;
  max-width: 72px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
