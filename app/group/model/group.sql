CREATE TABLE `group` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` smallint NOT NULL comment "1表示单聊, 2表示群聊",
  `status` smallint default 0 comment "1表示有效, 2表示无效(未同意), 3表示无效(拉黑)",
  `config` json comment "群聊配置", 
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_type` (`type`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `group_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `group_id` varchar(255) NOT NULL,
  `user_id` bigint NOT NULL,
  `alias_name` varchar(255) comment "用户对该群的备注名",
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_group_id` (`group_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;