<script setup lang="ts">
import type { CenterIcon } from '@/types'
import IconItem from './IconItem.vue'

defineProps<{
  icons: CenterIcon[]
  columns?: number
}>()
</script>

<template>
  <div class="center-icons" v-if="icons.length > 0">
    <div class="icons-grid" :style="{ gridTemplateColumns: `repeat(${columns || 5}, 1fr)` }">
      <IconItem
        v-for="icon in icons"
        :key="icon.id"
        :title="icon.title"
        :url="icon.url"
        :icon="icon.icon"
        :color="icon.color"
      />
    </div>
  </div>
</template>

<style scoped>
.center-icons {
  display: flex;
  justify-content: center;
  margin-top: 50px;
  position: relative;
  z-index: 10;
}

.icons-grid {
  display: grid;
  grid-template-columns: repeat(v-bind('columns || 12'), 1fr);
  gap: 20px;
  max-width: 1200px;
}
</style>
