-- 创建 member_address 表（用户地址表）
CREATE TABLE IF NOT EXISTS `member_address` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `member_id` bigint(20) NOT NULL COMMENT '用户 ID',
  `coin_id` bigint(20) NOT NULL COMMENT '币种 ID',
  `address` varchar(255) NOT NULL COMMENT '钱包地址',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `status` int(11) NOT NULL DEFAULT 0 COMMENT '状态：0 正常 1 非法',
  `create_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `delete_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_coin_id` (`coin_id`),
  KEY `idx_member_coin` (`member_id`, `coin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户地址表';

-- 创建 withdraw_record 表（提现记录表）
CREATE TABLE IF NOT EXISTS `withdraw_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `member_id` bigint(20) NOT NULL COMMENT '用户 ID',
  `coin_id` bigint(20) NOT NULL COMMENT '币种 ID',
  `total_amount` decimal(20,8) NOT NULL DEFAULT 0.00000000 COMMENT '提现金额',
  `fee` decimal(20,8) NOT NULL DEFAULT 0.00000000 COMMENT '手续费',
  `arrived_amount` decimal(20,8) NOT NULL DEFAULT 0.00000000 COMMENT '到账金额',
  `address` varchar(255) NOT NULL COMMENT '提现地址',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `transaction_number` varchar(255) DEFAULT '' COMMENT '交易哈希',
  `can_auto_withdraw` int(11) NOT NULL DEFAULT 0 COMMENT '是否自动提现：0 否 1 是',
  `isAuto` int(11) NOT NULL DEFAULT 0 COMMENT '是否自动：0 否 1 是',
  `status` int(11) NOT NULL DEFAULT 0 COMMENT '状态：0 审核中 1 成功 2 失败',
  `create_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `deal_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '处理时间',
  PRIMARY KEY (`id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_coin_id` (`coin_id`),
  KEY `idx_status` (`status`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='提现记录表';
