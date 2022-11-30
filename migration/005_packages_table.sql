-- +goose Up
CREATE TABLE `packages` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `product_id` varchar(255),
  `active_period` int,
  `total_internet` double,
  `main_internet` double,
  `night_internet` double,
  `social_media` double,
  `call` int,
  `sms` int,
  PRIMARY KEY (`id`),
  KEY `idx_packages_deleted_at` (`deleted_at`),
  KEY `fk_packages_product` (`product_id`),
  CONSTRAINT `fk_packages_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `packages`;