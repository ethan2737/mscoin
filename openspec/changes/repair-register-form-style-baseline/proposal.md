## Why

注册页表单样式未同步登录页的新表单基线，导致同一深色主题场景下输入框容器、聚焦态、错误态、prepend 组合控件等视觉风格混乱，严重影响用户体验和表单可用性。

## What Changes

- **注册页输入框样式基线**：迁移到登录页同一暗色风格，统一背景、描边、hover/focus/error 态
- **区号选择框 + 输入框组合**：修正 prepend 容器与输入框的视觉一体化处理
- **验证码行按钮样式**：改为暗色次级按钮，明确 hover/focus/disabled 状态
- **表单项间距**：统一纵向节奏，消除字段"断层感"
- **页面布局**：改为垂直居中卡片布局，与登录页对齐
- **协议勾选与人机验证区**：优化交互样式和视觉层级

## Capabilities

### New Capabilities

无新能力，纯前端样式修复。

### Modified Capabilities

无 spec 级行为变更，仅实现细节调整。

## Impact

- **前端代码**：`mscoin-frontend/src/pages-vue3/uc/Register.vue` 样式部分
- **依赖**：Element Plus 组件样式覆盖
- **影响范围**：仅注册页面视觉，不影响注册 API 或业务逻辑
