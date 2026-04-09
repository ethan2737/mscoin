# Local Business Flow Recovery Report

## Outcome

Local development now supports a complete baseline business flow across frontend and backend:

- Vue 3 frontend starts, builds, and serves successfully.
- Local login succeeds through the real `/uc/login` path using the repaired fallback captcha contract.
- Homepage market data renders from the repaired local `market/symbol-thumb` path.
- Help and announcement backend content contracts return `code: 0` and feed the Vue 3 CMS pages.
- Authenticated spot order submission succeeds through `/exchange/order/add`.
- Submitted orders become visible through `/exchange/order/current` and `/exchange/order/history`.

## Kafka Stabilization

The main local instability was not the broker itself, but poisoned consumer behavior and premature order state transitions:

1. `exchange` previously created Kafka-domain consumers too broadly and also advanced order status too early.
2. `ucenter` retried stale `add-exchange-order` messages in a tight local loop when the upstream order no longer existed.
3. Shared consumer groups were reused across unrelated local topics, increasing rebalance noise.

Applied fixes:

- `mscoin-backend/exchange/internal/svc/service_context.go`
  - Reused a singleton `KafkaDomain`.
- `mscoin-backend/exchange/internal/database/kafka.go`
  - Used topic-specific consumer groups.
- `mscoin-backend/exchange/internal/logic/order_logic.go`
  - Removed the premature `TRADING` status update from the order-add RPC path.
- `mscoin-backend/ucenter/internal/database/kafka.go`
  - Used topic-specific consumer groups for local consumers.
- `mscoin-backend/ucenter/internal/consumer/order.go`
  - Dropped stale orphaned order messages instead of requeueing them forever.
- `mscoin-backend/ucenter/etc/conf.yaml`
  - Switched local `ExchangeRpc` to direct endpoint mode for development stability.

Result:

- Terminal spam is gone.
- `exchange` and `ucenter` no longer loop on stale messages.
- New spot orders now freeze wallet funds first, then transition to `TRADING` via Kafka callback.

## Frontend Closure Fixes

Additional Vue 3 closure fixes were required after browser validation:

- `mscoin-frontend/src/pages-vue3/cms/HelpList.vue`
  - Fixed the `route.query` watcher bug that caused the help-list runtime error.
- `mscoin-frontend/src/pages-vue3/uc/EntrustCurrent.vue`
  - Defaulted the symbol selector to a real active pair, preferring `BTC/USDT`.
- `mscoin-frontend/src/pages-vue3/uc/EntrustHistory.vue`
  - Defaulted the symbol selector to a real active pair, preferring `BTC/USDT`.
  - Fixed clear-action reload behavior.

## Verification Evidence

### Browser-led validation

Executed with Playwright against `http://127.0.0.1:3000`:

- `/#/login`
  - Passed and redirected to `/#/uc/safe`
- `/#/`
  - Passed with `3` rendered market rows
- `/#/help/list?cate=1&cateTitle=Help`
  - Passed with non-empty rendered content
- `/#/notice/item/3001`
  - Passed without page errors
- Browser-submitted spot order
  - Returned order id `E1775718758337409208`
- `/#/uc/entrust/current`
  - Passed and showed active `BTC/USDT` entrust rows
- `/#/uc/entrust/history`
  - Passed and showed historical `BTC/USDT` entrust rows

Browser residual warnings:

- `tawk.to` third-party script CORS/load failure
- two generic `404` asset/resource requests

No page errors remained in the final pass.

### API checks

Final local checks all returned `code: 0`:

- `/uc/login`
- `/market/symbol-thumb`
- `/uc/ancillary/more/help/page`
- `/uc/announcement/more`
- `/exchange/order/current`
- `/exchange/order/history`

### Build check

- `pnpm run build`
  - Passed
  - Residual warnings only: Sass legacy API deprecation and large bundle warnings

## Remaining Local-only Limitations

These do not block the recovered baseline flow, but they still exist:

- Spot orders remain in a local `TRADING` baseline state unless a matching/completion path is explicitly driven by local liquidity.
- `NoticeItem` still shows a brief loading state before content settles.
- Third-party `tawk.to` is noisy in local development and not required for business validation.

## Readiness Classification

Local development is now ready for complete business-flow testing.

This closure target is satisfied for:

- anonymous homepage browsing
- content-page access
- local login
- market visibility
- authenticated spot order submission
- current entrust visibility
- history entrust visibility
