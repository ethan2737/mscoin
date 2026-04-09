## Context

The project has already recovered the Vue 3 frontend baseline, local login, and several core user-center and exchange entry paths, but the development environment still cannot execute a complete business flow end to end. The remaining blockers span multiple services:

- `ucenter-api` and `ucenter` do not expose the help/announcement content routes required by the active frontend
- `market/symbol-thumb` remains unstable in the local environment and breaks dependent page widgets
- trading validation has only reached route-level or empty-state checks; it has not yet proven a real local lifecycle from visible market data through order placement and order visibility

This is a cross-cutting recovery change touching backend contracts, local seed/baseline data, and browser validation. The desired outcome is not merely “pages load”, but “the development environment can run a complete business flow with repeatable local verification”.

## Goals / Non-Goals

**Goals:**
- Restore help and announcement backend capability required by the recovered Vue 3 frontend pages.
- Repair `market/symbol-thumb` so market and user-center pages can consistently load symbol metadata in local development.
- Seed the minimum local exchange data required to run a real order lifecycle in development.
- Validate a complete local business flow through a real browser, covering login, content, market visibility, order submission, current entrust, and history entrust.
- Publish a final change-local report that classifies the development environment as ready or not ready for complete business-flow testing.

**Non-Goals:**
- Rebuild the full production content management pipeline if a simpler local backend-backed baseline is sufficient.
- Seed a production-like order book or realistic market stream depth beyond what is required for local lifecycle verification.
- Remove all non-blocking console noise such as third-party chat script failures if they do not block business completion.
- Solve unrelated future enhancements outside the development-environment flow closure target.

## Decisions

1. Implement missing help/announcement capability in the local backend instead of replacing frontend pages with static mocks.
   Rationale: the stated goal is a development environment that runs the actual business flow, not a frontend-only demo shell.
   Alternatives considered:
   - Static/mock frontend content. Rejected because it would hide the missing backend contract and would not satisfy end-to-end validation.
   - Disable the affected pages locally. Rejected because these pages are part of the active recovered navigation surface.

2. Normalize on existing backend-wrapped response contracts rather than special-casing the frontend per page.
   Rationale: current active Vue 3 pages increasingly expect `{code,message,data}` envelopes; the remaining unstable pages should consume the canonical wrapped shape instead of maintaining mixed assumptions.
   Alternatives considered:
   - Introduce route-specific unwrapped responses. Rejected because it increases contract inconsistency.
   - Add frontend-only interceptors to coerce all shapes globally. Rejected because it obscures the actual endpoint contract and raises migration risk.

3. Treat local schema/data baselines as explicit change artifacts.
   Rationale: recent validation work showed that route repair alone is not enough; missing local tables or rows immediately collapse the flow back into backend `500`s. The minimum local schema and seed inputs must be documented and replayable.
   Alternatives considered:
   - Seed data manually without artifacts. Rejected because it is not repeatable.
   - Require full production snapshots locally. Rejected because it is heavier than needed for development-flow recovery.

4. Validate with a browser-led scenario matrix instead of API-only certification.
   Rationale: previous work already proved that API success alone misses frontend parsing and rendering assumptions. Final closure needs real-page navigation and state checks.
   Alternatives considered:
   - API-only validation. Rejected because it can still leave the visible flow broken.
   - Frontend-only screenshots. Rejected because they do not prove backend contract health.

## Risks / Trade-offs

- [Help/announcement data model is unclear in the current Go backend] → Mitigation: implement the thinnest backend-backed local capability needed by the active pages and document any simplified local constraints in the final report.
- [Repairing `market/symbol-thumb` may expose further missing market tables or incompatible seed assumptions] → Mitigation: include market schema/seed adjustments in the same recovery scope and validate directly against the canonical API before browser checks.
- [A real trade lifecycle may cascade into missing wallet, order, or Kafka-triggered side effects] → Mitigation: define a minimum local lifecycle goal and seed only the tables and supporting state required to prove order submission and visibility across current/history views.
- [The development environment may still pass the targeted flow while other pages remain noisy] → Mitigation: classify residual non-blocking noise separately so flow readiness is not confused with total system polish.

## Migration Plan

1. Audit and implement local help/announcement backend routes and supporting minimal data access/seed.
2. Repair `market/symbol-thumb` and any directly required market schema or seed dependencies.
3. Seed the minimum exchange/wallet/order baseline needed for a real local trading lifecycle.
4. Validate canonical APIs directly, then run a browser-led scenario matrix for the full flow.
5. Publish a final report summarizing pass/fail status, residual blockers, and the exact local baseline required.

Rollback strategy:
- Remove the newly added local seed artifacts and revert the backend route/data changes if they introduce unacceptable drift from intended local architecture.
- Keep the final report even on rollback so the failed assumptions remain documented.

## Open Questions

- Should help/announcement local data be stored in MySQL tables, served from a simplified local static dataset behind backend handlers, or reconstructed from another existing data source?
- What is the minimum acceptable “real trading lifecycle” for local closure: submit + current/history visibility only, or also cancellation/completion?
- If `market/symbol-thumb` depends on live stream-derived state, should the local baseline materialize static thumb rows or restore the full update path?
