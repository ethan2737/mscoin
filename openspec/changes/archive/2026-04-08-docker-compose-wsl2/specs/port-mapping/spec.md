## ADDED Requirements

### Requirement: 宿主机端口统一为 3 开头

系统 SHALL 将所有 Docker 服务的宿主机端口映射设置为 3 开头，避免与系统服务冲突。

#### Scenario: MySQL 端口映射
- **WHEN** MySQL 容器启动
- **THEN** 宿主机端口 33306 映射到容器端口 3306

#### Scenario: Redis 端口映射
- **WHEN** Redis 容器启动
- **THEN** 宿主机端口 33679 映射到容器端口 6379

#### Scenario: MongoDB 端口映射
- **WHEN** MongoDB 容器启动
- **THEN** 宿主机端口 37018 映射到容器端口 27017

#### Scenario: Etcd 端口映射
- **WHEN** Etcd 容器启动
- **THEN** 宿主机端口 32379 映射到容器端口 2379，32380 映射到 2380

#### Scenario: Kafka 端口映射
- **WHEN** Kafka 容器启动
- **THEN** 宿主机端口 39092 映射到容器端口 9092

#### Scenario: Zookeeper 端口映射
- **WHEN** Zookeeper 容器启动（Kafka 依赖）
- **THEN** 宿主机端口 32181 映射到容器端口 2181

---

### Requirement: 端口配置可修改

系统 SHALL 允许在 .env 文件中自定义各服务的端口映射。

#### Scenario: 修改端口
- **WHEN** 在 .env 中修改 `MYSQL_PORT=33307`
- **THEN** MySQL 容器使用新的端口 33307 映射

#### Scenario: 端口冲突检测
- **WHEN** 启动时端口已被占用
- **THEN** Docker Compose 报错提示端口冲突

---

### Requirement: 后端服务配置同步

系统 SHALL 确保后端服务的配置文件（etc/conf.yaml）使用正确的端口连接基础设施服务。

#### Scenario: 后端连接 MySQL
- **WHEN** 后端服务配置 DataSource
- **THEN** 使用 `127.0.0.1:33306` 连接 MySQL

#### Scenario: 后端连接 Redis
- **WHEN** 后端服务配置 CacheRedis
- **THEN** 使用 `127.0.0.1:33679` 连接 Redis

#### Scenario: 后端连接 MongoDB
- **WHEN** 后端服务配置 Mongo URL
- **THEN** 使用 `mongodb://localhost:37018` 连接 MongoDB

#### Scenario: 后端连接 Etcd
- **WHEN** 后端服务配置 Etcd 服务发现
- **THEN** 使用 `127.0.0.1:32379` 连接 Etcd

#### Scenario: 后端连接 Kafka
- **WHEN** 后端服务配置 Kafka 地址
- **THEN** 使用 `localhost:39092` 连接 Kafka
