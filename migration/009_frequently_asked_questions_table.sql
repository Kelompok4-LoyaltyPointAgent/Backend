-- +goose Up
CREATE TABLE `frequently_asked_questions` (
  `id` varchar(255) NOT NULL,
  `created_at` datetime(3),
  `updated_at` datetime(3),
  `deleted_at` datetime(3),
  `category` varchar(255),
  `question` text,
  `answer` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_frequently_asked_questions_category` (`category`),
  KEY `idx_frequently_asked_questions_deleted_at` (`deleted_at`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `frequently_asked_questions`;