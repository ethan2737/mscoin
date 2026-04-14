# MSCOIN

MSCOIN 是一个加密货币交易平台，采用 Go 微服务后端 + Vue.js 前端架构。

## 项目结构

```
mscoin/
├── mscoin-backend/         # Go 微服务后端
│   ├── ucenter/            # 用户中心（注册、登录、会员、资产、提现）
│   ├── ucenter-api/        # 用户服务 API 网关
│   ├── market/             # 行情数据和汇率
│   ├── market-api/         # 行情服务 API（含 WebSocket）
│   ├── exchange/           # 订单撮合引擎
│   ├── exchange-api/       # 交易所服务 API
│   ├── swap/               # 永续合约服务（gRPC）
│   ├── swap-api/           # 永续合约 HTTP API
│   ├── jobcenter/          # 后台任务（K 线、比特币交易、汇率）
│   ├── mscoin-common/      # 公共工具和模型
│   └── grpc-common/        # gRPC 客户端存根
├── mscoin-frontend/        # Vue.js 前端
│   ├── src/
│   │   ├── pages-vue3/     # Vue 3 页面组件（活动版本）
│   │   ├── pages/          # Vue 2 页面组件（历史版本）
│   │   ├── main-vue3.js    # Vue 3 入口文件
│   │   ├── config/         # 路由、状态管理、API 配置
│   │   ├── components/     # 可复用组件
│   │   └── assets/         # 静态资源、语言包
│   └── index.html          # 前端入口（指向 main-vue3.js）
├── scripts/                # 启动脚本和工具库
│   └── lib/
├── docs/                   # 项目文档
└── dev-up.cmd              # 一键启动脚本
```

## 技术栈

**后端**
- Go 1.19+ 使用 go-zero 框架
- gRPC 用于服务间通信
- GORM (MySQL) + MongoDB
- Kafka 消息队列
- Redis 缓存
- Etcd 服务发现

**前端（Vue 3）**
- Vue 3.5 + Vite 5
- Vue Router 5 + Vuex 4
- Element Plus
- Axios + Socket.io/StompJS WebSocket 通信
- Vue I18n 11（中/英）

**前端（Vue 2 - 历史版本）**
- Vue 2.7 + Webpack 3
- iView 3.x UI 组件库
- vue-router 3 + vuex 3

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
- 自动应用本地认证、行情和交易验证基线
- 按依赖顺序启动 `ucenter`、`market`、`exchange`、`swap`、`jobcenter`、`ucenter-api`、`market-api`、`exchange-api`、`swap-api`、前端 Vite
- 对前端可访问性、本地登录链路、首页行情快照、交易页盘口、最新成交和当前委托链路执行 smoke checks
- 输出前端地址和本地测试账号；如需持久化启动日志可使用 `.\dev-up.cmd -WriteLogs`

### 后端

```bash
cd mscoin-backend

# 同步工作区
go work sync

# 运行服务（手工回退方式）
go run ucenter/main.go -f ucenter/etc/conf.yaml
go run market/main.go -f market/etc/conf.yaml
go run exchange/main.go -f exchange/etc/conf.yaml
go run swap/main.go -f swap/etc/conf.yaml
go run jobcenter/main.go -f jobcenter/etc/conf.yaml

# API 服务
go run ucenter-api/main.go -f ucenter-api/etc/conf.yaml
go run market-api/main.go -f market-api/etc/conf.yaml
go run exchange-api/main.go -f exchange-api/etc/conf.yaml
go run swap-api/main.go -f swap-api/etc/conf.yaml
```

### 前端

```bash
cd mscoin-frontend

# 安装依赖
pnpm install

# 开发服务器（Vite + Vue 3）
pnpm run dev

# 生产构建
pnpm run build

# 预览构建结果
pnpm run preview
```

**Vue 3 迁移说明**

当前前端使用 Vue 3 + Vite 作为活动版本：
- 入口文件：`src/main-vue3.js`
- 路由配置：`src/config/routes-vue3.js`
- 状态管理：`src/config/store-vue3.js`
- 页面组件：`src/pages-vue3/`

Vue 2.7 + Webpack 方案作为历史回退参考保留在 `src/pages/` 和相关配置中。

当前前端恢复基线和验收记录见：
- `openspec/changes/archive/2026-04-09-unify-frontend-migration-path/frontend-migration-baseline.md`
- `openspec/changes/archive/2026-04-09-frontend-migration-validation-and-cleanup/frontend-recovery-validation.md`

### 基础设施（手工回退方式）

```bash
cd mscoin-backend
docker-compose up -d
```

启动服务：MySQL (3309), Redis (6379), Etcd (2379), MongoDB (27018), Kafka (9092), Kafdrop (9000)

## 项目文档

详细文档请查看 `docs/` 目录，包含以下业务流程文档：

**核心业务**
- `用户登录注册流程梳理.md` - 用户认证完整流程
- `用户中心 UC 业务流程梳理.md` - 用户中心功能详解
- `充提币业务流程梳理.md` - 充值与提现业务流转

**交易业务**
- `币币交易页入口与模块梳理.md` - 币币交易页面结构
- `币币交易盘口业务梳理.md` - 盘口数据流转机制
- `永续合约业务流程梳理.md` - 合约交易完整流程

**其他业务**
- `C2C 快捷买卖流程梳理.md` - C2C 交易流程
- `OTC 业务流程梳理.md` - OTC 交易流程
- `首页业务流程梳理.md` - 首页功能模块
- `Nginx 业务流程梳理.md` - Nginx 配置与请求流转

**设计方案**
- `行情数据对接实现方案.md` - 行情数据接入方案
- `业务域现状梳理.md` - 整体业务架构概览
- `公益创新室.md` - 公益创新项目说明

## 本地交易页验证

默认启动完成后，可直接访问：

- 前端首页：`http://127.0.0.1:3000`
- 现货交易页：`http://127.0.0.1:3000/#/exchange/BTC_USDT`

本地验证账号：

- 手机号：`13800000000`
- 密码：`123456`

当前默认基线保证：

- 首页行情列表可展示
- 交易页 K 线可渲染
- 盘口与最新成交接口有本地可验证数据
- 当前委托列表可读取到最小验证数据

## 开发

```bash
# 提交代码
git add .
git commit -m "描述"
git push origin master
```

## License

MIT
