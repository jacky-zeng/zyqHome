# ZyqHome 前端应用

Vue 3 + TypeScript + Vite 个人导航主页前端，支持自定义首页导航、管理后台和搜索引擎切换。

## 技术栈

| 技术 | 用途 | 版本 |
|------|------|------|
| **Vue 3** | 前端框架 (Composition API + `<script setup>`) | ^3.5 |
| **TypeScript** | 编程语言 | ~6.0 |
| **Vite** | 构建工具 | ^8.0 |
| **Pinia** | 状态管理 | ^3.0 |
| **Vue Router** | 路由管理 | ^4.6 |
| **Element Plus** | UI 组件库 | ^2.14 |
| **Axios** | HTTP 客户端 | ^1.16 |
| **lucide-vue-next** | 开源图标库 (~130 个图标) | ^1.0 |
| **vuedraggable** | 拖拽排序组件 | ^4.1 |

## 项目结构

```
webVue/
├── public/
│   ├── favicon.svg            # 网站图标
│   └── icons.svg              # SVG 图标精灵
├── src/
│   ├── main.ts                # 应用入口
│   ├── App.vue                # 根组件 (<router-view/>)
│   ├── api/                   # Axios 接口封装
│   │   ├── request.ts         # Axios 实例 + 拦截器 (Token/401)
│   │   ├── auth.ts            # 认证 API
│   │   ├── menu.ts            # 菜单 API
│   │   ├── icon.ts            # 图标 API
│   │   ├── config.ts          # 配置 API
│   │   ├── image.ts           # 图片 API
│   │   └── behavior.ts        # 行为追踪 API
│   ├── assets/styles/global.css # 全局样式
│   ├── components/
│   │   ├── homepage/          # 首页子组件
│   │   │   ├── PageBackground.vue # 全屏背景 (图片/渐变色)
│   │   │   ├── SearchBar.vue      # 搜索栏 (6 个搜索引擎)
│   │   │   ├── CenterIcons.vue    # 中心图标网格
│   │   │   ├── IconItem.vue       # 单个图标项
│   │   │   └── RightMenu.vue      # 侧边菜单栏
│   │   └── common/            # 通用组件
│   │       ├── IconPicker.vue     # lucide 图标选择器
│   │       └── ImagePicker.vue    # 图库图片选择器
│   ├── composables/           # 组合式函数
│   │   ├── useAuth.ts         # 认证状态工具
│   │   ├── useClock.ts        # 实时时钟 (每秒更新)
│   │   ├── useLunarCalendar.ts # 公历转农历 (1900-2100)
│   │   └── useTrackBehavior.ts # 用户行为自动上报
│   ├── icons/index.ts         # lucide 图标注册表
│   ├── router/index.ts        # 路由 + 导航守卫
│   ├── stores/                # Pinia 状态管理
│   │   ├── authStore.ts       # 认证状态 (Token/localStorage)
│   │   ├── configStore.ts     # 站点配置
│   │   ├── menuStore.ts       # 菜单数据
│   │   ├── iconStore.ts       # 图标数据
│   │   └── imageStore.ts      # 图片数据
│   ├── types/index.ts         # TypeScript 类型定义
│   └── views/
│       ├── HomePage.vue       # 首页 (导航主页)
│       ├── LoginPage.vue      # 管理员登录页
│       └── admin/
│           ├── AdminLayout.vue   # 后台布局 (侧边栏 + 内容区)
│           ├── AdminDashboard.vue# 仪表盘
│           ├── AdminMenus.vue    # 菜单管理
│           ├── AdminIcons.vue    # 图标管理
│           ├── AdminImages.vue   # 图片管理
│           └── AdminConfig.vue   # 站点配置
├── index.html
├── vite.config.ts             # Vite 配置 (路径别名 + 代理)
├── tsconfig*.ts               # TypeScript 配置
└── package.json
```

## 路由

| 路径 | 页面 | 需登录 |
|------|------|--------|
| `/` | HomePage — 导航首页 | 否 |
| `/login` | LoginPage — 管理员登录 | 否 |
| `/admin` | AdminDashboard — 仪表盘 | 是 |
| `/admin/menus` | AdminMenus — 菜单管理 | 是 |
| `/admin/icons` | AdminIcons — 图标管理 | 是 |
| `/admin/images` | AdminImages — 图片管理 | 是 |
| `/admin/config` | AdminConfig — 站点配置 | 是 |

未登录访问 `/admin*` 路由时自动跳转到 `/login?redirect=xxx`。

## 首页功能

- **实时时钟**：显示当前时间 + 农历日期
- **搜索引擎**：支持 Google / 百度 / Bing / GitHub / 知乎 / B 站
- **中心图标**：按菜单分类筛选的可配置快捷图标网格
- **侧边菜单**：左侧垂直排列的可配置菜单栏
- **全屏背景**：支持自定义背景图片或渐变色

## 启动方式

```bash
npm install
npm run dev
```

开发服务器默认监听 `:5173`，Vite 自动代理 `/api` 和 `/uploads` 到后端 `localhost:8090`。

## 构建

```bash
npm run build
```

构建产物输出到 `dist/` 目录，可直接部署到任意静态文件服务器。
