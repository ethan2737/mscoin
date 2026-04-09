# Ancillary Help And Announcement Contract Report

## Summary

This change audited the current frontend expectations and local Go backend routing for the help and announcement capability families. The result is not a frontend mapping problem. Both families are currently **missing implementation** in the local Go backend slice that serves `/uc/**`.

## Frontend Endpoint Expectations

### Help family

The active Vue 3 CMS pages currently expect the following backend routes:

- [Help.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/cms/Help.vue) calls `POST /uc/ancillary/more/help`
- [HelpList.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/cms/HelpList.vue) calls `POST /uc/ancillary/more/help/page`
- [HelpDetail.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/cms/HelpDetail.vue) calls:
  - `POST /uc/ancillary/more/help/page/top`
  - `POST /uc/ancillary/more/help/detail`

### Announcement family

The active Vue 3 pages currently expect two announcement endpoint families:

- [Index.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/index/Index.vue) uses `api.uc.announcement`
- [api.js](E:/Project/web3/mscoin/mscoin-frontend/src/config/api.js) defines `api.uc.announcement` as `POST /uc/announcement/page`
- [Notice.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/cms/Notice.vue) calls `POST /uc/ancillary/announcement`
- [NoticeItem.vue](E:/Project/web3/mscoin/mscoin-frontend/src/pages-vue3/cms/NoticeItem.vue) calls:
  - `POST /uc/ancillary/announcement`
  - `POST /uc/announcement/more`

## Backend Route Audit

The current local Go backend route registrations were inspected in:

- [routes_handler.go](E:/Project/web3/mscoin/mscoin-backend/ucenter-api/internal/handler/routes_handler.go)
- [ucenterapi.api](E:/Project/web3/mscoin/mscoin-backend/ucenter-api/ucenterapi.api)

Confirmed registered `ucenter-api` routes are limited to:

- register
- mobile code
- login
- login check
- asset wallet and transaction queries
- approve security setting
- withdraw support and records

No handlers, route registrations, or `.api` declarations were found for:

- `/uc/ancillary/more/help`
- `/uc/ancillary/more/help/page`
- `/uc/ancillary/more/help/page/top`
- `/uc/ancillary/more/help/detail`
- `/uc/announcement/page`
- `/uc/announcement/more`
- `/uc/ancillary/announcement`

No alternate handler ownership was found in the current Go backend services under:

- `ucenter-api`
- `ucenter`
- `market-api`
- `market`
- `exchange-api`
- `exchange`

## Runtime Evidence

Direct local backend checks against `http://127.0.0.1:8888` returned `404 Not Found`:

- `POST /uc/ancillary/more/help`
- `POST /uc/announcement/page`
- `POST /uc/ancillary/announcement`
- `POST /uc/announcement/more`

Browser validation against the recovered frontend also matched this classification:

- `/#/help` renders the shell but no help content
- `/#/notice/item/0` renders the shell and fallback empty-state copy
- No page-level runtime exception occurred; the failure is downstream of missing backend content endpoints

## Classification

### Help family

- Classification: **missing implementation**
- Canonical route decision: no canonical local Go route currently exists
- Evidence: frontend calls `/uc/ancillary/more/help*`, but no matching or alternate help handlers are registered anywhere in the current Go backend

### Announcement family

- Classification: **missing implementation**
- Canonical route decision: no canonical local Go route currently exists
- Evidence: frontend uses both `/uc/announcement/*` and `/uc/ancillary/announcement`, but the current Go backend registers neither family and contains no alternate announcement service route

## Follow-up Recommendation

The next repair should be a dedicated backend-facing change, not a frontend remap.

Recommended concrete action:

1. Define a single canonical local content contract for help and announcement under the current backend architecture.
2. Implement the missing `ucenter-api` handlers and corresponding `ucenter` RPC/domain/data access support for:
   - help list
   - help page
   - help page top
   - help detail
   - announcement list
   - announcement detail
3. After backend capability exists, normalize the frontend onto the chosen single announcement family and rerun local validation.

If content capability is intentionally out of scope for the Go backend, the alternative is to explicitly replace these pages with local mock/static content for local validation. That is a product decision, not a route-mapping fix.
