# AICoreOps

AI 驱动的云原生运维平台（微服务版）

## 目录

- [AICoreOps](#AICoreOps)
  - [目录](#目录)
  - [项目介绍](#项目介绍)
  - [快速开始](#快速开始)
    - [克隆项目](#克隆项目)
    - [运行服务](#运行服务)
  - [项目结构](#项目结构)
  - [许可证](#许可证)
  - [联系方式](#联系方式)
  - [致谢](#致谢)

## 项目介绍

AICoreOps 是一个基于 **go-zero** 的微服务架构的 AI 驱动云原生运维管理平台，旨在通过人工智能技术提升运维效率和智能化水平。平台包含以下核心模块：

1. **AIOps 模块**：通过机器学习和 AI 技术，分析系统日志、监控数据，提供智能告警、故障预测和根因分析。
2. **用户与权限**：管理用户、角色及权限，确保系统的安全和可控性。
3. **服务树与 CMDB**：提供可视化的服务树结构和配置管理数据库（CMDB），实现运维资源的全面管理。
4. **工单系统**：支持工单的创建、分配、处理和追踪，提高问题解决效率。
5. **Prometheus 集成**：实时监控系统性能，结合 AI 技术，进行异常预警和自动化响应。
6. **Kubernetes 管理**：支持 Kubernetes 集群的管理与监控，简化云端资源操作，集成 AI 进行自动化优化和资源调度。

## 快速开始

### 克隆项目

首先，将项目克隆到本地：

```bash
git clone https://github.com/GoSimplicity/AICoreOps.git
```

### 运行服务

#### 环境准备

确保已安装以下环境：

- **Go** 1.20 或更高版本
- **MySQL** 或其他数据库
- **Redis** 用于缓存
- **Docker**（可选，用于容器化运行）

#### 安装依赖

在项目根目录下运行以下命令以安装依赖：

```bash
go mod tidy
```

#### 配置环境

根据 `config` 文件夹中的示例配置文件，填写数据库、Redis 和其他必要的配置。

#### 启动服务

运行以下命令以启动微服务：

```bash
# 启动用户服务
go run services/user/cmd/main.go

# 启动权限服务
go run services/permission/cmd/main.go

# 启动运维服务
go run services/ops/cmd/main.go

# 启动 AI 模块服务
go run services/ai/cmd/main.go

# 启动网关
go run gateway/main.go
```

#### 前端项目

前端项目地址：<https://github.com/GoSimplicity/AICoreOps-web>

参见前端项目文档进行安装与启动。

## 项目结构

```text
AICoreOps/
├── LICENSE                      # 项目许可证文件
├── Makefile                     # 项目构建和管理脚本
├── README.md                    # 项目说明文档
├── common                       # 公共模块
│   ├── constants                # 常量定义模块
│   ├── middleware               # 中间件文件夹
│   └── utils                    # 工具模块
├── deploy                       # 部署相关文件
│   ├── cicd                     # CI/CD 文件夹
│   │   └── scripts              # 构建和部署脚本
│   └── k8s                      # Kubernetes 配置文件
├── docs                         # 项目文档
├── proto                        # gRPC 协议文件夹
│   ├── k8s                      # Kubernetes gRPC 协议
│   ├── prometheus               # Prometheus gRPC 协议
│   ├── role                     # 角色相关 gRPC 协议
│   ├── tree                     # 服务树相关 gRPC 协议
│   └── user                     # 用户相关 gRPC 协议
└── services                     # 微服务文件夹
    ├── aicoreops_api            # AICoreOps API 服务
    │   ├── etc                  # 配置文件夹
    │   └── internal             # 内部模块
    │       ├── config           # 配置模块
    │       │   └── config.go    # 配置文件
    │       ├── handler          # 处理器模块
    │       │   └── routes.go    # 路由配置
    │       ├── logic            # 业务逻辑
    │       ├── svc              # 服务上下文
    │       └── types            # 类型定义
    ├── k8s_rpc                  # Kubernetes RPC 服务
    ├── prometheus_rpc           # Prometheus RPC 服务
    ├── role_rpc                 # 角色 RPC 服务
    ├── tree_rpc                 # 服务树 RPC 服务
    └── user_rpc                 # 用户 RPC 服务
```

## 许可证

本项目使用 [Apache 2.0许可证](./LICENSE)，详情请查看 LICENSE 文件。

## 联系方式

如果有任何问题或建议，欢迎通过以下方式联系我：

- Email: [wzijian62@gmail.com](mailto:wzijian62@gmail.com)
- 微信：GoSimplicity（加我后可邀请进微信群交流）

## 致谢

感谢所有为本项目贡献代码、文档和建议的人！AICoreOps 的发展离不开社区的支持和贡献。