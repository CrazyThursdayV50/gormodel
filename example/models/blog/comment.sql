CREATE TABLE `comment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `post_id` int unsigned NOT NULL COMMENT '文章ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父评论ID，0表示顶级评论',
  `content` text NOT NULL COMMENT '评论内容',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-已发布，2-待审核，0-已删除',
  `like_count` int unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `ip` varchar(45) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `user_agent` varchar(255) NOT NULL DEFAULT '' COMMENT '用户代理',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_post` (`post_id`),
  KEY `idx_user` (`user_id`),
  KEY `idx_parent` (`parent_id`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 