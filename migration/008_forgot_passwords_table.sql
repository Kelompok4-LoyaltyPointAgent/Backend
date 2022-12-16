-- +goose Up
CREATE TABLE `forgot_passwords` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `user_id` varchar(255),
  `access_key` varchar(255),
  `expired_at` datetime(3),
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_forgot_passwords_access_key` (`access_key`),
  KEY `idx_forgot_passwords_deleted_at` (`deleted_at`),
  KEY `fk_forgot_passwords_user` (`user_id`),
  CONSTRAINT `fk_forgot_passwords_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `forgot_passwords`;