## Why

交易页当前只打通了币对列表和顶部摘要，核心交易体验没有跑通：K 线图不渲染、盘口与最新成交接口返回 404、深度图无法展示、委托区域缺少可验证的数据反馈，同时页面深色主题和默认白底组件混用，已经影响功能验证和业务演示。这个页面是“去交易”主链路的落点，必须尽快恢复到可展示、可下单、可验证的状态。

## What Changes

- 补齐交易页 K 线图初始化链路，确保图表实例能够在 Vue 3 页面正确挂载并拉取历史蜡烛数据。
- 补齐 `market-api` 的盘口、深度和最新成交接口，使交易页右侧盘口表格和深度图有真实数据来源。
- 打通交易页实时数据更新链路，确保币对切换后，K 线、盘口和最新成交能够同步刷新。
- 梳理交易页委托区域的数据依赖，保证当前委托、历史委托和撤单动作在本地联调环境中具备可验证反馈。
- 统一交易页深色主题样式，修复 Element Plus 默认白底组件与旧样式混用导致的视觉混乱和布局遮挡问题。

## Capabilities

### New Capabilities
- `exchange-trading-page-experience`: 定义交易页在 K 线、盘口、深度、成交、委托和主题布局上的完整可用性要求。

### Modified Capabilities

## Impact

- 前端页面：[Exchange.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/exchange/Exchange.vue)、[DepthGraph.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/exchange/DepthGraph.vue)、[bitrade.js](E:/Project/web3/mscoin/mscoin-frontend/src/assets/js/charting_library/datafeed/bitrade.js)、[exchange.css](E:/Project/web3/mscoin/mscoin-frontend/src/assets/css/exchange.css)
- 前端启动与运行时配置：[main-vue3.js](E:/Project/web3/mscoin/mscoin-frontend/src/main-vue3.js)、[runtime-vue3.js](E:/Project/web3/mscoin/mscoin-frontend/src/config/runtime-vue3.js)
- 行情 API：[routes_handler.go](E:/Project/web3/mscoin/mscoin-backend/market-api/internal/handler/routes_handler.go)、[market.go](E:/Project/web3/mscoin/mscoin-backend/market-api/internal/handler/market.go)、[market_logic.go](E:/Project/web3/mscoin/mscoin-backend/market-api/internal/logic/market_logic.go)
- 交易委托 API：[routes_handler.go](E:/Project/web3/mscoin/mscoin-backend/exchange-api/internal/handler/routes_handler.go)、[order.go](E:/Project/web3/mscoin/mscoin-backend/exchange-api/internal/handler/order.go)、[order_logic.go](E:/Project/web3/mscoin/mscoin-backend/exchange-api/internal/logic/order_logic.go)
- 本地联调与验证：可能涉及启动 smoke check 和最小数据基线补充
