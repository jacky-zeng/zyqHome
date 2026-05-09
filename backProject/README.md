# ZyqHome 后端服务

Go + Gin + GORM + MySQL + Redis 个人导航主页后端服务。

## 技术栈

- **框架**: Gin (HTTP 路由和中间件)
- **ORM**: GORM (数据库操作)
- **数据库**: MySQL 8.x
- **缓存**: Redis 7.x
- **认证**: JWT (golang-jwt) + bcrypt 密码加密

## 项目结构

```
backProject/
├── cmd/server/main.go        # 程序入口
├── internal/
│   ├── config/config.go      # 环境变量配置
│   ├── database/
│   │   ├── mysql.go          # MySQL 连接 + AutoMigrate
│   │   └── redis.go          # Redis 客户端
│   ├── middleware/
│   │   ├── auth.go           # JWT 认证中间件
│   │   └── cors.go           # CORS 跨域中间件
│   ├── model/                # 数据模型 (5 个表)
│   │   ├── user.go           # 管理员用户
│   │   ├── menu_item.go      # 右侧菜单项
│   │   ├── center_icon.go    # 中心图标
│   │   ├── site_config.go    # 站点配置
│   │   └── user_behavior.go  # 用户行为
│   ├── repository/           # 数据访问层
│   ├── service/              # 业务逻辑层 (含 Redis 缓存)
│   ├── handler/              # HTTP 接口处理器
│   └── router/router.go      # 路由注册
├── pkg/
│   ├── response/response.go  # 统一响应格式
│   └── jwt/jwt.go            # JWT 工具
└── uploads/                  # 上传文件目录
```

## API 接口

### 公开接口 (无需认证)

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | /api/config | 获取站点配置 |
| GET | /api/icons | 获取启用的图标 |
| GET | /api/menus | 获取启用的菜单树 |
| POST | /api/behavior/track | 记录用户行为 |
| POST | /api/auth/login | 管理员登录 |

### 管理接口 (需 JWT 认证)

| 方法 | 路径 | 说明 |
|---|---|---|
| GET / POST | /api/admin/menus | 菜单 CRUD |
| PUT / DELETE | /api/admin/menus/:id | 菜单单条操作 |
| PUT | /api/admin/menus/sort | 菜单批量排序 |
| GET / POST | /api/admin/icons | 图标 CRUD |
| PUT / DELETE | /api/admin/icons/:id | 图标单条操作 |
| PUT | /api/admin/icons/sort | 图标批量排序 |
| GET / PUT | /api/admin/config | 站点配置读写 |
| POST | /api/admin/upload | 图片上传 |
| GET | /api/admin/auth/me | 当前用户信息 |

## 启动方式

```bash
# 确保 MySQL 和 Redis 已运行

# 使用项目根目录的 .env 配置 (或设置环境变量)
go run cmd/server/main.go
```

服务器默认监听 `:8080`，首次启动会自动创建数据库表并添加默认管理员 `admin / admin123`。

## 数据库表

| 表名 | 说明 |
|---|---|
| users | 管理员用户 |
| menu_items | 右侧菜单项 (支持 parent_id 二级菜单) |
| center_icons | 中心图标 |
| site_configs | 站点 key-value 配置 |
| user_behaviors | 用户行为日志 (IP + 行为类型) |

## 缓存策略

- Redis 缓存键: `zyqhome:config:public`、`zyqhome:icons:active`、`zyqhome:menus:active`
- TTL: 3600 秒
- 管理员更新配置时自动清除缓存
