# Simple Go Web应用框架

一个基于Go语言的现代化Web应用框架，采用清晰的分层架构设计，集成了常用的中间件和工具，适合快速构建高性能、可扩展的后台服务。

## 项目特点

- 🌟 **分层架构**：采用清晰的分层设计，分离业务逻辑与基础设施
- 🚀 **高性能**：基于Gin框架构建，提供高性能的HTTP服务
- 🔄 **集成ORM**：使用GORM进行数据库操作，支持读写分离
- 🔐 **JWT认证**：内置JWT认证机制，支持令牌刷新和黑名单
- 📝 **结构化日志**：基于Zap的高性能日志系统，支持日志分割和轮转
- 🔍 **链路追踪**：集成OpenTelemetry进行分布式追踪
- 🗄️ **缓存支持**：Redis缓存集成，支持多种模式（单机、集群、哨兵）
- ⚙️ **热重载配置**：基于Viper的配置管理，支持配置热重载

## 目录结构

```
.
├── cmd/                  # 命令行工具
├── internal/             # 内部包，不对外暴露
│   ├── global/           # 全局变量和状态
│   ├── logic/            # 业务逻辑实现
│   └── types/            # 内部类型定义
├── model/                # 数据模型定义
├── pkg/                  # 可重用的包
│   ├── cache/            # 缓存实现
│   ├── config/           # 配置管理
│   ├── consts/           # 常量定义
│   ├── database/         # 数据库操作
│   ├── logger/           # 日志处理
│   └── resp/             # HTTP响应处理
├── resource/             # 资源文件
│   └── config/           # 配置文件
├── .gitignore            # Git忽略文件
├── go.mod                # Go模块文件
├── go.sum                # Go依赖校验文件
└── main.go               # 程序入口文件
```

## 技术栈

- **Web框架**: [Gin](https://github.com/gin-gonic/gin)
- **ORM框架**: [GORM](https://gorm.io/)
- **配置管理**: [Viper](https://github.com/spf13/viper)
- **日志框架**: [Zap](https://github.com/uber-go/zap)
- **缓存**: [Redis](https://github.com/redis/go-redis)
- **链路追踪**: [OpenTelemetry](https://opentelemetry.io/)

## 快速开始

### 前置要求

- Go 1.21+
- MySQL 5.7+
- Redis 6.0+

### 安装

1. 克隆仓库
```bash
git clone https://github.com/tokmz/simple.git
cd simple
```

2. 安装依赖
```bash
go mod download
```

3. 修改配置
```bash
# 复制并修改配置文件
cp resource/config/config.yaml resource/config/config.local.yaml
# 根据环境修改配置文件
```

4. 运行应用
```bash
go run main.go
```

### 目录结构说明

- **cmd/**: 包含CLI工具和应用入口
- **internal/**: 包含不对外导出的包
  - **global/**: 全局变量和状态管理
  - **logic/**: 业务逻辑的实现
  - **types/**: 内部类型和数据结构定义
- **model/**: 数据模型定义
- **pkg/**: 可被外部项目导入的公共包
  - **cache/**: 缓存实现
  - **config/**: 配置管理
  - **consts/**: 常量定义
  - **database/**: 数据库连接和操作
  - **logger/**: 日志工具
  - **resp/**: HTTP响应和错误处理

## 功能模块

### 用户角色权限管理

- 角色管理：创建、更新、删除角色
- 用户管理：创建和管理用户，分配角色
- 权限控制：基于角色的权限控制

### 配置管理

支持多环境配置：
- 开发环境
- 测试环境
- 生产环境

### 日志管理

- 分级日志：Debug, Info, Warn, Error, Fatal
- 日志轮转：基于大小和时间的日志分割
- 结构化日志：JSON格式输出，便于解析和搜索

## 开发指南

### 添加新路由

1. 在 `internal/types/dto` 中定义请求和响应结构
2. 在 `internal/logic` 中实现业务逻辑
3. 在路由配置中注册新路由

### 数据库操作

项目使用GORM作为ORM框架，支持：
- 读写分离
- 自动生成模型代码
- 事务管理

### 错误处理

统一的错误处理机制：
- 业务逻辑中返回预定义错误
- 通过中间件统一处理并转换为HTTP响应

## 生产部署

### 推荐环境

- Linux (Ubuntu 20.04+)
- 2核4GB以上服务器
- MySQL 5.7+
- Redis 6.0+

### 部署方式

1. 构建可执行文件
```bash
go build -o app main.go
```

2. 使用systemd或supervisor管理进程

3. 使用Nginx作为反向代理

## 贡献指南

1. Fork本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 许可证

[MIT License](LICENSE)

## 联系方式

- 项目维护者: [清风](https://github.com/tokmz)
- 邮箱: your-email@example.com 