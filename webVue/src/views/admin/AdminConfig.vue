<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useConfigStore } from '@/stores/configStore'
import ImagePicker from '@/components/common/ImagePicker.vue'

const configStore = useConfigStore()
const loading = ref(false)
const imagePickerVisible = ref(false)

const form = ref({
  site_title: 'ZyqHome',
  site_description: '我的导航',
  background_image: '',
  background_color: '#f0f2f5',
  default_search: 'google',
  search_placeholder: '搜索...',
  footer_text: '',
  icon_columns: '8',
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

function openImagePicker() {
  imagePickerVisible.value = true
}

function onImageSelected(url: string) {
  form.value.background_image = url
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
            <el-button @click="openImagePicker">从图库选择</el-button>
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

        <el-form-item label="底部文字">
          <el-input v-model="form.footer_text" placeholder="首页底部显示的文字" maxlength="50" show-word-limit />
        </el-form-item>

        <el-form-item label="图标列数">
          <el-input-number v-model="form.icon_columns" :min="3" :max="10" />
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Image Picker -->
    <ImagePicker v-model="imagePickerVisible" @select="onImageSelected" />
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
