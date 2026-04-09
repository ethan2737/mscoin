## 1. Contract Audit

- [x] 1.1 Confirm the exact fallback captcha payload currently submitted by the recovered Vue 3 login page.
- [x] 1.2 Locate the backend login verification point that rejects the local fallback captcha mode.

## 2. Contract Repair

- [x] 2.1 Implement a documented local-compatible captcha verification branch for `/uc/login`.
- [x] 2.2 Keep successful login token issuance and credential validation on the normal backend login path.
- [x] 2.3 Verify the seeded local validation account can authenticate through `/uc/login` without external token injection.

## 3. Browser Validation

- [x] 3.1 Re-run `/#/login` in a real browser and confirm authenticated navigation succeeds.
- [x] 3.2 Publish a change-local artifact describing the repaired local captcha contract and local login validation result.
