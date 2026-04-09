## ADDED Requirements

### Requirement: Frontend recovery SHALL have an explicit acceptance checklist
The repaired frontend baseline MUST define a repeatable acceptance checklist covering startup, production build, and critical navigation flows.

#### Scenario: Release-readiness review
- **WHEN** maintainers evaluate whether the repaired frontend is ready for normal development
- **THEN** they can run a documented checklist that verifies startup, build, and critical navigation behavior

### Requirement: Migration leftovers SHALL be resolved after verification
Backup files, stale entry artifacts, and other migration leftovers MUST be either removed or explicitly deferred once the active baseline passes validation.

#### Scenario: Cleanup decision
- **WHEN** an artifact is no longer part of the verified active baseline
- **THEN** it is either removed from the active tree or documented as an intentional deferred reference

### Requirement: Residual debt SHALL be documented
Non-blocking warnings, deferred incompatibilities, and incomplete follow-up items MUST be recorded at the end of frontend recovery.

#### Scenario: Residual warning review
- **WHEN** the team completes the frontend recovery acceptance pass
- **THEN** any remaining warnings or deferred debt are documented with their impact and follow-up expectation
