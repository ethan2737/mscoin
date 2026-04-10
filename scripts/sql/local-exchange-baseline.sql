CREATE TABLE IF NOT EXISTS exchange_order (
  id BIGINT NOT NULL AUTO_INCREMENT,
  order_id VARCHAR(128) NOT NULL,
  amount DECIMAL(32,16) NOT NULL DEFAULT 0,
  base_symbol VARCHAR(32) NOT NULL DEFAULT '',
  canceled_time BIGINT NOT NULL DEFAULT 0,
  coin_symbol VARCHAR(32) NOT NULL DEFAULT '',
  completed_time BIGINT NOT NULL DEFAULT 0,
  direction INT NOT NULL DEFAULT 0,
  member_id BIGINT NOT NULL DEFAULT 0,
  price DECIMAL(32,16) NOT NULL DEFAULT 0,
  status INT NOT NULL DEFAULT 4,
  symbol VARCHAR(64) NOT NULL DEFAULT '',
  time BIGINT NOT NULL DEFAULT 0,
  traded_amount DECIMAL(32,16) NOT NULL DEFAULT 0,
  turnover DECIMAL(32,16) NOT NULL DEFAULT 0,
  type INT NOT NULL DEFAULT 1,
  use_discount VARCHAR(16) NOT NULL DEFAULT '0',
  PRIMARY KEY (id),
  UNIQUE KEY uk_exchange_order_order_id (order_id),
  KEY idx_exchange_order_symbol_status (symbol, status),
  KEY idx_exchange_order_member_symbol_status (member_id, symbol, status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DELETE FROM exchange_order;

INSERT INTO exchange_order (
  order_id,
  amount,
  base_symbol,
  canceled_time,
  coin_symbol,
  completed_time,
  direction,
  member_id,
  price,
  status,
  symbol,
  time,
  traded_amount,
  turnover,
  type,
  use_discount
)
WITH seed_symbols AS (
  SELECT 'BTC/USDT' AS symbol, 'BTC' AS coin_symbol, 'USDT' AS base_symbol, 71769.9 AS mid_price
  UNION ALL SELECT 'ETH/USDT', 'ETH', 'USDT', 2193.73
  UNION ALL SELECT 'OKB/USDT', 'OKB', 'USDT', 84.06
  UNION ALL SELECT 'SOL/USDT', 'SOL', 'USDT', 82.96
  UNION ALL SELECT 'DOGE/USDT', 'DOGE', 'USDT', 0.09215
  UNION ALL SELECT 'XRP/USDT', 'XRP', 'USDT', 1.3382
  UNION ALL SELECT 'BCH/USDT', 'BCH', 'USDT', 438.9
  UNION ALL SELECT '1INCH/USDT', '1INCH', 'USDT', 0.0920
  UNION ALL SELECT 'AAVE/USDT', 'AAVE', 'USDT', 90.24
  UNION ALL SELECT 'ACH/USDT', 'ACH', 'USDT', 0.005837
  UNION ALL SELECT 'ADA/USDT', 'ADA', 'USDT', 0.2501
  UNION ALL SELECT 'AEVO/USDT', 'AEVO', 'USDT', 0.1364
  UNION ALL SELECT 'AGLD/USDT', 'AGLD', 'USDT', 1.48
  UNION ALL SELECT 'ALGO/USDT', 'ALGO', 'USDT', 0.1672
  UNION ALL SELECT 'APE/USDT', 'APE', 'USDT', 0.0887
  UNION ALL SELECT 'API3/USDT', 'API3', 'USDT', 0.28
  UNION ALL SELECT 'APT/USDT', 'APT', 'USDT', 0.8381
  UNION ALL SELECT 'AR/USDT', 'AR', 'USDT', 1.716
  UNION ALL SELECT 'ARB/USDT', 'ARB', 'USDT', 0.11132
  UNION ALL SELECT 'ATOM/USDT', 'ATOM', 'USDT', 1.821
  UNION ALL SELECT 'ETH/BTC', 'ETH', 'BTC', 0.03056
  UNION ALL SELECT 'OKB/BTC', 'OKB', 'BTC', 0.001171
  UNION ALL SELECT 'SOL/BTC', 'SOL', 'BTC', 0.001156
  UNION ALL SELECT 'BETH/ETH', 'BETH', 'ETH', 1.0008
  UNION ALL SELECT 'STETH/ETH', 'STETH', 'ETH', 0.9993
)
SELECT
  CONCAT('LOCAL-SEED-', REPLACE(symbol, '/', '-'), '-BUY-1'),
  1.25000000,
  base_symbol,
  0,
  coin_symbol,
  0,
  0,
  10001,
  ROUND(mid_price * 0.996, 8),
  0,
  symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 20 MINUTE)) * 1000,
  0,
  0,
  1,
  '0'
FROM seed_symbols
UNION ALL
SELECT
  CONCAT('LOCAL-SEED-', REPLACE(symbol, '/', '-'), '-BUY-2'),
  2.40000000,
  base_symbol,
  0,
  coin_symbol,
  0,
  0,
  10002,
  ROUND(mid_price * 0.992, 8),
  0,
  symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 16 MINUTE)) * 1000,
  0,
  0,
  1,
  '0'
FROM seed_symbols
UNION ALL
SELECT
  CONCAT('LOCAL-SEED-', REPLACE(symbol, '/', '-'), '-SELL-1'),
  1.10000000,
  base_symbol,
  0,
  coin_symbol,
  0,
  1,
  10001,
  ROUND(mid_price * 1.004, 8),
  0,
  symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 12 MINUTE)) * 1000,
  0,
  0,
  1,
  '0'
FROM seed_symbols
UNION ALL
SELECT
  CONCAT('LOCAL-SEED-', REPLACE(symbol, '/', '-'), '-SELL-2'),
  2.20000000,
  base_symbol,
  0,
  coin_symbol,
  0,
  1,
  10002,
  ROUND(mid_price * 1.008, 8),
  0,
  symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 9 MINUTE)) * 1000,
  0,
  0,
  1,
  '0'
FROM seed_symbols
UNION ALL
SELECT
  CONCAT('LOCAL-SEED-', REPLACE(symbol, '/', '-'), '-DONE-BUY'),
  0.90000000,
  base_symbol,
  0,
  coin_symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 6 MINUTE)) * 1000,
  0,
  10001,
  ROUND(mid_price * 1.001, 8),
  1,
  symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 7 MINUTE)) * 1000,
  0.90000000,
  ROUND(ROUND(mid_price * 1.001, 8) * 0.90000000, 8),
  1,
  '0'
FROM seed_symbols
UNION ALL
SELECT
  CONCAT('LOCAL-SEED-', REPLACE(symbol, '/', '-'), '-DONE-SELL'),
  1.40000000,
  base_symbol,
  0,
  coin_symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 3 MINUTE)) * 1000,
  1,
  10002,
  ROUND(mid_price * 0.999, 8),
  1,
  symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 4 MINUTE)) * 1000,
  1.40000000,
  ROUND(ROUND(mid_price * 0.999, 8) * 1.40000000, 8),
  1,
  '0'
FROM seed_symbols
UNION ALL
SELECT
  CONCAT('LOCAL-SEED-', REPLACE(symbol, '/', '-'), '-CANCELED'),
  0.75000000,
  base_symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 2 MINUTE)) * 1000,
  coin_symbol,
  0,
  0,
  10001,
  ROUND(mid_price * 0.985, 8),
  2,
  symbol,
  UNIX_TIMESTAMP(DATE_SUB(NOW(3), INTERVAL 5 MINUTE)) * 1000,
  0,
  0,
  1,
  '0'
FROM seed_symbols;
