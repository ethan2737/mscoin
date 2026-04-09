## ADDED Requirements

### Requirement: Local help and announcement pages have backend-supported content contracts
The system SHALL expose backend-supported local help and announcement routes that satisfy the active recovered Vue 3 CMS and homepage content pages without returning `404`.

#### Scenario: Help page family is loaded in local development
- **WHEN** the frontend requests the active local help routes used by `/#/help`, `/#/help/list`, and `/#/help/detail`
- **THEN** the backend returns successful page-compatible responses instead of `404`

#### Scenario: Announcement page family is loaded in local development
- **WHEN** the frontend requests the active local announcement routes used by the homepage and notice pages
- **THEN** the backend returns successful page-compatible responses instead of `404`

### Requirement: Local market symbol thumb contract is stable
The system SHALL return a successful local `market/symbol-thumb` response that the recovered Vue 3 pages can consume without backend `500` failures.

#### Scenario: Symbol thumb is requested for local market-dependent pages
- **WHEN** the frontend or validation tooling requests `POST /market/symbol-thumb` in the local environment
- **THEN** the backend returns `code: 0` with a frontend-compatible symbol collection payload

### Requirement: Local trading lifecycle is executable end to end
The system SHALL provide the minimum local market, wallet, and order baseline required to execute and observe a real spot trading lifecycle in development.

#### Scenario: Local user submits a spot order
- **WHEN** a valid local authenticated user submits a supported spot order through the recovered frontend
- **THEN** the backend accepts the order through the normal exchange route and returns a successful order submission response

#### Scenario: Submitted order is visible in user-center views
- **WHEN** a local order has been submitted successfully
- **THEN** the order is observable through the relevant current and/or history entrust views using backend-supported routes

### Requirement: Development environment flow closure is browser-validated
The system SHALL validate the repaired local environment through a real browser and classify whether the complete business flow is runnable in development.

#### Scenario: Full local business flow validation is executed
- **WHEN** the local recovery work is complete
- **THEN** a browser-led validation covers login, content pages, market visibility, order submission, and entrust visibility

#### Scenario: Final local validation artifact is published
- **WHEN** the browser-led local validation finishes
- **THEN** the change publishes a report that clearly states whether the development environment can run the complete target business flow and lists any residual non-blocking issues
