# Frontend Business/Data-Flow Audit

Last updated: 2026-04-09

## Scope

This audit covers the active Vue 3 + Vite frontend baseline defined in `docs/frontend-migration-baseline.md`, with emphasis on:

- shared app shell behavior in `mscoin-frontend/src/App.vue`
- route inventory in `mscoin-frontend/src/config/routes-vue3.js`
- key business flows: homepage, login, registration, help/notice, exchange entry
- runtime data dependencies, proxy contracts, and local backend availability

## Method

The audit combined:

- static route and code-path inspection
- dependency tracing from frontend caller to endpoint/socket contract
- local runtime checks for expected backend ports
- browser verification on the highest-risk flows using Playwright

## Environment Snapshot

- Active frontend dev server is reachable on `http://127.0.0.1:3000`.
- Only port `3000` was listening during the audit.
- Expected backend ports were not listening:
  - `8888` for `ucenter-api`
  - `8889` for `market-api`
  - `8890` for `exchange-api`
  - `9100` is not the current market realtime path; the active frontend should use `/socket.io` proxied to `8889`

## Route Inventory

| Route group | Primary routes | Business purpose | Current audit status |
| --- | --- | --- | --- |
| Homepage | `/`, `/index` | landing, market discovery, notice teaser, banner, entry to exchange | partially rendering; mixed real/fallback/static |
| Auth | `/login`, `/register`, `/mobile/register`, `/findpwd` | login, signup, password recovery | login renders with warnings; register has runtime contract break |
| CMS | `/help`, `/help/list`, `/help/detail`, `/notice`, `/notice/item/:id`, `/about` | help center, notices, about pages | requests still use hardcoded `http://localhost`; data contract not normalized |
| Exchange | `/exchange`, `/exchange/:pair` | spot trading, market list, depth, trade feed, wallet and order actions | mount-stage runtime failure; not usable |
| Swap | `/swapindex`, `/swapindex/:pair`, `/kline/:pair` | contract trading | inherits same old store/socket contract risks as exchange |
| OTC | `/otc`, `/otc/trade/*`, `/chat`, `/checkuser`, `/otc/ad/*` | OTC trading and chat | heavy `http://localhost` and socket legacy usage; not fully verified |
| Activity/Mining/Crowdfunding/CTC | `/activity`, `/mining`, `/crowdfunding`, `/ctc` | promo and auxiliary flows | many pages still hardcode `http://localhost`; not fully verified |
| User center | `/uc/**` | profile, security, wallet, orders, innovation, contract history | many screens still depend on old global/store contract |

## Shared Shell Audit

The root shell in `mscoin-frontend/src/App.vue` currently owns:

- header navigation and top-level route entry
- language dropdown
- login/register top menu
- logged-in user dropdown
- mobile drawer navigation
- login-state check on mount
- footer links and help shortcuts

Observed shell behaviors:

- Language dropdown is wired with Element Plus and responded in browser verification; two menu items opened successfully.
- Login state is checked in `checkLogin()` from `App.vue`, but current runtime still emits server-side `500` responses during shell boot because the local backend is not up.
- The shell still emits Vue warnings about duplicate `store` provide/plugin registration and deprecated router-guard `next()` usage, which indicates migration cleanup is incomplete even where the UI appears to work.

## UI Region Classification

### Homepage (`src/pages-vue3/index/Index.vue`)

| Region | Classification | Source |
| --- | --- | --- |
| announcement bar | API-backed with default fallback | `/uc/announcement/page`; falls back to default welcome text |
| market tabs/table | API-backed with local fallback | `/market/symbol-thumb-trend`; local `fallbackCoins` currently populate only `USDT` symbols |
| favorite toggle | auth-gated API | `/exchange/favor/add`, `/exchange/favor/delete`, `/exchange/favor/find` |
| feature cards | static | local image/text assets |
| banner/world-map section | API-backed with local fallback | `/uc/ancillary/system/advertise`; currently falls back to local `bannerbg.png` |
| realtime market updates | socket.io-backed | `/socket.io` via Vite proxy to `market-api` |

### Login (`src/pages-vue3/uc/Login.vue`)

| Region | Classification | Source |
| --- | --- | --- |
| title/labels | static but migration-polluted | hardcoded constants, some source text still contains encoding damage |
| captcha area | external dependency | `vaptcha` container and token callback |
| submit action | API-backed | `store.state.host + store.state.api.uc.login` |
| post-login navigation | route transition | `/uc/safe` |

### Register (`src/pages-vue3/uc/Register.vue`)

| Region | Classification | Source |
| --- | --- | --- |
| marquee recommendation strip | API-backed but broken by contract mismatch | `store.state.api.uc.getRandomly` |
| sms code send | API-backed | `/uc/mobile/code` |
| register submit | API-backed | `/uc/register/phone` |
| captcha area | external dependency | `vaptcha` |
| agreement link | static route composition | help detail route |

### Exchange (`src/pages-vue3/exchange/Exchange.vue`)

| Region | Classification | Source |
| --- | --- | --- |
| market pair list | API-backed | `/market/symbol-thumb` |
| favorites | auth-gated API | `/exchange/favor/*` |
| plate/trade/depth | API-backed + socket-backed | `/market/exchange-plate*`, `/market/latest-trade`, socket.io |
| current pair realtime | socket-backed | `io('http://localhost')` legacy path |
| wallet and orders | auth-gated API | exchange/uc routes from old API contract |

## Data Dependency Matrix

| Flow | Frontend caller | Dependency | Auth | Expected service |
| --- | --- | --- | --- | --- |
| homepage market list | `Index.vue` | `/market/symbol-thumb-trend` | optional token for favorites | `market-api:8889` |
| homepage banner | `Index.vue` | `/uc/ancillary/system/advertise` | no | `ucenter-api:8888` |
| homepage notice | `Index.vue` | `/uc/announcement/page` | no | `ucenter-api:8888` |
| homepage realtime | `Index.vue` | `/socket.io` | no | `market-api:8889` |
| login submit | `Login.vue` | `store.state.api.uc.login` | no | `ucenter-api:8888` |
| register random promo | `Register.vue` | `store.state.api.uc.getRandomly` | no | `ucenter-api:8888` |
| register sms | `Register.vue` | `/uc/mobile/code` | no | `ucenter-api:8888` |
| register submit | `Register.vue` | `/uc/register/phone` | no | `ucenter-api:8888` |
| help list | `HelpList.vue` | `/uc/ancillary/more/help/page` | no | `ucenter-api:8888` |
| notice detail | `NoticeItem.vue` | `/uc/announcement/more`, `/uc/ancillary/announcement` | no | `ucenter-api:8888` |
| exchange boot | `Exchange.vue` | `store.state.host + store.state.api.*` and `io('http://localhost')` | mixed | `market-api`, `exchange-api`, `ucenter-api` |

## Browser Verification Summary

### Homepage

- Rendered successfully.
- Language dropdown opened and exposed two items.
- `USDT` table showed three rows from local fallback data.
- External Tawk script failed by CORS, but did not block render.
- Market websocket handshake failed because the backend socket endpoint is unavailable in the current local environment.

### Login

- Page rendered visibly.
- Browser warnings included non-production-safe template refs and backend `500` responses.
- Inputs and CTA rendered, but full submit-path verification is blocked by backend availability and captcha dependency.

### Register

- Page rendered only partially.
- Runtime threw repeated `Cannot read properties of undefined (reading 'uc')`.
- `marquee` is unresolved in Vue 3.
- Root cause is code still expecting `store.state.host` and `store.state.api`, which do not exist in `store-vue3`.

### Exchange

- Page did not render the trading surface.
- Browser errors repeatedly showed `Cannot read properties of null (reading 'offsetWidth')`.
- Console also showed missing prop warnings and unhandled mounted-hook/render errors.
- Exchange remains unusable even before real backend data is considered.

## Confirmed Findings

### P0: Migration contract fracture between boot/runtime and migrated pages

- Evidence:
  - `mscoin-frontend/src/main-vue3.js:61` still sets global `host` to `http://localhost`
  - `mscoin-frontend/src/main-vue3.js:63` sets global `api = Api`
  - `mscoin-frontend/src/config/store-vue3.js:3` only defines `member`, `activeNav`, `lang`, `exchangeSkin`, and `loginTimes`
  - `mscoin-frontend/src/pages-vue3/uc/Register.vue:199` reads `store.state.api.uc.getRandomly`
  - `mscoin-frontend/src/pages-vue3/exchange/Exchange.vue:570` reads `store.state.host`
- Impact:
  - Register page throws immediately.
  - Exchange and swap families are structurally unsafe to mount.
  - Migrated pages are using a state contract that the Vue 3 store no longer provides.

### P0: Exchange flow is not operational

- Evidence:
  - Browser verification on `/#/exchange/btc_usdt` produced repeated `offsetWidth` null errors and no trading surface.
  - `mscoin-frontend/src/pages-vue3/exchange/Exchange.vue:1082` still opens `io('http://localhost')`.
- Impact:
  - The core trading flow cannot be considered recovered.
  - This blocks meaningful QA of pair switching, order forms, depth, trade feed, and wallet interactions.

### P1: Homepage currently mixes fallback, static, and real-data regions without clear contract boundaries

- Evidence:
  - `mscoin-frontend/src/pages-vue3/index/Index.vue:178` defines local `fallbackCoins`
  - `mscoin-frontend/src/pages-vue3/index/Index.vue:222` defines local `fallbackAds`
  - `mscoin-frontend/src/pages-vue3/index/Index.vue:397` requests `/market/symbol-thumb-trend`
  - `mscoin-frontend/src/pages-vue3/index/Index.vue:413` requests `/uc/ancillary/system/advertise`
- Impact:
  - Homepage appears partially healthy while still not proving real backend integration.
  - `BTC` and `ETH` tabs are empty because the current fallback only covers `USDT` pairs.
  - The world-map banner section is a fallback image, not verified business content.

### P1: Large parts of the Vue 3 page tree still hardcode `http://localhost` or legacy socket paths

- Evidence:
  - Static scan shows many active Vue 3 pages still define `const host = 'http://localhost'`.
  - OTC chat, swap, activity, mining, CMS, and user-center pages still contain legacy host or socket assumptions.
- Impact:
  - Even with Vite proxy configured correctly, these pages bypass it.
  - The app will continue to fail inconsistently until host/socket contracts are normalized.

### P1: Auth pages contain migration regressions beyond backend unavailability

- Evidence:
  - `mscoin-frontend/src/pages-vue3/uc/Register.vue:4` still uses `<marquee>`
  - `Register.vue` and `Login.vue` still use old validation/message patterns and non-ref form refs
  - source files contain visible encoding damage in constant definitions and validation messages
- Impact:
  - The pages may visually render, but they are not reliable or maintainable migration outputs.
  - Register is currently blocked by frontend runtime errors even before API validation starts.

### P2: Shared shell migration cleanup is incomplete

- Evidence:
  - Browser console shows duplicate `store` provide/plugin warnings.
  - Vue Router warns about deprecated `next()` use and alias parameter mismatch for `/announcement/:id`.
- Impact:
  - Not always fatal, but signals unstable foundation behavior in navigation and shell lifecycle.

### P3: External support widget fails by CORS

- Evidence:
  - Tawk embed request fails on homepage and register page.
- Impact:
  - Non-blocking for core business flows.
  - Should be isolated or deferred until core recovery is complete.

## Defect Classification

### Frontend regressions

- Vue 3 store contract no longer matches migrated page expectations
- exchange page mount/render failures
- register page runtime errors and unsupported template/component usage
- legacy `http://localhost` and raw socket paths across active Vue 3 pages

### Dependency gaps

- `ucenter-api`, `market-api`, and `exchange-api` were not running locally during the audit
- socket.io market handshake cannot complete without `market-api`
- auth submit, notice/help, and wallet/order flows cannot be fully validated without backend availability

### Unresolved contract questions

- Whether migrated auth pages should preserve legacy UX exactly or only restore functional behavior
- Whether homepage should continue to use temporary fallback content during recovery, and if so, how that state should be surfaced to QA

## Recommended Remediation Sequence

1. Rebuild the shared Vue 3 runtime contract.
   - Normalize `host`, `api`, auth token, and helper access so pages stop depending on removed `store.state.*` fields.
2. Fix core exchange boot path before auxiliary pages.
   - Exchange is the main product surface and is currently unusable.
3. Fix auth pages after the runtime contract is stabilized.
   - Remove `store.state.host/api` assumptions, `marquee`, and malformed/encoded constants.
4. Sweep all active Vue 3 pages for `http://localhost`, `ws://localhost`, and `io('http://localhost')`.
   - Convert them to the canonical proxy/runtime path.
5. Start local backend services and re-run QA against real data.
   - Only after frontend contract cleanup does backend integration testing become high signal.
6. Replace temporary homepage fallback data with explicit QA fixtures or real API responses.
   - QA must be able to tell whether a region is mocked, fallback, or live.

## Suggested Follow-up Changes

- `repair-vue3-runtime-contract`
  - normalize shared frontend runtime access to host/api/auth helpers
- `restore-auth-pages-vue3`
  - fix login/register/find-password behavior and remove migration regressions
- `restore-exchange-and-swap-runtime`
  - repair exchange/swap boot, data contracts, and socket usage
- `normalize-vue3-network-contracts`
  - remove hardcoded localhost usage across the active Vue 3 page tree
- `validate-frontend-against-local-backend`
  - re-run the audit with `8888/8889/8890` services up and capture true business-flow results
