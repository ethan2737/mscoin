## 1. Market Baseline Audit

- [x] 1.1 Identify the market tables and fields required by the homepage and exchange validation flows.
- [x] 1.2 Locate any existing schema or seed assets in the repo that can serve as the local market baseline.

## 2. Schema And Seed Repair

- [x] 2.1 Restore or create the minimum missing local market tables required by current frontend validation.
- [x] 2.2 Seed visible symbol and coin metadata rows for the agreed local validation baseline.
- [x] 2.3 Verify `market-api` returns non-empty or schema-safe responses for the homepage and exchange endpoints used by the frontend.

## 3. Reporting

- [x] 3.1 Re-run homepage and exchange frontend validation against the seeded local backend.
- [x] 3.2 Publish a change-local baseline artifact describing the local market schema and seed prerequisites.
