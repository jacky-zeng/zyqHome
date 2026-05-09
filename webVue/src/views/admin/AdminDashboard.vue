<script setup lang="ts">
</script>

<template>
  <div class="dashboard">
    <h2>仪表盘</h2>
    <el-alert
      title="欢迎使用管理后台"
      type="success"
      :closable="false"
      show-icon
      style="margin-bottom: 20px"
    />
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>
            <span>快捷操作</span>
          </template>
          <div style="display: flex; flex-direction: column; gap: 12px">
            <el-button type="primary" @click="$router.push('/admin/menus')">
              管理菜单
            </el-button>
            <el-button type="success" @click="$router.push('/admin/icons')">
              管理图标
            </el-button>
            <el-button type="warning" @click="$router.push('/admin/config')">
              站点配置
            </el-button>
          </div>
        </el-card>
      </el-col>
      <el-col :span="16">
        <el-card shadow="hover">
          <template #header>
            <span>使用说明</span>
          </template>
          <ul style="line-height: 2; padding-left: 20px">
            <li>菜单管理：配置右侧导航菜单，支持二级菜单和拖拽排序</li>
            <li>图标管理：配置首页中心图标，支持自定义颜色和拖拽排序</li>
            <li>站点配置：修改站点标题、背景图、搜索引擎等全局设置</li>
            <li>用户行为：系统自动记录用户访问和点击行为</li>
          </ul>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.dashboard h2 {
  margin-bottom: 20px;
  font-size: 24px;
  color: #333;
}
</style>
