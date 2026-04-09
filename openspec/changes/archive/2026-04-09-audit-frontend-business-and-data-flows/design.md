## Context

The active frontend baseline is now Vue 3 + Vite, but the repository still behaves like a migration in progress rather than a verified product surface. Several classes of problems are already visible:

- Homepage content mixes real API wiring, local fallback data, and static presentation without exposing which is which.
- Login and registration views show obvious encoding and migration regressions.
- Some interactions render but do not behave correctly, including menu/dropdown behavior reported during manual QA.
- Local backend services expected by the frontend are not consistently running, so current page inspection alone cannot distinguish frontend defects from missing dependencies.

The investigation must therefore work as a cross-cutting audit across routes, page components, shared app shell logic, API config, runtime proxying, and backend service expectations.

## Goals / Non-Goals

**Goals:**
- Produce a route-level inventory of the active Vue 3 frontend surface.
- Classify each audited UI region as static, API-backed, realtime-backed, auth-gated, or fallback/mocked.
- Trace key business flows end to end, including frontend trigger, route transition, API or websocket dependency, and expected backend service.
- Record findings in a reusable report that can drive later fix changes in a deterministic order.
- Make it possible for future QA to tell whether a page is truly healthy or merely rendering placeholders.

**Non-Goals:**
- Fix every issue discovered during the audit.
- Reconstruct the entire historical product behavior from legacy Vue 2 code during this change.
- Certify backend business correctness beyond the contracts the frontend currently depends on.

## Decisions

1. Audit by business flow, not by file tree.
   Rationale: users experience homepage browse, login, register, exchange entry, and user-center navigation as flows, not components. This reduces the chance of fixing local rendering while missing broken handoffs.
   Alternative considered: audit page files one by one. Rejected because it does not surface shared shell issues, route drift, or broken flow transitions clearly enough.

2. Separate display classification from defect classification.
   Rationale: each audited region first needs a source-of-truth label such as static, API-backed, websocket-backed, or fallback. Only after that can defects be described accurately.
   Alternative considered: list visible bugs only. Rejected because it leaves QA unable to determine what “correct” even means for each region.

3. Verify dependencies at the contract boundary.
   Rationale: for each flow, the audit should identify route entry, frontend caller, endpoint or socket path, required token state, dev proxy path, and expected backend service/port. This is the minimum boundary needed to separate frontend bugs from missing runtime dependencies.
   Alternative considered: stop at browser-network observations. Rejected because many current failures originate in hardcoded hosts, mismatched endpoint families, or backend services not started locally.

4. Treat the investigation report itself as a deliverable artifact.
   Rationale: the audit is only useful if it becomes a stable reference for later implementation changes and manual QA.
   Alternative considered: keep findings only in ad hoc chat summaries. Rejected because that is not durable enough for a multi-change migration thread.

## Risks / Trade-offs

- [The audit may uncover more defects than can be fixed quickly] -> Mitigation: require severity, reproduction notes, and remediation order in the report.
- [Backend services may remain unavailable during investigation] -> Mitigation: explicitly mark findings as frontend defect, dependency gap, or unresolved due to missing backend evidence.
- [Legacy Vue 2 behavior may differ from the intended Vue 3 target] -> Mitigation: use the selected Vue 3 baseline as the active contract and reference legacy Vue 2 code only to explain drift, not to silently redefine scope.
- [Manual QA can miss interaction bugs] -> Mitigation: combine code-path tracing with browser verification on the highest-risk routes and controls.

## Migration Plan

1. Define the audit scope around the active Vue 3 route tree and shared shell.
2. Build a route and business-flow matrix covering homepage, auth, notices/help, exchange, OTC, swap, and user center entry points.
3. For each audited flow, classify the UI regions and map the data dependencies to API/socket/backend contracts.
4. Verify the highest-risk flows in-browser and reconcile observed behavior with code-level expectations.
5. Publish a report that groups findings into frontend regressions, missing dependencies, and unresolved items requiring backend or product clarification.
6. Use that report to seed subsequent fix changes in priority order.

Rollback strategy: none for production behavior; this change is investigative and should only add audit artifacts and report outputs.

## Open Questions

- Which business flows are mandatory for “frontend usable” status versus acceptable deferred scope during migration recovery?
- Should the investigation report live under `docs/` as a persistent QA reference, or only under OpenSpec change artifacts before implementation begins?
- For login/register, is the intended target to preserve legacy UX exactly, or only restore a functionally correct Vue 3 equivalent?
