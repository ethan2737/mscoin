## ADDED Requirements

### Requirement: Frontend validation declares backend prerequisites
The frontend validation process SHALL declare the local backend services, ports, and prerequisite conditions required before business-flow validation begins.

#### Scenario: Validation starts without required backend services
- **WHEN** one or more required local backend services are unavailable
- **THEN** the validation result explicitly marks the affected flows as blocked by environment rather than silently passing or failing them

### Requirement: Validation records business-flow outcomes against local backend services
The validation process SHALL record pass/fail/blocked outcomes for the key frontend business flows exercised against local backend services.

#### Scenario: Homepage switches to live backend data
- **WHEN** homepage market, banner, or notice regions are validated with the required local services running
- **THEN** the result records whether those regions are using live backend data, fallback data, or remain blocked

#### Scenario: Exchange flow is validated
- **WHEN** exchange is exercised with the required local services running
- **THEN** the result records whether the page boots, whether market data loads, and whether realtime wiring reaches the expected backend path
