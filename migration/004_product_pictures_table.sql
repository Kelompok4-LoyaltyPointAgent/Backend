-- +goose Up
CREATE TABLE `product_pictures` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `url` text,
  `name` varchar(255),
  PRIMARY KEY (`id`),
  KEY `idx_product_pictures_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `product_pictures`;