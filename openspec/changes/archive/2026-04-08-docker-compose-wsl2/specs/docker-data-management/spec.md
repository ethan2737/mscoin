## ADDED Requirements

### Requirement: 项目内 Docker 数据目录

系统 SHALL 在项目根目录创建 `docker/` 目录，所有容器数据通过绑定挂载存储于此。

#### Scenario: 目录结构存在
- **WHEN** 运行启动脚本
- **THEN** 自动创建 `docker/{mysql,redis,mongo,etcd,kafka}/` 目录结构

#### Scenario: MySQL 数据挂载
- **WHEN** MySQL 容器启动
- **THEN** 数据目录挂载到 `./docker/mysql/data`

#### Scenario: Redis 数据挂载
- **WHEN** Redis 容器启动
- **THEN** 数据目录挂载到 `./docker/redis/data`

#### Scenario: MongoDB 数据挂载
- **WHEN** MongoDB 容器启动
- **THEN** 数据目录挂载到 `./docker/mongo/db`

#### Scenario: Etcd 数据挂载
- **WHEN** Etcd 容器启动
- **THEN** 数据目录挂载到 `./docker/etcd/data`

#### Scenario: Kafka 数据挂载
- **WHEN** Kafka 容器启动
- **THEN** 数据目录挂载到 `./docker/kafka/data`

---

### Requirement: 数据目录权限管理

系统 SHALL 确保 Docker 容器有权限读写挂载的数据目录。

#### Scenario: 首次启动权限
- **WHEN** 容器首次启动并写入数据
- **THEN** 无权限错误，容器正常写入

#### Scenario: 跨平台权限
- **WHEN** 在 WSL2 和 Windows 间切换使用
- **THEN** 数据目录权限正确，容器可读写

---

### Requirement: .gitignore 配置

系统 SHALL 将 Docker 数据目录添加到 .gitignore，避免提交二进制数据文件。

#### Scenario: Git 忽略数据
- **WHEN** 运行 `git status`
- **THEN** `docker/*/data/`、`docker/*/db/` 等目录不显示为未跟踪文件

#### Scenario: 保留目录结构
- **WHEN** 克隆仓库后首次启动
- **THEN** 启动脚本自动创建必要的目录结构
