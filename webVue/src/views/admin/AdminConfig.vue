<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useConfigStore } from '@/stores/configStore'

const configStore = useConfigStore()
const loading = ref(false)
const uploading = ref(false)

const form = ref({
  site_title: 'ZyqHome',
  site_description: '我的导航',
  background_image: '',
  background_color: '#f0f2f5',
  default_search: 'google',
  search_placeholder: '搜索...',
  show_center_icons: 'true',
  show_right_menu: 'true',
  icon_columns: '5',
})

onMounted(async () => {
  await configStore.fetchPublicConfig()
  if (configStore.publicConfig) {
    Object.assign(form.value, configStore.publicConfig)
  }
})

async function handleSave() {
  loading.value = true
  try {
    const result = await configStore.updateConfigs({ ...form.value })
    if (result.code === 0) {
      ElMessage.success('保存成功')
    } else {
      ElMessage.error(result.message || '保存失败')
    }
  } catch {
    ElMessage.error('保存失败')
  } finally {
    loading.value = false
  }
}

async function handleUpload(file: File) {
  uploading.value = true
  try {
    const result = await configStore.uploadImage(file)
    if (result.code === 0) {
      form.value.background_image = result.data.url
      ElMessage.success('上传成功')
    } else {
      ElMessage.error(result.message || '上传失败')
    }
  } catch {
    ElMessage.error('上传失败')
  } finally {
    uploading.value = false
  }
}

</script>

<template>
  <div class="admin-config">
    <div class="page-header">
      <h2>站点配置</h2>
      <el-button type="primary" :loading="loading" @click="handleSave">保存配置</el-button>
    </div>

    <el-card shadow="hover">
      <el-form label-width="120px">
        <el-form-item label="站点标题">
          <el-input v-model="form.site_title" placeholder="ZyqHome" />
        </el-form-item>

        <el-form-item label="站点描述">
          <el-input v-model="form.site_description" placeholder="我的导航" />
        </el-form-item>

        <el-form-item label="背景图片">
          <div class="bg-upload">
            <el-upload
              :show-file-list="false"
              :before-upload="(file: any) => { handleUpload(file); return false }"
              accept="image/*"
            >
              <el-button :loading="uploading">上传图片</el-button>
            </el-upload>
            <span class="hint">支持 jpg/png/gif/webp，最大 10MB</span>
          </div>
          <div v-if="form.background_image" class="bg-preview">
            <img :src="form.background_image" alt="背景预览" />
          </div>
        </el-form-item>

        <el-form-item label="背景颜色">
          <el-color-picker v-model="form.background_color" />
          <span class="hint" style="margin-left: 12px">无背景图时使用</span>
        </el-form-item>

        <el-form-item label="默认搜索引擎">
          <el-select v-model="form.default_search">
            <el-option label="Google" value="google" />
            <el-option label="百度" value="baidu" />
            <el-option label="Bing" value="bing" />
            <el-option label="GitHub" value="github" />
          </el-select>
        </el-form-item>

        <el-form-item label="搜索框占位符">
          <el-input v-model="form.search_placeholder" placeholder="搜索..." />
        </el-form-item>

        <el-form-item label="显示中心图标">
          <el-switch v-model="form.show_center_icons" :active-value="'true'" :inactive-value="'false'" />
        </el-form-item>

        <el-form-item label="显示右侧菜单">
          <el-switch v-model="form.show_right_menu" :active-value="'true'" :inactive-value="'false'" />
        </el-form-item>

        <el-form-item label="图标列数">
          <el-input-number v-model="form.icon_columns" :min="3" :max="8" />
        </el-form-item>
      </el-form>
    </el-card>
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

.bg-upload {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
}

.hint {
  font-size: 12px;
  color: #999;
}

.bg-preview {
  width: 200px;
  height: 112px;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid #e0e0e0;
}

.bg-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>
