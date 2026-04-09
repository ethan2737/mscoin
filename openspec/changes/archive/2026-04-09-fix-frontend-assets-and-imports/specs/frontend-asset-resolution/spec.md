## ADDED Requirements

### Requirement: Active frontend pages SHALL resolve local static assets
Every local static asset referenced by an active frontend page MUST resolve to an existing file in the active baseline.

#### Scenario: Image asset resolution
- **WHEN** the build tool processes an active page with a local image reference
- **THEN** the referenced image path resolves to an existing asset and does not fail the build

### Requirement: Active frontend pages SHALL resolve local module imports
Every relative import used by an active frontend page MUST resolve to an existing local module.

#### Scenario: Local component import resolution
- **WHEN** the build tool evaluates a relative import from an active page component
- **THEN** the import target exists and loads into the module graph without resolution errors

### Requirement: Build-blocking resource failures SHALL be detectable before runtime QA
The frontend recovery workflow MUST include a repeatable check that surfaces unresolved local assets and imports in the active baseline.

#### Scenario: Pre-QA integrity check
- **WHEN** maintainers run the documented resource-integrity check
- **THEN** they receive actionable failures for unresolved local assets or module imports before manual page testing begins
