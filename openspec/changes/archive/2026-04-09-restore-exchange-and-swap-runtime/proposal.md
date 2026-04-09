## Why

The audit identified exchange as a P0 broken flow: the trading surface does not render, mounted hooks throw, and realtime/socket usage still points at legacy paths. Swap inherits the same class of problems. Since exchange and swap are the platform’s core product surfaces, they need a dedicated recovery change instead of being treated as incidental cleanup.

## What Changes

- Restore the boot path, render path, and realtime path for Vue 3 exchange and swap pages.
- Repair mount-time assumptions that currently cause null access, missing props, and unusable tables/forms.
- Align market data, depth, trade-feed, and favorites behavior with the active Vue 3 runtime contract.
- Define the minimum verifiable trading-surface behavior required before broader backend validation.

## Capabilities

### New Capabilities
- `exchange-swap-runtime-recovery`: Defines the minimum functional runtime behavior for Vue 3 exchange and swap surfaces, including render safety, market boot data, and realtime channel wiring.

### Modified Capabilities
- None.

## Impact

- Affected code: `mscoin-frontend/src/pages-vue3/exchange/**`, `src/pages-vue3/swapindex/**`, related shared chart/depth components, and market socket wiring.
- Affected systems: spot trading UI, contract trading UI, favorites, market feed, depth/trade panels.
- Dependency on audit findings: `openspec/changes/archive/2026-04-09-audit-frontend-business-and-data-flows/report.md`.
