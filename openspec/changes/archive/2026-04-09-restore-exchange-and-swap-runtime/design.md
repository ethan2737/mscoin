## Context

The audit shows exchange currently fails before real business validation begins. Errors include null layout reads, missing prop expectations, and legacy socket wiring such as `io('http://localhost')`. Swap shares much of the same migration heritage and is likely to fail for the same reasons.

This change focuses on making exchange and swap mount safely, render key regions, and consume the normalized runtime contract introduced by the preceding runtime-contract change.

## Goals / Non-Goals

**Goals:**
- Make exchange and swap pages mount without fatal runtime errors.
- Restore market-list, depth/trade panels, and basic pair boot flow to a verifiable state.
- Normalize realtime wiring for these surfaces.
- Make the core trading pages ready for backend-backed validation.

**Non-Goals:**
- Full trading business certification against live backend services.
- Broader OTC or user-center fixes.
- Final performance or chart polish work.

## Decisions

1. Prioritize render safety before data completeness.
   Rationale: a trading surface that crashes on mount cannot be meaningfully integrated with backend services.
   Alternative considered: start with backend/socket validation first. Rejected because the page currently fails too early.

2. Treat exchange and swap together where they share runtime patterns.
   Rationale: both flows rely on similar market boot, pair selection, and realtime assumptions.
   Alternative considered: split spot and swap immediately. Rejected because the first recovery step is the same class of runtime stabilization.

3. Verify a minimum interactive surface rather than every trading branch.
   Rationale: the immediate goal is to restore a safe, inspectable surface; deeper order-flow validation belongs to the backend-validation change.
   Alternative considered: validate full order lifecycle now. Rejected as too dependent on backend state.

## Risks / Trade-offs

- [Stabilizing render paths may reveal hidden data-shape mismatches next] -> Mitigation: document these separately and continue forward into backend validation.
- [Spot and swap fixes may touch shared components with side effects] -> Mitigation: keep changes localized and verify both route families after each shared adjustment.
- [Partial recovery may still leave some order actions disabled] -> Mitigation: explicitly define the minimum acceptable interactive state in this change.

## Migration Plan

1. Normalize runtime and socket access for exchange and swap pages.
2. Fix mount-time null access and missing-prop assumptions.
3. Restore safe rendering of market list, pair summary, and depth/trade surfaces.
4. Verify core route entry for both spot and swap.

Rollback strategy: revert specific exchange/swap stabilization edits if they introduce broader routing or shell regressions, but prefer fixing forward.

## Open Questions

- Which shared charting/depth components are safe to repair in place versus wrap defensively?
- What is the minimum order-form behavior required before this change is considered done?
