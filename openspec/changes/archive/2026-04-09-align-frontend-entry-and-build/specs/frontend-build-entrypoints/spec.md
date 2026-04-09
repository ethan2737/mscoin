## ADDED Requirements

### Requirement: Frontend command surface SHALL be coherent
The frontend package SHALL expose development and build commands that reference existing files and target the same active runtime baseline.

#### Scenario: Build command execution
- **WHEN** a maintainer runs the documented production build command
- **THEN** the command resolves its Vite config and HTML entry without missing-file errors

### Requirement: Active entrypoint SHALL be unambiguous
The frontend SHALL designate one active HTML entry file and one active boot script for the selected baseline.

#### Scenario: Entrypoint inspection
- **WHEN** a maintainer reviews package scripts and Vite configuration
- **THEN** the active startup path maps to a single HTML entry and a single application boot module

### Requirement: Deferred startup artifacts SHALL not be referenced by active scripts
Inactive or experimental startup files MUST NOT be invoked by the active development or build commands.

#### Scenario: Deferred artifact isolation
- **WHEN** a file is marked deferred from the active baseline
- **THEN** no active package script or Vite config points to that file
