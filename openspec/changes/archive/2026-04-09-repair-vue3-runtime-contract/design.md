## Context

The current Vue 3 frontend has no stable app-level runtime contract. `main-vue3.js` still sets global properties such as `host = 'http://localhost'` and `api = Api`, while `store-vue3.js` no longer contains equivalent state. Meanwhile, migrated pages read `store.state.host`, `store.state.api`, provided dependencies, and direct constants interchangeably.

This produces a class of failures that are orthogonal to individual page logic. Until the contract is unified, page-level fixes will continue to rediscover the same mismatch in different screens.

## Goals / Non-Goals

**Goals:**
- Define one supported source for shared runtime values in the Vue 3 frontend.
- Make migrated pages consume app-level dependencies consistently.
- Remove the most dangerous ambiguity between store state, provided objects, and hardcoded globals.
- Create a stable base for subsequent auth, exchange, and network-path changes.

**Non-Goals:**
- Fix every page that still has business logic defects.
- Complete backend integration validation.
- Redesign app architecture beyond what is necessary to stabilize the current migration baseline.

## Decisions

1. Use an explicit shared runtime surface rather than implicit globals.
   Rationale: implicit global properties make it too easy for migrated pages to depend on stale values silently.
   Alternative considered: preserve legacy globalProperties and patch pages opportunistically. Rejected because it perpetuates ambiguity.

2. Make `store-vue3` the source of durable app state and a dedicated runtime helper/module the source of configuration access.
   Rationale: app state and configuration are different concerns; pages should not infer configuration from unrelated state.
   Alternative considered: put host/api back into Vuex state. Rejected because it couples static config to mutable store state unnecessarily.

3. Scope this change to the shared contract and the highest-risk consumers.
   Rationale: the whole page tree needs normalization eventually, but this change should primarily unblock the flows currently broken by contract mismatch.
   Alternative considered: sweep every page here. Rejected because that belongs to the broader network-contract normalization change.

## Risks / Trade-offs

- [Some pages may still rely on legacy assumptions after the contract is introduced] -> Mitigation: explicitly update the highest-risk consumers and record the remaining sweep for a follow-up change.
- [Introducing a new helper surface may overlap with old globals temporarily] -> Mitigation: document one canonical path and treat legacy access as deprecated immediately.
- [Changing contract access can expose previously hidden page bugs] -> Mitigation: accept this as progress and let later targeted fixes handle surfaced page-level defects.

## Migration Plan

1. Define the canonical runtime surface for config, auth token access, and shared helper exposure.
2. Update bootstrapping to expose that surface consistently.
3. Refactor the pages identified by the audit as structurally broken because of runtime mismatch.
4. Verify that register and exchange no longer fail solely because `host/api` are missing.

Rollback strategy: restore the previous boot/runtime wiring only if the new contract prevents the app from booting at all. Page-level regressions should be fixed forward.

## Open Questions

- Should the canonical runtime helper live under `src/config/` or a dedicated `src/runtime/` module?
- Which additional pages beyond register/exchange/swap should be upgraded in this change versus left to later targeted fixes?
