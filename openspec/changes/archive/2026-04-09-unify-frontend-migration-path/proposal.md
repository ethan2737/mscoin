## Why

The frontend upgrade currently mixes two incompatible directions: a rollback commit established Vue 2.7 plus Vite as the compatibility path, while the current working tree reintroduces Vue 3 entrypoints, dependencies, and migrated pages. Until the project selects one authoritative runtime and removes the parallel path, every startup or build investigation will be polluted by configuration drift instead of real defects.

## What Changes

- Define a single authoritative frontend migration baseline for the repository.
- Inventory and classify Vue 2, Vue 2.7 compatibility, and Vue 3 migration artifacts.
- Set repository-level rules for which entrypoints, dependencies, and component trees are active versus archival.
- Document the decision criteria and rollback boundary so follow-up fixes can build on a stable baseline.

## Capabilities

### New Capabilities
- `frontend-migration-baseline`: Establishes one supported frontend runtime path and marks all conflicting artifacts as either active, deferred, or removable.

### Modified Capabilities
- None.

## Impact

- Affected code: `mscoin-frontend/package.json`, Vite configs, entry HTML files, app boot files, and Vue 2/Vue 3 page trees.
- Affected systems: local development startup, production build selection, and follow-up frontend repair work.
- Affected dependencies: Vue, vue-router, vuex, vue-i18n, Element Plus, and Vite plugin selection.
