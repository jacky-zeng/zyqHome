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
    <div class="icon-bg" :style="{ background: icon.startsWith('http') || icon.startsWith('/') ? 'transparent' : color }">
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
  transition: transform 0.3s ease;
}

.icon-item:hover {
  transform: translateY(-5px);
}

.icon-bg {
  width: 64px;
  height: 64px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
  position: relative;
  overflow: hidden;
}

.icon-img {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  min-width: 100%;
  min-height: 100%;
  width: 165%;
  max-width: none;
  object-fit: cover;
  border-radius: 16px;
}

.icon-text {
  color: #fff;
  font-size: 28px;
  font-weight: 600;
}

.icon-title {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.8);
  text-align: center;
  max-width: 70px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
