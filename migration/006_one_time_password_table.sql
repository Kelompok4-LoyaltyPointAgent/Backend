-- +goose Up
CREATE TABLE `one_time_passwords` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `user_id` varchar(255),
  `pin` varchar(255),
  `expired_at` datetime(3),
  PRIMARY KEY (`id`),
  KEY `idx_one_time_passwords_deleted_at` (`deleted_at`),
  KEY `fk_one_time_passwords_user` (`user_id`),
  CONSTRAINT `fk_one_time_passwords_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `one_time_passwords`;