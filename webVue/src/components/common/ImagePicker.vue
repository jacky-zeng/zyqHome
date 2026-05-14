<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useImageStore } from '@/stores/imageStore'
import { useUserImageStore } from '@/stores/userImageStore'
import type { ImageItem } from '@/types'

const props = withDefaults(defineProps<{
  modelValue: boolean
  selectedUrl?: string
  mode?: 'admin' | 'user'
}>(), {
  selectedUrl: '',
  mode: 'admin',
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  select: [url: string]
}>()

const adminStore = useImageStore()
const userStore = useUserImageStore()
const store = computed(() => props.mode === 'admin' ? adminStore : userStore)

const loading = ref(false)
const currentCategory = ref('')
const selectedImage = ref('')

// Upload
const uploadDialogVisible = ref(false)
const uploading = ref(false)
const uploadFile = ref<File | null>(null)
const uploadCategory = ref('')
const uploadRef = ref()

// Edit (only available in admin mode)
const editDialogVisible = ref(false)
const editingItem = ref<ImageItem | null>(null)
const editCategory = ref('')
const editFilename = ref('')

onMounted(async () => {
  await Promise.all([
    loadImages(),
    store.value.fetchCategories(),
  ])
})

watch(() => props.modelValue, (val) => {
  if (val) {
    selectedImage.value = props.selectedUrl
    loadImages()
    store.value.fetchCategories()
  }
})

async function loadImages() {
  loading.value = true
  try {
    await store.value.fetchList(currentCategory.value, 1, 50)
  } finally {
    loading.value = false
  }
}

function onCategoryChange(category: string) {
  currentCategory.value = category
  loadImages()
}

function selectImage(url: string) {
  selectedImage.value = url
}

function confirm() {
  if (selectedImage.value) {
    emit('select', selectedImage.value)
  }
  emit('update:modelValue', false)
}

function close() {
  emit('update:modelValue', false)
}

const uploadTipText = computed(() =>
  props.mode === 'admin' ? '支持 jpg/png/gif/webp/svg/ico，最大 10MB' : '支持 jpg/png/gif/webp/svg/ico，最大 500KB'
)

function getImageUrl(url: string): string {
  if (url.startsWith('http')) return url
  return import.meta.env.VITE_API_BASE_URL ? `${import.meta.env.VITE_API_BASE_URL}${url}` : url
}

// Upload
function onFileChange(uploadFileObj: { raw: File }) {
  if (uploadFileObj.raw) {
    uploadFile.value = uploadFileObj.raw
  }
}

function openUploadDialog() {
  uploadFile.value = null
  uploadCategory.value = ''
  if (uploadRef.value) {
    uploadRef.value.clearFiles()
  }
  uploadDialogVisible.value = true
}

async function handleUpload() {
  if (!uploadFile.value) {
    ElMessage.warning('请选择要上传的图片')
    return
  }
  uploading.value = true
  try {
    const result = await store.value.upload(uploadFile.value, uploadCategory.value)
    if (result.code === 0) {
      ElMessage.success('上传成功')
      uploadDialogVisible.value = false
      if (uploadRef.value) uploadRef.value.clearFiles()
      await loadImages()
      await store.value.fetchCategories()
    } else {
      ElMessage.error(result.message || '上传失败')
    }
  } catch {
    ElMessage.error('上传失败')
  } finally {
    uploading.value = false
  }
}

// Edit (admin only)
function openEditDialog(item: ImageItem) {
  if (props.mode !== 'admin') return
  editingItem.value = item
  editCategory.value = item.category
  editFilename.value = item.filename
  editDialogVisible.value = true
}

async function handleEdit() {
  if (!editingItem.value) return
  if (!editFilename.value.trim()) {
    ElMessage.warning('文件名不能为空')
    return
  }
  try {
    const result = await store.value.update(editingItem.value.id, {
      category: editCategory.value,
      filename: editFilename.value.trim(),
    })
    if (result.code === 0) {
      ElMessage.success('更新成功')
      editDialogVisible.value = false
      await loadImages()
      await store.value.fetchCategories()
    } else {
      ElMessage.error(result.message || '更新失败')
    }
  } catch {
    ElMessage.error('更新失败')
  }
}

// Delete
async function handleDelete(item: ImageItem) {
  try {
    await ElMessageBox.confirm(`确定要删除「${item.filename}」吗？删除后不可恢复。`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    })
    const result = await store.value.remove(item.id)
    if (result.code === 0) {
      ElMessage.success('删除成功')
      if (selectedImage.value === item.url) {
        selectedImage.value = ''
      }
      await loadImages()
      await store.value.fetchCategories()
    } else {
      ElMessage.error(result.message || '删除失败')
    }
  } catch {
    // cancelled
  }
}

// Copy URL
function copyUrl(url: string) {
  navigator.clipboard.writeText(url).then(() => {
    ElMessage.success('URL 已复制')
  })
}
</script>

<template>
  <el-dialog
    :model-value="modelValue"
    :title="mode === 'admin' ? '图片管理' : '我的图库'"
    width="1125px"
    top="5vh"
    @close="close"
    destroy-on-close
  >
    <!-- Toolbar -->
    <div class="toolbar">
      <div class="filter-bar">
        <el-radio-group :model-value="currentCategory" @change="onCategoryChange" size="small">
          <el-radio-button label="" value="">全部</el-radio-button>
          <el-radio-button
            v-for="cat in store.categories"
            :key="cat"
            :label="cat"
            :value="cat"
          />
        </el-radio-group>
      </div>
      <el-button type="primary" size="small" @click="openUploadDialog">
        上传图片
      </el-button>
    </div>

    <!-- Image grid -->
    <div v-loading="loading" class="image-grid">
      <div
        v-for="item in store.images"
        :key="item.id"
        class="image-card"
        :class="{ selected: selectedImage === item.url }"
        @click="selectImage(item.url)"
      >
        <div class="image-wrapper">
          <el-image
            :src="getImageUrl(item.url)"
            :alt="item.filename"
            fit="cover"
            lazy
          >
            <template #error>
              <div class="image-error">加载失败</div>
            </template>
          </el-image>
        </div>
        <div class="image-actions-overlay">
          <el-button link size="small" @click.stop="copyUrl(item.url)">复制</el-button>
          <el-button v-if="mode === 'admin'" link size="small" @click.stop="openEditDialog(item)">编辑</el-button>
          <el-button link size="small" type="danger" @click.stop="handleDelete(item)">删除</el-button>
        </div>
        <div class="check-mark" v-if="selectedImage === item.url">
          <el-icon color="#409eff" size="20"><Check /></el-icon>
        </div>
      </div>
      <div v-if="store.images.length === 0 && !loading" class="empty-state">
        <el-empty description="暂无图片" :image-size="80" />
      </div>
    </div>

    <template #footer>
      <el-button @click="close">取消</el-button>
      <el-button type="primary" :disabled="!selectedImage" @click="confirm">确定</el-button>
    </template>
  </el-dialog>

  <!-- Upload Dialog -->
  <el-dialog
    v-model="uploadDialogVisible"
    title="上传图片"
    width="450px"
    append-to-body
    destroy-on-close
  >
    <el-form label-width="80px">
      <el-form-item label="选择图片" required>
        <el-upload
          ref="uploadRef"
          :show-file-list="true"
          :on-change="onFileChange"
          accept="image/jpeg,image/png,image/gif,image/webp,image/svg+xml,image/x-icon"
          :limit="1"
          :auto-upload="false"
        >
          <template #trigger>
            <el-button type="primary" size="small">选择文件</el-button>
          </template>
          <template #tip>
            <div class="upload-tip">{{ uploadTipText }}</div>
          </template>
        </el-upload>
      </el-form-item>
      <el-form-item label="分类">
        <el-select
          v-model="uploadCategory"
          placeholder="选择或输入分类"
          size="small"
          allow-create
          filterable
          clearable
        >
          <el-option
            v-for="cat in store.categories"
            :key="cat"
            :label="cat"
            :value="cat"
          />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button size="small" @click="uploadDialogVisible = false">取消</el-button>
      <el-button size="small" type="primary" :loading="uploading" @click="handleUpload">上传</el-button>
    </template>
  </el-dialog>

  <!-- Edit Dialog (admin only) -->
  <el-dialog
    v-if="mode === 'admin'"
    v-model="editDialogVisible"
    title="编辑图片"
    width="400px"
    append-to-body
    destroy-on-close
  >
    <el-form label-width="80px">
      <el-form-item label="文件名" required>
        <el-input v-model="editFilename" placeholder="文件名" size="small" />
      </el-form-item>
      <el-form-item label="分类">
        <el-select
          v-model="editCategory"
          placeholder="选择或输入分类"
          size="small"
          allow-create
          filterable
          clearable
        >
          <el-option
            v-for="cat in store.categories"
            :key="cat"
            :label="cat"
            :value="cat"
          />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button size="small" @click="editDialogVisible = false">取消</el-button>
      <el-button size="small" type="primary" @click="handleEdit">保存</el-button>
    </template>
  </el-dialog>
</template>

<style scoped>
.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  gap: 12px;
}

.filter-bar {
  flex: 1;
  overflow-x: auto;
}

.image-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
  gap: 12px;
  max-height: 600px;
  overflow-y: auto;
  min-height: 150px;
}

.empty-state {
  grid-column: 1 / -1;
  display: flex;
  justify-content: center;
  padding: 30px 0;
}

.image-card {
  position: relative;
  border: 2px solid #eee;
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
  transition: border-color 0.2s;
  aspect-ratio: 1;
}

.image-card:hover {
  border-color: #409eff;
}

.image-card.selected {
  border-color: #409eff;
}

.image-wrapper {
  width: 100%;
  height: 100%;
  background: #f5f5f5;
}

.image-wrapper .el-image {
  width: 100%;
  height: 100%;
}

.image-error {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #999;
  font-size: 12px;
}

.image-actions-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  gap: 2px;
  padding: 4px;
  background: linear-gradient(transparent, rgba(0, 0, 0, 0.6));
  opacity: 0;
  transition: opacity 0.2s;
}

.image-card:hover .image-actions-overlay {
  opacity: 1;
}

.image-actions-overlay .el-button {
  color: #fff;
  font-size: 11px;
}

.image-actions-overlay .el-button:hover {
  color: #409eff;
}

.image-actions-overlay .el-button--danger:hover {
  color: #f56c6c;
}

.check-mark {
  position: absolute;
  top: 4px;
  right: 4px;
  background: #fff;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.15);
}

.upload-tip {
  font-size: 12px;
  color: #999;
  margin-top: 4px;
}
</style>
