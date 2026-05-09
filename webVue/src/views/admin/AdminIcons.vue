<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useIconStore } from '@/stores/iconStore'
import type { CenterIcon } from '@/types'

const iconStore = useIconStore()
const icons = ref<CenterIcon[]>([])
const loading = ref(false)
const dialogVisible = ref(false)
const editingItem = ref<Partial<CenterIcon>>({})
const isEdit = ref(false)

const defaultForm: Partial<CenterIcon> = {
  title: '',
  url: '',
  icon: '',
  color: '#1890ff',
  sort_order: 0,
  is_active: true,
}

onMounted(async () => {
  await loadIcons()
})

async function loadIcons() {
  loading.value = true
  try {
    icons.value = await iconStore.fetchAllIcons()
  } finally {
    loading.value = false
  }
}

function openCreate() {
  isEdit.value = false
  editingItem.value = { ...defaultForm }
  dialogVisible.value = true
}

function openEdit(item: CenterIcon) {
  isEdit.value = true
  editingItem.value = { ...item }
  dialogVisible.value = true
}

async function handleSave() {
  if (!editingItem.value.title) {
    ElMessage.warning('请输入图标名称')
    return
  }
  try {
    if (isEdit.value && editingItem.value.id) {
      await iconStore.update(editingItem.value.id, editingItem.value)
      ElMessage.success('更新成功')
    } else {
      await iconStore.create(editingItem.value)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    await loadIcons()
  } catch {
    ElMessage.error('操作失败')
  }
}

async function handleDelete(id: number) {
  try {
    await ElMessageBox.confirm('确定要删除该图标吗？', '提示')
    await iconStore.remove(id)
    ElMessage.success('删除成功')
    await loadIcons()
  } catch {
    // cancelled
  }
}

</script>

<template>
  <div class="admin-icons">
    <div class="page-header">
      <h2>图标管理</h2>
      <el-button type="primary" @click="openCreate">添加图标</el-button>
    </div>

    <el-card shadow="hover">
      <el-table :data="icons" v-loading="loading" row-key="id" style="width: 100%">
        <el-table-column prop="id" label="ID" width="60" />
        <el-table-column label="预览" width="80">
          <template #default="{ row }">
            <div class="icon-preview" :style="{ background: row.color }">
              <span>{{ row.title.charAt(0) }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="名称" min-width="120" />
        <el-table-column prop="url" label="链接" min-width="200">
          <template #default="{ row }">
            <span class="url-text">{{ row.url || '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="color" label="颜色" width="100">
          <template #default="{ row }">
            <div class="color-display">
              <span class="color-dot" :style="{ background: row.color }"></span>
              <span>{{ row.color }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="sort_order" label="排序" width="70" />
        <el-table-column prop="is_active" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.is_active ? 'success' : 'info'" size="small">
              {{ row.is_active ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="160" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click="openEdit(row)">编辑</el-button>
            <el-button type="danger" link size="small" @click="handleDelete(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Create/Edit Dialog -->
    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑图标' : '添加图标'"
      width="500px"
    >
      <el-form label-width="80px">
        <el-form-item label="名称" required>
          <el-input v-model="editingItem.title" placeholder="图标名称" />
        </el-form-item>
        <el-form-item label="链接" required>
          <el-input v-model="editingItem.url" placeholder="跳转链接" />
        </el-form-item>
        <el-form-item label="图标URL">
          <el-input v-model="editingItem.icon" placeholder="图标图片URL，留空显示首字" />
        </el-form-item>
        <el-form-item label="背景色">
          <el-color-picker v-model="editingItem.color" />
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="editingItem.sort_order" :min="0" />
        </el-form-item>
        <el-form-item label="启用">
          <el-switch v-model="editingItem.is_active" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSave">保存</el-button>
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

.icon-preview {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 14px;
  font-weight: 600;
}

.color-display {
  display: flex;
  align-items: center;
  gap: 6px;
}

.color-dot {
  width: 16px;
  height: 16px;
  border-radius: 4px;
  display: inline-block;
}

.url-text {
  color: #409eff;
  font-size: 13px;
}
</style>
