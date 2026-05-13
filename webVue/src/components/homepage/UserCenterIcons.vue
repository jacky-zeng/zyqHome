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
  image_url: '',
  link_url: '',
})

function resetForm() {
  form.image_url = ''
  form.link_url = ''
}

function onImageSelected(url: string) {
  form.image_url = url
}

async function handleAdd() {
  if (!form.image_url) {
    ElMessage.warning('请选择或输入图片URL')
    return
  }
  if (!form.link_url) {
    ElMessage.warning('请输入跳转URL')
    return
  }

  saving.value = true
  try {
    const result = await userIconStore.create({
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
          <el-image :src="icon.image_url" fit="cover" class="icon-img" />
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
      width="450px"
      :close-on-click-modal="false"
      destroy-on-close
      @closed="resetForm"
    >
      <el-form label-width="90px">
        <el-form-item label="图片URL" required>
          <el-input v-model="form.image_url" placeholder="输入图片URL或从图库选择">
            <template #append>
              <el-button @click="showImagePicker = true">选择</el-button>
            </template>
          </el-input>
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

    <!-- 图片选择器（按需挂载，避免未登录时请求管理后台 API） -->
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

.icons-grid {
  display: grid;
  grid-template-columns: repeat(10, 1fr);
  gap: 12px;
}

.icon-item {
  position: relative;
  aspect-ratio: 1;
  border-radius: 12px;
  overflow: hidden;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(255, 255, 255, 0.15);
  transition: transform 0.2s, box-shadow 0.2s;
}

.icon-item:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.icon-link {
  display: block;
  width: 100%;
  height: 100%;
}

.icon-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.delete-btn {
  position: absolute;
  top: 2px;
  right: 2px;
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

.empty-hint {
  text-align: center;
  padding: 60px 20px;
  color: rgba(255, 255, 255, 0.4);
  font-size: 14px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  border: 1px dashed rgba(255, 255, 255, 0.15);
}
</style>
