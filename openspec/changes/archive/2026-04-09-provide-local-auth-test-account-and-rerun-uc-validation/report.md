# Local Auth Test Account And User-Center Validation Report

## Summary

This change established a reusable local authenticated validation baseline and reran selected user-center flows in a real browser. The local account and wallet baseline are now available. Authenticated validation is no longer blocked by missing credentials.

## Local Auth Baseline

The reusable local validation baseline is defined in [local-auth-baseline.sql](E:/Project/web3/mscoin/openspec/changes/provide-local-auth-test-account-and-rerun-uc-validation/local-auth-baseline.sql).

### Local validation account

- Database: `mscoin`
- Member id: `10001`
- Username: `local-tester`
- Login phone: `13800000000`
- Local-only password: `Passw0rd!Local`

Password details:

- Salt: `LocalAuthSeedSaltForMscoin2026`
- Hash algorithm: PBKDF2-HMAC-SHA512, 10000 iterations, 128-byte derived key

### Minimum wallet baseline

The seeded member wallet baseline includes:

- `BTC`
  - balance: `0.25`
  - frozen: `0.01`
- `USDT`
  - balance: `1250.50`
  - frozen: `50.25`

### Minimum order baseline

No exchange order rows were seeded in this change.

Reason:

- The selected user-center order-history flow is currently blocked earlier by a backend route contract problem (`404`), so adding order rows would not yet make the route pass.

## Equivalent Local Sign-In Path

The current recovered login page uses a frontend fallback captcha UI, but the Go backend still expects a real Vaptcha server/token pair. Because of that mismatch, successful interactive login through `/uc/login` is not yet a reliable local validation path.

For this change, authenticated validation used the documented equivalent local setup path:

1. Seed the local member and wallet baseline from [local-auth-baseline.sql](E:/Project/web3/mscoin/openspec/changes/provide-local-auth-test-account-and-rerun-uc-validation/local-auth-baseline.sql)
2. Generate a local JWT for `userId=10001` using secret `!@#$mscoin`
3. Set browser local storage:
   - `TOKEN=<generated jwt>`
   - `MEMBER=<serialized member payload>`

This equivalent path was verified by:

- `POST /uc/check/login` returning `{"code":0,"data":true}`

## Authenticated Flow Outcomes

Browser validation was executed against `http://127.0.0.1:3000` with real page loads after injecting the local auth baseline.

### `/#/uc/safe`

- Status: `pass`
- Result:
  - page rendered without `pageerror`
  - `POST /uc/approve/security/setting` returned `200`
  - authenticated security information rendered
- Notes:
  - the page still emits unrelated homepage announcement noise because `App.vue` requests `/uc/announcement/page`, which remains missing

### `/#/uc/money`

- Status: `pass`
- Result:
  - page rendered without `pageerror`
  - `POST /uc/asset/wallet` returned `200`
  - seeded `BTC` and `USDT` balances rendered in the asset table

### `/#/uc/entrust/history`

- Status: `fail`
- Result:
  - page rendered without `pageerror`
  - `POST /market/symbol-thumb` returned `200`
  - `POST /exchange/order/personal/history` returned `404`
  - UI showed `No Data`, but the root cause is not simply missing orders
- Classification:
  - backend route contract failure or stale frontend endpoint
  - not blocked by auth anymore

## Additional Runtime Findings

- `ucenter.rpc` and `ucenter-api` had to be restarted during this change to refresh local service discovery and make member/asset RPC calls resolve cleanly again
- third-party `tawk.to` script still fails due CORS; this remains non-blocking for authenticated validation
- authenticated pages still inherit missing announcement noise from `/uc/announcement/page`

## Remaining Blockers

1. Interactive login success path is still blocked by captcha contract mismatch
   - frontend uses fallback captcha UX
   - backend still requires real Vaptcha verification fields

2. Exchange user-center history path still fails after authentication
   - `POST /exchange/order/personal/history` returned `404`
   - this must be fixed before order seed data becomes meaningful for that route

3. Announcement route family is still missing
   - authenticated pages continue to emit `404` noise from `/uc/announcement/page`

## Recommended Next Steps

1. Create a dedicated change to reconcile local login captcha behavior with the recovered frontend fallback login flow.
2. Create a dedicated change to repair authenticated exchange user-center endpoints, starting with `/exchange/order/personal/history`.
3. After that route contract is corrected, seed one minimal personal entrust history dataset and rerun authenticated order validation.
