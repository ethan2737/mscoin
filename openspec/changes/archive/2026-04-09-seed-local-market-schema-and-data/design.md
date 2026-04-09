## Context

The latest local validation run showed that the frontend now reaches local market services successfully, but those services return empty or failing responses because the underlying market schema is incomplete. Specifically, `market.exchange_coin` and `market.coin` are missing, which prevents homepage and exchange flows from switching from empty placeholders to real backend-driven data.

## Goals / Non-Goals

**Goals:**
- Identify the minimum market tables and fields required by the currently validated frontend flows.
- Restore or create those schema objects locally.
- Seed enough market rows for homepage and exchange views to return meaningful non-empty responses.
- Produce a repeatable local seed baseline for future frontend validation.

**Non-Goals:**
- Build a full production-like market dataset.
- Resolve help/announcement API gaps.
- Resolve auth/test-account setup outside what is needed to verify market-backed flows.

## Decisions

1. Seed only the minimum market baseline required by active Vue 3 validation.
   Rationale: the user needs deterministic local verification, not a full production dataset.
   Alternative considered: restore the full historical schema and all rows. Rejected because it increases scope and setup time without improving current validation coverage.

2. Treat homepage and exchange as the acceptance targets for market seeding.
   Rationale: these are the two validated frontend surfaces currently blocked by missing market schema.
   Alternative considered: seed every market-adjacent API. Rejected because it dilutes focus and complicates verification.

3. Record schema and seed prerequisites as a durable local baseline artifact.
   Rationale: this problem will recur unless the expected local data contract is explicit.
   Alternative considered: rely on ad-hoc local database state. Rejected because it makes future validation non-repeatable.

## Risks / Trade-offs

- [Seed data may not match backend logic assumptions] -> Mitigation: verify through actual `market-api` responses used by the frontend.
- [Schema restoration may reveal additional missing tables beyond the initial two] -> Mitigation: expand the documented minimum baseline only when a validated flow actually depends on it.
- [Local data may go stale] -> Mitigation: keep the baseline intentionally small and document how to refresh it.

## Migration Plan

1. Audit the market service queries currently exercised by homepage and exchange validation.
2. Restore or create the minimum required local tables.
3. Insert baseline visible symbols and coin metadata rows.
4. Re-run homepage and exchange validation against the seeded local backend.
5. Publish a local market baseline artifact with schema and seed expectations.

Rollback strategy: remove only the newly added local seed rows or revert the local schema changes if they prove incompatible with existing local backend logic.

## Open Questions

- Are there existing SQL migration or seed assets elsewhere in the repo that should be used instead of creating a new baseline?
- Which exact symbol set should be considered the canonical local validation baseline: only `BTC/USDT`, or a broader homepage list?
