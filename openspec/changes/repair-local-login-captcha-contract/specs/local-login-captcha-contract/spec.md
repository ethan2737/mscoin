## ADDED Requirements

### Requirement: Local recovered login flow has a backend-compatible captcha contract
The system SHALL accept the recovered frontend fallback captcha request shape for local login validation when the request explicitly uses the documented local captcha mode.

#### Scenario: Fallback captcha login request is submitted locally
- **WHEN** the recovered Vue 3 login page submits `POST /uc/login` with the documented local fallback captcha payload
- **THEN** the backend accepts that captcha contract without requiring external Vaptcha verification

### Requirement: Local fallback captcha preserves the normal login success path
The system SHALL issue the normal authenticated login response after local captcha acceptance if the seeded member credentials are valid.

#### Scenario: Local validation account logs in successfully
- **WHEN** a valid local member submits correct credentials through the repaired login flow
- **THEN** `/uc/login` returns a successful login response with the normal token and member payload

### Requirement: Repaired local login flow is browser-validated
The system SHALL validate the repaired login flow in a real browser against the recovered frontend page.

#### Scenario: Browser login rerun after contract repair
- **WHEN** the local captcha contract repair is implemented
- **THEN** a real-browser run through `/#/login` completes authenticated navigation without relying on manual token injection
