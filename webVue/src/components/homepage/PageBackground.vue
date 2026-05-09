<script setup lang="ts">
defineProps<{
  imageUrl?: string
  bgColor?: string
}>()
</script>

<template>
  <div
    class="page-background"
    :style="{
      background: imageUrl
        ? `url(${imageUrl}) center/cover no-repeat`
        : bgColor || '#f0f2f5',
    }"
  >
    <div class="overlay" />
    <slot />
  </div>
</template>

<style scoped>
.page-background {
  position: relative;
  width: 100%;
  min-height: 100vh;
}

.overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.08);
  pointer-events: none;
}

.page-background > :not(.overlay) {
  position: relative;
  z-index: 1;
}
</style>
