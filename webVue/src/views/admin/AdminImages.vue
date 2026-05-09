<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useImageStore } from '@/stores/imageStore'
import type { ImageItem } from '@/types'

const imageStore = useImageStore()
const loading = ref(false)
const currentCategory = ref('')
const pageSize = ref(20)
const uploadDialogVisible = ref(false)
const editDialogVisible = ref(false)
const editingItem = ref<ImageItem | null>(null)
const editCategory = ref('')
const editFilename = ref('')
const newCategory = ref('')
const uploading = ref(false)
const uploadFile = ref<File | null>(null)
const uploadRef = ref()

onMounted(async () => {
  await Promise.all([
    loadImages(),
    imageStore.fetchCategories(),
  ])
})

async function loadImages() {
  loading.value = true
  try {
    await imageStore.fetchList(currentCategory.value, imageStore.page, pageSize.value)
  } finally {
    loading.value = false
  }
}

function onCategoryChange(category: string) {
  currentCategory.value = category
  imageStore.page = 1
  loadImages()
}

function onPageChange(p: number) {
  imageStore.page = p
  loadImages()
}

// Upload
function onFileChange(uploadFileObj: { raw: File }) {
  if (uploadFileObj.raw) {
    uploadFile.value = uploadFileObj.raw
  }
}

function openUploadDialog() {
  uploadFile.value = null
  newCategory.value = ''
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
    const result = await imageStore.upload(uploadFile.value, newCategory.value)
    if (result.code === 0) {
      ElMessage.success('上传成功')
      uploadDialogVisible.value = false
      if (uploadRef.value) uploadRef.value.clearFiles()
      await Promise.all([
        loadImages(),
        imageStore.fetchCategories(),
      ])
    } else {
      ElMessage.error(result.message || '上传失败')
    }
  } catch {
    ElMessage.error('上传失败')
  } finally {
    uploading.value = false
  }
}

// Edit
function openEditDialog(item: ImageItem) {
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
    const result = await imageStore.update(editingItem.value.id, {
      category: editCategory.value,
      filename: editFilename.value.trim(),
    })
    if (result.code === 0) {
      ElMessage.success('更新成功')
      editDialogVisible.value = false
      await Promise.all([
        loadImages(),
        imageStore.fetchCategories(),
      ])
    } else {
      ElMessage.error(result.message || '更新失败')
    }
  } catch {
    ElMessage.error('更新失败')
  }
}

// Delete
async function handleDelete(id: number) {
  try {
    await ElMessageBox.confirm('确定要删除该图片吗？删除后不可恢复。', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    const result = await imageStore.remove(id)
    if (result.code === 0) {
      ElMessage.success('删除成功')
      await Promise.all([
        loadImages(),
        imageStore.fetchCategories(),
      ])
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

function getImageUrl(url: string): string {
  if (url.startsWith('http')) return url
  return import.meta.env.VITE_API_BASE_URL ? `${import.meta.env.VITE_API_BASE_URL}${url}` : url
}
</script>

<template>
  <div class="admin-images">
    <div class="page-header">
      <h2>图片管理</h2>
      <el-button type="primary" @click="openUploadDialog">上传图片</el-button>
    </div>

    <!-- Category filter -->
    <div class="filter-bar">
      <el-radio-group :model-value="currentCategory" @change="onCategoryChange">
        <el-radio-button label="" value="">全部</el-radio-button>
        <el-radio-button
          v-for="cat in imageStore.categories"
          :key="cat"
          :label="cat"
          :value="cat"
        />
      </el-radio-group>
    </div>

    <!-- Image grid -->
    <el-card shadow="hover">
      <div v-loading="loading" class="image-grid">
        <div v-if="imageStore.images.length === 0 && !loading" class="empty-state">
          <el-empty description="暂无图片" />
        </div>
        <div
          v-for="item in imageStore.images"
          :key="item.id"
          class="image-card"
        >
          <div class="image-wrapper">
            <el-image
              :src="getImageUrl(item.url)"
              :alt="item.filename"
              fit="cover"
              class="image-preview"
              lazy
            >
              <template #error>
                <div class="image-error">加载失败</div>
              </template>
            </el-image>
          </div>
          <div class="image-info">
            <p class="image-filename" :title="item.filename">{{ item.filename }}</p>
            <el-tag size="small" type="info">{{ item.category }}</el-tag>
          </div>
          <div class="image-actions">
            <el-button type="primary" link size="small" @click.stop="copyUrl(item.url)">复制URL</el-button>
            <el-button type="primary" link size="small" @click.stop="openEditDialog(item)">编辑</el-button>
            <el-button type="danger" link size="small" @click.stop="handleDelete(item.id)">删除</el-button>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div v-if="imageStore.total > 0" class="pagination-wrapper">
        <el-pagination
          v-model:current-page="imageStore.page"
          :page-size="pageSize"
          :total="imageStore.total"
          layout="total, prev, pager, next"
          @current-change="onPageChange"
        />
      </div>
    </el-card>

    <!-- Upload Dialog -->
    <el-dialog
      v-model="uploadDialogVisible"
      title="上传图片"
      width="500px"
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
              <el-button type="primary">选择文件</el-button>
            </template>
            <template #tip>
              <div class="upload-tip">支持 jpg/png/gif/webp/svg/ico，最大 10MB</div>
            </template>
          </el-upload>
        </el-form-item>
        <el-form-item label="分类">
          <el-select
            v-model="newCategory"
            placeholder="选择或输入分类"
            allow-create
            filterable
            clearable
          >
            <el-option
              v-for="cat in imageStore.categories"
              :key="cat"
              :label="cat"
              :value="cat"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="uploadDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="uploading" @click="handleUpload">上传</el-button>
      </template>
    </el-dialog>

    <!-- Edit Dialog -->
    <el-dialog
      v-model="editDialogVisible"
      title="编辑图片"
      width="400px"
    >
      <el-form label-width="80px">
        <el-form-item label="文件名" required>
          <el-input v-model="editFilename" placeholder="文件名" />
        </el-form-item>
        <el-form-item label="分类">
          <el-select
            v-model="editCategory"
            placeholder="选择或输入分类"
            allow-create
            filterable
            clearable
          >
            <el-option
              v-for="cat in imageStore.categories"
              :key="cat"
              :label="cat"
              :value="cat"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="editDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleEdit">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.page-header h2 {
  font-size: 24px;
  color: #333;
}

.filter-bar {
  margin-bottom: 16px;
}

.image-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
  min-height: 200px;
}

.empty-state {
  grid-column: 1 / -1;
  display: flex;
  justify-content: center;
  padding: 40px 0;
}

.image-card {
  border: 1px solid #eee;
  border-radius: 8px;
  overflow: hidden;
  transition: box-shadow 0.3s;
  background: #fff;
}

.image-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.image-wrapper {
  width: 100%;
  height: 140px;
  overflow: hidden;
  background: #f5f5f5;
}

.image-preview {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.image-error {
  width: 100%;
  height: 140px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #999;
  font-size: 13px;
}

.image-info {
  padding: 8px 10px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}

.image-filename {
  font-size: 12px;
  color: #666;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
  margin: 0;
}

.image-actions {
  padding: 4px 10px 8px;
  display: flex;
  gap: 4px;
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  padding-top: 16px;
  border-top: 1px solid #eee;
}

.upload-tip {
  font-size: 12px;
  color: #999;
  margin-top: 4px;
}
</style>
