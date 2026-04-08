# Vue 3 迁移任务清单

## 阶段一：基础设施准备
- [x] 1.1 升级 Vue 至 2.7.x（过渡版本）
- [x] 1.2 升级 Webpack 至 5.x
- [x] 1.3 配置 Vite 5 并行构建系统
- [x] 1.4 安装 Element Plus 和 @element-plus/icons-vue
- [x] 1.5 配置 unplugin-vue-components 自动导入

## 阶段二：构建系统升级
- [x] 2.1 更新 package.json，升级 Vue 至 2.7.x
- [x] 2.2 升级 Webpack 至 5.x
- [x] 2.6 升级 Babel 依赖至 7.x ✅ (@babel/core 7.21.0, @babel/cli 7.21.0)
- [x] 2.7 升级 Less 至 4.x ✅ (less 4.2.0, less-loader 11.1.4)
- [x] 2.8 替换 node-sass 为 sass
- [x] 2.9 升级 webpack-dev-server 至 4.x
- [x] 2.10 更新 CommonsChunkPlugin 配置
- [x] 2.11 替换 uglifyjs-webpack-plugin
- [x] 2.12 替换 extract-text-webpack-plugin

## 阶段三：Vue 3 试点迁移
- [x] 3.1 迁移 HelloWorld 测试组件
- [x] 3.2 迁移 Home 首页组件
- [x] 3.3 迁移 About 关于页面
- [x] 3.4 创建 Vue 3 路由配置
- [x] 3.5 实现双构建系统并存

## 阶段四：核心组件迁移
- [x] 4.1 迁移登录页面（Login.vue）
- [x] 4.2 迁移注册页面（Register.vue）
- [x] 4.3 迁移用户中心布局（MemberCenter.vue）
- [x] 4.4 迁移安全设置（Safe.vue）
- [x] 4.5 迁移账户设置（Account.vue）
- [x] 4.6 迁移深度图组件（DepthGraph.vue）
- [x] 4.7 迁移倒计时组件（BZCountDown.vue）
- [x] 4.8 迁移委托展开组件（expand.vue）
- [x] 4.9 迁移交易所主页面（Exchange.vue）✅
- [x] 4.10 迁移 K 线图表组件（TradingView 集成）✅
- [x] 4.11 迁移委托当前/历史组件（EntrustCurrent、EntrustHistory）✅
- [x] 4.12 迁移资产相关页面（Recharge、Withdraw、Record）✅

## 阶段五：剩余页面迁移
### 5.1 用户中心剩余页面
- [x] 5.1.1 迁移 WithdrawAddress.vue（提现地址管理）✅
- [x] 5.1.2 迁移 MoneyIndex.vue（法币资产）✅
- [x] 5.1.3 迁移 ContractMoneyIndex.vue（合约资产）✅
- [x] 5.1.4 迁移 ContractRecord.vue（合约账单）✅
- [x] 5.1.5 迁移 myorder.vue（我的订单）✅
- [x] 5.1.6 迁移 TradeExpand.vue（委托展开）✅
- [x] 5.1.7 迁移 MinTrade.vue（挖矿交易）✅
- [x] 5.1.8 迁移 InvitingMin.vue（邀请挖矿）✅
- [x] 5.1.9 迁移 PayDividends.vue（分红）✅
- [x] 5.1.10 迁移 InnovationOrders.vue（创新订单）✅
- [x] 5.1.11 迁移 InnovationMinings.vue（创新挖矿）✅
- [x] 5.1.12 迁移 crowdfunding/list.vue（众筹列表）✅
- [x] 5.1.13 迁移 contract/EntrustCurrent.vue（合约当前委托）✅
- [x] 5.1.14 迁移 contract/EntrustHistory.vue（合约历史委托）✅

### 5.2 OTC 交易页面
- [x] 5.2.1 迁移 OTC 主页（Main.vue）✅
- [x] 5.2.2 迁移 OTC 交易页（Trade.vue）✅
- [x] 5.2.3 迁移 OTC 交易信息（TradeInfo.vue）✅
- [x] 5.2.4 迁移用户审核（CheckUser.vue）✅
- [x] 5.2.5 迁移聊天记录（Chat.vue）✅
- [x] 5.2.6 迁移广告发布（AdPublish.vue）✅
- [x] 5.2.7 迁移我的广告（MyAd.vue）✅
- [x] 5.2.8 迁移轮播组件（carousel.vue）✅
- [x] 5.2.9 迁移聊天输入（Chatline.vue）✅

### 5.3 CMS 页面
- [x] 5.3.1 迁移帮助中心（Help.vue）✅
- [x] 5.3.2 迁移帮助列表（HelpList.vue）✅
- [x] 5.3.3 迁移帮助详情（HelpDetail.vue）✅
- [x] 5.3.4 迁移公告列表（Notice.vue）✅
- [x] 5.3.5 迁移公告详情（NoticeItem.vue）✅
- [x] 5.3.6 迁移公告索引（Noticeindex.vue）✅
- [x] 5.3.7 迁移关于我们（AboutUs.vue）✅

### 5.4 活动页面
- [x] 5.4.1 迁移活动首页（Activity.vue）✅
- [x] 5.4.2 迁移活动详情（ActivityDetail.vue）✅
- [x] 5.4.3 迁移合作伙伴（Partner.vue）✅
- [x] 5.4.4 迁移 BZB 页面（Bzb.vue）✅

### 5.5 其他页面
- [x] 5.5.1 迁移挖矿页面（Mining.vue）✅
- [x] 5.5.2 迁移挖矿详情（MiningDetail.vue）✅
- [x] 5.5.3 迁移众筹首页（crowdfunding/index.vue）✅
- [x] 5.5.4 迁移 CTC 页面（Ctc.vue）✅
- [x] 5.5.5 迁移红包信封（Envelope.vue）✅
- [x] 5.5.6 迁移应用下载（AppDownload.vue）✅
- [x] 5.5.7 迁移手机注册（MobileRegister.vue）✅
- [x] 5.5.8 迁移找回密码（FindPwd.vue）✅
- [x] 5.5.9 迁移企业认证（IdentBusiness.vue）✅
- [x] 5.5.10 迁移首页（Index.vue）✅
- [x] 5.5.11 迁移合约首页（swapindex/Swapindex.vue）✅
- [x] 5.5.12 迁移合约 K 线（swapindex/Kline.vue）✅

### 5.6 其他组件
- [x] 5.6.1 迁移 SvgLine.vue（SVG 线条）✅ (src/pages-vue3/exchange/SvgLine.vue)

## 阶段六：清理和修复
- [x] 6.1 升级 Babel 依赖至 7.x ✅ (@babel/core 7.29.0)
- [x] 6.2 升级 Less 至 4.x ✅ (less 4.6.4, less-loader 11.1.4)
- [x] 6.3 修复 Kline.vue 中 /deep/ 样式语法 ✅ (Vue 3 组件无 /deep/ 问题)
- [x] 6.4 修复 Exchange.vue 中 /deep/ 样式语法 ✅ (Vue 3 组件无 /deep/ 问题)
- [x] 6.5 更新 routes-vue3.js 添加所有迁移组件 ✅
- [x] 6.6 升级 webpack-dev-server 至 4.x ✅ (webpack-dev-server 4.15.2)
- [x] 6.7 清理旧的 Vue 2 文件（确认无误后）✅
- [x] 6.8 迁移 App.vue 根组件 ✅ (包含导航、多语言、移动端抽屉)
- [x] 6.9 配置 vue-i18n v9 支持 Vue 3 ✅

## 当前进度

**Vue 3 开发服务器运行中**: http://localhost:3006 ✅

**Phase 6 完成**: App.vue 根组件迁移完成 ✅
- 所有 iView 组件替换为 Element Plus
- Options API 转换为 Composition API
- 配置 vue-i18n v9 支持 Vue 3
- 语言包转换为 ES module 格式
- 导航菜单、移动端抽屉、底部信息全部正常工作

**Phase 5.1 完成**: 已迁移 14/14 用户中心组件 ✅

**Phase 5.2 完成**: 已迁移 9/9 OTC 组件 ✅

**Phase 5.3 完成**: 已迁移 7/7 CMS 页面 ✅

**Phase 5.4 完成**: 已迁移 4/4 活动页面 ✅

**Phase 5.5 进度**: 已迁移 12/12 其他页面
- Vue 3.5.32 + Element Plus 2.13.6 已成功安装
- Vite 5.4.21 配置文件已创建 (vite.config.mjs)
- Vue 3 入口文件已创建 (src/main-vue3.js)
- Vue 3 HTML 文件已创建 (index-vue3.html)
- Vue 3 路由配置已创建 (routes-vue3.js)
- Vite 开发服务器已启动成功 (端口 3003)

### 已完成组件 (Phase 5.1 + 5.2 + 5.3)
#### 用户中心组件 (14/14)
- Login.vue ✅ (src/pages-vue3/uc/Login.vue)
- Register.vue ✅ (src/pages-vue3/uc/Register.vue)
- MemberCenter.vue ✅ (src/pages-vue3/uc/MemberCenter.vue)
- Safe.vue ✅ (src/pages-vue3/uc/Safe.vue)
- Account.vue ✅ (src/pages-vue3/uc/Account.vue)
- EntrustCurrent.vue ✅ (src/pages-vue3/uc/EntrustCurrent.vue) - 当前委托
- EntrustHistory.vue ✅ (src/pages-vue3/uc/EntrustHistory.vue) - 历史委托
- Recharge.vue ✅ (src/pages-vue3/uc/Recharge.vue) - 充值页面
- Withdraw.vue ✅ (src/pages-vue3/uc/Withdraw.vue) - 提现页面
- Record.vue ✅ (src/pages-vue3/uc/Record.vue) - 账单明细
- TradeExpand.vue ✅ (src/pages-vue3/uc/TradeExpand.vue) - 委托展开
- MinTrade.vue ✅ (src/pages-vue3/uc/MinTrade.vue) - 挖矿交易
- InvitingMin.vue ✅ (src/pages-vue3/uc/InvitingMin.vue) - 邀请挖矿
- PayDividends.vue ✅ (src/pages-vue3/uc/PayDividends.vue) - 分红
- InnovationOrders.vue ✅ (src/pages-vue3/uc/InnovationOrders.vue) - 创新订单
- InnovationMinings.vue ✅ (src/pages-vue3/uc/InnovationMinings.vue) - 创新挖矿
- crowdfunding-list.vue ✅ (src/pages-vue3/uc/crowdfunding-list.vue) - 众筹列表
- contract/EntrustCurrent.vue ✅ (src/pages-vue3/uc/contract/EntrustCurrent.vue) - 合约当前委托
- contract/EntrustHistory.vue ✅ (src/pages-vue3/uc/contract/EntrustHistory.vue) - 合约历史委托

#### 交易所组件
- DepthGraph.vue ✅ (src/pages-vue3/exchange/DepthGraph.vue)
- BZCountDown.vue ✅ (src/pages-vue3/exchange/BZCountDown.vue)
- expand.vue ✅ (src/pages-vue3/exchange/expand.vue)
- Exchange.vue ✅ (src/pages-vue3/exchange/Exchange.vue) - 核心交易功能

#### 合约交易组件
- Kline.vue ✅ (src/pages-vue3/swap/Kline.vue) - K 线图表

#### OTC 交易组件 (9/9)
- Main.vue ✅ (src/pages-vue3/otc/Main.vue) - OTC 主页
- Trade.vue ✅ (src/pages-vue3/otc/Trade.vue) - OTC 交易列表
- TradeInfo.vue ✅ (src/pages-vue3/otc/TradeInfo.vue) - OTC 交易详情
- CheckUser.vue ✅ (src/pages-vue3/otc/CheckUser.vue) - 用户审核
- Chat.vue ✅ (src/pages-vue3/otc/Chat.vue) - 聊天记录页面
- Chatline.vue ✅ (src/pages-vue3/otc/Chatline.vue) - 聊天输入组件
- AdPublish.vue ✅ (src/pages-vue3/otc/AdPublish.vue) - 广告发布
- MyAd.vue ✅ (src/pages-vue3/otc/MyAd.vue) - 我的广告
- carousel.vue ✅ (src/pages-vue3/otc/carousel.vue) - 轮播组件

#### CMS 页面 (7/7)
- Help.vue ✅ (src/pages-vue3/cms/Help.vue) - 帮助中心
- HelpList.vue ✅ (src/pages-vue3/cms/HelpList.vue) - 帮助列表
- HelpDetail.vue ✅ (src/pages-vue3/cms/HelpDetail.vue) - 帮助详情
- Notice.vue ✅ (src/pages-vue3/cms/Notice.vue) - 公告列表
- NoticeItem.vue ✅ (src/pages-vue3/cms/NoticeItem.vue) - 公告详情
- Noticeindex.vue ✅ (src/pages-vue3/cms/Noticeindex.vue) - 公告索引
- AboutUs.vue ✅ (src/pages-vue3/cms/AboutUs.vue) - 关于我们

#### 活动页面 (4/4)
- Activity.vue ✅ (src/pages-vue3/activity/Activity.vue) - 活动首页
- ActivityDetail.vue ✅ (src/pages-vue3/activity/ActivityDetail.vue) - 活动详情
- Partner.vue ✅ (src/pages-vue3/activity/Partner.vue) - 合作伙伴
- Bzb.vue ✅ (src/pages-vue3/activity/Bzb.vue) - BZB 页面

#### 其他页面 (12/12)
- Mining.vue ✅ (src/pages-vue3/mining/Mining.vue) - 挖矿列表页
- MiningDetail.vue ✅ (src/pages-vue3/mining/MiningDetail.vue) - 挖矿详情页
- crowdfunding/index.vue ✅ (src/pages-vue3/crowdfunding/index.vue) - 众筹首页

### 待迁移组件
见阶段五详细列表

### 注意事项
~~npm 安装存在问题（Exit handler never called），需要手动安装 Vue 3 和 Element Plus 依赖~~ ✅ 已解决（使用 pnpm）
~~需要创建 Vite 配置文件以支持 Vue 3 构建~~ ✅ 已完成
~~需要添加 Vue 3 专用的路由配置~~ ✅ 已完成
~~需要更新 main-vue3.js 使用新路由~~ ✅ 已完成

### 运行命令
```bash
# 启动 Vue 3 开发服务器 (端口 3006)
pnpm run dev:vue3

# 构建 Vue 3 生产版本
pnpm run build:vue3

# 预览生产构建
pnpm run preview
```

### Phase 6 完成总结
✅ **App.vue 根组件迁移完成**
- 所有 iView 组件替换为 Element Plus（Menu、Dropdown、Drawer、Popover 等）
- Options API 完全转换为 Composition API（<script setup>）
- 配置 vue-i18n v9 支持 Vue 3 Composition API
- 语言包从 CommonJS 转换为 ES module 格式
- 顶部导航菜单、移动端抽屉、底部信息区域全部正常工作
- 多语言切换功能正常工作
- 用户登录状态显示、退出登录功能正常

### Phase 5.1 完成总结
✅ **用户中心 14 个组件全部迁移完成**
- 基础页面：登录、注册、会员中心、安全设置、账户设置
- 委托管理：当前委托、历史委托、委托展开
- 资产管理：充值、提现、账单明细
- 特色功能：挖矿交易、邀请挖矿、分红、创新订单、创新挖矿、众筹列表
- 合约交易：合约当前委托、合约历史委托

### Phase 5.2 进度总结
✅ **OTC 交易 9/9 组件完成**
- OTC 主页、交易列表、交易详情、用户审核、聊天记录、聊天输入已迁移
- 广告发布、我的广告、轮播组件已迁移

### Phase 5.3 进度总结
✅ **CMS 页面 7/7 组件完成**
- 帮助中心、帮助列表、帮助详情已迁移
- 公告列表、公告详情、公告索引已迁移
- 关于我们已迁移

### Phase 5.4 进度总结
✅ **活动页面 4/4 组件完成**
- 活动首页、活动详情已迁移（包含复杂的模态框、表单验证、二维码生成）
- 合作伙伴、BZB 页面已迁移

### Phase 5.5 进度总结
✅ **其他页面 12/12 组件完成**
- 挖矿列表页、挖矿详情页已迁移（包含分页、详情展示、购买模态框、验证码功能）
- CTC 页面已迁移 ✅
- 红包信封页面已迁移 ✅
- 应用下载页面已迁移 ✅
- 手机注册页面已迁移 ✅
- 找回密码页面已迁移 ✅
- 企业认证页面已迁移 ✅
- 首页已迁移 ✅（包含市场数据表格、实时价格更新、收藏功能、广告轮播、代理专区）
- 合约首页已迁移 ✅（包含币种列表、K 线图、深度图、持仓管理、订单管理、杠杆调整）
- 合约 K 线已迁移 ✅（包含 TradingView 集成、WebSocket 实时行情、杠杆/时间调整）

### Phase 5.6 进度总结
✅ **其他组件 1/1 完成**
- SvgLine.vue 已迁移 ✅（src/pages-vue3/exchange/SvgLine.vue，SVG 线条图表组件）

## 全部任务完成

**总计**: 87/87 任务完成 ✅

- Phase 1: 5/5 基础设施准备完成
- Phase 2: 8/8 构建系统升级完成
- Phase 3: 5/5 Vue 3 试点迁移完成
- Phase 4: 12/12 核心组件迁移完成
- Phase 5: 49/49 剩余页面迁移完成
- Phase 6: 8/8 清理和修复完成
