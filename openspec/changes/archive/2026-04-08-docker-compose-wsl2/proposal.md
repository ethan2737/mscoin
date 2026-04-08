## Why

当前 MSCOIN 后端服务的 Docker Compose 配置指向原开发机器的绝对路径（`D:/go/project/microservice/docker/...`），无法在 WSL2 + Docker Desktop 环境中启动。这导致升级后的后端服务无法进行集成测试和性能基准验证（任务 6.6 未完成）。本变更旨在改造 Docker 配置，使其适配当前 Windows + WSL2 环境，并将数据目录纳入项目管理。

## What Changes

- **docker-compose.yml 改造**: 将绝对路径改为相对路径，适配 WSL2 + Docker Desktop 环境
- **项目数据目录重构**: 在 `docker/` 目录下管理所有 Docker 数据（MySQL、Redis、MongoDB、Etcd、Kafka）
- **端口映射调整**: 宿主机端口统一为 3 开头（如 3306→33060）
- **.env 配置更新**: 使用相对路径和 WSL2 兼容的配置
- **启动脚本**: 添加便捷的启动/停止脚本

## Capabilities

### New Capabilities

- `docker-compose-config`: Docker Compose 配置文件改造，支持 WSL2 环境
- `docker-data-management`: 项目内 Docker 数据目录管理，使用绑定挂载
- `port-mapping`: 端口映射规范，宿主机端口统一为 3 开头
- `startup-scripts`: 便捷的 Docker 服务启动/停止脚本

### Modified Capabilities

无（不修改现有功能规范，仅调整部署配置）

## Impact

- **受影响文件**:
  - `mscoin-backend/docker-compose.yml` - Compose 配置
  - `mscoin-backend/.env` - 环境变量
  - `docker/` - 新建数据目录
  - `scripts/docker-*` - 新建启动脚本
- **基础设施**: MySQL、Redis、MongoDB、Etcd、Kafka 的容器配置
- **开发流程**: 本地启动后端服务的流程简化
