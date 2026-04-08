## ADDED Requirements

### Requirement: Docker 服务启动脚本

系统 SHALL 提供一键启动 Docker Compose 服务的脚本。

#### Scenario: Windows 批处理启动
- **WHEN** 在 `mscoin-backend/` 目录运行 `docker-start.bat`
- **THEN** 执行 `docker-compose up -d`，启动所有服务

#### Scenario: WSL/Linux shell 启动
- **WHEN** 在 `mscoin-backend/` 目录运行 `docker-start.sh`
- **THEN** 执行 `docker-compose up -d`，启动所有服务

#### Scenario: 启动状态检查
- **WHEN** 脚本执行完成后
- **THEN** 显示各容器启动状态

---

### Requirement: Docker 服务停止脚本

系统 SHALL 提供一键停止 Docker Compose 服务的脚本。

#### Scenario: Windows 批处理停止
- **WHEN** 在 `mscoin-backend/` 目录运行 `docker-stop.bat`
- **THEN** 执行 `docker-compose down`，停止所有服务

#### Scenario: WSL/Linux shell 停止
- **WHEN** 在 `mscoin-backend/` 目录运行 `docker-stop.sh`
- **THEN** 执行 `docker-compose down`，停止所有服务

---

### Requirement: 目录自动创建

系统 SHALL 在启动时自动检查并创建必要的数据目录。

#### Scenario: 目录不存在时
- **WHEN** 首次运行启动脚本
- **THEN** 自动创建 `docker/` 下所有必要的子目录

#### Scenario: 目录已存在时
- **WHEN** 目录已存在
- **THEN** 跳过创建，直接启动服务

---

### Requirement: 脚本帮助信息

系统 SHALL 提供脚本使用说明和帮助信息。

#### Scenario: 无参数运行
- **WHEN** 运行脚本无参数
- **THEN** 显示简要使用说明

#### Scenario: 帮助参数
- **WHEN** 运行脚本带 `-h` 或 `--help` 参数
- **THEN** 显示详细帮助信息，包括可用命令和示例
