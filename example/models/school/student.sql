CREATE TABLE `student` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `age` int unsigned NOT NULL DEFAULT '0',
  `grade` int unsigned NOT NULL DEFAULT '1' COMMENT '年级',
  `class` int unsigned NOT NULL DEFAULT '1' COMMENT '班级',
  `sex` boolean NOT NULL DEFAULT '0' COMMENT 'male: 0; female: 1',
  `address` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(20) NOT NULL DEFAULT '',
  `email` varchar(100) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_grade_class` (`grade`, `class`),
  KEY `idx_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 