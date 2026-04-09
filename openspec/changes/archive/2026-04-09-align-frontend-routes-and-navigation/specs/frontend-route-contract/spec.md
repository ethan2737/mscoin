## ADDED Requirements

### Requirement: Active navigation links SHALL target reachable routes
Every navigation target rendered in the active app shell MUST resolve to a supported route or a documented redirect alias.

#### Scenario: Header navigation click
- **WHEN** a user clicks an active header or footer navigation link
- **THEN** the frontend navigates to a reachable destination instead of falling through to an unintended catch-all redirect

### Requirement: Legacy route handling SHALL be explicit
Any legacy URL still referenced by the active frontend MUST be either supported directly, redirected intentionally, or removed from the active UI.

#### Scenario: Legacy CMS path evaluation
- **WHEN** the active UI or a known deep link references a renamed CMS path
- **THEN** the path behavior is explicitly defined as direct support, redirect alias, or removal

### Requirement: Route definitions and app-shell navigation SHALL stay synchronized
The active route table and active navigation configuration MUST describe the same supported URL contract.

#### Scenario: Route contract review
- **WHEN** maintainers compare the active route table with app-shell navigation targets
- **THEN** every active target has matching route behavior in the route configuration
