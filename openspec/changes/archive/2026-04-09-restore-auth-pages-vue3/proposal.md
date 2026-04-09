## Why

The audit found that auth pages are not merely blocked by backend availability; they contain migration regressions of their own. Register throws because it still expects removed runtime fields, uses unsupported legacy constructs such as `<marquee>`, and contains visibly damaged text/constants. Login renders, but it still carries non-ref-safe form wiring and migration residue.

Auth needs a dedicated change so the platform can recover a stable entry path for real users and later backend verification.

## What Changes

- Restore Vue 3 login, register, and related auth entry pages to a stable, maintainable state.
- Remove unsupported legacy constructs and encoding-damaged constants from auth surfaces.
- Align auth submission paths with the supported Vue 3 runtime contract.
- Define the minimum auth-page behavior that must work before backend-backed validation begins.

## Capabilities

### New Capabilities
- `vue3-auth-page-recovery`: Defines the minimum stable runtime and UX behavior for Vue 3 login/register/auth-entry pages.

### Modified Capabilities
- None.

## Impact

- Affected code: `mscoin-frontend/src/pages-vue3/uc/Login.vue`, `Register.vue`, `MobileRegister.vue`, `FindPwd.vue`, related auth helpers/captcha integration.
- Affected systems: login entry, registration entry, SMS/captcha flows, post-login redirect path.
- Dependency on audit findings: `openspec/changes/archive/2026-04-09-audit-frontend-business-and-data-flows/report.md`.
