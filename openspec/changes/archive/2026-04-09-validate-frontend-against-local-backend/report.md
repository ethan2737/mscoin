# Frontend Validation Against Local Backend

## Validation Preconditions

### Required Local Services

The frontend was validated against these locally running services:

- `ucenter.rpc` on `127.0.0.1:8081`
- `market.rpc` on `127.0.0.1:8082`
- `exchange.rpc` on `127.0.0.1:8083`
- `ucenter-api` on `127.0.0.1:8888`
- `market-api` on `127.0.0.1:8889`
- `exchange-api` on `127.0.0.1:8890`
- Docker-backed dependencies already listening:
  - MySQL on `33306`
  - Redis on `33679`
  - Mongo on `37018`
  - etcd gateway on `32379`
  - Kafka on `39092`

### Frontend Runtime / Proxy Confirmation

The active frontend dev server was validated at `http://127.0.0.1:3000`.

Observed proxy-backed requests confirm the frontend is pointed at local services:

- `/uc/check/login`
- `/market/symbol-thumb-trend`
- `/market/symbol-thumb`
- `/market/coin-info`
- `/socket.io`

### Data / Account Preconditions

Meaningful authenticated validation still requires:

- At least one known-good local member account that can log in through `/uc/login`
- Seeded market tables for symbol and coin metadata
- Seeded help/announcement ancillary endpoints if those capabilities are expected locally
- Optional wallet/order seed data for user-center order views

## Flow Outcomes

### Homepage

Status: `partial`

- `GET /` loads without frontend runtime errors.
- `POST /uc/check/login` returns `{"code":0,"data":false}`.
- `POST /market/symbol-thumb-trend` returns `{"code":0,"data":[]}`.
- The homepage now renders an empty market state (`暂无相关币种`) instead of fallback market rows because the backend is reachable and returning an empty dataset.

Classification:

- Frontend pathing is healthy.
- Real homepage market content is blocked by missing/empty local market seed data.

### Auth

Status: `partial`

- Login page boots without `pageerror`.
- `POST /uc/check/login` returns healthy anonymous state (`false`).
- A success-path login could not be validated because no known valid local account was provided.
- Direct dummy login submission reached `/uc/login` and returned HTTP `400`, which confirms the endpoint is reachable but does not validate a successful auth path.

Classification:

- Anonymous auth checks pass.
- Successful authenticated flow is blocked by missing test credentials.

### Help / Notice

Status: `fail`

- Help page boots, but `POST /uc/ancillary/more/help` returns `404`.
- Direct request to `POST /uc/announcement/page` also returns `404`.
- Notice page renders shell content but does not have a confirmed working backend-backed data flow in the local environment.

Classification:

- This is not a frontend runtime failure.
- The local backend currently does not expose the ancillary/help/announcement routes expected by the frontend recovery baseline.

### Exchange

Status: `partial`

- `/#/exchange/btc_usdt` boots without `pageerror`.
- `POST /market/symbol-thumb` returns `{"code":0,"data":[]}`.
- `POST /market/coin-info` returns application code `500` with message indicating missing table `market.coin`.
- Core exchange shell and realtime bootstrap render.

Classification:

- Frontend exchange runtime is healthy.
- Live exchange content is blocked by missing market database tables and seed data.

### User-Center Entrust History

Status: `blocked`

- Route resolves and frontend no longer emits stale legacy endpoints.
- Anonymous access is redirected back into login flow because `/uc/check/login` reports `false`.
- `POST /market/symbol-thumb` succeeds with an empty dataset.
- Authenticated order-history success path cannot be validated without a local account and matching order data.

Classification:

- Frontend routing and endpoint selection are healthy.
- Business validation is blocked by missing authenticated test conditions.

## Realtime Validation

Status: `pass`

Observed on `/#/exchange/btc_usdt`:

- `/socket.io/?EIO=3&transport=polling...` returned HTTP `200`
- A websocket upgrade to `/socket.io/?EIO=3&transport=websocket...` succeeded

Notes:

- This confirms the expected local socket backend path is reachable through the frontend proxy.
- It does not prove that meaningful market events are streaming, only that handshake and connection setup succeed.

## Backend / Environment Issues Confirmed During Validation

1. Missing market tables
   - `market-rpc` log shows `Error 1146 (42S02): Table 'market.exchange_coin' doesn't exist`
   - `/market/coin-info` reports `Table 'market.coin' doesn't exist`

2. Missing ancillary/help/announcement HTTP routes
   - `/uc/ancillary/more/help` -> `404`
   - `/uc/announcement/page` -> `404`

3. Missing authenticated validation conditions
   - No known-good local login account was supplied
   - No confirmed seeded order/wallet data was available for user-center validation

4. Non-blocking noise
   - `tawk.to` external script CORS/load failure
   - Vue Router alias warning for `/announcement/:id`
   - Duplicate `store` provide warning

## Summary Matrix

- Homepage: `partial` -> frontend healthy, backend data empty
- Login page: `partial` -> anonymous path healthy, success path blocked by missing credentials
- Help: `fail` -> backend route missing (`404`)
- Notice: `fail` -> backend route family not available locally
- Exchange: `partial` -> frontend healthy, backend market schema incomplete
- User-center entrust history: `blocked` -> auth/data prerequisites missing
- Realtime socket path: `pass` -> handshake and websocket upgrade succeed

## Recommended Next Steps

1. Restore or seed local market schema and data (`market.exchange_coin`, `market.coin`, visible symbols).
2. Confirm whether ancillary/help/announcement APIs were intentionally omitted from the current Go backend or still need to be implemented/mapped.
3. Provide one known-good local test account and at least one seeded order dataset for authenticated validation.
4. Re-run this validation change after data and account prerequisites are satisfied.
