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
  '8f8fca248e1bafb5d6aaaec9bee6043deeae29ff6870118e9c8ffbd938cbd34a32f81e67e86d167075528a80c0b1c928d84b58c41d82a25b8fb0d65fc5713a245efcc97fbb878ce493a158aff62d131e8e9086b5bb4a5407db9ce341464e05b1ff17adb93ca750550f06efe38feb43086e8953a185b4b2ba140796e35a825ceb',
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
  (10001, 1, 'BTC', 'LOCAL-BTC-ADDRESS-10001', 0.25000000, 0.01000000, 0, 0, 0, 0)
ON DUPLICATE KEY UPDATE
  address = VALUES(address),
  balance = VALUES(balance),
  frozen_balance = VALUES(frozen_balance),
  release_balance = VALUES(release_balance),
  to_released = VALUES(to_released),
  is_lock = VALUES(is_lock),
  version = VALUES(version);
