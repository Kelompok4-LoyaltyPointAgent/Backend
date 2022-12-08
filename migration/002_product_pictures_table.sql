-- +goose Up
CREATE TABLE `product_pictures` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `icon_id` varchar(255),
  `url` text,
  `name` varchar(255),
  `type` varchar(255),
  PRIMARY KEY (`id`),
  KEY `idx_product_pictures_deleted_at` (`deleted_at`),
  KEY `idx_product_pictures_icon` (`icon_id`),
  CONSTRAINT `idx_product_pictures_icon` FOREIGN KEY (`icon_id`) REFERENCES `product_pictures` (`icon_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `product_pictures`;