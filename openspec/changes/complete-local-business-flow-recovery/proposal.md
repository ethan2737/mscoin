## Why

The frontend and backend have been recovered to a usable local baseline, but the development environment still cannot execute a complete business flow end to end. Help and announcement content APIs are missing, `market/symbol-thumb` remains unstable, and trading validation has only reached partial or empty-state checks instead of a full local lifecycle.

## What Changes

- Restore the local backend help and announcement capability families required by the recovered Vue 3 frontend pages.
- Repair the local `market/symbol-thumb` backend path so market and user-center pages can load symbol metadata consistently in development.
- Seed and validate the minimum local exchange data needed to execute a real trading lifecycle instead of only empty-table checks.
- Run a full browser-led local validation covering anonymous entry, login, market visibility, order submission, current entrust, history entrust, and supporting content pages.
- Publish a final change-local validation artifact classifying the development environment as ready for complete business-flow testing.

## Capabilities

### New Capabilities

- `local-business-flow-recovery`: restores the missing backend/content/data contracts and validates a complete business flow in the local development environment.

### Modified Capabilities

- None.

## Impact

- `mscoin-backend/ucenter-api` and `mscoin-backend/ucenter` help/announcement handlers, contracts, and persistence
- `mscoin-backend/market-api` and `mscoin-backend/market` symbol-thumb path and market seed/runtime data
- `mscoin-backend/exchange-api` / `mscoin-backend/exchange` local trading schema and lifecycle validation support
- `mscoin-frontend/src/pages-vue3/**` pages that depend on content, market symbol metadata, and authenticated trading flows
- Local MySQL seed data and browser-based validation artifacts under `openspec/changes/complete-local-business-flow-recovery/`
