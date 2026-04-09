## Context

Current auth pages combine runtime-contract failures, legacy markup, and migration cleanup debt. Register is blocked by undefined runtime keys, unresolved legacy `<marquee>`, and old message/validation patterns. Login renders but still has warnings that indicate the migration was not completed cleanly.

This change restores the auth surfaces as Vue 3 pages that are stable enough for subsequent real-backend validation.

## Goals / Non-Goals

**Goals:**
- Make login and register mount and behave without frontend runtime errors.
- Remove unsupported legacy constructs and obviously damaged text/constants.
- Ensure auth submit flows use the normalized runtime contract.
- Preserve necessary captcha/SMS/auth entry behavior needed for later integration testing.

**Non-Goals:**
- Guarantee backend acceptance of auth requests in a fully disconnected local environment.
- Redesign auth UX beyond what is necessary to restore a clean migration baseline.
- Recover every user-center screen in the same change.

## Decisions

1. Treat auth as a dedicated user-entry surface.
   Rationale: auth is the first meaningful business flow after homepage and deserves explicit recovery boundaries.
   Alternative considered: fix auth incidentally inside the runtime-contract change. Rejected because auth also has UX and markup-specific migration debt.

2. Remove unsupported legacy constructs rather than shim them indefinitely.
   Rationale: constructs like `<marquee>` are a clear sign of incomplete migration and are not worth preserving literally.
   Alternative considered: emulate legacy behavior exactly. Rejected unless product explicitly requires it later.

3. Normalize visible text and validation behavior during recovery.
   Rationale: encoding damage and malformed validation strings are user-facing defects, not optional cleanup.
   Alternative considered: postpone text cleanup until later localization work. Rejected because it leaves broken auth UX visible.

## Risks / Trade-offs

- [Auth cleanup may reveal backend contract mismatches after frontend errors are removed] -> Mitigation: accept that as the next stage and verify against backend later.
- [Removing legacy constructs may slightly alter page appearance] -> Mitigation: preserve intent and hierarchy while prioritizing functional stability.
- [Captcha dependencies can still block full automation] -> Mitigation: define frontend stability separately from end-to-end backend confirmation.

## Migration Plan

1. Update auth pages to the shared runtime contract.
2. Remove unsupported legacy markup and damaged constants/messages.
3. Stabilize form refs, validation, and submit boot paths.
4. Verify login and register render cleanly and no longer throw frontend runtime errors.

Rollback strategy: revert specific auth-page changes if they introduce broader shell regressions, but prefer iterative forward fixes.

## Open Questions

- Should the scrolling registration “recent winners” strip be preserved at all, or replaced with a simpler static/info block during recovery?
- Do auth pages need full parity with the legacy view or only functionally correct Vue 3 behavior?
