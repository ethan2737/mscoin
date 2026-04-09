# Vue 3 Network Contract Sweep

## Supported Rules

The active Vue 3 frontend now treats first-party traffic as same-origin traffic and relies on Vite proxy routing during local development.

- HTTP callers use relative first-party paths instead of absolute `http://localhost` URLs.
- Realtime callers use same-origin socket bootstrap (`io()`) instead of `io('http://localhost')`.
- Local Vite proxy targets now cover:
  - `/uc` -> `127.0.0.1:8888`
  - `/market` -> `127.0.0.1:8889`
  - `/exchange` -> `127.0.0.1:8890`
  - `/swap` -> `127.0.0.1:8890`
  - `/otc` -> `127.0.0.1:8888`
  - `/crowd` -> `127.0.0.1:8888`
  - `/api/market` -> `127.0.0.1:8889`
  - `/api/exchange` -> `127.0.0.1:8890`
  - `/api/swap` -> `127.0.0.1:8890`
  - `/socket.io` -> `127.0.0.1:8889`
- Explicit absolute URLs should remain only for third-party services, not first-party application APIs.

## Swept Page Families

This normalization sweep removed hardcoded localhost usage from the active Vue 3 tree in these families:

- CMS
- Activity and mining
- OTC
- Crowdfunding, CTC, and envelope pages
- Swap and swapindex pages
- User-center pages, including contract/entrust variants

## Verification Notes

Representative browser checks were run against:

- `/#/help`
- `/#/otc`
- `/#/activity`
- `/#/mining`
- `/#/findPwd`

Observed result:

- No representative route emitted requests to `http://localhost` or `ws://localhost`.
- No representative route threw a `pageerror`.
- Requests now consistently flow through `http://127.0.0.1:3000/...`, which means frontend traffic is using the normalized same-origin path before proxying.

## Remaining Endpoint-Family Mismatches

The sweep exposed endpoint drift that predates this change and still needs follow-up:

1. Legacy `/api/market/*` callers do not match current backend route names.
   - Frontend callers still use `/api/market/thumb`.
   - Current backend exposes `/market/symbol-thumb`.
   - Frontend callers still use `/api/market/coin/info`.
   - Current backend exposes `/market/coin-info`.

2. Legacy `/api/swap/*` callers remain unverified against the current backend.
   - `swap/Kline.vue` still calls `/api/swap/thumb` and `/api/swap/symbolInfo`.
   - Current backend route coverage for those exact compatibility paths was not found in this sweep.

3. Legacy cancel-order compatibility path remains unverified.
   - `uc/EntrustCurrent.vue` still calls `/api/exchange/order/cancel/:id`.
   - Current `exchange-api` handler file exposes `/order/current`, `/order/history`, and `/order/add`, but no matching cancel route was confirmed during this sweep.

4. Several normalized routes now surface backend failures directly instead of silently bypassing them.
   - `/uc/check/login`
   - `/uc/ancillary/more/help`
   - `/otc/coin`
   - `/uc/activity/page-query`
   - `/uc/mining/page-query`
   - `/uc/ancillary/system/app/version/0`

5. Non-network issues surfaced during browser verification and should be handled separately.
   - Vue Router alias param warning for `/announcement/:id`
   - Duplicate `store` provide warning in `App.vue`
   - `Activity.vue` unresolved `el-spinner`
   - External `tawk.to` CORS/load failure
   - `FindPwd.vue` eager async-validator warnings before input
