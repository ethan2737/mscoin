## ADDED Requirements

### Requirement: golang.org/x/crypto 安全漏洞修复

系统 SHALL 立即升级 golang.org/x/crypto 包至 v0.49.0 或更高版本，修复所有已知安全漏洞。

#### Scenario: 漏洞修复验证
- **WHEN** 升级完成后运行 `go list -m golang.org/x/crypto`
- **THEN** 显示版本 >= v0.49.0，无已知 CVE 漏洞

#### Scenario: 依赖传递检查
- **WHEN** 运行 `go mod graph` 检查依赖树
- **THEN** 所有间接依赖的 golang.org/x/crypto 版本均 >= v0.49.0

---

### Requirement: golang.org/x/net 安全更新

系统 SHALL 升级 golang.org/x/net 包至与 golang.org/x/crypto v0.49.0+ 兼容的最新版本。

#### Scenario: HTTP/2 安全修复
- **WHEN** 服务处理 HTTP/2 请求
- **THEN** 不受已知的 HTTP/2 快速重置攻击影响

#### Scenario: TLS 连接安全
- **WHEN** 建立 TLS 连接时
- **THEN** 使用最新的 TLS 配置，无已知漏洞

---

### Requirement: golang.org/x/text 和 golang.org/x/sys 更新

系统 SHALL 升级 golang.org/x/text 和 golang.org/x/sys 包至最新稳定版本。

#### Scenario: 兼容性验证
- **WHEN** 编译所有服务模块
- **THEN** 无 golang.org/x/* 包的编译错误或警告

---

### Requirement: 安全补丁验证流程

系统 SHALL 建立安全补丁验证流程，确保升级后服务正常运行。

#### Scenario: 服务启动验证
- **WHEN** 升级后启动每个服务（ucenter、market、exchange、jobcenter 及其 API）
- **THEN** 服务正常启动，注册到 Etcd，无 panic 或错误日志

#### Scenario: gRPC 通信验证
- **WHEN** API 服务调用后端 gRPC 服务
- **THEN** 调用成功，响应时间正常，无认证或连接错误
