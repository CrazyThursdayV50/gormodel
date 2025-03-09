CREATE TABLE `course` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '课程名称',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '课程代码',
  `credit` decimal(3,1) NOT NULL DEFAULT '0.0' COMMENT '学分',
  `hours` int unsigned NOT NULL DEFAULT '0' COMMENT '课时',
  `description` text COMMENT '课程描述',
  `teacher_id` int unsigned NOT NULL DEFAULT '0' COMMENT '教师ID',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-正常，0-停用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`),
  KEY `idx_teacher` (`teacher_id`),
  KEY `idx_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 