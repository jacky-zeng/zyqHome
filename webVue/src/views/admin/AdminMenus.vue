<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useMenuStore } from '@/stores/menuStore'
import type { MenuItem } from '@/types'
import { getIcon } from '@/icons'
import IconPicker from '@/components/common/IconPicker.vue'

const menuStore = useMenuStore()
const menus = ref<MenuItem[]>([])
const loading = ref(false)
const dialogVisible = ref(false)
const editingItem = ref<Partial<MenuItem>>({})
const isEdit = ref(false)
const iconPickerVisible = ref(false)

const defaultForm: Partial<MenuItem> = {
  title: '',
  url: '',
  icon: '',
  sort_order: 0,
  is_active: true,
  target: '_blank',
}

onMounted(async () => {
  await loadMenus()
})

async function loadMenus() {
  loading.value = true
  try {
    menus.value = await menuStore.fetchAllMenus()
  } finally {
    loading.value = false
  }
}

function openCreate() {
  isEdit.value = false
  editingItem.value = { ...defaultForm }
  dialogVisible.value = true
}

function openEdit(item: MenuItem) {
  isEdit.value = true
  editingItem.value = { ...item }
  dialogVisible.value = true
}

async function handleSave() {
  if (!editingItem.value.title) {
    ElMessage.warning('请输入菜单名称')
    return
  }
  try {
    if (isEdit.value && editingItem.value.id) {
      await menuStore.update(editingItem.value.id, editingItem.value)
      ElMessage.success('更新成功')
    } else {
      await menuStore.create(editingItem.value)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    await loadMenus()
  } catch {
    ElMessage.error('操作失败')
  }
}

async function handleDelete(id: number) {
  try {
    await ElMessageBox.confirm('确定要删除该菜单吗？', '提示')
    await menuStore.remove(id)
    ElMessage.success('删除成功')
    await loadMenus()
  } catch {
    // cancelled
  }
}

</script>

<template>
  <div class="admin-menus">
    <div class="page-header">
      <h2>菜单管理</h2>
      <el-button type="primary" @click="openCreate">添加菜单</el-button>
    </div>

    <el-card shadow="hover">
      <el-table :data="menus" v-loading="loading" row-key="id" style="width: 100%">
        <el-table-column prop="id" label="ID" width="60" />
        <el-table-column label="图标" width="70">
          <template #default="{ row }">
            <component
              v-if="row.icon && getIcon(row.icon)"
              :is="getIcon(row.icon)"
              :size="20"
              style="color: #666; vertical-align: middle;"
            />
            <span v-else class="icon-placeholder">-</span>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="名称" min-width="120" />
        <el-table-column prop="url" label="链接" min-width="200">
          <template #default="{ row }">
            <span class="url-text">{{ row.url || '-' }}</span>
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
      :title="isEdit ? '编辑菜单' : '添加菜单'"
      width="500px"
    >
      <el-form label-width="80px">
        <el-form-item label="名称" required>
          <el-input v-model="editingItem.title" placeholder="菜单名称" />
        </el-form-item>
        <el-form-item label="链接">
          <el-input v-model="editingItem.url" placeholder="链接地址，留空则为分组" />
        </el-form-item>
        <el-form-item label="图标">
          <div class="icon-picker-trigger">
            <el-button class="icon-select-btn" @click="iconPickerVisible = true">
              <template v-if="editingItem.icon && getIcon(editingItem.icon)">
                <component :is="getIcon(editingItem.icon)" :size="18" style="vertical-align: middle; margin-right: 4px;" />
                <span>{{ editingItem.icon }}</span>
              </template>
              <span v-else>选择图标</span>
            </el-button>
            <el-button v-if="editingItem.icon" size="small" type="danger" link @click="editingItem.icon = ''">
              清除
            </el-button>
          </div>
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="editingItem.sort_order" :min="0" />
        </el-form-item>
        <el-form-item label="打开方式">
          <el-select v-model="editingItem.target">
            <el-option label="新窗口(_blank)" value="_blank" />
            <el-option label="当前窗口(_self)" value="_self" />
          </el-select>
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

    <!-- Icon Picker -->
    <IconPicker v-model="iconPickerVisible" @select="(name: string) => editingItem.icon = name" />
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

.url-text {
  color: #409eff;
  font-size: 13px;
}

.icon-placeholder {
  color: #ccc;
}

.icon-picker-trigger {
  display: flex;
  align-items: center;
  gap: 8px;
}

.icon-select-btn {
  min-width: 120px;
}
</style>
