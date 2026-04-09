## Why

The audit found a broad P1 class of defects across the active Vue 3 page tree: many pages still hardcode `http://localhost`, `ws://localhost`, or `io('http://localhost')`, bypassing the canonical Vite proxy/runtime path. This means even otherwise-correct pages can fail inconsistently depending on environment.

We need a dedicated normalization change so network access across the active Vue 3 baseline becomes environment-consistent and testable.

## What Changes

- Remove hardcoded localhost HTTP and websocket paths from the active Vue 3 page tree.
- Normalize active pages to use the supported runtime host/proxy/socket contract.
- Document the supported frontend network path rules for future migration work.
- Identify any flows that still require explicit backend route mapping after normalization.

## Capabilities

### New Capabilities
- `vue3-network-contract-normalization`: Defines the supported HTTP and realtime path rules for the active Vue 3 page tree and removes hardcoded localhost access.

### Modified Capabilities
- None.

## Impact

- Affected code: active Vue 3 pages under `mscoin-frontend/src/pages-vue3/**`, shared helpers, socket callers, and supporting network access patterns.
- Affected systems: CMS, OTC, activity, mining, user center, swap, and any page still bypassing proxy/runtime rules.
- Dependency on audit findings: `openspec/changes/archive/2026-04-09-audit-frontend-business-and-data-flows/report.md`.
