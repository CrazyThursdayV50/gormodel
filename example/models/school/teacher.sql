CREATE TABLE `teacher` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `sex` boolean NOT NULL DEFAULT '0' COMMENT 'male: 0; femail: 1',
  `age` int unsigned NOT NULL DEFAULT '0' COMMENT 'age',
  `name` varchar(100) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_age` (`age`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
