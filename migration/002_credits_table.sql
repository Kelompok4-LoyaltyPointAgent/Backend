-- +goose Up
CREATE TABLE `credits` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `description` text,
  `active_period` int,
  `amount` int,
  PRIMARY KEY (`id`),
  KEY `idx_credits_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `credits`;