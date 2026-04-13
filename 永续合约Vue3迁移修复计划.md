# 永续合约 Vue3 迁移修复计划

> 目标：以 Vue3 页面树为唯一有效前端实现，完成永续合约前后端接口对齐、清理 Vue2 遗留逻辑与本地 mock 兜底，最终跑通真实端到端业务流程。

## 1. 背景与范围

本计划只以以下 Vue3 页面为主链路，不再以 Vue2 页面作为业务依据：

- `mscoin-frontend/src/pages-vue3/swapindex/Swapindex.vue`
- `mscoin-frontend/src/pages-vue3/swapindex/Kline.vue`
- `mscoin-frontend/src/pages-vue3/uc/ContractMoneyIndex.vue`
- `mscoin-frontend/src/pages-vue3/uc/contract/EntrustCurrent.vue`
- `mscoin-frontend/src/pages-vue3/uc/contract/EntrustHistory.vue`

本次修复覆盖四类问题：

1. Vue3 页面请求没有真正接入 `swap-api`
2. 前后端接口字段和响应结构不一致
3. 本地 mock 掩盖真实后端问题
4. Vue2 遗留页面、路由、接口逻辑继续干扰维护

不在本次计划内的内容：

- 重做永续业务产品规则
- 重构现货、OTC、众筹等非永续模块
- 重新设计整站路由体系

## 2. 现状结论

### 2.1 已确认问题

1. `vite.config.mjs` 当前仅把 `/swap` 代理到 `8086`，`/uc` 仍代理到 `8888`，导致 Vue3 用户中心永续页没有接入新服务。
2. `localAcceptanceMocks.mjs` 中仍拦截永续相关路径，开发环境会伪造成功结果。
3. Vue3 资产页提交 `direction`，后端 `swap-api` 要求 `type`，划转真实请求直接失败。
4. Vue3 当前委托/历史委托页消费的是分页结构 `content/totalElements`，而当前 `swap-api` 返回的是 `data.list`。
5. 强平扫描读取 Redis key 与 K 线任务写入 key 不一致，强平逻辑实际失效。
6. Vue2 页面和旧接口定义仍留在仓库中，容易误导后续开发继续补旧链路。

### 2.2 迁移原则

1. 前端只保留 Vue3 主链路
2. 后端接口以 Vue3 页面真实使用方式为准
3. 开发环境优先直连真实后端，mock 只用于显式验收模式
4. 每个阶段都要有可执行验证
5. 涉及资金、持仓、委托的链路必须补集成测试

## 3. 目标架构

### 3.1 前端目标

- Vue3 页面只调用一套真实永续接口
- 开发环境下：
  - `/swap/*` -> `swap-api`
  - `/uc/contract/*` -> `swap-api`
  - `/uc/contract-wallet*` -> `swap-api`
  - 非永续 `uc` 路径仍走 `ucenter-api`
- 永续相关 mock 默认关闭，只有显式开关时才启用

### 3.2 后端目标

- `swap-api` 统一提供 Vue3 永续页面需要的 HTTP 接口
- 返回结构与 Vue3 页面消费模型一致
- 强平、K 线、钱包划转、委托查询具备真实可验证能力
- 对用户身份、资金变更、状态流转增加必要校验

## 4. 分阶段修复方案

### 阶段一：冻结 Vue3 主链路契约

目标：先锁定“哪些 Vue3 页面是有效页面，哪些接口是必须支持的真实契约”。

产出：

- 一份 Vue3 永续页面与接口映射表
- 一份 Vue2 遗留页面清单
- 一份需要保留的后端接口清单

需要梳理的文件：

- `mscoin-frontend/src/pages-vue3/swapindex/Swapindex.vue`
- `mscoin-frontend/src/pages-vue3/swapindex/Kline.vue`
- `mscoin-frontend/src/pages-vue3/uc/ContractMoneyIndex.vue`
- `mscoin-frontend/src/pages-vue3/uc/contract/EntrustCurrent.vue`
- `mscoin-frontend/src/pages-vue3/uc/contract/EntrustHistory.vue`
- `mscoin-frontend/src/config/routes-vue3.js`
- `mscoin-frontend/src/config/api.js`

完成标准：

- 明确每个 Vue3 页面依赖的真实请求路径、请求字段、响应字段
- 标记哪些字段是页面渲染必需字段
- 标记哪些旧接口只是兼容残留

### 阶段二：修正开发代理与 mock 策略

目标：让 Vue3 永续页面在本地开发时优先走真实后端，而不是被旧代理或 mock 掩盖。

要改的文件：

- `mscoin-frontend/vite.config.mjs`
- `mscoin-frontend/dev/localAcceptanceMocks.mjs`

修复内容：

1. 为 `/uc/contract/*`、`/uc/contract-wallet*`、`/uc/asset/contract-transaction/all` 增加到 `swap-api` 的精准代理规则
2. 保留其余 `uc` 路径到 `8888`
3. 将永续相关 mock 改为默认关闭
4. 增加显式环境变量开关，例如：
   - `MSCOIN_ENABLE_SWAP_MOCK=true` 时才启用永续 mock
5. 移除与真实后端契约冲突的永续 mock 返回结构

完成标准：

- 通过 `http://127.0.0.1:3000` 访问 Vue3 永续页面时，请求实际到达 `8086`
- 无 token、字段缺失、协议不一致等问题不再被本地 mock 吞掉

### 阶段三：统一 `swap-api` 与 Vue3 页面的接口契约

目标：按 Vue3 页面实际消费方式统一请求字段、响应结构和错误返回。

要改的文件：

- `mscoin-backend/swap-api/internal/types/types.go`
- `mscoin-backend/swap-api/internal/handler/*.go`
- `mscoin-backend/swap-api/internal/logic/*.go`

重点修复项：

1. 钱包划转接口统一字段
   - 前后端统一为 `type` 或统一为 `direction`
   - 只能保留一种主字段
   - 如需兼容旧字段，只做过渡兼容，不作为最终契约
2. 当前委托、历史委托、持仓、钱包接口响应结构统一
   - 与 Vue3 页面使用的字段完全对齐
   - 必要时在后端做 ViewModel 转换，不把数据库模型直接透传到前端
3. 合约币种列表、当前杠杆、账单列表返回结构统一
4. 错误码、错误消息统一使用当前站点已有 `common.NewResult()` 风格

推荐做法：

- 在 `swap-api/internal/types/` 中新增面向页面的响应 DTO
- 在 `logic` 层做字段适配，不让 `handler` 和页面互相硬编码数据库字段

完成标准：

- Vue3 交易页、资产页、当前委托页、历史委托页不再依赖“猜字段”
- 真实请求可直接驱动页面渲染

### 阶段四：补齐真实业务链路

目标：确保 Vue3 永续的关键业务流程在真实后端下可跑通。

要覆盖的链路：

1. 合约列表与合约详情查询
2. 当前委托查询
3. 历史委托查询
4. 当前持仓查询
5. 合约钱包查询
6. 钱包划转
7. 下单
8. 撤单
9. 闪电平仓

重点修复项：

1. 强平价格缓存 key 对齐
   - `jobcenter/internal/logic/swap_kline.go`
   - `jobcenter/internal/logic/swap_liquidation.go`
2. 钱包与订单状态流转一致性
3. 缺省钱包、缺省持仓、空列表场景处理
4. 分页参数和默认值与 Vue3 页面保持一致

完成标准：

- 关键链路在真实数据库、Redis、Mongo、服务进程下可跑通
- 不依赖 `localAcceptanceMocks`

### 阶段五：清理 Vue2 遗留逻辑

目标：消除“仓库里还留着 Vue2 页面，后续又继续补旧代码”的维护风险。

清理范围：

- `mscoin-frontend/src/pages/swapindex/*`
- `mscoin-frontend/src/config/routes.js` 中与永续相关的旧路由
- 只服务于 Vue2 永续页面的组件、接口常量、工具函数

清理策略：

1. 先确认 `main-vue3.js` 和 `routes-vue3.js` 已完全覆盖线上入口
2. 搜索 Vue2 永续页面是否仍被任何文件引用
3. 将 Vue2 永续页面从构建主链路剥离
4. 删除未引用的 Vue2 永续页面与相关旧逻辑
5. 对删除动作补一份简短迁移说明

注意：

- 只删除永续相关的 Vue2 遗留逻辑
- 不扩大到无关 Vue2 历史页面，避免一次性清理过大

完成标准：

- 仓库内不再存在“永续业务有两套页面实现”的状态
- 新人阅读代码时不会再走错到 Vue2 永续页面

### 阶段六：建立端到端验证闭环

目标：把“能编译”升级为“业务真实可用”。

测试分层：

1. 后端单元测试
2. 后端集成测试
3. 前端接口联调测试
4. 浏览器端到端测试

建议补充的测试：

- `mscoin-backend/swap-api/internal/logic/`
  - 钱包划转测试
  - 委托列表转换测试
  - 持仓返回测试
  - 强平条件判断测试
- `mscoin-backend/jobcenter/internal/logic/`
  - K 线缓存写入 key 测试
  - 强平读取 key 测试
- 前端 E2E
  - 登录后进入 `/swapindex/1`
  - 查看合约信息与盘口
  - 打开当前委托/历史委托
  - 进入 `/uc/contract-money`
  - 发起一次钱包划转
  - 验证页面展示更新

建议执行命令：

```powershell
cd E:\Project\web3\mscoin\mscoin-backend\swap-api
go test ./...

cd E:\Project\web3\mscoin\mscoin-backend\jobcenter
go test ./...

cd E:\Project\web3\mscoin\mscoin-frontend
pnpm run build
pnpm run test
```

如果仓库现有测试框架不足，需要补一个最小可执行的浏览器验收脚本。

## 5. 执行顺序建议

建议按下面顺序执行，避免互相覆盖：

1. 先修代理与 mock
2. 再统一后端契约
3. 再修真实业务链路
4. 再删除 Vue2 永续遗留代码
5. 最后补齐 E2E 验证

原因：

- 如果先删 Vue2，不先修 Vue3 真实链路，会让验收更难
- 如果不先关掉 mock，很多契约问题不会暴露
- 如果不先统一接口结构，页面修修补补会越来越乱

## 6. 验收标准

### 6.1 功能验收

- Vue3 永续交易页可通过真实后端加载合约信息、盘口、成交、持仓、委托
- Vue3 合约资产页可通过真实后端查询资产并完成划转
- Vue3 当前委托页、历史委托页可查询真实数据
- 下单、撤单、闪电平仓至少跑通一条真实链路
- 强平任务可以读到价格并进入判断逻辑

### 6.2 代码验收

- Vue2 永续页面和路由不再出现在主链路中
- 永续相关 mock 默认关闭
- `vite.config.mjs` 代理规则按业务边界拆清楚
- `swap-api` 的 DTO 和 Vue3 页面字段一一对应

### 6.3 测试验收

- `swap-api` 新增或修正的逻辑有单元测试
- `jobcenter` 强平/K 线 key 对齐有测试
- 至少有一套 Vue3 永续端到端验收脚本
- 本地启动后，不依赖 mock 即可完成永续业务 smoke

## 7. 风险与应对

### 风险 1：旧接口兼容期过长

影响：

- 后端继续维护双协议，复杂度持续上升

应对：

- 在计划执行时明确“最终唯一契约”
- 兼容字段只保留一个过渡版本

### 风险 2：Vue3 页面字段依赖分散，改一处漏多处

影响：

- 页面部分模块能用，部分模块继续读旧字段

应对：

- 先建立页面字段映射表
- 后端统一输出 ViewModel

### 风险 3：删除 Vue2 代码时误删仍有引用的逻辑

影响：

- 构建失败或运行时报错

应对：

- 先全局搜索引用
- 分两步执行：先剥离引用，再删除文件

## 8. 计划交付物

本计划执行完成后，仓库中应至少新增或更新以下内容：

- Vue3 永续接口映射文档
- 修正后的 `vite.config.mjs`
- 清理后的 `localAcceptanceMocks.mjs`
- 对齐后的 `swap-api` DTO/logic/handler`
- 修正后的 `jobcenter` K 线与强平逻辑
- Vue2 永续遗留代码清理提交
- 后端与前端端到端测试脚本

## 9. 下一步建议

建议把这份计划继续拆成可执行任务，按以下 5 个工作包推进：

1. 工作包 A：Vue3 接口契约梳理
2. 工作包 B：代理与 mock 修复
3. 工作包 C：`swap-api` 契约与业务修复
4. 工作包 D：Vue2 永续遗留代码清理
5. 工作包 E：端到端测试与 smoke 脚本

如果你需要，我下一步可以直接基于这份文档继续拆成一份“逐任务执行计划”，每个工作包细化到文件级别和验证命令级别。
