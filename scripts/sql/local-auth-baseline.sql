CREATE TABLE IF NOT EXISTS member (
  id BIGINT NOT NULL AUTO_INCREMENT,
  username VARCHAR(64) NOT NULL DEFAULT '',
  mobile_phone VARCHAR(32) NOT NULL DEFAULT '',
  password VARCHAR(512) NOT NULL DEFAULT '',
  salt VARCHAR(128) NOT NULL DEFAULT '',
  login_count BIGINT NOT NULL DEFAULT 0,
  member_level BIGINT NOT NULL DEFAULT 0,
  real_name VARCHAR(64) NOT NULL DEFAULT '',
  country VARCHAR(64) NOT NULL DEFAULT '',
  avatar VARCHAR(255) NOT NULL DEFAULT '',
  promotion_code VARCHAR(64) NOT NULL DEFAULT '',
  super_partner VARCHAR(16) NOT NULL DEFAULT '0',
  status BIGINT NOT NULL DEFAULT 0,
  email VARCHAR(128) NOT NULL DEFAULT '',
  jy_password VARCHAR(255) NOT NULL DEFAULT '',
  id_number VARCHAR(64) NOT NULL DEFAULT '',
  real_name_status BIGINT NOT NULL DEFAULT 0,
  bank VARCHAR(128) NOT NULL DEFAULT '',
  ali_no VARCHAR(64) NOT NULL DEFAULT '',
  wechat VARCHAR(64) NOT NULL DEFAULT '',
  registration_time BIGINT NOT NULL DEFAULT 0,
  token VARCHAR(512) NOT NULL DEFAULT '',
  token_expire_time BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  UNIQUE KEY uk_member_mobile_phone (mobile_phone)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS member_wallet (
  id BIGINT NOT NULL AUTO_INCREMENT,
  address VARCHAR(255) NOT NULL DEFAULT '',
  balance DECIMAL(32,16) NOT NULL DEFAULT 0,
  frozen_balance DECIMAL(32,16) NOT NULL DEFAULT 0,
  release_balance DECIMAL(32,16) NOT NULL DEFAULT 0,
  is_lock INT NOT NULL DEFAULT 0,
  member_id BIGINT NOT NULL,
  version INT NOT NULL DEFAULT 0,
  coin_id BIGINT NOT NULL DEFAULT 0,
  to_released DECIMAL(32,16) NOT NULL DEFAULT 0,
  coin_name VARCHAR(32) NOT NULL DEFAULT '',
  address_private_key TEXT,
  PRIMARY KEY (id),
  UNIQUE KEY uk_member_wallet_member_coin (member_id, coin_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS member_transaction (
  id BIGINT NOT NULL AUTO_INCREMENT,
  address VARCHAR(255) NOT NULL DEFAULT '',
  amount DECIMAL(32,16) NOT NULL DEFAULT 0,
  create_time BIGINT NOT NULL DEFAULT 0,
  fee DECIMAL(32,16) NOT NULL DEFAULT 0,
  flag INT NOT NULL DEFAULT 0,
  member_id BIGINT NOT NULL,
  symbol VARCHAR(32) NOT NULL DEFAULT '',
  type INT NOT NULL DEFAULT 0,
  discount_fee VARCHAR(64) NOT NULL DEFAULT '',
  real_fee VARCHAR(64) NOT NULL DEFAULT '',
  PRIMARY KEY (id),
  KEY idx_member_transaction_member (member_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO member (
  id, username, mobile_phone, password, salt, login_count, member_level, real_name, country, avatar,
  promotion_code, super_partner, status, email, jy_password, id_number, real_name_status, bank, ali_no, wechat, registration_time
)
VALUES (
  10001,
  'local-tester',
  '13800000000',
  '38105f0f7f02321cac5665dcc91c8a41a7f23f09a1019bb941c7047d3a6e0c37be8733e1a9b1a183ca4829fc657869f82b3c9a468ba6fe1c8620dc60c6c4ee3db52f4e22bb92ec78d341f3ac604681288019e3783392e7cfeb1371bad79d3d673730a4ce96c5144d27eb9255e5d2c6a301234692a82bb2bc5815c6b71b736557',
  'LocalAuthSeedSaltForMscoin2026',
  0,
  0,
  'Local Tester',
  'CN',
  'https://mszlu.oss-cn-beijing.aliyuncs.com/mscoin/defaultavatar.png',
  'LOCALTEST',
  '0',
  0,
  'local-tester@example.com',
  '',
  '',
  0,
  '',
  '',
  '',
  UNIX_TIMESTAMP() * 1000
)
ON DUPLICATE KEY UPDATE
  username = VALUES(username),
  password = VALUES(password),
  salt = VALUES(salt),
  country = VALUES(country),
  avatar = VALUES(avatar),
  promotion_code = VALUES(promotion_code),
  email = VALUES(email);

INSERT INTO member_wallet (
  member_id, coin_id, coin_name, address, balance, frozen_balance, release_balance, to_released, is_lock, version
)
VALUES
  (10001, 2, 'USDT', 'LOCAL-USDT-ADDRESS-10001', 1250.50000000, 50.25000000, 0, 0, 0, 0),
  (10001, 1, 'BTC', 'LOCAL-BTC-ADDRESS-10001', 0.25000000, 0.01000000, 0, 0, 0, 0),
  (10001, 0, 'ETH', 'LOCAL-ETH-ADDRESS-10001', 25.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'SOL', 'LOCAL-SOL-ADDRESS-10001', 500.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'OKB', 'LOCAL-OKB-ADDRESS-10001', 300.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'DOGE', 'LOCAL-DOGE-ADDRESS-10001', 50000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'XRP', 'LOCAL-XRP-ADDRESS-10001', 5000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'BCH', 'LOCAL-BCH-ADDRESS-10001', 50.00000000, 0, 0, 0, 0, 0),
  (10001, 0, '1INCH', 'LOCAL-1INCH-ADDRESS-10001', 10000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'AAVE', 'LOCAL-AAVE-ADDRESS-10001', 100.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'ACH', 'LOCAL-ACH-ADDRESS-10001', 100000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'ADA', 'LOCAL-ADA-ADDRESS-10001', 20000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'AEVO', 'LOCAL-AEVO-ADDRESS-10001', 20000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'AGLD', 'LOCAL-AGLD-ADDRESS-10001', 5000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'ALGO', 'LOCAL-ALGO-ADDRESS-10001', 20000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'APE', 'LOCAL-APE-ADDRESS-10001', 20000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'API3', 'LOCAL-API3-ADDRESS-10001', 5000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'APT', 'LOCAL-APT-ADDRESS-10001', 5000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'AR', 'LOCAL-AR-ADDRESS-10001', 1000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'ARB', 'LOCAL-ARB-ADDRESS-10001', 10000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'ATOM', 'LOCAL-ATOM-ADDRESS-10001', 5000.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'BETH', 'LOCAL-BETH-ADDRESS-10001', 20.00000000, 0, 0, 0, 0, 0),
  (10001, 0, 'STETH', 'LOCAL-STETH-ADDRESS-10001', 20.00000000, 0, 0, 0, 0, 0)
ON DUPLICATE KEY UPDATE
  address = VALUES(address),
  balance = VALUES(balance),
  frozen_balance = VALUES(frozen_balance),
  release_balance = VALUES(release_balance),
  to_released = VALUES(to_released),
  is_lock = VALUES(is_lock),
  version = VALUES(version);
