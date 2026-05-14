<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserIconStore } from '@/stores/userIconStore'
import ImagePicker from '@/components/common/ImagePicker.vue'

const userIconStore = useUserIconStore()

const showAddDialog = ref(false)
const showImagePicker = ref(false)
const saving = ref(false)

const form = reactive({
  title: '',
  image_url: '',
  link_url: '',
})

function resetForm() {
  form.title = ''
  form.image_url = ''
  form.link_url = ''
}

function openImagePicker() {
  if (form.image_url) {
    // 重新选择时清空
    form.image_url = ''
  }
  showImagePicker.value = true
}

function onImageSelected(url: string) {
  form.image_url = url
}

async function handleAdd() {
  if (!form.title) {
    ElMessage.warning('请输入图标名称')
    return
  }
  if (!form.image_url) {
    ElMessage.warning('请选择图片')
    return
  }
  if (!form.link_url) {
    ElMessage.warning('请输入跳转URL')
    return
  }

  saving.value = true
  try {
    const result = await userIconStore.create({
      title: form.title,
      image_url: form.image_url,
      link_url: form.link_url,
    })
    if (result.code === 0) {
      ElMessage.success('添加成功')
      showAddDialog.value = false
      resetForm()
    } else {
      ElMessage.error(result.message || '添加失败')
    }
  } catch {
    ElMessage.error('添加失败，请检查网络连接')
  } finally {
    saving.value = false
  }
}

async function handleDelete(id: number) {
  try {
    await ElMessageBox.confirm('确定要删除此图标吗？', '提示')
  } catch {
    return
  }

  try {
    const result = await userIconStore.remove(id)
    if (result.code === 0) {
      ElMessage.success('删除成功')
    } else {
      ElMessage.error(result.message || '删除失败')
    }
  } catch {
    ElMessage.error('删除失败')
  }
}
</script>

<template>
  <div class="user-center">
    <div class="user-center-toolbar">
      <span class="section-title">我的图标</span>
      <el-button type="primary" size="small" @click="showAddDialog = true">
        + 添加图标
      </el-button>
    </div>

    <div v-if="userIconStore.icons.length > 0" class="icons-grid">
      <div
        v-for="icon in userIconStore.icons"
        :key="icon.id"
        class="icon-item"
      >
        <a :href="icon.link_url" target="_blank" rel="noopener noreferrer" class="icon-link">
          <div class="icon-bg">
            <img :src="icon.image_url" :alt="icon.title" class="icon-img" />
          </div>
          <span class="icon-title">{{ icon.title }}</span>
        </a>
        <button class="delete-btn" title="删除" @click="handleDelete(icon.id)">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="18" y1="6" x2="6" y2="18" />
            <line x1="6" y1="6" x2="18" y2="18" />
          </svg>
        </button>
      </div>
    </div>

    <div v-else class="empty-hint">
      暂无图标，点击"添加图标"按钮添加
    </div>

    <!-- 添加图标对话框 -->
    <el-dialog
      v-model="showAddDialog"
      title="添加图标"
      width="480px"
      :close-on-click-modal="false"
      destroy-on-close
      @closed="resetForm"
    >
      <el-form label-width="80px">
        <el-form-item label="名称" required>
          <el-input v-model="form.title" placeholder="请输入图标名称" />
        </el-form-item>
        <el-form-item label="图标" required>
          <div class="image-select-wrapper">
            <div v-if="form.image_url" class="selected-image" @click="openImagePicker">
              <img :src="form.image_url" alt="已选图片" />
              <div class="selected-image-overlay">点击更换</div>
            </div>
            <div v-else class="image-placeholder" @click="openImagePicker">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="3" y="3" width="18" height="18" rx="2" ry="2" />
                <circle cx="8.5" cy="8.5" r="1.5" />
                <polyline points="21 15 16 10 5 21" />
              </svg>
              <span>点击选择图片</span>
            </div>
          </div>
        </el-form-item>
        <el-form-item label="跳转URL" required>
          <el-input v-model="form.link_url" placeholder="输入点击后跳转的链接" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showAddDialog = false">取消</el-button>
        <el-button type="primary" :loading="saving" @click="handleAdd">添加</el-button>
      </template>
    </el-dialog>

    <!-- 图片选择器 -->
    <ImagePicker v-if="showImagePicker" v-model="showImagePicker" mode="user" @select="onImageSelected" />
  </div>
</template>

<style scoped>
.user-center {
  padding-top: 40px;
  width: 100%;
  max-width: 900px;
  margin: 0 auto;
}

.user-center-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.section-title {
  font-size: 18px;
  font-weight: 500;
  color: #fff;
  text-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
}

/* ===== 图标卡片样式 ===== */
.icons-grid {
  display: grid;
  grid-template-columns: repeat(10, 1fr);
  gap: 16px;
}

.icon-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  position: relative;
}

.icon-link {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  text-decoration: none;
  cursor: pointer;
  transition: transform 0.3s ease;
}

.icon-link:hover {
  transform: translateY(-5px);
}

.icon-bg {
  width: 64px;
  height: 64px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
}

.icon-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 12px;
}

.icon-title {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.8);
  text-align: center;
  max-width: 70px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  line-height: 1.3;
}

/* ===== 删除按钮 ===== */
.delete-btn {
  position: absolute;
  top: -4px;
  right: -4px;
  width: 22px;
  height: 22px;
  border-radius: 50%;
  border: none;
  background: rgba(0, 0, 0, 0.5);
  color: #fff;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.2s, background 0.2s;
  padding: 4px;
  z-index: 10;
}

.icon-item:hover .delete-btn {
  opacity: 1;
}

.delete-btn:hover {
  background: rgba(255, 80, 80, 0.8);
}

.delete-btn svg {
  width: 12px;
  height: 12px;
}

/* ===== 空状态 ===== */
.empty-hint {
  text-align: center;
  padding: 60px 20px;
  color: rgba(255, 255, 255, 0.4);
  font-size: 14px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  border: 1px dashed rgba(255, 255, 255, 0.15);
}

/* ===== 添加弹窗：图片选择 ===== */
.image-select-wrapper {
  display: flex;
  align-items: flex-start;
}

.selected-image {
  position: relative;
  width: 100px;
  height: 100px;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  border: 2px solid #409eff;
  transition: border-color 0.2s;
}

.selected-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.selected-image-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.6);
  color: #fff;
  font-size: 11px;
  text-align: center;
  padding: 4px 0;
  opacity: 0;
  transition: opacity 0.2s;
}

.selected-image:hover .selected-image-overlay {
  opacity: 1;
}

.image-placeholder {
  width: 100px;
  height: 100px;
  border-radius: 12px;
  border: 2px dashed #dcdfe6;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  cursor: pointer;
  transition: border-color 0.2s, background 0.2s;
  color: #999;
  background: #fafafa;
}

.image-placeholder:hover {
  border-color: #409eff;
  color: #409eff;
  background: #ecf5ff;
}

.image-placeholder svg {
  width: 28px;
  height: 28px;
}

.image-placeholder span {
  font-size: 12px;
}
</style>
