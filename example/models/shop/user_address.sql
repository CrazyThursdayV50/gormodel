CREATE TABLE `user_address` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(50) NOT NULL COMMENT '收货人姓名',
  `phone` varchar(20) NOT NULL COMMENT '收货人电话',
  `province` varchar(50) NOT NULL COMMENT '省',
  `city` varchar(50) NOT NULL COMMENT '市',
  `district` varchar(50) NOT NULL COMMENT '区/县',
  `detail` varchar(200) NOT NULL COMMENT '详细地址',
  `postal_code` varchar(20) NOT NULL DEFAULT '' COMMENT '邮政编码',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认地址',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`),
  KEY `idx_is_default` (`is_default`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 