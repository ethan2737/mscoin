## 1. Script Consolidation

- [x] 1.1 Audit frontend package scripts and identify every referenced entry HTML file, boot module, and Vite config.
- [x] 1.2 Update the active script set so all referenced artifacts exist and target the selected runtime baseline.

## 2. Entrypoint Alignment

- [x] 2.1 Align the active HTML entry file, application boot file, and Vite config into one documented startup path.
- [x] 2.2 Remove or quarantine stale script references that target missing or deferred artifacts.

## 3. Verification

- [x] 3.1 Verify the active development command starts with the selected baseline.
- [x] 3.2 Verify the active production build command completes without missing-config failures.

Verification note: `pnpm run build` now completes successfully against the canonical `vite.config.mjs` path. Remaining output is limited to Sass legacy API deprecation warnings and bundle-size warnings, not missing-config or entry-alignment failures.
