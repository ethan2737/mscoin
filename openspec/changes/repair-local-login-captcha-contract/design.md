## Context

The current recovered login page in [Login.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/uc/Login.vue) intentionally exposes a frontend fallback captcha action for local recovery. However, the Go backend login stack still expects the original Vaptcha request shape and always calls the remote verifier in `CaptchaDomain.Verify`. In practice, local browser login cannot succeed even when the page is otherwise healthy, because the backend has no local-compatible verification branch.

## Goals / Non-Goals

**Goals:**
- Make the recovered local login page complete a successful login through the real `/uc/login` path.
- Introduce a local-compatible captcha verification branch with explicit contract rules.
- Keep the local verification mode controlled and documented so later validation can use it repeatably.
- Verify the repaired flow in a real browser rather than only by direct API calls.

**Non-Goals:**
- Rebuild or replace the original production Vaptcha integration comprehensively.
- Solve broader authenticated business issues outside login, such as announcement routes or entrust history.
- Introduce permanent insecure login bypass behavior without an explicit local-only contract.

## Decisions

1. Add an explicit local captcha acceptance branch rather than trying to simulate Vaptcha traffic.
   Rationale: the recovered frontend already communicates local fallback intent, and emulating the original external service would add unnecessary moving parts.
   Alternative considered: run a local Vaptcha stub service. Rejected because it preserves the wrong dependency and complicates local setup.

2. Keep the contract centered on `/uc/login` instead of adding a separate local login endpoint.
   Rationale: local frontend validation should exercise the same login route and token issuance path as the real app.
   Alternative considered: add `/uc/login/local`. Rejected because it would create a parallel path that weakens the value of frontend validation.

3. Treat the local captcha mode as an explicit request contract, not an implicit environment bypass.
   Rationale: the backend should only accept fallback verification when the request explicitly matches the recovered local mode.
   Alternative considered: skip captcha verification whenever running locally. Rejected because it is too broad and would hide contract mistakes.

## Risks / Trade-offs

- [Local captcha acceptance could accidentally broaden login bypass behavior] → Mitigation: keep the contract explicit, narrow, and documented as local-only.
- [Frontend and backend may still serialize fallback captcha fields differently] → Mitigation: verify the final request shape against the actual Vue 3 login page and test in a real browser.
- [Production captcha restoration may later diverge from the local branch] → Mitigation: document the local contract clearly so future work can replace or remove it intentionally.

## Migration Plan

1. Confirm the exact request shape emitted by the recovered login page.
2. Update the backend login/captcha path to accept that documented local fallback contract.
3. Re-run browser login through `/#/login` with the seeded local account.
4. Record the resulting local login baseline in the change report.

Rollback strategy: remove the local fallback branch and return to the original verifier-only path if it causes unacceptable contract ambiguity.

## Open Questions

- Should the local captcha acceptance be keyed strictly off a fallback marker in `captcha`, or also require a local environment guard?
- Once local browser login succeeds, should token-injection validation remain documented as a secondary fallback or be retired?
