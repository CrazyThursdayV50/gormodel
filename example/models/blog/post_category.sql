CREATE TABLE `post_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `post_id` int unsigned NOT NULL COMMENT '文章ID',
  `category_id` int unsigned NOT NULL COMMENT '分类ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_post_category` (`post_id`, `category_id`),
  KEY `idx_post` (`post_id`),
  KEY `idx_category` (`category_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 