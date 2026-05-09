<script setup lang="ts">
import { ref, computed } from 'vue'
import { iconEntries } from '@/icons'

const props = withDefaults(defineProps<{
  modelValue: boolean
  selectedIcon?: string
}>(), {
  selectedIcon: '',
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  select: [iconName: string]
}>()

const search = ref('')
const selectedIcon = ref(props.selectedIcon)

const filteredIcons = computed(() => {
  if (!search.value) return iconEntries
  const q = search.value.toLowerCase()
  return iconEntries.filter(entry => entry.name.toLowerCase().includes(q))
})

function selectIcon(name: string) {
  selectedIcon.value = name
}

function confirm() {
  if (selectedIcon.value) {
    emit('select', selectedIcon.value)
  }
  emit('update:modelValue', false)
}

function close() {
  emit('update:modelValue', false)
}

function reset() {
  search.value = ''
  selectedIcon.value = props.selectedIcon
}
</script>

<template>
  <el-dialog
    :model-value="modelValue"
    title="选择图标"
    width="680px"
    top="5vh"
    @open="reset"
    @close="close"
    destroy-on-close
  >
    <el-input
      v-model="search"
      placeholder="搜索图标..."
      clearable
      class="search-input"
      prefix-icon="Search"
    />

    <div class="icon-grid" v-if="filteredIcons.length > 0">
      <div
        v-for="icon in filteredIcons"
        :key="icon.name"
        class="icon-item"
        :class="{ selected: selectedIcon === icon.name }"
        @click="selectIcon(icon.name)"
        :title="icon.name"
      >
        <div class="icon-preview">
          <component :is="icon.component" :size="24" />
        </div>
        <span class="icon-name">{{ icon.name }}</span>
      </div>
    </div>
    <div v-else class="empty-state">
      <el-empty :description="`未找到「${search}」相关图标`" :image-size="60" />
    </div>

    <template #footer>
      <el-button @click="close">取消</el-button>
      <el-button type="primary" :disabled="!selectedIcon" @click="confirm">确定</el-button>
    </template>
  </el-dialog>
</template>

<style scoped>
.search-input {
  margin-bottom: 16px;
}

.icon-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(80px, 1fr));
  gap: 8px;
  max-height: 420px;
  overflow-y: auto;
}

.icon-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 10px 4px 6px;
  border: 2px solid #eee;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.icon-item:hover {
  border-color: #409eff;
  background: #f0f7ff;
}

.icon-item.selected {
  border-color: #409eff;
  background: #e6f1ff;
}

.icon-preview {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  color: #555;
}

.icon-item.selected .icon-preview {
  color: #409eff;
}

.icon-name {
  font-size: 10px;
  color: #999;
  margin-top: 4px;
  text-align: center;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  width: 100%;
}

.icon-item.selected .icon-name {
  color: #409eff;
}

.empty-state {
  padding: 30px 0;
}
</style>
