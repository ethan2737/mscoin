## ADDED Requirements

### Requirement: Register send-code uses fixed verification code for local testing
The system SHALL write the fixed verification code `123456` to the registration Redis cache when handling the local registration send-code flow.

#### Scenario: Send code stores fixed verification code
- **WHEN** the client calls `/uc/mobile/code` with a valid phone number during local testing
- **THEN** the backend SHALL write `123456` into the `REGISTER::<phone>` cache entry instead of generating a random code

### Requirement: Register submit validates against the fixed cached code
The system SHALL continue to validate registration requests against the code stored in the registration Redis cache so that the fixed local test code remains part of the normal registration flow.

#### Scenario: Registration succeeds with fixed verification code
- **WHEN** the client first calls `/uc/mobile/code` for a phone number and then submits `/uc/register/phone` with code `123456`
- **THEN** the backend SHALL read the cached value from `REGISTER::<phone>` and treat the verification-code check as passed

#### Scenario: Registration fails with non-matching verification code
- **WHEN** the client first calls `/uc/mobile/code` for a phone number and then submits `/uc/register/phone` with a code other than `123456`
- **THEN** the backend SHALL reject the request with a verification-code mismatch error
