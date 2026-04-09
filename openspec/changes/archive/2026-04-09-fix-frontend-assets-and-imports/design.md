## Context

Once Vite reaches page transformation, it currently fails on missing page resources. The first confirmed defect is a missing banner image referenced by the OTC carousel component. Static scanning also shows unresolved relative imports that will block the module graph even if the first asset issue is fixed.

## Goals / Non-Goals

**Goals:**
- Eliminate build-blocking local asset and import resolution errors in the active frontend path.
- Ensure each active relative resource reference either exists or is intentionally replaced.
- Create a quick audit loop that can be rerun after fixes.

**Non-Goals:**
- Redesigning page visuals or replacing assets for aesthetic reasons.
- Fixing business logic defects unrelated to resource resolution.

## Decisions

1. Scope the audit to the selected active baseline only.
   Rationale: deferred artifacts may remain broken by design and should not block current recovery.
   Alternative considered: scan the entire repository. Rejected because it would mix inactive historical code into current acceptance.

2. Prefer correcting relative paths or replacing with existing canonical assets before introducing new files.
   Rationale: the migration already contains multiple overlapping asset trees, and adding more files can deepen confusion.

3. Treat unresolved local imports and static asset URLs as the same class of build integrity defect.
   Rationale: both break the module graph and should be cleared before runtime navigation testing.

## Risks / Trade-offs

- [A fixed path may still point to an obsolete visual asset] -> Mitigation: verify the page renders acceptably during the later QA change.
- [Inactive files may still contain broken references] -> Mitigation: limit enforcement to the active baseline and document exclusions.

## Migration Plan

1. Scan active pages for unresolved relative imports and local asset references.
2. Fix or replace missing targets with baseline-consistent resources.
3. Rebuild until Vite completes resource resolution for the active path.

Rollback strategy: revert individual path corrections that introduce worse regressions and replace them with safer canonical assets.

## Open Questions

- Are there baseline-specific asset directories that should be normalized after the immediate recovery work succeeds?
