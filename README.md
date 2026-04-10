# MSCOIN

MSCOIN 是一个加密货币交易平台，采用 Go 微服务后端 + Vue.js 前端架构。

## 项目结构

```
mscoin/
├── mscoin-backend/     # Go 微服务后端
│   ├── ucenter/        # 用户中心（注册、登录、会员、资产、提现）
│   ├── ucenter-api/    # 用户服务 API 网关
│   ├── market/         # 行情数据和汇率
│   ├── market-api/     # 行情服务 API（含 WebSocket）
│   ├── exchange/       # 订单撮合引擎
│   ├── exchange-api/   # 交易所服务 API
│   ├── jobcenter/      # 后台任务（K 线、比特币交易、汇率）
│   ├── mscoin-common/  # 公共工具和模型
│   └── grpc-common/    # gRPC 客户端存根
├── mscoin-frontend/    # Vue.js 前端
│   ├── src/pages/      # 页面组件
│   ├── src/components/ # 可复用组件
│   ├── src/config/     # 路由、状态管理、API 配置
│   └── src/assets/     # 静态资源
└── docs/               # 项目文档
```

## 技术栈

**后端**
- Go 1.19+ 使用 go-zero 框架
- gRPC 用于服务间通信
- GORM (MySQL) + MongoDB
- Kafka 消息队列
- Redis 缓存
- Etcd 服务发现

**前端**
- Vue 3 + Vite
- Vue Router 5 + Vuex 4
- Element Plus
- Socket.io/StompJS WebSocket 通信

## 快速开始

### 默认本地启动方式

在仓库根目录执行：

```powershell
.\dev-up.cmd
```

等价命令：

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\dev-up.ps1
```

这个命令会：

- 启动 Docker 基础设施依赖
- 自动应用本地认证与最小行情基线
- 按依赖顺序启动 `ucenter`、`market`、`jobcenter`、`ucenter-api`、`market-api`、前端 Vite
- 对前端可访问性、本地登录链路和行情快照链路执行 smoke checks
- 输出前端地址和本地测试账号；如需持久化启动日志可使用 `.\dev-up.cmd -WriteLogs`

详细说明见 [docs/local-dev-startup.md](./docs/local-dev-startup.md)。

### 后端

```bash
cd mscoin-backend

# 同步工作区
go work sync

# 运行服务（手工回退方式）
go run ucenter/main.go -f ucenter/etc/conf.yaml
go run market/main.go -f market/etc/conf.yaml
go run exchange/main.go -f exchange/etc/conf.yaml
```

### 前端

```bash
cd mscoin-frontend

# 安装依赖
pnpm install

# 开发服务器（手工回退方式）
pnpm run dev

# 生产构建
pnpm run build
```

当前前端恢复基线和验收记录见：

- `openspec/changes/archive/2026-04-09-unify-frontend-migration-path/frontend-migration-baseline.md`
- `openspec/changes/archive/2026-04-09-frontend-migration-validation-and-cleanup/frontend-recovery-validation.md`

### 基础设施（手工回退方式）

```bash
cd mscoin-backend
docker-compose up -d
```

启动服务：MySQL (3309), Redis (6379), Etcd (2379), MongoDB (27018), Kafka (9092)

## 项目文档

详细文档请查看 `docs/` 目录。

## 开发

```bash
# 提交代码
git add .
git commit -m "描述"
git push origin master
```

## License

MIT
