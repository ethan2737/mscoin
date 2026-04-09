## Why

The frontend cannot be built or started reliably because its entrypoint and build script surface is internally inconsistent. `index.html`, `index-vue3.html`, `dev`, `dev:vue3`, `build:vue3`, and Vite config references currently describe different startup paths, including a missing `vite.config.vue3.mjs`.

This change assumes the Vue 3 + Vite recovery baseline selected by `unify-frontend-migration-path`.

## What Changes

- Consolidate frontend startup and build scripts around the active baseline selected in the preceding baseline change.
- Define the supported HTML entry file, boot script, and Vite config contract for local development and production builds.
- Remove or quarantine stale script references that point to missing or inactive artifacts.
- Document the expected commands for development, build, and preview.

## Capabilities

### New Capabilities
- `frontend-build-entrypoints`: Defines one coherent set of frontend scripts, HTML entry files, and Vite configs for the supported runtime path.

### Modified Capabilities
- None.

## Impact

- Affected code: `mscoin-frontend/package.json`, `index.html`, `index-vue3.html`, Vite config files, and boot scripts such as `main-vue3.js`.
- Affected systems: local `pnpm run dev`, production build invocation, and preview workflow.
