## Context

At audit time, only the frontend dev server was running. This prevented full verification of API-backed and realtime-backed flows. The earlier recovery changes and the new repair changes are still only a partial answer until the frontend is exercised against the expected local backend services.

This change is intentionally validation-focused. It should happen after the structural frontend defects that currently invalidate the flows on their own have been addressed.

## Goals / Non-Goals

**Goals:**
- Define the required local backend readiness conditions for meaningful frontend verification.
- Validate the key frontend business flows against actual local services.
- Record which previously blocked flows now pass, fail, or remain environment-dependent.
- Produce a durable validation artifact for the recovery thread.

**Non-Goals:**
- Implement the frontend fixes themselves.
- Debug arbitrary backend business issues beyond what is needed to classify validation outcomes.
- Replace automated or manual backend service startup documentation.

## Decisions

1. Validate only after the structural frontend defects are repaired.
   Rationale: otherwise the results remain dominated by frontend-only failures.
   Alternative considered: validate immediately with broken pages. Rejected because the signal is too noisy.

2. Use business flows as validation units.
   Rationale: the user needs to know whether the product behaves correctly, not merely whether individual endpoints answer.
   Alternative considered: endpoint smoke tests only. Rejected because UI and handoff behavior matter.

3. Treat backend readiness as an explicit prerequisite.
   Rationale: the audit already showed that missing local services can masquerade as frontend failure.
   Alternative considered: silently skip missing services. Rejected because it weakens the validation result.

## Risks / Trade-offs

- [Backend services may still be unavailable or partially unhealthy] -> Mitigation: classify results clearly as blocked-by-environment versus frontend failure.
- [Validation may uncover additional backend contract drift] -> Mitigation: record it explicitly rather than hiding it under generic frontend failure.
- [Some flows require authenticated test data] -> Mitigation: define prerequisite accounts/data where necessary.

## Migration Plan

1. Confirm local backend services and ports required for validation.
2. Run flow-based validation across homepage, auth, notices/help, exchange, and selected user-center paths.
3. Record outcomes, including fallback regions that now switch to live data.
4. Publish the validation result as a durable recovery artifact.

Rollback strategy: not applicable; this is a validation/reporting change.

## Open Questions

- Which authenticated test accounts or seeded market data are available locally for validation?
- Should websocket validation require live market traffic or only a successful handshake/subscription path?
