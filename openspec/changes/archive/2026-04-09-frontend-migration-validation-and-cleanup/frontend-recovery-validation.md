# Frontend Recovery Validation

Last updated: 2026-04-09

## Purpose

This document defines the minimum acceptance checklist, cleanup boundary, and residual debt for the repaired Vue 3 + Vite frontend baseline.

## Acceptance Checklist

Run the following commands from `E:\Project\web3\mscoin\mscoin-frontend`:

1. `pnpm install`
   Pass condition: dependencies resolve without changing the active runtime path away from Vue 3 + Vite.

2. `pnpm run check:asset-resolution`
   Pass condition: the command reports `Asset resolution check passed for src/pages-vue3.`

3. `pnpm run dev`
   Pass condition: Vite boots successfully and exposes a local URL without immediate startup exceptions.

4. `pnpm run build`
   Pass condition: production assets are emitted to `dist-vite/` and the build exits with code `0`.

5. `pnpm run preview`
   Pass condition: the preview server serves the SPA shell for the critical routes listed below.

## Critical Route Set

The minimum route set for recovery sign-off is:

- `/`
- `/activity`
- `/help/list?cate=0&cateTitle=新手指南`
- `/notice/item/0`
- `/about`
- `/uc/innovation/orders`

Pass condition: each route returns HTTP `200` from the preview server and includes `<div id="app"></div>` in the response body.

## Pass/Fail Criteria

The frontend recovery thread can be closed only when all of the following are true:

- The active commands remain `pnpm run dev`, `pnpm run build`, and `pnpm run preview`.
- Asset resolution passes for `src/pages-vue3/**`.
- Dev startup succeeds on the Vue 3 + Vite entry path.
- Production build succeeds on the Vue 3 + Vite entry path.
- The critical route set resolves under preview without falling through to the catch-all in unexpected places.
- No backup files or alternate entry files remain wired into the active baseline.
- Residual warnings and deferred debt are documented explicitly.

The recovery thread remains open if any of the following occur:

- A required command exits non-zero.
- A critical route returns a non-`200` status or does not serve the SPA shell.
- The active startup/build commands depend on deferred legacy files.
- Cleanup candidates are left in place without being documented as intentional debt.

## Validation Evidence

Validation run performed on 2026-04-09:

- `pnpm run check:asset-resolution`: passed
- `pnpm run dev`: passed, Vite booted successfully on a local URL
- `pnpm run build`: passed
- `pnpm exec vite preview --config vite.config.mjs --host 127.0.0.1 --port 4173 --strictPort`: passed
- Preview route checks:
  - `/` -> `200`, SPA shell present
  - `/activity` -> `200`, SPA shell present
  - `/help/list?cate=0&cateTitle=新手指南` -> `200`, SPA shell present
  - `/notice/item/0` -> `200`, SPA shell present
  - `/about` -> `200`, SPA shell present
  - `/uc/innovation/orders` -> `200`, SPA shell present

## Cleanup Decisions

Removed from the repository because they no longer belong to the verified baseline:

- `mscoin-frontend/index-vue3.html`
- `mscoin-frontend/src/App.vue.original`
- `mscoin-frontend/src/App.vue.vue3-backup`
- `mscoin-frontend/vite-dev.out.log`
- `mscoin-frontend/vite-dev.err.log`
- `mscoin-frontend/.playwright-mcp/`

Retained as deferred reference only:

- `mscoin-frontend/src/config/routes.js`
- `mscoin-frontend/src/config/store.js`
- `mscoin-frontend/src/pages/**`
- `mscoin-frontend/build/**`
- `mscoin-frontend/config/**`
- `pnpm run dev:legacy-webpack`
- `pnpm run build:legacy-webpack`

## Residual Warnings and Deferred Debt

The following items are non-blocking for this recovery phase but remain open debt:

- Sass emits `Deprecation [legacy-js-api]` warnings during build. This points to the current Sass loader/tooling chain and should be addressed in a follow-up dependency/tooling cleanup.
- Vite reports large chunk warnings for `main` and `vendor`. The build succeeds, but bundle splitting should be improved before performance-focused work or release hardening.
- The dependency tree still contains Vue 2 and webpack-era packages because the deferred legacy path remains in the repository for reference. These packages should be removed only when the team is ready to retire the legacy path completely.

## Follow-up Rule

Any future frontend work MUST start from the Vue 3 + Vite baseline described in [`frontend-migration-baseline.md`](./frontend-migration-baseline.md) and this validation record instead of re-opening the Vue 2.7 rollback path.
