<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getMemberListApi, getMemberDetailApi, updateMemberStatusApi } from '@/api/member'
import { adminGetUserIconsApi } from '@/api/userIcon'
import type { User, UserIcon } from '@/types'

const members = ref<User[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(20)
const loading = ref(false)

const detailVisible = ref(false)
const detailUser = ref<User | null>(null)
const userIcons = ref<UserIcon[]>([])

async function fetchMembers() {
  loading.value = true
  try {
    const res = await getMemberListApi({ page: page.value, page_size: pageSize.value })
    if (res.data.code === 0) {
      members.value = res.data.data.items
      total.value = res.data.data.total
    }
  } catch {
    ElMessage.error('获取会员列表失败')
  } finally {
    loading.value = false
  }
}

async function showDetail(id: number) {
  try {
    const [userRes, iconRes] = await Promise.all([
      getMemberDetailApi(id),
      adminGetUserIconsApi(id),
    ])
    if (userRes.data.code === 0) {
      detailUser.value = userRes.data.data
      userIcons.value = iconRes.data.code === 0 ? iconRes.data.data : []
      detailVisible.value = true
    } else {
      ElMessage.error(userRes.data.message || '获取详情失败')
    }
  } catch {
    ElMessage.error('获取会员详情失败')
  }
}

async function toggleStatus(user: User) {
  const newStatus = user.status === 1 ? 0 : 1
  const actionText = newStatus === 0 ? '禁用' : '启用'

  try {
    await ElMessageBox.confirm(
      `确定要${actionText}该会员账号吗？${newStatus === 0 ? '禁用后该会员将无法登录。' : ''}`,
      '提示'
    )
  } catch {
    return
  }

  try {
    const res = await updateMemberStatusApi(user.id, newStatus)
    if (res.data.code === 0) {
      ElMessage.success(`${actionText}成功`)
      user.status = newStatus
    } else {
      ElMessage.error(res.data.message || `${actionText}失败`)
    }
  } catch {
    ElMessage.error(`${actionText}失败`)
  }
}

function onPageChange(newPage: number) {
  page.value = newPage
  fetchMembers()
}

onMounted(fetchMembers)
</script>

<template>
  <div class="members-page">
    <h2>会员管理</h2>

    <el-card shadow="hover">
      <el-table
        :data="members"
        v-loading="loading"
        stripe
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="60" />
        <el-table-column prop="username" label="用户名" width="120" />
        <el-table-column prop="nickname" label="昵称" width="120" />
        <el-table-column prop="email" label="邮箱" min-width="160" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'" size="small">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="last_login_ip" label="最后登录IP" width="140" />
        <el-table-column prop="last_login_at" label="最后登录时间" width="170" />
        <el-table-column prop="created_at" label="注册时间" width="170" />
        <el-table-column label="操作" width="160" fixed="right">
          <template #default="{ row }">
            <el-button text type="primary" size="small" @click="showDetail(row.id)">
              详情
            </el-button>
            <el-button
              text
              :type="row.status === 1 ? 'danger' : 'success'"
              size="small"
              @click="toggleStatus(row)"
            >
              {{ row.status === 1 ? '禁用' : '启用' }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-wrapper">
        <el-pagination
          v-model:current-page="page"
          :page-size="pageSize"
          :total="total"
          layout="total, prev, pager, next"
          @current-change="onPageChange"
        />
      </div>
    </el-card>

    <!-- 详情弹窗 -->
    <el-dialog
      v-model="detailVisible"
      title="会员详情"
      width="550px"
    >
      <el-descriptions v-if="detailUser" :column="2" border>
        <el-descriptions-item label="ID" :span="1">{{ detailUser.id }}</el-descriptions-item>
        <el-descriptions-item label="用户名" :span="1">{{ detailUser.username }}</el-descriptions-item>
        <el-descriptions-item label="昵称" :span="1">{{ detailUser.nickname || '-' }}</el-descriptions-item>
        <el-descriptions-item label="邮箱" :span="1">{{ detailUser.email || '-' }}</el-descriptions-item>
        <el-descriptions-item label="状态" :span="1">
          <el-tag :type="detailUser.status === 1 ? 'success' : 'danger'" size="small">
            {{ detailUser.status === 1 ? '正常' : '禁用' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="注册时间" :span="1">{{ detailUser.created_at }}</el-descriptions-item>
        <el-descriptions-item label="最后登录IP" :span="2">{{ detailUser.last_login_ip || '-' }}</el-descriptions-item>
        <el-descriptions-item label="最后登录时间" :span="2">{{ detailUser.last_login_at || '-' }}</el-descriptions-item>
      </el-descriptions>

      <!-- 自定义图标 -->
      <div v-if="userIcons.length > 0" style="margin-top:16px;">
        <el-divider />
        <h4 style="margin-bottom:12px;font-size:14px;color:#666;">自定义图标（{{ userIcons.length }} 个）</h4>
        <el-table :data="userIcons" stripe size="small" style="width:100%;">
          <el-table-column label="图片" width="70">
            <template #default="{ row }">
              <el-image :src="row.image_url" fit="cover" style="width:40px;height:40px;border-radius:4px;" />
            </template>
          </el-table-column>
          <el-table-column label="图片URL" min-width="180">
            <template #default="{ row }">
              <el-link :href="row.image_url" target="_blank" :ellipsis="true" style="font-size:12px;">{{ row.image_url }}</el-link>
            </template>
          </el-table-column>
          <el-table-column label="跳转URL" min-width="180">
            <template #default="{ row }">
              <el-link :href="row.link_url" target="_blank" :ellipsis="true" style="font-size:12px;">{{ row.link_url }}</el-link>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" width="160">
            <template #default="{ row }">{{ row.created_at }}</template>
          </el-table-column>
        </el-table>
      </div>
    </el-dialog>
  </div>
</template>

<style scoped>
.members-page h2 {
  margin-bottom: 20px;
  font-size: 24px;
  color: #333;
}

.pagination-wrapper {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}
</style>
