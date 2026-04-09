## Why

The root app shell and the Vue 3 route table currently disagree on multiple paths, including renamed CMS routes and missing feature routes. Even after the app boots, users will still hit redirects, dead links, or the wrong page if the route contract is not synchronized.

This change assumes the Vue 3 + Vite recovery baseline selected by `unify-frontend-migration-path`.

## What Changes

- Define the supported route contract for the active frontend baseline.
- Align app-shell navigation targets with the active route table.
- Decide how legacy or renamed URLs are handled: preserve, redirect, or remove.
- Identify route-level acceptance criteria for header, footer, and user-center navigation.

## Capabilities

### New Capabilities
- `frontend-route-contract`: Ensures active navigation links and the active route table describe the same reachable URL set.

### Modified Capabilities
- None.

## Impact

- Affected code: `src/App.vue`, route config files, page-level navigation links, and any compatibility redirects.
- Affected systems: browser navigation, deep links, header/footer menus, and user-center child routes.
