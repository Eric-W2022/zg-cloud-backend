# ZCloud-BG

ZCloud-BG 是一个基于Go-Gin-GORM框架的数字员工后端服务项目，采用了模块化的设计理念，以便于维护和扩展。

## 项目结构

以下是项目的目录结构：

```plaintext
zcloud-bg/
│
├── internal/                   # 内部包，业务逻辑
│   ├── handler/                # 处理 HTTP 请求的处理器
│   │   ├── auth.go             # 身份验证相关的处理器
│   │   └── user.go             # 用户相关的处理器
│   │
│   ├── model/                  # 数据模型（与数据库表相对应）
│   │   ├── user.go             # 用户模型
│   │   └── messages.go         # 消息模型
│   │
│   ├── service/                # 业务逻辑服务
│   │   ├── auth_service.go     # 身份验证服务
│   │   └── user_service.go     # 用户服务
│   │
│   └── repository/             # 数据访问层，与数据库交互
│       ├── user_repo.go        # 用户数据访问
│       └── token_repo.go       # 令牌数据访问
│
├── pkg/                        # 公共库（可由其他项目引用）
│   ├── jwt/                    # JWT 功能相关的工具
│   ├── router/                 # 路由相关配置
│   └── db/                     # 数据库相关的工具和配置
│
├── .env                        # 环境变量文件
│
├── main.go                     # 主程序入口
│
└── go.mod                      # Go 依赖管理文件


## 开发和部署

- 确保已经安装了 Go 语言环境。
- 克隆仓库后，在项目根目录运行 `go build` 或 `go run main.go` 来启动项目。
- 云端部署见Dockerfile

## 贡献

由智工科技团队参与开发维护


