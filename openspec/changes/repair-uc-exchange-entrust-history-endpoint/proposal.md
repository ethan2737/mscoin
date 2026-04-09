## Why

Authenticated user-center validation is no longer blocked by missing credentials, but `/#/uc/entrust/history` still fails because the frontend calls `/exchange/order/personal/history` and receives `404`. Until that route contract is repaired, order-history validation cannot distinguish between missing order data and a broken authenticated exchange endpoint.

## What Changes

- Confirm the canonical backend contract for authenticated user-center spot entrust history.
- Repair the frontend/backend route mismatch or missing handler for `/exchange/order/personal/history`.
- Re-run authenticated user-center history validation after the endpoint contract is restored.
- Document whether the remaining outcome is pass, fail, or blocked by missing order data after the route itself works.

## Capabilities

### New Capabilities
- `authenticated-uc-exchange-entrust-history`: Defines the backend-supported authenticated contract required for the user-center spot entrust history page to load under local validation.

### Modified Capabilities
- None.

## Impact

- Affected systems: user-center spot entrust history page, `exchange-api`, `exchange.rpc`, authenticated exchange order query path, and local authenticated frontend validation.
- Affected APIs: current frontend expectation `/exchange/order/personal/history` and any canonical replacement route chosen for the repaired contract.
- Dependency source: [report.md](E:/Project/web3/mscoin/openspec/changes/archive/2026-04-09-provide-local-auth-test-account-and-rerun-uc-validation/report.md).
