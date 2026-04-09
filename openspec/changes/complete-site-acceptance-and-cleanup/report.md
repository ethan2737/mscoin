# Full-Site Acceptance Report

Date: 2026-04-09
Change: `complete-site-acceptance-and-cleanup`

## Summary

This change closed the remaining acceptance gaps after the core spot-flow recovery and established a browser-verified local development baseline across the active Vue 3 site.

The local development environment can now be classified as:

- `spot`: `passed`
- `content`: `passed`
- `contract`: `passed_with_local_limitations`
- `otc`: `passed_with_local_limitations`
- `mining/crowdfunding`: `passed_with_local_limitations`
- `remaining user-center adjacent routes`: `passed_with_local_limitations`

The system is acceptance-ready for local development of the recovered active baseline. Remaining limitations are explicitly local-baseline limitations rather than unresolved unknown blockers.

## Acceptance Matrix

| Module | Status | Evidence | Notes |
| --- | --- | --- | --- |
| Home / market overview | `passed` | Browser verification on `/#/`; `/market/symbol-thumb` returns `code: 0` with 6 visible pairs | Local market baseline seeded earlier remains stable |
| Login / auth entry | `passed` | Real browser login through `/#/login` with local fallback captcha path; `/uc/login` returns `code: 0` | Token and member state are written normally |
| Spot current entrust | `passed` | Browser verification on `/#/uc/entrust/current`; `/exchange/order/current` returns `code: 0` | Stable authenticated rendering |
| Spot history entrust | `passed` | Browser verification on `/#/uc/entrust/history`; `/exchange/order/history` returns `code: 0` | Empty and populated states both render |
| Spot terminal lifecycle | `passed` | Real API path `submit -> cancel -> history` shows terminal `CANCELED` state | Local scope accepts cancel terminal lifecycle rather than matched-fill lifecycle |
| Help center | `passed` | Browser verification on `/#/help/list?cate=1&cateTitle=Help`; `/uc/ancillary/more/help/page` returns `code: 0` with 3 records | Content APIs now present locally |
| Announcement detail | `passed` | Browser verification on `/#/notice/item/3001`; `/uc/announcement/more` returns `code: 0` | Loading no longer lingers in a broken state |
| Contract trading route | `passed_with_local_limitations` | Browser verification on `/#/swapindex/1` | Uses local acceptance mock responses and a static chart placeholder on localhost |
| Contract user-center routes | `passed_with_local_limitations` | Browser verification on `/#/uc/contract/entrust/current` and `/#/uc/contract/entrust/history` | Page-level acceptance only; real contract backend/datafeed remains deferred |
| OTC list / trade | `passed_with_local_limitations` | Browser verification on `/#/otc` and `/#/otc/trade/usdt` | Local dev mocks provide stable browse baseline |
| Mining | `passed_with_local_limitations` | Browser verification on `/#/mining` | Local dev mock baseline only |
| Crowdfunding | `passed_with_local_limitations` | Browser verification on `/#/crowdfunding` | Local dev mock baseline only |
| Innovation minings | `passed_with_local_limitations` | Browser verification on `/#/uc/innovation/minings` | Stable empty-state rendering after API/path cleanup |

## Fixes Applied

### 1. Presentation and local-noise cleanup

- Reworked the announcement detail page to settle cleanly from initial loading into content:
  - `mscoin-frontend/src/pages-vue3/cms/NoticeItem.vue`
- Suppressed `tawk.to` only on `localhost` / `127.0.0.1` so local acceptance is not polluted by third-party CORS failures:
  - `mscoin-frontend/index.html`
- Cleaned active-page migration residue that was still visible during acceptance:
  - fixed broken OTC column labels in `mscoin-frontend/src/pages-vue3/otc/Trade.vue`
  - fixed invalid table structure and stale API paths in `mscoin-frontend/src/pages-vue3/uc/InnovationMinings.vue`
  - guarded crowdfunding swiper initialization in `mscoin-frontend/src/pages-vue3/crowdfunding/index.vue`

### 2. Local baseline closure for remaining domains

- Added a dedicated local acceptance mock plugin for active-but-not-fully-restored domains:
  - `mscoin-frontend/dev/localAcceptanceMocks.mjs`
  - enabled from `mscoin-frontend/vite.config.mjs`
- Added local-dev chart fallback for contract acceptance on localhost:
  - `mscoin-frontend/src/pages-vue3/swapindex/Swapindex.vue`
- Mocked residual homepage favorite lookup to prevent acceptance noise:
  - `/exchange/favor/find`

### 3. Spot terminal lifecycle closure

- Revalidated the local spot lifecycle and found a real blocker on cancel:
  - `/exchange/order/cancel/:orderId` returned business error `订单编号不能为空`
- Fixed local `exchange-api` cancel handling to accept the order id from the canonical path:
  - `mscoin-backend/exchange-api/internal/handler/order.go`
- Rebuilt and restarted `exchange-api.exe`
- Verified a real local order lifecycle:
  - login success
  - add order success
  - cancel success
  - history query returns same order with status `CANCELED`

This closes local terminal-state acceptance for the spot domain at the cancellation terminal state.

## Browser-Led Verification

Browser verification was executed against the local Vite site at `http://127.0.0.1:3000` using the local test account:

- account: `13800000000`
- password: `Passw0rd!Local`

Verified routes:

- `/#/`
- `/#/help/list?cate=1&cateTitle=Help`
- `/#/notice/item/3001`
- `/#/uc/entrust/current`
- `/#/uc/entrust/history`
- `/#/swapindex/1`
- `/#/uc/contract/entrust/current`
- `/#/uc/contract/entrust/history`
- `/#/otc`
- `/#/otc/trade/usdt`
- `/#/mining`
- `/#/crowdfunding`
- `/#/uc/innovation/minings`

Acceptance result:

- all routes rendered without `pageerror`
- no blocking `4xx`/`5xx` browser failures remained in the accepted flows
- announcement detail no longer remained in a misleading loading state
- local `tawk.to` noise no longer distorted acceptance results

## Build And Service Checks

Post-fix verification:

- `pnpm run build`: passed
- `/uc/login`: `code: 0`
- `/market/symbol-thumb`: `code: 0`
- `/uc/ancillary/more/help/page`: `code: 0`
- `/uc/announcement/more`: `code: 0`
- `/exchange/order/current`: `code: 0`
- `/exchange/order/history`: `code: 0`

Observed counts during final verification:

- market visible pairs: `6`
- help records: `3`
- current entrust rows: `9`
- history entrust rows: `10`

## Remaining Deferred Work

The following items remain explicitly deferred rather than unresolved:

- real contract backend/datafeed recovery beyond page-level local acceptance
- fully restored OTC backend rather than local browse-baseline mocks
- fully restored mining/crowdfunding backend rather than local browse-baseline mocks
- deeper matched/fill terminal lifecycle validation beyond the accepted cancel terminal state
- non-acceptance-critical residue such as Sass legacy warnings, bundle size warnings, and residual encoding cleanup outside the active accepted surfaces

## Conclusion

This change completes the site-wide acceptance pass for the active recovered Vue 3 baseline.

The local development environment can be considered end-to-end acceptance-ready for the recovered core business flow and surrounding active modules, with explicit local limitations documented for domains that currently rely on acceptance mocks instead of full backend restoration.
