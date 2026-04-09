## ADDED Requirements

### Requirement: Help and announcement endpoint families are contract-classified
The system SHALL classify the current help and announcement endpoint families as missing implementation, missing mapping, or explicitly unresolved based on evidence from the frontend call sites and backend codebase.

#### Scenario: Help route returns 404 locally
- **WHEN** the validated frontend requests the local help endpoint family and receives `404`
- **THEN** the change output states whether the backend implementation is absent, present under another mapping, or still unresolved

### Requirement: Contract confirmation identifies the intended canonical route or follow-up action
The system SHALL produce either the canonical backend route mapping for help and announcement flows or a clear implementation recommendation if no backend handler exists.

#### Scenario: Announcement capability is inspected
- **WHEN** the backend and frontend contract for announcements is reviewed
- **THEN** the result includes the intended route contract or an explicit recommendation to implement the missing backend capability
