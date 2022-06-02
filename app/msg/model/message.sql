CREATE TABLE `chat_msg` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `group_id` varchar(255) NOT NULL,
  `sender_id` bigint NOT NULL,
  `type` int DEFAULT 1 COMMENT "1文本, 2图片, 3视频, 4音频", 
  `content` varchar(2048) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `uuid` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uuid` (`uuid`),
  INDEX `idx_group_id` (`group_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;