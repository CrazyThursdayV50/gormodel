CREATE TABLE `category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '分类名称',
  `slug` varchar(50) NOT NULL COMMENT 'URL友好的名称',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父分类ID，0表示顶级分类',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-正常，0-禁用',
  `sort_order` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_slug` (`slug`),
  KEY `idx_parent` (`parent_id`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 