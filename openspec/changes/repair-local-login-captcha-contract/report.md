# repair-local-login-captcha-contract

## Summary

The recovered Vue 3 login page now completes a real local `/uc/login` success path without manual token injection.

## Contract Audit

### Confirmed frontend fallback payload

Browser capture against `/#/login` confirmed the current recovered page submits:

```json
{
  "username": "13800000000",
  "password": "Passw0rd!Local",
  "captcha": {
    "mode": "fallback",
    "passed": true
  }
}
```

### Confirmed backend rejection point

Two backend layers were blocking the local fallback mode:

1. `ucenter-api/internal/types/types.go`
   - `captcha.server` and `captcha.token` were treated as required fields by the request schema, so the frontend fallback payload was rejected before login logic ran.

2. `ucenter/internal/domain/captcha.go`
   - captcha verification always posted to the remote Vaptcha server and had no explicit local-compatible branch.

## Repair

### Local-compatible captcha contract

The backend now accepts the recovered frontend fallback contract by mapping:

- request `captcha.mode == "fallback"`
- request `captcha.passed == true`

to an internal local-only sentinel captcha:

- `server = "local-fallback"`
- `token = "passed"`

The `ucenter` captcha domain explicitly accepts only that sentinel pair as the local-compatible branch. All credential validation and token issuance still run through the normal `/uc/login` flow.

## Files Changed

- `mscoin-frontend/src/pages-vue3/uc/Login.vue`
- `mscoin-backend/ucenter-api/internal/types/types.go`
- `mscoin-backend/ucenter-api/internal/logic/login_logic.go`
- `mscoin-backend/ucenter/internal/domain/captcha.go`

## Validation

### API validation

- `POST /uc/login` with the local sentinel captcha returned `code: 0`
- successful response still included the normal token and member payload

### Real browser validation

Using Playwright against `http://127.0.0.1:3000/#/login`:

- filled seeded account `13800000000 / Passw0rd!Local`
- clicked the local fallback captcha button
- submitted the login form

Observed result:

- request body remained the frontend fallback shape with `mode/passed`
- `/uc/login` returned HTTP `200` with `code: 0`
- browser navigated to `http://127.0.0.1:3000/#/uc/safe`
- `localStorage.TOKEN` was set
- `localStorage.MEMBER` was set
- no `pageerror` occurred during the flow

## Residual Noise

- `tawk.to` still emits a third-party CORS/load error in the console
- this change does not modify production Vaptcha behavior beyond the explicit local fallback contract
