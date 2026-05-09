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
        : bgColor || 'linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%)',
    }"
  >
    <!-- 装饰渐变圆（仅无背景图时显示） -->
    <template v-if="!imageUrl">
      <div class="bg-blob bg-red" />
      <div class="bg-blob bg-orange" />
      <div class="bg-blob bg-cyan" />
      <div class="bg-blob bg-blue" />
    </template>

    <div class="overlay" />
    <slot />
  </div>
</template>

<style scoped>
.page-background {
  position: relative;
  width: 100%;
  min-height: 100vh;
  overflow: hidden;
}

.overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.08);
  pointer-events: none;
}

.page-background > :not(.overlay):not(.bg-blob) {
  position: relative;
  z-index: 1;
}

/* 装饰渐变圆 */
.bg-blob {
  position: absolute;
  border-radius: 50%;
  opacity: 0.4;
  filter: blur(80px);
  pointer-events: none;
  z-index: 0;
}

.bg-red {
  width: 500px;
  height: 500px;
  background: linear-gradient(135deg, #ff6b6b, #ee5a24);
  top: -100px;
  right: -100px;
}

.bg-orange {
  width: 600px;
  height: 600px;
  background: linear-gradient(135deg, #f39c12, #e74c3c);
  bottom: -200px;
  left: -100px;
}

.bg-cyan {
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, #00cec9, #0984e3);
  top: 20%;
  right: 10%;
}

.bg-blue {
  width: 450px;
  height: 450px;
  background: linear-gradient(135deg, #6c5ce7, #00b894);
  bottom: 0;
  right: -150px;
}
</style>
