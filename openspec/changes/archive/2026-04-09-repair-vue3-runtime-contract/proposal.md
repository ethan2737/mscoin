## Why

The audit confirmed a P0 migration contract fracture between the Vue 3 boot/runtime layer and many migrated pages. The app currently mixes global `host/api` helpers, `store.state.host/api` reads, and direct hardcoded URLs, which makes register, exchange, swap, and user-center flows structurally unstable before backend integration is even considered.

This change establishes one shared runtime contract so later page-level fixes can build on a coherent foundation.

## What Changes

- Define the canonical Vue 3 runtime contract for frontend pages, including network base access, API map access, auth token access, and shared helper exposure.
- Remove the mismatch between `main-vue3.js`, `store-vue3.js`, and migrated pages that still expect legacy `store.state.host/api` fields.
- Normalize how shared shell and migrated pages read app-level dependencies so they no longer depend on removed or implicit globals.
- Document the supported runtime-access pattern for later migration changes.

## Capabilities

### New Capabilities
- `vue3-runtime-contract`: Defines the single supported runtime contract for Vue 3 pages to access host, api definitions, auth state, and shared helpers.

### Modified Capabilities
- None.

## Impact

- Affected code: `mscoin-frontend/src/main-vue3.js`, `src/config/store-vue3.js`, `src/App.vue`, and migrated pages that currently read `store.state.host`, `store.state.api`, or implicit globals.
- Affected systems: register, exchange, swap, and user-center page boot behavior.
- Dependency on audit findings: `openspec/changes/archive/2026-04-09-audit-frontend-business-and-data-flows/report.md`.
