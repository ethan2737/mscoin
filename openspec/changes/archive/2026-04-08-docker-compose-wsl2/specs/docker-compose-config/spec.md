## ADDED Requirements

### Requirement: Docker Compose 配置支持 WSL2 环境

系统 SHALL 使用相对路径配置 Docker Compose，使其可在 WSL2 + Docker Desktop 环境中运行。

#### Scenario: 使用相对路径
- **WHEN** Docker Compose 读取 .env 文件
- **THEN** 路径变量指向项目内的相对路径（如 `./docker/mysql/data`）

#### Scenario: 容器启动成功
- **WHEN** 执行 `docker-compose up -d`
- **THEN** 所有容器正常启动，无路径错误

#### Scenario: 数据持久化
- **WHEN** 容器停止并重新启动
- **THEN** 数据文件保留在 `docker/` 目录，数据不丢失

---

### Requirement: Docker Compose 配置使用 .env 变量

系统 SHALL 将所有可配置项提取到 .env 文件中，docker-compose.yml 仅引用变量。

#### Scenario: 环境变量替换
- **WHEN** docker-compose.yml 使用 `${VARIABLE_NAME}` 语法
- **THEN** Compose 正确读取 .env 中的值并替换

#### Scenario: 端口配置
- **WHEN** 在 .env 中修改端口变量
- **THEN** 容器使用新的端口映射

---

### Requirement: 多环境兼容

系统 SHALL 同时支持 Windows 批处理和 WSL/Linux shell 启动方式。

#### Scenario: Windows 批处理启动
- **WHEN** 在 Windows CMD 或 PowerShell 中运行 `docker-start.bat`
- **THEN** Docker Compose 服务正常启动

#### Scenario: WSL shell 启动
- **WHEN** 在 WSL terminal 中运行 `docker-start.sh`
- **THEN** Docker Compose 服务正常启动

#### Scenario: 路径格式兼容
- **WHEN** 在 WSL2 环境中访问项目目录
- **THEN** 路径正确解析，无格式错误
