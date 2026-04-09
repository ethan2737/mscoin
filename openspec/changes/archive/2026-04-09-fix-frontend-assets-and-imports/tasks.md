## 1. Resource Audit

- [x] 1.1 Scan active frontend pages for unresolved relative imports and local asset references.
- [x] 1.2 Classify each failure as missing file, wrong relative path, or stale reference to deferred code.

## 2. Resource Repair

- [x] 2.1 Fix build-blocking local asset references using existing baseline-consistent assets where possible.
- [x] 2.2 Fix unresolved relative imports so active pages load into the module graph cleanly.

## 3. Verification

- [x] 3.1 Re-run the active production build until resource resolution errors are cleared.
- [x] 3.2 Record any deferred-path exclusions so later changes do not mistake them for active defects.

Deferred-path exclusions for this change:
- Active scanning and `check:asset-resolution` intentionally cover `mscoin-frontend/src/pages-vue3/**` only.
- Deferred legacy/reference paths excluded from this change: `mscoin-frontend/src/pages/**`, `mscoin-frontend/build/**`, `mscoin-frontend/config/**`, `mscoin-frontend/src/config/routes.js`, `mscoin-frontend/src/config/store.js`, and `mscoin-frontend/index-vue3.html`.
