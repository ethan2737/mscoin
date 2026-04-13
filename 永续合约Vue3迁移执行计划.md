# 永续合约 Vue3 迁移执行计划

> 目标：按可执行工作包推进永续合约模块修复，只保留 Vue3 主链路，清除 Vue2 永续遗留逻辑，打通真实前后端端到端流程。

## 执行总原则

1. 只以 Vue3 页面树为准
2. 先修代理和契约，再删遗留代码
3. 每个工作包都要求可验证
4. 不允许依赖永续 mock 证明“已跑通”

## 工作包 A：冻结 Vue3 主链路契约

### 目标

梳理 Vue3 永续页面实际依赖的请求、字段、响应结构，形成唯一契约基线。

### 涉及文件

- `mscoin-frontend/src/pages-vue3/swapindex/Swapindex.vue`
- `mscoin-frontend/src/pages-vue3/swapindex/Kline.vue`
- `mscoin-frontend/src/pages-vue3/uc/ContractMoneyIndex.vue`
- `mscoin-frontend/src/pages-vue3/uc/ContractRecord.vue`
- `mscoin-frontend/src/pages-vue3/uc/contract/EntrustCurrent.vue`
- `mscoin-frontend/src/pages-vue3/uc/contract/EntrustHistory.vue`
- `mscoin-frontend/src/config/routes-vue3.js`
- `mscoin-frontend/src/config/api.js`

### 执行步骤

1. 输出 Vue3 永续页面与接口映射表
2. 标记每个页面的必需字段
3. 标记当前仍在使用的 `/swap/*`、`/uc/contract/*`、`/uc/contract-wallet*` 路径
4. 标记 Vue2 永续页面和旧路由入口
5. 生成一份“Vue3 契约基线表”

### 交付物

- Vue3 永续页面接口映射文档
- Vue2 永续遗留清单

### 验证方式

人工核对以下页面的请求是否全部有明确来源：

- `/swapindex/:pair`
- `/uc/contract-money`
- `/uc/contract/entrust/current`
- `/uc/contract/entrust/history`
- `/uc/contract-record`

---

## 工作包 B：修复开发代理与 mock

### 目标

让 Vue3 永续页面在开发环境真实命中 `swap-api`，不再被旧代理或 mock 掩盖。

### 涉及文件

- `mscoin-frontend/vite.config.mjs`
- `mscoin-frontend/dev/localAcceptanceMocks.mjs`

### 执行步骤

1. 在 `vite.config.mjs` 中为以下路径加精准代理到 `http://127.0.0.1:8086`
   - `/swap`
   - `/uc/contract`
   - `/uc/contract-wallet`
   - `/uc/asset/contract-transaction/all`
2. 保持其他 `/uc` 路径继续走 `http://127.0.0.1:8888`
3. 将永续相关 mock 改为默认关闭
4. 增加显式环境变量开关控制永续 mock
5. 删除与真实接口冲突的永续 mock 返回结构

### 交付物

- 修正后的 Vite 代理规则
- 默认关闭的永续 mock 策略

### 验证命令

```powershell
cd E:\Project\web3\mscoin\mscoin-frontend
pnpm run dev
```

### 验证标准

使用浏览器或 `Invoke-RestMethod` 验证：

```powershell
Invoke-RestMethod -Uri 'http://127.0.0.1:3000/swap/symbol-thumb' -Method Get
Invoke-RestMethod -Uri 'http://127.0.0.1:3000/uc/contract/current' -Method Post -ContentType 'application/json' -Body '{"memberId":1001}'
```

要求：

- 请求实际命中 `8086`
- 不再返回本地 mock 结构
- 协议错误能真实暴露

---

## 工作包 C：统一 `swap-api` 与 Vue3 页面契约

### 目标

让 `swap-api` 的请求字段、响应结构、错误返回与 Vue3 页面真实消费方式对齐。

### 涉及文件

- `mscoin-backend/swap-api/internal/types/types.go`
- `mscoin-backend/swap-api/internal/handler/contract_wallet_handler.go`
- `mscoin-backend/swap-api/internal/handler/order_entrust_handler.go`
- `mscoin-backend/swap-api/internal/handler/uc_contract_handler.go`
- `mscoin-backend/swap-api/internal/logic/contract_wallet_logic.go`
- `mscoin-backend/swap-api/internal/logic/order_entrust_logic.go`
- `mscoin-backend/swap-api/internal/logic/uc_contract_logic.go`

### 执行步骤

1. 统一钱包划转字段
   - 选定最终主字段 `type` 或 `direction`
   - 前后端统一
2. 统一钱包接口响应结构
3. 统一当前委托、历史委托接口的分页结构
4. 统一持仓接口字段命名
5. 统一合约币种列表字段，保证 Vue3 页面可直接展示
6. 补充 DTO，隔离数据库模型与前端展示模型

### 重点问题

- `ContractMoneyIndex.vue` 当前与后端划转字段不一致
- `EntrustCurrent.vue` 和 `EntrustHistory.vue` 期待的是分页 `content`
- 当前后端返回 `data.list`，页面无法直接消费

### 交付物

- 对齐后的请求 DTO/响应 DTO
- 与 Vue3 页面一致的 HTTP 响应结构

### 验证命令

```powershell
cd E:\Project\web3\mscoin\mscoin-backend\swap-api
go test ./...
```

再执行：

```powershell
Invoke-RestMethod -Uri 'http://127.0.0.1:8086/uc/contract/current' -Method Post -ContentType 'application/json' -Body '{"memberId":1001}'
Invoke-RestMethod -Uri 'http://127.0.0.1:8086/uc/contract-wallet' -Method Post -ContentType 'application/json' -Body '{"memberId":1001,"unit":"USDT"}'
Invoke-RestMethod -Uri 'http://127.0.0.1:8086/uc/contract-wallet/transfer' -Method Post -ContentType 'application/json' -Body '{"memberId":1001,"unit":"USDT","type":1,"amount":10}'
```

### 验证标准

- 当前委托、历史委托、钱包、持仓接口均返回 Vue3 可直接消费的数据
- 划转接口字段不再报 `field type is not set`

---

## 工作包 D：修复真实业务链路

### 目标

确保真实业务流程可以在数据库、Redis、Mongo、jobcenter 环境下跑通。

### 涉及文件

- `mscoin-backend/jobcenter/internal/logic/swap_kline.go`
- `mscoin-backend/jobcenter/internal/logic/swap_liquidation.go`
- `mscoin-backend/swap-api/internal/logic/order_entrust_logic.go`
- `mscoin-backend/swap-api/internal/logic/contract_wallet_logic.go`

### 执行步骤

1. 对齐 K 线写缓存 key 与强平读缓存 key
2. 校正强平价格读取逻辑
3. 校正钱包划转与余额更新逻辑
4. 校正委托状态流转、撤单、闪电平仓链路
5. 补充空钱包、空持仓、空订单保护逻辑
6. 校验分页参数、默认值、状态过滤逻辑

### 交付物

- 可实际读取价格并进行强平判断的任务链路
- 可真实执行的钱包、委托、持仓业务链路

### 验证命令

```powershell
cd E:\Project\web3\mscoin\mscoin-backend\jobcenter
go test ./...
```

手工验证：

```powershell
Invoke-RestMethod -Uri 'http://127.0.0.1:8086/swap/symbol-thumb' -Method Get
Invoke-RestMethod -Uri 'http://127.0.0.1:8086/swap/order/position' -Method Post -ContentType 'application/json' -Body '{"memberId":1001}'
Invoke-RestMethod -Uri 'http://127.0.0.1:8086/swap/order-entrust/add' -Method Post -ContentType 'application/json' -Body '{"memberId":1001,"contractCoinId":1,"entrustType":1,"type":1,"direction":1,"leverage":10,"price":1000,"amount":0.1}'
```

### 验证标准

- 强平逻辑不再因价格读取失败而短路
- 下单、查持仓、查委托、划转至少能跑通一条真实链路

---

## 工作包 E：清理 Vue2 永续遗留逻辑

### 目标

去掉仓库中会误导后续开发的 Vue2 永续页面、旧路由和相关旧逻辑。

### 涉及文件

- `mscoin-frontend/src/pages/swapindex/*`
- `mscoin-frontend/src/config/routes.js`
- 仅服务于 Vue2 永续的旧组件、工具和接口定义

### 执行步骤

1. 搜索 Vue2 永续页面引用
2. 先断开引用，再删除文件
3. 从 `routes.js` 中移除永续相关旧路由
4. 删除只服务于 Vue2 永续的无用代码
5. 补一份清理说明

### 交付物

- 被剥离的 Vue2 永续页面与路由
- 仓库内唯一的永续前端实现为 Vue3

### 验证命令

```powershell
cd E:\Project\web3\mscoin
git grep "pages/swapindex"
git grep "/swapindex" mscoin-frontend/src/config/routes.js
cd E:\Project\web3\mscoin\mscoin-frontend
pnpm run build
```

### 验证标准

- Vue2 永续页面不再被任何主链路引用
- 构建仍然通过

---

## 工作包 F：端到端测试与 smoke

### 目标

建立真实后端场景下的永续端到端验证，不再依赖 mock 和人工猜测。

### 涉及文件

- `mscoin-backend/swap-api/internal/logic/order_entrust_logic_test.go`
- `mscoin-backend/swap-api/internal/logic/position_logic_test.go`
- `mscoin-backend/swap-api/internal/logic/risk_control_logic_test.go`
- 新增前端 E2E 脚本文件
- 必要时更新 `scripts/dev-up.ps1` 与 `scripts/lib/DevStartup.psm1`

### 执行步骤

1. 替换“算术自检型”测试为真实逻辑测试
2. 补 jobcenter 中 key 对齐相关测试
3. 增加 Vue3 永续 E2E 用例
4. 如有必要，将 `swap` 与 `swap-api` 纳入本地启动脚本
5. 为永续业务增加独立 smoke 检查

### 端到端最小覆盖场景

1. 启动基础设施与后端服务
2. 打开 `/swapindex/1`
3. 查询合约信息、合约列表
4. 查询当前委托、历史委托、钱包
5. 执行一次钱包划转
6. 验证页面数据更新

### 验证命令

```powershell
cd E:\Project\web3\mscoin\mscoin-backend\swap-api
go test ./...

cd E:\Project\web3\mscoin\mscoin-backend\jobcenter
go test ./...

cd E:\Project\web3\mscoin\mscoin-frontend
pnpm run build
pnpm run test
```

### 验证标准

- 永续主流程在真实后端下可执行
- smoke 不依赖 mock
- 测试能复现并拦截本次 review 提到的关键问题

---

## 推荐执行顺序

1. 工作包 A：冻结契约
2. 工作包 B：修代理与 mock
3. 工作包 C：统一 `swap-api` 契约
4. 工作包 D：修真实业务链路
5. 工作包 E：删 Vue2 永续遗留
6. 工作包 F：补 E2E 与 smoke

## 里程碑定义

### 里程碑 M1：开发链路真实可见

- 3000 下 Vue3 永续请求真实命中 8086
- 永续 mock 默认关闭

### 里程碑 M2：页面与接口契约统一

- 钱包、委托、持仓、币种列表接口与 Vue3 页面对齐

### 里程碑 M3：业务可跑通

- 至少一条真实下单/查询/划转链路可用
- 强平逻辑可读到价格

### 里程碑 M4：完成迁移

- Vue2 永续页面从主仓库逻辑上退出
- 真实 E2E 与 smoke 覆盖 Vue3 永续主流程

## 建议的提交节奏

1. `chore: route vue3 swap traffic to swap-api`
2. `refactor: align swap-api dto with vue3 contract pages`
3. `fix: repair swap kline and liquidation cache contract`
4. `refactor: remove legacy vue2 swap pages`
5. `test: add vue3 perpetual swap e2e coverage`

## 下一步

如果继续执行，建议先从工作包 B 开始，因为它能最快暴露真实问题。
如果你愿意，我下一步可以直接开始实施工作包 B 和 C。
