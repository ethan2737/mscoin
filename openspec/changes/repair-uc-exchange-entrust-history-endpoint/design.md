## Context

The recovered frontend now has a reusable local authenticated baseline, and real-browser validation proved that `/#/uc/entrust/history` is no longer blocked by auth. Instead, the page emits a `404` on `POST /exchange/order/personal/history`. Current `exchange-api` route registration exposes `/order/history` and `/order/current`, but not the `/order/personal/*` family expected by the user-center page. This is a contract repair problem first, not merely a missing data problem.

## Goals / Non-Goals

**Goals:**
- Determine and implement the canonical authenticated backend route for spot entrust history used by the user-center page.
- Remove the current `404` from authenticated `/#/uc/entrust/history`.
- Re-run browser validation after the contract is repaired and classify whether the page then passes or is only blocked by missing order rows.

**Non-Goals:**
- Rebuild all exchange user-center endpoints in one change.
- Seed a full trading/order lifecycle unless the repaired route requires minimal fixture data to validate.
- Solve unrelated announcement or login captcha issues.

## Decisions

1. Treat the current `404` as a route-contract defect before adding seed order data.
   Rationale: order fixtures are irrelevant until the authenticated history endpoint itself resolves successfully.
   Alternative considered: seed orders immediately. Rejected because the current route never reaches a backend handler.

2. Prefer convergence onto one canonical authenticated history route instead of carrying parallel legacy and repaired endpoint families unnecessarily.
   Rationale: user-center validation should not keep accumulating compatibility aliases without clear ownership.
   Alternative considered: preserve both `/order/history` and `/order/personal/history` long-term. Rejected unless compatibility is required during migration.

3. Validate through the actual Vue 3 user-center page after repair.
   Rationale: direct API checks alone would miss frontend request-shape or rendering assumptions.
   Alternative considered: validate only with backend calls. Rejected because this change exists to unblock frontend user-center certification.

## Risks / Trade-offs

- [The frontend route may rely on response structure that differs from the current backend history payload] → Mitigation: validate both API response shape and browser rendering after route repair.
- [Repairing the route may expose a second blocker from missing order rows] → Mitigation: classify data blockage separately after the `404` is removed.
- [Adding compatibility mapping may hide a better canonical route choice] → Mitigation: document the final chosen contract in the change report.

## Migration Plan

1. Audit the current user-center history request and current exchange-api route registrations.
2. Repair the route contract by implementing or remapping the authenticated history path.
3. Re-run authenticated browser validation on `/#/uc/entrust/history`.
4. If the route succeeds but data is empty, record the minimum follow-up seed requirement separately.

Rollback strategy: restore the prior route registration state if the repair breaks current exchange-api history behavior elsewhere.

## Open Questions

- Should the repaired contract standardize on `/exchange/order/history` for authenticated user-center use, or preserve `/exchange/order/personal/history` as the canonical frontend-facing route?
- Does the current user-center history page require minimal seeded completed/canceled orders to validate table rendering beyond route success?
