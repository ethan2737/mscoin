## ADDED Requirements

### Requirement: Authenticated user-center entrust history route is backend-supported
The system SHALL expose a backend-supported authenticated history route that satisfies the recovered Vue 3 user-center spot entrust history page without returning `404`.

#### Scenario: Authenticated history page requests entrust history
- **WHEN** an authenticated user opens `/#/uc/entrust/history`
- **THEN** the frontend history request resolves to a backend-supported route instead of returning `404`

### Requirement: Authenticated entrust history returns page-compatible response data
The system SHALL return a response shape that the recovered Vue 3 entrust history page can consume for rendering history rows or an empty-state result.

#### Scenario: History endpoint returns no matching orders
- **WHEN** the authenticated history route is called and no order rows exist for the user
- **THEN** the backend responds successfully with an empty result shape compatible with the frontend page

### Requirement: Repaired authenticated history flow is browser-validated
The system SHALL validate the repaired entrust history flow in a real browser using the reusable local authenticated baseline.

#### Scenario: Browser validation after route repair
- **WHEN** the endpoint contract repair is implemented
- **THEN** `/#/uc/entrust/history` loads without a `404` history request and the final outcome is classified as pass or blocked-by-data
