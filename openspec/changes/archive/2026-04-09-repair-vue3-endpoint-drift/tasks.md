## 1. Drift Audit

- [x] 1.1 Enumerate the active Vue 3 callers still using legacy `/api/market/*`, `/api/swap/*`, and `/api/exchange/*` compatibility paths.
- [x] 1.2 Confirm the current backend route names for each affected endpoint family from the backend handler files and note any unresolved families.

## 2. Caller Repair

- [x] 2.1 Replace confirmed stale market endpoint callers with canonical backend route names.
- [x] 2.2 Repair swap and exchange cancel-order callers using the chosen explicit mapping strategy for each unresolved or confirmed family.
- [x] 2.3 Update compatibility proxy configuration only where it is still needed after caller repair.

## 3. Verification

- [x] 3.1 Verify representative swap, market-derived user-center, and order-management routes in a real browser and confirm they no longer emit stale legacy endpoint names.
- [x] 3.2 Record the repaired endpoint mappings and any remaining unresolved backend-contract questions in a change-local report.
