## ADDED Requirements

### Requirement: Frontend runtime baseline SHALL be singular
The repository SHALL define exactly one supported frontend runtime baseline for active development and troubleshooting at any given time.

#### Scenario: Startup path selection
- **WHEN** maintainers inspect the active frontend configuration
- **THEN** they can identify one authoritative combination of Vue runtime, router version, store version, i18n version, and Vite plugin configuration

### Requirement: Active versus deferred artifacts SHALL be explicit
The repository SHALL label frontend boot artifacts, page trees, and config files as active or deferred according to the selected runtime baseline.

#### Scenario: Entry artifact classification
- **WHEN** a maintainer reviews entry HTML files, boot scripts, and Vite configs
- **THEN** each artifact is either part of the active startup path or explicitly documented as deferred from the active path

### Requirement: Follow-up frontend fixes SHALL depend on the selected baseline
Subsequent frontend repair changes MUST reference the selected runtime baseline instead of redefining compatibility assumptions.

#### Scenario: Downstream change planning
- **WHEN** a new frontend fix proposal is created after baseline selection
- **THEN** its scope assumes the chosen baseline and only addresses defects within that baseline
