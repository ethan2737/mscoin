## ADDED Requirements

### Requirement: go-zero 框架保守升级

系统 SHALL 升级 go-zero 框架至 v1.7.x 或 v1.8.x 版本，保持 API 兼容性。

#### Scenario: 核心 API 兼容
- **WHEN** 现有代码使用 go-zero 的 rest、rpc、core 包
- **THEN** 编译通过，无需修改或仅需少量修改

#### Scenario: 配置文件兼容
- **WHEN** 使用现有的 etc/conf.yaml 配置文件
- **THEN** 配置结构兼容，服务正常加载配置

#### Scenario: 中间件兼容
- **WHEN** 使用 go-zero 的中间件（如 auth、log、trace）
- **THEN** 中间件正常工作，无需重写

---

### Requirement: go-zero 服务发现兼容

系统 SHALL 保持基于 Etcd 的服务发现机制，兼容 go-zero 新版本。

#### Scenario: 服务注册
- **WHEN** RPC 服务启动时
- **THEN** 成功注册到 Etcd，键为 `<service>.rpc`

#### Scenario: 服务发现
- **WHEN** API 网关调用 RPC 服务时
- **THEN** 通过 Etcd 正确发现服务地址，建立 gRPC 连接

---

### Requirement: go-zero 日志和监控兼容

系统 SHALL 保持 go-zero 的日志和 Prometheus 监控功能。

#### Scenario: 日志输出
- **WHEN** 服务运行时产生日志
- **THEN** 日志格式正确，包含 trace、span 等信息

#### Scenario: Prometheus 指标
- **WHEN** 访问 Prometheus 指标端点
- **THEN** 输出正确的 HTTP 和 gRPC 指标数据

---

### Requirement: go-zero 配置升级验证

系统 SHALL 验证 go-zero 升级后的配置项变更。

#### Scenario: 检查破坏性变更
- **WHEN** 阅读 go-zero changelog 从 v1.4.4 到目标版本
- **THEN** 记录所有破坏性变更和迁移步骤

#### Scenario: 配置项验证
- **WHEN** 对比新旧版本的配置结构体
- **THEN** 确认所有配置项兼容或已更新
