## Context

The project has progressed from a broken mixed migration state to a locally usable baseline with a verified core spot-trading path. That work restored entry/build integrity, runtime contracts, endpoint normalization, local login, market data baselines, content APIs, and authenticated spot order visibility. However, the system still lacks a deliberate site-wide acceptance phase across all remaining active domains.

The remaining work is not a single bug. It is a cross-module closure pass with three kinds of risk:

- business-flow coverage gaps: some domains such as contract, OTC, crowdfunding, and other active modules have not yet been verified end to end
- polish and migration residue: loading flashes, style drift, encoding artifacts, and historical leftovers still weaken the recovered frontend
- environment-noise contamination: third-party integrations and non-critical console failures make local acceptance harder to interpret

The goal of this change is to turn the recovered local environment into a module-by-module accepted development baseline, not merely a collection of individually repaired pages.

## Goals / Non-Goals

**Goals:**
- Validate the remaining active business domains one by one using a browser-led acceptance matrix.
- Close the spot trading lifecycle gap enough to classify the local environment as passing or still blocked for terminal lifecycle verification.
- Remove or isolate the `tawk.to` local CORS noise when it interferes with acceptance clarity.
- Repair visible announcement-detail loading behavior that remains after the content API recovery.
- Clean up residual style, encoding, and legacy-page artifacts that still affect the active Vue 3 experience.
- Publish a final full-site acceptance report with module status, evidence, residual blockers, and explicitly deferred areas.

**Non-Goals:**
- Rebuild inactive or intentionally deferred legacy Vue 2 surfaces that are outside the active Vue 3 baseline.
- Eliminate every warning or console message if it does not affect business validation or active page quality.
- Create production-grade fixture completeness for every module when a smaller local acceptance baseline is sufficient.
- Expand scope into unrelated feature work beyond restoring and accepting the existing business modules.

## Decisions

1. Validate by module, not by technical layer.
   Rationale: the remaining risk is user-visible incompleteness across domains, so acceptance should follow business modules such as spot, contract, OTC, crowdfunding, content, and user center rather than splitting work purely into frontend/backend buckets.
   Alternatives considered:
   - Finish remaining issues as unrelated bug tickets. Rejected because it would lose the global acceptance view.
   - Run one generic smoke test only. Rejected because it would not expose module-specific blockers.

2. Treat third-party noise as an acceptance concern only when it degrades local verification.
   Rationale: `tawk.to` is not core business logic, but repeated CORS/load failures pollute debugging and can obscure real errors during acceptance.
   Alternatives considered:
   - Ignore all third-party noise. Rejected because it weakens final QA signal.
   - Remove all third-party integrations outright. Rejected because a local-only guard or graceful disable is usually enough.

3. Preserve explicit status categories in the final report.
   Rationale: the project needs a reliable closure artifact that distinguishes "fully accepted", "accepted with local limitations", "blocked by missing backend capability", and "deferred by scope".
   Alternatives considered:
   - Binary pass/fail only. Rejected because it hides important nuance.
   - Leave status implicit in prose. Rejected because it is hard to action.

4. Fix only what is necessary to make the active development baseline coherent.
   Rationale: this is a closure pass, not a new product phase. The work should target blockers, visible migration residue, and acceptance ambiguity.
   Alternatives considered:
   - Comprehensive redesign/polish sweep. Rejected because it exceeds the stated scope.
   - Report-only audit. Rejected because the user asked to make the remaining business domains done module by module.

## Risks / Trade-offs

- [Remaining domain flows may expose new backend or data gaps] -> Mitigation: classify each domain clearly and repair only the blockers required for active-baseline closure.
- [The matching terminal lifecycle may require more infrastructure than a local acceptance pass should own] -> Mitigation: define the minimum acceptable evidence for lifecycle closure and document any deeper matching limitations explicitly.
- [Style/encoding cleanup can sprawl across many pages] -> Mitigation: restrict cleanup to active Vue 3 pages and visible acceptance issues.
- [Silencing `tawk.to` locally could diverge from production behavior] -> Mitigation: use environment-aware disabling or graceful failure handling rather than deleting the integration globally.

## Migration Plan

1. Build the site-wide acceptance matrix for all active modules and classify known gaps.
2. Repair presentation and environment-noise issues that currently distort acceptance results.
3. Validate and repair the remaining business domains one by one, including contract, OTC, crowdfunding, content detail, and any still-active user-center surfaces.
4. Revisit the spot trading lifecycle and classify the terminal-state coverage achieved in local development.
5. Publish a final full-site acceptance report with evidence and residual deferred items.

Rollback strategy:
- Revert local-only guards or cleanup changes if they unintentionally alter intended production behavior.
- Preserve the final acceptance report even if some fixes are rolled back, so the uncovered gaps remain documented.

## Open Questions

- What is the minimum local evidence required to count the matching terminal lifecycle as accepted: current/history visibility, cancellation, partial fill, full fill, or all of them?
- Which remaining active routes should be considered in-scope if some business domains are present in navigation but still depend on missing local infrastructure?
- Should third-party chat be disabled only in development, or should the integration be made resilient across all environments?
