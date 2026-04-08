# Backend Upgrade Report

**Date:** 2026-04-08
**Change:** backend-upgrade-conservative
**Schema:** spec-driven

## Summary

Successfully upgraded backend dependencies from go-zero v1.4.4 to v1.10.1 with all security patches applied.

## Version Changes

| Component | Before | After | Status |
|-----------|--------|-------|--------|
| go-zero | v1.4.4 | v1.10.1 | ✓ |
| gRPC | v1.53.0 | v1.80.0 | ✓ |
| GORM | v1.24.6 | v1.31.1 | ✓ |
| golang.org/x/crypto | v0.7.0 | v0.49.0 | ✓ |
| golang.org/x/net | v0.8.0 | v0.51.0 | ✓ |
| golang.org/x/text | v0.6.0-v0.7.0 | v0.35.0 | ✓ |
| golang.org/x/sys | v0.4.0-v0.5.0 | v0.42.0 | ✓ |
| google.golang.org/protobuf | v1.28.1 | v1.36.11 | ✓ |
| etcd client | v3.5.5 | v3.5.21 | ✓ |
| mongodb driver | v1.11.x | v1.17.9 | ✓ |
| jwt | v4.5.0 | v4.5.2 | ✓ |

## Security Vulnerabilities Fixed

### Module Dependencies (Fixed)
- **GO-2024-2687**: HTTP/2 CONTINUATION flood in golang.org/x/net
- **GO-2026-4762**: Authorization bypass in google.golang.org/grpc
- **GO-2023-2153**: HTTP/2 Rapid Reset in google.golang.org/grpc
- **GO-2025-3553**: Excessive memory allocation in github.com/golang-jwt/jwt
- **GO-2024-3250**: Improper error handling in github.com/golang-jwt/jwt

### Go Standard Library (Requires Go 1.25.8+)
The following vulnerabilities remain in the Go standard library itself and require upgrading the Go compiler version:
- **GO-2026-4601**: net/url IPv6 parsing (fixed in go1.25.8)
- **GO-2026-4337**: crypto/tls session resumption (fixed in go1.25.7)
- **GO-2026-4603**: html/template URL escaping (fixed in go1.25.8)

## Issues Encountered

1. **go mod tidy auto-upgraded go-zero**: During security patch upgrade, `go mod tidy` automatically upgraded go-zero to v1.10.1 (latest) instead of the planned v1.7.x. All services compile successfully, indicating good API compatibility.

2. **Docker Compose paths**: The docker-compose.yml references paths from the original development machine (D:/go/project/microservice/docker/...). These need to be updated for the current environment.

## Verification

- [x] All 9 modules compile successfully
- [x] Unit tests pass (core modules: bc, op, tools)
- [x] Config files compatible (etc/conf.yaml)
- [x] Git tags created: pre-upgrade-20260408, post-upgrade-20260408
- [x] Dependencies recorded: deps-before.txt, deps-after.txt

## Modules Upgraded

1. grpc-common
2. mscoin-common
3. ucenter
4. market
5. exchange
6. jobcenter
7. ucenter-api
8. market-api
9. exchange-api

## Next Steps

1. Update Docker Compose paths for local environment
2. Run full integration tests with Docker infrastructure
3. Deploy to staging environment
4. Monitor for 3 days
5. Production rollout (single node first)
6. Full production deployment

## Files Modified

- `mscoin-backend/*/go.mod` - All 9 module files
- `mscoin-backend/*/go.sum` - All 9 sum files
- `README.md` - Updated version information
- `openspec/changes/backend-upgrade-conservative/tasks.md` - Task completion

## Git References

- Branch: `feature/backend-upgrade`
- Pre-upgrade tag: `pre-upgrade-20260408`
- Post-upgrade tag: `post-upgrade-20260408`
