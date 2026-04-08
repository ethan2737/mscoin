## Why

后端服务存在严重安全漏洞（golang.org/x/crypto v0.7.0 已知多个 CVE），且核心依赖 go-zero、gRPC、GORM 等落后 2-3 年版本，导致安全风险、性能落后、无法获得官方支持。采用保守升级策略，优先修复安全问题，核心框架适度升级，降低破坏性变更风险。

## What Changes

- **golang.org/x/crypto**: v0.7.0 → v0.49.0+（安全漏洞修复，**紧急**）
- **go-zero**: v1.4.4 → v1.7.x 或 v1.8.x（核心框架，适度升级）
- **gRPC**: v1.53.0 → v1.60.x（通信协议，中间版本）
- **GORM**: v1.24.6 → v1.25.x（数据库 ORM，小幅度升级）
- **go-redis**: v8.11.5 → v8.x 或 v9.x（根据 go-zero 兼容性决定）
- **etcd client**: v3.5.5 → v3.5.x 最新版（随 go-zero 依赖升级）
- **Go 语言版本**: 1.19/1.20 → 1.21+（部分依赖要求）

## Capabilities

### New Capabilities

- **security-patches**: 安全补丁更新，包括 golang.org/x/crypto、golang.org/x/net 等系统包
- **go-zero-upgrade**: go-zero 框架保守升级，保持 API 兼容
- **dependency-audit**: 依赖审计机制，定期检查和更新易受攻击的依赖

### Modified Capabilities

- **grpc-communication**: gRPC 通信协议升级，可能涉及负载均衡器行为变化
- **orm-database**: GORM 版本升级，API 基本兼容，修复已知 bug

## Impact

- **受影响服务**: ucenter、market、exchange、jobcenter 及其 API 网关
- **依赖变更**: 所有 go.mod 文件需更新，go.work 可能需要调整
- **代码修改**: 少量 API 调用变更（主要是 gRPC、go-redis）
- **测试要求**: 全量回归测试，特别是 gRPC 服务间调用
- **部署影响**: 需重启所有服务，建议灰度发布
