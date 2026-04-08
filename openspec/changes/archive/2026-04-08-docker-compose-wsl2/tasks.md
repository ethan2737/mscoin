## 1. 目录结构创建

- [x] 1.1 在项目根目录创建 `docker/` 目录
- [x] 1.2 创建 MySQL 数据目录：`docker/mysql/{data,conf,logs}`
- [x] 1.3 创建 Redis 数据目录：`docker/redis/{data,conf}`
- [x] 1.4 创建 MongoDB 数据目录：`docker/mongo/{db,log}`
- [x] 1.5 创建 Etcd 数据目录：`docker/etcd/data`
- [x] 1.6 创建 Kafka 数据目录：`docker/kafka/data`
- [x] 1.7 更新 `.gitignore`，添加 Docker 数据目录

## 2. Docker Compose 配置改造

- [x] 2.1 复制原 `docker-compose.yml` 为备份
- [x] 2.2 更新 `docker-compose.yml`，使用相对路径
- [x] 2.3 复制原 `.env` 为备份
- [x] 2.4 更新 `.env`，使用相对路径配置
- [x] 2.5 更新 `.env`，端口映射改为 3 开头

## 3. 端口映射配置

- [x] 3.1 配置 MySQL 端口：33306 → 3306
- [x] 3.2 配置 Redis 端口：33679 → 6379
- [x] 3.3 配置 MongoDB 端口：37018 → 27017
- [x] 3.4 配置 Etcd 端口：32379 → 2379, 32380 → 2380
- [x] 3.5 配置 Kafka 端口：39092 → 9092
- [x] 3.6 配置 Zookeeper 端口：32181 → 2181

## 4. 启动脚本创建

- [x] 4.1 创建 `docker-start.bat`（Windows）
- [x] 4.2 创建 `docker-stop.bat`（Windows）
- [x] 4.3 创建 `docker-start.sh`（WSL/Linux）
- [x] 4.4 创建 `docker-stop.sh`（WSL/Linux）
- [x] 4.5 添加执行权限：`chmod +x *.sh`

## 5. 后端服务配置更新

- [x] 5.1 更新 `ucenter/etc/conf.yaml` 中的基础设施端口
- [x] 5.2 更新 `market/etc/conf.yaml` 中的基础设施端口
- [x] 5.3 更新 `exchange/etc/conf.yaml` 中的基础设施端口
- [x] 5.4 更新 `jobcenter/etc/conf.yaml` 中的基础设施端口
- [x] 5.5 更新 API 服务的配置文件

## 6. 测试验证

- [x] 6.1 运行启动脚本，启动所有服务
- [x] 6.2 验证 MySQL 连接：`docker-compose ps`
- [x] 6.3 验证 Redis 连接
- [x] 6.4 验证 MongoDB 连接
- [x] 6.5 验证 Etcd 服务发现
- [x] 6.6 验证 Kafka 消息队列
- [x] 6.7 启动后端服务，验证基础设施连接

## 7. 文档和收尾

- [x] 7.1 更新 README，添加 Docker 启动说明
- [x] 7.2 编写 Docker 配置说明文档
- [x] 7.3 清理临时备份文件（可选）
