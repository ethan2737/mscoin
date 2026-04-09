# repair-uc-exchange-entrust-history-endpoint

## Summary

The recovered Vue 3 user-center entrust history page no longer calls the stale `POST /exchange/order/personal/history` route. It now uses the canonical authenticated history route and renders a successful empty-state result in the local baseline.

## Contract Audit

### Confirmed frontend request before repair

`/#/uc/entrust/history` was issuing:

```json
POST /exchange/order/personal/history
{
  "pageNo": 1,
  "pageSize": 10
}
```

### Confirmed backend-supported canonical route

`exchange-api` already exposed:

- `POST /exchange/order/history`
- `POST /exchange/order/current`

and delegated them to `exchange.rpc` history/current handlers.

That made the `personal/history` path a frontend-side contract drift, not a missing RPC implementation.

## Repair

### Frontend contract repair

The user-center history page was updated to:

- call `api.exchange.history`
- send auth headers through the shared runtime helper
- unwrap the current backend response shape from `response.data.data`
- update pagination correctly for empty history responses
- degrade gracefully when symbol list loading fails

### Local baseline repair

The canonical history route initially exposed a second local blocker:

- `exchange.exchange_order` table was missing

A minimal local schema baseline was added in:

- `local-exchange-order-baseline.sql`

and applied to the local `exchange` database so the history route can return a successful empty-page result.

## Files Changed

- `mscoin-frontend/src/pages-vue3/uc/EntrustHistory.vue`
- `openspec/changes/repair-uc-exchange-entrust-history-endpoint/local-exchange-order-baseline.sql`

## Validation

### Direct API validation

After repair, the canonical endpoint returns a frontend-compatible success shape:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "content": [],
    "totalElements": 0,
    "number": 0,
    "totalPages": 1,
    "hasNext": false,
    "isLast": true
  }
}
```

### Real browser validation

Using Playwright:

1. logged in through the repaired local `/uc/login` flow
2. opened `http://127.0.0.1:3000/#/uc/entrust/history`

Observed result:

- no `pageerror`
- page request used `POST /exchange/order/history`
- no request hit `/exchange/order/personal/history`
- `POST /exchange/order/history` returned HTTP `200` with `code: 0`
- the table rendered an empty state with `No Data`
- pagination rendered `Total 0`

## Classification

Post-repair result: **pass**

The entrust history route contract is repaired and browser-validated. The current empty table is a valid empty-result state, not a route failure.

## Residual Noise

- `POST /market/symbol-thumb` still returns `500` in the current local environment, so the symbol filter options are not populated
- `tawk.to` still emits a third-party CORS/load error in the browser console

Neither issue blocks the repaired entrust history endpoint itself.
