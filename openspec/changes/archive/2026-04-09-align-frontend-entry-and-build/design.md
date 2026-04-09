## Context

Current frontend scripts do not describe a single runnable application. `dev` uses the default Vite entry, `dev:vue3` explicitly opens `index-vue3.html`, and `build:vue3` points to a config file that does not exist. This produces failures before page-level code is even evaluated.

## Goals / Non-Goals

**Goals:**
- Make development and build commands deterministic.
- Ensure every referenced config and entry file exists and belongs to the active baseline.
- Reduce the number of supported startup paths to the minimum required.

**Non-Goals:**
- Fixing all downstream component runtime issues.
- Refactoring page-level logic beyond what is needed for bootstrapping.

## Decisions

1. Define one canonical command path for dev/build/preview.
   Rationale: ambiguous commands increase diagnosis cost and cause false negatives during verification.
   Alternative considered: keep parallel Vue 2/Vue 3 scripts. Rejected unless the baseline change explicitly requires dual support.

2. Align HTML entry files, boot scripts, and Vite config as a single contract.
   Rationale: changing one without the others is what created the current broken state.
   Alternative considered: fix only the missing config path. Rejected because other script-entry mismatches would remain.

3. Treat missing referenced artifacts as release-blocking defects.
   Rationale: a broken build command makes the change unverifiable and blocks all later fixes.

## Risks / Trade-offs

- [Removing alternate scripts may reduce experimentation flexibility] -> Mitigation: keep deferred scripts documented outside the active command path.
- [Selected baseline may still contain runtime issues after build works] -> Mitigation: limit this change to command and entry coherence, then hand off runtime defects to later changes.

## Migration Plan

1. Map each package script to the entry HTML file, boot file, and Vite config it uses.
2. Remove or update references to missing or deferred artifacts.
3. Verify that the active dev and build commands both execute against the same baseline.

Rollback strategy: restore the previous script set only if the new consolidated path cannot support the selected baseline and a better baseline decision is documented.

## Open Questions

- Does the repo still need a named migration-only script after baseline selection, or should all active scripts be baseline-neutral?
