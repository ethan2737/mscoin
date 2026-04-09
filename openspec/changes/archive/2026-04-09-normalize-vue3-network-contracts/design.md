## Context

Even after entry/build alignment and partial homepage repair, much of the Vue 3 page tree still bypasses the active runtime path. Static scanning shows many pages continue to hardcode direct localhost URLs and legacy socket paths. This undermines environment portability and prevents meaningful local/proxy-based verification.

This change performs the broad sweep that was intentionally kept out of earlier narrower fixes.

## Goals / Non-Goals

**Goals:**
- Remove hardcoded localhost network access from the active Vue 3 page tree.
- Standardize how HTTP and realtime paths are built across migrated pages.
- Reduce environment-specific failures and false negatives during QA.

**Non-Goals:**
- Rebuild every business flow end to end.
- Certify backend correctness.
- Replace broader runtime-contract work that belongs to the dedicated runtime change.

## Decisions

1. Sweep only the active Vue 3 path.
   Rationale: the baseline explicitly excludes legacy Vue 2 paths from active recovery work.
   Alternative considered: clean both `pages` and `pages-vue3`. Rejected because it expands scope without helping the active runtime.

2. Normalize network access after the shared runtime contract is defined.
   Rationale: a sweep without a target contract just replaces one inconsistency with another.
   Alternative considered: patch each page independently now. Rejected because it risks divergence.

3. Treat socket callers and HTTP callers as one network-contract problem.
   Rationale: both are currently violating the same environment abstraction boundary.
   Alternative considered: split websocket normalization into its own change. Rejected because pages commonly use both.

## Risks / Trade-offs

- [Wide sweep touches many files] -> Mitigation: restrict the sweep to active Vue 3 pages and verify representative flows after normalization.
- [Some pages may reveal endpoint-family drift once localhost is removed] -> Mitigation: record those as page-level follow-ups rather than preserving hardcoded paths.
- [Normalization may temporarily surface more backend 404/500 responses] -> Mitigation: that is preferable to silently bypassing the supported runtime path.

## Migration Plan

1. Define the supported host/proxy/socket rules from the runtime contract.
2. Sweep active Vue 3 pages for hardcoded localhost HTTP and websocket usage.
3. Replace them with the supported runtime/network access pattern.
4. Verify representative routes from each major page family still boot through the normalized path.

Rollback strategy: revert isolated page families if a sweep introduces regressions there, but keep the canonical network rules intact.

## Open Questions

- Which page families require additional proxy targets beyond `/uc`, `/market`, `/exchange`, and `/socket.io` after normalization?
- Should any page retain explicit absolute URLs for third-party services, or should those be isolated into dedicated config?
