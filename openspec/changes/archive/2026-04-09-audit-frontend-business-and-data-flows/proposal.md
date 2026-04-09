## Why

The current Vue 3 frontend can now boot, but it is still not verifiably correct as an application. Page shells render while key flows remain ambiguous: some regions are static placeholders, some are fallback data, some depend on backend APIs that are not running, and some migrated views are visibly broken or encoded incorrectly.

We need a system-level audit now so follow-up fixes are driven by business flow and data flow evidence rather than page-by-page guesswork.

## What Changes

- Inventory the active Vue 3 frontend by route, business flow, and data dependency.
- Classify visible regions on key pages as static content, API-backed content, websocket-backed content, or migration fallback content.
- Trace primary frontend flows including homepage discovery, language/menu interactions, login, registration, notices/help, exchange entry, and user-center entry points.
- Verify which backend services, endpoints, tokens, and realtime channels each flow requires, including expected ports and proxy behavior in local development.
- Produce a written investigation report with confirmed findings, severity, reproduction notes, and a recommended remediation sequence.

## Capabilities

### New Capabilities
- `frontend-business-data-flow-audit`: Defines how the frontend recovery effort inventories routes, maps UI regions to real data sources, verifies business flows, and records findings in a reusable investigation report.

### Modified Capabilities
- None.

## Impact

- Affected code: active Vue 3 entrypoints, `src/App.vue`, `src/config/routes-vue3.js`, `src/pages-vue3/**`, shared API config, and Vite proxy configuration.
- Affected systems: local frontend runtime, backend service dependencies on ports `8888`, `8889`, `8890`, websocket/socket.io market feeds, authentication/token flows, and investigation documentation.
- Affected outputs: an explicit audit report that distinguishes static UI, mocked/fallback UI, and real business data paths.
