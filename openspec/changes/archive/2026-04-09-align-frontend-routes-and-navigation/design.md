## Context

Current evidence shows that `App.vue` still links to URLs such as `/lab`, `/invite`, `/announcement/0`, `/helplist?...`, and `/about-us`, while `routes-vue3.js` defines different CMS and user-center paths and omits several of those URLs entirely. This mismatch creates runtime defects even if the build succeeds.

## Goals / Non-Goals

**Goals:**
- Make every active navigation link resolve to a reachable route under the selected baseline.
- Decide explicit behavior for legacy URLs that are still referenced in the UI or external bookmarks.
- Reduce accidental route drift between app shell and route definitions.

**Non-Goals:**
- Redesigning page IA or adding new product areas not already implied by the current app shell.
- Rewriting every page component to a new routing architecture.

## Decisions

1. Treat the route table as the canonical reachable URL contract.
   Rationale: navigation elements should point to supported routes, not define them implicitly.
   Alternative considered: preserve app-shell links as canonical and retrofit routes around them. Rejected because some current links clearly reference missing or renamed pages.

2. Resolve mismatches with one of three explicit outcomes: supported route, redirect alias, or removal.
   Rationale: this makes legacy path handling testable and prevents silent fallthrough to catch-all redirects.

3. Validate navigation in grouped flows: marketing shell, CMS, and user center.
   Rationale: route mismatches currently cluster around those areas and can be verified as coherent user journeys.

## Risks / Trade-offs

- [Preserving legacy aliases may increase route-table complexity] -> Mitigation: only keep aliases that protect real navigation paths or known deep links.
- [Removing stale links may expose missing product decisions] -> Mitigation: document each removed path and why it is out of scope.

## Migration Plan

1. Inventory all active navigation targets emitted by app shell and key pages.
2. Compare them with the active route table and classify each mismatch.
3. Implement direct routes, redirects, or link removals according to the selected baseline.
4. Verify grouped navigation flows end on the expected pages.

Rollback strategy: restore previous links only if their replacement route behavior proves incompatible with required user journeys.

## Open Questions

- Which legacy URLs, if any, must remain reachable for SEO, bookmarks, or external documentation?
