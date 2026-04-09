## 1. Contract Audit

- [x] 1.1 Confirm the exact frontend request path and payload currently used by `/#/uc/entrust/history`.
- [x] 1.2 Inspect current `exchange-api` and `exchange.rpc` history handlers to identify the intended canonical authenticated history route.

## 2. Endpoint Repair

- [x] 2.1 Implement or remap the authenticated spot entrust history route so the user-center page no longer receives `404`.
- [x] 2.2 Verify the repaired route returns a frontend-compatible success shape for empty or populated history results.

## 3. Authenticated Validation

- [x] 3.1 Re-run `/#/uc/entrust/history` in a real browser with the reusable local authenticated baseline.
- [x] 3.2 Classify the post-repair result as pass or blocked-by-data and publish a change-local validation artifact.
