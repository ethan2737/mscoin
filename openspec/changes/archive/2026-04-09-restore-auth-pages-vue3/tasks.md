## 1. Runtime And Markup Cleanup

- [x] 1.1 Update login/register/auth-entry pages to the normalized Vue 3 runtime contract.
- [x] 1.2 Remove unsupported legacy constructs, damaged constants, and malformed validation/messages from the auth surfaces.

## 2. Form Behavior Recovery

- [x] 2.1 Stabilize form refs, validation wiring, and submit-path boot logic for login.
- [x] 2.2 Stabilize register SMS/captcha/submit flow so the page no longer fails before backend validation begins.

## 3. Verification

- [x] 3.1 Verify `/login` renders cleanly and does not throw frontend runtime errors.
- [x] 3.2 Verify `/register` renders cleanly and does not throw frontend runtime errors.
