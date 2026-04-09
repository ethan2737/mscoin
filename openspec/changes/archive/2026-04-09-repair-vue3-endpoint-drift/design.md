## Context

The active Vue 3 frontend now uses same-origin network paths, so residual failures are easier to classify. The remaining drift is no longer about environment-specific hostnames; it is about path naming mismatches between migrated pages and the current Go backend handlers. The normalization report confirmed three drift families: legacy market aliases, legacy swap aliases, and a legacy exchange cancel-order alias.

## Goals / Non-Goals

**Goals:**
- Replace confirmed stale endpoint names in the active Vue 3 tree with the current backend route names.
- Keep one explicit compatibility strategy for unresolved endpoint families instead of letting pages drift independently.
- Produce a concrete endpoint mapping record that the later local-backend validation change can consume.

**Non-Goals:**
- Fix backend business failures or data correctness issues.
- Redesign all frontend data access patterns.
- Resolve unrelated Vue warnings, router warnings, or third-party script failures.

## Decisions

1. Repair only endpoint-name drift that affects the active Vue 3 baseline.
   Rationale: the goal is to make current migrated pages callable against the known backend, not to modernize every historical endpoint in the repo.
   Alternative considered: repair all frontend endpoint aliases globally. Rejected because it expands into legacy Vue 2 scope and slows the active recovery path.

2. Prefer direct alignment to confirmed backend route names over adding more proxy aliases.
   Rationale: if the backend handler already exposes a canonical route, the frontend should call that route directly so future debugging is straightforward.
   Alternative considered: preserve legacy frontend callers and keep translating them in Vite. Rejected because it hides drift and adds more compatibility surface to maintain.

3. For unconfirmed endpoint families, require an explicit mapping note in the change output.
   Rationale: some swap and cancel-order callers still reference paths whose backend equivalents were not confirmed during the sweep. Those should be resolved deliberately, not guessed piecemeal.
   Alternative considered: leave them unchanged without documentation. Rejected because it would preserve the same ambiguity that blocked validation.

## Risks / Trade-offs

- [A frontend caller may be moved to the wrong canonical backend path] -> Mitigation: only replace paths that were confirmed from backend handler files, and record any unresolved family separately.
- [Removing proxy aliases too early could break still-unmigrated callers] -> Mitigation: this change repairs callers first and treats proxy cleanup as optional follow-up, not as the primary success condition.
- [Some routes may still return backend 500s after drift repair] -> Mitigation: that is acceptable here because the purpose is contract alignment, not backend validation.

## Migration Plan

1. Enumerate active Vue 3 callers still using legacy compatibility paths.
2. Match each caller against confirmed backend routes from `market-api`, `exchange-api`, and any other relevant handler files.
3. Update confirmed callers to canonical routes and adjust compatibility proxy rules only where still necessary.
4. Verify representative pages boot without `pageerror` and emit the repaired canonical request paths.
5. Record unresolved endpoint families for the later backend validation pass.

Rollback strategy: revert the specific page-family path changes and temporary proxy edits if a repaired caller regresses before backend mapping is resolved.

## Open Questions

- What is the current supported backend contract for the swap thumb/symbol-info family: direct `/swap/*`, another service alias, or an intentionally retained compatibility path?
- Does the exchange backend still expose cancel-order under another handler file that was not included in the initial scan, or does the frontend need a new explicit mapping?
