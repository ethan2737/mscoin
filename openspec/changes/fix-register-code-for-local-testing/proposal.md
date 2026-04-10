## Why

当前本地测试注册流程时，短信验证码由后端动态生成并写入 Redis，测试人员需要额外读取缓存或依赖短信通道确认验证码，导致注册链路验证效率低且不稳定。为了让本地联调、冒烟测试和自动化验证可重复执行，需要将注册验证码发送逻辑改为固定值。

## What Changes

- 将注册发送验证码逻辑调整为固定返回并写入验证码 `123456`
- 保持注册接口、请求参数、Redis key 和验证码校验链路不变，只替换验证码生成策略
- 明确该固定验证码行为面向本地测试场景，便于后续 smoke check 和人工回归复用

## Capabilities

### New Capabilities
- `register-test-verification-code`: 定义本地测试环境下注册验证码发送与校验使用固定验证码 `123456` 的行为

### Modified Capabilities

## Impact

- 影响代码：`mscoin-backend/ucenter/internal/logic/register_logic.go`
- 影响流程：`/uc/mobile/code` 发送验证码、`/uc/register/phone` 注册校验
- 不涉及前端接口协议变更，不新增外部依赖
