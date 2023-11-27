my-gin-app/
│
├── cmd/                        # 应用程序的入口
│   └── main.go                 # 主程序入口
│
├── internal/                   # 内部包，业务逻辑
│   ├── handler/                # 处理 HTTP 请求的处理器
│   │   ├── auth.go             # 身份验证相关的处理器
│   │   └── user.go             # 用户相关的处理器
│   │
│   ├── model/                  # 数据模型（与数据库表相对应）
│   │   ├── user.go             # 用户模型
│   │   └── token.go            # 令牌模型
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
│   └── db/                     # 数据库相关的工具和配置
│
├── configs/                    # 配置文件
│   └── config.yaml             # YAML 格式的配置文件
│
├── migrations/                 # 数据库迁移脚本
│
├── tests/                      # 测试文件
│
├── .env                        # 环境变量文件
│
└── go.mod                      # Go 依赖管理文件
