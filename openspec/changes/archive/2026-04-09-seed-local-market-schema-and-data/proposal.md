## Why

Local-backend validation proved that the recovered frontend can now reach the market service, but the market service cannot provide meaningful data because required tables and seed rows are missing. Until the local market schema and baseline行情 data are restored, homepage and exchange verification remain blocked by environment state rather than frontend correctness.

## What Changes

- Restore or create the local market schema objects required by `market.rpc` and `market-api`.
- Seed the minimum visible symbol and coin metadata needed for homepage and exchange flows.
- Verify that local market APIs return non-empty, backend-backed responses for the baseline frontend routes.
- Record the exact schema/data prerequisites required for repeatable local frontend validation.

## Capabilities

### New Capabilities
- `local-market-schema-and-seed-baseline`: Defines the minimum local market schema and seed data required for the recovered frontend to render real homepage and exchange market content.

### Modified Capabilities
- None.

## Impact

- Affected systems: local MySQL market schema, `market.rpc`, `market-api`, homepage market list, exchange symbol metadata, and later frontend validation runs.
- Affected files: backend schema/seed assets, local setup notes, and validation artifacts generated after seeding.
- Dependency source: [report.md](E:/Project/web3/mscoin/openspec/changes/validate-frontend-against-local-backend/report.md).
