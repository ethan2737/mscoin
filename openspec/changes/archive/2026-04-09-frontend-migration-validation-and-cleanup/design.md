## Context

The frontend working tree includes artifacts such as `App.vue.original`, `App.vue.vue3-backup`, dual HTML entries, and toolchain warnings during build. The earlier OpenSpec frontend archive also lacks complete proposal and design artifacts. After tactical repairs, the project still needs a bounded cleanup and verification step to avoid drifting back into an ambiguous state.

## Goals / Non-Goals

**Goals:**
- Define objective acceptance checks for startup, build, and critical navigation flows.
- Remove or quarantine migration leftovers that are no longer needed after baseline recovery.
- Document deferred debt and remaining warnings that are acceptable versus those that block release.

**Non-Goals:**
- Performing a full product QA pass across every feature in one change.
- Large-scale refactors unrelated to stabilization.

## Decisions

1. Separate validation/cleanup from tactical bug-fix changes.
   Rationale: acceptance criteria should verify the result of earlier fixes, not change during those fixes.

2. Treat orphaned backup artifacts and stale startup files as cleanup candidates only after the active baseline is verified.
   Rationale: they may still be needed as temporary references during repair.

3. Record residual warnings explicitly instead of silently tolerating them.
   Rationale: warnings such as Sass legacy API notices may be non-blocking now but should remain visible debt.

## Risks / Trade-offs

- [Cleanup removes reference material still needed for migration work] -> Mitigation: only remove leftovers after the active baseline passes validation and preserve any deferred references intentionally.
- [Validation scope grows too large] -> Mitigation: focus on startup, build, and critical navigation journeys rather than exhaustive functional QA.

## Migration Plan

1. Define the minimum acceptance checklist after the preceding changes land.
2. Run the checklist and capture pass/fail evidence.
3. Remove or quarantine stale migration leftovers that no longer belong to the active baseline.
4. Document residual debt and close the recovery thread only if acceptance passes.

Rollback strategy: restore removed cleanup artifacts if validation reveals they are still required to operate or debug the active baseline.

## Open Questions

- Which warnings are acceptable as deferred debt for this recovery phase, and which must be eliminated before sign-off?
