# CLAUDE.md

本文档为 Claude Code (claude.ai/code) 在此代码仓库中工作提供指导。

无论是推理内容还是回复内容，都要使用中文回复。

## 项目概述

MSCOIN 是一个加密货币交易平台，采用 Go 微服务后端 + Vue.js 前端的架构。

**注意**: 本项目原为前后端分离的两个独立仓库，后合并为统一的 `mscoin` 仓库以便管理。
- 后端原始仓库：`mscoin-backend/`
- 前端原始仓库：`mscoin-frontend/`

## 架构

### 后端 (mscoin-backend/)

基于 go-zero 框架的 Go 微服务，使用 gRPC 进行服务间通信。服务模块：

- **ucenter** - 用户中心（注册、登录、会员、资产、提现）
- **ucenter-api** - 用户服务 HTTP API 网关
- **market** - 行情数据和汇率
- **market-api** - 行情服务 HTTP API（含 WebSocket 支持）
- **exchange** - 订单撮合引擎
- **exchange-api** - 交易所服务 HTTP API
- **jobcenter** - 后台任务（K 线、比特币交易、汇率）
- **mscoin-common** - 公共工具（加密、JWT、工具类、模型）
- **grpc-common** - 生成的 gRPC 客户端存根

每个服务遵循 go-zero 标准布局：
- `internal/config/` - 配置结构体
- `internal/handler/` - HTTP 处理器（API 服务）
- `internal/logic/` - 业务逻辑
- `internal/dao/` - 数据访问层
- `internal/repo/` - 仓库层
- `internal/model/` - 数据模型
- `internal/server/` - gRPC 服务器
- `internal/svc/` - 服务上下文
- `etc/conf.yaml` - 配置文件

### 前端 (mscoin-frontend/)

当前前端恢复基线：Vue 3 + Vite + `src/pages-vue3/`。

**基线规则**
- 活动入口：`mscoin-frontend/index.html`
- 活动启动文件：`mscoin-frontend/src/main-vue3.js`
- 活动路由/状态：`mscoin-frontend/src/config/routes-vue3.js` 和 `mscoin-frontend/src/config/store-vue3.js`
- 活动页面树：`mscoin-frontend/src/pages-vue3/`
- `6e05b70` 的 Vue 2.7 + Vite 方案只作为历史回退参考，不得再被接入当前 active startup/build path

**详细基线文档**
- `openspec/changes/archive/2026-04-09-unify-frontend-migration-path/frontend-migration-baseline.md`

核心目录：
- `src/pages/` - 页面组件
- `src/components/` - 可复用组件
- `src/config/` - 路由、状态管理、API 配置
- `src/assets/` - 静态资源、JS 工具、语言文件
- `build/` - Webpack 配置

## 命令

### 后端

```bash
# 默认推荐：从仓库根目录运行 .\dev-up.cmd

# 进入后端目录
cd mscoin-backend

# 同步工作区
go work sync

# 运行单个服务
go run ucenter/main.go -f ucenter/etc/conf.yaml
go run market/main.go -f market/etc/conf.yaml
go run exchange/main.go -f exchange/etc/conf.yaml
go run jobcenter/main.go -f jobcenter/etc/conf.yaml
# API 服务
go run ucenter-api/main.go -f ucenter-api/etc/conf.yaml
go run market-api/main.go -f market-api/etc/conf.yaml
go run exchange-api/main.go -f exchange-api/etc/conf.yaml

# 构建单个服务
go build -o <输出文件> <服务>/main.go
```

### 前端

```bash
# 默认推荐：从仓库根目录运行 .\dev-up.cmd

cd mscoin-frontend

# 安装依赖
npm install

# 开发服务器（端口 8080）
npm run dev

# 生产构建
npm run build

# 运行测试
npm run test        # 全部测试
npm run unit        # 单元测试 (Jest)
npm run e2e         # 端到端测试 (Nightwatch)

# 代码检查
npm run lint
```

### 基础设施

```bash
# 使用 Docker Compose 启动所有服务
cd mscoin-backend
docker-compose up -d

# 服务：MySQL (3309), Redis (6379), Etcd (2379), MongoDB (27018), Kafka (9092), Kafdrop (9000)
```

## 技术栈

**后端：**
- Go 1.19+ 使用 go-zero 框架
- gRPC 用于服务间通信
- GORM 作为 MySQL ORM
- MongoDB 存储时序数据（K 线、行情）
- Kafka 消息队列
- Redis 缓存
- Etcd 服务发现

**前端：**
- Vue 2.5 + Vue Router + Vuex
- iView 3.x UI 组件库
- Webpack 3 打包
- Socket.io/StompJS WebSocket 通信
- vue-i18n 国际化（中/英）
- vue-resource HTTP 请求

## 配置

- 后端服务使用 `etc/` 目录下的 YAML 配置文件
- 前端配置在 `config/dev.env.js` 和 `config/prod.env.js`
- Docker 环境变量在 `.env`

## 测试

- 前端使用 Jest 进行单元测试，Nightwatch 进行 E2E 测试
- 测试命令：`npm run unit` 单元测试，`npm run e2e` 端到端测试

## 流程梳理规则

当用户提出“梳理流程”“整理流程”“总结业务流程”“画流程图”“登录注册流程梳理”这类需求时，默认按以下规则输出，除非用户明确要求其他格式。

### 1. 信息来源

- 先阅读当前仓库中的前端页面、路由、store、API 调用代码。
- 再阅读对应后端 handler、logic、domain、RPC、缓存或数据库相关代码。
- 所有结论必须以“当前代码真实实现”为准，不得按理想方案推断。
- 如果前端展示与后端真实行为不一致，必须在文档中明确写出。

### 2. 文档结构

- 按业务主题分大章节。
- 如果一次梳理多个主流程，默认每个主流程单独成章。
- 例如登录注册场景，默认只分为“登录”“注册”两部分，不额外展开背景、风险、优化建议等章节，除非用户明确要求。
- 流程梳理类文档默认写到仓库根目录。
- 文件名默认直接使用流程名命名，例如：`用户登录注册流程梳理.md`、`充值流程梳理.md`。

### 3. 每个流程章节的固定写法

每个流程章节默认包含且只包含以下三部分：

1. `用户操作步骤`
2. `业务逻辑说明`
3. `流程图`

具体要求：

- `用户操作步骤`
  - 从用户视角按页面操作顺序逐步描述。
  - 写清入口、输入项、按钮、提示、跳转结果。

- `业务逻辑说明`
  - 按真实代码执行顺序描述前端校验、接口调用、后端处理、状态写入、页面跳转。
  - 语言要求简介、精炼。
  - 要覆盖关键判断条件，例如验证码是否通过、账号是否存在、密码是否正确、是否写入本地登录态。

- `流程图`
  - 必须使用 `mermaid`。
  - 必须尽量复刻代码中的动作、判断、分支，而不是只画产品层面的简化主干图。
  - 图中要体现成功、失败、回退、跳转等关键路径。

### 4. 语言风格

- 全文使用中文。
- 用简洁、概括性的表达，不写成冗长技术报告。
- 重点描述“用户如何操作”和“系统如何处理”。

### 5. 禁止项

- 不得把未来规划或理想设计写成当前现状。
- 不得补写代码中不存在的接口、校验或状态流转。
- 不要默认附带优化建议、问题清单、改造方案，除非用户明确要求。
- 不要输出缺少判断分支的过度简化流程图。

## GStack

**网页浏览**: 所有网页浏览任务必须使用 `/browse` skill (来自 gstack)，不得使用 `mcp__claude-in-chrome__*` 工具。

**可用 skills**:
- `/office-hours` - YC Office Hours 模式
- `/plan-ceo-review` - CEO/创始人模式计划审查
- `/plan-eng-review` - 工程经理模式计划审查
- `/plan-design-review` - 设计师视角计划审查
- `/design-consultation` - 设计咨询
- `/design-shotgun` - 生成多个 AI 设计变体
- `/design-html` - 生成生产就绪的 HTML/CSS
- `/review` - PR 审查
- `/ship` - 合并和部署工作流
- `/land-and-deploy` - 着陆和部署
- `/canary` - 部署后金丝雀监控
- `/benchmark` - 性能回归检测
- `/browse` - 快速无头浏览器 QA 测试
- `/connect-chrome` - 连接 Chrome 浏览器
- `/qa` - 系统性 QA 测试
- `/qa-only` - 仅报告 QA 测试
- `/design-review` - 设计视觉审查
- `/setup-browser-cookies` - 导入浏览器 Cookie
- `/setup-deploy` - 配置部署设置
- `/retro` - 每周工程回顾
- `/investigate` - 系统性调试
- `/document-release` - 发布后文档更新
- `/codex` - OpenAI Codex CLI 包装器
- `/cso` - 首席安全官模式
- `/autoplan` - 自动审查管道
- `/plan-devex-review` - 开发者体验计划审查
- `/devex-review` - 开发者体验审计
- `/careful` - 破坏性命令安全保护
- `/freeze` - 限制文件编辑范围
- `/guard` - 全面安全模式
- `/unfreeze` - 清除冻结边界
- `/gstack-upgrade` - 升级 gstack
- `/learn` - 管理项目学习
