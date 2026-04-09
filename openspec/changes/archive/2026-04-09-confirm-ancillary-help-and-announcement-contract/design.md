## Context

The recovered frontend now reaches local backend services through the intended proxy path, which made the help and announcement failures easier to classify. Those pages are no longer blocked by frontend runtime defects; they are blocked because the expected endpoints return `404`. We need a contract-confirmation change before deciding whether to implement backend handlers, re-map frontend routes, or both.

## Goals / Non-Goals

**Goals:**
- Confirm the intended backend ownership for help and announcement flows.
- Determine whether the current local backend already implements those capabilities under different routes.
- Produce a clear contract decision that unblocks either frontend mapping repair or backend implementation work.

**Non-Goals:**
- Implement the full help or announcement feature in this change.
- Seed market data or auth test accounts.
- Fix unrelated CMS rendering issues outside contract confirmation.

## Decisions

1. Treat `404` as a contract-classification problem before treating it as a coding problem.
   Rationale: if the capability moved services or routes, implementing duplicate handlers would be the wrong fix.
   Alternative considered: immediately implement the missing endpoints. Rejected because the root cause may be wrong mapping rather than absent logic.

2. Use both frontend call sites and backend handler search as evidence.
   Rationale: the contract question spans both sides of the boundary.
   Alternative considered: infer solely from frontend expectations. Rejected because it cannot distinguish missing implementation from moved ownership.

3. End with an explicit next-step recommendation.
   Rationale: the user needs a decision, not just another list of findings.
   Alternative considered: publish only a neutral audit. Rejected because it would not unblock follow-up work.

## Risks / Trade-offs

- [Backend codebase may have partial or renamed ancillary features that are hard to locate] -> Mitigation: search by both route patterns and domain terms, then record uncertainty explicitly if it remains.
- [Frontend may be calling historically valid routes that no longer fit the current backend architecture] -> Mitigation: classify that as mapping drift rather than preserving the old contract by default.
- [A missing implementation result may expand into broader backend work] -> Mitigation: keep this change focused on contract confirmation and leave implementation to a dedicated follow-up change if needed.

## Migration Plan

1. Enumerate the exact frontend requests used by help and announcement pages.
2. Search the current backend codebase for equivalent handler implementations or alternate route mappings.
3. Classify each family as missing implementation, missing mapping, or ambiguous.
4. Publish a change-local contract report with the recommended follow-up action.

Rollback strategy: not applicable; this is a contract-confirmation and reporting change.

## Open Questions

- If announcement and help capabilities were intentionally removed from the Go backend, where should those content flows live now?
- If the routes exist in another service or deployment slice, should the frontend be remapped locally or should the local backend bootstrap include that service?
