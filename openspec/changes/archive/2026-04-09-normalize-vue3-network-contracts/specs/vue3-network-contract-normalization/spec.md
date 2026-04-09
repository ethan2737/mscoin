## ADDED Requirements

### Requirement: Active Vue 3 pages do not hardcode localhost network paths
Active Vue 3 pages SHALL use the supported runtime/network contract for HTTP and realtime access. They MUST NOT hardcode `http://localhost`, `ws://localhost`, or `io('http://localhost')` for first-party application traffic.

#### Scenario: Migrated page makes an HTTP request
- **WHEN** an active Vue 3 page needs to call a first-party backend endpoint
- **THEN** it uses the supported runtime/proxy path rather than a hardcoded localhost URL

#### Scenario: Migrated page opens a realtime connection
- **WHEN** an active Vue 3 page needs a first-party realtime channel
- **THEN** it uses the supported socket/runtime path rather than a hardcoded localhost socket URL
