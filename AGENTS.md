# AGENTS.md

本文档为 Codex 在此代码仓库中工作提供指导。

无论是推理内容还是回复内容，都要使用中文回复。

## 项目概述

MSCOIN 是一个加密货币交易平台，采用 Go 微服务后端加 Vue 前端的架构。

注意：
- 本项目原为前后端分离的两个独立仓库，后合并为统一的 `mscoin` 仓库以便管理。
- 后端原始目录为 `mscoin-backend/`。
- 前端原始目录为 `mscoin-frontend/`。

## 架构

### 后端 (`mscoin-backend/`)

后端基于 go-zero 架构，服务间通过 gRPC 通信。核心模块如下：

- `ucenter`：用户中心，负责注册、登录、会员、资产、提现等能力。
- `ucenter-api`：用户服务 HTTP API 网关。
- `market`：行情数据和汇率能力。
- `market-api`：行情服务 HTTP API，包含 WebSocket 支持。
- `exchange`：订单撮合引擎。
- `exchange-api`：交易服务 HTTP API。
- `jobcenter`：后台任务，如 K 线、比特币交易、汇率同步。
- `mscoin-common`：公共工具，如加密、JWT、工具类、公共模型。
- `grpc-common`：生成的 gRPC 客户端存根。

典型目录结构：

- `internal/config/`：配置结构体。
- `internal/handler/`：HTTP 处理器，仅 API 服务包含。
- `internal/logic/`：业务逻辑。
- `internal/dao/`：数据访问层。
- `internal/repo/`：仓库层。
- `internal/model/`：数据模型。
- `internal/server/`：gRPC 服务实现。
- `internal/svc/`：服务上下文。
- `etc/conf.yaml`：服务配置文件。

### 前端 (`mscoin-frontend/`)

当前前端有效基线为 Vue 3 + Vite，主页面树位于 `src/pages-vue3/`。

基线规则：

- 活动入口：`mscoin-frontend/index.html`
- 活动启动文件：`mscoin-frontend/src/main-vue3.js`
- 活动路由与状态：`mscoin-frontend/src/config/routes-vue3.js`、`mscoin-frontend/src/config/store-vue3.js`
- 活动页面树：`mscoin-frontend/src/pages-vue3/`
- 历史 `Vue 2.7 + Vite` 方案仅作为回退参考，不得重新接入当前启动和构建主链路。

详细基线文档：

- `openspec/changes/archive/2026-04-09-unify-frontend-migration-path/frontend-migration-baseline.md`

核心目录：

- `src/pages/`：页面组件。
- `src/components/`：可复用组件。
- `src/config/`：路由、状态管理、API 配置。
- `src/assets/`：静态资源、JS 工具、语言文件。
- `build/`：历史构建配置。

## 命令

### 默认本地启动方式

```powershell
.\dev-up.cmd
```

等价 PowerShell 命令：

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\dev-up.ps1
```

默认启动脚本会统一完成以下工作：

- 启动 Docker 基础设施依赖。
- 应用本地登录、行情和交易验证所需的最小数据基线。
- 按依赖顺序启动 `ucenter`、`market`、`exchange`、`jobcenter`、`ucenter-api`、`market-api`、`exchange-api` 和前端开发服务器。
- 执行前端可访问性、登录链路、行情快照、交易页盘口、最新成交和当前委托链路 smoke checks。

补充说明：

- 前端默认地址为 `http://127.0.0.1:3000`。
- 本地验证账号为 `13800000000 / 123456`。
- 默认不会在仓库内创建日志文件；如需持久化启动日志，使用 `.\dev-up.cmd -WriteLogs`。
- 交易页本地验证入口为 `http://127.0.0.1:3000/#/exchange/BTC_USDT`。

### 后端（手工排障/兜底）

```bash
cd mscoin-backend

go work sync

go run ucenter/main.go -f ucenter/etc/conf.yaml
go run market/main.go -f market/etc/conf.yaml
go run exchange/main.go -f exchange/etc/conf.yaml
go run jobcenter/main.go -f jobcenter/etc/conf.yaml

go run ucenter-api/main.go -f ucenter-api/etc/conf.yaml
go run market-api/main.go -f market-api/etc/conf.yaml
go run exchange-api/main.go -f exchange-api/etc/conf.yaml

go build -o <输出文件> <服务>/main.go
```

### 前端（手工排障/兜底）

```bash
cd mscoin-frontend

pnpm install
pnpm run dev
pnpm run build
pnpm run test
pnpm run unit
pnpm run e2e
pnpm run lint
```

### 基础设施（手工排障/兜底）

```bash
cd mscoin-backend
docker compose up -d
```

默认依赖服务包括：

- MySQL
- Redis
- Etcd
- MongoDB
- Kafka
- Kafdrop

## 技术栈

后端：

- Go 1.19+
- go-zero
- gRPC
- GORM
- MySQL
- MongoDB
- Kafka
- Redis
- Etcd

前端：

- Vue 3
- Vite
- Vue Router
- Vuex
- WebSocket
- i18n

## 配置

- 后端服务配置位于各服务 `etc/` 目录下的 YAML 文件。
- 前端环境配置位于 `config/` 目录。
- Docker 环境变量位于 `mscoin-backend/.env`。

## 测试

- 前端测试命令：`pnpm run unit`、`pnpm run e2e`。
- 代码检查命令：`pnpm run lint`。
- 启动链路验证优先使用 `.\dev-up.cmd` 自带的 smoke checks。
- 如需浏览器自动化验证，可使用 `agent-browser`，命令帮助执行 `agent-browser --help`。

## 流程梳理规则

当用户提出“梳理流程”“整理流程”“总结业务流程”“画流程图”“登录注册流程梳理”这类需求时，默认按以下规则输出，除非用户明确要求其他格式。

### 1. 信息来源

- 先读当前仓库中的前端页面、路由、store、API 调用代码。
- 再读对应后端 handler、logic、domain、RPC 或数据库相关代码。
- 结论必须以“当前代码真实实现”为准，不得按理想设计臆测。
- 如果存在前端文案与后端实现不一致，要在文档中直接指出。

### 2. 文档结构

- 按业务主题分大章节。
- 如果用户一次要求梳理多个主流程，默认每个主流程单独成章。
- 例如登录注册场景，默认只分为“登录”“注册”两部分，不额外扩展背景、风险、优化建议等章节，除非用户明确要求。
- 流程梳理类文档默认落在仓库根目录。
- 文件名默认直接使用流程名命名，例如：`用户登录注册流程梳理.md`、`提现流程梳理.md`。

### 3. 每个流程章节的固定写法

每个流程章节默认包含且只包含以下三部分：

1. `用户操作步骤`
2. `业务逻辑说明`
3. `流程图`

具体要求：

- `用户操作步骤`
  - 站在用户视角，按页面操作顺序逐步描述。
  - 要写清入口页面、按钮点击、表单填写、跳转结果。

- `业务逻辑说明`
  - 按代码执行顺序描述前端校验、API 请求、后端处理、状态落地、跳转逻辑。
  - 语言要简洁、精炼，避免空泛术语。
  - 要说明关键判断条件，例如是否通过验证码、是否已注册、密码是否正确、是否写入本地 session。

- `流程图`
  - 必须使用 `mermaid`。
  - 必须尽量复刻真实代码逻辑，而不是只画产品视角的简图。
  - 图中要体现动作节点、判断节点、成功分支、失败分支、跳转分支。
  - 如果某个流程明显分前端/API/RPC 三段，也可以在同一张图里按顺序体现，但不要脱离当前实现。

### 4. 语言风格

- 全文使用中文。
- 用简介、精炼的语言总结，不写成长篇分析报告。
- 重点是“用户怎么操作”与“代码怎么执行”，不是泛泛解释系统设计。

### 5. 禁止项

- 不要把“当前未实现但未来可能支持”的逻辑写成现状。
- 不要脱离代码自行补全不存在的接口或判断。
- 不要默认加入优化建议、问题清单、改造方案，除非用户明确要求。
- 不要把流程图画成只有主干路径、缺少关键判断的示意图。
