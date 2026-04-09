## Why

The audit explicitly separated frontend regressions from dependency gaps, but many critical business flows still cannot be certified because the local backend services were not running during inspection. After runtime, auth, exchange, and network-path recovery work lands, we need one validation-focused change that exercises the frontend against actual local backend services on the expected ports.

## What Changes

- Define the local backend prerequisites required for meaningful frontend validation.
- Validate key business flows against running `ucenter-api`, `market-api`, and `exchange-api` services.
- Confirm which homepage, auth, notice/help, exchange, and user-center regions switch from fallback or placeholder behavior to real backend-backed behavior.
- Produce an updated validation record that distinguishes frontend-fixed issues from backend/environment issues.

## Capabilities

### New Capabilities
- `frontend-local-backend-validation`: Defines how the recovered Vue 3 frontend is validated against the expected local backend services and how the results are recorded.

### Modified Capabilities
- None.

## Impact

- Affected systems: local services on `8888`, `8889`, `8890`, frontend proxy behavior, auth/token flows, market feed, notices/help, exchange and wallet/order queries.
- Affected outputs: validation evidence and updated recovery status after frontend recovery changes are applied.
- Dependency on audit findings: `openspec/changes/archive/2026-04-09-audit-frontend-business-and-data-flows/report.md`.
