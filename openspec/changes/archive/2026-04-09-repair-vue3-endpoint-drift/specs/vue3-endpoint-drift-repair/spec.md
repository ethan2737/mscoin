## ADDED Requirements

### Requirement: Active Vue 3 pages use confirmed backend endpoint names
The system SHALL ensure that active Vue 3 pages call confirmed backend endpoint names instead of stale legacy aliases when the backend route has been identified from the current backend codebase.

#### Scenario: Market caller uses canonical market route
- **WHEN** an active Vue 3 page requests market symbol or coin metadata
- **THEN** it uses the confirmed current market route name instead of a stale legacy `/api/market/*` alias

### Requirement: Unresolved endpoint families use one explicit compatibility strategy
The system SHALL not leave active Vue 3 pages split across multiple accidental aliases for the same unresolved endpoint family. If a current backend route is still unconfirmed, the change SHALL document and preserve one explicit compatibility strategy until validation resolves it.

#### Scenario: Swap family remains partially unresolved
- **WHEN** the implementation cannot yet confirm the current backend route for a swap endpoint family
- **THEN** the frontend keeps one explicit documented mapping for that family instead of allowing page-specific drift

### Requirement: Endpoint drift repair produces a mapping record for validation
The system SHALL record the repaired frontend-to-backend endpoint mapping and any remaining unresolved drifts so later local-backend validation can verify the canonical contract.

#### Scenario: Repair change completes
- **WHEN** the endpoint drift repair work is finished
- **THEN** the change output includes a record of repaired endpoint mappings and any remaining unconfirmed endpoint families
