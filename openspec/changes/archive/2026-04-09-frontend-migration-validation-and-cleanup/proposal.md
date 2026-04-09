## Why

Even after baseline, build, asset, and route issues are fixed, the frontend upgrade will remain risky unless the team has a repeatable acceptance pass and cleanup boundary. The current migration already contains backup files, mixed entry artifacts, Sass deprecation noise, and incomplete OpenSpec records, so a final validation and cleanup change is needed before declaring the frontend stable.

This change assumes the Vue 3 + Vite recovery baseline selected by `unify-frontend-migration-path`.

## What Changes

- Define the verification checklist required to accept the repaired frontend baseline.
- Identify removable migration leftovers such as stale backups or unused entry artifacts once the baseline is verified.
- Record remaining warnings, deferred debt, and the conditions for closing the frontend recovery thread.
- Capture the post-repair documentation needed so future work does not repeat the same ambiguity.

## Capabilities

### New Capabilities
- `frontend-migration-verification`: Establishes the final acceptance, cleanup, and documentation contract for the repaired frontend baseline.

### Modified Capabilities
- None.

## Impact

- Affected code: frontend backup files, deferred configs, warning-producing style/tooling settings, and OpenSpec/frontend documentation.
- Affected systems: release readiness, QA handoff, and future migration planning.
