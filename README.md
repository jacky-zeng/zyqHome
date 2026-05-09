# ZyqHome — 个人导航主页

一个功能完整的个人浏览器起始页项目，支持自定义导航菜单、快捷图标、搜索引擎切换，并附带完整的管理后台。前后端完全分离部署。

## 技术栈

| 层级 | 技术 | 版本 |
|------|------|------|
| **后端** | Go + Gin + GORM + MySQL + Redis | Go 1.26 |
| **前端** | Vue 3 + TypeScript + Vite + Pinia + Element Plus | Vue 3.5 |
| **缓存** | Redis 7.x | — |
| **认证** | JWT + bcrypt | — |

## 项目结构

```
zyqHome/
├── backProject/          # Go 后端服务 (Gin 框架)
│   ├── cmd/server/       # 程序入口
│   ├── internal/         # 核心代码 (config/database/model/repository/service/handler/middleware/router)
│   ├── pkg/              # 工具包 (JWT / 统一响应)
│   └── uploads/          # 上传文件存储
└── webVue/               # Vue3 前端应用
    └── src/
        ├── api/          # Axios 接口封装
        ├── components/   # 可复用组件
        ├── composables/  # 组合式函数 (时钟/农历/行为追踪)
        ├── icons/        # lucide 图标注册表
        ├── router/       # Vue Router 路由
        ├── stores/       # Pinia 状态管理
        ├── types/        # TypeScript 类型定义
        └── views/        # 页面 (首页/登录/管理后台)
```

## 功能模块

| 模块 | 说明 |
|------|------|
| **首页展示** | 实时时钟 + 农历日期、多搜索引擎搜索栏、中心图标网格、侧边菜单栏、全屏背景 |
| **菜单管理** | 无限级树形菜单、拖拽排序、lucide 图标选择 |
| **图标管理** | 首页快捷图标 CRUD、自定义背景色、关联菜单动态筛选 |
| **图片管理** | 图片上传/分类/删除、从图库选择背景图 |
| **站点配置** | 11 项自定义配置（标题/背景/搜索引擎/页脚等） |
| **用户认证** | JWT 登录、路由守卫、Token 持久化 |
| **行为追踪** | 记录页面访问、点击和搜索行为（IP + 操作类型） |

## 快速启动

### 后端

```bash
cd backProject
# 确保 MySQL 和 Redis 已运行，配置 .env 文件
go run cmd/server/main.go
# 服务默认监听 :8090
```

### 前端

```bash
cd webVue
npm install
npm run dev
# 开发服务器默认监听 :5173，自动代理 /api 到后端
```

首次启动后端会自动创建数据库表并生成默认管理员账号 `admin / admin123`。
