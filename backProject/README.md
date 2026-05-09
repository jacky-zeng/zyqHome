# ZyqHome 后端服务

Go + Gin + GORM + MySQL + Redis 个人导航主页后端服务，提供 RESTful API 和管理后台接口。

## 技术栈

| 技术 | 用途 | 版本 |
|------|------|------|
| **Go** | 编程语言 | 1.26.1 |
| **Gin** | HTTP Web 框架 | v1.12 |
| **GORM** | ORM 框架 | v1.31 |
| **MySQL** | 数据库 | 8.x |
| **Redis** | 缓存 | 7.x |
| **golang-jwt** | JWT 认证 | v5 |
| **bcrypt** | 密码加密 | — |

## 项目结构

```
backProject/
├── cmd/server/main.go         # 程序入口
├── internal/
│   ├── config/config.go       # 环境变量配置加载
│   ├── database/
│   │   ├── mysql.go           # MySQL 连接 + GORM AutoMigrate
│   │   └── redis.go           # Redis 客户端初始化
│   ├── model/                 # 数据模型 (6 张表)
│   │   ├── user.go            # 管理员用户
│   │   ├── menu_item.go       # 菜单项 (无限级树形)
│   │   ├── center_icon.go     # 首页中心图标
│   │   ├── site_config.go     # 站点 KV 配置
│   │   ├── user_behavior.go   # 用户行为日志
│   │   └── image.go           # 图片资源
│   ├── repository/            # 数据访问层 (DAO)
│   ├── service/               # 业务逻辑层 (含 Redis 缓存)
│   ├── handler/               # HTTP 接口处理器
│   ├── middleware/
│   │   ├── auth.go            # JWT 认证中间件
│   │   └── cors.go            # CORS 跨域中间件
│   └── router/router.go       # 路由注册
├── pkg/
│   ├── response/response.go   # 统一 API 响应格式 {code, message, data}
│   └── jwt/jwt.go             # JWT Token 生成/解析工具
└── uploads/                   # 上传文件存储目录
```

## 数据库表

| 表名 | 说明 |
|------|------|
| `users` | 管理员用户 (username, password/bcrypt, nickname, avatar) |
| `menu_items` | 菜单项 (支持 parent_id 无限级树形、排序、启用/禁用) |
| `center_icons` | 首页中心图标 (背景色、关联菜单筛选、排序) |
| `site_configs` | 站点 KV 配置 (11 个配置项) |
| `user_behaviors` | 用户行为日志 (IP + 4 种行为类型 + 目标) |
| `images` | 图片资源 (分类管理、上传记录) |

## API 接口

### 公开接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/config` | 获取站点公开配置 |
| GET | `/api/icons?menu_id=` | 获取启用的图标列表 |
| GET | `/api/menus` | 获取启用的菜单树 |
| POST | `/api/behavior/track` | 记录用户行为 |
| POST | `/api/auth/login` | 管理员登录 |
| 静态 | `/uploads/*` | 静态文件服务 |

### 管理接口（需 JWT Bearer Token）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/admin/auth/me` | 获取当前用户信息 |
| GET | `/api/admin/dashboard` | 仪表盘统计 |
| **菜单管理** | | |
| GET | `/api/admin/menus` | 获取所有菜单（树形） |
| POST | `/api/admin/menus` | 创建菜单 |
| PUT | `/api/admin/menus/:id` | 更新菜单 |
| DELETE | `/api/admin/menus/:id` | 删除菜单 |
| PUT | `/api/admin/menus/sort` | 批量排序菜单 |
| **图标管理** | | |
| GET | `/api/admin/icons` | 获取所有图标 |
| POST | `/api/admin/icons` | 创建图标 |
| PUT | `/api/admin/icons/:id` | 更新图标 |
| DELETE | `/api/admin/icons/:id` | 删除图标 |
| PUT | `/api/admin/icons/sort` | 批量排序图标 |
| **图片管理** | | |
| GET | `/api/admin/images?page=&page_size=&category=` | 分页获取图片 |
| GET | `/api/admin/images/categories` | 获取所有分类 |
| POST | `/api/admin/images` | 上传图片 |
| PUT | `/api/admin/images/:id` | 更新图片信息 |
| DELETE | `/api/admin/images/:id` | 删除图片（同步删物理文件） |
| **站点配置** | | |
| GET | `/api/admin/config` | 获取所有配置 |
| PUT | `/api/admin/config` | 更新配置 |
| **通用** | | |
| POST | `/api/admin/upload` | 通用文件上传 (10MB 限制) |

## 缓存策略

- **缓存键**: `zyqhome:config:public`、`zyqhome:icons:active`、`zyqhome:menus:active`
- **TTL**: 3600 秒
- **失效策略**: 管理员更新对应数据时自动清除缓存

## 启动方式

```bash
# 确保 MySQL 和 Redis 已运行

# 配置环境变量（支持 .env 文件或系统环境变量）
cp .env.example .env
# 编辑 .env 填写数据库和 Redis 配置

# 启动服务
go run cmd/server/main.go
```

服务默认监听 `:8090`，首次启动自动创建数据表并生成默认管理员账号 `admin / admin123`。
