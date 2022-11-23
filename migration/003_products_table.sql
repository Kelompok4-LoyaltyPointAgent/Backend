-- +goose Up
CREATE TABLE `products` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `name` varchar(255),
  `type` varchar(255),
  `provider` varchar(255),
  `price` int unsigned,
  `price_points` int unsigned,
  `reward_points` int unsigned,
  `stock` int unsigned,
  `recommended` boolean,
  `product_picture_id` varchar(255),
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`),
  KEY `fk_products_product_picture` (`product_picture_id`),
  CONSTRAINT `fk_products_product_picture` FOREIGN KEY (`product_picture_id`) REFERENCES `product_pictures` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `products`;