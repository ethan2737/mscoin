## ADDED Requirements

### Requirement: Local market services have the minimum schema required by frontend validation
The system SHALL provide the minimum local market schema needed for homepage and exchange validation flows to execute without schema-level backend failures.

#### Scenario: Homepage or exchange requests market-backed data
- **WHEN** the local frontend calls the validated market APIs used by homepage or exchange
- **THEN** the local market backend does not fail because required market tables are missing

### Requirement: Local market baseline includes visible symbols and coin metadata
The system SHALL provide a minimum local seed dataset with visible symbols and coin metadata required for the recovered frontend to render real backend-backed market content.

#### Scenario: Homepage requests symbol trend data
- **WHEN** the frontend requests homepage market trend data against the local backend
- **THEN** the response contains backend-backed market entries rather than only an empty dataset caused by missing seed rows

### Requirement: Local market baseline is documented for repeatable validation
The system SHALL record the schema objects, seed rows, and verification endpoints required to recreate the local market validation baseline.

#### Scenario: Future contributor prepares local validation
- **WHEN** someone needs to reproduce the local market-backed frontend validation environment
- **THEN** they can follow the documented baseline without rediscovering missing schema and seed requirements
