# repair-exchange-page-data-and-ui

## Summary

交易页主链路已经恢复到可联调、可验证状态。Vue 3 交易页现在能够渲染 TradingView K 线，读取盘口、完整盘口和最新成交数据，并在本地验证基线下展示当前委托。页面的 Element Plus 白底表格和输入框也已统一到深色主题。

## Backend Recovery

### market-api 接口补齐

本次补齐了交易页依赖的行情接口：

- `POST /market/exchange-plate-mini`
- `POST /market/exchange-plate-full`
- `POST /market/latest-trade`

实现方式：

- 在 `routes_handler.go` 中注册兼容路由
- 在 `market.go` 中使用 `httpx.Parse` 解析 JSON body，修复之前 `ParseForm` 导致 `symbol` 丢失的问题
- 在 `market_logic.go` 中优先读取实时缓存，缓存为空时回退到本地 `exchange.exchange_order` 快照
- 在 `processor` 中补充盘口和最新成交缓存结构，并继续通过 websocket 推送交易页所需主题

### 本地交易联调基线

新增 `scripts/sql/local-exchange-baseline.sql`，为本地 `exchange` 数据库提供最小可验证订单集：

- 每个本地展示交易对包含买盘、卖盘、已成交和已撤销订单
- 当前委托和历史委托都能读取到真实记录
- 撤单接口可以在本地验证后再通过重新应用 baseline 恢复环境

`dev-up` 也纳入了：

- `exchange`
- `exchange-api`
- 交易页盘口、最新成交和当前委托 smoke checks

## Frontend Recovery

### K 线与实时链路

Vue 3 交易页恢复了旧版 TradingView 初始化逻辑，并继续复用现有 `bitrade.js` datafeed：

- 初次进入交易页时创建 widget
- 切换币对时销毁并重建 widget
- 语言切换时重新初始化 K 线
- websocket 重连时同步更新顶部摘要、盘口和最新成交

### 委托区调用修复

交易页前端统一改为：

- 通过共享 runtime helper 注入 `x-auth-token`
- 兼容 `response.data.code/data` 结构
- 当前委托、历史委托和撤单都走当前后端真实路由

### 深色主题修复

样式修复策略是收口到交易页组件内部，而不是继续堆叠历史 iView 规则：

- 在 `.exchange` 根节点显式设置 Element Plus 主题变量
- 对 `el-table`、`el-input`、空态和滚动容器做深色覆盖
- 修正 `scoped + :deep(...)` 选择器作用域，确保覆盖真实命中运行态节点

## Validation

### 自动化验证

已执行：

- `go test ./internal/config ./internal/database ./internal/logic ./internal/processor ./internal/model ./internal/handler ./internal/repo` in `mscoin-backend/market-api`
- `Invoke-Pester E:/Project/web3/mscoin/scripts/tests/DevStartup.Tests.ps1`
- `pnpm run build` in `mscoin-frontend`
- `cmd /c dev-up.cmd`

### 浏览器验证

通过 Chrome DevTools Protocol 对交易页做了真实页面校验：

- `K 线容器子节点数 = 1`
- `当前委托行数 = 2`
- `最新成交行数 = 2`
- `盘口行数 = 22`
- `左侧币对表背景 = rgb(25, 35, 48)`
- `盘口表背景 = rgb(25, 35, 48)`
- `委托表背景 = rgb(25, 35, 48)`
- `输入框背景 = rgb(39, 49, 62)`

这说明：

- K 线已渲染
- 盘口和最新成交已回填
- 当前委托可读
- 白底 Element Plus 组件已切回深色主题

## Files Changed

- `mscoin-backend/market-api/etc/conf.yaml`
- `mscoin-backend/market-api/internal/config/config.go`
- `mscoin-backend/market-api/internal/config/config_test.go`
- `mscoin-backend/market-api/internal/database/kafka.go`
- `mscoin-backend/market-api/internal/database/mysql.go`
- `mscoin-backend/market-api/internal/handler/market.go`
- `mscoin-backend/market-api/internal/handler/routes_handler.go`
- `mscoin-backend/market-api/internal/logic/market_logic.go`
- `mscoin-backend/market-api/internal/logic/market_logic_test.go`
- `mscoin-backend/market-api/internal/model/trade.go`
- `mscoin-backend/market-api/internal/processor/processor.go`
- `mscoin-backend/market-api/internal/processor/wshandler.go`
- `mscoin-backend/market-api/internal/repo/exchange_snapshot.go`
- `mscoin-backend/market-api/internal/svc/service_context.go`
- `mscoin-backend/market-api/internal/types/types.go`
- `mscoin-frontend/src/pages-vue3/exchange/Exchange.vue`
- `scripts/lib/DevStartup.psm1`
- `scripts/sql/local-exchange-baseline.sql`
- `scripts/tests/DevStartup.Tests.ps1`
- `README.md`
- `AGENTS.md`

## Result

本次 change 已达到“交易页主链路可联调、可浏览器验证、可重复启动验证”的目标，可以进入 archive 阶段。
