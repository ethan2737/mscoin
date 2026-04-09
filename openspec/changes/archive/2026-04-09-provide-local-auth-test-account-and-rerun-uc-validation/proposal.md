## Why

The latest local-backend validation confirmed that user-center routing and endpoint selection are now structurally healthy, but authenticated business flows remain blocked because no known-good local login account or seeded order data is available. We need a dedicated change that creates a repeatable local auth validation condition and then re-runs the user-center flow checks against it.

## What Changes

- Establish one usable local test account for frontend authentication validation.
- Confirm the minimum associated wallet/order data required to validate selected user-center flows.
- Re-run user-center validation with that account and classify which flows now pass, fail, or remain blocked.
- Publish an updated user-center validation artifact based on the authenticated local run.

## Capabilities

### New Capabilities
- `local-auth-test-account-and-uc-validation`: Defines the local authenticated test-account baseline required for user-center frontend validation and records the results of the rerun.

### Modified Capabilities
- None.

## Impact

- Affected systems: local auth/login flow, user-center validation routes, local member/account seed state, and recovery evidence for authenticated flows.
- Affected outputs: authenticated validation report and reusable local test-account setup notes.
- Dependency source: [report.md](E:/Project/web3/mscoin/openspec/changes/validate-frontend-against-local-backend/report.md).
