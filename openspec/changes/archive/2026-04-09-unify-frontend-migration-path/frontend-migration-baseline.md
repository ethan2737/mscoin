# Frontend Migration Baseline

Last updated: 2026-04-09

## Purpose

This document establishes the single active frontend recovery baseline for the repository so all follow-up fixes use the same assumptions.

## Comparison of Candidate Baselines

### Candidate A: `4055ada` Vue 3 migration spike

- Runtime: Vue 3.5 + Vite + Element Plus
- Entrypoints:
  - `index-vue3.html`
  - `src/main-vue3.js`
  - `src/config/routes-vue3.js`
  - `src/config/store-vue3.js`
  - `src/pages-vue3/**`
- Strengths:
  - Matches the long-term modernization target of moving off Vue 2.
  - Includes the migrated page tree that current recovery work is already editing.
- Weaknesses:
  - Was introduced as a migration spike rather than the last known compatible runtime.
  - Still contains unresolved build and runtime issues.

### Candidate B: `6e05b70` Vue 2.7 compatibility rollback

- Runtime: Vue 2.7 + Vite + iView compatibility
- Entrypoints:
  - `index.html`
  - Vue 2 runtime with `@vitejs/plugin-vue2`
- Strengths:
  - Explicitly optimized for short-term compatibility with the legacy codebase.
  - Commit message states the frontend service was running on port `3000`.
- Weaknesses:
  - Reintroduces legacy compatibility constraints that this upgrade thread is trying to leave behind.
  - Does not align with the current working tree, which is already centered on Vue 3 files and `pages-vue3`.

### Candidate C: Current working tree

- Runtime intent: Vue 3 + Vite
- Current active files in the working tree:
  - `mscoin-frontend/index.html` loads `src/main-vue3.js`
  - `mscoin-frontend/package.json` uses Vue 3, Vue Router 5, Vuex 4, and vue-i18n 11
  - `mscoin-frontend/vite.config.mjs` uses `@vitejs/plugin-vue`
  - `mscoin-frontend/src/config/routes-vue3.js`
  - `mscoin-frontend/src/config/store-vue3.js`
  - `mscoin-frontend/src/pages-vue3/**`
- Strengths:
  - Matches the active recovery work already present in the branch.
  - Keeps follow-up fixes focused on the actual migration target instead of a temporary rollback.
- Weaknesses:
  - Not yet runnable because entry/build mismatches and page-level errors still exist.

## Selected Baseline

The active frontend recovery baseline for this repository is:

- Runtime: Vue 3 + Vite
- UI direction: Element Plus for migrated surfaces
- Boot path:
  - `mscoin-frontend/index.html`
  - `mscoin-frontend/src/main-vue3.js`
  - `mscoin-frontend/src/config/routes-vue3.js`
  - `mscoin-frontend/src/config/store-vue3.js`
- Page tree:
  - `mscoin-frontend/src/pages-vue3/**`

## Decision Criteria

The baseline is selected using these criteria:

1. The thread objective is frontend upgrade completion, not legacy compatibility preservation.
2. The current branch already invests heavily in Vue 3 migration files; switching back to Vue 2.7 would discard the active work surface and create a second split.
3. The next OpenSpec changes in this recovery sequence target issues inside the Vue 3 boot path and `pages-vue3` tree.
4. A single forward-moving baseline is easier to stabilize than repeatedly toggling between a rollback runtime and a migration runtime.

## Active vs Deferred Artifacts

### Active artifacts for current recovery

- `mscoin-frontend/index.html`
- `mscoin-frontend/package.json` current Vue 3 dependency set
- `mscoin-frontend/pnpm-lock.yaml`
- `mscoin-frontend/vite.config.mjs` current Vue 3 config
- `mscoin-frontend/src/main-vue3.js`
- `mscoin-frontend/src/App.vue`
- `mscoin-frontend/src/config/routes-vue3.js`
- `mscoin-frontend/src/config/store-vue3.js`
- `mscoin-frontend/src/pages-vue3/**`
- Shared supporting directories that are still imported by the Vue 3 path:
  - `mscoin-frontend/src/assets/**`
  - `mscoin-frontend/src/components/**`
  - `mscoin-frontend/src/utils/**`
  - `mscoin-frontend/src/config/api.js`

### Deferred artifacts and references

- Commit `6e05b70` as the historical Vue 2.7 rollback reference
- Legacy Vue 2 route/store files when not imported by the Vue 3 boot path:
  - `mscoin-frontend/src/config/routes.js`
  - `mscoin-frontend/src/config/store.js`
- Legacy page tree retained for reference only:
  - `mscoin-frontend/src/pages/**`
- Legacy webpack build path retained for historical reference only:
  - `mscoin-frontend/build/**`
  - `mscoin-frontend/config/**`
- Deferred legacy commands retained for historical reference only:
  - `pnpm run dev:legacy-webpack`
  - `pnpm run build:legacy-webpack`

## Removed During Validation Cleanup

The validation/cleanup change removed these migration leftovers from the repository because they were no longer part of the verified active baseline:

- `mscoin-frontend/index-vue3.html`
- `mscoin-frontend/src/App.vue.original`
- `mscoin-frontend/src/App.vue.vue3-backup`
- `mscoin-frontend/vite-dev.out.log`
- `mscoin-frontend/vite-dev.err.log`
- `mscoin-frontend/.playwright-mcp/`

## Guardrails for Follow-up Changes

- Follow-up frontend fixes in this recovery sequence MUST assume the Vue 3 + Vite baseline defined here.
- No active script, config, or route repair may reintroduce Vue 2.7 compatibility as an implicit active path.
- Deferred artifacts may remain in the repository for reference, but they must not be wired into active startup/build commands.
- If this baseline proves non-viable, the team must open a new baseline revision change instead of mixing both runtime paths again.

## Canonical Commands

Use these commands for the active baseline:

- `pnpm run dev`
- `pnpm run build`
- `pnpm run preview`

Deferred legacy commands are retained only for historical reference:

- `pnpm run dev:legacy-webpack`
- `pnpm run build:legacy-webpack`

## Validation Record

The acceptance checklist, cleanup decisions, and residual warnings for this baseline are tracked in [`frontend-recovery-validation.md`](./frontend-recovery-validation.md).

## Follow-up Change Mapping

The next frontend recovery changes depend on this baseline:

1. `align-frontend-entry-and-build`
2. `fix-frontend-assets-and-imports`
3. `align-frontend-routes-and-navigation`
4. `frontend-migration-validation-and-cleanup`
