## Context

The repository has already recorded two contradictory frontend upgrade states. Commit `4055ada` introduced a Vue 3 plus Vite path with `pages-vue3`, while commit `6e05b70` moved the project back to Vue 2.7 plus Vite for compatibility. The current working tree again changes `package.json`, `vite.config.mjs`, `main-vue3.js`, `store-vue3.js`, and `App.vue` toward Vue 3. Because there is no active OpenSpec change and no complete archived design for the frontend migration, the team lacks a current source of truth.

## Goals / Non-Goals

**Goals:**
- Produce one explicit frontend baseline that all subsequent fixes assume.
- Identify which files are active startup/build artifacts and which are historical or experimental.
- Record acceptance criteria for moving from baseline selection into code repair.

**Non-Goals:**
- Fixing asset imports, route mismatches, or runtime bugs in this change.
- Completing the full Vue 3 migration in this document.
- Removing old code before the baseline decision is implemented and verified.

## Decisions

1. Use a baseline-first change before any tactical fix changes.
   Rationale: build and startup failures are currently ambiguous because the repo does not define whether Vue 2.7 or Vue 3 is authoritative.
   Alternative considered: fix individual errors first. Rejected because each fix may need to be redone once the runtime path is finalized.

2. Treat entrypoints, dependency versions, and router/store bootstrapping as the core baseline surface.
   Rationale: those files determine whether the app can even initialize, and they create the contract for all page-level fixes.
   Alternative considered: define baseline at component level. Rejected because components are downstream of boot configuration.

3. Preserve the non-selected path as explicitly deferred, not silently mixed in.
   Rationale: the repo may still need historical migration work for reference, but it cannot remain partially active.
   Alternative considered: immediate deletion of the non-selected path. Rejected because some migrated pages may still be useful references for later work.

## Risks / Trade-offs

- [Team selects the wrong baseline for current delivery pressure] -> Mitigation: require explicit evaluation criteria based on startup viability, dependency compatibility, and migration cost.
- [Deferred artifacts continue to leak into active build paths] -> Mitigation: identify a precise active file list and enforce it in follow-up build and routing changes.
- [Future migration work loses context] -> Mitigation: keep deferred artifacts discoverable and document why they are deferred.

## Migration Plan

1. Compare the current working tree, the Vue 3 migration commit, and the Vue 2.7 compatibility commit.
2. Choose one baseline as the supported path for current frontend recovery.
3. Mark active entry files, active dependency versions, and deferred artifact directories.
4. Feed the selected baseline into the follow-up build, asset, route, and validation changes.

Rollback strategy: if the selected baseline proves non-viable during subsequent changes, restore the previously recorded baseline decision and reopen this change with updated evidence rather than mixing both paths again.

## Open Questions

- Is the immediate objective to get the existing product running with minimal change, or to finish the Vue 3 migration now?
- Should archived Vue 3 pages remain in-tree as references if Vue 2.7 is selected, or move under a clearer quarantine path?
