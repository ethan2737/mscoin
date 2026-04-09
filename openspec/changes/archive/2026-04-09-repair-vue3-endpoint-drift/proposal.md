## Why

The network normalization sweep removed hardcoded localhost usage, but it also exposed a narrower P1 problem: several active Vue 3 pages still call legacy endpoint names that no longer match the current backend route surface. Until those endpoint drifts are repaired, proxy normalization alone cannot make swap, market, and order-management flows trustworthy.

## What Changes

- Audit the active Vue 3 callers that still depend on legacy `/api/market/*`, `/api/swap/*`, and `/api/exchange/*` compatibility paths.
- Align those callers to the current backend route contract where the backend route is known and confirmed.
- For endpoint families whose current backend route is still unclear, isolate and document the mismatch so the implementation can choose a single supported compatibility strategy instead of preserving accidental drift.
- Record the exact frontend-to-backend mapping used after the repair so later local-backend validation can test the canonical paths.

## Capabilities

### New Capabilities
- `vue3-endpoint-drift-repair`: Defines how active Vue 3 pages map legacy frontend endpoint calls to the current backend route contract and removes confirmed endpoint-name drift.

### Modified Capabilities
- None.

## Impact

- Affected code: active Vue 3 pages under `mscoin-frontend/src/pages-vue3/**`, Vite proxy compatibility paths, and any shared helpers introduced to normalize endpoint access.
- Affected systems: market data callers, swap callers, and user-center order-management callers that currently rely on legacy compatibility URLs.
- Dependency on prior findings: [report.md](E:/Project/web3/mscoin/openspec/changes/normalize-vue3-network-contracts/report.md).
