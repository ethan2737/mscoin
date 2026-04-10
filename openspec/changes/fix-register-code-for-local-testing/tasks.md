## 1. 后端验证码逻辑调整

- [x] 1.1 修改 `mscoin-backend/ucenter/internal/logic/register_logic.go` 中注册发送验证码逻辑，将写入 Redis 的验证码固定为 `123456`
- [x] 1.2 保持注册提交阶段继续读取 `REGISTER::<phone>` 并按缓存值校验，不额外引入绕过分支

## 2. 验证与回归

- [x] 2.1 重启本地 `ucenter` 服务，确认运行中的进程已加载固定验证码逻辑
- [x] 2.2 调用 `/uc/mobile/code` 后检查 Redis，确认 `REGISTER::<phone>` 的值为 `123456`
- [x] 2.3 使用验证码 `123456` 提交 `/uc/register/phone`，确认验证码校验通过；使用其他验证码提交时确认仍返回验证码错误
