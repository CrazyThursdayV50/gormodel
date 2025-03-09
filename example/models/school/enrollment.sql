CREATE TABLE `enrollment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `student_id` int unsigned NOT NULL COMMENT '学生ID',
  `course_id` int unsigned NOT NULL COMMENT '课程ID',
  `semester` varchar(20) NOT NULL DEFAULT '' COMMENT '学期，如：2023-2024-1',
  `score` decimal(5,2) DEFAULT NULL COMMENT '成绩',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：1-正常，2-退课，3-重修',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_student_course_semester` (`student_id`, `course_id`, `semester`),
  KEY `idx_student` (`student_id`),
  KEY `idx_course` (`course_id`),
  KEY `idx_semester` (`semester`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci; 