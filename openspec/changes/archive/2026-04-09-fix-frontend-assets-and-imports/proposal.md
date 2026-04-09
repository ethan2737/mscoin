## Why

After the entry/build layer is aligned, the next blocking failures are broken static asset references and invalid module imports inside migrated pages. Current evidence already shows a missing `banner.jpg` reference in `pages-vue3/otc/carousel.vue` and a broken local import in `pages-vue3/swap/Kline.vue`.

This change assumes the Vue 3 + Vite recovery baseline selected by `unify-frontend-migration-path`.

## What Changes

- Audit active-baseline Vue pages for broken relative asset paths and unresolved local imports.
- Repair missing references or replace them with valid assets/modules that match the selected baseline.
- Establish a repeatable check for unresolved local resources before runtime QA begins.

## Capabilities

### New Capabilities
- `frontend-asset-resolution`: Ensures active frontend pages resolve local static assets and local module imports without build-time failures.

### Modified Capabilities
- None.

## Impact

- Affected code: active Vue pages under `src/pages` or `src/pages-vue3`, local static assets under `src/assets`, and relative imports in component scripts/styles.
- Affected systems: Vite production build, module graph resolution, and image/style rendering.
