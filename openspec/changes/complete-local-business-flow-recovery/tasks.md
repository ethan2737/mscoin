## 1. Content Backend Restoration
- [x] 1.1 Audit the current `ucenter-api` and `ucenter` ancillary/help/announcement route surface, handler coverage, and storage dependencies for the local development environment.
- [x] 1.2 Implement the missing local help and announcement backend endpoints, or wire their canonical routes to the existing implementations when the capability already exists.
- [x] 1.3 Seed the minimum local content baseline required for `/#/help`, `/#/help/list`, and `/#/notice/item/:id` to render non-empty content in development.
- [x] 1.4 Verify the help and announcement API families return `code: 0` and frontend-compatible payloads without `404` responses.

## 2. Market Thumb Recovery
- [x] 2.1 Identify the root cause of the current `POST /market/symbol-thumb` `500` in the local environment, including any missing schema, cache, or dependency prerequisites.
- [x] 2.2 Repair the local `market/symbol-thumb` execution path so the endpoint returns a stable wrapped success response with frontend-compatible thumb rows.
- [x] 2.3 Add or update the minimum local market seed data required for `symbol-thumb` and related homepage or trading pair selectors to return usable results.
- [x] 2.4 Verify `POST /market/symbol-thumb` succeeds repeatedly after service restart and is consumed by the active Vue 3 frontend without browser-side runtime errors.

## 3. Trading Lifecycle Baseline
- [x] 3.1 Audit the local exchange flow prerequisites for a minimal real spot trading lifecycle, including test account balances, pair availability, order persistence, and matching or state-transition behavior.
- [x] 3.2 Add the minimum local exchange schema or seed data required to submit at least one authenticated spot order successfully in development.
- [x] 3.3 Verify a submitted spot order becomes visible in current entrust and history entrust views through the canonical backend contracts.
- [x] 3.4 Document any remaining local-only limitations of the trading lifecycle that are outside this change's closure target.

## 4. Browser-Led Flow Closure Validation
- [x] 4.1 Execute a browser-led validation pass covering login, homepage market data, help content, announcement content, spot order submission, current entrust, and history entrust.
- [x] 4.2 Record the exact validation evidence, API outcomes, remaining warnings, and readiness classification in a final report artifact stored in this change directory.
- [x] 4.3 Re-run build and critical local service checks after the final fixes to confirm the development environment remains stable.
