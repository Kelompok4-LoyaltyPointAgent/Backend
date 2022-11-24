-- +goose Up
CREATE TABLE `credits` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `product_id` varchar(255),
  `description` text,
  `active_period` int,
  `amount` int,
  PRIMARY KEY (`id`),
  KEY `idx_credits_deleted_at` (`deleted_at`),
  KEY `fk_credits_product` (`product_id`),
  CONSTRAINT `fk_credits_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `credits`;