# Vue 3 Endpoint Drift Repair Report

## Audited Active Callers

The active Vue 3 tree previously contained these stale compatibility callers:

- `swap/Kline.vue`
  - `/api/swap/thumb`
  - `/api/market/coin/info`
  - `/api/swap/symbolInfo`
- `uc/EntrustCurrent.vue`
  - `/api/market/thumb`
  - `/api/exchange/order/cancel/:id`
- `uc/EntrustHistory.vue`
  - `/api/market/thumb`

## Confirmed Backend Route Mapping

Confirmed from current backend handler files:

- Market canonical routes:
  - `/market/symbol-thumb`
  - `/market/symbol-info`
  - `/market/coin-info`
  - `/market/history`
- Exchange HTTP routes confirmed in `exchange-api`:
  - `/exchange/order/history`
  - `/exchange/order/current`
  - `/exchange/order/add`

## Implemented Repairs

1. Confirmed market drift was repaired to canonical route names.
   - `uc/EntrustCurrent.vue` now reads `api.market.thumb` -> `/market/symbol-thumb`
   - `uc/EntrustHistory.vue` now reads `api.market.thumb` -> `/market/symbol-thumb`
   - `swap/Kline.vue` now reads `api.market.coinInfo` -> `/market/coin-info`

2. Swap callers now use one explicit compatibility strategy.
   - `swap/Kline.vue` no longer hardcodes ad-hoc `/api/swap/*` aliases.
   - It now uses the shared runtime API map:
     - `api.swap.thumb` -> `/swap/symbol-thumb`
     - `api.swap.symbolInfo` -> `/swap/symbol-info`

3. Exchange cancel-order caller now uses one explicit compatibility strategy.
   - `uc/EntrustCurrent.vue` no longer hardcodes `/api/exchange/order/cancel/:id`.
   - It now uses `api.exchange.orderCancel` -> `/exchange/order/cancel/:id`.

4. Vite compatibility proxies for `/api/market`, `/api/swap`, and `/api/exchange` were removed.
   - Active Vue 3 pages no longer depend on those stale aliases.

## Remaining Unresolved Contract Questions

These items remain intentionally unresolved and should be validated against a running local backend:

1. Swap HTTP route family
   - The active frontend now consistently uses `/swap/symbol-thumb` and `/swap/symbol-info`.
   - A matching HTTP handler file was not confirmed during this change, so this remains a documented compatibility assumption pending backend validation.

2. Exchange cancel-order HTTP route
   - The active frontend now consistently uses `/exchange/order/cancel/:id`.
   - The current `exchange-api` handler file did not expose that route during this audit, so this remains a documented compatibility assumption pending backend validation.

## Verification Result

Verification performed:

- Static scan of `src/pages-vue3/**` found no remaining `/api/market/*`, `/api/swap/*`, or `/api/exchange/*` callers.
- `pnpm run build` passed after the repair.
- Real-browser verification on representative routes found no stale `/api/*` requests.

Representative browser observations:

- `/#/swapindex/1`
  - No stale `/api/*` requests
  - No `pageerror`
- `/#/uc/entrust/current`
  - No stale `/api/*` requests
  - Route still redirects into auth flow and surfaces a backend `500` via `uc/check/login`
- `/#/uc/entrust/history`
  - No stale `/api/*` requests
  - Requests now use canonical `/market/symbol-thumb`
  - Backend `500` remains on auth/order endpoints

## Follow-up Ownership

The remaining work belongs to local-backend validation, not endpoint drift repair:

- Confirm whether `/swap/symbol-thumb` and `/swap/symbol-info` are the intended HTTP contract.
- Confirm whether `/exchange/order/cancel/:id` still exists or needs a different canonical mapping.
- Diagnose backend `500` responses once `8888`, `8889`, and `8890` are running in a known-good state.
