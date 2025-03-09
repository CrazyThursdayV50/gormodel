CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `email` varchar(100) NOT NULL COMMENT '邮箱',
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像URL',
  `bio` text COMMENT '个人简介',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-正常，0-禁用',
  `role` tinyint NOT NULL DEFAULT '1' COMMENT '角色：1-普通用户，2-管理员',
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  UNIQUE KEY `uk_email` (`email`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 