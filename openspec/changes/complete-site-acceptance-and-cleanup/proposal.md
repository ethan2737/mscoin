## Why

The local development environment can now run the recovered core spot-trading path, but the project still lacks a site-wide acceptance pass across the remaining business domains. Several known issues remain outside the completed spot-flow recovery:

- the post-match terminal lifecycle has not been validated as a full business acceptance flow
- the announcement detail page still shows a brief `加载中...` state during content hydration
- the third-party `tawk.to` integration still adds noisy CORS failures in development
- residual style, encoding, and legacy page artifacts remain after the Vue 3 recovery
- contract, OTC, crowdfunding, and other business domains have not yet been validated end to end module by module

Without one final acceptance-focused change, the project risks stopping at "core spot flow works" while leaving other active modules partially recovered, visually degraded, or operationally unverified.

## What Changes

- Define a site-wide acceptance matrix covering spot lifecycle completion, announcement detail polish, customer-support script handling, style/encoding cleanup, and the remaining business modules.
- Validate the active modules one by one, including contract, OTC, crowdfunding, user-center adjacent flows, and any other recovered Vue 3 navigation surfaces still lacking closure.
- Repair module-level blockers that prevent a complete development-environment business walkthrough.
- Reduce or isolate non-business-blocking noise, such as third-party `tawk.to` failures, so local verification focuses on product behavior instead of avoidable console pollution.
- Publish a final full-site acceptance report that classifies each module as passed, partially passed, blocked by backend data, or deferred by scope.

## Capabilities

### New Capabilities

- `site-wide-acceptance-and-cleanup`: validates and closes the remaining active business domains and polish gaps required for a complete local development acceptance pass.

### Modified Capabilities

- None.

## Impact

- `mscoin-frontend/src/pages-vue3/**` active domain pages, shared layout, content pages, and residual migrated UI surfaces
- `mscoin-backend/**` services and local seed/runtime dependencies as needed to support end-to-end acceptance of remaining domains
- local browser-led QA artifacts and the final acceptance report under `openspec/changes/complete-site-acceptance-and-cleanup/`
