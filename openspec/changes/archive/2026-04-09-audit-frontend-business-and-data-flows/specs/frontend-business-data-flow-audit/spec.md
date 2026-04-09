## ADDED Requirements

### Requirement: Audit route and flow inventory
The investigation SHALL produce an inventory of the active Vue 3 frontend routes and the primary business flows they support. For each audited route or flow, the inventory MUST identify the owning page/component entry, the user-facing purpose, and the shared shell dependencies involved in reaching it.

#### Scenario: Homepage flow is inventoried
- **WHEN** the homepage is included in the audit scope
- **THEN** the audit records the route entry, major visible regions, and the user actions that can transition into downstream flows such as notices, exchange, or login

#### Scenario: Auth flow is inventoried
- **WHEN** login and registration are included in the audit scope
- **THEN** the audit records the route entry, form components, validation path, captcha dependency, submission path, and expected post-login navigation

### Requirement: Audit data-source classification
The investigation SHALL classify every audited UI region as one of: static presentation, API-backed data, realtime-backed data, auth-gated data, or fallback/mocked data. The audit MUST identify the frontend caller, endpoint or socket path, token requirement, proxy expectation, and backend service dependency for every non-static region.

#### Scenario: API-backed section is classified
- **WHEN** an audited region renders data from an HTTP request
- **THEN** the audit identifies the request path, the frontend caller, the required auth state, and the backend service expected to answer it

#### Scenario: Fallback-backed section is classified
- **WHEN** an audited region renders local fallback or placeholder data
- **THEN** the audit marks that region as fallback-backed and explains the condition under which real data should replace it

### Requirement: Investigation report captures actionable findings
The investigation SHALL produce a report that groups findings by severity and by defect class, including frontend regression, dependency gap, and unresolved contract question. Each finding MUST include affected route or component, observed behavior, expected behavior or contract, reproduction notes, and recommended next action.

#### Scenario: Broken interaction is reported
- **WHEN** a user-visible control renders but does not behave as expected
- **THEN** the report records the control, the failed interaction, the suspected code path or dependency, and the remediation priority

#### Scenario: Missing backend dependency is reported
- **WHEN** a flow cannot be fully verified because its backend dependency is unavailable
- **THEN** the report records the missing service, endpoint or socket contract, the blocked user flow, and the evidence that the failure is not solely a frontend rendering issue
