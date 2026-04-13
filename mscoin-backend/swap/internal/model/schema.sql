-- 合约配置表
CREATE TABLE IF NOT EXISTS `contract_coins` (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `symbol` VARCHAR(20) NOT NULL COMMENT '合约符号，如 BTCUSDT',
    `contract_type` TINYINT NOT NULL DEFAULT 1 COMMENT '合约类型：1=永续',
    `price_precision` TINYINT NOT NULL DEFAULT 2 COMMENT '价格精度',
    `quantity_precision` TINYINT NOT NULL DEFAULT 4 COMMENT '数量精度',
    `share_number` DECIMAL(20,8) NOT NULL COMMENT '每张面值',
    `maker_fee` DECIMAL(10,8) NOT NULL DEFAULT '0.0002' COMMENT 'maker 手续费率',
    `taker_fee` DECIMAL(10,8) NOT NULL DEFAULT '0.0006' COMMENT 'taker 手续费率',
    `min_leverage` TINYINT NOT NULL DEFAULT 1 COMMENT '最小杠杆',
    `max_leverage` TINYINT NOT NULL DEFAULT 100 COMMENT '最大杠杆',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1=交易，2=暂停，3=下架',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_symbol` (`symbol`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合约配置表';

-- 委托订单表
CREATE TABLE IF NOT EXISTS `contract_orders` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `member_id` INT UNSIGNED NOT NULL COMMENT '用户 ID',
    `contract_coin_id` INT UNSIGNED NOT NULL COMMENT '合约 ID',
    `entrust_type` TINYINT NOT NULL COMMENT '委托类型：1=开仓，2=平仓',
    `type` TINYINT NOT NULL COMMENT '价格类型：2=限价，3=计划委托',
    `direction` TINYINT NOT NULL COMMENT '方向：1=多，2=空',
    `leverage` TINYINT NOT NULL COMMENT '杠杆倍数',
    `price` DECIMAL(20,8) NOT NULL COMMENT '委托价格',
    `amount` DECIMAL(20,8) NOT NULL COMMENT '委托张数',
    `deal_amount` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '已成交张数',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1=待成交，2=已成交，3=已撤销，4=强平',
    `deal_amount_usdt` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '已成交 USDT',
    `fee` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '手续费',
    `fee_unit` VARCHAR(10) NOT NULL COMMENT '手续费币种',
    `profit` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '已实现盈亏',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `idx_member_id` (`member_id`),
    KEY `idx_contract_coin_id` (`contract_coin_id`),
    KEY `idx_status` (`status`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合约委托订单表';

-- 持仓表
CREATE TABLE IF NOT EXISTS `contract_positions` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `member_id` INT UNSIGNED NOT NULL COMMENT '用户 ID',
    `contract_coin_id` INT UNSIGNED NOT NULL COMMENT '合约 ID',
    `direction` TINYINT NOT NULL COMMENT '方向：1=多，2=空',
    `leverage` TINYINT NOT NULL COMMENT '杠杆倍数',
    `size` DECIMAL(20,8) NOT NULL COMMENT '持仓张数',
    `entry_price` DECIMAL(20,8) NOT NULL COMMENT '开仓均价',
    `margin` DECIMAL(20,8) NOT NULL COMMENT '保证金',
    `unrealized_pnl` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '未实现盈亏',
    `liquidation_price` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '强平价格',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_member_contract_direction` (`member_id`, `contract_coin_id`, `direction`),
    KEY `idx_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合约持仓表';

-- 合约钱包表
CREATE TABLE IF NOT EXISTS `contract_wallets` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `member_id` INT UNSIGNED NOT NULL COMMENT '用户 ID',
    `unit` VARCHAR(10) NOT NULL COMMENT '币种',
    `balance` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '可用余额',
    `frozen` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '冻结金额',
    `main_balance` DECIMAL(20,8) NOT NULL DEFAULT 0 COMMENT '主账户余额',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_member_unit` (`member_id`, `unit`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合约钱包表';

-- 合约账单流水表
CREATE TABLE IF NOT EXISTS `contract_transactions` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `member_id` INT UNSIGNED NOT NULL COMMENT '用户 ID',
    `unit` VARCHAR(10) NOT NULL COMMENT '币种',
    `type` TINYINT NOT NULL COMMENT '类型：0=划入，1=划出，2=结算，3=强平，4=手续费，5=盈亏',
    `amount` DECIMAL(20,8) NOT NULL COMMENT '金额',
    `balance` DECIMAL(20,8) NOT NULL COMMENT '变更后余额',
    `related_order_id` BIGINT UNSIGNED COMMENT '关联订单 ID',
    `remark` VARCHAR(255) COMMENT '备注',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    KEY `idx_member_id` (`member_id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合约账单流水表';
