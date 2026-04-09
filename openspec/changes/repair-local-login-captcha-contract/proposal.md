## Why

The recovered Vue 3 login page already uses a local fallback captcha UX, but the Go backend still hard-requires a real Vaptcha verification contract. That mismatch prevents a reliable local interactive login success path and forces authenticated validation to rely on equivalent token injection instead of the actual frontend login flow.

## What Changes

- Align the local frontend and backend login captcha contract so the recovered login page can complete a successful local login without external Vaptcha dependency.
- Define one explicit local login verification mode that the backend accepts when the frontend uses the recovered fallback captcha path.
- Preserve a clear boundary between local validation behavior and any future production captcha integration.
- Re-run real-browser login validation through `/#/login` after the contract is repaired.

## Capabilities

### New Capabilities
- `local-login-captcha-contract`: Defines the local-only captcha verification contract required for the recovered login page to complete a real authenticated login flow.

### Modified Capabilities
- None.

## Impact

- Affected systems: `mscoin-frontend` login page, `ucenter-api` login handler, `ucenter.rpc` login logic, captcha verification path, and local authenticated browser validation.
- Affected APIs: `POST /uc/login`.
- Affected validation: local interactive login flow, later user-center validation changes that currently rely on token injection.
