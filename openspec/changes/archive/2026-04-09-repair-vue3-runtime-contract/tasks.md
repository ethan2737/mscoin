## 1. Contract Definition

- [x] 1.1 Define the canonical Vue 3 runtime access pattern for config, API definitions, and auth/token helpers.
- [x] 1.2 Update the boot layer and supporting modules so the canonical runtime contract is exposed consistently.

## 2. High-Risk Consumer Migration

- [x] 2.1 Refactor the pages identified by the audit as structurally broken because of missing runtime fields, starting with register and exchange/swap entry flows.
- [x] 2.2 Remove direct dependence on `store.state.host` and `store.state.api` from the consumers updated in this change.

## 3. Verification

- [x] 3.1 Verify that register no longer throws because runtime contract keys are undefined.
- [x] 3.2 Verify that exchange/swap boot paths no longer fail solely because runtime config access is missing.
