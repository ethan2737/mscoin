## ADDED Requirements

### Requirement: Vue 3 pages use one canonical runtime contract
The frontend SHALL expose one canonical runtime contract for migrated Vue 3 pages to access configuration, API definitions, and auth-related helpers. Migrated pages MUST NOT depend on missing `store.state.host` or `store.state.api` fields.

#### Scenario: Migrated page needs API config
- **WHEN** a migrated Vue 3 page needs an endpoint definition or base-path helper
- **THEN** it reads that value from the canonical runtime contract rather than from removed Vuex fields or hardcoded globals

### Requirement: Boot and store layers are contract-consistent
The Vue 3 boot layer and store layer SHALL agree on the supported app-level dependency surface. The app MUST NOT boot into a state where core migrated pages fail because expected runtime keys are undefined.

#### Scenario: Register page mounts
- **WHEN** the register page mounts under the supported Vue 3 boot path
- **THEN** it does not throw solely because `host`, `api`, or equivalent runtime dependencies are undefined

#### Scenario: Exchange page mounts
- **WHEN** the exchange page mounts under the supported Vue 3 boot path
- **THEN** it does not fail solely because it is reading configuration from removed store fields
