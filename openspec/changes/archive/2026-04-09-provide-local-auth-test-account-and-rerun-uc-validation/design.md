## Context

The local validation run reached the auth endpoints successfully and confirmed anonymous-state behavior, but it could not certify user-center success paths because no working local credentials were available. That makes user-center results ambiguous: the frontend may be healthy, but there is no local authenticated condition to prove it. This change makes that condition explicit and reruns the relevant flows.

## Goals / Non-Goals

**Goals:**
- Establish a known-good local login account for frontend validation.
- Confirm the minimum data attached to that account needed for selected user-center checks.
- Re-run user-center validation with authenticated state and record the outcomes.

**Non-Goals:**
- Rebuild all user-center business data comprehensively.
- Fix unrelated backend feature gaps such as help/announcement contracts.
- Replace broader end-to-end QA beyond the selected local user-center flows.

## Decisions

1. Use one explicit local validation account rather than ad-hoc personal credentials.
   Rationale: future validation needs to be repeatable by other contributors.
   Alternative considered: validate with any available local account. Rejected because it does not create a reusable baseline.

2. Validate a focused user-center slice after login succeeds.
   Rationale: the immediate blocker is authenticated validation, not exhaustive user-center coverage.
   Alternative considered: attempt every authenticated route. Rejected because many will depend on broader seed data not required for current recovery decisions.

3. Record both account prerequisites and flow outcomes in one change-local artifact.
   Rationale: the next validation run should not require rediscovering how to prepare auth state.
   Alternative considered: keep account setup separate from validation results. Rejected because it splits closely related evidence.

## Risks / Trade-offs

- [A local account may exist but lack the data needed for the selected user-center flows] -> Mitigation: document the minimum wallet/order prerequisites alongside the account setup.
- [Credentials may be sensitive or unsuitable for repo storage] -> Mitigation: record the setup approach and safe local-only expectations rather than committing secret values if necessary.
- [Authenticated flows may still fail due to backend business issues] -> Mitigation: classify those distinctly from missing-account blockers.

## Migration Plan

1. Determine how a reusable local test account should be created or identified.
2. Confirm the minimum related wallet/order data required for selected user-center validation.
3. Perform authenticated login and re-run the targeted user-center flows in a real browser.
4. Publish an updated authenticated validation artifact with pass/fail/blocked results.

Rollback strategy: remove or disable the local-only test account data if it conflicts with existing local setup, while preserving the validation notes.

## Open Questions

- Should the local validation account be created through backend seed data, through the recovered register flow, or via an existing fixture?
- Which user-center flows beyond entrust history are most valuable once authenticated validation is available?
