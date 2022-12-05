-- +goose Up
CREATE TABLE `transactions` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `user_id` varchar(255),
  `product_id` varchar(255) NULL,
  `amount` double,
  `type` varchar(255),
  `status` varchar(255),
  `method` varchar(255),
  PRIMARY KEY (`id`),
  KEY `idx_transactions_deleted_at` (`deleted_at`),
  KEY `fk_transactions_user` (`user_id`),
  CONSTRAINT `fk_transactions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  KEY `fk_transactions_product` (`product_id`),
  CONSTRAINT `fk_transactions_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `transaction_details` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `transaction_id` varchar(255),
  `number` varchar(255),
  `email` varchar(255),
  PRIMARY KEY (`id`),
  KEY `idx_transaction_details_deleted_at` (`deleted_at`),
  KEY `fk_transaction_details_transaction` (`transaction_id`),
  CONSTRAINT `fk_transaction_details_transaction` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- +goose Down
DROP TABLE IF EXISTS `transaction_details`;
DROP TABLE IF EXISTS `transactions`;