-- +goose Up
CREATE TABLE `feedbacks` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `user_id` varchar(255),
  `is_information_helpful` boolean,
  `is_article_helpful` boolean,
  `is_article_easy_to_find` boolean,
  `is_design_good` boolean,
  `review` text,
  PRIMARY KEY (`id`),
  KEY `fk_feedbacks_user` (`user_id`),
  CONSTRAINT `fk_feedbacks_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  KEY `idx_feedbacks_deleted_at` (`deleted_at`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `feedbacks`;
