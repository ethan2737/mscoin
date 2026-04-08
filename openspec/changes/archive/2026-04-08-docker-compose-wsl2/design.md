## Context

**当前状态：**
- Docker Compose 配置在 `mscoin-backend/docker-compose.yml`
- 环境变量在 `mscoin-backend/.env`
- 路径配置指向原开发机器：`D:/go/project/microservice/docker/...`
- 当前环境：Windows 11 + WSL2 + Docker Desktop

**基础设施服务：**
- MySQL 8.0 - 数据库
- Redis 6.2 - 缓存
- MongoDB 4.2 - 时序数据
- Etcd 3.5 - 服务发现
- Kafka - 消息队列

## Goals / Non-Goals

**Goals:**
- Docker Compose 可在当前 WSL2 环境启动
- 数据目录位于项目根目录 `docker/` 下，便于管理
- 宿主机端口映射统一为 3 开头
- 提供便捷的启动/停止脚本
- 支持 `docker-compose up -d` 一键启动

**Non-Goals:**
- 不变更容器镜像版本
- 不修改服务配置文件（etc/conf.yaml）
- 不引入新的容器服务
- 不变更网络配置模式

## Decisions

### 1. 数据目录位置

**决策**: 在项目根目录创建 `docker/` 目录，所有数据通过绑定挂载存储

**目录结构:**
```
mscoin/
├── docker/
│   ├── mysql/
│   │   ├── data/      # MySQL 数据文件
│   │   ├── conf/      # MySQL 配置文件
│   │   └── logs/      # MySQL 日志
│   ├── redis/
│   │   ├── data/      # Redis 数据文件
│   │   └── conf/      # Redis 配置文件
│   ├── mongo/
│   │   ├── db/        # MongoDB 数据文件
│   │   └── log/       # MongoDB 日志
│   ├── etcd/
│   │   └── data/      # Etcd 数据文件
│   └── kafka/
│       └── data/      # Kafka 数据文件
└── mscoin-backend/
    ├── docker-compose.yml
    └── .env
```

**理由:**
- 数据与代码同仓，便于版本控制和备份
- 避免绝对路径依赖，支持多环境
- 清理项目时一并删除 Docker 数据

**替代方案考虑:**
- 保持原路径：不可行，原路径不存在
- 使用 Docker Volume：数据不易访问和备份

### 2. 端口映射规范

**决策**: 宿主机端口统一为 3 开头

| 服务 | 容器端口 | 原宿主机端口 | 新宿主机端口 |
|------|----------|--------------|--------------|
| MySQL | 3306 | 3309 | 33306 |
| Redis | 6379 | 6379 | 33679 |
| MongoDB | 27017 | 27018 | 37018 |
| Etcd | 2379/2380 | 2379/2380 | 32379/32380 |
| Kafka | 9092 | 9092 | 39092 |
| Zookeeper | 2181 | 2181 | 32181 |

**理由:**
- 避免与系统服务冲突（如本地 MySQL 默认 3306）
- 统一命名规范，易于识别
- 3 开头便于记忆（3 = infrastructure）

### 3. WSL2 兼容性

**决策**: 使用 Unix 风格路径，兼容 WSL2

**配置要点:**
- `.env` 文件使用相对路径（如 `./docker/mysql/data`）
- `docker-compose.yml` 引用 `.env` 中的变量
- Windows 路径使用 WSL2 互操作路径（如 `/mnt/d/...`）或改用项目内路径

### 4. 启动脚本

**决策**: 在 `mscoin-backend/` 目录添加批处理脚本

**脚本列表:**
- `docker-start.bat` - Windows 下启动所有服务
- `docker-stop.bat` - Windows 下停止所有服务
- `docker-start.sh` - WSL/Linux 下启动所有服务
- `docker-stop.sh` - WSL/Linux 下停止所有服务

## Risks / Trade-offs

| 风险 | 影响 | 缓解措施 |
|------|------|----------|
| 数据目录初始化 | 首次启动需创建目录结构 | 启动脚本自动创建 |
| 端口冲突 | 3 开头端口可能被占用 | 启动前检查，冲突时报错提示 |
| WSL2 内存限制 | Docker 容器可能内存不足 | 配置 `.wslconfig` 限制 Docker 内存 |
| 文件权限问题 | WSL2 与 Windows 文件权限差异 | 使用 Docker 用户运行，避免权限问题 |
| 数据备份 | 数据与代码混在一起 | 文档说明如何排除 Docker 数据 |

## Migration Plan

### 步骤 1: 创建目录结构
```bash
mkdir -p docker/{mysql/{data,conf,logs},redis/{data,conf},mongo/{db,log},etcd/data,kafka/data}
```

### 步骤 2: 更新 .env 文件
- 将绝对路径改为相对路径
- 更新端口映射为 3 开头

### 步骤 3: 更新 docker-compose.yml
- 使用 `.env` 中的变量
- 调整卷挂载路径

### 步骤 4: 创建启动脚本
- Windows 批处理脚本
- WSL/Linux shell 脚本

### 步骤 5: 测试验证
- 启动所有服务
- 验证各服务端口可访问
- 后端服务可正常连接

### 回滚步骤
1. 停止 Docker Compose：`docker-compose down`
2. 恢复原 `.env` 备份
3. 恢复原 `docker-compose.yml` 备份

## Open Questions

1. **是否需要数据迁移**: 如果原 Docker 数据需要保留，是否需要迁移脚本？
2. **.gitignore 配置**: Docker 数据目录是否应加入 `.gitignore`？（建议加入）
