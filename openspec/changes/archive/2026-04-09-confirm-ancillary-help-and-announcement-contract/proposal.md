## Why

Local-backend validation showed that the help and announcement frontend flows now reach the expected local backend, but both endpoint families return `404`. Before those pages can be repaired or validated further, we need to determine whether the current Go backend intentionally omits those capabilities or whether the frontend is calling the wrong mapped routes.

## What Changes

- Audit the current frontend expectation for help and announcement flows.
- Confirm whether the local backend contains implementations for those capability families under different routes or services.
- Distinguish between “missing implementation” and “missing/incorrect route mapping”.
- Record the canonical contract decision and identify the next repair step for whichever case is true.

## Capabilities

### New Capabilities
- `ancillary-help-announcement-contract-confirmation`: Defines how the local frontend and backend contract for help and announcement flows is confirmed and classified as missing implementation or missing mapping.

### Modified Capabilities
- None.

## Impact

- Affected systems: `ucenter-api`, any ancillary/help/announcement backend handlers, CMS help/notice frontend routes, and subsequent recovery changes that depend on those contracts.
- Affected outputs: contract confirmation report and follow-up implementation direction.
- Dependency source: [report.md](E:/Project/web3/mscoin/openspec/changes/validate-frontend-against-local-backend/report.md).
