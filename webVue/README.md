# ZyqHome 前端页面

Vue 3 + TypeScript + Vite 个人导航主页前端。

## 技术栈

- **框架**: Vue 3 (Composition API + `<script setup>`)
- **构建**: Vite
- **路由**: Vue Router 4
- **状态管理**: Pinia
- **UI 库**: Element Plus
- **HTTP**: Axios
- **拖拽**: vuedraggable (next)

## 项目结构

```
webVue/
├── src/
│   ├── api/                     # API 请求封装
│   │   ├── request.ts           # Axios 实例 + 拦截器
│   │   ├── auth.ts              # 认证 API
│   │   ├── menu.ts              # 菜单 API
│   │   ├── icon.ts              # 图标 API
│   │   ├── config.ts            # 配置 API
│   │   └── behavior.ts          # 行为追踪 API
│   ├── assets/styles/           # 全局样式
│   ├── components/
│   │   ├── homepage/            # 首页子组件
│   │   │   ├── SearchBar.vue      # 搜索框
│   │   │   ├── CenterIcons.vue    # 图标网格
│   │   │   ├── IconItem.vue       # 单个图标
│   │   │   ├── RightMenu.vue      # 右侧菜单栏
│   │   │   └── PageBackground.vue # 背景管理
│   │   └── admin/               # 后台组件
│   ├── composables/             # 组合式函数
│   ├── router/index.ts          # 路由定义 + 导航守卫
│   ├── stores/                  # Pinia 状态管理
│   │   ├── authStore.ts         # 用户认证
│   │   ├── menuStore.ts         # 菜单
│   │   ├── iconStore.ts         # 图标
│   │   └── configStore.ts       # 配置
│   ├── types/index.ts           # TypeScript 类型
│   ├── views/                   # 页面组件
│   │   ├── HomePage.vue         # 首页
│   │   ├── LoginPage.vue        # 登录页
│   │   └── admin/               # 管理后台
│   │       ├── AdminLayout.vue    # 后台布局
│   │       ├── AdminDashboard.vue # 仪表盘
│   │       ├── AdminMenus.vue     # 菜单管理
│   │       ├── AdminIcons.vue     # 图标管理
│   │       └── AdminConfig.vue    # 站点配置
│   └── main.ts                  # 程序入口
├── index.html
└── vite.config.ts               # Vite 配置 (路径别名 + 代理)
```

## 路由

| 路径 | 页面 | 说明 | 需登录 |
|---|---|---|---|
| / | HomePage | 公开首页 | 否 |
| /login | LoginPage | 管理员登录 | 否 |
| /admin | AdminDashboard | 仪表盘 | 是 |
| /admin/menus | AdminMenus | 菜单管理 | 是 |
| /admin/icons | AdminIcons | 图标管理 | 是 |
| /admin/config | AdminConfig | 站点配置 | 是 |

## 启动方式

```bash
npm install
npm run dev
```

开发环境默认监听 `:5173`，自动代理 `/api` 请求到 `localhost:8080`。

## 构建

```bash
npm run build
```

构建产物输出到 `dist/` 目录。
