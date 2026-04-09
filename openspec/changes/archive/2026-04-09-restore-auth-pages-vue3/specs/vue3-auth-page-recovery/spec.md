## ADDED Requirements

### Requirement: Login and register pages mount without frontend runtime failures
The Vue 3 auth entry pages SHALL render without fatal runtime exceptions caused by missing runtime contract fields, unsupported legacy constructs, or malformed migration artifacts.

#### Scenario: Register route opens
- **WHEN** the user navigates to `/register`
- **THEN** the page renders without throwing because `api`, `host`, or equivalent auth dependencies are undefined

#### Scenario: Login route opens
- **WHEN** the user navigates to `/login`
- **THEN** the page renders without fatal frontend runtime exceptions and shows the expected form controls

### Requirement: Auth pages use supported Vue 3 patterns
Recovered auth pages SHALL use supported Vue 3-compatible markup and form behavior instead of unresolved legacy constructs that break render or maintenance.

#### Scenario: Register page contains legacy marquee strip
- **WHEN** the register page needs to present promotional or informational content
- **THEN** it uses a supported Vue 3-compatible pattern rather than unresolved legacy-only markup
