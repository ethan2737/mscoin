# Seed Local Market Schema And Data Report

## Summary

This change restored the minimum local `market` MySQL baseline required for the recovered Vue 3 frontend to render backend-backed homepage and exchange content.

## Audit Findings

- Required market tables were confirmed from backend models and DAO queries:
  - `market.coin`
  - `market.exchange_coin`
- Required query fields were confirmed from `market.rpc` usage:
  - `exchange_coin`: `symbol`, `base_symbol`, `coin_symbol`, `visible`, `base_coin_scale`, `coin_scale`, `enable`, `exchangeable`, `sort`, `zone`
  - `coin`: `unit`, `name`, `name_cn`, `status`, `usd_rate`, `withdraw_scale`, `infolink`, `information`
- No reusable schema or seed asset was found in-repo for these tables. This change therefore publishes an explicit local baseline artifact at [local-market-baseline.sql](E:/Project/web3/mscoin/openspec/changes/seed-local-market-schema-and-data/local-market-baseline.sql).

## Local Baseline Applied

The following baseline was applied to MySQL database `market`:

- Created table `coin`
- Created table `exchange_coin`
- Seeded coins:
  - `BTC`
  - `USDT`
  - `ETH`
  - `SOL`
- Seeded visible symbols:
  - `BTC/USDT`
  - `ETH/USDT`
  - `SOL/USDT`
  - `ETH/BTC`
  - `SOL/BTC`
  - `SOL/ETH`

Applied row counts after seeding:

- `coin_count = 4`
- `exchange_coin_count = 6`

## Backend Validation

Verified against local `market-api` at `http://127.0.0.1:8889`:

- `POST /market/coin-info` with `unit=BTC` returns `code: 0`
- `POST /market/symbol-info` with `symbol=BTC/USDT` returns `code: 0`
- `POST /market/symbol-thumb-trend` returns 6 visible symbols after restarting `market-api`
- `POST /market/symbol-thumb` returns the same warmed cache set after restarting `market-api`

Important local prerequisite:

- `market-api` must be restarted after first-time seeding because its thumb cache is initialized on process startup. If it starts before `exchange_coin` exists, `/market/symbol-thumb*` may remain empty until restart.

## Frontend Validation

Browser validation was re-run against the local backend and Vite dev server at `http://127.0.0.1:3000`:

- Homepage `USDT` tab showed:
  - `BTC/USDT`
  - `ETH/USDT`
  - `SOL/USDT`
- Homepage `BTC` tab showed:
  - `ETH/BTC`
  - `SOL/BTC`
- Homepage `ETH` tab showed:
  - `SOL/ETH`
- Exchange route `/#/exchange/btc_usdt` rendered `BTC/USDT`
- No page-level runtime exception was observed during homepage/exchange validation

## Residual Notes

- Seeded market data is intentionally minimal and price fields remain zero-valued until real kline/trade data flows through Kafka and Mongo-backed market processing.
- Non-blocking console noise still exists from unrelated missing assets and third-party scripts, but it does not block homepage or exchange rendering for this local baseline.
